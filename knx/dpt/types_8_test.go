// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 8.xxx (V16)
func TestDPT_8(t *testing.T) {
	type DPT8 struct {
		Dpv       DatapointValue
		Min       int16
		MinStr    string
		Middle    int16
		MiddleStr string
		Max       int16
		MaxStr    string
	}
	var types_8 = []DPT8{
		{new(DPT_8001), -32768, "-32768 pulses", 0, "0 pulses", 32767, "32767 pulses"},
		{new(DPT_8012), -32768, "-32768 m", 0, "0 m", 32767, "32767 m"},
	}
	for _, e := range types_8 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default (0) value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		err := e.Dpv.Unpack(packV16(e.Min))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MinStr {
			t.Errorf("%#v has wrong smallest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err = e.Dpv.Unpack(packV16(e.Max))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong largest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
