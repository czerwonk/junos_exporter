package systemstatisticsIPv6

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

const prefix string = "junos_systemstatistics_ipv6_"

// Metrics to collect for the feature
var (
	totalPacketsReceivedDesc                  *prometheus.Desc
	ip6PacketsWithSizeSmallerThanMinimumDesc  *prometheus.Desc
	packetsWithDatasizeLessThanDataLengthDesc *prometheus.Desc
	ip6PacketsWithBadOptionsDesc              *prometheus.Desc
	ip6PacketsWithIncorrectVersionNumberDesc  *prometheus.Desc
	ip6FragmentsReceivedDesc                  *prometheus.Desc
	duplicateOrOutOfSpaceFragmentsDroppedDesc *prometheus.Desc
	ip6FragmentsDroppedAfterTimeoutDesc       *prometheus.Desc
	fragmentsThatExceededLimitDesc            *prometheus.Desc
	ip6PacketsReassembledOkDesc               *prometheus.Desc
	ip6PacketsForThisHostDesc                 *prometheus.Desc
	ip6PacketsForwardedDesc                   *prometheus.Desc
	ip6PacketsNotForwardableDesc              *prometheus.Desc
	ip6RedirectsSentDesc                      *prometheus.Desc
	ip6PacketsSentFromThisHostDesc            *prometheus.Desc
	ip6PacketsSentWithFabricatedIpHeaderDesc  *prometheus.Desc
	ip6OutputPacketsDroppedDueToNoBufsDesc    *prometheus.Desc
	ip6OutputPacketsDiscardedDueToNoRouteDesc *prometheus.Desc
	ip6OutputDatagramsFragmentedDesc          *prometheus.Desc
	ip6FragmentsCreatedDesc                   *prometheus.Desc
	ip6DatagramsThatCanNotBeFragmentedDesc    *prometheus.Desc
	packetsThatViolatedScopeRulesDesc         *prometheus.Desc
	multicastPacketsWhichWeDoNotJoinDesc      *prometheus.Desc
	ip6nhTcpDesc                              *prometheus.Desc
	ip6nhUdpDesc                              *prometheus.Desc
	ip6nhIcmp6Desc                            *prometheus.Desc
	packetsWhoseHeadersAreNotContinuousDesc   *prometheus.Desc
	tunnelingPacketsThatCanNotFindGifDesc     *prometheus.Desc
	packetsDiscardedDueToTooMayHeadersDesc    *prometheus.Desc
	failuresOfSourceAddressSelectionDesc      *prometheus.Desc
	headerTypeLinkLocalsDesc                  *prometheus.Desc
	headerTypeGlobalsDesc                     *prometheus.Desc
	forwardCacheHitDesc                       *prometheus.Desc
	forwardCacheMissDesc                      *prometheus.Desc
	ip6PacketsDestinedToDeadNextHopDesc       *prometheus.Desc
	ip6OptionPacketsDroppedDueToRateLimitDesc *prometheus.Desc
	ip6PacketsDroppedDesc                     *prometheus.Desc
	packetsDroppedDueToBadProtocolDesc        *prometheus.Desc
	transitRePacketDroppedOnMgmtInterfaceDesc *prometheus.Desc
	packetUsedFirstNexthopInEcmpUnilistDesc   *prometheus.Desc
)

func init() {
	labels := []string{"target", "protocol"}
	totalPacketsReceivedDesc = prometheus.NewDesc(prefix+"total_packets_received", "Total number of packets received", labels, nil)
	ip6PacketsWithSizeSmallerThanMinimumDesc = prometheus.NewDesc(prefix+"ip6_packets_with_size_smaller_than_minimum", "Number of packets received with size smaller than minimum", labels, nil)
	packetsWithDatasizeLessThanDataLengthDesc = prometheus.NewDesc(prefix+"packets_with_datasize_less_than_data_length", "Number of packets received with data length less than data length", labels, nil)
	ip6PacketsWithBadOptionsDesc = prometheus.NewDesc(prefix+"ip6_packets_with_bad_options", "Number of packets received with bad options", labels, nil)
	ip6PacketsWithIncorrectVersionNumberDesc = prometheus.NewDesc(prefix+"ip6_packets_with_incorrect_version_number", "Number of packets received with incorrect version number", labels, nil)
	ip6FragmentsReceivedDesc = prometheus.NewDesc(prefix+"ip6_fragments_received", "Number of fragments received", labels, nil)
	duplicateOrOutOfSpaceFragmentsDroppedDesc = prometheus.NewDesc(prefix+"duplicate_or_out_of_space_fragments_dropped", "Number of duplicate or out of space fragments dropped", labels, nil)
	ip6FragmentsDroppedAfterTimeoutDesc = prometheus.NewDesc(prefix+"ip6_fragments_dropped_after_timeout", "Number of fragments dropped after timeout", labels, nil)
	fragmentsThatExceededLimitDesc = prometheus.NewDesc(prefix+"fragments_that_exceeded_limit", "Number of fragments that exceeded limit", labels, nil)
	ip6PacketsReassembledOkDesc = prometheus.NewDesc(prefix+"ip6_packets_reassembled_ok", "Number of packets reassembled ok", labels, nil)
	ip6PacketsForThisHostDesc = prometheus.NewDesc(prefix+"ip6_packets_for_this_host", "Number of packets for this host", labels, nil)
	ip6PacketsForwardedDesc = prometheus.NewDesc(prefix+"ip6_packets_forwarded", "Number of packets forwarded", labels, nil)
	ip6PacketsNotForwardableDesc = prometheus.NewDesc(prefix+"ip6_packets_not_forwardable", "Number of packets not forwardable", labels, nil)
	ip6RedirectsSentDesc = prometheus.NewDesc(prefix+"ip6_redirects_sent", "Number of redirects sent", labels, nil)
	ip6PacketsSentFromThisHostDesc = prometheus.NewDesc(prefix+"ip6_packets_sent_from_this_host", "Number of packets sent from this host", labels, nil)
	ip6PacketsSentWithFabricatedIpHeaderDesc = prometheus.NewDesc(prefix+"ip6_packets_sent_with_fabricated_ip_header", "Number of packets sent with fabricated ip header", labels, nil)
	ip6OutputPacketsDroppedDueToNoBufsDesc = prometheus.NewDesc(prefix+"ip6_output_packets_dropped_due_to_no_bufs", "Number of output packets dropped due to no bufs", labels, nil)
	ip6OutputPacketsDiscardedDueToNoRouteDesc = prometheus.NewDesc(prefix+"ip6_output_packets_discarded_due_to_no_route", "Number of output packets discarded due to no route", labels, nil)
	ip6OutputDatagramsFragmentedDesc = prometheus.NewDesc(prefix+"ip6_output_datagrams_fragmented", "Number of output datagrams fragmented", labels, nil)
	ip6FragmentsCreatedDesc = prometheus.NewDesc(prefix+"ip6_fragments_created", "Number of fragments created", labels, nil)
	ip6DatagramsThatCanNotBeFragmentedDesc = prometheus.NewDesc(prefix+"ip6_datagrams_that_can_not_be_fragmented", "Number of datagrams that can not be fragmented", labels, nil)
	packetsThatViolatedScopeRulesDesc = prometheus.NewDesc(prefix+"packets_that_violated_scope_rules", "Number of packets that violated scope rules", labels, nil)
	multicastPacketsWhichWeDoNotJoinDesc = prometheus.NewDesc(prefix+"multicast_packets_which_we_do_not_join", "Number of multicast packets which we do not join", labels, nil)
	ip6nhTcpDesc = prometheus.NewDesc(prefix+"ip6nh_tcp", "Number of packets with next header tcp", labels, nil)
	ip6nhUdpDesc = prometheus.NewDesc(prefix+"ip6nh_udp", "Number of packets with next header udp", labels, nil)
	ip6nhIcmp6Desc = prometheus.NewDesc(prefix+"ip6nh_icmp6", "Number of packets with next header icmp6", labels, nil)
	packetsWhoseHeadersAreNotContinuousDesc = prometheus.NewDesc(prefix+"packets_whose_headers_are_not_continuous", "Number of packets whose headers are not continuous", labels, nil)
	tunnelingPacketsThatCanNotFindGifDesc = prometheus.NewDesc(prefix+"tunneling_packets_that_can_not_find_gif", "Number of tunneling packets that can not find gif", labels, nil)
	packetsDiscardedDueToTooMayHeadersDesc = prometheus.NewDesc(prefix+"packets_discarded_due_to_too_may_headers", "Number of packets discarded due to too may headers", labels, nil)
	failuresOfSourceAddressSelectionDesc = prometheus.NewDesc(prefix+"failures_of_source_address_selection", "Number of failures of source address selection", labels, nil)
	l := []string{"target", "protocol", "header_type"}
	headerTypeLinkLocalsDesc = prometheus.NewDesc(prefix+"header_type_link_locals", "Number of packets with header type link locals", l, nil)
	headerTypeGlobalsDesc = prometheus.NewDesc(prefix+"header_type_globals", "Number of packets with header type globals", l, nil)
	forwardCacheHitDesc = prometheus.NewDesc(prefix+"forward_cache_hit", "Number of forward cache hit", labels, nil)
	forwardCacheMissDesc = prometheus.NewDesc(prefix+"forward_cache_miss", "Number of forward cache miss", labels, nil)
	ip6PacketsDestinedToDeadNextHopDesc = prometheus.NewDesc(prefix+"ip6_packets_destined_to_dead_next_hop", "Number of packets destined to dead next hop", labels, nil)
	ip6OptionPacketsDroppedDueToRateLimitDesc = prometheus.NewDesc(prefix+"ip6_option_packets_dropped_due_to_rate_limit", "Number of option packets dropped due to rate limit", labels, nil)
	ip6PacketsDroppedDesc = prometheus.NewDesc(prefix+"ip6_packets_dropped", "Number of packets dropped", labels, nil)
	packetsDroppedDueToBadProtocolDesc = prometheus.NewDesc(prefix+"packets_dropped_due_to_bad_protocol", "Number of packets dropped due to bad protocol", labels, nil)
	transitRePacketDroppedOnMgmtInterfaceDesc = prometheus.NewDesc(prefix+"transit_re_packet_dropped_on_mgmt_interface", "Number of transit re packet dropped on mgmt interface", labels, nil)
	packetUsedFirstNexthopInEcmpUnilistDesc = prometheus.NewDesc(prefix+"packet_used_first_nexthop_in_ecmp_unilist", "Number of packet used first nexthop in ecmp unilist", labels, nil)
}

type systemstatisticsIPv6Collector struct{}

func NewCollector() collector.RPCCollector {
	return &systemstatisticsIPv6Collector{}
}

func (c *systemstatisticsIPv6Collector) Name() string {
	return "systemstatisticsIPv6"
}

func (c *systemstatisticsIPv6Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- totalPacketsReceivedDesc
	ch <- ip6PacketsWithSizeSmallerThanMinimumDesc
	ch <- packetsWithDatasizeLessThanDataLengthDesc
	ch <- ip6PacketsWithBadOptionsDesc
	ch <- ip6PacketsWithIncorrectVersionNumberDesc
	ch <- ip6FragmentsReceivedDesc
	ch <- duplicateOrOutOfSpaceFragmentsDroppedDesc
	ch <- ip6FragmentsDroppedAfterTimeoutDesc
	ch <- fragmentsThatExceededLimitDesc
	ch <- ip6PacketsReassembledOkDesc
	ch <- ip6PacketsForThisHostDesc
	ch <- ip6PacketsForwardedDesc
	ch <- ip6PacketsNotForwardableDesc
	ch <- ip6RedirectsSentDesc
	ch <- ip6PacketsSentFromThisHostDesc
	ch <- ip6PacketsSentWithFabricatedIpHeaderDesc
	ch <- ip6OutputPacketsDroppedDueToNoBufsDesc
	ch <- ip6OutputPacketsDiscardedDueToNoRouteDesc
	ch <- ip6OutputDatagramsFragmentedDesc
	ch <- ip6FragmentsCreatedDesc
	ch <- ip6DatagramsThatCanNotBeFragmentedDesc
	ch <- packetsThatViolatedScopeRulesDesc
	ch <- multicastPacketsWhichWeDoNotJoinDesc
	ch <- ip6nhTcpDesc
	ch <- ip6nhUdpDesc
	ch <- ip6nhIcmp6Desc
	ch <- packetsWhoseHeadersAreNotContinuousDesc
	ch <- tunnelingPacketsThatCanNotFindGifDesc
	ch <- packetsDiscardedDueToTooMayHeadersDesc
	ch <- failuresOfSourceAddressSelectionDesc
	ch <- headerTypeLinkLocalsDesc
	ch <- headerTypeGlobalsDesc
	ch <- forwardCacheHitDesc
	ch <- forwardCacheMissDesc
	ch <- ip6PacketsDestinedToDeadNextHopDesc
	ch <- ip6OptionPacketsDroppedDueToRateLimitDesc
	ch <- ip6PacketsDroppedDesc
	ch <- packetsDroppedDueToBadProtocolDesc
	ch <- transitRePacketDroppedOnMgmtInterfaceDesc
	ch <- packetUsedFirstNexthopInEcmpUnilistDesc
}

func (c *systemstatisticsIPv6Collector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var s StatisticsIPv6
	err := client.RunCommandAndParse("show system statistics ip6", &s)
	if err != nil {
		return err
	}
	c.collectSystemStatisticsIPv6(ch, labelValues, s)
	return nil
}

func (c *systemstatisticsIPv6Collector) collectSystemStatisticsIPv6(ch chan<- prometheus.Metric, labelValues []string, s StatisticsIPv6) {
	labels := append(labelValues, "ipv6")
	ch <- prometheus.MustNewConstMetric(totalPacketsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip6.TotalPacketsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsWithSizeSmallerThanMinimumDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsWithSizeSmallerThanMinimum, labels...)
	ch <- prometheus.MustNewConstMetric(packetsWithDatasizeLessThanDataLengthDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsWithDatasizeLessThanDataLength, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsWithBadOptionsDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsWithBadOptions, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsWithIncorrectVersionNumberDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsWithIncorrectVersionNumber, labels...)
	ch <- prometheus.MustNewConstMetric(ip6FragmentsReceivedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6FragmentsReceived, labels...)
	ch <- prometheus.MustNewConstMetric(duplicateOrOutOfSpaceFragmentsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip6.DuplicateOrOutOfSpaceFragmentsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(ip6FragmentsDroppedAfterTimeoutDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6FragmentsDroppedAfterTimeout, labels...)
	ch <- prometheus.MustNewConstMetric(fragmentsThatExceededLimitDesc, prometheus.CounterValue, s.Statistics.Ip6.FragmentsThatExceededLimit, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsReassembledOkDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsReassembledOk, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsForThisHostDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsForThisHost, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsForwardedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsForwarded, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsNotForwardableDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsNotForwardable, labels...)
	ch <- prometheus.MustNewConstMetric(ip6RedirectsSentDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6RedirectsSent, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsSentFromThisHostDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsSentFromThisHost, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsSentWithFabricatedIpHeaderDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsSentWithFabricatedIpHeader, labels...)
	ch <- prometheus.MustNewConstMetric(ip6OutputPacketsDroppedDueToNoBufsDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OutputPacketsDroppedDueToNoBufs, labels...)
	ch <- prometheus.MustNewConstMetric(ip6OutputPacketsDiscardedDueToNoRouteDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OutputPacketsDiscardedDueToNoRoute, labels...)
	ch <- prometheus.MustNewConstMetric(ip6OutputDatagramsFragmentedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OutputDatagramsFragmented, labels...)
	ch <- prometheus.MustNewConstMetric(ip6FragmentsCreatedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6FragmentsCreated, labels...)
	ch <- prometheus.MustNewConstMetric(ip6DatagramsThatCanNotBeFragmentedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6DatagramsThatCanNotBeFragmented, labels...)
	ch <- prometheus.MustNewConstMetric(packetsThatViolatedScopeRulesDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsThatViolatedScopeRules, labels...)
	ch <- prometheus.MustNewConstMetric(multicastPacketsWhichWeDoNotJoinDesc, prometheus.CounterValue, s.Statistics.Ip6.MulticastPacketsWhichWeDoNotJoin, labels...)
	ch <- prometheus.MustNewConstMetric(ip6nhTcpDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6nhTcp, labels...)
	ch <- prometheus.MustNewConstMetric(ip6nhUdpDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6nhUdp, labels...)
	ch <- prometheus.MustNewConstMetric(ip6nhIcmp6Desc, prometheus.CounterValue, s.Statistics.Ip6.Ip6nhIcmp6, labels...)
	ch <- prometheus.MustNewConstMetric(packetsWhoseHeadersAreNotContinuousDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsWhoseHeadersAreNotContinuous, labels...)
	ch <- prometheus.MustNewConstMetric(tunnelingPacketsThatCanNotFindGifDesc, prometheus.CounterValue, s.Statistics.Ip6.TunnelingPacketsThatCanNotFindGif, labels...)
	ch <- prometheus.MustNewConstMetric(packetsDiscardedDueToTooMayHeadersDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsDiscardedDueToTooMayHeaders, labels...)
	ch <- prometheus.MustNewConstMetric(failuresOfSourceAddressSelectionDesc, prometheus.CounterValue, s.Statistics.Ip6.FailuresOfSourceAddressSelection, labels...)
	for _, header := range s.Statistics.Ip6.HeaderType {
		labels := append(labelValues, "ipv6", header.HeaderForSourceAddressSelection)
		ch <- prometheus.MustNewConstMetric(headerTypeLinkLocalsDesc, prometheus.CounterValue, header.LinkLocals, labels...)
		ch <- prometheus.MustNewConstMetric(headerTypeGlobalsDesc, prometheus.CounterValue, header.Globals, labels...)
	}
	ch <- prometheus.MustNewConstMetric(forwardCacheHitDesc, prometheus.CounterValue, s.Statistics.Ip6.ForwardCacheHit, labels...)
	ch <- prometheus.MustNewConstMetric(forwardCacheMissDesc, prometheus.CounterValue, s.Statistics.Ip6.ForwardCacheMiss, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsDestinedToDeadNextHopDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsDestinedToDeadNextHop, labels...)
	ch <- prometheus.MustNewConstMetric(ip6OptionPacketsDroppedDueToRateLimitDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6OptionPacketsDroppedDueToRateLimit, labels...)
	ch <- prometheus.MustNewConstMetric(ip6PacketsDroppedDesc, prometheus.CounterValue, s.Statistics.Ip6.Ip6PacketsDropped, labels...)
	ch <- prometheus.MustNewConstMetric(packetsDroppedDueToBadProtocolDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketsDroppedDueToBadProtocol, labels...)
	ch <- prometheus.MustNewConstMetric(transitRePacketDroppedOnMgmtInterfaceDesc, prometheus.CounterValue, s.Statistics.Ip6.TransitRePacketDroppedOnMgmtInterface, labels...)
	ch <- prometheus.MustNewConstMetric(packetUsedFirstNexthopInEcmpUnilistDesc, prometheus.CounterValue, s.Statistics.Ip6.PacketUsedFirstNexthopInEcmpUnilist, labels...)
}
