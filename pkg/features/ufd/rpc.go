// SPDX-License-Identifier: MIT

package ufd

import "encoding/xml"

// Junos emits the UFD response in two shapes. On a single-RE non-VC device:
//
//   <rpc-reply>
//     <uplink-failure-detection-information>...</...>
//   </rpc-reply>
//
// On a multi-RE chassis or Virtual Chassis the body is wrapped:
//
//   <rpc-reply>
//     <multi-routing-engine-results>
//       <multi-routing-engine-item>
//         <re-name>fpc0</re-name>
//         <uplink-failure-detection-information>...</...>
//       </multi-routing-engine-item>
//       ...
//     </multi-routing-engine-results>
//   </rpc-reply>
//
// The QFX YANG declares the body as anyxml so the device chooses which case at
// runtime; both shapes must be handled. Same pattern as the alarm collector.

type singleEngineResult struct {
	XMLName xml.Name   `xml:"rpc-reply"`
	Groups  []ufdGroup `xml:"uplink-failure-detection-information>ufd-group-information>ufd-group"`
}

type multiEngineResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Engines struct {
		Items []routingEngine `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
}

type routingEngine struct {
	Name   string     `xml:"re-name"`
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
	// validate the shape.
	DebounceTimeLeft []string `xml:"link-to-disable-list>debounce-time-left"`
}
