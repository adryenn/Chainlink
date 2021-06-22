package vrf

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/vrf_coordinator_v2"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	httypes "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
	"go.uber.org/multierr"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type pendingRequest struct {
	confirmedAtBlock uint64
	req              *vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested
	lb               log.Broadcast
}

type listenerV2 struct {
	cfg             Config
	l               logger.Logger
	abi             abi.ABI
	ethClient       eth.Client
	logBroadcaster  log.Broadcaster
	txm             bulletprooftxmanager.TxManager
	headBroadcaster httypes.HeadBroadcasterRegistry
	coordinator     *vrf_coordinator_v2.VRFCoordinatorV2
	pipelineRunner  pipeline.Runner
	pipelineORM     pipeline.ORM
	vorm            keystore.VRFORM
	job             job.Job
	db              *gorm.DB
	vrfks           *keystore.VRF
	gethks          *keystore.Eth
	mbLogs          *utils.Mailbox
	chStop          chan struct{}
	waitOnStop      chan struct{}
	latestHead      uint64
	pendingLogs     []pendingRequest
	utils.StartStopOnce
}

func (lsn *listenerV2) Start() error {
	return lsn.StartOnce("VRFListenerV2", func() error {
		// Take the larger of the global vs specific.
		// Note that the v2 vrf requests specify their own confirmation requirements.
		// We wait for max(minConfs, request required confs) to be safe.
		minConfs := lsn.cfg.MinIncomingConfirmations()
		if lsn.job.VRFSpec.Confirmations > lsn.cfg.MinIncomingConfirmations() {
			minConfs = lsn.job.VRFSpec.Confirmations
		}
		unsubscribeLogs := lsn.logBroadcaster.Register(lsn, log.ListenerOpts{
			Contract: lsn.coordinator.Address(),
			LogsWithTopics: map[common.Hash][][]log.Topic{
				vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested{}.Topic(): {
					{
						log.Topic(lsn.job.VRFSpec.PublicKey.MustHash()),
					},
				},
			},
			// Do not specify min confirmations, as it varies from request to request.
		})

		// Subscribe to the head broadcaster for handling
		// per request conf requirements.
		unsubscribeHeadBroadcaster := lsn.headBroadcaster.Subscribe(lsn)

		// Goroutine 1 receives logs from the log broadcaster and stores them in
		//
		go lsn.run([]func(){unsubscribeLogs, unsubscribeHeadBroadcaster}, minConfs)
		return nil
	})
}

func (lsn *listenerV2) Connect(head *models.Head) error {
	lsn.latestHead = uint64(head.Number)
	return nil
}

func (lsn *listenerV2) OnNewLongestChain(ctx context.Context, head models.Head) {
	// Check if any v2 logs are ready for processing.
	lsn.l.Debugw("VRFListenerV2: new longest chain")
	lsn.latestHead = uint64(head.Number)
	var remainingLogs []pendingRequest
	for _, pl := range lsn.pendingLogs {
		if pl.confirmedAtBlock <= lsn.latestHead {
			// Note below makes API calls and opens a database transaction
			lsn.ProcessV2VRFRequest(pl.req, pl.lb)
		} else {
			remainingLogs = append(remainingLogs, pl)
		}
	}
	lsn.pendingLogs = remainingLogs
}

func (lsn *listenerV2) run(unsubscribeLogs []func(), minConfs uint32) {
	lsn.l.Infow("VRFListenerV2: listening for run requests",
		"minConfs", minConfs)
	for {
		select {
		case <-lsn.chStop:
			for _, us := range unsubscribeLogs {
				us()
			}
			lsn.waitOnStop <- struct{}{}
			return
		case <-lsn.mbLogs.Notify():
			// Process all the logs in the queue if one is added
			for {
				i, exists := lsn.mbLogs.Retrieve()
				if !exists {
					break
				}
				lb, ok := i.(log.Broadcast)
				if !ok {
					panic(fmt.Sprintf("VRFListenerV2: invariant violated, expected log.Broadcast got %T", i))
				}
				alreadyConsumed, err := lsn.logBroadcaster.WasAlreadyConsumed(lsn.db, lb)
				if err != nil {
					lsn.l.Errorw("VRFListenerV2: could not determine if log was already consumed", "error", err, "txHash", lb.RawLog().TxHash)
					continue
				} else if alreadyConsumed {
					continue
				}
				req, err := lsn.coordinator.ParseRandomWordsRequested(lb.RawLog())
				if err != nil {
					lsn.l.Errorw("VRFListenerV2: failed to parse log", "err", err, "txHash", lb.RawLog().TxHash)
					lsn.l.ErrorIf(lsn.logBroadcaster.MarkConsumed(lsn.db, lb), "failed to mark consumed")
					return
				}
				lsn.pendingLogs = append(lsn.pendingLogs, pendingRequest{
					confirmedAtBlock: req.Raw.BlockNumber + req.MinimumRequestConfirmations,
					req:              req,
					lb:               lb,
				})
			}
		}
	}
}

func (lsn *listenerV2) ProcessV2VRFRequest(req *vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested, lb log.Broadcast) {
	// Check if the vrf req has already been fulfilled
	callback, err := lsn.coordinator.SCallbacks(nil, req.PreSeed)
	if err != nil {
		lsn.l.Errorw("VRFListenerV2: unable to check if already fulfilled, processing anyways", "err", err, "txHash", req.Raw.TxHash)
	} else if utils.IsEmpty(callback[:]) {
		// If seedAndBlockNumber is zero then the response has been fulfilled
		// and we should skip it
		lsn.l.Infow("VRFListenerV2: request already fulfilled", "txHash", req.Raw.TxHash, "subID", req.SubId, "callback", callback)
		lsn.l.ErrorIf(lsn.logBroadcaster.MarkConsumed(lsn.db, lb), "failed to mark consumed")
		return
	}

	s := time.Now()
	proof, err1 := lsn.LogToProof(req, lb)
	gasLimit, err2 := lsn.computeTxGasLimit(req.CallbackGasLimit, proof)
	vrfCoordinatorPayload, _, err3 := lsn.ProcessLogV2(proof)
	err = multierr.Combine(err1, err2, err3)
	logger.Infow("estimated gas limit for tx", "gasLimit", gasLimit, "callbackLimit", req.CallbackGasLimit)
	f := time.Now()
	err = postgres.GormTransactionWithDefaultContext(lsn.db, func(tx *gorm.DB) error {
		if err == nil {
			// No errors processing the log, submit a transaction
			var etx models.EthTx
			var from common.Address
			from, err = lsn.gethks.GetRoundRobinAddress()
			if err != nil {
				return err
			}
			etx, err = lsn.txm.CreateEthTransaction(tx,
				from,
				lsn.coordinator.Address(),
				vrfCoordinatorPayload,
				gasLimit,
				&models.EthTxMetaV2{
					JobID: lsn.job.ID,
					//RequestID:     req.PreSeed,
					RequestTxHash: lb.RawLog().TxHash,
				})
			if err != nil {
				return err
			}
			// TODO: Once we have eth tasks supported, we can use the pipeline directly
			// and be able to save errored proof generations. Until then only save
			// successful runs and log errors.
			_, err = lsn.pipelineRunner.InsertFinishedRun(tx, pipeline.Run{
				PipelineSpecID: lsn.job.PipelineSpecID,
				Errors:         []null.String{{}},
				Outputs: pipeline.JSONSerializable{
					Val: []interface{}{fmt.Sprintf("queued tx from %v to %v txdata %v",
						etx.FromAddress,
						etx.ToAddress,
						hex.EncodeToString(etx.EncodedPayload))},
				},
				Meta: pipeline.JSONSerializable{
					Val: map[string]interface{}{"eth_tx_id": etx.ID},
				},
				CreatedAt:  s,
				FinishedAt: &f,
			}, nil, false)
			if err != nil {
				return errors.Wrap(err, "VRFListenerV2: failed to insert finished run")
			}
		}
		// Always mark consumed regardless of whether the proof failed or not.
		err = lsn.logBroadcaster.MarkConsumed(tx, lb)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		lsn.l.Errorw("VRFListenerV2 failed to save run", "err", err)
	}
}

// Compute the gasLimit required for the fulfillment transaction
// such that the user gets their requested amount of gas.
func (lsn *listenerV2) computeTxGasLimit(requestedCallbackGas uint64, proof []byte) (uint64, error) {
	vrfCoordinatorArgs, err := lsn.abi.Methods["getRandomnessFromProof"].Inputs.PackValues(
		[]interface{}{
			proof[:], // geth expects slice, even if arg is constant-length
		})
	if err != nil {
		lsn.l.Errorw("VRFListenerV2: error building fulfill args", "err", err)
		return 0, err
	}
	to := lsn.coordinator.Address()
	variableVerifyGas, err := lsn.ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &to,
		Data: append(lsn.abi.Methods["getRandomnessFromProof"].ID, vrfCoordinatorArgs...),
	})
	if err != nil {
		return 0, err
	}
	// Gas for everything other than "getRandomnessFromProof" calls and oracle payments
	// We will have a hard upper bound on that - worse case is first request etc.
	// TODO: seems to be variation with the sim and also seems to cost a more than 6k as suggested in v1?
	staticVerifyGas := uint64(22000)
	return variableVerifyGas + requestedCallbackGas + staticVerifyGas, nil
}

func (lsn *listenerV2) LogToProof(req *vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested, lb log.Broadcast) ([]byte, error) {
	lsn.l.Infow("VRFListenerV2: received log request",
		"log", lb.String(),
		"reqID", req.PreSeed.String(),
		"keyHash", hex.EncodeToString(req.KeyHash[:]),
		"txHash", req.Raw.TxHash,
		"blockNumber", req.Raw.BlockNumber,
		"seed", req.PreSeed.String())
	// Validate the key against the spec
	kh, err := lsn.job.VRFSpec.PublicKey.Hash()
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(req.KeyHash[:], kh[:]) {
		return nil, errors.New(fmt.Sprintf("invalid key hash %v expected %v", hex.EncodeToString(req.KeyHash[:]), hex.EncodeToString(kh[:])))
	}

	// req.PreSeed is uint256(keccak256(abi.encode(keyHash, msg.sender, nonce)))
	preSeed, err := BigToSeed(req.PreSeed)
	if err != nil {
		return nil, errors.New("unable to parse preseed")
	}
	seed := PreSeedData{
		PreSeed:   preSeed,
		BlockHash: req.Raw.BlockHash,
		BlockNum:  req.Raw.BlockNumber,
		// V2 only fields
		SubId:            req.SubId,
		CallbackGasLimit: req.CallbackGasLimit,
		NumWords:         req.NumWords,
		Sender:           req.Sender,
	}
	lsn.l.Infow("generating proof", "pk", lsn.job.VRFSpec.PublicKey.String(), "seed", preSeed, "blockHash", req.Raw.BlockHash.String(), "sender", req.Sender)
	solidityProof, err := GenerateProofResponse(lsn.vrfks, lsn.job.VRFSpec.PublicKey, seed)
	if err != nil {
		lsn.l.Errorw("VRFListenerV2: error generating proof", "err", err)
		return nil, err
	}
	return solidityProof[:], nil
}

func (lsn *listenerV2) ProcessLogV2(solidityProof []byte) ([]byte, *vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested, error) {
	vrfCoordinatorArgs, err := lsn.abi.Methods["fulfillRandomWords"].Inputs.PackValues(
		[]interface{}{
			solidityProof[:], // geth expects slice, even if arg is constant-length
		})
	if err != nil {
		lsn.l.Errorw("VRFListenerV2: error building fulfill args", "err", err)
		return nil, nil, err
	}

	return append(lsn.abi.Methods["fulfillRandomWords"].ID, vrfCoordinatorArgs...), nil, nil
}

// Close complies with job.Service
func (lsn *listenerV2) Close() error {
	return lsn.StopOnce("VRFListenerV2", func() error {
		close(lsn.chStop)
		<-lsn.waitOnStop
		return nil
	})
}

func (lsn *listenerV2) HandleLog(lb log.Broadcast) {
	wasOverCapacity := lsn.mbLogs.Deliver(lb)
	if wasOverCapacity {
		logger.Error("VRFListenerV2: log mailbox is over capacity - dropped the oldest log")
	}
}

// JobID complies with log.Listener
func (*listenerV2) JobID() models.JobID {
	return models.NilJobID
}

// Job complies with log.Listener
func (lsn *listenerV2) JobIDV2() int32 {
	return lsn.job.ID
}

// IsV2Job complies with log.Listener
func (*listenerV2) IsV2Job() bool {
	return true
}
