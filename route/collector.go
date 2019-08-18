package route

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_routes_"

var (
	totalRoutesDesc      *prometheus.Desc
	activeRoutesDesc     *prometheus.Desc
	maxRoutesDesc        *prometheus.Desc
	protocolRoutes       *prometheus.Desc
	protocolActiveRoutes *prometheus.Desc
)

func init() {
	l := []string{"target", "table"}
	totalRoutesDesc = prometheus.NewDesc(prefix+"total_count", "Number of routes in table", l, nil)
	activeRoutesDesc = prometheus.NewDesc(prefix+"active_count", "Number of active routes in table", l, nil)
	maxRoutesDesc = prometheus.NewDesc(prefix+"max_count", "Max. number of routes", l, nil)

	l = append(l, "protocol")
	protocolRoutes = prometheus.NewDesc(prefix+"protocol_count", "Number of routes by protocol in table", l, nil)
	protocolActiveRoutes = prometheus.NewDesc(prefix+"protocol_active_count", "Number of active routes by protocol in table", l, nil)
}

type routeCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &routeCollector{}
}

// Describe describes the metrics
func (*routeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- totalRoutesDesc
	ch <- activeRoutesDesc
	ch <- maxRoutesDesc
	ch <- protocolRoutes
	ch <- protocolActiveRoutes
}

// Collect collects metrics from JunOS
func (c *routeCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RouteRpc{}
	err := client.RunCommandAndParse("show route summary", &x)
	if err != nil {
		return err
	}

	for _, t := range x.Information.Tables {
		c.collectForTable(t, ch, labelValues)
	}

	return nil
}

func (c *routeCollector) collectForTable(table RouteTable, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, table.Name)

	ch <- prometheus.MustNewConstMetric(totalRoutesDesc, prometheus.GaugeValue, float64(table.TotalRoutes), l...)
	ch <- prometheus.MustNewConstMetric(activeRoutesDesc, prometheus.GaugeValue, float64(table.ActiveRoutes), l...)
	ch <- prometheus.MustNewConstMetric(maxRoutesDesc, prometheus.GaugeValue, float64(table.MaxRoutes), l...)

	for _, proto := range table.Protocols {
		lp := append(l, proto.Name)
		ch <- prometheus.MustNewConstMetric(protocolRoutes, prometheus.GaugeValue, float64(proto.Routes), lp...)
		ch <- prometheus.MustNewConstMetric(protocolActiveRoutes, prometheus.GaugeValue, float64(proto.ActiveRoutes), lp...)
	}
}
