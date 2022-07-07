// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

func abs[T float64 | float32 | ~int64](x T) T {
	if x < 0.0 {
		return -x
	}
	return x
}
