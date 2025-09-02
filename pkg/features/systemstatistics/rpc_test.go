package systemstatistics

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsIPv4Unmarshaling(t *testing.T) {
	IPv4XMLDataCase1, _ := os.Open("testsFiles/IPV4/ipv4TestDataCase1.xml")
	IPv4DataCase1, _ := ioutil.ReadAll(IPv4XMLDataCase1)
	IPv4XMLDataCase2, _ := os.Open("testsFiles/IPV4/ipv4TestDataCase2.xml")
	IPv4DataCase2, _ := ioutil.ReadAll(IPv4XMLDataCase2)
	IPv4XMLDataCase3, _ := os.Open("testsFiles/IPV4/ipv4TestDataCase3.xml")
	IPv4DataCase3, _ := ioutil.ReadAll(IPv4XMLDataCase3)

	type testCase struct {
		name     string
		xmlInput string
		expect   func(t *testing.T, got SystemStatistics)
	}

	testsIPV4 := []testCase{
		{
			name:     "complete_ipv4_statistics",
			xmlInput: string(IPv4DataCase1),
			expect: func(t *testing.T, got SystemStatistics) {
				ip := got.Statistics.Ip
				assert.Equal(t, float64(1000), ip.PacketsReceived)
				assert.Equal(t, float64(5), ip.BadHeaderChecksums)
				assert.Equal(t, float64(10), ip.PacketsWithSizeSmallerThanMinimum)
				assert.Equal(t, float64(2), ip.PacketsWithDataSizeLessThanDatalength)
				assert.Equal(t, float64(3), ip.PacketsWithHeaderLengthLessThanDataSize)
				assert.Equal(t, float64(1), ip.PacketsWithDataLengthLessThanHeaderlength)
				assert.Equal(t, float64(0), ip.PacketsWithIncorrectVersionNumber)
				assert.Equal(t, float64(0), ip.PacketsDestinedToDeadNextHop)
				assert.Equal(t, float64(50), ip.FragmentsReceived)
				assert.Equal(t, float64(2), ip.FragmentsDroppedDueToOutofspaceOrDup)
				assert.Equal(t, float64(1), ip.FragmentsDroppedDueToQueueoverflow)
				assert.Equal(t, float64(0), ip.FragmentsDroppedAfterTimeout)
				assert.Equal(t, float64(48), ip.PacketsReassembledOk)
				assert.Equal(t, float64(500), ip.PacketsForThisHost)
				assert.Equal(t, float64(5), ip.PacketsForUnknownOrUnsupportedProtocol)
				assert.Equal(t, float64(400), ip.PacketsForwarded)
				assert.Equal(t, float64(10), ip.PacketsNotForwardable)
				assert.Equal(t, float64(2), ip.RedirectsSent)
				assert.Equal(t, float64(800), ip.PacketsSentFromThisHost)
				assert.Equal(t, float64(0), ip.PacketsSentWithFabricatedIpHeader)
				assert.Equal(t, float64(3), ip.OutputPacketsDroppedDueToNoBufs)
				assert.Equal(t, float64(1), ip.OutputPacketsDiscardedDueToNoRoute)
				assert.Equal(t, float64(20), ip.OutputDatagramsFragmented)
				assert.Equal(t, float64(40), ip.FragmentsCreated)
				assert.Equal(t, float64(2), ip.DatagramsThatCanNotBeFragmented)
				assert.Equal(t, float64(1), ip.PacketsWithBadOptions)
				assert.Equal(t, float64(15), ip.PacketsWithOptionsHandledWithoutError)
				assert.Equal(t, float64(0), ip.StrictSourceAndRecordRouteOptions)
				assert.Equal(t, float64(2), ip.LooseSourceAndRecordRouteOptions)
				assert.Equal(t, float64(5), ip.RecordRouteOptions)
				assert.Equal(t, float64(3), ip.TimestampOptions)
				assert.Equal(t, float64(1), ip.TimestampAndAddressOptions)
				assert.Equal(t, float64(0), ip.TimestampAndPrespecifiedAddressOptions)
				assert.Equal(t, float64(0), ip.OptionPacketsDroppedDueToRateLimit)
				assert.Equal(t, float64(4), ip.RouterAlertOptions)
				assert.Equal(t, float64(8), ip.MulticastPacketsDropped)
				assert.Equal(t, float64(12), ip.PacketsDropped)
				assert.Equal(t, float64(0), ip.TransitRePacketsDroppedOnMgmtInterface)
				assert.Equal(t, float64(25), ip.PacketsUsedFirstNexthopInEcmpUnilist)
				assert.Equal(t, float64(100), ip.IncomingTtpoipPacketsReceived)
				assert.Equal(t, float64(2), ip.IncomingTtpoipPacketsDropped)
				assert.Equal(t, float64(95), ip.OutgoingTtpoipPacketsSent)
				assert.Equal(t, float64(1), ip.OutgoingTtpoipPacketsDropped)
				assert.Equal(t, float64(3), ip.IncomingRawipPacketsDroppedNoSocketBuffer)
				assert.Equal(t, float64(200), ip.IncomingVirtualNodePacketsDelivered)
				assert.Equal(t, "user@router>", got.Cli.Banner)
			},
		},
		{
			name:     "empty_ipv4_statistics",
			xmlInput: string(IPv4DataCase2),
			expect: func(t *testing.T, got SystemStatistics) {
				assert.Equal(t, "user@router>", got.Cli.Banner)
			},
		},
		{
			name:     "high_values_ipv4_statistics",
			xmlInput: string(IPv4DataCase3),
			expect: func(t *testing.T, got SystemStatistics) {
				ip := got.Statistics.Ip
				assert.Equal(t, float64(999999999), ip.PacketsReceived)
				assert.Equal(t, float64(12345), ip.BadHeaderChecksums)
				assert.Equal(t, float64(54321), ip.PacketsWithSizeSmallerThanMinimum)
				assert.Equal(t, float64(1111), ip.PacketsWithDataSizeLessThanDatalength)
				assert.Equal(t, float64(2222), ip.PacketsWithHeaderLengthLessThanDataSize)
				assert.Equal(t, float64(3333), ip.PacketsWithDataLengthLessThanHeaderlength)
				assert.Equal(t, float64(4444), ip.PacketsWithIncorrectVersionNumber)
				assert.Equal(t, float64(5555), ip.PacketsDestinedToDeadNextHop)
				assert.Equal(t, float64(888888), ip.FragmentsReceived)
				assert.Equal(t, float64(6666), ip.FragmentsDroppedDueToOutofspaceOrDup)
				assert.Equal(t, float64(7777), ip.FragmentsDroppedDueToQueueoverflow)
				assert.Equal(t, float64(8888), ip.FragmentsDroppedAfterTimeout)
				assert.Equal(t, float64(777777), ip.PacketsReassembledOk)
				assert.Equal(t, float64(555555), ip.PacketsForThisHost)
				assert.Equal(t, float64(9999), ip.PacketsForUnknownOrUnsupportedProtocol)
				assert.Equal(t, float64(444444), ip.PacketsForwarded)
				assert.Equal(t, float64(11111), ip.PacketsNotForwardable)
				assert.Equal(t, float64(12121), ip.RedirectsSent)
				assert.Equal(t, float64(666666), ip.PacketsSentFromThisHost)
				assert.Equal(t, float64(13131), ip.PacketsSentWithFabricatedIpHeader)
				assert.Equal(t, float64(14141), ip.OutputPacketsDroppedDueToNoBufs)
				assert.Equal(t, float64(15151), ip.OutputPacketsDiscardedDueToNoRoute)
				assert.Equal(t, float64(16161), ip.OutputDatagramsFragmented)
				assert.Equal(t, float64(17171), ip.FragmentsCreated)
				assert.Equal(t, float64(18181), ip.DatagramsThatCanNotBeFragmented)
				assert.Equal(t, float64(19191), ip.PacketsWithBadOptions)
				assert.Equal(t, float64(20202), ip.PacketsWithOptionsHandledWithoutError)
				assert.Equal(t, float64(21212), ip.StrictSourceAndRecordRouteOptions)
				assert.Equal(t, float64(22222), ip.LooseSourceAndRecordRouteOptions)
				assert.Equal(t, float64(23232), ip.RecordRouteOptions)
				assert.Equal(t, float64(24242), ip.TimestampOptions)
				assert.Equal(t, float64(25252), ip.TimestampAndAddressOptions)
				assert.Equal(t, float64(26262), ip.TimestampAndPrespecifiedAddressOptions)
				assert.Equal(t, float64(27272), ip.OptionPacketsDroppedDueToRateLimit)
				assert.Equal(t, float64(28282), ip.RouterAlertOptions)
				assert.Equal(t, float64(29292), ip.MulticastPacketsDropped)
				assert.Equal(t, float64(30303), ip.PacketsDropped)
				assert.Equal(t, float64(31313), ip.TransitRePacketsDroppedOnMgmtInterface)
				assert.Equal(t, float64(32323), ip.PacketsUsedFirstNexthopInEcmpUnilist)
				assert.Equal(t, float64(33333), ip.IncomingTtpoipPacketsReceived)
				assert.Equal(t, float64(34343), ip.IncomingTtpoipPacketsDropped)
				assert.Equal(t, float64(35353), ip.OutgoingTtpoipPacketsSent)
				assert.Equal(t, float64(36363), ip.OutgoingTtpoipPacketsDropped)
				assert.Equal(t, float64(37373), ip.IncomingRawipPacketsDroppedNoSocketBuffer)
				assert.Equal(t, float64(38383), ip.IncomingVirtualNodePacketsDelivered)
				assert.Equal(t, "admin@high-traffic-router>", got.Cli.Banner)
			},
		},
	}

	for _, tc := range testsIPV4 {
		t.Run(tc.name, func(t *testing.T) {
			var result SystemStatistics
			err := xml.Unmarshal([]byte(tc.xmlInput), &result)
			assert.NoError(t, err, "unmarshal should not return error")
			tc.expect(t, result)
		})
	}
}

// Tests for the IPv6 sub-structure (Ip6) of SystemStatistics. We use inline XML to focus on Ip6.
func TestStatisticsIPv6Unmarshaling(t *testing.T) {
	IPv6XMLDataCase1, _ := os.Open("testsFiles/IPv6/ipv6TestDataCase1.xml")
	IPv6DataCase1, _ := ioutil.ReadAll(IPv6XMLDataCase1)
	type testCase struct {
		name     string
		xmlInput string
		expect   func(t *testing.T, got SystemStatistics)
	}
	tests := []testCase{
		{
			name:     "complete_ipv6_statistics",
			xmlInput: string(IPv6DataCase1),
			expect: func(t *testing.T, got SystemStatistics) {
				ip6 := got.Statistics.Ip6
				assert.Equal(t, float64(100), ip6.TotalPacketsReceived)
				assert.Equal(t, float64(1), ip6.Ip6PacketsWithSizeSmallerThanMinimum)
				assert.Equal(t, float64(2), ip6.PacketsWithDatasizeLessThanDataLength)
				assert.Equal(t, float64(3), ip6.Ip6PacketsWithBadOptions)
				assert.Equal(t, float64(4), ip6.Ip6PacketsWithIncorrectVersionNumber)
				assert.Equal(t, float64(5), ip6.Ip6FragmentsReceived)
				assert.Equal(t, float64(6), ip6.DuplicateOrOutOfSpaceFragmentsDropped)
				assert.Equal(t, float64(7), ip6.Ip6FragmentsDroppedAfterTimeout)
				assert.Equal(t, float64(8), ip6.FragmentsThatExceededLimit)
				assert.Equal(t, float64(9), ip6.Ip6PacketsReassembledOk)
				assert.Equal(t, float64(10), ip6.Ip6PacketsForThisHost)
				assert.Equal(t, float64(11), ip6.Ip6PacketsForwarded)
				assert.Equal(t, float64(12), ip6.Ip6PacketsNotForwardable)
				assert.Equal(t, float64(13), ip6.Ip6RedirectsSent)
				assert.Equal(t, float64(14), ip6.Ip6PacketsSentFromThisHost)
				assert.Equal(t, float64(15), ip6.Ip6PacketsSentWithFabricatedIpHeader)
				assert.Equal(t, float64(16), ip6.Ip6OutputPacketsDroppedDueToNoBufs)
				assert.Equal(t, float64(17), ip6.Ip6OutputPacketsDiscardedDueToNoRoute)
				assert.Equal(t, float64(18), ip6.Ip6OutputDatagramsFragmented)
				assert.Equal(t, float64(19), ip6.Ip6FragmentsCreated)
				assert.Equal(t, float64(20), ip6.Ip6DatagramsThatCanNotBeFragmented)
				assert.Equal(t, float64(21), ip6.PacketsThatViolatedScopeRules)
				assert.Equal(t, float64(22), ip6.MulticastPacketsWhichWeDoNotJoin)
				assert.Equal(t, float64(23), ip6.Ip6nhTcp)
				assert.Equal(t, float64(24), ip6.Ip6nhUdp)
				assert.Equal(t, float64(25), ip6.Ip6nhIcmp6)
				assert.Equal(t, float64(26), ip6.PacketsWhoseHeadersAreNotContinuous)
				assert.Equal(t, float64(27), ip6.TunnelingPacketsThatCanNotFindGif)
				assert.Equal(t, float64(28), ip6.PacketsDiscardedDueToTooMayHeaders)
				assert.Equal(t, float64(29), ip6.FailuresOfSourceAddressSelection)
				assert.Equal(t, 2, len(ip6.HeaderType))
				var defLink, defGlob, polLink, polGlob float64
				for _, h := range ip6.HeaderType {
					switch h.HeaderForSourceAddressSelection {
					case "default":
						defLink = h.LinkLocals
						defGlob = h.Globals
					case "policy":
						polLink = h.LinkLocals
						polGlob = h.Globals
					}
				}
				assert.Equal(t, float64(30), defLink)
				assert.Equal(t, float64(31), defGlob)
				assert.Equal(t, float64(32), polLink)
				assert.Equal(t, float64(33), polGlob)
				assert.Equal(t, float64(34), ip6.ForwardCacheHit)
				assert.Equal(t, float64(35), ip6.ForwardCacheMiss)
				assert.Equal(t, float64(36), ip6.Ip6PacketsDestinedToDeadNextHop)
				assert.Equal(t, float64(37), ip6.Ip6OptionPacketsDroppedDueToRateLimit)
				assert.Equal(t, float64(38), ip6.Ip6PacketsDropped)
				assert.Equal(t, float64(39), ip6.PacketsDroppedDueToBadProtocol)
				assert.Equal(t, float64(40), ip6.TransitRePacketDroppedOnMgmtInterface)
				assert.Equal(t, float64(41), ip6.PacketUsedFirstNexthopInEcmpUnilist)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got SystemStatistics
			err := xml.Unmarshal([]byte(tc.xmlInput), &got)
			assert.NoError(t, err, "unmarshal should not return error")
			tc.expect(t, got)
		})
	}
}

func TestStatisticsUDPUnmarshaling(t *testing.T) {
	UDPXMLDataCase1, _ := os.Open("testsFiles/UDP/UDPTestDataCase1.xml")
	UDPDataCase1, _ := ioutil.ReadAll(UDPXMLDataCase1)
	type testCase struct {
		name     string
		xmlInput string
		expect   func(t *testing.T, got SystemStatistics)
	}

	tests := []testCase{
		{
			name:     "complete_udp_statistics",
			xmlInput: string(UDPDataCase1),
			expect: func(t *testing.T, got SystemStatistics) {
				udp := got.Statistics.Udp
				assert.Equal(t, float64(100), udp.DatagramsReceived)
				assert.Equal(t, float64(1), udp.DatagramsWithIncompleteHeader)
				assert.Equal(t, float64(2), udp.DatagramsWithBadDatalengthField)
				assert.Equal(t, float64(3), udp.DatagramsWithBadChecksum)
				assert.Equal(t, float64(4), udp.DatagramsDroppedDueToNoSocket)
				assert.Equal(t, float64(5), udp.BroadcastOrMulticastDatagramsDroppedDueToNoSocket)
				assert.Equal(t, float64(6), udp.DatagramsDroppedDueToFullSocketBuffers)
				assert.Equal(t, float64(7), udp.DatagramsNotForHashedPcb)
				assert.Equal(t, float64(8), udp.DatagramsDelivered)
				assert.Equal(t, float64(9), udp.DatagramsOutput)
				assert.Equal(t, "user@router>", got.Cli.Banner)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got SystemStatistics
			err := xml.Unmarshal([]byte(tc.xmlInput), &got)
			assert.NoError(t, err, "unmarshal should not return error")
			tc.expect(t, got)
		})
	}
}

func TestStatisticsTCPUnmarshaling(t *testing.T) {
	TCPXMLDataCase1, _ := os.Open("testsFiles/TCP/TCPTestDataCase1.xml")
	TCPDataCase1, _ := ioutil.ReadAll(TCPXMLDataCase1)
	type testCase struct {
		name     string
		xmlInput string
		expect   func(t *testing.T, got SystemStatistics)
	}

	tests := []testCase{
		{
			name:     "complete_tcp_statistics",
			xmlInput: string(TCPDataCase1),
			expect: func(t *testing.T, got SystemStatistics) {
				tcp := got.Statistics.Tcp
				assert.Equal(t, float64(1000), tcp.PacketsSent)
				assert.Equal(t, float64(900), tcp.SentDataPackets)
				assert.Equal(t, float64(123456), tcp.DataPacketsBytes)
				assert.Equal(t, float64(10), tcp.SentDataPacketsRetransmitted)
				assert.Equal(t, float64(2048), tcp.RetransmittedBytes)
				assert.Equal(t, float64(50), tcp.SentAckOnlyPackets)
				assert.Equal(t, float64(51), tcp.SentPacketsDelayed)
				assert.Equal(t, float64(52), tcp.SentUrgOnlyPackets)
				assert.Equal(t, float64(53), tcp.SentWindowProbePackets)
				assert.Equal(t, float64(54), tcp.SentWindowUpdatePackets)
				assert.Equal(t, float64(55), tcp.SentControlPackets)
				assert.Equal(t, float64(56), tcp.PacketsReceived)
				assert.Equal(t, float64(57), tcp.ReceivedAcks)
				assert.Equal(t, float64(58), tcp.AcksBytes)
				assert.Equal(t, float64(59), tcp.ReceivedDuplicateAcks)
				assert.Equal(t, float64(60), tcp.ReceivedAcksForUnsentData)
				assert.Equal(t, float64(61), tcp.PacketsReceivedInSequence)
				assert.Equal(t, float64(62), tcp.InSequenceBytes)
				assert.Equal(t, float64(63), tcp.ReceivedCompletelyDuplicatePacket)
				assert.Equal(t, float64(64), tcp.DuplicateInBytes)
				assert.Equal(t, float64(65), tcp.ReceivedOldDuplicatePackets)
				assert.Equal(t, float64(66), tcp.ReceivedPacketsWithSomeDupliacteData)
				assert.Equal(t, float64(67), tcp.SomeDuplicateInBytes)
				assert.Equal(t, float64(68), tcp.ReceivedOutOfOrderPackets)
				assert.Equal(t, float64(69), tcp.OutOfOrderInBytes)
				assert.Equal(t, float64(70), tcp.ReceivedPacketsOfDataAfterWindow)
				assert.Equal(t, float64(71), tcp.Bytes)
				assert.Equal(t, float64(72), tcp.ReceivedWindowProbes)
				assert.Equal(t, float64(73), tcp.ReceivedWindowUpdatePackets)
				assert.Equal(t, float64(74), tcp.PacketsReceivedAfterClose)
				assert.Equal(t, float64(75), tcp.ReceivedDiscardedForBadChecksum)
				assert.Equal(t, float64(76), tcp.ReceivedDiscardedForBadHeaderOffset)
				assert.Equal(t, float64(77), tcp.ReceivedDiscardedBecausePacketTooShort)
				assert.Equal(t, float64(78), tcp.ConnectionRequests)
				assert.Equal(t, float64(79), tcp.ConnectionAccepts)
				assert.Equal(t, float64(80), tcp.BadConnectionAttempts)
				assert.Equal(t, float64(81), tcp.ListenQueueOverflows)
				assert.Equal(t, float64(82), tcp.BadRstWindow)
				assert.Equal(t, float64(83), tcp.ConnectionsEstablished)
				assert.Equal(t, float64(84), tcp.ConnectionsClosed)
				assert.Equal(t, float64(85), tcp.Drops)
				assert.Equal(t, float64(86), tcp.ConnectionsUpdatedRttOnClose)
				assert.Equal(t, float64(87), tcp.ConnectionsUpdatedVarianceOnClose)
				assert.Equal(t, float64(88), tcp.ConnectionsUpdatedSsthreshOnClose)
				assert.Equal(t, float64(89), tcp.EmbryonicConnectionsDropped)
				assert.Equal(t, float64(90), tcp.SegmentsUpdatedRtt)
				assert.Equal(t, float64(91), tcp.Attempts)
				assert.Equal(t, float64(92), tcp.RetransmitTimeouts)
				assert.Equal(t, float64(93), tcp.ConnectionsDroppedByRetransmitTimeout)
				assert.Equal(t, float64(94), tcp.PersistTimeouts)
				assert.Equal(t, float64(95), tcp.ConnectionsDroppedByPersistTimeout)
				assert.Equal(t, float64(96), tcp.KeepaliveTimeouts)
				assert.Equal(t, float64(97), tcp.KeepaliveProbesSent)
				assert.Equal(t, float64(98), tcp.KeepaliveConnectionsDropped)
				assert.Equal(t, float64(99), tcp.AckHeaderPredictions)
				assert.Equal(t, float64(100), tcp.DataPacketHeaderPredictions)
				assert.Equal(t, float64(101), tcp.SyncacheEntriesAdded)
				assert.Equal(t, float64(102), tcp.Retransmitted)
				assert.Equal(t, float64(103), tcp.Dupsyn)
				assert.Equal(t, float64(104), tcp.Dropped)
				assert.Equal(t, float64(105), tcp.Completed)
				assert.Equal(t, float64(106), tcp.BucketOverflow)
				assert.Equal(t, float64(107), tcp.CacheOverflow)
				assert.Equal(t, float64(108), tcp.Reset)
				assert.Equal(t, float64(109), tcp.Stale)
				assert.Equal(t, float64(110), tcp.Aborted)
				assert.Equal(t, float64(111), tcp.Badack)
				assert.Equal(t, float64(112), tcp.Unreach)
				assert.Equal(t, float64(113), tcp.ZoneFailures)
				assert.Equal(t, float64(114), tcp.CookiesSent)
				assert.Equal(t, float64(115), tcp.CookiesReceived)
				assert.Equal(t, float64(116), tcp.SackRecoveryEpisodes)
				assert.Equal(t, float64(117), tcp.SegmentRetransmits)
				assert.Equal(t, float64(118), tcp.ByteRetransmits)
				assert.Equal(t, float64(119), tcp.SackOptionsReceived)
				assert.Equal(t, float64(120), tcp.SackOpitionsSent)
				assert.Equal(t, float64(121), tcp.SackScoreboardOverflow)
				assert.Equal(t, float64(122), tcp.AcksSentInResponseButNotExactRsts)
				assert.Equal(t, float64(123), tcp.AcksSentInResponseToSynsOnEstablishedConnections)
				assert.Equal(t, float64(124), tcp.RcvPacketsDroppedDueToBadAddress)
				assert.Equal(t, float64(125), tcp.OutOfSequenceSegmentDrops)
				assert.Equal(t, float64(126), tcp.RstPackets)
				assert.Equal(t, float64(127), tcp.IcmpPacketsIgnored)
				assert.Equal(t, float64(128), tcp.SendPacketsDropped)
				assert.Equal(t, float64(129), tcp.RcvPacketsDropped)
				assert.Equal(t, float64(130), tcp.OutgoingSegmentsDropped)
				assert.Equal(t, float64(131), tcp.ReceivedSynfinDropped)
				assert.Equal(t, float64(132), tcp.ReceivedIpsecDropped)
				assert.Equal(t, float64(133), tcp.ReceivedMacDropped)
				assert.Equal(t, float64(134), tcp.ReceivedMinttlExceeded)
				assert.Equal(t, float64(135), tcp.ListenstateBadflagsDropped)
				assert.Equal(t, float64(136), tcp.FinwaitstateBadflagsDropped)
				assert.Equal(t, float64(137), tcp.ReceivedDosAttack)
				assert.Equal(t, float64(138), tcp.ReceivedBadSynack)
				assert.Equal(t, float64(139), tcp.SyncacheZoneFull)
				assert.Equal(t, float64(140), tcp.ReceivedRstFirewallfilter)
				assert.Equal(t, float64(141), tcp.ReceivedNoackTimewait)
				assert.Equal(t, float64(142), tcp.ReceivedNoTimewaitState)
				assert.Equal(t, float64(143), tcp.ReceivedRstTimewaitState)
				assert.Equal(t, float64(144), tcp.ReceivedTimewaitDrops)
				assert.Equal(t, float64(145), tcp.ReceivedBadaddrTimewaitState)
				assert.Equal(t, float64(146), tcp.ReceivedAckoffInSynSentrcvd)
				assert.Equal(t, float64(147), tcp.ReceivedBadaddrFirewall)
				assert.Equal(t, float64(148), tcp.ReceivedNosynSynSent)
				assert.Equal(t, float64(149), tcp.ReceivedBadrstSynSent)
				assert.Equal(t, float64(150), tcp.ReceivedBadrstListenState)
				assert.Equal(t, float64(151), tcp.OptionMaxsegmentLength)
				assert.Equal(t, float64(152), tcp.OptionWindowLength)
				assert.Equal(t, float64(153), tcp.OptionTimestampLength)
				assert.Equal(t, float64(154), tcp.OptionMd5Length)
				assert.Equal(t, float64(155), tcp.OptionAuthLength)
				assert.Equal(t, float64(156), tcp.OptionSackpermittedLength)
				assert.Equal(t, float64(157), tcp.OptionSackLength)
				assert.Equal(t, float64(158), tcp.OptionAuthoptionLength)
				assert.Equal(t, "user@router>", got.Cli.Banner)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got SystemStatistics
			err := xml.Unmarshal([]byte(tc.xmlInput), &got)
			assert.NoError(t, err, "unmarshal should not return error")
			tc.expect(t, got)
		})
	}
}
