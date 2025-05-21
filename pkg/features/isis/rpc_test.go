// SPDX-License-Identifier: MIT

package isis

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testParseXML(t *testing.T) {
	resultBackupData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S3.9/junos">
    <isis-spf-information xmlns="http://xml.juniper.net/junos/23.4R0/junos-routing">
        <isis-spf>
            <isis-spf-results-header>
                <level>1</level>
            </isis-spf-results-header>
            <node-count>0</node-count>
        </isis-spf>
        <isis-spf>
            <isis-spf-results-header>
                <level>2</level>
            </isis-spf-results-header>
                <node-id>bbbb01.01.00</node-id>
                <node-address>0xd498000</node-address>
                <next-hop-element>
                    <interface-name>et-0/0/6.0</interface-name>
                    <isis-next-hop-type>IPV6</isis-next-hop-type>
                    <isis-next-hop>bbbb01.01</isis-next-hop>
                    <snpa> 20:20:20:20:20:20</snpa>
                </next-hop-element>
                <backup-root>bbbb01.01</backup-root>
                <backup-root-metric>50</backup-root-metric>
                <metric>0</metric>
                <backup-root-preference>0x0</backup-root-preference>
                <no-coverage-reason-element>
                    <isis-next-hop-type>IPV6</isis-next-hop-type>
                    <no-coverage-reason>Primary next-hop link fate sharing</no-coverage-reason>
                </no-coverage-reason-element>
                <backup-root>bbbb01.01</backup-root>
                <backup-root-metric>90</backup-root-metric>
                <metric>40</metric>
                <backup-root-preference>0x0</backup-root-preference>
                <track-item>bbbb01.01.00-00</track-item>
                <backup-next-hop-element>
                    <interface-name>et-0/0/1.0</interface-name>
                    <isis-next-hop-type>IPV6</isis-next-hop-type>
                    <isis-backup-prefix-refcount>4</isis-backup-prefix-refcount>
                    <isis-next-hop>bb01.ams01</isis-next-hop>
                    <snpa> 30:30:30:30:30:30</snpa>
                </backup-next-hop-element>
                <backup-root>drdr01.01</backup-root>
                <backup-root-metric>220</backup-root-metric>
                <metric>170</metric>
                <backup-root-preference>0x0</backup-root-preference>
                <track-item>bbbb01.01.00-00</track-item>
                <track-item>bbbb01.01.00-00</track-item>
                <no-coverage-reason-element>
                    <isis-next-hop-type>IPV6</isis-next-hop-type>
                    <no-coverage-reason>Interface is already covered</no-coverage-reason>
                </no-coverage-reason-element>
                <backup-root>drdr01.01</backup-root>
                <backup-root-metric>170</backup-root-metric>
                <metric>220</metric>
                <backup-root-preference>0x0</backup-root-preference>
                <track-item>bbbb01.01.00-00</track-item>
                <track-item>bbbb01.01.00-00</track-item>
                <no-coverage-reason-element>
                    <isis-next-hop-type>IPV6</isis-next-hop-type>
                    <no-coverage-reason>Interface is already covered</no-coverage-reason>
                </no-coverage-reason-element>
            </isis-backup-spf-result>
            <node-count>12</node-count>
        </isis-spf>
    </isis-spf-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	resultsCoverageData := `
	<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S3.9/junos">
<isis-backup-coverage-information xmlns="http://xml.juniper.net/junos/23.4R0/junos-routing">
	<isis-backup-coverage>
	<isis-topology-id>IPV4 Unicast</isis-topology-id>
	<level>2</level>
	<isis-node-coverage>97.77%</isis-node-coverage>
	<isis-route-coverage-ipv4>0.00%</isis-route-coverage-ipv4>
	<isis-route-coverage-ipv6>99.99%</isis-route-coverage-ipv6>
	<isis-route-coverage-clns>0.00%</isis-route-coverage-clns>
	<isis-route-coverage-ipv4-mpls>0.00%</isis-route-coverage-ipv4-mpls>
	<isis-route-coverage-ipv6-mpls>0.00%</isis-route-coverage-ipv6-mpls>
	<isis-route-coverage-ipv4-mpls-sspf>0.00%</isis-route-coverage-ipv4-mpls-sspf>
	<isis-route-coverage-ipv6-mpls-sspf>0.00%</isis-route-coverage-ipv6-mpls-sspf>
	</isis-backup-coverage>
	</isis-backup-coverage-information>
	<cli>
	<banner></banner>
	</cli>
	</rpc-reply>`

	var resultsSPF backupSPF
	var resultsCoverage backupCoverage
	// need to check wha everything passes
	err := xml.Unmarshal([]byte(resultBackupData), &resultsSPF)
	assert.NoError(t, err)
	assert.Len(t, resultsSPF.IsisSpfInformation.IsisSpf, 1)
	assert.Len(t, resultsSPF.IsisSpfInformation.IsisSpf[0].IsisBackupSpfResult[0], 1)

	assert.Equal(t, "30:30:30:30:30:30", resultsSPF.IsisSpfInformation.IsisSpf[0].IsisBackupSpfResult[0].BackupNextHopElement.SNPA)
	assert.Empty(t, resultsSPF.IsisSpfInformation.IsisSpf[0].IsisBackupSpfResult[1])

	err = xml.Unmarshal([]byte(resultsCoverageData), &resultsCoverage)
	assert.NoError(t, err)
	assert.Len(t, resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage, 1)

	assert.Equal(t, "97.77%", resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage.IsisNodeCoverage)
	assert.Equal(t, "99.99%", resultsCoverage.IsisBackupCoverageInformation.IsisBackupCoverage.IsisRouteCoverageIpv6)
}
