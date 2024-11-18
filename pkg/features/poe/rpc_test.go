package poe

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	resultData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/15.1R7/junos">
    <poe>
        <interface-information>
            <interface-name>ge-0/0/0</interface-name>
            <interface-enabled>Enabled</interface-enabled>
            <interface-status>ON</interface-status>
            <interface-power-limit>30.0W</interface-power-limit>
            <interface-lldp-negotiation-power>  </interface-lldp-negotiation-power>
            <interface-priority>Low</interface-priority>
            <interface-lldp-negotiation-priority>  </interface-lldp-negotiation-priority>
            <interface-power>5.3W</interface-power>
            <interface-asterisk> </interface-asterisk>
            <interface-class>4</interface-class>
        </interface-information>
        <interface-information>
            <interface-name>ge-0/0/23</interface-name>
            <interface-enabled>Enabled</interface-enabled>
            <interface-status>OFF</interface-status>
            <interface-power-limit>15.4W</interface-power-limit>
            <interface-lldp-negotiation-power>  </interface-lldp-negotiation-power>
            <interface-priority>Low</interface-priority>
            <interface-lldp-negotiation-priority>  </interface-lldp-negotiation-priority>
            <interface-power>0.0W</interface-power>
            <interface-asterisk> </interface-asterisk>
            <interface-class>not-applicable</interface-class>
        </interface-information>
    </poe>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>
`

	var result poeInterfaceResult

	err := xml.Unmarshal([]byte(resultData), &result)
	assert.NoError(t, err)

	assert.Equal(t, "ge-0/0/0", result.Poe.InterfaceInformation[0].Name)
	assert.Equal(t, "Enabled", result.Poe.InterfaceInformation[0].Enabled)
	assert.Equal(t, "ON", result.Poe.InterfaceInformation[0].Status)
	assert.Equal(t, "30.0W", result.Poe.InterfaceInformation[0].PowerLimit)
	assert.Equal(t, "5.3W", result.Poe.InterfaceInformation[0].Power)
	assert.Equal(t, "4", result.Poe.InterfaceInformation[0].Class)

	assert.Equal(t, "ge-0/0/23", result.Poe.InterfaceInformation[1].Name)
	assert.Equal(t, "Enabled", result.Poe.InterfaceInformation[1].Enabled)
	assert.Equal(t, "OFF", result.Poe.InterfaceInformation[1].Status)
	assert.Equal(t, "15.4W", result.Poe.InterfaceInformation[1].PowerLimit)
	assert.Equal(t, "0.0W", result.Poe.InterfaceInformation[1].Power)
	assert.Equal(t, "not-applicable", result.Poe.InterfaceInformation[1].Class)
}
