// SPDX-License-Identifier: MIT

package mac

import (
	"fmt"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_mac_table_"

var (
	totalCount   *prometheus.Desc
	recieveCount *prometheus.Desc
	dynamicCount *prometheus.Desc
	floodCount   *prometheus.Desc
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
func (c *macCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x result
	err := client.RunCommandAndParse("show ethernet-switching table summary", &x)
	if err != nil {
		return err
	}

	if x.NewMacdb != nil {
		// new XML element l2ng-l2ald-rtb-macdb is present
		ch <- prometheus.MustNewConstMetric(totalCount, prometheus.GaugeValue, float64(x.NewMacdb.TableSummary.TotalMacCount), labelValues...)
		ch <- prometheus.MustNewConstMetric(recieveCount, prometheus.GaugeValue, 0, labelValues...)
		ch <- prometheus.MustNewConstMetric(dynamicCount, prometheus.GaugeValue, 0, labelValues...)
		ch <- prometheus.MustNewConstMetric(floodCount, prometheus.GaugeValue, 0, labelValues...)
		return nil
	}

	if x.OldInformation != nil {
		// old XML element ethernet-switching-table-information is present
		entry := x.OldInformation.Table.Entry
		ch <- prometheus.MustNewConstMetric(totalCount, prometheus.GaugeValue, float64(entry.TotalCount), labelValues...)
		ch <- prometheus.MustNewConstMetric(recieveCount, prometheus.GaugeValue, float64(entry.ReceiveCount), labelValues...)
		ch <- prometheus.MustNewConstMetric(dynamicCount, prometheus.GaugeValue, float64(entry.DynamicCount), labelValues...)
		ch <- prometheus.MustNewConstMetric(floodCount, prometheus.GaugeValue, float64(entry.FloodCount), labelValues...)
		return nil
	}

	return fmt.Errorf("neither old nor new MAC table data found in XML")
}
