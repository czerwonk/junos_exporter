package rpc

import (
	"encoding/xml"
	"fmt"

	"log"

	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/ospf"
	"github.com/czerwonk/junos_exporter/route"
)

type RpcClient struct {
	conn  *connector.SshConnection
	debug bool
}

func NewClient(ssh *connector.SshConnection, debug bool) *RpcClient {
	return &RpcClient{conn: ssh, debug: debug}
}

func (c *RpcClient) AlarmCounter() (*alarm.AlarmCounter, error) {
	red := 0
	yellow := 0

	cmds := []string{"show system alarms", "show chassis alarms"}

	for _, cmd := range cmds {
		var a = AlarmRpc{}
		err := c.runCommandAndParse(cmd, &a)
		if err != nil {
			return nil, err
		}

		for _, d := range a.Information.Details {
			if d.Class == "Major" {
				red++
			} else if d.Class == "Minor" {
				yellow++
			}
		}
	}

	return &alarm.AlarmCounter{RedCount: float64(red), YellowCount: float64(yellow)}, nil
}

func (c *RpcClient) InterfaceStats() ([]*interfaces.InterfaceStats, error) {
	var x = InterfaceRpc{}
	err := c.runCommandAndParse("show interfaces statistics detail", &x)
	if err != nil {
		return nil, err
	}

	stats := make([]*interfaces.InterfaceStats, 0)
	for _, phy := range x.Information.Interfaces {
		s := &interfaces.InterfaceStats{
			IsPhysical:     true,
			Name:           phy.Name,
			Description:    phy.Description,
			Mac:            phy.MacAddress,
			ReceiveDrops:   float64(phy.InputErrors.Drops),
			ReceiveErrors:  float64(phy.InputErrors.Errors),
			ReceiveBytes:   float64(phy.Stats.InputBytes),
			TransmitDrops:  float64(phy.OutputErrors.Drops),
			TransmitErrors: float64(phy.OutputErrors.Errors),
			TransmitBytes:  float64(phy.Stats.OutputBytes),
		}

		stats = append(stats, s)

		for _, log := range phy.LogicalInterfaces {
			sl := &interfaces.InterfaceStats{
				IsPhysical:    false,
				Name:          log.Name,
				Description:   log.Description,
				Mac:           phy.MacAddress,
				ReceiveBytes:  float64(log.Stats.InputBytes),
				TransmitBytes: float64(log.Stats.OutputBytes),
			}

			stats = append(stats, sl)
		}
	}

	return stats, nil
}

func (c *RpcClient) BgpSessions() ([]*bgp.BgpSession, error) {
	var x = BgpRpc{}
	err := c.runCommandAndParse("show bgp summary", &x)
	if err != nil {
		return nil, err
	}

	sessions := make([]*bgp.BgpSession, 0)
	for _, peer := range x.Information.Peers {
		s := &bgp.BgpSession{
			Ip:               peer.Ip,
			Up:               peer.State == "Established",
			Asn:              peer.Asn,
			Flaps:            float64(peer.Flaps),
			InputMessages:    float64(peer.InputMessages),
			OutputMessages:   float64(peer.OutputMessages),
			AcceptedPrefixes: float64(peer.Rib.AcceptedPrefixes),
			ActivePrefixes:   float64(peer.Rib.ActivePrefixes),
			ReceivedPrefixes: float64(peer.Rib.ReceivedPrefixes),
			RejectedPrefixes: float64(peer.Rib.RejectedPrefixes),
		}

		sessions = append(sessions, s)
	}

	return sessions, nil
}

func (c *RpcClient) OspfAreas() ([]*ospf.OspfArea, error) {
	var x = Ospf3Rpc{}
	err := c.runCommandAndParse("show ospf3 overview", &x)
	if err != nil {
		return nil, err
	}

	areas := make([]*ospf.OspfArea, 0)
	for _, area := range x.Information.Overview.Areas {
		a := &ospf.OspfArea{
			Name:      area.Name,
			Neighbors: float64(area.Neighbors.NeighborsUp),
		}

		areas = append(areas, a)
	}

	return areas, nil
}

func (c *RpcClient) RoutingTables() ([]*route.RoutingTable, error) {
	var x = RouteRpc{}
	err := c.runCommandAndParse("show route summary", &x)
	if err != nil {
		return nil, err
	}

	tables := make([]*route.RoutingTable, 0)
	for _, table := range x.Information.Tables {
		t := &route.RoutingTable{
			Name:         table.Name,
			MaxRoutes:    float64(table.MaxRoutes),
			ActiveRoutes: float64(table.ActiveRoutes),
			TotalRoutes:  float64(table.TotalRoutes),
			Protocols:    make([]*route.ProtocolRouteCount, 0),
		}

		for _, proto := range table.Protocols {
			p := &route.ProtocolRouteCount{
				Name:         proto.Name,
				Routes:       float64(proto.Routes),
				ActiveRoutes: float64(proto.ActiveRoutes),
			}

			t.Protocols = append(t.Protocols, p)
		}

		tables = append(tables, t)
	}

	return tables, nil
}

func (c *RpcClient) runCommandAndParse(cmd string, obj interface{}) error {
	if c.debug {
		log.Printf("Running command on %s: %s\n", c.conn.Host, cmd)
	}

	b, err := c.conn.RunCommand(fmt.Sprintf("%s | display xml", cmd))
	if err != nil {
		return err
	}

	if c.debug {
		log.Printf("Output for %s: %s\n", c.conn.Host, string(b))
	}

	err = xml.Unmarshal(b, obj)
	return err
}
