package firewall

type FirewallRpc struct {
	Information struct {
		Filters []Filter `xml:"filter-information"`
	} `xml:"firewall-information"`
}

type Filter struct {
	Name         string               `xml:"filter-name"`
	Counters    []FilterCounter       `xml:"counter"`
	Policers    []FilterPolicer       `xml:"policer"`
}

type FilterCounter struct {
	Name         string `xml:"counter-name"`
	Packets      int64  `xml:"packet-count"`
	Bytes        int64  `xml:"byte-count"`
}

type FilterPolicer struct {
	Name         string `xml:"policer-name"`
	Packets      int64  `xml:"packet-count"`
	Bytes        int64  `xml:"byte-count"`
}
