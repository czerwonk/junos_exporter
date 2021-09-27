package environment

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name                            string                          `xml:"re-name"`
	EnvironmentComponentInformation EnvironmentComponentInformation `xml:"environment-component-information"`
	EnvironmentInformation          EnvironmentInformation          `xml:"environment-information"`
}

type EnvironmentComponentInformation struct {
	EnvironmentComponentItem []EnvironmentComponentItem `xml:"environment-component-item"`
}

type EnvironmentComponentItem struct {
	Name            string `xml:"name"`
	State           string `xml:"state"`
	FanSpeedReading []struct {
		FanName  string `xml:"fan-name"`
		FanSpeed string `xml:"fan-speed"`
	} `xml:"fan-speed-reading"`
	DcInformation struct {
		DcDetail struct {
			DcVoltage     float64 `xml:"dc-voltage,omitempty"`
			DcCurrent     float64 `xml:"dc-current,omitempty"`
			DcPower       float64 `xml:"dc-power,omitempty"`
			DcLoad        float64 `xml:"dc-load,omitempty"`
			Str3DcVoltage float64 `xml:"str3-dc-voltage,omitempty"`
		} `xml:"dc-detail"`
	} `xml:"dc-information"`
}

type EnvironmentInformation struct {
	Items []EnvironmentItemRpc `xml:"environment-item"`
}

type EnvironmentItemRpc struct {
	Name        string `xml:"name"`
	Class       string `xml:"class"`
	Status      string `xml:"status"`
	Temperature *struct {
		Value float64 `xml:"celsius,attr"`
	} `xml:"temperature,omitempty"`
}

type RpcReplyNoRE struct {
	XMLName                         xml.Name                        `xml:"rpc-reply"`
	EnvironmentComponentInformation EnvironmentComponentInformation `xml:"environment-component-information"`
	EnvironmentInformation          EnvironmentInformation          `xml:"environment-information"`
}
