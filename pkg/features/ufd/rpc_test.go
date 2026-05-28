// SPDX-License-Identifier: MIT

package ufd

import (
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
// no debounce-interval element, second group in failure-action Active
// (uplinks lost their "*" marker).
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

// Synthetic multi-RE / Virtual Chassis sample. The QFX UFD YANG declares
// the response can be wrapped in <multi-routing-engine-results> with a
// <multi-routing-engine-item> per RE / VC member. Each RE reports its own
// UFD groups; we merge them.
const multiEngineSample = `<rpc-reply>
  <multi-routing-engine-results>
    <multi-routing-engine-item>
      <re-name>fpc0</re-name>
      <uplink-failure-detection-information>
        <ufd-group-information>
          <ufd-group>
            <ufd-group-name>vc-group-1</ufd-group-name>
            <link-to-monitor-list>
              <uplink-interface>et-0/0/55*</uplink-interface>
            </link-to-monitor-list>
            <link-to-disable-list>
              <downlink-interface>xe-0/0/46*</downlink-interface>
            </link-to-disable-list>
            <failure-action>Inactive</failure-action>
          </ufd-group>
        </ufd-group-information>
      </uplink-failure-detection-information>
    </multi-routing-engine-item>
    <multi-routing-engine-item>
      <re-name>fpc1</re-name>
      <uplink-failure-detection-information>
        <ufd-group-information>
          <ufd-group>
            <ufd-group-name>vc-group-2</ufd-group-name>
            <link-to-monitor-list>
              <uplink-interface>et-1/0/55*</uplink-interface>
            </link-to-monitor-list>
            <link-to-disable-list>
              <downlink-interface>xe-1/0/46*</downlink-interface>
            </link-to-disable-list>
            <failure-action>Inactive</failure-action>
          </ufd-group>
        </ufd-group-information>
      </uplink-failure-detection-information>
    </multi-routing-engine-item>
  </multi-routing-engine-results>
</rpc-reply>`

func TestParseUFDHealthy(t *testing.T) {
	groups, err := parseGroups([]byte(healthySample))
	if err != nil {
		t.Fatalf("parseGroups failed: %v", err)
	}

	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}

	g := groups[0]
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
	groups, err := parseGroups([]byte(multiGroupSample))
	if err != nil {
		t.Fatalf("parseGroups failed: %v", err)
	}

	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}

	if groups[0].Name != "group-a" {
		t.Errorf("group 0 name: got %q", groups[0].Name)
	}
	if groups[0].FailureAction != "Inactive" {
		t.Errorf("group 0 failure-action: got %q", groups[0].FailureAction)
	}
	if groups[0].DebounceInterval != 0 {
		t.Errorf("group 0 debounce default: got %v want 0", groups[0].DebounceInterval)
	}

	if groups[1].Name != "group-b" {
		t.Errorf("group 1 name: got %q", groups[1].Name)
	}
	if groups[1].FailureAction != "Active" {
		t.Errorf("group 1 failure-action: got %q want Active", groups[1].FailureAction)
	}
	if len(groups[1].Uplinks) != 2 {
		t.Errorf("group 1 uplinks: got %#v", groups[1].Uplinks)
	}
}

func TestParseUFDMultiEngine(t *testing.T) {
	groups, err := parseGroups([]byte(multiEngineSample))
	if err != nil {
		t.Fatalf("parseGroups failed: %v", err)
	}

	if len(groups) != 2 {
		t.Fatalf("expected 2 merged groups (one per RE), got %d", len(groups))
	}

	names := []string{groups[0].Name, groups[1].Name}
	wantNames := map[string]bool{"vc-group-1": false, "vc-group-2": false}
	for _, n := range names {
		if _, ok := wantNames[n]; !ok {
			t.Errorf("unexpected group name %q", n)
		}
		wantNames[n] = true
	}
	for n, seen := range wantNames {
		if !seen {
			t.Errorf("expected to see group %q but didn't", n)
		}
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
