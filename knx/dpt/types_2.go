// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import "fmt"

// DPT_2001 represents DPT 2.001 (G) / DPT_Switch_Control.
type DPT_2001 struct {
	Control bool
	Value   bool
}

func (d DPT_2001) Pack() []byte {
	return packB2(d.Control, d.Value)
}

func (d *DPT_2001) Unpack(data []byte) error {
	return unpackB2(data, &d.Control, &d.Value)
}

func (d DPT_2001) Unit() string {
	return ""
}

func (d DPT_2001) String() string {
	return fmt.Sprintf("c=%t, v=%t", d.Control, d.Value)
}

// DPT_2002 represents DPT 2.002 (G) / DPT_Bool_Control.
type DPT_2002 struct {
	Control bool
	Value   bool
}

func (d DPT_2002) Pack() []byte {
	return packB2(d.Control, d.Value)
}

func (d *DPT_2002) Unpack(data []byte) error {
	return unpackB2(data, &d.Control, &d.Value)
}

func (d DPT_2002) Unit() string {
	return ""
}

func (d DPT_2002) String() string {
	return fmt.Sprintf("c=%t, v=%t", d.Control, d.Value)
}
