package systemstatistics

import (
	"encoding/xml"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsIPv4Unmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}

	testsIPV4 := []testCase{
		{
			name:    "complete_ipv4_statistics",
			xmlFile: "testsFiles/IPv4/ipv4TestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Ip: IP{
						PacketsReceived:                           1000,
						BadHeaderChecksums:                        1001,
						PacketsWithSizeSmallerThanMinimum:         1002,
						PacketsWithDataSizeLessThanDatalength:     1003,
						PacketsWithHeaderLengthLessThanDataSize:   1004,
						PacketsWithDataLengthLessThanHeaderlength: 1005,
						PacketsWithIncorrectVersionNumber:         1006,
						PacketsDestinedToDeadNextHop:              1007,
						FragmentsReceived:                         1008,
						FragmentsDroppedDueToOutofspaceOrDup:      1009,
						FragmentsDroppedDueToQueueoverflow:        1010,
						FragmentsDroppedAfterTimeout:              1011,
						PacketsReassembledOk:                      1012,
						PacketsForThisHost:                        1013,
						PacketsForUnknownOrUnsupportedProtocol:    1014,
						PacketsForwarded:                          1015,
						PacketsNotForwardable:                     1016,
						RedirectsSent:                             1017,
						PacketsSentFromThisHost:                   1018,
						PacketsSentWithFabricatedIpHeader:         1019,
						OutputPacketsDroppedDueToNoBufs:           1020,
						OutputPacketsDiscardedDueToNoRoute:        1021,
						OutputDatagramsFragmented:                 1022,
						FragmentsCreated:                          1023,
						DatagramsThatCanNotBeFragmented:           1024,
						PacketsWithBadOptions:                     1025,
						PacketsWithOptionsHandledWithoutError:     1026,
						StrictSourceAndRecordRouteOptions:         1027,
						LooseSourceAndRecordRouteOptions:          1028,
						RecordRouteOptions:                        1029,
						TimestampOptions:                          1030,
						TimestampAndAddressOptions:                1031,
						TimestampAndPrespecifiedAddressOptions:    1032,
						OptionPacketsDroppedDueToRateLimit:        1033,
						RouterAlertOptions:                        1034,
						MulticastPacketsDropped:                   1035,
						PacketsDropped:                            1036,
						TransitRePacketsDroppedOnMgmtInterface:    1037,
						PacketsUsedFirstNexthopInEcmpUnilist:      1038,
						IncomingTtpoipPacketsReceived:             1039,
						IncomingTtpoipPacketsDropped:              1040,
						OutgoingTtpoipPacketsSent:                 1041,
						OutgoingTtpoipPacketsDropped:              1042,
						IncomingRawipPacketsDroppedNoSocketBuffer: 1043,
						IncomingVirtualNodePacketsDelivered:       1044,
					},
				},
			},
		},
	}

	for _, tc := range testsIPV4 {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {

			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {

			}

			result.Statistics.Ip.Text = ""
			assert.Equal(t, tc.expect.Statistics.Ip, result.Statistics.Ip, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}

func TestStatisticsIPv6Unmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}

	testsIPV6 := []testCase{
		{
			name:    "complete_ipv6_statistics",
			xmlFile: "testsFiles/IPv6/ipv6TestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Ip6: IP6{
						TotalPacketsReceived:                  2000,
						Ip6PacketsWithSizeSmallerThanMinimum:  2001,
						PacketsWithDatasizeLessThanDataLength: 2002,
						Ip6PacketsWithBadOptions:              2003,
						Ip6PacketsWithIncorrectVersionNumber:  2004,
						Ip6FragmentsReceived:                  2005,
						DuplicateOrOutOfSpaceFragmentsDropped: 2006,
						Ip6FragmentsDroppedAfterTimeout:       2007,
						FragmentsThatExceededLimit:            2008,
						Ip6PacketsReassembledOk:               2009,
						Ip6PacketsForThisHost:                 2010,
						Ip6PacketsForwarded:                   2011,
						Ip6PacketsNotForwardable:              2012,
						Ip6RedirectsSent:                      2013,
						Ip6PacketsSentFromThisHost:            2014,
						Ip6PacketsSentWithFabricatedIpHeader:  2015,
						Ip6OutputPacketsDroppedDueToNoBufs:    2016,
						Ip6OutputPacketsDiscardedDueToNoRoute: 2017,
						Ip6OutputDatagramsFragmented:          2018,
						Ip6FragmentsCreated:                   2019,
						Ip6DatagramsThatCanNotBeFragmented:    2020,
						PacketsThatViolatedScopeRules:         2021,
						MulticastPacketsWhichWeDoNotJoin:      2022,
						Ip6nhTcp:                              2023,
						Ip6nhUdp:                              2024,
						Ip6nhIcmp6:                            2025,
						PacketsWhoseHeadersAreNotContinuous:   2026,
						TunnelingPacketsThatCanNotFindGif:     2027,
						PacketsDiscardedDueToTooMayHeaders:    2028,
						FailuresOfSourceAddressSelection:      2029,
						HeaderType: []HeaderType{
							{
								LinkLocals: 2030,
								Globals:    2031,
							},
							{
								LinkLocals: 2100,
								Globals:    2101,
							},
						},
						ForwardCacheHit:                       2032,
						ForwardCacheMiss:                      2033,
						Ip6PacketsDestinedToDeadNextHop:       2034,
						Ip6OptionPacketsDroppedDueToRateLimit: 2035,
						Ip6PacketsDropped:                     2036,
						PacketsDroppedDueToBadProtocol:        2037,
						TransitRePacketDroppedOnMgmtInterface: 2038,
						PacketUsedFirstNexthopInEcmpUnilist:   2039,
					},
				},
			},
		},
	}

	for _, tc := range testsIPV6 {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {

			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {

			}
			result.Statistics.Ip6.Text = ""
			for i, _ := range result.Statistics.Ip6.HeaderType {
				result.Statistics.Ip6.HeaderType[i].Text = ""
				tc.expect.Statistics.Ip6.HeaderType[i].HeaderForSourceAddressSelection = ""
			}
			assert.Equal(t, tc.expect.Statistics.Ip6, result.Statistics.Ip6, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}

func TestStatisticsUDPUnmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}

	testsUDP := []testCase{
		{
			name:    "complete_udp_statistics",
			xmlFile: "testsFiles/UDP/UDPTestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Udp: UDP{
						DatagramsReceived: 3000,
						DatagramsWithIncompleteHeader: 3001,
						DatagramsWithBadDatalengthField: 3002,
						DatagramsWithBadChecksum: 3003,
						DatagramsDroppedDueToNoSocket: 3004,
						BroadcastOrMulticastDatagramsDroppedDueToNoSocket: 3005,
						DatagramsDroppedDueToFullSocketBuffers: 3006,
						DatagramsNotForHashedPcb: 3007,
						DatagramsDelivered: 3008,
						DatagramsOutput: 3009,
					},
				},
			},
		},
	}

	for _, tc := range testsUDP {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {

			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {

			}

			result.Statistics.Udp.Text = ""
			assert.Equal(t, tc.expect.Statistics.Udp, result.Statistics.Udp, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}

/*
func TestStatisticsUDPUnmarshaling(t *testing.T) {
	UDPXMLDataCase1, _ := os.Open("testsFiles/UDP/UDPTestDataCase1.xml")
	UDPDataCase1, _ := io.ReadAll(UDPXMLDataCase1)
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
				//assert.Equal(t, "user@router>", got.Cli.Banner)
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
*/


func TestStatisticsTCPUnmarshaling(t *testing.T) {
	TCPXMLDataCase1, _ := os.Open("testsFiles/TCP/TCPTestDataCase1.xml")
	TCPDataCase1, _ := io.ReadAll(TCPXMLDataCase1)
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
				//assert.Equal(t, "user@router>", got.Cli.Banner)
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
