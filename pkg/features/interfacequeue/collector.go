// SPDX-License-Identifier: MIT

package interfacequeue

import (
	"regexp"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/dynamiclabels"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_queues_"

type description struct {
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

func newDescriptions(dynLabels dynamiclabels.Labels) *description {
	d := &description{}

	l := []string{"target", "name", "description"}
	l = append(l, "queue_number")
	l = append(l, "forwarding_class")
	l = append(l, dynLabels.Keys()...)

	d.queuedPackets = prometheus.NewDesc(prefix+"queued_packets_count", "Number of queued packets", l, nil)
	d.queuedBytes = prometheus.NewDesc(prefix+"queued_bytes_count", "Number of bytes of queued packets", l, nil)
	d.transferedPackets = prometheus.NewDesc(prefix+"transfered_packets_count", "Number of transfered packets", l, nil)
	d.transferedBytes = prometheus.NewDesc(prefix+"transfered_bytes_count", "Number of bytes of transfered packets", l, nil)
	d.rateLimitDropPackets = prometheus.NewDesc(prefix+"rate_limit_drop_packets_count", "Number of packets droped by rate limit", l, nil)
	d.rateLimitDropBytes = prometheus.NewDesc(prefix+"rate_limit_drop_bytes_count", "Number of bytes droped by rate limit", l, nil)
	d.redPackets = prometheus.NewDesc(prefix+"red_packets_count", "Number of queued packets", l, nil)
	d.redBytes = prometheus.NewDesc(prefix+"red_bytes_count", "Number of bytes of queued packets", l, nil)
	d.redPacketsLow = prometheus.NewDesc(prefix+"red_packets_low_count", "Number of queued packets", l, nil)
	d.redBytesLow = prometheus.NewDesc(prefix+"red_bytes_low_count", "Number of bytes of queued packets", l, nil)
	d.redPacketsMediumLow = prometheus.NewDesc(prefix+"red_packets_medium_low_count", "Number of queued packets", l, nil)
	d.redBytesMediumLow = prometheus.NewDesc(prefix+"red_bytes_medium_low_count", "Number of bytes of queued packets", l, nil)
	d.redPacketsMediumHigh = prometheus.NewDesc(prefix+"red_packets_medium_high_count", "Number of queued packets", l, nil)
	d.redBytesMediumHigh = prometheus.NewDesc(prefix+"red_bytes_medium_high_count", "Number of bytes of queued packets", l, nil)
	d.redPacketsHigh = prometheus.NewDesc(prefix+"red_packets_high_count", "Number of queued packets", l, nil)
	d.redBytesHigh = prometheus.NewDesc(prefix+"red_bytes_high_count", "Number of bytes of queued packets", l, nil)
	d.tailDropPackets = prometheus.NewDesc(prefix+"tail_drop_packets_count", "Number of tail droped packets", l, nil)
	d.totalDropPackets = prometheus.NewDesc(prefix+"drop_packets_count", "Number of packets droped", l, nil)
	d.totalDropBytes = prometheus.NewDesc(prefix+"drop_bytes_count", "Number of bytes droped", l, nil)

	return d
}

// NewCollector creates an queue collector instance
func NewCollector(descRe *regexp.Regexp) collector.RPCCollector {
	c := &interfaceQueueCollector{
		descriptionRe: descRe,
	}

	return c
}

type interfaceQueueCollector struct {
	descriptionRe *regexp.Regexp
}

// Name returns the name of the collector
func (*interfaceQueueCollector) Name() string {
	return "Interface Queues"
}

// Describe describes the metrics
func (c *interfaceQueueCollector) Describe(ch chan<- *prometheus.Desc) {
	d := newDescriptions(nil)
	ch <- d.queuedBytes
	ch <- d.queuedPackets
	ch <- d.transferedBytes
	ch <- d.transferedPackets
	ch <- d.rateLimitDropBytes
	ch <- d.rateLimitDropPackets
	ch <- d.redPackets
	ch <- d.redBytes
	ch <- d.redPacketsLow
	ch <- d.redBytesLow
	ch <- d.redPacketsMediumLow
	ch <- d.redBytesMediumLow
	ch <- d.redPacketsMediumHigh
	ch <- d.redBytesMediumHigh
	ch <- d.redPacketsHigh
	ch <- d.redBytesHigh
	ch <- d.tailDropPackets
	ch <- d.totalDropBytes
	ch <- d.totalDropPackets
}

// Collect collects metrics from JunOS
func (c *interfaceQueueCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	q := result{}

	err := client.RunCommandAndParse("show interfaces queue", &q)
	if err != nil {
		return err
	}

	for _, iface := range q.InterfaceInformation.Interfaces {
		c.collectForInterface(iface, ch, labelValues)
	}

	return nil
}

func (c *interfaceQueueCollector) collectForInterface(iface physicalInterface, ch chan<- prometheus.Metric, labelValues []string) {
	lv := append(labelValues, []string{iface.Name, iface.Description}...)
	dynLabels := dynamiclabels.ParseDescription(iface.Description, c.descriptionRe)

	for _, q := range iface.QueueCounters.Queues {
		c.collectForQueue(q, ch, lv, dynLabels)
	}
}

func (c *interfaceQueueCollector) collectForQueue(queue queue, ch chan<- prometheus.Metric, labelValues []string, dynLabels dynamiclabels.Labels) {
	l := append(labelValues, queue.Number)
	l = append(l, queue.ForwaringClassName)
	l = append(l, dynLabels.Values()...)

	d := newDescriptions(dynLabels)
	ch <- prometheus.MustNewConstMetric(d.queuedPackets, prometheus.CounterValue, float64(queue.QueuedPackets), l...)
	ch <- prometheus.MustNewConstMetric(d.queuedBytes, prometheus.CounterValue, float64(queue.QueuedBytes), l...)
	ch <- prometheus.MustNewConstMetric(d.transferedPackets, prometheus.CounterValue, float64(queue.TransferedPackets), l...)
	ch <- prometheus.MustNewConstMetric(d.transferedBytes, prometheus.CounterValue, float64(queue.TransferedBytes), l...)
	ch <- prometheus.MustNewConstMetric(d.rateLimitDropPackets, prometheus.CounterValue, float64(queue.RateLimitDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(d.rateLimitDropBytes, prometheus.CounterValue, float64(queue.RateLimitDropBytes), l...)
	ch <- prometheus.MustNewConstMetric(d.redPackets, prometheus.CounterValue, float64(queue.RedPackets), l...)
	ch <- prometheus.MustNewConstMetric(d.redBytes, prometheus.CounterValue, float64(queue.RedBytes), l...)
	ch <- prometheus.MustNewConstMetric(d.redPacketsLow, prometheus.CounterValue, float64(queue.RedPacketsLow), l...)
	ch <- prometheus.MustNewConstMetric(d.redBytesLow, prometheus.CounterValue, float64(queue.RedBytesLow), l...)
	ch <- prometheus.MustNewConstMetric(d.redPacketsMediumLow, prometheus.CounterValue, float64(queue.RedPacketsMediumLow), l...)
	ch <- prometheus.MustNewConstMetric(d.redBytesMediumLow, prometheus.CounterValue, float64(queue.RedBytesMediumLow), l...)
	ch <- prometheus.MustNewConstMetric(d.redPacketsMediumHigh, prometheus.CounterValue, float64(queue.RedPacketsMediumHigh), l...)
	ch <- prometheus.MustNewConstMetric(d.redBytesMediumHigh, prometheus.CounterValue, float64(queue.RedBytesMediumHigh), l...)
	ch <- prometheus.MustNewConstMetric(d.redPacketsHigh, prometheus.CounterValue, float64(queue.RedPacketsHigh), l...)
	ch <- prometheus.MustNewConstMetric(d.redBytesHigh, prometheus.CounterValue, float64(queue.RedBytesHigh), l...)
	ch <- prometheus.MustNewConstMetric(d.tailDropPackets, prometheus.CounterValue, float64(queue.TailDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(d.totalDropPackets, prometheus.CounterValue, float64(queue.TotalDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(d.totalDropBytes, prometheus.CounterValue, float64(queue.TotalDropBytes), l...)
}
