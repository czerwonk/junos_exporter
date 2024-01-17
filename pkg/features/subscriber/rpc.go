package subscriber

type subcsribers_information struct {
	SubscribersInformation struct {
		Subscriber []Subscriber `xml:"subscriber"`
	} `xml:"subscribers-information"`
}

type Subscriber struct {
	AccessType          string `xml:"access-type"`
	Interface           string `xml:"interface"`
	AgentCircuitId      string `xml:"agent-circuit-id"`
	AgentRemoteId       string `xml:"agent-remote-id"`
	UnderlyingInterface string `xml:"underlying-interface"`
}

type InterfaceInformation struct {
	LogicalInterfaces []LogicalInterface `xml:"interface-information>physical-interface>logical-interface"`
}

type LogicalInterface struct {
	Name                  string `xml:"name"`
	DemuxUnderlyingIfName string `xml:"demux-information>demux-interface>demux-underlying-interface-name"`
}
