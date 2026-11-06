// SPDX-License-Identifier: MIT
package ntp

import (
	"encoding/xml"
	"regexp"
	"strings"
)

type rpcReply struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Output  struct {
		Text string `xml:",chardata"`
	} `xml:"output"`
	NTPStatus *ntpStatus `xml:"ntp-status"`
}

// ntpStatus is the structured reply of newer JunOS releases.
type ntpStatus struct {
	AssocID   string `xml:"associd"`
	Leap      string `xml:"leap"`
	Stratum   string `xml:"stratum"`
	Precision string `xml:"precision"`
	RootDelay string `xml:"rootdelay"`
	RefID     string `xml:"refid"`
	TC        string `xml:"tc"`
	Offset    string `xml:"offset"`
	SysJitter string `xml:"sys-jitter"`
	ClkJitter string `xml:"clk-jitter"`
}

// metrics returns the NTP variables keyed like the text output of older releases.
func (r *rpcReply) metrics() map[string]string {
	metrics := parseNTPOutput(r.Output.Text)
	if len(metrics) > 0 || r.NTPStatus == nil {
		return metrics
	}

	for k, v := range map[string]string{
		"associd":    r.NTPStatus.AssocID,
		"leap":       r.NTPStatus.Leap,
		"stratum":    r.NTPStatus.Stratum,
		"precision":  r.NTPStatus.Precision,
		"rootdelay":  r.NTPStatus.RootDelay,
		"refid":      r.NTPStatus.RefID,
		"tc":         r.NTPStatus.TC,
		"offset":     r.NTPStatus.Offset,
		"sys_jitter": r.NTPStatus.SysJitter,
		"clk_jitter": r.NTPStatus.ClkJitter,
	} {
		if v = strings.TrimSpace(v); v != "" {
			metrics[k] = v
		}
	}
	return metrics
}

type parseResult struct {
	AssocID      string
	Stratum      float64
	RefID        string
	Offset       float64
	SysJitter    float64
	ClkJitter    float64
	RootDelay    float64
	Leap         string
	Precision    float64
	PollInterval float64
}

func parseNTPOutput(output string) map[string]string {
	re := regexp.MustCompile(`(\w+)=("[^"]+"|\S+)`)
	matches := re.FindAllStringSubmatch(output, -1)

	metrics := make(map[string]string)
	for _, m := range matches {
		key := m[1]
		value := strings.Trim(m[2], "\", ")
		metrics[key] = value
	}
	return metrics
}
