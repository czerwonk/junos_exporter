package ipsec

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test multi routing engine
func TestParseSRXOutputMultiRE(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R3/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>node0</re-name>
            
            <ipsec-security-associations-information junos:style="brief">
                <total-active-tunnels>2</total-active-tunnels>
                <total-ipsec-sas>4</total-ipsec-sas>
                <ipsec-security-associations-block>
                    <sa-block-state>up</sa-block-state>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>bb0d675a</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>24979</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>a9f5fbf3</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>24979</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>f2f52201</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>25347</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>7e5c22ff</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>25347</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                </ipsec-security-associations-block>
                <ipsec-security-associations-block>
                    <sa-block-state>up</sa-block-state>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>9bb78e12</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>2801</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>b8ce3a62</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>2801</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>8a2f7ff3</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>28556</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>d78c51e8</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>28556</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                </ipsec-security-associations-block>
            </ipsec-security-associations-information>
        </multi-routing-engine-item>
    </multi-routing-engine-results>
    <cli>
        <banner>{secondary:node1}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	// test routing engine 0
	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].IpSec)

	assert.Equal(t, "node0", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	assert.Equal(t, 2, rpc.MultiRoutingEngineResults.RoutingEngine[0].IpSec.ActiveTunnels, "total-active-tunnels")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].IpSec.SecurityAssociations[0]

	assert.Equal(t, "up", f.State, "state")

	r := f.SecurityAssociations[0]

	assert.Equal(t, int64(231076), r.TunnelIndex, "sa-tunnel-index")

	r = f.SecurityAssociations[1]

	assert.Equal(t, "ESP", r.Protocol, "sa-protocol")

	assert.Equal(t, "1.1.1.1", r.RemoteGateway, "sa-remote-gateway")
}

// Test no multi routing engine
func TestParseSRXOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R3/junos">
            <ipsec-security-associations-information junos:style="brief">
                <total-active-tunnels>2</total-active-tunnels>
                <total-ipsec-sas>4</total-ipsec-sas>
                <ipsec-security-associations-block>
                    <sa-block-state>up</sa-block-state>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>bb0d675a</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>24979</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>a9f5fbf3</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>24979</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>f2f52201</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>25347</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>231076</sa-tunnel-index>
                        <sa-spi>7e5c22ff</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>1.1.1.1</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>25347</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                </ipsec-security-associations-block>
                <ipsec-security-associations-block>
                    <sa-block-state>up</sa-block-state>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>9bb78e12</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>2801</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>b8ce3a62</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>2801</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&lt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>8a2f7ff3</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>28556</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                    <ipsec-security-associations>
                        <sa-direction>&gt;</sa-direction>
                        <sa-tunnel-index>131074</sa-tunnel-index>
                        <sa-spi>d78c51e8</sa-spi>
                        <sa-aux-spi>0</sa-aux-spi>
                        <sa-remote-gateway>2.2.2.2</sa-remote-gateway>
                        <sa-port>401</sa-port>
                        <sa-vpn-monitoring-state>-</sa-vpn-monitoring-state>
                        <sa-protocol>ESP</sa-protocol>
                        <sa-esp-encryption-algorithm>3des</sa-esp-encryption-algorithm>
                        <sa-hmac-algorithm>sha256</sa-hmac-algorithm>
                        <sa-hard-lifetime>28556</sa-hard-lifetime>
                        <sa-lifesize-remaining>unlim</sa-lifesize-remaining>
                        <sa-virtual-system>root</sa-virtual-system>
                    </ipsec-security-associations>
                </ipsec-security-associations-block>
            </ipsec-security-associations-information>
    <cli>
        <banner>{secondary:node1}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].IpSec)

	assert.Equal(t, "N/A", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	assert.Equal(t, 2, rpc.MultiRoutingEngineResults.RoutingEngine[0].IpSec.ActiveTunnels, "total-active-tunnels")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].IpSec.SecurityAssociations[0]

	assert.Equal(t, "up", f.State, "state")

	r := f.SecurityAssociations[0]

	assert.Equal(t, int64(231076), r.TunnelIndex, "sa-tunnel-index")

	r = f.SecurityAssociations[1]

	assert.Equal(t, "ESP", r.Protocol, "sa-protocol")
}

// Test no multi routing engine
func TestParseConfigOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R3/junos">
    <configuration>
            <security>
                <ipsec>
                    <proposal>
                        <name>test-ipsec1</name>
                        <protocol>esp</protocol>
                        <authentication-algorithm>hmac-sha-256-128</authentication-algorithm>
                        <encryption-algorithm>3des-cbc</encryption-algorithm>
                        <lifetime-seconds>28800</lifetime-seconds>
                    </proposal>
                    <policy>
                        <name>test-ipsec1</name>
                        <proposals>test-ipsec1</proposals>
                    </policy>
                    <vpn>
                        <name>vpn1</name>
                        <bind-interface>en0.0</bind-interface>
                        <ike>
                            <gateway>gateway1</gateway>
                            <ipsec-policy>test-ipsec1</ipsec-policy>
                        </ike>
                        <establish-tunnels>immediately</establish-tunnels>
                    </vpn>
                    <vpn>
                        <name>vpn2</name>
                        <bind-interface>en0.0</bind-interface>
                        <ike>
                            <gateway>gateway2</gateway>
                            <ipsec-policy>test-ipsec2</ipsec-policy>
                        </ike>
                        <establish-tunnels>immediately</establish-tunnels>
                    </vpn>
                </ipsec>
            </security>
    </configuration>
</rpc-reply>`

	rpc := ConfigurationSecurityIpsec{}
	err := xml.Unmarshal([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.Configuration.Security.Ipsec)

	assert.Equal(t, "test-ipsec1", rpc.Configuration.Security.Ipsec.Proposal.Name)

	assert.Equal(t, "test-ipsec1", rpc.Configuration.Security.Ipsec.Policy.Name)

	f := rpc.Configuration.Security.Ipsec.Vpn

	assert.Equal(t, "vpn1", f[0].Name, "name")

	assert.Equal(t, "vpn2", f[1].Name, "name")

	assert.Equal(t, "test-ipsec1", f[0].Ike.IpsecPolicy, "ipsec-policy")
	assert.Equal(t, "gateway2", f[1].Ike.Gateway, "gateway")

	assert.Equal(t, 2, len(f), "configured vpns")
}
