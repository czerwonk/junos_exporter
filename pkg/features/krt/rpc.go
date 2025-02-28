package krt

import "encoding/xml"

type resultKRT struct {
	XMLName             xml.Name `xml:"rpc-reply"`
	Text                string   `xml:",chardata"`
	Junos               string   `xml:"junos,attr"`
	KrtQueueInformation struct {
		Text     string `xml:",chardata"`
		Xmlns    string `xml:"xmlns,attr"`
		KrtQueue []struct {
			Text            string  `xml:",chardata"`
			KrtqType        string  `xml:"krtq-type"`
			KrtqQueueLength float64 `xml:"krtq-queue-length"`
		} `xml:"krt-queue"`
	} `xml:"krt-queue-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
