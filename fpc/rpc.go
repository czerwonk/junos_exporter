package fpc

type FPCRpc struct {
	Information struct {
		FPCs []FPC `xml:"fpc"`
	} `xml:"fpc-information"`
}

type FPC struct {
	Slot        int    `xml:"slot"`
	State       string `xml:"state"`
	Temperature struct {
		Celsius int `xml:"celsius,attr"`
	} `xml:"temperature,omitempty"`
	MemoryDramSize    uint `xml:"memory-dram-size,omitempty"`
	MemoryRldramSize  uint `xml:"memory-rldram-size,omitempty"`
	MemoryDdrDramSize uint `xml:"memory-ddr-dram-size,omitempty"`
	MemorySdramSize   uint `xml:"memory-sdram-size,omitempty"`
	MemorySramSize    uint `xml:"memory-sram-size,omitempty"`
	StartTime         struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"start-time"`
	UpTime struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"up-time"`
	MaxPowerConsumption uint `xml:"max-power-consumption,omitempty"`
}
