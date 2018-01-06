package interfaces

import "github.com/prometheus/client_golang/prometheus"

const prefix = "junos_interface_"

var (
	receiveBytesDesc   *prometheus.Desc
	receiveErrorsDesc  *prometheus.Desc
	receiveDropsDesc   *prometheus.Desc
	transmitBytesDesc  *prometheus.Desc
	transmitErrorsDesc *prometheus.Desc
	transmitDropsDesc  *prometheus.Desc
)

func init() {
	l := []string{"name", "description", "mac", "target"}
	receiveBytesDesc = prometheus.NewDesc(prefix+"interface_receive_bytes", "Received data in bytes", l, nil)
	receiveErrorsDesc = prometheus.NewDesc(prefix+"interface_receive_errors", "Number of errors caused by incoming packets", l, nil)
	receiveDropsDesc = prometheus.NewDesc(prefix+"interface_receive_drops", "Number of dropped incoming packets", l, nil)
	transmitBytesDesc = prometheus.NewDesc(prefix+"interface_transmit_bytes", "Transmitted data in bytes", l, nil)
	transmitErrorsDesc = prometheus.NewDesc(prefix+"interface_transmit_errors", "Number of errors caused by outgoing packets", l, nil)
	transmitDropsDesc = prometheus.NewDesc(prefix+"interface_transmit_drops", "Number of dropped outgoing packets", l, nil)
}

type InterfaceCollector struct {

}

func (*InterfaceCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- receiveBytesDesc
	ch <- receiveErrorsDesc
	ch <- receiveDropsDesc
	ch <- transmitBytesDesc
	ch <- transmitDropsDesc
	ch <- transmitErrorsDesc
}

func (c *InterfaceCollector) Collect(datasource InterfaceStatsDatasource, ch chan<- prometheus.Metric) error {
	stats, err := datasource.InterfaceStats()
	if err != nil {
		return err
	}

	for _, s := range stats {
		c.collectForInterface(s, ch)
	}

	return nil
}

func (*InterfaceCollector) collectForInterface(s *InterfaceStats, ch chan<- prometheus.Metric) {
	
}