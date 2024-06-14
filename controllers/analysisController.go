package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/common/stderr"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-raadhshenshahhaseeb/services/optionsContract"
)

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	StrikePrice    float64   `json:"strike_price"`
	Type           string    `json:"type"`
	Bid            float64   `json:"bid"`
	Ask            float64   `json:"ask"`
	LongShort      string    `json:"long_short"`
	ExpirationDate time.Time `json:"expiration_date"`
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type analysisController struct {
	svc optionsContract.Service
}

type AnalysisController interface {
	AnalysisHandler(ctx *gin.Context)
}

func NewAnalysisController(svc optionsContract.Service) AnalysisController {
	return &analysisController{svc: svc}
}

func (a *analysisController) AnalysisHandler(ctx *gin.Context) {
	req := make([]OptionsContract, 0)

	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(int(stderr.ErrBindingTheRequest), "unable to perform analysis: internal server error", nil))
		return
	}

	validationCode := ValidateOptionsArray(req)

	if validationCode != 0 {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(int(validationCode), "invalid request", nil))
		return
	}

	optsObj := make([]*optionsContract.Options, len(req))

	for i, v := range req {
		optsObj[i] = &optionsContract.Options{
			StrikePrice:    v.StrikePrice,
			Type:           v.Type,
			Bid:            v.Bid,
			Ask:            v.Ask,
			LongShort:      v.LongShort,
			ExpirationDate: v.ExpirationDate,
		}
	}

	resp, err := a.svc.Analysis(optsObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewStandardResponse(int(validationCode), "unable to perform the analysis", nil))
		return
	}

	xy := make([]XYValue, len(resp.XYValues))

	for i, v := range resp.XYValues {
		xy[i] = XYValue{
			X: v.X,
			Y: v.Y,
		}
	}

	ctx.JSON(http.StatusOK, NewStandardResponse(stderr.SuccessCode(), "success", &AnalysisResponse{
		XYValues:        xy,
		MaxProfit:       resp.MaxProfit,
		MaxLoss:         resp.MaxLoss,
		BreakEvenPoints: resp.BreakEvenPoints,
	}))
}

func ValidateOptionsArray(opts []OptionsContract) stderr.ErrorCode {
	currentTime := time.Now()
	validPositions := map[string]bool{"long": true, "short": true}
	validTypes := map[string]bool{"call": true, "put": true}

	for _, o := range opts {
		if o.Bid > o.Ask {
			return stderr.ErrBidCannotBeGreaterThanAsk
		}

		if o.Bid < 0 || o.Ask < 0 {
			return stderr.ErrBidAndAskCannotBeNegative
		}

		if o.ExpirationDate.Before(currentTime) {
			return stderr.ErrExpiryCannotBeInPast
		}

		if o.StrikePrice < 0 {
			return stderr.ErrStrikeCannotBeNegative
		}

		if !validPositions[strings.ToLower(o.LongShort)] {
			return stderr.ErrPositionNotSupportedOrInvalid
		}

		if !validTypes[strings.ToLower(o.Type)] {
			return stderr.ErrQuoteNotSupportedOrInvalid
		}
	}

	return 0
}
