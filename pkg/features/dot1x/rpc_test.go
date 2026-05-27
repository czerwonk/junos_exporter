// SPDX-License-Identifier: MIT

package dot1x

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/22.4R0/junos">
    <dot1x-interface-information>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/1.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/3.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/4.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/5.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/6.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/7.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/8.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/9.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/10.0</interface-name>
            <user-mac-address>5C:11:DD:55:47:99</user-mac-address>
            <authenticated-method>Mac Radius</authenticated-method>
            <authenticated-vlan>23</authenticated-vlan>
            <authenticated-voip-vlan>-</authenticated-voip-vlan>
            <user-name>5c11dd554799</user-name>
            <state>Authenticated</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/27.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/28.0</interface-name>
            <user-mac-address>F4:CC:DD:55:44:77</user-mac-address>
            <authenticated-method>Radius</authenticated-method>
            <authenticated-vlan>66</authenticated-vlan>
            <authenticated-voip-vlan>-</authenticated-voip-vlan>
            <user-name>testuser</user-name>
            <state>Authenticated</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/29.0</interface-name>
            <state>Initialize</state>   
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-0/0/30.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/10.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/11.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/12.0</interface-name>
            <state>Initialize</state>
        </interface>                    
        <interface junos:style="extensive">
            <interface-name>ge-1/0/13.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/14.0</interface-name>
            <state>Connecting</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/15.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/16.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/17.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/18.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/19.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/20.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/21.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/22.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/23.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/24.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/25.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/26.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/27.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/28.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/29.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/32.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/36.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/44.0</interface-name>
            <state>Initialize</state>
        </interface>
        <interface junos:style="extensive">
            <interface-name>ge-1/0/46.0</interface-name>
            <state>Initialize</state>
        </interface>
    </dot1x-interface-information>
    <cli>
        <banner>{master:1}</banner>
    </cli>
</rpc-reply>`

	rpc := result{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	// test rtt
	assert.Equal(t, "ge-0/0/1.0", rpc.Results.Interfaces[0].InterfaceName, "interface-name")
	assert.Equal(t, "Authenticated", rpc.Results.Interfaces[8].State, "state")

	//<measurement-max>194</measurement-max>
	assert.Equal(t, "Mac Radius", rpc.Results.Interfaces[8].AuthenticatedMethod, "authenticated-method")

}

func parseXML(b []byte, res *result) error {
	err := xml.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	return nil
}
