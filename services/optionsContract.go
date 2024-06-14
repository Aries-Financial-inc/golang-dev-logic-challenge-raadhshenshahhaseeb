package services

import (
	"fmt"
	"math"
	"strings"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/common"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/model"
)

type OptionsSvc interface {
	Analysis(opts []*model.Options) (*OptionsAnalysis, error)
}

type optionSvc struct {
	MarketSvc MarketPriceSvc
}

type OptionsAnalysis struct {
	XYValues        []*XYValue
	MaxProfit       float64
	MaxLoss         float64
	BreakEvenPoints []float64
}

type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// NewOptionsSvc initializes a new options service
func NewOptionsSvc(mktSvc MarketPriceSvc) OptionsSvc {
	return &optionSvc{MarketSvc: mktSvc}
}

func (o *optionSvc) Analysis(opts []*model.Options) (*OptionsAnalysis, error) {
	analysis := new(OptionsAnalysis)

	token, err := o.MarketSvc.Get()
	if err != nil {
		return nil, fmt.Errorf("unable to get the token price: %w", err)
	}

	analysis.MaxProfit, analysis.MaxLoss, analysis.BreakEvenPoints = o.CalculateMaxProfitAndLoss(opts, token.Price)
	analysis.XYValues = o.CalculateXYValues(opts, token.Price)

	return analysis, nil
}

// CalculateXYValues calculates the x and y values for risk and reward graph
func (o *optionSvc) CalculateXYValues(opts []*model.Options, price float64) []*XYValue {

	return nil
}

func (o *optionSvc) CalculateMaxProfitAndLoss(opts []*model.Options, price float64) (float64, float64, []float64) {
	maxProfit := 0.0
	maxLoss := 0.0
	breakEvenPoints := make([]float64, 0)

	for _, opt := range opts {
		gains := 0.0

		if strings.Compare(strings.ToUpper(opt.LongShort), common.QUOTE_BUY) == 0 {
			if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_CALL) == 0 {
				gains = math.Max(price-opt.StrikePrice, 0) - opt.Ask
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice+opt.Ask)
			} else if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_PUT) == 0 {
				gains = math.Max(opt.StrikePrice-price, 0) - opt.Ask
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice-opt.Ask)
			}

		} else if strings.Compare(strings.ToUpper(opt.LongShort), common.QUOTE_SELL) == 0 {
			if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_CALL) == 0 {
				gains = opt.Bid - math.Max(price-opt.StrikePrice, 0)
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice+opt.Bid)
			} else if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_PUT) == 0 {
				gains = opt.Bid - math.Max(opt.StrikePrice-price, 0)
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice-opt.Bid)
			}
		}

		if gains < 0 {
			maxLoss -= -gains
		} else if gains > 0 {
			maxProfit += gains
		}
	}

	return maxProfit, maxLoss, breakEvenPoints
}
