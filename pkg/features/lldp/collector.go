// SPDX-License-Identifier: MIT

package lldp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_lldp_"

var (
	lldpPeer *prometheus.Desc
)

func init() {
	l := []string{"target", "local_interface", "parent_interface", "remote_port_info", "remote_system_name"}
	lldpPeer = prometheus.NewDesc(prefix+"peer", "LLDP peer information (1: connected)", l, nil)
}

type lldpCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &lldpCollector{}
}

// Name returns the name of the collector
func (*lldpCollector) Name() string {
	return "LLDP"
}

// Describe describes the metrics
func (*lldpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- lldpPeer
}

// Collect collects metrics from JunOS
func (c *lldpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}
	err := client.RunCommandAndParse("show lldp neighbors", &x)
	if err != nil {
		return err
	}

	for _, neighbor := range x.Information.Neighbors {
		l := labelValues
		l = append(l, neighbor.LocalPortID, neighbor.LocalParentInterfaceName, neighbor.RemotePortID, neighbor.RemoteSystemName)
		ch <- prometheus.MustNewConstMetric(lldpPeer, prometheus.GaugeValue, 1.0, l...)
	}

	return nil
} 
