// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import "fmt"

// DPT_5001 represents DPT 5.001 (G) / DPT_Scaling.
type DPT_5001 float32

func (d DPT_5001) Pack() []byte {
	if d <= 0 {
		return packU8(uint8(0))
	} else if d >= 100 {
		return packU8(uint8(255))
	} else {
		return packU8(uint8(d * 2.55))
	}
}

func (d *DPT_5001) Unpack(data []byte) error {
	var value uint8

	if err := unpackU8(data, &value); err != nil {
		return err
	}

	*d = DPT_5001(value) / 2.55

	return nil
}

func (d DPT_5001) Unit() string {
	return "%"
}

func (d DPT_5001) String() string {
	return fmt.Sprintf("%.2f %%", float32(d))
}

// DPT_5003 represents DPT 5.003 (G) / DPT_Angle.
type DPT_5003 float32

func (d DPT_5003) Pack() []byte {
	if d <= 0 {
		return packU8(uint8(0))
	} else if d >= 360 {
		return packU8(uint8(255))
	} else {
		return packU8(uint8(d * 255 / 360))
	}
}

func (d *DPT_5003) Unpack(data []byte) error {
	var value uint8

	if err := unpackU8(data, &value); err != nil {
		return err
	}

	*d = DPT_5003(float32(value) * 360 / 255)

	return nil
}

func (d DPT_5003) Unit() string {
	return "°"
}

func (d DPT_5003) String() string {
	return fmt.Sprintf("%d°", int(d))
}

// DPT_5004 represents DPT 5.004 (FB) / DPT_Percent_U8.
type DPT_5004 uint8

func (d DPT_5004) Pack() []byte {
	return packU8(d)
}

func (d *DPT_5004) Unpack(data []byte) error {
	return unpackU8(data, d)
}

func (d DPT_5004) Unit() string {
	return "%"
}

func (d DPT_5004) String() string {
	return fmt.Sprintf("%d %%", d)
}

// DPT_5005 represents DPT 5.005 (G) / DPT_DecimalFactor.
type DPT_5005 uint8

func (d DPT_5005) Pack() []byte {
	return packU8(d)
}

func (d *DPT_5005) Unpack(data []byte) error {
	return unpackU8(data, d)
}

func (d DPT_5005) Unit() string {
	return ""
}

func (d DPT_5005) String() string {
	return fmt.Sprintf("%d", d)
}

// DPT_5006 represents DPT 5.006 (G) / DPT_Tariff.
type DPT_5006 uint8

func (d DPT_5006) Pack() []byte {
	return packU8(d)
}

func (d *DPT_5006) Unpack(data []byte) error {
	if err := unpackU8(data, d); err != nil {
		return err
	}

	return nil
}

func (d DPT_5006) Unit() string {
	return ""
}

func (d DPT_5006) String() string {
	if !d.IsValid() {
		return "Reserved"
	}
	return fmt.Sprintf("%d", d)
}

func (d DPT_5006) IsValid() bool {
	// 255 is reserved; shall not be used (transmitted or received)
	return d != 255
}

// DPT_5010 represents DPT 5.010 (G) / DPT_Value_1_Ucount.
type DPT_5010 uint8

func (d DPT_5010) Pack() []byte {
	return packU8(d)
}

func (d *DPT_5010) Unpack(data []byte) error {
	return unpackU8(data, d)
}

func (d DPT_5010) Unit() string {
	return "counter pulses"
}

func (d DPT_5010) String() string {
	return fmt.Sprintf("%d counter pulses", d)
}
