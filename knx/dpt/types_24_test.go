// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package dpt

import (
	"fmt"
	"strings"
	"testing"
)

// Test DPT 24.001 with some special inputs
func TestDPT_24001(t *testing.T) {
	var dst DPT_24001

	// Test with ASCII string
	src := DPT_24001("KNX is OK")
	buf := src.Pack()

	if fmt.Sprintf("%x", buf) != "004b4e58206973204f4b00" {
		t.Errorf("Error packing [%s] => [%x]", src, buf)
	}

	err := dst.Unpack(buf)
	if err != nil {
		t.Errorf("Error unpacking [%x]", buf)
	}

	// Test with an empty string
	src = DPT_24001("")
	buf = src.Pack()
	err = dst.Unpack(buf)

	if err != nil {
		t.Errorf("Error unpacking [%x] => [%s]", buf, dst)
	}

	if dst.String() != src.String() {
		t.Errorf("Error comparing [%s]/%d => [%s]/%d", src, len(src), dst, len(dst))
	}

	// Test with a very long string
	src = DPT_24001(strings.Repeat("|01234567890123456789", 50))
	buf = src.Pack()
	err = dst.Unpack(buf)

	if err != nil {
		t.Errorf("Error unpacking [%x] => [%s]", buf, dst)
	}

	if dst.String() != string(src) {
		t.Errorf("Error comparing [%s]/%d => [%s]/%d", src, len(src), dst, len(dst))
	}
	// Test with Latin1 chars
	src = DPT_24001("Latin1:çäüöàéè")
	buf = src.Pack()
	err = dst.Unpack(buf)

	if err != nil {
		t.Errorf("Unpacking error [%x] => [%s]", buf, dst)
	}

	if dst.String() != src.String() {
		t.Errorf("Unexpected ISO-8859-1 chars [%s] => [%s]", src, dst)
	}

	// Test with mixed chars, 1, 2 and 3 bytes
	// Length = 17 = 3 * 2 + 5 * 1 + 2 * 3
	src = DPT_24001("çäühello你好")
	expected := "çäühello  "
	buf = src.Pack()
	err = dst.Unpack(buf)

	if err != nil {
		t.Errorf("Error unpacking [%s] => [%s]", src, dst)
	}

	if dst.String() != expected {
		t.Errorf("Error when comparing [%s] => [%s]", src, dst)
	}

	// 13 = 3 * 2 + 8 is the length in byte, hence in ISO chars,
	// because non ISO will be replaced by 0x20 (space).
	// 10 is the length in rune len([]rune(expected))
	if len(dst) != len(expected) {
		t.Errorf("Lost or gained something [%s] => [%s]", src, dst)
	}

	// A very long string will be truncated. Still possible to send it!
	src = DPT_24001(`Lorem Ipsum is simply dummy text of the printing and typesetting industry. 
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer 
took a galley of type and scrambled it to make a type specimen book. It has survived not only five 
centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was 
popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and 
more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`)

	// length in bytes!
	if len(src) > 255 && src.IsValid() {
		t.Errorf("Should not be valid [%s]", src)
	}
}
