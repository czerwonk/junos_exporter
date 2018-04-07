package main

import (
	"strings"
	"time"

	"sync"

	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/environment"
	"github.com/czerwonk/junos_exporter/interfacediagnostics"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/isis"
	"github.com/czerwonk/junos_exporter/ospf"
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
	interfaceCollector            *interfaces.Collector
	alarmCollector                *alarm.Collector
	bgpCollector                  *bgp.Collector
	ospfCollector                 *ospf.Collector
	isisCollector                 *isis.Collector
	routeCollector                *route.Collector
	routingEngineCollector        *routingengine.Collector
	environmentCollector          *environment.Collector
	interfaceDiagnosticsCollector *interfacediagnostics.Collector
}

// Describe implements prometheus.Collector interface
func (c *junosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc

	c.interfaceCollector.Describe(ch)
	c.alarmCollector.Describe(ch)
	c.bgpCollector.Describe(ch)
	c.ospfCollector.Describe(ch)
	c.isisCollector.Describe(ch)
	c.routeCollector.Describe(ch)
	c.environmentCollector.Describe(ch)
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

	rpc := rpc.NewClient(conn, *debug, *alarmFilter)
	for k, c := range c.collectors(rpc, ch, l) {
		err = c()
		if err != nil && err.Error() != "EOF" {
			log.Errorln(k + ": " + err.Error())
		}
	}
}

func (c *junosCollector) collectors(rpc *rpc.RpcClient, ch chan<- prometheus.Metric, labelValues []string) map[string]func() error {
	m := map[string]func() error{
		"interface": func() error { return c.interfaceCollector.Collect(rpc, ch, labelValues) },
		"alarm":     func() error { return c.alarmCollector.Collect(rpc, ch, labelValues) },
	}

	if *routesEnabled {
		m["routes"] = func() error { return c.routeCollector.Collect(rpc, ch, labelValues) }
	}

	if *bgpEnabled {
		m["bgp"] = func() error { return c.bgpCollector.Collect(rpc, ch, labelValues) }
	}

	if *ospfEnabled {
		m["ospf"] = func() error { return c.ospfCollector.Collect(rpc, ch, labelValues) }
	}

	if *isisEnabled {
		m["isis"] = func() error { return c.isisCollector.Collect(rpc, ch, labelValues) }
	}

	if *routingEngineEnabled {
		m["routing-engine"] = func() error { return c.routingEngineCollector.Collect(rpc, ch, labelValues) }
	}

	if *environmentEnabled {
		m["environment"] = func() error { return c.environmentCollector.Collect(rpc, ch, labelValues) }
	}

	if *ifDiagnEnabled {
		m["interface_diagnostics"] = func() error { return c.interfaceDiagnosticsCollector.Collect(rpc, ch, labelValues) }
	}

	return m
}
