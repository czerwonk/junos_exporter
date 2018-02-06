package isis

import "github.com/prometheus/client_golang/prometheus"

const prefix string = "junos_isis_"

var (
	upCount    *prometheus.Desc
	totalCount *prometheus.Desc
)

func init() {
	l := []string{"target"}
	upCount = prometheus.NewDesc(prefix+"up", "Number of ISIS Adjacencies in state up", l, nil)
	totalCount = prometheus.NewDesc(prefix+"total", "Number of ISIS Adjacencies", l, nil)
}

type IsisCollector struct {
}

func (*IsisCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upCount
	ch <- totalCount
}

func (c *IsisCollector) Collect(datasource IsisDatasource, ch chan<- prometheus.Metric, labelValues []string) error {
	adjancies, err := datasource.IsisAdjancies()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(upCount, prometheus.GaugeValue, adjancies.Up, labelValues...)
	ch <- prometheus.MustNewConstMetric(totalCount, prometheus.GaugeValue, adjancies.Total, labelValues...)

	return nil
}
