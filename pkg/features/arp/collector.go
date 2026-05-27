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
	arpEntriesCountDesc = prometheus.NewDesc(prefix+"entries", "Amount of ARP entries on an interface", l, nil)
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
	var res results
	err := client.RunCommandAndParse("show arp no-resolve", &res)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show arp no-resolve'")
	}

	interfaces := make(map[string]float64)
	for _, a := range res.ArpTableInformation.ArpTableEntry {
		interfaces[a.InterfaceName] += 1
	}

	for key, value := range interfaces {
		labels := append(labelValues, key)
		ch <- prometheus.MustNewConstMetric(arpEntriesCountDesc, prometheus.GaugeValue, value, labels...)
	}

	return nil
}
