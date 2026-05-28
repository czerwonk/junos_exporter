// SPDX-License-Identifier: MIT

package ufd

import (
	"encoding/xml"
	"testing"
)

// Synthetic healthy-state sample modelled on real EX-series output:
// single group, one uplink (currently marked active with "*"),
// two downlinks (one marked, one not), failure-action Inactive,
// explicit debounce-interval.
const healthySample = `<rpc-reply>
  <uplink-failure-detection-information>
    <ufd-group-information>
      <ufd-group>
        <ufd-group-name>uplink-group-1</ufd-group-name>
        <link-to-monitor-list>
          <uplink-interface>et-0/0/55*</uplink-interface>
        </link-to-monitor-list>
        <link-to-disable-list>
          <downlink-interface>xe-0/0/46*</downlink-interface>
          <downlink-interface>xe-0/0/47</downlink-interface>
        </link-to-disable-list>
        <failure-action>Inactive</failure-action>
        <debounce-interval>60</debounce-interval>
      </ufd-group>
    </ufd-group-information>
  </uplink-failure-detection-information>
</rpc-reply>`

// Synthetic multi-group sample modelled on real older-Junos output:
// two <ufd-group> siblings inside ONE <ufd-group-information> wrapper,
// no debounce-interval element (older Junos / not configured),
// second group in failure-action Active (uplinks lost their "*" marker).
const multiGroupSample = `<rpc-reply>
  <uplink-failure-detection-information>
    <ufd-group-information>
      <ufd-group>
        <ufd-group-name>group-a</ufd-group-name>
        <link-to-monitor-list>
          <uplink-interface>ae22*</uplink-interface>
        </link-to-monitor-list>
        <link-to-disable-list>
          <downlink-interface>xe-0/0/20*</downlink-interface>
          <downlink-interface>xe-0/0/21</downlink-interface>
        </link-to-disable-list>
        <failure-action>Inactive</failure-action>
      </ufd-group>
      <ufd-group>
        <ufd-group-name>group-b</ufd-group-name>
        <link-to-monitor-list>
          <uplink-interface>et-0/0/25</uplink-interface>
          <uplink-interface>et-0/0/26</uplink-interface>
        </link-to-monitor-list>
        <link-to-disable-list>
          <downlink-interface>et-0/0/27</downlink-interface>
        </link-to-disable-list>
        <failure-action>Active</failure-action>
      </ufd-group>
    </ufd-group-information>
  </uplink-failure-detection-information>
</rpc-reply>`

func TestParseUFDHealthy(t *testing.T) {
	var r result
	if err := xml.Unmarshal([]byte(healthySample), &r); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(r.Groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(r.Groups))
	}

	g := r.Groups[0]
	if g.Name != "uplink-group-1" {
		t.Errorf("group name: got %q", g.Name)
	}
	if g.FailureAction != "Inactive" {
		t.Errorf("failure-action: got %q", g.FailureAction)
	}
	if g.DebounceInterval != 60 {
		t.Errorf("debounce-interval: got %v", g.DebounceInterval)
	}
	if len(g.Uplinks) != 1 || g.Uplinks[0] != "et-0/0/55*" {
		t.Errorf("uplinks: got %#v", g.Uplinks)
	}
	if len(g.Downlinks) != 2 || g.Downlinks[0] != "xe-0/0/46*" || g.Downlinks[1] != "xe-0/0/47" {
		t.Errorf("downlinks: got %#v", g.Downlinks)
	}
}

func TestParseUFDMultiGroup(t *testing.T) {
	var r result
	if err := xml.Unmarshal([]byte(multiGroupSample), &r); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(r.Groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(r.Groups))
	}

	// First group: healthy, single uplink, two downlinks
	if r.Groups[0].Name != "group-a" {
		t.Errorf("group 0 name: got %q", r.Groups[0].Name)
	}
	if r.Groups[0].FailureAction != "Inactive" {
		t.Errorf("group 0 failure-action: got %q", r.Groups[0].FailureAction)
	}
	if r.Groups[0].DebounceInterval != 0 {
		t.Errorf("group 0 debounce default: got %v want 0", r.Groups[0].DebounceInterval)
	}
	if len(r.Groups[0].Uplinks) != 1 {
		t.Errorf("group 0 uplinks: got %#v", r.Groups[0].Uplinks)
	}

	// Second group: triggered failure, multiple uplinks (no "*"), single downlink
	if r.Groups[1].Name != "group-b" {
		t.Errorf("group 1 name: got %q", r.Groups[1].Name)
	}
	if r.Groups[1].FailureAction != "Active" {
		t.Errorf("group 1 failure-action: got %q want Active", r.Groups[1].FailureAction)
	}
	if len(r.Groups[1].Uplinks) != 2 {
		t.Errorf("group 1 uplinks: got %#v", r.Groups[1].Uplinks)
	}
	if len(r.Groups[1].Downlinks) != 1 {
		t.Errorf("group 1 downlinks: got %#v", r.Groups[1].Downlinks)
	}
}

func TestSplitMark(t *testing.T) {
	cases := []struct {
		in       string
		wantName string
		wantMark float64
	}{
		{"et-0/0/55*", "et-0/0/55", 1.0},
		{"xe-0/0/47", "xe-0/0/47", 0.0},
		{"ae22*", "ae22", 1.0},
		{"  et-0/0/1*  ", "et-0/0/1", 1.0},
		{"", "", 0.0},
	}
	for _, c := range cases {
		name, mark := splitMark(c.in)
		if name != c.wantName || mark != c.wantMark {
			t.Errorf("splitMark(%q) = (%q, %v); want (%q, %v)", c.in, name, mark, c.wantName, c.wantMark)
		}
	}
}
