package interfaces

type InterfaceRpc struct {
	Information struct {
		Interfaces []PhyInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type PhyInterface struct {
	Name              string         `xml:"name"`
	AdminStatus       string         `xml:"admin-status"`
	OperStatus        string         `xml:"oper-status"`
	Description       string         `xml:"description"`
	MacAddress        string         `xml:"current-physical-address"`
	Speed             string         `xml:"speed"`
	Stats             TrafficStat    `xml:"traffic-statistics"`
	LogicalInterfaces []LogInterface `xml:"logical-interface"`
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
	EthernetMacStatistics EthernetMacStat `xml:"ethernet-mac-statistics"`
	EthernetFecStatistics EthernetFecStat `xml:"ethernet-fec-statistics"`
}

type LogInterface struct {
	Name        string         `xml:"name"`
	Description string         `xml:"description"`
	Stats       TrafficStat    `xml:"traffic-statistics"`
	LagStats    LagTrafficStat `xml:"lag-traffic-statistics"`
}

type TrafficStat struct {
	InputBytes    uint64   `xml:"input-bytes"`
	InputPackets  uint64   `xml:"input-packets"`
	OutputBytes   uint64   `xml:"output-bytes"`
	OutputPackets uint64   `xml:"output-packets"`
	IPv6Traffic   IPv6Stat `xml:"ipv6-transit-statistics"`
}

type IPv6Stat struct {
	InputBytes    uint64 `xml:"input-bytes"`
	InputPackets  uint64 `xml:"input-packets"`
	OutputBytes   uint64 `xml:"output-bytes"`
	OutputPackets uint64 `xml:"output-packets"`
}

type LagTrafficStat struct {
	Stats TrafficStat `xml:"lag-bundle"`
	Links []struct {
		Name string `xml:"name"`
	} `xml:"lag-link"`
}

type EthernetMacStat struct {
	InputUnicasts         uint64 `xml:"input-unicasts"`
	InputBroadcasts       uint64 `xml:"input-broadcasts"`
	InputMulticasts       uint64 `xml:"input-multicasts"`
	InputCrcErrors        uint64 `xml:"input-crc-errors"`
	OutputUnicasts        uint64 `xml:"output-unicasts"`
	OutputBroadcasts      uint64 `xml:"output-broadcasts"`
	OutputMulticasts      uint64 `xml:"output-multicasts"`
	OutputCrcErrors       uint64 `xml:"output-crc-errors"`
	InputOversizedFrames  uint64 `xml:"input-oversized-frames"`
	InputJabberFrames     uint64 `xml:"input-jabber-frames"`
	InputFragmentFrames   uint64 `xml:"input-fragment-frames"`
	InputVlanTaggedFrames uint64 `xml:"input-vlan-tagged-frames"`
	InputCodeViolations   uint64 `xml:"input-code-violations"`
	InputTotalErrors      uint64 `xml:"input-total-errors"`
	OutputTotalErrors     uint64 `xml:"output-total-errors"`
}

type EthernetFecStat struct {
	NumberfecCcwCount      uint64 `xml:"fec_ccw_count"`
	NumberfecNccwCount     uint64 `xml:"fec_nccw_count"`
	NumberfecCcwErrorRate  uint64 `xml:"fec_ccw_error_rate"`
	NumberfecNccwErrorRate uint64 `xml:"fec_nccw_error_rate"`
}
