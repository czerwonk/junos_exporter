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

	tcpPacketsSent                                      *prometheus.Desc
	tcpSentDataPackets                                  *prometheus.Desc
	tcpDataPacketsBytes                                 *prometheus.Desc
	tcpSentDataPacketsRetransmitted                     *prometheus.Desc
	tcpRetransmittedBytes                               *prometheus.Desc
	tcpSentDataUnnecessaryRetransmitted                 *prometheus.Desc
	tcpSentResendsByMtuDiscovery                        *prometheus.Desc
	tcpSentAckOnlyPackets                               *prometheus.Desc
	tcpSentPacketsDelayed                               *prometheus.Desc
	tcpSentUrgOnlyPackets                               *prometheus.Desc
	tcpSentWindowProbePackets                           *prometheus.Desc
	tcpSentWindowUpdatePackets                          *prometheus.Desc
	tcpSentControlPackets                               *prometheus.Desc
	tcpPacketsReceived                                  *prometheus.Desc
	tcpReceivedAcks                                     *prometheus.Desc
	tcpAcksBytes                                        *prometheus.Desc
	tcpReceivedDuplicateAcks                            *prometheus.Desc
	tcpReceivedAcksForUnsentData                        *prometheus.Desc
	tcpPacketsReceivedInSequence                        *prometheus.Desc
	tcpInSequenceBytes                                  *prometheus.Desc
	tcpReceivedCompletelyDuplicatePacket                *prometheus.Desc
	tcpDuplicateInBytes                                 *prometheus.Desc
	tcpReceivedOldDuplicatePackets                      *prometheus.Desc
	tcpReceivedPacketsWithSomeDupliacteData             *prometheus.Desc
	tcpSomeDuplicateInBytes                             *prometheus.Desc
	tcpReceivedOutOfOrderPackets                        *prometheus.Desc
	tcpOutOfOrderInBytes                                *prometheus.Desc
	tcpReceivedPacketsOfDataAfterWindow                 *prometheus.Desc
	tcpBytes                                            *prometheus.Desc
	tcpReceivedWindowProbes                             *prometheus.Desc
	tcpReceivedWindowUpdatePackets                      *prometheus.Desc
	tcpPacketsReceivedAfterClose                        *prometheus.Desc
	tcpReceivedDiscardedForBadChecksum                  *prometheus.Desc
	tcpReceivedDiscardedForBadHeaderOffset              *prometheus.Desc
	tcpReceivedDiscardedBecausePacketTooShort           *prometheus.Desc
	tcpConnectionRequests                               *prometheus.Desc
	tcpConnectionAccepts                                *prometheus.Desc
	tcpBadConnectionAttempts                            *prometheus.Desc
	tcpListenQueueOverflows                             *prometheus.Desc
	tcpBadRstWindow                                     *prometheus.Desc
	tcpConnectionsEstablished                           *prometheus.Desc
	tcpConnectionsClosed                                *prometheus.Desc
	tcpDrops                                            *prometheus.Desc
	tcpConnectionsUpdatedRttOnClose                     *prometheus.Desc
	tcpConnectionsUpdatedVarianceOnClose                *prometheus.Desc
	tcpConnectionsUpdatedSsthreshOnClose                *prometheus.Desc
	tcpEmbryonicConnectionsDropped                      *prometheus.Desc
	tcpSegmentsUpdatedRtt                               *prometheus.Desc
	tcpAttempts                                         *prometheus.Desc
	tcpRetransmitTimeouts                               *prometheus.Desc
	tcpConnectionsDroppedByRetransmitTimeout            *prometheus.Desc
	tcpPersistTimeouts                                  *prometheus.Desc
	tcpConnectionsDroppedByPersistTimeout               *prometheus.Desc
	tcpKeepaliveTimeouts                                *prometheus.Desc
	tcpKeepaliveProbesSent                              *prometheus.Desc
	tcpKeepaliveConnectionsDropped                      *prometheus.Desc
	tcpAckHeaderPredictions                             *prometheus.Desc
	tcpDataPacketHeaderPredictions                      *prometheus.Desc
	tcpSyncacheEntriesAdded                             *prometheus.Desc
	tcpRetransmitted                                    *prometheus.Desc
	tcpDupsyn                                           *prometheus.Desc
	tcpDropped                                          *prometheus.Desc
	tcpCompleted                                        *prometheus.Desc
	tcpBucketOverflow                                   *prometheus.Desc
	tcpCacheOverflow                                    *prometheus.Desc
	tcpReset                                            *prometheus.Desc
	tcpStale                                            *prometheus.Desc
	tcpAborted                                          *prometheus.Desc
	tcpBadack                                           *prometheus.Desc
	tcpUnreach                                          *prometheus.Desc
	tcpZoneFailures                                     *prometheus.Desc
	tcpCookiesSent                                      *prometheus.Desc
	tcpCookiesReceived                                  *prometheus.Desc
	tcpSackRecoveryEpisodes                             *prometheus.Desc
	tcpSegmentRetransmits                               *prometheus.Desc
	tcpByteRetransmits                                  *prometheus.Desc
	tcpSackOptionsReceived                              *prometheus.Desc
	tcpSackOptionsSent                                  *prometheus.Desc
	tcpSackScoreboardOverflow                           *prometheus.Desc
	tcpAcksSentInResponseButNotExactRsts                *prometheus.Desc
	tcpAcksSentInResponseToSynsOnEstablishedConnections *prometheus.Desc
	tcpRcvPacketsDroppedDueToBadAddress                 *prometheus.Desc
	tcpOutOfSequenceSegmentDrops                        *prometheus.Desc
	tcpRstPackets                                       *prometheus.Desc
	tcpIcmpPacketsIgnored                               *prometheus.Desc
	tcpSendPacketsDropped                               *prometheus.Desc
	tcpRcvPacketsDropped                                *prometheus.Desc
	tcpOutgoingSegmentsDropped                          *prometheus.Desc
	tcpReceivedSynfinDropped                            *prometheus.Desc
	tcpReceivedIpsecDropped                             *prometheus.Desc
	tcpReceivedMacDropped                               *prometheus.Desc
	tcpReceivedMinttlExceeded                           *prometheus.Desc
	tcpListenstateBadflagsDropped                       *prometheus.Desc
	tcpFinwaitstateBadflagsDropped                      *prometheus.Desc
	tcpReceivedDosAttack                                *prometheus.Desc
	tcpReceivedBadSynack                                *prometheus.Desc
	tcpSyncacheZoneFull                                 *prometheus.Desc
	tcpReceivedRstFirewallfilter                        *prometheus.Desc
	tcpReceivedNoackTimewait                            *prometheus.Desc
	tcpReceivedNoTimewaitState                          *prometheus.Desc
	tcpReceivedRstTimewaitState                         *prometheus.Desc
	tcpReceivedTimewaitDrops                            *prometheus.Desc
	tcpReceivedBadaddrTimewaitState                     *prometheus.Desc
	tcpReceivedAckoffInSynSentrcvd                      *prometheus.Desc
	tcpReceivedBadaddrFirewall                          *prometheus.Desc
	tcpReceivedNosynSynSent                             *prometheus.Desc
	tcpReceivedBadrstSynSent                            *prometheus.Desc
	tcpReceivedBadrstListenState                        *prometheus.Desc
	tcpOptionMaxsegmentLength                           *prometheus.Desc
	tcpOptionWindowLength                               *prometheus.Desc
	tcpOptionTimestampLength                            *prometheus.Desc
	tcpOptionMd5Length                                  *prometheus.Desc
	tcpOptionAuthLength                                 *prometheus.Desc
	tcpOptionSackpermittedLength                        *prometheus.Desc
	tcpOptionSackLength                                 *prometheus.Desc
	tcpOptionAuthoptionLength                           *prometheus.Desc
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

	labelsTCP := []string{"target", "protocol"}
	tcpPacketsSent = prometheus.NewDesc(prefix+"tcp_packets_sent", "Number of TCP packets sent", labelsTCP, nil)
	tcpSentDataPackets = prometheus.NewDesc(prefix+"tcp_sent_data_packets", "Number of TCP sent data packets", labelsTCP, nil)
	tcpDataPacketsBytes = prometheus.NewDesc(prefix+"tcp_data_packets_bytes", "Number of TCP data packets bytes", labelsTCP, nil)
	tcpSentDataPacketsRetransmitted = prometheus.NewDesc(prefix+"tcp_sent_data_packets_retransmitted", "Number of TCP sent data packets retransmitted", labelsTCP, nil)
	tcpRetransmittedBytes = prometheus.NewDesc(prefix+"tcp_retransmitted_bytes", "Number of TCP retransmitted bytes", labelsTCP, nil)
	tcpSentDataUnnecessaryRetransmitted = prometheus.NewDesc(prefix+"tcp_sent_data_unnecessary_retransmitted", "Number of tcp data unnecessary retransmitted packets", labelsTCP, nil)
	tcpSentResendsByMtuDiscovery = prometheus.NewDesc(prefix+"tcp_sent_resends_by_mtu_discovery", "Number of tcp sent resends by mtu discovery", labelsTCP, nil)
	tcpSentAckOnlyPackets = prometheus.NewDesc(prefix+"tcp_sent_ack_only_packets", "Number of tcp sent ack only packets", labelsTCP, nil)
	tcpSentPacketsDelayed = prometheus.NewDesc(prefix+"tcp_sent_packets_delayed", "Number of tcp sent packets delayed", labelsTCP, nil)
	tcpSentUrgOnlyPackets = prometheus.NewDesc(prefix+"tcp_sent_urg_only_packets", "Number of tcp sent urg only packets", labelsTCP, nil)
	tcpSentWindowProbePackets = prometheus.NewDesc(prefix+"tcp_sent_window_probe_packets", "Number of tcp sent window probe packets", labelsTCP, nil)
	tcpSentWindowUpdatePackets = prometheus.NewDesc(prefix+"tcp_sent_window_update_packets", "Number of tcp sent window update packets", labelsTCP, nil)
	tcpSentControlPackets = prometheus.NewDesc(prefix+"tcp_sent_control_packets", "Number of tcp sent control packets", labelsTCP, nil)
	tcpPacketsReceived = prometheus.NewDesc(prefix+"tcp_packets_received", "Number of TCP packets received", labelsTCP, nil)
	tcpReceivedAcks = prometheus.NewDesc(prefix+"tcp_received_acks", "Number of TCP received acks", labelsTCP, nil)
	tcpAcksBytes = prometheus.NewDesc(prefix+"tcp_acks_bytes", "Number of TCP acks bytes", labelsTCP, nil)
	tcpReceivedDuplicateAcks = prometheus.NewDesc(prefix+"tcp_received_duplicate_acks", "Number of TCP received duplicate acks", labelsTCP, nil)
	tcpReceivedAcksForUnsentData = prometheus.NewDesc(prefix+"tcp_received_acks_for_unsent_data", "Number of TCP received acks for unsent data", labelsTCP, nil)
	tcpPacketsReceivedInSequence = prometheus.NewDesc(prefix+"tcp_packets_received_in_sequence", "Number of TCP packets received in sequence", labelsTCP, nil)
	tcpInSequenceBytes = prometheus.NewDesc(prefix+"tcp_in_sequence_bytes", "Number of TCP in sequence bytes", labelsTCP, nil)
	tcpReceivedCompletelyDuplicatePacket = prometheus.NewDesc(prefix+"tcp_received_completely_duplicate_packet", "Number of TCP received completely duplicate packet", labelsTCP, nil)
	tcpDuplicateInBytes = prometheus.NewDesc(prefix+"tcp_duplicate_in_bytes", "Number of TCP duplicate in bytes", labelsTCP, nil)
	tcpReceivedOldDuplicatePackets = prometheus.NewDesc(prefix+"tcp_received_old_duplicate_packets", "Number of TCP received old duplicate packets", labelsTCP, nil)
	tcpReceivedPacketsWithSomeDupliacteData = prometheus.NewDesc(prefix+"tcp_received_packet_with_some_duplicate_data", "Number of TCP received packet with some duplicate data", labelsTCP, nil)
	tcpSomeDuplicateInBytes = prometheus.NewDesc(prefix+"tcp_some_duplicate_in_bytes", "Number of TCP some duplicate in bytes", labelsTCP, nil)
	tcpReceivedOutOfOrderPackets = prometheus.NewDesc(prefix+"tcp_received_out_of_order_packets", "Number of TCP received out of order packets", labelsTCP, nil)
	tcpOutOfOrderInBytes = prometheus.NewDesc(prefix+"tcp_out_of_order_in_bytes", "Number of TCP out of order in bytes", labelsTCP, nil)
	tcpReceivedPacketsOfDataAfterWindow = prometheus.NewDesc(prefix+"tcp_received_packets_of_data_after_window", "Number of TCP received packets of data after window", labelsTCP, nil)
	tcpBytes = prometheus.NewDesc(prefix+"tcp_bytes", "Number of TCP bytes", labelsTCP, nil)
	tcpReceivedWindowProbes = prometheus.NewDesc(prefix+"tcp_received_window_probes", "Number of TCP received window probes", labelsTCP, nil)
	tcpReceivedWindowUpdatePackets = prometheus.NewDesc(prefix+"tcp_received_window_update_packets", "Number of TCP received window update packets", labelsTCP, nil)
	tcpPacketsReceivedAfterClose = prometheus.NewDesc(prefix+"tcp_packets_received_after_close", "Number of TCP packets received after close", labelsTCP, nil)
	tcpReceivedDiscardedForBadChecksum = prometheus.NewDesc(prefix+"tcp_received_discarded_for_bad_checksum", "Number of TCP received discarded for bad checksum", labelsTCP, nil)
	tcpReceivedDiscardedForBadHeaderOffset = prometheus.NewDesc(prefix+"tcp_received_discarded_for_bad_header_offset", "Number of TCP received discarded for bad header offset", labelsTCP, nil)
	tcpReceivedDiscardedBecausePacketTooShort = prometheus.NewDesc(prefix+"tcp_received_discarded_because_packet_too_short", "Number of TCP received discarded because packet too short", labelsTCP, nil)
	tcpConnectionRequests = prometheus.NewDesc(prefix+"tcp_connection_requests", "Number of TCP connection requests", labelsTCP, nil)
	tcpConnectionAccepts = prometheus.NewDesc(prefix+"tcp_connection_accepts", "Number of TCP connection accepts", labelsTCP, nil)
	tcpBadConnectionAttempts = prometheus.NewDesc(prefix+"tcp_bad_connection_attempts", "Number of TCP bad connection attempts", labelsTCP, nil)
	tcpListenQueueOverflows = prometheus.NewDesc(prefix+"tcp_listen_queue_overflows", "Number of TCP listen queue overflows", labelsTCP, nil)
	tcpBadRstWindow = prometheus.NewDesc(prefix+"tcp_bad_rst_window", "Number of TCP bad rst window", labelsTCP, nil)
	tcpConnectionsEstablished = prometheus.NewDesc(prefix+"tcp_connections_established", "Number of TCP connections established", labelsTCP, nil)
	tcpConnectionsClosed = prometheus.NewDesc(prefix+"tcp_connections_closed", "Number of TCP connections closed", labelsTCP, nil)
	tcpDrops = prometheus.NewDesc(prefix+"tcp_drops", "Number of TCP drops", labelsTCP, nil)
	tcpConnectionsUpdatedRttOnClose = prometheus.NewDesc(prefix+"tcp_connections_updated_rtt_on_close", "Number of TCP connections updated rtt on close", labelsTCP, nil)
	tcpConnectionsUpdatedVarianceOnClose = prometheus.NewDesc(prefix+"tcp_connections_updated_variance_on_close", "Number of TCP connections updated variance on close", labelsTCP, nil)
	tcpConnectionsUpdatedSsthreshOnClose = prometheus.NewDesc(prefix+"tcp_connections_updated_ssthresh_on_close", "Number of TCP connections updated ssthresh on close", labelsTCP, nil)
	tcpEmbryonicConnectionsDropped = prometheus.NewDesc(prefix+"tcp_embryonic_connections_dropped", "Number of TCP embryonic connections dropped", labelsTCP, nil)
	tcpSegmentsUpdatedRtt = prometheus.NewDesc(prefix+"tcp_segments_updated_rtt", "Number of TCP segments updated rtt", labelsTCP, nil)
	tcpAttempts = prometheus.NewDesc(prefix+"tcp_attempts", "Number of TCP attempts", labelsTCP, nil)
	tcpRetransmitTimeouts = prometheus.NewDesc(prefix+"tcp_retransmit_timeouts", "Number of TCP retransmit timeouts", labelsTCP, nil)
	tcpConnectionsDroppedByRetransmitTimeout = prometheus.NewDesc(prefix+"tcp_connections_dropped_by_retransmit_timeout", "Number of TCP connections dropped by retransmit timeout", labelsTCP, nil)
	tcpPersistTimeouts = prometheus.NewDesc(prefix+"tcp_persist_timeouts", "Number of TCP persist timeouts", labelsTCP, nil)
	tcpConnectionsDroppedByPersistTimeout = prometheus.NewDesc(prefix+"tcp_connections_dropped_by_persist_timeout", "Number of TCP connections dropped by persist timeout", labelsTCP, nil)
	tcpKeepaliveTimeouts = prometheus.NewDesc(prefix+"tcp_keepalive_timeouts", "Number of TCP keepalive timeouts", labelsTCP, nil)
	tcpKeepaliveProbesSent = prometheus.NewDesc(prefix+"tcp_keepalive_probes_sent", "Number of TCP keepalive probes sent", labelsTCP, nil)
	tcpKeepaliveConnectionsDropped = prometheus.NewDesc(prefix+"tcp_keepalive_connections_dropped", "Number of TCP keepalive connections dropped", labelsTCP, nil)
	tcpAckHeaderPredictions = prometheus.NewDesc(prefix+"tcp_ack_header_predictions", "Number of TCP ack header predictions", labelsTCP, nil)
	tcpDataPacketHeaderPredictions = prometheus.NewDesc(prefix+"tcp_data_packet_header_predictions", "Number of TCP data packet header predictions", labelsTCP, nil)
	tcpSyncacheEntriesAdded = prometheus.NewDesc(prefix+"tcp_syncache_entries_added", "Number of TCP syncache entries added", labelsTCP, nil)
	tcpRetransmitted = prometheus.NewDesc(prefix+"tcp_retransmitted", "Number of TCP retransmitted", labelsTCP, nil)
	tcpDupsyn = prometheus.NewDesc(prefix+"tcp_dupsyn", "Number of TCP dupsyn", labelsTCP, nil)
	tcpDropped = prometheus.NewDesc(prefix+"tcp_dropped", "Number of TCP dropped", labelsTCP, nil)
	tcpCompleted = prometheus.NewDesc(prefix+"tcp_completed", "Number of TCP completed", labelsTCP, nil)
	tcpBucketOverflow = prometheus.NewDesc(prefix+"tcp_bucket_overflow", "Number of TCP bucket overflow", labelsTCP, nil)
	tcpCacheOverflow = prometheus.NewDesc(prefix+"tcp_cache_overflow", "Number of TCP cache overflow", labelsTCP, nil)
	tcpReset = prometheus.NewDesc(prefix+"tcp_reset", "Number of TCP reset", labelsTCP, nil)
	tcpStale = prometheus.NewDesc(prefix+"tcp_stale", "Number of TCP stale", labelsTCP, nil)
	tcpAborted = prometheus.NewDesc(prefix+"tcp_aborted", "Number of TCP aborted", labelsTCP, nil)
	tcpBadack = prometheus.NewDesc(prefix+"tcp_badack", "Number of TCP badack", labelsTCP, nil)
	tcpUnreach = prometheus.NewDesc(prefix+"tcp_unreach", "Number of TCP unreach", labelsTCP, nil)
	tcpZoneFailures = prometheus.NewDesc(prefix+"tcp_zone_failures", "Number of TCP zone failures", labelsTCP, nil)
	tcpCookiesSent = prometheus.NewDesc(prefix+"tcp_cookies_sent", "Number of TCP cookies sent", labelsTCP, nil)
	tcpCookiesReceived = prometheus.NewDesc(prefix+"tcp_cookies_received", "Number of TCP cookies received", labelsTCP, nil)
	tcpSackRecoveryEpisodes = prometheus.NewDesc(prefix+"tcp_sack_recovery_episodes", "Number of TCP sack recovery episodes", labelsTCP, nil)
	tcpSegmentRetransmits = prometheus.NewDesc(prefix+"tcp_segment_retransmits", "Number of TCP segment retransmits", labelsTCP, nil)
	tcpByteRetransmits = prometheus.NewDesc(prefix+"tcp_byte_retransmits", "Number of TCP byte retransmits", labelsTCP, nil)
	tcpSackOptionsReceived = prometheus.NewDesc(prefix+"tcp_sack_options_received", "Number of TCP sack options received", labelsTCP, nil)
	tcpSackOptionsSent = prometheus.NewDesc(prefix+"tcp_sack_options_sent", "Number of TCP sack options sent", labelsTCP, nil)
	tcpSackScoreboardOverflow = prometheus.NewDesc(prefix+"tcp_sack_scoreboard_overflow", "Number of TCP sack scoreboard overflow", labelsTCP, nil)
	tcpAcksSentInResponseButNotExactRsts = prometheus.NewDesc(prefix+"tcp_acks_sent_in_response_but_not_exact_rsts", "Number of TCP acks sent in response but not exact rsts", labelsTCP, nil)
	tcpAcksSentInResponseToSynsOnEstablishedConnections = prometheus.NewDesc(prefix+"tcp_ack_sent_in_response_to_syns_on_established_connections", "Number of TCP acks sent in response to syns on established connections", labelsTCP, nil)
	tcpRcvPacketsDroppedDueToBadAddress = prometheus.NewDesc(prefix+"tcp_rcv_packets_dropped_due_to_bad_address", "Number of TCP rcv packets dropped due to bad address", labelsTCP, nil)
	tcpOutOfSequenceSegmentDrops = prometheus.NewDesc(prefix+"tcp_out_of_sequence_segment_drops", "Number of TCP out of sequence segment drops", labelsTCP, nil)
	tcpRstPackets = prometheus.NewDesc(prefix+"tcp_rst_packets", "Number of TCP rst packets", labelsTCP, nil)
	tcpIcmpPacketsIgnored = prometheus.NewDesc(prefix+"tcp_icmp_packets_ignored", "Number of TCP icmp packets ignored", labelsTCP, nil)
	tcpSendPacketsDropped = prometheus.NewDesc(prefix+"tcp_send_packets_dropped", "Number of TCP send packets dropped", labelsTCP, nil)
	tcpRcvPacketsDropped = prometheus.NewDesc(prefix+"tcp_rcv_packets_dropped", "Number of TCP rcv packets dropped", labelsTCP, nil)
	tcpOutgoingSegmentsDropped = prometheus.NewDesc(prefix+"tcp_outgoing_segments_dropped", "Number of TCP outgoing segments dropped", labelsTCP, nil)
	tcpReceivedSynfinDropped = prometheus.NewDesc(prefix+"tcp_received_synfin_dropped", "Number of TCP received synfin dropped", labelsTCP, nil)
	tcpReceivedIpsecDropped = prometheus.NewDesc(prefix+"tcp_received_ipsec_dropped", "Number of TCP received ipsec dropped", labelsTCP, nil)
	tcpReceivedMacDropped = prometheus.NewDesc(prefix+"tcp_received_mac_dropped", "Number of TCP received mac dropped", labelsTCP, nil)
	tcpReceivedMinttlExceeded = prometheus.NewDesc(prefix+"tcp_received_minttl_exceeded", "Number of TCP received minttl exceeded", labelsTCP, nil)
	tcpListenstateBadflagsDropped = prometheus.NewDesc(prefix+"tcp_listenstate_badflags_dropped", "Number of TCP listenstate badflags dropped", labelsTCP, nil)
	tcpFinwaitstateBadflagsDropped = prometheus.NewDesc(prefix+"tcp_finwaitstate_badflags_dropped", "Number of TCP finwaitstate badflags dropped", labelsTCP, nil)
	tcpReceivedDosAttack = prometheus.NewDesc(prefix+"tcp_received_dos_attack", "Number of tcp received dos attack", labelsTCP, nil)
	tcpReceivedBadSynack = prometheus.NewDesc(prefix+"tcp_received_bad_synack", "Number of tcp received bad synack", labelsTCP, nil)
	tcpSyncacheZoneFull = prometheus.NewDesc(prefix+"tcp_syncache_zone_full", "Number of TCP syncache zone full", labelsTCP, nil)
	tcpReceivedRstFirewallfilter = prometheus.NewDesc(prefix+"tcp_received_rst_firewallfilter", "Number of TCP received rst firewallfilter", labelsTCP, nil)
	tcpReceivedNoackTimewait = prometheus.NewDesc(prefix+"tcp_received_noack_timewait", "Number of TCP received noack timewait", labelsTCP, nil)
	tcpReceivedNoTimewaitState = prometheus.NewDesc(prefix+"tcp_received_no_timewait_state", "Number of TCP received no timewait state", labelsTCP, nil)
	tcpReceivedRstTimewaitState = prometheus.NewDesc(prefix+"tcp_received_rst_timewait_state", "Number of TCP received rst timewait state", labelsTCP, nil)
	tcpReceivedTimewaitDrops = prometheus.NewDesc(prefix+"tcp_received_timewait_drops", "Number of TCP received timewait drops", labelsTCP, nil)
	tcpReceivedBadaddrTimewaitState = prometheus.NewDesc(prefix+"tcp_received_badaddr_timewait_state", "Number of TCP received badaddr timewait state", labelsTCP, nil)
	tcpReceivedAckoffInSynSentrcvd = prometheus.NewDesc(prefix+"tcp_received_ackoff_insyn_sentrcvd", "Number of TCP received ackoff in syn sentrcvd", labelsTCP, nil)
	tcpReceivedBadaddrFirewall = prometheus.NewDesc(prefix+"tcp_received_badaddr_firewall", "Number of TCP received badaddr firewall", labelsTCP, nil)
	tcpReceivedNosynSynSent = prometheus.NewDesc(prefix+"tcp_received_nosyn_synsent", "Number of TCP received nosyn synsent", labelsTCP, nil)
	tcpReceivedBadrstSynSent = prometheus.NewDesc(prefix+"tcp_received_badrst_synsent", "Number of TCP received badrst synsent", labelsTCP, nil)
	tcpReceivedBadrstListenState = prometheus.NewDesc(prefix+"tcp_received_badrst_listenstate", "Number of TCP received badrst listenstate", labelsTCP, nil)
	tcpOptionMaxsegmentLength = prometheus.NewDesc(prefix+"tcp_option_maxsegment_length", "Number of TCP option maxsegment length", labelsTCP, nil)
	tcpOptionWindowLength = prometheus.NewDesc(prefix+"tcp_option_window_length", "Number of TCP option window length", labelsTCP, nil)
	tcpOptionTimestampLength = prometheus.NewDesc(prefix+"tcp_option_timestamp_length", "Number of TCP option timestamp length", labelsTCP, nil)
	tcpOptionMd5Length = prometheus.NewDesc(prefix+"tcp_option_md5_length", "Number of TCP option md5 length", labelsTCP, nil)
	tcpOptionAuthLength = prometheus.NewDesc(prefix+"tcp_option_auth_length", "Number of TCP option auth length", labelsTCP, nil)
	tcpOptionSackpermittedLength = prometheus.NewDesc(prefix+"tcp_option_sackpermitted_length", "Number of TCP option sackpermitted length", labelsTCP, nil)
	tcpOptionSackLength = prometheus.NewDesc(prefix+"tcp_option_sack_length", "Number of TCP option sack length", labelsTCP, nil)
	tcpOptionAuthoptionLength = prometheus.NewDesc(prefix+"tcp_option_authoption_length", "Number of TCP option authoption length", labelsTCP, nil)

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

	ch <- tcpPacketsSent
	ch <- tcpSentDataPackets
	ch <- tcpDataPacketsBytes
	ch <- tcpSentDataPacketsRetransmitted
	ch <- tcpRetransmittedBytes
	ch <- tcpSentDataUnnecessaryRetransmitted
	ch <- tcpSentResendsByMtuDiscovery
	ch <- tcpSentAckOnlyPackets
	ch <- tcpSentPacketsDelayed
	ch <- tcpSentUrgOnlyPackets
	ch <- tcpSentWindowProbePackets
	ch <- tcpSentWindowUpdatePackets
	ch <- tcpSentControlPackets
	ch <- tcpPacketsReceived
	ch <- tcpReceivedAcks
	ch <- tcpAcksBytes
	ch <- tcpReceivedDuplicateAcks
	ch <- tcpReceivedAcksForUnsentData
	ch <- tcpPacketsReceivedInSequence
	ch <- tcpInSequenceBytes
	ch <- tcpReceivedCompletelyDuplicatePacket
	ch <- tcpDuplicateInBytes
	ch <- tcpReceivedOldDuplicatePackets
	ch <- tcpReceivedPacketsWithSomeDupliacteData
	ch <- tcpSomeDuplicateInBytes
	ch <- tcpReceivedOutOfOrderPackets
	ch <- tcpOutOfOrderInBytes
	ch <- tcpReceivedPacketsOfDataAfterWindow
	ch <- tcpBytes
	ch <- tcpReceivedWindowProbes
	ch <- tcpReceivedWindowUpdatePackets
	ch <- tcpPacketsReceivedAfterClose
	ch <- tcpReceivedDiscardedForBadChecksum
	ch <- tcpReceivedDiscardedForBadHeaderOffset
	ch <- tcpReceivedDiscardedBecausePacketTooShort
	ch <- tcpConnectionRequests
	ch <- tcpConnectionAccepts
	ch <- tcpBadConnectionAttempts
	ch <- tcpListenQueueOverflows
	ch <- tcpBadRstWindow
	ch <- tcpConnectionsEstablished
	ch <- tcpConnectionsClosed
	ch <- tcpDrops
	ch <- tcpConnectionsUpdatedRttOnClose
	ch <- tcpConnectionsUpdatedVarianceOnClose
	ch <- tcpConnectionsUpdatedSsthreshOnClose
	ch <- tcpEmbryonicConnectionsDropped
	ch <- tcpSegmentsUpdatedRtt
	ch <- tcpAttempts
	ch <- tcpRetransmitTimeouts
	ch <- tcpConnectionsDroppedByRetransmitTimeout
	ch <- tcpPersistTimeouts
	ch <- tcpConnectionsDroppedByPersistTimeout
	ch <- tcpKeepaliveTimeouts
	ch <- tcpKeepaliveProbesSent
	ch <- tcpKeepaliveConnectionsDropped
	ch <- tcpAckHeaderPredictions
	ch <- tcpDataPacketHeaderPredictions
	ch <- tcpSyncacheEntriesAdded
	ch <- tcpRetransmitted
	ch <- tcpDupsyn
	ch <- tcpDropped
	ch <- tcpCompleted
	ch <- tcpBucketOverflow
	ch <- tcpCacheOverflow
	ch <- tcpReset
	ch <- tcpStale
	ch <- tcpAborted
	ch <- tcpBadack
	ch <- tcpUnreach
	ch <- tcpZoneFailures
	ch <- tcpCookiesSent
	ch <- tcpCookiesReceived
	ch <- tcpSackRecoveryEpisodes
	ch <- tcpSegmentRetransmits
	ch <- tcpByteRetransmits
	ch <- tcpSackOptionsReceived
	ch <- tcpSackOptionsSent
	ch <- tcpSackScoreboardOverflow
	ch <- tcpAcksSentInResponseButNotExactRsts
	ch <- tcpAcksSentInResponseToSynsOnEstablishedConnections
	ch <- tcpRcvPacketsDroppedDueToBadAddress
	ch <- tcpOutOfSequenceSegmentDrops
	ch <- tcpRstPackets
	ch <- tcpIcmpPacketsIgnored
	ch <- tcpSendPacketsDropped
	ch <- tcpRcvPacketsDropped
	ch <- tcpOutgoingSegmentsDropped
	ch <- tcpReceivedSynfinDropped
	ch <- tcpReceivedIpsecDropped
	ch <- tcpReceivedMacDropped
	ch <- tcpReceivedMinttlExceeded
	ch <- tcpListenstateBadflagsDropped
	ch <- tcpFinwaitstateBadflagsDropped
	ch <- tcpReceivedDosAttack
	ch <- tcpReceivedBadSynack
	ch <- tcpSyncacheZoneFull
	ch <- tcpReceivedRstFirewallfilter
	ch <- tcpReceivedNoackTimewait
	ch <- tcpReceivedNoTimewaitState
	ch <- tcpReceivedRstTimewaitState
	ch <- tcpReceivedTimewaitDrops
	ch <- tcpReceivedBadaddrTimewaitState
	ch <- tcpReceivedAckoffInSynSentrcvd
	ch <- tcpReceivedBadaddrFirewall
	ch <- tcpReceivedNosynSynSent
	ch <- tcpReceivedBadrstSynSent
	ch <- tcpReceivedBadrstListenState
	ch <- tcpOptionMaxsegmentLength
	ch <- tcpOptionWindowLength
	ch <- tcpOptionTimestampLength
	ch <- tcpOptionMd5Length
	ch <- tcpOptionAuthLength
	ch <- tcpOptionSackpermittedLength
	ch <- tcpOptionSackLength
	ch <- tcpOptionAuthoptionLength

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
	c.collectSystemStatisticsTCP(ch, labelValues, s)
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

func (c *systemstatisticsCollector) collectSystemStatisticsTCP(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	l := append(labelValues, "tcp")
	ch <- prometheus.MustNewConstMetric(tcpPacketsSent, prometheus.CounterValue, s.Statistics.Tcp.PacketsSent, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentDataPackets, prometheus.CounterValue, s.Statistics.Tcp.SentDataPackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpDataPacketsBytes, prometheus.CounterValue, s.Statistics.Tcp.DataPacketsBytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentDataPacketsRetransmitted, prometheus.CounterValue, s.Statistics.Tcp.SentDataPacketsRetransmitted, l...)
	ch <- prometheus.MustNewConstMetric(tcpRetransmittedBytes, prometheus.CounterValue, s.Statistics.Tcp.RetransmittedBytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentDataUnnecessaryRetransmitted, prometheus.CounterValue, s.Statistics.Tcp.SentDataUnnecessaryRetransmitted, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentResendsByMtuDiscovery, prometheus.CounterValue, s.Statistics.Tcp.SentResendsByMtuDiscovery, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentAckOnlyPackets, prometheus.CounterValue, s.Statistics.Tcp.SentAckOnlyPackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentPacketsDelayed, prometheus.CounterValue, s.Statistics.Tcp.SentPacketsDelayed, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentUrgOnlyPackets, prometheus.CounterValue, s.Statistics.Tcp.SentUrgOnlyPackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentWindowProbePackets, prometheus.CounterValue, s.Statistics.Tcp.SentWindowProbePackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentWindowUpdatePackets, prometheus.CounterValue, s.Statistics.Tcp.SentWindowUpdatePackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpSentControlPackets, prometheus.CounterValue, s.Statistics.Tcp.SentControlPackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpPacketsReceived, prometheus.CounterValue, s.Statistics.Tcp.PacketsReceived, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedAcks, prometheus.CounterValue, s.Statistics.Tcp.ReceivedAcks, l...)
	ch <- prometheus.MustNewConstMetric(tcpAcksBytes, prometheus.CounterValue, s.Statistics.Tcp.AcksBytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedDuplicateAcks, prometheus.CounterValue, s.Statistics.Tcp.ReceivedDuplicateAcks, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedAcksForUnsentData, prometheus.CounterValue, s.Statistics.Tcp.ReceivedAcksForUnsentData, l...)
	ch <- prometheus.MustNewConstMetric(tcpPacketsReceivedInSequence, prometheus.CounterValue, s.Statistics.Tcp.PacketsReceivedInSequence, l...)
	ch <- prometheus.MustNewConstMetric(tcpInSequenceBytes, prometheus.CounterValue, s.Statistics.Tcp.InSequenceBytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedCompletelyDuplicatePacket, prometheus.CounterValue, s.Statistics.Tcp.ReceivedCompletelyDuplicatePacket, l...)
	ch <- prometheus.MustNewConstMetric(tcpDuplicateInBytes, prometheus.CounterValue, s.Statistics.Tcp.DuplicateInBytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedOldDuplicatePackets, prometheus.CounterValue, s.Statistics.Tcp.ReceivedOldDuplicatePackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpSomeDuplicateInBytes, prometheus.CounterValue, s.Statistics.Tcp.SomeDuplicateInBytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedOutOfOrderPackets, prometheus.CounterValue, s.Statistics.Tcp.ReceivedOutOfOrderPackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedPacketsOfDataAfterWindow, prometheus.CounterValue, s.Statistics.Tcp.ReceivedPacketsOfDataAfterWindow, l...)
	ch <- prometheus.MustNewConstMetric(tcpBytes, prometheus.CounterValue, s.Statistics.Tcp.Bytes, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedWindowProbes, prometheus.CounterValue, s.Statistics.Tcp.ReceivedWindowProbes, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedWindowUpdatePackets, prometheus.CounterValue, s.Statistics.Tcp.ReceivedWindowUpdatePackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpPacketsReceivedAfterClose, prometheus.CounterValue, s.Statistics.Tcp.PacketsReceivedAfterClose, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedDiscardedForBadChecksum, prometheus.CounterValue, s.Statistics.Tcp.ReceivedDiscardedForBadChecksum, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedDiscardedForBadHeaderOffset, prometheus.CounterValue, s.Statistics.Tcp.ReceivedDiscardedForBadHeaderOffset, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedDiscardedBecausePacketTooShort, prometheus.CounterValue, s.Statistics.Tcp.ReceivedDiscardedBecausePacketTooShort, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionRequests, prometheus.CounterValue, s.Statistics.Tcp.ConnectionRequests, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionAccepts, prometheus.CounterValue, s.Statistics.Tcp.ConnectionAccepts, l...)
	ch <- prometheus.MustNewConstMetric(tcpBadConnectionAttempts, prometheus.CounterValue, s.Statistics.Tcp.BadConnectionAttempts, l...)
	ch <- prometheus.MustNewConstMetric(tcpListenQueueOverflows, prometheus.CounterValue, s.Statistics.Tcp.ListenQueueOverflows, l...)
	ch <- prometheus.MustNewConstMetric(tcpBadRstWindow, prometheus.CounterValue, s.Statistics.Tcp.BadRstWindow, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsEstablished, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsEstablished, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsClosed, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsClosed, l...)
	ch <- prometheus.MustNewConstMetric(tcpDrops, prometheus.CounterValue, s.Statistics.Tcp.Drops, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsUpdatedRttOnClose, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsUpdatedRttOnClose, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsUpdatedVarianceOnClose, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsUpdatedVarianceOnClose, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsUpdatedSsthreshOnClose, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsUpdatedSsthreshOnClose, l...)
	ch <- prometheus.MustNewConstMetric(tcpEmbryonicConnectionsDropped, prometheus.CounterValue, s.Statistics.Tcp.EmbryonicConnectionsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpSegmentsUpdatedRtt, prometheus.CounterValue, s.Statistics.Tcp.SegmentsUpdatedRtt, l...)
	ch <- prometheus.MustNewConstMetric(tcpAttempts, prometheus.CounterValue, s.Statistics.Tcp.Attempts, l...)
	ch <- prometheus.MustNewConstMetric(tcpRetransmitTimeouts, prometheus.CounterValue, s.Statistics.Tcp.RetransmitTimeouts, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsDroppedByRetransmitTimeout, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsDroppedByRetransmitTimeout, l...)
	ch <- prometheus.MustNewConstMetric(tcpPersistTimeouts, prometheus.CounterValue, s.Statistics.Tcp.PersistTimeouts, l...)
	ch <- prometheus.MustNewConstMetric(tcpConnectionsDroppedByPersistTimeout, prometheus.CounterValue, s.Statistics.Tcp.ConnectionsDroppedByPersistTimeout, l...)
	ch <- prometheus.MustNewConstMetric(tcpKeepaliveTimeouts, prometheus.CounterValue, s.Statistics.Tcp.KeepaliveTimeouts, l...)
	ch <- prometheus.MustNewConstMetric(tcpKeepaliveProbesSent, prometheus.CounterValue, s.Statistics.Tcp.KeepaliveProbesSent, l...)
	ch <- prometheus.MustNewConstMetric(tcpKeepaliveConnectionsDropped, prometheus.CounterValue, s.Statistics.Tcp.KeepaliveConnectionsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpAckHeaderPredictions, prometheus.CounterValue, s.Statistics.Tcp.AckHeaderPredictions, l...)
	ch <- prometheus.MustNewConstMetric(tcpDataPacketHeaderPredictions, prometheus.CounterValue, s.Statistics.Tcp.DataPacketHeaderPredictions, l...)
	ch <- prometheus.MustNewConstMetric(tcpSyncacheEntriesAdded, prometheus.CounterValue, s.Statistics.Tcp.SyncacheEntriesAdded, l...)
	ch <- prometheus.MustNewConstMetric(tcpRetransmitted, prometheus.CounterValue, s.Statistics.Tcp.Retransmitted, l...)
	ch <- prometheus.MustNewConstMetric(tcpDupsyn, prometheus.CounterValue, s.Statistics.Tcp.Dupsyn, l...)
	ch <- prometheus.MustNewConstMetric(tcpDropped, prometheus.CounterValue, s.Statistics.Tcp.Dropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpCompleted, prometheus.CounterValue, s.Statistics.Tcp.Completed, l...)
	ch <- prometheus.MustNewConstMetric(tcpBucketOverflow, prometheus.CounterValue, s.Statistics.Tcp.BucketOverflow, l...)
	ch <- prometheus.MustNewConstMetric(tcpCacheOverflow, prometheus.CounterValue, s.Statistics.Tcp.CacheOverflow, l...)
	ch <- prometheus.MustNewConstMetric(tcpReset, prometheus.CounterValue, s.Statistics.Tcp.Reset, l...)
	ch <- prometheus.MustNewConstMetric(tcpStale, prometheus.CounterValue, s.Statistics.Tcp.Stale, l...)
	ch <- prometheus.MustNewConstMetric(tcpAborted, prometheus.CounterValue, s.Statistics.Tcp.Aborted, l...)
	ch <- prometheus.MustNewConstMetric(tcpBadack, prometheus.CounterValue, s.Statistics.Tcp.Badack, l...)
	ch <- prometheus.MustNewConstMetric(tcpUnreach, prometheus.CounterValue, s.Statistics.Tcp.Unreach, l...)
	ch <- prometheus.MustNewConstMetric(tcpZoneFailures, prometheus.CounterValue, s.Statistics.Tcp.ZoneFailures, l...)
	ch <- prometheus.MustNewConstMetric(tcpCookiesSent, prometheus.CounterValue, s.Statistics.Tcp.CookiesSent, l...)
	ch <- prometheus.MustNewConstMetric(tcpCookiesReceived, prometheus.CounterValue, s.Statistics.Tcp.CookiesReceived, l...)
	ch <- prometheus.MustNewConstMetric(tcpSackRecoveryEpisodes, prometheus.CounterValue, s.Statistics.Tcp.SackRecoveryEpisodes, l...)
	ch <- prometheus.MustNewConstMetric(tcpSegmentRetransmits, prometheus.CounterValue, s.Statistics.Tcp.SegmentRetransmits, l...)
	ch <- prometheus.MustNewConstMetric(tcpByteRetransmits, prometheus.CounterValue, s.Statistics.Tcp.ByteRetransmits, l...)
	ch <- prometheus.MustNewConstMetric(tcpSackOptionsReceived, prometheus.CounterValue, s.Statistics.Tcp.SackOptionsReceived, l...)
	ch <- prometheus.MustNewConstMetric(tcpSackOptionsSent, prometheus.CounterValue, s.Statistics.Tcp.SackOptionsReceived, l...)
	ch <- prometheus.MustNewConstMetric(tcpSackScoreboardOverflow, prometheus.CounterValue, s.Statistics.Tcp.SackScoreboardOverflow, l...)
	ch <- prometheus.MustNewConstMetric(tcpAcksSentInResponseButNotExactRsts, prometheus.CounterValue, s.Statistics.Tcp.AcksSentInResponseButNotExactRsts, l...)
	ch <- prometheus.MustNewConstMetric(tcpAcksSentInResponseToSynsOnEstablishedConnections, prometheus.CounterValue, s.Statistics.Tcp.AcksSentInResponseToSynsOnEstablishedConnections, l...)
	ch <- prometheus.MustNewConstMetric(tcpRcvPacketsDroppedDueToBadAddress, prometheus.CounterValue, s.Statistics.Tcp.RcvPacketsDroppedDueToBadAddress, l...)
	ch <- prometheus.MustNewConstMetric(tcpOutOfSequenceSegmentDrops, prometheus.CounterValue, s.Statistics.Tcp.OutOfSequenceSegmentDrops, l...)
	ch <- prometheus.MustNewConstMetric(tcpRstPackets, prometheus.CounterValue, s.Statistics.Tcp.RstPackets, l...)
	ch <- prometheus.MustNewConstMetric(tcpIcmpPacketsIgnored, prometheus.CounterValue, s.Statistics.Tcp.IcmpPacketsIgnored, l...)
	ch <- prometheus.MustNewConstMetric(tcpSendPacketsDropped, prometheus.CounterValue, s.Statistics.Tcp.SendPacketsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpRcvPacketsDropped, prometheus.CounterValue, s.Statistics.Tcp.RcvPacketsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpOutgoingSegmentsDropped, prometheus.CounterValue, s.Statistics.Tcp.OutgoingSegmentsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedSynfinDropped, prometheus.CounterValue, s.Statistics.Tcp.ReceivedSynfinDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedIpsecDropped, prometheus.CounterValue, s.Statistics.Tcp.ReceivedIpsecDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedMacDropped, prometheus.CounterValue, s.Statistics.Tcp.ReceivedMacDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedMinttlExceeded, prometheus.CounterValue, s.Statistics.Tcp.ReceivedMinttlExceeded, l...)
	ch <- prometheus.MustNewConstMetric(tcpListenstateBadflagsDropped, prometheus.CounterValue, s.Statistics.Tcp.ListenstateBadflagsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpFinwaitstateBadflagsDropped, prometheus.CounterValue, s.Statistics.Tcp.FinwaitstateBadflagsDropped, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedDosAttack, prometheus.CounterValue, s.Statistics.Tcp.ReceivedDosAttack, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedBadSynack, prometheus.CounterValue, s.Statistics.Tcp.ReceivedBadSynack, l...)
	ch <- prometheus.MustNewConstMetric(tcpSyncacheZoneFull, prometheus.CounterValue, s.Statistics.Tcp.SyncacheZoneFull, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedRstFirewallfilter, prometheus.CounterValue, s.Statistics.Tcp.ReceivedRstFirewallfilter, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedNoackTimewait, prometheus.CounterValue, s.Statistics.Tcp.ReceivedNoackTimewait, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedNoTimewaitState, prometheus.CounterValue, s.Statistics.Tcp.ReceivedNoTimewaitState, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedRstTimewaitState, prometheus.CounterValue, s.Statistics.Tcp.ReceivedRstTimewaitState, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedTimewaitDrops, prometheus.CounterValue, s.Statistics.Tcp.ReceivedTimewaitDrops, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedBadaddrTimewaitState, prometheus.CounterValue, s.Statistics.Tcp.ReceivedBadaddrTimewaitState, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedAckoffInSynSentrcvd, prometheus.CounterValue, s.Statistics.Tcp.ReceivedAckoffInSynSentrcvd, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedBadaddrFirewall, prometheus.CounterValue, s.Statistics.Tcp.ReceivedBadaddrFirewall, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedNosynSynSent, prometheus.CounterValue, s.Statistics.Tcp.ReceivedNosynSynSent, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedBadrstSynSent, prometheus.CounterValue, s.Statistics.Tcp.ReceivedBadrstSynSent, l...)
	ch <- prometheus.MustNewConstMetric(tcpReceivedBadrstListenState, prometheus.CounterValue, s.Statistics.Tcp.ReceivedBadrstListenState, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionMaxsegmentLength, prometheus.CounterValue, s.Statistics.Tcp.OptionMaxsegmentLength, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionWindowLength, prometheus.CounterValue, s.Statistics.Tcp.OptionWindowLength, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionTimestampLength, prometheus.CounterValue, s.Statistics.Tcp.OptionTimestampLength, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionMd5Length, prometheus.CounterValue, s.Statistics.Tcp.OptionMd5Length, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionAuthLength, prometheus.CounterValue, s.Statistics.Tcp.OptionAuthLength, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionSackpermittedLength, prometheus.CounterValue, s.Statistics.Tcp.OptionSackpermittedLength, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionSackLength, prometheus.CounterValue, s.Statistics.Tcp.OptionSackLength, l...)
	ch <- prometheus.MustNewConstMetric(tcpOptionAuthoptionLength, prometheus.CounterValue, s.Statistics.Tcp.OptionAuthoptionLength, l...)
}
