// SPDX-License-Identifier: MIT

package interfacequeue

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/interfacelabels"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_queues_"

// NewCollector creates an queue collector instance
func NewCollector(labels *interfacelabels.DynamicLabelManager) collector.RPCCollector {
	c := &interfaceQueueCollector{
		labels: labels,
	}
	c.init()

	return c
}

type interfaceQueueCollector struct {
	labels               *interfacelabels.DynamicLabelManager
	queuedPackets        *prometheus.Desc
	queuedBytes          *prometheus.Desc
	transferedPackets    *prometheus.Desc
	transferedBytes      *prometheus.Desc
	rateLimitDropPackets *prometheus.Desc
	rateLimitDropBytes   *prometheus.Desc
	redPackets           *prometheus.Desc
	redBytes             *prometheus.Desc
	redPacketsLow        *prometheus.Desc
	redBytesLow          *prometheus.Desc
	redPacketsMediumLow  *prometheus.Desc
	redBytesMediumLow    *prometheus.Desc
	redPacketsMediumHigh *prometheus.Desc
	redBytesMediumHigh   *prometheus.Desc
	redPacketsHigh       *prometheus.Desc
	redBytesHigh         *prometheus.Desc
	tailDropPackets      *prometheus.Desc
	totalDropPackets     *prometheus.Desc
	totalDropBytes       *prometheus.Desc
}

// Name returns the name of the collector
func (*interfaceQueueCollector) Name() string {
	return "Interface Queues"
}

func (c *interfaceQueueCollector) init() {
	l := []string{"target", "name", "description"}
	l = append(l, c.labels.LabelNames()...)
	l = append(l, "queue_number")

	c.queuedPackets = prometheus.NewDesc(prefix+"queued_packets_count", "Number of queued packets", l, nil)
	c.queuedBytes = prometheus.NewDesc(prefix+"queued_bytes_count", "Number of bytes of queued packets", l, nil)
	c.transferedPackets = prometheus.NewDesc(prefix+"transfered_packets_count", "Number of transfered packets", l, nil)
	c.transferedBytes = prometheus.NewDesc(prefix+"transfered_bytes_count", "Number of bytes of transfered packets", l, nil)
	c.rateLimitDropPackets = prometheus.NewDesc(prefix+"rate_limit_drop_packets_count", "Number of packets droped by rate limit", l, nil)
	c.rateLimitDropBytes = prometheus.NewDesc(prefix+"rate_limit_drop_bytes_count", "Number of bytes droped by rate limit", l, nil)
	c.redPackets = prometheus.NewDesc(prefix+"red_packets_count", "Number of queued packets", l, nil)
	c.redBytes = prometheus.NewDesc(prefix+"red_bytes_count", "Number of bytes of queued packets", l, nil)
	c.redPacketsLow = prometheus.NewDesc(prefix+"red_packets_low_count", "Number of queued packets", l, nil)
	c.redBytesLow = prometheus.NewDesc(prefix+"red_bytes_low_count", "Number of bytes of queued packets", l, nil)
	c.redPacketsMediumLow = prometheus.NewDesc(prefix+"red_packets_medium_low_count", "Number of queued packets", l, nil)
	c.redBytesMediumLow = prometheus.NewDesc(prefix+"red_bytes_medium_low_count", "Number of bytes of queued packets", l, nil)
	c.redPacketsMediumHigh = prometheus.NewDesc(prefix+"red_packets_medium_high_count", "Number of queued packets", l, nil)
	c.redBytesMediumHigh = prometheus.NewDesc(prefix+"red_bytes_medium_high_count", "Number of bytes of queued packets", l, nil)
	c.redPacketsHigh = prometheus.NewDesc(prefix+"red_packets_high_count", "Number of queued packets", l, nil)
	c.redBytesHigh = prometheus.NewDesc(prefix+"red_bytes_high_count", "Number of bytes of queued packets", l, nil)
	c.tailDropPackets = prometheus.NewDesc(prefix+"tail_drop_packets_count", "Number of tail droped packets", l, nil)
	c.totalDropPackets = prometheus.NewDesc(prefix+"drop_packets_count", "Number of packets droped", l, nil)
	c.totalDropBytes = prometheus.NewDesc(prefix+"drop_bytes_count", "Number of bytes droped", l, nil)
}

// Describe describes the metrics
func (c *interfaceQueueCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.queuedBytes
	ch <- c.queuedPackets
	ch <- c.transferedBytes
	ch <- c.transferedPackets
	ch <- c.rateLimitDropBytes
	ch <- c.rateLimitDropPackets
	ch <- c.redPackets
	ch <- c.redBytes
	ch <- c.redPacketsLow
	ch <- c.redBytesLow
	ch <- c.redPacketsMediumLow
	ch <- c.redBytesMediumLow
	ch <- c.redPacketsMediumHigh
	ch <- c.redBytesMediumHigh
	ch <- c.redPacketsHigh
	ch <- c.redBytesHigh
	ch <- c.tailDropPackets
	ch <- c.totalDropBytes
	ch <- c.totalDropPackets
}

// Collect collects metrics from JunOS
func (c *interfaceQueueCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	q := result{}

	err := client.RunCommandAndParse("show interfaces queue", &q)
	if err != nil {
		return err
	}

	for _, iface := range q.InterfaceInformation.Interfaces {
		c.collectForInterface(iface, client.Device(), ch, labelValues)
	}

	return nil
}

func (c *interfaceQueueCollector) collectForInterface(iface physicalInterface, device *connector.Device, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, iface.Name, iface.Description)
	l = append(l, c.labels.ValuesForInterface(device, iface.Name)...)

	for _, q := range iface.QueueCounters.Queues {
		c.collectForQueue(q, ch, l)
	}
}

func (c *interfaceQueueCollector) collectForQueue(queue queue, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, queue.Number)

	ch <- prometheus.MustNewConstMetric(c.queuedPackets, prometheus.CounterValue, float64(queue.QueuedPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.queuedBytes, prometheus.CounterValue, float64(queue.QueuedBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.transferedPackets, prometheus.CounterValue, float64(queue.TransferedPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.transferedBytes, prometheus.CounterValue, float64(queue.TransferedBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.rateLimitDropPackets, prometheus.CounterValue, float64(queue.RateLimitDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.rateLimitDropBytes, prometheus.CounterValue, float64(queue.RateLimitDropBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.redPackets, prometheus.CounterValue, float64(queue.RedPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.redBytes, prometheus.CounterValue, float64(queue.RedBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.redPacketsLow, prometheus.CounterValue, float64(queue.RedPacketsLow), l...)
	ch <- prometheus.MustNewConstMetric(c.redBytesLow, prometheus.CounterValue, float64(queue.RedBytesLow), l...)
	ch <- prometheus.MustNewConstMetric(c.redPacketsMediumLow, prometheus.CounterValue, float64(queue.RedPacketsMediumLow), l...)
	ch <- prometheus.MustNewConstMetric(c.redBytesMediumLow, prometheus.CounterValue, float64(queue.RedBytesMediumLow), l...)
	ch <- prometheus.MustNewConstMetric(c.redPacketsMediumHigh, prometheus.CounterValue, float64(queue.RedPacketsMediumHigh), l...)
	ch <- prometheus.MustNewConstMetric(c.redBytesMediumHigh, prometheus.CounterValue, float64(queue.RedBytesMediumHigh), l...)
	ch <- prometheus.MustNewConstMetric(c.redPacketsHigh, prometheus.CounterValue, float64(queue.RedPacketsHigh), l...)
	ch <- prometheus.MustNewConstMetric(c.redBytesHigh, prometheus.CounterValue, float64(queue.RedBytesHigh), l...)
	ch <- prometheus.MustNewConstMetric(c.tailDropPackets, prometheus.CounterValue, float64(queue.TailDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.totalDropPackets, prometheus.CounterValue, float64(queue.TotalDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.totalDropBytes, prometheus.CounterValue, float64(queue.TotalDropBytes), l...)
}
