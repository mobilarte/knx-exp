// Copyright 2025 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 12.xxx (U32)
func TestDPT_12(t *testing.T) {
	type DPT12 struct {
		Dpv       DatapointValue
		Min       uint32
		MinStr    string
		Middle    uint32
		MiddleStr string
		Max       uint32
		MaxStr    string
	}
	var types_12 = []DPT12{
		{new(DPT_12001), 0, "0 counter pulses", 0, "0 counter pulses", 4294967295, "4294967295 counter pulses"},
		{new(DPT_12100), 0, "0 s", 0, "0 s", 4294967295, "4294967295 s"},
		{new(DPT_12101), 0, "0 min", 0, "0 min", 4294967295, "4294967295 min"},
	}

	for _, e := range types_12 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default (0) value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		err := e.Dpv.Unpack(packU32(e.Min))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MinStr {
			t.Errorf("%#v has wrong smallest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err = e.Dpv.Unpack(packU32(e.Max))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong largest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
