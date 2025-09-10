package systemstatistics

import (
	"encoding/xml"
	"log"
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

	tests := []testCase{
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in IPv4 testing due to: ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in IPv4 testing due to: ", err)
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

	tests := []testCase{
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in IPv6 testing due to: ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in IPv6 testing due to: ", err)
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

	tests := []testCase{
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in UDP testing due to: ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in UDP testing due to: ", err)
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

	tests := []testCase{
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in TCP testing due to; ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in TCP testing due to: ", err)
			}
			result.Statistics.Tcp.Text = ""
			assert.Equal(t, tc.expect.Statistics.Tcp, result.Statistics.Tcp, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}

func TestStatisticsARPUnmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}

	tests := []testCase{
		{
			name:    "complete_arp_statistics",
			xmlFile: "testsFiles/ARP/ARPTestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Arp: ARP{
						DatagramsReceived:                                        5000,
						ArpRequestsReceived:                                      5001,
						ArpRepliesReceived:                                       5002,
						ResolutionRequestReceived:                                5003,
						ResolutionRequestDropped:                                 5004,
						UnrestrictedProxyRequests:                                5005,
						RestrictedProxyRequests:                                  5006,
						ReceivedProxyRequests:                                    5007,
						ProxyRequestsNotProxied:                                  5008,
						RestrictedProxyRequestsNotProxied:                        5009,
						DatagramsWithBogusInterface:                              5010,
						DatagramsWithIncorrectLength:                             5011,
						DatagramsForNonIpProtocol:                                5012,
						DatagramsWithUnsupportedOpcode:                           5013,
						DatagramsWithBadProtocolAddressLength:                    5014,
						DatagramsWithBadHardwareAddressLength:                    5015,
						DatagramsWithMulticastSourceAddress:                      5016,
						DatagramsWithMulticastTargetAddress:                      5017,
						DatagramsWithMyOwnHardwareAddress:                        5018,
						DatagramsForAnAddressNotOnTheInterface:                   5019,
						DatagramsWithABroadcastSourceAddress:                     5020,
						DatagramsWithSourceAddressDuplicateToMine:                5021,
						DatagramsWhichWereNotForMe:                               5022,
						PacketsDiscardedWaitingForResolution:                     5023,
						PacketsSentAfterWaitingForResolution:                     5024,
						ArpRequestsSent:                                          5025,
						ArpRepliesSent:                                           5026,
						RequestsForMemoryDenied:                                  5027,
						RequestsDroppedOnEntry:                                   5028,
						RequestsDroppedDuringRetry:                               5029,
						RequestsDroppedDueToInterfaceDeletion:                    5030,
						RequestsOnUnnumberedInterfaces:                           5031,
						NewRequestsOnUnnumberedInterfaces:                        5032,
						RepliesFromUnnumberedInterfaces:                          5033,
						RequestsOnUnnumberedInterfaceWithNonSubnettedDonor:       5034,
						RepliesFromUnnumberedInterfaceWithNonSubnettedDonor:      5035,
						ArpPacketsRejectedAsFamilyIsConfiguredWithDenyArp:        5036,
						ArpResponsePacketsAreRejectedOnMcAeIclInterface:          5037,
						ArpRepliesAreRejectedAsSourceAndDestinationIsSame:        5038,
						ArpProbeForProxyAddressReachableFromTheIncomingInterface: 5039,
						ArpRequestDiscardedForVrrpSourceAddress:                  5040,
						SelfArpRequestPacketReceivedOnIrbInterface:               5041,
						ProxyArpRequestDiscardedAsSourceIpIsAProxyTarget:         5042,
						ArpPacketsAreDroppedAsNexthopAllocationFailed:            5043,
						ArpPacketsReceivedFromPeerVrrpRouterAndDiscarded:         5044,
						ArpPacketsAreRejectedAsTargetIpArpResolveIsInProgress:    5045,
						GratArpPacketsAreIgnoredAsMacAddressIsNotChanged:         5046,
						ArpPacketsAreDroppedFromPeerVrrp:                         5047,
						ArpPacketsAreDroppedAsDriverCallFailed:                   5048,
						ArpPacketsAreDroppedAsSourceIsNotValidated:               5049,
						ArpSystemMax:  5050,
						ArpPublicMax:  5051,
						ArpIriMax:     5052,
						ArpMgtMax:     5053,
						ArpPublicCnt:  5054,
						ArpIriCnt:     5055,
						ArpMgtCnt:     5056,
						ArpSystemDrop: 5057,
						ArpPublicDrop: 5058,
						ArpIriDrop:    5059,
						ArpMgtDrop:    5060,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in ARP testing due to: ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in ARP testing due to: ", err)
			}

			result.Statistics.Arp.Text = ""
			assert.Equal(t, tc.expect.Statistics.Arp, result.Statistics.Arp, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}

func TestStatisticsICMPUnmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}
	tests := []testCase{
		{
			name:    "complete_icmp_statistics",
			xmlFile: "testsFiles/ICMP/ICMPTestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Icmp: ICMP{
						DropsDueToRateLimit:                        6000,
						CallsToIcmpError:                           6001,
						ErrorsNotGeneratedBecauseOldMessageWasIcmp: 6002,
						Histogram: []ICMPHistogram{
							{
								IcmpEchoReply:                    6003,
								DestinationUnreachable:           6004,
								IcmpEcho:                         6005,
								TimeStampReply:                   6006,
								TimeExceeded:                     6007,
								TimeStamp:                        6008,
								AddressMaskRequest:               6009,
								AnEndpointChangedItsCookieSecret: 6010,
							},
							{
								IcmpEchoReply:                    6011,
								DestinationUnreachable:           6012,
								IcmpEcho:                         6013,
								TimeStampReply:                   6014,
								TimeExceeded:                     6015,
								TimeStamp:                        6016,
								AddressMaskRequest:               6017,
								AnEndpointChangedItsCookieSecret: 6018,
							},
						},
						MessagesWithBadCodeFields:                                6019,
						MessagesLessThanTheMinimumLength:                         6020,
						MessagesWithBadChecksum:                                  6021,
						MessagesWithBadSourceAddress:                             6022,
						MessagesWithBadLength:                                    6023,
						EchoDropsWithBroadcastOrMulticastDestinatonAddress:       6024,
						TimestampDropsWithBroadcastOrMulticastDestinationAddress: 6025,
						MessageResponsesGenerated:                                6026,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in ICMP testing due to: ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in ICMP testing due to: ", err)
			}
			for i, _ := range result.Statistics.Icmp.Histogram {
				result.Statistics.Icmp.Histogram[i].Text = ""
				result.Statistics.Icmp.Histogram[i].TypeOfHistogram = ""
			}
			result.Statistics.Icmp.Text = ""
			assert.Equal(t, tc.expect.Statistics.Icmp, result.Statistics.Icmp, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}

func TestStatisticsICMP6Unmarshaling(t *testing.T) {
	type testCase struct {
		name    string
		xmlFile string
		expect  SystemStatistics
	}
	tests := []testCase{
		{
			name:    "complete_icmp6_statistics",
			xmlFile: "testsFiles/ICMP6/ICMP6TestDataCase1.xml",
			expect: SystemStatistics{
				Statistics: Statistics{
					Icmp6: ICMP6{
						CallsToIcmp6Error: 7000,
						ErrorsNotGeneratedBecauseOldMessageWasIcmpError: 7001,
						ErrorsNotGeneratedBecauseRateLimitation:         7002,
						OutputHistogram: ICMP6OutputHistogram{
							UnreachableIcmp6Packets: 7003,
							Icmp6Echo:               7004,
							Icmp6EchoReply:          7005,
							NeighborSolicitation:    7006,
							NeighborAdvertisement:   7007,
						},
						Icmp6MessagesWithBadCodeFields: 7008,
						MessagesLessThanMinimumLength:  7009,
						BadChecksums:                   7010,
						Icmp6MessagesWithBadLength:     7011,
						InputHistogram: ICMP6InputHistogram{
							UnreachableIcmp6Packets:        7012,
							PacketTooBig:                   7013,
							TimeExceededIcmp6Packets:       7014,
							Icmp6Echo:                      7015,
							Icmp6EchoReply:                 7016,
							RouterSolicitationIcmp6Packets: 7017,
							NeighborSolicitation:           7018,
							NeighborAdvertisement:          7019,
						},
						NoRoute:                               7020,
						AdministrativelyProhibited:            7021,
						BeyondScope:                           7022,
						AddressUnreachable:                    7023,
						PortUnreachable:                       7024,
						PacketTooBig:                          7025,
						TimeExceedTransit:                     7026,
						TimeExceedReassembly:                  7027,
						ErroneousHeaderField:                  7028,
						UnrecognizedNextHeader:                7029,
						UnrecognizedOption:                    7030,
						Redirect:                              7031,
						Unknown:                               7032,
						Icmp6MessageResponsesGenerated:        7033,
						MessagesWithTooManyNdOptions:          7034,
						NdSystemMax:                           7035,
						NdPublicMax:                           7036,
						NdIriMax:                              7037,
						NdMgtMax:                              7038,
						NdPublicCnt:                           7039,
						NdIriCnt:                              7040,
						NdMgtCnt:                              7041,
						NdSystemDrop:                          7042,
						NdPublicDrop:                          7043,
						NdIriDrop:                             7044,
						NdMgtDrop:                             7045,
						Nd6NdpProxyRequests:                   7046,
						Nd6DadProxyRequests:                   7047,
						Nd6NdpProxyResponses:                  7048,
						Nd6DadProxyConflicts:                  7049,
						Nd6DupProxyResponses:                  7050,
						Nd6NdpProxyResolveCnt:                 7051,
						Nd6DadProxyResolveCnt:                 7052,
						Nd6DadProxyEqmacDrop:                  7053,
						Nd6DadProxyNomacDrop:                  7054,
						Nd6NdpProxyUnrRequests:                7055,
						Nd6DadProxyUnrRequests:                7056,
						Nd6NdpProxyUnrResponses:               7057,
						Nd6DadProxyUnrConflicts:               7058,
						Nd6DadProxyUnrResponses:               7059,
						Nd6NdpProxyUnrResolveCnt:              7060,
						Nd6DadProxyUnrResolveCnt:              7061,
						Nd6DadProxyUnrEqportDrop:              7062,
						Nd6DadProxyUnrNomacDrop:               7063,
						Nd6RequestsDroppedOnEntry:             7064,
						Nd6RequestsDroppedDuringRetry:         7065,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fc, err := os.ReadFile(tc.xmlFile)
			if err != nil {
				log.Fatal("failed to read xml file in ICMP6 testing due to: ", err)
			}
			var result SystemStatistics
			err = xml.Unmarshal(fc, &result)
			if err != nil {
				log.Fatal("failed to unmarshal xml file in ICMP6 testing due to: ", err)
			}
			result.Statistics.Icmp6.Text = ""
			result.Statistics.Icmp6.HistogramOfErrorMessagesToBeGenerated = ""
			result.Statistics.Icmp6.InputHistogram.Text = ""
			result.Statistics.Icmp6.OutputHistogram.Text = ""
			result.Statistics.Icmp6.InputHistogram.HistogramType = ""
			result.Statistics.Icmp6.OutputHistogram.HistogramType = ""
			result.Statistics.Icmp6.InputHistogram.Style = ""
			result.Statistics.Icmp6.OutputHistogram.Style = ""
			assert.Equal(t, tc.expect.Statistics.Icmp6, result.Statistics.Icmp6, tc.name)
			assert.NoError(t, err, "unmarshal should not return error")
		})
	}
}
