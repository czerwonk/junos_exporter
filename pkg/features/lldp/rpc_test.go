// SPDX-License-Identifier: MIT

package lldp

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLLDPNeighbors(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/15.1R7/junos">
    <lldp-neighbors-information>
        <lldp-neighbor-information>
            <lldp-local-port-id>ge-0/0/0</lldp-local-port-id>
            <lldp-local-parent-interface-name>ge-0/0/0</lldp-local-parent-interface-name>
            <lldp-remote-chassis-id>aa:bb:cc:dd:ee:ff</lldp-remote-chassis-id>
            <lldp-remote-port-id>eth0</lldp-remote-port-id>
            <lldp-remote-system-name>router1.example.com</lldp-remote-system-name>
        </lldp-neighbor-information>
        <lldp-neighbor-information>
            <lldp-local-port-id>ge-0/0/1</lldp-local-port-id>
            <lldp-local-parent-interface-name>ge-0/0/1</lldp-local-parent-interface-name>
            <lldp-remote-chassis-id>11:22:33:44:55:66</lldp-remote-chassis-id>
            <lldp-remote-port-id>eth1</lldp-remote-port-id>
            <lldp-remote-system-name>switch2.example.com</lldp-remote-system-name>
        </lldp-neighbor-information>
    </lldp-neighbors-information>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>`

	var res result
	err := xml.Unmarshal([]byte(body), &res)
	assert.NoError(t, err)

	assert.Len(t, res.Information.Neighbors, 2)

	// First neighbor
	assert.Equal(t, "ge-0/0/0", res.Information.Neighbors[0].LocalPortID)
	assert.Equal(t, "ge-0/0/0", res.Information.Neighbors[0].LocalParentInterfaceName)
	assert.Equal(t, "aa:bb:cc:dd:ee:ff", res.Information.Neighbors[0].RemoteChassisID)
	assert.Equal(t, "eth0", res.Information.Neighbors[0].RemotePortID)
	assert.Equal(t, "router1.example.com", res.Information.Neighbors[0].RemoteSystemName)

	// Second neighbor
	assert.Equal(t, "ge-0/0/1", res.Information.Neighbors[1].LocalPortID)
	assert.Equal(t, "ge-0/0/1", res.Information.Neighbors[1].LocalParentInterfaceName)
	assert.Equal(t, "11:22:33:44:55:66", res.Information.Neighbors[1].RemoteChassisID)
	assert.Equal(t, "eth1", res.Information.Neighbors[1].RemotePortID)
	assert.Equal(t, "switch2.example.com", res.Information.Neighbors[1].RemoteSystemName)
}

func TestParseLLDPLocalInformation(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/15.1R7/junos">
    <lldp-local-info>
        <lldp-local-interface-info>
            <lldp-local-interface-name>ge-0/0/0</lldp-local-interface-name>
            <lldp-parent-local-interface-name>ge-0/0/0</lldp-parent-local-interface-name>
            <lldp-local-interface-id>ge-0/0/0</lldp-local-interface-id>
            <lldp-local-interface-description>Connection to router1</lldp-local-interface-description>
            <lldp-local-interface-status>Up</lldp-local-interface-status>
        </lldp-local-interface-info>
        <lldp-local-interface-info>
            <lldp-local-interface-name>ge-0/0/1</lldp-local-interface-name>
            <lldp-parent-local-interface-name>ge-0/0/1</lldp-parent-local-interface-name>
            <lldp-local-interface-id>ge-0/0/1</lldp-local-interface-id>
            <lldp-local-interface-description>Connection to switch2</lldp-local-interface-description>
            <lldp-local-interface-status>Up</lldp-local-interface-status>
        </lldp-local-interface-info>
        <lldp-local-interface-info>
            <lldp-local-interface-name>fxp0</lldp-local-interface-name>
            <lldp-parent-local-interface-name>fxp0</lldp-parent-local-interface-name>
            <lldp-local-interface-id>fxp0</lldp-local-interface-id>
            <lldp-local-interface-description>Management interface</lldp-local-interface-description>
            <lldp-local-interface-status>Up</lldp-local-interface-status>
        </lldp-local-interface-info>
        <lldp-local-interface-info>
            <lldp-local-interface-name>ge-0/0/2</lldp-local-interface-name>
            <lldp-parent-local-interface-name>ge-0/0/2</lldp-parent-local-interface-name>
            <lldp-local-interface-id>ge-0/0/2</lldp-local-interface-id>
            <lldp-local-interface-description>Unused port</lldp-local-interface-description>
            <lldp-local-interface-status>Down</lldp-local-interface-status>
        </lldp-local-interface-info>
    </lldp-local-info>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>`

	var res localResult
	err := xml.Unmarshal([]byte(body), &res)
	assert.NoError(t, err)

	assert.Len(t, res.Information.LocalInterfaces, 4)

	// First interface (active)
	assert.Equal(t, "ge-0/0/0", res.Information.LocalInterfaces[0].InterfaceName)
	assert.Equal(t, "ge-0/0/0", res.Information.LocalInterfaces[0].ParentInterfaceName)
	assert.Equal(t, "ge-0/0/0", res.Information.LocalInterfaces[0].InterfaceID)
	assert.Equal(t, "Connection to router1", res.Information.LocalInterfaces[0].InterfaceDescription)
	assert.Equal(t, "Up", res.Information.LocalInterfaces[0].InterfaceStatus)

	// Second interface (active)
	assert.Equal(t, "ge-0/0/1", res.Information.LocalInterfaces[1].InterfaceName)
	assert.Equal(t, "ge-0/0/1", res.Information.LocalInterfaces[1].ParentInterfaceName)
	assert.Equal(t, "ge-0/0/1", res.Information.LocalInterfaces[1].InterfaceID)
	assert.Equal(t, "Connection to switch2", res.Information.LocalInterfaces[1].InterfaceDescription)
	assert.Equal(t, "Up", res.Information.LocalInterfaces[1].InterfaceStatus)

	// Management interface (should be filtered by collector)
	assert.Equal(t, "fxp0", res.Information.LocalInterfaces[2].InterfaceName)
	assert.Equal(t, "fxp0", res.Information.LocalInterfaces[2].ParentInterfaceName)
	assert.Equal(t, "fxp0", res.Information.LocalInterfaces[2].InterfaceID)
	assert.Equal(t, "Management interface", res.Information.LocalInterfaces[2].InterfaceDescription)
	assert.Equal(t, "Up", res.Information.LocalInterfaces[2].InterfaceStatus)

	// Down interface (should not be reported by collector)
	assert.Equal(t, "ge-0/0/2", res.Information.LocalInterfaces[3].InterfaceName)
	assert.Equal(t, "ge-0/0/2", res.Information.LocalInterfaces[3].ParentInterfaceName)
	assert.Equal(t, "ge-0/0/2", res.Information.LocalInterfaces[3].InterfaceID)
	assert.Equal(t, "Unused port", res.Information.LocalInterfaces[3].InterfaceDescription)
	assert.Equal(t, "Down", res.Information.LocalInterfaces[3].InterfaceStatus)
}