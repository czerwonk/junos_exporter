// SPDX-License-Identifier: MIT

package interfaces

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/dynamiclabels"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_"

type description struct {
	receiveBytesDesc            *prometheus.Desc
	receivePacketsDesc          *prometheus.Desc
	receiveErrorsDesc           *prometheus.Desc
	receiveDropsDesc            *prometheus.Desc
	interfaceSpeedDesc          *prometheus.Desc
	interfaceBPDUErrorDesc      *prometheus.Desc
	transmitBytesDesc           *prometheus.Desc
	transmitPacketsDesc         *prometheus.Desc
	transmitErrorsDesc          *prometheus.Desc
	transmitDropsDesc           *prometheus.Desc
	ipv6receiveBytesDesc        *prometheus.Desc
	ipv6receivePacketsDesc      *prometheus.Desc
	ipv6transmitBytesDesc       *prometheus.Desc
	ipv6transmitPacketsDesc     *prometheus.Desc
	adminStatusDesc             *prometheus.Desc
	operStatusDesc              *prometheus.Desc
	errorStatusDesc             *prometheus.Desc
	lastFlappedDesc             *prometheus.Desc
	receiveUnicastsDesc         *prometheus.Desc
	receiveBroadcastsDesc       *prometheus.Desc
	receiveMulticastsDesc       *prometheus.Desc
	receiveCRCErrorsDesc        *prometheus.Desc
	transmitUnicastsDesc        *prometheus.Desc
	transmitBroadcastsDesc      *prometheus.Desc
	transmitMulticastsDesc      *prometheus.Desc
	transmitCRCErrorsDesc       *prometheus.Desc
	fecCcwCountDesc             *prometheus.Desc
	fecNccwCountDesc            *prometheus.Desc
	fecCcwErrorRateDesc         *prometheus.Desc
	fecNccwErrorRateDesc        *prometheus.Desc
	receiveOversizedFramesDesc  *prometheus.Desc
	receiveJabberFramesDesc     *prometheus.Desc
	receiveFragmentFramesDesc   *prometheus.Desc
	receiveVlanTaggedFramesDesc *prometheus.Desc
	receiveCodeViolationsDesc   *prometheus.Desc
	receiveTotalErrorsDesc      *prometheus.Desc
	transmitTotalErrorsDesc     *prometheus.Desc
	mtuDesc                     *prometheus.Desc
	fecModeDesc                 *prometheus.Desc
}

func newDescriptions(dynLabels dynamiclabels.Labels) *description {
	d := &description{}
	l := []string{"target", "name", "description", "mac"}
	l = append(l, dynLabels.Keys()...)

	d.receiveBytesDesc = prometheus.NewDesc(prefix+"receive_bytes", "Received data in bytes", l, nil)
	d.receivePacketsDesc = prometheus.NewDesc(prefix+"receive_packets_total", "Received packets", l, nil)
	d.receiveErrorsDesc = prometheus.NewDesc(prefix+"receive_errors", "Number of errors caused by incoming packets", l, nil)
	d.receiveDropsDesc = prometheus.NewDesc(prefix+"receive_drops", "Number of dropped incoming packets", l, nil)
	d.interfaceSpeedDesc = prometheus.NewDesc(prefix+"speed", "speed in in bps", l, nil)
	d.interfaceBPDUErrorDesc = prometheus.NewDesc(prefix+"error_bpdublock", "Flag which tells that there's a BPDU_Block on the interface (bool)", l, nil)
	d.transmitBytesDesc = prometheus.NewDesc(prefix+"transmit_bytes", "Transmitted data in bytes", l, nil)
	d.transmitPacketsDesc = prometheus.NewDesc(prefix+"transmit_packets_total", "Transmitted packets", l, nil)
	d.transmitErrorsDesc = prometheus.NewDesc(prefix+"transmit_errors", "Number of errors caused by outgoing packets", l, nil)
	d.transmitDropsDesc = prometheus.NewDesc(prefix+"transmit_drops", "Number of dropped outgoing packets", l, nil)
	d.ipv6receiveBytesDesc = prometheus.NewDesc(prefix+"IPv6_receive_bytes_total", "Received IPv6 data in bytes", l, nil)
	d.ipv6receivePacketsDesc = prometheus.NewDesc(prefix+"IPv6_receive_packets_total", "Received IPv6 packets", l, nil)
	d.ipv6transmitBytesDesc = prometheus.NewDesc(prefix+"IPv6_transmit_bytes_total", "Transmitted IPv6 data in bytes", l, nil)
	d.ipv6transmitPacketsDesc = prometheus.NewDesc(prefix+"IPv6_transmit_packets_total", "Transmitted IPv6 packets", l, nil)
	d.adminStatusDesc = prometheus.NewDesc(prefix+"admin_up", "Admin operational status", l, nil)
	d.operStatusDesc = prometheus.NewDesc(prefix+"up", "Interface operational status", l, nil)
	d.errorStatusDesc = prometheus.NewDesc(prefix+"error_status", "Admin and operational status differ", l, nil)
	d.lastFlappedDesc = prometheus.NewDesc(prefix+"last_flapped_seconds", "Seconds since last flapped (-1 if never)", l, nil)
	d.receiveUnicastsDesc = prometheus.NewDesc(prefix+"receive_unicasts_packets", "Received unicast packets", l, nil)
	d.receiveBroadcastsDesc = prometheus.NewDesc(prefix+"receive_broadcasts_packets", "Received broadcast packets", l, nil)
	d.receiveMulticastsDesc = prometheus.NewDesc(prefix+"receive_multicasts_packets", "Received multicast packets", l, nil)
	d.receiveCRCErrorsDesc = prometheus.NewDesc(prefix+"receive_errors_crc_packets", "Number of CRC error incoming packets", l, nil)
	d.transmitUnicastsDesc = prometheus.NewDesc(prefix+"transmit_unicasts_packets", "Transmitted unicast packets", l, nil)
	d.transmitBroadcastsDesc = prometheus.NewDesc(prefix+"transmit_broadcasts_packets", "Transmitted broadcast packets", l, nil)
	d.transmitMulticastsDesc = prometheus.NewDesc(prefix+"transmit_multicasts_packets", "Transmitted multicast packets", l, nil)
	d.transmitCRCErrorsDesc = prometheus.NewDesc(prefix+"transmit_errors_crc_packets", "Number of CRC error outgoing packets", l, nil)
	d.fecCcwCountDesc = prometheus.NewDesc(prefix+"fec_ccw_count", "Number FEC Corrected Errors", l, nil)
	d.fecNccwCountDesc = prometheus.NewDesc(prefix+"fec_nccw_count", "Number FEC Uncorrected Errors", l, nil)
	d.fecCcwErrorRateDesc = prometheus.NewDesc(prefix+"fec_ccw_error_rate", "Number FEC Corrected Errors Rate", l, nil)
	d.fecNccwErrorRateDesc = prometheus.NewDesc(prefix+"fec_nccw_error_rate", "Number FEC Uncorrected Errors Rate", l, nil)
	d.receiveOversizedFramesDesc = prometheus.NewDesc(prefix+"receive_oversized_frames", "Number of received Oversize Frames", l, nil)
	d.receiveJabberFramesDesc = prometheus.NewDesc(prefix+"receive_jabber_frames", "Number of received Jabber Frames", l, nil)
	d.receiveFragmentFramesDesc = prometheus.NewDesc(prefix+"receive_fragment_frames", "Number of received Fragment Frames", l, nil)
	d.receiveVlanTaggedFramesDesc = prometheus.NewDesc(prefix+"receive_vlan_tagged_frames", "Number of received Vlan Tagged Frames", l, nil)
	d.receiveCodeViolationsDesc = prometheus.NewDesc(prefix+"receive_code_violations", "Number of received Code Violations", l, nil)
	d.receiveTotalErrorsDesc = prometheus.NewDesc(prefix+"receive_total_errors", "Number of received Total Errors", l, nil)
	d.transmitTotalErrorsDesc = prometheus.NewDesc(prefix+"transmit_total_errors", "Number of transmitted Total Errors", l, nil)
	d.mtuDesc = prometheus.NewDesc(prefix+"mtu", "configured MTU", l, nil)
	d.fecModeDesc = prometheus.NewDesc(prefix+"fec_mode", "Mode of FEC. 0 for none, 1 for default, 2 for fec74, 3 for fec91, 4 for fec108", l, nil)
	return d
}

// Collector collects interface metrics
type interfaceCollector struct {
	descriptionRe *regexp.Regexp
}

// NewCollector creates a new collector
func NewCollector(descRe *regexp.Regexp) collector.RPCCollector {
	c := &interfaceCollector{
		descriptionRe: descRe,
	}

	return c
}

// Name returns the name of the collector
func (*interfaceCollector) Name() string {
	return "Interfaces"
}

// Describe describes the metrics
func (*interfaceCollector) Describe(ch chan<- *prometheus.Desc) {
	d := newDescriptions(nil)
	ch <- d.receiveBytesDesc
	ch <- d.receivePacketsDesc
	ch <- d.receiveErrorsDesc
	ch <- d.receiveDropsDesc
	ch <- d.interfaceSpeedDesc
	ch <- d.interfaceBPDUErrorDesc
	ch <- d.transmitBytesDesc
	ch <- d.transmitPacketsDesc
	ch <- d.transmitDropsDesc
	ch <- d.transmitErrorsDesc
	ch <- d.ipv6receiveBytesDesc
	ch <- d.ipv6receivePacketsDesc
	ch <- d.ipv6transmitBytesDesc
	ch <- d.ipv6transmitPacketsDesc
	ch <- d.adminStatusDesc
	ch <- d.operStatusDesc
	ch <- d.errorStatusDesc
	ch <- d.lastFlappedDesc
	ch <- d.receiveUnicastsDesc
	ch <- d.receiveBroadcastsDesc
	ch <- d.receiveMulticastsDesc
	ch <- d.receiveCRCErrorsDesc
	ch <- d.transmitUnicastsDesc
	ch <- d.transmitBroadcastsDesc
	ch <- d.transmitMulticastsDesc
	ch <- d.transmitCRCErrorsDesc
	ch <- d.fecCcwCountDesc
	ch <- d.fecNccwCountDesc
	ch <- d.fecCcwErrorRateDesc
	ch <- d.fecNccwErrorRateDesc
	ch <- d.receiveOversizedFramesDesc
	ch <- d.receiveJabberFramesDesc
	ch <- d.receiveFragmentFramesDesc
	ch <- d.receiveVlanTaggedFramesDesc
	ch <- d.receiveCodeViolationsDesc
	ch <- d.receiveTotalErrorsDesc
	ch <- d.transmitTotalErrorsDesc
	ch <- d.mtuDesc
	ch <- d.fecModeDesc
}

// Collect collects metrics from JunOS
func (c *interfaceCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	stats, err := c.interfaceStats(client)
	if err != nil {
		return err
	}

	for _, s := range stats {
		c.collectForInterface(s, ch, labelValues)
	}

	return nil
}

func (c *interfaceCollector) interfaceStats(client collector.Client) ([]*interfaceStats, error) {
	var x = result{}
	err := client.RunCommandAndParse("show interfaces extensive", &x)
	if err != nil {
		return nil, err
	}

	stats := make([]*interfaceStats, 0)
	for _, phy := range x.Information.Interfaces {
		s := &interfaceStats{
			IsPhysical:              true,
			Name:                    phy.Name,
			AdminStatus:             phy.AdminStatus == "up",
			OperStatus:              phy.OperStatus == "up",
			ErrorStatus:             !(phy.AdminStatus == phy.OperStatus),
			Description:             phy.Description,
			Mac:                     phy.MacAddress,
			ReceiveDrops:            float64(phy.InputErrors.Drops),
			ReceiveErrors:           float64(phy.InputErrors.Errors),
			ReceiveBytes:            float64(phy.Stats.InputBytes),
			ReceivePackets:          float64(phy.Stats.InputPackets),
			Speed:                   phy.Speed,
			BPDUError:               phy.BPDUError == "detected",
			TransmitDrops:           float64(phy.OutputErrors.Drops),
			TransmitErrors:          float64(phy.OutputErrors.Errors),
			TransmitBytes:           float64(phy.Stats.OutputBytes),
			TransmitPackets:         float64(phy.Stats.OutputPackets),
			IPv6ReceiveBytes:        float64(phy.Stats.IPv6Traffic.InputBytes),
			IPv6ReceivePackets:      float64(phy.Stats.IPv6Traffic.InputPackets),
			IPv6TransmitBytes:       float64(phy.Stats.IPv6Traffic.OutputBytes),
			IPv6TransmitPackets:     float64(phy.Stats.IPv6Traffic.OutputPackets),
			LastFlapped:             -1,
			ReceiveUnicasts:         float64(phy.MACStatistics.InputUnicasts),
			ReceiveBroadcasts:       float64(phy.MACStatistics.InputBroadcasts),
			ReceiveMulticasts:       float64(phy.MACStatistics.InputMulticasts),
			ReceiveCRCErrors:        float64(phy.MACStatistics.InputCRCErrors),
			TransmitUnicasts:        float64(phy.MACStatistics.OutputUnicasts),
			TransmitBroadcasts:      float64(phy.MACStatistics.OutputBroadcasts),
			TransmitMulticasts:      float64(phy.MACStatistics.OutputMulticasts),
			TransmitCRCErrors:       float64(phy.MACStatistics.OutputCRCErrors),
			FecCcwCount:             float64(phy.FECStatistics.NumberfecCcwCount),
			FecNccwCount:            float64(phy.FECStatistics.NumberfecNccwCount),
			FecCcwErrorRate:         float64(phy.FECStatistics.NumberfecCcwErrorRate),
			FecNccwErrorRate:        float64(phy.FECStatistics.NumberfecNccwErrorRate),
			ReceiveOversizedFrames:  float64(phy.MACStatistics.InputOversizedFrames),
			ReceiveJabberFrames:     float64(phy.MACStatistics.InputJabberFrames),
			ReceiveFragmentFrames:   float64(phy.MACStatistics.InputFragmentFrames),
			ReceiveVlanTaggedFrames: float64(phy.MACStatistics.InputVlanTaggedFrames),
			ReceiveCodeViolations:   float64(phy.MACStatistics.InputCodeViolations),
			ReceiveTotalErrors:      float64(phy.MACStatistics.InputTotalErrors),
			TransmitTotalErrors:     float64(phy.MACStatistics.OutputTotalErrors),
			MTU:                     phy.MTU,
			FECMode:                 convertFECModeToFloat64(strings.ToLower(strings.TrimRight(phy.EthernetFecMode.EnabledFecMode, "\n"))),
		}

		if phy.InterfaceFlapped.Value != "Never" {
			s.LastFlapped = float64(phy.InterfaceFlapped.Seconds)
		}

		stats = append(stats, s)

		for _, log := range phy.LogicalInterfaces {
			var s trafficStat
			if (log.Stats != trafficStat{}) {
				s = log.Stats
			} else {
				s = log.LagStats.Stats
			}
			sl := &interfaceStats{
				IsPhysical:          false,
				Name:                log.Name,
				Description:         log.Description,
				Mac:                 phy.MacAddress,
				ReceiveBytes:        float64(s.InputBytes),
				ReceivePackets:      float64(s.InputPackets),
				TransmitBytes:       float64(s.OutputBytes),
				TransmitPackets:     float64(s.OutputPackets),
				IPv6ReceiveBytes:    float64(s.IPv6Traffic.InputBytes),
				IPv6ReceivePackets:  float64(s.IPv6Traffic.InputPackets),
				IPv6TransmitBytes:   float64(s.IPv6Traffic.OutputBytes),
				IPv6TransmitPackets: float64(s.IPv6Traffic.OutputPackets),
			}

			stats = append(stats, sl)
		}
	}

	return stats, nil
}

func (c *interfaceCollector) collectForInterface(s *interfaceStats, ch chan<- prometheus.Metric, labelValues []string) {
	lv := append(labelValues, []string{s.Name, s.Description, s.Mac}...)
	dynLabels := dynamiclabels.ParseDescription(s.Description, c.descriptionRe)
	lv = append(lv, dynLabels.Values()...)
	d := newDescriptions(dynLabels)

	ch <- prometheus.MustNewConstMetric(d.receiveBytesDesc, prometheus.CounterValue, s.ReceiveBytes, lv...)
	ch <- prometheus.MustNewConstMetric(d.receivePacketsDesc, prometheus.CounterValue, s.ReceivePackets, lv...)
	ch <- prometheus.MustNewConstMetric(d.transmitBytesDesc, prometheus.CounterValue, s.TransmitBytes, lv...)
	ch <- prometheus.MustNewConstMetric(d.transmitPacketsDesc, prometheus.CounterValue, s.TransmitPackets, lv...)
	ch <- prometheus.MustNewConstMetric(d.ipv6receiveBytesDesc, prometheus.CounterValue, s.IPv6ReceiveBytes, lv...)
	ch <- prometheus.MustNewConstMetric(d.ipv6receivePacketsDesc, prometheus.CounterValue, s.IPv6ReceivePackets, lv...)
	ch <- prometheus.MustNewConstMetric(d.ipv6transmitBytesDesc, prometheus.CounterValue, s.IPv6TransmitBytes, lv...)
	ch <- prometheus.MustNewConstMetric(d.ipv6transmitPacketsDesc, prometheus.CounterValue, s.IPv6TransmitPackets, lv...)

	if s.IsPhysical {
		adminUp := 0
		if s.AdminStatus {
			adminUp = 1
		}
		operUp := 0
		if s.OperStatus {
			operUp = 1
		}
		err := 0
		if s.ErrorStatus {
			err = 1
		}

		speed := "0"
		if strings.Contains(strings.ToLower(s.Speed), "mbps") {
			speed = strings.Replace(strings.ToLower(s.Speed), "mbps", "000000", 1)
		}
		if strings.Contains(s.Speed, "Gbps") {
			speed = strings.Replace(s.Speed, "Gbps", "000000000", 1)
		}
		if strings.Contains(s.Speed, "Auto") || strings.Contains(s.Speed, "Unspecified") {
			//some cards have just 'Auto' as speed, let's check if it's Gigabit
			if strings.Contains(s.Name, "ge-") {
				speed = "1000000000"
			} else if strings.Contains(s.Name, "xe-") {
				speed = "10000000000"
			} else {
				speed = strings.Replace(s.Speed, "Auto", "0", 1)
				speed = strings.Replace(speed, "Unspecified", "0", 1)
			}
		}
		if strings.Contains(s.Speed, "Unlimited") {
			speed = strings.Replace(s.Speed, "Unlimited", "0", 1)
		}

		// Trimming all white spaces in the entire string
		speed = strings.ReplaceAll(speed, " ", "")

		sp64, _ := strconv.ParseFloat(speed, 64)

		if s.BPDUError {
			ch <- prometheus.MustNewConstMetric(d.interfaceBPDUErrorDesc, prometheus.GaugeValue, float64(1), lv...)
		}

		mtu := s.MTU
		if strings.Contains(s.MTU, "Unlimited") {
			mtu = "65535"
		}
		mtu64, _ := strconv.ParseFloat(mtu, 64)
		ch <- prometheus.MustNewConstMetric(d.adminStatusDesc, prometheus.GaugeValue, float64(adminUp), lv...)
		ch <- prometheus.MustNewConstMetric(d.operStatusDesc, prometheus.GaugeValue, float64(operUp), lv...)
		ch <- prometheus.MustNewConstMetric(d.errorStatusDesc, prometheus.GaugeValue, float64(err), lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitErrorsDesc, prometheus.CounterValue, s.TransmitErrors, lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitDropsDesc, prometheus.CounterValue, s.TransmitDrops, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveErrorsDesc, prometheus.CounterValue, s.ReceiveErrors, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveDropsDesc, prometheus.CounterValue, s.ReceiveDrops, lv...)
		ch <- prometheus.MustNewConstMetric(d.interfaceSpeedDesc, prometheus.GaugeValue, float64(sp64), lv...)
		ch <- prometheus.MustNewConstMetric(d.mtuDesc, prometheus.GaugeValue, float64(mtu64), lv...)

		if s.LastFlapped != 0 {
			ch <- prometheus.MustNewConstMetric(d.lastFlappedDesc, prometheus.GaugeValue, s.LastFlapped, lv...)
		}

		ch <- prometheus.MustNewConstMetric(d.receiveUnicastsDesc, prometheus.CounterValue, s.ReceiveUnicasts, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveBroadcastsDesc, prometheus.CounterValue, s.ReceiveBroadcasts, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveMulticastsDesc, prometheus.CounterValue, s.ReceiveMulticasts, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveCRCErrorsDesc, prometheus.CounterValue, s.ReceiveCRCErrors, lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitUnicastsDesc, prometheus.CounterValue, s.TransmitUnicasts, lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitBroadcastsDesc, prometheus.CounterValue, s.TransmitBroadcasts, lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitMulticastsDesc, prometheus.CounterValue, s.TransmitMulticasts, lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitCRCErrorsDesc, prometheus.CounterValue, s.TransmitCRCErrors, lv...)
		ch <- prometheus.MustNewConstMetric(d.fecCcwCountDesc, prometheus.CounterValue, s.FecCcwCount, lv...)
		ch <- prometheus.MustNewConstMetric(d.fecNccwCountDesc, prometheus.CounterValue, s.FecNccwCount, lv...)
		ch <- prometheus.MustNewConstMetric(d.fecCcwErrorRateDesc, prometheus.CounterValue, s.FecCcwErrorRate, lv...)
		ch <- prometheus.MustNewConstMetric(d.fecNccwErrorRateDesc, prometheus.CounterValue, s.FecNccwErrorRate, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveOversizedFramesDesc, prometheus.CounterValue, s.ReceiveOversizedFrames, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveJabberFramesDesc, prometheus.CounterValue, s.ReceiveJabberFrames, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveFragmentFramesDesc, prometheus.CounterValue, s.ReceiveFragmentFrames, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveVlanTaggedFramesDesc, prometheus.CounterValue, s.ReceiveVlanTaggedFrames, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveCodeViolationsDesc, prometheus.CounterValue, s.ReceiveCodeViolations, lv...)
		ch <- prometheus.MustNewConstMetric(d.receiveTotalErrorsDesc, prometheus.CounterValue, s.ReceiveTotalErrors, lv...)
		ch <- prometheus.MustNewConstMetric(d.transmitTotalErrorsDesc, prometheus.CounterValue, s.TransmitTotalErrors, lv...)
		ch <- prometheus.MustNewConstMetric(d.fecModeDesc, prometheus.CounterValue, s.FECMode, lv...)
	}
}

func convertFECModeToFloat64(s string) float64 {
	switch s {
	case "none":
		return 0
	case "fec74":
		return 2
	case "fec91":
		return 3
	case "fec108":
		return 4
	default:
		return 1
	}
}
