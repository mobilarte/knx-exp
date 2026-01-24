// Copyright 2017 Ole Krüger.
// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"

	"math"
	"math/rand"
)

// Test DPT 14.xxx (F32)
func TestDPT_14(t *testing.T) {
	type DPT14 struct {
		Dpv       DatapointValue
		Min       float32
		MinStr    string
		Middle    float32
		MiddleStr string
		Max       float32
		MaxStr    string
	}

	var types_14 = []DPT14{
		{new(DPT_14000), math.SmallestNonzeroFloat32, "0.00 m/s²", 0,
			"0.00 m/s²", math.MaxFloat32, "340282346638528859811704183484516925440.00 m/s²"},
	}
	for _, e := range types_14 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default (0) value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		err := e.Dpv.Unpack(packF32(e.Min))
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%s", e.Dpv) != e.MinStr {
			t.Errorf("%#v has wrong smallest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err = e.Dpv.Unpack(packF32(e.Max))
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong largest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}

func TestDPT_14P(t *testing.T) {
	var dst DPT_14001

	for i := 1; i <= 1000; i++ {
		value := rand.Float32()
		src := DPT_14001(value)
		buf := src.Pack()

		bitrep := fmt.Sprintf("%08b%08b%08b%08b", buf[1], buf[2], buf[3], buf[4])
		if fmt.Sprintf("%032b", math.Float32bits(value)) != bitrep {
			t.Errorf("%#v has bad packing [%s]. Should be [%032b].", src, bitrep, math.Float32bits(value))
		}

		err := dst.Unpack(buf)
		if err != nil {
			t.Errorf("%#v error unpacking [%s].", dst, bitrep)
		}

		if math.Float32bits(value) != math.Float32bits(float32(dst)) {
			t.Errorf("%#v has bad unpacking [%s]. Should be [%032b].", src, bitrep, math.Float32bits(value))
		}
	}
}
