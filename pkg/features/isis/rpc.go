// SPDX-License-Identifier: MIT

package isis

import "encoding/xml"

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

type interfaces struct {
	XMLName                  xml.Name `xml:"rpc-reply"`
	Text                     string   `xml:",chardata"`
	Junos                    string   `xml:"junos,attr"`
	IsisInterfaceInformation struct {
		Text          string `xml:",chardata"`
		Xmlns         string `xml:"xmlns,attr"`
		IsisInterface []struct {
			InterfaceName      string  `xml:"interface-name"`
			LSPInterval        float64 `xml:"lsp-interval"`
			CSNPInterval       float64 `xml:"csnp-interval"`
			HelloPadding       string  `xml:"hello-padding"`
			MaxHelloSize       float64 `xml:"max-hello-size"`
			InterfaceLevelData struct {
				Level             string  `xml:"level"`
				AdjacencyCount    float64 `xml:"adjacency-count"`
				InterfacePriority float64 `xml:"interface-priority"`
				Metric            float64 `xml:"metric"`
				HelloTime         float64 `xml:"hello-time"`
				HoldTime          float64 `xml:"holdtime"`
				Passive           string  `xml:"passive"`
			} `xml:"interface-level-data"`
		} `xml:"isis-interface"`
	} `xml:"isis-interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
