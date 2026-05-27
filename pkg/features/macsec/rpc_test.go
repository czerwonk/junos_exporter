package macsec

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
							ConnectivityAssociationName: "cc12.evc12-es12.lfg12",
							CipherSuite:                 "XXX-YYY-ZZZ-65000",
							Encryption:                  "on",
							Offset:                      0,
							IncludeSci:                  "no",
							ReplayProtect:               "off",
							ReplayProtectWindow:         0,
						},
						OutboundSecureChannel: &OutboundSecureChannel{
							Sci:                  "AA:AA:AA:AA:AA:AA/1",
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
							Sci: "AA:AA:AA:AA:AA:AB/1",
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
							ConnectivityAssociationName: "cc12.bnt12-cc12.evc12",
							CipherSuite:                 "XXX-YYY-ZZZ-65000",
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
							ConnectivityAssociationName: "cc12.evc12-cc12.mis12",
							CipherSuite:                 "XXX-YYY-ZZZ-65000",
							Encryption:                  "on",
							Offset:                      0,
							IncludeSci:                  "no",
							ReplayProtect:               "off",
							ReplayProtectWindow:         0,
						},
						OutboundSecureChannel: &OutboundSecureChannel{
							Sci:                  "AA:AA:AA:AA:AA:AC/1",
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
							Sci: "AA:AA:AA:AA:AA:AD/1",
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

func TestParseShowSecurityMacsecStatistics(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		expected  *ShowSecMacsecStats
	}{
		{
			name:      "Test #1",
			inputFile: "show_security_macsec_statistics.xml",
			expected: &ShowSecMacsecStats{
				XMLName: xml.Name{Space: "", Local: "rpc-reply"},
				MacsecStatistics: MacsecStatistics{
					Interfaces: []string{"et-0/0/0", "et-0/0/1", "et-0/0/6"},
					SecureChannelSent: []SecureChannelSentStats{
						{
							EncryptedPackets: 1,
							EncryptedBytes:   2,
							ProtectedPackets: 3,
							ProtectedBytes:   4,
						},
						{
							EncryptedPackets: 2000,
							EncryptedBytes:   3000,
							ProtectedPackets: 0,
							ProtectedBytes:   0,
						},
						{
							EncryptedPackets: 8000,
							EncryptedBytes:   9000,
							ProtectedPackets: 0,
							ProtectedBytes:   0,
						},
					},
					SecureAssociationSent: []SecureAssociationSentStats{
						{
							EncryptedPackets: 5,
							ProtectedPackets: 6,
						},
						{
							EncryptedPackets: 4000,
							ProtectedPackets: 0,
						},
						{
							EncryptedPackets: 10000,
							ProtectedPackets: 0,
						},
					},
					SecureChannelReceived: []SecureChannelReceivedStats{
						{
							OkPackets:      9,
							ValidatedBytes: 1000,
							DecryptedBytes: 2000,
						},
						{
							OkPackets:      5000,
							ValidatedBytes: 0,
							DecryptedBytes: 6000,
						},
						{
							OkPackets:      11000,
							ValidatedBytes: 0,
							DecryptedBytes: 12000,
						},
					},
					SecureAssociationReceived: []SecureAssociationReceivedStats{
						{
							OkPackets:      3000,
							ValidatedBytes: 0,
							DecryptedBytes: 0,
						},
						{
							OkPackets:      7000,
							ValidatedBytes: 0,
							DecryptedBytes: 0,
						},
						{
							OkPackets:      13000,
							ValidatedBytes: 0,
							DecryptedBytes: 0,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		fmt.Println(test.name)
		fc, err := os.ReadFile("testdata/" + test.inputFile)
		if err != nil {
			t.Fatalf("Failed to read test file: %v", err)
		}

		//unmarshal to the intermediate struct
		var temp *ShowSecMacsecStats
		err = xml.Unmarshal(fc, &temp)
		if err != nil {
			t.Fatalf("Failed to parse data into temporary struct: %v", err)
		}

		assert.Equal(t, test.expected, temp, test.name)
	}
}
