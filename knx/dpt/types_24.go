// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"unicode"
)

// DPT_24001 represents DPT 24.001 (G) / DPT_VarString_8859_1.
// Non-ISO-8859-1 chars will be replaced with a space = 0x20.
// String longer than 254 chars will be truncated.
// KNX needs 1 byte for length.
type DPT_24001 string

func (d DPT_24001) Pack() []byte {
	r := []rune(d)
	length := len(r)

	var buffer = make([]byte, length+2)

	for i := 0; i < length; i++ {
		if r[i] > unicode.MaxLatin1 {
			buffer[i+1] = 0x20
		} else {
			buffer[i+1] = byte(r[i])
		}
	}

	buffer[length+1] = 0x0

	return buffer
}

func (d *DPT_24001) Unpack(data []byte) error {
	buffer := make([]rune, len(data[1:len(data)-1]))

	for i, c := range data[1 : len(data)-1] {
		buffer[i] = rune(c)
	}

	*d = DPT_24001(buffer)

	return nil
}

func (d DPT_24001) Unit() string {
	return ""
}

func (d DPT_24001) IsValid() bool {
	for _, c := range d {
		if c > unicode.MaxLatin1 {
			return false
		}
	}
	// Payload larger than that will be truncated
	return len(d) <= 255
}

func (d DPT_24001) String() string {
	return string(d)
}
