package fpc

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name                  `xml:"rpc-reply"`
	Results multiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type multiRoutingEngineResults struct {
	RoutingEngines []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name string `xml:"re-name"`
	FPCs fpcs   `xml:"fpc-information"`
}

type fpcs struct {
	FPC []fpc `xml:"fpc"`
}

type fpc struct {
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
	CPUTotal                uint `xml:"cpu-total,omitempty"`
	CPUInterrupt            uint `xml:"cpu-interrupt,omitempty"`
	CPU1minAVG              uint `xml:"cpu-1min-avg,omitempty"`
	CPU5minAVG              uint `xml:"cpu-5min-avg,omitempty"`
	CPU15minAVG             uint `xml:"cpu-15min-avg,omitempty"`
	MemoryHeapUtilization   uint `xml:"memory-heap-utilization,omitempty"`
	MemoryBufferUtilization uint `xml:"memory-buffer-utilization,omitempty"`
	StartTime               struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"start-time,omitempty"`
	UpTime struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"up-time,omitempty"`
	MaxPowerConsumption uint  `xml:"max-power-consumption,omitempty"`
	Pics                []pic `xml:"pic,omitempty"`
}
type pic struct {
	PicSlot  int    `xml:"pic-slot"`
	PicState string `xml:"pic-state"`
	PicType  string `xml:"pic-type"`
}

type singeEngineResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	FPCs    fpcs     `xml:"fpc-information"`
}
