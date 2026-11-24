// SPDX-License-Identifier: MIT

package environment

import (
	"context"
	"encoding/xml"
	"errors"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

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

// TestEnvironmentPEMItemsStr3DcCurrent covers newer Junos releases (22.4R3 and
// later) that report the PEM current as <str3-dc-current> next to
// <str3-dc-voltage> instead of <dc-current>.
func TestEnvironmentPEMItemsStr3DcCurrent(t *testing.T) {
	cl := &fakeClient{
		replies: map[string]string{
			"show chassis environment pem": `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/22.4R0/junos">
  <environment-component-information xmlns="http://xml.juniper.net/junos/22.4R0/junos-chassis">
    <environment-component-item>
      <name>PEM 0</name>
      <state>Online</state>
      <dc-information>
        <dc-input>OK</dc-input>
        <dc-detail>
          <str3-dc-voltage>12.00</str3-dc-voltage>
          <str3-dc-current>9.00</str3-dc-current>
        </dc-detail>
      </dc-information>
    </environment-component-item>
  </environment-component-information>
</rpc-reply>`,
		},
	}

	c := &environmentCollector{}
	ch := make(chan prometheus.Metric, 16)

	if err := c.environmentPEMItems(cl, ch, []string{"target"}); err != nil {
		t.Fatalf("environmentPEMItems returned error: %v", err)
	}
	close(ch)

	got := map[string]float64{}
	for m := range ch {
		var pb dto.Metric
		if err := m.Write(&pb); err != nil {
			t.Fatal(err)
		}
		got[m.Desc().String()] = pb.GetGauge().GetValue()
	}

	if v := got[dcVoltageDesc.String()]; v != 12 {
		t.Errorf("pem_voltage = %v, want 12", v)
	}
	if v := got[dcCurrentDesc.String()]; v != 9 {
		t.Errorf("pem_current = %v, want 9", v)
	}
}
