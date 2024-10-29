// This plugin for MACsec collects metrics from the command "show security macsec connections".
package macsec

import (
	"fmt"
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"strings"
)

const prefix string = "junos_macsec_"

// Metrics to collect for the feature
var (
	macsecConnectionInformationDesc      *prometheus.Desc
	macsecInterfaceCommonInformationDesc *prometheus.Desc
	macsecOffsetDesc                     *prometheus.Desc
	macsecInboundPacketsDesc             *prometheus.Desc
	macsecOutboundPacketsDesc            *prometheus.Desc
)

// Initialize metrics descriptions
func init() {
	//labels := []string{"interface_name", "ca_name", "cipher_suite", "sci"}
	labels := []string{"host", "interface_name", "cipher_suit", "outgoing_packet_number", "sci", "created_since"}
	macsecConnectionInformationDesc = prometheus.NewDesc(prefix+"connection_info", "Interfaces that have macsec", labels, nil)
	macsecInterfaceCommonInformationDesc = prometheus.NewDesc(prefix+"amount_of_connections", "Information of specific interface", labels, nil)
	macsecOffsetDesc = prometheus.NewDesc(prefix+"offset", "Information regarding the offset", labels, nil)
	macsecInboundPacketsDesc = prometheus.NewDesc(prefix+"inbound_packets", "Information of inbound packets", labels, nil)
	macsecOutboundPacketsDesc = prometheus.NewDesc(prefix+"outbound_packets", "Information of outbound packets", labels, nil)
}

// macsecCollector collects MACsec metrics
type macsecCollector struct{}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &macsecCollector{}
}

// Name returns the name of the collector
func (*macsecCollector) Name() string {
	return "MACsec"
}

// Describe describes the metrics
func (*macsecCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- macsecConnectionInformationDesc
	ch <- macsecInterfaceCommonInformationDesc
	ch <- macsecOffsetDesc
	ch <- macsecInboundPacketsDesc
	ch <- macsecOutboundPacketsDesc
}

// Collect collects metrics from JunOS
func (c *macsecCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	fmt.Printf("\n INSIDE COLLECT")
	var x result
	err := client.RunCommandAndParse("show security macsec connections", &x)
	if err != nil {
		return errors.Wrap(err, "failed to run command")
	}

	c.collectForSessions(x, ch, labelValues)
	return nil
}

// collectForSessions collects metrics for the sessions
func (c *macsecCollector) collectForSessions(sessions result, ch chan<- prometheus.Metric, labelValues []string) {
	//var pn int
	//var err error
	//var labels []string
	var addinionalLables []string
	for _, iface := range sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation {
		fmt.Printf("\n iface is %v", iface)
		//labels := append(labelValues, iface.InterfaceName, iface.ConnectivityAssociationName, iface.CipherSuite /*, iface.Text*/)

		// Collecting the number of connections
		//ch <- prometheus.MustNewConstMetric(macsecInterfaceCommonInformationDesc, prometheus.GaugeValue, float64(len(sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation)), labels...)

		for _, interfaceInfo := range sessions.MacsecConnectionInformation.OutboundSecureChannel {
			pn, err := getPacketsNumber(interfaceInfo.OutgoingPacketNumber)
			if err != nil {
				fmt.Printf("\n packet number is non-numerical. Maybe unmarshaling issues \n")
			}
			fmt.Printf(" \\ņ label values are %s \n", labelValues)
			//for _, outChannel := range sessions.MacsecConnectionInformation.OutboundSecureChannel {
			//	labels = append(labelValues, iface.InterfaceName, iface.CipherSuite, outChannel.OutgoingPacketNumber, outChannel.Sci, outChannel.OutboundSecureAssociation.CreateTime.Text)
			//}
			for _, outChannel := range sessions.MacsecConnectionInformation.OutboundSecureChannel {
				addinionalLables = append(labelValues, iface.InterfaceName, iface.CipherSuite, outChannel.OutgoingPacketNumber, outChannel.Sci, outChannel.OutboundSecureAssociation.CreateTime.Text)
			}
			// Collecting outbound packets
			fmt.Printf(" \\n labels are %s \n", labels)
			//ch <- prometheus.MustNewConstMetric(macsecOutboundPacketsDesc, prometheus.GaugeValue, float64(pn), labels...)
		}
		ch <- prometheus.MustNewConstMetric(macsecOutboundPacketsDesc, prometheus.GaugeValue, float64(pn), labels...)
	}
}

// stateToFloat converts the status string to a float value
func stateToFloat(status string) float64 {
	if status == "inuse" {
		return 1
	}
	return 0
}

// getNumberOfConnections returns the number of connections
func getNumberOfConnections(connections []string) (int, error) {
	if len(connections) == 0 {
		return 0, errors.New("No connections")
	}
	return len(connections), nil
}

// getInterfaceNumber converts interface name to number
func getInterfaceNumber(nameAsString string) (int, error) {
	result := strings.SplitAfter(nameAsString, "/")
	i, err := strconv.Atoi(result[len(result)-1])
	if err != nil {
		return 0, err
	}
	return i, nil
}

// getPacketsNumber converts packet number string to integer
func getPacketsNumber(packetsAsString string) (int, error) {
	i, err := strconv.Atoi(packetsAsString)
	if err != nil {
		return 0, err
	}
	return i, nil
}

/*
func getOutboundLabels(sessions result) (int, string, string, err){
	for _, iface := range sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation {
		packet_num := iface.
	}

	}
*/
