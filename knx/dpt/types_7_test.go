// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 7.xxx (U16)
func TestDPT_7(t *testing.T) {
	type DPT7 struct {
		Dpv    DatapointValue
		Min    uint16
		MinStr string
		Max    uint16
		MaxStr string
	}

	var types_7 = []DPT7{
		{new(DPT_7001), 0, "0 pulses", 65535, "65535 pulses"},
		{new(DPT_7002), 0, "0 ms", 65535, "65535 ms"},
		{new(DPT_7005), 0, "0 s", 65535, "65535 s"},
		{new(DPT_7006), 0, "0 min", 65535, "65535 min"},
		{new(DPT_7007), 0, "0 h", 65535, "65535 h"},
		{new(DPT_7010), 0, "0", 65535, "65535"},
		{new(DPT_7011), 0, "0 mm", 65535, "65535 mm"},
		{new(DPT_7012), 0, "0", 65535, "65535 mA"},
		{new(DPT_7013), 0, "0 lux", 65535, "65535 lux"},
		{new(DPT_7600), 0, "0 K", 65535, "65535 K"},
	}
	for _, e := range types_7 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MinStr {
			t.Errorf("%#v has wrong default value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err := e.Dpv.Unpack(packU16(e.Max))
		if err != nil {
			if fmt.Sprintf("%v", err) != "payload is not valid" {
				t.Errorf("%v", err)
			}
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong true value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}

	type DPT7SP struct {
		Dpv    DatapointValue
		Min    float32
		MinStr string
		Max    float32
		MaxStr string
	}

	var types_7sp = []DPT7SP{
		{new(DPT_7003), 0, "0 ms", 655.35, "655350 ms"},
		{new(DPT_7004), 0, "0 ms", 6553.5, "6553500 ms"},
	}

	for _, e := range types_7sp {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MinStr {
			t.Errorf("%#v has wrong default value \"%v\"! Should be \"%s\".", e.Dpv, e.Dpv, e.MinStr)
		}

		// Pack the largest number, then unpack it
		err := e.Dpv.Unpack(DPT_7001(65535).Pack())
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong true value \"%v\"! Should be \"%s\".", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
