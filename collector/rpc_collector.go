package collector

import (
	"github.com/czerwonk/junos_exporter/rpc"

	"github.com/prometheus/client_golang/prometheus"
)

// RPCCollector collects metrics from JunOS using rpc.Client
type RPCCollector interface {

	// Describe describes the metrics
	Describe(ch chan<- *prometheus.Desc)

	// Collect collects metrics from JunOS
	Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error
}
