// SPDX-License-Identifier: MIT

package macsec

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseXML tests the XML parsing of the MACsec connection information
func TestParseXML(t *testing.T) {
	resultIntData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.2R2-S1.3/junos">
    <macsec-connection-information>
        <macsec-interface-common-information>
            <interface-name>et-0/0/0</interface-name>
            <connectivity-association-name>bb01.dub01-dr01.kef01</connectivity-association-name>
            <cipher-suite>GCM-AES-XPN-128</cipher-suite>
            <encryption>on</encryption>
            <offset>0</offset>
            <include-sci>no</include-sci>
            <replay-protect>off</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="1300258">2w1d 01:10:58</create-time>
        <outbound-secure-channel>
            <sci>B4:F9:5D:8C:A7:91/1</sci>
            <outgoing-packet-number>29462517698</outgoing-packet-number>
            <outbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="1300258">2w1d 01:10:58</create-time>
            </outbound-secure-association>
        </outbound-secure-channel>
        <inbound-secure-channel>
            <sci>B4:F9:5D:0D:24:71/1</sci>
            <inbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="1300258">2w1d 01:10:58</create-time>
            </inbound-secure-association>
        </inbound-secure-channel>
        <macsec-interface-common-information>
            <interface-name>et-0/0/1</interface-name>
            <connectivity-association-name>bb01.ams01-bb01.dub01</connectivity-association-name>
            <cipher-suite>GCM-AES-XPN-128</cipher-suite>
            <encryption>on</encryption>
            <offset>0</offset>
            <include-sci>no</include-sci>
            <replay-protect>off</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="784806">1w2d 02:00:06</create-time>
        <outbound-secure-channel>
            <sci>B4:F9:5D:8C:A7:99/1</sci>
            <outgoing-packet-number>185812349</outgoing-packet-number>
            <outbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="784806">1w2d 02:00:06</create-time>
            </outbound-secure-association>
        </outbound-secure-channel>
        <inbound-secure-channel>
            <sci>20:93:39:38:51:19/1</sci>
            <inbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="784806">1w2d 02:00:06</create-time>
            </inbound-secure-association>
        </inbound-secure-channel>
        <macsec-interface-common-information>
            <interface-name>et-0/0/6</interface-name>
            <connectivity-association-name>bb01.dub01-bb01.lhr01</connectivity-association-name>
            <cipher-suite>GCM-AES-XPN-128</cipher-suite>
            <encryption>on</encryption>
            <offset>0</offset>
            <include-sci>no</include-sci>
            <replay-protect>off</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="309851">3d 14:04:11</create-time>
        <outbound-secure-channel>
            <sci>B4:F9:5D:8C:A7:AC/1</sci>
            <outgoing-packet-number>4543932225</outgoing-packet-number>
            <outbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="309851">3d 14:04:11</create-time>
            </outbound-secure-association>
        </outbound-secure-channel>
        <inbound-secure-channel>
            <sci>20:93:39:36:89:3C/1</sci>
            <inbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="309851">3d 14:04:11</create-time>
            </inbound-secure-association>
        </inbound-secure-channel>
        <macsec-interface-common-information>
            <interface-name>et-0/0/7</interface-name>
            <connectivity-association-name>bb01.dub01-dr01.fal01</connectivity-association-name>
            <cipher-suite>GCM-AES-XPN-128</cipher-suite>
            <encryption>on</encryption>
            <offset>0</offset>
            <include-sci>no</include-sci>
            <replay-protect>off</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="160845">1d 20:40:45</create-time>
        <outbound-secure-channel>
            <sci>B4:F9:5D:8C:A7:B4/1</sci>
            <outgoing-packet-number>571962</outgoing-packet-number>
            <outbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="160845">1d 20:40:45</create-time>
            </outbound-secure-association>
        </outbound-secure-channel>
        <inbound-secure-channel>
            <sci>B4:F9:5D:B6:77:20/1</sci>
            <inbound-secure-association>
                <association-number>0</association-number>
                <association-number-status>inuse</association-number-status>
                <create-time junos:seconds="160845">1d 20:40:45</create-time>
            </inbound-secure-association>
        </inbound-secure-channel>
    </macsec-connection-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	resultStatsData := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.2R2-S1.3/junos">
    <macsec-statistics>
        <interface-name>et-0/0/0</interface-name>
        <secure-channel-sent>
            <encrypted-packets>29470457654</encrypted-packets>
            <encrypted-bytes>15371387814258</encrypted-bytes>
            <protected-packets>0</protected-packets>
            <protected-bytes>0</protected-bytes>
        </secure-channel-sent>
        <secure-association-sent>
            <encrypted-packets>29470457654</encrypted-packets>
            <protected-packets>0</protected-packets>
        </secure-association-sent>
        <secure-channel-received>
            <ok-packets>52215340924</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>17104313476786</decrypted-bytes>
        </secure-channel-received>
        <secure-association-received>
            <ok-packets>52215340924</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>0</decrypted-bytes>
        </secure-association-received>
        <interface-name>et-0/0/1</interface-name>
        <secure-channel-sent>
            <encrypted-packets>185933514</encrypted-packets>
            <encrypted-bytes>52358181842</encrypted-bytes>
            <protected-packets>0</protected-packets>
            <protected-bytes>0</protected-bytes>
        </secure-channel-sent>
        <secure-association-sent>
            <encrypted-packets>185933514</encrypted-packets>
            <protected-packets>0</protected-packets>
        </secure-association-sent>
        <secure-channel-received>
            <ok-packets>1404697306</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>625824199534</decrypted-bytes>
        </secure-channel-received>
        <secure-association-received>
            <ok-packets>1404697306</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>0</decrypted-bytes>
        </secure-association-received>
        <interface-name>et-0/0/6</interface-name>
        <secure-channel-sent>
            <encrypted-packets>4609309585</encrypted-packets>
            <encrypted-bytes>1578896710717</encrypted-bytes>
            <protected-packets>0</protected-packets>
            <protected-bytes>0</protected-bytes>
        </secure-channel-sent>
        <secure-association-sent>
            <encrypted-packets>4609309585</encrypted-packets>
            <protected-packets>0</protected-packets>
        </secure-association-sent>      
        <secure-channel-received>
            <ok-packets>98753547</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>81188722443</decrypted-bytes>
        </secure-channel-received>
        <secure-association-received>
            <ok-packets>98753547</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>0</decrypted-bytes>
        </secure-association-received>
        <interface-name>et-0/0/7</interface-name>
        <secure-channel-sent>
            <encrypted-packets>583486</encrypted-packets>
            <encrypted-bytes>139631615</encrypted-bytes>
            <protected-packets>0</protected-packets>
            <protected-bytes>0</protected-bytes>
        </secure-channel-sent>
        <secure-association-sent>
            <encrypted-packets>583486</encrypted-packets>
            <protected-packets>0</protected-packets>
        </secure-association-sent>
        <secure-channel-received>
            <ok-packets>483806</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>56295047</decrypted-bytes>
        </secure-channel-received>
        <secure-association-received>
            <ok-packets>483806</ok-packets>
            <validated-bytes>0</validated-bytes>
            <decrypted-bytes>0</decrypted-bytes>
        </secure-association-received>
    </macsec-statistics>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	var resultInt resultInt
	var resultStats resultStats

	// Parse the XML data for Interfaces
	err := xml.Unmarshal([]byte(resultIntData), &resultInt)
	assert.NoError(t, err)

	// Validate the parsed data
	assert.Len(t, resultInt.MacsecConnectionInformation, 4)

	// Validate first connection
	assert.Equal(t, "et-0/0/1", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].InterfaceName)
	assert.Equal(t, "bb01.ams01-bb01.dub01", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].ConnectivityAssociationName)
	assert.Equal(t, "GCM-AES-XPN-128", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].CipherSuite)
	assert.Equal(t, "inuse", resultInt.MacsecConnectionInformation.OutboundSecureChannel[0].OutboundSecureAssociation.AssociationNumberStatus)
	assert.Equal(t, "20:93:39:38:51:19/1", resultInt.MacsecConnectionInformation.OutboundSecureChannel[0].Sci)
	assert.Len(t, resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation, 4)

	// Validate second connection
	assert.Equal(t, "et-0/0/6", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].InterfaceName)
	assert.Equal(t, "bb01.dub01-bb01.lhr01", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].ConnectivityAssociationName)
	assert.Equal(t, "GCM-AES-XPN-128", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].CipherSuite)
	assert.Equal(t, "inuse", resultInt.MacsecConnectionInformation.OutboundSecureChannel[0].OutboundSecureAssociation.AssociationNumberStatus)
	assert.Equal(t, "B4:F9:5D:8C:A7:AC/1", resultInt.MacsecConnectionInformation.OutboundSecureChannel[0].Sci)
	assert.Len(t, resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation, 4)

	// Parse the XML data for statistics
	err = xml.Unmarshal([]byte(resultStatsData), &resultStats)
	assert.NoError(t, err)

	// Validate the parsed data
	assert.Len(t, resultStats.MacsecStatistics, 4)

	//validate first connection
	assert.Equal(t, "et-0/0/0", resultStats.MacsecStatistics.Interfaces[0])
	assert.Equal(t, "29470457654", resultStats.MacsecStatistics.SecureChannelSent[0].EncryptedPackets)
	assert.Equal(t, "15371387814258", resultStats.MacsecStatistics.SecureChannelSent[0].EncryptedBytes)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureChannelSent[0].ProtectedPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureChannelSent[0].ProtectedBytes)
	assert.Equal(t, "29470457654", resultStats.MacsecStatistics.SecureAssociationSent[0].EncryptedPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureAssociationSent[0].ProtectedPackets)
	assert.Equal(t, "52215340924", resultStats.MacsecStatistics.SecureChannelReceived[0].OkPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureChannelReceived[0].ValidatedBytes)
	assert.Equal(t, "17104313476786", resultStats.MacsecStatistics.SecureChannelReceived[0].DecryptedBytes)
	assert.Equal(t, "52215340924", resultStats.MacsecStatistics.SecureAssociationReceived[0].OkPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureAssociationReceived[0].ValidatedBytes)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureAssociationReceived[0].DecryptedBytes)

	//validate third connection
	assert.Equal(t, "et-0/0/6", resultStats.MacsecStatistics.Interfaces[0])
	assert.Equal(t, "4609309585", resultStats.MacsecStatistics.SecureChannelSent[0].EncryptedPackets)
	assert.Equal(t, "1578896710717", resultStats.MacsecStatistics.SecureChannelSent[0].EncryptedBytes)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureChannelSent[0].ProtectedPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureChannelSent[0].ProtectedBytes)
	assert.Equal(t, "4609309585", resultStats.MacsecStatistics.SecureAssociationSent[0].EncryptedPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureAssociationSent[0].ProtectedPackets)
	assert.Equal(t, "98753547", resultStats.MacsecStatistics.SecureChannelReceived[0].OkPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureChannelReceived[0].ValidatedBytes)
	assert.Equal(t, "81188722443", resultStats.MacsecStatistics.SecureChannelReceived[0].DecryptedBytes)
	assert.Equal(t, "98753547", resultStats.MacsecStatistics.SecureAssociationReceived[0].OkPackets)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureAssociationReceived[0].ValidatedBytes)
	assert.Equal(t, "0", resultStats.MacsecStatistics.SecureAssociationReceived[0].DecryptedBytes)
}
