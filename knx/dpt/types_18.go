// Copyright 2017 Ole Krüger.
// Copyright 2025 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_18001 represents DPT 18.001 (G) / DPT_SceneControl.
type DPT_18001 uint8

func (d DPT_18001) Pack() []byte {
	if d.IsValid() {
		return packU8(d)
	} else {
		return packU8(uint8(0x3F))
	}
}

func (d *DPT_18001) Unpack(data []byte) error {
	if err := unpackU8(data, d); err != nil {
		return err
	}

	if !d.IsValid() {
		return ErrOutOfRange
	}

	return nil
}

func (d DPT_18001) Unit() string {
	return ""
}

func (d DPT_18001) IsValid() bool {
	return d <= 0x3F || (d >= 0x80 && d <= 0xBF)
}

// String returns a string description.
// KNX Association recommends to display the scene numbers [1..64].
// Displaying value like ETS.
// See note 6 of the KNX Specifications v2.1.
func (d DPT_18001) String() string {
	if d >= 0x80 && d <= 0xBF {
		return fmt.Sprintf("learn %d", uint8(d)-128+1)
	} else if d <= 0x3F {
		return fmt.Sprintf("activate %d", uint8(d)+1)
	}

	return "invalid payload"
}
