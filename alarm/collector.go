package alarm

import (
	"regexp"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_alarms_"

var (
	alarmsYellowCount *prometheus.Desc
	alarmsRedCount    *prometheus.Desc
)

func init() {
	l := []string{"target"}
	alarmsYellowCount = prometheus.NewDesc(prefix+"yellow_count", "Number of yollow alarms (not silenced)", l, nil)
	alarmsRedCount = prometheus.NewDesc(prefix+"red_count", "Number of red alarms (not silenced)", l, nil)
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
func (c *alarmCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	counter, err := c.alarmCounter(client)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(alarmsYellowCount, prometheus.GaugeValue, counter.YellowCount, labelValues...)
	ch <- prometheus.MustNewConstMetric(alarmsRedCount, prometheus.GaugeValue, counter.RedCount, labelValues...)

	return nil
}

func (c *alarmCollector) alarmCounter(client *rpc.Client) (*AlarmCounter, error) {
	red := 0
	yellow := 0

	cmds := []string{
		"show system alarms",
		"show chassis alarms",
	}

	messages := make(map[string]interface{})
	for _, cmd := range cmds {
		var a = AlarmRpc{}
		err := client.RunCommandAndParse(cmd, &a)
		if err != nil {
			return nil, err
		}

		for _, d := range a.Information.Details {
			if _, found := messages[d.Description]; found {
				continue
			}

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

	return &AlarmCounter{RedCount: float64(red), YellowCount: float64(yellow)}, nil
}

func (c *alarmCollector) shouldFilterAlarm(a *AlarmDetails) bool {
	if c.filter == nil {
		return false
	}

	return c.filter.MatchString(a.Description) || c.filter.MatchString(a.Type)
}
