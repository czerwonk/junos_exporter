// SPDX-License-Identifier: MIT

package lldp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_lldp_"

var (
	lldpPeer *prometheus.Desc
)

func init() {
	l := []string{"target", "local_interface", "parent_interface", "remote_port_info", "remote_system_name", "interface_status"}
	lldpPeer = prometheus.NewDesc(prefix+"peer", "LLDP peer information (1: connected)", l, nil)
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

	// Create a map of common management interfaces to filter out
	mgmtInterfaces := make(map[string]bool)
	commonMgmtInterfaces := []string{"fxp0", "fxp1", "fxp2", "em0", "em1", "em2", "me0", "me1", "me2", "mgmt0", "mgmt1"}
	for _, iface := range commonMgmtInterfaces {
		mgmtInterfaces[iface] = true
	}

	// Report all UP interfaces that are not management interfaces
	for _, iface := range localResult.Information.LocalInterfaces {
		// Only report interfaces that are UP in the local information
		if iface.InterfaceStatus == "Up" {
			// Skip interfaces that are common management interfaces
			if mgmtInterfaces[iface.InterfaceName] || mgmtInterfaces[iface.ParentInterfaceName] {
				continue
			}

			// Determine interface state based on whether it has active neighbors
			state := 0.0
			if activeInterfaces[iface.InterfaceName] {
				state = 1.0
			}

			// For interfaces without neighbors, use empty values for remote fields
			remotePortInfo := ""
			remoteSystemName := ""
			if activeInterfaces[iface.InterfaceName] {
				// Find the neighbor info for this interface
				for _, neighbor := range neighborResult.Information.Neighbors {
					if neighbor.LocalPortID == iface.InterfaceName {
						remotePortInfo = neighbor.RemotePortID
						remoteSystemName = neighbor.RemoteSystemName
						break
					}
				}
			}

			l := labelValues
			l = append(l, iface.InterfaceName, iface.ParentInterfaceName, remotePortInfo, remoteSystemName, iface.InterfaceStatus)
			ch <- prometheus.MustNewConstMetric(lldpPeer, prometheus.GaugeValue, state, l...)
		}
	}

	return nil
}
