package common

import (
	"errors"
	"math/big"
	"regexp"
)

func max(a, b *big.Float) *big.Float {
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
