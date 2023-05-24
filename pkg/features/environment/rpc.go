// SPDX-License-Identifier: MIT

package environment

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name           `xml:"rpc-reply"`
	Results multiEngineResults `xml:"multi-routing-engine-results"`
}

type multiEngineResults struct {
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name                            string                          `xml:"re-name"`
	EnvironmentComponentInformation environmentComponentInformation `xml:"environment-component-information"`
	EnvironmentInformation          environmentInformation          `xml:"environment-information"`
}

type environmentComponentInformation struct {
	EnvironmentComponentItem []environmentComponentItem `xml:"environment-component-item"`
}

type environmentComponentItem struct {
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

type environmentInformation struct {
	Items []environmentItem `xml:"environment-item"`
}

type environmentItem struct {
	Name        string `xml:"name"`
	Class       string `xml:"class"`
	Status      string `xml:"status"`
	Temperature *struct {
		Value float64 `xml:"celsius,attr"`
	} `xml:"temperature,omitempty"`
}

type singleEngineResult struct {
	XMLName                         xml.Name                        `xml:"rpc-reply"`
	EnvironmentComponentInformation environmentComponentInformation `xml:"environment-component-information"`
	EnvironmentInformation          environmentInformation          `xml:"environment-information"`
}
