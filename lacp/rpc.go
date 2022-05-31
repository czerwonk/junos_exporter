package lacp

type lacpRpc struct {
	Information struct {
		LacpInterfaces []lacpInterface `xml:"lacp-interface-information"`
	} `xml:"lacp-interface-information-list"`
}

type lacpInterface struct {
	LagLacpHeader struct {
		Name string `xml:"aggregate-name"`
	} `xml:"lag-lacp-header"`
	LagLacpStates    []LagLacpStateStruct    `xml:"lag-lacp-state"`
	LagLacpProtocols []LagLacpProtocolStruct `xml:"lag-lacp-protocol"`
}

type LagLacpStateStruct struct {
}

type LagLacpProtocolStruct struct {
	Member       string `xml:"name"`
	LacpMuxState string `xml:"lacp-mux-state"`
}
