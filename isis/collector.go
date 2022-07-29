package isis

import (
	"strconv"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_isis_"

var (
	upCount    *prometheus.Desc
	totalCount *prometheus.Desc
	adjState   *prometheus.Desc
)

func init() {
	l := []string{"target"}
	upCount = prometheus.NewDesc(prefix+"up_count", "Number of ISIS Adjacencies in state up", l, nil)
	totalCount = prometheus.NewDesc(prefix+"total_count", "Number of ISIS Adjacencies", l, nil)
	l = append(l, "interface_name", "sysem_name", "level")
	adjState = prometheus.NewDesc(prefix+"adjacency_state", "The ISIS Adjacency state (1 = UP, 0 = DOWN)", l, nil)
}

type isisCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &isisCollector{}
}

// Name returns the name of the collector
func (*isisCollector) Name() string {
	return "ISIS"
}

// Describe describes the metrics
func (*isisCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upCount
	ch <- totalCount
}

// Collect collects metrics from JunOS
func (c *isisCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	adjancies, err := c.isisAdjancies(client)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(upCount, prometheus.GaugeValue, adjancies.Up, labelValues...)
	ch <- prometheus.MustNewConstMetric(totalCount, prometheus.GaugeValue, adjancies.Total, labelValues...)
	if adjancies.Adjacencies != nil {
		for _, adj := range adjancies.Adjacencies {
			localLabelvalues := append(labelValues, adj.InterfaceName, adj.SystemName, strconv.Itoa(int(adj.Level)))
			state := 0.0
			switch adj.AdjacencyState {
			case "Down":
				state = 0.0
			case "Up":
				state = 1.0
			case "New":
				state = 2.0
			case "One-way":
				state = 3.0
			case "Initializing":
				state = 4.0
			case "Rejected":
				state = 5.0
			}
			ch <- prometheus.MustNewConstMetric(adjState, prometheus.GaugeValue, state, localLabelvalues...)
		}

	}

	return nil
}

func (c *isisCollector) isisAdjancies(client *rpc.Client) (*IsisAdjacencies, error) {
	up := 0
	total := 0

	var x = IsisRpc{}
	err := client.RunCommandAndParse("show isis adjacency", &x)
	if err != nil {
		return nil, err
	}

	for _, adjacency := range x.Information.Adjacencies {
		if adjacency.AdjacencyState == "Up" {
			up++
		}
		total++
	}

	return &IsisAdjacencies{Up: float64(up), Total: float64(total), Adjacencies: x.Information.Adjacencies}, nil
}
