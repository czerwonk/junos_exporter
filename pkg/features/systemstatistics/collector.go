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

	arpDatagramsReceivedDesc                                     *prometheus.Desc
	arpRequestsReceivedDesc                                      *prometheus.Desc
	arpRepliesReceivedDesc                                       *prometheus.Desc
	arpResolutionRequestReceivedDesc                             *prometheus.Desc
	arpResolutionRequestDroppedDesc                              *prometheus.Desc
	arpUnrestrictedProxyRequestsDesc                             *prometheus.Desc
	arpRestrictedProxyRequestsDesc                               *prometheus.Desc
	arpReceivedProxyRequestsDesc                                 *prometheus.Desc
	arpProxyRequestsNotProxiedDesc                               *prometheus.Desc
	arpRestrictedProxyRequestsNotProxiedDesc                     *prometheus.Desc
	arpDatagramsWithBogusInterfaceDesc                           *prometheus.Desc
	arpDatagramsWithIncorrectLengthDesc                          *prometheus.Desc
	arpDatagramsForNonIpProtocolDesc                             *prometheus.Desc
	arpDatagramsWithUnsupportedOpcodeDesc                        *prometheus.Desc
	arpDatagramsWithBadProtocolAddressLengthDesc                 *prometheus.Desc
	arpDatagramsWithBadHardwareAddressLengthDesc                 *prometheus.Desc
	arpDatagramsWithMulticastSourceAddressDesc                   *prometheus.Desc
	arpDatagramsWithMulticastTargetAddressDesc                   *prometheus.Desc
	arpDatagramsWithMyOwnHardwareAddressDesc                     *prometheus.Desc
	arpDatagramsForAnAddressNotOnTheInterfaceDesc                *prometheus.Desc
	arpDatagramsWithABroadcastSourceAddressDesc                  *prometheus.Desc
	arpDatagramsWithSourceAddressDuplicateToMineDesc             *prometheus.Desc
	arpDatagramsWhichWereNotForMeDesc                            *prometheus.Desc
	arpPacketsDiscardedWaitingForResolutionDesc                  *prometheus.Desc
	arpPacketsSentAfterWaitingForResolutionDesc                  *prometheus.Desc
	arpRequestsSentDesc                                          *prometheus.Desc
	arpRepliesSentDesc                                           *prometheus.Desc
	arpRequestsForMemoryDeniedDesc                               *prometheus.Desc
	arpRequestsDroppedOnEntryDesc                                *prometheus.Desc
	arpRequestsDroppedDuringRetryDesc                            *prometheus.Desc
	arpRequestsDroppedDueToInterfaceDeletionDesc                 *prometheus.Desc
	arpRequestsOnUnnumberedInterfacesDesc                        *prometheus.Desc
	arpNewRequestsOnUnnumberedInterfacesDesc                     *prometheus.Desc
	arpRepliesFromUnnumberedInterfacesDesc                       *prometheus.Desc
	arpRequestsOnUnnumberedInterfaceWithNonSubnettedDonorDesc    *prometheus.Desc
	arpRepliesFromUnnumberedInterfaceWithNonSubnettedDonorDesc   *prometheus.Desc
	arpPacketsRejectedAsFamilyIsConfiguredWithDenyArpDesc        *prometheus.Desc
	arpResponsePacketsAreRejectedOnMcAeIclInterfaceDesc          *prometheus.Desc
	arpRepliesAreRejectedAsSourceAndDestinationIsSameDesc        *prometheus.Desc
	arpProbeForProxyAddressReachableFromTheIncomingInterfaceDesc *prometheus.Desc
	arpRequestDiscardedForVrrpSourceAddressDesc                  *prometheus.Desc
	arpSelfArpRequestPacketReceivedOnIrbInterfaceDesc            *prometheus.Desc
	arpProxyArpRequestDiscardedAsSourceIpIsAProxyTargetDesc      *prometheus.Desc
	arpPacketsAreDroppedAsNexthopAllocationFailedDesc            *prometheus.Desc
	arpPacketsReceivedFromPeerVrrpRouterAndDiscardedDesc         *prometheus.Desc
	arpPacketsAreRejectedAsTargetIpArpResolveIsInProgressDesc    *prometheus.Desc
	arpGratArpPacketsAreIgnoredAsMacAddressIsNotChangedDesc      *prometheus.Desc
	arpPacketsAreDroppedFromPeerVrrpDesc                         *prometheus.Desc
	arpPacketsAreDroppedAsDriverCallFailedDesc                   *prometheus.Desc
	arpPacketsAreDroppedAsSourceIsNotValidatedDesc               *prometheus.Desc
	arpSystemMaxDesc                                             *prometheus.Desc
	arpPublicMaxDesc                                             *prometheus.Desc
	arpIriMaxDesc                                                *prometheus.Desc
	arpMgtMaxDesc                                                *prometheus.Desc
	arpPublicCntDesc                                             *prometheus.Desc
	arpIriCntDesc                                                *prometheus.Desc
	arpMgtCntDesc                                                *prometheus.Desc
	arpSystemDropDesc                                            *prometheus.Desc
	arpPublicDropDesc                                            *prometheus.Desc
	arpIriDropDesc                                               *prometheus.Desc
	arpMgtDropDesc                                               *prometheus.Desc

	icmpDropsDueToRateLimitDesc                                      *prometheus.Desc
	icmpCallsToIcmpErrorDesc                                         *prometheus.Desc
	icmpErrorsNotGeneratedBecauseOldMessageWasIcmp                   *prometheus.Desc
	icmpIcmpEchoReplyDesc                                            *prometheus.Desc
	icmpDestinationUnreachableDesc                                   *prometheus.Desc
	icmpIcmpEchoDesc                                                 *prometheus.Desc
	icmpTimeStampReplyDesc                                           *prometheus.Desc
	icmpTimeExceededDesc                                             *prometheus.Desc
	icmpTimeStampDesc                                                *prometheus.Desc
	icmpAddressMaskRequestDesc                                       *prometheus.Desc
	icmpAnEndpointChangedItsCookieSecretDesc                         *prometheus.Desc
	icmpMessagesWithBadCodeFieldsDesc                                *prometheus.Desc
	icmpMessagesLessThanTheMinimumLengthDesc                         *prometheus.Desc
	icmpMessagesWithBadChecksumDesc                                  *prometheus.Desc
	icmpMessagesWithBadSourceAddressDesc                             *prometheus.Desc
	icmpMessagesWithBadLengthDesc                                    *prometheus.Desc
	icmpEchoDropsWithBroadcastOrMulticastDestinatonAddressDesc       *prometheus.Desc
	icmpTimestampDropsWithBroadcastOrMulticastDestinationAddressDesc *prometheus.Desc
	icmpMessageResponsesGeneratedDesc                                *prometheus.Desc

	icmp6CallsToIcmp6ErrorDesc                               *prometheus.Desc
	icmp6ErrorsNotGeneratedBecauseOldMessageWasIcmpErrorDesc *prometheus.Desc
	icmp6ErrorsNotGeneratedBecauseRateLimitationDesc         *prometheus.Desc
	icmp6UnreachableIcmp6PacketsOutputHistoDesc              *prometheus.Desc
	icmp6Icmp6EchoOutputHistoDesc                            *prometheus.Desc
	icmp6Icmp6EchoReplyOutputHistoDesc                       *prometheus.Desc
	icmp6NeighborSolicitationOutputHistoDesc                 *prometheus.Desc
	icmp6NeighborAdvertisementOutputHistoDesc                *prometheus.Desc
	icmp6Icmp6MessagesWithBadCodeFieldsDesc                  *prometheus.Desc
	icmp6MessagesLessThanMinimumLengthDesc                   *prometheus.Desc
	icmp6BadChecksumsDesc                                    *prometheus.Desc
	icmp6Icmp6MessagesWithBadLengthDesc                      *prometheus.Desc
	icmp6UnreachableIcmp6PacketInputHistosDesc               *prometheus.Desc
	icmp6PacketTooBigInputHistoDesc                          *prometheus.Desc
	icmp6TimeExceededIcmp6PacketsInputHistoDesc              *prometheus.Desc
	icmp6Icmp6EchoInputHistoDesc                             *prometheus.Desc
	icmp6Icmp6EchoReplyInputHistoDesc                        *prometheus.Desc
	icmp6RouterSolicitationIcmp6PacketsInputHistoDesc        *prometheus.Desc
	icmp6NeighborSolicitationInputHistoDesc                  *prometheus.Desc
	icmp6NeighborAdvertisementInputHistoDesc                 *prometheus.Desc
	icmp6NoRouteDesc                                         *prometheus.Desc
	icmp6AdministrativelyProhibitedDesc                      *prometheus.Desc
	icmp6BeyondScopeDesc                                     *prometheus.Desc
	icmp6AddressUnreachableDesc                              *prometheus.Desc
	icmp6PortUnreachableDesc                                 *prometheus.Desc
	icmp6PacketTooBigDesc                                    *prometheus.Desc
	icmp6TimeExceedTransitDesc                               *prometheus.Desc
	icmp6TimeExceedReassemblyDesc                            *prometheus.Desc
	icmp6ErroneousHeaderFieldDesc                            *prometheus.Desc
	icmp6UnrecognizedNextHeaderDesc                          *prometheus.Desc
	icmp6UnrecognizedOptionDesc                              *prometheus.Desc
	icmp6RedirectDesc                                        *prometheus.Desc
	icmp6UnknownDesc                                         *prometheus.Desc
	icmp6Icmp6MessageResponsesGeneratedDesc                  *prometheus.Desc
	icmp6MessagesWithTooManyNdOptionsDesc                    *prometheus.Desc
	icmp6NdSystemMaxDesc                                     *prometheus.Desc
	icmp6NdPublicMaxDesc                                     *prometheus.Desc
	icmp6NdIriMaxDesc                                        *prometheus.Desc
	icmp6NdMgtMaxDesc                                        *prometheus.Desc
	icmp6NdPublicCntDesc                                     *prometheus.Desc
	icmp6NdIriCntDesc                                        *prometheus.Desc
	icmp6NdMgtCntDesc                                        *prometheus.Desc
	icmp6NdSystemDropDesc                                    *prometheus.Desc
	icmp6NdPublicDropDesc                                    *prometheus.Desc
	icmp6NdIriDropDesc                                       *prometheus.Desc
	icmp6NdMgtDropDesc                                       *prometheus.Desc
	icmp6Nd6NdpProxyRequestsDesc                             *prometheus.Desc
	icmp6Nd6DadProxyRequestsDesc                             *prometheus.Desc
	icmp6Nd6NdpProxyResponsesDesc                            *prometheus.Desc
	icmp6Nd6DadProxyConflictsDesc                            *prometheus.Desc
	icmp6Nd6DupProxyResponsesDesc                            *prometheus.Desc
	icmp6Nd6NdpProxyResolveCntDesc                           *prometheus.Desc
	icmp6Nd6DadProxyResolveCntDesc                           *prometheus.Desc
	icmp6Nd6DadProxyEqmacDropDesc                            *prometheus.Desc
	icmp6Nd6DadProxyNomacDropDesc                            *prometheus.Desc
	icmp6Nd6NdpProxyUnrRequestsDesc                          *prometheus.Desc
	icmp6Nd6DadProxyUnrRequestsDesc                          *prometheus.Desc
	icmp6Nd6NdpProxyUnrResponsesDesc                         *prometheus.Desc
	icmp6Nd6DadProxyUnrConflictsDesc                         *prometheus.Desc
	icmp6Nd6DadProxyUnrResponsesDesc                         *prometheus.Desc
	icmp6Nd6NdpProxyUnrResolveCntDesc                        *prometheus.Desc
	icmp6Nd6DadProxyUnrResolveCntDesc                        *prometheus.Desc
	icmp6Nd6DadProxyUnrEqportDropDesc                        *prometheus.Desc
	icmp6Nd6DadProxyUnrNomacDropDesc                         *prometheus.Desc
	icmp6Nd6RequestsDroppedOnEntryDesc                       *prometheus.Desc
	icmp6Nd6RequestsDroppedDuringRetryDesc                   *prometheus.Desc

	mplsTotalMplsPacketsReceivedDesc                  *prometheus.Desc
	mplsPacketsForwardedDesc                          *prometheus.Desc
	mplsPacketsDroppedDesc                            *prometheus.Desc
	mplsPacketsWithHeaderTooSmallDesc                 *prometheus.Desc
	mplsAfterTaggingPacketsCanNotFitLinkMtuDesc       *prometheus.Desc
	mplsPacketsWithIpv4ExplicitNullTagDesc            *prometheus.Desc
	mplsPacketsWithIpv4ExplicitNullChecksumErrorsDesc *prometheus.Desc
	mplsPacketsWithRouterAlertTagDesc                 *prometheus.Desc
	mplsLspPingPacketsDesc                            *prometheus.Desc
	mplsPacketsWithTtlExpiredDesc                     *prometheus.Desc
	mplsPacketsWithTagEncodingErrorDesc               *prometheus.Desc
	mplsPacketsDiscardedDueToNoRouteDesc              *prometheus.Desc
	mplsPacketsUsedFirstNexthopInEcmpUnilistDesc      *prometheus.Desc
	mplsPacketsDroppedDueToIflDownDesc                *prometheus.Desc
	mplsPacketsDroppedAtMplsSocketSendDesc            *prometheus.Desc
	mplsPacketsForwardedAtMplsSocketSendDesc          *prometheus.Desc
	mplsPacketsDroppedAtP2mpCnhOutputDesc             *prometheus.Desc
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

	labelsARP := []string{"target", "protocol"}
	arpDatagramsReceivedDesc = prometheus.NewDesc(prefix+"arp_datagrams_received", "Number of ARP datagrams received", labelsARP, nil)
	arpRequestsReceivedDesc = prometheus.NewDesc(prefix+"arp_requests_received", "Number of ARP requests received", labelsARP, nil)
	arpRepliesReceivedDesc = prometheus.NewDesc(prefix+"arp_replies_received", "Number of ARP replies received", labelsARP, nil)
	arpResolutionRequestReceivedDesc = prometheus.NewDesc(prefix+"arp_resolution_request_received", "Number of ARP resolution request received", labelsARP, nil)
	arpResolutionRequestDroppedDesc = prometheus.NewDesc(prefix+"arp_resolution_request_dropped", "Number of ARP resolution request dropped", labelsARP, nil)
	arpUnrestrictedProxyRequestsDesc = prometheus.NewDesc(prefix+"arp_unrestricted_proxy_requests", "Number of ARP unrestricted proxy requests", labelsARP, nil)
	arpRestrictedProxyRequestsDesc = prometheus.NewDesc(prefix+"arp_restricted_proxy_requests", "Number of ARP restricted proxy requests", labelsARP, nil)
	arpReceivedProxyRequestsDesc = prometheus.NewDesc(prefix+"arp_received_proxy_requests", "Number of ARP received proxy requests", labelsARP, nil)
	arpProxyRequestsNotProxiedDesc = prometheus.NewDesc(prefix+"arp_proxy_requests_not_proxied", "Number of ARP proxy requests not proxied", labelsARP, nil)
	arpRestrictedProxyRequestsNotProxiedDesc = prometheus.NewDesc(prefix+"arp_restricted_proxy_requests_not_proxied", "Number of ARP restricted proxy requests not proxied", labelsARP, nil)
	arpDatagramsWithBogusInterfaceDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_bogus_interface", "Number of ARP datagrams with bogus interface", labelsARP, nil)
	arpDatagramsWithIncorrectLengthDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_incorrect_length", "Number of ARP datagrams with incorrect length", labelsARP, nil)
	arpDatagramsForNonIpProtocolDesc = prometheus.NewDesc(prefix+"arp_datagrams_for_non_ip_protocol", "Number of ARP datagrams for non ip protocol", labelsARP, nil)
	arpDatagramsWithUnsupportedOpcodeDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_unsupported_opcode", "Number of ARP datagrams with unsupported opcode", labelsARP, nil)
	arpDatagramsWithBadProtocolAddressLengthDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_bad_protocol_address_length", "Number of ARP datagrams with bad protocol address length", labelsARP, nil)
	arpDatagramsWithBadHardwareAddressLengthDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_bad_hardware_address_length", "Number of ARP datagrams with bad hardware address length", labelsARP, nil)
	arpDatagramsWithMulticastSourceAddressDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_multicast_source_address", "Number of ARP datagrams with multicast source address", labelsARP, nil)
	arpDatagramsWithMulticastTargetAddressDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_multicast_target_address", "Number of ARP datagrams with multicast target address", labelsARP, nil)
	arpDatagramsWithMyOwnHardwareAddressDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_my_own_hardware_address", "Number of ARP datagrams with my own hardware address", labelsARP, nil)
	arpDatagramsForAnAddressNotOnTheInterfaceDesc = prometheus.NewDesc(prefix+"arp_datagrams_for_an_address_not_on_interface", "Number of ARP datagrams for an address not on the interface", labelsARP, nil)
	arpDatagramsWithABroadcastSourceAddressDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_a_broadcast_source_address", "Number of ARP datagrams with a broadcast source address", labelsARP, nil)
	arpDatagramsWithSourceAddressDuplicateToMineDesc = prometheus.NewDesc(prefix+"arp_datagrams_with_source_address_duplicate_to_mine", "Number of ARP datagrams with source address duplicate to mine", labelsARP, nil)
	arpDatagramsWhichWereNotForMeDesc = prometheus.NewDesc(prefix+"arp_datagrams_which_were_not_for_me", "Number of ARP datagrams which were not for me", labelsARP, nil)
	arpPacketsDiscardedWaitingForResolutionDesc = prometheus.NewDesc(prefix+"arp_packets_discarded_waiting_for_resolution", "Number of ARP packets discarded waiting for resolution", labelsARP, nil)
	arpPacketsSentAfterWaitingForResolutionDesc = prometheus.NewDesc(prefix+"arp_packets_sent_after_waiting_for_resolution", "Number of ARP packets sent after waiting for resolution", labelsARP, nil)
	arpRequestsSentDesc = prometheus.NewDesc(prefix+"arp_requests_sent", "Number of ARP requests sent", labelsARP, nil)
	arpRepliesSentDesc = prometheus.NewDesc(prefix+"arp_replies_sent", "Number of ARP replies sent", labelsARP, nil)
	arpRequestsForMemoryDeniedDesc = prometheus.NewDesc(prefix+"arp_requests_for_memory_denied", "Number of ARP requests for memory denied", labelsARP, nil)
	arpRequestsDroppedOnEntryDesc = prometheus.NewDesc(prefix+"arp_requests_dropped_on_entry", "Number of ARP requests dropped on entry", labelsARP, nil)
	arpRequestsDroppedDuringRetryDesc = prometheus.NewDesc(prefix+"arp_requests_dropped_during_retry", "Number of ARP requests dropped during retry", labelsARP, nil)
	arpRequestsDroppedDueToInterfaceDeletionDesc = prometheus.NewDesc(prefix+"arp_requests_dropped_due_to_interface_deletion", "Number of ARP requests dropped due to interface deletion", labelsARP, nil)
	arpRequestsOnUnnumberedInterfacesDesc = prometheus.NewDesc(prefix+"arp_requests_on_unnumbered_interfaces", "Number of ARP requests on unnumbered interfaces", labelsARP, nil)
	arpNewRequestsOnUnnumberedInterfacesDesc = prometheus.NewDesc(prefix+"arp_new_requests_on_unnumbered_interfaces", "Number of ARP requests on unnumbered interfaces", labelsARP, nil)
	arpRepliesFromUnnumberedInterfacesDesc = prometheus.NewDesc(prefix+"arp_replies_from_unnumbered_interfaces", "Number of ARP replies from unnumbered interfaces", labelsARP, nil)
	arpRequestsOnUnnumberedInterfaceWithNonSubnettedDonorDesc = prometheus.NewDesc(prefix+"arp_requests_on_unnumbered_interface_with_non_subnetted_donor", "Number of ARP requests on unnumbered interface with non subnetted donor", labelsARP, nil)
	arpRepliesFromUnnumberedInterfaceWithNonSubnettedDonorDesc = prometheus.NewDesc(prefix+"arp_replies_from_unnumbered_interface_with_non_subnet_donor", "Number of ARP replies from unnumbered interface with non subnetted donor", labelsARP, nil)
	arpPacketsRejectedAsFamilyIsConfiguredWithDenyArpDesc = prometheus.NewDesc(prefix+"arp_packets_rejected_as_family_is_configured_with_deny", "Number of ARP packets rejected as family is configured with deny", labelsARP, nil)
	arpResponsePacketsAreRejectedOnMcAeIclInterfaceDesc = prometheus.NewDesc(prefix+"arp_response_packets_are_rejected_on_McAeIcl_interface", "Number of ARP response packets are rejected on McAeIcl interface", labelsARP, nil)
	arpRepliesAreRejectedAsSourceAndDestinationIsSameDesc = prometheus.NewDesc(prefix+"arp_replies_are_rejected_as_source_and_destination_is_same", "Number of ARP replies tha tare rejected due to source and destination being same ", labelsARP, nil)
	arpProbeForProxyAddressReachableFromTheIncomingInterfaceDesc = prometheus.NewDesc(prefix+"arp_probe_for_proxy_address_reachable_from_the_incoming_interface", "Number of ARP probes for proxy address reachable from the incoming interface", labelsARP, nil)
	arpRequestDiscardedForVrrpSourceAddressDesc = prometheus.NewDesc(prefix+"arp_request_discarded_for_vrrp_source_address", "Number of ARP request discarded for vrrp source address", labelsARP, nil)
	arpSelfArpRequestPacketReceivedOnIrbInterfaceDesc = prometheus.NewDesc(prefix+"arp_self_arp_request_packet_received_on_irb_interface", "Number of ARP self arp request packet received on irb interface", labelsARP, nil)
	arpProxyArpRequestDiscardedAsSourceIpIsAProxyTargetDesc = prometheus.NewDesc(prefix+"arp_proxy_arp_request_discarded_as_source_ip_is_a_proxy_target", "Number of ARP proxy arp request discarded as source ip is a proxy target", labelsARP, nil)
	arpPacketsAreDroppedAsNexthopAllocationFailedDesc = prometheus.NewDesc(prefix+"arp_packets_are_dropped_as_nexthop_allocation_failed", "Number of ARP packets are dropped as nexthop allocation failed", labelsARP, nil)
	arpPacketsReceivedFromPeerVrrpRouterAndDiscardedDesc = prometheus.NewDesc(prefix+"apr_packets_received_from_peer_vrrp_router_and_discarded", "NUmber of ARP packets received from peer vrrp router and discarded", labelsARP, nil)
	arpPacketsAreRejectedAsTargetIpArpResolveIsInProgressDesc = prometheus.NewDesc(prefix+"arp_packets_are_rejected_as_target_ip_arp_resolve_is_in_progress", "Number of ARP packets are rejected as target ip arp resolve is in progress", labelsARP, nil)
	arpGratArpPacketsAreIgnoredAsMacAddressIsNotChangedDesc = prometheus.NewDesc(prefix+"arp_grat_arp_packets_are_ignored_as_mac_address_is_not_changed", "Number of ARP grat arp packets are ignored as mac address is not changed", labelsARP, nil)
	arpPacketsAreDroppedFromPeerVrrpDesc = prometheus.NewDesc(prefix+"apr_packets_are_dropped_from_peer_vrrp", "NUmber of ARP packets are dropped from peer vrrp", labelsARP, nil)
	arpPacketsAreDroppedAsDriverCallFailedDesc = prometheus.NewDesc(prefix+"arp_packets_are_dropped_as_driver_call_failed", "Number of ARP packets are dropped as driver call failed", labelsARP, nil)
	arpPacketsAreDroppedAsSourceIsNotValidatedDesc = prometheus.NewDesc(prefix+"arp_packets_are_dropped_as_source_is_not_validated", "Number of ARP packets are dropped as source is not validated", labelsARP, nil)
	arpSystemMaxDesc = prometheus.NewDesc(prefix+"arp_system_max", "Number of ARP system max", labelsARP, nil)
	arpPublicMaxDesc = prometheus.NewDesc(prefix+"arp_public_max", "Number of ARP public max", labelsARP, nil)
	arpIriMaxDesc = prometheus.NewDesc(prefix+"arp_iri_max", "Number of ARP iri max", labelsARP, nil)
	arpMgtMaxDesc = prometheus.NewDesc(prefix+"arp_mgnt_max", "Number of ARP mgnt max", labelsARP, nil)
	arpPublicCntDesc = prometheus.NewDesc(prefix+"arp_public_cnt", "Number of ARP public cnt", labelsARP, nil)
	arpIriCntDesc = prometheus.NewDesc(prefix+"arp_iri_cnt", "Number of ARP iri cnt", labelsARP, nil)
	arpMgtCntDesc = prometheus.NewDesc(prefix+"arp_mgnt_cnt", "Number of ARP mgnt cnt", labelsARP, nil)
	arpSystemDropDesc = prometheus.NewDesc(prefix+"arp_system_drop", "Number of ARP system drop", labelsARP, nil)
	arpPublicDropDesc = prometheus.NewDesc(prefix+"arp_public_drop", "Number of ARP public drop", labelsARP, nil)
	arpIriDropDesc = prometheus.NewDesc(prefix+"arp_iri_drop", "Number of ARP iri drop", labelsARP, nil)
	arpMgtDropDesc = prometheus.NewDesc(prefix+"arp_mgnt_drop", "Number of ARP mgnt drop", labelsARP, nil)

	labelsICMP := []string{"target", "protocol"}
	labelsICMPHistogram := []string{"target", "protocol", "histogram_type"}
	icmpDropsDueToRateLimitDesc = prometheus.NewDesc(prefix+"icmp_drops_due_to_rate_limit", "Number of ICMP drops due to rate limit", labelsICMP, nil)
	icmpCallsToIcmpErrorDesc = prometheus.NewDesc(prefix+"icmp_calls_to_icmp_error", "Number of ICMP calls to icmp error", labelsICMP, nil)
	icmpErrorsNotGeneratedBecauseOldMessageWasIcmp = prometheus.NewDesc(prefix+"icmp_errors_not_generated_because_old_message_was_icmp", "Number of ICMP errors not generated because old message was icmp", labelsICMP, nil)
	icmpIcmpEchoReplyDesc = prometheus.NewDesc(prefix+"icmp_echo_reply", "Number of ICMP echo reply", labelsICMPHistogram, nil)
	icmpDestinationUnreachableDesc = prometheus.NewDesc(prefix+"icmp_destination_unreachable", "Number of icmp destination unrechable", labelsICMPHistogram, nil)
	icmpIcmpEchoDesc = prometheus.NewDesc(prefix+"icmp_echo", "Number of icmp echos", labelsICMPHistogram, nil)
	icmpTimeStampReplyDesc = prometheus.NewDesc(prefix+"icmp_time_stamp_reply", "Number of icmp time stamp reply", labelsICMPHistogram, nil)
	icmpTimeExceededDesc = prometheus.NewDesc(prefix+"icmp_time_exceeded", "Number of icmp time exceeded", labelsICMPHistogram, nil)
	icmpTimeStampDesc = prometheus.NewDesc(prefix+"icmp_time_stamp", "Number of icmp time stamps", labelsICMPHistogram, nil)
	icmpAddressMaskRequestDesc = prometheus.NewDesc(prefix+"icmp_address_mask_request", "Number of icmp address mask requests", labelsICMPHistogram, nil)
	icmpAnEndpointChangedItsCookieSecretDesc = prometheus.NewDesc(prefix+"icmp_an_endpoint_changed_its_cookie_secret", "Number of icmp that changed its cookie secret an ednpoint", labelsICMPHistogram, nil)
	icmpMessagesWithBadCodeFieldsDesc = prometheus.NewDesc(prefix+"icmp_messages_with_bad_code_fields", "Number of icmp messages with bad code fields", labelsICMP, nil)
	icmpMessagesLessThanTheMinimumLengthDesc = prometheus.NewDesc(prefix+"icmp_messages_less_than_the_minimum_length", "Number of icmp messages less than the minimum length", labelsICMP, nil)
	icmpMessagesWithBadChecksumDesc = prometheus.NewDesc(prefix+"icmp_messages_with_bad_checksum", "Number of icmp messages with bad checksum", labelsICMP, nil)
	icmpMessagesWithBadSourceAddressDesc = prometheus.NewDesc(prefix+"icmp_messages_with_nad_source-address", "Number of icmp messages with bad source address", labelsICMP, nil)
	icmpMessagesWithBadLengthDesc = prometheus.NewDesc(prefix+"icmp_messages_with_bad_length", "Number of icmp messages with bad length", labelsICMP, nil)
	icmpEchoDropsWithBroadcastOrMulticastDestinatonAddressDesc = prometheus.NewDesc(prefix+"icmp_echo_drops_with_broadcast_or_multicast_destination_address", "Number of icmp echo drops with broadcast or multicast destination address", labelsICMP, nil)
	icmpTimestampDropsWithBroadcastOrMulticastDestinationAddressDesc = prometheus.NewDesc(prefix+"icmp_timestamp_drops_with_broadcast_or_multicast_destination_address", "Number of icmp timestamp drops with broadcast or multicast destination address", labelsICMP, nil)
	icmpMessageResponsesGeneratedDesc = prometheus.NewDesc(prefix+"icmp_message_responses_generated", "Number of icmp message responses generated", labelsICMP, nil)

	labelsICMP6 := []string{"target", "protocol"}
	icmp6CallsToIcmp6ErrorDesc = prometheus.NewDesc(prefix+"icmp6_calles_to_icmp6_error", "Number of icmp6 calls to icmp6 error", labelsICMP6, nil)
	icmp6ErrorsNotGeneratedBecauseOldMessageWasIcmpErrorDesc = prometheus.NewDesc(prefix+"icmp6_errors_not_generated_because_old_message_was_icmp_error", "Number of icmp6 errors not generated because old message was icmp error", labelsICMP6, nil)
	icmp6ErrorsNotGeneratedBecauseRateLimitationDesc = prometheus.NewDesc(prefix+"icmp6_errors_not_generated_because_rate_limitation", "Number of icmp6 errors not generated becasue of rate limition", labelsICMP6, nil)
	labelsICMP6Histogram := []string{"target", "protocol", "histogram_type"}
	icmp6UnreachableIcmp6PacketsOutputHistoDesc = prometheus.NewDesc(prefix+"icmp6_unreachable_icmp6_packets", "Number of icmp6 unreachable icmp6 packets", labelsICMP6Histogram, nil)
	icmp6Icmp6EchoOutputHistoDesc = prometheus.NewDesc(prefix+"icmp6_icmp6_echo", "Number of icmp6 echos", labelsICMP6Histogram, nil)
	icmp6Icmp6EchoReplyOutputHistoDesc = prometheus.NewDesc(prefix+"icmp6_icmp6_echo_reply", "Number of icmp6 echo replies", labelsICMP6Histogram, nil)
	icmp6NeighborSolicitationOutputHistoDesc = prometheus.NewDesc(prefix+"icmp6_neighbor_solicitation", "Number of icmp6 neighbor solicitation", labelsICMP6Histogram, nil)
	icmp6NeighborAdvertisementOutputHistoDesc = prometheus.NewDesc(prefix+"icmp6_neighbor_advertisement", "Number of icmp6 neighbor advertisement", labelsICMP6Histogram, nil)
	icmp6Icmp6MessagesWithBadCodeFieldsDesc = prometheus.NewDesc(prefix+"icmp6_messages_with_bad_code_fields", "Number of icmp6 messages with bad code fields", labelsICMP6, nil)
	icmp6MessagesLessThanMinimumLengthDesc = prometheus.NewDesc(prefix+"icmp6_messages_less_than_minimum_length", "Number of icmp6 messages less than minimum length", labelsICMP6, nil)
	icmp6BadChecksumsDesc = prometheus.NewDesc(prefix+"icmp6_bad_checksums", "Number of icmp6 bad checksums", labelsICMP6, nil)
	icmp6Icmp6MessagesWithBadLengthDesc = prometheus.NewDesc(prefix+"icmp6_messages_with_bad_length", "Number of icmp6 messages with bad length", labelsICMP6, nil)
	icmp6UnreachableIcmp6PacketInputHistosDesc = prometheus.NewDesc(prefix+"icmp6_unreachable_icmp6_packet_input_histogram", "Number of icmp6 unreachable icmp6 packets input histogram", labelsICMP6Histogram, nil)
	icmp6PacketTooBigInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_packet_too_big_input_histogram", "Number of icmp6 packets too big input histogram", labelsICMP6Histogram, nil)
	icmp6TimeExceededIcmp6PacketsInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_time_exceeded_icmp6_packets_input_histogram", "Number of icmp6 time exceeded packets input histogram", labelsICMP6Histogram, nil)
	icmp6Icmp6EchoInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_icmp6_echo_input_histogram", "Number of icmp6 echos input histogram", labelsICMP6Histogram, nil)
	icmp6Icmp6EchoReplyInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_echo_reply_input_histogram", "Number of icmp6 echo replies input histogram", labelsICMP6Histogram, nil)
	icmp6RouterSolicitationIcmp6PacketsInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_router_solicitation_icmp6_packets_input_histogram", "Number of icmp6 router solicitation packets input histogram", labelsICMP6Histogram, nil)
	icmp6NeighborSolicitationInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_neighbor_solicitation_input_histogram", "Number of icmp6 neighbor solicitation packets input histogram", labelsICMP6Histogram, nil)
	icmp6NeighborAdvertisementInputHistoDesc = prometheus.NewDesc(prefix+"icmp6_neighbor_advertisement_input_histogram", "Number of icmp6 neighbor advertisement packets input histogram", labelsICMP6Histogram, nil)
	icmp6NoRouteDesc = prometheus.NewDesc(prefix+"icmp6_no_route", "Number ofr icmp6 without route", labelsICMP6, nil)
	icmp6AdministrativelyProhibitedDesc = prometheus.NewDesc(prefix+"icmp6_administratively_prohibited", "Number of icmp6 adiminiostratively prohibited", labelsICMP6, nil)
	icmp6BeyondScopeDesc = prometheus.NewDesc(prefix+"icmp6_beyond_scope", "Number of icmp6 beyond scope", labelsICMP6, nil)
	icmp6AddressUnreachableDesc = prometheus.NewDesc(prefix+"icmp6_address_unreachable", "Number of icmp6 address unreachable", labelsICMP6, nil)
	icmp6PortUnreachableDesc = prometheus.NewDesc(prefix+"icmp6_port_unreachable", "Number of icmp6 port unreachable", labelsICMP6, nil)
	icmp6PacketTooBigDesc = prometheus.NewDesc(prefix+"icmp6_packet_too_big", "Number of ICMP6 packets that are too big", labelsICMP6, nil)
	icmp6TimeExceedTransitDesc = prometheus.NewDesc(prefix+"icmp6_time_exceed_transit", "Number of icmp6 packets whose time exceed transit", labelsICMP6, nil)
	icmp6TimeExceedReassemblyDesc = prometheus.NewDesc(prefix+"icmp6_time_exceed_reassembly", "Number of icmp6 packets whose time exceed reassembly", labelsICMP6, nil)
	icmp6ErroneousHeaderFieldDesc = prometheus.NewDesc(prefix+"icmp6_errors_on_header_file", "Number of icmp6 with errors on header file", labelsICMP6, nil)
	icmp6UnrecognizedNextHeaderDesc = prometheus.NewDesc(prefix+"icmp6_unrecognized_next_header", "Number of icmp6 with unrecognized next header", labelsICMP6, nil)
	icmp6UnrecognizedOptionDesc = prometheus.NewDesc(prefix+"icmp6_unrecognized_option", "Number of icmp6 with unrecognized option", labelsICMP6, nil)
	icmp6RedirectDesc = prometheus.NewDesc(prefix+"icmp6_redirect", "Number of icmp6 redirect", labelsICMP6, nil)
	icmp6UnknownDesc = prometheus.NewDesc(prefix+"icmp6_unknown", "Number of icmp6 unknown", labelsICMP6, nil)
	icmp6Icmp6MessageResponsesGeneratedDesc = prometheus.NewDesc(prefix+"icmp6_message_responses_generated", "Number of icmp6 message responses generated", labelsICMP6, nil)
	icmp6MessagesWithTooManyNdOptionsDesc = prometheus.NewDesc(prefix+"icmp6_message_with_to_many_nd_options", "Number of icmp6 messages with too many nd options", labelsICMP6, nil)
	icmp6NdSystemMaxDesc = prometheus.NewDesc(prefix+"icmp6_nd_system_max", "Number of icmp6 nd system max", labelsICMP6, nil)
	icmp6NdPublicMaxDesc = prometheus.NewDesc(prefix+"icmp6_nd_public_max", "Number of icmp6 nd public max", labelsICMP6, nil)
	icmp6NdIriMaxDesc = prometheus.NewDesc(prefix+"icmp6_nd_iri_max", "Number of icmp6 nd iri max", labelsICMP6, nil)
	icmp6NdMgtMaxDesc = prometheus.NewDesc(prefix+"icmp6_nd_mgt_max", "Number of icmp6 nd mgt max", labelsICMP6, nil)
	icmp6NdPublicCntDesc = prometheus.NewDesc(prefix+"icmp6_nd_public_cnt", "Number of icmp6 nd public cnt", labelsICMP6, nil)
	icmp6NdIriCntDesc = prometheus.NewDesc(prefix+"icmp6_nd_iri_cnt", "Number of icmp6 nd iri cnt", labelsICMP6, nil)
	icmp6NdMgtCntDesc = prometheus.NewDesc(prefix+"icmp6_nd_mgt_cnt", "Number of icmp6 nt mgt cnt", labelsICMP6, nil)
	icmp6NdSystemDropDesc = prometheus.NewDesc(prefix+"icmp6_nd_system_drop", "Number of icmp6 nd system drop", labelsICMP6, nil)
	icmp6NdPublicDropDesc = prometheus.NewDesc(prefix+"icmp6_nd_public_drop", "Number of icmp6 nd public drop", labelsICMP6, nil)
	icmp6NdIriDropDesc = prometheus.NewDesc(prefix+"icmp6_nd_iri_drop", "Number of icmp6 nd iri drops", labelsICMP6, nil)
	icmp6NdMgtDropDesc = prometheus.NewDesc(prefix+"icmp6_nd_mgt_drop", "Number of icmp6 nd mgt drop", labelsICMP6, nil)
	icmp6Nd6NdpProxyRequestsDesc = prometheus.NewDesc(prefix+"icmp6_ndp_proxy_requests", "Number of icmp6 ndp proxy requests", labelsICMP6, nil)
	icmp6Nd6DadProxyRequestsDesc = prometheus.NewDesc(prefix+"icmp6_nd_dad_proxy_requests", "Number of icmp6 nd dad proxy requests", labelsICMP6, nil)
	icmp6Nd6NdpProxyResponsesDesc = prometheus.NewDesc(prefix+"icmp6_nd6_ndp_proxy_responses", "Number of icmp6 ndp proxy responses", labelsICMP6, nil)
	icmp6Nd6DadProxyConflictsDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_conflicts", "Number of icmp6 nd6 dad proxy conflicts", labelsICMP6, nil)
	icmp6Nd6DupProxyResponsesDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dup_proxy_response", "Number of icmp6 nd6 dup proxy responses", labelsICMP6, nil)
	icmp6Nd6NdpProxyResolveCntDesc = prometheus.NewDesc(prefix+"icmp6_nd6_ndp_proxy_resolve_cnt", "Number of icmp6 nd6 ndp proxy resolve cnt", labelsICMP6, nil)
	icmp6Nd6DadProxyResolveCntDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_resolve", "Number of icmp nd6 dad proxy resolve", labelsICMP6, nil)
	icmp6Nd6DadProxyEqmacDropDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_eqmac_drop", "Number of icmp6 nd6 dad proxy eqmac drop", labelsICMP6, nil)
	icmp6Nd6DadProxyNomacDropDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_nomac_drop", "Number of icmp6 nd6 dad proxy nomac drop", labelsICMP6, nil)
	icmp6Nd6NdpProxyUnrRequestsDesc = prometheus.NewDesc(prefix+"icmp6_nd6_ndp_proxy_unr_requests", "Number of icmp6 nd6 ndp proxy unr requests", labelsICMP6, nil)
	icmp6Nd6DadProxyUnrRequestsDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_unr_requests", "Number of icmp6 nd6 dad proxy unr requests", labelsICMP6, nil)
	icmp6Nd6NdpProxyUnrResponsesDesc = prometheus.NewDesc(prefix+"icmp6_nd6_ndp_proxy_unr_responses", "Number of icmp6 nd6 ndp proxy unr responses", labelsICMP6, nil)
	icmp6Nd6DadProxyUnrConflictsDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_unr_conflicts", "Number of icmp6 nd6 dad proxy unr conflicts", labelsICMP6, nil)
	icmp6Nd6DadProxyUnrResponsesDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_unr_responses", "Number of icmp6 nd6 dad proxy unr responses", labelsICMP6, nil)
	icmp6Nd6NdpProxyUnrResolveCntDesc = prometheus.NewDesc(prefix+"icmp6_nd6_ndp_proxy_unr_resolve_cnt", "Number of icmp6 nd6 ndp proxy unr resolve cnt", labelsICMP6, nil)
	icmp6Nd6DadProxyUnrResolveCntDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_unr_resolve_cnt", "Number of icmp6 nd6 dad proxy unr resolve cnt", labelsICMP6, nil)
	icmp6Nd6DadProxyUnrEqportDropDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_unr_eqport_drop", "Number of icmp6 nd6 dad proxy unr eqport drop", labelsICMP6, nil)
	icmp6Nd6DadProxyUnrNomacDropDesc = prometheus.NewDesc(prefix+"icmp6_nd6_dad_proxy_unr_nomac_droop", "Number of icmp6 nd6 dad proxy unr nomac drop", labelsICMP6, nil)
	icmp6Nd6RequestsDroppedOnEntryDesc = prometheus.NewDesc(prefix+"icmp6_nd6_requests_dropped_on_entry", "Number of icmp6 nd6 requests dropped on entry", labelsICMP6, nil)
	icmp6Nd6RequestsDroppedDuringRetryDesc = prometheus.NewDesc(prefix+"icmp6_nd6_requests_dropped_during_retry", "Number of icmp6 nd6 requests dropped during retry", labelsICMP6, nil)

	labelsMPLS := []string{"target", "protocol"}
	mplsTotalMplsPacketsReceivedDesc = prometheus.NewDesc(prefix+"mpls_total_mpls_packets_received", "Number of mpls packets received", labelsMPLS, nil)
	mplsPacketsForwardedDesc = prometheus.NewDesc(prefix+"mpls_packets_forwarded", "Number of mpls packets forwarded", labelsMPLS, nil)
	mplsPacketsDroppedDesc = prometheus.NewDesc(prefix+"mpls_packets_dropped", "Number of mpls packets dropped", labelsMPLS, nil)
	mplsPacketsWithHeaderTooSmallDesc = prometheus.NewDesc(prefix+"mpls_packets_with_header_too_small", "Number of mpls packets with header too small", labelsMPLS, nil)
	mplsAfterTaggingPacketsCanNotFitLinkMtuDesc = prometheus.NewDesc(prefix+"mpls_after_tagging_packets_can_not_fit_link_mtu", "Number of mpls after tagging packets can not fit link mtu", labelsMPLS, nil)
	mplsPacketsWithIpv4ExplicitNullTagDesc = prometheus.NewDesc(prefix+"mpls_packets_with_ipv4_explicit_null_tag", "Number of mpls packets with ipv4 explicit null tag", labelsMPLS, nil)
	mplsPacketsWithIpv4ExplicitNullChecksumErrorsDesc = prometheus.NewDesc(prefix+"mpls_packets_with_ipv4_explicit_null_checksum_errors", "Number of mpls packets with ipv4 explicit null checksum errors", labelsMPLS, nil)
	mplsPacketsWithRouterAlertTagDesc = prometheus.NewDesc(prefix+"mpls_packets_with_router_alert_tag", "Number of mpls packets with router alert tag", labelsMPLS, nil)
	mplsLspPingPacketsDesc = prometheus.NewDesc(prefix+"mpls_lsp_ping_packets", "Number of mpls lsp ping packets", labelsMPLS, nil)
	mplsPacketsWithTtlExpiredDesc = prometheus.NewDesc(prefix+"mpls_packets_with_ttl_expired", "Number of mpls packets with ttl expired", labelsMPLS, nil)
	mplsPacketsWithTagEncodingErrorDesc = prometheus.NewDesc(prefix+"mpls_packets_with_tag_encoding_error", "Number of mpls packets with tag encoding error", labelsMPLS, nil)
	mplsPacketsDiscardedDueToNoRouteDesc = prometheus.NewDesc(prefix+"mpls_packets_discarded_due_to_no_route", "Number of mpls packets discarded due to no route", labelsMPLS, nil)
	mplsPacketsUsedFirstNexthopInEcmpUnilistDesc = prometheus.NewDesc(prefix+"mpls_packets_used_first_next_hop_in_ecmp_unilist", "Number of mpls packets used first nexthop in ecmp unilist", labelsMPLS, nil)
	mplsPacketsDroppedDueToIflDownDesc = prometheus.NewDesc(prefix+"mpls_packets_dropped_due_to_ifl_down", "Number of mpls packets dropped due to ifl down", labelsMPLS, nil)
	mplsPacketsDroppedAtMplsSocketSendDesc = prometheus.NewDesc(prefix+"mpls_packets_dropped_at_mpls_socket_send", "Number of mpls packets dropped at mpls socket send", labelsMPLS, nil)
	mplsPacketsForwardedAtMplsSocketSendDesc = prometheus.NewDesc(prefix+"mpls_packets_forwarded_at_mpls_socket_send", "Number of mpls packets forwarded at mpls socket send", labelsMPLS, nil)
	mplsPacketsDroppedAtP2mpCnhOutputDesc = prometheus.NewDesc(prefix+"mpls_packets_dropped_at_p2mp_cnh_output", "Number of mpls packets dropped at p2mp cnh output", labelsMPLS, nil)
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

	ch <- arpDatagramsReceivedDesc
	ch <- arpRequestsReceivedDesc
	ch <- arpRepliesReceivedDesc
	ch <- arpResolutionRequestReceivedDesc
	ch <- arpResolutionRequestDroppedDesc
	ch <- arpUnrestrictedProxyRequestsDesc
	ch <- arpRestrictedProxyRequestsDesc
	ch <- arpReceivedProxyRequestsDesc
	ch <- arpProxyRequestsNotProxiedDesc
	ch <- arpDatagramsWithBogusInterfaceDesc
	ch <- arpDatagramsWithIncorrectLengthDesc
	ch <- arpDatagramsForNonIpProtocolDesc
	ch <- arpDatagramsWithUnsupportedOpcodeDesc
	ch <- arpDatagramsWithBadProtocolAddressLengthDesc
	ch <- arpDatagramsWithBadHardwareAddressLengthDesc
	ch <- arpDatagramsWithMulticastSourceAddressDesc
	ch <- arpDatagramsWithMulticastTargetAddressDesc
	ch <- arpDatagramsWithMyOwnHardwareAddressDesc
	ch <- arpDatagramsForAnAddressNotOnTheInterfaceDesc
	ch <- arpDatagramsWithABroadcastSourceAddressDesc
	ch <- arpDatagramsWithSourceAddressDuplicateToMineDesc
	ch <- arpDatagramsWhichWereNotForMeDesc
	ch <- arpPacketsDiscardedWaitingForResolutionDesc
	ch <- arpPacketsSentAfterWaitingForResolutionDesc
	ch <- arpRequestsSentDesc
	ch <- arpRepliesSentDesc
	ch <- arpRequestsForMemoryDeniedDesc
	ch <- arpRequestsDroppedOnEntryDesc
	ch <- arpRequestsDroppedDuringRetryDesc
	ch <- arpRequestsDroppedDueToInterfaceDeletionDesc
	ch <- arpRequestsOnUnnumberedInterfacesDesc
	ch <- arpNewRequestsOnUnnumberedInterfacesDesc
	ch <- arpRepliesFromUnnumberedInterfacesDesc
	ch <- arpRequestsOnUnnumberedInterfaceWithNonSubnettedDonorDesc
	ch <- arpRepliesFromUnnumberedInterfaceWithNonSubnettedDonorDesc
	ch <- arpPacketsRejectedAsFamilyIsConfiguredWithDenyArpDesc
	ch <- arpResponsePacketsAreRejectedOnMcAeIclInterfaceDesc
	ch <- arpRepliesAreRejectedAsSourceAndDestinationIsSameDesc
	ch <- arpProbeForProxyAddressReachableFromTheIncomingInterfaceDesc
	ch <- arpRequestDiscardedForVrrpSourceAddressDesc
	ch <- arpSelfArpRequestPacketReceivedOnIrbInterfaceDesc
	ch <- arpProxyArpRequestDiscardedAsSourceIpIsAProxyTargetDesc
	ch <- arpPacketsAreDroppedAsNexthopAllocationFailedDesc
	ch <- arpPacketsReceivedFromPeerVrrpRouterAndDiscardedDesc
	ch <- arpPacketsAreRejectedAsTargetIpArpResolveIsInProgressDesc
	ch <- arpGratArpPacketsAreIgnoredAsMacAddressIsNotChangedDesc
	ch <- arpPacketsAreDroppedFromPeerVrrpDesc
	ch <- arpPacketsAreDroppedAsDriverCallFailedDesc
	ch <- arpPacketsAreDroppedAsSourceIsNotValidatedDesc
	ch <- arpSystemMaxDesc
	ch <- arpPublicMaxDesc
	ch <- arpIriMaxDesc
	ch <- arpMgtMaxDesc
	ch <- arpPublicCntDesc
	ch <- arpIriCntDesc
	ch <- arpMgtCntDesc
	ch <- arpSystemDropDesc
	ch <- arpPublicDropDesc
	ch <- arpIriDropDesc
	ch <- arpMgtDropDesc

	ch <- icmpDropsDueToRateLimitDesc
	ch <- icmpCallsToIcmpErrorDesc
	ch <- icmpErrorsNotGeneratedBecauseOldMessageWasIcmp
	ch <- icmpIcmpEchoReplyDesc
	ch <- icmpDestinationUnreachableDesc
	ch <- icmpIcmpEchoDesc
	ch <- icmpTimeStampReplyDesc
	ch <- icmpTimeExceededDesc
	ch <- icmpTimeStampDesc
	ch <- icmpAddressMaskRequestDesc
	ch <- icmpAnEndpointChangedItsCookieSecretDesc
	ch <- icmpMessagesWithBadCodeFieldsDesc
	ch <- icmpMessagesLessThanTheMinimumLengthDesc
	ch <- icmpMessagesWithBadChecksumDesc
	ch <- icmpMessagesWithBadSourceAddressDesc
	ch <- icmpMessagesWithBadLengthDesc
	ch <- icmpEchoDropsWithBroadcastOrMulticastDestinatonAddressDesc
	ch <- icmpTimestampDropsWithBroadcastOrMulticastDestinationAddressDesc
	ch <- icmpMessageResponsesGeneratedDesc
	ch <- icmp6CallsToIcmp6ErrorDesc
	ch <- icmp6ErrorsNotGeneratedBecauseOldMessageWasIcmpErrorDesc
	ch <- icmp6ErrorsNotGeneratedBecauseRateLimitationDesc
	ch <- icmp6UnreachableIcmp6PacketsOutputHistoDesc
	ch <- icmp6Icmp6EchoOutputHistoDesc
	ch <- icmp6Icmp6EchoReplyOutputHistoDesc
	ch <- icmp6NeighborSolicitationOutputHistoDesc
	ch <- icmp6NeighborAdvertisementOutputHistoDesc
	ch <- icmp6Icmp6MessagesWithBadCodeFieldsDesc
	ch <- icmp6MessagesLessThanMinimumLengthDesc
	ch <- icmp6BadChecksumsDesc
	ch <- icmp6Icmp6MessagesWithBadLengthDesc
	ch <- icmp6UnreachableIcmp6PacketInputHistosDesc
	ch <- icmp6PacketTooBigInputHistoDesc
	ch <- icmp6TimeExceededIcmp6PacketsInputHistoDesc
	ch <- icmp6Icmp6EchoInputHistoDesc
	ch <- icmp6Icmp6EchoReplyInputHistoDesc
	ch <- icmp6RouterSolicitationIcmp6PacketsInputHistoDesc
	ch <- icmp6NeighborSolicitationInputHistoDesc
	ch <- icmp6NeighborAdvertisementInputHistoDesc
	ch <- icmp6NoRouteDesc
	ch <- icmp6AdministrativelyProhibitedDesc
	ch <- icmp6BeyondScopeDesc
	ch <- icmp6AddressUnreachableDesc
	ch <- icmp6PortUnreachableDesc
	ch <- icmp6PacketTooBigDesc
	ch <- icmp6TimeExceedTransitDesc
	ch <- icmp6TimeExceedReassemblyDesc
	ch <- icmp6ErroneousHeaderFieldDesc
	ch <- icmp6UnrecognizedNextHeaderDesc
	ch <- icmp6UnrecognizedOptionDesc
	ch <- icmp6RedirectDesc
	ch <- icmp6UnknownDesc
	ch <- icmp6Icmp6MessageResponsesGeneratedDesc
	ch <- icmp6MessagesWithTooManyNdOptionsDesc
	ch <- icmp6NdSystemMaxDesc
	ch <- icmp6NdPublicMaxDesc
	ch <- icmp6NdIriMaxDesc
	ch <- icmp6NdMgtMaxDesc
	ch <- icmp6NdPublicCntDesc
	ch <- icmp6NdIriCntDesc
	ch <- icmp6NdMgtCntDesc
	ch <- icmp6NdSystemDropDesc
	ch <- icmp6NdPublicDropDesc
	ch <- icmp6NdIriDropDesc
	ch <- icmp6NdMgtDropDesc
	ch <- icmp6Nd6NdpProxyRequestsDesc
	ch <- icmp6Nd6DadProxyRequestsDesc
	ch <- icmp6Nd6NdpProxyResponsesDesc
	ch <- icmp6Nd6DadProxyConflictsDesc
	ch <- icmp6Nd6DupProxyResponsesDesc
	ch <- icmp6Nd6NdpProxyResolveCntDesc
	ch <- icmp6Nd6DadProxyResolveCntDesc
	ch <- icmp6Nd6DadProxyEqmacDropDesc
	ch <- icmp6Nd6DadProxyNomacDropDesc
	ch <- icmp6Nd6NdpProxyUnrRequestsDesc
	ch <- icmp6Nd6DadProxyUnrRequestsDesc
	ch <- icmp6Nd6NdpProxyUnrResponsesDesc
	ch <- icmp6Nd6DadProxyUnrConflictsDesc
	ch <- icmp6Nd6DadProxyUnrResponsesDesc
	ch <- icmp6Nd6NdpProxyUnrResolveCntDesc
	ch <- icmp6Nd6DadProxyUnrResolveCntDesc
	ch <- icmp6Nd6DadProxyUnrEqportDropDesc
	ch <- icmp6Nd6DadProxyUnrNomacDropDesc
	ch <- icmp6Nd6RequestsDroppedOnEntryDesc
	ch <- icmp6Nd6RequestsDroppedDuringRetryDesc

	ch <- mplsTotalMplsPacketsReceivedDesc
	ch <- mplsPacketsForwardedDesc
	ch <- mplsPacketsDroppedDesc
	ch <- mplsPacketsWithHeaderTooSmallDesc
	ch <- mplsAfterTaggingPacketsCanNotFitLinkMtuDesc
	ch <- mplsPacketsWithIpv4ExplicitNullTagDesc
	ch <- mplsPacketsWithIpv4ExplicitNullChecksumErrorsDesc
	ch <- mplsPacketsWithRouterAlertTagDesc
	ch <- mplsLspPingPacketsDesc
	ch <- mplsPacketsWithTtlExpiredDesc
	ch <- mplsPacketsWithTagEncodingErrorDesc
	ch <- mplsPacketsDiscardedDueToNoRouteDesc
	ch <- mplsPacketsUsedFirstNexthopInEcmpUnilistDesc
	ch <- mplsPacketsDroppedDueToIflDownDesc
	ch <- mplsPacketsDroppedAtMplsSocketSendDesc
	ch <- mplsPacketsForwardedAtMplsSocketSendDesc
	ch <- mplsPacketsDroppedAtP2mpCnhOutputDesc
}

func (c *systemstatisticsCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var s SystemStatistics
	err := client.RunCommandAndParse("show system statistics ip", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsIPV4(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics ip6", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsIPV6(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics udp", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsUDP(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics tcp", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsTCP(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics arp", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsARP(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics icmp", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsICMP(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics icmp6", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsICMP6(ch, labelValues, s)

	err = client.RunCommandAndParse("show system statistics mpls", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsMPLS(ch, labelValues, s)
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

func (c *systemstatisticsCollector) collectSystemStatisticsARP(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	labels := append(labelValues, "ARP")
	ch <- prometheus.MustNewConstMetric(arpDatagramsReceivedDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsReceivedDesc, prometheus.CounterValue, s.Statistics.Arp.ArpRequestsSent, labels...)
	ch <- prometheus.MustNewConstMetric(arpRepliesReceivedDesc, prometheus.CounterValue, s.Statistics.Arp.ArpRepliesReceived, labels...)
	ch <- prometheus.MustNewConstMetric(arpResolutionRequestReceivedDesc, prometheus.CounterValue, s.Statistics.Arp.ResolutionRequestReceived, labels...)
	ch <- prometheus.MustNewConstMetric(arpResolutionRequestDroppedDesc, prometheus.CounterValue, s.Statistics.Arp.ResolutionRequestDropped, labels...)
	ch <- prometheus.MustNewConstMetric(arpUnrestrictedProxyRequestsDesc, prometheus.CounterValue, s.Statistics.Arp.UnrestrictedProxyRequests, labels...)
	ch <- prometheus.MustNewConstMetric(arpRestrictedProxyRequestsDesc, prometheus.CounterValue, s.Statistics.Arp.RestrictedProxyRequests, labels...)
	ch <- prometheus.MustNewConstMetric(arpReceivedProxyRequestsDesc, prometheus.CounterValue, s.Statistics.Arp.ReceivedProxyRequests, labels...)
	ch <- prometheus.MustNewConstMetric(arpProxyRequestsNotProxiedDesc, prometheus.CounterValue, s.Statistics.Arp.ProxyRequestsNotProxied, labels...)
	ch <- prometheus.MustNewConstMetric(arpRestrictedProxyRequestsNotProxiedDesc, prometheus.CounterValue, s.Statistics.Arp.RestrictedProxyRequestsNotProxied, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithBogusInterfaceDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithBogusInterface, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithIncorrectLengthDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithIncorrectLength, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsForNonIpProtocolDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsForNonIpProtocol, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithUnsupportedOpcodeDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithUnsupportedOpcode, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithBadProtocolAddressLengthDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithBadProtocolAddressLength, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithBadHardwareAddressLengthDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithBadHardwareAddressLength, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithMulticastSourceAddressDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithMulticastSourceAddress, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithMulticastTargetAddressDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithMulticastTargetAddress, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithMyOwnHardwareAddressDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithMyOwnHardwareAddress, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsForAnAddressNotOnTheInterfaceDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsForAnAddressNotOnTheInterface, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithABroadcastSourceAddressDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithABroadcastSourceAddress, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWithSourceAddressDuplicateToMineDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWithSourceAddressDuplicateToMine, labels...)
	ch <- prometheus.MustNewConstMetric(arpDatagramsWhichWereNotForMeDesc, prometheus.CounterValue, s.Statistics.Arp.DatagramsWhichWereNotForMe, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsDiscardedWaitingForResolutionDesc, prometheus.CounterValue, s.Statistics.Arp.PacketsDiscardedWaitingForResolution, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsSentAfterWaitingForResolutionDesc, prometheus.CounterValue, s.Statistics.Arp.PacketsSentAfterWaitingForResolution, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsSentDesc, prometheus.CounterValue, s.Statistics.Arp.ArpRequestsSent, labels...)
	ch <- prometheus.MustNewConstMetric(arpRepliesSentDesc, prometheus.CounterValue, s.Statistics.Arp.ArpRepliesSent, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsForMemoryDeniedDesc, prometheus.CounterValue, s.Statistics.Arp.RequestsForMemoryDenied, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsDroppedOnEntryDesc, prometheus.CounterValue, s.Statistics.Arp.RequestsDroppedOnEntry, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsDroppedDuringRetryDesc, prometheus.CounterValue, s.Statistics.Arp.RequestsDroppedDuringRetry, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsDroppedDueToInterfaceDeletionDesc, prometheus.CounterValue, s.Statistics.Arp.RequestsDroppedDueToInterfaceDeletion, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsOnUnnumberedInterfacesDesc, prometheus.CounterValue, s.Statistics.Arp.RequestsOnUnnumberedInterfaces, labels...)
	ch <- prometheus.MustNewConstMetric(arpNewRequestsOnUnnumberedInterfacesDesc, prometheus.CounterValue, s.Statistics.Arp.NewRequestsOnUnnumberedInterfaces, labels...)
	ch <- prometheus.MustNewConstMetric(arpRepliesFromUnnumberedInterfacesDesc, prometheus.CounterValue, s.Statistics.Arp.RepliesFromUnnumberedInterfaces, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestsOnUnnumberedInterfaceWithNonSubnettedDonorDesc, prometheus.CounterValue, s.Statistics.Arp.RequestsOnUnnumberedInterfaceWithNonSubnettedDonor, labels...)
	ch <- prometheus.MustNewConstMetric(arpRepliesFromUnnumberedInterfaceWithNonSubnettedDonorDesc, prometheus.CounterValue, s.Statistics.Arp.RepliesFromUnnumberedInterfaceWithNonSubnettedDonor, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsRejectedAsFamilyIsConfiguredWithDenyArpDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsRejectedAsFamilyIsConfiguredWithDenyArp, labels...)
	ch <- prometheus.MustNewConstMetric(arpResponsePacketsAreRejectedOnMcAeIclInterfaceDesc, prometheus.CounterValue, s.Statistics.Arp.ArpResponsePacketsAreRejectedOnMcAeIclInterface, labels...)
	ch <- prometheus.MustNewConstMetric(arpRepliesAreRejectedAsSourceAndDestinationIsSameDesc, prometheus.CounterValue, s.Statistics.Arp.ArpRepliesAreRejectedAsSourceAndDestinationIsSame, labels...)
	ch <- prometheus.MustNewConstMetric(arpProbeForProxyAddressReachableFromTheIncomingInterfaceDesc, prometheus.CounterValue, s.Statistics.Arp.ArpProbeForProxyAddressReachableFromTheIncomingInterface, labels...)
	ch <- prometheus.MustNewConstMetric(arpRequestDiscardedForVrrpSourceAddressDesc, prometheus.CounterValue, s.Statistics.Arp.ArpRequestDiscardedForVrrpSourceAddress, labels...)
	ch <- prometheus.MustNewConstMetric(arpSelfArpRequestPacketReceivedOnIrbInterfaceDesc, prometheus.CounterValue, s.Statistics.Arp.SelfArpRequestPacketReceivedOnIrbInterface, labels...)
	ch <- prometheus.MustNewConstMetric(arpProxyArpRequestDiscardedAsSourceIpIsAProxyTargetDesc, prometheus.CounterValue, s.Statistics.Arp.ProxyArpRequestDiscardedAsSourceIpIsAProxyTarget, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsAreDroppedAsNexthopAllocationFailedDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsAreDroppedAsNexthopAllocationFailed, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsReceivedFromPeerVrrpRouterAndDiscardedDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsReceivedFromPeerVrrpRouterAndDiscarded, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsAreRejectedAsTargetIpArpResolveIsInProgressDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsAreRejectedAsTargetIpArpResolveIsInProgress, labels...)
	ch <- prometheus.MustNewConstMetric(arpGratArpPacketsAreIgnoredAsMacAddressIsNotChangedDesc, prometheus.CounterValue, s.Statistics.Arp.GratArpPacketsAreIgnoredAsMacAddressIsNotChanged, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsAreDroppedFromPeerVrrpDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsAreDroppedFromPeerVrrp, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsAreDroppedAsDriverCallFailedDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsAreDroppedAsDriverCallFailed, labels...)
	ch <- prometheus.MustNewConstMetric(arpPacketsAreDroppedAsSourceIsNotValidatedDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPacketsAreDroppedAsSourceIsNotValidated, labels...)
	ch <- prometheus.MustNewConstMetric(arpSystemMaxDesc, prometheus.CounterValue, s.Statistics.Arp.ArpSystemMax, labels...)
	ch <- prometheus.MustNewConstMetric(arpPublicMaxDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPublicMax, labels...)
	ch <- prometheus.MustNewConstMetric(arpIriMaxDesc, prometheus.CounterValue, s.Statistics.Arp.ArpIriMax, labels...)
	ch <- prometheus.MustNewConstMetric(arpMgtMaxDesc, prometheus.CounterValue, s.Statistics.Arp.ArpMgtMax, labels...)
	ch <- prometheus.MustNewConstMetric(arpPublicCntDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPublicCnt, labels...)
	ch <- prometheus.MustNewConstMetric(arpIriCntDesc, prometheus.CounterValue, s.Statistics.Arp.ArpIriCnt, labels...)
	ch <- prometheus.MustNewConstMetric(arpMgtCntDesc, prometheus.CounterValue, s.Statistics.Arp.ArpMgtCnt, labels...)
	ch <- prometheus.MustNewConstMetric(arpSystemDropDesc, prometheus.CounterValue, s.Statistics.Arp.ArpSystemDrop, labels...)
	ch <- prometheus.MustNewConstMetric(arpPublicDropDesc, prometheus.CounterValue, s.Statistics.Arp.ArpPublicDrop, labels...)
	ch <- prometheus.MustNewConstMetric(arpIriDropDesc, prometheus.CounterValue, s.Statistics.Arp.ArpIriDrop, labels...)
	ch <- prometheus.MustNewConstMetric(arpMgtDropDesc, prometheus.CounterValue, s.Statistics.Arp.ArpMgtDrop, labels...)
}

func (c *systemstatisticsCollector) collectSystemStatisticsICMP(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	labels := append(labelValues, "ICMP")
	ch <- prometheus.MustNewConstMetric(icmpDropsDueToRateLimitDesc, prometheus.CounterValue, s.Statistics.Icmp.DropsDueToRateLimit, labels...)
	ch <- prometheus.MustNewConstMetric(icmpCallsToIcmpErrorDesc, prometheus.CounterValue, s.Statistics.Icmp.CallsToIcmpError, labels...)
	ch <- prometheus.MustNewConstMetric(icmpErrorsNotGeneratedBecauseOldMessageWasIcmp, prometheus.CounterValue, s.Statistics.Icmp.ErrorsNotGeneratedBecauseOldMessageWasIcmp, labels...)
	for _, histogram := range s.Statistics.Icmp.Histogram {
		labels := append(labelValues, "ICMP", histogram.TypeOfHistogram)
		ch <- prometheus.MustNewConstMetric(icmpIcmpEchoReplyDesc, prometheus.CounterValue, histogram.IcmpEchoReply, labels...)
		ch <- prometheus.MustNewConstMetric(icmpDestinationUnreachableDesc, prometheus.CounterValue, histogram.DestinationUnreachable, labels...)
		ch <- prometheus.MustNewConstMetric(icmpIcmpEchoDesc, prometheus.CounterValue, histogram.IcmpEcho, labels...)
		ch <- prometheus.MustNewConstMetric(icmpTimeStampReplyDesc, prometheus.CounterValue, histogram.TimeStampReply, labels...)
		ch <- prometheus.MustNewConstMetric(icmpTimeExceededDesc, prometheus.CounterValue, histogram.TimeExceeded, labels...)
		ch <- prometheus.MustNewConstMetric(icmpTimeStampDesc, prometheus.CounterValue, histogram.TimeStamp, labels...)
		ch <- prometheus.MustNewConstMetric(icmpAddressMaskRequestDesc, prometheus.CounterValue, histogram.AddressMaskRequest, labels...)
		ch <- prometheus.MustNewConstMetric(icmpAnEndpointChangedItsCookieSecretDesc, prometheus.CounterValue, histogram.AnEndpointChangedItsCookieSecret, labels...)
	}
	ch <- prometheus.MustNewConstMetric(icmpMessagesWithBadCodeFieldsDesc, prometheus.CounterValue, s.Statistics.Icmp.MessagesWithBadCodeFields, labels...)
	ch <- prometheus.MustNewConstMetric(icmpMessagesLessThanTheMinimumLengthDesc, prometheus.CounterValue, s.Statistics.Icmp.MessagesLessThanTheMinimumLength, labels...)
	ch <- prometheus.MustNewConstMetric(icmpMessagesWithBadChecksumDesc, prometheus.CounterValue, s.Statistics.Icmp.MessagesWithBadChecksum, labels...)
	ch <- prometheus.MustNewConstMetric(icmpMessagesWithBadSourceAddressDesc, prometheus.CounterValue, s.Statistics.Icmp.MessagesWithBadSourceAddress, labels...)
	ch <- prometheus.MustNewConstMetric(icmpMessagesWithBadLengthDesc, prometheus.CounterValue, s.Statistics.Icmp.MessagesWithBadLength, labels...)
	ch <- prometheus.MustNewConstMetric(icmpEchoDropsWithBroadcastOrMulticastDestinatonAddressDesc, prometheus.CounterValue, s.Statistics.Icmp.EchoDropsWithBroadcastOrMulticastDestinatonAddress, labels...)
	ch <- prometheus.MustNewConstMetric(icmpTimestampDropsWithBroadcastOrMulticastDestinationAddressDesc, prometheus.CounterValue, s.Statistics.Icmp.TimestampDropsWithBroadcastOrMulticastDestinationAddress, labels...)
	ch <- prometheus.MustNewConstMetric(icmpMessageResponsesGeneratedDesc, prometheus.CounterValue, s.Statistics.Icmp.MessageResponsesGenerated, labels...)
}

func (c *systemstatisticsCollector) collectSystemStatisticsICMP6(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	labels := append(labelValues, "ICMP6")
	ch <- prometheus.MustNewConstMetric(icmp6CallsToIcmp6ErrorDesc, prometheus.CounterValue, s.Statistics.Icmp6.CallsToIcmp6Error, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6ErrorsNotGeneratedBecauseOldMessageWasIcmpErrorDesc, prometheus.CounterValue, s.Statistics.Icmp6.ErrorsNotGeneratedBecauseOldMessageWasIcmpError, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6ErrorsNotGeneratedBecauseRateLimitationDesc, prometheus.CounterValue, s.Statistics.Icmp6.ErrorsNotGeneratedBecauseRateLimitation, labels...)
	labels = append(labels, "Output Histogram")
	ch <- prometheus.MustNewConstMetric(icmp6UnreachableIcmp6PacketsOutputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.OutputHistogram.UnreachableIcmp6Packets, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6EchoOutputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.OutputHistogram.Icmp6Echo, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6EchoReplyOutputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.OutputHistogram.Icmp6EchoReply, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NeighborSolicitationOutputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.OutputHistogram.NeighborSolicitation, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NeighborAdvertisementOutputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.OutputHistogram.NeighborAdvertisement, labels...)
	labels = labels[:len(labels)-1]
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6MessagesWithBadCodeFieldsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Icmp6MessagesWithBadCodeFields, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6MessagesLessThanMinimumLengthDesc, prometheus.CounterValue, s.Statistics.Icmp6.MessagesLessThanMinimumLength, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6BadChecksumsDesc, prometheus.CounterValue, s.Statistics.Icmp6.BadChecksums, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6MessagesWithBadLengthDesc, prometheus.CounterValue, s.Statistics.Icmp6.Icmp6MessagesWithBadLength, labels...)
	labels = append(labels, "Input Histogram")
	ch <- prometheus.MustNewConstMetric(icmp6UnreachableIcmp6PacketInputHistosDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.UnreachableIcmp6Packets, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6PacketTooBigInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.PacketTooBig, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6TimeExceededIcmp6PacketsInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.TimeExceededIcmp6Packets, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6EchoInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.Icmp6EchoReply, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6EchoReplyInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.Icmp6EchoReply, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6RouterSolicitationIcmp6PacketsInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.RouterSolicitationIcmp6Packets, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NeighborSolicitationInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.NeighborSolicitation, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NeighborAdvertisementInputHistoDesc, prometheus.CounterValue, s.Statistics.Icmp6.InputHistogram.NeighborAdvertisement, labels...)
	labels = labels[:len(labels)-1]
	ch <- prometheus.MustNewConstMetric(icmp6NoRouteDesc, prometheus.CounterValue, s.Statistics.Icmp6.NoRoute, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6AdministrativelyProhibitedDesc, prometheus.CounterValue, s.Statistics.Icmp6.AdministrativelyProhibited, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6BeyondScopeDesc, prometheus.CounterValue, s.Statistics.Icmp6.BeyondScope, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6AddressUnreachableDesc, prometheus.CounterValue, s.Statistics.Icmp6.AddressUnreachable, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6PortUnreachableDesc, prometheus.CounterValue, s.Statistics.Icmp6.PortUnreachable, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6PacketTooBigDesc, prometheus.CounterValue, s.Statistics.Icmp6.PacketTooBig, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6TimeExceedTransitDesc, prometheus.CounterValue, s.Statistics.Icmp6.TimeExceedTransit, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6TimeExceedReassemblyDesc, prometheus.CounterValue, s.Statistics.Icmp6.TimeExceedReassembly, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6ErroneousHeaderFieldDesc, prometheus.CounterValue, s.Statistics.Icmp6.ErroneousHeaderField, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6UnrecognizedNextHeaderDesc, prometheus.CounterValue, s.Statistics.Icmp6.UnrecognizedNextHeader, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6UnrecognizedOptionDesc, prometheus.CounterValue, s.Statistics.Icmp6.UnrecognizedOption, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6RedirectDesc, prometheus.CounterValue, s.Statistics.Icmp6.Redirect, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6UnknownDesc, prometheus.CounterValue, s.Statistics.Icmp6.Unknown, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Icmp6MessageResponsesGeneratedDesc, prometheus.CounterValue, s.Statistics.Icmp6.Icmp6MessageResponsesGenerated, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6MessagesWithTooManyNdOptionsDesc, prometheus.CounterValue, s.Statistics.Icmp6.MessagesWithTooManyNdOptions, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdSystemMaxDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdSystemMax, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdPublicMaxDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdPublicMax, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdIriMaxDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdIriMax, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdMgtMaxDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdMgtMax, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdPublicCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdPublicCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdIriCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdIriCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdMgtCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdMgtCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdSystemDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdSystemDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdPublicDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdPublicDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdIriDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdIriDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6NdMgtDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.NdMgtDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6NdpProxyRequestsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6NdpProxyRequests, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyRequestsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyRequests, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6NdpProxyResponsesDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6NdpProxyResponses, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyConflictsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyConflicts, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DupProxyResponsesDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DupProxyResponses, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6NdpProxyResolveCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6NdpProxyResolveCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyResolveCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyResolveCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyEqmacDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyEqmacDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyNomacDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyNomacDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6NdpProxyUnrRequestsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6NdpProxyUnrRequests, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyUnrRequestsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyUnrRequests, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6NdpProxyUnrResponsesDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6NdpProxyUnrResponses, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyUnrConflictsDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyUnrConflicts, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyUnrResponsesDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyUnrResponses, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6NdpProxyUnrResolveCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6NdpProxyUnrResolveCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyUnrResolveCntDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyUnrResolveCnt, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6DadProxyUnrNomacDropDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6DadProxyUnrNomacDrop, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6RequestsDroppedOnEntryDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6RequestsDroppedOnEntry, labels...)
	ch <- prometheus.MustNewConstMetric(icmp6Nd6RequestsDroppedDuringRetryDesc, prometheus.CounterValue, s.Statistics.Icmp6.Nd6RequestsDroppedDuringRetry, labels...)
}

func (c *systemstatisticsCollector) collectSystemStatisticsMPLS(ch chan<- prometheus.Metric, labelValues []string, s SystemStatistics) {
	labels := append(labelValues, "MPLS")
	ch <- prometheus.MustNewConstMetric(mplsTotalMplsPacketsReceivedDesc, prometheus.CounterValue, s.Statistics.Mpls.TotalMplsPacketsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsForwardedDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsForwarded, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsWithHeaderTooSmallDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsWithHeaderTooSmall, labels...)
	ch <- prometheus.MustNewConstMetric(mplsAfterTaggingPacketsCanNotFitLinkMtuDesc, prometheus.CounterValue, s.Statistics.Mpls.AfterTaggingPacketsCanNotFitLinkMtu, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsWithIpv4ExplicitNullTagDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsWithIpv4ExplicitNullTag, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsWithIpv4ExplicitNullChecksumErrorsDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsWithIpv4ExplicitNullChecksumErrors, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsWithRouterAlertTagDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsWithRouterAlertTag, labels...)
	ch <- prometheus.MustNewConstMetric(mplsLspPingPacketsDesc, prometheus.CounterValue, s.Statistics.Mpls.LspPingPackets, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsWithTtlExpiredDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsWithTtlExpired, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsWithTagEncodingErrorDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsWithTagEncodingError, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsDiscardedDueToNoRouteDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsDiscardedDueToNoRoute, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsUsedFirstNexthopInEcmpUnilistDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsUsedFirstNexthopInEcmpUnilist, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsDroppedDueToIflDownDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsDroppedDueToIflDown, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsDroppedAtMplsSocketSendDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsDroppedAtMplsSocketSend, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsForwardedAtMplsSocketSendDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsForwardedAtMplsSocketSend, labels...)
	ch <- prometheus.MustNewConstMetric(mplsPacketsDroppedAtP2mpCnhOutputDesc, prometheus.CounterValue, s.Statistics.Mpls.PacketsDroppedAtP2mpCnhOutput, labels...)
}
