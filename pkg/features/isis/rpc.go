// SPDX-License-Identifier: MIT

package isis

type result struct {
	Information struct {
		Adjacencies []adjacency `xml:"isis-adjacency"`
	} `xml:"isis-adjacency-information"`
}

type adjacency struct {
	InterfaceName  string `xml:"interface-name"`
	SystemName     string `xml:"system-name"`
	Level          int64  `xml:"level"`
	AdjacencyState string `xml:"adjacency-state"`
	Holdtime       int64  `xml:"holdtime"`
	SNPA           string `xml:"snpa"`
}
