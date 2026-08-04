// SPDX-License-Identifier: MIT

package mnha

import (
	"context"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

const srgSampleGroup1Offline = `<rpc-reply>
    <chassis-high-availability-default-srg-info>
        <current-state>OFFLINE</current-state>
        <peer-id>2</peer-id>
    </chassis-high-availability-default-srg-info>
</rpc-reply>`

// mnhaMockClient is a minimal collector.Client fake, modelled on the one used
// in pkg/features/interfaces/collector_test.go. It serves canned XML bodies
// keyed by the exact command string and records every command it was asked
// to run, so tests can assert on dispatch (which commands, how many times).
type mnhaMockClient struct {
	commands  []string
	responses map[string]string
	failing   map[string]error
}

func (m *mnhaMockClient) RunCommandAndParse(cmd string, obj any) error {
	m.commands = append(m.commands, cmd)

	if err, ok := m.failing[cmd]; ok {
		return err
	}

	body, ok := m.responses[cmd]
	if !ok {
		return fmt.Errorf("mnhaMockClient: no canned response for %q", cmd)
	}

	return xml.Unmarshal([]byte(body), obj)
}

func (m *mnhaMockClient) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	m.commands = append(m.commands, cmd)

	body, ok := m.responses[cmd]
	if !ok {
		return fmt.Errorf("mnhaMockClient: no canned response for %q", cmd)
	}

	return parser([]byte(body))
}

func (m *mnhaMockClient) IsSatelliteEnabled() bool {
	return false
}

func (m *mnhaMockClient) IsScrapingLicenseEnabled() bool {
	return false
}

func (m *mnhaMockClient) Device() *connector.Device {
	return &connector.Device{Host: "srx1"}
}

func (m *mnhaMockClient) Context() context.Context {
	return context.Background()
}

func drain(ch chan prometheus.Metric) []prometheus.Metric {
	close(ch)
	var out []prometheus.Metric
	for m := range ch {
		out = append(out, m)
	}
	return out
}

const detailCmd = "show chassis high-availability information detail"

func srgCmd(id int) string {
	return fmt.Sprintf("show chassis high-availability services-redundancy-group %d", id)
}

func TestCollectDefaultGroupReusesDetailResponse(t *testing.T) {
	client := &mnhaMockClient{
		responses: map[string]string{
			detailCmd: detailSample,
		},
	}

	col := NewCollector(nil).(*mnhaCollector) // nil -> defaults to [0]
	ch := make(chan prometheus.Metric, 64)

	err := col.Collect(client, ch, []string{"srx1"})
	assert.NoError(t, err)

	metrics := drain(ch)
	assert.Equal(t, []string{detailCmd}, client.commands, "SRG 0 must be served from the detail response, not a second round-trip")

	srgMetrics := findByDesc(metrics, srgStateDesc)
	assert.Len(t, srgMetrics, 1)
	labels := labelMap(t, srgMetrics[0])
	assert.Equal(t, "0", labels["srg_id"])
	assert.Equal(t, "2", labels["peer_id"], "from the embedded default-srg-info block, peer-id 2")
	assert.Equal(t, "ONLINE", labels["state"])
}

func TestCollectMultipleGroupsQueriesOnlyNonDefault(t *testing.T) {
	client := &mnhaMockClient{
		responses: map[string]string{
			detailCmd: detailSample,
			srgCmd(1): srgSampleGroup1Offline,
		},
	}

	col := NewCollector([]int{0, 1}).(*mnhaCollector)
	ch := make(chan prometheus.Metric, 64)

	err := col.Collect(client, ch, []string{"srx1"})
	assert.NoError(t, err)

	metrics := drain(ch)
	assert.Equal(t, []string{detailCmd, srgCmd(1)}, client.commands, "group 0 must not trigger its own RPC when detail already succeeded")

	srgMetrics := findByDesc(metrics, srgStateDesc)
	assert.Len(t, srgMetrics, 2)

	byGroup := map[string]map[string]string{}
	for _, m := range srgMetrics {
		l := labelMap(t, m)
		byGroup[l["srg_id"]] = l
	}

	assert.Equal(t, "ONLINE", byGroup["0"]["state"])
	assert.Equal(t, "OFFLINE", byGroup["1"]["state"])
}

func TestCollectFallsBackToExplicitQueryWhenDetailFails(t *testing.T) {
	client := &mnhaMockClient{
		responses: map[string]string{
			srgCmd(0): srgSample,
		},
		failing: map[string]error{
			detailCmd: fmt.Errorf("permission denied"),
		},
	}

	col := NewCollector(nil).(*mnhaCollector)
	ch := make(chan prometheus.Metric, 64)

	err := col.Collect(client, ch, []string{"srx1"})
	assert.Error(t, err, "the detail command's error must still surface")
	assert.Equal(t, []string{detailCmd, srgCmd(0)}, client.commands, "must fall back to an explicit SRG 0 query once detail fails")

	metrics := drain(ch)
	srgMetrics := findByDesc(metrics, srgStateDesc)
	assert.Len(t, srgMetrics, 1, "SRG state should still be collected despite the detail failure")
}

func TestCollectJoinsErrorsFromBothCalls(t *testing.T) {
	client := &mnhaMockClient{
		responses: map[string]string{},
		failing: map[string]error{
			detailCmd: fmt.Errorf("detail failed"),
			srgCmd(0): fmt.Errorf("srg failed"),
		},
	}

	col := NewCollector(nil).(*mnhaCollector)
	ch := make(chan prometheus.Metric, 64)

	err := col.Collect(client, ch, []string{"srx1"})
	assert.Error(t, err)
	assert.ErrorContains(t, err, "detail failed")
	assert.ErrorContains(t, err, "srg failed")

	metrics := drain(ch)
	assert.Empty(t, metrics)
}
