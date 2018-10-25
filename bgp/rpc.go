package bgp

type BgpRpc struct {
	Information struct {
		Peers []BgpPeer `xml:"bgp-peer"`
	} `xml:"bgp-information"`
}

type BgpPeer struct {
	Ip             string `xml:"peer-address"`
	Asn            string `xml:"peer-as"`
	State          string `xml:"peer-state"`
	Description    string `xml:"description"`
	Flaps          int64  `xml:"flap-count"`
	InputMessages  int64  `xml:"input-messages"`
	OutputMessages int64  `xml:"output-messages"`
	Rib            struct {
		ActivePrefixes   int64 `xml:"active-prefix-count"`
		ReceivedPrefixes int64 `xml:"received-prefix-count"`
		AcceptedPrefixes int64 `xml:"accepted-prefix-count"`
		RejectedPrefixes int64 `xml:"suppressed-prefix-count"`
	} `xml:"bgp-rib"`
}
