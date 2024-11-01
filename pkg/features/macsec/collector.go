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
	macsecInterfaceDesc *prometheus.Desc
	macsecStatsDesc     *prometheus.Desc
)

// Initialize metrics descriptions
func init() {
	labelsInterface := []string{"host", "interface_name", "ca_name", "cipher_suit", "outgoing_packet_number", "sci", "created_since", "outbound_channel_status"}
	labelsStats := []string{"host", "interface", "secure_channel_sent_bytes_encrypted", "secure_channel_sent_pyckets_encrypted", "secure_channel_sent_bytes_protected", "secure_channel_sent_packets_protected", "secure_association_sent_packets_encrypted",
		"secure_association_sent_packets_protected", "secure_channel_receive_packets_accepted",
		"secure_channel_received_bytes_validated", "secure_channel_received_bytes_decrypted",
		"secure_association_received_packets_accepted", "secure_association_received_bytes_validated", "secure_association_received_bytes_decrypted"}
	macsecInterfaceDesc = prometheus.NewDesc(prefix+"interface_info", "Information regarding interface", labelsInterface, nil)
	macsecStatsDesc = prometheus.NewDesc(prefix+"stats_info", "Information regarding stats", labelsStats, nil)
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
	ch <- macsecInterfaceDesc
	ch <- macsecStatsDesc
}

// Collect collects metrics from JunOS
func (c *macsecCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var i resultInt
	err := client.RunCommandAndParse("show security macsec connections", &i)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show security macsec connections'")
	}
	c.collectForInterfaces(i, ch, labelValues)

	var s resultStats
	err = client.RunCommandAndParse("show security macsec statistics", &s)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show security macsec statistics'")
	}
	c.collectForStats(s, ch, labelValues)
	return nil
}

// collectForSessions collects metrics for the sessions
func (c *macsecCollector) collectForInterfaces(sessions resultInt, ch chan<- prometheus.Metric, labelValues []string) {
	for interfaceCounter := 0; interfaceCounter < (len(sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation)); interfaceCounter++ {
		labels := append(labelValues,
			sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation[interfaceCounter].InterfaceName,
			sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation[interfaceCounter].ConnectivityAssociationName,
			sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation[interfaceCounter].CipherSuite,
			sessions.MacsecConnectionInformation.OutboundSecureChannel[interfaceCounter].OutgoingPacketNumber,
			sessions.MacsecConnectionInformation.OutboundSecureChannel[interfaceCounter].Sci,
			sessions.MacsecConnectionInformation.OutboundSecureChannel[interfaceCounter].OutboundSecureAssociation.CreateTime.Text,
			sessions.MacsecConnectionInformation.OutboundSecureChannel[interfaceCounter].OutboundSecureAssociation.AssociationNumberStatus)
		pn, err := getPacketsNumber(sessions.MacsecConnectionInformation.OutboundSecureChannel[interfaceCounter].OutgoingPacketNumber)
		if err != nil {
			fmt.Printf("\n packet number is non-numerical. Maybe unmarshaling issues \n")
		}
		ch <- prometheus.MustNewConstMetric(macsecInterfaceDesc, prometheus.GaugeValue, float64(pn), labels...)
	}
}

func (c *macsecCollector) collectForStats(sessions resultStats, ch chan<- prometheus.Metric, labelValues []string) {
	for interfaceCounter := 0; interfaceCounter < (len(sessions.MacsecStatistics.Interfaces)); interfaceCounter++ {
		labels := append(labelValues,
			sessions.MacsecStatistics.Interfaces[interfaceCounter],
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].ProtectedPackets), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].ProtectedBytes), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].EncryptedBytes), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].ProtectedPackets), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureAssociationSent[interfaceCounter].ProtectedPackets), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureAssociationSent[interfaceCounter].EncryptedPackets), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelReceived[interfaceCounter].OkPackets), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelReceived[interfaceCounter].ValidatedBytes), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureChannelReceived[interfaceCounter].DecryptedBytes), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureAssociationReceived[interfaceCounter].OkPackets), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureAssociationReceived[interfaceCounter].ValidatedBytes), 10),
			strconv.FormatInt(int64(sessions.MacsecStatistics.SecureAssociationReceived[interfaceCounter].DecryptedBytes), 10))
		ch <- prometheus.MustNewConstMetric(macsecStatsDesc, prometheus.GaugeValue, float64(interfaceCounter), labels...)
	}
}

// stateToFloat converts the status string to a float value
func stateToFloat(status string) float64 {
	if status == "inuse" {
		return 1
	}
	return 0
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
