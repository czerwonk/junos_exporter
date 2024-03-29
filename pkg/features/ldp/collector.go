// SPDX-License-Identifier: MIT

package ldp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ldpNeighborDesc     *prometheus.Desc
	ldpSessionDesc      *prometheus.Desc
	ldpSessionCountDesc *prometheus.Desc
	ldpStateMap         = map[string]int{
		"Operational": 1,
		"Nonexistant": 0,
	}
)

func init() {
	ldprefix := "junos_ldp_"

	lSession := []string{"target", "neighbor"}
	l := []string{"target"}

	ldpNeighborDesc = prometheus.NewDesc(ldprefix+"neighbor_count", "Number of LDP Neighbors", l, nil)

	ldpSessionCountDesc = prometheus.NewDesc(ldprefix+"session_count", "Number of LDP Sessions", l, nil)

	ldpSessionDesc = prometheus.NewDesc(ldprefix+"session_state", "State of LDP Sessions", lSession, nil)
}

// Collector collects ldpv3 metrics
type ldpCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &ldpCollector{}
}

// Name returns the name of the collector
func (*ldpCollector) Name() string {
	return "LDP"
}

// Describe describes the metrics
func (*ldpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- ldpNeighborDesc
	ch <- ldpSessionCountDesc
	ch <- ldpSessionDesc
}

// Collect collects metrics from JunOS
func (c *ldpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collectLDPSessions(client, ch, labelValues)
	if err != nil {
		return err
	}

	return c.collectLDPMetrics(client, ch, labelValues)
}

func (c *ldpCollector) collectLDPMetrics(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}
	err := client.RunCommandAndParse("show ldp neighbor", &x)
	if err != nil {
		return err
	}

	neighbors := x.Information.Neighbors
	ch <- prometheus.MustNewConstMetric(ldpNeighborDesc, prometheus.GaugeValue, float64(len(neighbors)), labelValues...)

	return nil
}

func (c *ldpCollector) collectLDPSessions(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = sessionResult{}
	err := client.RunCommandAndParse("show ldp session", &x)
	if err != nil {
		return err
	}

	sessions := x.Information.Sessions
	sessionCount := len(sessions)

	for _, sess := range sessions {
		l := append(labelValues, sess.NeighborAddress)
		ch <- prometheus.MustNewConstMetric(ldpSessionDesc, prometheus.GaugeValue, float64(ldpStateMap[sess.State]), l...)
	}
	ch <- prometheus.MustNewConstMetric(ldpSessionCountDesc, prometheus.GaugeValue, float64(sessionCount), labelValues...)

	return nil
}
