package subscriber

type subcsribers_information struct {
	SubscribersInformation struct {
		Subscriber []subscriber `xml:"subscriber"`
	} `xml:"subscribers-information"`
}

type subscriber struct {
	AccessType     string `xml:"access-type"`
	Interface      string `xml:"interface"`
	AgentCircuitId string `xml:"agent-circuit-id"`
	AgentRemoteId  string `xml:"agent-remote-id"`
}
