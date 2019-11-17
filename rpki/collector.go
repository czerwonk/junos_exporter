package rpki

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_rpki_"

var (
	// Session metrics
	upDesc              *prometheus.Desc
	flapsDesc           *prometheus.Desc
	ipv4PrefixCountDesc *prometheus.Desc
	ipv6PrefixCountDesc *prometheus.Desc

	// Statistics metrics
	memoryUtilizationDesc    *prometheus.Desc
	originResultsValidDesc   *prometheus.Desc
	originResultsInvalidDesc *prometheus.Desc
	originResultsUnknownDesc *prometheus.Desc
)

func init() {
	lSession := []string{"target", "ip"}
	upDesc = prometheus.NewDesc(prefix+"session_state", "Session is up (1 = Up)", lSession, nil)
	flapsDesc = prometheus.NewDesc(prefix+"session_flap_count", "Number of session flaps", lSession, nil)
	ipv4PrefixCountDesc = prometheus.NewDesc(prefix+"session_ipv4_prefix_count", "Number of IPv4 route validation records", lSession, nil)
	ipv6PrefixCountDesc = prometheus.NewDesc(prefix+"session_ipv6_prefix_count", "Number of IPv6 route validation records", lSession, nil)

	lStats := []string{"target"}
	stats_prefix := prefix + "statistics_"
	memoryUtilizationDesc = prometheus.NewDesc(stats_prefix+"memory", "Memory utilization of RV database (in bytes)", lStats, nil)
	originResultsValidDesc = prometheus.NewDesc(stats_prefix+"origin_valid", "Origin validation result of valid", lStats, nil)
	originResultsInvalidDesc = prometheus.NewDesc(stats_prefix+"origin_invalid", "Origin validation result of invalid", lStats, nil)
	originResultsUnknownDesc = prometheus.NewDesc(stats_prefix+"origin_unknown", "Origin validation result of unknown", lStats, nil)
}

type rpkiCollector struct {
}

// Name returns the name of the collector
func (*rpkiCollector) Name() string {
	return "RPKI"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &rpkiCollector{}
}

// Describe describes the metrics
func (*rpkiCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- flapsDesc
	ch <- ipv4PrefixCountDesc
	ch <- ipv6PrefixCountDesc
	ch <- memoryUtilizationDesc
	ch <- originResultsValidDesc
	ch <- originResultsInvalidDesc
	ch <- originResultsUnknownDesc
}

// Collect collects metrics from JunOS
func (c *rpkiCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collectSessions(client, ch, labelValues)
	if err != nil {
		return err
	}

	err = c.collectStatistics(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *rpkiCollector) collectSessions(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RpkiSessionRpc{}
	err := client.RunCommandAndParse("show validation session", &x)
	if err != nil {
		return err
	}

	for _, session := range x.Information.RpkiSessions {
		c.collectForSession(session, ch, labelValues)
	}

	return nil
}

func (c *rpkiCollector) collectForSession(s RpkiSession, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.IpAddress}...)

	up := 0
	if s.SessionState == "Up" {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(up), l...)
	ch <- prometheus.MustNewConstMetric(flapsDesc, prometheus.GaugeValue, float64(s.SessionFlaps), l...)
	ch <- prometheus.MustNewConstMetric(ipv4PrefixCountDesc, prometheus.GaugeValue, float64(s.Ipv4PrefixCount), l...)
	ch <- prometheus.MustNewConstMetric(ipv6PrefixCountDesc, prometheus.GaugeValue, float64(s.Ipv6PrefixCount), l...)
}

func (c *rpkiCollector) collectStatistics(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RpkiStatisticsRpc{}

	err := client.RunCommandAndParse("show validation statistics", &x)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(memoryUtilizationDesc, prometheus.GaugeValue, float64(x.Information.Statistics.MemoryUtilization), labelValues...)
	ch <- prometheus.MustNewConstMetric(originResultsValidDesc, prometheus.GaugeValue, float64(x.Information.Statistics.OriginResultsValid), labelValues...)
	ch <- prometheus.MustNewConstMetric(originResultsInvalidDesc, prometheus.GaugeValue, float64(x.Information.Statistics.OriginResultsInvalid), labelValues...)
	ch <- prometheus.MustNewConstMetric(originResultsUnknownDesc, prometheus.GaugeValue, float64(x.Information.Statistics.OriginResultsUnknown), labelValues...)

	return nil
}
