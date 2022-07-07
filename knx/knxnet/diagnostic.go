// Copyright (c) 2022 mobilarte.
// Licensed under the MIT license which can be found in the LICENSE file.

package knxnet

import (
	"fmt"
	"net"

	"github.com/mobilarte/knx-exp/knx/util"
)

type SelectorType uint8

const (
	// Programming mode selects the devices in Programming Mode.
	PrgModeSelector SelectorType = 0x01
	// MAC selects a device via MAC address.
	MACSelector SelectorType = 0x02
)

// NewDiagnosticReq creates a new Diagnostic Request, addr defines where
// KNXnet/IP server should send the response to.
func NewDiagnosticReq(addr net.Addr) (*DiagnosticReq, error) {
	req := &DiagnosticReq{}

	var err error

	req.HostInfo, err = HostInfoFromAddress(addr)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type Selector struct {
	Length     uint8
	Type       SelectorType
	MACAddress []uint8
}

// A DescriptionReq requests a description from a particular KNXnet/IP Server via unicast.
type DiagnosticReq struct {
	HostInfo
	Selector
}

// Service returns the service identifier for a Description Request.
func (DiagnosticReq) Service() ServiceID {
	return DiagnosticReqService
}

func (req *DiagnosticReq) SetSelector(progMode bool, macAddr net.HardwareAddr) {
	if progMode {
		req.Selector.Length = 2
		req.Selector.Type = PrgModeSelector
	} else if macAddr != nil {
		req.Selector.Length = 8
		req.Selector.Type = MACSelector
		req.MACAddress = macAddr
	} else {
		req.Selector.Length = 2
	}
}

// Size returns the size of HostInfo plus the variable size of the selector.
func (req DiagnosticReq) Size() uint {
	return req.HostInfo.Size() + uint(req.Selector.Length)
}

// Pack copies the DiagnosticReq structure to the buffer.
func (req DiagnosticReq) Pack(buffer []byte) {
	util.PackSome(
		buffer,
		byte(8),
		uint8(req.HostInfo.Protocol),
		req.HostInfo.Address[:],
		uint16(req.HostInfo.Port))

	if req.Type == PrgModeSelector {
		util.PackSome(
			buffer[8:],
			req.Selector.Length,
			uint8(req.Selector.Type),
		)
	} else {
		util.PackSome(
			buffer[8:],
			req.Selector.Length,
			uint8(req.Selector.Type),
			req.Selector.MACAddress,
		)
	}
}

// DiagnosticRes is a Diagnostic Response from a KNXnet/IP server.
type DiagnosticRes struct {
	HostInfo
	Selector
	DescriptionBlock
}

// Service returns the service identifier for Description Response.
func (DiagnosticRes) Service() ServiceID {
	return DiagnosticResService
}

// Size returns the packed size of a Description Response.
func (res DiagnosticRes) Size() uint {
	return res.HostInfo.Size() + uint(res.Selector.Length)
}

// Pack assembles the Description Response structure in the given buffer.
func (res *DiagnosticRes) Pack(buffer []byte) {
	util.PackSome(buffer, res.HostInfo, res.Selector)
}

// Unpack parses the given service payload in order to initialize the Description Response.
func (res *DiagnosticRes) Unpack(data []byte) (n uint, err error) {
	fmt.Printf("%#v\n", data)
	return 0, nil
	//return (*DescriptionBlock)(res).Unpack(data)
}
