// SPDX-License-Identifier: MIT

package environment

import (
	"context"
	"encoding/xml"
	"errors"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

// fakeClient returns a different canned reply per command, so the base and
// satellite RPCs can be given different shapes. A non-nil err fails every RPC
// instead.
type fakeClient struct {
	replies   map[string]string
	satellite bool
	err       error
}

func (c *fakeClient) replyFor(cmd string) []byte {
	// Longest match wins so "show chassis environment satellite" is not served
	// by the "show chassis environment" entry.
	best := ""
	for k := range c.replies {
		if strings.Contains(cmd, k) && len(k) > len(best) {
			best = k
		}
	}

	return []byte(c.replies[best])
}

func (c *fakeClient) RunCommandAndParse(cmd string, obj any) error {
	if c.err != nil {
		return c.err
	}

	return xml.Unmarshal(c.replyFor(cmd), obj)
}

func (c *fakeClient) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	if c.err != nil {
		return c.err
	}

	return parser(c.replyFor(cmd))
}

func (c *fakeClient) IsSatelliteEnabled() bool { return c.satellite }

func (c *fakeClient) IsScrapingLicenseEnabled() bool { return false }

func (c *fakeClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}

func (c *fakeClient) Context() context.Context { return context.Background() }

// TestEnvironmentItemsSatelliteWithEmptyBaseResult covers the case where the
// base RPC returns a multi-RE wrapper with no items but the satellite RPC
// returns one. The merge indexes the base result without checking its length.
func TestEnvironmentItemsSatelliteWithEmptyBaseResult(t *testing.T) {
	cl := &fakeClient{
		satellite: true,
		replies: map[string]string{
			"show chassis environment": `<rpc-reply>
  <multi-routing-engine-results>
  </multi-routing-engine-results>
</rpc-reply>`,
			"show chassis environment satellite": `<rpc-reply>
  <multi-routing-engine-results>
    <multi-routing-engine-item>
      <re-name>satellite-1</re-name>
      <environment-information>
        <environment-item>
          <name>FPC 0 Sensor</name>
          <class>Temp</class>
          <status>OK</status>
        </environment-item>
      </environment-information>
    </multi-routing-engine-item>
  </multi-routing-engine-results>
</rpc-reply>`,
		},
	}

	c := &environmentCollector{}
	ch := make(chan prometheus.Metric, 16)

	if err := c.environmentItems(cl, ch, []string{"target"}); err != nil {
		t.Fatalf("environmentItems returned error: %v", err)
	}
}

// TestEnvironmentItemsPropagatesRPCError pins that a failed RPC surfaces as an
// error rather than an empty successful scrape.
func TestEnvironmentItemsPropagatesRPCError(t *testing.T) {
	want := errors.New("rpc failed")
	cl := &fakeClient{err: want}

	c := &environmentCollector{}
	ch := make(chan prometheus.Metric, 16)

	err := c.environmentItems(cl, ch, []string{"target"})
	if !errors.Is(err, want) {
		t.Errorf("environmentItems error = %v, want %v", err, want)
	}
}
