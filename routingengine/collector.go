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
	reStatus           *prometheus.Desc
	uptime             *prometheus.Desc
)

func init() {
	l := []string{"target", "slot"}
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
	uptime = prometheus.NewDesc(prefix+"uptime_seconds", "Seconds since boot", l, nil)
	reStatus = prometheus.NewDesc(prefix+"status", "Status of routing-engine (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", l, nil)
}

type routingEngineCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &routingEngineCollector{}
}

// Name returns the name of the collector
func (*routingEngineCollector) Name() string {
	return "Routing Engine"
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
	ch <- reStatus
	ch <- uptime
}

// Collect collects metrics from JunOS
func (c *routingEngineCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RoutingEngineRpc{}
	err := client.RunCommandAndParse("show chassis routing-engine", &x)
	if err != nil {
		return err
	}

	for _, re := range x.Information.RouteEngine {
		c.collectForSlot(re, ch, labelValues)
	}

	return nil
}

func (c *routingEngineCollector) collectForSlot(re RouteEngine, ch chan<- prometheus.Metric, labelValues []string) error {
	l := append(labelValues, re.Slot)

	ch <- prometheus.MustNewConstMetric(temperature, prometheus.GaugeValue, re.Temperature.Value, l...)
	ch <- prometheus.MustNewConstMetric(memoryUtilization, prometheus.GaugeValue, re.MemoryUtilization, l...)
	ch <- prometheus.MustNewConstMetric(cpuTemperature, prometheus.GaugeValue, re.CPUTemperature.Value, l...)
	ch <- prometheus.MustNewConstMetric(cpuUser, prometheus.GaugeValue, re.CPUUser, l...)
	ch <- prometheus.MustNewConstMetric(cpuBackground, prometheus.GaugeValue, re.CPUBackground, l...)
	ch <- prometheus.MustNewConstMetric(cpuSystem, prometheus.GaugeValue, re.CPUSystem, l...)
	ch <- prometheus.MustNewConstMetric(cpuInterrupt, prometheus.GaugeValue, re.CPUInterrupt, l...)
	ch <- prometheus.MustNewConstMetric(cpuIdle, prometheus.GaugeValue, re.CPUIdle, l...)
	ch <- prometheus.MustNewConstMetric(loadAverageOne, prometheus.GaugeValue, re.LoadAverageOne, l...)
	ch <- prometheus.MustNewConstMetric(loadAverageFive, prometheus.GaugeValue, re.LoadAverageFive, l...)
	ch <- prometheus.MustNewConstMetric(loadAverageFifteen, prometheus.GaugeValue, re.LoadAverageFifteen, l...)
	ch <- prometheus.MustNewConstMetric(uptime, prometheus.CounterValue, float64(re.UpTime.Seconds), l...)

	statusValues := map[string]int{
		"OK":      1,
		"Testing": 2,
		"Failed":  3,
		"Absent":  4,
		"Present": 5,
	}
	ch <- prometheus.MustNewConstMetric(reStatus, prometheus.GaugeValue, float64(statusValues[re.Status]), l...)

	return nil
}
