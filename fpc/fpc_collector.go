package fpc

import (
	"strconv"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_fpc_"

var (
	upDesc          *prometheus.Desc
	temperatureDesc *prometheus.Desc
	memoryDesc      *prometheus.Desc
	uptimeDesc      *prometheus.Desc
	powerDesc       *prometheus.Desc
)

type fpcCollector struct {
}

func init() {
	l := []string{"target", "slot"}
	upDesc = prometheus.NewDesc(prefix+"up", "Status of the linecard (1 = Online)", l, nil)
	temperatureDesc = prometheus.NewDesc(prefix+"temperature_celsius", "Temperature in degree celsius", l, nil)
	uptimeDesc = prometheus.NewDesc(prefix+"uptime_seconds", "Seconds since boot", l, nil)
	powerDesc = prometheus.NewDesc(prefix+"max_power_consumption_watt", "Maximum power consumption in Watt", l, nil)

	l = append(l, "memory_type")
	memoryDesc = prometheus.NewDesc(prefix+"memory_bytes", "Memory size in bytes", l, nil)
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &fpcCollector{}
}

// Describe describes the metrics
func (*fpcCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- temperatureDesc
	ch <- memoryDesc
	ch <- uptimeDesc
	ch <- powerDesc
}

// Collect collects metrics from JunOS
func (c *fpcCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := FPCRpc{}
	err := client.RunCommandAndParse("show chassis fpc detail", &r)
	if err != nil {
		return err
	}

	for _, f := range r.Information.FPCs {
		c.collectForFPC(ch, labelValues, &f)
	}

	return nil
}

func (c *fpcCollector) collectForFPC(ch chan<- prometheus.Metric, labelValues []string, fpc *FPC) {
	up := 0
	if fpc.State == "Online" {
		up = 1
	}

	l := append(labelValues, strconv.Itoa(fpc.Slot))
	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(up), l...)
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

func (c *fpcCollector) collectMemory(memory uint, memType string, ch chan<- prometheus.Metric, labelValues []string) {
	if memory > 0 {
		l := append(labelValues, memType)
		ch <- prometheus.MustNewConstMetric(memoryDesc, prometheus.GaugeValue, float64(memory), l...)
	}
}
