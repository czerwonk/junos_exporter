package mac

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_mac_table_"

var (
	totalCount    *prometheus.Desc
	recieveCount  *prometheus.Desc
	dynamicCount  *prometheus.Desc
	floodCount    *prometheus.Desc
)

func init() {
	l := []string{"target"}
	totalCount = prometheus.NewDesc(prefix+"total_count", "Number of entries in table", l, nil)
	recieveCount = prometheus.NewDesc(prefix+"recieve_count", "Number of L3 recieve route entries in table", l, nil)
	dynamicCount = prometheus.NewDesc(prefix+"dynamic_count", "Number of dynamic entries in table", l, nil)
	floodCount = prometheus.NewDesc(prefix+"flood_count", "Number of flood entries in table", l, nil)
}

type macCollector struct {
}

// Name returns the name of the collector
func (*macCollector) Name() string {
	return "Mac"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &macCollector{}
}

// Describe describes the metrics
func (*macCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- totalCount
	ch <- recieveCount
	ch <- dynamicCount
	ch <- floodCount
}

// Collect collects metrics from JunOS
func (c *macCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = MacRpc{}
	err := client.RunCommandAndParse("show ethernet-switching table summary", &x)
	if err != nil {
		return err
	}

	entry := x.Information.Table.Entry
	ch <- prometheus.MustNewConstMetric(totalCount, prometheus.GaugeValue, float64(entry.TotalCount), labelValues...)
	ch <- prometheus.MustNewConstMetric(recieveCount, prometheus.GaugeValue, float64(entry.ReceiveCount), labelValues...)
	ch <- prometheus.MustNewConstMetric(dynamicCount, prometheus.GaugeValue, float64(entry.DynamicCount), labelValues...)
	ch <- prometheus.MustNewConstMetric(floodCount, prometheus.GaugeValue, float64(entry.FloodCount), labelValues...)

	return nil
}
