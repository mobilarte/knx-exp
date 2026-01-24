// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package util

import (
	crand "crypto/rand"
	"math"
	"math/big"
)

// Randint64 returns a random int64.
func Randint64() int64 {
	val, err := crand.Int(crand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0
	}

	return val.Int64()
}
