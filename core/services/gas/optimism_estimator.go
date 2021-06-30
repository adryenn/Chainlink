package gas

import (
	"context"
	"math/big"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

var _ Estimator = &optimismEstimator{}

// TODO: Separate tracker/estimator? can it be done?
type optimismEstimator struct {
	config Config
}

func NewOptimismEstimator(config Config) Estimator {
	return &optimismEstimator{config}
}

func (o *optimismEstimator) Start() error {
	panic("not implemented")
}
func (o *optimismEstimator) Close() error {
	panic("not implemented")
}

func (o *optimismEstimator) EstimateGas(_ []byte, gasLimit uint64) (gasPrice *big.Int, chainSpecificGasLimit uint64, err error) {
	panic("not implemented")
}

func (o *optimismEstimator) BumpGas(originalGasPrice *big.Int, originalGasLimit uint64) (gasPrice *big.Int, gasLimit uint64, err error) {
	panic("not implemented")
}

func (o *optimismEstimator) OnNewLongestChain(_ context.Context, _ models.Head) {
	panic("not implemented") // TODO: Implement
}
