package fpc

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name string `xml:"re-name"`
	FPCs FPCs   `xml:"fpc-information"`
}

type FPCs struct {
	FPC []FPC `xml:"fpc"`
}

type FPC struct {
	Slot        int    `xml:"slot"`
	State       string `xml:"state"`
	Description string `xml:"description,omitempty"`
	Temperature struct {
		Celsius int `xml:"celsius,attr"`
	} `xml:"temperature,omitempty"`
	MemoryDramSize          uint `xml:"memory-dram-size,omitempty"`
	MemoryRldramSize        uint `xml:"memory-rldram-size,omitempty"`
	MemoryDdrDramSize       uint `xml:"memory-ddr-dram-size,omitempty"`
	MemorySdramSize         uint `xml:"memory-sdram-size,omitempty"`
	MemorySramSize          uint `xml:"memory-sram-size,omitempty"`
	CpuTotal                uint `xml:"cpu-total,omitempty"`
	CpuInterrupt            uint `xml:"cpu-interrupt,omitempty"`
	Cpu1min_avg             uint `xml:"cpu-1min-avg,omitempty"`
	Cpu5min_avg             uint `xml:"cpu-5min-avg,omitempty"`
	Cpu15min_avg            uint `xml:"cpu-15min-avg,omitempty"`
	MemoryHeapUtilization   uint `xml:"memory-heap-utilization,omitempty"`
	MemoryBufferUtilization uint `xml:"memory-buffer-utilization,omitempty"`
	StartTime               struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"start-time,omitempty"`
	UpTime struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"up-time,omitempty"`
	MaxPowerConsumption uint  `xml:"max-power-consumption,omitempty"`
	Pics                []PIC `xml:"pic,omitempty"`
}
type PIC struct {
	PicSlot  int    `xml:"pic-slot"`
	PicState string `xml:"pic-state"`
	PicType  string `xml:"pic-type"`
}

type RpcReplyNoRE struct {
	XMLName xml.Name `xml:"rpc-reply"`
	FPCs    FPCs     `xml:"fpc-information"`
}
