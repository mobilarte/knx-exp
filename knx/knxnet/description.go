// Copyright 2022 Martin Müller
// Licensed under the MIT license which can be found in the LICENSE file.

package knxnet

import (
	"net"

	"github.com/mobilarte/knx-exp/knx/util"
)

// NewDescriptionReq creates a new Description Request, addr defines where
// the KNXnet/IP server should send the response to.
func NewDescriptionReq(addr net.Addr) (req *DescriptionReq, err error) {
	req = &DescriptionReq{}

	req.HostInfo, err = HostInfoFromAddress(addr)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// A DescriptionReq requests a description from a particular KNXnet/IP server via unicast.
type DescriptionReq struct {
	HostInfo
}

// Service returns the service identifier for a Description Request.
func (DescriptionReq) Service() ServiceID {
	return DescrReqService
}

// A DescriptionRes is a Description Response from a KNXnet/IP server.
type DescriptionRes DescriptionBlock

// Service returns the service identifier for a Description Response.
func (DescriptionRes) Service() ServiceID {
	return DescrResService
}

// Size returns the packed size of a Description Response.
func (res DescriptionRes) Size() uint {
	return res.DeviceHardware.Size() + res.SupportedServices.Size()
}

// Pack assembles the Description Response structure in the given buffer.
func (res *DescriptionRes) Pack(buffer []byte) {
	util.PackSome(buffer, res.DeviceHardware, res.SupportedServices)
}

// Unpack parses the given service payload in order to initialize the Description Response.
func (res *DescriptionRes) Unpack(data []byte) (n uint, err error) {
	return (*DescriptionBlock)(res).Unpack(data)
}
