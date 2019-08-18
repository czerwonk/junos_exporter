package firewall

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_firewall_filter_"

var (
	counterPackets       *prometheus.Desc
	counterBytes         *prometheus.Desc
	policerPackets       *prometheus.Desc
	policerBytes         *prometheus.Desc
)

func init() {
	l := []string{"target", "filter", "counter"}

	counterPackets = prometheus.NewDesc(prefix+"counter_packets", "Number of packets matching counter in firewall filter", l, nil)
	counterBytes = prometheus.NewDesc(prefix+"counter_bytes", "Number of bytes matching counter in firewall filter", l, nil)
  policerPackets = prometheus.NewDesc(prefix+"policer_packets", "Number of packets matching policer in firewall filter", l, nil)
	policerBytes = prometheus.NewDesc(prefix+"policer_bytes", "Number of bytes matching policer in firewall filter", l, nil)
}

type firewallCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &firewallCollector{}
}

// Describe describes the metrics
func (*firewallCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- counterPackets
	ch <- counterBytes
	ch <- policerPackets
	ch <- policerBytes
}

// Collect collects metrics from JunOS
func (c *firewallCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = FirewallRpc{}
	err := client.RunCommandAndParse("show firewall filter regex .*", &x)
	if err != nil {
		return err
	}

	for _, t := range x.Information.Filters {
		c.collectForFilter(t, ch, labelValues)
	}

	return nil
}

func (c *firewallCollector) collectForFilter(filter Filter, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, filter.Name)

	for _, counter := range filter.Counters {
		lp := append(l, counter.Name)
		ch <- prometheus.MustNewConstMetric(counterPackets, prometheus.GaugeValue, float64(counter.Packets), lp...)
		ch <- prometheus.MustNewConstMetric(counterBytes, prometheus.GaugeValue, float64(counter.Bytes), lp...)
	}

	for _, policer := range filter.Policers {
		lp := append(l, policer.Name)
		ch <- prometheus.MustNewConstMetric(policerPackets, prometheus.GaugeValue, float64(policer.Packets), lp...)
		ch <- prometheus.MustNewConstMetric(policerBytes, prometheus.GaugeValue, float64(policer.Bytes), lp...)
	}
}
