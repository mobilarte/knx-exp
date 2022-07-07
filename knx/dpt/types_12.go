// Copyright 2017 Ole Krüger.
// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_12001 represents DPT 12.001 (G) / DPT_Value_4_Ucount.
type DPT_12001 uint32

func (d DPT_12001) Pack() []byte {
	return packU32(d)
}

func (d *DPT_12001) Unpack(data []byte) error {
	return unpackU32(data, d)
}

func (d DPT_12001) Unit() string {
	return "counter pulses"
}

func (d DPT_12001) String() string {
	return fmt.Sprintf("%d counter pulses", uint32(d))
}

// DPT_12100 represents DPT 12.100 (G) / DPT_LongTimePeriod_Sec.
type DPT_12100 uint32

func (d DPT_12100) Pack() []byte {
	return packU32(d)
}

func (d *DPT_12100) Unpack(data []byte) error {
	return unpackU32(data, d)
}

func (d DPT_12100) Unit() string {
	return "s"
}

func (d DPT_12100) String() string {
	return fmt.Sprintf("%d s", uint32(d))
}

// DPT_12101 represents DPT 12.100 (G) / DPT_LongTimePeriod_Min.
type DPT_12101 uint32

func (d DPT_12101) Pack() []byte {
	return packU32(d)
}

func (d *DPT_12101) Unpack(data []byte) error {
	return unpackU32(data, d)
}

func (d DPT_12101) Unit() string {
	return "min"
}

func (d DPT_12101) String() string {
	return fmt.Sprintf("%d min", uint32(d))
}

// DPT_12102 represents DPT 12.100 (G) / DPT_LongTimePeriod_Hrs.
type DPT_12102 uint32

func (d DPT_12102) Pack() []byte {
	return packU32(d)
}

func (d *DPT_12102) Unpack(data []byte) error {
	return unpackU32(data, d)
}

func (d DPT_12102) Unit() string {
	return "h"
}

func (d DPT_12102) String() string {
	return fmt.Sprintf("%d h", uint32(d))
}

// DPT_121200 represents DPT 12.1200 / DPT_VolumeLiquid_Litre.
type DPT_121200 uint32

func (d DPT_121200) Pack() []byte {
	return packU32(d)
}

func (d *DPT_121200) Unpack(data []byte) error {
	return unpackU32(data, d)
}

func (d DPT_121200) Unit() string {
	return "litre"
}

func (d DPT_121200) String() string {
	return fmt.Sprintf("%d litre", uint32(d))
}

// DPT_121201 represents DPT 12.1201 / DPT_Volume_m3.
type DPT_121201 uint32

func (d DPT_121201) Pack() []byte {
	return packU32(d)
}

func (d *DPT_121201) Unpack(data []byte) error {
	return unpackU32(data, d)
}

func (d DPT_121201) Unit() string {
	return "m³"
}

func (d DPT_121201) String() string {
	return fmt.Sprintf("%d m³", uint32(d))
}
