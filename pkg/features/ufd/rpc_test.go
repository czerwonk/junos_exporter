// SPDX-License-Identifier: MIT

package ufd

import (
	"encoding/xml"
	"testing"
)

// Synthetic healthy-state sample modelled on real EX-series output:
// single group, one uplink (currently marked active with "*"),
// two downlinks (one marked, one not), failure-action Inactive.
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

func TestParseUFDHealthy(t *testing.T) {
	var r result
	if err := xml.Unmarshal([]byte(healthySample), &r); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(r.Information.GroupInfos) != 1 {
		t.Fatalf("expected 1 group-info, got %d", len(r.Information.GroupInfos))
	}

	g := r.Information.GroupInfos[0].Group
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

func TestSplitMark(t *testing.T) {
	cases := []struct {
		in       string
		wantName string
		wantMark float64
	}{
		{"et-0/0/55*", "et-0/0/55", 1.0},
		{"xe-0/0/47", "xe-0/0/47", 0.0},
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
