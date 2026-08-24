// SPDX-License-Identifier: MIT

package securitynat

import "encoding/xml"

// poolUsageResult is the response to "show security nat resource-usage source-pool all".
type poolUsageResult struct {
	XMLName xml.Name      `xml:"rpc-reply"`
	Info    poolUsageInfo `xml:"source-resource-usage-pool-information"`
}

type poolUsageInfo struct {
	Entries []poolUsageEntry `xml:"resource-usage-entry"`
}

// poolUsageEntry is the resource usage of a single source NAT pool, e.g.:
//
//	<resource-usage-entry junos:style="all-pat-pool">
//	    <resource-usage-pool-name>POOL-A</resource-usage-pool-name>
//	    <resource-usage-port-ol-factor>1</resource-usage-port-ol-factor>
//	    <resource-usage-peak-usage>31%</resource-usage-peak-usage>
//	    <resource-usage-peak-date-time junos:seconds="1786981339">…</resource-usage-peak-date-time>
//	    <resource-usage-total-address>1</resource-usage-total-address>
//	    <resource-usage-total-used>8967</resource-usage-total-used>
//	    <resource-usage-total-avail>55545</resource-usage-total-avail>
//	    <resource-usage-total-total>64512</resource-usage-total-total>
//	    <resource-usage-total-usage>13%</resource-usage-total-usage>
//	</resource-usage-entry>
//
// The unit of the used/avail/total triple depends on the pool style: ports for
// port-translating (PAT) pools, addresses otherwise. Hence the neutral
// "resource" naming, mirroring the RPC itself.
type poolUsageEntry struct {
	Style        string  `xml:"style,attr"`
	PoolName     string  `xml:"resource-usage-pool-name"`
	PortOlFactor float64 `xml:"resource-usage-port-ol-factor"`
	PeakUsage    string  `xml:"resource-usage-peak-usage"`
	PeakDateTime struct {
		Seconds int64 `xml:"seconds,attr"`
	} `xml:"resource-usage-peak-date-time"`
	TotalAddress float64 `xml:"resource-usage-total-address"`
	TotalUsed    float64 `xml:"resource-usage-total-used"`
	TotalAvail   float64 `xml:"resource-usage-total-avail"`
	TotalTotal   float64 `xml:"resource-usage-total-total"`
	TotalUsage   string  `xml:"resource-usage-total-usage"`
}

// ruleResult is the response to "show security nat source rule all".
type ruleResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Info    ruleInfo `xml:"source-nat-rule-detail-information"`
}

type ruleInfo struct {
	TotalRules struct {
		Total float64 `xml:"total-src-rules"`
	} `xml:"total-source-nat-rules"`
	TotalRefAddr struct {
		V4 float64 `xml:"total-source-nat-rule-ref-addr-num-v4"`
		V6 float64 `xml:"total-source-nat-rule-ref-addr-num-v6"`
	} `xml:"total-source-nat-rule-ref-addr-num"`
	Rules []ruleEntry `xml:"source-nat-rule-entry"`
}

type ruleEntry struct {
	Name    string `xml:"rule-name"`
	SetName string `xml:"rule-set-name"`
	ID      string `xml:"rule-id"`
	// A rule can match traffic from several zones/interfaces, so Junos repeats
	// the context-name element once per entry.
	FromContextName []string   `xml:"rule-from-context-name"`
	ToContextName   []string   `xml:"rule-to-context-name"`
	Action          ruleAction `xml:"source-nat-rule-action-entry"`
	// Hits is absent for rules Junos has no counters for, which is not the same
	// as a rule with zero hits, so keep it nullable.
	Hits *ruleHits `xml:"source-nat-rule-hits-entry"`
}

type ruleAction struct {
	Pool string `xml:"source-nat-rule-action"`
}

type ruleHits struct {
	TranslationHits float64 `xml:"rule-translation-hits"`
	SuccessHits     float64 `xml:"succ-hits"`
	ConcurrentHits  float64 `xml:"concurrent-hits"`
}
