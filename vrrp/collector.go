package vrrp

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_vrrp_"

var (
	vrrpState  *prometheus.Desc
)

func init() {
	l := []string{"target", "interface", "group", "local_interface_address", "virtual_ip_address"}
	vrrpState = prometheus.NewDesc(prefix+"state", "VRRP state (1: init, 2: backup, 3: master)", l, nil)
}

type vrrpCollector struct {
}

// Name returns the name of the collector
func (*vrrpCollector) Name() string {
	return "VRRP"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &vrrpCollector{}
}

// Describe describes the metrics
func (*vrrpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- vrrpState
}


// Collect collects metrics from JunOS
func (c *vrrpCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
        statusValues := map[string]int{
		"init":   1,
		"backup": 2,
		"master": 3,
	}

        var x = VrrpRpc{}
        err := client.RunCommandAndParse("show vrrp summary", &x)
	if err != nil {
		return err
	}

	for _, iface := range x.Information.Interfaces {
                l := labelValues
                l = append(l, iface.Interface, iface.Group, iface.LocalInterfaceAddress, iface.VirtualIpAddress)
                ch <- prometheus.MustNewConstMetric(vrrpState, prometheus.GaugeValue, float64(statusValues[iface.VrrpState]), l...)
	}

	return nil
}
