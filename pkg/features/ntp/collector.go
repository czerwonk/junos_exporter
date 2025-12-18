package ntp

import (
	"log"
	"math"
	"regexp"
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

// Used for old format of output
func parseNTPOutput(output string) map[string]string {
	metrics := make(map[string]string)

	re := regexp.MustCompile(`(\w+)=("[^"]*"|\S+?)(?:,|\s|$)`)
	matches := re.FindAllStringSubmatch(output, -1)

	for _, m := range matches {
		if len(m) == 3 {
			key := strings.ToLower(m[1])
			value := strings.Trim(m[2], "\", ")
			metrics[key] = value
		}
	}

	// DEBUG: Number of found metrics
	log.Printf("NTP parsed %d metrics from output length %d", len(metrics), len(output))

	return metrics
}

// Used for old format of output
func parseTextLeap(output string) float64 {
	switch {
	case strings.Contains(output, "leap_none"):
		return 0
	case strings.Contains(output, "leap_addsec"):
		return 1
	case strings.Contains(output, "leap_delsec"):
		return 2
	case strings.Contains(output, "leap_alarm"):
		return 3
	default:
		return -1
	}
}

// Used for new format of output
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

// Used for old and new format of output
func mustParseFloat(s string) float64 {
	s = strings.Trim(s, "+,\" ")
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

// Used for old and new format of output
func exportMetric(ch chan<- prometheus.Metric, desc *prometheus.Desc, value float64, labels []string) {
	ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, value, labels...)
}


func (c *ntpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var reply rpcReply

	err := client.RunCommandAndParse("show ntp status", &reply)
	if err != nil {
		return errors.Wrap(err, "failed to execute NTP command")
	}

	labels := labelValues

	// If output of command follows new XML design containing extra value for stratum
	if reply.NtpStatus.Stratum != "" {
		server := reply.NtpStatus.Refid
		if server == "" {
			server = "unknown"
		}
		labels = append(labels, server)

		exportMetric(ch, ntpStratumDesc, mustParseFloat(reply.NtpStatus.Stratum), labels)
		exportMetric(ch, ntpOffsetDesc, mustParseFloat(reply.NtpStatus.Offset), labels)
		exportMetric(ch, ntpSysJitterDesc, mustParseFloat(reply.NtpStatus.SysJitter), labels)
		exportMetric(ch, ntpRootDelayDesc, mustParseFloat(reply.NtpStatus.Rootdelay), labels)
		exportMetric(ch, ntpPrecisionDesc, mustParseFloat(reply.NtpStatus.Precision), labels)
		exportMetric(ch, ntpClkJitterDesc, mustParseFloat(reply.NtpStatus.ClkJitter), labels)
		exportMetric(ch, ntpLeapDesc, parseLeap(reply.NtpStatus.Leap), labels)

		// tc → Poll-Intervall is a logarithm value and needs scepial treatment
		tc := mustParseFloat(reply.NtpStatus.AssocID)
		if tc == 0 {
			tc = 10
		}
		exportMetric(ch, ntpPollDesc, math.Pow(2, tc), labels)

		log.Printf("NTP(XML) metrics exported for %s", server)
		return nil
	}

	// command output follows old design containg all values within one output
	log.Printf("NTP reply length: %d", len(reply.Output.Text))

	if reply.Output.Text == "" {
		return errors.New("no ntp output or ntp-status found")
	}

	metrics := parseNTPOutput(reply.Output.Text)
	if len(metrics) == 0 {
		return errors.New("no NTP metrics parsed")
	}

	server := metrics["refid"]
	if server == "" {
		server = "unknown"
	}
	labels = append(labels, server)

	exportMetric(ch, ntpStratumDesc, mustParseFloat(metrics["stratum"]), labels)
	exportMetric(ch, ntpOffsetDesc, mustParseFloat(metrics["offset"]), labels)
	exportMetric(ch, ntpSysJitterDesc, mustParseFloat(metrics["jitter"]), labels)
	exportMetric(ch, ntpRootDelayDesc, mustParseFloat(metrics["rootdelay"]), labels)
	exportMetric(ch, ntpPrecisionDesc, mustParseFloat(metrics["precision"]), labels)
	exportMetric(ch, ntpClkJitterDesc, mustParseFloat(metrics["jitter"]), labels)
	exportMetric(ch, ntpLeapDesc, parseTextLeap(reply.Output.Text),	labels)

	tc := mustParseFloat(metrics["poll"])
	exportMetric(ch, ntpPollDesc, math.Pow(2, tc), labels)

	log.Printf("NTP(TEXT) metrics exported for %s", server)
	return nil
}



// HELPER für DEBUG truncate
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
