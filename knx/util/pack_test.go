// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package util

import (
	"bytes"
	"io"
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

func TestPackStringBoundaries(t *testing.T) {
	testCases := []struct {
		name     string
		maxLen   uint
		input    string
		want     []byte
	}{
		{name: "exact capacity", maxLen: 4, input: "1234", want: []byte{'1', '2', '3', 0}},
		{name: "over capacity", maxLen: 4, input: "12345", want: []byte{'1', '2', '3', 0}},
		{name: "zero capacity", maxLen: 0, input: "", want: nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			buffer := bytes.Repeat([]byte{0xFF}, int(testCase.maxLen))
			n, err := PackString(buffer, testCase.maxLen, testCase.input)
			if err != nil {
				t.Fatalf("PackString() error = %v", err)
			}
			if n != testCase.maxLen {
				t.Fatalf("PackString() wrote %d bytes, want %d", n, testCase.maxLen)
			}
			if !bytes.Equal(buffer, testCase.want) {
				t.Fatalf("PackString() = %v, want %v", buffer, testCase.want)
			}
		})
	}

	if _, err := PackString(make([]byte, 3), 4, "x"); err != io.ErrShortBuffer {
		t.Fatalf("PackString() short buffer error = %v, want %v", err, io.ErrShortBuffer)
	}
}
