package interfacequeue

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_queues_"

var (
	queuedPackets        *prometheus.Desc
	queuedBytes          *prometheus.Desc
	transferedPackets    *prometheus.Desc
	transferedBytes      *prometheus.Desc
	rateLimitDropPackets *prometheus.Desc
	rateLimitDropBytes   *prometheus.Desc
	totalDropPackets     *prometheus.Desc
	totalDropBytes       *prometheus.Desc
)

func init() {
	l := []string{"target", "name", "description", "queue_number"}
	queuedPackets = prometheus.NewDesc(prefix+"queued_packets_count", "Number of queued packets", l, nil)
	queuedBytes = prometheus.NewDesc(prefix+"queued_bytes_count", "Number of bytes of queued packets", l, nil)
	transferedPackets = prometheus.NewDesc(prefix+"transfered_packets_count", "Number of transfered packets", l, nil)
	transferedBytes = prometheus.NewDesc(prefix+"transfered_bytes_count", "Number of bytes of transfered packets", l, nil)
	rateLimitDropPackets = prometheus.NewDesc(prefix+"rate_limit_drop_packets_count", "Number of packets droped by rate limit", l, nil)
	rateLimitDropBytes = prometheus.NewDesc(prefix+"rate_limit_drop_bytes_count", "Number of bytes droped by rate limit", l, nil)
	totalDropPackets = prometheus.NewDesc(prefix+"drop_packets_count", "Number of packets droped", l, nil)
	totalDropBytes = prometheus.NewDesc(prefix+"drop_bytes_count", "Number of bytes droped", l, nil)
}

// NewCollector creates an queue collector instance
func NewCollector() collector.RPCCollector {
	return &interfaceQueueCollector{}
}

type interfaceQueueCollector struct {
}

// Describe describes the metrics
func (*interfaceQueueCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- queuedBytes
	ch <- queuedPackets
	ch <- transferedBytes
	ch <- transferedPackets
	ch <- rateLimitDropBytes
	ch <- rateLimitDropPackets
	ch <- totalDropBytes
	ch <- totalDropPackets
}

// Collect collects metrics from JunOS
func (c *interfaceQueueCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	q := InterfaceQueueRPC{}

	err := client.RunCommandAndParse("show interfaces queue", &q)
	if err != nil {
		return err
	}

	for _, iface := range q.InterfaceInformation.Interfaces {
		c.collectForInterface(iface, ch, labelValues)
	}

	return nil
}

func (c *interfaceQueueCollector) collectForInterface(iface PhysicalInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, iface.Name, iface.Description)

	for _, q := range iface.QueueCounters.Queues {
		c.collectForQueue(q, ch, l)
	}
}

func (c *interfaceQueueCollector) collectForQueue(queue Queue, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, queue.Number)

	ch <- prometheus.MustNewConstMetric(queuedPackets, prometheus.CounterValue, float64(queue.QueuedPackets), l...)
	ch <- prometheus.MustNewConstMetric(queuedBytes, prometheus.CounterValue, float64(queue.QueuedBytes), l...)
	ch <- prometheus.MustNewConstMetric(transferedPackets, prometheus.CounterValue, float64(queue.TransferedPackets), l...)
	ch <- prometheus.MustNewConstMetric(transferedBytes, prometheus.CounterValue, float64(queue.TransferedBytes), l...)
	ch <- prometheus.MustNewConstMetric(rateLimitDropPackets, prometheus.CounterValue, float64(queue.RateLimitDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(rateLimitDropBytes, prometheus.CounterValue, float64(queue.RateLimitDropBytes), l...)
	ch <- prometheus.MustNewConstMetric(totalDropBytes, prometheus.CounterValue, float64(queue.TotalDropPackets), l...)
	ch <- prometheus.MustNewConstMetric(totalDropPackets, prometheus.CounterValue, float64(queue.TotalDropBytes), l...)
}
