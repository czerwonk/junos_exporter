// SPDX-License-Identifier: MIT

package lldp

import (
	"strings"

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

	// Get interfaces in mgmt_junos routing instance to filter them out
	var mgmtResult = routingInstanceResult{}
	err = client.RunCommandAndParse("show interfaces routing-instance mgmt_junos", &mgmtResult)
	if err != nil {
		// If the command fails (e.g., mgmt_junos doesn't exist), continue without filtering
		// This is not a critical error
	}

	// Create a map of interfaces with active neighbors
	activeInterfaces := make(map[string]bool)
	for _, neighbor := range neighborResult.Information.Neighbors {
		activeInterfaces[neighbor.LocalPortID] = true
	}

	// Create a map of interfaces in mgmt_junos routing instance
	mgmtInterfaces := make(map[string]bool)
	for _, info := range mgmtResult.Information {
		for _, phy := range info.PhysicalInterfaces {
			for _, logical := range phy.LogicalInterfaces {
				// Extract physical interface name from logical interface name (e.g., fxp0.0 -> fxp0)
				if logical.Name != "" {
					// Split on dot and take the first part
					parts := strings.Split(logical.Name, ".")
					if len(parts) > 0 {
						physicalName := parts[0]
						mgmtInterfaces[physicalName] = true
					}
				}
			}
		}
	}

	// Fallback: Add common management interfaces if the routing instance command didn't work
	if len(mgmtInterfaces) == 0 {
		commonMgmtInterfaces := []string{"fxp0", "fxp1", "fxp2", "em0", "em1", "em2", "me0", "me1", "me2"}
		for _, iface := range commonMgmtInterfaces {
			mgmtInterfaces[iface] = true
		}
	}

	// Report all UP interfaces that are not in mgmt_junos
	for _, iface := range localResult.Information.LocalInterfaces {
		// Only report interfaces that are UP in the local information
		if iface.InterfaceStatus == "Up" {
			// Skip interfaces that are in the mgmt_junos routing instance
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
