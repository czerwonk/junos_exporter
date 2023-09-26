// SPDX-License-Identifier: MIT

package subscriber

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_subscriber_info"

var subscriberInfoDesc *prometheus.Desc

func init() {
	l := []string{"target", "interface", "agent_circuit_id", "agent_remote_id"}
	subscriberInfoDesc = prometheus.NewDesc(prefix+"", "Subscriber Detail", l, nil)
}

// Name implements collector.RPCCollector.
func (*subcsribers_information) Name() string {
	return "Subscriber Detail"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &subcsribers_information{}
}

// Describe describes the metrics
func (*subcsribers_information) Describe(ch chan<- *prometheus.Desc) {
	ch <- subscriberInfoDesc
}

// Collect collects metrics from JunOS
func (c *subcsribers_information) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = subcsribers_information{}
	err := client.RunCommandAndParse("show subscribers client-type dhcp detail", &x)
	if err != nil {
		return err
	}

	for _, subscriber := range x.SubscribersInformation.Subscriber {
		labels := append(labelValues, subscriber.Interface, subscriber.AgentCircuitID, subscriber.AgentRemoteID)
		ch <- prometheus.MustNewConstMetric(subscriberInfoDesc, prometheus.CounterValue, 1, labels...)
	}

	return nil
}
