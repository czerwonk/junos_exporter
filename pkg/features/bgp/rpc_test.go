// SPDX-License-Identifier: MIT

package bgp

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

func TestParseOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZ/junos">
    <bgp-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-routing">
        <bgp-peer junos:style="detail">
            <peer-address>10.0.0.1+179</peer-address>
            <peer-as>64496</peer-as>
            <local-address>10.0.1.1+33333</local-address>
            <local-as>64498</local-as>
            <peer-group>group1</peer-group>
            <peer-cfg-rti>CRI</peer-cfg-rti>
            <peer-fwd-rti>CRI</peer-fwd-rti>
            <peer-type>External</peer-type>
            <peer-state>Established</peer-state>
            <peer-flags>Sync</peer-flags>
            <last-state>OpenConfirm</last-state>
            <last-event>Refresh</last-event>
            <last-error>Hold Timer Expired Error</last-error>
            <bgp-option-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-routing">
                <export-policy>
                    BGP-EXPORT-1
                </export-policy>
                <import-policy>
                    BGP-IMPORT-1
                </import-policy>
                <bgp-options>Multihop Preference LocalAddress HoldTime PeerAS LocalAS Refresh</bgp-options>
                <bgp-options2></bgp-options2>
                <bgp-options-extended>GracefulShutdownRcv</bgp-options-extended>
                <local-address>10.0.1.1</local-address>
                <holdtime>10</holdtime>
                <preference>170</preference>
                <gshut-recv-local-preference>0</gshut-recv-local-preference>
                <local-as>64498</local-as>
                <local-system-as>0</local-system-as>
            </bgp-option-information>
            <flap-count>19</flap-count>
            <last-flap-event>HoldTime</last-flap-event>
            <bgp-error>
                <name>Hold Timer Expired Error</name>
                <send-count>21</send-count>
                <receive-count>1</receive-count>
            </bgp-error>
            <peer-id>10.0.0.1</peer-id>
            <local-id>10.0.1.1</local-id>
            <active-holdtime>10</active-holdtime>
            <keepalive-interval>3</keepalive-interval>
            <group-index>2</group-index>
            <peer-index>0</peer-index>
            <snmp-index>0</snmp-index>
            <bgp-peer-iosession>
                <iosession-thread-name>bgpio-0</iosession-thread-name>
                <iosession-state>Enabled</iosession-state>
            </bgp-peer-iosession>
            <bgp-bfd>
                <bfd-configuration-state>disabled</bfd-configuration-state>
                <bfd-operational-state>down</bfd-operational-state>
            </bgp-bfd>
            <peer-restart-nlri-configured>inet-unicast</peer-restart-nlri-configured>
            <nlri-type-peer>inet-unicast</nlri-type-peer>
            <nlri-type-session>inet-unicast</nlri-type-session>
            <peer-refresh-capability>2</peer-refresh-capability>
            <peer-stale-route-time-configured>120</peer-stale-route-time-configured>
            <peer-restart-time-received>120</peer-restart-time-received>
            <peer-restart-nlri-received>inet-unicast</peer-restart-nlri-received>
            <peer-restart-nlri-can-save-state>inet-unicast</peer-restart-nlri-can-save-state>
            <peer-restart-nlri-state-saved></peer-restart-nlri-state-saved>
            <peer-restart-nlri-negotiated>inet-unicast</peer-restart-nlri-negotiated>
            <peer-end-of-rib-received>inet-unicast</peer-end-of-rib-received>
            <peer-end-of-rib-sent>inet-unicast</peer-end-of-rib-sent>
            <peer-end-of-rib-scheduled></peer-end-of-rib-scheduled>
            <peer-no-llgr-helper/>
            <peer-4byte-as-capability-advertised>64496</peer-4byte-as-capability-advertised>
            <peer-addpath-not-supported/>
            <bgp-rib junos:style="detail">
                <name>CRI.inet.0</name>
                <rib-bit>20001</rib-bit>
                <bgp-rib-state>BGP restart is complete</bgp-rib-state>
                <vpn-rib-state>VPN restart is complete</vpn-rib-state>
                <send-state>in sync</send-state>
                <active-prefix-count>9</active-prefix-count>
                <received-prefix-count>9</received-prefix-count>
                <accepted-prefix-count>9</accepted-prefix-count>
                <suppressed-prefix-count>0</suppressed-prefix-count>
                <advertised-prefix-count>1</advertised-prefix-count>
            </bgp-rib>
            <last-received>0</last-received>
            <last-sent>2</last-sent>
            <last-checked>896903</last-checked>
            <input-messages>349663</input-messages>
            <input-updates>4</input-updates>
            <input-refreshes>3</input-refreshes>
            <input-octets>6643729</input-octets>
            <output-messages>330655</output-messages>
            <output-updates>4</output-updates>
            <output-refreshes>0</output-refreshes>
            <output-octets>6282561</output-octets>
            <bgp-output-queue>
                <number>1</number>
                <count>0</count>
                <table-name>CRI.inet.0</table-name>
                <rib-adv-nlri>inet-unicast</rib-adv-nlri>
            </bgp-output-queue>
        </bgp-peer>
        <bgp-peer junos:style="detail">
            <peer-address>10.0.0.2+179</peer-address>
            <peer-as>64497</peer-as>
            <local-address>10.0.1.1+44444</local-address>
            <local-as>64498</local-as>
            <peer-group>group2</peer-group>
            <peer-cfg-rti>CRI</peer-cfg-rti>
            <peer-fwd-rti>CRI</peer-fwd-rti>
            <peer-type>External</peer-type>
            <peer-state>Established</peer-state>
            <peer-flags>Sync</peer-flags>
            <last-state>OpenConfirm</last-state>
            <last-event>Refresh</last-event>
            <last-error>Cease</last-error>
            <bgp-option-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-routing">
                <export-policy>
                    BGP-EXPORT-2
                </export-policy>
                <import-policy>
                    BGP-IMPORT-2
                </import-policy>
                <bgp-options>Multihop Preference LocalAddress HoldTime PeerAS LocalAS Refresh</bgp-options>
                <bgp-options2></bgp-options2>
                <bgp-options-extended>GracefulShutdownRcv</bgp-options-extended>
                <local-address>10.0.1.1</local-address>
                <holdtime>10</holdtime>
                <preference>170</preference>
                <gshut-recv-local-preference>0</gshut-recv-local-preference>
                <local-as>64498</local-as>
                <local-system-as>0</local-system-as>
            </bgp-option-information>
            <flap-count>7</flap-count>
            <last-flap-event>HoldTime</last-flap-event>
            <bgp-error>
                <name>Hold Timer Expired Error</name>
                <send-count>5</send-count>
                <receive-count>1</receive-count>
            </bgp-error>
            <bgp-error>
                <name>Cease</name>
                <send-count>3</send-count>
                <receive-count>1</receive-count>
            </bgp-error>
            <peer-id>10.0.0.2</peer-id>
            <local-id>10.0.1.1</local-id>
            <active-holdtime>10</active-holdtime>
            <keepalive-interval>3</keepalive-interval>
            <group-index>3</group-index>
            <peer-index>0</peer-index>
            <snmp-index>1</snmp-index>
            <bgp-peer-iosession>
                <iosession-thread-name>bgpio-0</iosession-thread-name>
                <iosession-state>Enabled</iosession-state>
            </bgp-peer-iosession>
            <bgp-bfd>
                <bfd-configuration-state>disabled</bfd-configuration-state>
                <bfd-operational-state>down</bfd-operational-state>
            </bgp-bfd>
            <peer-restart-nlri-configured>inet-unicast</peer-restart-nlri-configured>
            <nlri-type-peer>inet-unicast</nlri-type-peer>
            <nlri-type-session>inet-unicast</nlri-type-session>
            <peer-refresh-capability>2</peer-refresh-capability>
            <peer-stale-route-time-configured>120</peer-stale-route-time-configured>
            <peer-restart-time-received>120</peer-restart-time-received>
            <peer-restart-nlri-received>inet-unicast</peer-restart-nlri-received>
            <peer-restart-nlri-can-save-state>inet-unicast</peer-restart-nlri-can-save-state>
            <peer-restart-nlri-state-saved></peer-restart-nlri-state-saved>
            <peer-restart-nlri-negotiated>inet-unicast</peer-restart-nlri-negotiated>
            <peer-end-of-rib-received>inet-unicast</peer-end-of-rib-received>
            <peer-end-of-rib-sent>inet-unicast</peer-end-of-rib-sent>
            <peer-end-of-rib-scheduled></peer-end-of-rib-scheduled>
            <peer-no-llgr-helper/>
            <peer-4byte-as-capability-advertised>64497</peer-4byte-as-capability-advertised>
            <peer-addpath-not-supported/>
            <bgp-rib junos:style="detail">
                <name>CRI.inet.0</name>
                <rib-bit>20000</rib-bit>
                <bgp-rib-state>BGP restart is complete</bgp-rib-state>
                <vpn-rib-state>VPN restart is complete</vpn-rib-state>
                <send-state>in sync</send-state>
                <active-prefix-count>0</active-prefix-count>
                <received-prefix-count>9</received-prefix-count>
                <accepted-prefix-count>9</accepted-prefix-count>
                <suppressed-prefix-count>0</suppressed-prefix-count>
                <advertised-prefix-count>1</advertised-prefix-count>
            </bgp-rib>
            <last-received>3</last-received>
            <last-sent>1</last-sent>
            <last-checked>1247359</last-checked>
            <input-messages>486332</input-messages>
            <input-updates>3</input-updates>
            <input-refreshes>1</input-refreshes>
            <input-octets>9240448</input-octets>
            <output-messages>459847</output-messages>
            <output-updates>2</output-updates>
            <output-refreshes>0</output-refreshes>
            <output-octets>8737177</output-octets>
            <bgp-output-queue>
                <number>1</number>
                <count>0</count>
                <table-name>CRI.inet.0</table-name>
                <rib-adv-nlri>inet-unicast</rib-adv-nlri>
            </bgp-output-queue>
        </bgp-peer>
    </bgp-information>
    <cli>
        <banner>{primary:node0}</banner>
    </cli>
</rpc-reply>`

	var b, err = rpc.UnpackRpcReply([]byte(body))

	if err != nil {
		t.Fatal(err)
	}

	result := information{}
	err = xml.Unmarshal(b, &result)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(result.Peers), "bgp-peer count")

	assert.Equal(t, "64496", result.Peers[0].ASN, "peer-as")
	assert.Equal(t, "group1", result.Peers[0].Group, "peer-group")
}
