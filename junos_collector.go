// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/czerwonk/junos_exporter/internal/config"
	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/dynamiclabels"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	log "github.com/sirupsen/logrus"
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
	collectors *collectors
	ctx        context.Context
}

func newJunosCollector(ctx context.Context, devices []*connector.Device, logicalSystem string) *junosCollector {
	clients := make(map[*connector.Device]*rpc.Client)

	for _, d := range devices {
		cl, err := clientForDevice(d, connManager)
		if err != nil {
			log.Errorf("Could not connect to %s: %s", d, err)
			continue
		}

		clients[d] = cl
	}

	return &junosCollector{
		devices:    devices,
		collectors: collectorsForDevices(devices, cfg, logicalSystem),
		clients:    clients,
		ctx:        ctx,
	}
}

func deviceInterfaceRegex(cfg *config.Config, host string) *regexp.Regexp {
	dc := cfg.FindDeviceConfig(host)

	if dc != nil {
		return dc.IfDescReg
	}

	if cfg.IfDescReg != nil {
		return cfg.IfDescReg
	}

	return dynamiclabels.DefaultInterfaceDescRegex()
}

func clientForDevice(device *connector.Device, connManager *connector.SSHConnectionManager) (*rpc.Client, error) {
	conn, err := connManager.Connect(device)
	if err != nil {
		return nil, err
	}

	opts := []rpc.ClientOption{}
	if *debug {
		opts = append(opts, rpc.WithDebug())
	}

	if cfg.Features.Satellite {
		opts = append(opts, rpc.WithSatellite())
	}

	if cfg.Features.License {
		opts = append(opts, rpc.WithLicenseInformation())
	}

	c := rpc.NewClient(conn, opts...)
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
	ctx, span := tracer.Start(c.ctx, "Collect")
	defer span.End()

	wg := &sync.WaitGroup{}

	wg.Add(len(c.devices))
	for _, d := range c.devices {
		go c.collectForHost(ctx, d, ch, wg)
	}

	wg.Wait()
}

func (c *junosCollector) collectForHost(ctx context.Context, device *connector.Device, ch chan<- prometheus.Metric, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, span := tracer.Start(ctx, "CollectForHost", trace.WithAttributes(
		attribute.String("host", device.Host),
	))
	defer span.End()

	l := []string{device.Host}

	t := time.Now()
	defer func() {
		ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, time.Since(t).Seconds(), l...)
	}()

	cl, found := c.clients[device]
	if !found {
		ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 0, l...)
		return
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, l...)

	for _, col := range c.collectors.collectorsForDevice(device) {
		ctx, sp := tracer.Start(ctx, "CollectForHostWithCollector", trace.WithAttributes(
			attribute.String("collector", col.Name()),
		))

		cta := &clientTracingAdapter{
			cl:  cl,
			ctx: ctx,
		}

		ct := time.Now()
		err := col.Collect(cta, ch, l)

		if err != nil && err.Error() != "EOF" {
			sp.RecordError(err)
			sp.SetStatus(codes.Error, err.Error())
			log.Errorln(col.Name() + ": " + err.Error())
		}

		ch <- prometheus.MustNewConstMetric(scrapeCollectorDurationDesc, prometheus.GaugeValue, time.Since(ct).Seconds(), append(l, col.Name())...)
		sp.End()
	}
}
