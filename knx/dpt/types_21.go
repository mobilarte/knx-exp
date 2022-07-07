// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
)

// DPT_21001 represents DPT 21.001 (G) / DPT_StatusGen.
type DPT_21001 uint8

func (d DPT_21001) Pack() []byte {
	return packU8(d)
}

func (d *DPT_21001) Unpack(data []byte) error {
	if err := unpackU8(data, d); err != nil {
		return err
	}

	if d.IsValid() {
		return nil
	}
	return ErrBadReservedBits
}

func (d DPT_21001) Unit() string {
	return ""
}

func (d DPT_21001) IsValid() bool {
	// At most 5 bits set.
	return d <= 0x1F
}

func (d DPT_21001) String() string {
	// Print the rightmost 5 bits
	// b0 - OutOfService
	// b1 - Fault
	// b2 - Overriden
	// b3 - InAlarm
	// b4 - AlarmUnAck
	return fmt.Sprintf("%05b", d)
}

// DPT_21002 represents DPT 21.002 (G) / DPT_Device_Control.
type DPT_21002 uint8

func (d DPT_21002) Pack() []byte {
	return packU8(d)
}

func (d *DPT_21002) Unpack(data []byte) error {
	if err := unpackU8(data, d); err != nil {
		return err
	}

	if d.IsValid() {
		return nil
	}
	return ErrBadReservedBits
}

func (d DPT_21002) Unit() string {
	return ""
}

func (d DPT_21002) IsValid() bool {
	// At most 3 bits set.
	return d <= 0x7
}

func (d DPT_21002) String() string {
	// Print the rightmost 3 bits
	// b0 - UserStopped
	// b1 - OwnIA
	// b2 - VerifyMode
	return fmt.Sprintf("%03b", d)
}
