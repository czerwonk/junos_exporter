package ntp

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_ntp_"

var (
	ntpStratumDesc   *prometheus.Desc
	ntpOffsetDesc    *prometheus.Desc
	ntpSysJitterDesc *prometheus.Desc
	ntpClkJitterDesc *prometheus.Desc
	ntpRootDelayDesc *prometheus.Desc
	ntpLeapDesc      *prometheus.Desc
	ntpPrecisionDesc *prometheus.Desc
	ntpPollDesc      *prometheus.Desc
)

func init() {
	l := []string{"target", "server"}
	ntpStratumDesc = prometheus.NewDesc(prefix+"stratum", "NTP stratum level (0: reference clock, 1-15: hops to refernce clock, 16: not syncronized)", l, nil)
	ntpOffsetDesc = prometheus.NewDesc(prefix+"offset", "Time offset in msec", l, nil)
	ntpSysJitterDesc = prometheus.NewDesc(prefix+"system_jitter", "System jitter in msec", l, nil)
	ntpClkJitterDesc = prometheus.NewDesc(prefix+"clock_jitter", "Clock jitter in msec", l, nil)
	ntpRootDelayDesc = prometheus.NewDesc(prefix+"root_delay", "Root delay in msec", l, nil)
	ntpLeapDesc = prometheus.NewDesc(prefix+"leap", "Leap indicator (00=ok, 01: last minute with 61 seconds, 10: last minute with 59 seconds, 11: not syncronized)", l, nil)
	ntpPrecisionDesc = prometheus.NewDesc(prefix+"precision", "Clock precision (should be -20 to -22)", l, nil)
	ntpPollDesc = prometheus.NewDesc(prefix+"poll_interval", "Poll interval in seconds", l, nil)
}

type ntpCollector struct{}

func NewCollector() collector.RPCCollector {
	return &ntpCollector{}
}

func (c *ntpCollector) Name() string {
	return "ntp"
}

func (c *ntpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- ntpStratumDesc
	ch <- ntpOffsetDesc
	ch <- ntpSysJitterDesc
	ch <- ntpClkJitterDesc
	ch <- ntpRootDelayDesc
	ch <- ntpLeapDesc
	ch <- ntpPrecisionDesc
	ch <- ntpPollDesc
}

func (c *ntpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var reply rpcReply

	err := client.RunCommandAndParse("show ntp status | display xml", &reply)
	if err != nil {
		return errors.Wrap(err, "failed to execute NTP command")
	}

	// Hier wird das parseResult direkt aus den Metriken erzeugt
	metrics := parseNTPOutput(reply.Output.Text)
	if len(metrics) == 0 {
		return errors.New("no NTP metrics parsed")
	}

	tc := mustParseFloat(metrics["tc"])
	if tc == 0 {
		tc = 10
	}

	// Konvertierung der Metriken in parseResult
	result := &parseResult{
		AssocID:      metrics["associd"],
		Stratum:      mustParseFloat(metrics["stratum"]),
		RefID:        metrics["refid"],
		Offset:       mustParseFloat(metrics["offset"]),
		SysJitter:    mustParseFloat(metrics["sys_jitter"]),
		ClkJitter:    mustParseFloat(metrics["clk_jitter"]),
		RootDelay:    mustParseFloat(metrics["rootdelay"]),
		Leap:         metrics["leap"],
		Precision:    mustParseFloat(metrics["precision"]),
		PollInterval: math.Pow(2, tc), // 2^10 = 1024
	}

	server := result.RefID
	if server == "" {
		server = "unknown"
	}

	labels := append(labelValues, server)

	exportMetric(ch, ntpStratumDesc, result.Stratum, labels)
	exportMetric(ch, ntpOffsetDesc, result.Offset, labels)
	exportMetric(ch, ntpSysJitterDesc, result.SysJitter, labels)
	exportMetric(ch, ntpClkJitterDesc, result.ClkJitter, labels)
	exportMetric(ch, ntpRootDelayDesc, result.RootDelay, labels)
	exportMetric(ch, ntpLeapDesc, parseLeap(result.Leap), labels)
	exportMetric(ch, ntpPrecisionDesc, result.Precision, labels)
	exportMetric(ch, ntpPollDesc, result.PollInterval, labels)

	return nil
}

func exportMetric(ch chan<- prometheus.Metric, desc *prometheus.Desc, value float64, labels []string) {
	ch <- prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		value,
		labels...,
	)
}

func parseLeap(leap string) float64 {
	leap = strings.TrimSpace(leap)
	switch leap {
	case "00":
		return 0
	case "01":
		return 1
	case "10":
		return 2
	case "11":
		return 3
	default:
		return -1
	}
}

func mustParseFloat(s string) float64 {
	s = strings.Trim(s, "+,\" ") // Kommas entfernen
	if s == "" || s == "-" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Parse error for '%s': %v", s, err)
		return 0
	}
	return f
}
