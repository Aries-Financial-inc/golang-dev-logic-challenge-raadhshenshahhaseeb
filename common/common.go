package common

import (
	"errors"
	"math/big"
	"regexp"
)

func Max(a, b *big.Float) *big.Float {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

func CheckString(s string) error {
	regex := regexp.MustCompile("^[a-z]+$")

	if !regex.MatchString(s) {
		return errors.New("string contains invalid characters")
	}
	return nil
}

const (
	QUOTE_BUY  = "LONG"
	QUOTE_SELL = "SHORT"

	TYPE_CALL = "CALL"
	TYPE_PUT  = "PUT"
)
