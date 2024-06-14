package optionsContract

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/common"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/marketPrice"
)

type Options struct {
	StrikePrice    float64   `json:"strike_price"`
	Type           string    `json:"type"`
	Bid            float64   `json:"bid"`
	Ask            float64   `json:"ask"`
	LongShort      string    `json:"long_short"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type OptionsSvc interface {
	Analysis(opts []*Options) (*OptionsAnalysis, error)
	AnalysisCalculation(opts []*Options, price float64) (float64, float64, []float64, []XYValue)
}

type optionSvc struct {
	MarketSvc marketPrice.MarketPriceSvc
}

type OptionsAnalysis struct {
	XYValues        []XYValue
	MaxProfit       float64
	MaxLoss         float64
	BreakEvenPoints []float64
}

type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// NewOptionsSvc initializes a new options service
func NewOptionsSvc(mktSvc marketPrice.MarketPriceSvc) OptionsSvc {
	return &optionSvc{MarketSvc: mktSvc}
}

func (o *optionSvc) Analysis(opts []*Options) (*OptionsAnalysis, error) {
	analysis := new(OptionsAnalysis)

	token, err := o.MarketSvc.Get()
	if err != nil {
		return nil, fmt.Errorf("unable to get the token price: %w", err)
	}

	analysis.MaxProfit, analysis.MaxLoss, analysis.BreakEvenPoints, analysis.XYValues = o.AnalysisCalculation(opts, token.Price)

	return analysis, nil
}

func (o *optionSvc) AnalysisCalculation(opts []*Options, price float64) (float64, float64, []float64, []XYValue) {
	maxProfit := 0.0
	maxLoss := 0.0
	breakEvenPoints := make([]float64, 0)
	riskRewardGraph := make([]XYValue, 0)

	for _, opt := range opts {
		gains := 0.0

		if strings.Compare(strings.ToUpper(opt.LongShort), common.QUOTE_BUY) == 0 {
			if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_CALL) == 0 {
				// Gains = max(Current Price - Strike Price) - Ask Price
				gains = math.Max(price-opt.StrikePrice, 0) - opt.Ask
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice+opt.Ask)
			} else if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_PUT) == 0 {
				// Gains = max(Strike Price - Current Price) - Ask Price
				gains = math.Max(opt.StrikePrice-price, 0) - opt.Ask
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice-opt.Ask)
			}

		} else if strings.Compare(strings.ToUpper(opt.LongShort), common.QUOTE_SELL) == 0 {
			if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_CALL) == 0 {
				// Gains = Bid Price - max(Current Price - Strike Price)
				gains = opt.Bid - math.Max(price-opt.StrikePrice, 0)
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice+opt.Bid)
			} else if strings.Compare(strings.ToUpper(opt.Type), common.TYPE_PUT) == 0 {
				// Gains = Bid Price - max(Strike Price - Current Price)
				gains = opt.Bid - math.Max(opt.StrikePrice-price, 0)
				breakEvenPoints = append(breakEvenPoints, opt.StrikePrice-opt.Bid)
			}
		}

		riskRewardGraph = append(riskRewardGraph, XYValue{
			X: price,
			Y: gains,
		})

		if gains < 0 {
			maxLoss += -gains
		} else if gains > 0 {
			maxProfit += gains
		}
	}

	return maxProfit, maxLoss, breakEvenPoints, riskRewardGraph
}
