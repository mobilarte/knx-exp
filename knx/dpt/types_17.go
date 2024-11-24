// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_17001 represents DPT 17.001 (G) / DPT_SceneNumber.
type DPT_17001 uint8

func (d DPT_17001) Pack() []byte {
	if d.IsValid() {
		return packU8(d)
	} else {
		return packU8(uint8(63))
	}
}

func (d *DPT_17001) Unpack(data []byte) error {
	var value uint8

	if err := unpackU8(data, &value); err != nil {
		return err
	}

	if value <= 63 {
		*d = DPT_17001(value)
		return nil
	} else {
		*d = DPT_17001(63)
		return nil
	}
}

func (d DPT_17001) Unit() string {
	return ""
}

func (d DPT_17001) IsValid() bool {
	return d <= 0x3F
}

func (d DPT_17001) String() string {
	return fmt.Sprintf("%d", uint8(d))
}
