// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_8001 represents DPT 8.001 (G) / DPT_Value_2_Count.
type DPT_8001 int16

func (d DPT_8001) Pack() []byte {
	return packV16(d)
}

func (d *DPT_8001) Unpack(data []byte) error {
	return unpackV16(data, d)
}

func (d DPT_8001) Unit() string {
	return "pulses"
}

func (d DPT_8001) IsValid() bool {
	return int16(d) != 0x7FFF
}

func (d DPT_8001) String() string {
	return fmt.Sprintf("%d pulses", d)
}

// DPT_8002 represents DPT 8.002 (G) / DPT_DeltaTimeMsec.
type DPT_8002 int16

func (d DPT_8002) Pack() []byte {
	return packV16(d)
}

func (d *DPT_8002) Unpack(data []byte) error {
	return unpackV16(data, d)
}

func (d DPT_8002) Unit() string {
	return "ms"
}

func (d DPT_8002) String() string {
	return fmt.Sprintf("%d ms", d)
}

// DPT_8003 represents DPT 8.002 (G) / DPT_DeltaTime10Msec.
type DPT_8003 float32

func (d DPT_8003) Pack() []byte {
	return packV16(int16(d * 100))
}

func (d *DPT_8003) Unpack(data []byte) error {
	var value int16

	if err := unpackV16(data, &value); err != nil {
		return err
	}

	*d = DPT_8003(float32(value) / 100)

	return nil
}

func (d DPT_8003) Unit() string {
	return "ms"
}

func (d DPT_8003) IsValid() bool {
	// TBD
	return true
}

func (d DPT_8003) String() string {
	return fmt.Sprintf("%f ms", d)
}

// DPT_8004 represents DPT 8.004 (G) / DPT_DeltaTime100Msec.
type DPT_8004 float32

func (d DPT_8004) Pack() []byte {
	return packV16(int16(d * 10))
}

func (d *DPT_8004) Unpack(data []byte) error {
	var value int16

	if err := unpackV16(data, &value); err != nil {
		return err
	}

	*d = DPT_8004(float32(value) / 10)

	return nil
}

func (d DPT_8004) Unit() string {
	return "ms"
}

func (d DPT_8004) String() string {
	return fmt.Sprintf("%f ms", d)
}

// DPT_8005 represents DPT 8.005 (G) / DPT_DeltaTimeSec.
type DPT_8005 int16

func (d DPT_8005) Pack() []byte {
	return packV16(d)
}

func (d *DPT_8005) Unpack(data []byte) error {
	return unpackV16(data, d)
}

func (d DPT_8005) Unit() string {
	return "s"
}

func (d DPT_8005) String() string {
	return fmt.Sprintf("%d s", d)
}

// DPT_8006 represents DPT 8.005 (G) / DPT_DeltaTimeMin.
type DPT_8006 int16

func (d DPT_8006) Pack() []byte {
	return packV16(d)
}

func (d *DPT_8006) Unpack(data []byte) error {
	return unpackV16(data, d)
}

func (d DPT_8006) Unit() string {
	return "min"
}

func (d DPT_8006) String() string {
	return fmt.Sprintf("%d min", d)
}

// DPT_8007 represents DPT 8.007 (G) / DPT_DeltaTimeHour.
type DPT_8007 int16

func (d DPT_8007) Pack() []byte {
	return packV16(d)
}

func (d *DPT_8007) Unpack(data []byte) error {
	return unpackV16(data, d)
}

func (d DPT_8007) Unit() string {
	return "h"
}

func (d DPT_8007) String() string {
	return fmt.Sprintf("%d h", d)
}

// DPT_8010 represents DPT 8.010 (G) / DPT_Percent_V16.
type DPT_8010 float64

func (d DPT_8010) Pack() []byte {
	var approx int16
	b := float64(d) * 100
	if b < -32768 {
		approx = -32768
	} else if b > 32767 {
		approx = 32767
	} else {
		approx = int16(b)
	}
	return packV16(approx)
}

func (d *DPT_8010) Unpack(data []byte) error {
	var b int16
	err := unpackV16(data, &b)
	if err != nil {
		return err
	}
	*d = DPT_8010(float64(b) / 100)
	return nil
}

func (d DPT_8010) Unit() string {
	return "%"
}

func (d DPT_8010) IsValid() bool {
	return float64(d) >= -327.68 && float64(d) <= 327.67
}

func (d DPT_8010) String() string {
	return fmt.Sprintf("%.2f %%", d)
}

// DPT_8011 represents DPT 8.011 (FB) / DPT_Rotation_Angle.
type DPT_8011 int16

func (d DPT_8011) Pack() []byte {
	return packV16(int16(d))
}

func (d *DPT_8011) Unpack(data []byte) error {
	return unpackV16(data, (*int16)(d))
}

func (d DPT_8011) Unit() string {
	return "°"
}

func (d DPT_8011) String() string {
	return fmt.Sprintf("%d°", int16(d))
}

// DPT_8012 represents DPT 8.012 (FB) / DPT_Length_m.
type DPT_8012 int16

func (d DPT_8012) Pack() []byte {
	return packV16(d)
}

func (d *DPT_8012) Unpack(data []byte) error {
	return unpackV16(data, d)
}

func (d *DPT_8012) Unit() string {
	return "m"
}

func (d DPT_8012) String() string {
	return fmt.Sprintf("%d m", d)
}
