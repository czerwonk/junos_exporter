package systemstatisticsIPv4

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

const prefix string = "junos_systemstatistics_ipv4_"

// Metrics to collect for the feature
var (
	packetsReceivedDesc                             *prometheus.Desc
	badHeaderChecksumsDesc                          *prometheus.Desc
	packetsWithSizeSmallerThanMinimumDesc           *prometheus.Desc
	packetsWithDataSizeLessThanDatalengthDesc       *prometheus.Desc
	packetsWithHeaderLengthLessThanHeaderlengthDesc *prometheus.Desc
	packetsWithIncorrectVersionNumberDesc           *prometheus.Desc
	packetsDestinedToDeadNextHopDesc                *prometheus.Desc
	fragmentsReceivedDesc                           *prometheus.Desc
	fragmentsDroppedDueToOutspaceOfDUPDesc          *prometheus.Desc
	fragmentsDroppedDueToQueueoverflowDesc          *prometheus.Desc
	fragmentsDroppedAfterTimeoutDesc                *prometheus.Desc
	packetsReassembledOKDesc                        *prometheus.Desc
	packetsForThisHostDesc                          *prometheus.Desc
	packetsForUnknownOrUnsupportedProtocolDesc      *prometheus.Desc
	packetsForwardedDesc                            *prometheus.Desc
	packetsNotForwardableDesc                       *prometheus.Desc
	redirectsSentDesc                               *prometheus.Desc
	packetsSentFromThisHostDesc                     *prometheus.Desc
	packetsSentWithFabricatedIPHeaderDesc           *prometheus.Desc
	outputPacketsDroppedDueToNoBufsDesc             *prometheus.Desc
	outputPacketsDiscardedDueToNoRouteDesc          *prometheus.Desc
	outputDatagramsFragmentedDesc                   *prometheus.Desc
	fragmentsCreatedDesc                            *prometheus.Desc
	datagramsThatCanNotBeFragmentedDesc             *prometheus.Desc
	packetsWithBadOptionsDesc                       *prometheus.Desc
	packetsWithOptionsHandledWithoutErrorDesc       *prometheus.Desc
	strictSourceAndRecordRouteOptionsDesc           *prometheus.Desc
	looseSourceAndRecordRouteOptionsDesc            *prometheus.Desc
	recordRouteOptionsDesc                          *prometheus.Desc
	timestampOptionsDesc                            *prometheus.Desc
	timestampAndAddressOptionsDesc                  *prometheus.Desc
	timestampAndPrespecifiedAddressOptionsDesc      *prometheus.Desc
	optionPacketsDroppedDueToRateLimitDesc          *prometheus.Desc
	routerAlertOptionDesc                           *prometheus.Desc
	multicastPacketsDroppedDesc                     *prometheus.Desc
	packetsDroppedDesc                              *prometheus.Desc
	transitREPacketsDroppedonMGMTInterfaceDesc      *prometheus.Desc
	packetsUsedFirstNexthopInECMPUnilistDesc        *prometheus.Desc
	incomingTtpoipPacketsReceivedDesc               *prometheus.Desc
	incomingTtpoipPacketsDroppedDesc                *prometheus.Desc
	outgoingTtpoipPacketsSentDesc                   *prometheus.Desc
	outgoingTtpoipPacketsDroppedDesc                *prometheus.Desc
	incomingRawIPPacketsDroppedNoSocketBufferDesc   *prometheus.Desc
	incomingVirtualNodePacketsDeliveredDesc         *prometheus.Desc
)

func init() {
	l := []string{"target"}
	packetsReceivedDesc = prometheus.NewDesc(prefix+"packets_received", "Number of packets received", l, nil)
	badHeaderChecksumsDesc = prometheus.NewDesc(prefix+"bad_header_checksums", "Number of packets received with bad header checksums", l, nil)
	packetsWithSizeSmallerThanMinimumDesc = prometheus.NewDesc(prefix+"packets_with_size_smaller_than_minimum", "Number of packets received with size smaller than minimum", l, nil)
	packetsWithDataSizeLessThanDatalengthDesc = prometheus.NewDesc(prefix+"packets_with_data_size_less_than_datalength", "Number of packets received with data size less than data length", l, nil)
	packetsWithHeaderLengthLessThanHeaderlengthDesc = prometheus.NewDesc(prefix+"packets_with_header_length_less_than_headerlength", "Number of packets received with header length less than header length", l, nil)
	packetsWithIncorrectVersionNumberDesc = prometheus.NewDesc(prefix+"packets_with_incorrect_version_number", "Number of packets received with incorrect version number", l, nil)
	packetsDestinedToDeadNextHopDesc = prometheus.NewDesc(prefix+"packets_destined_to_dead_next_hop", "Number of packets received destined to dead next hop", l, nil)
	fragmentsReceivedDesc = prometheus.NewDesc(prefix+"fragments_received", "Number of fragments received", l, nil)
	fragmentsDroppedDueToOutspaceOfDUPDesc = prometheus.NewDesc(prefix+"fragments_dropped_due_to_outspace_of_dup", "Number of fragments dropped due to outspace of DUP", l, nil)
	fragmentsDroppedDueToQueueoverflowDesc = prometheus.NewDesc(prefix+"fragments_dropped_due_to_queueoverflow", "Number of fragments dropped due to queue overflow", l, nil)
	fragmentsDroppedAfterTimeoutDesc = prometheus.NewDesc(prefix+"fragments_dropped_after_timeout", "Number of fragments dropped after timeout", l, nil)
	packetsReassembledOKDesc = prometheus.NewDesc(prefix+"packets_reassembled_ok", "Number of packets reassembled OK", l, nil)
	packetsForThisHostDesc = prometheus.NewDesc(prefix+"packets_for_this_host", "Number of packets for this host", l, nil)
	packetsForUnknownOrUnsupportedProtocolDesc = prometheus.NewDesc(prefix+"packets_for_unknown_or_unsupported_protocol", "Number of packets for unknown or unsupported protocol", l, nil)
	packetsForwardedDesc = prometheus.NewDesc(prefix+"packets_forwarded", "Number of packets forwarded", l, nil)
	packetsNotForwardableDesc = prometheus.NewDesc(prefix+"packets_not_forwardable", "Number of packets not forwardable", l, nil)
	redirectsSentDesc = prometheus.NewDesc(prefix+"redirects_sent", "Number of redirects sent", l, nil)
	packetsSentFromThisHostDesc = prometheus.NewDesc(prefix+"packets_sent_from_this_host", "Number of packets sent from this host", l, nil)
	packetsSentWithFabricatedIPHeaderDesc = prometheus.NewDesc(prefix+"packets_sent_with_fabricated_ip_header", "Number of packets sent with fabricated IP header", l, nil)
	outputPacketsDroppedDueToNoBufsDesc = prometheus.NewDesc(prefix+"output_packets_dropped_due_to_no_bufs", "Number of output packets dropped due to no bufs", l, nil)
	outputPacketsDiscardedDueToNoRouteDesc = prometheus.NewDesc(prefix+"output_packets_discarded_due_to_no_route", "Number of output packets discarded due to no route", l, nil)
	outputDatagramsFragmentedDesc = prometheus.NewDesc(prefix+"output_datagrams_fragmented", "Number of output datagrams fragmented", l, nil)
	fragmentsCreatedDesc = prometheus.NewDesc(prefix+"fragments_created", "Number of fragments created", l, nil)
	datagramsThatCanNotBeFragmentedDesc = prometheus.NewDesc(prefix+"datagrams_that_can_not_be_fragmented", "Number of datagrams that can not be fragmented", l, nil)
	packetsWithBadOptionsDesc = prometheus.NewDesc(prefix+"packets_with_bad_options", "Number of packets with bad options", l, nil)
	packetsWithOptionsHandledWithoutErrorDesc = prometheus.NewDesc(prefix+"packets_with_options_handled_without_error", "Number of packets with options handled without error", l, nil)
	strictSourceAndRecordRouteOptionsDesc = prometheus.NewDesc(prefix+"strict_source_and_record_route_options", "Number of packets with strict source and record route options", l, nil)
	looseSourceAndRecordRouteOptionsDesc = prometheus.NewDesc(prefix+"loose_source_and_record_route_options", "Number of packets with loose source and record route options", l, nil)
	recordRouteOptionsDesc = prometheus.NewDesc(prefix+"record_route_options", "Number of packets with record route options", l, nil)
	timestampOptionsDesc = prometheus.NewDesc(prefix+"timestamp_options", "Number of packets with timestamp options", l, nil)
	timestampAndAddressOptionsDesc = prometheus.NewDesc(prefix+"timestamp_and_address_options", "Number of packets with timestamp and address options", l, nil)
	timestampAndPrespecifiedAddressOptionsDesc = prometheus.NewDesc(prefix+"timestamp_and_prespecified_address_options", "Number of packets with timestamp and prespecified address options", l, nil)
	optionPacketsDroppedDueToRateLimitDesc = prometheus.NewDesc(prefix+"option_packets_dropped_due_to_rate_limit", "Number of option packets dropped due to rate limit", l, nil)
	routerAlertOptionDesc = prometheus.NewDesc(prefix+"router_alert_option", "Number of packets with router alert option", l, nil)
	multicastPacketsDroppedDesc = prometheus.NewDesc(prefix+"multicast_packets_dropped", "Number of multicast packets dropped", l, nil)
	packetsDroppedDesc = prometheus.NewDesc(prefix+"packets_dropped", "Number of packets dropped", l, nil)
	transitREPacketsDroppedonMGMTInterfaceDesc = prometheus.NewDesc(prefix+"transit_re_packets_droppedon_mgt_interface", "Number of transit RE packets dropped on MGMT interface", l, nil)
	packetsUsedFirstNexthopInECMPUnilistDesc = prometheus.NewDesc(prefix+"packets_used_first_nexthop_in_ecmp_unilist", "Number of packets used first nexthop in ECMP unilist", l, nil)
	incomingTtpoipPacketsReceivedDesc = prometheus.NewDesc(prefix+"incoming_ttpoip_packets_received", "Number of incoming TTPoIP packets received", l, nil)
	incomingTtpoipPacketsDroppedDesc = prometheus.NewDesc(prefix+"incoming_ttpoip_packets_dropped", "Number of incoming TTPoIP packets dropped", l, nil)
	outgoingTtpoipPacketsSentDesc = prometheus.NewDesc(prefix+"outgoing_ttpoip_packets_sent", "Number of outgoing TTPoIP packets sent", l, nil)
	outgoingTtpoipPacketsDroppedDesc = prometheus.NewDesc(prefix+"outgoing_ttpoip_packets_dropped", "Number of outgoing TTPoIP packets dropped", l, nil)
	incomingRawIPPacketsDroppedNoSocketBufferDesc = prometheus.NewDesc(prefix+"incoming_raw_ip_packets_dropped_no_socket_buffer", "Number of incoming raw IP packets dropped due to no socket buffer", l, nil)
	incomingVirtualNodePacketsDeliveredDesc = prometheus.NewDesc(prefix+"incoming_virtual_node_packets_delivered", "Number of incoming virtual node packets delivered", l, nil)
}

type systemstatisticsIPv4Collector struct{}

func NewCollector() collector.RPCCollector {
	return &systemstatisticsIPv4Collector{}
}

func (c *systemstatisticsIPv4Collector) Name() string {
	return "systemstatisticsIPv4"
}

func (c *systemstatisticsIPv4Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- packetsReceivedDesc
	ch <- badHeaderChecksumsDesc
	ch <- packetsWithSizeSmallerThanMinimumDesc
	ch <- packetsWithDataSizeLessThanDatalengthDesc
	ch <- packetsWithHeaderLengthLessThanHeaderlengthDesc
	ch <- packetsWithIncorrectVersionNumberDesc
	ch <- packetsDestinedToDeadNextHopDesc
	ch <- fragmentsReceivedDesc
	ch <- fragmentsDroppedDueToOutspaceOfDUPDesc
	ch <- fragmentsDroppedDueToQueueoverflowDesc
	ch <- fragmentsDroppedAfterTimeoutDesc
	ch <- packetsReassembledOKDesc
	ch <- packetsForThisHostDesc
	ch <- packetsForUnknownOrUnsupportedProtocolDesc
	ch <- packetsForwardedDesc
	ch <- packetsNotForwardableDesc
	ch <- redirectsSentDesc
	ch <- packetsSentFromThisHostDesc
	ch <- packetsSentWithFabricatedIPHeaderDesc
	ch <- outputPacketsDroppedDueToNoBufsDesc
	ch <- outputPacketsDiscardedDueToNoRouteDesc
	ch <- outputDatagramsFragmentedDesc
	ch <- fragmentsCreatedDesc
	ch <- datagramsThatCanNotBeFragmentedDesc
	ch <- packetsWithBadOptionsDesc
	ch <- packetsWithOptionsHandledWithoutErrorDesc
	ch <- strictSourceAndRecordRouteOptionsDesc
	ch <- looseSourceAndRecordRouteOptionsDesc
	ch <- recordRouteOptionsDesc
	ch <- timestampOptionsDesc
	ch <- timestampAndAddressOptionsDesc
	ch <- timestampAndPrespecifiedAddressOptionsDesc
	ch <- optionPacketsDroppedDueToRateLimitDesc
	ch <- routerAlertOptionDesc
	ch <- multicastPacketsDroppedDesc
	ch <- packetsDroppedDesc
	ch <- transitREPacketsDroppedonMGMTInterfaceDesc
	ch <- packetsUsedFirstNexthopInECMPUnilistDesc
	ch <- incomingTtpoipPacketsReceivedDesc
	ch <- incomingTtpoipPacketsDroppedDesc
	ch <- outgoingTtpoipPacketsSentDesc
	ch <- outgoingTtpoipPacketsDroppedDesc
	ch <- incomingRawIPPacketsDroppedNoSocketBufferDesc
	ch <- incomingVirtualNodePacketsDeliveredDesc
}

func (c *systemstatisticsIPv4Collector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var s StatisticsIPv4
	err := client.RunCommandAndParse("show system statistics ipv", &s)
	if err != nil {
		return err
	}

}
