package fpc

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_fpc_"

var (
	// fpc detail + fpc
	upDesc          *prometheus.Desc
	temperatureDesc *prometheus.Desc
	memoryDesc      *prometheus.Desc
	uptimeDesc      *prometheus.Desc
	powerDesc       *prometheus.Desc
	picstatusDesc   *prometheus.Desc

	// fpc only
	cpuTotalDesc                *prometheus.Desc
	cpuInterruptDesc            *prometheus.Desc
	cpuAvgDesc                  *prometheus.Desc
	memoryHeapUtilizationDesc   *prometheus.Desc
	memoryBufferUtilizationDesc *prometheus.Desc
)

type fpcCollector struct {
}

func init() {
	l := []string{"target", "re_name", "slot"}
	upDesc = prometheus.NewDesc(prefix+"up", "Status of the linecard (1 = Online)", l, nil)
	temperatureDesc = prometheus.NewDesc(prefix+"temperature_celsius", "Temperature in degree celsius", l, nil)
	uptimeDesc = prometheus.NewDesc(prefix+"uptime_seconds", "Seconds since boot", l, nil)
	powerDesc = prometheus.NewDesc(prefix+"max_power_consumption_watt", "Maximum power consumption in Watt", l, nil)

	cpuTotalDesc = prometheus.NewDesc(prefix+"cpu_total", "Overall CPU utilization in percent", l, nil)
	cpuInterruptDesc = prometheus.NewDesc(prefix+"cpu_interrupts", "Number of CPU interrupts", l, nil)
	memoryHeapUtilizationDesc = prometheus.NewDesc(prefix+"mem_heap_utilization_percent", "Heap usage percent", l, nil)
	memoryBufferUtilizationDesc = prometheus.NewDesc(prefix+"mem_buffers_utilization_percent", "Buffers usage percent", l, nil)
	cpuAvgDesc = prometheus.NewDesc(prefix+"cpu_load_avg", "CPU load", append(l, "interval"), nil)

	l = append(l, "memory_type")
	memoryDesc = prometheus.NewDesc(prefix+"memory_bytes", "Memory size in bytes", l, nil)

	l_pic := []string{"target", "re_name", "fpc_slot", "pic_slot", "pic_type"}
	picstatusDesc = prometheus.NewDesc(prefix+"pic_status", "Status of the PIC (1 = Online, 0 = Offline)", l_pic, nil)
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &fpcCollector{}
}

// Name returns the name of the collector
func (*fpcCollector) Name() string {
	return "FPC"
}

// Describe describes the metrics
func (*fpcCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- temperatureDesc
	ch <- memoryDesc
	ch <- uptimeDesc
	ch <- powerDesc
	ch <- picstatusDesc
	ch <- cpuTotalDesc
	ch <- cpuInterruptDesc
	ch <- memoryHeapUtilizationDesc
	ch <- memoryBufferUtilizationDesc
	ch <- cpuAvgDesc
}

// Collect collects metrics from JunOS
func (c *fpcCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.CollectFPCDetail(client, ch, labelValues)
	if err != nil {
		return err
	}

	err = c.CollectFPC(client, ch, labelValues)
	if err != nil {
		return err
	}

	err = c.CollectPIC(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

// CollectFPC collects metrics from JunOS
func (c *fpcCollector) CollectFPCDetail(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := RpcReply{}
	err := client.RunCommandAndParseWithParser("show chassis fpc detail", func(b []byte) error {
		return parseXML(b, &r)
	})

	if err != nil {
		return err
	}

	for _, r := range r.MultiRoutingEngineResults.RoutingEngine {
		labels := append(labelValues, r.Name)
		for _, f := range r.FPCs.FPC {
			c.collectForFPCDetail(ch, labels, &f)
		}
	}

	return nil
}

// Collect collects metrics from JunOS
func (c *fpcCollector) CollectFPC(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := RpcReply{}
	err := client.RunCommandAndParseWithParser("show chassis fpc", func(b []byte) error {
		return parseXML(b, &r)
	})
	if err != nil {
		return err
	}

	for _, r := range r.MultiRoutingEngineResults.RoutingEngine {
		labels := append(labelValues, r.Name)
		for _, f := range r.FPCs.FPC {
			c.collectForFPC(ch, labels, &f)
		}
	}
	return nil
}

func (c *fpcCollector) CollectPIC(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := RpcReply{}
	err := client.RunCommandAndParseWithParser("show chassis fpc pic-status", func(b []byte) error {
		return parseXML(b, &r)
	})
	if err != nil {
		return err
	}
	for _, r := range r.MultiRoutingEngineResults.RoutingEngine {
		labels := append(labelValues, r.Name)
		for _, f := range r.FPCs.FPC {
			for _, p := range f.Pics {
				c.collectForPIC(ch, labels, &f, &p)
			}
		}
	}
	return nil
}

func (c *fpcCollector) collectForFPCDetail(ch chan<- prometheus.Metric, labelValues []string, fpc *FPC) {
	l := append(labelValues, strconv.Itoa(fpc.Slot))
	ch <- prometheus.MustNewConstMetric(uptimeDesc, prometheus.CounterValue, float64(fpc.UpTime.Seconds), l...)

	if fpc.Temperature.Celsius > 0 {
		ch <- prometheus.MustNewConstMetric(temperatureDesc, prometheus.GaugeValue, float64(fpc.Temperature.Celsius), l...)
	}

	if fpc.MaxPowerConsumption > 0 {
		ch <- prometheus.MustNewConstMetric(powerDesc, prometheus.GaugeValue, float64(fpc.MaxPowerConsumption), l...)
	}

	c.collectMemory(fpc.MemorySramSize, "sram", ch, l)
	c.collectMemory(fpc.MemoryDramSize, "dram", ch, l)
	c.collectMemory(fpc.MemoryDdrDramSize, "ddr-dram", ch, l)
	c.collectMemory(fpc.MemorySdramSize, "sdram", ch, l)
	c.collectMemory(fpc.MemoryRldramSize, "rl-dram", ch, l)
}

func (c *fpcCollector) collectForFPC(ch chan<- prometheus.Metric, labelValues []string, fpc *FPC) {
	up := 0
	if fpc.State == "Online" {
		up = 1
	}
	l := append(labelValues, strconv.Itoa(fpc.Slot))

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(up), l...)
	if up == 1 {
		ch <- prometheus.MustNewConstMetric(cpuTotalDesc, prometheus.GaugeValue, float64(fpc.CpuTotal), l...)
		ch <- prometheus.MustNewConstMetric(cpuInterruptDesc, prometheus.GaugeValue, float64(fpc.CpuInterrupt), l...)
		ch <- prometheus.MustNewConstMetric(memoryHeapUtilizationDesc, prometheus.GaugeValue, float64(fpc.MemoryHeapUtilization), l...)
		ch <- prometheus.MustNewConstMetric(memoryBufferUtilizationDesc, prometheus.GaugeValue, float64(fpc.MemoryBufferUtilization), l...)
		ch <- prometheus.MustNewConstMetric(cpuAvgDesc, prometheus.GaugeValue, float64(fpc.Cpu1min_avg), append(l, "1min")...)
		ch <- prometheus.MustNewConstMetric(cpuAvgDesc, prometheus.GaugeValue, float64(fpc.Cpu5min_avg), append(l, "5min")...)
		ch <- prometheus.MustNewConstMetric(cpuAvgDesc, prometheus.GaugeValue, float64(fpc.Cpu15min_avg), append(l, "15min")...)
	}
}

func (c *fpcCollector) collectMemory(memory uint, memType string, ch chan<- prometheus.Metric, labelValues []string) {
	if memory > 0 {
		l := append(labelValues, memType)
		// values returned by the command are in MiB
		ch <- prometheus.MustNewConstMetric(memoryDesc, prometheus.GaugeValue, float64(memory)*1024*1024, l...)
	}
}

func (c *fpcCollector) collectForPIC(ch chan<- prometheus.Metric, labelValues []string, fpc *FPC, pic *PIC) {
	picup := 0
	if pic.PicState == "Online" {
		picup = 1
	}

	l_pic := append(labelValues, strconv.Itoa(fpc.Slot), strconv.Itoa(pic.PicSlot), pic.PicType)
	ch <- prometheus.MustNewConstMetric(picstatusDesc, prometheus.GaugeValue, float64(picup), l_pic...)
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
		{
			Name: "N/A",
			FPCs: fi.FPCs,
		},
	}
	return nil
}
