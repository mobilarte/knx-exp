// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package cemi

import (
	"bytes"
	crand "crypto/rand"
	"testing"

	"github.com/mobilarte/knx-exp/knx/util"
)

func makeRandInfoSegment() []byte {
	n := util.Randint64() % 256

	buffer := make([]byte, n+1)

	buffer[0] = byte(n)
	if _, err := crand.Read(buffer[1:]); err != nil {
		panic(err)
	}

	return buffer
}

func TestInfo_Unpack(t *testing.T) {
	for range 100 {
		data := makeRandInfoSegment()
		info := Info{}

		num, err := info.Unpack(data)
		if err != nil {
			t.Error("Unexpected error:", err, data)
			continue
		}

		if num != uint(len(data)) {
			t.Error("Unexpected length:", num, len(data), data)
		}

		if !bytes.Equal([]byte(info), data[1:]) {
			t.Error("Unexpected result:", info, data)
		}
	}
}

func TestInfo_UnpackTruncated(t *testing.T) {
	var info Info
	if _, err := info.Unpack([]byte{2, 0xAA}); err == nil {
		t.Fatal("expected an error for truncated info data")
	}
}

func TestControl2Hops(t *testing.T) {
	for hops := uint8(0); hops <= 7; hops++ {
		control := Control2Hops(hops) | Control2GroupAddr
		if got := control.Hops(); got != hops {
			t.Errorf("Hops() = %d, want %d", got, hops)
		}
	}
}

func TestLRawIndMessageCode(t *testing.T) {
	if got := (LRawInd{}).MessageCode(); got != LRawIndCode {
		t.Fatalf("MessageCode() = %v, want %v", got, LRawIndCode)
	}
}

func TestUnpack(t *testing.T) {
	ldataCodes := []MessageCode{LDataReqCode, LDataConCode, LDataIndCode}

	for range 100 {
		code := ldataCodes[util.Randint64()%3]
		data := append([]byte{byte(code)}, makeRandLData()...)

		var msg Message

		num, err := Unpack(data, &msg)
		if err != nil {
			t.Error("Unexpected error:", err, data)
			continue
		}

		if num != uint(len(data)) {
			t.Error("Unexpected length:", num, len(data), data)
		}
	}
}
