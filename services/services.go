package services

import (
	"fmt"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/marketPrice"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/optionsContract"
)

type New struct {
	MktPriceService        marketPrice.Service
	OptionsContractService optionsContract.Service
}

func Init() (*New, error) {
	mktPricingServices, err := marketPrice.NewMarketPriceSvc(&marketPrice.Token{
		Symbol:  "Test",
		Chain:   "TestChain",
		Network: "TestNetwork",
		Price:   100,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to bootstrap market pricing services: %w", err)
	}

	optsServices := optionsContract.NewOptionsSvc(mktPricingServices)

	return &New{
		MktPriceService:        mktPricingServices,
		OptionsContractService: optsServices,
	}, nil
}
