// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package util

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Randint64() int64 {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0
	}
	return val.Int64()
}
