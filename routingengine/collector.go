package routingengine

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_route_engine_"

var (
	temperature             *prometheus.Desc
	memoryUtilization       *prometheus.Desc
	memoryBufferUtilization *prometheus.Desc
	cpuTemperature          *prometheus.Desc
	cpuUser                 *prometheus.Desc
	cpuBackground           *prometheus.Desc
	cpuSystem               *prometheus.Desc
	cpuInterrupt            *prometheus.Desc
	cpuIdle                 *prometheus.Desc
	loadAverageOne          *prometheus.Desc
	loadAverageFive         *prometheus.Desc
	loadAverageFifteen      *prometheus.Desc
	reStatus                *prometheus.Desc
	uptime                  *prometheus.Desc
	memorySystemTotal       *prometheus.Desc
	memorySystemTotalUsed   *prometheus.Desc
	memorySystemTotalUtil   *prometheus.Desc
	memoryControlPlane      *prometheus.Desc
	memoryControlPlaneUsed  *prometheus.Desc
	memoryControlPlaneUtil  *prometheus.Desc
	memoryDataPlane         *prometheus.Desc
	memoryDataPlaneUsed     *prometheus.Desc
	memoryDataPlaneUtil     *prometheus.Desc
	mastershipState         *prometheus.Desc
	mastershipPriority      *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name", "slot"}
	temperature = prometheus.NewDesc(prefix+"temp", "Temperature of the air flowing past the Routing Engine (in degrees C)", l, nil)
	memoryUtilization = prometheus.NewDesc(prefix+"memory_utilization_percent", "Percent of Routing Engine memory being used", l, nil)
	cpuTemperature = prometheus.NewDesc(prefix+"cpu_temp", "Temperature of the CPU (in degrees C)", l, nil)
	cpuUser = prometheus.NewDesc(prefix+"cpu_user_percent", "Percent of CPU time being used by user processes", l, nil)
	cpuBackground = prometheus.NewDesc(prefix+"cpu_background_percent", "Percent of CPU time being used by background processes", l, nil)
	cpuSystem = prometheus.NewDesc(prefix+"cpu_system_percent", "Percent of CPU time being used by kernel processes", l, nil)
	cpuInterrupt = prometheus.NewDesc(prefix+"cpu_interrupt_percent", "Percent of CPU time being used by interrupts", l, nil)
	cpuIdle = prometheus.NewDesc(prefix+"cpu_idle_percent", "Percent of CPU time that is idle", l, nil)
	loadAverageOne = prometheus.NewDesc(prefix+"load_average_one", "Routing Engine load averages for the last 1 minute", l, nil)
	loadAverageFive = prometheus.NewDesc(prefix+"load_average_five", "Routing Engine load averages for the last 5 minutes", l, nil)
	loadAverageFifteen = prometheus.NewDesc(prefix+"load_average_fifteen", "Routing Engine load averages for the last 15 minutes", l, nil)
	uptime = prometheus.NewDesc(prefix+"uptime_seconds", "Seconds since boot", l, nil)
	reStatus = prometheus.NewDesc(prefix+"status", "Status of routing-engine (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", l, nil)

	memorySystemTotal = prometheus.NewDesc(prefix+"memory_system_total_bytes", "Total System memory", l, nil)
	memorySystemTotalUsed = prometheus.NewDesc(prefix+"memory_system_total_used_bytes", "System memory utilized", l, nil)
	memoryControlPlane = prometheus.NewDesc(prefix+"memory_control_plane_bytes", "Total Control Plane memory", l, nil)
	memoryControlPlaneUsed = prometheus.NewDesc(prefix+"memory_control_plane_used_bytes", "Control Plane utilized", l, nil)
	memoryDataPlane = prometheus.NewDesc(prefix+"memory_data_plane_bytes", "Total Data Plane memory", l, nil)
	memoryDataPlaneUsed = prometheus.NewDesc(prefix+"memory_data_plane_used_bytes", "Data Plane memory utilized", l, nil)

	l = []string{"target", "re_name", "slot", "mastership"}
	mastershipState = prometheus.NewDesc(prefix+"mastership_state", "Mastership state", l, nil)
	mastershipPriority = prometheus.NewDesc(prefix+"mastership_priority", "Mastership priority", l, nil)
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
	ch <- memorySystemTotal
	ch <- memorySystemTotalUsed
	ch <- memoryControlPlane
	ch <- memoryControlPlaneUsed
	ch <- memoryDataPlane
	ch <- memoryDataPlaneUsed
	ch <- mastershipState
	ch <- mastershipPriority
}

// Collect collects metrics from JunOS
func (c *routingEngineCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RpcReply{}
	err := client.RunCommandAndParseWithParser("show chassis routing-engine", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, re := range x.MultiRoutingEngineResults.RoutingEngine {
		labelValues := append(labelValues, re.Name)
		for _, re_ := range re.RouteEngineInformation.RouteEngines {
			c.collectForSlot(re_, ch, labelValues)
		}
	}

	return nil
}

func (c *routingEngineCollector) collectForSlot(re RouteEngine, ch chan<- prometheus.Metric, labelValues []string) error {
	if re.Slot == "" {
		re.Slot = "N/A"
	}
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

	if re.MemorySystemTotal > 0 {
		ch <- prometheus.MustNewConstMetric(memorySystemTotal, prometheus.GaugeValue, float64(re.MemorySystemTotal)*1024*1024, l...)
	}
	if re.MemorySystemTotalUsed > 0 {
		ch <- prometheus.MustNewConstMetric(memorySystemTotalUsed, prometheus.GaugeValue, float64(re.MemorySystemTotalUsed)*1024*1024, l...)
	}
	if re.MemoryControlPlane > 0 {
		ch <- prometheus.MustNewConstMetric(memoryControlPlane, prometheus.GaugeValue, float64(re.MemoryControlPlane)*1024*1024, l...)
	}
	if re.MemoryControlPlaneUsed > 0 {
		ch <- prometheus.MustNewConstMetric(memoryControlPlaneUsed, prometheus.GaugeValue, float64(re.MemoryControlPlaneUsed)*1024*1024, l...)
	}
	if re.MemoryDataPlane > 0 {
		ch <- prometheus.MustNewConstMetric(memoryDataPlane, prometheus.GaugeValue, float64(re.MemoryDataPlane)*1024*1024, l...)
	}
	if re.MemoryDataPlaneUsed > 0 {
		ch <- prometheus.MustNewConstMetric(memoryDataPlaneUsed, prometheus.GaugeValue, float64(re.MemoryDataPlaneUsed)*1024*1024, l...)
	}

	if re.MastershipState != "" {
		ch <- prometheus.MustNewConstMetric(mastershipState, prometheus.GaugeValue, float64(1), append(l, re.MastershipState)...)
	}

	if re.MastershipPriority != "" {
		ch <- prometheus.MustNewConstMetric(mastershipPriority, prometheus.GaugeValue, float64(1), append(l, re.MastershipPriority)...)
	}

	return nil
}

func parseXML(b []byte, res *RpcReply) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := RpcReplyNoRE{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.MultiRoutingEngineResults.RoutingEngine = []RoutingEngine{
		RoutingEngine{
			Name: "N/A",
			RouteEngineInformation: fi.RouteEngineInformation,
		},
	}

	return nil
}
