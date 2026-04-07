// SPDX-License-Identifier: MIT

package virtualchassis

import (
	"strconv"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_virtual_chassis_"

var (
    chassisMode           *prometheus.Desc
    memberStatusDesc      *prometheus.Desc
    memberPriorityDesc    *prometheus.Desc
    memberMixedModeDesc   *prometheus.Desc
    vcpStatusDesc         *prometheus.Desc
)

func init() {
    // target label is automatically provided by the exporter framework
    l := []string{"target"}

    chassisMode = prometheus.NewDesc(
        prefix+"mode",
        "Virtual Chassis mode (1 for Enabled)",
        append(l, "virtual_chassis_id", "mode", "preprovisioned"), nil,
    )

    memberStatusDesc = prometheus.NewDesc(
        prefix+"member_status",
        "Status of the member (1 = Present/Prsnt, 0 = Otherwise)",
        append(l, "member_id", "role", "model", "serial_number"), nil,
    )

    memberPriorityDesc = prometheus.NewDesc(
        prefix+"member_priority",
        "Mastership priority of the member (0-255)",
        append(l, "member_id", "serial_number"), nil,
    )

    memberMixedModeDesc = prometheus.NewDesc(
        prefix+"member_mixed_mode",
        "Boolean for mixed mode status (N = Not in mixed mode)",
        append(l, "member_id", "serial_number"), nil,
    )

    vcpStatusDesc = prometheus.NewDesc(
        prefix+"vcp_status",
        "Status of the Virtual Chassis Port (1 = Up, 0 = Down/Disabled)",
        append(l, "port_name", "neighbor_id", "member_node"), nil,
    )
}

type virtualChassisCollector struct{}

func NewCollector() collector.RPCCollector {
    return &virtualChassisCollector{}
}

func (*virtualChassisCollector) Name() string {
    return "Virtual Chassis"
}

func (c *virtualChassisCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- chassisMode
    ch <- memberStatusDesc
    ch <- memberPriorityDesc
    ch <- memberMixedModeDesc
    ch <- vcpStatusDesc
}

func (c *virtualChassisCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
    // Collect VC member information.
    var reply VirtualChassisReply
    err := client.RunCommandAndParse("show virtual-chassis status", &reply)
    if err != nil {
        return err
    }

    // Determine which info block to use
    info := reply.VirtChassInfo.PreProvisVCInfo
    provisionedLabel := "true"
    
    if info.VirtChassID == "" {
        info = reply.VirtChassInfo.NonProvisVCInfo
        provisionedLabel = "false"
    }

    // Export Global Mode
    modeVal := 0.0
    if info.VirtChassMode == "Enabled" {
        modeVal = 1.0
    }

    ch <- prometheus.MustNewConstMetric(
        chassisMode,
        prometheus.GaugeValue,
        modeVal,
        append(labelValues, info.VirtChassID, info.VirtChassMode, provisionedLabel)...,
    )

    // Export Per-Member Metrics
    for _, m := range reply.VirtChassInfo.MemberList {
        mID := strconv.Itoa(m.MemberID)

        // Member Status (1 = Healthy/Present)
        status := 0.0
        if m.MemberStatus == "Prsnt" || m.MemberStatus == "NotPreprevisioned" {
            status = 1.0
        }

        ch <- prometheus.MustNewConstMetric(
            memberStatusDesc,
            prometheus.GaugeValue,
            status,
            append(labelValues, mID, m.MemberRole, m.MemberModel, m.MemberSerial)...,
        )

        // Member Priority
        ch <- prometheus.MustNewConstMetric(
            memberPriorityDesc,
            prometheus.GaugeValue,
            float64(m.MemberPriority),
            append(labelValues, mID, m.MemberSerial)...,
        )

        // Mixed Mode
        mixedmode := 0.0
        if m.MemberMixedMode == "Y" {
            mixedmode = 1.0
        }
        ch <- prometheus.MustNewConstMetric(
            memberMixedModeDesc,
            prometheus.GaugeValue,
            mixedmode,
            append(labelValues, mID, m.MemberSerial)...,
        )
    }

    // Collect VC Port information
    var vcpReply VirtualChassisPortReply
    // "all-members" for every switch in the stack
    port_err := client.RunCommandAndParse("show virtual-chassis vc-port all-members", &vcpReply)
    if port_err != nil {
        return port_err
    }


    // Loop through each member switch's results
    for _, re := range vcpReply.Results.Items {
        // Loop through the ports for that specific member
        for _, p := range re.VCPPortInfo.PortList {
            vcpStatus := 0.0
            if p.PortStatus == "Up" {
                vcpStatus = 1.0
            }

            ch <- prometheus.MustNewConstMetric(
                vcpStatusDesc,
                prometheus.GaugeValue,
                vcpStatus,
		append(labelValues, p.PortName, p.Neighbor, re.Name)...,
            )
        }
    }

    return nil
}
