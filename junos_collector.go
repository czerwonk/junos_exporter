package main

import (
	"strings"
	"sync"
	"time"

	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/environment"
	"github.com/czerwonk/junos_exporter/interfacediagnostics"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/isis"
	"github.com/czerwonk/junos_exporter/l2circuit"
	"github.com/czerwonk/junos_exporter/ospf"
	"github.com/czerwonk/junos_exporter/route"
	"github.com/czerwonk/junos_exporter/routingengine"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/czerwonk/junos_exporter/storage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const prefix = "junos_"

var (
	scrapeDurationDesc *prometheus.Desc
	upDesc             *prometheus.Desc
)

func init() {
	upDesc = prometheus.NewDesc(prefix+"up", "Scrape of target was successful", []string{"target"}, nil)
	scrapeDurationDesc = prometheus.NewDesc(prefix+"collector_duration_seconds", "Duration of a collector scrape for one target", []string{"target"}, nil)
}

type junosCollector struct {
	targets           []string
	collectors        map[string]collector.RPCCollector
	connectionManager *connector.SSHConnectionManager
}

func newJunosCollector(targets []string, connectionManager *connector.SSHConnectionManager) *junosCollector {
	collectors := collectors()
	return &junosCollector{targets, collectors, connectionManager}
}

func collectors() map[string]collector.RPCCollector {
	m := map[string]collector.RPCCollector{
		"alarm": alarm.NewCollector(*alarmFilter),
	}

	f := &cfg.Features

	if f.Interfaces {
		m["interfaces"] = interfaces.NewCollector()
	}

	if f.Routes {
		m["routes"] = route.NewCollector()
	}

	if f.BGP {
		m["bgp"] = bgp.NewCollector()
	}

	if f.OSPF {
		m["ospf"] = ospf.NewCollector()
	}

	if f.ISIS {
		m["isis"] = isis.NewCollector()
	}

	if f.L2Circuit {
		m["l2circuit"] = l2circuit.NewCollector()
	}

	if f.RoutingEngine {
		m["routing-engine"] = routingengine.NewCollector()
	}

	if f.Environment {
		m["environment"] = environment.NewCollector()
	}

	if f.InterfaceDiagnostic {
		m["interface-diagnostics"] = interfacediagnostics.NewCollector()
	}

	if f.Storage {
		m["storage"] = storage.NewCollector()
	}

	return m
}

// Describe implements prometheus.Collector interface
func (c *junosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc

	for _, col := range c.collectors {
		col.Describe(ch)
	}
}

// Collect implements prometheus.Collector interface
func (c *junosCollector) Collect(ch chan<- prometheus.Metric) {
	hosts := c.targets
	wg := &sync.WaitGroup{}

	wg.Add(len(hosts))
	for _, h := range hosts {
		go c.collectForHost(strings.Trim(h, " "), ch, wg)
	}

	wg.Wait()
}

func (c *junosCollector) collectForHost(host string, ch chan<- prometheus.Metric, wg *sync.WaitGroup) {
	defer wg.Done()

	l := []string{host}

	t := time.Now()
	defer func() {
		ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(t).Seconds(), l...)
	}()

	conn, err := c.connectionManager.Connect(host)
	if err != nil {
		log.Errorf("Could not connect to %s: %v", host, err)
		ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 0, l...)
		return
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, l...)

	rpc := rpc.NewClient(conn, *debug)
	for k, col := range c.collectors {
		err = col.Collect(rpc, ch, l)
		if err != nil && err.Error() != "EOF" {
			log.Errorln(k + ": " + err.Error())
		}
	}
}
