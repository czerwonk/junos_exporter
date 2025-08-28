package systemstatistics

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsIPv4Unmarshaling(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected SystemStatistics
	}{
		{
			name: "complete_ipv4_statistics",
			xmlInput: `<rpc-reply junos:style="normal">
				<statistics>
					<ip>
						<packets-received>1000</packets-received>
						<bad-header-checksums>5</bad-header-checksums>
						<packets-with-size-smaller-than-minimum>10</packets-with-size-smaller-than-minimum>
						<packets-with-data-size-less-than-datalength>2</packets-with-data-size-less-than-datalength>
						<packets-with-header-length-less-than-data-size>3</packets-with-header-length-less-than-data-size>
						<packets-with-data-length-less-than-headerlength>1</packets-with-data-length-less-than-headerlength>
						<packets-with-incorrect-version-number>0</packets-with-incorrect-version-number>
						<packets-destined-to-dead-next-hop>0</packets-destined-to-dead-next-hop>
						<fragments-received>50</fragments-received>
						<fragments-dropped-due-to-outofspace-or-dup>2</fragments-dropped-due-to-outofspace-or-dup>
						<fragments-dropped-due-to-queueoverflow>1</fragments-dropped-due-to-queueoverflow>
						<fragments-dropped-after-timeout>0</fragments-dropped-after-timeout>
						<packets-reassembled-ok>48</packets-reassembled-ok>
						<packets-for-this-host>500</packets-for-this-host>
						<packets-for-unknown-or-unsupported-protocol>5</packets-for-unknown-or-unsupported-protocol>
						<packets-forwarded>400</packets-forwarded>
						<packets-not-forwardable>10</packets-not-forwardable>
						<redirects-sent>2</redirects-sent>
						<packets-sent-from-this-host>800</packets-sent-from-this-host>
						<packets-sent-with-fabricated-ip-header>0</packets-sent-with-fabricated-ip-header>
						<output-packets-dropped-due-to-no-bufs>3</output-packets-dropped-due-to-no-bufs>
						<output-packets-discarded-due-to-no-route>1</output-packets-discarded-due-to-no-route>
						<output-datagrams-fragmented>20</output-datagrams-fragmented>
						<fragments-created>40</fragments-created>
						<datagrams-that-can-not-be-fragmented>2</datagrams-that-can-not-be-fragmented>
						<packets-with-bad-options>1</packets-with-bad-options>
						<packets-with-options-handled-without-error>15</packets-with-options-handled-without-error>
						<strict-source-and-record-route-options>0</strict-source-and-record-route-options>
						<loose-source-and-record-route-options>2</loose-source-and-record-route-options>
						<record-route-options>5</record-route-options>
						<timestamp-options>3</timestamp-options>
						<timestamp-and-address-options>1</timestamp-and-address-options>
						<timestamp-and-prespecified-address-options>0</timestamp-and-prespecified-address-options>
						<option-packets-dropped-due-to-rate-limit>0</option-packets-dropped-due-to-rate-limit>
						<router-alert-options>4</router-alert-options>
						<multicast-packets-dropped>8</multicast-packets-dropped>
						<packets-dropped>12</packets-dropped>
						<transit-re-packets-dropped-on-mgmt-interface>0</transit-re-packets-dropped-on-mgmt-interface>
						<packets-used-first-nexthop-in-ecmp-unilist>25</packets-used-first-nexthop-in-ecmp-unilist>
						<incoming-ttpoip-packets-received>100</incoming-ttpoip-packets-received>
						<incoming-ttpoip-packets-dropped>2</incoming-ttpoip-packets-dropped>
						<outgoing-ttpoip-packets-sent>95</outgoing-ttpoip-packets-sent>
						<outgoing-ttpoip-packets-dropped>1</outgoing-ttpoip-packets-dropped>
						<incoming-rawip-packets-dropped-no-socket-buffer>3</incoming-rawip-packets-dropped-no-socket-buffer>
						<incoming-virtual-node-packets-delivered>200</incoming-virtual-node-packets-delivered>
					</ip>
				</statistics>
				<cli>
					<banner>user@router></banner>
				</cli>
			</rpc-reply>`,
			expected: SystemStatistics{
				Statistics: struct {
					Text string `xml:",chardata"`
					Tcp  struct {
						Text                                             string  `xml:",chardata"`
						PacketsSent                                      float64 `xml:"packets-sent"`
						SentDataPackets                                  float64 `xml:"sent-data-packets"`
						DataPacketsBytes                                 float64 `xml:"data-packets-bytes"`
						SentDataPacketsRetransmitted                     float64 `xml:"sent-data-packets-retransmitted"`
						RetransmittedBytes                               float64 `xml:"retransmitted-bytes"`
						SentDataUnnecessaryRetransmitted                 float64 `xml:"sent-data-unnecessary-retransmitted"`
						SentResendsByMtuDiscovery                        float64 `xml:"sent-resends-by-mtu-discovery"`
						SentAckOnlyPackets                               float64 `xml:"sent-ack-only-packets"`
						SentPacketsDelayed                               float64 `xml:"sent-packets-delayed"`
						SentUrgOnlyPackets                               float64 `xml:"sent-urg-only-packets"`
						SentWindowProbePackets                           float64 `xml:"sent-window-probe-packets"`
						SentWindowUpdatePackets                          float64 `xml:"sent-window-update-packets"`
						SentControlPackets                               float64 `xml:"sent-control-packets"`
						PacketsReceived                                  float64 `xml:"packets-received"`
						ReceivedAcks                                     float64 `xml:"received-acks"`
						AcksBytes                                        float64 `xml:"acks-bytes"`
						ReceivedDuplicateAcks                            float64 `xml:"received-duplicate-acks"`
						ReceivedAcksForUnsentData                        float64 `xml:"received-acks-for-unsent-data"`
						PacketsReceivedInSequence                        float64 `xml:"packets-received-in-sequence"`
						InSequenceBytes                                  float64 `xml:"in-sequence-bytes"`
						ReceivedCompletelyDuplicatePacket                float64 `xml:"received-completely-duplicate-packet"`
						DuplicateInBytes                                 float64 `xml:"duplicate-in-bytes"`
						ReceivedOldDuplicatePackets                      float64 `xml:"received-old-duplicate-packets"`
						ReceivedPacketsWithSomeDupliacteData             float64 `xml:"received-packets-with-some-dupliacte-data"`
						SomeDuplicateInBytes                             float64 `xml:"some-duplicate-in-bytes"`
						ReceivedOutOfOrderPackets                        float64 `xml:"received-out-of-order-packets"`
						OutOfOrderInBytes                                float64 `xml:"out-of-order-in-bytes"`
						ReceivedPacketsOfDataAfterWindow                 float64 `xml:"received-packets-of-data-after-window"`
						Bytes                                            float64 `xml:"bytes"`
						ReceivedWindowProbes                             float64 `xml:"received-window-probes"`
						ReceivedWindowUpdatePackets                      float64 `xml:"received-window-update-packets"`
						PacketsReceivedAfterClose                        float64 `xml:"packets-received-after-close"`
						ReceivedDiscardedForBadChecksum                  float64 `xml:"received-discarded-for-bad-checksum"`
						ReceivedDiscardedForBadHeaderOffset              float64 `xml:"received-discarded-for-bad-header-offset"`
						ReceivedDiscardedBecausePacketTooShort           float64 `xml:"received-discarded-because-packet-too-short"`
						ConnectionRequests                               float64 `xml:"connection-requests"`
						ConnectionAccepts                                float64 `xml:"connection-accepts"`
						BadConnectionAttempts                            float64 `xml:"bad-connection-attempts"`
						ListenQueueOverflows                             float64 `xml:"listen-queue-overflows"`
						BadRstWindow                                     float64 `xml:"bad-rst-window"`
						ConnectionsEstablished                           float64 `xml:"connections-established"`
						ConnectionsClosed                                float64 `xml:"connections-closed"`
						Drops                                            float64 `xml:"drops"`
						ConnectionsUpdatedRttOnClose                     float64 `xml:"connections-updated-rtt-on-close"`
						ConnectionsUpdatedVarianceOnClose                float64 `xml:"connections-updated-variance-on-close"`
						ConnectionsUpdatedSsthreshOnClose                float64 `xml:"connections-updated-ssthresh-on-close"`
						EmbryonicConnectionsDropped                      float64 `xml:"embryonic-connections-dropped"`
						SegmentsUpdatedRtt                               float64 `xml:"segments-updated-rtt"`
						Attempts                                         float64 `xml:"attempts"`
						RetransmitTimeouts                               float64 `xml:"retransmit-timeouts"`
						ConnectionsDroppedByRetransmitTimeout            float64 `xml:"connections-dropped-by-retransmit-timeout"`
						PersistTimeouts                                  float64 `xml:"persist-timeouts"`
						ConnectionsDroppedByPersistTimeout               float64 `xml:"connections-dropped-by-persist-timeout"`
						KeepaliveTimeouts                                float64 `xml:"keepalive-timeouts"`
						KeepaliveProbesSent                              float64 `xml:"keepalive-probes-sent"`
						KeepaliveConnectionsDropped                      float64 `xml:"keepalive-connections-dropped"`
						AckHeaderPredictions                             float64 `xml:"ack-header-predictions"`
						DataPacketHeaderPredictions                      float64 `xml:"data-packet-header-predictions"`
						SyncacheEntriesAdded                             float64 `xml:"syncache-entries-added"`
						Retransmitted                                    float64 `xml:"retransmitted"`
						Dupsyn                                           float64 `xml:"dupsyn"`
						Dropped                                          float64 `xml:"dropped"`
						Completed                                        float64 `xml:"completed"`
						BucketOverflow                                   float64 `xml:"bucket-overflow"`
						CacheOverflow                                    float64 `xml:"cache-overflow"`
						Reset                                            float64 `xml:"reset"`
						Stale                                            float64 `xml:"stale"`
						Aborted                                          float64 `xml:"aborted"`
						Badack                                           float64 `xml:"badack"`
						Unreach                                          float64 `xml:"unreach"`
						ZoneFailures                                     float64 `xml:"zone-failures"`
						CookiesSent                                      float64 `xml:"cookies-sent"`
						CookiesReceived                                  float64 `xml:"cookies-received"`
						SackRecoveryEpisodes                             float64 `xml:"sack-recovery-episodes"`
						SegmentRetransmits                               float64 `xml:"segment-retransmits"`
						ByteRetransmits                                  float64 `xml:"byte-retransmits"`
						SackOptionsReceived                              float64 `xml:"sack-options-received"`
						SackOpitionsSent                                 float64 `xml:"sack-opitions-sent"`
						SackScoreboardOverflow                           float64 `xml:"sack-scoreboard-overflow"`
						AcksSentInResponseButNotExactRsts                float64 `xml:"acks-sent-in-response-but-not-exact-rsts"`
						AcksSentInResponseToSynsOnEstablishedConnections float64 `xml:"acks-sent-in-response-to-syns-on-established-connections"`
						RcvPacketsDroppedDueToBadAddress                 float64 `xml:"rcv-packets-dropped-due-to-bad-address"`
						OutOfSequenceSegmentDrops                        float64 `xml:"out-of-sequence-segment-drops"`
						RstPackets                                       float64 `xml:"rst-packets"`
						IcmpPacketsIgnored                               float64 `xml:"icmp-packets-ignored"`
						SendPacketsDropped                               float64 `xml:"send-packets-dropped"`
						RcvPacketsDropped                                float64 `xml:"rcv-packets-dropped"`
						OutgoingSegmentsDropped                          float64 `xml:"outgoing-segments-dropped"`
						ReceivedSynfinDropped                            float64 `xml:"received-synfin-dropped"`
						ReceivedIpsecDropped                             float64 `xml:"received-ipsec-dropped"`
						ReceivedMacDropped                               float64 `xml:"received-mac-dropped"`
						ReceivedMinttlExceeded                           float64 `xml:"received-minttl-exceeded"`
						ListenstateBadflagsDropped                       float64 `xml:"listenstate-badflags-dropped"`
						FinwaitstateBadflagsDropped                      float64 `xml:"finwaitstate-badflags-dropped"`
						ReceivedDosAttack                                float64 `xml:"received-dos-attack"`
						ReceivedBadSynack                                float64 `xml:"received-bad-synack"`
						SyncacheZoneFull                                 float64 `xml:"syncache-zone-full"`
						ReceivedRstFirewallfilter                        float64 `xml:"received-rst-firewallfilter"`
						ReceivedNoackTimewait                            float64 `xml:"received-noack-timewait"`
						ReceivedNoTimewaitState                          float64 `xml:"received-no-timewait-state"`
						ReceivedRstTimewaitState                         float64 `xml:"received-rst-timewait-state"`
						ReceivedTimewaitDrops                            float64 `xml:"received-timewait-drops"`
						ReceivedBadaddrTimewaitState                     float64 `xml:"received-badaddr-timewait-state"`
						ReceivedAckoffInSynSentrcvd                      float64 `xml:"received-ackoff-in-syn-sentrcvd"`
						ReceivedBadaddrFirewall                          float64 `xml:"received-badaddr-firewall"`
						ReceivedNosynSynSent                             float64 `xml:"received-nosyn-syn-sent"`
						ReceivedBadrstSynSent                            float64 `xml:"received-badrst-syn-sent"`
						ReceivedBadrstListenState                        float64 `xml:"received-badrst-listen-state"`
						OptionMaxsegmentLength                           float64 `xml:"option-maxsegment-length"`
						OptionWindowLength                               float64 `xml:"option-window-length"`
						OptionTimestampLength                            float64 `xml:"option-timestamp-length"`
						OptionMd5Length                                  float64 `xml:"option-md5-length"`
						OptionAuthLength                                 float64 `xml:"option-auth-length"`
						OptionSackpermittedLength                        float64 `xml:"option-sackpermitted-length"`
						OptionSackLength                                 float64 `xml:"option-sack-length"`
						OptionAuthoptionLength                           float64 `xml:"option-authoption-length"`
					} `xml:"tcp"`
					Udp struct {
						Text                                              string  `xml:",chardata"`
						DatagramsReceived                                 float64 `xml:"datagrams-received"`
						DatagramsWithIncompleteHeader                     float64 `xml:"datagrams-with-incomplete-header"`
						DatagramsWithBadDatalengthField                   float64 `xml:"datagrams-with-bad-datalength-field"`
						DatagramsWithBadChecksum                          float64 `xml:"datagrams-with-bad-checksum"`
						DatagramsDroppedDueToNoSocket                     float64 `xml:"datagrams-dropped-due-to-no-socket"`
						BroadcastOrMulticastDatagramsDroppedDueToNoSocket float64 `xml:"broadcast-or-multicast-datagrams-dropped-due-to-no-socket"`
						DatagramsDroppedDueToFullSocketBuffers            float64 `xml:"datagrams-dropped-due-to-full-socket-buffers"`
						DatagramsNotForHashedPcb                          float64 `xml:"datagrams-not-for-hashed-pcb"`
						DatagramsDelivered                                float64 `xml:"datagrams-delivered"`
						DatagramsOutput                                   float64 `xml:"datagrams-output"`
					} `xml:"udp"`
					Ip struct {
						Text                                      string  `xml:",chardata"`
						PacketsReceived                           float64 `xml:"packets-received"`
						BadHeaderChecksums                        float64 `xml:"bad-header-checksums"`
						PacketsWithSizeSmallerThanMinimum         float64 `xml:"packets-with-size-smaller-than-minimum"`
						PacketsWithDataSizeLessThanDatalength     float64 `xml:"packets-with-data-size-less-than-datalength"`
						PacketsWithHeaderLengthLessThanDataSize   float64 `xml:"packets-with-header-length-less-than-data-size"`
						PacketsWithDataLengthLessThanHeaderlength float64 `xml:"packets-with-data-length-less-than-headerlength"`
						PacketsWithIncorrectVersionNumber         float64 `xml:"packets-with-incorrect-version-number"`
						PacketsDestinedToDeadNextHop              float64 `xml:"packets-destined-to-dead-next-hop"`
						FragmentsReceived                         float64 `xml:"fragments-received"`
						FragmentsDroppedDueToOutofspaceOrDup      float64 `xml:"fragments-dropped-due-to-outofspace-or-dup"`
						FragmentsDroppedDueToQueueoverflow        float64 `xml:"fragments-dropped-due-to-queueoverflow"`
						FragmentsDroppedAfterTimeout              float64 `xml:"fragments-dropped-after-timeout"`
						PacketsReassembledOk                      float64 `xml:"packets-reassembled-ok"`
						PacketsForThisHost                        float64 `xml:"packets-for-this-host"`
						PacketsForUnknownOrUnsupportedProtocol    float64 `xml:"packets-for-unknown-or-unsupported-protocol"`
						PacketsForwarded                          float64 `xml:"packets-forwarded"`
						PacketsNotForwardable                     float64 `xml:"packets-not-forwardable"`
						RedirectsSent                             float64 `xml:"redirects-sent"`
						PacketsSentFromThisHost                   float64 `xml:"packets-sent-from-this-host"`
						PacketsSentWithFabricatedIpHeader         float64 `xml:"packets-sent-with-fabricated-ip-header"`
						OutputPacketsDroppedDueToNoBufs           float64 `xml:"output-packets-dropped-due-to-no-bufs"`
						OutputPacketsDiscardedDueToNoRoute        float64 `xml:"output-packets-discarded-due-to-no-route"`
						OutputDatagramsFragmented                 float64 `xml:"output-datagrams-fragmented"`
						FragmentsCreated                          float64 `xml:"fragments-created"`
						DatagramsThatCanNotBeFragmented           float64 `xml:"datagrams-that-can-not-be-fragmented"`
						PacketsWithBadOptions                     float64 `xml:"packets-with-bad-options"`
						PacketsWithOptionsHandledWithoutError     float64 `xml:"packets-with-options-handled-without-error"`
						StrictSourceAndRecordRouteOptions         float64 `xml:"strict-source-and-record-route-options"`
						LooseSourceAndRecordRouteOptions          float64 `xml:"loose-source-and-record-route-options"`
						RecordRouteOptions                        float64 `xml:"record-route-options"`
						TimestampOptions                          float64 `xml:"timestamp-options"`
						TimestampAndAddressOptions                float64 `xml:"timestamp-and-address-options"`
						TimestampAndPrespecifiedAddressOptions    float64 `xml:"timestamp-and-prespecified-address-options"`
						OptionPacketsDroppedDueToRateLimit        float64 `xml:"option-packets-dropped-due-to-rate-limit"`
						RouterAlertOptions                        float64 `xml:"router-alert-options"`
						MulticastPacketsDropped                   float64 `xml:"multicast-packets-dropped"`
						PacketsDropped                            float64 `xml:"packets-dropped"`
						TransitRePacketsDroppedOnMgmtInterface    float64 `xml:"transit-re-packets-dropped-on-mgmt-interface"`
						PacketsUsedFirstNexthopInEcmpUnilist      float64 `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
						IncomingTtpoipPacketsReceived             float64 `xml:"incoming-ttpoip-packets-received"`
						IncomingTtpoipPacketsDropped              float64 `xml:"incoming-ttpoip-packets-dropped"`
						OutgoingTtpoipPacketsSent                 float64 `xml:"outgoing-ttpoip-packets-sent"`
						OutgoingTtpoipPacketsDropped              float64 `xml:"outgoing-ttpoip-packets-dropped"`
						IncomingRawipPacketsDroppedNoSocketBuffer float64 `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
						IncomingVirtualNodePacketsDelivered       float64 `xml:"incoming-virtual-node-packets-delivered"`
					} `xml:"ip"`
					Icmp struct {
						Text                                       string `xml:",chardata"`
						DropsDueToRateLimit                        string `xml:"drops-due-to-rate-limit"`
						CallsToIcmpError                           string `xml:"calls-to-icmp-error"`
						ErrorsNotGeneratedBecauseOldMessageWasIcmp string `xml:"errors-not-generated-because-old-message-was-icmp"`
						Histogram                                  []struct {
							Text                             string `xml:",chardata"`
							TypeOfHistogram                  string `xml:"type-of-histogram"`
							IcmpEchoReply                    string `xml:"icmp-echo-reply"`
							DestinationUnreachable           string `xml:"destination-unreachable"`
							IcmpEcho                         string `xml:"icmp-echo"`
							TimeStampReply                   string `xml:"time-stamp-reply"`
							TimeExceeded                     string `xml:"time-exceeded"`
							TimeStamp                        string `xml:"time-stamp"`
							AddressMaskRequest               string `xml:"address-mask-request"`
							AnEndpointChangedItsCookiesecret string `xml:"an-endpoint-changed-its-cookiesecret"`
						} `xml:"histogram"`
						MessagesWithBadCodeFields                                string `xml:"messages-with-bad-code-fields"`
						MessagesLessThanTheMinimumLength                         string `xml:"messages-less-than-the-minimum-length"`
						MessagesWithBadChecksum                                  string `xml:"messages-with-bad-checksum"`
						MessagesWithBadSourceAddress                             string `xml:"messages-with-bad-source-address"`
						MessagesWithBadLength                                    string `xml:"messages-with-bad-length"`
						EchoDropsWithBroadcastOrMulticastDestinatonAddress       string `xml:"echo-drops-with-broadcast-or-multicast-destinaton-address"`
						TimestampDropsWithBroadcastOrMulticastDestinationAddress string `xml:"timestamp-drops-with-broadcast-or-multicast-destination-address"`
						MessageResponsesGenerated                                string `xml:"message-responses-generated"`
					} `xml:"icmp"`
					Arp struct {
						Text                                                     string `xml:",chardata"`
						DatagramsReceived                                        string `xml:"datagrams-received"`
						ArpRequestsReceived                                      string `xml:"arp-requests-received"`
						ArpRepliesReceived                                       string `xml:"arp-replies-received"`
						ResolutionRequestReceived                                string `xml:"resolution-request-received"`
						ResolutionRequestDropped                                 string `xml:"resolution-request-dropped"`
						UnrestrictedProxyRequests                                string `xml:"unrestricted-proxy-requests"`
						RestrictedProxyRequests                                  string `xml:"restricted-proxy-requests"`
						ReceivedProxyRequests                                    string `xml:"received-proxy-requests"`
						ProxyRequestsNotProxied                                  string `xml:"proxy-requests-not-proxied"`
						RestrictedProxyRequestsNotProxied                        string `xml:"restricted-proxy-requests-not-proxied"`
						DatagramsWithBogusInterface                              string `xml:"datagrams-with-bogus-interface"`
						DatagramsWithIncorrectLength                             string `xml:"datagrams-with-incorrect-length"`
						DatagramsForNonIpProtocol                                string `xml:"datagrams-for-non-ip-protocol"`
						DatagramsWithUnsupportedOpcode                           string `xml:"datagrams-with-unsupported-opcode"`
						DatagramsWithBadProtocolAddressLength                    string `xml:"datagrams-with-bad-protocol-address-length"`
						DatagramsWithBadHardwareAddressLength                    string `xml:"datagrams-with-bad-hardware-address-length"`
						DatagramsWithMulticastSourceAddress                      string `xml:"datagrams-with-multicast-source-address"`
						DatagramsWithMulticastTargetAddress                      string `xml:"datagrams-with-multicast-target-address"`
						DatagramsWithMyOwnHardwareAddress                        string `xml:"datagrams-with-my-own-hardware-address"`
						DatagramsForAnAddressNotOnTheInterface                   string `xml:"datagrams-for-an-address-not-on-the-interface"`
						DatagramsWithABroadcastSourceAddress                     string `xml:"datagrams-with-a-broadcast-source-address"`
						DatagramsWithSourceAddressDuplicateToMine                string `xml:"datagrams-with-source-address-duplicate-to-mine"`
						DatagramsWhichWereNotForMe                               string `xml:"datagrams-which-were-not-for-me"`
						PacketsDiscardedWaitingForResolution                     string `xml:"packets-discarded-waiting-for-resolution"`
						PacketsSentAfterWaitingForResolution                     string `xml:"packets-sent-after-waiting-for-resolution"`
						ArpRequestsSent                                          string `xml:"arp-requests-sent"`
						ArpRepliesSent                                           string `xml:"arp-replies-sent"`
						RequestsForMemoryDenied                                  string `xml:"requests-for-memory-denied"`
						RequestsDroppedOnEntry                                   string `xml:"requests-dropped-on-entry"`
						RequestsDroppedDuringRetry                               string `xml:"requests-dropped-during-retry"`
						RequestsDroppedDueToInterfaceDeletion                    string `xml:"requests-dropped-due-to-interface-deletion"`
						RequestsOnUnnumberedInterfaces                           string `xml:"requests-on-unnumbered-interfaces"`
						NewRequestsOnUnnumberedInterfaces                        string `xml:"new-requests-on-unnumbered-interfaces"`
						RepliesFromUnnumberedInterfaces                          string `xml:"replies-from-unnumbered-interfaces"`
						RequestsOnUnnumberedInterfaceWithNonSubnettedDonor       string `xml:"requests-on-unnumbered-interface-with-non-subnetted-donor"`
						RepliesFromUnnumberedInterfaceWithNonSubnettedDonor      string `xml:"replies-from-unnumbered-interface-with-non-subnetted-donor"`
						ArpPacketsRejectedAsFamilyIsConfiguredWithDenyArp        string `xml:"arp-packets-rejected-as-family-is-configured-with-deny-arp"`
						ArpResponsePacketsAreRejectedOnMcAeIclInterface          string `xml:"arp-response-packets-are-rejected-on-mc-ae-icl-interface"`
						ArpRepliesAreRejectedAsSourceAndDestinationIsSame        string `xml:"arp-replies-are-rejected-as-source-and-destination-is-same"`
						ArpProbeForProxyAddressReachableFromTheIncomingInterface string `xml:"arp-probe-for-proxy-address-reachable-from-the-incoming-interface"`
						ArpRequestDiscardedForVrrpSourceAddress                  string `xml:"arp-request-discarded-for-vrrp-source-address"`
						SelfArpRequestPacketReceivedOnIrbInterface               string `xml:"self-arp-request-packet-received-on-irb-interface"`
						ProxyArpRequestDiscardedAsSourceIpIsAProxyTarget         string `xml:"proxy-arp-request-discarded-as-source-ip-is-a-proxy-target"`
						ArpPacketsAreDroppedAsNexthopAllocationFailed            string `xml:"arp-packets-are-dropped-as-nexthop-allocation-failed"`
						ArpPacketsReceivedFromPeerVrrpRouterAndDiscarded         string `xml:"arp-packets-received-from-peer-vrrp-router-and-discarded"`
						ArpPacketsAreRejectedAsTargetIpArpResolveIsInProgress    string `xml:"arp-packets-are-rejected-as-target-ip-arp-resolve-is-in-progress"`
						GratArpPacketsAreIgnoredAsMacAddressIsNotChanged         string `xml:"grat-arp-packets-are-ignored-as-mac-address-is-not-changed"`
						ArpPacketsAreDroppedFromPeerVrrp                         string `xml:"arp-packets-are-dropped-from-peer-vrrp"`
						ArpPacketsAreDroppedAsDriverCallFailed                   string `xml:"arp-packets-are-dropped-as-driver-call-failed"`
						ArpPacketsAreDroppedAsSourceIsNotValidated               string `xml:"arp-packets-are-dropped-as-source-is-not-validated"`
						ArpSystemMax                                             string `xml:"arp-system-max"`
						ArpPublicMax                                             string `xml:"arp-public-max"`
						ArpIriMax                                                string `xml:"arp-iri-max"`
						ArpMgtMax                                                string `xml:"arp-mgt-max"`
						ArpPublicCnt                                             string `xml:"arp-public-cnt"`
						ArpIriCnt                                                string `xml:"arp-iri-cnt"`
						ArpMgtCnt                                                string `xml:"arp-mgt-cnt"`
						ArpSystemDrop                                            string `xml:"arp-system-drop"`
						ArpPublicDrop                                            string `xml:"arp-public-drop"`
						ArpIriDrop                                               string `xml:"arp-iri-drop"`
						ArpMgtDrop                                               string `xml:"arp-mgt-drop"`
					} `xml:"arp"`
					Ip6 struct {
						Text                                  string `xml:",chardata"`
						TotalPacketsReceived                  float64 `xml:"total-packets-received"`
						Ip6PacketsWithSizeSmallerThanMinimum  float64 `xml:"ip6-packets-with-size-smaller-than-minimum"`
						PacketsWithDatasizeLessThanDataLength float64 `xml:"packets-with-datasize-less-than-data-length"`
						Ip6PacketsWithBadOptions              float64 `xml:"ip6-packets-with-bad-options"`
						Ip6PacketsWithIncorrectVersionNumber  float64 `xml:"ip6-packets-with-incorrect-version-number"`
						Ip6FragmentsReceived                  float64 `xml:"ip6-fragments-received"`
						DuplicateOrOutOfSpaceFragmentsDropped float64 `xml:"duplicate-or-out-of-space-fragments-dropped"`
						Ip6FragmentsDroppedAfterTimeout       float64 `xml:"ip6-fragments-dropped-after-timeout"`
						FragmentsThatExceededLimit            float64 `xml:"fragments-that-exceeded-limit"`
						Ip6PacketsReassembledOk               float64 `xml:"ip6-packets-reassembled-ok"`
						Ip6PacketsForThisHost                 float64 `xml:"ip6-packets-for-this-host"`
						Ip6PacketsForwarded                   float64 `xml:"ip6-packets-forwarded"`
						Ip6PacketsNotForwardable              float64 `xml:"ip6-packets-not-forwardable"`
						Ip6RedirectsSent                      float64 `xml:"ip6-redirects-sent"`
						Ip6PacketsSentFromThisHost            float64 `xml:"ip6-packets-sent-from-this-host"`
						Ip6PacketsSentWithFabricatedIpHeader  float64 `xml:"ip6-packets-sent-with-fabricated-ip-header"`
						Ip6OutputPacketsDroppedDueToNoBufs    float64 `xml:"ip6-output-packets-dropped-due-to-no-bufs"`
						Ip6OutputPacketsDiscardedDueToNoRoute float64 `xml:"ip6-output-packets-discarded-due-to-no-route"`
						Ip6OutputDatagramsFragmented          float64 `xml:"ip6-output-datagrams-fragmented"`
						Ip6FragmentsCreated                   float64 `xml:"ip6-fragments-created"`
						Ip6DatagramsThatCanNotBeFragmented    float64 `xml:"ip6-datagrams-that-can-not-be-fragmented"`
						PacketsThatViolatedScopeRules         float64 `xml:"packets-that-violated-scope-rules"`
						MulticastPacketsWhichWeDoNotJoin      float64 `xml:"multicast-packets-which-we-do-not-join"`
						Histogram                             float64 `xml:"histogram"`
						Ip6nhTcp                              float64 `xml:"ip6nh-tcp"`
						Ip6nhUdp                              float64 `xml:"ip6nh-udp"`
						Ip6nhIcmp6                            float64 `xml:"ip6nh-icmp6"`
						PacketsWhoseHeadersAreNotContinuous   float64 `xml:"packets-whose-headers-are-not-continuous"`
						TunnelingPacketsThatCanNotFindGif     float64 `xml:"tunneling-packets-that-can-not-find-gif"`
						PacketsDiscardedDueToTooMayHeaders    float64 `xml:"packets-discarded-due-to-too-may-headers"`
						FailuresOfSourceAddressSelection      float64 `xml:"failures-of-source-address-selection"`
						HeaderType                            []struct {
							Text                            string  `xml:",chardata"`
							HeaderForSourceAddressSelection string `xml:"header-for-source-address-selection"`
							LinkLocals                      float64 `xml:"link-locals"`
							Globals                         float64 `xml:"globals"`
							AddressScope                    float64 `xml:"address-scope"`
							HexValue                        float64 `xml:"hex-value"`
						} `xml:"header-type"`
						ForwardCacheHit                       float64 `xml:"forward-cache-hit"`
						ForwardCacheMiss                      float64 `xml:"forward-cache-miss"`
						Ip6PacketsDestinedToDeadNextHop       float64 `xml:"ip6-packets-destined-to-dead-next-hop"`
						Ip6OptionPacketsDroppedDueToRateLimit float64 `xml:"ip6-option-packets-dropped-due-to-rate-limit"`
						Ip6PacketsDropped                     float64 `xml:"ip6-packets-dropped"`
						PacketsDroppedDueToBadProtocol        float64 `xml:"packets-dropped-due-to-bad-protocol"`
						TransitRePacketDroppedOnMgmtInterface float64 `xml:"transit-re-packet-dropped-on-mgmt-interface"`
						PacketUsedFirstNexthopInEcmpUnilist   float64 `xml:"packet-used-first-nexthop-in-ecmp-unilist"`
					} `xml:"ip6"`
					Icmp6 struct {
						Text                                            string `xml:",chardata"`
						ProtocolName                                    string `xml:"protocol-name"`
						CallsToIcmp6Error                               string `xml:"calls-to-icmp6-error"`
						ErrorsNotGeneratedBecauseOldMessageWasIcmpError string `xml:"errors-not-generated-because-old-message-was-icmp-error"`
						ErrorsNotGeneratedBecauseRateLimitation         string `xml:"errors-not-generated-because-rate-limitation"`
						OutputHistogram                                 struct {
							Text                    string `xml:",chardata"`
							Style                   string `xml:"style,attr"`
							HistogramType           string `xml:"histogram-type"`
							UnreachableIcmp6Packets string `xml:"unreachable-icmp6-packets"`
							Icmp6Echo               string `xml:"icmp6-echo"`
							Icmp6EchoReply          string `xml:"icmp6-echo-reply"`
							NeighborSolicitation    string `xml:"neighbor-solicitation"`
							NeighborAdvertisement   string `xml:"neighbor-advertisement"`
						} `xml:"output-histogram"`
						Icmp6MessagesWithBadCodeFields string `xml:"icmp6-messages-with-bad-code-fields"`
						MessagesLessThanMinimumLength  string `xml:"messages-less-than-minimum-length"`
						BadChecksums                   string `xml:"bad-checksums"`
						Icmp6MessagesWithBadLength     string `xml:"icmp6-messages-with-bad-length"`
						InputHistogram                 struct {
							Text                           string `xml:",chardata"`
							Style                          string `xml:"style,attr"`
							HistogramType                  string `xml:"histogram-type"`
							UnreachableIcmp6Packets        string `xml:"unreachable-icmp6-packets"`
							PacketTooBig                   string `xml:"packet-too-big"`
							TimeExceededIcmp6Packets       string `xml:"time-exceeded-icmp6-packets"`
							Icmp6Echo                      string `xml:"icmp6-echo"`
							Icmp6EchoReply                 string `xml:"icmp6-echo-reply"`
							RouterSolicitationIcmp6Packets string `xml:"router-solicitation-icmp6-packets"`
							NeighborSolicitation           string `xml:"neighbor-solicitation"`
							NeighborAdvertisement          string `xml:"neighbor-advertisement"`
						} `xml:"input-histogram"`
						HistogramOfErrorMessagesToBeGenerated string `xml:"histogram-of-error-messages-to-be-generated"`
						NoRoute                               string `xml:"no-route"`
						AdministrativelyProhibited            string `xml:"administratively-prohibited"`
						BeyondScope                           string `xml:"beyond-scope"`
						AddressUnreachable                    string `xml:"address-unreachable"`
						PortUnreachable                       string `xml:"port-unreachable"`
						PacketTooBig                          string `xml:"packet-too-big"`
						TimeExceedTransit                     string `xml:"time-exceed-transit"`
						TimeExceedReassembly                  string `xml:"time-exceed-reassembly"`
						ErroneousHeaderField                  string `xml:"erroneous-header-field"`
						UnrecognizedNextHeader                string `xml:"unrecognized-next-header"`
						UnrecognizedOption                    string `xml:"unrecognized-option"`
						Redirect                              string `xml:"redirect"`
						Unknown                               string `xml:"unknown"`
						Icmp6MessageResponsesGenerated        string `xml:"icmp6-message-responses-generated"`
						MessagesWithTooManyNdOptions          string `xml:"messages-with-too-many-nd-options"`
						NdSystemMax                           string `xml:"nd-system-max"`
						NdPublicMax                           string `xml:"nd-public-max"`
						NdIriMax                              string `xml:"nd-iri-max"`
						NdMgtMax                              string `xml:"nd-mgt-max"`
						NdPublicCnt                           string `xml:"nd-public-cnt"`
						NdIriCnt                              string `xml:"nd-iri-cnt"`
						NdMgtCnt                              string `xml:"nd-mgt-cnt"`
						NdSystemDrop                          string `xml:"nd-system-drop"`
						NdPublicDrop                          string `xml:"nd-public-drop"`
						NdIriDrop                             string `xml:"nd-iri-drop"`
						NdMgtDrop                             string `xml:"nd-mgt-drop"`
						Nd6NdpProxyRequests                   string `xml:"nd6-ndp-proxy-requests"`
						Nd6DadProxyRequests                   string `xml:"nd6-dad-proxy-requests"`
						Nd6NdpProxyResponses                  string `xml:"nd6-ndp-proxy-responses"`
						Nd6DadProxyConflicts                  string `xml:"nd6-dad-proxy-conflicts"`
						Nd6DupProxyResponses                  string `xml:"nd6-dup-proxy-responses"`
						Nd6NdpProxyResolveCnt                 string `xml:"nd6-ndp-proxy-resolve-cnt"`
						Nd6DadProxyResolveCnt                 string `xml:"nd6-dad-proxy-resolve-cnt"`
						Nd6DadProxyEqmacDrop                  string `xml:"nd6-dad-proxy-eqmac-drop"`
						Nd6DadProxyNomacDrop                  string `xml:"nd6-dad-proxy-nomac-drop"`
						Nd6NdpProxyUnrRequests                string `xml:"nd6-ndp-proxy-unr-requests"`
						Nd6DadProxyUnrRequests                string `xml:"nd6-dad-proxy-unr-requests"`
						Nd6NdpProxyUnrResponses               string `xml:"nd6-ndp-proxy-unr-responses"`
						Nd6DadProxyUnrConflicts               string `xml:"nd6-dad-proxy-unr-conflicts"`
						Nd6DadProxyUnrResponses               string `xml:"nd6-dad-proxy-unr-responses"`
						Nd6NdpProxyUnrResolveCnt              string `xml:"nd6-ndp-proxy-unr-resolve-cnt"`
						Nd6DadProxyUnrResolveCnt              string `xml:"nd6-dad-proxy-unr-resolve-cnt"`
						Nd6DadProxyUnrEqportDrop              string `xml:"nd6-dad-proxy-unr-eqport-drop"`
						Nd6DadProxyUnrNomacDrop               string `xml:"nd6-dad-proxy-unr-nomac-drop"`
						Nd6RequestsDroppedOnEntry             string `xml:"nd6-requests-dropped-on-entry"`
						Nd6RequestsDroppedDuringRetry         string `xml:"nd6-requests-dropped-during-retry"`
					} `xml:"icmp6"`
					Mpls struct {
						Text                                      string `xml:",chardata"`
						TotalMplsPacketsReceived                  string `xml:"total-mpls-packets-received"`
						PacketsForwarded                          string `xml:"packets-forwarded"`
						PacketsDropped                            string `xml:"packets-dropped"`
						PacketsWithHeaderTooSmall                 string `xml:"packets-with-header-too-small"`
						AfterTaggingPacketsCanNotFitLinkMtu       string `xml:"after-tagging-packets-can-not-fit-link-mtu"`
						PacketsWithIpv4ExplicitNullTag            string `xml:"packets-with-ipv4-explicit-null-tag"`
						PacketsWithIpv4ExplicitNullChecksumErrors string `xml:"packets-with-ipv4-explicit-null-checksum-errors"`
						PacketsWithRouterAlertTag                 string `xml:"packets-with-router-alert-tag"`
						LspPingPackets                            string `xml:"lsp-ping-packets"`
						PacketsWithTtlExpired                     string `xml:"packets-with-ttl-expired"`
						PacketsWithTagEncodingError               string `xml:"packets-with-tag-encoding-error"`
						PacketsDiscardedDueToNoRoute              string `xml:"packets-discarded-due-to-no-route"`
						PacketsUsedFirstNexthopInEcmpUnilist      string `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
						PacketsDroppedDueToIflDown                string `xml:"packets-dropped-due-to-ifl-down"`
						PacketsDroppedAtMplsSocketSend            string `xml:"packets-dropped-at-mpls-socket-send"`
						PacketsForwardedAtMplsSocketSend          string `xml:"packets-forwarded-at-mpls-socket-send"`
						PacketsDroppedAtP2mpCnhOutput             string `xml:"packets-dropped-at-p2mp-cnh-output"`
					} `xml:"mpls"`
				}{
					Ip: struct {
						Text                                      string  `xml:",chardata"`
						PacketsReceived                           float64 `xml:"packets-received"`
						BadHeaderChecksums                        float64 `xml:"bad-header-checksums"`
						PacketsWithSizeSmallerThanMinimum         float64 `xml:"packets-with-size-smaller-than-minimum"`
						PacketsWithDataSizeLessThanDatalength     float64 `xml:"packets-with-data-size-less-than-datalength"`
						PacketsWithHeaderLengthLessThanDataSize   float64 `xml:"packets-with-header-length-less-than-data-size"`
						PacketsWithDataLengthLessThanHeaderlength float64 `xml:"packets-with-data-length-less-than-headerlength"`
						PacketsWithIncorrectVersionNumber         float64 `xml:"packets-with-incorrect-version-number"`
						PacketsDestinedToDeadNextHop              float64 `xml:"packets-destined-to-dead-next-hop"`
						FragmentsReceived                         float64 `xml:"fragments-received"`
						FragmentsDroppedDueToOutofspaceOrDup      float64 `xml:"fragments-dropped-due-to-outofspace-or-dup"`
						FragmentsDroppedDueToQueueoverflow        float64 `xml:"fragments-dropped-due-to-queueoverflow"`
						FragmentsDroppedAfterTimeout              float64 `xml:"fragments-dropped-after-timeout"`
						PacketsReassembledOk                      float64 `xml:"packets-reassembled-ok"`
						PacketsForThisHost                        float64 `xml:"packets-for-this-host"`
						PacketsForUnknownOrUnsupportedProtocol    float64 `xml:"packets-for-unknown-or-unsupported-protocol"`
						PacketsForwarded                          float64 `xml:"packets-forwarded"`
						PacketsNotForwardable                     float64 `xml:"packets-not-forwardable"`
						RedirectsSent                             float64 `xml:"redirects-sent"`
						PacketsSentFromThisHost                   float64 `xml:"packets-sent-from-this-host"`
						PacketsSentWithFabricatedIpHeader         float64 `xml:"packets-sent-with-fabricated-ip-header"`
						OutputPacketsDroppedDueToNoBufs           float64 `xml:"output-packets-dropped-due-to-no-bufs"`
						OutputPacketsDiscardedDueToNoRoute        float64 `xml:"output-packets-discarded-due-to-no-route"`
						OutputDatagramsFragmented                 float64 `xml:"output-datagrams-fragmented"`
						FragmentsCreated                          float64 `xml:"fragments-created"`
						DatagramsThatCanNotBeFragmented           float64 `xml:"datagrams-that-can-not-be-fragmented"`
						PacketsWithBadOptions                     float64 `xml:"packets-with-bad-options"`
						PacketsWithOptionsHandledWithoutError     float64 `xml:"packets-with-options-handled-without-error"`
						StrictSourceAndRecordRouteOptions         float64 `xml:"strict-source-and-record-route-options"`
						LooseSourceAndRecordRouteOptions          float64 `xml:"loose-source-and-record-route-options"`
						RecordRouteOptions                        float64 `xml:"record-route-options"`
						TimestampOptions                          float64 `xml:"timestamp-options"`
						TimestampAndAddressOptions                float64 `xml:"timestamp-and-address-options"`
						TimestampAndPrespecifiedAddressOptions    float64 `xml:"timestamp-and-prespecified-address-options"`
						OptionPacketsDroppedDueToRateLimit        float64 `xml:"option-packets-dropped-due-to-rate-limit"`
						RouterAlertOptions                        float64 `xml:"router-alert-options"`
						MulticastPacketsDropped                   float64 `xml:"multicast-packets-dropped"`
						PacketsDropped                            float64 `xml:"packets-dropped"`
						TransitRePacketsDroppedOnMgmtInterface    float64 `xml:"transit-re-packets-dropped-on-mgmt-interface"`
						PacketsUsedFirstNexthopInEcmpUnilist      float64 `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
						IncomingTtpoipPacketsReceived             float64 `xml:"incoming-ttpoip-packets-received"`
						IncomingTtpoipPacketsDropped              float64 `xml:"incoming-ttpoip-packets-dropped"`
						OutgoingTtpoipPacketsSent                 float64 `xml:"outgoing-ttpoip-packets-sent"`
						OutgoingTtpoipPacketsDropped              float64 `xml:"outgoing-ttpoip-packets-dropped"`
						IncomingRawipPacketsDroppedNoSocketBuffer float64 `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
						IncomingVirtualNodePacketsDelivered       float64 `xml:"incoming-virtual-node-packets-delivered"`
					}{
						PacketsReceived:                           1000,
						BadHeaderChecksums:                        5,
						PacketsWithSizeSmallerThanMinimum:         10,
						PacketsWithDataSizeLessThanDatalength:     2,
						PacketsWithHeaderLengthLessThanDataSize:   3,
						PacketsWithDataLengthLessThanHeaderlength: 1,
						PacketsWithIncorrectVersionNumber:         0,
						PacketsDestinedToDeadNextHop:              0,
						FragmentsReceived:                         50,
						FragmentsDroppedDueToOutofspaceOrDup:      2,
						FragmentsDroppedDueToQueueoverflow:        1,
						FragmentsDroppedAfterTimeout:              0,
						PacketsReassembledOk:                      48,
						PacketsForThisHost:                        500,
						PacketsForUnknownOrUnsupportedProtocol:    5,
						PacketsForwarded:                          400,
						PacketsNotForwardable:                     10,
						RedirectsSent:                             2,
						PacketsSentFromThisHost:                   800,
						PacketsSentWithFabricatedIpHeader:         0,
						OutputPacketsDroppedDueToNoBufs:           3,
						OutputPacketsDiscardedDueToNoRoute:        1,
						OutputDatagramsFragmented:                 20,
						FragmentsCreated:                          40,
						DatagramsThatCanNotBeFragmented:           2,
						PacketsWithBadOptions:                     1,
						PacketsWithOptionsHandledWithoutError:     15,
						StrictSourceAndRecordRouteOptions:         0,
						LooseSourceAndRecordRouteOptions:          2,
						RecordRouteOptions:                        5,
						TimestampOptions:                          3,
						TimestampAndAddressOptions:                1,
						TimestampAndPrespecifiedAddressOptions:    0,
						OptionPacketsDroppedDueToRateLimit:        0,
						RouterAlertOptions:                        4,
						MulticastPacketsDropped:                   8,
						PacketsDropped:                            12,
						TransitRePacketsDroppedOnMgmtInterface:    0,
						PacketsUsedFirstNexthopInEcmpUnilist:      25,
						IncomingTtpoipPacketsReceived:             100,
						IncomingTtpoipPacketsDropped:              2,
						OutgoingTtpoipPacketsSent:                 95,
						OutgoingTtpoipPacketsDropped:              1,
						IncomingRawipPacketsDroppedNoSocketBuffer: 3,
						IncomingVirtualNodePacketsDelivered:       200,
					},
				},
				Cli: struct {
					Text   string `xml:",chardata"`
					Banner string `xml:"banner"`
				}{
					Banner: "user@router>",
				},
			},
		},
		{
			name: "empty_ipv4_statistics",
			xmlInput: `<rpc-reply junos:style="normal">
				<statistics>
					<ip>
						<packets-received>0</packets-received>
						<bad-header-checksums>0</bad-header-checksums>
						<packets-with-size-smaller-than-minimum>0</packets-with-size-smaller-than-minimum>
						<packets-with-data-size-less-than-datalength>0</packets-with-data-size-less-than-datalength>
						<packets-with-header-length-less-than-data-size>0</packets-with-header-length-less-than-data-size>
						<packets-with-data-length-less-than-headerlength>0</packets-with-data-length-less-than-headerlength>
						<packets-with-incorrect-version-number>0</packets-with-incorrect-version-number>
						<packets-destined-to-dead-next-hop>0</packets-destined-to-dead-next-hop>
						<fragments-received>0</fragments-received>
						<fragments-dropped-due-to-outofspace-or-dup>0</fragments-dropped-due-to-outofspace-or-dup>
						<fragments-dropped-due-to-queueoverflow>0</fragments-dropped-due-to-queueoverflow>
						<fragments-dropped-after-timeout>0</fragments-dropped-after-timeout>
						<packets-reassembled-ok>0</packets-reassembled-ok>
						<packets-for-this-host>0</packets-for-this-host>
						<packets-for-unknown-or-unsupported-protocol>0</packets-for-unknown-or-unsupported-protocol>
						<packets-forwarded>0</packets-forwarded>
						<packets-not-forwardable>0</packets-not-forwardable>
						<redirects-sent>0</redirects-sent>
						<packets-sent-from-this-host>0</packets-sent-from-this-host>
						<packets-sent-with-fabricated-ip-header>0</packets-sent-with-fabricated-ip-header>
						<output-packets-dropped-due-to-no-bufs>0</output-packets-dropped-due-to-no-bufs>
						<output-packets-discarded-due-to-no-route>0</output-packets-discarded-due-to-no-route>
						<output-datagrams-fragmented>0</output-datagrams-fragmented>
						<fragments-created>0</fragments-created>
						<datagrams-that-can-not-be-fragmented>0</datagrams-that-can-not-be-fragmented>
						<packets-with-bad-options>0</packets-with-bad-options>
						<packets-with-options-handled-without-error>0</packets-with-options-handled-without-error>
						<strict-source-and-record-route-options>0</strict-source-and-record-route-options>
						<loose-source-and-record-route-options>0</loose-source-and-record-route-options>
						<record-route-options>0</record-route-options>
						<timestamp-options>0</timestamp-options>
						<timestamp-and-address-options>0</timestamp-and-address-options>
						<timestamp-and-prespecified-address-options>0</timestamp-and-prespecified-address-options>
						<option-packets-dropped-due-to-rate-limit>0</option-packets-dropped-due-to-rate-limit>
						<router-alert-options>0</router-alert-options>
						<multicast-packets-dropped>0</multicast-packets-dropped>
						<packets-dropped>0</packets-dropped>
						<transit-re-packets-dropped-on-mgmt-interface>0</transit-re-packets-dropped-on-mgmt-interface>
						<packets-used-first-nexthop-in-ecmp-unilist>0</packets-used-first-nexthop-in-ecmp-unilist>
						<incoming-ttpoip-packets-received>0</incoming-ttpoip-packets-received>
						<incoming-ttpoip-packets-dropped>0</incoming-ttpoip-packets-dropped>
						<outgoing-ttpoip-packets-sent>0</outgoing-ttpoip-packets-sent>
						<outgoing-ttpoip-packets-dropped>0</outgoing-ttpoip-packets-dropped>
						<incoming-rawip-packets-dropped-no-socket-buffer>0</incoming-rawip-packets-dropped-no-socket-buffer>
						<incoming-virtual-node-packets-delivered>0</incoming-virtual-node-packets-delivered>
					</ip>
				</statistics>
				<cli>
					<banner>user@router></banner>
				</cli>
			</rpc-reply>`,
			expected: SystemStatistics{
				Cli: struct {
					Text   string `xml:",chardata"`
					Banner string `xml:"banner"`
				}{
					Banner: "user@router>",
				},
			},
		},
		{
			name: "high_values_ipv4_statistics",
			xmlInput: `<rpc-reply junos:style="normal">
				<statistics>
					<ip>
						<packets-received>999999999</packets-received>
						<bad-header-checksums>12345</bad-header-checksums>
						<packets-with-size-smaller-than-minimum>54321</packets-with-size-smaller-than-minimum>
						<packets-with-data-size-less-than-datalength>1111</packets-with-data-size-less-than-datalength>
						<packets-with-header-length-less-than-data-size>2222</packets-with-header-length-less-than-data-size>
						<packets-with-data-length-less-than-headerlength>3333</packets-with-data-length-less-than-headerlength>
						<packets-with-incorrect-version-number>4444</packets-with-incorrect-version-number>
						<packets-destined-to-dead-next-hop>5555</packets-destined-to-dead-next-hop>
						<fragments-received>888888</fragments-received>
						<fragments-dropped-due-to-outofspace-or-dup>6666</fragments-dropped-due-to-outofspace-or-dup>
						<fragments-dropped-due-to-queueoverflow>7777</fragments-dropped-due-to-queueoverflow>
						<fragments-dropped-after-timeout>8888</fragments-dropped-after-timeout>
						<packets-reassembled-ok>777777</packets-reassembled-ok>
						<packets-for-this-host>555555</packets-for-this-host>
						<packets-for-unknown-or-unsupported-protocol>9999</packets-for-unknown-or-unsupported-protocol>
						<packets-forwarded>444444</packets-forwarded>
						<packets-not-forwardable>11111</packets-not-forwardable>
						<redirects-sent>12121</redirects-sent>
						<packets-sent-from-this-host>666666</packets-sent-from-this-host>
						<packets-sent-with-fabricated-ip-header>13131</packets-sent-with-fabricated-ip-header>
						<output-packets-dropped-due-to-no-bufs>14141</output-packets-dropped-due-to-no-bufs>
						<output-packets-discarded-due-to-no-route>15151</output-packets-discarded-due-to-no-route>
						<output-datagrams-fragmented>16161</output-datagrams-fragmented>
						<fragments-created>17171</fragments-created>
						<datagrams-that-can-not-be-fragmented>18181</datagrams-that-can-not-be-fragmented>
						<packets-with-bad-options>19191</packets-with-bad-options>
						<packets-with-options-handled-without-error>20202</packets-with-options-handled-without-error>
						<strict-source-and-record-route-options>21212</strict-source-and-record-route-options>
						<loose-source-and-record-route-options>22222</loose-source-and-record-route-options>
						<record-route-options>23232</record-route-options>
						<timestamp-options>24242</timestamp-options>
						<timestamp-and-address-options>25252</timestamp-and-address-options>
						<timestamp-and-prespecified-address-options>26262</timestamp-and-prespecified-address-options>
						<option-packets-dropped-due-to-rate-limit>27272</option-packets-dropped-due-to-rate-limit>
						<router-alert-options>28282</router-alert-options>
						<multicast-packets-dropped>29292</multicast-packets-dropped>
						<packets-dropped>30303</packets-dropped>
						<transit-re-packets-dropped-on-mgmt-interface>31313</transit-re-packets-dropped-on-mgmt-interface>
						<packets-used-first-nexthop-in-ecmp-unilist>32323</packets-used-first-nexthop-in-ecmp-unilist>
						<incoming-ttpoip-packets-received>33333</incoming-ttpoip-packets-received>
						<incoming-ttpoip-packets-dropped>34343</incoming-ttpoip-packets-dropped>
						<outgoing-ttpoip-packets-sent>35353</outgoing-ttpoip-packets-sent>
						<outgoing-ttpoip-packets-dropped>36363</outgoing-ttpoip-packets-dropped>
						<incoming-rawip-packets-dropped-no-socket-buffer>37373</incoming-rawip-packets-dropped-no-socket-buffer>
						<incoming-virtual-node-packets-delivered>38383</incoming-virtual-node-packets-delivered>
					</ip>
				</statistics>
				<cli>
					<banner>admin@high-traffic-router></banner>
				</cli>
			</rpc-reply>`,
			expected: SystemStatistics{
				Statistics: struct {
					Text string `xml:",chardata"`
					Tcp  struct {
						Text                                             string  `xml:",chardata"`
						PacketsSent                                      float64 `xml:"packets-sent"`
						SentDataPackets                                  float64 `xml:"sent-data-packets"`
						DataPacketsBytes                                 float64 `xml:"data-packets-bytes"`
						SentDataPacketsRetransmitted                     float64 `xml:"sent-data-packets-retransmitted"`
						RetransmittedBytes                               float64 `xml:"retransmitted-bytes"`
						SentDataUnnecessaryRetransmitted                 float64 `xml:"sent-data-unnecessary-retransmitted"`
						SentResendsByMtuDiscovery                        float64 `xml:"sent-resends-by-mtu-discovery"`
						SentAckOnlyPackets                               float64 `xml:"sent-ack-only-packets"`
						SentPacketsDelayed                               float64 `xml:"sent-packets-delayed"`
						SentUrgOnlyPackets                               float64 `xml:"sent-urg-only-packets"`
						SentWindowProbePackets                           float64 `xml:"sent-window-probe-packets"`
						SentWindowUpdatePackets                          float64 `xml:"sent-window-update-packets"`
						SentControlPackets                               float64 `xml:"sent-control-packets"`
						PacketsReceived                                  float64 `xml:"packets-received"`
						ReceivedAcks                                     float64 `xml:"received-acks"`
						AcksBytes                                        float64 `xml:"acks-bytes"`
						ReceivedDuplicateAcks                            float64 `xml:"received-duplicate-acks"`
						ReceivedAcksForUnsentData                        float64 `xml:"received-acks-for-unsent-data"`
						PacketsReceivedInSequence                        float64 `xml:"packets-received-in-sequence"`
						InSequenceBytes                                  float64 `xml:"in-sequence-bytes"`
						ReceivedCompletelyDuplicatePacket                float64 `xml:"received-completely-duplicate-packet"`
						DuplicateInBytes                                 float64 `xml:"duplicate-in-bytes"`
						ReceivedOldDuplicatePackets                      float64 `xml:"received-old-duplicate-packets"`
						ReceivedPacketsWithSomeDupliacteData             float64 `xml:"received-packets-with-some-dupliacte-data"`
						SomeDuplicateInBytes                             float64 `xml:"some-duplicate-in-bytes"`
						ReceivedOutOfOrderPackets                        float64 `xml:"received-out-of-order-packets"`
						OutOfOrderInBytes                                float64 `xml:"out-of-order-in-bytes"`
						ReceivedPacketsOfDataAfterWindow                 float64 `xml:"received-packets-of-data-after-window"`
						Bytes                                            float64 `xml:"bytes"`
						ReceivedWindowProbes                             float64 `xml:"received-window-probes"`
						ReceivedWindowUpdatePackets                      float64 `xml:"received-window-update-packets"`
						PacketsReceivedAfterClose                        float64 `xml:"packets-received-after-close"`
						ReceivedDiscardedForBadChecksum                  float64 `xml:"received-discarded-for-bad-checksum"`
						ReceivedDiscardedForBadHeaderOffset              float64 `xml:"received-discarded-for-bad-header-offset"`
						ReceivedDiscardedBecausePacketTooShort           float64 `xml:"received-discarded-because-packet-too-short"`
						ConnectionRequests                               float64 `xml:"connection-requests"`
						ConnectionAccepts                                float64 `xml:"connection-accepts"`
						BadConnectionAttempts                            float64 `xml:"bad-connection-attempts"`
						ListenQueueOverflows                             float64 `xml:"listen-queue-overflows"`
						BadRstWindow                                     float64 `xml:"bad-rst-window"`
						ConnectionsEstablished                           float64 `xml:"connections-established"`
						ConnectionsClosed                                float64 `xml:"connections-closed"`
						Drops                                            float64 `xml:"drops"`
						ConnectionsUpdatedRttOnClose                     float64 `xml:"connections-updated-rtt-on-close"`
						ConnectionsUpdatedVarianceOnClose                float64 `xml:"connections-updated-variance-on-close"`
						ConnectionsUpdatedSsthreshOnClose                float64 `xml:"connections-updated-ssthresh-on-close"`
						EmbryonicConnectionsDropped                      float64 `xml:"embryonic-connections-dropped"`
						SegmentsUpdatedRtt                               float64 `xml:"segments-updated-rtt"`
						Attempts                                         float64 `xml:"attempts"`
						RetransmitTimeouts                               float64 `xml:"retransmit-timeouts"`
						ConnectionsDroppedByRetransmitTimeout            float64 `xml:"connections-dropped-by-retransmit-timeout"`
						PersistTimeouts                                  float64 `xml:"persist-timeouts"`
						ConnectionsDroppedByPersistTimeout               float64 `xml:"connections-dropped-by-persist-timeout"`
						KeepaliveTimeouts                                float64 `xml:"keepalive-timeouts"`
						KeepaliveProbesSent                              float64 `xml:"keepalive-probes-sent"`
						KeepaliveConnectionsDropped                      float64 `xml:"keepalive-connections-dropped"`
						AckHeaderPredictions                             float64 `xml:"ack-header-predictions"`
						DataPacketHeaderPredictions                      float64 `xml:"data-packet-header-predictions"`
						SyncacheEntriesAdded                             float64 `xml:"syncache-entries-added"`
						Retransmitted                                    float64 `xml:"retransmitted"`
						Dupsyn                                           float64 `xml:"dupsyn"`
						Dropped                                          float64 `xml:"dropped"`
						Completed                                        float64 `xml:"completed"`
						BucketOverflow                                   float64 `xml:"bucket-overflow"`
						CacheOverflow                                    float64 `xml:"cache-overflow"`
						Reset                                            float64 `xml:"reset"`
						Stale                                            float64 `xml:"stale"`
						Aborted                                          float64 `xml:"aborted"`
						Badack                                           float64 `xml:"badack"`
						Unreach                                          float64 `xml:"unreach"`
						ZoneFailures                                     float64 `xml:"zone-failures"`
						CookiesSent                                      float64 `xml:"cookies-sent"`
						CookiesReceived                                  float64 `xml:"cookies-received"`
						SackRecoveryEpisodes                             float64 `xml:"sack-recovery-episodes"`
						SegmentRetransmits                               float64 `xml:"segment-retransmits"`
						ByteRetransmits                                  float64 `xml:"byte-retransmits"`
						SackOptionsReceived                              float64 `xml:"sack-options-received"`
						SackOpitionsSent                                 float64 `xml:"sack-opitions-sent"`
						SackScoreboardOverflow                           float64 `xml:"sack-scoreboard-overflow"`
						AcksSentInResponseButNotExactRsts                float64 `xml:"acks-sent-in-response-but-not-exact-rsts"`
						AcksSentInResponseToSynsOnEstablishedConnections float64 `xml:"acks-sent-in-response-to-syns-on-established-connections"`
						RcvPacketsDroppedDueToBadAddress                 float64 `xml:"rcv-packets-dropped-due-to-bad-address"`
						OutOfSequenceSegmentDrops                        float64 `xml:"out-of-sequence-segment-drops"`
						RstPackets                                       float64 `xml:"rst-packets"`
						IcmpPacketsIgnored                               float64 `xml:"icmp-packets-ignored"`
						SendPacketsDropped                               float64 `xml:"send-packets-dropped"`
						RcvPacketsDropped                                float64 `xml:"rcv-packets-dropped"`
						OutgoingSegmentsDropped                          float64 `xml:"outgoing-segments-dropped"`
						ReceivedSynfinDropped                            float64 `xml:"received-synfin-dropped"`
						ReceivedIpsecDropped                             float64 `xml:"received-ipsec-dropped"`
						ReceivedMacDropped                               float64 `xml:"received-mac-dropped"`
						ReceivedMinttlExceeded                           float64 `xml:"received-minttl-exceeded"`
						ListenstateBadflagsDropped                       float64 `xml:"listenstate-badflags-dropped"`
						FinwaitstateBadflagsDropped                      float64 `xml:"finwaitstate-badflags-dropped"`
						ReceivedDosAttack                                float64 `xml:"received-dos-attack"`
						ReceivedBadSynack                                float64 `xml:"received-bad-synack"`
						SyncacheZoneFull                                 float64 `xml:"syncache-zone-full"`
						ReceivedRstFirewallfilter                        float64 `xml:"received-rst-firewallfilter"`
						ReceivedNoackTimewait                            float64 `xml:"received-noack-timewait"`
						ReceivedNoTimewaitState                          float64 `xml:"received-no-timewait-state"`
						ReceivedRstTimewaitState                         float64 `xml:"received-rst-timewait-state"`
						ReceivedTimewaitDrops                            float64 `xml:"received-timewait-drops"`
						ReceivedBadaddrTimewaitState                     float64 `xml:"received-badaddr-timewait-state"`
						ReceivedAckoffInSynSentrcvd                      float64 `xml:"received-ackoff-in-syn-sentrcvd"`
						ReceivedBadaddrFirewall                          float64 `xml:"received-badaddr-firewall"`
						ReceivedNosynSynSent                             float64 `xml:"received-nosyn-syn-sent"`
						ReceivedBadrstSynSent                            float64 `xml:"received-badrst-syn-sent"`
						ReceivedBadrstListenState                        float64 `xml:"received-badrst-listen-state"`
						OptionMaxsegmentLength                           float64 `xml:"option-maxsegment-length"`
						OptionWindowLength                               float64 `xml:"option-window-length"`
						OptionTimestampLength                            float64 `xml:"option-timestamp-length"`
						OptionMd5Length                                  float64 `xml:"option-md5-length"`
						OptionAuthLength                                 float64 `xml:"option-auth-length"`
						OptionSackpermittedLength                        float64 `xml:"option-sackpermitted-length"`
						OptionSackLength                                 float64 `xml:"option-sack-length"`
						OptionAuthoptionLength                           float64 `xml:"option-authoption-length"`
					} `xml:"tcp"`
					Udp struct {
						Text                                              string  `xml:",chardata"`
						DatagramsReceived                                 float64 `xml:"datagrams-received"`
						DatagramsWithIncompleteHeader                     float64 `xml:"datagrams-with-incomplete-header"`
						DatagramsWithBadDatalengthField                   float64 `xml:"datagrams-with-bad-datalength-field"`
						DatagramsWithBadChecksum                          float64 `xml:"datagrams-with-bad-checksum"`
						DatagramsDroppedDueToNoSocket                     float64 `xml:"datagrams-dropped-due-to-no-socket"`
						BroadcastOrMulticastDatagramsDroppedDueToNoSocket float64 `xml:"broadcast-or-multicast-datagrams-dropped-due-to-no-socket"`
						DatagramsDroppedDueToFullSocketBuffers            float64 `xml:"datagrams-dropped-due-to-full-socket-buffers"`
						DatagramsNotForHashedPcb                          float64 `xml:"datagrams-not-for-hashed-pcb"`
						DatagramsDelivered                                float64 `xml:"datagrams-delivered"`
						DatagramsOutput                                   float64 `xml:"datagrams-output"`
					} `xml:"udp"`
					Ip struct {
						Text                                      string  `xml:",chardata"`
						PacketsReceived                           float64 `xml:"packets-received"`
						BadHeaderChecksums                        float64 `xml:"bad-header-checksums"`
						PacketsWithSizeSmallerThanMinimum         float64 `xml:"packets-with-size-smaller-than-minimum"`
						PacketsWithDataSizeLessThanDatalength     float64 `xml:"packets-with-data-size-less-than-datalength"`
						PacketsWithHeaderLengthLessThanDataSize   float64 `xml:"packets-with-header-length-less-than-data-size"`
						PacketsWithDataLengthLessThanHeaderlength float64 `xml:"packets-with-data-length-less-than-headerlength"`
						PacketsWithIncorrectVersionNumber         float64 `xml:"packets-with-incorrect-version-number"`
						PacketsDestinedToDeadNextHop              float64 `xml:"packets-destined-to-dead-next-hop"`
						FragmentsReceived                         float64 `xml:"fragments-received"`
						FragmentsDroppedDueToOutofspaceOrDup      float64 `xml:"fragments-dropped-due-to-outofspace-or-dup"`
						FragmentsDroppedDueToQueueoverflow        float64 `xml:"fragments-dropped-due-to-queueoverflow"`
						FragmentsDroppedAfterTimeout              float64 `xml:"fragments-dropped-after-timeout"`
						PacketsReassembledOk                      float64 `xml:"packets-reassembled-ok"`
						PacketsForThisHost                        float64 `xml:"packets-for-this-host"`
						PacketsForUnknownOrUnsupportedProtocol    float64 `xml:"packets-for-unknown-or-unsupported-protocol"`
						PacketsForwarded                          float64 `xml:"packets-forwarded"`
						PacketsNotForwardable                     float64 `xml:"packets-not-forwardable"`
						RedirectsSent                             float64 `xml:"redirects-sent"`
						PacketsSentFromThisHost                   float64 `xml:"packets-sent-from-this-host"`
						PacketsSentWithFabricatedIpHeader         float64 `xml:"packets-sent-with-fabricated-ip-header"`
						OutputPacketsDroppedDueToNoBufs           float64 `xml:"output-packets-dropped-due-to-no-bufs"`
						OutputPacketsDiscardedDueToNoRoute        float64 `xml:"output-packets-discarded-due-to-no-route"`
						OutputDatagramsFragmented                 float64 `xml:"output-datagrams-fragmented"`
						FragmentsCreated                          float64 `xml:"fragments-created"`
						DatagramsThatCanNotBeFragmented           float64 `xml:"datagrams-that-can-not-be-fragmented"`
						PacketsWithBadOptions                     float64 `xml:"packets-with-bad-options"`
						PacketsWithOptionsHandledWithoutError     float64 `xml:"packets-with-options-handled-without-error"`
						StrictSourceAndRecordRouteOptions         float64 `xml:"strict-source-and-record-route-options"`
						LooseSourceAndRecordRouteOptions          float64 `xml:"loose-source-and-record-route-options"`
						RecordRouteOptions                        float64 `xml:"record-route-options"`
						TimestampOptions                          float64 `xml:"timestamp-options"`
						TimestampAndAddressOptions                float64 `xml:"timestamp-and-address-options"`
						TimestampAndPrespecifiedAddressOptions    float64 `xml:"timestamp-and-prespecified-address-options"`
						OptionPacketsDroppedDueToRateLimit        float64 `xml:"option-packets-dropped-due-to-rate-limit"`
						RouterAlertOptions                        float64 `xml:"router-alert-options"`
						MulticastPacketsDropped                   float64 `xml:"multicast-packets-dropped"`
						PacketsDropped                            float64 `xml:"packets-dropped"`
						TransitRePacketsDroppedOnMgmtInterface    float64 `xml:"transit-re-packets-dropped-on-mgmt-interface"`
						PacketsUsedFirstNexthopInEcmpUnilist      float64 `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
						IncomingTtpoipPacketsReceived             float64 `xml:"incoming-ttpoip-packets-received"`
						IncomingTtpoipPacketsDropped              float64 `xml:"incoming-ttpoip-packets-dropped"`
						OutgoingTtpoipPacketsSent                 float64 `xml:"outgoing-ttpoip-packets-sent"`
						OutgoingTtpoipPacketsDropped              float64 `xml:"outgoing-ttpoip-packets-dropped"`
						IncomingRawipPacketsDroppedNoSocketBuffer float64 `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
						IncomingVirtualNodePacketsDelivered       float64 `xml:"incoming-virtual-node-packets-delivered"`
					} `xml:"ip"`
					Icmp struct {
						Text                                       string `xml:",chardata"`
						DropsDueToRateLimit                        string `xml:"drops-due-to-rate-limit"`
						CallsToIcmpError                           string `xml:"calls-to-icmp-error"`
						ErrorsNotGeneratedBecauseOldMessageWasIcmp string `xml:"errors-not-generated-because-old-message-was-icmp"`
						Histogram                                  []struct {
							Text                             string `xml:",chardata"`
							TypeOfHistogram                  string `xml:"type-of-histogram"`
							IcmpEchoReply                    string `xml:"icmp-echo-reply"`
							DestinationUnreachable           string `xml:"destination-unreachable"`
							IcmpEcho                         string `xml:"icmp-echo"`
							TimeStampReply                   string `xml:"time-stamp-reply"`
							TimeExceeded                     string `xml:"time-exceeded"`
							TimeStamp                        string `xml:"time-stamp"`
							AddressMaskRequest               string `xml:"address-mask-request"`
							AnEndpointChangedItsCookiesecret string `xml:"an-endpoint-changed-its-cookiesecret"`
						} `xml:"histogram"`
						MessagesWithBadCodeFields                                string `xml:"messages-with-bad-code-fields"`
						MessagesLessThanTheMinimumLength                         string `xml:"messages-less-than-the-minimum-length"`
						MessagesWithBadChecksum                                  string `xml:"messages-with-bad-checksum"`
						MessagesWithBadSourceAddress                             string `xml:"messages-with-bad-source-address"`
						MessagesWithBadLength                                    string `xml:"messages-with-bad-length"`
						EchoDropsWithBroadcastOrMulticastDestinatonAddress       string `xml:"echo-drops-with-broadcast-or-multicast-destinaton-address"`
						TimestampDropsWithBroadcastOrMulticastDestinationAddress string `xml:"timestamp-drops-with-broadcast-or-multicast-destination-address"`
						MessageResponsesGenerated                                string `xml:"message-responses-generated"`
					} `xml:"icmp"`
					Arp struct {
						Text                                                     string `xml:",chardata"`
						DatagramsReceived                                        string `xml:"datagrams-received"`
						ArpRequestsReceived                                      string `xml:"arp-requests-received"`
						ArpRepliesReceived                                       string `xml:"arp-replies-received"`
						ResolutionRequestReceived                                string `xml:"resolution-request-received"`
						ResolutionRequestDropped                                 string `xml:"resolution-request-dropped"`
						UnrestrictedProxyRequests                                string `xml:"unrestricted-proxy-requests"`
						RestrictedProxyRequests                                  string `xml:"restricted-proxy-requests"`
						ReceivedProxyRequests                                    string `xml:"received-proxy-requests"`
						ProxyRequestsNotProxied                                  string `xml:"proxy-requests-not-proxied"`
						RestrictedProxyRequestsNotProxied                        string `xml:"restricted-proxy-requests-not-proxied"`
						DatagramsWithBogusInterface                              string `xml:"datagrams-with-bogus-interface"`
						DatagramsWithIncorrectLength                             string `xml:"datagrams-with-incorrect-length"`
						DatagramsForNonIpProtocol                                string `xml:"datagrams-for-non-ip-protocol"`
						DatagramsWithUnsupportedOpcode                           string `xml:"datagrams-with-unsupported-opcode"`
						DatagramsWithBadProtocolAddressLength                    string `xml:"datagrams-with-bad-protocol-address-length"`
						DatagramsWithBadHardwareAddressLength                    string `xml:"datagrams-with-bad-hardware-address-length"`
						DatagramsWithMulticastSourceAddress                      string `xml:"datagrams-with-multicast-source-address"`
						DatagramsWithMulticastTargetAddress                      string `xml:"datagrams-with-multicast-target-address"`
						DatagramsWithMyOwnHardwareAddress                        string `xml:"datagrams-with-my-own-hardware-address"`
						DatagramsForAnAddressNotOnTheInterface                   string `xml:"datagrams-for-an-address-not-on-the-interface"`
						DatagramsWithABroadcastSourceAddress                     string `xml:"datagrams-with-a-broadcast-source-address"`
						DatagramsWithSourceAddressDuplicateToMine                string `xml:"datagrams-with-source-address-duplicate-to-mine"`
						DatagramsWhichWereNotForMe                               string `xml:"datagrams-which-were-not-for-me"`
						PacketsDiscardedWaitingForResolution                     string `xml:"packets-discarded-waiting-for-resolution"`
						PacketsSentAfterWaitingForResolution                     string `xml:"packets-sent-after-waiting-for-resolution"`
						ArpRequestsSent                                          string `xml:"arp-requests-sent"`
						ArpRepliesSent                                           string `xml:"arp-replies-sent"`
						RequestsForMemoryDenied                                  string `xml:"requests-for-memory-denied"`
						RequestsDroppedOnEntry                                   string `xml:"requests-dropped-on-entry"`
						RequestsDroppedDuringRetry                               string `xml:"requests-dropped-during-retry"`
						RequestsDroppedDueToInterfaceDeletion                    string `xml:"requests-dropped-due-to-interface-deletion"`
						RequestsOnUnnumberedInterfaces                           string `xml:"requests-on-unnumbered-interfaces"`
						NewRequestsOnUnnumberedInterfaces                        string `xml:"new-requests-on-unnumbered-interfaces"`
						RepliesFromUnnumberedInterfaces                          string `xml:"replies-from-unnumbered-interfaces"`
						RequestsOnUnnumberedInterfaceWithNonSubnettedDonor       string `xml:"requests-on-unnumbered-interface-with-non-subnetted-donor"`
						RepliesFromUnnumberedInterfaceWithNonSubnettedDonor      string `xml:"replies-from-unnumbered-interface-with-non-subnetted-donor"`
						ArpPacketsRejectedAsFamilyIsConfiguredWithDenyArp        string `xml:"arp-packets-rejected-as-family-is-configured-with-deny-arp"`
						ArpResponsePacketsAreRejectedOnMcAeIclInterface          string `xml:"arp-response-packets-are-rejected-on-mc-ae-icl-interface"`
						ArpRepliesAreRejectedAsSourceAndDestinationIsSame        string `xml:"arp-replies-are-rejected-as-source-and-destination-is-same"`
						ArpProbeForProxyAddressReachableFromTheIncomingInterface string `xml:"arp-probe-for-proxy-address-reachable-from-the-incoming-interface"`
						ArpRequestDiscardedForVrrpSourceAddress                  string `xml:"arp-request-discarded-for-vrrp-source-address"`
						SelfArpRequestPacketReceivedOnIrbInterface               string `xml:"self-arp-request-packet-received-on-irb-interface"`
						ProxyArpRequestDiscardedAsSourceIpIsAProxyTarget         string `xml:"proxy-arp-request-discarded-as-source-ip-is-a-proxy-target"`
						ArpPacketsAreDroppedAsNexthopAllocationFailed            string `xml:"arp-packets-are-dropped-as-nexthop-allocation-failed"`
						ArpPacketsReceivedFromPeerVrrpRouterAndDiscarded         string `xml:"arp-packets-received-from-peer-vrrp-router-and-discarded"`
						ArpPacketsAreRejectedAsTargetIpArpResolveIsInProgress    string `xml:"arp-packets-are-rejected-as-target-ip-arp-resolve-is-in-progress"`
						GratArpPacketsAreIgnoredAsMacAddressIsNotChanged         string `xml:"grat-arp-packets-are-ignored-as-mac-address-is-not-changed"`
						ArpPacketsAreDroppedFromPeerVrrp                         string `xml:"arp-packets-are-dropped-from-peer-vrrp"`
						ArpPacketsAreDroppedAsDriverCallFailed                   string `xml:"arp-packets-are-dropped-as-driver-call-failed"`
						ArpPacketsAreDroppedAsSourceIsNotValidated               string `xml:"arp-packets-are-dropped-as-source-is-not-validated"`
						ArpSystemMax                                             string `xml:"arp-system-max"`
						ArpPublicMax                                             string `xml:"arp-public-max"`
						ArpIriMax                                                string `xml:"arp-iri-max"`
						ArpMgtMax                                                string `xml:"arp-mgt-max"`
						ArpPublicCnt                                             string `xml:"arp-public-cnt"`
						ArpIriCnt                                                string `xml:"arp-iri-cnt"`
						ArpMgtCnt                                                string `xml:"arp-mgt-cnt"`
						ArpSystemDrop                                            string `xml:"arp-system-drop"`
						ArpPublicDrop                                            string `xml:"arp-public-drop"`
						ArpIriDrop                                               string `xml:"arp-iri-drop"`
						ArpMgtDrop                                               string `xml:"arp-mgt-drop"`
					} `xml:"arp"`
					Ip6 struct {
						Text                                  string `xml:",chardata"`
						TotalPacketsReceived                  float64 `xml:"total-packets-received"`
						Ip6PacketsWithSizeSmallerThanMinimum  float64 `xml:"ip6-packets-with-size-smaller-than-minimum"`
						PacketsWithDatasizeLessThanDataLength float64 `xml:"packets-with-datasize-less-than-data-length"`
						Ip6PacketsWithBadOptions              float64 `xml:"ip6-packets-with-bad-options"`
						Ip6PacketsWithIncorrectVersionNumber  float64 `xml:"ip6-packets-with-incorrect-version-number"`
						Ip6FragmentsReceived                  float64 `xml:"ip6-fragments-received"`
						DuplicateOrOutOfSpaceFragmentsDropped float64 `xml:"duplicate-or-out-of-space-fragments-dropped"`
						Ip6FragmentsDroppedAfterTimeout       float64 `xml:"ip6-fragments-dropped-after-timeout"`
						FragmentsThatExceededLimit            float64 `xml:"fragments-that-exceeded-limit"`
						Ip6PacketsReassembledOk               float64 `xml:"ip6-packets-reassembled-ok"`
						Ip6PacketsForThisHost                 float64 `xml:"ip6-packets-for-this-host"`
						Ip6PacketsForwarded                   float64 `xml:"ip6-packets-forwarded"`
						Ip6PacketsNotForwardable              float64 `xml:"ip6-packets-not-forwardable"`
						Ip6RedirectsSent                      float64 `xml:"ip6-redirects-sent"`
						Ip6PacketsSentFromThisHost            float64 `xml:"ip6-packets-sent-from-this-host"`
						Ip6PacketsSentWithFabricatedIpHeader  float64 `xml:"ip6-packets-sent-with-fabricated-ip-header"`
						Ip6OutputPacketsDroppedDueToNoBufs    float64 `xml:"ip6-output-packets-dropped-due-to-no-bufs"`
						Ip6OutputPacketsDiscardedDueToNoRoute float64 `xml:"ip6-output-packets-discarded-due-to-no-route"`
						Ip6OutputDatagramsFragmented          float64 `xml:"ip6-output-datagrams-fragmented"`
						Ip6FragmentsCreated                   float64 `xml:"ip6-fragments-created"`
						Ip6DatagramsThatCanNotBeFragmented    float64 `xml:"ip6-datagrams-that-can-not-be-fragmented"`
						PacketsThatViolatedScopeRules         float64 `xml:"packets-that-violated-scope-rules"`
						MulticastPacketsWhichWeDoNotJoin      float64 `xml:"multicast-packets-which-we-do-not-join"`
						Histogram                             float64 `xml:"histogram"`
						Ip6nhTcp                              float64 `xml:"ip6nh-tcp"`
						Ip6nhUdp                              float64 `xml:"ip6nh-udp"`
						Ip6nhIcmp6                            float64 `xml:"ip6nh-icmp6"`
						PacketsWhoseHeadersAreNotContinuous   float64 `xml:"packets-whose-headers-are-not-continuous"`
						TunnelingPacketsThatCanNotFindGif     float64 `xml:"tunneling-packets-that-can-not-find-gif"`
						PacketsDiscardedDueToTooMayHeaders    float64 `xml:"packets-discarded-due-to-too-may-headers"`
						FailuresOfSourceAddressSelection      float64 `xml:"failures-of-source-address-selection"`
						HeaderType                            []struct {
							Text                            string  `xml:",chardata"`
							HeaderForSourceAddressSelection string `xml:"header-for-source-address-selection"`
							LinkLocals                      float64 `xml:"link-locals"`
							Globals                         float64 `xml:"globals"`
							AddressScope                    float64 `xml:"address-scope"`
							HexValue                        float64 `xml:"hex-value"`
						} `xml:"header-type"`
						ForwardCacheHit                       float64 `xml:"forward-cache-hit"`
						ForwardCacheMiss                      float64 `xml:"forward-cache-miss"`
						Ip6PacketsDestinedToDeadNextHop       float64 `xml:"ip6-packets-destined-to-dead-next-hop"`
						Ip6OptionPacketsDroppedDueToRateLimit float64 `xml:"ip6-option-packets-dropped-due-to-rate-limit"`
						Ip6PacketsDropped                     float64 `xml:"ip6-packets-dropped"`
						PacketsDroppedDueToBadProtocol        float64 `xml:"packets-dropped-due-to-bad-protocol"`
						TransitRePacketDroppedOnMgmtInterface float64 `xml:"transit-re-packet-dropped-on-mgmt-interface"`
						PacketUsedFirstNexthopInEcmpUnilist   float64 `xml:"packet-used-first-nexthop-in-ecmp-unilist"`
					} `xml:"ip6"`
					Icmp6 struct {
						Text                                            string `xml:",chardata"`
						ProtocolName                                    string `xml:"protocol-name"`
						CallsToIcmp6Error                               string `xml:"calls-to-icmp6-error"`
						ErrorsNotGeneratedBecauseOldMessageWasIcmpError string `xml:"errors-not-generated-because-old-message-was-icmp-error"`
						ErrorsNotGeneratedBecauseRateLimitation         string `xml:"errors-not-generated-because-rate-limitation"`
						OutputHistogram                                 struct {
							Text                    string `xml:",chardata"`
							Style                   string `xml:"style,attr"`
							HistogramType           string `xml:"histogram-type"`
							UnreachableIcmp6Packets string `xml:"unreachable-icmp6-packets"`
							Icmp6Echo               string `xml:"icmp6-echo"`
							Icmp6EchoReply          string `xml:"icmp6-echo-reply"`
							NeighborSolicitation    string `xml:"neighbor-solicitation"`
							NeighborAdvertisement   string `xml:"neighbor-advertisement"`
						} `xml:"output-histogram"`
						Icmp6MessagesWithBadCodeFields string `xml:"icmp6-messages-with-bad-code-fields"`
						MessagesLessThanMinimumLength  string `xml:"messages-less-than-minimum-length"`
						BadChecksums                   string `xml:"bad-checksums"`
						Icmp6MessagesWithBadLength     string `xml:"icmp6-messages-with-bad-length"`
						InputHistogram                 struct {
							Text                           string `xml:",chardata"`
							Style                          string `xml:"style,attr"`
							HistogramType                  string `xml:"histogram-type"`
							UnreachableIcmp6Packets        string `xml:"unreachable-icmp6-packets"`
							PacketTooBig                   string `xml:"packet-too-big"`
							TimeExceededIcmp6Packets       string `xml:"time-exceeded-icmp6-packets"`
							Icmp6Echo                      string `xml:"icmp6-echo"`
							Icmp6EchoReply                 string `xml:"icmp6-echo-reply"`
							RouterSolicitationIcmp6Packets string `xml:"router-solicitation-icmp6-packets"`
							NeighborSolicitation           string `xml:"neighbor-solicitation"`
							NeighborAdvertisement          string `xml:"neighbor-advertisement"`
						} `xml:"input-histogram"`
						HistogramOfErrorMessagesToBeGenerated string `xml:"histogram-of-error-messages-to-be-generated"`
						NoRoute                               string `xml:"no-route"`
						AdministrativelyProhibited            string `xml:"administratively-prohibited"`
						BeyondScope                           string `xml:"beyond-scope"`
						AddressUnreachable                    string `xml:"address-unreachable"`
						PortUnreachable                       string `xml:"port-unreachable"`
						PacketTooBig                          string `xml:"packet-too-big"`
						TimeExceedTransit                     string `xml:"time-exceed-transit"`
						TimeExceedReassembly                  string `xml:"time-exceed-reassembly"`
						ErroneousHeaderField                  string `xml:"erroneous-header-field"`
						UnrecognizedNextHeader                string `xml:"unrecognized-next-header"`
						UnrecognizedOption                    string `xml:"unrecognized-option"`
						Redirect                              string `xml:"redirect"`
						Unknown                               string `xml:"unknown"`
						Icmp6MessageResponsesGenerated        string `xml:"icmp6-message-responses-generated"`
						MessagesWithTooManyNdOptions          string `xml:"messages-with-too-many-nd-options"`
						NdSystemMax                           string `xml:"nd-system-max"`
						NdPublicMax                           string `xml:"nd-public-max"`
						NdIriMax                              string `xml:"nd-iri-max"`
						NdMgtMax                              string `xml:"nd-mgt-max"`
						NdPublicCnt                           string `xml:"nd-public-cnt"`
						NdIriCnt                              string `xml:"nd-iri-cnt"`
						NdMgtCnt                              string `xml:"nd-mgt-cnt"`
						NdSystemDrop                          string `xml:"nd-system-drop"`
						NdPublicDrop                          string `xml:"nd-public-drop"`
						NdIriDrop                             string `xml:"nd-iri-drop"`
						NdMgtDrop                             string `xml:"nd-mgt-drop"`
						Nd6NdpProxyRequests                   string `xml:"nd6-ndp-proxy-requests"`
						Nd6DadProxyRequests                   string `xml:"nd6-dad-proxy-requests"`
						Nd6NdpProxyResponses                  string `xml:"nd6-ndp-proxy-responses"`
						Nd6DadProxyConflicts                  string `xml:"nd6-dad-proxy-conflicts"`
						Nd6DupProxyResponses                  string `xml:"nd6-dup-proxy-responses"`
						Nd6NdpProxyResolveCnt                 string `xml:"nd6-ndp-proxy-resolve-cnt"`
						Nd6DadProxyResolveCnt                 string `xml:"nd6-dad-proxy-resolve-cnt"`
						Nd6DadProxyEqmacDrop                  string `xml:"nd6-dad-proxy-eqmac-drop"`
						Nd6DadProxyNomacDrop                  string `xml:"nd6-dad-proxy-nomac-drop"`
						Nd6NdpProxyUnrRequests                string `xml:"nd6-ndp-proxy-unr-requests"`
						Nd6DadProxyUnrRequests                string `xml:"nd6-dad-proxy-unr-requests"`
						Nd6NdpProxyUnrResponses               string `xml:"nd6-ndp-proxy-unr-responses"`
						Nd6DadProxyUnrConflicts               string `xml:"nd6-dad-proxy-unr-conflicts"`
						Nd6DadProxyUnrResponses               string `xml:"nd6-dad-proxy-unr-responses"`
						Nd6NdpProxyUnrResolveCnt              string `xml:"nd6-ndp-proxy-unr-resolve-cnt"`
						Nd6DadProxyUnrResolveCnt              string `xml:"nd6-dad-proxy-unr-resolve-cnt"`
						Nd6DadProxyUnrEqportDrop              string `xml:"nd6-dad-proxy-unr-eqport-drop"`
						Nd6DadProxyUnrNomacDrop               string `xml:"nd6-dad-proxy-unr-nomac-drop"`
						Nd6RequestsDroppedOnEntry             string `xml:"nd6-requests-dropped-on-entry"`
						Nd6RequestsDroppedDuringRetry         string `xml:"nd6-requests-dropped-during-retry"`
					} `xml:"icmp6"`
					Mpls struct {
						Text                                      string `xml:",chardata"`
						TotalMplsPacketsReceived                  string `xml:"total-mpls-packets-received"`
						PacketsForwarded                          string `xml:"packets-forwarded"`
						PacketsDropped                            string `xml:"packets-dropped"`
						PacketsWithHeaderTooSmall                 string `xml:"packets-with-header-too-small"`
						AfterTaggingPacketsCanNotFitLinkMtu       string `xml:"after-tagging-packets-can-not-fit-link-mtu"`
						PacketsWithIpv4ExplicitNullTag            string `xml:"packets-with-ipv4-explicit-null-tag"`
						PacketsWithIpv4ExplicitNullChecksumErrors string `xml:"packets-with-ipv4-explicit-null-checksum-errors"`
						PacketsWithRouterAlertTag                 string `xml:"packets-with-router-alert-tag"`
						LspPingPackets                            string `xml:"lsp-ping-packets"`
						PacketsWithTtlExpired                     string `xml:"packets-with-ttl-expired"`
						PacketsWithTagEncodingError               string `xml:"packets-with-tag-encoding-error"`
						PacketsDiscardedDueToNoRoute              string `xml:"packets-discarded-due-to-no-route"`
						PacketsUsedFirstNexthopInEcmpUnilist      string `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
						PacketsDroppedDueToIflDown                string `xml:"packets-dropped-due-to-ifl-down"`
						PacketsDroppedAtMplsSocketSend            string `xml:"packets-dropped-at-mpls-socket-send"`
						PacketsForwardedAtMplsSocketSend          string `xml:"packets-forwarded-at-mpls-socket-send"`
						PacketsDroppedAtP2mpCnhOutput             string `xml:"packets-dropped-at-p2mp-cnh-output"`
					} `xml:"mpls"`
				}{
					Ip: struct {
						Text                                      string  `xml:",chardata"`
						PacketsReceived                           float64 `xml:"packets-received"`
						BadHeaderChecksums                        float64 `xml:"bad-header-checksums"`
						PacketsWithSizeSmallerThanMinimum         float64 `xml:"packets-with-size-smaller-than-minimum"`
						PacketsWithDataSizeLessThanDatalength     float64 `xml:"packets-with-data-size-less-than-datalength"`
						PacketsWithHeaderLengthLessThanDataSize   float64 `xml:"packets-with-header-length-less-than-data-size"`
						PacketsWithDataLengthLessThanHeaderlength float64 `xml:"packets-with-data-length-less-than-headerlength"`
						PacketsWithIncorrectVersionNumber         float64 `xml:"packets-with-incorrect-version-number"`
						PacketsDestinedToDeadNextHop              float64 `xml:"packets-destined-to-dead-next-hop"`
						FragmentsReceived                         float64 `xml:"fragments-received"`
						FragmentsDroppedDueToOutofspaceOrDup      float64 `xml:"fragments-dropped-due-to-outofspace-or-dup"`
						FragmentsDroppedDueToQueueoverflow        float64 `xml:"fragments-dropped-due-to-queueoverflow"`
						FragmentsDroppedAfterTimeout              float64 `xml:"fragments-dropped-after-timeout"`
						PacketsReassembledOk                      float64 `xml:"packets-reassembled-ok"`
						PacketsForThisHost                        float64 `xml:"packets-for-this-host"`
						PacketsForUnknownOrUnsupportedProtocol    float64 `xml:"packets-for-unknown-or-unsupported-protocol"`
						PacketsForwarded                          float64 `xml:"packets-forwarded"`
						PacketsNotForwardable                     float64 `xml:"packets-not-forwardable"`
						RedirectsSent                             float64 `xml:"redirects-sent"`
						PacketsSentFromThisHost                   float64 `xml:"packets-sent-from-this-host"`
						PacketsSentWithFabricatedIpHeader         float64 `xml:"packets-sent-with-fabricated-ip-header"`
						OutputPacketsDroppedDueToNoBufs           float64 `xml:"output-packets-dropped-due-to-no-bufs"`
						OutputPacketsDiscardedDueToNoRoute        float64 `xml:"output-packets-discarded-due-to-no-route"`
						OutputDatagramsFragmented                 float64 `xml:"output-datagrams-fragmented"`
						FragmentsCreated                          float64 `xml:"fragments-created"`
						DatagramsThatCanNotBeFragmented           float64 `xml:"datagrams-that-can-not-be-fragmented"`
						PacketsWithBadOptions                     float64 `xml:"packets-with-bad-options"`
						PacketsWithOptionsHandledWithoutError     float64 `xml:"packets-with-options-handled-without-error"`
						StrictSourceAndRecordRouteOptions         float64 `xml:"strict-source-and-record-route-options"`
						LooseSourceAndRecordRouteOptions          float64 `xml:"loose-source-and-record-route-options"`
						RecordRouteOptions                        float64 `xml:"record-route-options"`
						TimestampOptions                          float64 `xml:"timestamp-options"`
						TimestampAndAddressOptions                float64 `xml:"timestamp-and-address-options"`
						TimestampAndPrespecifiedAddressOptions    float64 `xml:"timestamp-and-prespecified-address-options"`
						OptionPacketsDroppedDueToRateLimit        float64 `xml:"option-packets-dropped-due-to-rate-limit"`
						RouterAlertOptions                        float64 `xml:"router-alert-options"`
						MulticastPacketsDropped                   float64 `xml:"multicast-packets-dropped"`
						PacketsDropped                            float64 `xml:"packets-dropped"`
						TransitRePacketsDroppedOnMgmtInterface    float64 `xml:"transit-re-packets-dropped-on-mgmt-interface"`
						PacketsUsedFirstNexthopInEcmpUnilist      float64 `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
						IncomingTtpoipPacketsReceived             float64 `xml:"incoming-ttpoip-packets-received"`
						IncomingTtpoipPacketsDropped              float64 `xml:"incoming-ttpoip-packets-dropped"`
						OutgoingTtpoipPacketsSent                 float64 `xml:"outgoing-ttpoip-packets-sent"`
						OutgoingTtpoipPacketsDropped              float64 `xml:"outgoing-ttpoip-packets-dropped"`
						IncomingRawipPacketsDroppedNoSocketBuffer float64 `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
						IncomingVirtualNodePacketsDelivered       float64 `xml:"incoming-virtual-node-packets-delivered"`
					}{
						PacketsReceived:                           999999999,
						BadHeaderChecksums:                        12345,
						PacketsWithSizeSmallerThanMinimum:         54321,
						PacketsWithDataSizeLessThanDatalength:     1111,
						PacketsWithHeaderLengthLessThanDataSize:   2222,
						PacketsWithDataLengthLessThanHeaderlength: 3333,
						PacketsWithIncorrectVersionNumber:         4444,
						PacketsDestinedToDeadNextHop:              5555,
						FragmentsReceived:                         888888,
						FragmentsDroppedDueToOutofspaceOrDup:      6666,
						FragmentsDroppedDueToQueueoverflow:        7777,
						FragmentsDroppedAfterTimeout:              8888,
						PacketsReassembledOk:                      777777,
						PacketsForThisHost:                        555555,
						PacketsForUnknownOrUnsupportedProtocol:    9999,
						PacketsForwarded:                          444444,
						PacketsNotForwardable:                     11111,
						RedirectsSent:                             12121,
						PacketsSentFromThisHost:                   666666,
						PacketsSentWithFabricatedIpHeader:         13131,
						OutputPacketsDroppedDueToNoBufs:           14141,
						OutputPacketsDiscardedDueToNoRoute:        15151,
						OutputDatagramsFragmented:                 16161,
						FragmentsCreated:                          17171,
						DatagramsThatCanNotBeFragmented:           18181,
						PacketsWithBadOptions:                     19191,
						PacketsWithOptionsHandledWithoutError:     20202,
						StrictSourceAndRecordRouteOptions:         21212,
						LooseSourceAndRecordRouteOptions:          22222,
						RecordRouteOptions:                        23232,
						TimestampOptions:                          24242,
						TimestampAndAddressOptions:                25252,
						TimestampAndPrespecifiedAddressOptions:    26262,
						OptionPacketsDroppedDueToRateLimit:        27272,
						RouterAlertOptions:                        28282,
						MulticastPacketsDropped:                   29292,
						PacketsDropped:                            30303,
						TransitRePacketsDroppedOnMgmtInterface:    31313,
						PacketsUsedFirstNexthopInEcmpUnilist:      32323,
						IncomingTtpoipPacketsReceived:             33333,
						IncomingTtpoipPacketsDropped:              34343,
						OutgoingTtpoipPacketsSent:                 35353,
						OutgoingTtpoipPacketsDropped:              36363,
						IncomingRawipPacketsDroppedNoSocketBuffer: 37373,
						IncomingVirtualNodePacketsDelivered:       38383,
					},
				},
				Cli: struct {
					Text   string `xml:",chardata"`
					Banner string `xml:"banner"`
				}{
					Banner: "admin@high-traffic-router>",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SystemStatistics
			err := xml.Unmarshal([]byte(tt.xmlInput), &result)
			assert.NoError(t, err, "unmarshal should not return error")

			// Test all IP struct fields
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsReceived, result.Statistics.Ip.PacketsReceived, "PacketsReceived should match")
			assert.Equal(t, tt.expected.Statistics.Ip.BadHeaderChecksums, result.Statistics.Ip.BadHeaderChecksums, "BadHeaderChecksums should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithSizeSmallerThanMinimum, result.Statistics.Ip.PacketsWithSizeSmallerThanMinimum, "PacketsWithSizeSmallerThanMinimum should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithDataSizeLessThanDatalength, result.Statistics.Ip.PacketsWithDataSizeLessThanDatalength, "PacketsWithDataSizeLessThanDatalength should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithHeaderLengthLessThanDataSize, result.Statistics.Ip.PacketsWithHeaderLengthLessThanDataSize, "PacketsWithHeaderLengthLessThanDataSize should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithDataLengthLessThanHeaderlength, result.Statistics.Ip.PacketsWithDataLengthLessThanHeaderlength, "PacketsWithDataLengthLessThanHeaderlength should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithIncorrectVersionNumber, result.Statistics.Ip.PacketsWithIncorrectVersionNumber, "PacketsWithIncorrectVersionNumber should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsDestinedToDeadNextHop, result.Statistics.Ip.PacketsDestinedToDeadNextHop, "PacketsDestinedToDeadNextHop should match")
			assert.Equal(t, tt.expected.Statistics.Ip.FragmentsReceived, result.Statistics.Ip.FragmentsReceived, "FragmentsReceived should match")
			assert.Equal(t, tt.expected.Statistics.Ip.FragmentsDroppedDueToOutofspaceOrDup, result.Statistics.Ip.FragmentsDroppedDueToOutofspaceOrDup, "FragmentsDroppedDueToOutofspaceOrDup should match")
			assert.Equal(t, tt.expected.Statistics.Ip.FragmentsDroppedDueToQueueoverflow, result.Statistics.Ip.FragmentsDroppedDueToQueueoverflow, "FragmentsDroppedDueToQueueoverflow should match")
			assert.Equal(t, tt.expected.Statistics.Ip.FragmentsDroppedAfterTimeout, result.Statistics.Ip.FragmentsDroppedAfterTimeout, "FragmentsDroppedAfterTimeout should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsReassembledOk, result.Statistics.Ip.PacketsReassembledOk, "PacketsReassembledOk should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsForThisHost, result.Statistics.Ip.PacketsForThisHost, "PacketsForThisHost should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsForUnknownOrUnsupportedProtocol, result.Statistics.Ip.PacketsForUnknownOrUnsupportedProtocol, "PacketsForUnknownOrUnsupportedProtocol should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsForwarded, result.Statistics.Ip.PacketsForwarded, "PacketsForwarded should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsNotForwardable, result.Statistics.Ip.PacketsNotForwardable, "PacketsNotForwardable should match")
			assert.Equal(t, tt.expected.Statistics.Ip.RedirectsSent, result.Statistics.Ip.RedirectsSent, "RedirectsSent should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsSentFromThisHost, result.Statistics.Ip.PacketsSentFromThisHost, "PacketsSentFromThisHost should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsSentWithFabricatedIpHeader, result.Statistics.Ip.PacketsSentWithFabricatedIpHeader, "PacketsSentWithFabricatedIpHeader should match")
			assert.Equal(t, tt.expected.Statistics.Ip.OutputPacketsDroppedDueToNoBufs, result.Statistics.Ip.OutputPacketsDroppedDueToNoBufs, "OutputPacketsDroppedDueToNoBufs should match")
			assert.Equal(t, tt.expected.Statistics.Ip.OutputPacketsDiscardedDueToNoRoute, result.Statistics.Ip.OutputPacketsDiscardedDueToNoRoute, "OutputPacketsDiscardedDueToNoRoute should match")
			assert.Equal(t, tt.expected.Statistics.Ip.OutputDatagramsFragmented, result.Statistics.Ip.OutputDatagramsFragmented, "OutputDatagramsFragmented should match")
			assert.Equal(t, tt.expected.Statistics.Ip.FragmentsCreated, result.Statistics.Ip.FragmentsCreated, "FragmentsCreated should match")
			assert.Equal(t, tt.expected.Statistics.Ip.DatagramsThatCanNotBeFragmented, result.Statistics.Ip.DatagramsThatCanNotBeFragmented, "DatagramsThatCanNotBeFragmented should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithBadOptions, result.Statistics.Ip.PacketsWithBadOptions, "PacketsWithBadOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsWithOptionsHandledWithoutError, result.Statistics.Ip.PacketsWithOptionsHandledWithoutError, "PacketsWithOptionsHandledWithoutError should match")
			assert.Equal(t, tt.expected.Statistics.Ip.StrictSourceAndRecordRouteOptions, result.Statistics.Ip.StrictSourceAndRecordRouteOptions, "StrictSourceAndRecordRouteOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.LooseSourceAndRecordRouteOptions, result.Statistics.Ip.LooseSourceAndRecordRouteOptions, "LooseSourceAndRecordRouteOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.RecordRouteOptions, result.Statistics.Ip.RecordRouteOptions, "RecordRouteOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.TimestampOptions, result.Statistics.Ip.TimestampOptions, "TimestampOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.TimestampAndAddressOptions, result.Statistics.Ip.TimestampAndAddressOptions, "TimestampAndAddressOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.TimestampAndPrespecifiedAddressOptions, result.Statistics.Ip.TimestampAndPrespecifiedAddressOptions, "TimestampAndPrespecifiedAddressOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.OptionPacketsDroppedDueToRateLimit, result.Statistics.Ip.OptionPacketsDroppedDueToRateLimit, "OptionPacketsDroppedDueToRateLimit should match")
			assert.Equal(t, tt.expected.Statistics.Ip.RouterAlertOptions, result.Statistics.Ip.RouterAlertOptions, "RouterAlertOptions should match")
			assert.Equal(t, tt.expected.Statistics.Ip.MulticastPacketsDropped, result.Statistics.Ip.MulticastPacketsDropped, "MulticastPacketsDropped should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsDropped, result.Statistics.Ip.PacketsDropped, "PacketsDropped should match")
			assert.Equal(t, tt.expected.Statistics.Ip.TransitRePacketsDroppedOnMgmtInterface, result.Statistics.Ip.TransitRePacketsDroppedOnMgmtInterface, "TransitRePacketsDroppedOnMgmtInterface should match")
			assert.Equal(t, tt.expected.Statistics.Ip.PacketsUsedFirstNexthopInEcmpUnilist, result.Statistics.Ip.PacketsUsedFirstNexthopInEcmpUnilist, "PacketsUsedFirstNexthopInEcmpUnilist should match")
			assert.Equal(t, tt.expected.Statistics.Ip.IncomingTtpoipPacketsReceived, result.Statistics.Ip.IncomingTtpoipPacketsReceived, "IncomingTtpoipPacketsReceived should match")
			assert.Equal(t, tt.expected.Statistics.Ip.IncomingTtpoipPacketsDropped, result.Statistics.Ip.IncomingTtpoipPacketsDropped, "IncomingTtpoipPacketsDropped should match")
			assert.Equal(t, tt.expected.Statistics.Ip.OutgoingTtpoipPacketsSent, result.Statistics.Ip.OutgoingTtpoipPacketsSent, "OutgoingTtpoipPacketsSent should match")
			assert.Equal(t, tt.expected.Statistics.Ip.OutgoingTtpoipPacketsDropped, result.Statistics.Ip.OutgoingTtpoipPacketsDropped, "OutgoingTtpoipPacketsDropped should match")
			assert.Equal(t, tt.expected.Statistics.Ip.IncomingRawipPacketsDroppedNoSocketBuffer, result.Statistics.Ip.IncomingRawipPacketsDroppedNoSocketBuffer, "IncomingRawipPacketsDroppedNoSocketBuffer should match")
			assert.Equal(t, tt.expected.Statistics.Ip.IncomingVirtualNodePacketsDelivered, result.Statistics.Ip.IncomingVirtualNodePacketsDelivered, "IncomingVirtualNodePacketsDelivered should match")
		})
	}
}