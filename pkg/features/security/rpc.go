// SPDX-License-Identifier: MIT

package security

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name       `xml:"rpc-reply"`
	Results routingEngines `xml:"multi-routing-engine-results"`
}

type routingEngines struct {
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name               string                     `xml:"re-name"`
	PerformanceSummary securityPerformanceSummary `xml:"performance-summary-information"`
}

type securityPerformanceSummary struct {
	PerformanceStatistics []securityPerformanceStatistics `xml:"performance-summary-statistics"`
}

type securityPerformanceStatistics struct {
	FPCNumber   int64 `xml:"fpc-number"`
	PICNumber   int64 `xml:"pic-number"`
	CPUUtil     int64 `xml:"spu-cpu-utilization"`
	MemoryUtil  int64 `xml:"spu-memory-utilization"`
	CurrentFlow int64 `xml:"spu-current-flow-session"`
	MaxFlow     int64 `xml:"spu-max-flow-session"`
	CurrentCP   int64 `xml:"spu-current-cp-session"`
	MaxCP       int64 `xml:"spu-max-cp-session"`
}

type singleEngineResult struct {
	XMLName            xml.Name                   `xml:"rpc-reply"`
	PerformanceSummary securityPerformanceSummary `xml:"performance-summary-information"`
}
