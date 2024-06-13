package model

import "time"

type Options struct {
	// StrikePrice is a predetermined price at which the shares of stock will be exchanged if the option is exercised.
	StrikePrice    float64   `json:"strike_price"`
	Type           string    `json:"type"`
	Bid            float64   `json:"bid"`
	Ask            float64   `json:"ask"`
	LongShort      string    `json:"long_short"`
	ExpirationDate time.Time `json:"expiration_date"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
