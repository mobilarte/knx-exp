// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_7001 represents DPT 7.001 (G) / DPT_Value_2_Ucount.
type DPT_7001 uint16

func (d DPT_7001) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7001) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7001) Unit() string {
	return "pulses"
}

func (d DPT_7001) String() string {
	return fmt.Sprintf("%d pulses", uint16(d))
}

// DPT_7002 represents DPT 7.002 (FB) / DPT_Time Period MSec.
type DPT_7002 uint16

func (d DPT_7002) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7002) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7002) Unit() string {
	return "ms"
}

func (d DPT_7002) String() string {
	return fmt.Sprintf("%d ms", uint16(d))
}

// DPT_7003 represents DPT 7.003 (G) / DPT_TimePeriod10MSec.
type DPT_7003 float32

func (d DPT_7003) Pack() []byte {
	return packU16(uint16(d * 100))
}

func (d *DPT_7003) Unpack(data []byte) error {
	var buf uint16
	err := unpackU16(data, &buf)
	if err == nil {
		*d = (DPT_7003)(float32(buf) * 10)
		return nil
	}
	return err
}

func (d DPT_7003) Unit() string {
	return "ms"
}

func (d DPT_7003) String() string {
	return fmt.Sprintf("%d ms", uint32(d))
}

func (d DPT_7003) IsValid() bool {
	return d >= 0 && d <= 655.35
}

// DPT_7004 represents DPT 7.004 (G) / DPT_TimePeriod100MSec.
type DPT_7004 float32

func (d DPT_7004) Pack() []byte {
	return packU16(uint16(d * 100))
}

func (d *DPT_7004) Unpack(data []byte) error {
	var buf uint16
	err := unpackU16(data, &buf)
	if err == nil {
		*d = (DPT_7004)(float32(buf) * 100)
		return nil
	}
	return err
}

func (d DPT_7004) Unit() string {
	return "ms"
}

func (d DPT_7004) String() string {
	return fmt.Sprintf("%d ms", uint32(d))
}

func (d DPT_7004) IsValid() bool {
	return d >= 0 && d <= 6553.5
}

// DPT_7005 represents DPT 7.005 (G) / DPT_TimePeriodSec.
type DPT_7005 uint16

func (d DPT_7005) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7005) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7005) Unit() string {
	return "s"
}

func (d DPT_7005) String() string {
	return fmt.Sprintf("%d s", uint16(d))
}

// DPT_7006 represents DPT 7.006 (G) / DPT_TimePeriodMin.
type DPT_7006 uint16

func (d DPT_7006) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7006) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7006) Unit() string {
	return "min"
}

func (d DPT_7006) String() string {
	return fmt.Sprintf("%d min", uint16(d))
}

// DPT_7007 represents DPT 7.007 (G) / DPT_TimePeriodHrs.
type DPT_7007 uint16

func (d DPT_7007) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7007) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7007) Unit() string {
	return "h"
}

func (d DPT_7007) String() string {
	return fmt.Sprintf("%d h", uint16(d))
}

// DPT_7010 represents DPT 7.010 (FB) / DPT_PropData_Type.
type DPT_7010 uint16

func (d DPT_7010) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7010) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7010) Unit() string {
	return ""
}

func (d DPT_7010) String() string {
	return fmt.Sprintf("%d", uint16(d))
}

// DPT_7011 represents DPT 7.011 (FB SAB)/ DPT_Length_mm.
type DPT_7011 uint16

func (d DPT_7011) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7011) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7011) Unit() string {
	return "mm"
}

func (d DPT_7011) String() string {
	return fmt.Sprintf("%d mm", uint16(d))
}

// DPT_7012 represents DPT 7.012 (FB) / DPT_UEICurrentmA.
type DPT_7012 uint16

func (d DPT_7012) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7012) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7012) Unit() string {
	return "mA"
}

func (d DPT_7012) String() string {
	if d == 0 {
		return "0"
	}
	return fmt.Sprintf("%d mA", uint16(d))
}

// DPT_7013 represents DPT 7.013 (FB) / DPT_Brightness.
type DPT_7013 uint16

func (d DPT_7013) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7013) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7013) Unit() string {
	return "lx"
}

func (d DPT_7013) String() string {
	return fmt.Sprintf("%d lx", uint16(d))
}

// DPT_7600 represents DPT 7.600 (FB) / DPT_Absolute_Colour_Temperature.
type DPT_7600 uint16

func (d DPT_7600) Pack() []byte {
	return packU16(d)
}

func (d *DPT_7600) Unpack(data []byte) error {
	return unpackU16(data, d)
}

func (d DPT_7600) Unit() string {
	return "K"
}

func (d DPT_7600) String() string {
	return fmt.Sprintf("%d K", uint16(d))
}
