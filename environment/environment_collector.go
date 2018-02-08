package environment

import "github.com/prometheus/client_golang/prometheus"

const prefix string = "junos_environment_"

var (
	temperaturesDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "item"}
	temperaturesDesc = prometheus.NewDesc(prefix+"item_temp", "Temperature of the air flowing past", l, nil)
}

type EnvironmentCollector struct {
}

func (*EnvironmentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperaturesDesc
}

func (c *EnvironmentCollector) Collect(datasource EnvironmentDatasource, ch chan<- prometheus.Metric, labelValues []string) error {
	items, err := datasource.EnvironmentItems()
	if err != nil {
		return err
	}

	for _, item := range items {
		l := append(labelValues, item.Name)

		ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, item.Temperature, l...)
	}

	return nil
}
