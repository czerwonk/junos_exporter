package krt

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

const prefix string = "junos_krt"

var (
	queueLengthDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "krtq_type"}
	queueLengthDesc = prometheus.NewDesc(prefix+"krt_queue_length", "KRT queue length", l, nil)
}

type krtCollector struct{}

func NewCollector() collector.RPCCollector { return &krtCollector{} }

func (c *krtCollector) Name() string { return "krt" }

func (c *krtCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- queueLengthDesc
}

func (c *krtCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var k resultKRT
	err := client.RunCommandAndParse("show krt queue", &k)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show krt queue'")
	}
	c.collectKRT(k, ch, labelValues)
	return nil
}

func (c *krtCollector) collectKRT(k resultKRT, ch chan<- prometheus.Metric, labelValues []string) {
	for _, q := range k.KrtQueueInformation.KrtQueue {
		labels := append(labelValues, q.KrtqType)
		ch <- prometheus.MustNewConstMetric(queueLengthDesc, prometheus.GaugeValue, q.KrtqQueueLength, labels...)
	}
}
