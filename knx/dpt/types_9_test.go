// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"testing"
)

// Test DPT 9.00x (F16)
func TestDPT_9001(t *testing.T) {
	type DPT9 struct {
		Dpv       DatapointValue
		Min       float64
		MinStr    string
		Middle    float64
		MiddleStr string
		Max       float64
		MaxStr    string
	}
	var types_9 = []DPT9{
		{new(DPT_9001), 20.48, "20.48 °C", 0, "0.00 °C", 670433.28, "670433.28 °C"},
		{new(DPT_9001), -273, "-273.00 °C", 0, "0.00 °C", 670433.28, "670433.28 °C"},
		{new(DPT_9002), -671088.64, "-671088.64 K", 0, "0.00 K", 670433.28, "670433.28 K"},
		{new(DPT_9003), -671088.64, "-671088.64 K/h", 0, "0.00 K/h", 670433.28, "670433.28 K/h"},
		{new(DPT_9004), 0, "0.00 lx", 0, "0.00 lx", 670433.28, "670433.28 lx"},
		{new(DPT_9005), 0, "0.00 m/s", 0, "0.00 m/s", 670433.28, "670433.28 m/s"},
		{new(DPT_9006), 0, "0.00 Pa", 0, "0.00 Pa", 670433.28, "670433.28 Pa"},
		{new(DPT_9007), 0, "0.00 %", 0, "0.00 %", 670433.28, "670433.28 %"},
		{new(DPT_9008), 0, "0.00 ppm", 0, "0.00 ppm", 670433.28, "670433.28 ppm"},
		{new(DPT_9009), 0, "0.00 m³/h", 0, "0.00 m³/h", 670433.28, "670433.28 m³/h"},
		{new(DPT_9010), -671088.64, "-671088.64 s", 0, "0.00 s", 670433.28, "670433.28 s"},
		{new(DPT_9011), -671088.64, "-671088.64 ms", 0, "0.00 ms", 670433.28, "670433.28 ms"},
		{new(DPT_9020), -671088.64, "-671088.64 mV", 0, "0.00 mV", 670433.28, "670433.28 mV"},
		{new(DPT_9021), -671088.64, "-671088.64 mA", 0, "0.00 mA", 670433.28, "670433.28 mA"},
		{new(DPT_9022), -671088.64, "-671088.64 W/m²", 0, "0.00 W/m²", 670433.28, "670433.28 W/m²"},
		{new(DPT_9023), -671088.64, "-671088.64 K/%", 0, "0.00 K/%", 670433.28, "670433.28 K/%"},
		{new(DPT_9024), -671088.64, "-671088.64 kW", 0, "0.00 kW", 670433.28, "670433.28 kW"},
		{new(DPT_9025), -671088.64, "-671088.64 l/h", 0, "0.00 l/h", 670433.28, "670433.28 l/h"},
		{new(DPT_9026), -671088.64, "-671088.64 l/m²", 0, "0.00 l/m²", 670433.28, "670433.28 l/m²"},
		{new(DPT_9027), -459.6, "-459.60 °F", 0, "0.00 °F", 670433.28, "670433.28 °F"},
		{new(DPT_9028), 0, "0.00 km/h", 0, "0.00 km/h", 670433.28, "670433.28 km/h"},
		{new(DPT_9029), 0, "0.00 g/m³", 0, "0.00 g/m³", 670433.28, "670433.28 g/m³"},
		{new(DPT_9030), 0, "0.00 μg/m³", 0, "0.00 μg/m³", 670433.28, "670433.28 μg/m³"},
	}

	for _, e := range types_9 {
		src := e.Dpv
		if fmt.Sprintf("%s", src) != e.MiddleStr {
			t.Errorf("%#v has wrong default (0) value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MiddleStr)
		}

		err := e.Dpv.Unpack(packF16(e.Min))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MinStr {
			t.Errorf("%#v has wrong smallest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MinStr)
		}

		err = e.Dpv.Unpack(packF16(e.Max))
		if err != nil {
			t.Errorf("%v", err)
		}
		if fmt.Sprintf("%s", e.Dpv) != e.MaxStr {
			t.Errorf("%#v has wrong largest value [%v]. Should be [%s].", e.Dpv, e.Dpv, e.MaxStr)
		}
	}
}
