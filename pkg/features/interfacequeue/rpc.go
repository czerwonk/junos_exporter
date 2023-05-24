// SPDX-License-Identifier: MIT

package interfacequeue

type result struct {
	InterfaceInformation struct {
		Interfaces []physicalInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type physicalInterface struct {
	Name          string `xml:"name"`
	Description   string `xml:"description"`
	QueueCounters struct {
		Queues []queue `xml:"queue"`
	} `xml:"queue-counters"`
}

type queue struct {
	Number               string `xml:"queue-number"`
	QueuedPackets        uint64 `xml:"queue-counters-queued-packets"`
	QueuedBytes          uint64 `xml:"queue-counters-queued-bytes"`
	TransferedPackets    uint64 `xml:"queue-counters-trans-packets"`
	TransferedBytes      uint64 `xml:"queue-counters-trans-bytes"`
	RateLimitDropPackets uint64 `xml:"queue-counters-rate-limit-drop-packets"`
	RateLimitDropBytes   uint64 `xml:"queue-counters-rate-limit-drop-bytes"`
	RedPackets           uint64 `xml:"queue-counters-red-packets"`
	RedBytes             uint64 `xml:"queue-counters-red-bytes"`
	RedPacketsLow        uint64 `xml:"queue-counters-red-packets-low"`
	RedBytesLow          uint64 `xml:"queue-counters-red-bytes-low"`
	RedPacketsMediumLow  uint64 `xml:"queue-counters-red-packets-medium-low"`
	RedBytesMediumLow    uint64 `xml:"queue-counters-red-bytes-medium-low"`
	RedPacketsMediumHigh uint64 `xml:"queue-counters-red-packets-medium-high"`
	RedBytesMediumHigh   uint64 `xml:"queue-counters-red-bytes-medium-high"`
	RedPacketsHigh       uint64 `xml:"queue-counters-red-packets-high"`
	RedBytesHigh         uint64 `xml:"queue-counters-red-bytes-high"`
	TailDropPackets      uint64 `xml:"queue-counters-tail-drop-packets"`
	TotalDropPackets     uint64 `xml:"queue-counters-total-drop-packets"`
	TotalDropBytes       uint64 `xml:"queue-counters-total-drop-bytes"`
}
