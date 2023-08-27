// SPDX-License-Identifier: MIT

package rpm

import (
	"encoding/xml"
)

type result struct {
	XMLName xml.Name `xml:"probe-results"`
	Probes  []probe  `xml:"probe-test-results"`
}

type probe struct {
	Owner     string `xml:"owner"`
	Name      string `xml:"test-name"`
	Address   string `xml:"target-address"`
	Type      string `xml:"probe-type"`
	Interface string `xml:"destination-interface"`
	Size      int64  `xml:"test-size"`
	Last      struct {
		Results genericResults `xml:"probe-test-generic-results"`
	} `xml:"probe-last-test-results"`
	Global struct {
		Results genericResults `xml:"probe-test-generic-results"`
	} `xml:"probe-test-global-results"`
}

type genericResults struct {
	Scope       string  `xml:"results-scope"`
	Sent        int64   `xml:"probes-sent"`
	Responses   int64   `xml:"probe-responses"`
	LossPercent float64 `xml:"loss-percentage"`
	RTT         struct {
		Summary struct {
			Samples int64 `xml:"samples"`
			Min     int64 `xml:"min-delay"`
			Max     int64 `xml:"max-delay"`
			Avg     int64 `xml:"avg-delay"`
			Jitter  int64 `xml:"jitter-delay"`
			Stddev  int64 `xml:"stddev-delay"`
			Sum     int64 `xml:"sum-delay"`
		} `xml:"probe-summary-results"`
	} `xml:"probe-test-rtt"`
}
