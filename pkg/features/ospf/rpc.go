// SPDX-License-Identifier: MIT

package ospf

type v2Result struct {
	Information struct {
		Overview struct {
			Areas []area `xml:"ospf-area-overview"`
		} `xml:"ospf-overview"`
	} `xml:"ospf-overview-information"`
}

type v3Result struct {
	Information struct {
		Overview struct {
			Areas []area `xml:"ospf-area-overview"`
		} `xml:"ospf-overview"`
	} `xml:"ospf3-overview-information"`
}

type area struct {
	Name      string `xml:"ospf-area"`
	Neighbors struct {
		NeighborsUp int64 `xml:"ospf-nbr-up-count"`
	} `xml:"ospf-nbr-overview"`
}
