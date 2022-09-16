// Copyright 2017 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

// Described in 03_08_07 KNXnetIP Remote Configuration and Diagnosis v01.01.02 AS.pdf

package knx

import (
	"net"
	"time"

	"github.com/mobilarte/knx-exp/knx/knxnet"
)

func DiagnosticWithMAC(multicastDiscoveryAddress string, macAddr net.HardwareAddr, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	return DiagnosticOnInterface(nil, multicastDiscoveryAddress, macAddr, false, searchTimeout)
}

func DiagnosticInProgMode(multicastDiscoveryAddress string, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	return DiagnosticOnInterface(nil, multicastDiscoveryAddress, nil, true, searchTimeout)
}

// DiagnosticOnInterface sends the diagnostic request on a specified interface. If the
// interface is nil, the system-assigned multicast interface is used.
func DiagnosticOnInterface(ifi *net.Interface, multicastDiscoveryAddress string, macAddr net.HardwareAddr,
	progMode bool, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	socket, err := knxnet.ListenRouterOnInterface(ifi, multicastDiscoveryAddress, false)

	if err != nil {
		return nil, err
	}
	defer socket.Close()

	req, err := knxnet.NewDiagnosticReq(socket.LocalAddr())

	if err != nil {
		return nil, err
	}

	req.SetSelector(progMode, macAddr)

	if err := socket.Send(req); err != nil {
		return nil, err
	}
	results := []*knxnet.DiagnosticRes{}
	timeout := time.After(searchTimeout)

loop:
	for {
		select {
		case msg := <-socket.Inbound():
			diagnosticRes, ok := msg.(*knxnet.DiagnosticRes)
			if !ok {
				continue
			}
			results = append(results, diagnosticRes)

		case <-timeout:
			break loop
		}
	}

	return results, nil
}
