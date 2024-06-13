package services

import "github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/model"

type OptionsSvc interface {
}

type optionSvc struct {
	MarketSvc MarketPriceSvc
}

type OptionsAnalysis struct {
	XYValues        []*model.XYValue
	MaxProfit       float64
	MaxLoss         float64
	BreakEvenPoints []float64
}

// NewOptionsSvc initializes a new options service
func NewOptionsSvc(mktSvc MarketPriceSvc) OptionsSvc {
	return &optionSvc{MarketSvc: mktSvc}
}

func (o *optionSvc) Analysis(opts []*model.Options) *OptionsAnalysis {
	analysis := new(OptionsAnalysis)

	analysis.MaxProfit = o.CalculateMaxProfit(opts)
	analysis.MaxLoss = o.CalculateMaxLoss(opts)
	analysis.XYValues = o.CalculateXYValues(opts)
	analysis.BreakEvenPoints = o.CalculateBreakEvenPoints(opts)

	return analysis
}

// CalculateXYValues calculates the x and y values for risk and reward graph
func (o *optionSvc) CalculateXYValues(contracts []*model.Options) []*model.XYValue {

	return nil
}

func (o *optionSvc) CalculateMaxProfit(contracts []*model.Options) float64 {
	// Your code here
	return 0
}

func (o *optionSvc) CalculateMaxLoss(contracts []*model.Options) float64 {
	// Your code here
	return 0
}

func (o *optionSvc) CalculateBreakEvenPoints(contracts []*model.Options) []float64 {
	// Your code here
	return nil
}
