// SPDX-License-Identifier: MIT

package alarm

import (
	"encoding/xml"
)

type multiEngineResult struct {
	XMLName        xml.Name        `xml:"multi-routing-engine-results"`
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name      string           `xml:"re-name"`
	AlarmInfo alarmInformation `xml:"alarm-information"`
}

type alarmInformation struct {
	XMLName xml.Name  `xml:"alarm-information"`
	Details []details `xml:"alarm-detail"`
}

type details struct {
	Class       string `xml:"alarm-class"`
	Description string `xml:"alarm-description"`
	Type        string `xml:"alarm-type"`
}
