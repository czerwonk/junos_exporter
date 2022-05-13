package environment

import (
	"encoding/xml"
	"log"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_environment_"

var (
	temperaturesDesc *prometheus.Desc
	powerSupplyDesc  *prometheus.Desc
	pemDesc          *prometheus.Desc
	fanDesc          *prometheus.Desc
	dcVoltageDesc    *prometheus.Desc
	dcCurrentDesc    *prometheus.Desc
	dcPowerDesc      *prometheus.Desc
	dcLoadDesc       *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name", "item"}
	temperaturesDesc = prometheus.NewDesc(prefix+"item_temp", "Temperature of the air flowing past", l, nil)
	powerSupplyDesc = prometheus.NewDesc(prefix+"power_up", "Status of power supplies (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", append(l, "status"), nil)

	pemDesc = prometheus.NewDesc(prefix+"pem_state", "State of PEM module. 1 - Online, 2 - Present, 3 - Empty", append(l, "state"), nil)
	dcVoltageDesc = prometheus.NewDesc(prefix+"pem_voltage", "PEM voltage value", l, nil)
	dcCurrentDesc = prometheus.NewDesc(prefix+"pem_current", "PEM current value", l, nil)
	dcPowerDesc = prometheus.NewDesc(prefix+"pem_power_usage", "PEM power usage in W", l, nil)
	dcLoadDesc = prometheus.NewDesc(prefix+"pem_power_load_percent", "PEM power usage percent of total", l, nil)

	l = []string{"target", "re_name", "item", "fan_name"}
	fanDesc = prometheus.NewDesc(prefix+"pem_fanspeed", "Fan speed in RPM", l, nil)
}

type environmentCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &environmentCollector{}
}

// Name returns the name of the collector
func (*environmentCollector) Name() string {
	return "Environment"
}

// Describe describes the metrics
func (*environmentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperaturesDesc
	ch <- fanDesc
	ch <- dcPowerDesc
}

// Collect collects metrics from JunOS
func (c *environmentCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	c.environmentItems(client, ch, labelValues)
	c.environmentPEMItems(client, ch, labelValues)

	return nil
}

func (c *environmentCollector) environmentItems(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RpcReply{}

	statusValues := map[string]int{
		"OK":      1,
		"Testing": 2,
		"Failed":  3,
		"Absent":  4,
		"Present": 5,
	}

//	err := client.RunCommandAndParseWithParser("show chassis environment", func(b []byte) error {
	err := client.RunCommandAndParseWithParser("<get-environment-information/>", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return nil
	}

	// gather satellite data
	if client.Satellite {
		var y = RpcReply{}
		err = client.RunCommandAndParseWithParser("<get-chassis-environment-satellite-information/>", func(b []byte) error {
			if string(b[:]) == "\nerror: syntax error, expecting <command>: satellite\n" {
				log.Printf("system doesn't seem to have satellite enabled")
				return nil
			}

			return parseXML(b, &y)
		})
		if err != nil {
			return nil
		} else {
			// add satellite details (only if y.MultiRoutingEngineResults.RoutingEngine has elements)
			if len(y.MultiRoutingEngineResults.RoutingEngine) > 0 {
				x.MultiRoutingEngineResults.RoutingEngine[0].EnvironmentInformation.Items = append(x.MultiRoutingEngineResults.RoutingEngine[0].EnvironmentInformation.Items, y.MultiRoutingEngineResults.RoutingEngine[0].EnvironmentInformation.Items...)
			}
		}
	}

	for _, re := range x.MultiRoutingEngineResults.RoutingEngine {
		l := labelValues
		for _, item := range re.EnvironmentInformation.Items {
			l = append(labelValues, re.Name)
			if strings.Contains(item.Name, "Power Supply") || strings.Contains(item.Name, "PEM") {
				l = append(l, item.Name, item.Status)
				ch <- prometheus.MustNewConstMetric(powerSupplyDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), l...)
			} else if item.Temperature != nil {
				l = append(l, item.Name)
				ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, item.Temperature.Value, l...)
			}
		}
	}

	return nil
}

func (c *environmentCollector) environmentPEMItems(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RpcReply{}

	stateValues := map[string]int{
		"Online":  1,
		"Present": 2,
		"Empty":   3,
	}

//	err := client.RunCommandAndParseWithParser("show chassis environment pem", func(b []byte) error {
	err := client.RunCommandAndParseWithParser("<get-environment-pem-information/>", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, re := range x.MultiRoutingEngineResults.RoutingEngine {
		for _, e := range re.EnvironmentComponentInformation.EnvironmentComponentItem {
			l := append(labelValues, re.Name, e.Name)

			ch <- prometheus.MustNewConstMetric(pemDesc, prometheus.GaugeValue, float64(stateValues[e.State]), append(l, e.State)...)

			for _, f := range e.FanSpeedReading {
				rpms, err_ := strconv.ParseFloat(strings.TrimSuffix(f.FanSpeed, " RPM"), 64)
				if err_ != nil {
					return err_
				}
				ch <- prometheus.MustNewConstMetric(fanDesc, prometheus.GaugeValue, rpms, append(l, f.FanName)...)
			}
			voltage := 0.0
			if e.DcInformation.DcDetail.DcVoltage > 0 {
				voltage = e.DcInformation.DcDetail.DcVoltage
			}

			if e.DcInformation.DcDetail.Str3DcVoltage > 0 {

				voltage = e.DcInformation.DcDetail.Str3DcVoltage
			}

			if voltage > 0 {
				ch <- prometheus.MustNewConstMetric(dcVoltageDesc, prometheus.GaugeValue, voltage, l...)
				ch <- prometheus.MustNewConstMetric(dcCurrentDesc, prometheus.GaugeValue, e.DcInformation.DcDetail.DcCurrent, l...)
				ch <- prometheus.MustNewConstMetric(dcPowerDesc, prometheus.GaugeValue, e.DcInformation.DcDetail.DcPower, l...)
				ch <- prometheus.MustNewConstMetric(dcLoadDesc, prometheus.GaugeValue, e.DcInformation.DcDetail.DcLoad, l...)
			}
		}
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
		{
			Name:                            "N/A",
			EnvironmentComponentInformation: fi.EnvironmentComponentInformation,
			EnvironmentInformation:          fi.EnvironmentInformation,
		},
	}
	return nil
}
