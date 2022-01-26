package security

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name               string                     `xml:"re-name"`
	PerformanceSummary SecurityPerformanceSummary `xml:"performance-summary-information"`
}

type SecurityPerformanceSummary struct {
	PerformanceStatistics []SecurityPerformanceStatistics `xml:"performance-summary-statistics"`
}

type SecurityPerformanceStatistics struct {
	FPCNumber   int64 `xml:"fpc-number"`
	PICNumber   int64 `xml:"pic-number"`
	CPUUtil     int64 `xml:"spu-cpu-utilization"`
	MemoryUtil  int64 `xml:"spu-memory-utilization"`
	CurrentFlow int64 `xml:"spu-current-flow-session"`
	MaxFlow     int64 `xml:"spu-max-flow-session"`
	CurrentCP   int64 `xml:"spu-current-cp-session"`
	MaxCP       int64 `xml:"spu-max-cp-session"`
}

type RpcReplyNoRE struct {
	XMLName            xml.Name                   `xml:"rpc-reply"`
	PerformanceSummary SecurityPerformanceSummary `xml:"performance-summary-information"`
}
