package ospfv3

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_ospf3_"

var (
	upDesc        *prometheus.Desc
	neighborsDesc *prometheus.Desc
)

func init() {
	l := []string{"target"}
	upDesc = prometheus.NewDesc(prefix+"up", "OSPF is up and running (1 = up)", l, nil)

	l = append(l, "area")
	neighborsDesc = prometheus.NewDesc(prefix+"neighbors_count", "Number of neighbors", l, nil)
}

// Collector collects OSPFv3 metrics
type ospfv3Collector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &ospfv3Collector{}
}

// Describe describes the metrics
func (*ospfv3Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- neighborsDesc
}

// Collect collects metrics from JunOS
func (c *ospfv3Collector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	areas, err := c.ospfAreas(client)
	if err != nil {
		return err
	}

	up := 0
	if len(areas) > 0 {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(up), labelValues...)

	for _, area := range areas {
		c.collectForArea(area, ch, labelValues)
	}

	return nil
}

func (c *ospfv3Collector) ospfAreas(client *rpc.Client) ([]*OspfArea, error) {
	var x = Ospf3Rpc{}
	err := client.RunCommandAndParse("show ospf3 overview", &x)
	if err != nil {
		return nil, err
	}

	areas := make([]*OspfArea, 0)
	for _, area := range x.Information.Overview.Areas {
		a := &OspfArea{
			Name:      area.Name,
			Neighbors: float64(area.Neighbors.NeighborsUp),
		}

		areas = append(areas, a)
	}

	return areas, nil
}

func (c *ospfv3Collector) collectForArea(area *OspfArea, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, area.Name)

	ch <- prometheus.MustNewConstMetric(neighborsDesc, prometheus.GaugeValue, area.Neighbors, l...)
}
