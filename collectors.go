package main

import (
	"github.com/czerwonk/junos_exporter/accounting"
	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/config"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/environment"
	"github.com/czerwonk/junos_exporter/firewall"
	"github.com/czerwonk/junos_exporter/fpc"
	"github.com/czerwonk/junos_exporter/interfacediagnostics"
	"github.com/czerwonk/junos_exporter/interfacelabels"
	"github.com/czerwonk/junos_exporter/interfacequeue"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/ipsec"
	"github.com/czerwonk/junos_exporter/isis"
	"github.com/czerwonk/junos_exporter/l2circuit"
	"github.com/czerwonk/junos_exporter/ldp"
	"github.com/czerwonk/junos_exporter/mac"
	"github.com/czerwonk/junos_exporter/nat"
	"github.com/czerwonk/junos_exporter/ospf"
	"github.com/czerwonk/junos_exporter/power"
	"github.com/czerwonk/junos_exporter/route"
	"github.com/czerwonk/junos_exporter/routingengine"
	"github.com/czerwonk/junos_exporter/rpki"
	"github.com/czerwonk/junos_exporter/rpm"
	"github.com/czerwonk/junos_exporter/storage"
	"github.com/czerwonk/junos_exporter/system"
	"github.com/czerwonk/junos_exporter/vrrp"
)

type collectors struct {
	logicalSystem string
	dynamicLabels *interfacelabels.DynamicLabels
	collectors    map[string]collector.RPCCollector
	devices       map[string][]collector.RPCCollector
	cfg           *config.Config
}

func collectorsForDevices(devices []*connector.Device, cfg *config.Config, logicalSystem string, dynamicLabels *interfacelabels.DynamicLabels) *collectors {
	c := &collectors{
		logicalSystem: logicalSystem,
		dynamicLabels: dynamicLabels,
		collectors:    make(map[string]collector.RPCCollector),
		devices:       make(map[string][]collector.RPCCollector),
		cfg:           cfg,
	}

	for _, d := range devices {
		c.initCollectorsForDevices(d)
	}

	return c
}

func (c *collectors) initCollectorsForDevices(device *connector.Device) {
	f := c.cfg.FeaturesForDevice(device.Host)

	c.devices[device.Host] = make([]collector.RPCCollector, 0)

	c.addCollectorIfEnabledForDevice(device, "routingengine", f.RoutingEngine, routingengine.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "accounting", f.Accounting, accounting.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "alarm", f.Alarm, func() collector.RPCCollector {
		return alarm.NewCollector(*alarmFilter)
	})
	c.addCollectorIfEnabledForDevice(device, "bgp", f.BGP, func() collector.RPCCollector {
		return bgp.NewCollector(c.logicalSystem)
	})
	c.addCollectorIfEnabledForDevice(device, "env", f.Environment, environment.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "firewall", f.Firewall, firewall.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "fpc", f.FPC, fpc.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ifacediag", f.InterfaceDiagnostic, func() collector.RPCCollector {
		return interfacediagnostics.NewCollector(c.dynamicLabels)
	})
	c.addCollectorIfEnabledForDevice(device, "ifacequeue", f.InterfaceQueue, func() collector.RPCCollector {
		return interfacequeue.NewCollector(c.dynamicLabels)
	})
	c.addCollectorIfEnabledForDevice(device, "iface", f.Interfaces, func() collector.RPCCollector {
		return interfaces.NewCollector(c.dynamicLabels)
	})
	c.addCollectorIfEnabledForDevice(device, "ipsec", f.IPSec, ipsec.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "isis", f.ISIS, isis.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "l2c", f.L2Circuit, l2circuit.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ldp", f.LDP, ldp.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "nat", f.NAT, nat.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "ospf", f.OSPF, func() collector.RPCCollector {
		return ospf.NewCollector(c.logicalSystem)
	})
	c.addCollectorIfEnabledForDevice(device, "routes", f.Routes, route.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "rpki", f.RPKI, rpki.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "rpm", f.RPM, rpm.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "storage", f.Storage, storage.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "system", f.System, system.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "power", f.Power, power.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "mac", f.MAC, mac.NewCollector)
	c.addCollectorIfEnabledForDevice(device, "vrrp", f.VRRP, vrrp.NewCollector)
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
