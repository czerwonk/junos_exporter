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
			PacketsReceived                           float64    `xml:"packets-received"`
			BadHeaderChecksums                        float64    `xml:"bad-header-checksums"`
			PacketsWithSizeSmallerThanMinimum         float64    `xml:"packets-with-size-smaller-than-minimum"`
			PacketsWithDataSizeLessThanDatalength     float64    `xml:"packets-with-data-size-less-than-datalength"`
			PacketsWithHeaderLengthLessThanDataSize   float64    `xml:"packets-with-header-length-less-than-data-size"`
			PacketsWithDataLengthLessThanHeaderlength float64    `xml:"packets-with-data-length-less-than-headerlength"`
			PacketsWithIncorrectVersionNumber         float64    `xml:"packets-with-incorrect-version-number"`
			PacketsDestinedToDeadNextHop              float64    `xml:"packets-destined-to-dead-next-hop"`
			FragmentsReceived                         float64    `xml:"fragments-received"`
			FragmentsDroppedDueToOutofspaceOrDup      float64    `xml:"fragments-dropped-due-to-outofspace-or-dup"`
			FragmentsDroppedDueToQueueoverflow        float64    `xml:"fragments-dropped-due-to-queueoverflow"`
			FragmentsDroppedAfterTimeout              float64    `xml:"fragments-dropped-after-timeout"`
			PacketsReassembledOk                      float64    `xml:"packets-reassembled-ok"`
			PacketsForThisHost                        float64    `xml:"packets-for-this-host"`
			PacketsForUnknownOrUnsupportedProtocol    float64    `xml:"packets-for-unknown-or-unsupported-protocol"`
			PacketsForwarded                          float64    `xml:"packets-forwarded"`
			PacketsNotForwardable                     float64    `xml:"packets-not-forwardable"`
			RedirectsSent                             float64    `xml:"redirects-sent"`
			PacketsSentFromThisHost                   float64    `xml:"packets-sent-from-this-host"`
			PacketsSentWithFabricatedIpHeader         float64    `xml:"packets-sent-with-fabricated-ip-header"`
			OutputPacketsDroppedDueToNoBufs           float64    `xml:"output-packets-dropped-due-to-no-bufs"`
			OutputPacketsDiscardedDueToNoRoute        float64    `xml:"output-packets-discarded-due-to-no-route"`
			OutputDatagramsFragmented                 float64    `xml:"output-datagrams-fragmented"`
			FragmentsCreated                          float64    `xml:"fragments-created"`
			DatagramsThatCanNotBeFragmented           float64    `xml:"datagrams-that-can-not-be-fragmented"`
			PacketsWithBadOptions                     float64    `xml:"packets-with-bad-options"`
			PacketsWithOptionsHandledWithoutError     float64    `xml:"packets-with-options-handled-without-error"`
			StrictSourceAndRecordRouteOptions         float64    `xml:"strict-source-and-record-route-options"`
			LooseSourceAndRecordRouteOptions          float64    `xml:"loose-source-and-record-route-options"`
			RecordRouteOptions                        float64    `xml:"record-route-options"`
			TimestampOptions                          float64    `xml:"timestamp-options"`
			TimestampAndAddressOptions                float64    `xml:"timestamp-and-address-options"`
			TimestampAndPrespecifiedAddressOptions    float64    `xml:"timestamp-and-prespecified-address-options"`
			OptionPacketsDroppedDueToRateLimit        float64    `xml:"option-packets-dropped-due-to-rate-limit"`
			RouterAlertOptions                        float64    `xml:"router-alert-options"`
			MulticastPacketsDropped                   float64    `xml:"multicast-packets-dropped"`
			PacketsDropped                            float64    `xml:"packets-dropped"`
			TransitRePacketsDroppedOnMgmtInterface    float64    `xml:"transit-re-packets-dropped-on-mgmt-interface"`
			PacketsUsedFirstNexthopInEcmpUnilist      float64    `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
			IncomingTtpoipPacketsReceived             float64    `xml:"incoming-ttpoip-packets-received"`
			IncomingTtpoipPacketsDropped              float64    `xml:"incoming-ttpoip-packets-dropped"`
			OutgoingTtpoipPacketsSent                 float64    `xml:"outgoing-ttpoip-packets-sent"`
			OutgoingTtpoipPacketsDropped              float64    `xml:"outgoing-ttpoip-packets-dropped"`
			IncomingRawipPacketsDroppedNoSocketBuffer float64    `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
			IncomingVirtualNodePacketsDelivered       float64    `xml:"incoming-virtual-node-packets-delivered"`
		} `xml:"ip"`
	} `xml:"statistics"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
