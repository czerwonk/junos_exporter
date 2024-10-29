package macsec

import (
	"encoding/xml"
)

type result struct {
	XMLName                     xml.Name `xml:"rpc-reply"`
	Text                        string   `xml:",chardata"`
	Junos                       string   `xml:"junos,attr"`
	MacsecConnectionInformation struct {
		Text                             string `xml:",chardata"`
		MacsecInterfaceCommonInformation []struct {
			Text                        string `xml:",chardata"`
			InterfaceName               string `xml:"interface-name"`
			ConnectivityAssociationName string `xml:"connectivity-association-name"`
			CipherSuite                 string `xml:"cipher-suite"`
			Encryption                  string `xml:"encryption"`
			Offset                      string `xml:"offset"`
			IncludeSci                  string `xml:"include-sci"`
			ReplayProtect               string `xml:"replay-protect"`
			ReplayProtectWindow         string `xml:"replay-protect-window"`
		} `xml:"macsec-interface-common-information"`
		CreateTime []struct {
			Text    string `xml:",chardata"`
			Seconds string `xml:"seconds,attr"`
		} `xml:"create-time"`
		OutboundSecureChannel []struct {
			Text                      string `xml:",chardata"`
			Sci                       string `xml:"sci"`
			OutgoingPacketNumber      string `xml:"outgoing-packet-number"`
			OutboundSecureAssociation struct {
				Text                    string `xml:",chardata"`
				AssociationNumber       string `xml:"association-number"`
				AssociationNumberStatus string `xml:"association-number-status"`
				CreateTime              struct {
					Text    string `xml:",chardata"`
					Seconds string `xml:"seconds,attr"`
				} `xml:"create-time"`
			} `xml:"outbound-secure-association"`
		} `xml:"outbound-secure-channel"`
		InboundSecureChannel []struct {
			Text                     string `xml:",chardata"`
			Sci                      string `xml:"sci"`
			InboundSecureAssociation struct {
				Text                    string `xml:",chardata"`
				AssociationNumber       string `xml:"association-number"`
				AssociationNumberStatus string `xml:"association-number-status"`
				CreateTime              struct {
					Text    string `xml:",chardata"`
					Seconds string `xml:"seconds,attr"`
				} `xml:"create-time"`
			} `xml:"inbound-secure-association"`
		} `xml:"inbound-secure-channel"`
	} `xml:"macsec-connection-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
