package marketPriceMock

import (
	"errors"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/marketPrice"
)

type marketPriceMock struct {
	get         func() (*marketPrice.Token, error)
	changePrice func(price float64) (*marketPrice.Token, error)
}

func (m *marketPriceMock) Get() (*marketPrice.Token, error) {
	if m.get != nil {
		return m.get()
	}

	return nil, errors.New("not implemented")
}

func (m *marketPriceMock) ChangePrice(price float64) (*marketPrice.Token, error) {
	if m.changePrice != nil {
		return m.changePrice(price)
	}

	return nil, errors.New("not implemented")
}

// Option is the option passed to the mock service
type Option interface {
	apply(*marketPriceMock)
}

type optionFunc func(mock *marketPriceMock)

func (f optionFunc) apply(o *marketPriceMock) { f(o) }

func New(opts ...Option) marketPrice.MarketPriceSvc {
	mock := new(marketPriceMock)
	for _, o := range opts {
		o.apply(mock)
	}
	return mock
}

func WithGet(f func() (*marketPrice.Token, error)) Option {
	return optionFunc(func(mock *marketPriceMock) {
		mock.get = f
	})
}

func WithChangePrice(f func(price float64) (*marketPrice.Token, error)) Option {
	return optionFunc(func(mock *marketPriceMock) {
		mock.changePrice = f
	})
}
