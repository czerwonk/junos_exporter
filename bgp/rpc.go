package bgp

type BGPRPC struct {
	Information struct {
		Peers []BGPPeer `xml:"bgp-peer"`
	} `xml:"bgp-information"`
}

type BGPPeer struct {
	IP             string `xml:"peer-address"`
	ASN            string `xml:"peer-as"`
	State          string `xml:"peer-state"`
	Group          string `xml:"peer-group"`
	Description    string `xml:"description"`
	Flaps          int64  `xml:"flap-count"`
	InputMessages  int64  `xml:"input-messages"`
	OutputMessages int64  `xml:"output-messages"`
	RIBs           []RIB  `xml:"bgp-rib"`
}

type RIB struct {
	Name               string `xml:"name"`
	ActivePrefixes     int64  `xml:"active-prefix-count"`
	ReceivedPrefixes   int64  `xml:"received-prefix-count"`
	AcceptedPrefixes   int64  `xml:"accepted-prefix-count"`
	RejectedPrefixes   int64  `xml:"suppressed-prefix-count"`
	AdvertisedPrefixes int64  `xml:"advertised-prefix-count"`
}
