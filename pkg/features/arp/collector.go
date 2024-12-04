package arp

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

const prefix string = "junos_arp_"

var (
	arpEntriesCountDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "interface"}
	arpEntriesCountDesc = prometheus.NewDesc(prefix+"arp_entries_count_per_interface", "Amount of ARP entries on an interface", l, nil)
}

type arpCollector struct{}

func NewCollector() collector.RPCCollector {
	return &arpCollector{}
}

func (c *arpCollector) Name() string {
	return "arp"
}

func (c *arpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- arpEntriesCountDesc
}

func (c *arpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var arps results
	err := client.RunCommandAndParse("show arp no-resolve", &arps)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show arp no-resolve'")
	}
	interfaces_map := make(map[string]float64)
	for _, a := range arps.ArpTableInformation.ArpTableEntry {
		interfaces_map[a.InterfaceName] += 1
	}
	for key, value := range interfaces_map {
		labels := append(labelValues, key)
		ch <- prometheus.MustNewConstMetric(arpEntriesCountDesc, prometheus.CounterValue, value, labels...)
	}
	return nil
}
