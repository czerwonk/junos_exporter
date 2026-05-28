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
// mac-database-status-table, and the per-interface tables.
type evpnInstance struct {
	Name                            string         `xml:"evpn-instance-name"`
	RouteDistinguisher              string         `xml:"route-distinguisher"`
	EncapType                       string         `xml:"evpn-encap-type"`
	ControlWord                     string         `xml:"evpn-instance-control-word"`
	DuplicateMACDetectionThreshold  float64        `xml:"duplicate-mac-detection-threshold"`
	DuplicateMACDetectionWindow     float64        `xml:"duplicate-mac-detection-window"`
	MACDatabase                     *macDatabase   `xml:"mac-database-status-table"`
	LocalInterfaces                 float64        `xml:"local-interfaces"`
	LocalInterfacesUp               float64        `xml:"local-interfaces-up"`
	IRBInterfaces                   float64        `xml:"irb-interfaces"`
	IRBInterfacesUp                 float64        `xml:"irb-interfaces-up"`
	NumProtectInterfaces            float64        `xml:"num-protect-interfaces"`
	NumBridgeDomains                float64        `xml:"num-bridge-domains"`
	NumNeighbors                    float64        `xml:"evpn-num-neighbors"`
	Neighbors                       []evpnNeighbor `xml:"evpn-neighbor"`
	NumESI                          float64        `xml:"evpn-num-esi"`
	RouterID                        string         `xml:"evpn-router-id"`
	SourceVTEPAddr                  string         `xml:"evpn-source-vtep-ipaddr"`
	SMETForwarding                  string         `xml:"evpn-smet-forwarding"`
}

type macDatabase struct {
	LocalMACs               float64 `xml:"local-macs"`
	RemoteMACs              float64 `xml:"remote-macs"`
	LocalMACIPs             float64 `xml:"local-mac-ips"`
	RemoteMACIPs            float64 `xml:"remote-mac-ips"`
	LocalDefaultGatewayMACs float64 `xml:"local-default-gateway-macs"`
	RemoteDefaultGatewayMACs float64 `xml:"remote-default-gateway-macs"`
}

type evpnNeighbor struct {
	Routes evpnNeighborRoutes `xml:"evpn-neighbor-route-information"`
}

type evpnNeighborRoutes struct {
	Address                    string  `xml:"evpn-neighbor-address"`
	NumMACRoutes               float64 `xml:"evpn-num-mac-routes"`
	NumMACIPRoutes             float64 `xml:"evpn-num-mac-ip-routes"`
	NumAutoDiscoveryRoutes     float64 `xml:"evpn-num-ethernet-autodiscovery-routes"`
	NumInclusiveMulticastRoutes float64 `xml:"evpn-num-inclusive-multicast-routes"`
	NumEthernetSegmentRoutes   float64 `xml:"evpn-num-ethernet-segment-routes"`
}
