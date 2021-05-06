package power

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name                  string                `xml:"re-name"`
	PowerUsageInformation PowerUsageInformation `xml:"power-usage-information"`
}

type PowerUsageInformation struct {
	PowerUsageItem   []PowerUsageItem `xml:"power-usage-item"`
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

type PowerUsageItem struct {
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

type RpcReplyNoRE struct {
	XMLName               xml.Name              `xml:"rpc-reply"`
	PowerUsageInformation PowerUsageInformation `xml:"power-usage-information"`
}
