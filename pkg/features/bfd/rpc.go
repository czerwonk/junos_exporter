package bfd

type result struct {
	Information struct {
		BfdSessions  []session `xml:"bfd-session"`
		Sessions     int64     `xml:"sessions"`
		Clients      int64     `xml:"clients"`
		CumTransRate float64   `xml:"cumulative-transmission-rate"`
		CumRecRate   float64   `xml:"cumulative-reception-rate"`
	} `xml:"bfd-session-information"`
}

type session struct {
	Neighbor  string `xml:"session-neighbor"`
	State     string `xml:"session-state"`
	Interface string `xml:"session-interface"`
	Client    struct {
		Name string `xml:"client-name"`
	} `xml:"bfd-client"`
}
