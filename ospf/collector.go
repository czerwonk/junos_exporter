package ospf

import (
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ospfUpDesc         *prometheus.Desc
	ospf3UpDesc        *prometheus.Desc
	ospfNeighborsDesc  *prometheus.Desc
	ospf3NeighborsDesc *prometheus.Desc
)

func init() {
	ospfPrefix := "junos_ospf_"
	ospf3Prefix := "junos_ospf3_"

	l := []string{"target"}
	ospfUpDesc = prometheus.NewDesc(ospfPrefix+"up", "OSPF is up and running (1 = up)", l, nil)
	ospf3UpDesc = prometheus.NewDesc(ospf3Prefix+"up", "OSPFv3 is up and running (1 = up)", l, nil)

	l = append(l, "area")
	ospfNeighborsDesc = prometheus.NewDesc(ospfPrefix+"neighbors_count", "Number of neighbors", l, nil)
	ospf3NeighborsDesc = prometheus.NewDesc(ospf3Prefix+"neighbors_count", "Number of neighbors", l, nil)
}

// Collector collects OSPFv3 metrics
type ospfCollector struct {
	LogicalSystem string
}

// NewCollector creates a new collector
func NewCollector(logicalSystem string) collector.RPCCollector {
	return &ospfCollector{LogicalSystem: logicalSystem}
}

// Describe describes the metrics
func (*ospfCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- ospfUpDesc
	ch <- ospf3UpDesc
	ch <- ospfNeighborsDesc
	ch <- ospf3NeighborsDesc
}

// Collect collects metrics from JunOS
func (c *ospfCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collectOSPFMetrics(client, ch, labelValues)
	if err != nil {
		return err
	}

	return c.collectOSPFv3Metrics(client, ch, labelValues)
}

func (c *ospfCollector) collectOSPFMetrics(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = OspfRpc{}
	var cmd strings.Builder
	cmd.WriteString("show ospf overview")
	if c.LogicalSystem != "" {
		cmd.WriteString(" logical-system " + c.LogicalSystem)
	}

	err := client.RunCommandAndParse(cmd.String(), &x)
	if err != nil {
		return err
	}

	areas := x.Information.Overview.Areas

	up := 0
	if len(areas) > 0 {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(ospfUpDesc, prometheus.GaugeValue, float64(up), labelValues...)

	for _, a := range areas {
		l := append(labelValues, a.Name)
		ch <- prometheus.MustNewConstMetric(ospfNeighborsDesc, prometheus.GaugeValue, float64(a.Neighbors.NeighborsUp), l...)
	}

	return nil
}

func (c *ospfCollector) collectOSPFv3Metrics(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = Ospf3Rpc{}
	var cmd strings.Builder
	cmd.WriteString("show ospf3 overview")
	if c.LogicalSystem != "" {
		cmd.WriteString(" logical-system " + c.LogicalSystem)
	}

	err := client.RunCommandAndParse(cmd.String(), &x)
	if err != nil {
		return err
	}

	areas := x.Information.Overview.Areas

	up := 0
	if len(areas) > 0 {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(ospf3UpDesc, prometheus.GaugeValue, float64(up), labelValues...)

	for _, a := range areas {
		l := append(labelValues, a.Name)
		ch <- prometheus.MustNewConstMetric(ospf3NeighborsDesc, prometheus.GaugeValue, float64(a.Neighbors.NeighborsUp), l...)
	}

	return nil
}
