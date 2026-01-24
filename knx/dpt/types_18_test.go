// Copyright 2017 Ole Krüger.
// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"testing"
)

// Test DPT 18.001 (B1r1U6)
func TestDPT_18(t *testing.T) {
	type DPT18 struct {
		Dpv   DatapointValue
		Value uint8
		Out   string
	}

	var types_18 = []DPT18{
		{new(DPT_18001), 0, "activate 1"},
		{new(DPT_18001), 63, "activate 64"},
		{new(DPT_18001), 96, "not valid"},
		{new(DPT_18001), 128, "learn 1"},
		{new(DPT_18001), 191, "learn 64"},
		{new(DPT_18001), 243, "not valid"},
	}

	for _, e := range types_18 {
		src := DPT_18001(e.Value)
		if src.IsValid() {
			if src.String() != e.Out {
				t.Errorf("%#v has wrong value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.Out)
			}

			var dst DPT_18001

			buf := src.Pack()

			err := dst.Unpack(buf)
			if err != nil {
				t.Errorf("Error unpacking")
			}

			if dst.String() != src.String() {
				t.Errorf("%#v not identical after unpacking [%s]. Should be [%s]", e.Dpv, dst.String(), src.String())
			}
		} else if e.Out != "not valid" {
			t.Errorf("%#v not recognized as invalid [%d].", e.Dpv, e.Value)
		}
	}
}
