// SPDX-License-Identifier: MIT

package interfaces

type result struct {
	Information struct {
		Interfaces []phyInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type phyInterface struct {
	Name              string         `xml:"name"`
	AdminStatus       string         `xml:"admin-status"`
	OperStatus        string         `xml:"oper-status"`
	Description       string         `xml:"description"`
	MacAddress        string         `xml:"current-physical-address"`
	Speed             string         `xml:"speed"`
	BPDUError         string         `xml:"bpdu-error"`
	Stats             trafficStat    `xml:"traffic-statistics"`
	LogicalInterfaces []logInterface `xml:"logical-interface"`
	InputErrors       struct {
		Drops  uint64 `xml:"input-drops"`
		Errors uint64 `xml:"input-errors"`
	} `xml:"input-error-list"`
	OutputErrors struct {
		Drops  uint64 `xml:"output-drops"`
		Errors uint64 `xml:"output-errors"`
	} `xml:"output-error-list"`
	InterfaceFlapped struct {
		Seconds uint64 `xml:"seconds,attr"`
		Value   string `xml:",chardata"`
	} `xml:"interface-flapped"`
	MACStatistics ethernetMACStat `xml:"ethernet-mac-statistics"`
	FECStatistics ethernetFECStat `xml:"ethernet-fec-statistics"`
	MTU           string          `xml:"mtu"`
}

type logInterface struct {
	Name        string         `xml:"name"`
	Description string         `xml:"description"`
	Stats       trafficStat    `xml:"traffic-statistics"`
	LagStats    lagTrafficStat `xml:"lag-traffic-statistics"`
}

type trafficStat struct {
	InputBytes    uint64   `xml:"input-bytes"`
	InputPackets  uint64   `xml:"input-packets"`
	OutputBytes   uint64   `xml:"output-bytes"`
	OutputPackets uint64   `xml:"output-packets"`
	IPv6Traffic   ipv6Stat `xml:"ipv6-transit-statistics"`
}

type ipv6Stat struct {
	InputBytes    uint64 `xml:"input-bytes"`
	InputPackets  uint64 `xml:"input-packets"`
	OutputBytes   uint64 `xml:"output-bytes"`
	OutputPackets uint64 `xml:"output-packets"`
}

type lagTrafficStat struct {
	Stats trafficStat `xml:"lag-bundle"`
	Links []struct {
		Name string `xml:"name"`
	} `xml:"lag-link"`
}

type ethernetMACStat struct {
	InputUnicasts         uint64 `xml:"input-unicasts"`
	InputBroadcasts       uint64 `xml:"input-broadcasts"`
	InputMulticasts       uint64 `xml:"input-multicasts"`
	InputCRCErrors        uint64 `xml:"input-crc-errors"`
	OutputUnicasts        uint64 `xml:"output-unicasts"`
	OutputBroadcasts      uint64 `xml:"output-broadcasts"`
	OutputMulticasts      uint64 `xml:"output-multicasts"`
	OutputCRCErrors       uint64 `xml:"output-crc-errors"`
	InputOversizedFrames  uint64 `xml:"input-oversized-frames"`
	InputJabberFrames     uint64 `xml:"input-jabber-frames"`
	InputFragmentFrames   uint64 `xml:"input-fragment-frames"`
	InputVlanTaggedFrames uint64 `xml:"input-vlan-tagged-frames"`
	InputCodeViolations   uint64 `xml:"input-code-violations"`
	InputTotalErrors      uint64 `xml:"input-total-errors"`
	OutputTotalErrors     uint64 `xml:"output-total-errors"`
}

type ethernetFECStat struct {
	NumberfecCcwCount      uint64 `xml:"fec_ccw_count"`
	NumberfecNccwCount     uint64 `xml:"fec_nccw_count"`
	NumberfecCcwErrorRate  uint64 `xml:"fec_ccw_error_rate"`
	NumberfecNccwErrorRate uint64 `xml:"fec_nccw_error_rate"`
}
