// SPDX-License-Identifier: MIT

package power

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMultiREOutputSRX(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18XX/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>node0</re-name>
            
            <power-usage-information>
                <power-usage-item>
                    <name>PEM 0</name>
                    <state>Online</state>
                    <pem-capacity-detail>
                        <capacity-actual>0</capacity-actual>
                        <capacity-max>0</capacity-max>
                    </pem-capacity-detail>
                    <dc-output-detail>
                        <dc-power>120</dc-power>
                        <zone>0</zone>
                        <dc-current>10</dc-current>
                        <dc-voltage>12</dc-voltage>
                        <dc-load>4</dc-load>
                    </dc-output-detail>
                </power-usage-item>
                <power-usage-item>
                    <name>PEM 1</name>
                    <state>Online</state>
                    <pem-capacity-detail>
                        <capacity-actual>0</capacity-actual>
                        <capacity-max>0</capacity-max>
                    </pem-capacity-detail>
                    <dc-output-detail>
                        <dc-power>120</dc-power>
                        <zone>0</zone>
                        <dc-current>10</dc-current>
                        <dc-voltage>12</dc-voltage>
                        <dc-load>4</dc-load>
                    </dc-output-detail>
                </power-usage-item>
                <power-usage-system>
                    <power-usage-zone-information>
                        <zone>0</zone>
                        <capacity-actual>0</capacity-actual>
                        <capacity-max>0</capacity-max>
                        <capacity-allocated>240</capacity-allocated>
                        <capacity-remaining>0</capacity-remaining>
                        <capacity-actual-usage>240</capacity-actual-usage>
                    </power-usage-zone-information>
                    <capacity-sys-actual>0</capacity-sys-actual>
                    <capacity-sys-max>0</capacity-sys-max>
                    <capacity-sys-remaining>0</capacity-sys-remaining>
                </power-usage-system>
            </power-usage-information>
        </multi-routing-engine-item>
        
        <multi-routing-engine-item>
            
            <re-name>node1</re-name>
            
            <power-usage-information>
                <power-usage-item>
                    <name>PEM 0</name>
                    <state>Online</state>
                    <pem-capacity-detail>
                        <capacity-actual>0</capacity-actual>
                        <capacity-max>0</capacity-max>
                    </pem-capacity-detail>
                    <dc-output-detail>
                        <dc-power>120</dc-power>
                        <zone>0</zone>
                        <dc-current>10</dc-current>
                        <dc-voltage>12</dc-voltage>
                        <dc-load>4</dc-load>
                    </dc-output-detail>
                </power-usage-item>
                <power-usage-item>
                    <name>PEM 1</name>
                    <state>Online</state>
                    <pem-capacity-detail>
                        <capacity-actual>0</capacity-actual>
                        <capacity-max>0</capacity-max>
                    </pem-capacity-detail>
                    <dc-output-detail>
                        <dc-power>120</dc-power>
                        <zone>0</zone>
                        <dc-current>10</dc-current>
                        <dc-voltage>12</dc-voltage>
                        <dc-load>4</dc-load>
                    </dc-output-detail>
                </power-usage-item>
                <power-usage-system>
                    <power-usage-zone-information>
                        <zone>0</zone>
                        <capacity-actual>0</capacity-actual>
                        <capacity-max>0</capacity-max>
                        <capacity-allocated>240</capacity-allocated>
                        <capacity-remaining>0</capacity-remaining>
                        <capacity-actual-usage>240</capacity-actual-usage>
                    </power-usage-zone-information>
                    <capacity-sys-actual>0</capacity-sys-actual>
                    <capacity-sys-max>0</capacity-sys-max>
                    <capacity-sys-remaining>0</capacity-sys-remaining>
                </power-usage-system>
            </power-usage-information>
        </multi-routing-engine-item>
        
    </multi-routing-engine-results>
    <cli>
        <banner>{secondary:node1}</banner>
    </cli>
</rpc-reply>`

	rpc := multiRoutingEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.Results.RoutingEngine[0].PowerUsageInformation)

	// test first routing engine
	assert.Equal(t, "node0", rpc.Results.RoutingEngine[0].Name, "re-name")

	p := rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageItem[1]

	assert.Equal(t, "PEM 1", p.Name, "name")
	assert.Equal(t, "Online", p.State, "state")
	assert.Equal(t, int(120), p.DcOutputDetail.DcPower, "dc-power")
	assert.Equal(t, "0", p.DcOutputDetail.Zone, "zone")
	assert.Equal(t, int(10), p.DcOutputDetail.DcCurrent, "dc-current")
	assert.Equal(t, int(12), p.DcOutputDetail.DcVoltage, "dc-voltage")
	assert.Equal(t, int(4), p.DcOutputDetail.DcLoad, "dc-load")

	s := rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.PowerUsageZoneInformation[0]

	assert.Equal(t, "0", s.Zone, "zone")
	assert.Equal(t, int(0), s.CapacityActual, "capacity-actual")
	assert.Equal(t, int(0), s.CapacityMax, "capacity-max")
	assert.Equal(t, int(240), s.CapacityAllocated, "capacity-allocated")
	assert.Equal(t, int(0), s.CapacityRemaining, "capacity-remaining")
	assert.Equal(t, int(240), s.CapacityActualUsage, "capacity-actual-usage")

	assert.Equal(t, int(0), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysActual, "capacity-sys-usage")
	assert.Equal(t, int(0), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysMax, "capacity-sys-max")
	assert.Equal(t, int(0), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysRemaining, "capacity-sys-remaining")

	// test the second routing engine
	assert.Equal(t, "node1", rpc.Results.RoutingEngine[1].Name, "re-name")

	p = rpc.Results.RoutingEngine[1].PowerUsageInformation.PowerUsageItem[0]

	assert.Equal(t, "PEM 0", p.Name, "name")
	assert.Equal(t, "Online", p.State, "state")
	assert.Equal(t, int(120), p.DcOutputDetail.DcPower, "dc-power")
	assert.Equal(t, "0", p.DcOutputDetail.Zone, "zone")
	assert.Equal(t, int(10), p.DcOutputDetail.DcCurrent, "dc-current")
	assert.Equal(t, int(12), p.DcOutputDetail.DcVoltage, "dc-voltage")
	assert.Equal(t, int(4), p.DcOutputDetail.DcLoad, "dc-load")

	s = rpc.Results.RoutingEngine[1].PowerUsageInformation.PowerUsageSystem.PowerUsageZoneInformation[0]

	assert.Equal(t, "0", s.Zone, "zone")
	assert.Equal(t, int(0), s.CapacityActual, "capacity-actual")
	assert.Equal(t, int(0), s.CapacityMax, "capacity-max")
	assert.Equal(t, int(240), s.CapacityAllocated, "capacity-allocated")
	assert.Equal(t, int(0), s.CapacityRemaining, "capacity-remaining")
	assert.Equal(t, int(240), s.CapacityActualUsage, "capacity-actual-usage")

	assert.Equal(t, int(0), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysActual, "capacity-sys-usage")
	assert.Equal(t, int(0), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysMax, "capacity-sys-max")
	assert.Equal(t, int(0), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysRemaining, "capacity-sys-remaining")
}

func TestParseNoMultiREOutputMX(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/17XX/junos">
    <power-usage-information>
        <power-usage-item>
            <name>PEM 0</name>
            <state>Online</state>
            <dc-input-detail>
                <dc-input>OK</dc-input>
                <dc-expect-feed>1</dc-expect-feed>
                <dc-actual-feed>1</dc-actual-feed>
                <reference-voltage>48.0 V input</reference-voltage>
                <actual-voltage>56500</actual-voltage>
            </dc-input-detail>
            <pem-capacity-detail>
                <capacity-actual>2440</capacity-actual>
                <capacity-max>2440</capacity-max>
            </pem-capacity-detail>
            <dc-output-detail>
                <dc-power>448</dc-power>
                <zone>0</zone>
                <dc-current>8</dc-current>
                <dc-voltage>56</dc-voltage>
                <dc-load>18</dc-load>
            </dc-output-detail>
        </power-usage-item>
        <power-usage-item>
            <name>PEM 1</name>
            <state>Online</state>
            <dc-input-detail>
                <dc-input>OK</dc-input>
                <dc-expect-feed>1</dc-expect-feed>
                <dc-actual-feed>1</dc-actual-feed>
                <reference-voltage>48.0 V input</reference-voltage>
                <actual-voltage>56500</actual-voltage>
            </dc-input-detail>
            <pem-capacity-detail>
                <capacity-actual>2440</capacity-actual>
                <capacity-max>2440</capacity-max>
            </pem-capacity-detail>
            <dc-output-detail>
                <dc-power>168</dc-power>
                <zone>1</zone>
                <dc-current>3</dc-current>
                <dc-voltage>56</dc-voltage>
                <dc-load>6</dc-load>
            </dc-output-detail>
        </power-usage-item>
        <power-usage-item>
            <name>PEM 2</name>
            <state>Online</state>
            <dc-input-detail>
                <dc-input>OK</dc-input>
                <dc-expect-feed>1</dc-expect-feed>
                <dc-actual-feed>1</dc-actual-feed>
                <reference-voltage>48.0 V input</reference-voltage>
                <actual-voltage>57000</actual-voltage>
            </dc-input-detail>
            <pem-capacity-detail>
                <capacity-actual>2440</capacity-actual>
                <capacity-max>2440</capacity-max>
            </pem-capacity-detail>
            <dc-output-detail>
                <dc-power>448</dc-power>
                <zone>0</zone>
                <dc-current>8</dc-current>
                <dc-voltage>56</dc-voltage>
                <dc-load>18</dc-load>
            </dc-output-detail>
        </power-usage-item>
        <power-usage-item>
            <name>PEM 3</name>
            <state>Online</state>
            <dc-input-detail>
                <dc-input>OK</dc-input>
                <dc-expect-feed>1</dc-expect-feed>
                <dc-actual-feed>1</dc-actual-feed>
                <reference-voltage>48.0 V input</reference-voltage>
                <actual-voltage>56500</actual-voltage>
            </dc-input-detail>
            <pem-capacity-detail>
                <capacity-actual>2440</capacity-actual>
                <capacity-max>2440</capacity-max>
            </pem-capacity-detail>
            <dc-output-detail>
                <dc-power>57</dc-power>
                <zone>1</zone>
                <dc-current>1</dc-current>
                <dc-voltage>57</dc-voltage>
                <dc-load>2</dc-load>
            </dc-output-detail>
        </power-usage-item>
        <power-usage-system>
            <power-usage-zone-information>
                <zone>0</zone>
                <capacity-actual>2440</capacity-actual>
                <capacity-max>2440</capacity-max>
                <capacity-allocated>1520</capacity-allocated>
                <capacity-remaining>920</capacity-remaining>
                <capacity-actual-usage>896</capacity-actual-usage>
            </power-usage-zone-information>
            <power-usage-zone-information>
                <zone>1</zone>
                <capacity-actual>2440</capacity-actual>
                <capacity-max>2440</capacity-max>
                <capacity-allocated>465</capacity-allocated>
                <capacity-remaining>1975</capacity-remaining>
                <capacity-actual-usage>225</capacity-actual-usage>
            </power-usage-zone-information>
            <capacity-sys-actual>4880</capacity-sys-actual>
            <capacity-sys-max>4880</capacity-sys-max>
            <capacity-sys-remaining>2895</capacity-sys-remaining>
        </power-usage-system>
    </power-usage-information>
    <cli>
        <banner>{master}</banner>
    </cli>
</rpc-reply>`

	rpc := multiRoutingEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.Results.RoutingEngine[0].PowerUsageInformation)

	assert.Equal(t, "N/A", rpc.Results.RoutingEngine[0].Name, "re-name")

	p := rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageItem[1]

	assert.Equal(t, "PEM 1", p.Name, "name")
	assert.Equal(t, "Online", p.State, "state")
	assert.Equal(t, int(168), p.DcOutputDetail.DcPower, "dc-power")
	assert.Equal(t, "1", p.DcOutputDetail.Zone, "zone")
	assert.Equal(t, int(3), p.DcOutputDetail.DcCurrent, "dc-current")
	assert.Equal(t, int(56), p.DcOutputDetail.DcVoltage, "dc-voltage")
	assert.Equal(t, int(6), p.DcOutputDetail.DcLoad, "dc-load")

	s := rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.PowerUsageZoneInformation[1]

	assert.Equal(t, "1", s.Zone, "zone")
	assert.Equal(t, int(2440), s.CapacityActual, "capacity-actual")
	assert.Equal(t, int(2440), s.CapacityMax, "capacity-max")
	assert.Equal(t, int(465), s.CapacityAllocated, "capacity-allocated")
	assert.Equal(t, int(1975), s.CapacityRemaining, "capacity-remaining")
	assert.Equal(t, int(225), s.CapacityActualUsage, "capacity-actual-usage")

	assert.Equal(t, int(4880), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysActual, "capacity-sys-usage")
	assert.Equal(t, int(4880), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysMax, "capacity-sys-max")
	assert.Equal(t, int(2895), rpc.Results.RoutingEngine[0].PowerUsageInformation.PowerUsageSystem.CapacitySysRemaining, "capacity-sys-remaining")
}
