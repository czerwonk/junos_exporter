package macsec

//
//import (
//	"encoding/xml"
//)
//
//// Result represents the top-level structure of the XML response
//type result struct {
//	XMLName                     xml.Name                    `xml:"rpc-reply"`
//	MacsecConnectionInformation MacsecConnectionInformation `xml:"macsec-connection-information"`
//	CLI                         CLI                         `xml:"cli"`
//}
//
//// MacsecConnectionInformation represents the macsec connection information
//type MacsecConnectionInformation struct {
//	MacsecInterfaceCommonInformation []MacsecInterfaceCommonInformation `xml:"macsec-interface-common-information"`
//}
//
//// MacsecInterfaceCommonInformation represents information about a MACsec interface
//type MacsecInterfaceCommonInformation struct {
//	InterfaceName               string                `xml:"interface-name"`
//	ConnectivityAssociationName string                `xml:"connectivity-association-name"`
//	CipherSuite                 string                `xml:"cipher-suite"`
//	Encryption                  string                `xml:"encryption"`
//	Offset                      int                   `xml:"offset"`
//	IncludeSCI                  string                `xml:"include-sci"`
//	ReplayProtect               string                `xml:"replay-protect"`
//	ReplayProtectWindow         int                   `xml:"replay-protect-window"`
//	CreateTime                  CreateTime            `xml:"create-time"`
//	OutboundSecureChannel       OutboundSecureChannel `xml:"outbound-secure-channel"`
//	InboundSecureChannel        InboundSecureChannel  `xml:"inbound-secure-channel"`
//}
//
//// CreateTime represents the creation time of the connection
//type CreateTime struct {
//	JunosSeconds string `xml:"junos:seconds,attr"`
//	Time         string `xml:",chardata"`
//}
//
//// OutboundSecureChannel represents the outbound secure channel details
//type OutboundSecureChannel struct {
//	SCI                       string                    `xml:"sci"`
//	OutgoingPacketNumber      int                       `xml:"outgoing-packet-number"`
//	OutboundSecureAssociation OutboundSecureAssociation `xml:"outbound-secure-association"`
//}
//
//// InboundSecureChannel represents the inbound secure channel details
//type InboundSecureChannel struct {
//	SCI                      string                   `xml:"sci"`
//	InboundSecureAssociation InboundSecureAssociation `xml:"inbound-secure-association"`
//}
//
//// OutboundSecureAssociation represents the outbound secure association details
//type OutboundSecureAssociation struct {
//	AssociationNumber       int        `xml:"association-number"`
//	AssociationNumberStatus string     `xml:"association-number-status"`
//	CreateTime              CreateTime `xml:"create-time"`
//}
//
//// InboundSecureAssociation represents the inbound secure association details
//type InboundSecureAssociation struct {
//	AssociationNumber       int        `xml:"association-number"`
//	AssociationNumberStatus string     `xml:"association-number-status"`
//	CreateTime              CreateTime `xml:"create-time"`
//}
//
//// CLI represents the CLI section of the response
//type CLI struct {
//	Banner string `xml:"banner"`
//}
