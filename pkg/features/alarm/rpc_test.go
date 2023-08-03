// SPDX-License-Identifier: MIT

package alarm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test multi routing engine
func TestParseOutputMultiRESystemAlarms(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZ/junos">
    <multi-routing-engine-results>

        <multi-routing-engine-item>

            <re-name>node0</re-name>

            <alarm-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-alarm">
                <alarm-summary>
                    <active-alarm-count>2</active-alarm-count>
                </alarm-summary>
                <alarm-detail>
                    <alarm-time junos:seconds="1684172206">2023-05-15 17:36:46 UTC</alarm-time>
                    <alarm-class>Minor</alarm-class>
                    <alarm-description>Autorecovery information needs to be saved</alarm-description>
                    <alarm-short-description>autorecovery-save-r</alarm-short-description>
                    <alarm-type>Autorecovery</alarm-type>
                </alarm-detail>
                <alarm-detail>
                    <alarm-time junos:seconds="1684172206">2023-05-15 17:36:46 UTC</alarm-time>
                    <alarm-class>Minor</alarm-class>
                    <alarm-description>Rescue configuration is not set</alarm-description>
                    <alarm-short-description>no-rescue</alarm-short-description>
                    <alarm-type>Configuration</alarm-type>
                </alarm-detail>
            </alarm-information>
        </multi-routing-engine-item>

        <multi-routing-engine-item>

            <re-name>node1</re-name>

            <alarm-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-alarm">
                <alarm-summary>
                    <active-alarm-count>2</active-alarm-count>
                </alarm-summary>
                <alarm-detail>
                    <alarm-time junos:seconds="1685967777">2023-06-05 12:22:57 UTC</alarm-time>
                    <alarm-class>Minor</alarm-class>
                    <alarm-description>Autorecovery information needs to be saved</alarm-description>
                    <alarm-short-description>autorecovery-save-r</alarm-short-description>
                    <alarm-type>Autorecovery</alarm-type>
                </alarm-detail>
                <alarm-detail>
                    <alarm-time junos:seconds="1685967777">2023-06-05 12:22:57 UTC</alarm-time>
                    <alarm-class>Minor</alarm-class>
                    <alarm-description>Rescue configuration is not set</alarm-description>
                    <alarm-short-description>no-rescue</alarm-short-description>
                    <alarm-type>Configuration</alarm-type>
                </alarm-detail>
            </alarm-information>
        </multi-routing-engine-item>

    </multi-routing-engine-results>
    <cli>
        <banner>{primary:node0}</banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "node0", rpc.Information.RoutingEngines[0].Name, "re-name")
	assert.Equal(t, "Autorecovery", rpc.Information.RoutingEngines[0].AlarmInfo.Details[0].Type, "alarm-type")
	assert.Equal(t, "Minor", rpc.Information.RoutingEngines[0].AlarmInfo.Details[0].Class, "alarm-class")
	assert.Equal(t, "Autorecovery information needs to be saved", rpc.Information.RoutingEngines[0].AlarmInfo.Details[0].Description, "alarm-description")

	assert.Equal(t, "node1", rpc.Information.RoutingEngines[1].Name, "re-name")
	assert.Equal(t, "Configuration", rpc.Information.RoutingEngines[1].AlarmInfo.Details[1].Type, "alarm-type")
	assert.Equal(t, "Minor", rpc.Information.RoutingEngines[1].AlarmInfo.Details[1].Class, "alarm-class")
	assert.Equal(t, "Rescue configuration is not set", rpc.Information.RoutingEngines[1].AlarmInfo.Details[1].Description, "alarm-description")
}

// Test no multi routing engine
func TestParseOutputSingleRESystemAlarms(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZ/junos">
    <alarm-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-alarm">
        <alarm-summary>
            <active-alarm-count>2</active-alarm-count>
        </alarm-summary>
        <alarm-detail>
            <alarm-time junos:seconds="1635533090">
                2021-10-29 18:44:50 UTC
            </alarm-time>
            <alarm-class>Minor</alarm-class>
            <alarm-description>Autorecovery information needs to be saved</alarm-description>
            <alarm-short-description>autorecovery-save-r</alarm-short-description>
            <alarm-type>Autorecovery</alarm-type>
        </alarm-detail>
        <alarm-detail>
            <alarm-time junos:seconds="1635533089">
                2021-10-29 18:44:49 UTC
            </alarm-time>
            <alarm-class>Minor</alarm-class>
            <alarm-description>Rescue configuration is not set</alarm-description>
            <alarm-short-description>no-rescue</alarm-short-description>
            <alarm-type>Configuration</alarm-type>
        </alarm-detail>
    </alarm-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "N/A", rpc.Information.RoutingEngines[0].Name, "re-name")
	assert.Equal(t, "Autorecovery", rpc.Information.RoutingEngines[0].AlarmInfo.Details[0].Type, "alarm-type")
	assert.Equal(t, "Minor", rpc.Information.RoutingEngines[0].AlarmInfo.Details[0].Class, "alarm-class")
	assert.Equal(t, "Autorecovery information needs to be saved", rpc.Information.RoutingEngines[0].AlarmInfo.Details[0].Description, "alarm-description")
}
