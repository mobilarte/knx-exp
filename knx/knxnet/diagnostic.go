// Copyright 2025 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package knxnet

import (
	"net"

	"github.com/mobilarte/knx-exp/knx/util"
)

type DiagnosticSelector uint8

const (
	// Programming mode selects the devices that are in Programming Mode.
	PrgModeSelector DiagnosticSelector = 0x01
	// MAC selects a device via its MAC address.
	MACSelector DiagnosticSelector = 0x02
)

type Selector struct {
	Length       uint8
	SelectorType DiagnosticSelector
	MACAddress   []uint8
}

func (sel *Selector) Set(progMode bool, macAddr net.HardwareAddr) {
	if progMode {
		sel.Length = 2
		sel.SelectorType = PrgModeSelector
	} else if macAddr != nil {
		sel.Length = 8
		sel.SelectorType = MACSelector
		sel.MACAddress = macAddr
	} else {
		sel.Length = 2
	}
}

// DiagnosticReq describes a diagnostic request from a particular
// KNXnet/IP Server via multicast or broadcast.
type DiagnosticReq struct {
	HostInfo
	Selector
}

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

// Service returns the service identifier for a Diagnostic Request.
func (DiagnosticReq) Service() ServiceID {
	return DiagnosticReqService
}

func (req *DiagnosticReq) SetSelector(progMode bool, macAddr net.HardwareAddr) {
	req.Set(progMode, macAddr)
}

// Size returns the size of HostInfo plus the variable size of the selector.
func (req DiagnosticReq) Size() uint {
	// TBD may be wrong!
	return req.HostInfo.Size() + uint(req.Length)
}

// Pack copies the DiagnosticReq structure to the buffer.
func (req DiagnosticReq) Pack(buffer []byte) {
	util.PackSome(
		buffer,
		byte(8),
		uint8(req.Protocol),
		req.Address[:],
		uint16(req.Port))

	if req.SelectorType == PrgModeSelector {
		util.PackSome(
			buffer[8:],
			req.Length,
			uint8(req.SelectorType),
		)
	} else {
		util.PackSome(
			buffer[8:],
			req.Length,
			uint8(req.SelectorType),
			req.MACAddress,
		)
	}
}

// DiagnosticRes is a Diagnostic Response from a KNXnet/IP server.
type DiagnosticRes struct {
	HostInfo
	Selector
	DescriptionBlock
}

// Service returns the service identifier for Diagnostic Response.
func (DiagnosticRes) Service() ServiceID {
	return DiagnosticResService
}

// Size returns the packed size of a Diagnostic Response.
func (res DiagnosticRes) Size() uint {
	// TBD may be wrong!
	return res.HostInfo.Size() + uint(res.Length)
}

// Pack assembles the Diagnostic Response structure in the given buffer.
func (res *DiagnosticRes) Pack(buffer []byte) {
	util.PackSome(buffer, res.HostInfo, res.Selector)
}

// Unpack parses the given service payload in order to initialize the Diagnostic Response.
func (res *DiagnosticRes) Unpack(data []byte) (n uint, err error) {
	return (*DescriptionBlock)(&res.DescriptionBlock).Unpack(data)
}

type BasicConfigurationReq struct {
	HostInfo
	Selector
	IPConfig IpConfigDIB
}

func NewBasicConfigReq(addr net.Addr) (*BasicConfigurationReq, error) {
	req := &BasicConfigurationReq{}
	req.IPConfig.DefaultGateway = Address{192, 168, 1, 2}

	var err error

	req.HostInfo, err = HostInfoFromAddress(addr)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (BasicConfigurationReq) Service() ServiceID {
	return BasicConfReqService
}

func (req *BasicConfigurationReq) SetSelector(progMode bool, macAddr net.HardwareAddr) {
	req.Set(progMode, macAddr)
}

func (req *BasicConfigurationReq) Pack(buffer []byte) {
	util.PackSome(buffer, req.HostInfo, req.Selector)
}

func (req *BasicConfigurationReq) Unpack(data []byte) (n uint, err error) {
	//return (*DescriptionBlock)(&res.DescriptionBlock).Unpack(data)
	return
}
