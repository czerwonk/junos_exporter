// SPDX-License-Identifier: MIT

package main

import (
	"regexp"

	"github.com/czerwonk/junos_exporter/pkg/features/ddosprotection"
	"github.com/czerwonk/junos_exporter/pkg/features/poe"

	"github.com/czerwonk/junos_exporter/internal/config"
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/features/accounting"
	"github.com/czerwonk/junos_exporter/pkg/features/alarm"
	"github.com/czerwonk/junos_exporter/pkg/features/arp"
	"github.com/czerwonk/junos_exporter/pkg/features/bfd"
	"github.com/czerwonk/junos_exporter/pkg/features/bgp"
	"github.com/czerwonk/junos_exporter/pkg/features/environment"
	"github.com/czerwonk/junos_exporter/pkg/features/firewall"
	"github.com/czerwonk/junos_exporter/pkg/features/fpc"
	"github.com/czerwonk/junos_exporter/pkg/features/interfacediagnostics"
	"github.com/czerwonk/junos_exporter/pkg/features/interfacequeue"
	"github.com/czerwonk/junos_exporter/pkg/features/interfaces"
	"github.com/czerwonk/junos_exporter/pkg/features/ipsec"
	"github.com/czerwonk/junos_exporter/pkg/features/isis"
	"github.com/czerwonk/junos_exporter/pkg/features/krt"
	"github.com/czerwonk/junos_exporter/pkg/features/l2circuit"
	"github.com/czerwonk/junos_exporter/pkg/features/l2vpn"
	"github.com/czerwonk/junos_exporter/pkg/features/lacp"
	"github.com/czerwonk/junos_exporter/pkg/features/ldp"
	"github.com/czerwonk/junos_exporter/pkg/features/mac"
	"github.com/czerwonk/junos_exporter/pkg/features/macsec"
	"github.com/czerwonk/junos_exporter/pkg/features/mplslsp"
	"github.com/czerwonk/junos_exporter/pkg/features/nat"
	"github.com/czerwonk/junos_exporter/pkg/features/nat2"
	"github.com/czerwonk/junos_exporter/pkg/features/ospf"
	"github.com/czerwonk/junos_exporter/pkg/features/power"
	"github.com/czerwonk/junos_exporter/pkg/features/route"
	"github.com/czerwonk/junos_exporter/pkg/features/routingengine"
	"github.com/czerwonk/junos_exporter/pkg/features/rpki"
	"github.com/czerwonk/junos_exporter/pkg/features/rpm"
	"github.com/czerwonk/junos_exporter/pkg/features/security"
	"github.com/czerwonk/junos_exporter/pkg/features/securityike"
	"github.com/czerwonk/junos_exporter/pkg/features/securitypolicies"
	"github.com/czerwonk/junos_exporter/pkg/features/storage"
	"github.com/czerwonk/junos_exporter/pkg/features/subscriber"
	"github.com/czerwonk/junos_exporter/pkg/features/system"
	"github.com/czerwonk/junos_exporter/pkg/features/vpws"
	"github.com/czerwonk/junos_exporter/pkg/features/vrrp"
)

type collectors struct {
	logicalSystem string
	collectors    map[string]collector.RPCCollector
	devices       map[string][]collector.RPCCollector
	cfg           *config.Config
}

func collectorsForDevices(devices []*connector.Device, cfg *config.Config, logicalSystem string) *collectors {
	c := &collectors{
		logicalSystem: logicalSystem,
		collectors:    make(map[string]collector.RPCCollector),
		devices:       make(map[string][]collector.RPCCollector),
		cfg:           cfg,
	}

	for _, d := range devices {
		c.initCollectorsForDevices(d, deviceInterfaceRegex(cfg, d.Host))
	}

	return c
}

func (c *collectors) initCollectorsForDevices(device *connector.Device, descRe *regexp.Regexp) {
	f := c.cfg.FeaturesForDevice(device.Host)

	c.devices[device.Host] = make([]collector.RPCCollector, 0)

	c.addCollectorIfEnabledForDevice(device, "routingengine", f.RoutingEngine, routingengine.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "accounting", f.Accounting, accounting.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "alarm", f.Alarm, func() collector.RPCCollector {
		return alarm.NewCollector(*alarmFilter)
	})
	c.addCollectorIfEnabledForDevice(device, "bfd", f.BFD, bfd.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "bgp", f.BGP, func() collector.RPCCollector {
		return bgp.NewCollector(c.logicalSystem, descRe)
	})
	c.addCollectorIfEnabledForDevice(device, "env", f.Environment, environment.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "firewall", f.Firewall, firewall.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "fpc", f.FPC, fpc.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ifacediag", f.InterfaceDiagnostic, func() collector.RPCCollector {
		return interfacediagnostics.NewCollector(descRe)
	})
	c.addCollectorIfEnabledForDevice(device, "ifacequeue", f.InterfaceQueue, func() collector.RPCCollector {
		return interfacequeue.NewCollector(descRe)
	})
	c.addCollectorIfEnabledForDevice(device, "iface", f.Interfaces, func() collector.RPCCollector {
		return interfaces.NewCollector(descRe)
	})
	c.addCollectorIfEnabledForDevice(device, "ipsec", f.IPSec, ipsec.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "isis", f.ISIS, isis.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "l2c", f.L2Circuit, l2circuit.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "l2vpn", f.L2Vpn, l2vpn.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "lacp", f.LACP, lacp.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ldp", f.LDP, ldp.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "nat", f.NAT, nat.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "nat2", f.NAT2, nat2.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ospf", f.OSPF, func() collector.RPCCollector {
		return ospf.NewCollector(c.logicalSystem)
	})
	c.addCollectorIfEnabledForDevice(device, "routes", f.Routes, route.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "rpki", f.RPKI, rpki.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "rpm", f.RPM, rpm.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "security", f.Security, security.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "security_ike", f.SecurityIKE, securityike.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "security_policies", f.SecurityPolicies, securitypolicies.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "storage", f.Storage, storage.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "system", (f.System || f.License), system.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "power", f.Power, power.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "mac", f.MAC, mac.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "vrrp", f.VRRP, vrrp.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "vpws", f.VPWS, vpws.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "mpls_lsp", f.MPLSLSP, mplslsp.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "subscriber", f.Subscriber, subscriber.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "macsec", f.MACSec, macsec.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "arp", f.ARP, arp.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "poe", f.Poe, poe.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ddosprotection", f.DDOSProtection, ddosprotection.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "krt", f.KRT, krt.NewCollector)
}

func (c *collectors) addCollectorIfEnabledForDevice(device *connector.Device, key string, enabled bool, newCollector func() collector.RPCCollector) {
	if !enabled {
		return
	}

	col, found := c.collectors[key]
	if !found {
		col = newCollector()
		c.collectors[key] = col
	}

	c.devices[device.Host] = append(c.devices[device.Host], col)
}

func (c *collectors) allEnabledCollectors() []collector.RPCCollector {
	collectors := make([]collector.RPCCollector, len(c.collectors))

	i := 0
	for _, collector := range c.collectors {
		collectors[i] = collector
		i++
	}

	return collectors
}

func (c *collectors) collectorsForDevice(device *connector.Device) []collector.RPCCollector {
	cols, found := c.devices[device.Host]
	if !found {
		return []collector.RPCCollector{}
	}

	return cols
}
