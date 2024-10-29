// SPDX-License-Identifier: MIT

package macsec

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseXML tests the XML parsing of the MACsec connection information
func TestParseXML(t *testing.T) {
	xmlData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.2R2-S1.3/junos">
    <macsec-connection-information>
        <macsec-interface-common-information>
            <interface-name>et-0/0/0</interface-name>
            <connectivity-association-name>bb01.arn01-bb01.ham02</connectivity-association-name>
            <cipher-suite>GCM-AES-XPN-128</cipher-suite>
            <encryption>on</encryption>
            <offset>0</offset>
            <include-sci>no</include-sci>
            <replay-protect>off</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="438477">5d 01:47:57</create-time>
        <outbound-secure-channel>
            <sci>B4:F9:5D:8C:D7:91/1</sci>
            <outgoing-packet-number>5876812</outgoing-packet-number>
            <outbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="438477">5d 01:47:57</create-time>
            </outbound-secure-association>
        </outbound-secure-channel>
        <inbound-secure-channel>
            <sci>20:93:39:EE:6B:DD/1</sci>
            <inbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="438477">5d 01:47:57</create-time>
            </inbound-secure-association>
        </inbound-secure-channel>
        <macsec-interface-common-information>
            <interface-name>et-0/0/7</interface-name>
            <connectivity-association-name>bb01.arn01-bb01.cph01</connectivity-association-name>
            <cipher-suite>GCM-AES-XPN-128</cipher-suite>
            <encryption>on</encryption>
            <offset>0</offset>
            <include-sci>no</include-sci>
            <replay-protect>off</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="622759">1w0d 04:59:19</create-time>
        <outbound-secure-channel>
            <sci>B4:F9:5D:8C:D7:B4/1</sci>
            <outgoing-packet-number>49402478</outgoing-packet-number>
            <outbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="622759">1w0d 04:59:19</create-time>
            </outbound-secure-association>
        </outbound-secure-channel>
        <inbound-secure-channel>        
            <sci>9C:5A:80:18:83:54/1</sci>
            <inbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="622759">1w0d 04:59:19</create-time>
            </inbound-secure-association>
        </inbound-secure-channel>
    </macsec-connection-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	var result MultiConnectionResult

	// Parse the XML data
	err := xml.Unmarshal([]byte(xmlData), &result)
	assert.NoError(t, err)

	// Validate the parsed data
	assert.Len(t, result.Connections, 2)

	// Validate first connection
	assert.Equal(t, "et-0/0/0", result.Connections[0].InterfaceName)
	assert.Equal(t, "bb01.arn01-bb01.ham02", result.Connections[0].CAName)
	assert.Equal(t, "GCM-AES-XPN-128", result.Connections[0].CipherSuite)
	assert.Equal(t, "on", result.Connections[0].Encryption)
	assert.Len(t, result.Connections[0].Outbound.SecureAssociations, 1)
	assert.Equal(t, "inuse", result.Connections[0].Outbound.SecureAssociations[0].Status)

	// Validate second connection
	assert.Equal(t, "et-0/0/7", result.Connections[1].InterfaceName)
	assert.Equal(t, "bb01.arn01-bb01.cph01", result.Connections[1].CAName)
	assert.Equal(t, "GCM-AES-XPN-128", result.Connections[1].CipherSuite)
	assert.Equal(t, "on", result.Connections[1].Encryption)
	assert.Len(t, result.Connections[1].Inbound.SecureAssociations, 1)
	assert.Equal(t, "inuse", result.Connections[1].Inbound.SecureAssociations[0].Status)
}
