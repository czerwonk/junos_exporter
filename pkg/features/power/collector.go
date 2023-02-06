package power

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_power_"

var (
	capacityActualDesc       *prometheus.Desc
	capacityMaxDesc          *prometheus.Desc
	capacityAllocatedDesc    *prometheus.Desc
	capacityRemainingDesc    *prometheus.Desc
	capacityActualUsageDesc  *prometheus.Desc
	capacitySysActualDesc    *prometheus.Desc
	capacitySysMaxDesc       *prometheus.Desc
	capacitySysRemainingDesc *prometheus.Desc
	pemPowerStateDesc        *prometheus.Desc
	dcPowerDesc              *prometheus.Desc
	dcCurrentDesc            *prometheus.Desc
	dcVoltageDesc            *prometheus.Desc
	dcLoadDesc               *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name"}
	capacitySysActualDesc = prometheus.NewDesc(prefix+"capacity_sys_actual_usage", "Actual power usage for the system, in watts", l, nil)
	capacitySysMaxDesc = prometheus.NewDesc(prefix+"capacity_sys_max", "Maximum power capacity for the system, in watts", l, nil)
	capacitySysRemainingDesc = prometheus.NewDesc(prefix+"capacity_sys_remaining", "Remaining capacity for the system, in watts", l, nil)

	l = []string{"target", "re_name", "zone"}
	capacityActualDesc = prometheus.NewDesc(prefix+"capacity_actual", "Power capacity applicable for the zone, in watts", l, nil)
	capacityMaxDesc = prometheus.NewDesc(prefix+"capacity_max", "Maximum power capacity applicable for the zone, in watts", l, nil)
	capacityAllocatedDesc = prometheus.NewDesc(prefix+"capacity_allocated", "Actual capacity allocated for the zone, in watts", l, nil)
	capacityRemainingDesc = prometheus.NewDesc(prefix+"capacity_remaining", "Remaining capacity for the zone, in watts", l, nil)
	capacityActualUsageDesc = prometheus.NewDesc(prefix+"capacity_actual_usage", "Actual power usage for the zone, in watts", l, nil)

	l = []string{"target", "re_name", "name", "zone"}
	dcPowerDesc = prometheus.NewDesc(prefix+"pem_power_usage", "PEM power usage in W", l, nil)
	dcCurrentDesc = prometheus.NewDesc(prefix+"pem_current", "PEM current value", l, nil)
	dcVoltageDesc = prometheus.NewDesc(prefix+"pem_voltage", "PEM voltage value", l, nil)
	dcLoadDesc = prometheus.NewDesc(prefix+"pem_power_load_percent", "PEM power usage percent of total", l, nil)

	pemPowerStateDesc = prometheus.NewDesc(prefix+"pem_power_state", "PEM power state. 1 - Online, 2 - Present, 3 - Empty", append(l, "state"), nil)
}

type powerCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &powerCollector{}
}

// Name returns the name of the collector
func (*powerCollector) Name() string {
	return "Power"
}

// Describe describes the metrics
func (*powerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pemPowerStateDesc
	ch <- capacityActualDesc
	ch <- capacityMaxDesc
	ch <- capacityAllocatedDesc
	ch <- capacityRemainingDesc
	ch <- capacityActualUsageDesc
	ch <- capacitySysActualDesc
	ch <- capacitySysMaxDesc
	ch <- capacitySysRemainingDesc
	ch <- dcPowerDesc
	ch <- dcCurrentDesc
	ch <- dcVoltageDesc
	ch <- dcLoadDesc
}

// Collect collects metrics from JunOS
func (c *powerCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	stateValues := map[string]int{
		"Online":  1,
		"Present": 2,
		"Empty":   3,
	}

	var x = multiRoutingEngineResult{}
	err := client.RunCommandAndParseWithParser("show chassis power", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, re := range x.Results.RoutingEngine {
		l := append(labelValues, re.Name)

		if re.PowerUsageInformation.PowerUsageSystem.CapacitySysActual > 0 {
			ch <- prometheus.MustNewConstMetric(capacitySysActualDesc, prometheus.GaugeValue, float64(re.PowerUsageInformation.PowerUsageSystem.CapacitySysActual), l...)
		}
		if re.PowerUsageInformation.PowerUsageSystem.CapacitySysMax > 0 {
			ch <- prometheus.MustNewConstMetric(capacitySysMaxDesc, prometheus.GaugeValue, float64(re.PowerUsageInformation.PowerUsageSystem.CapacitySysMax), l...)
		}
		if re.PowerUsageInformation.PowerUsageSystem.CapacitySysRemaining > 0 {
			ch <- prometheus.MustNewConstMetric(capacitySysRemainingDesc, prometheus.GaugeValue, float64(re.PowerUsageInformation.PowerUsageSystem.CapacitySysRemaining), l...)
		}

		for _, z := range re.PowerUsageInformation.PowerUsageSystem.PowerUsageZoneInformation {
			zl := append(l, z.Zone)

			ch <- prometheus.MustNewConstMetric(capacityActualDesc, prometheus.GaugeValue, float64(z.CapacityActual), zl...)
			ch <- prometheus.MustNewConstMetric(capacityMaxDesc, prometheus.GaugeValue, float64(z.CapacityMax), zl...)
			ch <- prometheus.MustNewConstMetric(capacityAllocatedDesc, prometheus.GaugeValue, float64(z.CapacityAllocated), zl...)
			ch <- prometheus.MustNewConstMetric(capacityRemainingDesc, prometheus.GaugeValue, float64(z.CapacityRemaining), zl...)
			ch <- prometheus.MustNewConstMetric(capacityActualUsageDesc, prometheus.GaugeValue, float64(z.CapacityActualUsage), zl...)
		}
		for _, p := range re.PowerUsageInformation.PowerUsageItem {
			pl := append(l, p.Name, p.DcOutputDetail.Zone)

			ch <- prometheus.MustNewConstMetric(dcPowerDesc, prometheus.GaugeValue, float64(p.DcOutputDetail.DcPower), pl...)
			ch <- prometheus.MustNewConstMetric(dcCurrentDesc, prometheus.GaugeValue, float64(p.DcOutputDetail.DcCurrent), pl...)
			ch <- prometheus.MustNewConstMetric(dcVoltageDesc, prometheus.GaugeValue, float64(p.DcOutputDetail.DcVoltage), pl...)
			ch <- prometheus.MustNewConstMetric(dcLoadDesc, prometheus.GaugeValue, float64(p.DcOutputDetail.DcLoad), pl...)

			ch <- prometheus.MustNewConstMetric(pemPowerStateDesc, prometheus.GaugeValue, float64(stateValues[p.State]), append(pl, p.State)...)
		}
	}

	return nil
}

func parseXML(b []byte, res *multiRoutingEngineResult) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := singleRoutingEngineResult{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.Results.RoutingEngine = []routingEngine{
		{
			Name:                  "N/A",
			PowerUsageInformation: fi.PowerUsageInformation,
		},
	}
	return nil
}
