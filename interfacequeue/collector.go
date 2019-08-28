package interfacequeue

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/interfacelabels"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_queues_"

// NewCollector creates an queue collector instance
func NewCollector(labels *interfacelabels.DynamicLabels) collector.RPCCollector {
	c := &interfaceQueueCollector{
		labels: labels,
	}
	c.init()

	return c
}

type interfaceQueueCollector struct {
	labels               *interfacelabels.DynamicLabels
	queuedPackets        *prometheus.Desc
	queuedBytes          *prometheus.Desc
	transferedPackets    *prometheus.Desc
	transferedBytes      *prometheus.Desc
	rateLimitDropPackets *prometheus.Desc
	rateLimitDropBytes   *prometheus.Desc
	totalDropPackets     *prometheus.Desc
	totalDropBytes       *prometheus.Desc
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
	ch <- c.totalDropBytes
	ch <- c.totalDropPackets
}

// Collect collects metrics from JunOS
func (c *interfaceQueueCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	q := InterfaceQueueRPC{}

	err := client.RunCommandAndParse("show interfaces queue", &q)
	if err != nil {
		return err
	}

	for _, iface := range q.InterfaceInformation.Interfaces {
		c.collectForInterface(iface, client.Device(), ch, labelValues)
	}

	return nil
}

func (c *interfaceQueueCollector) collectForInterface(iface PhysicalInterface, device *connector.Device, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, iface.Name, iface.Description)
	l = append(l, c.labels.ValuesForInterface(device, iface.Name)...)

	for _, q := range iface.QueueCounters.Queues {
		c.collectForQueue(q, ch, l)
	}
}

func (c *interfaceQueueCollector) collectForQueue(queue Queue, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, queue.Number)

	ch <- prometheus.MustNewConstMetric(c.queuedPackets, prometheus.CounterValue, float64(queue.QueuedPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.queuedBytes, prometheus.CounterValue, float64(queue.QueuedBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.transferedPackets, prometheus.CounterValue, float64(queue.TransferedPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.transferedBytes, prometheus.CounterValue, float64(queue.TransferedBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.rateLimitDropPackets, prometheus.CounterValue, float64(queue.RateLimitDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.rateLimitDropBytes, prometheus.CounterValue, float64(queue.RateLimitDropBytes), l...)
	ch <- prometheus.MustNewConstMetric(c.totalDropBytes, prometheus.CounterValue, float64(queue.TotalDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(c.totalDropPackets, prometheus.CounterValue, float64(queue.TotalDropBytes), l...)
}
