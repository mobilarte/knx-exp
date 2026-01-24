// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package util

import (
	"testing"
)

func TestPackString(t *testing.T) {
	testCases := []struct {
		MaxLen   uint
		Data     string
		Expected []byte
	}{
		{
			MaxLen: 30,
			Data:   "ABB IPS/S2.1",
			Expected: []byte{0x41, 0x42, 0x42, 0x20, 0x49, 0x50, 0x53, 0x2f,
				0x53, 0x32, 0x2e, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00},
		},
	}

	for _, testCase := range testCases {
		buffer := make([]byte, testCase.MaxLen)
		for i := range buffer {
			buffer[i] = 0xff
		}

		n, err := PackString(buffer, testCase.MaxLen, testCase.Data)
		if err == nil {
			if testCase.MaxLen != n {
				t.Errorf("Produced bytes not equal")
			}

			if len(testCase.Expected) != len(buffer) {
				t.Errorf("Packed buffer not equal")
			}

			for i, v := range testCase.Expected {
				if v != buffer[i] {
					t.Errorf("Packed buffer not equal")
					break
				}
			}

			for i, v := range buffer {
				if v != testCase.Expected[i] {
					t.Errorf("Packed buffer not equal")
					break
				}
			}
		}
		/*
			if assert.Nil(t, err) {
				assert.Equal(t, testCase.MaxLen, n, "Produced bytes not equal")
				assert.Equal(t, testCase.Expected, buffer, "Packed buffer not equal")
			}
		*/
	}
}
