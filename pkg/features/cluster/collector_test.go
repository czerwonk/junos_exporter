// SPDX-License-Identifier: MIT

package cluster

import (
	"context"
	"errors"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

// fakeClient fails every RPC with a fixed error.
type fakeClient struct {
	err error
}

func (c *fakeClient) RunCommandAndParse(_ string, _ any) error { return c.err }

func (c *fakeClient) RunCommandAndParseWithParser(_ string, _ rpc.Parser) error { return c.err }

func (c *fakeClient) IsSatelliteEnabled() bool { return false }

func (c *fakeClient) IsScrapingLicenseEnabled() bool { return false }

func (c *fakeClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}

func (c *fakeClient) Context() context.Context { return context.Background() }

// TestCollectPropagatesRPCError pins that a failed RPC surfaces as an error
// rather than an empty successful scrape.
func TestCollectPropagatesRPCError(t *testing.T) {
	want := errors.New("rpc failed")
	cl := &fakeClient{err: want}

	c := NewCollector()
	ch := make(chan prometheus.Metric, 16)

	err := c.Collect(cl, ch, []string{"target"})
	if !errors.Is(err, want) {
		t.Errorf("Collect error = %v, want %v", err, want)
	}
}
