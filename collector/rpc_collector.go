package collector

import (
	"github.com/czerwonk/junos_exporter/rpc"

	"github.com/prometheus/client_golang/prometheus"
)

// RPCCollector collects metrics from JunOS using rpc.Client
type RPCCollector interface {
	// Name returns an human readable name for logging and debugging purposes
	Name() string

	// Describe describes the metrics
	Describe(ch chan<- *prometheus.Desc)

	// Collect collects metrics from JunOS
	Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error
}
