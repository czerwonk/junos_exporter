package systemstatistics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

const prefix string = "junos_systemstatistics_"

var (
	ipv4PacketsReceivedDesc                           *prometheus.Desc
	ipv4BadHeaderChecksumsDesc                        *prometheus.Desc
	ipv4PacketsWithSizeSmallerThanMinimumDesc         *prometheus.Desc
	ipv4PacketsWithDataSizeLessThanDatalengthDesc     *prometheus.Desc
	ipv4PacketsWithHeaderLengthLessThanDataSizeDesc   *prometheus.Desc
	ipv4PacketsWithIncorrectVersionNumberDesc         *prometheus.Desc
	ipv4PacketsDestinedToDeadNextHopDesc              *prometheus.Desc
	ipv4FragmentsReceivedDesc                         *prometheus.Desc
	ipv4FragmentsDroppedDueToOutspaceOrDUPDesc        *prometheus.Desc
	ipv4FragmentsDroppedDueToQueueoverflowDesc        *prometheus.Desc
	ipv4FragmentsDroppedAfterTimeoutDesc              *prometheus.Desc
	ipv4PacketsReassembledOKDesc                      *prometheus.Desc
	ipv4PacketsForThisHostDesc                        *prometheus.Desc
	ipv4PacketsForUnknownOrUnsupportedProtocolDesc    *prometheus.Desc
	ipv4PacketsForwardedDesc                          *prometheus.Desc
	ipv4PacketsNotForwardableDesc                     *prometheus.Desc
	ipv4RedirectsSentDesc                             *prometheus.Desc
	ipv4PacketsSentFromThisHostDesc                   *prometheus.Desc
	ipv4PacketsSentWithFabricatedIPHeaderDesc         *prometheus.Desc
	ipv4OutputPacketsDroppedDueToNoBufsDesc           *prometheus.Desc
	ipv4OutputPacketsDiscardedDueToNoRouteDesc        *prometheus.Desc
	ipv4OutputDatagramsFragmentedDesc                 *prometheus.Desc
	ipv4FragmentsCreatedDesc                          *prometheus.Desc
	ipv4DatagramsThatCanNotBeFragmentedDesc           *prometheus.Desc
	ipv4PacketsWithBadOptionsDesc                     *prometheus.Desc
	ipv4PacketsWithOptionsHandledWithoutErrorDesc     *prometheus.Desc
	ipv4StrictSourceAndRecordRouteOptionsDesc         *prometheus.Desc
	ipv4LooseSourceAndRecordRouteOptionsDesc          *prometheus.Desc
	ipv4RecordRouteOptionsDesc                        *prometheus.Desc
	ipv4TimestampOptionsDesc                          *prometheus.Desc
	ipv4TimestampAndAddressOptionsDesc                *prometheus.Desc
	ipv4TimestampAndPrespecifiedAddressOptionsDesc    *prometheus.Desc
	ipv4OptionPacketsDroppedDueToRateLimitDesc        *prometheus.Desc
	ipv4RouterAlertOptionDesc                         *prometheus.Desc
	ipv4MulticastPacketsDroppedDesc                   *prometheus.Desc
	ipv4PacketsDroppedDesc                            *prometheus.Desc
	ipv4TransitREPacketsDroppedonMGMTInterfaceDesc    *prometheus.Desc
	ipv4PacketsUsedFirstNexthopInECMPUnilistDesc      *prometheus.Desc
	ipv4IncomingTtpoipPacketsReceivedDesc             *prometheus.Desc
	ipv4IncomingTtpoipPacketsDroppedDesc              *prometheus.Desc
	ipv4OutgoingTtpoipPacketsSentDesc                 *prometheus.Desc
	ipv4OutgoingTtpoipPacketsDroppedDesc              *prometheus.Desc
	ipv4IncomingRawIPPacketsDroppedNoSocketBufferDesc *prometheus.Desc
	ipv4IncomingVirtualNodePacketsDeliveredDesc       *prometheus.Desc

	ipv6TotalPacketsReceivedDesc                  *prometheus.Desc
	ipv6PacketsWithSizeSmallerThanMinimumDesc     *prometheus.Desc
	ipv6PacketsWithDatasizeLessThanDataLengthDesc *prometheus.Desc
	ipv6PacketsWithBadOptionsDesc                 *prometheus.Desc
	ipv6PacketsWithIncorrectVersionNumberDesc     *prometheus.Desc
	ipv6FragmentsReceivedDesc                     *prometheus.Desc
	ipv6DuplicateOrOutOfSpaceFragmentsDroppedDesc *prometheus.Desc
	ipv6FragmentsDroppedAfterTimeoutDesc          *prometheus.Desc
	ipv6FragmentsThatExceededLimitDesc            *prometheus.Desc
	ipv6PacketsReassembledOkDesc                  *prometheus.Desc
	ipv6PacketsForThisHostDesc                    *prometheus.Desc
	ipv6PacketsForwardedDesc                      *prometheus.Desc
	ipv6PacketsNotForwardableDesc                 *prometheus.Desc
	ipv6RedirectsSentDesc                         *prometheus.Desc
	ipv6PacketsSentFromThisHostDesc               *prometheus.Desc
	ipv6PacketsSentWithFabricatedIpHeaderDesc     *prometheus.Desc
	ipv6OutputPacketsDroppedDueToNoBufsDesc       *prometheus.Desc
	ipv6OutputPacketsDiscardedDueToNoRouteDesc    *prometheus.Desc
	ipv6OutputDatagramsFragmentedDesc             *prometheus.Desc
	ipv6FragmentsCreatedDesc                      *prometheus.Desc
	ipv6DatagramsThatCanNotBeFragmentedDesc       *prometheus.Desc
	ipv6PacketsThatViolatedScopeRulesDesc         *prometheus.Desc
	ipv6MulticastPacketsWhichWeDoNotJoinDesc      *prometheus.Desc
	ipv6NhTcpDesc                                 *prometheus.Desc
	ipv6NhUdpDesc                                 *prometheus.Desc
	ipv6NhIcmp6Desc                               *prometheus.Desc
	ipv6PacketsWhoseHeadersAreNotContinuousDesc   *prometheus.Desc
	ipv6TunnelingPacketsThatCanNotFindGifDesc     *prometheus.Desc
	ipv6PacketsDiscardedDueToTooMayHeadersDesc    *prometheus.Desc
	ipv6FailuresOfSourceAddressSelectionDesc      *prometheus.Desc
	ipv6HeaderTypeLinkLocalsDesc                  *prometheus.Desc
	ipv6HeaderTypeGlobalsDesc                     *prometheus.Desc
	ipv6ForwardCacheHitDesc                       *prometheus.Desc
	ipv6ForwardCacheMissDesc                      *prometheus.Desc
	ipv6PacketsDestinedToDeadNextHopDesc          *prometheus.Desc
	ipv6OptionPacketsDroppedDueToRateLimitDesc    *prometheus.Desc
	ipv6PacketsDroppedDesc                        *prometheus.Desc
	ipv6PacketsDroppedDueToBadProtocolDesc        *prometheus.Desc
	ipv6TransitRePacketDroppedOnMgmtInterfaceDesc *prometheus.Desc
	ipv6PacketUsedFirstNexthopInEcmpUnilistDesc   *prometheus.Desc

	udpDatagramsReceivedDesc                                 *prometheus.Desc
	udpDatagramsWithIncompleteHeaderDesc                     *prometheus.Desc
	udpDatagramsWithBadDatalengthFieldDesc                   *prometheus.Desc
	udpDatagramsWithBadChecksumDesc                          *prometheus.Desc
	udpDatagramsDroppedDueToNoSocketDesc                     *prometheus.Desc
	udpBroadcastOrMulticastDatagramsDroppedDueToNoSocketDesc *prometheus.Desc
	udpDatagramsDroppedDueToFullSocketBuffersDesc            *prometheus.Desc
	udpDatagramsNotForHashedPcbDesc                          *prometheus.Desc
	udpDatagramsDeliveredDesc                                *prometheus.Desc
	udpDatagramsOutputDesc                                   *prometheus.Desc
)

func init() {
	labelsIPV4 := []string{"target", "protocol"}
	ipv4PacketsReceivedDesc = prometheus.NewDesc(prefix+"ipv4_packets_received", "Number of packets received", labelsIPV4, nil)
	ipv4BadHeaderChecksumsDesc = prometheus.NewDesc(prefix+"ipv4_bad_header_checksums", "Number of packets received with bad header checksums", labelsIPV4, nil)
	ipv4PacketsWithSizeSmallerThanMinimumDesc = prometheus.NewDesc(prefix+"ipv4_packets_with_size_smaller_than_minimum", "Number of packets received with size smaller than minimum", labelsIPV4, nil)
	ipv4PacketsWithDataSizeLessThanDatalengthDesc = prometheus.NewDesc(prefix+"ipv4_packets_with_data_size_less_than_datalength", "Number of packets received with data size less than data length", labelsIPV4, nil)
	ipv4PacketsWithHeaderLengthLessThanDataSizeDesc = prometheus.NewDesc(prefix+"ipv4_packets_with_header_length_less_than_data_size", "Number of packets received with header length less than data size", labelsIPV4, nil)
	ipv4PacketsWithIncorrectVersionNumberDesc = prometheus.NewDesc(prefix+"ipv4_packets_with_incorrect_version_number", "Number of packets received with incorrect version number", labelsIPV4, nil)
	ipv4PacketsDestinedToDeadNextHopDesc = prometheus.NewDesc(prefix+"ipv4_packets_destined_to_dead_next_hop", "Number of packets received destined to dead next hop", labelsIPV4, nil)
	ipv4FragmentsReceivedDesc = prometheus.NewDesc(prefix+"ipv4_fragments_received", "Number of fragments received", labelsIPV4, nil)
	ipv4FragmentsDroppedDueToOutspaceOrDUPDesc = prometheus.NewDesc(prefix+"ipv4_fragments_dropped_due_to_outspace_or_dup", "Number of fragments dropped due to outspace or DUP", labelsIPV4, nil)
	ipv4FragmentsDroppedDueToQueueoverflowDesc = prometheus.NewDesc(prefix+"ipv4_fragments_dropped_due_to_queueoverflow", "Number of fragments dropped due to queue overflow", labelsIPV4, nil)
	ipv4FragmentsDroppedAfterTimeoutDesc = prometheus.NewDesc(prefix+"ipv4_fragments_dropped_after_timeout", "Number of fragments dropped after timeout", labelsIPV4, nil)
	ipv4PacketsReassembledOKDesc = prometheus.NewDesc(prefix+"ipv4_packets_reassembled_ok", "Number of packets reassembled OK", labelsIPV4, nil)
	ipv4PacketsForThisHostDesc = prometheus.NewDesc(prefix+"ipv4_packets_for_this_host", "Number of packets for this host", labelsIPV4, nil)
	ipv4PacketsForUnknownOrUnsupportedProtocolDesc = prometheus.NewDesc(prefix+"ipv4_packets_for_unknown_or_unsupported_protocol", "Number of packets for unknown or unsupported protocol", labelsIPV4, nil)
	ipv4PacketsForwardedDesc = prometheus.NewDesc(prefix+"ipv4_packets_forwarded", "Number of packets forwarded", labelsIPV4, nil)
	ipv4PacketsNotForwardableDesc = prometheus.NewDesc(prefix+"ipv4_packets_not_forwardable", "Number of packets not forwardable", labelsIPV4, nil)
	ipv4RedirectsSentDesc = prometheus.NewDesc(prefix+"ipv4_redirects_sent", "Number of redirects sent", labelsIPV4, nil)
	ipv4PacketsSentFromThisHostDesc = prometheus.NewDesc(prefix+"ipv4_packets_sent_from_this_host", "Number of packets sent from this host", labelsIPV4, nil)
	ipv4PacketsSentWithFabricatedIPHeaderDesc = prometheus.NewDesc(prefix+"ipv4_packets_sent_with_fabricated_ip_header", "Number of packets sent with fabricated IP header", labelsIPV4, nil)
	ipv4OutputPacketsDroppedDueToNoBufsDesc = prometheus.NewDesc(prefix+"ipv4_output_packets_dropped_due_to_no_bufs", "Number of output packets dropped due to no bufs", labelsIPV4, nil)
	ipv4OutputPacketsDiscardedDueToNoRouteDesc = prometheus.NewDesc(prefix+"ipv4_output_packets_discarded_due_to_no_route", "Number of output packets discarded due to no route", labelsIPV4, nil)
	ipv4OutputDatagramsFragmentedDesc = prometheus.NewDesc(prefix+"ipv4_output_datagrams_fragmented", "Number of output datagrams fragmented", labelsIPV4, nil)
	ipv4FragmentsCreatedDesc = prometheus.NewDesc(prefix+"ipv4_fragments_created", "Number of fragments created", labelsIPV4, nil)
	ipv4DatagramsThatCanNotBeFragmentedDesc = prometheus.NewDesc(prefix+"ipv4_datagrams_that_can_not_be_fragmented", "Number of datagrams that can not be fragmented", labelsIPV4, nil)
	ipv4PacketsWithBadOptionsDesc = prometheus.NewDesc(prefix+"ipv4_packets_with_bad_options", "Number of packets with bad options", labelsIPV4, nil)
	ipv4PacketsWithOptionsHandledWithoutErrorDesc = prometheus.NewDesc(prefix+"ipv4_packets_with_options_handled_without_error", "Number of packets with options handled without error", labelsIPV4, nil)
	ipv4StrictSourceAndRecordRouteOptionsDesc = prometheus.NewDesc(prefix+"ipv4_strict_source_and_record_route_options", "Number of packets with strict source and record route options", labelsIPV4, nil)
	ipv4LooseSourceAndRecordRouteOptionsDesc = prometheus.NewDesc(prefix+"ipv4_loose_source_and_record_route_options", "Number of packets with loose source and record route options", labelsIPV4, nil)
	ipv4RecordRouteOptionsDesc = prometheus.NewDesc(prefix+"ipv4_record_route_options", "Number of packets with record route options", labelsIPV4, nil)
	ipv4TimestampOptionsDesc = prometheus.NewDesc(prefix+"ipv4_timestamp_options", "Number of packets with timestamp options", labelsIPV4, nil)
	ipv4TimestampAndAddressOptionsDesc = prometheus.NewDesc(prefix+"ipv4_timestamp_and_address_options", "Number of packets with timestamp and address options", labelsIPV4, nil)
	ipv4TimestampAndPrespecifiedAddressOptionsDesc = prometheus.NewDesc(prefix+"ipv4_timestamp_and_prespecified_address_options", "Number of packets with timestamp and prespecified address options", labelsIPV4, nil)
	ipv4OptionPacketsDroppedDueToRateLimitDesc = prometheus.NewDesc(prefix+"ipv4_option_packets_dropped_due_to_rate_limit", "Number of option packets dropped due to rate limit", labelsIPV4, nil)
	ipv4RouterAlertOptionDesc = prometheus.NewDesc(prefix+"ipv4_router_alert_option", "Number of packets with router alert option", labelsIPV4, nil)
	ipv4MulticastPacketsDroppedDesc = prometheus.NewDesc(prefix+"ipv4_multicast_packets_dropped", "Number of multicast packets dropped", labelsIPV4, nil)
	ipv4PacketsDroppedDesc = prometheus.NewDesc(prefix+"ipv4_packets_dropped", "Number of packets dropped", labelsIPV4, nil)
	ipv4TransitREPacketsDroppedonMGMTInterfaceDesc = prometheus.NewDesc(prefix+"ipv4_transit_re_packets_droppedon_mgt_interface", "Number of transit RE packets dropped on MGMT interface", labelsIPV4, nil)
	ipv4PacketsUsedFirstNexthopInECMPUnilistDesc = prometheus.NewDesc(prefix+"ipv4_packets_used_first_nexthop_in_ecmp_unilist", "Number of packets used first nexthop in ECMP unilist", labelsIPV4, nil)
	ipv4IncomingTtpoipPacketsReceivedDesc = prometheus.NewDesc(prefix+"ipv4_incoming_ttpoip_packets_received", "Number of incoming TTPoIP packets received", labelsIPV4, nil)
	ipv4IncomingTtpoipPacketsDroppedDesc = prometheus.NewDesc(prefix+"ipv4_incoming_ttpoip_packets_dropped", "Number of incoming TTPoIP packets dropped", labelsIPV4, nil)
	ipv4OutgoingTtpoipPacketsSentDesc = prometheus.NewDesc(prefix+"ipv4_outgoing_ttpoip_packets_sent", "Number of outgoing TTPoIP packets sent", labelsIPV4, nil)
	ipv4OutgoingTtpoipPacketsDroppedDesc = prometheus.NewDesc(prefix+"ipv4_outgoing_ttpoip_packets_dropped", "Number of outgoing TTPoIP packets dropped", labelsIPV4, nil)
	ipv4IncomingRawIPPacketsDroppedNoSocketBufferDesc = prometheus.NewDesc(prefix+"ipv4_incoming_raw_ip_packets_dropped_no_socket_buffer", "Number of incoming raw IP packets dropped due to no socket buffer", labelsIPV4, nil)
	ipv4IncomingVirtualNodePacketsDeliveredDesc = prometheus.NewDesc(prefix+"ipv4_incoming_virtual_node_packets_delivered", "Number of incoming virtual node packets delivered", labelsIPV4, nil)

	labelsIPV6 := []string{"target", "protocol"}
	ipv6TotalPacketsReceivedDesc = prometheus.NewDesc(prefix+"ipv6_total_packets_received", "Total number of packets received", labelsIPV6, nil)
	ipv6PacketsWithSizeSmallerThanMinimumDesc = prometheus.NewDesc(prefix+"ipv6_packets_with_size_smaller_than_minimum", "Number of packets received with size smaller than minimum", labelsIPV6, nil)
	ipv6PacketsWithDatasizeLessThanDataLengthDesc = prometheus.NewDesc(prefix+"ipv6_packets_with_datasize_less_than_data_length", "Number of packets received with data length less than data length", labelsIPV6, nil)
	ipv6PacketsWithBadOptionsDesc = prometheus.NewDesc(prefix+"ipv6_packets_with_bad_options", "Number of packets received with bad options", labelsIPV6, nil)
	ipv6PacketsWithIncorrectVersionNumberDesc = prometheus.NewDesc(prefix+"ipv6_packets_with_incorrect_version_number", "Number of packets received with incorrect version number", labelsIPV6, nil)
	ipv6FragmentsReceivedDesc = prometheus.NewDesc(prefix+"ipv6_fragments_received", "Number of fragments received", labelsIPV6, nil)
	ipv6DuplicateOrOutOfSpaceFragmentsDroppedDesc = prometheus.NewDesc(prefix+"ipv6_duplicate_or_out_of_space_fragments_dropped", "Number of duplicate or out of space fragments dropped", labelsIPV6, nil)
	ipv6FragmentsDroppedAfterTimeoutDesc = prometheus.NewDesc(prefix+"ipv6_fragments_dropped_after_timeout", "Number of fragments dropped after timeout", labelsIPV6, nil)
	ipv6FragmentsThatExceededLimitDesc = prometheus.NewDesc(prefix+"ipv6_fragments_that_exceeded_limit", "Number of fragments that exceeded limit", labelsIPV6, nil)
	ipv6PacketsReassembledOkDesc = prometheus.NewDesc(prefix+"ipv6_packets_reassembled_ok", "Number of packets reassembled ok", labelsIPV6, nil)
	ipv6PacketsForThisHostDesc = prometheus.NewDesc(prefix+"ipv6_packets_for_this_host", "Number of packets for this host", labelsIPV6, nil)
	ipv6PacketsForwardedDesc = prometheus.NewDesc(prefix+"ipv6_packets_forwarded", "Number of packets forwarded", labelsIPV6, nil)
	ipv6PacketsNotForwardableDesc = prometheus.NewDesc(prefix+"ipv6_packets_not_forwardable", "Number of packets not forwardable", labelsIPV6, nil)
	ipv6RedirectsSentDesc = prometheus.NewDesc(prefix+"ipv6_redirects_sent", "Number of redirects sent", labelsIPV6, nil)
	ipv6PacketsSentFromThisHostDesc = prometheus.NewDesc(prefix+"ipv6_packets_sent_from_this_host", "Number of packets sent from this host", labelsIPV6, nil)
	ipv6PacketsSentWithFabricatedIpHeaderDesc = prometheus.NewDesc(prefix+"ipv6_packets_sent_with_fabricated_ip_header", "Number of packets sent with fabricated ip header", labelsIPV6, nil)
	ipv6OutputPacketsDroppedDueToNoBufsDesc = prometheus.NewDesc(prefix+"ipv6_output_packets_dropped_due_to_no_bufs", "Number of output packets dropped due to no bufs", labelsIPV6, nil)
	ipv6OutputPacketsDiscardedDueToNoRouteDesc = prometheus.NewDesc(prefix+"ipv6_output_packets_discarded_due_to_no_route", "Number of output packets discarded due to no route", labelsIPV6, nil)
	ipv6OutputDatagramsFragmentedDesc = prometheus.NewDesc(prefix+"ipv6_output_datagrams_fragmented", "Number of output datagrams fragmented", labelsIPV6, nil)
	ipv6FragmentsCreatedDesc = prometheus.NewDesc(prefix+"ipv6_fragments_created", "Number of fragments created", labelsIPV6, nil)
	ipv6DatagramsThatCanNotBeFragmentedDesc = prometheus.NewDesc(prefix+"ipv6_datagrams_that_can_not_be_fragmented", "Number of datagrams that can not be fragmented", labelsIPV6, nil)
	ipv6PacketsThatViolatedScopeRulesDesc = prometheus.NewDesc(prefix+"ipv6_packets_that_violated_scope_rules", "Number of packets that violated scope rules", labelsIPV6, nil)
	ipv6MulticastPacketsWhichWeDoNotJoinDesc = prometheus.NewDesc(prefix+"ipv6_multicast_packets_which_we_do_not_join", "Number of multicast packets which we do not join", labelsIPV6, nil)
	ipv6NhTcpDesc = prometheus.NewDesc(prefix+"ipv6_nh_tcp", "Number of packets with next header tcp", labelsIPV6, nil)
	ipv6NhUdpDesc = prometheus.NewDesc(prefix+"ipv6_nh_udp", "Number of packets with next header udp", labelsIPV6, nil)
	ipv6NhIcmp6Desc = prometheus.NewDesc(prefix+"ipv6_nh_icmp6", "Number of packets with next header icmp6", labelsIPV6, nil)
	ipv6PacketsWhoseHeadersAreNotContinuousDesc = prometheus.NewDesc(prefix+"ipv6_packets_whose_headers_are_not_continuous", "Number of packets whose headers are not continuous", labelsIPV6, nil)
	ipv6TunnelingPacketsThatCanNotFindGifDesc = prometheus.NewDesc(prefix+"ipv6_tunneling_packets_that_can_not_find_gif", "Number of tunneling packets that can not find gif", labelsIPV6, nil)
	ipv6PacketsDiscardedDueToTooMayHeadersDesc = prometheus.NewDesc(prefix+"ipv6_packets_discarded_due_to_too_may_headers", "Number of packets discarded due to too may headers", labelsIPV6, nil)
	ipv6FailuresOfSourceAddressSelectionDesc = prometheus.NewDesc(prefix+"ipv6_failures_of_source_address_selection", "Number of failures of source address selection", labelsIPV6, nil)
	labelsIPV6Header := []string{"target", "protocol", "header_type"}
	ipv6HeaderTypeLinkLocalsDesc = prometheus.NewDesc(prefix+"ipv6_header_type_link_locals", "Number of packets with header type link locals", labelsIPV6Header, nil)
	ipv6HeaderTypeGlobalsDesc = prometheus.NewDesc(prefix+"ipv6_header_type_globals", "Number of packets with header type globals", labelsIPV6Header, nil)
	ipv6ForwardCacheHitDesc = prometheus.NewDesc(prefix+"ipv6_forward_cache_hit", "Number of forward cache hit", labelsIPV6, nil)
	ipv6ForwardCacheMissDesc = prometheus.NewDesc(prefix+"ipv6_forward_cache_miss", "Number of forward cache miss", labelsIPV6, nil)
	ipv6PacketsDestinedToDeadNextHopDesc = prometheus.NewDesc(prefix+"ipv6_packets_destined_to_dead_next_hop", "Number of packets destined to dead next hop", labelsIPV6, nil)
	ipv6OptionPacketsDroppedDueToRateLimitDesc = prometheus.NewDesc(prefix+"ipv6_option_packets_dropped_due_to_rate_limit", "Number of option packets dropped due to rate limit", labelsIPV6, nil)
	ipv6PacketsDroppedDesc = prometheus.NewDesc(prefix+"ipv6_packets_dropped", "Number of packets dropped", labelsIPV6, nil)
	ipv6PacketsDroppedDueToBadProtocolDesc = prometheus.NewDesc(prefix+"ipv6_packets_dropped_due_to_bad_protocol", "Number of packets dropped due to bad protocol", labelsIPV6, nil)
	ipv6TransitRePacketDroppedOnMgmtInterfaceDesc = prometheus.NewDesc(prefix+"ipv6_transit_re_packet_dropped_on_mgmt_interface", "Number of transit re packet dropped on mgmt interface", labelsIPV6, nil)
	ipv6PacketUsedFirstNexthopInEcmpUnilistDesc = prometheus.NewDesc(prefix+"ipv6_packet_used_first_nexthop_in_ecmp_unilist", "Number of packet used first nexthop in ecmp unilist", labelsIPV6, nil)

	labelsUDP := []string{"target", "protocol"}
	udpDatagramsReceivedDesc = prometheus.NewDesc(prefix+"udp_datagrams_received", "Number of UDP datagrams received", labelsUDP, nil)
	udpDatagramsWithIncompleteHeaderDesc = prometheus.NewDesc(prefix+"udp_datagrams_with_incomplete_header", "Number of UDP datagrams with incomplete header", labelsUDP, nil)
	udpDatagramsWithBadDatalengthFieldDesc = prometheus.NewDesc(prefix+"udp_datagrams_with_bad_datalength_field", "Number of UDP datagrams with bad datalength field", labelsUDP, nil)
	udpDatagramsWithBadChecksumDesc = prometheus.NewDesc(prefix+"udp_datagrams_with_bad_checksum", "Number of UDP datagrams with bad checksum", labelsUDP, nil)
	udpDatagramsDroppedDueToNoSocketDesc = prometheus.NewDesc(prefix+"udp_datagrams_dropped_due_to_no_socket", "Number of UDP datagrams dropped due to no socket", labelsUDP, nil)
	udpBroadcastOrMulticastDatagramsDroppedDueToNoSocketDesc = prometheus.NewDesc(prefix+"udp_broadcast_or_multicast_datagrams_dropped_due_to_no_socket", "Number of UDP broadcast or multicast datagrams dropped due to no socket", labelsUDP, nil)
	udpDatagramsDroppedDueToFullSocketBuffersDesc = prometheus.NewDesc(prefix+"udp_datagrams_dropped_due_to_full_socket_buffers", "Number of UDP datagrams dropped due to full socket buffers", labelsUDP, nil)
	udpDatagramsNotForHashedPcbDesc = prometheus.NewDesc(prefix+"udp_datagrams_not_for_hashed_pcb", "Number of UDP datagrams not for hashed pcb", labelsUDP, nil)
	udpDatagramsDeliveredDesc = prometheus.NewDesc(prefix+"udp_datagrams_delivered", "Number of UDP datagrams delivered", labelsUDP, nil)
	udpDatagramsOutputDesc = prometheus.NewDesc(prefix+"udp_datagrams_output", "Number of UDP datagrams output", labelsUDP, nil)
}

type systemstatisticsCollector struct{}

func NewCollector() collector.RPCCollector {
	return &systemstatisticsCollector{}
}

func (c *systemstatisticsCollector) Name() string {
	return "systemstatistics"
}

func (c *systemstatisticsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- ipv4PacketsReceivedDesc
	ch <- ipv4BadHeaderChecksumsDesc
	ch <- ipv4PacketsWithSizeSmallerThanMinimumDesc
	ch <- ipv4PacketsWithDataSizeLessThanDatalengthDesc
	ch <- ipv4PacketsWithHeaderLengthLessThanDataSizeDesc
	ch <- ipv4PacketsWithIncorrectVersionNumberDesc
	ch <- ipv4PacketsDestinedToDeadNextHopDesc
	ch <- ipv4FragmentsReceivedDesc
	ch <- ipv4FragmentsDroppedDueToOutspaceOrDUPDesc
	ch <- ipv4FragmentsDroppedDueToQueueoverflowDesc
	ch <- ipv4FragmentsDroppedAfterTimeoutDesc
	ch <- ipv4PacketsReassembledOKDesc
	ch <- ipv4PacketsForThisHostDesc
	ch <- ipv4PacketsForUnknownOrUnsupportedProtocolDesc
	ch <- ipv4PacketsForwardedDesc
	ch <- ipv4PacketsNotForwardableDesc
	ch <- ipv4RedirectsSentDesc
	ch <- ipv4PacketsSentFromThisHostDesc
	ch <- ipv4PacketsSentWithFabricatedIPHeaderDesc
	ch <- ipv4OutputPacketsDroppedDueToNoBufsDesc
	ch <- ipv4OutputPacketsDiscardedDueToNoRouteDesc
	ch <- ipv4OutputDatagramsFragmentedDesc
	ch <- ipv4FragmentsCreatedDesc
	ch <- ipv4DatagramsThatCanNotBeFragmentedDesc
	ch <- ipv4PacketsWithBadOptionsDesc
	ch <- ipv4PacketsWithOptionsHandledWithoutErrorDesc
	ch <- ipv4StrictSourceAndRecordRouteOptionsDesc
	ch <- ipv4LooseSourceAndRecordRouteOptionsDesc
	ch <- ipv4RecordRouteOptionsDesc
	ch <- ipv4TimestampOptionsDesc
	ch <- ipv4TimestampAndAddressOptionsDesc
	ch <- ipv4TimestampAndPrespecifiedAddressOptionsDesc
	ch <- ipv4OptionPacketsDroppedDueToRateLimitDesc
	ch <- ipv4RouterAlertOptionDesc
	ch <- ipv4MulticastPacketsDroppedDesc
	ch <- ipv4PacketsDroppedDesc
	ch <- ipv4TransitREPacketsDroppedonMGMTInterfaceDesc
	ch <- ipv4PacketsUsedFirstNexthopInECMPUnilistDesc
	ch <- ipv4IncomingTtpoipPacketsReceivedDesc
	ch <- ipv4IncomingTtpoipPacketsDroppedDesc
	ch <- ipv4OutgoingTtpoipPacketsSentDesc
	ch <- ipv4OutgoingTtpoipPacketsDroppedDesc
	ch <- ipv4IncomingRawIPPacketsDroppedNoSocketBufferDesc
	ch <- ipv4IncomingVirtualNodePacketsDeliveredDesc

	ch <- ipv6TotalPacketsReceivedDesc
	ch <- ipv6PacketsWithSizeSmallerThanMinimumDesc
	ch <- ipv6PacketsWithDatasizeLessThanDataLengthDesc
	ch <- ipv6PacketsWithBadOptionsDesc
	ch <- ipv6PacketsWithIncorrectVersionNumberDesc
	ch <- ipv6FragmentsReceivedDesc
	ch <- ipv6DuplicateOrOutOfSpaceFragmentsDroppedDesc
	ch <- ipv6FragmentsDroppedAfterTimeoutDesc
	ch <- ipv6FragmentsThatExceededLimitDesc
	ch <- ipv6PacketsReassembledOkDesc
	ch <- ipv6PacketsForThisHostDesc
	ch <- ipv6PacketsForwardedDesc
	ch <- ipv6PacketsNotForwardableDesc
	ch <- ipv6RedirectsSentDesc
	ch <- ipv6PacketsSentFromThisHostDesc
	ch <- ipv6PacketsSentWithFabricatedIpHeaderDesc
	ch <- ipv6OutputPacketsDroppedDueToNoBufsDesc
	ch <- ipv6OutputPacketsDiscardedDueToNoRouteDesc
	ch <- ipv6OutputDatagramsFragmentedDesc
	ch <- ipv6FragmentsCreatedDesc
	ch <- ipv6DatagramsThatCanNotBeFragmentedDesc
	ch <- ipv6PacketsThatViolatedScopeRulesDesc
	ch <- ipv6MulticastPacketsWhichWeDoNotJoinDesc
	ch <- ipv6NhTcpDesc
	ch <- ipv6NhUdpDesc
	ch <- ipv6NhIcmp6Desc
	ch <- ipv6PacketsWhoseHeadersAreNotContinuousDesc
	ch <- ipv6TunnelingPacketsThatCanNotFindGifDesc
	ch <- ipv6PacketsDiscardedDueToTooMayHeadersDesc
	ch <- ipv6FailuresOfSourceAddressSelectionDesc
	ch <- ipv6HeaderTypeLinkLocalsDesc
	ch <- ipv6HeaderTypeGlobalsDesc
	ch <- ipv6ForwardCacheHitDesc
	ch <- ipv6ForwardCacheMissDesc
	ch <- ipv6PacketsDestinedToDeadNextHopDesc
	ch <- ipv6OptionPacketsDroppedDueToRateLimitDesc
	ch <- ipv6PacketsDroppedDesc
	ch <- ipv6PacketsDroppedDueToBadProtocolDesc
	ch <- ipv6TransitRePacketDroppedOnMgmtInterfaceDesc
	ch <- ipv6PacketUsedFirstNexthopInEcmpUnilistDesc

	ch <- udpDatagramsReceivedDesc
	ch <- udpDatagramsWithIncompleteHeaderDesc
	ch <- udpDatagramsWithBadDatalengthFieldDesc
	ch <- udpDatagramsWithBadChecksumDesc
	ch <- udpDatagramsDroppedDueToNoSocketDesc
	ch <- udpBroadcastOrMulticastDatagramsDroppedDueToNoSocketDesc
	ch <- udpDatagramsDroppedDueToFullSocketBuffersDesc
	ch <- udpDatagramsNotForHashedPcbDesc
	ch <- udpDatagramsDeliveredDesc
	ch <- udpDatagramsOutputDesc
}

func (c *systemstatisticsCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var s SystemStatistics
	err := client.RunCommandAndParse("show system statistics ip", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsIPV4(ch, labelValues, s)
	c.collectSystemStatisticsIPV6(ch, labelValues, s)
	c.collectSystemStatisticsUDP(ch, labelValues, s)
	return nil
}

func (c *systemstatisticsCollector) collectSystemStatisticsIPV4(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	labels := append(labelValues, "ipv4")
	ch <- prometheus.MustNewConstMetric(ipv4PacketsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4BadHeaderChecksumsDesc, prometheus.CounterValue, s.Statistics.Ip.BadHeaderChecksums, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsWithSizeSmallerThanMinimumDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsWithSizeSmallerThanMinimum, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsWithDataSizeLessThanDatalengthDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsWithDataSizeLessThanDatalength, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsWithHeaderLengthLessThanDataSizeDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsWithHeaderLengthLessThanDataSize, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsWithIncorrectVersionNumberDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsWithIncorrectVersionNumber, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsDestinedToDeadNextHopDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsDestinedToDeadNextHop, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4FragmentsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip.FragmentsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4FragmentsDroppedDueToOutspaceOrDUPDesc, prometheus.CounterValue, s.Statistics.Ip.FragmentsDroppedDueToOutofspaceOrDup, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4FragmentsDroppedDueToQueueoverflowDesc, prometheus.CounterValue, s.Statistics.Ip.FragmentsDroppedDueToQueueoverflow, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4FragmentsDroppedAfterTimeoutDesc, prometheus.CounterValue, s.Statistics.Ip.FragmentsDroppedAfterTimeout, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsReassembledOKDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsReassembledOk, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsForThisHostDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsForThisHost, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsForUnknownOrUnsupportedProtocolDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsForUnknownOrUnsupportedProtocol, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsForwardedDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsForwarded, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsNotForwardableDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsNotForwardable, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4RedirectsSentDesc, prometheus.CounterValue, s.Statistics.Ip.RedirectsSent, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsSentFromThisHostDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsSentFromThisHost, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsSentWithFabricatedIPHeaderDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsSentWithFabricatedIpHeader, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4OutputPacketsDroppedDueToNoBufsDesc, prometheus.CounterValue, s.Statistics.Ip.OutputPacketsDroppedDueToNoBufs, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4OutputPacketsDiscardedDueToNoRouteDesc, prometheus.CounterValue, s.Statistics.Ip.OutputPacketsDiscardedDueToNoRoute, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4OutputDatagramsFragmentedDesc, prometheus.CounterValue, s.Statistics.Ip.OutputDatagramsFragmented, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4FragmentsCreatedDesc, prometheus.CounterValue, s.Statistics.Ip.FragmentsCreated, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4DatagramsThatCanNotBeFragmentedDesc, prometheus.CounterValue, s.Statistics.Ip.DatagramsThatCanNotBeFragmented, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsWithBadOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsWithBadOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsWithOptionsHandledWithoutErrorDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsWithOptionsHandledWithoutError, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4StrictSourceAndRecordRouteOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.StrictSourceAndRecordRouteOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4LooseSourceAndRecordRouteOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.LooseSourceAndRecordRouteOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4RecordRouteOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.RecordRouteOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4TimestampOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.TimestampOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4TimestampAndAddressOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.TimestampAndAddressOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4TimestampAndPrespecifiedAddressOptionsDesc, prometheus.CounterValue, s.Statistics.Ip.TimestampAndPrespecifiedAddressOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4OptionPacketsDroppedDueToRateLimitDesc, prometheus.CounterValue, s.Statistics.Ip.OptionPacketsDroppedDueToRateLimit, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4RouterAlertOptionDesc, prometheus.CounterValue, s.Statistics.Ip.RouterAlertOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4MulticastPacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip.MulticastPacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4TransitREPacketsDroppedonMGMTInterfaceDesc, prometheus.CounterValue, s.Statistics.Ip.TransitRePacketsDroppedOnMgmtInterface, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4PacketsUsedFirstNexthopInECMPUnilistDesc, prometheus.CounterValue, s.Statistics.Ip.PacketsUsedFirstNexthopInEcmpUnilist, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4IncomingTtpoipPacketsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip.IncomingTtpoipPacketsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4IncomingTtpoipPacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip.IncomingTtpoipPacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4OutgoingTtpoipPacketsSentDesc, prometheus.CounterValue, s.Statistics.Ip.OutgoingTtpoipPacketsSent, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4OutgoingTtpoipPacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip.OutgoingTtpoipPacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4IncomingRawIPPacketsDroppedNoSocketBufferDesc, prometheus.CounterValue, s.Statistics.Ip.IncomingRawipPacketsDroppedNoSocketBuffer, labels...)
	ch <- prometheus.MustNewConstMetric(ipv4IncomingVirtualNodePacketsDeliveredDesc, prometheus.CounterValue, s.Statistics.Ip.IncomingVirtualNodePacketsDelivered, labels...)
}

func (c *systemstatisticsCollector) collectSystemStatisticsIPV6(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	labels := append(labelValues, "ipv6")
	ch <- prometheus.MustNewConstMetric(ipv6TotalPacketsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip6.TotalPacketsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsWithSizeSmallerThanMinimumDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsWithSizeSmallerThanMinimum, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsWithDatasizeLessThanDataLengthDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsWithDatasizeLessThanDataLength, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsWithBadOptionsDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsWithBadOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsWithIncorrectVersionNumberDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsWithIncorrectVersionNumber, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6FragmentsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6FragmentsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6DuplicateOrOutOfSpaceFragmentsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip6.DuplicateOrOutOfSpaceFragmentsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6FragmentsDroppedAfterTimeoutDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6FragmentsDroppedAfterTimeout, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6FragmentsThatExceededLimitDesc, prometheus.CounterValue, s.Statistics.Ip6.FragmentsThatExceededLimit, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsReassembledOkDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsReassembledOk, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsForThisHostDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsForThisHost, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsForwardedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsForwarded, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsNotForwardableDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsNotForwardable, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6RedirectsSentDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6RedirectsSent, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsSentFromThisHostDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsSentFromThisHost, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsSentWithFabricatedIpHeaderDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsSentWithFabricatedIpHeader, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6OutputPacketsDroppedDueToNoBufsDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OutputPacketsDroppedDueToNoBufs, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6OutputPacketsDiscardedDueToNoRouteDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OutputPacketsDiscardedDueToNoRoute, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6OutputDatagramsFragmentedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OutputDatagramsFragmented, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6FragmentsCreatedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6FragmentsCreated, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6DatagramsThatCanNotBeFragmentedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6DatagramsThatCanNotBeFragmented, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsThatViolatedScopeRulesDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsThatViolatedScopeRules, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6MulticastPacketsWhichWeDoNotJoinDesc, prometheus.CounterValue, s.Statistics.Ip6.MulticastPacketsWhichWeDoNotJoin, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6NhTcpDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6nhTcp, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6NhUdpDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6nhUdp, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6NhIcmp6Desc, prometheus.CounterValue, s.Statistics.Ip6.Ip6nhIcmp6, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsWhoseHeadersAreNotContinuousDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsWhoseHeadersAreNotContinuous, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6TunnelingPacketsThatCanNotFindGifDesc, prometheus.CounterValue, s.Statistics.Ip6.TunnelingPacketsThatCanNotFindGif, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsDiscardedDueToTooMayHeadersDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsDiscardedDueToTooMayHeaders, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6FailuresOfSourceAddressSelectionDesc, prometheus.CounterValue, s.Statistics.Ip6.FailuresOfSourceAddressSelection, labels...)
	for _, header := range s.Statistics.Ip6.HeaderType {
		labels := append(labelValues, "ipv6", header.HeaderForSourceAddressSelection)
		ch <- prometheus.MustNewConstMetric(ipv6HeaderTypeLinkLocalsDesc, prometheus.CounterValue, header.LinkLocals, labels...)
		ch <- prometheus.MustNewConstMetric(ipv6HeaderTypeGlobalsDesc, prometheus.CounterValue, header.Globals, labels...)
	}
	ch <- prometheus.MustNewConstMetric(ipv6ForwardCacheHitDesc, prometheus.CounterValue, s.Statistics.Ip6.ForwardCacheHit, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6ForwardCacheMissDesc, prometheus.CounterValue, s.Statistics.Ip6.ForwardCacheMiss, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsDestinedToDeadNextHopDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsDestinedToDeadNextHop, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6OptionPacketsDroppedDueToRateLimitDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OptionPacketsDroppedDueToRateLimit, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketsDroppedDueToBadProtocolDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsDroppedDueToBadProtocol, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6TransitRePacketDroppedOnMgmtInterfaceDesc, prometheus.CounterValue, s.Statistics.Ip6.TransitRePacketDroppedOnMgmtInterface, labels...)
	ch <- prometheus.MustNewConstMetric(ipv6PacketUsedFirstNexthopInEcmpUnilistDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketUsedFirstNexthopInEcmpUnilist, labels...)

}

func (c *systemstatisticsCollector) collectSystemStatisticsUDP(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	l := append(labelValues, "udp")
	ch <- prometheus.MustNewConstMetric(udpDatagramsReceivedDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsReceived, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsWithIncompleteHeaderDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsWithIncompleteHeader, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsWithBadDatalengthFieldDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsWithBadDatalengthField, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsWithBadChecksumDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsWithBadChecksum, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsDroppedDueToNoSocketDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsDroppedDueToNoSocket, l...)
	ch <- prometheus.MustNewConstMetric(udpBroadcastOrMulticastDatagramsDroppedDueToNoSocketDesc, prometheus.CounterValue, s.Statistics.Udp.BroadcastOrMulticastDatagramsDroppedDueToNoSocket, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsDroppedDueToFullSocketBuffersDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsDroppedDueToFullSocketBuffers, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsNotForHashedPcbDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsNotForHashedPcb, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsDeliveredDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsDelivered, l...)
	ch <- prometheus.MustNewConstMetric(udpDatagramsOutputDesc, prometheus.CounterValue, s.Statistics.Udp.DatagramsOutput, l...)
}