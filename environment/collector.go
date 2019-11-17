package environment

import (
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_environment_"

var (
	temperaturesDesc *prometheus.Desc
	powerSupplyDesc  *prometheus.Desc
)

func init() {
	l := []string{"target", "item"}
	temperaturesDesc = prometheus.NewDesc(prefix+"item_temp", "Temperature of the air flowing past", l, nil)
	l = append(l, "status")
	powerSupplyDesc = prometheus.NewDesc(prefix+"power_up", "Status of power supplies (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", l, nil)
}

type environmentCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &environmentCollector{}
}

// Name returns the name of the collector
func (*environmentCollector) Name() string {
	return "Environment"
}

// Describe describes the metrics
func (*environmentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperaturesDesc
}

// Collect collects metrics from JunOS
func (c *environmentCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	tempItems, powerItems, err := c.environmentItems(client)
	if err != nil {
		return err
	}

	for _, item := range tempItems {
		l := append(labelValues, item.Name)

		ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, item.Temperature, l...)
	}

	statusValues := map[string]int{
		"OK":      1,
		"Testing": 2,
		"Failed":  3,
		"Absent":  4,
		"Present": 5,
	}
	for _, item := range powerItems {
		l := append(labelValues, item.Name, item.Status)

		ch <- prometheus.MustNewConstMetric(powerSupplyDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), l...)
	}

	return nil
}

func (c *environmentCollector) environmentItems(client *rpc.Client) ([]*temperatureItem, []*powerItem, error) {
	var x = EnvironmentRpc{}
	err := client.RunCommandAndParse("show chassis environment", &x)
	if err != nil {
		return nil, nil, err
	}

	// remove duplicates
	temperatureList := make(map[string]float64)
	powersupplyList := make(map[string]string)
	for _, item := range x.Information.Items {
		if strings.Contains(item.Name, "Power Supply") || strings.Contains(item.Name, "PEM") {
			powersupplyList[item.Name] = item.Status
		} else if item.Temperature != nil {
			temperatureList[item.Name] = float64(item.Temperature.Value)
		}
	}

	temperatureItems := make([]*temperatureItem, 0)
	for name, value := range temperatureList {
		i := &temperatureItem{
			Name:        name,
			Temperature: value,
		}
		temperatureItems = append(temperatureItems, i)
	}

	powerItems := make([]*powerItem, 0)
	for name, value := range powersupplyList {
		i := &powerItem{
			Name:   name,
			Status: value,
		}
		powerItems = append(powerItems, i)
	}

	return temperatureItems, powerItems, nil
}
