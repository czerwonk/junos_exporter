package main

import (
	"strings"
	"time"

	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/prometheus/client_golang/prometheus"
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

type JunosCollector struct {
	interfaceCollector *interfaces.InterfaceCollector
	alarmCollector     *alarm.AlarmCollector
	bgpCollector       *bgp.BgpCollector
}

func (c *JunosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc

	c.interfaceCollector.Describe(ch)
	c.alarmCollector.Describe(ch)
	c.bgpCollector.Describe(ch)
}

func (c *JunosCollector) Collect(ch chan<- prometheus.Metric) {
	for _, h := range strings.Split(*sshHosts, ",") {
		go c.collectForHost(strings.Trim(h, " "), ch)
	}
}

func (c *JunosCollector) collectForHost(host string, ch chan<- prometheus.Metric) {
	l := []string{host}

	t := time.Now()
	defer func() {
		ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(t).Seconds(), l...)
	}()

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, l...)

	//c.interfaceCollector.Collect()
	//c.alarmCollector.Collect()
	//c.bgpCollector.Collect()
}
