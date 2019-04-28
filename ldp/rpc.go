package ldp

type LDPRpc struct {
	Information ldpInformation `xml:"ldp-neighbor-information"`
}

type ldpInformation struct {
	Neighbors []ldpNeighbor `xml:"ldp-neighbor"`
}

type ldpNeighbor struct {
	Address string `xml:"ldp-neighbor-address"`
}

type LDPSessionRpc struct {
	Information ldpSessionInformation `xml:"ldp-session-information"`
}

type ldpSessionInformation struct {
	Sessions []ldpSession `xml:"ldp-session"`
}

type ldpSession struct {
	NeighborAddress string `xml:"ldp-neighbor-address"`
	State           string `xml:"ldp-connection-state"`
}
