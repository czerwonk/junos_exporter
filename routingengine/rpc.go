package routingengine

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name                   string                 `xml:"re-name"`
	RouteEngineInformation RouteEngineInformation `xml:"route-engine-information"`
}

type RouteEngineInformation struct {
	RouteEngines []RouteEngine `xml:"route-engine"`
}

type RouteEngine struct {
	Slot              string                 `xml:"slot,omitempty"`
	Status            string                 `xml:"status"`
	Temperature       RouteEngineTemperature `xml:"temperature"`
	MemoryUtilization float64                `xml:"memory-buffer-utilization"`
	CPUTemperature    RouteEngineTemperature `xml:"cpu-temperature"`
	CPUUser           float64                `xml:"cpu-user"`
	CPUBackground     float64                `xml:"cpu-background"`
	CPUSystem         float64                `xml:"cpu-system"`
	CPUInterrupt      float64                `xml:"cpu-interrupt"`
	CPUIdle           float64                `xml:"cpu-idle"`
	UpTime            struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"up-time"`
	LoadAverageOne     float64 `xml:"load-average-one"`
	LoadAverageFive    float64 `xml:"load-average-five"`
	LoadAverageFifteen float64 `xml:"load-average-fifteen"`

	MemorySystemTotal      float64 `xml:"memory-system-total,omitempty"`
	MemorySystemTotalUsed  float64 `xml:"memory-system-total-used,omitempty"`
	MemoryControlPlane     float64 `xml:"memory-control-plane,omitempty"`
	MemoryControlPlaneUsed float64 `xml:"memory-control-plane-used,omitempty"`
	MemoryDataPlane        float64 `xml:"memory-data-plane,omitempty"`
	MemoryDataPlaneUsed    float64 `xml:"memory-data-plane-used,omitempty"`
	MastershipState        string  `xml:"mastership-state,omitempty"`
	MastershipPriority     string  `xml:"mastership-priority,omitempty"`
}

type RouteEngineTemperature struct {
	Value float64 `xml:"celsius,attr"`
}

type RpcReplyNoRE struct {
	XMLName                xml.Name               `xml:"rpc-reply"`
	RouteEngineInformation RouteEngineInformation `xml:"route-engine-information"`
}
