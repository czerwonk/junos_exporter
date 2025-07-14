package systemstatisticsIPv4

import "encoding/xml"

type StatisticsIPv4 struct {
	XMLName    xml.Name `xml:"rpc-reply"`
	Text       string   `xml:",chardata"`
	Junos      string   `xml:"junos,attr"`
	Statistics struct {
		Text string `xml:",chardata"`
		Ip   struct {
			Text                                      string `xml:",chardata"`
			PacketsReceived                           int    `xml:"packets-received"`
			BadHeaderChecksums                        int    `xml:"bad-header-checksums"`
			PacketsWithSizeSmallerThanMinimum         int    `xml:"packets-with-size-smaller-than-minimum"`
			PacketsWithDataSizeLessThanDatalength     int    `xml:"packets-with-data-size-less-than-datalength"`
			PacketsWithHeaderLengthLessThanDataSize   int    `xml:"packets-with-header-length-less-than-data-size"`
			PacketsWithDataLengthLessThanHeaderlength int    `xml:"packets-with-data-length-less-than-headerlength"`
			PacketsWithIncorrectVersionNumber         int    `xml:"packets-with-incorrect-version-number"`
			PacketsDestinedToDeadNextHop              int    `xml:"packets-destined-to-dead-next-hop"`
			FragmentsReceived                         int    `xml:"fragments-received"`
			FragmentsDroppedDueToOutofspaceOrDup      int    `xml:"fragments-dropped-due-to-outofspace-or-dup"`
			FragmentsDroppedDueToQueueoverflow        int    `xml:"fragments-dropped-due-to-queueoverflow"`
			FragmentsDroppedAfterTimeout              int    `xml:"fragments-dropped-after-timeout"`
			PacketsReassembledOk                      int    `xml:"packets-reassembled-ok"`
			PacketsForThisHost                        int    `xml:"packets-for-this-host"`
			PacketsForUnknownOrUnsupportedProtocol    int    `xml:"packets-for-unknown-or-unsupported-protocol"`
			PacketsForwarded                          int    `xml:"packets-forwarded"`
			PacketsNotForwardable                     int    `xml:"packets-not-forwardable"`
			RedirectsSent                             int    `xml:"redirects-sent"`
			PacketsSentFromThisHost                   int    `xml:"packets-sent-from-this-host"`
			PacketsSentWithFabricatedIpHeader         int    `xml:"packets-sent-with-fabricated-ip-header"`
			OutputPacketsDroppedDueToNoBufs           int    `xml:"output-packets-dropped-due-to-no-bufs"`
			OutputPacketsDiscardedDueToNoRoute        int    `xml:"output-packets-discarded-due-to-no-route"`
			OutputDatagramsFragmented                 int    `xml:"output-datagrams-fragmented"`
			FragmentsCreated                          int    `xml:"fragments-created"`
			DatagramsThatCanNotBeFragmented           int    `xml:"datagrams-that-can-not-be-fragmented"`
			PacketsWithBadOptions                     int    `xml:"packets-with-bad-options"`
			PacketsWithOptionsHandledWithoutError     int    `xml:"packets-with-options-handled-without-error"`
			StrictSourceAndRecordRouteOptions         int    `xml:"strict-source-and-record-route-options"`
			LooseSourceAndRecordRouteOptions          int    `xml:"loose-source-and-record-route-options"`
			RecordRouteOptions                        int    `xml:"record-route-options"`
			TimestampOptions                          int    `xml:"timestamp-options"`
			TimestampAndAddressOptions                int    `xml:"timestamp-and-address-options"`
			TimestampAndPrespecifiedAddressOptions    int    `xml:"timestamp-and-prespecified-address-options"`
			OptionPacketsDroppedDueToRateLimit        int    `xml:"option-packets-dropped-due-to-rate-limit"`
			RouterAlertOptions                        int    `xml:"router-alert-options"`
			MulticastPacketsDropped                   int    `xml:"multicast-packets-dropped"`
			PacketsDropped                            int    `xml:"packets-dropped"`
			TransitRePacketsDroppedOnMgmtInterface    int    `xml:"transit-re-packets-dropped-on-mgmt-interface"`
			PacketsUsedFirstNexthopInEcmpUnilist      int    `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
			IncomingTtpoipPacketsReceived             int    `xml:"incoming-ttpoip-packets-received"`
			IncomingTtpoipPacketsDropped              int    `xml:"incoming-ttpoip-packets-dropped"`
			OutgoingTtpoipPacketsSent                 int    `xml:"outgoing-ttpoip-packets-sent"`
			OutgoingTtpoipPacketsDropped              int    `xml:"outgoing-ttpoip-packets-dropped"`
			IncomingRawipPacketsDroppedNoSocketBuffer int    `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
			IncomingVirtualNodePacketsDelivered       int    `xml:"incoming-virtual-node-packets-delivered"`
		} `xml:"ip"`
	} `xml:"statistics"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
