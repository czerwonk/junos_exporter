package systemstatisticsIPv4

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsIPv4Unmarshaling(t *testing.T) {
	//There is an array of tests called tests which has 3 elements
	//each element includes the name, the xml input(example of device config)
	//and expected result. First element is regular values, second one is zero values and
	//third one is different rather bigger values
	tests := []struct {
		name     string
		xmlInput string
		expected StatisticsIPv4
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
			expected: StatisticsIPv4{
				Statistics: struct {
					Text string `xml:",chardata"`
					Ip   struct {
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
			expected: StatisticsIPv4{
				Statistics: struct {
					Text string `xml:",chardata"`
					Ip   struct {
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
						// All fields are zero by default
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
			expected: StatisticsIPv4{
				Statistics: struct {
					Text string `xml:",chardata"`
					Ip   struct {
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
			var result StatisticsIPv4
			err := xml.Unmarshal([]byte(tt.xmlInput), &result)
			assert.NoError(t, err, "unmarshal should not return error")

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
