package main

import (
	"regexp"
	"sync"
	"time"

	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/interfacelabels"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const prefix = "junos_"

var (
	scrapeCollectorDurationDesc *prometheus.Desc
	scrapeDurationDesc          *prometheus.Desc
	upDesc                      *prometheus.Desc
	defaultIfDescReg            *regexp.Regexp
)

func init() {
	upDesc = prometheus.NewDesc(prefix+"up", "Scrape of target was successful", []string{"target"}, nil)
	scrapeDurationDesc = prometheus.NewDesc(prefix+"collector_duration_seconds", "Duration of a collector scrape for one target", []string{"target"}, nil)
	scrapeCollectorDurationDesc = prometheus.NewDesc(prefix+"collect_duration_seconds", "Duration of a scrape by collector and target", []string{"target", "collector"}, nil)
	defaultIfDescReg = regexp.MustCompile(`\[([^=\]]+)(=[^\]]+)?\]`)
}

type junosCollector struct {
	devices    []*connector.Device
	clients    map[*connector.Device]*rpc.Client
	collectors *collectors
}

func newJunosCollector(devices []*connector.Device, connectionManager *connector.SSHConnectionManager, logicalSystem string) *junosCollector {
	l := interfacelabels.NewDynamicLabels()

	clients := make(map[*connector.Device]*rpc.Client)

	for index, d := range devices {
		cl, err := clientForDevice(d, connManager)
		if err != nil {
			log.Errorf("Could not connect to %s: %s", d, err)
			continue
		}

		clients[d] = cl

		if *dynamicIfaceLabels {
			regex := defaultIfDescReg
			if cfg.Devices[index].IfDescReg != "" {
				regex = regexp.MustCompile(cfg.Devices[index].IfDescReg)
			} else if cfg.IfDescReg != "" {
				regex = regexp.MustCompile(cfg.IfDescReg)
			}

			err = l.CollectDescriptions(d, cl, regex)

			if err != nil {
				log.Errorf("Could not get interface descriptions %s: %s", d, err)
				continue
			}
		}
	}

	return &junosCollector{
		devices:    devices,
		collectors: collectorsForDevices(devices, cfg, logicalSystem, l),
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

	if cfg.Features.Satellite {
		c.EnableSatellite()
	}

	return c, nil
}

// Describe implements prometheus.Collector interface
func (c *junosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- scrapeDurationDesc
	ch <- scrapeCollectorDurationDesc

	for _, col := range c.collectors.allEnabledCollectors() {
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

	for _, col := range c.collectors.collectorsForDevice(device) {
		ct := time.Now()
		err := col.Collect(rpc, ch, l)

		if err != nil && err.Error() != "EOF" {
			log.Errorln(col.Name() + ": " + err.Error())
		}

		ch <- prometheus.MustNewConstMetric(scrapeCollectorDurationDesc, prometheus.GaugeValue, time.Since(ct).Seconds(), append(l, col.Name())...)
	}
}
