// SPDX-License-Identifier: MIT

package evpnipprefix

import "encoding/xml"

// Response shape for `show evpn ip-prefix-database`:
//
//   <rpc-reply>
//     <evpn-ip-prefix-database-information>
//       <evpn-pfxdb-l3-context>                  ← repeats per L3 context
//         <context-name>vrf-a</context-name>
//
//         <evpn-pfxdb-ip-table>                  ← appears once per direction
//           <table-description>IPv4->EVPN | IPv6->EVPN</table-description>
//           <evpn-pfxdb-ip-entry>                ← repeats per local prefix
//             <entry-prefix>...</entry-prefix>
//             <entry-evpn-route-status>...</entry-evpn-route-status>
//           </evpn-pfxdb-ip-entry>
//         </evpn-pfxdb-ip-table>
//
//         <evpn-pfxdb-evpn-ip-table>             ← appears once per direction
//           <table-description>EVPN->IPv4 | EVPN->IPv6</table-description>
//           <evpn-pfxdb-evpn-ip-entry>           ← repeats per remote prefix
//             <entry-prefix>...</entry-prefix>
//             <entry-etag>0</entry-etag>
//             <evpn-pfxdb-evpn-ip-adv>           ← repeats per advertising PE
//               <route-distinguisher>...</route-distinguisher>
//               <adv-vni>...</adv-vni>
//               <adv-router-mac>...</adv-router-mac>
//               <adv-bgp-nexthop>...</adv-bgp-nexthop>
//               <adv-ip-route-status>Accepted|Rejected</adv-ip-route-status>
//               <adv-ip-route-error>n/a|...</adv-ip-route-error>
//             </evpn-pfxdb-evpn-ip-adv>
//           </evpn-pfxdb-evpn-ip-entry>
//         </evpn-pfxdb-evpn-ip-table>
//       </evpn-pfxdb-l3-context>
//     </evpn-ip-prefix-database-information>
//   </rpc-reply>
//
// Each L3 context emits exactly four <evpn-pfxdb-*-table> entries (two
// local-origin, two remote-received) distinguished by <table-description>.

type singleEngineResult struct {
	XMLName  xml.Name      `xml:"rpc-reply"`
	Contexts []pfxL3Context `xml:"evpn-ip-prefix-database-information>evpn-pfxdb-l3-context"`
}

type multiEngineResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Engines struct {
		Items []routingEngine `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
}

type routingEngine struct {
	Name     string         `xml:"re-name"`
	Contexts []pfxL3Context `xml:"evpn-ip-prefix-database-information>evpn-pfxdb-l3-context"`
}

type pfxL3Context struct {
	Name           string             `xml:"context-name"`
	LocalTables    []pfxLocalTable    `xml:"evpn-pfxdb-ip-table"`
	RemoteTables   []pfxRemoteTable   `xml:"evpn-pfxdb-evpn-ip-table"`
}

type pfxLocalTable struct {
	Description string         `xml:"table-description"`
	Entries     []pfxLocalEntry `xml:"evpn-pfxdb-ip-entry"`
}

type pfxLocalEntry struct {
	Prefix string `xml:"entry-prefix"`
	Status string `xml:"entry-evpn-route-status"`
}

type pfxRemoteTable struct {
	Description string           `xml:"table-description"`
	Entries     []pfxRemoteEntry `xml:"evpn-pfxdb-evpn-ip-entry"`
}

type pfxRemoteEntry struct {
	Prefix         string                 `xml:"entry-prefix"`
	ETag           string                 `xml:"entry-etag"`
	Advertisements []pfxRemoteAdvertisement `xml:"evpn-pfxdb-evpn-ip-adv"`
}

type pfxRemoteAdvertisement struct {
	RouteDistinguisher string `xml:"route-distinguisher"`
	VNI                string `xml:"adv-vni"`
	RouterMAC          string `xml:"adv-router-mac"`
	BGPNexthop         string `xml:"adv-bgp-nexthop"`
	Status             string `xml:"adv-ip-route-status"`
	Error              string `xml:"adv-ip-route-error"`
}
