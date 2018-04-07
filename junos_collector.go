package main

import (
	"strings"
	"time"

	"sync"

	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/ospfv3"
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/environment"
	"github.com/czerwonk/junos_exporter/interfacediagnostics"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/isis"
	"github.com/czerwonk/junos_exporter/route"
	"github.com/czerwonk/junos_exporter/routingengine"
	"github.com/czerwonk/junos_exporter/rpc"
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
	collectors map[string]collector.RPCCollector
}

func newJunosCollector() *junosCollector {
	collectors := collectors()
	return &junosCollector{collectors}
}

func collectors() map[string]collector.RPCCollector {
	m := map[string]collector.RPCCollector{
		"interface": interfaces.NewCollector(),
		"alarm":     alarm.NewCollector(*alarmFilter),
	}

	if *routesEnabled {
		m["routes"] = route.NewCollector()
	}

	if *bgpEnabled {
		m["bgp"] = bgp.NewCollector()
	}

	if *ospfEnabled {
		m["ospf"] = ospfv3.NewCollector()
	}

	if *isisEnabled {
		m["isis"] = isis.NewCollector()
	}

	if *routingEngineEnabled {
		m["routing-engine"] = routingengine.NewCollector()
	}

	if *environmentEnabled {
		m["environment"] = environment.NewCollector()
	}

	if *ifDiagnEnabled {
		m["interface_diagnostics"] = interfacediagnostics.NewCollector()
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
	hosts := strings.Split(*sshHosts, ",")
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

	conn, err := connector.NewSshConnection(host, *sshUsername, *sshKeyFile)
	if err != nil {
		log.Errorln(err)
		ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 0, l...)
		return
	}
	defer conn.Close()

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, l...)

	rpc := rpc.NewClient(conn, *debug)
	for k, col := range c.collectors {
		err = col.Collect(rpc, ch, l)
		if err != nil && err.Error() != "EOF" {
			log.Errorln(k + ": " + err.Error())
		}
	}
}
