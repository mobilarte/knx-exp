// Copyright 2025 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

package knx

import (
	"time"

	"github.com/mobilarte/knx-exp/knx/knxnet"
)

// Describe a single KNXnet/IP server. Uses unicast UDP, address format is "ip:port".
func DescribeTunnel(address string, searchTimeout time.Duration) (*knxnet.DescriptionRes, error) {
	socket, err := knxnet.DialTunnelUDP(address)
	if err != nil {
		return nil, err
	}
	defer socket.Close()

	req, err := knxnet.NewDescriptionReq(socket.LocalAddr())
	if err != nil {
		return nil, err
	}

	if err := socket.Send(req); err != nil {
		return nil, err
	}

	timeout := time.After(searchTimeout)

	for {
		select {
		case msg := <-socket.Inbound():
			descriptionRes, ok := msg.(*knxnet.DescriptionRes)
			if ok {
				return descriptionRes, nil
			}

		case <-timeout:
			return nil, nil
		}
	}
}
