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
	tables, err := c.routingTables(client)
	if err != nil {
		return err
	}

	for _, table := range tables {
		c.collectForTable(table, ch, labelValues)
	}

	return nil
}

func (c *routeCollector) routingTables(client *rpc.Client) ([]*RoutingTable, error) {
	var x = RouteRpc{}
	err := client.RunCommandAndParse("show route summary", &x)
	if err != nil {
		return nil, err
	}

	tables := make([]*RoutingTable, 0)
	for _, table := range x.Information.Tables {
		t := &RoutingTable{
			Name:         table.Name,
			MaxRoutes:    float64(table.MaxRoutes),
			ActiveRoutes: float64(table.ActiveRoutes),
			TotalRoutes:  float64(table.TotalRoutes),
			Protocols:    make([]*ProtocolRouteCount, 0),
		}

		for _, proto := range table.Protocols {
			p := &ProtocolRouteCount{
				Name:         proto.Name,
				Routes:       float64(proto.Routes),
				ActiveRoutes: float64(proto.ActiveRoutes),
			}

			t.Protocols = append(t.Protocols, p)
		}

		tables = append(tables, t)
	}

	return tables, nil
}

func (c *routeCollector) collectForTable(table *RoutingTable, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, table.Name)

	ch <- prometheus.MustNewConstMetric(totalRoutesDesc, prometheus.GaugeValue, table.TotalRoutes, l...)
	ch <- prometheus.MustNewConstMetric(activeRoutesDesc, prometheus.GaugeValue, table.ActiveRoutes, l...)
	ch <- prometheus.MustNewConstMetric(maxRoutesDesc, prometheus.GaugeValue, table.MaxRoutes, l...)

	for _, proto := range table.Protocols {
		lp := append(l, proto.Name)
		ch <- prometheus.MustNewConstMetric(protocolRoutes, prometheus.GaugeValue, proto.Routes, lp...)
		ch <- prometheus.MustNewConstMetric(protocolActiveRoutes, prometheus.GaugeValue, proto.ActiveRoutes, lp...)
	}
}
