// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 6.xxx (V8)
func TestDPT_6(t *testing.T) {
	type DPT6 struct {
		Dpv       DatapointValue
		Min       int8
		MinStr    string
		Middle    int8
		MiddleStr string
		Max       int8
		MaxStr    string
	}

	var types_6 = []DPT6{
		{new(DPT_6001), -128, "-128 %", 0, "0 %", 127, "127 %"},
		{new(DPT_6010), -128, "-128 counter pulses", 0, "0 counter pulses", 127, "127 counter pulses"},
	}

	for _, e := range types_6 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		e.Dpv.Unpack(packV8(e.Min))
		if fmt.Sprintf("%s", src) != e.MinStr {
			t.Errorf("%#v has wrong min value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		e.Dpv.Unpack(packV8(e.Max))
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong max value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
