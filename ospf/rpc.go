package ospf

type OspfRpc struct {
	Information struct {
		Overview struct {
			Areas []OspfAreaRpc `xml:"ospf-area-overview"`
		} `xml:"ospf-overview"`
	} `xml:"ospf-overview-information"`
}

type Ospf3Rpc struct {
	Information struct {
		Overview struct {
			Areas []OspfAreaRpc `xml:"ospf-area-overview"`
		} `xml:"ospf-overview"`
	} `xml:"ospf3-overview-information"`
}

type OspfAreaRpc struct {
	Name      string `xml:"ospf-area"`
	Neighbors struct {
		NeighborsUp int64 `xml:"ospf-nbr-up-count"`
	} `xml:"ospf-nbr-overview"`
}
