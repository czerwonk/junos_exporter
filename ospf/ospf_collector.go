package ospf

import "github.com/prometheus/client_golang/prometheus"

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
type Collector struct {
}

// Describe describes the metrics
func (*Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- neighborsDesc
}

// Collect collects metrics from datasource
func (c *Collector) Collect(datasource OspfDatasource, ch chan<- prometheus.Metric, labelValues []string) error {
	areas, err := datasource.OspfAreas()
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

func (c *Collector) collectForArea(area *OspfArea, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, area.Name)

	ch <- prometheus.MustNewConstMetric(neighborsDesc, prometheus.GaugeValue, area.Neighbors, l...)
}
