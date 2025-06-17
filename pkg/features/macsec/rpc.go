package macsec

import (
	"encoding/xml"
)

type ShowSecMacsecConns struct {
	InnerXML                    []byte                         `xml:",innerxml"`
	MacsecConnectionInformation []*MacsecConnectionInformation `xml:"macsec-connection-information"`
}

type MacsecConnectionInformation struct {
	MacsecInterfaceCommonInformation *MacsecInterfaceCommonInformation `xml:"macsec-interface-common-information"`
	OutboundSecureChannel            *OutboundSecureChannel            `xml:"outbound-secure-channel"`
	InboundSecureChannel             *InboundSecureChannel             `xml:"inbound-secure-channel"`
}

type MacsecInterfaceCommonInformation struct {
	InterfaceName               string `xml:"interface-name"`
	ConnectivityAssociationName string `xml:"connectivity-association-name"`
	CipherSuite                 string `xml:"cipher-suite"`
	Encryption                  string `xml:"encryption"`
	Offset                      int    `xml:"offset"`
	IncludeSci                  string `xml:"include-sci"`
	ReplayProtect               string `xml:"replay-protect"`
	ReplayProtectWindow         int    `xml:"replay-protect-window"`
}

type OutboundSecureChannel struct {
	Sci                       string                     `xml:"sci"`
	OutgoingPacketNumber      int                        `xml:"outgoing-packet-number"`
	OutboundSecureAssociation *OutboundSecureAssociation `xml:"outbound-secure-association"`
}

type OutboundSecureAssociation struct {
	AssociationNumber       int         `xml:"association-number"`
	AssociationNumberStatus string      `xml:"association-number-status"`
	CreateTime              *CreateTime `xml:"create-time"`
}

type CreateTime struct {
	Seconds int `xml:"seconds,attr"`
}

type InboundSecureChannel struct {
	Sci                      string                    `xml:"sci"`
	InboundSecureAssociation *InboundSecureAssociation `xml:"inbound-secure-association"`
}

type InboundSecureAssociation struct {
	AssociationNumber       int         `xml:"association-number"`
	AssociationNumberStatus string      `xml:"association-number-status"`
	CreateTime              *CreateTime `xml:"create-time"`
}

// structure for the statistics reply
type resultStats struct {
	XMLName          xml.Name         `xml:"rpc-reply"`
	MacsecStatistics MacsecStatistics `xml:"macsec-statistics"`
}

// Struct for macsec statistics
type MacsecStatistics struct {
	Interfaces                []string                         `xml:"interface-name"`
	SecureChannelSent         []SecureChannelSentStats         `xml:"secure-channel-sent"`
	SecureAssociationSent     []SecureAssociationSentStats     `xml:"secure-association-sent"`
	SecureChannelReceived     []SecureChannelReceivedStats     `xml:"secure-channel-received"`
	SecureAssociationReceived []SecureAssociationReceivedStats `xml:"secure-association-received"`
}

// Struct for secure channel sent statistics
type SecureChannelSentStats struct {
	EncryptedPackets uint64 `xml:"encrypted-packets"`
	EncryptedBytes   uint64 `xml:"encrypted-bytes"`
	ProtectedPackets uint64 `xml:"protected-packets"`
	ProtectedBytes   uint64 `xml:"protected-bytes"`
}

// Struct for secure association sent statistics
type SecureAssociationSentStats struct {
	EncryptedPackets uint64 `xml:"encrypted-packets"`
	ProtectedPackets uint64 `xml:"protected-packets"`
}

// Struct for secure channel received statistics
type SecureChannelReceivedStats struct {
	OkPackets      uint64 `xml:"ok-packets"`
	ValidatedBytes uint64 `xml:"validated-bytes"`
	DecryptedBytes uint64 `xml:"decrypted-bytes"`
}

// Struct for secure association received statistics
type SecureAssociationReceivedStats struct {
	OkPackets      uint64 `xml:"ok-packets"`
	ValidatedBytes uint64 `xml:"validated-bytes"`
	DecryptedBytes uint64 `xml:"decrypted-bytes"`
}

type intStruct struct {
	XMLName              xml.Name `xml:"rpc-reply"`
	Text                 string   `xml:",chardata"`
	Junos                string   `xml:"junos,attr"`
	InterfaceInformation struct {
		Text              string `xml:",chardata"`
		Xmlns             string `xml:"xmlns,attr"`
		Style             string `xml:"style,attr"`
		PhysicalInterface struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			AdminStatus struct {
				Text   string `xml:",chardata"`
				Format string `xml:"format,attr"`
			} `xml:"admin-status"`
			Description                 string `xml:"description"`
			CurrentPhysicalAddress      string `xml:"current-physical-address"`
			HardwarePhysicalAddress     string `xml:"hardware-physical-address"`
			InterfaceTransmitStatistics string `xml:"interface-transmit-statistics"`
		} `xml:"physical-interface"`
	} `xml:"interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type allInterfaces struct {
	XMLName              xml.Name `xml:"rpc-reply"`
	Text                 string   `xml:",chardata"`
	Junos                string   `xml:"junos,attr"`
	InterfaceInformation struct {
		Text              string `xml:",chardata"`
		Xmlns             string `xml:"xmlns,attr"`
		Style             string `xml:"style,attr"`
		PhysicalInterface []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			AdminStatus struct {
				Text   string `xml:",chardata"`
				Format string `xml:"format,attr"`
			} `xml:"admin-status"`
			CurrentPhysicalAddress struct {
				Text   string `xml:",chardata"`
				Format string `xml:"format,attr"`
			} `xml:"current-physical-address"`
			HardwarePhysicalAddress struct {
				Text   string `xml:",chardata"`
				Format string `xml:"format,attr"`
			} `xml:"hardware-physical-address"`
			PhysicalInformation string `xml:"physical-information"`
		} `xml:"physical-interface"`
	} `xml:"interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
