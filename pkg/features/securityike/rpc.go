// SPDX-License-Identifier: MIT

package securityike

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name       `xml:"rpc-reply"`
	Results routingEngines `xml:"multi-routing-engine-results"`
}

type routingEngines struct {
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name                      string                    `xml:"re-name"`
	IKEActivePeersInformation ikeActivePeersInformation `xml:"ike-active-peers-information"`
}

type ikeActivePeersInformation struct {
	IKEActivePeers []ikeActivePeer `xml:"ike-active-peers"`
}

type ikeActivePeer struct {
	IKESARemoteAddress     string `xml:"ike-sa-remote-address"`
	IKESARemotePort        int    `xml:"ike-sa-remote-port"`
	IKEIKEID               string `xml:"ike-ike-id"`
	IKEXAuthUsername       string `xml:"ike-xauth-username"`
	IKEXAuthUserAssignedIP string `xml:"ike-xauth-user-assigned-ip"`
}

type singleEngineResult struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	IKEActivePeersInformation ikeActivePeersInformation `xml:"ike-active-peers-information"`
}
