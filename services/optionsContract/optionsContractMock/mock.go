package optionsContractMock

import (
	"errors"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/optionsContract"
)

type optionsContractMock struct {
	analysis            func(opts []*optionsContract.Options) (*optionsContract.OptionsAnalysis, error)
	analysisCalculation func(opts []*optionsContract.Options, price float64) (float64, float64, []float64, []optionsContract.XYValue)
}

func (o *optionsContractMock) Analysis(opts []*optionsContract.Options) (*optionsContract.OptionsAnalysis, error) {
	if o.analysis != nil {
		return o.analysis(opts)
	}
	return nil, errors.New("not implemented")
}

func (o *optionsContractMock) AnalysisCalculation(opts []*optionsContract.Options, price float64) (float64, float64, []float64, []optionsContract.XYValue) {
	if o.analysis != nil {
		return o.analysisCalculation(opts, price)
	}
	return 0, 0, nil, nil
}

// Option is the option passed to the mock service
type Option interface {
	apply(*optionsContractMock)
}

type optionFunc func(mock *optionsContractMock)

func (f optionFunc) apply(o *optionsContractMock) { f(o) }

func New(opts ...Option) optionsContract.OptionsSvc {
	mock := new(optionsContractMock)
	for _, o := range opts {
		o.apply(mock)
	}
	return mock
}

func WithAnalysis(f func(opts []*optionsContract.Options) (*optionsContract.OptionsAnalysis, error)) Option {
	return optionFunc(func(mock *optionsContractMock) {
		mock.analysis = f
	})
}

func WithAnalysisCalculation(f func(opts []*optionsContract.Options, price float64) (float64, float64, []float64, []optionsContract.XYValue)) Option {
	return optionFunc(func(mock *optionsContractMock) {
		mock.analysisCalculation = f
	})
}
