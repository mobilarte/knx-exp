// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package knxnet

import (
	"testing"

	"github.com/mobilarte/knx-exp/knx/cemi"
	"github.com/mobilarte/knx-exp/knx/util"
)

func BenchmarkPack(b *testing.B) {
	b.ReportAllocs()

	lData := cemi.LData{
		Control1: cemi.Control1NoRepeat | cemi.Control1NoSysBroadcast |
			cemi.Control1StdFrame | cemi.Control1WantAck | cemi.Control1Prio(cemi.PrioLow),
		Control2:    cemi.Control2GroupAddr | cemi.Control2Hops(6),
		Source:      0,
		Destination: 0x1337,
		Data: &cemi.AppData{
			Command: cemi.GroupValueWrite,
			Data:    []byte{0, 0x13, 0x37},
		},
	}

	req := &TunnelReq{
		Channel:   1,
		SeqNumber: 0,
		Payload:   &cemi.LDataReq{LData: lData},
	}

	for b.Loop() {
		util.AllocAndPack(req)
	}
}

func TestUnpackRejectsPacketLengthMismatch(t *testing.T) {
	packet := AllocAndPack(&ConnStateRes{Channel: 1})
	packet[5]--

	var service Service
	if _, err := Unpack(packet, &service); err != ErrPacketLength {
		t.Fatalf("expected ErrPacketLength, got %v", err)
	}
}

func TestConnResUnpackRejectsTruncatedPrefix(t *testing.T) {
	var response ConnRes
	if _, err := response.Unpack(nil); err == nil {
		t.Fatal("expected an error for a truncated connection response")
	}
}

func TestDescriptionBlockUnpackRejectsInvalidDIBLengths(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{name: "length exceeds packet", data: []byte{16, byte(DescriptionTypeIPConfig)}},
		{name: "zero-length unsupported DIB", data: []byte{0, 0x77}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var block DescriptionBlock
			if _, err := block.Unpack(test.data); err == nil {
				t.Fatal("expected an error for an invalid DIB length")
			}
		})
	}
}

func TestBasicConfigurationReqUnpack(t *testing.T) {
	data := []byte{8, byte(UDP4), 192, 168, 1, 2, 0x0e, 0x57, 2, 0}
	var req BasicConfigurationReq

	n, err := req.Unpack(data)
	if err != nil {
		t.Fatalf("unexpected unpack error: %v", err)
	}
	if n != uint(len(data)) {
		t.Fatalf("unexpected number of bytes consumed: %d", n)
	}
	if req.SelectorType != 0 {
		t.Fatalf("unexpected selector type: %d", req.SelectorType)
	}
}

func TestBasicConfigurationReqUnpackRejectsTruncatedData(t *testing.T) {
	var req BasicConfigurationReq
	if _, err := req.Unpack(nil); err == nil {
		t.Fatal("expected an error for a truncated basic-configuration request")
	}
}
