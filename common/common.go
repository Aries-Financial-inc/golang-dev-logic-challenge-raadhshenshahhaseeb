package common

import "math/big"

func max(a, b *big.Float) *big.Float {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}
