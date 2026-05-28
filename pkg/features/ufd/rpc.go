// SPDX-License-Identifier: MIT

package ufd

// Junos puts every <ufd-group> as a sibling inside a single
// <ufd-group-information> wrapper, regardless of how many groups exist.
// The published YANG model nominally describes ufd-group-information as a
// list, but real-world emission flattens to one wrapper with many children;
// the xpath below covers both shapes.
type result struct {
	Groups []ufdGroup `xml:"uplink-failure-detection-information>ufd-group-information>ufd-group"`
}

type ufdGroup struct {
	Name             string   `xml:"ufd-group-name"`
	FailureAction    string   `xml:"failure-action"`
	DebounceInterval float64  `xml:"debounce-interval"`
	Uplinks          []string `xml:"link-to-monitor-list>uplink-interface"`
	Downlinks        []string `xml:"link-to-disable-list>downlink-interface"`
	// DebounceTimeLeft appears only while a failure is being debounced.
	// Not yet exposed as a metric — need a debouncing-state sample to
	// validate the shape (the only triggered sample so far had no entry).
	DebounceTimeLeft []string `xml:"link-to-disable-list>debounce-time-left"`
}
