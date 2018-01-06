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

type AlarmCollector struct {
}

func (*AlarmCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- alarmsYellowCount
	ch <- alarmsRedCount
}

func (c *AlarmCollector) Collect(datasource AlarmDatasource, ch chan<- prometheus.Metric, labelValues []string) error {
	counter, err := datasource.AlarmCounter()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(alarmsYellowCount, prometheus.GaugeValue, counter.YellowCount, labelValues...)
	ch <- prometheus.MustNewConstMetric(alarmsRedCount, prometheus.GaugeValue, counter.RedCount, labelValues...)

	return nil
}
