// SPDX-License-Identifier: MIT

package power

import "encoding/xml"

type multiRoutingEngineResult struct {
	XMLName xml.Name       `xml:"rpc-reply"`
	Results routingEngines `xml:"multi-routing-engine-results"`
}

type routingEngines struct {
	RoutingEngine []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name                  string                `xml:"re-name"`
	PowerUsageInformation powerUsageInformation `xml:"power-usage-information"`
}

type powerUsageInformation struct {
	PowerUsageItem   []powerUsageItem `xml:"power-usage-item"`
	PowerUsageSystem struct {
		PowerUsageZoneInformation []struct {
			Zone                string `xml:"zone"`
			CapacityActual      int    `xml:"capacity-actual"`
			CapacityMax         int    `xml:"capacity-max"`
			CapacityAllocated   int    `xml:"capacity-allocated"`
			CapacityRemaining   int    `xml:"capacity-remaining"`
			CapacityActualUsage int    `xml:"capacity-actual-usage"`
		} `xml:"power-usage-zone-information"`
		CapacitySysActual    int `xml:"capacity-sys-actual"`
		CapacitySysMax       int `xml:"capacity-sys-max"`
		CapacitySysRemaining int `xml:"capacity-sys-remaining"`
	} `xml:"power-usage-system"`
}

type powerUsageItem struct {
	Name           string `xml:"name"`
	State          string `xml:"state"`
	DcOutputDetail struct {
		DcPower   int    `xml:"dc-power"`
		Zone      string `xml:"zone"`
		DcCurrent int    `xml:"dc-current"`
		DcVoltage int    `xml:"dc-voltage"`
		DcLoad    int    `xml:"dc-load"`
	} `xml:"dc-output-detail"`
}

type singleRoutingEngineResult struct {
	XMLName               xml.Name              `xml:"rpc-reply"`
	PowerUsageInformation powerUsageInformation `xml:"power-usage-information"`
}
