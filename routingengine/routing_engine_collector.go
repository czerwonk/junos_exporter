package routingengine

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_route_engine_"

var (
	temperature        *prometheus.Desc
	memoryUtilization  *prometheus.Desc
	cpuTemperature     *prometheus.Desc
	cpuUser            *prometheus.Desc
	cpuBackground      *prometheus.Desc
	cpuSystem          *prometheus.Desc
	cpuInterrupt       *prometheus.Desc
	cpuIdle            *prometheus.Desc
	loadAverageOne     *prometheus.Desc
	loadAverageFive    *prometheus.Desc
	loadAverageFifteen *prometheus.Desc
)

func init() {
	l := []string{"target"}
	temperature = prometheus.NewDesc(prefix+"temp", "Temperature of the air flowing past the Routing Engine (in degrees C)", l, nil)
	memoryUtilization = prometheus.NewDesc(prefix+"memory_utilization", "Percentage of Routing Engine memory being used", l, nil)
	cpuTemperature = prometheus.NewDesc(prefix+"cpu_temp", "Temperature of the CPU (in degrees C)", l, nil)
	cpuUser = prometheus.NewDesc(prefix+"cpu_user_percent", "Percentage of CPU time being used by user processes", l, nil)
	cpuBackground = prometheus.NewDesc(prefix+"cpu_background_percent", "Percentage of CPU time being used by background processes", l, nil)
	cpuSystem = prometheus.NewDesc(prefix+"cpu_system_percent", "Percentage of CPU time being used by kernel processes", l, nil)
	cpuInterrupt = prometheus.NewDesc(prefix+"cpu_interrupt_percent", "Percentage of CPU time being used by interrupts", l, nil)
	cpuIdle = prometheus.NewDesc(prefix+"cpu_idle_percent", "Percentage of CPU time that is idle", l, nil)
	loadAverageOne = prometheus.NewDesc(prefix+"load_average_one", "Routing Engine load averages for the last 1 minute", l, nil)
	loadAverageFive = prometheus.NewDesc(prefix+"load_average_five", "Routing Engine load averages for the last 5 minutes", l, nil)
	loadAverageFifteen = prometheus.NewDesc(prefix+"load_average_fifteen", "Routing Engine load averages for the last 15 minutes", l, nil)
}

type routingEngineCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &routingEngineCollector{}
}

// Describe describes the metrics
func (*routingEngineCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperature
	ch <- memoryUtilization
	ch <- cpuTemperature
	ch <- cpuUser
	ch <- cpuBackground
	ch <- cpuSystem
	ch <- cpuInterrupt
	ch <- cpuIdle
	ch <- loadAverageOne
	ch <- loadAverageFive
	ch <- loadAverageFifteen
}

// Collect collects metrics from JunOS
func (c *routingEngineCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	stats, err := c.routeEngineStats(client)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(temperature, prometheus.GaugeValue, stats.Temperature, labelValues...)
	ch <- prometheus.MustNewConstMetric(memoryUtilization, prometheus.GaugeValue, stats.MemoryUtilization, labelValues...)
	ch <- prometheus.MustNewConstMetric(cpuTemperature, prometheus.GaugeValue, stats.CPUTemperature, labelValues...)
	ch <- prometheus.MustNewConstMetric(cpuUser, prometheus.GaugeValue, stats.CPUUser, labelValues...)
	ch <- prometheus.MustNewConstMetric(cpuBackground, prometheus.GaugeValue, stats.CPUBackground, labelValues...)
	ch <- prometheus.MustNewConstMetric(cpuSystem, prometheus.GaugeValue, stats.CPUSystem, labelValues...)
	ch <- prometheus.MustNewConstMetric(cpuInterrupt, prometheus.GaugeValue, stats.CPUInterrupt, labelValues...)
	ch <- prometheus.MustNewConstMetric(cpuIdle, prometheus.GaugeValue, stats.CPUIdle, labelValues...)
	ch <- prometheus.MustNewConstMetric(loadAverageOne, prometheus.GaugeValue, stats.LoadAverageOne, labelValues...)
	ch <- prometheus.MustNewConstMetric(loadAverageFive, prometheus.GaugeValue, stats.LoadAverageFive, labelValues...)
	ch <- prometheus.MustNewConstMetric(loadAverageFifteen, prometheus.GaugeValue, stats.LoadAverageFifteen, labelValues...)

	return nil
}

func (c *routingEngineCollector) routeEngineStats(client *rpc.Client) (*RouteEngineStats, error) {
	var x = RoutingEngineRpc{}
	err := client.RunCommandAndParse("show chassis routing-engine", &x)
	if err != nil {
		return nil, err
	}

	r := &RouteEngineStats{
		Temperature:        float64(x.Information.RouteEngine.Temperature.Value),
		MemoryUtilization:  float64(x.Information.RouteEngine.MemoryUtilization),
		CPUTemperature:     float64(x.Information.RouteEngine.CPUTemperature.Value),
		CPUUser:            float64(x.Information.RouteEngine.CPUUser),
		CPUBackground:      float64(x.Information.RouteEngine.CPUBackground),
		CPUSystem:          float64(x.Information.RouteEngine.CPUSystem),
		CPUInterrupt:       float64(x.Information.RouteEngine.CPUInterrupt),
		CPUIdle:            float64(x.Information.RouteEngine.CPUIdle),
		LoadAverageOne:     float64(x.Information.RouteEngine.LoadAverageOne),
		LoadAverageFive:    float64(x.Information.RouteEngine.LoadAAverageFive),
		LoadAverageFifteen: float64(x.Information.RouteEngine.LoadAverageFifteen),
	}

	return r, nil
}
