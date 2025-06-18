package macsec

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMacsec_DataDriven(t *testing.T) {
	connectionsData, err := os.ReadFile("testdata/show_security_macsec_connections.xml")
	if err != nil {
		t.Fatalf("Failed to read connections test file: %v", err)
	}

	statisticsData, err := os.ReadFile("testdata/show_security_macsec_statistics.xml")
	if err != nil {
		t.Fatalf("Failed to read statistics test file: %v", err)
	}

	tests := []struct {
		name       string
		xmlData    []byte
		resultType string
		validate   func(t *testing.T, data interface{})
	}{
		{
			name:       "Parse MACsec Connections",
			xmlData:    connectionsData,
			resultType: "connections",
			validate: func(t *testing.T, data interface{}) {
				result := data.(*ShowSecMacsecConns)

				//define expected values for each interface
				expectedConnections := map[string]struct {
					caName             string
					encryption         string
					includeSci         string
					replayProtect      string
					offset             int
					hasOutboundChannel bool
					sci                string
					packetNumber       int
				}{
					"et-0/0/0": {
						caName:             "bb01.dub01-dr01.kef01",
						encryption:         "on",
						includeSci:         "no",
						replayProtect:      "off",
						offset:             0,
						hasOutboundChannel: true,
						sci:                "B4:F9:5D:8C:A7:91/1",
						packetNumber:       29462517698,
					},
					"et-0/0/1": {
						caName:             "bb01.ams01-bb01.dub01",
						encryption:         "off",
						includeSci:         "yes",
						replayProtect:      "on",
						offset:             0,
						hasOutboundChannel: false,
					},
					"et-0/0/6": {
						caName:             "bb01.dub01-bb01.lhr01",
						encryption:         "on",
						includeSci:         "no",
						replayProtect:      "off",
						offset:             0,
						hasOutboundChannel: true,
						sci:                "B4:F9:5D:8C:A7:AC/1",
						packetNumber:       4543932225,
					},
				}

				for _, conn := range result.MacsecConnectionInformation {
					interfaceName := conn.MacsecInterfaceCommonInformation.InterfaceName
					// Skip interfaces not in our test set
					expected, ok := expectedConnections[interfaceName]
					if !ok {
						continue
					}

					//test common interface information
					assert.Equal(t, expected.caName,
						conn.MacsecInterfaceCommonInformation.ConnectivityAssociationName,
						"CA name mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.encryption,
						conn.MacsecInterfaceCommonInformation.Encryption,
						"Encryption mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.includeSci,
						conn.MacsecInterfaceCommonInformation.IncludeSci,
						"Include-SCI mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.replayProtect,
						conn.MacsecInterfaceCommonInformation.ReplayProtect,
						"Replay-protect mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.offset,
						conn.MacsecInterfaceCommonInformation.Offset,
						"Offset mismatch for interface %s", interfaceName)

					// Test outbound secure channel
					if expected.hasOutboundChannel {
						assert.NotNil(t, conn.OutboundSecureChannel,
							"Interface %s should have outbound secure channel", interfaceName)

						assert.Equal(t, expected.sci,
							conn.OutboundSecureChannel.Sci,
							"SCI mismatch for interface %s", interfaceName)

						assert.Equal(t, expected.packetNumber,
							conn.OutboundSecureChannel.OutgoingPacketNumber,
							"Outgoing packet number mismatch for interface %s", interfaceName)

						assert.Equal(t, "inuse",
							conn.OutboundSecureChannel.OutboundSecureAssociation.AssociationNumberStatus,
							"Association number status mismatch for interface %s", interfaceName)
					} else if interfaceName == "et-0/0/1" {
						//et-0/0/1 should not have outbound secure channel
						assert.Nil(t, conn.OutboundSecureChannel,
							"Interface %s should not have outbound secure channel", interfaceName)
					}
				}
			},
		},
		{
			name:       "Parse MACsec Statistics",
			xmlData:    statisticsData,
			resultType: "statistics",
			validate: func(t *testing.T, data interface{}) {
				result := data.(*resultStats)
				stats := result.MacsecStatistics

				//define expected values for each interface
				expectedStats := map[string]struct {
					encryptedPackets uint64
					encryptedBytes   uint64
					protectedPackets uint64
					protectedBytes   uint64
					okPackets        uint64
					validatedBytes   uint64
					decryptedBytes   uint64
				}{
					"et-0/0/0": {
						encryptedPackets: 29470457654,
						encryptedBytes:   15371387814258,
						protectedPackets: 0,
						protectedBytes:   0,
						okPackets:        52215340924,
						validatedBytes:   0,
						decryptedBytes:   17104313476786,
					},
					"et-0/0/1": {
						encryptedPackets: 185933514,
						encryptedBytes:   52358181842,
						protectedPackets: 0,
						protectedBytes:   0,
						okPackets:        1404697306,
						validatedBytes:   0,
						decryptedBytes:   625824199534,
					},
					"et-0/0/6": {
						encryptedPackets: 4609309585,
						encryptedBytes:   1578896710717,
						protectedPackets: 0,
						protectedBytes:   0,
						okPackets:        98753547,
						validatedBytes:   0,
						decryptedBytes:   81188722443,
					},
				}

				// Find and validate statistics for each interface
				for i, interfaceName := range stats.Interfaces {
					// Skip interfaces not in our test set
					expected, ok := expectedStats[interfaceName]
					if !ok {
						continue
					}

					// Verify secure channel sent metrics
					assert.Equal(t, expected.encryptedPackets,
						stats.SecureChannelSent[i].EncryptedPackets,
						"Encrypted packet count mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.encryptedBytes,
						stats.SecureChannelSent[i].EncryptedBytes,
						"Encrypted bytes mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.protectedPackets,
						stats.SecureChannelSent[i].ProtectedPackets,
						"Protected packets mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.protectedBytes,
						stats.SecureChannelSent[i].ProtectedBytes,
						"Protected bytes mismatch for interface %s", interfaceName)

					// Verify secure association sent metrics
					assert.Equal(t, expected.encryptedPackets,
						stats.SecureAssociationSent[i].EncryptedPackets,
						"Secure association encrypted packets mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.protectedPackets,
						stats.SecureAssociationSent[i].ProtectedPackets,
						"Secure association protected packets mismatch for interface %s", interfaceName)

					// Verify secure channel received metrics
					assert.Equal(t, expected.okPackets,
						stats.SecureChannelReceived[i].OkPackets,
						"Received packet count mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.validatedBytes,
						stats.SecureChannelReceived[i].ValidatedBytes,
						"Validated bytes mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.decryptedBytes,
						stats.SecureChannelReceived[i].DecryptedBytes,
						"Decrypted bytes mismatch for interface %s", interfaceName)

					// Verify secure association received metrics
					assert.Equal(t, expected.okPackets,
						stats.SecureAssociationReceived[i].OkPackets,
						"Secure association received packet count mismatch for interface %s", interfaceName)

					assert.Equal(t, expected.validatedBytes,
						stats.SecureAssociationReceived[i].ValidatedBytes,
						"Secure association validated bytes mismatch for interface %s", interfaceName)

					assert.Equal(t, uint64(0),
						stats.SecureAssociationReceived[i].DecryptedBytes,
						"Secure association decrypted bytes mismatch for interface %s", interfaceName)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			switch tt.resultType {
			case "connections":
				result, err := ParseShowSecurityMacsecConnections(tt.xmlData)
				assert.NoError(t, err, "Should parse connections XML without error")
				tt.validate(t, result)
			case "statistics":
				var result resultStats
				err = xml.Unmarshal(tt.xmlData, &result)
				assert.NoError(t, err, "Should parse statistics XML without error")
				tt.validate(t, &result)
			}
		})
	}
}
