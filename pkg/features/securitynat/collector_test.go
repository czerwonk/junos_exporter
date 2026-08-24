// SPDX-License-Identifier: MIT

package securitynat

import (
	"context"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

// mockClient is a minimal collector.Client fake, modelled on the one used in
// pkg/features/mnha/collect_test.go. It serves canned XML bodies keyed by the
// exact command string and records every command it was asked to run.
type mockClient struct {
	commands  []string
	responses map[string]string
	failing   map[string]error
}

func (m *mockClient) RunCommandAndParse(cmd string, obj any) error {
	m.commands = append(m.commands, cmd)

	if err, ok := m.failing[cmd]; ok {
		return err
	}

	body, ok := m.responses[cmd]
	if !ok {
		return fmt.Errorf("mockClient: no canned response for %q", cmd)
	}

	return xml.Unmarshal([]byte(body), obj)
}

func (m *mockClient) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	m.commands = append(m.commands, cmd)

	body, ok := m.responses[cmd]
	if !ok {
		return fmt.Errorf("mockClient: no canned response for %q", cmd)
	}

	return parser([]byte(body))
}

func (m *mockClient) IsSatelliteEnabled() bool {
	return false
}

func (m *mockClient) IsScrapingLicenseEnabled() bool {
	return false
}

func (m *mockClient) Device() *connector.Device {
	return &connector.Device{Host: "srx1"}
}

func (m *mockClient) Context() context.Context {
	return context.Background()
}

func clientWithSamples() *mockClient {
	return &mockClient{
		responses: map[string]string{
			poolUsageRPC: poolUsageSample,
			ruleRPC:      ruleSample,
		},
	}
}

func collect(t *testing.T, client *mockClient) []prometheus.Metric {
	t.Helper()

	ch := make(chan prometheus.Metric, 128)
	err := NewCollector().Collect(client, ch, []string{"srx1"})
	assert.NoError(t, err)

	return drain(ch)
}

func drain(ch chan prometheus.Metric) []prometheus.Metric {
	close(ch)

	var out []prometheus.Metric
	for m := range ch {
		out = append(out, m)
	}

	return out
}

func findByDesc(metrics []prometheus.Metric, desc *prometheus.Desc) []prometheus.Metric {
	var out []prometheus.Metric
	for _, m := range metrics {
		if m.Desc() == desc {
			out = append(out, m)
		}
	}

	return out
}

func labelMap(t *testing.T, m prometheus.Metric) map[string]string {
	t.Helper()

	var dm dto.Metric
	if err := m.Write(&dm); err != nil {
		t.Fatalf("writing metric: %v", err)
	}

	out := make(map[string]string, len(dm.Label))
	for _, lp := range dm.Label {
		out[lp.GetName()] = lp.GetValue()
	}

	return out
}

func metricValue(t *testing.T, m prometheus.Metric) float64 {
	t.Helper()

	var dm dto.Metric
	if err := m.Write(&dm); err != nil {
		t.Fatalf("writing metric: %v", err)
	}

	if dm.Gauge != nil {
		return dm.Gauge.GetValue()
	}

	return dm.Counter.GetValue()
}

// byLabel indexes metrics of one family by the value of a single label.
func byLabel(t *testing.T, metrics []prometheus.Metric, label string) map[string]float64 {
	t.Helper()

	out := make(map[string]float64, len(metrics))
	for _, m := range metrics {
		out[labelMap(t, m)[label]] = metricValue(t, m)
	}

	return out
}

func TestCollectPoolUsage(t *testing.T) {
	metrics := collect(t, clientWithSamples())

	assert.Equal(t, map[string]float64{"POOL-A": 13, "POOL-B": 0},
		byLabel(t, findByDesc(metrics, poolUsagePercentDesc), "pool_name"))
	assert.Equal(t, map[string]float64{"POOL-A": 31, "POOL-B": 0},
		byLabel(t, findByDesc(metrics, poolPeakUsagePercentDesc), "pool_name"))
	assert.Equal(t, map[string]float64{"POOL-A": 8967, "POOL-B": 0},
		byLabel(t, findByDesc(metrics, poolResourceUsedDesc), "pool_name"))
	assert.Equal(t, map[string]float64{"POOL-A": 55545, "POOL-B": 64512},
		byLabel(t, findByDesc(metrics, poolResourceAvailableDesc), "pool_name"))
	assert.Equal(t, map[string]float64{"POOL-A": 64512, "POOL-B": 64512},
		byLabel(t, findByDesc(metrics, poolResourceTotalDesc), "pool_name"))

	used := findByDesc(metrics, poolResourceUsedDesc)
	assert.Equal(t, map[string]string{"target": "srx1", "pool_name": "POOL-A", "style": "all-pat-pool"},
		labelMap(t, used[0]))
}

// POOL-B never recorded a peak, so it must not claim one at the unix epoch.
func TestCollectSkipsUnsetPeakTimestamp(t *testing.T) {
	metrics := collect(t, clientWithSamples())

	peaks := byLabel(t, findByDesc(metrics, poolPeakUsageTimestampSecondsDesc), "pool_name")
	assert.Equal(t, map[string]float64{"POOL-A": 1786981339}, peaks)
}

func TestCollectRules(t *testing.T) {
	metrics := collect(t, clientWithSamples())

	total := findByDesc(metrics, rulesCountDesc)
	assert.Len(t, total, 1)
	assert.Equal(t, float64(2), metricValue(t, total[0]))
	assert.Equal(t, map[string]string{"target": "srx1"}, labelMap(t, total[0]))

	assert.Equal(t, map[string]float64{"ipv4": 10, "ipv6": 0},
		byLabel(t, findByDesc(metrics, ruleReferencedAddressCountDesc), "ip_version"))

	assert.Equal(t, map[string]float64{"RULE-A": 228555982, "RULE-B": 915},
		byLabel(t, findByDesc(metrics, ruleTranslationHitsTotalDesc), "rule_name"))
	assert.Equal(t, map[string]float64{"RULE-A": 203932849, "RULE-B": 915},
		byLabel(t, findByDesc(metrics, ruleTranslationSuccessHitsTotalDesc), "rule_name"))
	assert.Equal(t, map[string]float64{"RULE-A": 8230, "RULE-B": 0},
		byLabel(t, findByDesc(metrics, ruleConcurrentHitsDesc), "rule_name"))

	hits := findByDesc(metrics, ruleTranslationHitsTotalDesc)
	for _, m := range hits {
		l := labelMap(t, m)
		if l["rule_name"] != "RULE-A" {
			continue
		}

		assert.Equal(t, map[string]string{
			"target":        "srx1",
			"rule_name":     "RULE-A",
			"rule_set_name": "RULESET-A",
			"rule_id":       "1",
			"pool_name":     "POOL-A",
		}, l, "the zones a rule matches must stay off the counters")
	}
}

// A rule matching several zones gets one info series per zone pair, so the zones
// are queryable without parsing a delimited label value.
func TestCollectRuleInfoFansOutOverZones(t *testing.T) {
	metrics := collect(t, clientWithSamples())

	info := findByDesc(metrics, ruleInfoDesc)
	assert.Len(t, info, 3, "RULE-A matches 2 from-zones, RULE-B matches 1")

	got := make(map[string]struct{}, len(info))
	for _, m := range info {
		assert.Equal(t, float64(1), metricValue(t, m))

		l := labelMap(t, m)
		got[fmt.Sprintf("%s|%s|%s", l["rule_name"], l["from_zone"], l["to_zone"])] = struct{}{}
	}

	assert.Equal(t, map[string]struct{}{
		"RULE-A|ZONE-TRUST-1|ZONE-INTERNET": {},
		"RULE-A|ZONE-TRUST-2|ZONE-INTERNET": {},
		"RULE-B|ZONE-TRUST-1|ZONE-PARTNER":  {},
	}, got)
}

// Rules Junos reports no counters for still need to show up as configured, so
// the info series is emitted while the counters are skipped.
func TestCollectRuleWithoutHitsEmitsInfoOnly(t *testing.T) {
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

	client := clientWithSamples()
	client.responses[ruleRPC] = body

	metrics := collect(t, client)

	assert.Empty(t, findByDesc(metrics, ruleTranslationHitsTotalDesc))
	assert.Empty(t, findByDesc(metrics, ruleConcurrentHitsDesc))

	info := findByDesc(metrics, ruleInfoDesc)
	assert.Len(t, info, 1)
	assert.Equal(t, map[string]string{
		"target":        "srx1",
		"rule_name":     "RULE-NO-HITS",
		"rule_set_name": "RULESET-A",
		"rule_id":       "7",
		"from_zone":     "",
		"to_zone":       "ZONE-INTERNET",
	}, labelMap(t, info[0]))
}

// A device without any source NAT configured answers both RPCs with empty
// bodies, which must not be treated as an error.
func TestCollectWithoutAnyNATConfigured(t *testing.T) {
	client := &mockClient{
		responses: map[string]string{
			poolUsageRPC: `<rpc-reply><source-resource-usage-pool-information></source-resource-usage-pool-information></rpc-reply>`,
			ruleRPC:      `<rpc-reply><source-nat-rule-detail-information></source-nat-rule-detail-information></rpc-reply>`,
		},
	}

	metrics := collect(t, client)

	assert.Empty(t, findByDesc(metrics, poolResourceUsedDesc))
	assert.Empty(t, findByDesc(metrics, ruleInfoDesc))

	rules := findByDesc(metrics, rulesCountDesc)
	assert.Len(t, rules, 1)
	assert.Equal(t, float64(0), metricValue(t, rules[0]))
}

func TestCollectQueriesBothRPCs(t *testing.T) {
	client := clientWithSamples()
	collect(t, client)

	assert.Equal(t, []string{poolUsageRPC, ruleRPC}, client.commands)
}

// One failing RPC must not hide the other's error, nor abort the scrape.
func TestCollectJoinsErrorsFromBothCalls(t *testing.T) {
	client := &mockClient{
		responses: map[string]string{},
		failing: map[string]error{
			poolUsageRPC: fmt.Errorf("pool usage failed"),
			ruleRPC:      fmt.Errorf("rule query failed"),
		},
	}

	ch := make(chan prometheus.Metric, 128)
	err := NewCollector().Collect(client, ch, []string{"srx1"})

	assert.ErrorContains(t, err, "pool usage failed")
	assert.ErrorContains(t, err, "rule query failed")
}

func TestCollectStillReportsRulesWhenPoolUsageFails(t *testing.T) {
	client := clientWithSamples()
	client.failing = map[string]error{poolUsageRPC: fmt.Errorf("pool usage failed")}

	ch := make(chan prometheus.Metric, 128)
	err := NewCollector().Collect(client, ch, []string{"srx1"})
	assert.ErrorContains(t, err, "pool usage failed")

	metrics := drain(ch)
	assert.Empty(t, findByDesc(metrics, poolResourceUsedDesc))
	assert.Len(t, findByDesc(metrics, ruleTranslationHitsTotalDesc), 2, "a failing pool query must not drop rule metrics")
}

// Every metric this collector describes must be registerable together, which
// catches duplicate names and label-set mismatches between Describe and Collect.
func TestMetricsAreConsistentWithRegistry(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()

	c := &wrappedCollector{client: clientWithSamples()}
	assert.NoError(t, reg.Register(c))

	_, err := reg.Gather()
	assert.NoError(t, err)
}

// wrappedCollector adapts the RPCCollector to prometheus.Collector so the
// pedantic registry can validate the emitted metrics.
type wrappedCollector struct {
	client *mockClient
}

func (w *wrappedCollector) Describe(ch chan<- *prometheus.Desc) {
	NewCollector().Describe(ch)
}

func (w *wrappedCollector) Collect(ch chan<- prometheus.Metric) {
	if err := NewCollector().Collect(w.client, ch, []string{"srx1"}); err != nil {
		panic(err)
	}
}
