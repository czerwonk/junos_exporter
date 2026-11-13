// SPDX-License-Identifier: MIT

package security

import (
	"context"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

type fakeClient struct {
	reply string
}

func (c *fakeClient) RunCommandAndParse(_ string, _ any) error { return nil }

func (c *fakeClient) RunCommandAndParseWithParser(_ string, parser rpc.Parser) error {
	return parser([]byte(c.reply))
}

func (c *fakeClient) IsSatelliteEnabled() bool { return false }

func (c *fakeClient) IsScrapingLicenseEnabled() bool { return false }

func (c *fakeClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}

func (c *fakeClient) Context() context.Context { return context.Background() }

// TestCollectBranchSRXWithoutCPSessions covers branch SRX devices (SRX3xx),
// which report "N/A" for the central point session counters.
func TestCollectBranchSRXWithoutCPSessions(t *testing.T) {
	cl := &fakeClient{reply: `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S5.5/junos">
    <performance-summary-information>
        <performance-summary-statistics>
            <fpc-number>0</fpc-number>
            <pic-number>0</pic-number>
            <spu-cpu-utilization>2</spu-cpu-utilization>
            <spu-memory-utilization>61</spu-memory-utilization>
            <spu-current-flow-session>18</spu-current-flow-session>
            <spu-max-flow-session>65536</spu-max-flow-session>
            <spu-current-cp-session>N/A</spu-current-cp-session>
            <spu-max-cp-session>N/A</spu-max-cp-session>
        </performance-summary-statistics>
    </performance-summary-information>
</rpc-reply>`}

	c := &securityCollector{}
	ch := make(chan prometheus.Metric, 16)

	if err := c.Collect(cl, ch, []string{"target"}); err != nil {
		t.Fatalf("Collect returned error: %v", err)
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

	if v, ok := got[cpuUtilization.String()]; !ok || v != 2 {
		t.Errorf("cpu_utilization = %v (present %v), want 2", v, ok)
	}
	if v, ok := got[maxFlowSession.String()]; !ok || v != 65536 {
		t.Errorf("max_flow_session = %v (present %v), want 65536", v, ok)
	}
	if _, ok := got[currentCpSession.String()]; ok {
		t.Error("current_cp_session exported for an N/A value")
	}
	if _, ok := got[maxCpSession.String()]; ok {
		t.Error("max_cp_session exported for an N/A value")
	}
}

func TestCollectCPSessions(t *testing.T) {
	cl := &fakeClient{reply: `<rpc-reply>
    <performance-summary-information>
        <performance-summary-statistics>
            <fpc-number>1</fpc-number>
            <pic-number>0</pic-number>
            <spu-cpu-utilization>5</spu-cpu-utilization>
            <spu-memory-utilization>40</spu-memory-utilization>
            <spu-current-flow-session>100</spu-current-flow-session>
            <spu-max-flow-session>4194304</spu-max-flow-session>
            <spu-current-cp-session>120</spu-current-cp-session>
            <spu-max-cp-session>14680064</spu-max-cp-session>
        </performance-summary-statistics>
    </performance-summary-information>
</rpc-reply>`}

	c := &securityCollector{}
	ch := make(chan prometheus.Metric, 16)

	if err := c.Collect(cl, ch, []string{"target"}); err != nil {
		t.Fatalf("Collect returned error: %v", err)
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

	if v := got[currentCpSession.String()]; v != 120 {
		t.Errorf("current_cp_session = %v, want 120", v)
	}
	if v := got[maxCpSession.String()]; v != 14680064 {
		t.Errorf("max_cp_session = %v, want 14680064", v)
	}
}
