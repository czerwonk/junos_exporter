// SPDX-License-Identifier: MIT

package alarm

import (
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_alarms_"

var (
	alarmsYellowCount *prometheus.Desc
	alarmsRedCount    *prometheus.Desc
	alarmDetails      *prometheus.Desc
)

func init() {
	l := []string{"target"}
	alarmsYellowCount = prometheus.NewDesc(prefix+"yellow_count", "Number of yellow alarms (not silenced)", l, nil)
	alarmsRedCount = prometheus.NewDesc(prefix+"red_count", "Number of red alarms (not silenced)", l, nil)
	l = append(l, "class", "type", "description")
	alarmDetails = prometheus.NewDesc(prefix+"set", "Alarm active with the details provided in labels", l, nil)
}

type alarmCollector struct {
	filter *regexp.Regexp
}

// NewCollector creates a new collector
func NewCollector(alarmsFilter string) collector.RPCCollector {
	c := &alarmCollector{}

	if len(alarmsFilter) > 0 {
		c.filter = regexp.MustCompile(alarmsFilter)
	}

	return c
}

// Name returns the name of the collector
func (*alarmCollector) Name() string {
	return "Alarm"
}

// Describe describes the metrics
func (*alarmCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- alarmsYellowCount
	ch <- alarmsRedCount
}

// Collect collects metrics from JunOS
func (c *alarmCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	counter, alarms, err := c.alarmCounter(client)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(alarmsYellowCount, prometheus.GaugeValue, counter.yellow, labelValues...)
	ch <- prometheus.MustNewConstMetric(alarmsRedCount, prometheus.GaugeValue, counter.red, labelValues...)
	if alarms != nil {
		for _, alarm := range *alarms {
			localLabelvalues := append(labelValues, alarm.Class, alarm.Type, alarm.Description)
			ch <- prometheus.MustNewConstMetric(alarmDetails, prometheus.GaugeValue, 1, localLabelvalues...)
		}
	}

	return nil
}

func (c *alarmCollector) alarmCounter(client collector.Client) (*alarmCounter, *[]details, error) {
	red := 0
	yellow := 0

	cmds := []string{
		"show system alarms",
		"show chassis alarms",
	}

	var alarms []details

	messages := make(map[string]interface{})
	for _, cmd := range cmds {
		var a = multiEngineResult{}
		err := client.RunCommandAndParseWithParser(cmd, func(b []byte) error {
			return parseXML(b, &a)
		})
		if err != nil {
			return nil, nil, err
		}

		for _, engine := range a.RoutingEngines {
			for _, d := range engine.AlarmInfo.Details {
				if _, found := messages[d.Description]; found {
					continue
				}

				alarms = append(alarms, d)

				if c.shouldFilterAlarm(&d) {
					continue
				}

				if d.Class == "Major" {
					red++
				} else if d.Class == "Minor" {
					yellow++
				}

				messages[d.Description] = nil
			}
		}
	}

	return &alarmCounter{red: float64(red), yellow: float64(yellow)}, &alarms, nil
}

func (c *alarmCollector) shouldFilterAlarm(a *details) bool {
	if c.filter == nil {
		return false
	}

	return c.filter.MatchString(a.Description) || c.filter.MatchString(a.Type)
}

func parseXML(b []byte, res *multiEngineResult) error {
	if strings.Contains(string(b), "<multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	se := alarmInformation{}

	err := xml.Unmarshal(b, &se)
	if err != nil {
		return err
	}

	res.RoutingEngines = []routingEngine{
		{
			Name:      "N/A",
			AlarmInfo: se,
		},
	}

	return nil
}
