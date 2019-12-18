package routingengine

type RoutingEngineRpc struct {
	Information struct {
		RouteEngine []RouteEngine `xml:"route-engine"`
	} `xml:"route-engine-information"`
}

type RouteEngine struct {
	Slot              string                 `xml:"slot"`
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
}

type RouteEngineTemperature struct {
	Value float64 `xml:"celsius,attr"`
}
