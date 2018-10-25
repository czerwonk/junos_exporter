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
	Stats             TrafficStat    `xml:"traffic-statistics"`
	LogicalInterfaces []LogInterface `xml:"logical-interface"`
	InputErrors       struct {
		Drops  int64 `xml:"input-drops"`
		Errors int64 `xml:"input-errors"`
	} `xml:"input-error-list"`
	OutputErrors struct {
		Drops  int64 `xml:"output-drops"`
		Errors int64 `xml:"output-errors"`
	} `xml:"output-error-list"`
}

type LogInterface struct {
	Name        string      `xml:"name"`
	Description string      `xml:"description"`
	Stats       TrafficStat `xml:"traffic-statistics"`
}

type TrafficStat struct {
	InputBytes    int64    `xml:"input-bytes"`
	InputPackets  int64    `xml:"input-packets"`
	OutputBytes   int64    `xml:"output-bytes"`
	OutputPackets int64    `xml:"output-packets"`
	IPv6Traffic   IPv6Stat `xml:"ipv6-transit-statistics"`
}

type IPv6Stat struct {
	InputBytes    int64 `xml:"input-bytes"`
	InputPackets  int64 `xml:"input-packets"`
	OutputBytes   int64 `xml:"output-bytes"`
	OutputPackets int64 `xml:"output-packets"`
}
