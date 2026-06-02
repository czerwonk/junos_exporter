// SPDX-License-Identifier: MIT

package evpn

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const prefix = "junos_evpn_"

var (
	// ============== Per-instance descriptors (from show evpn instance) ==============
	instanceInfo                  *prometheus.Desc
	instanceNeighborCount         *prometheus.Desc
	instanceESICount              *prometheus.Desc
	instanceLocalInterfaces       *prometheus.Desc
	instanceLocalInterfacesUp     *prometheus.Desc
	instanceIRBInterfaces         *prometheus.Desc
	instanceIRBInterfacesUp       *prometheus.Desc
	instanceProtectInterfaces     *prometheus.Desc
	instanceBridgeDomains         *prometheus.Desc
	instanceLocalMACs             *prometheus.Desc
	instanceRemoteMACs            *prometheus.Desc
	instanceLocalMACIPs           *prometheus.Desc
	instanceRemoteMACIPs          *prometheus.Desc
	instanceLocalDefaultGwMACs    *prometheus.Desc
	instanceRemoteDefaultGwMACs   *prometheus.Desc
	instanceDuplicateMACThreshold *prometheus.Desc
	instanceDuplicateMACWindow    *prometheus.Desc

	// ============== Per-neighbor descriptors (from show evpn instance) ==============
	neighborMACRoutes                *prometheus.Desc
	neighborMACIPRoutes              *prometheus.Desc
	neighborAutoDiscoveryRoutes      *prometheus.Desc
	neighborInclusiveMulticastRoutes *prometheus.Desc
	neighborEthernetSegmentRoutes    *prometheus.Desc

	// ============== Detail descriptors (from show evpn instance) ==============
	interfaceStatus              *prometheus.Desc
	irbStatus                    *prometheus.Desc
	bridgeDomainInterfaceCount   *prometheus.Desc
	bridgeDomainInterfaceUpCount *prometheus.Desc
	esiResolved                  *prometheus.Desc
	esiRemotePECount             *prometheus.Desc
	esiDesignatedForwarderInfo   *prometheus.Desc

	// ============== Duplicate-MAC descriptors (from show evpn database state duplicate) ==============
	duplicateMACTotal *prometheus.Desc
	duplicateMACCount *prometheus.Desc

	// ============== L3 context descriptors (from show evpn l3-context) ==============
	l3ContextVNI   *prometheus.Desc
	l3ContextCount *prometheus.Desc
)

func init() {
	il := []string{"target", "instance"}
	infoLabels := []string{"target", "instance", "rd", "encap", "router_id", "source_vtep"}
	neighborLabels := []string{"target", "instance", "neighbor"}
	interfaceLabels := []string{"target", "instance", "interface", "esi", "mode", "etree_role"}
	irbLabels := []string{"target", "instance", "irb_interface", "vni_id", "l3_context"}
	bdLabels := []string{"target", "instance", "vlan_id", "domain_id", "irb_interface", "mode", "mac_sync"}
	esiCountLabels := []string{"target", "instance", "esi"}
	esiDFLabels := []string{"target", "instance", "esi", "designated_forwarder", "backup_forwarder", "df_algorithm", "local_interface"}
	l3CtxLabels := []string{"target", "context", "type", "advertisement_mode", "router_mac", "encapsulation"}
	stateLabels := []string{"target"}

	instanceInfo = prometheus.NewDesc(prefix+"instance_info",
		"Per-EVI metadata (always 1). Labels carry route-distinguisher, encapsulation, router-id and source VTEP address.",
		infoLabels, nil)
	instanceNeighborCount = prometheus.NewDesc(prefix+"instance_neighbor_count",
		"Number of EVPN neighbors (PEs) for this instance", il, nil)
	instanceESICount = prometheus.NewDesc(prefix+"instance_esi_count",
		"Number of Ethernet Segment Identifiers known by this instance", il, nil)
	instanceLocalInterfaces = prometheus.NewDesc(prefix+"instance_local_interfaces",
		"Number of local interfaces in this EVPN instance", il, nil)
	instanceLocalInterfacesUp = prometheus.NewDesc(prefix+"instance_local_interfaces_up",
		"Number of local interfaces currently up in this EVPN instance", il, nil)
	instanceIRBInterfaces = prometheus.NewDesc(prefix+"instance_irb_interfaces",
		"Number of IRB interfaces in this EVPN instance", il, nil)
	instanceIRBInterfacesUp = prometheus.NewDesc(prefix+"instance_irb_interfaces_up",
		"Number of IRB interfaces currently up in this EVPN instance", il, nil)
	instanceProtectInterfaces = prometheus.NewDesc(prefix+"instance_protect_interfaces",
		"Number of protect (backup) interfaces in this EVPN instance", il, nil)
	instanceBridgeDomains = prometheus.NewDesc(prefix+"instance_bridge_domains",
		"Number of bridge domains in this EVPN instance", il, nil)
	instanceLocalMACs = prometheus.NewDesc(prefix+"instance_local_mac_count",
		"Number of MACs learned locally in this EVPN instance", il, nil)
	instanceRemoteMACs = prometheus.NewDesc(prefix+"instance_remote_mac_count",
		"Number of MACs learned from remote PEs in this EVPN instance", il, nil)
	instanceLocalMACIPs = prometheus.NewDesc(prefix+"instance_local_mac_ip_count",
		"Number of local MAC+IP bindings in this EVPN instance", il, nil)
	instanceRemoteMACIPs = prometheus.NewDesc(prefix+"instance_remote_mac_ip_count",
		"Number of remote MAC+IP bindings in this EVPN instance", il, nil)
	instanceLocalDefaultGwMACs = prometheus.NewDesc(prefix+"instance_local_default_gateway_mac_count",
		"Number of local default-gateway MACs in this EVPN instance", il, nil)
	instanceRemoteDefaultGwMACs = prometheus.NewDesc(prefix+"instance_remote_default_gateway_mac_count",
		"Number of remote default-gateway MACs in this EVPN instance", il, nil)
	instanceDuplicateMACThreshold = prometheus.NewDesc(prefix+"instance_duplicate_mac_threshold",
		"Configured duplicate-MAC detection threshold for this EVPN instance", il, nil)
	instanceDuplicateMACWindow = prometheus.NewDesc(prefix+"instance_duplicate_mac_window_seconds",
		"Configured duplicate-MAC detection window in seconds for this EVPN instance", il, nil)

	neighborMACRoutes = prometheus.NewDesc(prefix+"neighbor_mac_routes",
		"Per-neighbor count of EVPN Type-2 MAC routes (without IP)", neighborLabels, nil)
	neighborMACIPRoutes = prometheus.NewDesc(prefix+"neighbor_mac_ip_routes",
		"Per-neighbor count of EVPN Type-2 MAC+IP routes", neighborLabels, nil)
	neighborAutoDiscoveryRoutes = prometheus.NewDesc(prefix+"neighbor_ethernet_autodiscovery_routes",
		"Per-neighbor count of EVPN Type-1 Ethernet auto-discovery routes", neighborLabels, nil)
	neighborInclusiveMulticastRoutes = prometheus.NewDesc(prefix+"neighbor_inclusive_multicast_routes",
		"Per-neighbor count of EVPN Type-3 inclusive-multicast routes", neighborLabels, nil)
	neighborEthernetSegmentRoutes = prometheus.NewDesc(prefix+"neighbor_ethernet_segment_routes",
		"Per-neighbor count of EVPN Type-4 Ethernet-segment routes", neighborLabels, nil)

	// Detail (Phase A): per-interface, per-IRB, per-bridge-domain, per-ESI.
	interfaceStatus = prometheus.NewDesc(prefix+"interface_status",
		"EVPN interface status (0: down, 1: up). Labels carry the ESI, mode (single-homed / all-active / active-standby) and E-tree role.",
		interfaceLabels, nil)
	irbStatus = prometheus.NewDesc(prefix+"irb_status",
		"IRB interface status within an EVPN instance (0: down, 1: up). Labels carry the IRB VNI and L3 context.",
		irbLabels, nil)
	bridgeDomainInterfaceCount = prometheus.NewDesc(prefix+"bridge_domain_interface_count",
		"Total interfaces in this bridge-domain", bdLabels, nil)
	bridgeDomainInterfaceUpCount = prometheus.NewDesc(prefix+"bridge_domain_interface_up_count",
		"Interfaces currently up in this bridge-domain", bdLabels, nil)
	esiResolved = prometheus.NewDesc(prefix+"esi_resolved",
		"EVPN ESI resolution state (0: unresolved, 1: resolved). Join with junos_evpn_esi_designated_forwarder_info on (target, instance, esi) for DF/local-interface labels.",
		esiCountLabels, nil)
	esiRemotePECount = prometheus.NewDesc(prefix+"esi_remote_pe_count",
		"Number of remote PEs known for this Ethernet Segment", esiCountLabels, nil)
	esiDesignatedForwarderInfo = prometheus.NewDesc(prefix+"esi_designated_forwarder_info",
		"EVPN ESI designated-forwarder election state (gauge=1 info-pattern). Labels churn on DF election events; join with junos_evpn_esi_resolved for clean state queries.",
		esiDFLabels, nil)

	// Duplicate-MAC: target-level total always emitted, per-instance only when > 0.
	duplicateMACTotal = prometheus.NewDesc(prefix+"duplicate_mac_total",
		"Total MAC entries currently suppressed by duplicate-MAC detection across all EVIs on this device. Non-zero indicates a forwarding loop or split-brain.",
		stateLabels, nil)
	duplicateMACCount = prometheus.NewDesc(prefix+"duplicate_mac_count",
		"MAC entries currently suppressed by duplicate-MAC detection in a specific EVPN instance. Emitted only when count > 0.",
		il, nil)

	// L3 context.
	l3ContextVNI = prometheus.NewDesc(prefix+"l3_context_vni",
		"EVPN L3 context (VRF) VNI. Labels carry context type, advertisement mode, router MAC and encapsulation.",
		l3CtxLabels, nil)
	l3ContextCount = prometheus.NewDesc(prefix+"l3_context_count",
		"Number of EVPN L3 contexts configured on this device", []string{"target"}, nil)
}

type evpnCollector struct{}

// Name returns the name of the collector.
func (*evpnCollector) Name() string { return "evpn" }

// NewCollector creates a new collector.
func NewCollector() collector.RPCCollector { return new(evpnCollector) }

// Describe describes the metrics.
func (*evpnCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, d := range []*prometheus.Desc{
		instanceInfo,
		instanceNeighborCount, instanceESICount,
		instanceLocalInterfaces, instanceLocalInterfacesUp,
		instanceIRBInterfaces, instanceIRBInterfacesUp,
		instanceProtectInterfaces, instanceBridgeDomains,
		instanceLocalMACs, instanceRemoteMACs,
		instanceLocalMACIPs, instanceRemoteMACIPs,
		instanceLocalDefaultGwMACs, instanceRemoteDefaultGwMACs,
		instanceDuplicateMACThreshold, instanceDuplicateMACWindow,
		neighborMACRoutes, neighborMACIPRoutes,
		neighborAutoDiscoveryRoutes, neighborInclusiveMulticastRoutes,
		neighborEthernetSegmentRoutes,
		interfaceStatus, irbStatus,
		bridgeDomainInterfaceCount, bridgeDomainInterfaceUpCount,
		esiResolved, esiRemotePECount, esiDesignatedForwarderInfo,
		duplicateMACTotal, duplicateMACCount,
		l3ContextVNI, l3ContextCount,
	} {
		ch <- d
	}
}

// Collect issues three RPCs:
//
//  1. show evpn instance extensive — REQUIRED. Provides per-EVI state,
//     per-neighbor route counts, and (via the detail tables) per-interface,
//     per-IRB, per-bridge-domain, and per-ESI status. A failure here aborts
//     the scrape because the rest of the collector depends on knowing which
//     EVIs exist.
//
//  2. show evpn database state duplicate — BEST EFFORT. Returns MAC entries
//     suppressed by duplicate-MAC detection (the only true EVPN error
//     signal). A failure here logs a warning and leaves the metrics absent.
//
//  3. show evpn l3-context — BEST EFFORT. Returns L3 (IRB) context info.
//     Empty when EVPN-IRB is not configured.
func (c *evpnCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	// --- RPC 1: instance (required) ---
	var instances []evpnInstance
	if err := client.RunCommandAndParseWithParser("show evpn instance extensive", func(b []byte) error {
		var perr error
		instances, perr = parseInstances(b)
		return perr
	}); err != nil {
		return err
	}

	for _, in := range instances {
		c.emitInstance(ch, labelValues, in)
	}

	// --- RPC 2: duplicate-MAC (best effort) ---
	var dupInstances []duplicateInstance
	if err := client.RunCommandAndParseWithParser("show evpn database state duplicate", func(b []byte) error {
		var perr error
		dupInstances, perr = parseDuplicateMACs(b)
		return perr
	}); err != nil {
		log.Warnf("evpn: 'show evpn database state duplicate' failed: %v", err)
	} else {
		c.emitDuplicateMACs(ch, labelValues, dupInstances)
	}

	// --- RPC 3: l3-context (best effort) ---
	var contexts []l3Context
	if err := client.RunCommandAndParseWithParser("show evpn l3-context", func(b []byte) error {
		var perr error
		contexts, perr = parseL3Contexts(b)
		return perr
	}); err != nil {
		log.Warnf("evpn: 'show evpn l3-context' failed: %v", err)
	} else {
		c.emitL3Contexts(ch, labelValues, contexts)
	}

	return nil
}

// emitInstance emits all per-EVI and per-neighbor metrics derived from one
// <evpn-instance> entry, plus the detail tables (interfaces, IRBs,
// bridge-domains, ESIs) when present.
func (c *evpnCollector) emitInstance(ch chan<- prometheus.Metric, labelValues []string, in evpnInstance) {
	base := append(labelValues, in.Name)

	ch <- prometheus.MustNewConstMetric(
		instanceInfo, prometheus.GaugeValue, 1,
		append(append([]string{}, base...), in.RouteDistinguisher, strings.TrimSpace(in.EncapType), in.RouterID, in.SourceVTEPAddr)...,
	)

	ch <- prometheus.MustNewConstMetric(instanceNeighborCount, prometheus.GaugeValue, in.NumNeighbors, base...)
	ch <- prometheus.MustNewConstMetric(instanceESICount, prometheus.GaugeValue, in.NumESI, base...)
	ch <- prometheus.MustNewConstMetric(instanceLocalInterfaces, prometheus.GaugeValue, in.LocalInterfaces, base...)
	ch <- prometheus.MustNewConstMetric(instanceLocalInterfacesUp, prometheus.GaugeValue, in.LocalInterfacesUp, base...)
	ch <- prometheus.MustNewConstMetric(instanceIRBInterfaces, prometheus.GaugeValue, in.IRBInterfaces, base...)
	ch <- prometheus.MustNewConstMetric(instanceIRBInterfacesUp, prometheus.GaugeValue, in.IRBInterfacesUp, base...)
	ch <- prometheus.MustNewConstMetric(instanceProtectInterfaces, prometheus.GaugeValue, in.NumProtectInterfaces, base...)
	ch <- prometheus.MustNewConstMetric(instanceBridgeDomains, prometheus.GaugeValue, in.NumBridgeDomains, base...)
	ch <- prometheus.MustNewConstMetric(instanceDuplicateMACThreshold, prometheus.GaugeValue, in.DuplicateMACDetectionThreshold, base...)
	ch <- prometheus.MustNewConstMetric(instanceDuplicateMACWindow, prometheus.GaugeValue, in.DuplicateMACDetectionWindow, base...)

	if in.MACDatabase != nil {
		db := in.MACDatabase
		ch <- prometheus.MustNewConstMetric(instanceLocalMACs, prometheus.GaugeValue, db.LocalMACs, base...)
		ch <- prometheus.MustNewConstMetric(instanceRemoteMACs, prometheus.GaugeValue, db.RemoteMACs, base...)
		ch <- prometheus.MustNewConstMetric(instanceLocalMACIPs, prometheus.GaugeValue, db.LocalMACIPs, base...)
		ch <- prometheus.MustNewConstMetric(instanceRemoteMACIPs, prometheus.GaugeValue, db.RemoteMACIPs, base...)
		ch <- prometheus.MustNewConstMetric(instanceLocalDefaultGwMACs, prometheus.GaugeValue, db.LocalDefaultGatewayMACs, base...)
		ch <- prometheus.MustNewConstMetric(instanceRemoteDefaultGwMACs, prometheus.GaugeValue, db.RemoteDefaultGatewayMACs, base...)
	}

	for _, n := range in.Neighbors {
		r := n.Routes
		if r.Address == "" {
			continue
		}
		nl := append(append([]string{}, base...), r.Address)
		ch <- prometheus.MustNewConstMetric(neighborMACRoutes, prometheus.GaugeValue, r.NumMACRoutes, nl...)
		ch <- prometheus.MustNewConstMetric(neighborMACIPRoutes, prometheus.GaugeValue, r.NumMACIPRoutes, nl...)
		ch <- prometheus.MustNewConstMetric(neighborAutoDiscoveryRoutes, prometheus.GaugeValue, r.NumAutoDiscoveryRoutes, nl...)
		ch <- prometheus.MustNewConstMetric(neighborInclusiveMulticastRoutes, prometheus.GaugeValue, r.NumInclusiveMulticastRoutes, nl...)
		ch <- prometheus.MustNewConstMetric(neighborEthernetSegmentRoutes, prometheus.GaugeValue, r.NumEthernetSegmentRoutes, nl...)
	}

	// --- detail tables ---
	for _, ifce := range in.Interfaces {
		l := append(append([]string{}, base...), ifce.Name, ifce.ESI, ifce.Mode, ifce.EtreeRole)
		ch <- prometheus.MustNewConstMetric(interfaceStatus, prometheus.GaugeValue, upDown(ifce.Status), l...)
	}
	for _, irb := range in.IRBInterfaceTbl {
		l := append(append([]string{}, base...), irb.Name, irb.VNI, irb.L3Context)
		ch <- prometheus.MustNewConstMetric(irbStatus, prometheus.GaugeValue, upDown(irb.Status), l...)
	}
	for _, bd := range in.BridgeDomains {
		l := append(append([]string{}, base...), bd.VLANID, bd.DomainID, bd.IRBInterface, bd.Mode, bd.MACSyncStatus)
		ch <- prometheus.MustNewConstMetric(bridgeDomainInterfaceCount, prometheus.GaugeValue, bd.Interfaces, l...)
		ch <- prometheus.MustNewConstMetric(bridgeDomainInterfaceUpCount, prometheus.GaugeValue, bd.InterfacesUp, l...)
	}
	for _, esi := range in.ESIs {
		// _resolved is a pure-state gauge with stable labels — operators
		// alert on this directly without worrying about DF-election churn.
		resolved := 0.0
		if strings.Contains(strings.ToLower(esi.Status), "resolved") {
			resolved = 1.0
		}
		esiLabels := append(append([]string{}, base...), esi.Value)
		ch <- prometheus.MustNewConstMetric(esiResolved, prometheus.GaugeValue, resolved, esiLabels...)

		if esi.RemotePEInfo != nil {
			ch <- prometheus.MustNewConstMetric(esiRemotePECount, prometheus.GaugeValue, esi.RemotePEInfo.Count, esiLabels...)
		}

		// _designated_forwarder_info is an info-pattern series carrying the
		// labels that *do* change on DF-election events (DF/backup IPs,
		// local-interface bind, algorithm). It deliberately churns when those
		// rotate — operators query `changes(... [window])` on this series to
		// detect DF instability, and `group_left(...)` it onto _resolved for
		// dashboards. Emitted only when at least one of the labels is
		// populated, to avoid noise on remote ESIs with no local state.
		localIfName, dfAddr, backupAddr, dfAlgo := "", "", "", ""
		if esi.LocalIntf != nil {
			localIfName = esi.LocalIntf.Name
		}
		if esi.DFInfo != nil {
			dfAddr = esi.DFInfo.DesignatedForwarder
			backupAddr = esi.DFInfo.BackupForwarder
			dfAlgo = esi.DFInfo.Algorithm
		}
		if dfAddr != "" || backupAddr != "" || dfAlgo != "" || localIfName != "" {
			dfLabels := append(append([]string{}, base...), esi.Value, dfAddr, backupAddr, dfAlgo, localIfName)
			ch <- prometheus.MustNewConstMetric(esiDesignatedForwarderInfo, prometheus.GaugeValue, 1, dfLabels...)
		}
	}
}

// emitDuplicateMACs emits the target-level total (always) and per-instance
// breakdown (only when count > 0).
func (c *evpnCollector) emitDuplicateMACs(ch chan<- prometheus.Metric, labelValues []string, dupInstances []duplicateInstance) {
	total := 0
	for _, in := range dupInstances {
		n := len(in.Entries)
		if n == 0 {
			continue
		}
		total += n
		il := append(append([]string{}, labelValues...), in.Name)
		ch <- prometheus.MustNewConstMetric(duplicateMACCount, prometheus.GaugeValue, float64(n), il...)
	}
	ch <- prometheus.MustNewConstMetric(duplicateMACTotal, prometheus.GaugeValue, float64(total), labelValues...)
}

// emitL3Contexts emits one _vni series per context plus a target-level _count.
func (c *evpnCollector) emitL3Contexts(ch chan<- prometheus.Metric, labelValues []string, contexts []l3Context) {
	for _, ctx := range contexts {
		l := append(append([]string{}, labelValues...), ctx.Name, ctx.Type, ctx.AdvertisementMode, ctx.RouterMAC, ctx.Encapsulation)
		ch <- prometheus.MustNewConstMetric(l3ContextVNI, prometheus.GaugeValue, ctx.VNI, l...)
	}
	ch <- prometheus.MustNewConstMetric(l3ContextCount, prometheus.GaugeValue, float64(len(contexts)), labelValues...)
}

// upDown converts "Up" / "Down" to 1.0 / 0.0 (case-insensitive).
func upDown(s string) float64 {
	if strings.EqualFold(strings.TrimSpace(s), "Up") {
		return 1.0
	}
	return 0.0
}

// parseInstances handles single-RE and multi-RE response shapes for the
// 'show evpn instance' RPC. Merging the per-RE slices; re-name is not
// exposed as a label (matches alarm/ufd precedent).
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

// parseDuplicateMACs handles both response shapes for
// 'show evpn database state duplicate'.
func parseDuplicateMACs(b []byte) ([]duplicateInstance, error) {
	if strings.Contains(string(b), "<multi-routing-engine-results") {
		var m duplicateMultiResult
		if err := xml.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		var out []duplicateInstance
		for _, re := range m.Engines.Items {
			out = append(out, re.Instances...)
		}
		return out, nil
	}
	var s duplicateSingleResult
	if err := xml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s.Instances, nil
}

// parseL3Contexts handles both response shapes for 'show evpn l3-context'.
func parseL3Contexts(b []byte) ([]l3Context, error) {
	if strings.Contains(string(b), "<multi-routing-engine-results") {
		var m l3ContextMultiResult
		if err := xml.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		var out []l3Context
		for _, re := range m.Engines.Items {
			out = append(out, re.Contexts...)
		}
		return out, nil
	}
	var s l3ContextSingleResult
	if err := xml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s.Contexts, nil
}
