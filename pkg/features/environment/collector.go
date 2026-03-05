// SPDX-License-Identifier: MIT

package environment

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

const prefix string = "junos_environment_"

var (
	temperaturesDesc *prometheus.Desc
	powerSupplyDesc  *prometheus.Desc
	fanStatusDesc    *prometheus.Desc
	fanAirflowDesc   *prometheus.Desc
	pemDesc          *prometheus.Desc
	fanDesc          *prometheus.Desc
	dcVoltageDesc    *prometheus.Desc
	dcCurrentDesc    *prometheus.Desc
	dcPowerDesc      *prometheus.Desc
	dcLoadDesc       *prometheus.Desc
	dcOutputDesc     *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name", "item"}
	temperaturesDesc = prometheus.NewDesc(prefix+"item_temp", "Temperature of the air flowing past", l, nil)
	powerSupplyDesc = prometheus.NewDesc(prefix+"power_up", "Status of power supplies (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", append(l, "status"), nil)
	fanStatusDesc = prometheus.NewDesc(prefix+"fan_up", "Status of fans (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", append(l, "status"), nil)
	fanAirflowDesc = prometheus.NewDesc(prefix+"fan_airflow_up", "Status of	fan airflows (1 OK, 2 Testing, 3 Failed, 4 Absent, 5 Present)", append(l, "status"), nil)

	pemDesc = prometheus.NewDesc(prefix+"pem_state", "State of PEM module. 1 - Online, 2 - Present, 3 - Empty", append(l, "state"), nil)
	dcVoltageDesc = prometheus.NewDesc(prefix+"pem_voltage", "PEM voltage value", l, nil)
	dcCurrentDesc = prometheus.NewDesc(prefix+"pem_current", "PEM current value", l, nil)
	dcPowerDesc = prometheus.NewDesc(prefix+"pem_power_usage", "PEM power usage in W", l, nil)
	dcLoadDesc = prometheus.NewDesc(prefix+"pem_power_load_percent", "PEM power usage percent of total", l, nil)
	dcOutputDesc = prometheus.NewDesc(prefix+"pem_dc_output", "PSM DC output status (1 OK, 0 not OK)", l, nil)

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
	ch <- dcOutputDesc
}

// Collect collects metrics from JunOS
func (c *environmentCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var v showVersionResult
	err := client.RunCommandAndParse("show version", &v)
	if err != nil {
		//return errors.Wrap(err, "failed to run command 'show version'")
		return errors.Wrap(err, "failed to run command 'show version'")
	}
    //QFX5220 have a  slightly different xml for environment information, so we need to check the product model before collecting environment metrics
    fmt.Printf("model is %s", v.SoftwareInformation.ProductModel)
    if strings.Contains(strings.ToLower(v.SoftwareInformation.ProductModel), "qfx5220"){
        c.environmentItemsForSomeSwitchModels(client, ch, labelValues)
        c.environmentPEMItemsForSomeSwitchModels(client, ch, labelValues)
	} else {
	    c.environmentItems(client, ch, labelValues)
        c.environmentPEMItems(client, ch, labelValues)
	}
	return nil
}

func (c *environmentCollector) environmentItems(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	x := multiEngineResult{}

	statusValues := map[string]int{
		"OK":      1,
		"Testing": 2,
		"Failed":  3,
		"Absent":  4,
		"Present": 5,
	}

	err := client.RunCommandAndParseWithParser("show chassis environment", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return nil
	}

	if client.IsSatelliteEnabled() {
		var y = multiEngineResult{}
		err = client.RunCommandAndParseWithParser("show chassis environment satellite", func(b []byte) error {
			if string(b[:]) == "\nerror: syntax error, expecting <command>: satellite\n" {
				log.Printf("system doesn't seem to have satellite enabled")
				return nil
			}

			return parseXML(b, &y)
		})
		if err != nil {
			return nil
		}

		if len(y.Results.RoutingEngines) > 0 {
			x.Results.RoutingEngines[0].EnvironmentInformation.Items = append(x.Results.RoutingEngines[0].EnvironmentInformation.Items, y.Results.RoutingEngines[0].EnvironmentInformation.Items...)
		}
	}

	for _, re := range x.Results.RoutingEngines {
		l := labelValues
		for _, item := range re.EnvironmentInformation.Items {
			l = append(labelValues, re.Name)
			if strings.Contains(item.Name, "Power Supply") || strings.Contains(item.Name, "PEM") || strings.Contains(item.Name, "PSM") {
				l = append(l, item.Name, item.Status)
				ch <- prometheus.MustNewConstMetric(powerSupplyDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), l...)
			} else if strings.Contains(item.Name, "Fan") {
				l = append(l, item.Name, item.Status)
				if strings.Contains(item.Name, "Airflow") {
					ch <- prometheus.MustNewConstMetric(fanAirflowDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), l...)
				} else {
					ch <- prometheus.MustNewConstMetric(fanStatusDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), l...)
				}
			} else if item.Temperature != nil {
				l = append(l, item.Name)
				ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, item.Temperature.Value, l...)
			}
		}
	}

	return nil
}



func (c *environmentCollector) environmentPEMItems(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = multiEngineResult{}

	stateValues := map[string]int{
		"Online":  1,
		"Present": 2,
		"Empty":   3,
		"Offline": 4,
	}

	err := client.RunCommandAndParseWithParser("show chassis environment pem", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		err := client.RunCommandAndParseWithParser("show chassis environment psm", func(b []byte) error {
			return parseXML(b, &x)
		})
		if err != nil {
			return err
		}
	}

	for _, re := range x.Results.RoutingEngines {
		for _, e := range re.EnvironmentComponentInformation.EnvironmentComponentItem {
			l := append(labelValues, re.Name, e.Name)

			ch <- prometheus.MustNewConstMetric(pemDesc, prometheus.GaugeValue, float64(stateValues[e.State]), append(l, e.State)...)

			for _, f := range e.FanSpeedReading {
				rpms, err := strconv.ParseFloat(strings.TrimSuffix(f.FanSpeed, " RPM"), 64)
				if err != nil {
					return fmt.Errorf("could not parse fan speed value to float: %s", f.FanSpeed)
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

func (c *environmentCollector) environmentItemsForSomeSwitchModels(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	x := environmentResultSomeSwitches{}

	statusValues := map[string]int{
		"OK":      1,
		"Testing": 2,
		"Failed":  3,
		"Absent":  4,
		"Present": 5,
	}

	err := client.RunCommandAndParseWithParser("show chassis environment", func(b []byte) error {
		return xml.Unmarshal(b, &x)
	})
	if err != nil {
		return nil
	}

	reName := "N/A"
	for _, item := range x.EnvironmentInformation.EnvironmentItem {
		l := append(labelValues, reName)

		if strings.Contains(item.Name, "Power Supply") || strings.Contains(item.Name, "PEM") || strings.Contains(item.Name, "PSM") {
			ch <- prometheus.MustNewConstMetric(powerSupplyDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), append(l, item.Name, item.Status)...)
		} else if strings.Contains(item.Name, "Fan") {
			if strings.Contains(item.Name, "Airflow") {
				ch <- prometheus.MustNewConstMetric(fanAirflowDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), append(l, item.Name, item.Status)...)
			} else {
				ch <- prometheus.MustNewConstMetric(fanStatusDesc, prometheus.GaugeValue, float64(statusValues[item.Status]), append(l, item.Name, item.Status)...)
			}
		} else if item.Temperature.Celsius != "" {
			tempVal, err := strconv.ParseFloat(item.Temperature.Celsius, 64)
			if err != nil {
				return fmt.Errorf("could not parse temperature value to float: %s", item.Temperature.Celsius)
			}
			ch <- prometheus.MustNewConstMetric(temperaturesDesc, prometheus.GaugeValue, tempVal, append(l, item.Name)...)
		}
	}

	return nil
}

func (c *environmentCollector) environmentPEMItemsForSomeSwitchModels(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	x := multiEngineResultSomeSwitches{}

	stateValues := map[string]int{
		"Online":  1,
		"Present": 2,
		"Empty":   3,
		"Offline": 4,
	}

	err := client.RunCommandAndParseWithParser("show chassis environment pem", func(b []byte) error {
		return xml.Unmarshal(b, &x)
	})
	if err != nil {
		err := client.RunCommandAndParseWithParser("show chassis environment psm", func(b []byte) error {
			return xml.Unmarshal(b, &x)
		})
		if err != nil {
			return err
		}
	}

	reName := "N/A"
	for _, item := range x.EnvironmentComponentInformation.EnvironmentComponentItem {
		l := append(labelValues, reName, item.Name)

		ch <- prometheus.MustNewConstMetric(pemDesc, prometheus.GaugeValue, float64(stateValues[item.State]), append(l, item.State)...)

		fan1Speed := item.PsmInformation.FanSpeedReadingPsm.Fan1Speed
		if fan1Speed != "" {
			rpms, err := strconv.ParseFloat(strings.TrimSuffix(fan1Speed, " RPM"), 64)
			if err != nil {
				return fmt.Errorf("could not parse fan speed value to float: %s", fan1Speed)
			}
			ch <- prometheus.MustNewConstMetric(fanDesc, prometheus.GaugeValue, rpms, append(l, item.PsmInformation.FanSpeedReadingPsm.Fan1Name)...)
		}

		dcOutputVal := 0.0
		fmt.Printf("value of dc output is %s", item.PsmInformation.PsmStatus.DcOutput)
		if strings.EqualFold(strings.ToLower(item.PsmInformation.PsmStatus.DcOutput), "ok") {
		    fmt.Printf("inside dc output metric")
			dcOutputVal = 1.0
		}
		ch <- prometheus.MustNewConstMetric(dcOutputDesc, prometheus.GaugeValue, dcOutputVal, l...)
	}

	return nil
}

func parseXML(b []byte, res *multiEngineResult) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := singleEngineResult{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.Results.RoutingEngines = []routingEngine{
		{
			Name:                            "N/A",
			EnvironmentComponentInformation: fi.EnvironmentComponentInformation,
			EnvironmentInformation:          fi.EnvironmentInformation,
		},
	}
	return nil
}
