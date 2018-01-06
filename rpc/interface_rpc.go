package rpc

type InterfaceRpc struct {
	Information struct {
		Interfaces []PhyInterface `xml:"physical-interface"`
  } `xml:"interface-information"`
}

type PhyInterface struct {
	Name string `xml:"name"`
	Description string `xml:"description"`
	MacAddress string `xml:"current-physical-address"`
	Stats TrafficStat `xml:"traffic-statistics"`
	LogicalInterfaces []LogInterface `xml:"logical-interface"`
}

type LogInterface struct {
	Name string `xml:"name"`
	Description string `xml:"description"`
	Stats TrafficStat `xml:"traffic-statistics"`
}

type TrafficStat struct {
	InputBytes int64 `xml:"input-bytes"`
	OutputBytes int64 `xml:"output-bytes"`
}