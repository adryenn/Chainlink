package gas

import (
	"context"
	"math/big"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

var _ Estimator = &fixedPriceEstimator{}

// TODO: Separate tracker/estimator? can it be done?
type fixedPriceEstimator struct {
	config Config
}

func NewFixedPriceEstimator(config Config) Estimator {
	return &fixedPriceEstimator{config}
}

func (f *fixedPriceEstimator) Start() error {
	return nil
}
func (f *fixedPriceEstimator) Close() error {
	return nil
}

func (f *fixedPriceEstimator) OnNewLongestChain(_ context.Context, _ models.Head) {
	panic("not implemented") // TODO: Implement
}

func (f *fixedPriceEstimator) EstimateGas(_ []byte, gasLimit uint64) (gasPrice *big.Int, chainSpecificGasLimit uint64, err error) {
	gasPrice = f.config.EthGasPriceDefault()
	chainSpecificGasLimit = applyMultiplier(gasLimit, f.config.EthGasLimitMultiplier())
	return
}

// , "ethGasPriceDefault", store.Config.EthGasPriceDefault())
func (f *fixedPriceEstimator) BumpGas(originalGasPrice *big.Int, originalGasLimit uint64) (gasPrice *big.Int, gasLimit uint64, err error) {
	return bumpGasPriceOnly(f.config, originalGasPrice, originalGasLimit)
}
