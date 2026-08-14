// SPDX-License-Identifier: MIT

package mnha

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_mnha_"

var (
	nodeStatusDesc            *prometheus.Desc
	peerBFDUpDesc             *prometheus.Desc
	peerFwdBFDUpDesc          *prometheus.Desc
	peerEncryptionEnabledDesc *prometheus.Desc
	peerDetectTimeMsDesc      *prometheus.Desc
	coldSyncCompleteDesc      *prometheus.Desc
	peerPacketErrorsTotalDesc *prometheus.Desc
	peerPacketsTotalDesc      *prometheus.Desc
	coldSyncCompletedTotal    *prometheus.Desc
	coldSyncFailedTotal       *prometheus.Desc
	spuMonitoringEnabledDesc  *prometheus.Desc
	spuUpDesc                 *prometheus.Desc
	npcUpDesc                 *prometheus.Desc
	spuDownDesc               *prometheus.Desc
	npcDownDesc               *prometheus.Desc
	chinfoBlobErrorsTotal     *prometheus.Desc
	loopbackCheckSuccessDesc  *prometheus.Desc
	hwMonitoringEnabledDesc   *prometheus.Desc
	hwCtrlPlaneErrorsTotal    *prometheus.Desc
	hwDataPlaneErrorsTotal    *prometheus.Desc
	srgStateDesc              *prometheus.Desc
)

var detectTimeRe = regexp.MustCompile(`^\s*(\d+)\s*\*\s*(\d+)\s*ms\s*$`)

func init() {
	l := []string{"target", "grid_id", "local_id", "local_ip", "local_fwd_ip", "status"}
	nodeStatusDesc = prometheus.NewDesc(prefix+"node_status", "MNHA local node status (1 online, 0 offline, -1 unknown)", l, nil)

	l = []string{"target", "peer_id", "peer_ip", "local_interface"}
	peerBFDUpDesc = prometheus.NewDesc(prefix+"peer_bfd_up", "MNHA peer BFD session state over the ICL (1 up, 0 down)", l, nil)

	l = []string{"target", "peer_id", "peer_fwd_ip", "local_fwd_interface"}
	peerFwdBFDUpDesc = prometheus.NewDesc(prefix+"peer_fwd_bfd_up", "MNHA peer forwarding-path BFD session state (1 up, 0 down)", l, nil)

	l = []string{"target", "peer_id"}
	peerEncryptionEnabledDesc = prometheus.NewDesc(prefix+"peer_encryption_enabled", "1 if the MNHA ICL is encrypted, 0 otherwise", l, nil)
	peerDetectTimeMsDesc = prometheus.NewDesc(prefix+"peer_detect_time_milliseconds", "MNHA peer BFD failure detection time in milliseconds", l, nil)

	l = []string{"target", "peer_id", "status"}
	coldSyncCompleteDesc = prometheus.NewDesc(prefix+"cold_sync_complete", "1 if MNHA cold synchronization with the peer is complete, 0 otherwise", l, nil)

	l = []string{"target", "peer_id", "direction"}
	peerPacketErrorsTotalDesc = prometheus.NewDesc(prefix+"peer_packet_errors_total", "Total MNHA ICL packet errors by direction (send, receive)", l, nil)

	l = []string{"target", "peer_id", "packet_type", "direction"}
	peerPacketsTotalDesc = prometheus.NewDesc(prefix+"peer_packets_total", "Total MNHA ICL packets by type and direction (send, receive)", l, nil)

	l = []string{"target"}
	coldSyncCompletedTotal = prometheus.NewDesc(prefix+"cold_sync_completed_total", "Total number of completed MNHA cold synchronizations", l, nil)
	coldSyncFailedTotal = prometheus.NewDesc(prefix+"cold_sync_failed_total", "Total number of failed MNHA cold synchronizations", l, nil)
	spuMonitoringEnabledDesc = prometheus.NewDesc(prefix+"spu_monitoring_enabled", "1 if MNHA SPU monitoring is enabled, 0 otherwise", l, nil)
	spuUpDesc = prometheus.NewDesc(prefix+"spu_up", "Number of SPUs reported up by MNHA monitoring", l, nil)
	npcUpDesc = prometheus.NewDesc(prefix+"npc_up", "Number of NPCs reported up by MNHA monitoring", l, nil)
	spuDownDesc = prometheus.NewDesc(prefix+"spu_down", "Number of SPUs reported down by MNHA monitoring", l, nil)
	npcDownDesc = prometheus.NewDesc(prefix+"npc_down", "Number of NPCs reported down by MNHA monitoring", l, nil)
	chinfoBlobErrorsTotal = prometheus.NewDesc(prefix+"chinfo_blob_errors_total", "Total MNHA chassis-info blob errors", l, nil)
	hwMonitoringEnabledDesc = prometheus.NewDesc(prefix+"hardware_monitoring_enabled", "1 if MNHA hardware monitoring is enabled, 0 otherwise", l, nil)
	hwCtrlPlaneErrorsTotal = prometheus.NewDesc(prefix+"hardware_control_plane_errors_total", "Total MNHA hardware monitoring control-plane errors", l, nil)
	hwDataPlaneErrorsTotal = prometheus.NewDesc(prefix+"hardware_data_plane_errors_total", "Total MNHA hardware monitoring data-plane errors", l, nil)

	l = []string{"target", "pfe_name", "check"}
	loopbackCheckSuccessDesc = prometheus.NewDesc(prefix+"loopback_check_success", "1 if the MNHA PFE loopback self-test succeeded (loopback, nexthop, mbuf), 0 otherwise", l, nil)

	l = []string{"target", "srg_id", "peer_id", "state"}
	srgStateDesc = prometheus.NewDesc(prefix+"srg_state", "MNHA services-redundancy-group state (1 online, 0 offline, -1 unknown)", l, nil)
}

type mnhaCollector struct {
	srgIDs []int
}

// NewCollector creates a new collector for MNHA (Mixed/Multi-Node High Availability).
// srgIDs configures which services-redundancy-groups are polled individually
// via "show chassis high-availability services-redundancy-group <id>".
// If empty, group 0 (the default group) is used.
func NewCollector(srgIDs []int) collector.RPCCollector {
	if len(srgIDs) == 0 {
		srgIDs = []int{0}
	}

	return &mnhaCollector{srgIDs: srgIDs}
}

// Name returns the name of the collector
func (*mnhaCollector) Name() string {
	return "MNHA"
}

// Describe describes the metrics
func (*mnhaCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- nodeStatusDesc
	ch <- peerBFDUpDesc
	ch <- peerFwdBFDUpDesc
	ch <- peerEncryptionEnabledDesc
	ch <- peerDetectTimeMsDesc
	ch <- coldSyncCompleteDesc
	ch <- peerPacketErrorsTotalDesc
	ch <- peerPacketsTotalDesc
	ch <- coldSyncCompletedTotal
	ch <- coldSyncFailedTotal
	ch <- spuMonitoringEnabledDesc
	ch <- spuUpDesc
	ch <- npcUpDesc
	ch <- spuDownDesc
	ch <- npcDownDesc
	ch <- chinfoBlobErrorsTotal
	ch <- loopbackCheckSuccessDesc
	ch <- hwMonitoringEnabledDesc
	ch <- hwCtrlPlaneErrorsTotal
	ch <- hwDataPlaneErrorsTotal
	ch <- srgStateDesc
}

// Collect collects metrics from JunOS
func (c *mnhaCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var errs []error

	var res detailResult
	detailErr := client.RunCommandAndParse("show chassis high-availability information detail", &res)
	if detailErr != nil {
		errs = append(errs, detailErr)
	} else {
		collectDetail(&res, ch, labelValues)
	}

	for _, id := range c.srgIDs {
		var srg srgBlock

		// SRG 0 (the default group) is already reported by the "detail" call
		// above (chassis-high-availability-default-srg-info), so avoid a
		// second SSH round-trip for the common case of a single default group.
		if id == 0 && detailErr == nil {
			srg = res.Info.DefaultSRG
		} else {
			var result srgResult
			cmd := fmt.Sprintf("show chassis high-availability services-redundancy-group %d", id)
			if err := client.RunCommandAndParse(cmd, &result); err != nil {
				errs = append(errs, err)
				continue
			}
			srg = result.SRG
		}

		state := strings.ToUpper(strings.TrimSpace(srg.CurrentState))
		l := append(labelValues, strconv.Itoa(id), strconv.Itoa(srg.PeerID), state)
		ch <- prometheus.MustNewConstMetric(srgStateDesc, prometheus.GaugeValue, onlineOfflineValue(state), l...)
	}

	return errors.Join(errs...)
}

func collectDetail(res *detailResult, ch chan<- prometheus.Metric, labelValues []string) {
	info := res.Info
	node := info.NodeInfo
	peer := node.Peer

	nodeStatus := strings.ToUpper(strings.TrimSpace(node.NodeStatus))
	l := append(labelValues, strconv.Itoa(node.GridID), strconv.Itoa(node.LocalID), node.LocalIP, node.LocalFwdIP, nodeStatus)
	ch <- prometheus.MustNewConstMetric(nodeStatusDesc, prometheus.GaugeValue, onlineOfflineValue(nodeStatus), l...)

	peerID := strconv.Itoa(peer.PeerID)

	l = append(labelValues, peerID, peer.PeerIPAddress, peer.LocalInterface)
	ch <- prometheus.MustNewConstMetric(peerBFDUpDesc, prometheus.GaugeValue, upDownValue(peer.BFDStatus), l...)

	l = append(labelValues, peerID, peer.PeerFwdIPAddress, peer.LocalFwdInterface)
	ch <- prometheus.MustNewConstMetric(peerFwdBFDUpDesc, prometheus.GaugeValue, upDownValue(peer.PeerFwdBFDStatus), l...)

	l = append(labelValues, peerID)
	ch <- prometheus.MustNewConstMetric(peerEncryptionEnabledDesc, prometheus.GaugeValue, yesNoValue(peer.EncryptionStatus), l...)

	if ms, ok := parseDetectTimeMs(peer.DetectTime); ok {
		ch <- prometheus.MustNewConstMetric(peerDetectTimeMsDesc, prometheus.GaugeValue, ms, l...)
	}

	coldSyncStatus := strings.ToUpper(strings.TrimSpace(peer.ColdSyncStatus))
	l = append(labelValues, peerID, coldSyncStatus)
	ch <- prometheus.MustNewConstMetric(coldSyncCompleteDesc, prometheus.GaugeValue, boolValue(coldSyncStatus == "COMPLETE"), l...)

	l = append(labelValues, peerID, "send")
	ch <- prometheus.MustNewConstMetric(peerPacketErrorsTotalDesc, prometheus.CounterValue, float64(peer.PacketStats.SendErrCount), l...)
	l = append(labelValues, peerID, "receive")
	ch <- prometheus.MustNewConstMetric(peerPacketErrorsTotalDesc, prometheus.CounterValue, float64(peer.PacketStats.RecvErrCount), l...)

	for _, s := range peer.PacketStats.Stats {
		l = append(labelValues, peerID, s.Type, "send")
		ch <- prometheus.MustNewConstMetric(peerPacketsTotalDesc, prometheus.CounterValue, float64(s.SendCount), l...)
		l = append(labelValues, peerID, s.Type, "receive")
		ch <- prometheus.MustNewConstMetric(peerPacketsTotalDesc, prometheus.CounterValue, float64(s.RecvCount), l...)
	}

	ch <- prometheus.MustNewConstMetric(coldSyncCompletedTotal, prometheus.CounterValue, float64(info.ColdSyncMonitoring.Statistics.Completed), labelValues...)
	ch <- prometheus.MustNewConstMetric(coldSyncFailedTotal, prometheus.CounterValue, float64(info.ColdSyncMonitoring.Statistics.Failed), labelValues...)

	ch <- prometheus.MustNewConstMetric(spuMonitoringEnabledDesc, prometheus.GaugeValue, enabledDisabledValue(info.SPUMonitoring.Status.State), labelValues...)
	ch <- prometheus.MustNewConstMetric(spuUpDesc, prometheus.GaugeValue, float64(info.SPUMonitoring.Statistics.SPUUp), labelValues...)
	ch <- prometheus.MustNewConstMetric(npcUpDesc, prometheus.GaugeValue, float64(info.SPUMonitoring.Statistics.NPCUp), labelValues...)
	ch <- prometheus.MustNewConstMetric(spuDownDesc, prometheus.GaugeValue, float64(info.SPUMonitoring.Statistics.SPUDown), labelValues...)
	ch <- prometheus.MustNewConstMetric(npcDownDesc, prometheus.GaugeValue, float64(info.SPUMonitoring.Statistics.NPCDown), labelValues...)
	ch <- prometheus.MustNewConstMetric(chinfoBlobErrorsTotal, prometheus.CounterValue, float64(info.SPUMonitoring.Statistics.ChinfoBlobErrs), labelValues...)

	ch <- prometheus.MustNewConstMetric(hwMonitoringEnabledDesc, prometheus.GaugeValue, enabledDisabledValue(info.HardwareMonitoring.Status.ActivationStatus), labelValues...)
	ch <- prometheus.MustNewConstMetric(hwCtrlPlaneErrorsTotal, prometheus.CounterValue, float64(info.HardwareMonitoring.Status.CtrlPlaneErrors), labelValues...)
	ch <- prometheus.MustNewConstMetric(hwDataPlaneErrorsTotal, prometheus.CounterValue, float64(info.HardwareMonitoring.Status.DataPlaneErrors), labelValues...)

	for _, check := range info.LoopbackInformation.Checks {
		cl := append(labelValues, check.PFEName, "loopback")
		ch <- prometheus.MustNewConstMetric(loopbackCheckSuccessDesc, prometheus.GaugeValue, boolValue(strings.EqualFold(check.LoopbackStatus, "Success")), cl...)

		cl = append(labelValues, check.PFEName, "nexthop")
		ch <- prometheus.MustNewConstMetric(loopbackCheckSuccessDesc, prometheus.GaugeValue, boolValue(strings.EqualFold(check.NexthopStatus, "Success")), cl...)

		cl = append(labelValues, check.PFEName, "mbuf")
		ch <- prometheus.MustNewConstMetric(loopbackCheckSuccessDesc, prometheus.GaugeValue, boolValue(strings.EqualFold(check.MbufStatus, "Success")), cl...)
	}
}

func onlineOfflineValue(status string) float64 {
	switch status {
	case "ONLINE":
		return 1
	case "OFFLINE":
		return 0
	default:
		return -1
	}
}

func upDownValue(status string) float64 {
	return boolValue(strings.EqualFold(strings.TrimSpace(status), "UP"))
}

func yesNoValue(status string) float64 {
	return boolValue(strings.EqualFold(strings.TrimSpace(status), "YES"))
}

func enabledDisabledValue(status string) float64 {
	return boolValue(strings.EqualFold(strings.TrimSpace(status), "Enabled"))
}

func boolValue(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

// parseDetectTimeMs parses Junos' "<count> * <interval>ms" detect-time format,
// e.g. "3 * 1000ms" (3 missed intervals of 1000ms) into a total in milliseconds.
func parseDetectTimeMs(s string) (float64, bool) {
	m := detectTimeRe.FindStringSubmatch(s)
	if m == nil {
		return 0, false
	}

	count, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return 0, false
	}

	interval, err := strconv.ParseFloat(m[2], 64)
	if err != nil {
		return 0, false
	}

	return count * interval, true
}
