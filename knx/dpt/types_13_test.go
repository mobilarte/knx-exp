// Copyright 2017 Ole Krüger.
// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 13.xxx (V32)
func TestDPT_13(t *testing.T) {
	type DPT13 struct {
		Dpv       DatapointValue
		Min       int32
		MinStr    string
		Middle    int32
		MiddleStr string
		Max       int32
		MaxStr    string
	}

	var types_13 = []DPT13{
		{new(DPT_13001), -2147483648, "-2147483648 counter pulses", 0, "0 counter pulses",
			2147483647, "2147483647 counter pulses"},
		{new(DPT_13002), -2147483648, "-2147483648 m³/h", 0, "0 m³/h", 2147483647, "2147483647 m³/h"},
		{new(DPT_13010), -2147483648, "-2147483648 Wh", 0, "0 Wh", 2147483647, "2147483647 Wh"},
		{new(DPT_13011), -2147483648, "-2147483648 VAh", 0, "0 VAh", 2147483647, "2147483647 VAh"},
		{new(DPT_13012), -2147483648, "-2147483648 VARh", 0, "0 VARh", 2147483647, "2147483647 VARh"},
		{new(DPT_13013), -2147483648, "-2147483648 kWh", 0, "0 kWh", 2147483647, "2147483647 kWh"},
		{new(DPT_13014), -2147483648, "-2147483648 kVAh", 0, "0 kVAh", 2147483647, "2147483647 kVAh"},
		{new(DPT_13015), -2147483648, "-2147483648 kVARh", 0, "0 kVARh", 2147483647, "2147483647 kVARh"},
		{new(DPT_13016), -2147483648, "-2147483648 MWh", 0, "0 MWh", 2147483647, "2147483647 MWh"},
	}
	for _, e := range types_13 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default (0) value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		err := e.Dpv.Unpack(packV32(e.Min))
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%s", e.Dpv) != e.MinStr {
			t.Errorf("%#v has wrong smallest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err = e.Dpv.Unpack(packV32(e.Max))
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong largest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
