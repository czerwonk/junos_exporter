package interfaces

import (
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/interfacelabels"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_"

// Collector collects interface metrics
type interfaceCollector struct {
	labels                  *interfacelabels.DynamicLabels
	receiveBytesDesc        *prometheus.Desc
	receivePacketsDesc      *prometheus.Desc
	receiveErrorsDesc       *prometheus.Desc
	receiveDropsDesc        *prometheus.Desc
	interfaceSpeedDesc      *prometheus.Desc
	transmitBytesDesc       *prometheus.Desc
	transmitPacketsDesc     *prometheus.Desc
	transmitErrorsDesc      *prometheus.Desc
	transmitDropsDesc       *prometheus.Desc
	ipv6receiveBytesDesc    *prometheus.Desc
	ipv6receivePacketsDesc  *prometheus.Desc
	ipv6transmitBytesDesc   *prometheus.Desc
	ipv6transmitPacketsDesc *prometheus.Desc
	adminStatusDesc         *prometheus.Desc
	operStatusDesc          *prometheus.Desc
	errorStatusDesc         *prometheus.Desc
	lastFlappedDesc         *prometheus.Desc
}

// NewCollector creates a new collector
func NewCollector(labels *interfacelabels.DynamicLabels) collector.RPCCollector {
	c := &interfaceCollector{
		labels: labels,
	}
	c.init()

	return c
}

// Name returns the name of the collector
func (*interfaceCollector) Name() string {
	return "Interfaces"
}

func (c *interfaceCollector) init() {
	l := []string{"target", "name", "description", "mac"}
	l = append(l, c.labels.LabelNames()...)

	c.receiveBytesDesc = prometheus.NewDesc(prefix+"receive_bytes", "Received data in bytes", l, nil)
	c.receivePacketsDesc = prometheus.NewDesc(prefix+"receive_packets_total", "Received packets", l, nil)
	c.receiveErrorsDesc = prometheus.NewDesc(prefix+"receive_errors", "Number of errors caused by incoming packets", l, nil)
	c.receiveDropsDesc = prometheus.NewDesc(prefix+"receive_drops", "Number of dropped incoming packets", l, nil)
	c.interfaceSpeedDesc = prometheus.NewDesc(prefix+"speed", "speed in in bps", l, nil)
	c.transmitBytesDesc = prometheus.NewDesc(prefix+"transmit_bytes", "Transmitted data in bytes", l, nil)
	c.transmitPacketsDesc = prometheus.NewDesc(prefix+"transmit_packets_total", "Transmitted packets", l, nil)
	c.transmitErrorsDesc = prometheus.NewDesc(prefix+"transmit_errors", "Number of errors caused by outgoing packets", l, nil)
	c.transmitDropsDesc = prometheus.NewDesc(prefix+"transmit_drops", "Number of dropped outgoing packets", l, nil)
	c.ipv6receiveBytesDesc = prometheus.NewDesc(prefix+"IPv6_receive_bytes_total", "Received IPv6 data in bytes", l, nil)
	c.ipv6receivePacketsDesc = prometheus.NewDesc(prefix+"IPv6_receive_packets_total", "Received IPv6 packets", l, nil)
	c.ipv6transmitBytesDesc = prometheus.NewDesc(prefix+"IPv6_transmit_bytes_total", "Transmitted IPv6 data in bytes", l, nil)
	c.ipv6transmitPacketsDesc = prometheus.NewDesc(prefix+"IPv6_transmit_packets_total", "Transmitted IPv6 packets", l, nil)
	c.adminStatusDesc = prometheus.NewDesc(prefix+"admin_up", "Admin operational status", l, nil)
	c.operStatusDesc = prometheus.NewDesc(prefix+"up", "Interface operational status", l, nil)
	c.errorStatusDesc = prometheus.NewDesc(prefix+"error_status", "Admin and operational status differ", l, nil)
	c.lastFlappedDesc = prometheus.NewDesc(prefix+"last_flapped_seconds", "Seconds since last flapped (-1 if never)", l, nil)
}

// Describe describes the metrics
func (c *interfaceCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.receiveBytesDesc
	ch <- c.receivePacketsDesc
	ch <- c.receiveErrorsDesc
	ch <- c.receiveDropsDesc
	ch <- c.interfaceSpeedDesc
	ch <- c.transmitBytesDesc
	ch <- c.transmitPacketsDesc
	ch <- c.transmitDropsDesc
	ch <- c.transmitErrorsDesc
	ch <- c.ipv6receiveBytesDesc
	ch <- c.ipv6receivePacketsDesc
	ch <- c.ipv6transmitBytesDesc
	ch <- c.ipv6transmitPacketsDesc
	ch <- c.adminStatusDesc
	ch <- c.operStatusDesc
	ch <- c.errorStatusDesc
	ch <- c.lastFlappedDesc
}

// Collect collects metrics from JunOS
func (c *interfaceCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	stats, err := c.interfaceStats(client)
	if err != nil {
		return err
	}

	for _, s := range stats {
		c.collectForInterface(s, client.Device(), ch, labelValues)
	}

	return nil
}

func (c *interfaceCollector) interfaceStats(client *rpc.Client) ([]*InterfaceStats, error) {
	var x = InterfaceRpc{}
	err := client.RunCommandAndParse("show interfaces statistics detail \"*-*\"", &x)
	if err != nil {
		return nil, err
	}

	stats := make([]*InterfaceStats, 0)
	for _, phy := range x.Information.Interfaces {
		s := &InterfaceStats{
			IsPhysical:          true,
			Name:                phy.Name,
			AdminStatus:         phy.AdminStatus == "up",
			OperStatus:          phy.OperStatus == "up",
			ErrorStatus:         !(phy.AdminStatus == phy.OperStatus),
			Description:         phy.Description,
			Mac:                 phy.MacAddress,
			ReceiveDrops:        float64(phy.InputErrors.Drops),
			ReceiveErrors:       float64(phy.InputErrors.Errors),
			ReceiveBytes:        float64(phy.Stats.InputBytes),
			ReceivePackets:      float64(phy.Stats.InputPackets),
			Speed:               phy.Speed,
			TransmitDrops:       float64(phy.OutputErrors.Drops),
			TransmitErrors:      float64(phy.OutputErrors.Errors),
			TransmitBytes:       float64(phy.Stats.OutputBytes),
			TransmitPackets:     float64(phy.Stats.OutputPackets),
			IPv6ReceiveBytes:    float64(phy.Stats.IPv6Traffic.InputBytes),
			IPv6ReceivePackets:  float64(phy.Stats.IPv6Traffic.InputPackets),
			IPv6TransmitBytes:   float64(phy.Stats.IPv6Traffic.OutputBytes),
			IPv6TransmitPackets: float64(phy.Stats.IPv6Traffic.OutputPackets),
			LastFlapped:         -1,
		}

		if phy.InterfaceFlapped.Value != "Never" {
			s.LastFlapped = float64(phy.InterfaceFlapped.Seconds)
		}

		stats = append(stats, s)

		for _, log := range phy.LogicalInterfaces {
			var s TrafficStat
			if (log.Stats != TrafficStat{}) {
				s = log.Stats
			} else {
				s = log.LagStats.Stats
			}
			sl := &InterfaceStats{
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

func (c *interfaceCollector) collectForInterface(s *InterfaceStats, device *connector.Device, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Name, s.Description, s.Mac}...)
	l = append(l, c.labels.ValuesForInterface(device, s.Name)...)

	ch <- prometheus.MustNewConstMetric(c.receiveBytesDesc, prometheus.GaugeValue, s.ReceiveBytes, l...)
	ch <- prometheus.MustNewConstMetric(c.receivePacketsDesc, prometheus.GaugeValue, s.ReceivePackets, l...)
	ch <- prometheus.MustNewConstMetric(c.transmitBytesDesc, prometheus.GaugeValue, s.TransmitBytes, l...)
	ch <- prometheus.MustNewConstMetric(c.transmitPacketsDesc, prometheus.GaugeValue, s.TransmitPackets, l...)
	ch <- prometheus.MustNewConstMetric(c.ipv6receiveBytesDesc, prometheus.GaugeValue, s.IPv6ReceiveBytes, l...)
	ch <- prometheus.MustNewConstMetric(c.ipv6receivePacketsDesc, prometheus.GaugeValue, s.IPv6ReceivePackets, l...)
	ch <- prometheus.MustNewConstMetric(c.ipv6transmitBytesDesc, prometheus.GaugeValue, s.IPv6TransmitBytes, l...)
	ch <- prometheus.MustNewConstMetric(c.ipv6transmitPacketsDesc, prometheus.GaugeValue, s.IPv6TransmitPackets, l...)

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
		if strings.Contains(s.Speed, "mbps") {
			speed = strings.Replace(s.Speed, "mbps", "000000", 1)
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
				speed = strings.Replace(s.Speed, "Unspecified", "0", 1)
			}
		}
		if strings.Contains(s.Speed, "Unlimited") {
			speed = strings.Replace(s.Speed, "Unlimited", "0", 1)
		}

		sp64, _ := strconv.ParseFloat(speed, 64)

		ch <- prometheus.MustNewConstMetric(c.adminStatusDesc, prometheus.GaugeValue, float64(adminUp), l...)
		ch <- prometheus.MustNewConstMetric(c.operStatusDesc, prometheus.GaugeValue, float64(operUp), l...)
		ch <- prometheus.MustNewConstMetric(c.errorStatusDesc, prometheus.GaugeValue, float64(err), l...)
		ch <- prometheus.MustNewConstMetric(c.transmitErrorsDesc, prometheus.GaugeValue, s.TransmitErrors, l...)
		ch <- prometheus.MustNewConstMetric(c.transmitDropsDesc, prometheus.GaugeValue, s.TransmitDrops, l...)
		ch <- prometheus.MustNewConstMetric(c.receiveErrorsDesc, prometheus.GaugeValue, s.ReceiveErrors, l...)
		ch <- prometheus.MustNewConstMetric(c.receiveDropsDesc, prometheus.GaugeValue, s.ReceiveDrops, l...)
		ch <- prometheus.MustNewConstMetric(c.interfaceSpeedDesc, prometheus.GaugeValue, float64(sp64), l...)

		if s.LastFlapped != 0 {
			ch <- prometheus.MustNewConstMetric(c.lastFlappedDesc, prometheus.GaugeValue, s.LastFlapped, l...)
		}
	}
}
