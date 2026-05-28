// SPDX-License-Identifier: MIT

package ufd

type result struct {
	Information struct {
		GroupInfos []groupInfo `xml:"ufd-group-information"`
	} `xml:"uplink-failure-detection-information"`
}

type groupInfo struct {
	Group ufdGroup `xml:"ufd-group"`
}

type ufdGroup struct {
	Name             string   `xml:"ufd-group-name"`
	FailureAction    string   `xml:"failure-action"`
	DebounceInterval float64  `xml:"debounce-interval"`
	Uplinks          []string `xml:"link-to-monitor-list>uplink-interface"`
	Downlinks        []string `xml:"link-to-disable-list>downlink-interface"`
	// DebounceTimeLeft appears only while a failure is being debounced.
	// YANG: 'debounce-time-left' under link-to-disable-list. Not yet exposed
	// as a metric (need a triggered-state sample to validate the shape).
	DebounceTimeLeft []string `xml:"link-to-disable-list>debounce-time-left"`
}
