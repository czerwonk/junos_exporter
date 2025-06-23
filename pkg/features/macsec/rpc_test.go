package macsec

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
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
	type xmlStruct struct {
		XMLName          xml.Name `xml:"rpc-reply"`
		MacsecStatistics struct {
			InterfaceName     []string `xml:"interface-name"`
			SecureChannelSent []struct {
				EncryptedPackets string `xml:"encrypted-packets"`
				EncryptedBytes   string `xml:"encrypted-bytes"`
				ProtectedPackets string `xml:"protected-packets"`
				ProtectedBytes   string `xml:"protected-bytes"`
			} `xml:"secure-channel-sent"`
			SecureAssociationSent []struct {
				EncryptedPackets string `xml:"encrypted-packets"`
				ProtectedPackets string `xml:"protected-packets"`
			} `xml:"secure-association-sent"`
			SecureChannelReceived []struct {
				OkPackets      string `xml:"ok-packets"`
				ValidatedBytes string `xml:"validated-bytes"`
				DecryptedBytes string `xml:"decrypted-bytes"`
			} `xml:"secure-channel-received"`
			SecureAssociationReceived []struct {
				OkPackets      string `xml:"ok-packets"`
				ValidatedBytes string `xml:"validated-bytes"`
				DecryptedBytes string `xml:"decrypted-bytes"`
			} `xml:"secure-association-received"`
		} `xml:"macsec-statistics"`
	}

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
		var temp xmlStruct
		err = xml.Unmarshal(fc, &temp)
		if err != nil {
			t.Fatalf("Failed to parse data into temporary struct: %v", err)
		}

		//create final struct and convert the data
		result := &ShowSecMacsecStats{
			XMLName: temp.XMLName,
			MacsecStatistics: MacsecStatistics{
				Interfaces: temp.MacsecStatistics.InterfaceName,
			},
		}

		//convert and add data
		for _, raw := range temp.MacsecStatistics.SecureChannelSent {
			encPackets, _ := strconv.ParseUint(raw.EncryptedPackets, 10, 64)
			encBytes, _ := strconv.ParseUint(raw.EncryptedBytes, 10, 64)
			protPackets, _ := strconv.ParseUint(raw.ProtectedPackets, 10, 64)
			protBytes, _ := strconv.ParseUint(raw.ProtectedBytes, 10, 64)

			result.MacsecStatistics.SecureChannelSent = append(
				result.MacsecStatistics.SecureChannelSent,
				SecureChannelSentStats{
					EncryptedPackets: encPackets,
					EncryptedBytes:   encBytes,
					ProtectedPackets: protPackets,
					ProtectedBytes:   protBytes,
				},
			)
		}

		for _, raw := range temp.MacsecStatistics.SecureAssociationSent {
			encPackets, _ := strconv.ParseUint(raw.EncryptedPackets, 10, 64)
			protPackets, _ := strconv.ParseUint(raw.ProtectedPackets, 10, 64)

			result.MacsecStatistics.SecureAssociationSent = append(
				result.MacsecStatistics.SecureAssociationSent,
				SecureAssociationSentStats{
					EncryptedPackets: encPackets,
					ProtectedPackets: protPackets,
				},
			)
		}

		for _, raw := range temp.MacsecStatistics.SecureChannelReceived {
			okPackets, _ := strconv.ParseUint(raw.OkPackets, 10, 64)
			valBytes, _ := strconv.ParseUint(raw.ValidatedBytes, 10, 64)
			decBytes, _ := strconv.ParseUint(raw.DecryptedBytes, 10, 64)

			result.MacsecStatistics.SecureChannelReceived = append(
				result.MacsecStatistics.SecureChannelReceived,
				SecureChannelReceivedStats{
					OkPackets:      okPackets,
					ValidatedBytes: valBytes,
					DecryptedBytes: decBytes,
				},
			)
		}

		for _, raw := range temp.MacsecStatistics.SecureAssociationReceived {
			okPackets, _ := strconv.ParseUint(raw.OkPackets, 10, 64)
			valBytes, _ := strconv.ParseUint(raw.ValidatedBytes, 10, 64)
			decBytes, _ := strconv.ParseUint(raw.DecryptedBytes, 10, 64)

			result.MacsecStatistics.SecureAssociationReceived = append(
				result.MacsecStatistics.SecureAssociationReceived,
				SecureAssociationReceivedStats{
					OkPackets:      okPackets,
					ValidatedBytes: valBytes,
					DecryptedBytes: decBytes,
				},
			)
		}
		assert.Equal(t, test.expected, result, test.name)
	}
}
