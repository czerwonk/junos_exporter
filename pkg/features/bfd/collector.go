// SPDX-License-Identifier: MIT

package bfd

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_bfd_"

var (
	bfdState    *prometheus.Desc
	bfdStateMap = map[string]int{
		"Down": 0,
		"Up":   1,
	}
)

func init() {
	l := []string{"target", "neighbor", "interface", "client"}
	bfdState = prometheus.NewDesc(prefix+"state", "bfd state (0: down, 1:up)", l, nil)
}

type bfdCollector struct {
}

// Name returns the name of the collector
func (*bfdCollector) Name() string {
	return "bfd"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &bfdCollector{}
}

// Describe describes the metrics
func (*bfdCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- bfdState
}

// Collect collects metrics from JunOS
func (c *bfdCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var res = result{}
	err := client.RunCommandAndParse("show bfd session extensive", &res)
	if err != nil {
		return err
	}

	for _, bfds := range res.Information.BfdSessions {
		l := append(labelValues, bfds.Neighbor, bfds.Interface, bfds.Client.Name)
		ch <- prometheus.MustNewConstMetric(bfdState, prometheus.GaugeValue, float64(bfdStateMap[bfds.State]), l...)
	}

	return nil
}
