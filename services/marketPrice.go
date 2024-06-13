package services

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/common"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/domain/marketPriceDomain"
)

type Token struct {
	Symbol  string
	Chain   string
	Network string
	Price   float64
}

type MarketPriceSvc interface {
	Get() (*Token, error)
	ChangePrice(price float64) (*Token, error)
}

type marketPriceSvc struct {
	domain marketPriceDomain.Domain
}

// NewMarketPriceSvc is an over-simplified pricing service that only has one token to help with market price
func NewMarketPriceSvc(token *Token) (MarketPriceSvc, error) {
	// much of the validations happening in this func are done since this is a mock implementation

	if token == nil {
		return nil, fmt.Errorf("nil token: token cannot be nil")
	}

	if err := common.CheckString(strings.ToLower(token.Symbol)); err != nil {
		return nil, fmt.Errorf("invalid symbol: %w", err)
	}
	if err := common.CheckString(strings.ToLower(token.Chain)); err != nil {
		return nil, fmt.Errorf("invalid chain: %w", err)
	}
	if err := common.CheckString(strings.ToLower(token.Network)); err != nil {
		return nil, fmt.Errorf("invalid network: %w", err)
	}

	if big.NewFloat(token.Price).Cmp(big.NewFloat(0)) < 0 {
		return nil, fmt.Errorf("invalid price: cannot be less than 0")
	}

	return &marketPriceSvc{domain: marketPriceDomain.InitDomain(
		&marketPriceDomain.Token{
			Symbol:  token.Symbol,
			Chain:   token.Chain,
			Network: token.Network,
			Price:   token.Price,
		},
	)}, nil
}

func (m *marketPriceSvc) Get() (*Token, error) {
	tokenDO := m.domain.Get()

	return &Token{
		Symbol:  tokenDO.Symbol,
		Chain:   tokenDO.Chain,
		Network: tokenDO.Network,
		Price:   tokenDO.Price,
	}, nil
}

func (m *marketPriceSvc) ChangePrice(price float64) (*Token, error) {
	if big.NewFloat(price).Cmp(big.NewFloat(0)) < 0 {
		return nil, errors.New("price cannot be below 0")
	}

	tokenDO := m.domain.ChangePrice(price)

	return &Token{
		Symbol:  tokenDO.Symbol,
		Chain:   tokenDO.Chain,
		Network: tokenDO.Network,
		Price:   tokenDO.Price,
	}, nil
}
