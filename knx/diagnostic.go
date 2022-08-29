// Licensed under the MIT license which can be found in the LICENSE file.

package knx

import (
	"net"
	"time"

	"github.com/mobilarte/knx-exp/knx/knxnet"
)

// Diagnostic all KNXnet/IP servers.
func Diagnostic(multicastDiscoveryAddress string, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	return DiagnosticOnInterface(nil, multicastDiscoveryAddress, nil, false, searchTimeout)
}

func DiagnosticWithMAC(multicastDiscoveryAddress string, macAddr net.HardwareAddr, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	return DiagnosticOnInterface(nil, multicastDiscoveryAddress, macAddr, false, searchTimeout)
}

func DiagnosticInProgMode(multicastDiscoveryAddress string, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	return DiagnosticOnInterface(nil, multicastDiscoveryAddress, nil, true, searchTimeout)
}

// DiscoverOnInterface discovers all KNXnet/IP Server on a specific interface. If the
// interface is nil, the system-assigned multicast interface is used.
func DiagnosticOnInterface(ifi *net.Interface, multicastDiscoveryAddress string, macAddr net.HardwareAddr,
	progMode bool, searchTimeout time.Duration) ([]*knxnet.DiagnosticRes, error) {
	socket, err := knxnet.ListenRouterOnInterface(ifi, multicastDiscoveryAddress, false)

	if err != nil {
		return nil, err
	}
	defer socket.Close()

	addr := socket.LocalAddr()

	req, err := knxnet.NewDiagnosticReq(addr)

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
