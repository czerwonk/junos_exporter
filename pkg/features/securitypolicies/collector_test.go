// SPDX-License-Identifier: MIT

package securitypolicies

import (
	"context"
	"encoding/xml"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

// fakeClient feeds a canned RPC reply to the collector under test.
type fakeClient struct {
	reply []byte
}

func (c *fakeClient) RunCommandAndParse(_ string, obj any) error {
	return xml.Unmarshal(c.reply, obj)
}

func (c *fakeClient) RunCommandAndParseWithParser(_ string, parser rpc.Parser) error {
	return parser(c.reply)
}

func (c *fakeClient) IsSatelliteEnabled() bool { return false }

func (c *fakeClient) IsScrapingLicenseEnabled() bool { return false }

func (c *fakeClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}

func (c *fakeClient) Context() context.Context { return context.Background() }

// A multi-RE reply carrying the wrapper element but no routing-engine items.
// Junos returns this shape when the RPC succeeds with nothing to report.
const emptyMultiREReply = `<rpc-reply>
  <multi-routing-engine-results>
  </multi-routing-engine-results>
</rpc-reply>`

func TestCollectStatsWithEmptyMultiREResult(t *testing.T) {
	c := &securityPolicyCollector{}
	cl := &fakeClient{reply: []byte(emptyMultiREReply)}
	ch := make(chan prometheus.Metric, 16)

	if err := c.CollectStats(cl, ch, []string{"target"}); err != nil {
		t.Fatalf("CollectStats returned error: %v", err)
	}

	if len(ch) != 0 {
		t.Errorf("CollectStats emitted %d metrics for an empty result, want 0", len(ch))
	}
}

func TestCollectHitsWithEmptyMultiREResult(t *testing.T) {
	c := &securityPolicyCollector{}
	cl := &fakeClient{reply: []byte(emptyMultiREReply)}
	ch := make(chan prometheus.Metric, 16)

	if err := c.CollectHits(cl, ch, []string{"target"}); err != nil {
		t.Fatalf("CollectHits returned error: %v", err)
	}

	if len(ch) != 0 {
		t.Errorf("CollectHits emitted %d metrics for an empty result, want 0", len(ch))
	}
}
