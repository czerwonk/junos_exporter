package systemstatisticsIPv6

import "encoding/xml"

type StatisticsIPv6 struct {
	XMLName    xml.Name `xml:"rpc-reply"`
	Text       string   `xml:",chardata"`
	Junos      string   `xml:"junos,attr"`
	Statistics struct {
		Text string `xml:",chardata"`
		Ip6  struct {
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
			Histogram                             string `xml:"histogram"`
			Ip6nhTcp                              float64 `xml:"ip6nh-tcp"`
			Ip6nhUdp                              float64 `xml:"ip6nh-udp"`
			Ip6nhIcmp6                            float64 `xml:"ip6nh-icmp6"`
			PacketsWhoseHeadersAreNotContinuous   float64 `xml:"packets-whose-headers-are-not-continuous"`
			TunnelingPacketsThatCanNotFindGif     float64 `xml:"tunneling-packets-that-can-not-find-gif"`
			PacketsDiscardedDueToTooMayHeaders    float64 `xml:"packets-discarded-due-to-too-may-headers"`
			FailuresOfSourceAddressSelection      float64 `xml:"failures-of-source-address-selection"`
			HeaderType                            []struct {
				Text                            string `xml:",chardata"`
				HeaderForSourceAddressSelection string `xml:"header-for-source-address-selection"`
				LinkLocals                      float64 `xml:"link-locals"`
				Globals                         float64 `xml:"globals"`
				AddressScope                    string `xml:"address-scope"`
				HexValue                        string `xml:"hex-value"`
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
	} `xml:"statistics"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

