package bfd

type bfdRpc struct {
	Information struct {
		BfdSessions []bfdSession `xml:"bfd-session"`
		sessions    int64 `xml:"sessions"`
		clients     int64 `xml:"clients"`
		CumTransRate  float64 `xml:"cumulative-transmission-rate"`
		CumRecRate  float64 `xml:"cumulative-reception-rate"`
	} `xml:"bfd-session-information"`
}

type bfdSession struct {
	Neighbor           string    `xml:"session-neighbor"`
	State              string    `xml:"session-state"`
	Interface          string    `xml:"session-interface"`
	Client             struct {
		Name       string    `xml:"client-name"`
	} `xml:"bfd-client"`
}

