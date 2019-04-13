package firewall

type FirewallRpc struct {
	Filter struct {
		Counters []FilterType `xml:"filter-information"`
		// Policers []FilterType.Policer `xml:"filter-information"`
	} `xml:"firewall-information"`
}

type FilterType struct {

	Counter       struct {
		Name            string `xml:"counter-name"`
		Packets         int64  `xml:"packet-count"`
		Bytes           int64  `xml:"byte-count"`
	} `xml:"counter"`

	// Policer       struct {
	// 	Name            string `xml:"policer-name"`
	// 	Packets         int64  `xml:"packet-count"`
	// 	Bytes           int64  `xml:"byte-count"`
	// } `xml:"policer"`
}
