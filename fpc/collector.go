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
	picstatusDesc   *prometheus.Desc
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

	l_pic := []string{"target", "fpc_slot", "pic_slot", "pic_type"}
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
}

// Collect collects metrics from JunOS
func (c *fpcCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.CollectFPC(client, ch, labelValues)
	if err != nil {
		return err
	}

	err = c.CollectPIC(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

// Collect collects metrics from JunOS
func (c *fpcCollector) CollectFPC(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
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

func (c *fpcCollector) CollectPIC(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	r := FPCRpc{}
	err := client.RunCommandAndParse("show chassis fpc pic-status", &r)
	if err != nil {
		return err
	}
	for _, f := range r.Information.FPCs {
		for _, p := range f.Pics {
			c.collectForPIC(ch, labelValues, &f, &p)
		}
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

func (c *fpcCollector) collectForPIC(ch chan<- prometheus.Metric, labelValues []string, fpc *FPC, pic *PIC) {
	picup := 0
	if pic.PicState == "Online" {
		picup = 1
	}

	l_pic := append(labelValues, strconv.Itoa(fpc.Slot), strconv.Itoa(pic.PicSlot), pic.PicType)
	ch <- prometheus.MustNewConstMetric(picstatusDesc, prometheus.GaugeValue, float64(picup), l_pic...)
}
