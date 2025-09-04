package systemstatistics

import (
	"encoding/xml"
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
						DatagramsReceived:                                 3000,
						DatagramsWithIncompleteHeader:                     3001,
						DatagramsWithBadDatalengthField:                   3002,
						DatagramsWithBadChecksum:                          3003,
						DatagramsDroppedDueToNoSocket:                     3004,
						BroadcastOrMulticastDatagramsDroppedDueToNoSocket: 3005,
						DatagramsDroppedDueToFullSocketBuffers:            3006,
						DatagramsNotForHashedPcb:                          3007,
						DatagramsDelivered:                                3008,
						DatagramsOutput:                                   3009,
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

func TestStatisticsTCPUnmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}

	testsTCP := []testCase{
		{
			name:    "complete_tcp_statistics",
			xmlFile: "testsFiles/TCP/TCPTestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Tcp: TCP{
						PacketsSent:                            4000,
						SentDataPackets:                        4001,
						DataPacketsBytes:                       4002,
						SentDataPacketsRetransmitted:           4003,
						RetransmittedBytes:                     4004,
						SentDataUnnecessaryRetransmitted:       4005,
						SentResendsByMtuDiscovery:              4006,
						SentAckOnlyPackets:                     4007,
						SentPacketsDelayed:                     4008,
						SentUrgOnlyPackets:                     4009,
						SentWindowProbePackets:                 4010,
						SentWindowUpdatePackets:                4011,
						SentControlPackets:                     4012,
						PacketsReceived:                        4013,
						ReceivedAcks:                           4014,
						AcksBytes:                              4015,
						ReceivedDuplicateAcks:                  4016,
						ReceivedAcksForUnsentData:              4017,
						PacketsReceivedInSequence:              4018,
						InSequenceBytes:                        4019,
						ReceivedCompletelyDuplicatePacket:      4020,
						DuplicateInBytes:                       4021,
						ReceivedOldDuplicatePackets:            4022,
						ReceivedPacketsWithSomeDupliacteData:   4023,
						SomeDuplicateInBytes:                   4024,
						ReceivedOutOfOrderPackets:              4025,
						OutOfOrderInBytes:                      4026,
						ReceivedPacketsOfDataAfterWindow:       4027,
						Bytes:                                  4028,
						ReceivedWindowProbes:                   4029,
						ReceivedWindowUpdatePackets:            4030,
						PacketsReceivedAfterClose:              4031,
						ReceivedDiscardedForBadChecksum:        4032,
						ReceivedDiscardedForBadHeaderOffset:    4033,
						ReceivedDiscardedBecausePacketTooShort: 4034,
						ConnectionRequests:                     4035,
						ConnectionAccepts:                      4036,
						BadConnectionAttempts:                  4037,
						ListenQueueOverflows:                   4038,
						BadRstWindow:                           4039,
						ConnectionsEstablished:                 4040,
						ConnectionsClosed:                      4041,
						Drops:                                  4042,
						ConnectionsUpdatedRttOnClose:           4043,
						ConnectionsUpdatedVarianceOnClose:      4044,
						ConnectionsUpdatedSsthreshOnClose:      4045,
						EmbryonicConnectionsDropped:            4046,
						SegmentsUpdatedRtt:                     4047,
						Attempts:                               4048,
						RetransmitTimeouts:                     4049,
						ConnectionsDroppedByRetransmitTimeout:  4050,
						PersistTimeouts:                        4051,
						ConnectionsDroppedByPersistTimeout:     4052,
						KeepaliveTimeouts:                      4053,
						KeepaliveProbesSent:                    4054,
						KeepaliveConnectionsDropped:            4055,
						AckHeaderPredictions:                   4056,
						DataPacketHeaderPredictions:            4057,
						SyncacheEntriesAdded:                   4058,
						Retransmitted:                          4059,
						Dupsyn:                                 4060,
						Dropped:                                4061,
						Completed:                              4062,
						BucketOverflow:                         4063,
						CacheOverflow:                          4064,
						Reset:                                  4065,
						Stale:                                  4066,
						Aborted:                                4067,
						Badack:                                 4068,
						Unreach:                                4069,
						ZoneFailures:                           4070,
						CookiesSent:                            4071,
						CookiesReceived:                        4072,
						SackRecoveryEpisodes:                   4073,
						SegmentRetransmits:                     4074,
						ByteRetransmits:                        4075,
						SackOptionsReceived:                    4076,
						SackOpitionsSent:                       4077,
						SackScoreboardOverflow:                 4078,
						AcksSentInResponseButNotExactRsts:      4079,
						AcksSentInResponseToSynsOnEstablishedConnections: 4080,
						RcvPacketsDroppedDueToBadAddress:                 4081,
						OutOfSequenceSegmentDrops:                        4082,
						RstPackets:                                       4083,
						IcmpPacketsIgnored:                               4084,
						SendPacketsDropped:                               4085,
						RcvPacketsDropped:                                4086,
						OutgoingSegmentsDropped:                          4087,
						ReceivedSynfinDropped:                            4088,
						ReceivedIpsecDropped:                             4089,
						ReceivedMacDropped:                               4090,
						ReceivedMinttlExceeded:                           4091,
						ListenstateBadflagsDropped:                       4092,
						FinwaitstateBadflagsDropped:                      4093,
						ReceivedDosAttack:                                4094,
						ReceivedBadSynack:                                4095,
						SyncacheZoneFull:                                 4096,
						ReceivedRstFirewallfilter:                        4097,
						ReceivedNoackTimewait:                            4098,
						ReceivedNoTimewaitState:                          4099,
						ReceivedRstTimewaitState:                         4100,
						ReceivedTimewaitDrops:                            4101,
						ReceivedBadaddrTimewaitState:                     4102,
						ReceivedAckoffInSynSentrcvd:                      4103,
						ReceivedBadaddrFirewall:                          4104,
						ReceivedNosynSynSent:                             4105,
						ReceivedBadrstSynSent:                            4106,
						ReceivedBadrstListenState:                        4107,
						OptionMaxsegmentLength:                           4108,
						OptionWindowLength:                               4109,
						OptionTimestampLength:                            4110,
						OptionMd5Length:                                  4111,
						OptionAuthLength:                                 4112,
						OptionSackpermittedLength:                        4113,
						OptionSackLength:                                 4114,
						OptionAuthoptionLength:                           4115,
					},
				},
			},
		},
	}

	for _, tc := range testsTCP {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {

			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {

			}
			result.Statistics.Tcp.Text = ""
			assert.Equal(t, tc.expect.Statistics.Tcp, result.Statistics.Tcp, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}
