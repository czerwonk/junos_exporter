// SPDX-License-Identifier: MIT

package evpn

import "testing"

// healthySingleEVI - synthetic shape modelled on real EX-class extensive
// output but populated entirely with RFC 5737 documentation IPs and
// synthetic identifiers. Single EVI with one neighbor and a small MAC
// database.
const healthySingleEVI = `<rpc-reply>
  <evpn-instance-information>
    <evpn-instance>
      <evpn-instance-name>EVI-A</evpn-instance-name>
      <route-distinguisher>65000:1</route-distinguisher>
      <evpn-encap-type>VXLAN</evpn-encap-type>
      <evpn-instance-control-word>enabled</evpn-instance-control-word>
      <duplicate-mac-detection-threshold>5</duplicate-mac-detection-threshold>
      <duplicate-mac-detection-window>180</duplicate-mac-detection-window>
      <mac-database-status-table>
        <local-macs>10</local-macs>
        <remote-macs>4</remote-macs>
        <local-mac-ips>10</local-mac-ips>
        <remote-mac-ips>4</remote-mac-ips>
        <local-default-gateway-macs>1</local-default-gateway-macs>
        <remote-default-gateway-macs>0</remote-default-gateway-macs>
      </mac-database-status-table>
      <local-interfaces>3</local-interfaces>
      <local-interfaces-up>3</local-interfaces-up>
      <irb-interfaces>1</irb-interfaces>
      <irb-interfaces-up>1</irb-interfaces-up>
      <num-protect-interfaces>0</num-protect-interfaces>
      <num-bridge-domains>1</num-bridge-domains>
      <evpn-num-neighbors>1</evpn-num-neighbors>
      <evpn-neighbor>
        <evpn-neighbor-route-information>
          <evpn-neighbor-address>192.0.2.1</evpn-neighbor-address>
          <evpn-num-mac-routes>4</evpn-num-mac-routes>
          <evpn-num-mac-ip-routes>4</evpn-num-mac-ip-routes>
          <evpn-num-ethernet-autodiscovery-routes>2</evpn-num-ethernet-autodiscovery-routes>
          <evpn-num-inclusive-multicast-routes>1</evpn-num-inclusive-multicast-routes>
          <evpn-num-ethernet-segment-routes>0</evpn-num-ethernet-segment-routes>
        </evpn-neighbor-route-information>
      </evpn-neighbor>
      <evpn-num-esi>2</evpn-num-esi>
      <evpn-router-id>192.0.2.10</evpn-router-id>
      <evpn-source-vtep-ipaddr>192.0.2.10</evpn-source-vtep-ipaddr>
      <evpn-smet-forwarding>Disabled</evpn-smet-forwarding>
    </evpn-instance>
  </evpn-instance-information>
</rpc-reply>`

// multiEVIMixed - synthetic two-EVI sample. The first EVI mimics
// __default_evpn__ (no encap, no mac-database, no per-PE breakdown -
// just the aggregate neighbour count). The second is a fully populated
// VXLAN EVI with two neighbors.
const multiEVIMixed = `<rpc-reply>
  <evpn-instance-information>
    <evpn-instance>
      <evpn-instance-name>__default_evpn__</evpn-instance-name>
      <route-distinguisher>65000:0</route-distinguisher>
      <num-bridge-domains>0</num-bridge-domains>
      <evpn-num-neighbors>1</evpn-num-neighbors>
      <evpn-neighbor>
        <evpn-neighbor-route-information>
          <evpn-neighbor-address>192.0.2.1</evpn-neighbor-address>
          <evpn-num-mac-routes>0</evpn-num-mac-routes>
          <evpn-num-mac-ip-routes>0</evpn-num-mac-ip-routes>
          <evpn-num-ethernet-autodiscovery-routes>0</evpn-num-ethernet-autodiscovery-routes>
          <evpn-num-inclusive-multicast-routes>0</evpn-num-inclusive-multicast-routes>
          <evpn-num-ethernet-segment-routes>2</evpn-num-ethernet-segment-routes>
        </evpn-neighbor-route-information>
      </evpn-neighbor>
    </evpn-instance>
    <evpn-instance>
      <evpn-instance-name>EVI-B</evpn-instance-name>
      <route-distinguisher>65000:2</route-distinguisher>
      <evpn-encap-type>VXLAN</evpn-encap-type>
      <evpn-instance-control-word>enabled</evpn-instance-control-word>
      <mac-database-status-table>
        <local-macs>50</local-macs>
        <remote-macs>20</remote-macs>
        <local-mac-ips>48</local-mac-ips>
        <remote-mac-ips>20</remote-mac-ips>
        <local-default-gateway-macs>0</local-default-gateway-macs>
        <remote-default-gateway-macs>0</remote-default-gateway-macs>
      </mac-database-status-table>
      <local-interfaces>6</local-interfaces>
      <local-interfaces-up>5</local-interfaces-up>
      <irb-interfaces>0</irb-interfaces>
      <irb-interfaces-up>0</irb-interfaces-up>
      <num-protect-interfaces>0</num-protect-interfaces>
      <num-bridge-domains>1</num-bridge-domains>
      <evpn-num-neighbors>2</evpn-num-neighbors>
      <evpn-neighbor>
        <evpn-neighbor-route-information>
          <evpn-neighbor-address>192.0.2.1</evpn-neighbor-address>
          <evpn-num-mac-routes>15</evpn-num-mac-routes>
          <evpn-num-mac-ip-routes>15</evpn-num-mac-ip-routes>
          <evpn-num-ethernet-autodiscovery-routes>4</evpn-num-ethernet-autodiscovery-routes>
          <evpn-num-inclusive-multicast-routes>1</evpn-num-inclusive-multicast-routes>
          <evpn-num-ethernet-segment-routes>2</evpn-num-ethernet-segment-routes>
        </evpn-neighbor-route-information>
      </evpn-neighbor>
      <evpn-neighbor>
        <evpn-neighbor-route-information>
          <evpn-neighbor-address>192.0.2.2</evpn-neighbor-address>
          <evpn-num-mac-routes>5</evpn-num-mac-routes>
          <evpn-num-mac-ip-routes>5</evpn-num-mac-ip-routes>
          <evpn-num-ethernet-autodiscovery-routes>0</evpn-num-ethernet-autodiscovery-routes>
          <evpn-num-inclusive-multicast-routes>1</evpn-num-inclusive-multicast-routes>
          <evpn-num-ethernet-segment-routes>0</evpn-num-ethernet-segment-routes>
        </evpn-neighbor-route-information>
      </evpn-neighbor>
      <evpn-num-esi>4</evpn-num-esi>
      <evpn-router-id>192.0.2.10</evpn-router-id>
      <evpn-source-vtep-ipaddr>192.0.2.10</evpn-source-vtep-ipaddr>
      <evpn-smet-forwarding>Disabled</evpn-smet-forwarding>
    </evpn-instance>
  </evpn-instance-information>
</rpc-reply>`

// multiEngineSample - synthetic multi-RE / Virtual Chassis wrap. Each
// RE reports its own EVPN instance; the dispatch logic must flatten
// them into a single list.
const multiEngineSample = `<rpc-reply>
  <multi-routing-engine-results>
    <multi-routing-engine-item>
      <re-name>fpc0</re-name>
      <evpn-instance-information>
        <evpn-instance>
          <evpn-instance-name>EVI-A</evpn-instance-name>
          <route-distinguisher>65000:1</route-distinguisher>
          <evpn-encap-type>VXLAN</evpn-encap-type>
          <evpn-num-neighbors>1</evpn-num-neighbors>
          <evpn-num-esi>0</evpn-num-esi>
        </evpn-instance>
      </evpn-instance-information>
    </multi-routing-engine-item>
    <multi-routing-engine-item>
      <re-name>fpc1</re-name>
      <evpn-instance-information>
        <evpn-instance>
          <evpn-instance-name>EVI-B</evpn-instance-name>
          <route-distinguisher>65000:2</route-distinguisher>
          <evpn-encap-type>VXLAN</evpn-encap-type>
          <evpn-num-neighbors>1</evpn-num-neighbors>
          <evpn-num-esi>0</evpn-num-esi>
        </evpn-instance>
      </evpn-instance-information>
    </multi-routing-engine-item>
  </multi-routing-engine-results>
</rpc-reply>`

func TestParseSingleEVI(t *testing.T) {
	inst, err := parseInstances([]byte(healthySingleEVI))
	if err != nil {
		t.Fatalf("parseInstances: %v", err)
	}
	if len(inst) != 1 {
		t.Fatalf("expected 1 instance, got %d", len(inst))
	}
	g := inst[0]
	if g.Name != "EVI-A" {
		t.Errorf("name: got %q", g.Name)
	}
	if g.RouteDistinguisher != "65000:1" {
		t.Errorf("rd: got %q", g.RouteDistinguisher)
	}
	if g.EncapType != "VXLAN" {
		t.Errorf("encap: got %q", g.EncapType)
	}
	if g.NumNeighbors != 1 {
		t.Errorf("num-neighbors: got %v", g.NumNeighbors)
	}
	if g.NumESI != 2 {
		t.Errorf("num-esi: got %v", g.NumESI)
	}
	if g.LocalInterfaces != 3 || g.LocalInterfacesUp != 3 {
		t.Errorf("local-interfaces: got %v/%v", g.LocalInterfaces, g.LocalInterfacesUp)
	}
	if g.MACDatabase == nil {
		t.Fatal("expected mac-database-status-table to parse")
	}
	if g.MACDatabase.LocalMACs != 10 || g.MACDatabase.RemoteMACs != 4 {
		t.Errorf("mac counts: got %v/%v", g.MACDatabase.LocalMACs, g.MACDatabase.RemoteMACs)
	}
	if g.MACDatabase.LocalMACIPs != 10 || g.MACDatabase.RemoteMACIPs != 4 {
		t.Errorf("mac-ip counts: got %v/%v", g.MACDatabase.LocalMACIPs, g.MACDatabase.RemoteMACIPs)
	}
	if len(g.Neighbors) != 1 {
		t.Fatalf("expected 1 neighbor, got %d", len(g.Neighbors))
	}
	r := g.Neighbors[0].Routes
	if r.Address != "192.0.2.1" {
		t.Errorf("neighbor addr: got %q", r.Address)
	}
	if r.NumMACRoutes != 4 || r.NumMACIPRoutes != 4 {
		t.Errorf("type-2 counts: got %v/%v", r.NumMACRoutes, r.NumMACIPRoutes)
	}
	if r.NumAutoDiscoveryRoutes != 2 {
		t.Errorf("type-1 count: got %v", r.NumAutoDiscoveryRoutes)
	}
	if r.NumInclusiveMulticastRoutes != 1 {
		t.Errorf("type-3 count: got %v", r.NumInclusiveMulticastRoutes)
	}
	if g.DuplicateMACDetectionThreshold != 5 {
		t.Errorf("dup-mac threshold: got %v", g.DuplicateMACDetectionThreshold)
	}
	if g.DuplicateMACDetectionWindow != 180 {
		t.Errorf("dup-mac window: got %v", g.DuplicateMACDetectionWindow)
	}
}

func TestParseMultiEVIMixed(t *testing.T) {
	inst, err := parseInstances([]byte(multiEVIMixed))
	if err != nil {
		t.Fatalf("parseInstances: %v", err)
	}
	if len(inst) != 2 {
		t.Fatalf("expected 2 instances, got %d", len(inst))
	}

	// Default-evpn-style instance: no mac-database, no encap.
	def := inst[0]
	if def.Name != "__default_evpn__" {
		t.Errorf("default name: got %q", def.Name)
	}
	if def.EncapType != "" {
		t.Errorf("default encap expected empty, got %q", def.EncapType)
	}
	if def.MACDatabase != nil {
		t.Errorf("default mac-database expected nil, got %#v", def.MACDatabase)
	}
	if def.NumNeighbors != 1 {
		t.Errorf("default num-neighbors: got %v", def.NumNeighbors)
	}
	if len(def.Neighbors) != 1 || def.Neighbors[0].Routes.NumEthernetSegmentRoutes != 2 {
		t.Errorf("default neighbor ES routes: got %#v", def.Neighbors)
	}

	// Fully populated VXLAN EVI with two PEs.
	full := inst[1]
	if full.Name != "EVI-B" {
		t.Errorf("EVI-B name: got %q", full.Name)
	}
	if len(full.Neighbors) != 2 {
		t.Fatalf("EVI-B neighbors: got %d", len(full.Neighbors))
	}
	if full.NumESI != 4 {
		t.Errorf("EVI-B num-esi: got %v", full.NumESI)
	}
	if full.MACDatabase == nil || full.MACDatabase.LocalMACs != 50 {
		t.Errorf("EVI-B local-macs: got %#v", full.MACDatabase)
	}
}

func TestParseMultiEngine(t *testing.T) {
	inst, err := parseInstances([]byte(multiEngineSample))
	if err != nil {
		t.Fatalf("parseInstances: %v", err)
	}
	if len(inst) != 2 {
		t.Fatalf("expected 2 merged instances, got %d", len(inst))
	}
	names := map[string]bool{"EVI-A": false, "EVI-B": false}
	for _, in := range inst {
		if _, ok := names[in.Name]; !ok {
			t.Errorf("unexpected instance name %q", in.Name)
			continue
		}
		names[in.Name] = true
	}
	for n, seen := range names {
		if !seen {
			t.Errorf("expected to see instance %q but didn't", n)
		}
	}
}

// =================================================================
// Detail tables (Phase A) — exercise the same parseInstances() path
// but populate the per-interface / per-IRB / per-bridge-domain /
// per-ESI subtables. Synthetic data — RFC 5737 IPs, IANA documentation
// MACs.
// =================================================================
const detailedEVI = `<rpc-reply>
  <evpn-instance-information>
    <evpn-instance>
      <evpn-instance-name>EVI-A</evpn-instance-name>
      <route-distinguisher>65000:1</route-distinguisher>
      <evpn-encap-type>VXLAN</evpn-encap-type>
      <local-interfaces>2</local-interfaces>
      <local-interfaces-up>1</local-interfaces-up>
      <evpn-interface-status-table>
        <evpn-interface>
          <evpn-interface-name>ae10.0</evpn-interface-name>
          <evpn-interface-esi>01:23:45:00:00:00:00:00:00:01</evpn-interface-esi>
          <evpn-interface-mode>all-active</evpn-interface-mode>
          <evpn-interface-status>Up</evpn-interface-status>
          <evpn-interface-etree-role>Root</evpn-interface-etree-role>
        </evpn-interface>
        <evpn-interface>
          <evpn-interface-name>ge-0/0/0.0</evpn-interface-name>
          <evpn-interface-esi>00:00:00:00:00:00:00:00:00:00</evpn-interface-esi>
          <evpn-interface-mode>single-homed</evpn-interface-mode>
          <evpn-interface-status>Down</evpn-interface-status>
          <evpn-interface-etree-role>Root</evpn-interface-etree-role>
        </evpn-interface>
      </evpn-interface-status-table>
      <irb-interfaces>1</irb-interfaces>
      <irb-interfaces-up>1</irb-interfaces-up>
      <irb-interface-status-table>
        <irb-interface>
          <irb-interface-name>irb.10</irb-interface-name>
          <irb-interface-vni-id>10001</irb-interface-vni-id>
          <irb-interface-status>Up</irb-interface-status>
          <irb-interface-l3-context>vrf-a</irb-interface-l3-context>
        </irb-interface>
      </irb-interface-status-table>
      <num-bridge-domains>1</num-bridge-domains>
      <bridge-domain-status-table>
        <bridge-domain>
          <vlan-id>10</vlan-id>
          <domain-id>10001</domain-id>
          <interfaces>3</interfaces>
          <interfaces-up>2</interfaces-up>
          <irb-interface>irb.10</irb-interface>
          <mode>Extended</mode>
          <mac-sync-status>Enabled</mac-sync-status>
        </bridge-domain>
      </bridge-domain-status-table>
      <evpn-num-esi>2</evpn-num-esi>
      <evpn-esi>
        <evpn-esi-value>01:23:45:00:00:00:00:00:00:01</evpn-esi-value>
        <evpn-esi-status>Resolved by IFL ae10.0</evpn-esi-status>
        <evpn-esi-local-intf-information>
          <evpn-esi-local-intf-name>ae10.0</evpn-esi-local-intf-name>
          <evpn-esi-local-intf-status>Up/Forwarding</evpn-esi-local-intf-status>
        </evpn-esi-local-intf-information>
        <evpn-esi-remote-pe-information>
          <evpn-esi-num-remote-pe>1</evpn-esi-num-remote-pe>
        </evpn-esi-remote-pe-information>
        <evpn-esi-df-information>
          <esi-df-election-algorithm>MOD based</esi-df-election-algorithm>
          <esi-designated-forwarder>192.0.2.1</esi-designated-forwarder>
          <esi-backup-forwarder>192.0.2.10</esi-backup-forwarder>
        </evpn-esi-df-information>
      </evpn-esi>
      <evpn-esi>
        <evpn-esi-value>05:ab:cd:00:00:00:00:00:01:00</evpn-esi-value>
        <evpn-esi-status></evpn-esi-status>
        <evpn-esi-remote-pe-information>
        </evpn-esi-remote-pe-information>
      </evpn-esi>
    </evpn-instance>
  </evpn-instance-information>
</rpc-reply>`

func TestParseDetailTables(t *testing.T) {
	inst, err := parseInstances([]byte(detailedEVI))
	if err != nil {
		t.Fatalf("parseInstances: %v", err)
	}
	if len(inst) != 1 {
		t.Fatalf("expected 1 instance, got %d", len(inst))
	}
	g := inst[0]

	if len(g.Interfaces) != 2 {
		t.Errorf("interfaces: got %d, want 2", len(g.Interfaces))
	}
	if g.Interfaces[0].Name != "ae10.0" || g.Interfaces[0].Status != "Up" {
		t.Errorf("interface[0]: got %+v", g.Interfaces[0])
	}
	if g.Interfaces[1].Status != "Down" {
		t.Errorf("interface[1] status: got %q", g.Interfaces[1].Status)
	}

	if len(g.IRBInterfaceTbl) != 1 {
		t.Errorf("irb: got %d, want 1", len(g.IRBInterfaceTbl))
	}
	if g.IRBInterfaceTbl[0].VNI != "10001" || g.IRBInterfaceTbl[0].L3Context != "vrf-a" {
		t.Errorf("irb[0]: got %+v", g.IRBInterfaceTbl[0])
	}

	if len(g.BridgeDomains) != 1 {
		t.Errorf("bridge-domains: got %d", len(g.BridgeDomains))
	}
	if g.BridgeDomains[0].Interfaces != 3 || g.BridgeDomains[0].InterfacesUp != 2 {
		t.Errorf("bridge-domain counts: got %v/%v", g.BridgeDomains[0].Interfaces, g.BridgeDomains[0].InterfacesUp)
	}

	if len(g.ESIs) != 2 {
		t.Errorf("esi: got %d", len(g.ESIs))
	}
	// First ESI is resolved with full DF info.
	if g.ESIs[0].DFInfo == nil || g.ESIs[0].DFInfo.DesignatedForwarder != "192.0.2.1" {
		t.Errorf("esi[0].df: got %+v", g.ESIs[0].DFInfo)
	}
	if g.ESIs[0].RemotePEInfo == nil || g.ESIs[0].RemotePEInfo.Count != 1 {
		t.Errorf("esi[0].remote-pe-count: got %+v", g.ESIs[0].RemotePEInfo)
	}
	// Second ESI has no DF info, no remote PEs — minimal entry.
	if g.ESIs[1].DFInfo != nil {
		t.Errorf("esi[1].df: expected nil, got %+v", g.ESIs[1].DFInfo)
	}
}

// =================================================================
// show evpn database state duplicate fixtures
// =================================================================
const duplicateEmpty = `<rpc-reply>
  <evpn-database-information>
  </evpn-database-information>
</rpc-reply>`

const duplicatePopulated = `<rpc-reply>
  <evpn-database-information>
    <evpn-database-instance>
      <instance-name>EVI-A</instance-name>
      <mac-entry>
        <vni-id>10001</vni-id>
        <mac-address>00:00:5e:00:53:01</mac-address>
      </mac-entry>
      <mac-entry>
        <vni-id>10001</vni-id>
        <mac-address>00:00:5e:00:53:02</mac-address>
      </mac-entry>
    </evpn-database-instance>
    <evpn-database-instance>
      <instance-name>EVI-B</instance-name>
      <mac-entry>
        <vni-id>10002</vni-id>
        <mac-address>00:00:5e:00:53:11</mac-address>
      </mac-entry>
    </evpn-database-instance>
  </evpn-database-information>
</rpc-reply>`

func TestParseDuplicateMACsEmpty(t *testing.T) {
	inst, err := parseDuplicateMACs([]byte(duplicateEmpty))
	if err != nil {
		t.Fatalf("parseDuplicateMACs: %v", err)
	}
	if len(inst) != 0 {
		t.Errorf("expected 0 instances, got %d", len(inst))
	}
}

func TestParseDuplicateMACsPopulated(t *testing.T) {
	inst, err := parseDuplicateMACs([]byte(duplicatePopulated))
	if err != nil {
		t.Fatalf("parseDuplicateMACs: %v", err)
	}
	if len(inst) != 2 {
		t.Fatalf("expected 2 instances, got %d", len(inst))
	}
	if inst[0].Name != "EVI-A" || len(inst[0].Entries) != 2 {
		t.Errorf("inst[0]: name=%q entries=%d", inst[0].Name, len(inst[0].Entries))
	}
	if inst[1].Name != "EVI-B" || len(inst[1].Entries) != 1 {
		t.Errorf("inst[1]: name=%q entries=%d", inst[1].Name, len(inst[1].Entries))
	}
}

// =================================================================
// show evpn l3-context fixtures
// =================================================================
const l3ContextEmpty = `<rpc-reply>
  <evpn-l3-context-information>
  </evpn-l3-context-information>
</rpc-reply>`

const l3ContextPopulated = `<rpc-reply>
  <evpn-l3-context-information>
    <evpn-l3-context>
      <context-name>vrf-a</context-name>
      <context-type>cfg</context-type>
      <context-advertisement-mode>direct</context-advertisement-mode>
      <context-router-mac>00:00:5e:00:53:42</context-router-mac>
      <context-encapsulation>VXLAN</context-encapsulation>
      <context-vni>10001</context-vni>
    </evpn-l3-context>
    <evpn-l3-context>
      <context-name>vrf-b</context-name>
      <context-type>cfg</context-type>
      <context-advertisement-mode>direct</context-advertisement-mode>
      <context-router-mac>00:00:5e:00:53:43</context-router-mac>
      <context-encapsulation>MPLS</context-encapsulation>
      <context-vni>20002</context-vni>
    </evpn-l3-context>
  </evpn-l3-context-information>
</rpc-reply>`

func TestParseL3ContextsEmpty(t *testing.T) {
	c, err := parseL3Contexts([]byte(l3ContextEmpty))
	if err != nil {
		t.Fatalf("parseL3Contexts: %v", err)
	}
	if len(c) != 0 {
		t.Errorf("expected 0 contexts, got %d", len(c))
	}
}

func TestParseL3ContextsPopulated(t *testing.T) {
	c, err := parseL3Contexts([]byte(l3ContextPopulated))
	if err != nil {
		t.Fatalf("parseL3Contexts: %v", err)
	}
	if len(c) != 2 {
		t.Fatalf("expected 2 contexts, got %d", len(c))
	}
	if c[0].Name != "vrf-a" || c[0].VNI != 10001 || c[0].Encapsulation != "VXLAN" {
		t.Errorf("ctx[0]: got %+v", c[0])
	}
	if c[1].Name != "vrf-b" || c[1].VNI != 20002 || c[1].Encapsulation != "MPLS" {
		t.Errorf("ctx[1]: got %+v", c[1])
	}
}
