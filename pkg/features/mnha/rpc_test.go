// SPDX-License-Identifier: MIT

package mnha

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sample captured from a real SRX pair: "show chassis high-availability information detail | display xml"
const detailSample = `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/24.4R2-S2.6/junos">
    <chassis-high-availability-detail-information>
        <chassis-high-availability-node-info>
            <node-status>ONLINE</node-status>
            <grid-id>0</grid-id>
            <local-id>1</local-id>
            <local-ip>192.168.12.250</local-ip>
            <local-fwd-ip>192.168.12.249</local-fwd-ip>
            <chassis-high-availability-peer-info-detail>
                <high-availability-peer-desc></high-availability-peer-desc>
                <high-availability-peer-id>2</high-availability-peer-id>
                <high-availability-peer-ip-address>192.168.12.251</high-availability-peer-ip-address>
                <high-availability-local-interface>lo0.1</high-availability-local-interface>
                <high-availability-peer-rt-instance>mnha_icl</high-availability-peer-rt-instance>
                <high-availability-peer-encryption-status>NO</high-availability-peer-encryption-status>
                <high-availability-peer-bfd-status>UP</high-availability-peer-bfd-status>
                <high-availability-peer-detect-time>3 * 1000ms</high-availability-peer-detect-time>
                <cold-sync-status>COMPLETE</cold-sync-status>
                <high-availability-peer-fwd-ip-address>192.168.12.252</high-availability-peer-fwd-ip-address>
                <high-availability-local-fwd-interface>lo0.1</high-availability-local-fwd-interface>
                <high-availability-peer-fwd-bfd-status>UP</high-availability-peer-fwd-bfd-status>
                <peer-internal-interface>N/A</peer-internal-interface>
                <peer-internal-local-ip>N/A</peer-internal-local-ip>
                <peer-internal-peer-ip>N/A</peer-internal-peer-ip>
                <peer-internal-routing-instance>N/A</peer-internal-routing-instance>
                <local-fwd-ip>192.168.12.249</local-fwd-ip>
                <high-availability-peer-packet-stats>
                    <packet-send-err-cnt>0</packet-send-err-cnt>
                    <packet-recv-err-cnt>0</packet-recv-err-cnt>
                    <peer-packet-send-rcv-stats>
                        <peer-packet-stats-list>
                            <peer-packet-type>SRG Status Msg</peer-packet-type>
                            <peer-packet-send-cnt>0</peer-packet-send-cnt>
                            <peer-packet-rcv-cnt>0</peer-packet-rcv-cnt>
                        </peer-packet-stats-list>
                        <peer-packet-stats-list>
                            <peer-packet-type>SRG Status Ack</peer-packet-type>
                            <peer-packet-send-cnt>0</peer-packet-send-cnt>
                            <peer-packet-rcv-cnt>0</peer-packet-rcv-cnt>
                        </peer-packet-stats-list>
                        <peer-packet-stats-list>
                            <peer-packet-type>Attribute Msg</peer-packet-type>
                            <peer-packet-send-cnt>3</peer-packet-send-cnt>
                            <peer-packet-rcv-cnt>2</peer-packet-rcv-cnt>
                        </peer-packet-stats-list>
                        <peer-packet-stats-list>
                            <peer-packet-type>Attribute Ack</peer-packet-type>
                            <peer-packet-send-cnt>2</peer-packet-send-cnt>
                            <peer-packet-rcv-cnt>2</peer-packet-rcv-cnt>
                        </peer-packet-stats-list>
                    </peer-packet-send-rcv-stats>
                </high-availability-peer-packet-stats>
            </chassis-high-availability-peer-info-detail>
        </chassis-high-availability-node-info>
        <high-availability-peer-status-events>
            <high-availability-peer-status-event>May 20 10:08:24.090 : HA Peer 192.168.12.252 BFD conn came up</high-availability-peer-status-event>
        </high-availability-peer-status-events>
        <high-availability-hw-upgrade-events>
        </high-availability-hw-upgrade-events>
        <cold-synchronization-monitoring-information>
            <cold-synchronization-status-information>
                <cold-synchronization-completed>N/A</cold-synchronization-completed>
                <cold-synchronization-not-completed>N/A</cold-synchronization-not-completed>
                <cold-synchronization-unknown>N/A</cold-synchronization-unknown>
                <cold-synchronization-weight>0</cold-synchronization-weight>
            </cold-synchronization-status-information>
            <cold-synchronization-progress>CS Prereq               1 of 1 SPUs completed</cold-synchronization-progress>
            <cold-synchronization-statistics>
                <current-cold-synchronization-completed>4</current-cold-synchronization-completed>
                <current-cold-synchronization-failed>1</current-cold-synchronization-failed>
            </cold-synchronization-statistics>
            <cold-synchronization-events>
                <cold-synchronization-event>May 20 10:08:25.783 : Cold sync for PFE flowd is RTO sync in process</cold-synchronization-event>
            </cold-synchronization-events>
        </cold-synchronization-monitoring-information>
        <spu-monitoring-information>
            <spu-monitoring-status>
                <spu-monitoring-state>Enabled</spu-monitoring-state>
                <spu-monitoring-weight>0</spu-monitoring-weight>
                <spu-monitoring-spu-state>
                </spu-monitoring-spu-state>
            </spu-monitoring-status>
            <spu-monitoring-statistics>
                <spu-up-count>1</spu-up-count>
                <npc-up-count>0</npc-up-count>
                <spu-down-count>0</spu-down-count>
                <npc-down-count>0</npc-down-count>
                <chinfo-blob-error-count>0</chinfo-blob-error-count>
            </spu-monitoring-statistics>
            <spu-monitoring-events>
            </spu-monitoring-events>
        </spu-monitoring-information>
        <loopback-information>
            <loopback-information-list>
                <pfe-name>flowd</pfe-name>
                <loopback-status>Success</loopback-status>
                <nexthop-status>Success</nexthop-status>
                <mbuf-status>Success</mbuf-status>
            </loopback-information-list>
        </loopback-information>
        <high-availability-hardware-monitoring-information>
            <high-availability-hardware-monitoring-status-information>
                <high-availability-hardware-monitoring-activation-status>Enabled</high-availability-hardware-monitoring-activation-status>
                <high-availability-hardware-monitoring-ctrl-plane-errors>0</high-availability-hardware-monitoring-ctrl-plane-errors>
                <high-availability-hardware-monitoring-data-plane-errors>0</high-availability-hardware-monitoring-data-plane-errors>
            </high-availability-hardware-monitoring-status-information>
            <hardware-monitoring-events>
                <hardware-monitoring-event>May 20 10:04:14.401 : hw-mon errors read      Ctrl Plane  errors:0 Data plane  errors:0</hardware-monitoring-event>
            </hardware-monitoring-events>
        </high-availability-hardware-monitoring-information>
        <chassis-high-availability-default-srg-info>
            <current-state>ONLINE</current-state>
            <peer-id>2</peer-id>
        </chassis-high-availability-default-srg-info>
        <chassis-high-availability-srg-info-detail>
        </chassis-high-availability-srg-info-detail>
    </chassis-high-availability-detail-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

// Sample captured from a real SRX pair: "show chassis high-availability services-redundancy-group 0 | display xml"
const srgSample = `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/24.4R2-S2.6/junos">
    <chassis-high-availability-default-srg-info>
        <current-state>ONLINE</current-state>
        <peer-id>1</peer-id>
    </chassis-high-availability-default-srg-info>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

func TestParseDetailResult(t *testing.T) {
	var res detailResult
	err := xml.Unmarshal([]byte(detailSample), &res)
	if err != nil {
		t.Fatal(err)
	}

	node := res.Info.NodeInfo
	assert.Equal(t, "ONLINE", node.NodeStatus)
	assert.Equal(t, 0, node.GridID)
	assert.Equal(t, 1, node.LocalID)
	assert.Equal(t, "192.168.12.250", node.LocalIP)
	assert.Equal(t, "192.168.12.249", node.LocalFwdIP)

	peer := node.Peer
	assert.Equal(t, 2, peer.PeerID)
	assert.Equal(t, "192.168.12.251", peer.PeerIPAddress)
	assert.Equal(t, "lo0.1", peer.LocalInterface)
	assert.Equal(t, "mnha_icl", peer.RoutingInstance)
	assert.Equal(t, "NO", peer.EncryptionStatus)
	assert.Equal(t, "UP", peer.BFDStatus)
	assert.Equal(t, "3 * 1000ms", peer.DetectTime)
	assert.Equal(t, "COMPLETE", peer.ColdSyncStatus)
	assert.Equal(t, "192.168.12.252", peer.PeerFwdIPAddress)
	assert.Equal(t, "lo0.1", peer.LocalFwdInterface)
	assert.Equal(t, "UP", peer.PeerFwdBFDStatus)

	assert.Equal(t, 0, peer.PacketStats.SendErrCount)
	assert.Equal(t, 0, peer.PacketStats.RecvErrCount)
	assert.Len(t, peer.PacketStats.Stats, 4)
	assert.Equal(t, "SRG Status Msg", peer.PacketStats.Stats[0].Type)
	assert.Equal(t, "Attribute Msg", peer.PacketStats.Stats[2].Type)
	assert.Equal(t, 3, peer.PacketStats.Stats[2].SendCount)
	assert.Equal(t, 2, peer.PacketStats.Stats[2].RecvCount)

	assert.Equal(t, 4, res.Info.ColdSyncMonitoring.Statistics.Completed)
	assert.Equal(t, 1, res.Info.ColdSyncMonitoring.Statistics.Failed)

	assert.Equal(t, "Enabled", res.Info.SPUMonitoring.Status.State)
	assert.Equal(t, 1, res.Info.SPUMonitoring.Statistics.SPUUp)
	assert.Equal(t, 0, res.Info.SPUMonitoring.Statistics.NPCUp)
	assert.Equal(t, 0, res.Info.SPUMonitoring.Statistics.SPUDown)
	assert.Equal(t, 0, res.Info.SPUMonitoring.Statistics.NPCDown)
	assert.Equal(t, 0, res.Info.SPUMonitoring.Statistics.ChinfoBlobErrs)

	assert.Len(t, res.Info.LoopbackInformation.Checks, 1)
	assert.Equal(t, "flowd", res.Info.LoopbackInformation.Checks[0].PFEName)
	assert.Equal(t, "Success", res.Info.LoopbackInformation.Checks[0].LoopbackStatus)
	assert.Equal(t, "Success", res.Info.LoopbackInformation.Checks[0].NexthopStatus)
	assert.Equal(t, "Success", res.Info.LoopbackInformation.Checks[0].MbufStatus)

	assert.Equal(t, "Enabled", res.Info.HardwareMonitoring.Status.ActivationStatus)
	assert.Equal(t, 0, res.Info.HardwareMonitoring.Status.CtrlPlaneErrors)
	assert.Equal(t, 0, res.Info.HardwareMonitoring.Status.DataPlaneErrors)

	assert.Equal(t, "ONLINE", res.Info.DefaultSRG.CurrentState)
	assert.Equal(t, 2, res.Info.DefaultSRG.PeerID)
}

func TestParseSRGResult(t *testing.T) {
	var res srgResult
	err := xml.Unmarshal([]byte(srgSample), &res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "ONLINE", res.SRG.CurrentState)
	assert.Equal(t, 1, res.SRG.PeerID)
}
