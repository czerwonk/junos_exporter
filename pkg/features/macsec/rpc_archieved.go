// SPDX-License-Identifier: MIT

package macsec

/*
	type result struct {
		Information struct {
			Sessions []sessionInfos `xml:"sessions"`
		} `xml:"macsec-summary-information"`
	}
*/
type sessionInfos struct {
	InterfaceName  string           `xml:"interface-name"`
	ConnectionName string           `xml:"connection-name"`
	Cipher         string           `xml:"cipher"`
	ChannelInfo    []macsecChannels `xml:"channel-info"`
}

type macsecChannels struct {
	Name   string `xml:"macsec-channel-name"`
	SCId   string `xml:"macsec-channel-scid-name"`
	Status string `xml:"macsec-channel-status"`
}

// MultiConnectionResult represents the top-level structure for multiple MACsec connections.
/*
type MultiConnectionResult struct {
	Connections []Connection `xml:"connection"`
}

// Connection represents a single MACsec connection.
type Connection struct {
	InterfaceName string   `xml:"interface-name"`
	CAName        string   `xml:"ca-name"`
	CipherSuite   string   `xml:"cipher-suite"`
	Encryption    string   `xml:"encryption"`
	Outbound      Outbound `xml:"outbound-secure-channels"`
	Inbound       Inbound  `xml:"inbound-secure-channels"`
}

// Outbound represents the outbound secure channels for a connection.
type Outbound struct {
	SCId               string              `xml:"sc-id"`
	OutgoingPacketNum  int64               `xml:"outgoing-packet-number"`
	SecureAssociations []SecureAssociation `xml:"secure-associations"`
}

// Inbound represents the inbound secure channels for a connection.
type Inbound struct {
	SCId               string              `xml:"sc-id"`
	SecureAssociations []SecureAssociation `xml:"secure-associations"`
}

// SecureAssociation represents a secure association.
type SecureAssociation struct {
	AN     int    `xml:"an"`
	Status string `xml:"status"`
}

/*
package macsec

type result struct {
	Information struct {
		connections []neighbor `xml:"macsec-interface-common-information"`
	} `xml:"macsec-connection-information"`
}

type neighbor struct {
	InterfaceName  string `xml:"interface-name"`
	ConnectionName string `xml:"connectivity-association-name"`
	Cipher         string `xml:"cipher-suite"`
	Status         int64  `xml:"association-number-status"`
}


*/
