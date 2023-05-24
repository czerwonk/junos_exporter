// SPDX-License-Identifier: MIT

package lacp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_lacp_"

var (
	lacpMuxState    *prometheus.Desc
	lacpMuxStateMap = map[string]int{
		"Detached":                1,
		"Waiting":                 2,
		"Attached":                3,
		"Collecting":              4,
		"Distributing":            5,
		"Collecting distributing": 6,
	}
)

func init() {
	l := []string{"target", "aggregate", "name"}
	lacpMuxState = prometheus.NewDesc(prefix+"muxstate", "lacp mux state (1: detached, 2: waiting, 3: attached, 4: collecting, 5: distributing, 6: collecting distribuging)", l, nil)
}

type lacpCollector struct {
}

// Name returns the name of the collector
func (*lacpCollector) Name() string {
	return "lacp"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &lacpCollector{}
}

// Describe describes the metrics
func (*lacpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- lacpMuxState
}

// Collect collects metrics from JunOS
func (c *lacpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}
	err := client.RunCommandAndParse("show lacp interfaces", &x)
	if err != nil {
		return err
	}

	for _, iface := range x.Information.LacpInterfaces {
		for _, member := range iface.LagLACPProtocols {
			l := append(labelValues, iface.LagLACPHeader.Name, member.Member)
			ch <- prometheus.MustNewConstMetric(lacpMuxState, prometheus.GaugeValue, float64(lacpMuxStateMap[member.LacpMuxState]), l...)
		}
	}

	return nil
}
