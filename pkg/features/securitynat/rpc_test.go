// SPDX-License-Identifier: MIT

package securitynat

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The sample XML in this package is a sanitized version of real "show security
// nat ... | display xml" output: pool, rule and zone names are replaced with
// placeholders and addresses come from the RFC 5737 documentation ranges. The
// structure and the element names are unchanged.

const poolUsageSample = `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/24.4R2-S2.6/junos">
    <source-resource-usage-pool-information xmlns="http://xml.juniper.net/junos/24.4R0/junos-nat">
        <resource-usage-entry junos:style="all-pat-pool">
            <resource-usage-pool-name>POOL-A</resource-usage-pool-name>
            <resource-usage-total-pool-num>2</resource-usage-total-pool-num>
            <resource-usage-port-ol-factor>1</resource-usage-port-ol-factor>
            <resource-usage-peak-usage>31%</resource-usage-peak-usage>
            <resource-usage-peak-date-time junos:seconds="1786981339">
                2026-08-17 18:42:19 UTC</resource-usage-peak-date-time>
            <resource-usage-total-address>1</resource-usage-total-address>
            <resource-usage-total-used>8967</resource-usage-total-used>
            <resource-usage-total-avail>55545</resource-usage-total-avail>
            <resource-usage-total-total>64512</resource-usage-total-total>
            <resource-usage-total-usage>13%</resource-usage-total-usage>
        </resource-usage-entry>
        <resource-usage-entry junos:style="all-pat-pool">
            <resource-usage-pool-name>POOL-B</resource-usage-pool-name>
            <resource-usage-total-pool-num>2</resource-usage-total-pool-num>
            <resource-usage-port-ol-factor>1</resource-usage-port-ol-factor>
            <resource-usage-peak-usage>0%</resource-usage-peak-usage>
            <resource-usage-peak-date-time junos:seconds="0">
                </resource-usage-peak-date-time>
            <resource-usage-total-address>1</resource-usage-total-address>
            <resource-usage-total-used>0</resource-usage-total-used>
            <resource-usage-total-avail>64512</resource-usage-total-avail>
            <resource-usage-total-total>64512</resource-usage-total-total>
            <resource-usage-total-usage>0%</resource-usage-total-usage>
        </resource-usage-entry>
    </source-resource-usage-pool-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

func TestParsePoolUsageResult(t *testing.T) {
	var res poolUsageResult
	err := xml.Unmarshal([]byte(poolUsageSample), &res)
	assert.NoError(t, err)

	assert.Len(t, res.Info.Entries, 2)

	a := res.Info.Entries[0]
	assert.Equal(t, "all-pat-pool", a.Style, "the junos: prefixed style attribute must be picked up")
	assert.Equal(t, "POOL-A", a.PoolName)
	assert.Equal(t, float64(1), a.PortOlFactor)
	assert.Equal(t, "31%", a.PeakUsage)
	assert.Equal(t, int64(1786981339), a.PeakDateTime.Seconds)
	assert.Equal(t, float64(1), a.TotalAddress)
	assert.Equal(t, float64(8967), a.TotalUsed)
	assert.Equal(t, float64(55545), a.TotalAvail)
	assert.Equal(t, float64(64512), a.TotalTotal)
	assert.Equal(t, "13%", a.TotalUsage)

	b := res.Info.Entries[1]
	assert.Equal(t, "POOL-B", b.PoolName)
	assert.Equal(t, "0%", b.TotalUsage)
	assert.Equal(t, int64(0), b.PeakDateTime.Seconds, "a pool without a recorded peak reports no timestamp")
}

const ruleSample = `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/24.4R2-S2.6/junos">
    <source-nat-rule-detail-information xmlns="http://xml.juniper.net/junos/24.4R0/junos-nat">
        <total-source-nat-rules>
            <total-src-rules>2</total-src-rules>
        </total-source-nat-rules>
        <total-source-nat-rule-ref-addr-num>
            <total-source-nat-rule-ref-addr-num-v4>10</total-source-nat-rule-ref-addr-num-v4>
            <total-source-nat-rule-ref-addr-num-v6>0</total-source-nat-rule-ref-addr-num-v6>
        </total-source-nat-rule-ref-addr-num>
        <source-nat-rule-entry>
            <rule-name>RULE-A</rule-name>
            <rule-set-name>RULESET-A</rule-set-name>
            <rule-id>1</rule-id>
            <rule-matching-position>1</rule-matching-position>
            <rule-from-context>zone</rule-from-context>
            <rule-from-context-name>ZONE-TRUST-1</rule-from-context-name>
            <rule-from-context-name>ZONE-TRUST-2</rule-from-context-name>
            <rule-to-context>zone</rule-to-context>
            <rule-to-context-name>ZONE-INTERNET</rule-to-context-name>
            <source-address-range-entry>
                <rule-source-address-low-range>0.0.0.0</rule-source-address-low-range>
                <rule-source-address-high-range>255.255.255.255</rule-source-address-high-range>
            </source-address-range-entry>
            <destination-address-range-entry>
                <rule-destination-address-low-range>0.0.0.0</rule-destination-address-low-range>
                <rule-destination-address-high-range>255.255.255.255</rule-destination-address-high-range>
            </destination-address-range-entry>
            <destination-port-entry>
            </destination-port-entry>
            <source-port-entry>
            </source-port-entry>
            <src-nat-protocol-entry>
            </src-nat-protocol-entry>
            <source-nat-rule-action-entry>
                <source-nat-rule-action>POOL-A</source-nat-rule-action>
                <persistent-nat-type>N/A </persistent-nat-type>
                <persistent-nat-mapping-type>address-port-mapping </persistent-nat-mapping-type>
                <persistent-nat-timeout>0</persistent-nat-timeout>
                <persistent-nat-max-session>0</persistent-nat-max-session>
                <persistent-nat-blocksess>disabled </persistent-nat-blocksess>
            </source-nat-rule-action-entry>
            <source-nat-rule-hits-entry>
                <rule-translation-hits>228555982</rule-translation-hits>
                <succ-hits>203932849</succ-hits>
                <concurrent-hits>8230</concurrent-hits>
            </source-nat-rule-hits-entry>
        </source-nat-rule-entry>
        <source-nat-rule-entry>
            <rule-name>RULE-B</rule-name>
            <rule-set-name>RULESET-B</rule-set-name>
            <rule-id>2</rule-id>
            <rule-matching-position>2</rule-matching-position>
            <rule-from-context>zone</rule-from-context>
            <rule-from-context-name>ZONE-TRUST-1</rule-from-context-name>
            <rule-to-context>zone</rule-to-context>
            <rule-to-context-name>ZONE-PARTNER</rule-to-context-name>
            <source-address-range-entry>
                <rule-source-address-low-range>192.0.2.1</rule-source-address-low-range>
                <rule-source-address-high-range>192.0.2.1</rule-source-address-high-range>
            </source-address-range-entry>
            <destination-address-range-entry>
                <rule-destination-address-low-range>198.51.100.0</rule-destination-address-low-range>
                <rule-destination-address-high-range>198.51.100.255</rule-destination-address-high-range>
            </destination-address-range-entry>
            <destination-port-entry>
            </destination-port-entry>
            <source-port-entry>
            </source-port-entry>
            <src-nat-protocol-entry>
            </src-nat-protocol-entry>
            <source-nat-rule-action-entry>
                <source-nat-rule-action>POOL-B</source-nat-rule-action>
                <persistent-nat-type>N/A </persistent-nat-type>
                <persistent-nat-mapping-type>address-port-mapping </persistent-nat-mapping-type>
                <persistent-nat-timeout>0</persistent-nat-timeout>
                <persistent-nat-max-session>0</persistent-nat-max-session>
                <persistent-nat-blocksess>disabled </persistent-nat-blocksess>
            </source-nat-rule-action-entry>
            <source-nat-rule-hits-entry>
                <rule-translation-hits>915</rule-translation-hits>
                <succ-hits>915</succ-hits>
                <concurrent-hits>0</concurrent-hits>
            </source-nat-rule-hits-entry>
        </source-nat-rule-entry>
    </source-nat-rule-detail-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

func TestParseRuleResult(t *testing.T) {
	var res ruleResult
	err := xml.Unmarshal([]byte(ruleSample), &res)
	assert.NoError(t, err)

	assert.Equal(t, float64(2), res.Info.TotalRules.Total)
	assert.Equal(t, float64(10), res.Info.TotalRefAddr.V4)
	assert.Equal(t, float64(0), res.Info.TotalRefAddr.V6)

	assert.Len(t, res.Info.Rules, 2)

	a := res.Info.Rules[0]
	assert.Equal(t, "RULE-A", a.Name)
	assert.Equal(t, "RULESET-A", a.SetName)
	assert.Equal(t, "1", a.ID)
	assert.Equal(t, []string{"ZONE-TRUST-1", "ZONE-TRUST-2"}, a.FromContextName, "every repeated context element must be kept")
	assert.Equal(t, []string{"ZONE-INTERNET"}, a.ToContextName)
	assert.Equal(t, "POOL-A", a.Action.Pool)

	assert.NotNil(t, a.Hits)
	assert.Equal(t, float64(228555982), a.Hits.TranslationHits)
	assert.Equal(t, float64(203932849), a.Hits.SuccessHits)
	assert.Equal(t, float64(8230), a.Hits.ConcurrentHits)

	b := res.Info.Rules[1]
	assert.Equal(t, "RULE-B", b.Name)
	assert.Equal(t, "POOL-B", b.Action.Pool)
	assert.Equal(t, float64(915), b.Hits.TranslationHits)
	assert.Equal(t, float64(0), b.Hits.ConcurrentHits)
}

// A rule Junos reports no counters for must be distinguishable from a rule with
// zero hits, so the hits element stays nil instead of unmarshalling to zeroes.
func TestParseRuleResultWithoutHits(t *testing.T) {
	const body = `<rpc-reply>
    <source-nat-rule-detail-information>
        <total-source-nat-rules>
            <total-src-rules>1</total-src-rules>
        </total-source-nat-rules>
        <source-nat-rule-entry>
            <rule-name>RULE-NO-HITS</rule-name>
            <rule-set-name>RULESET-A</rule-set-name>
            <rule-id>7</rule-id>
            <rule-to-context-name>ZONE-INTERNET</rule-to-context-name>
            <source-nat-rule-action-entry>
                <source-nat-rule-action>interface</source-nat-rule-action>
            </source-nat-rule-action-entry>
        </source-nat-rule-entry>
    </source-nat-rule-detail-information>
</rpc-reply>`

	var res ruleResult
	err := xml.Unmarshal([]byte(body), &res)
	assert.NoError(t, err)

	assert.Len(t, res.Info.Rules, 1)
	assert.Nil(t, res.Info.Rules[0].Hits)
	assert.Empty(t, res.Info.Rules[0].FromContextName)
}

func TestParsePercent(t *testing.T) {
	v, ok := parsePercent("31%")
	assert.True(t, ok)
	assert.Equal(t, float64(31), v)

	v, ok = parsePercent(" 0% ")
	assert.True(t, ok)
	assert.Equal(t, float64(0), v)

	_, ok = parsePercent("N/A")
	assert.False(t, ok)

	_, ok = parsePercent("")
	assert.False(t, ok)
}
