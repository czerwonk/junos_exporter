package main

import (
	"sync"
	"time"

	"github.com/czerwonk/junos_exporter/accounting"
	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/collector"
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
	"github.com/czerwonk/junos_exporter/nat"
	"github.com/czerwonk/junos_exporter/ospf"
	"github.com/czerwonk/junos_exporter/route"
	"github.com/czerwonk/junos_exporter/routingengine"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/czerwonk/junos_exporter/rpki"
	"github.com/czerwonk/junos_exporter/storage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const prefix = "junos_"

var (
	scrapeCollectorDurationDesc *prometheus.Desc
	scrapeDurationDesc          *prometheus.Desc
	upDesc                      *prometheus.Desc
)

func init() {
	upDesc = prometheus.NewDesc(prefix+"up", "Scrape of target was successful", []string{"target"}, nil)
	scrapeDurationDesc = prometheus.NewDesc(prefix+"collector_duration_seconds", "Duration of a collector scrape for one target", []string{"target"}, nil)
	scrapeCollectorDurationDesc = prometheus.NewDesc(prefix+"collect_duration_seconds", "Duration of a scrape by collector and target", []string{"target", "collector"}, nil)
}

type junosCollector struct {
	devices    []*connector.Device
	clients    map[*connector.Device]*rpc.Client
	collectors map[string]collector.RPCCollector
}

func newJunosCollector(devices []*connector.Device, connectionManager *connector.SSHConnectionManager, logicalSystem string) *junosCollector {
	l := interfacelabels.NewDynamicLabels()

	clients := make(map[*connector.Device]*rpc.Client)

	for _, d := range devices {
		cl, err := clientForDevice(d, connManager)
		if err != nil {
			log.Errorf("Could not connect to %s: %s", d, err)
			continue
		}

		clients[d] = cl

		if *dynamicIfaceLabels {
			err = l.CollectDescriptions(d, cl)
			if err != nil {
				log.Errorf("Could not get interface descriptions %s: %s", d, err)
				continue
			}
		}
	}

	return &junosCollector{
		devices:    devices,
		collectors: collectors(l, logicalSystem),
		clients:    clients,
	}
}

func clientForDevice(device *connector.Device, connManager *connector.SSHConnectionManager) (*rpc.Client, error) {
	conn, err := connManager.Connect(device)
	if err != nil {
		return nil, err
	}

	c := rpc.NewClient(conn)

	if *debug {
		c.EnableDebug()
	}

	return c, nil
}

func collectors(ifaceLabels *interfacelabels.DynamicLabels, logicalSystem string) map[string]collector.RPCCollector {
	m := make(map[string]collector.RPCCollector)

	f := &cfg.Features

	if f.Alarm {
		m["alarm"] = alarm.NewCollector(*alarmFilter)
	}

	if f.Interfaces {
		m["interfaces"] = interfaces.NewCollector(ifaceLabels)
	}

	if f.Routes {
		m["routes"] = route.NewCollector()
	}

	if f.BGP {
		m["bgp"] = bgp.NewCollector(logicalSystem)
	}

	if f.OSPF {
		m["ospf"] = ospf.NewCollector(logicalSystem)
	}

	if f.ISIS {
		m["isis"] = isis.NewCollector()
	}

	if f.IPSec {
		m["ipsec"] = ipsec.NewCollector()
	}

	if f.LDP {
		m["ldp"] = ldp.NewCollector()
	}

	if f.L2Circuit {
		m["l2circuit"] = l2circuit.NewCollector()
	}

	if f.NAT {
		m["nat"] = nat.NewCollector()
	}

	if f.RoutingEngine {
		m["routing-engine"] = routingengine.NewCollector()
	}

	if f.Environment {
		m["environment"] = environment.NewCollector()
	}

	if f.Firewall {
		m["firewall"] = firewall.NewCollector()
	}

	if f.InterfaceDiagnostic {
		m["interface-diagnostics"] = interfacediagnostics.NewCollector(ifaceLabels)
	}

	if f.Storage {
		m["storage"] = storage.NewCollector()
	}

	if f.Accounting {
		m["accounting"] = accounting.NewCollector()
	}

	if f.FPC {
		m["fpc"] = fpc.NewCollector()
	}

	if f.InterfaceQueue {
		m["interface_queue"] = interfacequeue.NewCollector(ifaceLabels)
	}

	if f.RPKI {
		m["rpki"] = rpki.NewCollector()
	}

	return m
}

// Describe implements prometheus.Collector interface
func (c *junosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc
	ch <- scrapeCollectorDurationDesc

	for _, col := range c.collectors {
		col.Describe(ch)
	}
}

// Collect implements prometheus.Collector interface
func (c *junosCollector) Collect(ch chan<- prometheus.Metric) {
	wg := &sync.WaitGroup{}

	wg.Add(len(c.devices))
	for _, d := range c.devices {
		go c.collectForHost(d, ch, wg)
	}

	wg.Wait()
}

func (c *junosCollector) collectForHost(device *connector.Device, ch chan<- prometheus.Metric, wg *sync.WaitGroup) {
	defer wg.Done()

	l := []string{device.Host}

	t := time.Now()
	defer func() {
		ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(t).Seconds(), l...)
	}()

	rpc, found := c.clients[device]
	if !found {
		ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 0, l...)
		return
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, l...)

	for k, col := range c.collectors {
		ct := time.Now()
		err := col.Collect(rpc, ch, l)

		if err != nil && err.Error() != "EOF" {
			log.Errorln(k + ": " + err.Error())
		}

		ch <- prometheus.MustNewConstMetric(scrapeCollectorDurationDesc, prometheus.GaugeValue, time.Since(ct).Seconds(), append(l, k)...)
	}
}
