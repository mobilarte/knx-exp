// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"testing"
)

// Test DPT 3.00x B₁U₃
func TestDPT_3007(t *testing.T) {
	var (
		buf      []byte
		src, dst DPT_3007
	)

	for _, c := range []bool{true, false} {
		for v := 0; v < 8; v++ {
			src = DPT_3007{C: c, StepCode: uint8(v)}
			buf = src.Pack()

			err := dst.Unpack(buf)
			if err != nil {
				t.Errorf("Error unpacking [%x]", buf)
			}

			if dst.C != c {
				t.Errorf("%#v has wrong direction [%t]. Should be [%t].", dst, dst.C, c)
			}

			if dst.StepCode != uint8(v) {
				t.Errorf("%#v has wrong increase [%v]. Should be [%d].", dst, dst.StepCode, uint8(v))
			}
		}
	}
}
