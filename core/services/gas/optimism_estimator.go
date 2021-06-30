package gas

import (
	"context"
	"math/big"
	"sync"
	"time"

	optimismfees "github.com/ethereum-optimism/go-optimistic-ethereum-utils/fees"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/tidwall/gjson"
	"go.uber.org/multierr"
)

// It's always 0.015 GWei
// See: https://www.notion.so/How-to-pay-Fees-in-Optimistic-Ethereum-f706f4e5b13e460fa5671af48ce9a695
const optimisml1GasPrice = 15000000000

var _ Estimator = &optimismEstimator{}

type optimismRPCClient interface {
	Call(result interface{}, method string, args ...interface{}) error
}

// TODO: Separate tracker/estimator? can it be done?
type optimismEstimator struct {
	config     Config
	client     optimismRPCClient
	pollPeriod time.Duration

	gasPriceMu sync.RWMutex
	l1GasPrice *big.Int
	l2GasPrice *big.Int

	chInitialised chan struct{}
	chStop        chan struct{}
	chDone        chan struct{}
}

// NewOptimismEstimator returns a new optimism estimator
func NewOptimismEstimator(config Config, client optimismRPCClient) Estimator {
	return &optimismEstimator{
		config,
		client,
		10 * time.Second,
		sync.RWMutex{},
		nil,
		nil,
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
	}
}

func (o *optimismEstimator) Start() error {
	go o.run()
	<-o.chInitialised
	return nil
}
func (o *optimismEstimator) Close() error {
	close(o.chStop)
	<-o.chDone
	return nil
}

func (o *optimismEstimator) run() {
	defer close(o.chDone)

	t := o.refreshPrices()
	close(o.chInitialised)

	for {
		select {
		case <-o.chStop:
			return
		case <-t.C:
			t = o.refreshPrices()
		}
	}
}

type gasPriceResponse struct {
	l1GasPrice *big.Int
	l2GasPrice *big.Int
}

func (g *gasPriceResponse) UnmarshalJSON(b []byte) error {
	var l1Hex string = gjson.GetBytes(b, "l1GasPrice").Str
	var l2Hex string = gjson.GetBytes(b, "l2GasPrice").Str
	var l1 *hexutil.Big
	var l2 *hexutil.Big
	if err := multierr.Combine(l1.UnmarshalText([]byte(l1Hex)), l2.UnmarshalText([]byte(l2Hex))); err != nil {
		return err
	}
	g.l1GasPrice = l1.ToInt()
	g.l2GasPrice = l2.ToInt()
	return nil
}

func (o *optimismEstimator) refreshPrices() *time.Timer {
	var res gasPriceResponse
	o.client.Call(&res, "rollup_gasPrices")

	o.gasPriceMu.Lock()
	defer o.gasPriceMu.Unlock()
	o.l1GasPrice, o.l2GasPrice = res.l1GasPrice, res.l2GasPrice
	return time.NewTimer(utils.WithJitter(o.pollPeriod))
}

func (o *optimismEstimator) EstimateGas(calldata []byte, gasLimit uint64) (gasPrice *big.Int, chainSpecificGasLimit uint64, err error) {
	return o.calcGas(calldata, gasLimit)
}

func (o *optimismEstimator) BumpGas(originalGasPrice *big.Int, originalGasLimit uint64) (gasPrice *big.Int, gasLimit uint64, err error) {
	return nil, 0, errors.New("bump gas is not supported for optimism")
}

func (o *optimismEstimator) OnNewLongestChain(_ context.Context, _ models.Head) { return }

func (o *optimismEstimator) calcGas(calldata []byte, l2GasLimit uint64) (chainSpecificGasPrice *big.Int, chainSpecificGasLimit uint64, err error) {
	l1GasPrice, l2GasPrice := o.getGasPrices()

	optimismGasLimitBig := optimismfees.EncodeTxGasLimit(calldata, l1GasPrice, big.NewInt(int64(l2GasLimit)), l2GasPrice)
	if !optimismGasLimitBig.IsInt64() {
		logger.Errorw(
			"Optimism: unable to represent gas limit as Int64, this is an unexpected error and should be reported to the Chainlink team",
			"calldata", calldata,
			"l2GasLimit", l2GasLimit,
			"l1GasPrice", l1GasPrice,
			"l2GasPrice", l2GasPrice,
			"optimismGasLimitBig", optimismGasLimitBig.String(),
		)
		return nil, 0, errors.New("gas limit overflows int64")
	}
	chainSpecificGasLimit = uint64(optimismGasLimitBig.Int64())

	return big.NewInt(optimisml1GasPrice), chainSpecificGasLimit, nil
}

func (o *optimismEstimator) getGasPrices() (l1GasPrice, l2GasPrice *big.Int) {
	o.gasPriceMu.RLock()
	defer o.gasPriceMu.RUnlock()
	return o.l1GasPrice, o.l2GasPrice
}
