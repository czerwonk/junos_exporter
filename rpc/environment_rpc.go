package rpc

type EnvironmentRpc struct {
	Information struct {
		Items []EnvironmentItem `xml:"environment-item"`
	} `xml:"environment-information"`
}

type EnvironmentItem struct {
	Name        string `xml:"name"`
	Class       string `xml:"class"`
	Status      string `xml:"status"`
	Temperature *struct {
		Value float64 `xml:"celsius,attr"`
	} `xml:"temperature,omitempty"`
}
