// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import "fmt"

// DPT_3007 represents DPT 3.007 (FB) / DPT_Control_Dimming.
type DPT_3007 struct {
	Increase bool
	Value    uint8
}

func (d DPT_3007) Pack() []byte {
	return packB1U3(d.Increase, d.Value)
}

func (d *DPT_3007) Unpack(data []byte) error {
	var inc bool
	var val uint8
	if err := unpackB1U3(data, &inc, &val); err != nil {
		return err
	}
	*d = DPT_3007{
		Increase: inc,
		Value:    val,
	}
	return nil
}

func (d DPT_3007) Unit() string {
	return ""
}

func (d *DPT_3007) IsValid() bool {
	return d.Value < 8
}

func (d DPT_3007) String() string {
	if d.Increase {
		return fmt.Sprintf("Increase by %d", d.Value)
	} else {
		return fmt.Sprintf("Decrease by %d", d.Value)
	}
}
