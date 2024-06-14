package controllers

import (
	"fmt"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services"
)

type Controllers struct {
	AnalysisController
}

func New() (*Controllers, error) {
	newServices, err := services.Init()
	if err != nil {
		return nil, fmt.Errorf("unable to bootstrap services: %w", err)
	}

	newAnalysisController := NewAnalysisController(newServices.OptionsContractService)

	return &Controllers{newAnalysisController}, nil
}
