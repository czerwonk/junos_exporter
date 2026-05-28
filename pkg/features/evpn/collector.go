// SPDX-License-Identifier: MIT

package evpn

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_evpn_"

var (
	// Per-instance descriptors.
	instanceInfo                   *prometheus.Desc
	instanceNeighborCount          *prometheus.Desc
	instanceESICount               *prometheus.Desc
	instanceLocalInterfaces        *prometheus.Desc
	instanceLocalInterfacesUp      *prometheus.Desc
	instanceIRBInterfaces          *prometheus.Desc
	instanceIRBInterfacesUp        *prometheus.Desc
	instanceProtectInterfaces      *prometheus.Desc
	instanceBridgeDomains          *prometheus.Desc
	instanceLocalMACs              *prometheus.Desc
	instanceRemoteMACs             *prometheus.Desc
	instanceLocalMACIPs            *prometheus.Desc
	instanceRemoteMACIPs           *prometheus.Desc
	instanceLocalDefaultGwMACs     *prometheus.Desc
	instanceRemoteDefaultGwMACs    *prometheus.Desc
	instanceDuplicateMACThreshold  *prometheus.Desc
	instanceDuplicateMACWindow     *prometheus.Desc

	// Per-neighbor descriptors.
	neighborMACRoutes               *prometheus.Desc
	neighborMACIPRoutes             *prometheus.Desc
	neighborAutoDiscoveryRoutes     *prometheus.Desc
	neighborInclusiveMulticastRoutes *prometheus.Desc
	neighborEthernetSegmentRoutes   *prometheus.Desc
)

func init() {
	instanceLabels := []string{"target", "instance"}
	infoLabels := []string{"target", "instance", "rd", "encap", "router_id", "source_vtep"}
	neighborLabels := []string{"target", "instance", "neighbor"}

	instanceInfo = prometheus.NewDesc(
		prefix+"instance_info",
		"Per-EVI metadata (always 1). Labels carry route-distinguisher, encapsulation, router-id and source VTEP address.",
		infoLabels, nil,
	)
	instanceNeighborCount = prometheus.NewDesc(
		prefix+"instance_neighbor_count",
		"Number of EVPN neighbors (PEs) for this instance",
		instanceLabels, nil,
	)
	instanceESICount = prometheus.NewDesc(
		prefix+"instance_esi_count",
		"Number of Ethernet Segment Identifiers known by this instance",
		instanceLabels, nil,
	)
	instanceLocalInterfaces = prometheus.NewDesc(
		prefix+"instance_local_interfaces",
		"Number of local interfaces in this EVPN instance",
		instanceLabels, nil,
	)
	instanceLocalInterfacesUp = prometheus.NewDesc(
		prefix+"instance_local_interfaces_up",
		"Number of local interfaces currently up in this EVPN instance",
		instanceLabels, nil,
	)
	instanceIRBInterfaces = prometheus.NewDesc(
		prefix+"instance_irb_interfaces",
		"Number of IRB interfaces in this EVPN instance",
		instanceLabels, nil,
	)
	instanceIRBInterfacesUp = prometheus.NewDesc(
		prefix+"instance_irb_interfaces_up",
		"Number of IRB interfaces currently up in this EVPN instance",
		instanceLabels, nil,
	)
	instanceProtectInterfaces = prometheus.NewDesc(
		prefix+"instance_protect_interfaces",
		"Number of protect (backup) interfaces in this EVPN instance",
		instanceLabels, nil,
	)
	instanceBridgeDomains = prometheus.NewDesc(
		prefix+"instance_bridge_domains",
		"Number of bridge domains in this EVPN instance",
		instanceLabels, nil,
	)
	instanceLocalMACs = prometheus.NewDesc(
		prefix+"instance_local_mac_count",
		"Number of MACs learned locally in this EVPN instance",
		instanceLabels, nil,
	)
	instanceRemoteMACs = prometheus.NewDesc(
		prefix+"instance_remote_mac_count",
		"Number of MACs learned from remote PEs in this EVPN instance",
		instanceLabels, nil,
	)
	instanceLocalMACIPs = prometheus.NewDesc(
		prefix+"instance_local_mac_ip_count",
		"Number of local MAC+IP bindings in this EVPN instance",
		instanceLabels, nil,
	)
	instanceRemoteMACIPs = prometheus.NewDesc(
		prefix+"instance_remote_mac_ip_count",
		"Number of remote MAC+IP bindings in this EVPN instance",
		instanceLabels, nil,
	)
	instanceLocalDefaultGwMACs = prometheus.NewDesc(
		prefix+"instance_local_default_gateway_mac_count",
		"Number of local default-gateway MACs in this EVPN instance",
		instanceLabels, nil,
	)
	instanceRemoteDefaultGwMACs = prometheus.NewDesc(
		prefix+"instance_remote_default_gateway_mac_count",
		"Number of remote default-gateway MACs in this EVPN instance",
		instanceLabels, nil,
	)
	instanceDuplicateMACThreshold = prometheus.NewDesc(
		prefix+"instance_duplicate_mac_threshold",
		"Configured duplicate-MAC detection threshold for this EVPN instance",
		instanceLabels, nil,
	)
	instanceDuplicateMACWindow = prometheus.NewDesc(
		prefix+"instance_duplicate_mac_window_seconds",
		"Configured duplicate-MAC detection window in seconds for this EVPN instance",
		instanceLabels, nil,
	)

	neighborMACRoutes = prometheus.NewDesc(
		prefix+"neighbor_mac_routes",
		"Per-neighbor count of EVPN Type-2 MAC routes (without IP)",
		neighborLabels, nil,
	)
	neighborMACIPRoutes = prometheus.NewDesc(
		prefix+"neighbor_mac_ip_routes",
		"Per-neighbor count of EVPN Type-2 MAC+IP routes",
		neighborLabels, nil,
	)
	neighborAutoDiscoveryRoutes = prometheus.NewDesc(
		prefix+"neighbor_ethernet_autodiscovery_routes",
		"Per-neighbor count of EVPN Type-1 Ethernet auto-discovery routes",
		neighborLabels, nil,
	)
	neighborInclusiveMulticastRoutes = prometheus.NewDesc(
		prefix+"neighbor_inclusive_multicast_routes",
		"Per-neighbor count of EVPN Type-3 inclusive-multicast routes",
		neighborLabels, nil,
	)
	neighborEthernetSegmentRoutes = prometheus.NewDesc(
		prefix+"neighbor_ethernet_segment_routes",
		"Per-neighbor count of EVPN Type-4 Ethernet-segment routes",
		neighborLabels, nil,
	)
}

type evpnCollector struct{}

// Name returns the name of the collector.
func (*evpnCollector) Name() string { return "evpn" }

// NewCollector creates a new collector.
func NewCollector() collector.RPCCollector { return &evpnCollector{} }

// Describe describes the metrics.
func (*evpnCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- instanceInfo
	ch <- instanceNeighborCount
	ch <- instanceESICount
	ch <- instanceLocalInterfaces
	ch <- instanceLocalInterfacesUp
	ch <- instanceIRBInterfaces
	ch <- instanceIRBInterfacesUp
	ch <- instanceProtectInterfaces
	ch <- instanceBridgeDomains
	ch <- instanceLocalMACs
	ch <- instanceRemoteMACs
	ch <- instanceLocalMACIPs
	ch <- instanceRemoteMACIPs
	ch <- instanceLocalDefaultGwMACs
	ch <- instanceRemoteDefaultGwMACs
	ch <- instanceDuplicateMACThreshold
	ch <- instanceDuplicateMACWindow

	ch <- neighborMACRoutes
	ch <- neighborMACIPRoutes
	ch <- neighborAutoDiscoveryRoutes
	ch <- neighborInclusiveMulticastRoutes
	ch <- neighborEthernetSegmentRoutes
}

// Collect collects metrics from Junos. The extensive modifier is required:
// the basic 'show evpn instance' output is too thin (omits route-distinguisher,
// encapsulation, the per-PE route breakdown, and the full
// mac-database-status-table, and renames remaining scalars).
func (c *evpnCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var instances []evpnInstance
	err := client.RunCommandAndParseWithParser("show evpn instance extensive", func(b []byte) error {
		var perr error
		instances, perr = parseInstances(b)
		return perr
	})
	if err != nil {
		return err
	}

	for _, in := range instances {
		il := append(labelValues, in.Name)

		ch <- prometheus.MustNewConstMetric(
			instanceInfo, prometheus.GaugeValue, 1,
			append(il, in.RouteDistinguisher, strings.TrimSpace(in.EncapType), in.RouterID, in.SourceVTEPAddr)...,
		)

		ch <- prometheus.MustNewConstMetric(instanceNeighborCount, prometheus.GaugeValue, in.NumNeighbors, il...)
		ch <- prometheus.MustNewConstMetric(instanceESICount, prometheus.GaugeValue, in.NumESI, il...)
		ch <- prometheus.MustNewConstMetric(instanceLocalInterfaces, prometheus.GaugeValue, in.LocalInterfaces, il...)
		ch <- prometheus.MustNewConstMetric(instanceLocalInterfacesUp, prometheus.GaugeValue, in.LocalInterfacesUp, il...)
		ch <- prometheus.MustNewConstMetric(instanceIRBInterfaces, prometheus.GaugeValue, in.IRBInterfaces, il...)
		ch <- prometheus.MustNewConstMetric(instanceIRBInterfacesUp, prometheus.GaugeValue, in.IRBInterfacesUp, il...)
		ch <- prometheus.MustNewConstMetric(instanceProtectInterfaces, prometheus.GaugeValue, in.NumProtectInterfaces, il...)
		ch <- prometheus.MustNewConstMetric(instanceBridgeDomains, prometheus.GaugeValue, in.NumBridgeDomains, il...)
		ch <- prometheus.MustNewConstMetric(instanceDuplicateMACThreshold, prometheus.GaugeValue, in.DuplicateMACDetectionThreshold, il...)
		ch <- prometheus.MustNewConstMetric(instanceDuplicateMACWindow, prometheus.GaugeValue, in.DuplicateMACDetectionWindow, il...)

		if in.MACDatabase != nil {
			db := in.MACDatabase
			ch <- prometheus.MustNewConstMetric(instanceLocalMACs, prometheus.GaugeValue, db.LocalMACs, il...)
			ch <- prometheus.MustNewConstMetric(instanceRemoteMACs, prometheus.GaugeValue, db.RemoteMACs, il...)
			ch <- prometheus.MustNewConstMetric(instanceLocalMACIPs, prometheus.GaugeValue, db.LocalMACIPs, il...)
			ch <- prometheus.MustNewConstMetric(instanceRemoteMACIPs, prometheus.GaugeValue, db.RemoteMACIPs, il...)
			ch <- prometheus.MustNewConstMetric(instanceLocalDefaultGwMACs, prometheus.GaugeValue, db.LocalDefaultGatewayMACs, il...)
			ch <- prometheus.MustNewConstMetric(instanceRemoteDefaultGwMACs, prometheus.GaugeValue, db.RemoteDefaultGatewayMACs, il...)
		}

		for _, n := range in.Neighbors {
			r := n.Routes
			if r.Address == "" {
				continue
			}
			nl := append(append([]string{}, il...), r.Address)
			ch <- prometheus.MustNewConstMetric(neighborMACRoutes, prometheus.GaugeValue, r.NumMACRoutes, nl...)
			ch <- prometheus.MustNewConstMetric(neighborMACIPRoutes, prometheus.GaugeValue, r.NumMACIPRoutes, nl...)
			ch <- prometheus.MustNewConstMetric(neighborAutoDiscoveryRoutes, prometheus.GaugeValue, r.NumAutoDiscoveryRoutes, nl...)
			ch <- prometheus.MustNewConstMetric(neighborInclusiveMulticastRoutes, prometheus.GaugeValue, r.NumInclusiveMulticastRoutes, nl...)
			ch <- prometheus.MustNewConstMetric(neighborEthernetSegmentRoutes, prometheus.GaugeValue, r.NumEthernetSegmentRoutes, nl...)
		}
	}

	return nil
}

// parseInstances handles both single-RE and multi-RE response shapes.
// Same dispatch pattern as the alarm and ufd collectors. When multiple REs
// report instances they are merged; re-name is not exposed as a label today
// (matches alarm/ufd precedent).
func parseInstances(b []byte) ([]evpnInstance, error) {
	if strings.Contains(string(b), "<multi-routing-engine-results") {
		var m multiEngineResult
		if err := xml.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		var out []evpnInstance
		for _, re := range m.Engines.Items {
			out = append(out, re.Instances...)
		}
		return out, nil
	}

	var s singleEngineResult
	if err := xml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s.Instances, nil
}
