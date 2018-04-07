package alarm

import "github.com/prometheus/client_golang/prometheus"

const prefix = "junos_alarms_"

var (
	alarmsYellowCount *prometheus.Desc
	alarmsRedCount    *prometheus.Desc
)

func init() {
	l := []string{"target"}
	alarmsYellowCount = prometheus.NewDesc(prefix+"yellow_count", "Number of yollow alarms (not silenced)", l, nil)
	alarmsRedCount = prometheus.NewDesc(prefix+"red_count", "Number of red alarms (not silenced)", l, nil)
}

// Collector collects alarm metrics
type Collector struct {
}

// Describe describes the metrics
func (*Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- alarmsYellowCount
	ch <- alarmsRedCount
}

// Collect collects metrics from datasource
func (c *Collector) Collect(datasource AlarmDatasource, ch chan<- prometheus.Metric, labelValues []string) error {
	counter, err := datasource.AlarmCounter()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(alarmsYellowCount, prometheus.GaugeValue, counter.YellowCount, labelValues...)
	ch <- prometheus.MustNewConstMetric(alarmsRedCount, prometheus.GaugeValue, counter.RedCount, labelValues...)

	return nil
}
