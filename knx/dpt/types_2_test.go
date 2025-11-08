// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 2.xxx (B₂)
func TestDPT_2(t *testing.T) {
	type DPT2 struct {
		Dpv     DatapointValue
		Control bool
		Value   bool
		OnPrint string
	}

	var types_2 = []DPT2{
		{new(DPT_2001), false, false, "c=false, v=false"},
		{new(DPT_2001), true, false, "c=true, v=false"},
		{new(DPT_2001), false, true, "c=false, v=true"},
		{new(DPT_2001), true, true, "c=true, v=true"},
		{new(DPT_2002), false, false, "c=false, v=false"},
		{new(DPT_2002), true, false, "c=true, v=false"},
		{new(DPT_2002), false, true, "c=false, v=true"},
		{new(DPT_2002), true, true, "c=true, v=true"},
	}

	for _, e := range types_2 {
		if e.Dpv.(fmt.Stringer).String() != "c=false, v=false" {
			t.Errorf("%#v has wrong default value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.OnPrint)
		}

		_ = e.Dpv.Unpack(packB2(e.Control, e.Value))
		if fmt.Sprintf("%s", e.Dpv) != e.OnPrint {
			t.Errorf("%#v has wrong value [%v] after unpack. Should be [%s].", e.Dpv, e.Dpv, e.OnPrint)
		}
	}
}
