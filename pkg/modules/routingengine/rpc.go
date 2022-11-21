package routingengine

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name       `xml:"rpc-reply"`
	Results routingEngines `xml:"multi-routing-engine-results"`
}

type routingEngines struct {
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name        string       `xml:"re-name"`
	Information routeEngines `xml:"route-engine-information"`
}

type routeEngines struct {
	RouteEngines []routeEngine `xml:"route-engine"`
}

type routeEngine struct {
	Slot              string           `xml:"slot,omitempty"`
	Status            string           `xml:"status"`
	Temperature       temperatureValue `xml:"temperature"`
	MemoryUtilization float64          `xml:"memory-buffer-utilization"`
	CPUTemperature    temperatureValue `xml:"cpu-temperature"`
	CPUUser           float64          `xml:"cpu-user"`
	CPUBackground     float64          `xml:"cpu-background"`
	CPUSystem         float64          `xml:"cpu-system"`
	CPUInterrupt      float64          `xml:"cpu-interrupt"`
	CPUIdle           float64          `xml:"cpu-idle"`
	CPUUser1          float64          `xml:"cpu-user1"`
	CPUBackground1    float64          `xml:"cpu-background1"`
	CPUSystem1        float64          `xml:"cpu-system1"`
	CPUInterrupt1     float64          `xml:"cpu-interrupt1"`
	CPUIdle1          float64          `xml:"cpu-idle1"`
	CPUUser2          float64          `xml:"cpu-user2"`
	CPUBackground2    float64          `xml:"cpu-background2"`
	CPUSystem2        float64          `xml:"cpu-system2"`
	CPUInterrupt2     float64          `xml:"cpu-interrupt2"`
	CPUIdle2          float64          `xml:"cpu-idle2"`
	CPUUser3          float64          `xml:"cpu-user3"`
	CPUBackground3    float64          `xml:"cpu-background3"`
	CPUSystem3        float64          `xml:"cpu-system3"`
	CPUInterrupt3     float64          `xml:"cpu-interrupt3"`
	CPUIdle3          float64          `xml:"cpu-idle3"`
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

type temperatureValue struct {
	Value float64 `xml:"celsius,attr"`
}

type singleRoutingEngineResult struct {
	XMLName     xml.Name     `xml:"rpc-reply"`
	Information routeEngines `xml:"route-engine-information"`
}
