// SPDX-License-Identifier: MIT

package mnha

import "encoding/xml"

// detailResult is the response to "show chassis high-availability information detail".
type detailResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Info    struct {
		NodeInfo            nodeInfo        `xml:"chassis-high-availability-node-info"`
		ColdSyncMonitoring  coldSyncMonitor `xml:"cold-synchronization-monitoring-information"`
		SPUMonitoring       spuMonitoring   `xml:"spu-monitoring-information"`
		LoopbackInformation loopbackInfo    `xml:"loopback-information"`
		HardwareMonitoring  hwMonitoring    `xml:"high-availability-hardware-monitoring-information"`
		DefaultSRG          srgBlock        `xml:"chassis-high-availability-default-srg-info"`
	} `xml:"chassis-high-availability-detail-information"`
}

type nodeInfo struct {
	NodeStatus string   `xml:"node-status"`
	GridID     int      `xml:"grid-id"`
	LocalID    int      `xml:"local-id"`
	LocalIP    string   `xml:"local-ip"`
	LocalFwdIP string   `xml:"local-fwd-ip"`
	Peer       peerInfo `xml:"chassis-high-availability-peer-info-detail"`
}

type peerInfo struct {
	PeerID            int             `xml:"high-availability-peer-id"`
	PeerIPAddress     string          `xml:"high-availability-peer-ip-address"`
	LocalInterface    string          `xml:"high-availability-local-interface"`
	RoutingInstance   string          `xml:"high-availability-peer-rt-instance"`
	EncryptionStatus  string          `xml:"high-availability-peer-encryption-status"`
	BFDStatus         string          `xml:"high-availability-peer-bfd-status"`
	DetectTime        string          `xml:"high-availability-peer-detect-time"`
	ColdSyncStatus    string          `xml:"cold-sync-status"`
	PeerFwdIPAddress  string          `xml:"high-availability-peer-fwd-ip-address"`
	LocalFwdInterface string          `xml:"high-availability-local-fwd-interface"`
	PeerFwdBFDStatus  string          `xml:"high-availability-peer-fwd-bfd-status"`
	PacketStats       peerPacketStats `xml:"high-availability-peer-packet-stats"`
}

type peerPacketStats struct {
	SendErrCount int                  `xml:"packet-send-err-cnt"`
	RecvErrCount int                  `xml:"packet-recv-err-cnt"`
	Stats        []peerPacketTypeStat `xml:"peer-packet-send-rcv-stats>peer-packet-stats-list"`
}

type peerPacketTypeStat struct {
	Type      string `xml:"peer-packet-type"`
	SendCount int    `xml:"peer-packet-send-cnt"`
	RecvCount int    `xml:"peer-packet-rcv-cnt"`
}

type coldSyncMonitor struct {
	Statistics coldSyncStatistics `xml:"cold-synchronization-statistics"`
}

type coldSyncStatistics struct {
	Completed int `xml:"current-cold-synchronization-completed"`
	Failed    int `xml:"current-cold-synchronization-failed"`
}

type spuMonitoring struct {
	Status     spuMonitoringStatus `xml:"spu-monitoring-status"`
	Statistics spuMonitoringStats  `xml:"spu-monitoring-statistics"`
}

type spuMonitoringStatus struct {
	State string `xml:"spu-monitoring-state"`
}

type spuMonitoringStats struct {
	SPUUp          int `xml:"spu-up-count"`
	NPCUp          int `xml:"npc-up-count"`
	SPUDown        int `xml:"spu-down-count"`
	NPCDown        int `xml:"npc-down-count"`
	ChinfoBlobErrs int `xml:"chinfo-blob-error-count"`
}

type loopbackInfo struct {
	Checks []loopbackCheck `xml:"loopback-information-list"`
}

type loopbackCheck struct {
	PFEName        string `xml:"pfe-name"`
	LoopbackStatus string `xml:"loopback-status"`
	NexthopStatus  string `xml:"nexthop-status"`
	MbufStatus     string `xml:"mbuf-status"`
}

type hwMonitoring struct {
	Status hwMonitoringStatus `xml:"high-availability-hardware-monitoring-status-information"`
}

type hwMonitoringStatus struct {
	ActivationStatus string `xml:"high-availability-hardware-monitoring-activation-status"`
	CtrlPlaneErrors  int    `xml:"high-availability-hardware-monitoring-ctrl-plane-errors"`
	DataPlaneErrors  int    `xml:"high-availability-hardware-monitoring-data-plane-errors"`
}

// srgBlock is the state of a single services-redundancy-group, e.g.:
//
//	<chassis-high-availability-default-srg-info>
//	  <current-state>ONLINE</current-state>
//	  <peer-id>2</peer-id>
//	</chassis-high-availability-default-srg-info>
type srgBlock struct {
	CurrentState string `xml:"current-state"`
	PeerID       int    `xml:"peer-id"`
}

// srgResult is the response to "show chassis high-availability services-redundancy-group <id>".
//
// Verified against a real device for group 0 only, where the reply wraps the
// state in <chassis-high-availability-default-srg-info> - the same element
// used for the default group inside the "detail" output. Non-default groups
// (id > 0) are assumed to use the same element name; this has not been
// confirmed against real output and may need a second struct if Junos uses a
// different wrapper (e.g. a non-"default" element) for those.
type srgResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	SRG     srgBlock `xml:"chassis-high-availability-default-srg-info"`
}
