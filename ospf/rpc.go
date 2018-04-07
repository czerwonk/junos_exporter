package ospf

type Ospf3Rpc struct {
	Information struct {
		Overview struct {
			Areas []Ospf3Area `xml:"ospf-area-overview"`
		} `xml:"ospf-overview"`
	} `xml:"ospf3-overview-information"`
}

type Ospf3Area struct {
	Name      string `xml:"ospf-area"`
	Neighbors struct {
		NeighborsUp int64 `xml:"ospf-nbr-up-count"`
	} `xml:"ospf-nbr-overview"`
}
