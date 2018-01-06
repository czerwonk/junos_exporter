package rpc

import (
	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/connector"
)

type RpcClient struct {
	conn *connector.SshConnection
}

func NewClient(ssh *connector.SshConnection) *RpcClient {
	return &RpcClient{conn: ssh}
}

func (*RpcClient) AlarmCounter() (*alarm.AlarmCounter, error) {
	return &alarm.AlarmCounter{}, nil
}

func (*RpcClient) InterfaceStats() ([]*interfaces.InterfaceStats, error) {
	return make([]*interfaces.InterfaceStats, 0), nil
}

func (*RpcClient) BgpSessions() ([]*bgp.BgpSession, error) {
	return make([]*bgp.BgpSession, 0), nil
}