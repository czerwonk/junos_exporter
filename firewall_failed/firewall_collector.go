package firewall

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_firewall_filter_"

var (
	packetsDesc          *prometheus.Desc
	bytesDesc            *prometheus.Desc
)

func init() {
	l := []string{"target", "name", "type", "description"}
	packetsDesc = prometheus.NewDesc(prefix+"packets", "Number of packets counted in firewall filter term", l, nil)
	bytesDesc = prometheus.NewDesc(prefix+"packets", "Number of bytes counted in firewall filter term", l, nil)
}

type firewallCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &firewallCollector{}
}

// Describe describes the metrics
func (*firewallCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- packetsDesc
	ch <- bytesDesc
}

// Collect collects metrics from JunOS
func (c *firewallCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	filters, err := c.filterTypes(client)
	if err != nil {
		return err
	}

	for _, s := range filters {
		c.collectForFilter(s, ch, labelValues)
	}

	return nil
}

func (c *firewallCollector) filterTypes(client *rpc.Client) ([]*FilterType, error) {
	var x = FirewallRpc{}
	err := client.RunCommandAndParse("show firewall filter regex .*", &x)
	if err != nil {
		return nil, err
	}

	filters := make([]*FilterType, 0)
	for _, counter := range x.Filter.Counters {
		s := &FilterType{
			Packets:      int64(counter.Packets),
			Bytes:        int64(counter.Bytes),
		}

		counters = append(counters, s)
	}
	// policers := make([]*FilterType, 0)
	// for _, policer := range x.FirewallInformation.Policers {
	// 	s := &FirewallInformation{
	// 		Packets:      int64(policer.packets),
	// 		Bytes:        int64(policer.bytes),
	// 	}
	//
	// 	policers = append(policers, s)
	// }

	return counters, nil
}

func (*firewallCollector) collectForFilter(s *FilterType, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Name, s.Type}...)

	up := 0
	if s.Up {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(bytesDesc, prometheus.GaugeValue, float64(s.Bytes), l...)
	ch <- prometheus.MustNewConstMetric(packetsDesc, prometheus.GaugeValue, float64(s.Packets), l...)
}
