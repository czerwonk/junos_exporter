// SPDX-License-Identifier: MIT

package evpnipprefix

import "testing"

const emptyResponse = `<rpc-reply>
  <evpn-ip-prefix-database-information>
  </evpn-ip-prefix-database-information>
</rpc-reply>`

// populated — one L3 context with all four tables present, IPv4 and IPv6,
// mix of Accepted and Rejected advertisements. RFC 5737 / RFC 3849
// documentation ranges throughout.
const populated = `<rpc-reply>
  <evpn-ip-prefix-database-information>
    <evpn-pfxdb-l3-context>
      <context-name>vrf-a</context-name>
      <evpn-pfxdb-ip-table>
        <table-description>IPv4-&gt;EVPN</table-description>
        <evpn-pfxdb-ip-entry>
          <entry-prefix>192.0.2.0/24</entry-prefix>
          <entry-evpn-route-status>Created</entry-evpn-route-status>
        </evpn-pfxdb-ip-entry>
        <evpn-pfxdb-ip-entry>
          <entry-prefix>198.51.100.0/24</entry-prefix>
          <entry-evpn-route-status>Created</entry-evpn-route-status>
        </evpn-pfxdb-ip-entry>
        <evpn-pfxdb-ip-entry>
          <entry-prefix>203.0.113.0/24</entry-prefix>
          <entry-evpn-route-status>Created</entry-evpn-route-status>
        </evpn-pfxdb-ip-entry>
      </evpn-pfxdb-ip-table>
      <evpn-pfxdb-ip-table>
        <table-description>IPv6-&gt;EVPN</table-description>
        <evpn-pfxdb-ip-entry>
          <entry-prefix>2001:db8::/32</entry-prefix>
          <entry-evpn-route-status>Created</entry-evpn-route-status>
        </evpn-pfxdb-ip-entry>
      </evpn-pfxdb-ip-table>
      <evpn-pfxdb-evpn-ip-table>
        <table-description>EVPN-&gt;IPv4</table-description>
        <evpn-pfxdb-evpn-ip-entry>
          <entry-prefix>192.0.2.0/24</entry-prefix>
          <entry-etag>0</entry-etag>
          <evpn-pfxdb-evpn-ip-adv>
            <route-distinguisher>65000:1</route-distinguisher>
            <adv-vni>10001</adv-vni>
            <adv-router-mac>00:00:5e:00:53:01</adv-router-mac>
            <adv-bgp-nexthop>192.0.2.1</adv-bgp-nexthop>
            <adv-ip-route-status>Accepted</adv-ip-route-status>
            <adv-ip-route-error>n/a</adv-ip-route-error>
          </evpn-pfxdb-evpn-ip-adv>
          <evpn-pfxdb-evpn-ip-adv>
            <route-distinguisher>65000:2</route-distinguisher>
            <adv-vni>10001</adv-vni>
            <adv-router-mac>00:00:5e:00:53:02</adv-router-mac>
            <adv-bgp-nexthop>192.0.2.2</adv-bgp-nexthop>
            <adv-ip-route-status>Accepted</adv-ip-route-status>
            <adv-ip-route-error>n/a</adv-ip-route-error>
          </evpn-pfxdb-evpn-ip-adv>
        </evpn-pfxdb-evpn-ip-entry>
        <evpn-pfxdb-evpn-ip-entry>
          <entry-prefix>203.0.113.0/24</entry-prefix>
          <entry-etag>0</entry-etag>
          <evpn-pfxdb-evpn-ip-adv>
            <route-distinguisher>65000:3</route-distinguisher>
            <adv-vni>10001</adv-vni>
            <adv-router-mac>00:00:5e:00:53:03</adv-router-mac>
            <adv-bgp-nexthop>192.0.2.3</adv-bgp-nexthop>
            <adv-ip-route-status>Rejected</adv-ip-route-status>
            <adv-ip-route-error>policy</adv-ip-route-error>
          </evpn-pfxdb-evpn-ip-adv>
        </evpn-pfxdb-evpn-ip-entry>
      </evpn-pfxdb-evpn-ip-table>
      <evpn-pfxdb-evpn-ip-table>
        <table-description>EVPN-&gt;IPv6</table-description>
      </evpn-pfxdb-evpn-ip-table>
    </evpn-pfxdb-l3-context>
  </evpn-ip-prefix-database-information>
</rpc-reply>`

const multiEngineSample = `<rpc-reply>
  <multi-routing-engine-results>
    <multi-routing-engine-item>
      <re-name>fpc0</re-name>
      <evpn-ip-prefix-database-information>
        <evpn-pfxdb-l3-context>
          <context-name>vrf-a</context-name>
        </evpn-pfxdb-l3-context>
      </evpn-ip-prefix-database-information>
    </multi-routing-engine-item>
    <multi-routing-engine-item>
      <re-name>fpc1</re-name>
      <evpn-ip-prefix-database-information>
        <evpn-pfxdb-l3-context>
          <context-name>vrf-b</context-name>
        </evpn-pfxdb-l3-context>
      </evpn-ip-prefix-database-information>
    </multi-routing-engine-item>
  </multi-routing-engine-results>
</rpc-reply>`

func TestParseEmpty(t *testing.T) {
	c, err := parseContexts([]byte(emptyResponse))
	if err != nil {
		t.Fatalf("parseContexts: %v", err)
	}
	if len(c) != 0 {
		t.Errorf("expected 0 contexts, got %d", len(c))
	}
}

func TestParsePopulated(t *testing.T) {
	c, err := parseContexts([]byte(populated))
	if err != nil {
		t.Fatalf("parseContexts: %v", err)
	}
	if len(c) != 1 {
		t.Fatalf("expected 1 context, got %d", len(c))
	}
	ctx := c[0]
	if ctx.Name != "vrf-a" {
		t.Errorf("context name: got %q", ctx.Name)
	}
	if len(ctx.LocalTables) != 2 {
		t.Errorf("local-tables: got %d, want 2", len(ctx.LocalTables))
	}
	if len(ctx.LocalTables[0].Entries) != 3 {
		t.Errorf("IPv4 local entries: got %d, want 3", len(ctx.LocalTables[0].Entries))
	}
	if len(ctx.LocalTables[1].Entries) != 1 {
		t.Errorf("IPv6 local entries: got %d, want 1", len(ctx.LocalTables[1].Entries))
	}
	if len(ctx.RemoteTables) != 2 {
		t.Errorf("remote-tables: got %d, want 2", len(ctx.RemoteTables))
	}
	if len(ctx.RemoteTables[0].Entries) != 2 {
		t.Errorf("IPv4 remote entries: got %d, want 2", len(ctx.RemoteTables[0].Entries))
	}
	if len(ctx.RemoteTables[1].Entries) != 0 {
		t.Errorf("IPv6 remote entries: got %d, want 0", len(ctx.RemoteTables[1].Entries))
	}

	// Walk the per-PE advertisements to validate status counts.
	accepted, rejected := 0, 0
	for _, e := range ctx.RemoteTables[0].Entries {
		for _, adv := range e.Advertisements {
			switch adv.Status {
			case "Accepted":
				accepted++
			case "Rejected":
				rejected++
			}
		}
	}
	if accepted != 2 || rejected != 1 {
		t.Errorf("advertisement statuses: got %d accepted / %d rejected, want 2/1", accepted, rejected)
	}
}

func TestParseMultiEngine(t *testing.T) {
	c, err := parseContexts([]byte(multiEngineSample))
	if err != nil {
		t.Fatalf("parseContexts: %v", err)
	}
	if len(c) != 2 {
		t.Fatalf("expected 2 merged contexts, got %d", len(c))
	}
	names := map[string]bool{"vrf-a": false, "vrf-b": false}
	for _, ctx := range c {
		if _, ok := names[ctx.Name]; ok {
			names[ctx.Name] = true
		} else {
			t.Errorf("unexpected context name %q", ctx.Name)
		}
	}
	for n, seen := range names {
		if !seen {
			t.Errorf("expected to see context %q", n)
		}
	}
}

func TestAfiFromDescription(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"IPv4->EVPN", "v4"},
		{"IPv6->EVPN", "v6"},
		{"EVPN->IPv4", "v4"},
		{"EVPN->IPv6", "v6"},
		{"", ""},
	}
	for _, c := range cases {
		got := afiFromDescription(c.in)
		if got != c.want {
			t.Errorf("afiFromDescription(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}
