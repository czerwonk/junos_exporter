package environment

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_environment_"

var (
	temperaturesDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "item"}
	temperaturesDesc = prometheus.NewDesc(prefix+"item_temp", "Temperature of the air flowing past", l, nil)
}

type environmentCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &environmentCollector{}
}

// Describe describes the metrics
func (*environmentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperaturesDesc
}

// Collect collects metrics from JunOS
func (c *environmentCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	items, err := c.environmentItems(client)
	if err != nil {
		return err
	}

	for _, item := range items {
		l := append(labelValues, item.Name)

		ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, item.Temperature, l...)
	}

	return nil
}

func (c *environmentCollector) environmentItems(client *rpc.Client) ([]*EnvironmentItem, error) {
	var x = EnvironmentRpc{}
	err := client.RunCommandAndParse("show chassis environment", &x)
	if err != nil {
		return nil, err
	}

	// remove duplicates
	list := make(map[string]float64)
	for _, item := range x.Information.Items {
		if item.Temperature != nil {
			list[item.Name] = float64(item.Temperature.Value)
		}
	}

	items := make([]*EnvironmentItem, 0)
	for name, value := range list {
		i := &EnvironmentItem{
			Name:        name,
			Temperature: value,
		}
		items = append(items, i)
	}

	return items, nil
}
