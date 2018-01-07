package main

import (
	"strings"
	"time"

	"sync"

	"github.com/czerwonk/junos_exporter/alarm"
	"github.com/czerwonk/junos_exporter/bgp"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/interfaces"
	"github.com/czerwonk/junos_exporter/ospf"
	"github.com/czerwonk/junos_exporter/route"
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

type JunosCollector struct {
	interfaceCollector *interfaces.InterfaceCollector
	alarmCollector     *alarm.AlarmCollector
	bgpCollector       *bgp.BgpCollector
	ospfCollector      *ospf.OspfCollector
	routeCollector     *route.RouteCollector
}

func (c *JunosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc

	c.interfaceCollector.Describe(ch)
	c.alarmCollector.Describe(ch)
	c.bgpCollector.Describe(ch)
	c.ospfCollector.Describe(ch)
	c.routeCollector.Describe(ch)
}

func (c *JunosCollector) Collect(ch chan<- prometheus.Metric) {
	hosts := strings.Split(*sshHosts, ",")
	wg := &sync.WaitGroup{}

	wg.Add(len(hosts))
	for _, h := range hosts {
		go c.collectForHost(strings.Trim(h, " "), ch, wg)
	}

	wg.Wait()
}

func (c *JunosCollector) collectForHost(host string, ch chan<- prometheus.Metric, wg *sync.WaitGroup) {
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
	collectors := map[string]func() error{
		"interface": func() error { return c.interfaceCollector.Collect(rpc, ch, l) },
		"alarm":     func() error { return c.alarmCollector.Collect(rpc, ch, l) },
		"route":     func() error { return c.routeCollector.Collect(rpc, ch, l) },
		"bgp":       func() error { return c.bgpCollector.Collect(rpc, ch, l) },
		"ospf":      func() error { return c.ospfCollector.Collect(rpc, ch, l) },
	}

	for k, c := range collectors {
		err = c()
		if err != nil && err.Error() != "EOF" {
			log.Errorln(k + ": " + err.Error())
		}
	}
}
