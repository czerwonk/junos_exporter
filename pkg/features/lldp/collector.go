// SPDX-License-Identifier: MIT

package lldp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_lldp_"

var (
	lldpPeer           *prometheus.Desc
	lldpInterfaceState *prometheus.Desc
)

func init() {
	l := []string{"target", "local_interface", "parent_interface", "remote_port_info", "remote_system_name"}
	lldpPeer = prometheus.NewDesc(prefix+"peer", "LLDP peer information (1: connected)", l, nil)
	
	l = []string{"target", "local_interface", "parent_interface", "interface_description", "interface_status"}
	lldpInterfaceState = prometheus.NewDesc(prefix+"interface_state", "LLDP interface state (1: up, 0: down)", l, nil)
}

type lldpCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &lldpCollector{}
}

// Name returns the name of the collector
func (*lldpCollector) Name() string {
	return "LLDP"
}

// Describe describes the metrics
func (*lldpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- lldpPeer
	ch <- lldpInterfaceState
}

// Collect collects metrics from JunOS
func (c *lldpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	// Get LLDP neighbors (active connections)
	var neighborResult = result{}
	err := client.RunCommandAndParse("show lldp neighbors", &neighborResult)
	if err != nil {
		return err
	}

	// Get LLDP local information (all LLDP-enabled interfaces)
	var localResult = localResult{}
	err = client.RunCommandAndParse("show lldp local-information", &localResult)
	if err != nil {
		return err
	}

	// Create a map of interfaces with active neighbors
	activeInterfaces := make(map[string]bool)
	for _, neighbor := range neighborResult.Information.Neighbors {
		activeInterfaces[neighbor.LocalPortID] = true
	}

	// Report active LLDP neighbors
	for _, neighbor := range neighborResult.Information.Neighbors {
		l := labelValues
		l = append(l, neighbor.LocalPortID, neighbor.LocalParentInterfaceName, neighbor.RemotePortID, neighbor.RemoteSystemName)
		ch <- prometheus.MustNewConstMetric(lldpPeer, prometheus.GaugeValue, 1.0, l...)
	}

	// Report all LLDP-enabled interfaces and their states
	for _, localInfo := range localResult.Information.LocalInfo {
		for _, iface := range localInfo.LocalInterfaces {
			// Determine interface state (1: up, 0: down)
			state := 0.0
			if iface.InterfaceStatus == "Up" {
				state = 1.0
			}

			l := labelValues
			l = append(l, iface.InterfaceName, iface.ParentInterfaceName, iface.InterfaceDescription, iface.InterfaceStatus)
			ch <- prometheus.MustNewConstMetric(lldpInterfaceState, prometheus.GaugeValue, state, l...)
		}
	}

	return nil
} 
