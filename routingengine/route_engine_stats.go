package routingengine

type RouteEngineStats struct {
	Temperature        float64
	MemoryUtilization  float64
	CPUTemperature     float64
	CPUUser            float64
	CPUBackground      float64
	CPUSystem          float64
	CPUInterrupt       float64
	CPUIdle            float64
	LoadAverageOne     float64
	LoadAverageFive    float64
	LoadAverageFifteen float64
}
