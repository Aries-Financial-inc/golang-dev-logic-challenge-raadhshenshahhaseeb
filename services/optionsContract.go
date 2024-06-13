package services

import "github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/model"

type Options interface {
}

type options struct {
}

// NewOptions initializes a new options service
func NewOptions() Options {
	return &options{}
}

func (o *options) Analysis(opts []*model.Options) {

}

// calculateXYValues calculates the x and y values for risk and reward graph
func calculateXYValues(contracts []*model.Options) []*model.XYValue {
	// Your code here
	return nil
}

func calculateMaxProfit(contracts []*model.Options) float64 {
	// Your code here
	return 0
}

func calculateMaxLoss(contracts []*model.Options) float64 {
	// Your code here
	return 0
}

func calculateBreakEvenPoints(contracts []*model.Options) []float64 {
	// Your code here
	return nil
}
