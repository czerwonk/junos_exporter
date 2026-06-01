// SPDX-License-Identifier: MIT

package evpn

import "encoding/xml"

// Junos emits the EVPN instance response in two shapes. On a single-RE
// non-VC device:
//
//   <rpc-reply>
//     <evpn-instance-information>
//       <evpn-instance>...</evpn-instance>
//       <evpn-instance>...</evpn-instance>
//     </evpn-instance-information>
//   </rpc-reply>
//
// On a multi-RE chassis or Virtual Chassis the body is wrapped:
//
//   <rpc-reply>
//     <multi-routing-engine-results>
//       <multi-routing-engine-item>
//         <re-name>fpc0</re-name>
//         <evpn-instance-information>...</evpn-instance-information>
//       </multi-routing-engine-item>
//       ...
//     </multi-routing-engine-results>
//   </rpc-reply>
//
// Both shapes must be handled. Same dispatch pattern as the alarm and
// ufd collectors.

type singleEngineResult struct {
	XMLName   xml.Name       `xml:"rpc-reply"`
	Instances []evpnInstance `xml:"evpn-instance-information>evpn-instance"`
}

type multiEngineResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Engines struct {
		Items []routingEngine `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
}

type routingEngine struct {
	Name      string         `xml:"re-name"`
	Instances []evpnInstance `xml:"evpn-instance-information>evpn-instance"`
}

// evpnInstance models one <evpn-instance> entry. Fields are mostly
// optional - the default-evpn instance omits encapsulation, the
// mac-database-status-table, and the per-interface tables. The detail
// sub-tables (Interfaces, IRBInterfaces, BridgeDomains, ESIs) are
// populated when the device returns the extensive output and are only
// consumed by the collector when the EVPN feature is enabled.
type evpnInstance struct {
	Name                           string         `xml:"evpn-instance-name"`
	RouteDistinguisher             string         `xml:"route-distinguisher"`
	EncapType                      string         `xml:"evpn-encap-type"`
	DuplicateMACDetectionThreshold float64        `xml:"duplicate-mac-detection-threshold"`
	DuplicateMACDetectionWindow    float64        `xml:"duplicate-mac-detection-window"`
	MACDatabase                    *macDatabase   `xml:"mac-database-status-table"`
	LocalInterfaces                float64        `xml:"local-interfaces"`
	LocalInterfacesUp              float64        `xml:"local-interfaces-up"`
	IRBInterfaces                  float64        `xml:"irb-interfaces"`
	IRBInterfacesUp                float64        `xml:"irb-interfaces-up"`
	NumProtectInterfaces           float64        `xml:"num-protect-interfaces"`
	NumBridgeDomains               float64        `xml:"num-bridge-domains"`
	NumNeighbors                   float64        `xml:"evpn-num-neighbors"`
	Neighbors                      []evpnNeighbor `xml:"evpn-neighbor"`
	NumESI                         float64        `xml:"evpn-num-esi"`
	RouterID                       string         `xml:"evpn-router-id"`
	SourceVTEPAddr                 string         `xml:"evpn-source-vtep-ipaddr"`

	// Detail tables — populated by extensive output when supported by the device.
	// The collector parses these tables whenever they are present.
	Interfaces      []evpnInterface `xml:"evpn-interface-status-table>evpn-interface"`
	IRBInterfaceTbl []irbInterface  `xml:"irb-interface-status-table>irb-interface"`
	BridgeDomains   []bridgeDomain  `xml:"bridge-domain-status-table>bridge-domain"`
	ESIs            []evpnESI       `xml:"evpn-esi"`
}

// evpnInterface — entries inside <evpn-interface-status-table>.
type evpnInterface struct {
	Name      string `xml:"evpn-interface-name"`
	ESI       string `xml:"evpn-interface-esi"`
	Mode      string `xml:"evpn-interface-mode"`
	Status    string `xml:"evpn-interface-status"`
	EtreeRole string `xml:"evpn-interface-etree-role"`
}

// irbInterface — entries inside <irb-interface-status-table>.
type irbInterface struct {
	Name      string `xml:"irb-interface-name"`
	VNI       string `xml:"irb-interface-vni-id"`
	Status    string `xml:"irb-interface-status"`
	L3Context string `xml:"irb-interface-l3-context"`
}

// bridgeDomain — entries inside <bridge-domain-status-table>.
type bridgeDomain struct {
	VLANID        string  `xml:"vlan-id"`
	DomainID      string  `xml:"domain-id"`
	Interfaces    float64 `xml:"interfaces"`
	InterfacesUp  float64 `xml:"interfaces-up"`
	IRBInterface  string  `xml:"irb-interface"`
	Mode          string  `xml:"mode"`
	MACSyncStatus string  `xml:"mac-sync-status"`
}

// evpnESI — entries at the <evpn-esi> level. The nested DF-election
// and remote-PE blocks are flattened into labels.
type evpnESI struct {
	Value        string            `xml:"evpn-esi-value"`
	Status       string            `xml:"evpn-esi-status"`
	LocalIntf    *evpnESILocalIntf `xml:"evpn-esi-local-intf-information"`
	RemotePEInfo *evpnESIRemotePE  `xml:"evpn-esi-remote-pe-information"`
	DFInfo       *evpnESIDFInfo    `xml:"evpn-esi-df-information"`
}

type evpnESILocalIntf struct {
	Name   string `xml:"evpn-esi-local-intf-name"`
	Status string `xml:"evpn-esi-local-intf-status"`
}

type evpnESIRemotePE struct {
	Count float64 `xml:"evpn-esi-num-remote-pe"`
}

type evpnESIDFInfo struct {
	Algorithm           string `xml:"esi-df-election-algorithm"`
	DesignatedForwarder string `xml:"esi-designated-forwarder"`
	BackupForwarder     string `xml:"esi-backup-forwarder"`
}

// =================================================================
// show evpn database state duplicate
// =================================================================
//
// Returns the same <evpn-database-information> envelope as the
// unfiltered command, with only entries currently suppressed for
// duplicate-MAC detection. Empty envelope on healthy fabrics.

type duplicateSingleResult struct {
	XMLName   xml.Name            `xml:"rpc-reply"`
	Instances []duplicateInstance `xml:"evpn-database-information>evpn-database-instance"`
}

type duplicateMultiResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Engines struct {
		Items []duplicateRoutingEngine `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
}

type duplicateRoutingEngine struct {
	Name      string              `xml:"re-name"`
	Instances []duplicateInstance `xml:"evpn-database-information>evpn-database-instance"`
}

type duplicateInstance struct {
	Name    string              `xml:"instance-name"`
	Entries []duplicateMACEntry `xml:"mac-entry"`
}

type duplicateMACEntry struct {
	MACAddress string `xml:"mac-address"`
}

// =================================================================
// show evpn l3-context
// =================================================================

type l3ContextSingleResult struct {
	XMLName  xml.Name    `xml:"rpc-reply"`
	Contexts []l3Context `xml:"evpn-l3-context-information>evpn-l3-context"`
}

type l3ContextMultiResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Engines struct {
		Items []l3ContextRoutingEngine `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
}

type l3ContextRoutingEngine struct {
	Name     string      `xml:"re-name"`
	Contexts []l3Context `xml:"evpn-l3-context-information>evpn-l3-context"`
}

type l3Context struct {
	Name              string  `xml:"context-name"`
	Type              string  `xml:"context-type"`
	AdvertisementMode string  `xml:"context-advertisement-mode"`
	RouterMAC         string  `xml:"context-router-mac"`
	Encapsulation     string  `xml:"context-encapsulation"`
	VNI               float64 `xml:"context-vni"`
}

type macDatabase struct {
	LocalMACs                float64 `xml:"local-macs"`
	RemoteMACs               float64 `xml:"remote-macs"`
	LocalMACIPs              float64 `xml:"local-mac-ips"`
	RemoteMACIPs             float64 `xml:"remote-mac-ips"`
	LocalDefaultGatewayMACs  float64 `xml:"local-default-gateway-macs"`
	RemoteDefaultGatewayMACs float64 `xml:"remote-default-gateway-macs"`
}

type evpnNeighbor struct {
	Routes evpnNeighborRoutes `xml:"evpn-neighbor-route-information"`
}

type evpnNeighborRoutes struct {
	Address                     string  `xml:"evpn-neighbor-address"`
	NumMACRoutes                float64 `xml:"evpn-num-mac-routes"`
	NumMACIPRoutes              float64 `xml:"evpn-num-mac-ip-routes"`
	NumAutoDiscoveryRoutes      float64 `xml:"evpn-num-ethernet-autodiscovery-routes"`
	NumInclusiveMulticastRoutes float64 `xml:"evpn-num-inclusive-multicast-routes"`
	NumEthernetSegmentRoutes    float64 `xml:"evpn-num-ethernet-segment-routes"`
}
