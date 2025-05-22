// SPDX-License-Identifier: MIT

package isis

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseXML_DataDriven(t *testing.T) {
	backupCoverageXMLData, _ := os.Open("backupCoverageXMLData.xml")
	backupSPFXMLData, _ := os.Open("backupSPFXMLData.xml")
	backupCoverageData, _ := ioutil.ReadAll(backupCoverageXMLData)
	backupSPFData, _ := ioutil.ReadAll(backupSPFXMLData)
	tests := []struct {
		name       string
		xmlData    string
		resultType string
		validate   func(t *testing.T, data interface{})
	}{
		{
			name:       "Parse Backup SPF Data",
			xmlData:    string(backupSPFData),
			resultType: "spf",
			validate: func(t *testing.T, data interface{}) {
				resultsSPF := data.(*backupSPF)
				assert.Len(t, resultsSPF.IsisSpfInformation.IsisSpf, 2)
				assert.Len(t, resultsSPF.IsisSpfInformation.IsisSpf[1].IsisBackupSpfResult, 1)
				assert.Equal(t, "30:30:30:30:30:30", strings.TrimSpace(resultsSPF.IsisSpfInformation.IsisSpf[1].IsisBackupSpfResult[0].BackupNextHopElement.SNPA))
			},
		},
		{
			name:       "Parse Backup Coverage Data",
			xmlData:    string(backupCoverageData),
			resultType: "coverage",
			validate: func(t *testing.T, data interface{}) {
				resultsCoverage := data.(*backupCoverage)
				assert.Len(t, resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage.Level, 1)
				assert.Len(t, resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage.IsisRouteCoverageIpv6, 6)
				assert.Equal(t, "97.77%", resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage.IsisNodeCoverage)
				assert.Equal(t, "99.99%", resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage.IsisRouteCoverageIpv6)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			switch tt.resultType {
			case "spf":
				var result backupSPF
				err = xml.Unmarshal([]byte(tt.xmlData), &result)
				assert.Equal(t, nil, err)
				assert.NoError(t, err)
				tt.validate(t, &result)
			case "coverage":
				var result backupCoverage
				err = xml.Unmarshal([]byte(tt.xmlData), &result)
				assert.Equal(t, nil, err)
				assert.NoError(t, err)
				tt.validate(t, &result)
			}
		})
	}
}
