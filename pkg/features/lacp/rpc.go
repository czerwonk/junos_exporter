package lacp

type result struct {
	Information struct {
		LacpInterfaces []lacpInterface `xml:"lacp-interface-information"`
	} `xml:"lacp-interface-information-list"`
}

type lacpInterface struct {
	LagLACPHeader struct {
		Name string `xml:"aggregate-name"`
	} `xml:"lag-lacp-header"`
	LagLACPProtocols []lagLACPProtocol `xml:"lag-lacp-protocol"`
}

type lagLACPProtocol struct {
	Member       string `xml:"name"`
	LacpMuxState string `xml:"lacp-mux-state"`
}
