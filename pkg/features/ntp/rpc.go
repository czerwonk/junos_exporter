// In rpc.go NUR folgendes belassen:
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
