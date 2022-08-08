package bgp

type BGPRPC struct {
	Information struct {
		Peers []BGPPeer `xml:"bgp-peer"`
	} `xml:"bgp-information"`
}

type BGPPeer struct {
	CFG_RTI        string               `xml:"peer-cfg-rti"`
	IP             string               `xml:"peer-address"`
	ASN            string               `xml:"peer-as"`
	State          string               `xml:"peer-state"`
	Group          string               `xml:"peer-group"`
	Description    string               `xml:"description"`
	Flaps          int64                `xml:"flap-count"`
	InputMessages  int64                `xml:"input-messages"`
	OutputMessages int64                `xml:"output-messages"`
	RIBs           []RIB                `xml:"bgp-rib"`
	BGPOI          BGPOptionInformation `xml:"bgp-option-information"`
}

type RIB struct {
	Name               string `xml:"name"`
	ActivePrefixes     int64  `xml:"active-prefix-count"`
	ReceivedPrefixes   int64  `xml:"received-prefix-count"`
	AcceptedPrefixes   int64  `xml:"accepted-prefix-count"`
	RejectedPrefixes   int64  `xml:"suppressed-prefix-count"`
	AdvertisedPrefixes int64  `xml:"advertised-prefix-count"`
}

type BGPOptionInformation struct {
	ExportPolicy    string      `xml:"export-policy"`
	ImportPolicy    string      `xml:"import-policy"`
	AddressFamilies string      `xml:"address-families"`
	LocalAddress    string      `xml:"local-address"`
	Holdtime        int64       `xml:"holdtime"`
	MetricOut       int64       `xml:"metric-out"`
	Preference      int64       `xml:"preference"`
	PrefixLimit     PrefixLimit `xml:"prefix-limit"`
	LocalAs         int64       `xml:"local-as"`
	LocalSystemAs   int64       `xml:"local-system-as"`
}

type PrefixLimit struct {
	NlriType          string `xml:"nlri-type"`
	PrefixCount       int64  `xml:"prefix-count"`
	LimitAction       string `xml:"limit-action"`
	WarningPercentage int64  `xml:"warning-percentage"`
}
