// Copyright 2024 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import "fmt"

// DPT_3007 represents DPT 3.007 (FB) / DPT_Control_Dimming.
type DPT_3007 struct {
	C        bool
	StepCode uint8
}

func (d DPT_3007) Pack() []byte {
	return packB1U3(d.C, d.StepCode)
}

func (d *DPT_3007) Unpack(data []byte) error {
	var inc bool
	var val uint8
	if err := unpackB1U3(data, &inc, &val); err != nil {
		return err
	}
	*d = DPT_3007{
		C:        inc,
		StepCode: val,
	}
	return nil
}

func (d DPT_3007) Unit() string {
	return ""
}

func (d *DPT_3007) IsValid() bool {
	return d.StepCode < 8
}

func (d DPT_3007) String() string {
	if d.C {
		return fmt.Sprintf("Increase by %d", d.StepCode)
	} else {
		return fmt.Sprintf("Decrease by %d", d.StepCode)
	}
}

// DPT_3008 represents DPT 3.008 (FB) / DPT_Control_Blinds.
type DPT_3008 struct {
	C        bool
	StepCode uint8
}

func (d DPT_3008) Pack() []byte {
	return packB1U3(d.C, d.StepCode)
}

func (d *DPT_3008) Unpack(data []byte) error {
	var inc bool
	var val uint8
	if err := unpackB1U3(data, &inc, &val); err != nil {
		return err
	}
	*d = DPT_3008{
		C:        inc,
		StepCode: val,
	}
	return nil
}

func (d DPT_3008) Unit() string {
	return ""
}

func (d *DPT_3008) IsValid() bool {
	return d.StepCode < 8
}

func (d DPT_3008) String() string {
	if d.C {
		return fmt.Sprintf("Increase by %d", d.StepCode)
	} else {
		return fmt.Sprintf("Decrease by %d", d.StepCode)
	}
}
