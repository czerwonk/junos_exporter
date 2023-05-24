// SPDX-License-Identifier: MIT

package firewall

type result struct {
	Information struct {
		Filters []filter `xml:"filter-information"`
	} `xml:"firewall-information"`
}

type filter struct {
	Name     string          `xml:"filter-name"`
	Counters []filterCounter `xml:"counter"`
	Policers []filterPolicer `xml:"policer"`
}

type filterCounter struct {
	Name    string `xml:"counter-name"`
	Packets int64  `xml:"packet-count"`
	Bytes   int64  `xml:"byte-count"`
}

type filterPolicer struct {
	Name    string `xml:"policer-name"`
	Packets int64  `xml:"packet-count"`
	Bytes   int64  `xml:"byte-count"`
}
