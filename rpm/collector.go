package rpm

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_rpm_probe_results_"

var (
	currRTTMinDesc    *prometheus.Desc
	currRTTMaxDesc    *prometheus.Desc
	currRTTAvgDesc    *prometheus.Desc
	currRTTJitterDesc *prometheus.Desc
	currRTTStddevDesc *prometheus.Desc
	currRTTSumDesc    *prometheus.Desc
	totalSentDesc     *prometheus.Desc
	totalReceivedDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "owner", "name", "address", "type", "interface"}
	totalSentDesc = prometheus.NewDesc(prefix+"sent_total", "Number of probes sent within the current test", l, nil)
	totalReceivedDesc = prometheus.NewDesc(prefix+"received_total", "Number of probe responses received within the current test", l, nil)
	currRTTMinDesc = prometheus.NewDesc(prefix+"rtt_min_current", "Minimum RTT for the most recently completed test, in microseconds", l, nil)
	currRTTMaxDesc = prometheus.NewDesc(prefix+"rtt_max_current", "Maximum RTT for the most recently completed test, in microseconds", l, nil)
	currRTTAvgDesc = prometheus.NewDesc(prefix+"rtt_avg_current", "Average RTT for the most recently completed test, in microseconds", l, nil)
	currRTTJitterDesc = prometheus.NewDesc(prefix+"rtt_jitter_current", "Peak-to-peak difference, in microseconds", l, nil)
	currRTTStddevDesc = prometheus.NewDesc(prefix+"rtt_stddev_current", "Standard deviation, in microseconds", l, nil)
	currRTTSumDesc = prometheus.NewDesc(prefix+"rtt_sum_current", "Statistical sum", l, nil)
}

type rpmCollector struct{}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &rpmCollector{}
}

// Name returns the name of the collector
func (*rpmCollector) Name() string {
	return "RPM"
}

// Describe describes the metrics
func (*rpmCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- totalSentDesc
	ch <- totalReceivedDesc
	ch <- currRTTMinDesc
	ch <- currRTTMaxDesc
	ch <- currRTTAvgDesc
	ch <- currRTTJitterDesc
	ch <- currRTTStddevDesc
	ch <- currRTTSumDesc
}

// Collect collects metrics from JunOS
func (c *rpmCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collect(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *rpmCollector) collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RPMRPC{}

	err := client.RunCommandAndParse("show services rpm probe-results", &x)
	if err != nil {
		return err
	}

	for _, probe := range x.Results.Probes {
		c.collectForProbe(probe, ch, labelValues)
	}

	return nil
}

func (c *rpmCollector) collectForProbe(p RPMProbe, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{p.Owner, p.Name, p.Address, p.Type, p.Interface}...)

	ch <- prometheus.MustNewConstMetric(totalSentDesc, prometheus.GaugeValue, float64(p.Global.Results.Sent), l...)
	ch <- prometheus.MustNewConstMetric(totalReceivedDesc, prometheus.GaugeValue, float64(p.Global.Results.Responses), l...)
	ch <- prometheus.MustNewConstMetric(currRTTMinDesc, prometheus.GaugeValue, float64(p.Last.Results.RTT.Summary.Min), l...)
	ch <- prometheus.MustNewConstMetric(currRTTMaxDesc, prometheus.GaugeValue, float64(p.Last.Results.RTT.Summary.Max), l...)
	ch <- prometheus.MustNewConstMetric(currRTTAvgDesc, prometheus.GaugeValue, float64(p.Last.Results.RTT.Summary.Avg), l...)
	ch <- prometheus.MustNewConstMetric(currRTTJitterDesc, prometheus.GaugeValue, float64(p.Last.Results.RTT.Summary.Jitter), l...)
	ch <- prometheus.MustNewConstMetric(currRTTStddevDesc, prometheus.GaugeValue, float64(p.Last.Results.RTT.Summary.Stddev), l...)
	ch <- prometheus.MustNewConstMetric(currRTTSumDesc, prometheus.GaugeValue, float64(p.Last.Results.RTT.Summary.Sum), l...)
}
