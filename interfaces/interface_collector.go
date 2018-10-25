package interfaces

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_"

var (
	receiveBytesDesc        *prometheus.Desc
	receivePacketsDesc      *prometheus.Desc
	receiveErrorsDesc       *prometheus.Desc
	receiveDropsDesc        *prometheus.Desc
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
)

func init() {
	l := []string{"target", "name", "description", "mac"}
	receiveBytesDesc = prometheus.NewDesc(prefix+"receive_bytes", "Received data in bytes", l, nil)
	receivePacketsDesc = prometheus.NewDesc(prefix+"receive_packets_total", "Received packets", l, nil)
	receiveErrorsDesc = prometheus.NewDesc(prefix+"receive_errors", "Number of errors caused by incoming packets", l, nil)
	receiveDropsDesc = prometheus.NewDesc(prefix+"receive_drops", "Number of dropped incoming packets", l, nil)
	transmitBytesDesc = prometheus.NewDesc(prefix+"transmit_bytes", "Transmitted data in bytes", l, nil)
	transmitPacketsDesc = prometheus.NewDesc(prefix+"transmit_packets_total", "Transmitted packets", l, nil)
	transmitErrorsDesc = prometheus.NewDesc(prefix+"transmit_errors", "Number of errors caused by outgoing packets", l, nil)
	transmitDropsDesc = prometheus.NewDesc(prefix+"transmit_drops", "Number of dropped outgoing packets", l, nil)
	ipv6receiveBytesDesc = prometheus.NewDesc(prefix+"IPv6_receive_bytes_total", "Received IPv6 data in bytes", l, nil)
	ipv6receivePacketsDesc = prometheus.NewDesc(prefix+"IPv6_receive_packets_total", "Received IPv6 packets", l, nil)
	ipv6transmitBytesDesc = prometheus.NewDesc(prefix+"IPv6_transmit_bytes_total", "Transmitted IPv6 data in bytes", l, nil)
	ipv6transmitPacketsDesc = prometheus.NewDesc(prefix+"IPv6_transmit_packets_total", "Transmitted IPv6 packets", l, nil)
	adminStatusDesc = prometheus.NewDesc(prefix+"admin_up", "Admin operational status", l, nil)
	operStatusDesc = prometheus.NewDesc(prefix+"up", "Interface operational status", l, nil)
	errorStatusDesc = prometheus.NewDesc(prefix+"error_status", "Admin and operational status differ", l, nil)
}

// Collector collects interface metrics
type interfaceCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &interfaceCollector{}
}

// Describe describes the metrics
func (*interfaceCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- receiveBytesDesc
	ch <- receivePacketsDesc
	ch <- receiveErrorsDesc
	ch <- receiveDropsDesc
	ch <- transmitBytesDesc
	ch <- transmitPacketsDesc
	ch <- transmitDropsDesc
	ch <- transmitErrorsDesc
	ch <- ipv6receiveBytesDesc
	ch <- ipv6receivePacketsDesc
	ch <- ipv6transmitBytesDesc
	ch <- ipv6transmitPacketsDesc
	ch <- adminStatusDesc
	ch <- operStatusDesc
	ch <- errorStatusDesc
}

// Collect collects metrics from JunOS
func (c *interfaceCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	stats, err := c.interfaceStats(client)
	if err != nil {
		return err
	}

	for _, s := range stats {
		c.collectForInterface(s, ch, labelValues)
	}

	return nil
}

func (c *interfaceCollector) interfaceStats(client *rpc.Client) ([]*InterfaceStats, error) {
	var x = InterfaceRpc{}
	err := client.RunCommandAndParse("show interfaces statistics detail", &x)
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
			TransmitDrops:       float64(phy.OutputErrors.Drops),
			TransmitErrors:      float64(phy.OutputErrors.Errors),
			TransmitBytes:       float64(phy.Stats.OutputBytes),
			TransmitPackets:     float64(phy.Stats.OutputPackets),
			IPv6ReceiveBytes:    float64(phy.Stats.IPv6Traffic.InputBytes),
			IPv6ReceivePackets:  float64(phy.Stats.IPv6Traffic.InputPackets),
			IPv6TransmitBytes:   float64(phy.Stats.IPv6Traffic.OutputBytes),
			IPv6TransmitPackets: float64(phy.Stats.IPv6Traffic.OutputPackets),
		}

		stats = append(stats, s)

		for _, log := range phy.LogicalInterfaces {
			sl := &InterfaceStats{
				IsPhysical:          false,
				Name:                log.Name,
				Description:         log.Description,
				Mac:                 phy.MacAddress,
				ReceiveBytes:        float64(log.Stats.InputBytes),
				ReceivePackets:      float64(log.Stats.InputPackets),
				TransmitBytes:       float64(log.Stats.OutputBytes),
				TransmitPackets:     float64(log.Stats.OutputPackets),
				IPv6ReceiveBytes:    float64(log.Stats.IPv6Traffic.InputBytes),
				IPv6ReceivePackets:  float64(log.Stats.IPv6Traffic.InputPackets),
				IPv6TransmitBytes:   float64(log.Stats.IPv6Traffic.OutputBytes),
				IPv6TransmitPackets: float64(log.Stats.IPv6Traffic.OutputPackets),
			}

			stats = append(stats, sl)
		}
	}

	return stats, nil
}

func (*interfaceCollector) collectForInterface(s *InterfaceStats, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{s.Name, s.Description, s.Mac}...)
	ch <- prometheus.MustNewConstMetric(receiveBytesDesc, prometheus.GaugeValue, s.ReceiveBytes, l...)
	ch <- prometheus.MustNewConstMetric(receivePacketsDesc, prometheus.GaugeValue, s.ReceivePackets, l...)
	ch <- prometheus.MustNewConstMetric(transmitBytesDesc, prometheus.GaugeValue, s.TransmitBytes, l...)
	ch <- prometheus.MustNewConstMetric(transmitPacketsDesc, prometheus.GaugeValue, s.TransmitPackets, l...)
	ch <- prometheus.MustNewConstMetric(ipv6receiveBytesDesc, prometheus.GaugeValue, s.IPv6ReceiveBytes, l...)
	ch <- prometheus.MustNewConstMetric(ipv6receivePacketsDesc, prometheus.GaugeValue, s.IPv6ReceivePackets, l...)
	ch <- prometheus.MustNewConstMetric(ipv6transmitBytesDesc, prometheus.GaugeValue, s.IPv6TransmitBytes, l...)
	ch <- prometheus.MustNewConstMetric(ipv6transmitPacketsDesc, prometheus.GaugeValue, s.IPv6TransmitPackets, l...)

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

		ch <- prometheus.MustNewConstMetric(adminStatusDesc, prometheus.GaugeValue, float64(adminUp), l...)
		ch <- prometheus.MustNewConstMetric(operStatusDesc, prometheus.GaugeValue, float64(operUp), l...)
		ch <- prometheus.MustNewConstMetric(errorStatusDesc, prometheus.GaugeValue, float64(err), l...)
		ch <- prometheus.MustNewConstMetric(transmitErrorsDesc, prometheus.GaugeValue, s.TransmitErrors, l...)
		ch <- prometheus.MustNewConstMetric(transmitDropsDesc, prometheus.GaugeValue, s.TransmitDrops, l...)
		ch <- prometheus.MustNewConstMetric(receiveErrorsDesc, prometheus.GaugeValue, s.ReceiveErrors, l...)
		ch <- prometheus.MustNewConstMetric(receiveDropsDesc, prometheus.GaugeValue, s.ReceiveDrops, l...)
	}
}
