package common

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
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

func FindProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parentDir
	}
}
