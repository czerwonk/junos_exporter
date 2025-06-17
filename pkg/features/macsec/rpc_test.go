// SPDX-License-Identifier: MIT

package macsec

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseXML tests the XML parsing of the MACsec connection information
/*func TestParseXML(t *testing.T) {
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
            <encryption>off</encryption>
            <offset>0</offset>
            <include-sci>yes</include-sci>
            <replay-protect>on</replay-protect>
            <replay-protect-window>0</replay-protect-window>
        </macsec-interface-common-information>
        <create-time junos:seconds="784806">1w2d 02:00:06</create-time>
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
	//expected := resultInt{}
	var resultInt resultInt
	var resultStats resultStats
	//list_of_interfaces := [4]string{"et-0/0/0", "set-0/0/1", "et-0/0/6", "et-0/0/7"}
	// Parse the XML data for Interfaces
	err := xml.Unmarshal([]byte(resultIntData), &resultInt)
	assert.NoError(t, err)

	// Validate the parsed data
	assert.Len(t, resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation, 4)

	// Validate first connection
	assert.Equal(t, "et-0/0/0", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].InterfaceName)
	assert.Equal(t, "on", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].Encryption)
	assert.Equal(t, "0", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].Offset)
	assert.Equal(t, "off", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].ReplayProtect)
	assert.Equal(t, "no", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[0].IncludeSci)
	assert.Equal(t, "29462517698", resultInt.MacsecConnectionInformation.OutboundSecureChannel[0].OutgoingPacketNumber)

	// Validate second connection
	assert.Equal(t, "et-0/0/1", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[1].InterfaceName)
	assert.Equal(t, "off", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[1].Encryption)
	assert.Equal(t, "0", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[1].Offset)
	assert.Equal(t, "on", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[1].ReplayProtect)
	assert.Equal(t, "yes", resultInt.MacsecConnectionInformation.MacsecInterfaceCommonInformation[1].IncludeSci)
	//assert.Equal(t, "185812349", resultInt.MacsecConnectionInformation.OutboundSecureChannel[1].OutgoingPacketNumber)

	//testing outbound-secure-channel edge case when it is missing on one of the interfaces
	//assert.Equal(t, "4543932225", resultInt.MacsecConnectionInformation.OutboundSecureChannel[2].OutgoingPacketNumber)

	//for _, connection := range resultInt.MacsecConnectionInformation.OutboundSecureChannel {
	//	fmt.Println(connection.OutgoingPacketNumber)
	//}
	// Parse the XML data for statistics
	err = xml.Unmarshal([]byte(resultStatsData), &resultStats)
	assert.NoError(t, err)

	// Validate the parsed data
	assert.Len(t, resultStats.MacsecStatistics.Interfaces, 4)
	assert.Len(t, resultStats.MacsecStatistics.SecureChannelSent, 4)
	assert.Len(t, resultStats.MacsecStatistics.SecureChannelReceived, 4)
	assert.Len(t, resultStats.MacsecStatistics.SecureAssociationSent, 4)
	assert.Len(t, resultStats.MacsecStatistics.SecureAssociationReceived, 4)
	//validates the first interface
	assert.Equal(t, uint64(29470457654), resultStats.MacsecStatistics.SecureChannelSent[0].EncryptedPackets)
	assert.Equal(t, uint64(15371387814258), resultStats.MacsecStatistics.SecureChannelSent[0].EncryptedBytes)
	assert.Equal(t, uint64(0), resultStats.MacsecStatistics.SecureChannelSent[0].ProtectedPackets)
	assert.Equal(t, uint64(0), resultStats.MacsecStatistics.SecureChannelSent[0].ProtectedBytes)
	assert.Equal(t, uint64(29470457654), resultStats.MacsecStatistics.SecureAssociationSent[0].EncryptedPackets)
	assert.Equal(t, uint64(0), resultStats.MacsecStatistics.SecureAssociationSent[0].ProtectedPackets)
	assert.Equal(t, uint64(52215340924), resultStats.MacsecStatistics.SecureChannelReceived[0].OkPackets)
	assert.Equal(t, uint64(0), resultStats.MacsecStatistics.SecureChannelReceived[0].ValidatedBytes)
	assert.Equal(t, uint64(17104313476786), resultStats.MacsecStatistics.SecureChannelReceived[0].DecryptedBytes)
	assert.Equal(t, uint64(52215340924), resultStats.MacsecStatistics.SecureAssociationReceived[0].OkPackets)
	assert.Equal(t, uint64(0), resultStats.MacsecStatistics.SecureAssociationReceived[0].ValidatedBytes)
	assert.Equal(t, uint64(0), resultStats.MacsecStatistics.SecureAssociationReceived[0].DecryptedBytes)

}*/

func TestParseShowSecurityMacsecConnections(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		expected  *ShowSecMacsecConns
	}{
		{
			name:      "Test #1",
			inputFile: "show_security_macsec_connections.xml",
			expected: &ShowSecMacsecConns{
				MacsecConnectionInformation: []*MacsecConnectionInformation{
					{
						MacsecInterfaceCommonInformation: &MacsecInterfaceCommonInformation{
							InterfaceName:               "et-0/0/0",
							ConnectivityAssociationName: "bb01.dub01-dr01.kef01",
							CipherSuite:                 "GCM-AES-XPN-128",
							Encryption:                  "on",
							Offset:                      0,
							IncludeSci:                  "no",
							ReplayProtect:               "off",
							ReplayProtectWindow:         0,
						},
						OutboundSecureChannel: &OutboundSecureChannel{
							Sci:                  "B4:F9:5D:8C:A7:91/1",
							OutgoingPacketNumber: 29462517698,
							OutboundSecureAssociation: &OutboundSecureAssociation{
								AssociationNumber:       0,
								AssociationNumberStatus: "inuse",
								CreateTime: &CreateTime{
									Seconds: 1300258,
								},
							},
						},
						InboundSecureChannel: &InboundSecureChannel{
							Sci: "B4:F9:5D:0D:24:71/1",
							InboundSecureAssociation: &InboundSecureAssociation{
								AssociationNumber:       0,
								AssociationNumberStatus: "inuse",
								CreateTime: &CreateTime{
									Seconds: 1300258,
								},
							},
						},
					},
					{
						MacsecInterfaceCommonInformation: &MacsecInterfaceCommonInformation{
							InterfaceName:               "et-0/0/1",
							ConnectivityAssociationName: "bb01.ams01-bb01.dub01",
							CipherSuite:                 "GCM-AES-XPN-128",
							Encryption:                  "off",
							Offset:                      0,
							IncludeSci:                  "yes",
							ReplayProtect:               "on",
							ReplayProtectWindow:         0,
						},
					},
					{
						MacsecInterfaceCommonInformation: &MacsecInterfaceCommonInformation{
							InterfaceName:               "et-0/0/6",
							ConnectivityAssociationName: "bb01.dub01-bb01.lhr01",
							CipherSuite:                 "GCM-AES-XPN-128",
							Encryption:                  "on",
							Offset:                      0,
							IncludeSci:                  "no",
							ReplayProtect:               "off",
							ReplayProtectWindow:         0,
						},
						OutboundSecureChannel: &OutboundSecureChannel{
							Sci:                  "B4:F9:5D:8C:A7:AC/1",
							OutgoingPacketNumber: 4543932225,
							OutboundSecureAssociation: &OutboundSecureAssociation{
								AssociationNumber:       0,
								AssociationNumberStatus: "inuse",
								CreateTime: &CreateTime{
									Seconds: 309851,
								},
							},
						},
						InboundSecureChannel: &InboundSecureChannel{
							Sci: "20:93:39:36:89:3C/1",
							InboundSecureAssociation: &InboundSecureAssociation{
								AssociationNumber:       0,
								AssociationNumberStatus: "inuse",
								CreateTime: &CreateTime{
									Seconds: 309851,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		fc, err := os.ReadFile("testdata/" + test.inputFile)
		if err != nil {
			panic(err)
		}

		res, err := ParseShowSecurityMacsecConnections(fc)
		if err != nil {
			panic(err)
		}

		res.InnerXML = nil
		assert.Equal(t, test.expected, res, test.name)
	}
}
