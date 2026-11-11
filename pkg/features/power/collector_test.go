// SPDX-License-Identifier: MIT

package power

import (
	"context"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

type fakeClient struct {
	replies map[string]string
	called  []string
}

func (c *fakeClient) RunCommandAndParse(cmd string, obj any) error {
	c.called = append(c.called, cmd)
	body, ok := c.replies[cmd]
	if !ok {
		return fmt.Errorf("unexpected command %q", cmd)
	}
	return xml.Unmarshal([]byte(body), obj)
}

func (c *fakeClient) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	c.called = append(c.called, cmd)
	body, ok := c.replies[cmd]
	if !ok {
		return fmt.Errorf("unexpected command %q", cmd)
	}
	return parser([]byte(body))
}

func (c *fakeClient) IsSatelliteEnabled() bool       { return false }
func (c *fakeClient) IsScrapingLicenseEnabled() bool { return false }
func (c *fakeClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}
func (c *fakeClient) Context() context.Context { return context.Background() }

const chassisHardwareEX4000 = `<rpc-reply>
  <chassis-inventory>
    <chassis>
      <description>EX4000-24MP</description>
    </chassis>
  </chassis-inventory>
</rpc-reply>`

const chassisHardwareEX4300 = `<rpc-reply>
  <chassis-inventory>
    <chassis>
      <description>EX4300-48P</description>
    </chassis>
  </chassis-inventory>
</rpc-reply>`

const chassisHardwareMX = `<rpc-reply>
  <chassis-inventory>
    <chassis>
      <description>MX960</description>
    </chassis>
  </chassis-inventory>
</rpc-reply>`

// commandErrorReply is the shape EX4000 returns for "show chassis power",
// which is not a valid command on that platform (the CLI autocompletes it
// to "show chassis power-budget-statistics").
const commandErrorReply = `<command>
  show chassis power
</command>`

const powerBudgetReply = `<rpc-reply>
  <power-budget-information>
    <line-card-slot>0</line-card-slot>
    <psu-slot>0</psu-slot>
    <psu-type>JPSU-580W-AC</psu-type>
    <power-supplied-psu>580</power-supplied-psu>
    <power-supply-state>Online</power-supply-state>
    <psu-redundancy-config>N+0</psu-redundancy-config>
    <total-power-supplied>580</total-power-supplied>
    <base-power>60</base-power>
    <actual-power-used>2</actual-power-used>
  </power-budget-information>
</rpc-reply>`

const chassisPowerReply = `<rpc-reply>
  <power-usage-information>
    <power-usage-item>
      <name>PEM 0</name>
      <state>Online</state>
      <dc-output-detail>
        <dc-power>120</dc-power>
        <zone>0</zone>
        <dc-current>10</dc-current>
        <dc-voltage>12</dc-voltage>
        <dc-load>4</dc-load>
      </dc-output-detail>
    </power-usage-item>
    <power-usage-system>
      <capacity-sys-actual>120</capacity-sys-actual>
      <capacity-sys-max>240</capacity-sys-max>
      <capacity-sys-remaining>120</capacity-sys-remaining>
    </power-usage-system>
  </power-usage-information>
</rpc-reply>`

func TestUsesPowerBudgetStatistics(t *testing.T) {
	cases := []struct {
		platform string
		want     bool
	}{
		{"EX4000-24MP", true},
		{"EX4000-48MP", true},
		{"EX4300-48P", true},
		{"EX4300", true},
		{"MX960", false},
		{"EX4400-24T", false},
		{"SRX345", false},
		{"", false},
	}
	for _, tc := range cases {
		if got := usesPowerBudgetStatistics(tc.platform); got != tc.want {
			t.Errorf("usesPowerBudgetStatistics(%q) = %v, want %v", tc.platform, got, tc.want)
		}
	}
}

func TestCollectUsesPowerBudgetOnEX4000(t *testing.T) {
	cl := &fakeClient{
		replies: map[string]string{
			"show chassis hardware":                chassisHardwareEX4000,
			"show chassis power-budget-statistics": powerBudgetReply,
			// If the collector still issues this, the <command> root
			// would reproduce the original scrape error.
			"show chassis power": commandErrorReply,
		},
	}

	ch := make(chan prometheus.Metric, 16)
	err := NewCollector().Collect(cl, ch, []string{"switch1"})
	if err != nil {
		t.Fatalf("Collect returned error: %v", err)
	}

	for _, cmd := range cl.called {
		if cmd == "show chassis power" {
			t.Fatalf("Collect issued %q on EX4000; want power-budget-statistics only", cmd)
		}
	}
	if got := len(ch); got == 0 {
		t.Fatal("Collect emitted no budget metrics")
	}
}

func TestCollectUsesPowerBudgetOnEX4300(t *testing.T) {
	cl := &fakeClient{
		replies: map[string]string{
			"show chassis hardware":                chassisHardwareEX4300,
			"show chassis power-budget-statistics": powerBudgetReply,
			"show chassis power":                   commandErrorReply,
		},
	}

	ch := make(chan prometheus.Metric, 16)
	if err := NewCollector().Collect(cl, ch, []string{"switch1"}); err != nil {
		t.Fatalf("Collect returned error: %v", err)
	}
	for _, cmd := range cl.called {
		if cmd == "show chassis power" {
			t.Fatalf("Collect issued %q on EX4300", cmd)
		}
	}
}

func TestCollectUsesChassisPowerOnMX(t *testing.T) {
	cl := &fakeClient{
		replies: map[string]string{
			"show chassis hardware": chassisHardwareMX,
			"show chassis power":    chassisPowerReply,
		},
	}

	ch := make(chan prometheus.Metric, 16)
	if err := NewCollector().Collect(cl, ch, []string{"router1"}); err != nil {
		t.Fatalf("Collect returned error: %v", err)
	}
	for _, cmd := range cl.called {
		if cmd == "show chassis power-budget-statistics" {
			t.Fatalf("Collect issued %q on MX960", cmd)
		}
	}
	if got := len(ch); got == 0 {
		t.Fatal("Collect emitted no chassis power metrics")
	}
}

func TestCollectEX4000WouldFailOnShowChassisPower(t *testing.T) {
	// Pin the original failure mode: unmarshalling the EX4000
	// "show chassis power" reply into an rpc-reply struct errors.
	var res singleRoutingEngineResult
	err := xml.Unmarshal([]byte(commandErrorReply), &res)
	if err == nil {
		t.Fatal("expected unmarshal of <command> reply to fail")
	}
}
