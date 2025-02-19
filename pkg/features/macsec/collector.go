// This plugin for MACsec collects metrics from the command "show security macsec connections".
package macsec

import (
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const prefix string = "junos_macsec_"

// Metrics to collect for the feature
var (
	macsecTXPacketCountDesc                       *prometheus.Desc
	macsecTXChannelStatusDesc                     *prometheus.Desc
	macsecIncludeSCIDesc                          *prometheus.Desc
	macsecReplayProtectDesc                       *prometheus.Desc
	macsecKeyServerOffsetDesc                     *prometheus.Desc
	macsecEncryptionDesc                          *prometheus.Desc
	macsecSecureChannelTXEncryptedPacketsDesc     *prometheus.Desc
	macsecSecureChannelTXEncryptedBytessDesc      *prometheus.Desc
	macsecSecureChannelTXProtectedPacketsDesc     *prometheus.Desc
	macsecSecureChannelTXProtectedBytesDesc       *prometheus.Desc
	macsecSecureAssociationTXEncryptedPacketsDesc *prometheus.Desc
	macsecSecureAssociationTXProtectedPacketsDesc *prometheus.Desc
	macsecSecureChannelRXAcceptedPacketsDesc      *prometheus.Desc
	macsecSecureChannelRXValidatedBytesDesc       *prometheus.Desc
	macsecSecureChannelRXDecryptedBytesDesc       *prometheus.Desc
	macsecSecureAssociationRXAcceptedPacketsDesc  *prometheus.Desc
	macsecSecureAssociationRXValidatedBytesDesc   *prometheus.Desc
	macsecSecureAssociationRXDecryptedBytesDesc   *prometheus.Desc
)

// Initialize metrics descriptions
func init() {
	labelsInterface := []string{"target", "interface", "ca"}
	labelsStats := []string{"target", "interface"}
	macsecTXPacketCountDesc = prometheus.NewDesc(prefix+"interface_transmit_packet_count", "Information regarding transmitted packets by interface", labelsInterface, nil)
	macsecTXChannelStatusDesc = prometheus.NewDesc(prefix+"tx_channel_status", "Information regarding the status of outbound channel secure association. 1 for inuse", labelsInterface, nil)
	macsecIncludeSCIDesc = prometheus.NewDesc(prefix+"sci", "Information regarding if sci is included in the interface. 0 for not included, 1 for included, 2 for unknown", labelsInterface, nil)
	macsecReplayProtectDesc = prometheus.NewDesc(prefix+"replay_protect", "Information if replay protect is on or off. 0 for off, 1 for on, 2 for unknown", labelsInterface, nil)
	macsecKeyServerOffsetDesc = prometheus.NewDesc(prefix+"key_server_offset", "Information regarding key server offset", labelsInterface, nil)
	macsecEncryptionDesc = prometheus.NewDesc(prefix+"encryption", "Information regarding encryption. 0 for off, 1 for on, 2 for unknown", labelsInterface, nil)
	macsecSecureChannelTXEncryptedPacketsDesc = prometheus.NewDesc(prefix+"statistics_secure_channel_tx_encrypted_packets_count", "Amount of secure channel sent encrypted packets", labelsStats, nil)
	macsecSecureChannelTXEncryptedBytessDesc = prometheus.NewDesc(prefix+"statistics_secure_channel_tx_encrypted_bytes_count", "Amount of secure channel sent encrypted bytes", labelsStats, nil)
	macsecSecureChannelTXProtectedPacketsDesc = prometheus.NewDesc(prefix+"statistics_secure_channel_tx_protected_packets_count", "Amount of secure channel sent protected packets", labelsStats, nil)
	macsecSecureChannelTXProtectedBytesDesc = prometheus.NewDesc(prefix+"statistics_secure_channel_tx_protected_bytes_count", "Amount of secure channel sent protected bytes", labelsStats, nil)
	macsecSecureAssociationTXEncryptedPacketsDesc = prometheus.NewDesc(prefix+"statistics_secure_association_tx_encrypted_packets_count", "Amount of secure association sent encrypted packets", labelsStats, nil)
	macsecSecureAssociationTXProtectedPacketsDesc = prometheus.NewDesc(prefix+"statistics_secure_association_tx_protected_packets_count", "Amount of secure association sent protected packets", labelsStats, nil)
	macsecSecureChannelRXAcceptedPacketsDesc = prometheus.NewDesc(prefix+"statistics_secure_channel_rx_accepted_packets_count", "Amount of secure channel received accepted packets", labelsStats, nil)
	macsecSecureChannelRXValidatedBytesDesc = prometheus.NewDesc(prefix+"secure_channel_rx_validated_bytes_count", "Amount of secure channel received validated bytes", labelsStats, nil)
	macsecSecureChannelRXDecryptedBytesDesc = prometheus.NewDesc(prefix+"secure_channel_rx_decrypted_bytes_count", "Amount of secure channel received decrypted bytes", labelsStats, nil)
	macsecSecureAssociationRXAcceptedPacketsDesc = prometheus.NewDesc(prefix+"statistics_secure_association_rx_accepted_packets_count", "Amount of secure association received accepted packets", labelsStats, nil)
	macsecSecureAssociationRXValidatedBytesDesc = prometheus.NewDesc(prefix+"statistics_secure_association_rx_validated_bytes_count", "Amount of secure association received validated bytes", labelsStats, nil)
	macsecSecureAssociationRXDecryptedBytesDesc = prometheus.NewDesc(prefix+"statistics_secure_association_rx_decrypted_bytes_count", "Amount of secure association received decrypted bytes", labelsStats, nil)
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
	ch <- macsecTXPacketCountDesc
	ch <- macsecTXChannelStatusDesc
	ch <- macsecIncludeSCIDesc
	ch <- macsecReplayProtectDesc
	ch <- macsecKeyServerOffsetDesc
	ch <- macsecEncryptionDesc
	ch <- macsecSecureChannelTXEncryptedPacketsDesc
	ch <- macsecSecureChannelTXEncryptedBytessDesc
	ch <- macsecSecureChannelTXProtectedPacketsDesc
	ch <- macsecSecureChannelTXProtectedBytesDesc
	ch <- macsecSecureAssociationTXEncryptedPacketsDesc
	ch <- macsecSecureAssociationTXProtectedPacketsDesc
	ch <- macsecSecureChannelRXAcceptedPacketsDesc
	ch <- macsecSecureChannelRXValidatedBytesDesc
	ch <- macsecSecureChannelRXDecryptedBytesDesc
	ch <- macsecSecureAssociationRXAcceptedPacketsDesc
	ch <- macsecSecureAssociationRXValidatedBytesDesc
	ch <- macsecSecureAssociationRXDecryptedBytesDesc
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
	for c, mici := range sessions.MacsecConnectionInformation.MacsecInterfaceCommonInformation {
		labels := append(labelValues,
			mici.InterfaceName,
			mici.ConnectivityAssociationName)
		pn, err := strconv.Atoi(sessions.MacsecConnectionInformation.OutboundSecureChannel[c].OutgoingPacketNumber)
		if err != nil {
			log.Errorf("unable to convert outgoing packets number: %q", sessions.MacsecConnectionInformation.OutboundSecureChannel[c].OutgoingPacketNumber)
		}
		sci := convertYesNoToInt(strings.TrimRight(mici.IncludeSci, "\n"))
		rp := convertOnOffToInt(strings.TrimRight(mici.ReplayProtect, "\n"))
		kso, err := strconv.Atoi(mici.Offset)
		if err != nil {
			log.Errorf("unable to convert offset: %q", mici.Offset)
		}
		status := stateToFloat(sessions.MacsecConnectionInformation.OutboundSecureChannel[c].OutboundSecureAssociation.AssociationNumberStatus)
		enc := convertOnOffToInt(strings.TrimRight(mici.Encryption, "\n"))
		ch <- prometheus.MustNewConstMetric(macsecTXPacketCountDesc, prometheus.CounterValue, float64(pn), labels...)
		ch <- prometheus.MustNewConstMetric(macsecIncludeSCIDesc, prometheus.GaugeValue, float64(sci), labels...)
		ch <- prometheus.MustNewConstMetric(macsecReplayProtectDesc, prometheus.GaugeValue, float64(rp), labels...)
		ch <- prometheus.MustNewConstMetric(macsecKeyServerOffsetDesc, prometheus.GaugeValue, float64(kso), labels...)
		ch <- prometheus.MustNewConstMetric(macsecEncryptionDesc, prometheus.GaugeValue, float64(enc), labels...)
		ch <- prometheus.MustNewConstMetric(macsecTXChannelStatusDesc, prometheus.GaugeValue, status, labels...)
	}
}

func (c *macsecCollector) collectForStats(sessions resultStats, ch chan<- prometheus.Metric, labelValues []string) {
	for interfaceCounter := 0; interfaceCounter < (len(sessions.MacsecStatistics.Interfaces)); interfaceCounter++ {
		labels := append(labelValues,
			sessions.MacsecStatistics.Interfaces[interfaceCounter])
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelTXEncryptedPacketsDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].EncryptedPackets), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelTXEncryptedBytessDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].EncryptedBytes), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelTXProtectedPacketsDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].ProtectedPackets), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelTXProtectedBytesDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelSent[interfaceCounter].ProtectedBytes), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureAssociationTXEncryptedPacketsDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureAssociationSent[interfaceCounter].EncryptedPackets), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureAssociationTXProtectedPacketsDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureAssociationSent[interfaceCounter].ProtectedPackets), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelRXAcceptedPacketsDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelReceived[interfaceCounter].OkPackets), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelRXValidatedBytesDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelReceived[interfaceCounter].ValidatedBytes), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureChannelRXDecryptedBytesDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureChannelReceived[interfaceCounter].DecryptedBytes), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureAssociationRXAcceptedPacketsDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureAssociationReceived[interfaceCounter].OkPackets), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureAssociationRXValidatedBytesDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureAssociationReceived[interfaceCounter].ValidatedBytes), labels...)
		ch <- prometheus.MustNewConstMetric(macsecSecureAssociationRXDecryptedBytesDesc, prometheus.CounterValue, float64(sessions.MacsecStatistics.SecureAssociationReceived[interfaceCounter].DecryptedBytes), labels...)
	}
}

// stateToFloat converts the status string to a float value
func stateToFloat(status string) float64 {
	if strings.TrimRight(status, "\n") == "inuse" {
		return 1
	}
	return 0
}

// convertYesNoToInt returns 0, 1 or 2  depending on the input string value
func convertYesNoToInt(s string) int {
	switch s {
	case "no":
		return 0
	case "yes":
		return 1
	default:
		return 2
	}
}

// convertOnOffToInt returns 0, 1 or 2  depending on the input string value
func convertOnOffToInt(s string) int {
	switch s {
	case "off":
		return 0
	case "on":
		return 1
	default:
		return 2
	}
}
