// SPDX-License-Identifier: MIT

package interfaces

type interfaceStats struct {
	Name                    string
	AdminStatus             bool
	OperStatus              bool
	ErrorStatus             bool
	Description             string
	Mac                     string
	IsPhysical              bool
	Speed                   string
	BPDUError               bool
	ReceiveBytes            float64
	ReceivePackets          float64
	ReceiveErrors           float64
	ReceiveDrops            float64
	TransmitBytes           float64
	TransmitPackets         float64
	TransmitErrors          float64
	TransmitDrops           float64
	IPv6ReceiveBytes        float64
	IPv6ReceivePackets      float64
	IPv6TransmitBytes       float64
	IPv6TransmitPackets     float64
	LastFlapped             float64
	ReceiveUnicasts         float64
	ReceiveBroadcasts       float64
	ReceiveMulticasts       float64
	ReceiveCRCErrors        float64
	TransmitUnicasts        float64
	TransmitBroadcasts      float64
	TransmitMulticasts      float64
	TransmitCRCErrors       float64
	FecCcwCount             float64
	FecNccwCount            float64
	FecCcwErrorRate         float64
	FecNccwErrorRate        float64
	ReceiveOversizedFrames  float64
	ReceiveJabberFrames     float64
	ReceiveFragmentFrames   float64
	ReceiveVlanTaggedFrames float64
	ReceiveCodeViolations   float64
	ReceiveTotalErrors      float64
	TransmitTotalErrors     float64
	MTU                     string
}
