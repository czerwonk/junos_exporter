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
