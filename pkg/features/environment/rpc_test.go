package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMultiREOutputQFXPEM(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/19XX/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>fpc0</re-name>
            
            <environment-component-information xmlns="http://xml.juniper.net/junos/19XX/junos-chassis">
                <environment-component-item>
                    <name>FPC 0 PEM 0</name>
                    <state>Online</state>
                    <airflow-direction>
                        <airflow>Airflow</airflow>
                        <direction>Front to Back</direction>
                    </airflow-direction>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 0</temperature-name>
                        <temperature junos:celsius="41">OK   41 degrees C / 105 degrees F</temperature>
                    </temperature-reading>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 1</temperature-name>
                        <temperature junos:celsius="47">OK   47 degrees C / 116 degrees F</temperature>
                    </temperature-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 0</fan-name>
                        <fan-speed>5120 RPM</fan-speed>
                    </fan-speed-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 1</fan-name>
                        <fan-speed>5120 RPM</fan-speed>
                    </fan-speed-reading>
                    <dc-information>
                        <dc-input>OK</dc-input>
                        <dc-detail>
                            <dc-voltage>12</dc-voltage>
                            <dc-current>8</dc-current>
                            <dc-power>96</dc-power>
                            <dc-load>11</dc-load>
                        </dc-detail>
                    </dc-information>
                </environment-component-item>
                <environment-component-item>
                    <name>FPC 0 PEM 1</name>
                    <state>Online</state>
                    <airflow-direction>
                        <airflow>Airflow</airflow>
                        <direction>Front to Back</direction>
                    </airflow-direction>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 0</temperature-name>
                        <temperature junos:celsius="40">OK   40 degrees C / 104 degrees F</temperature>
                    </temperature-reading>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 1</temperature-name>
                        <temperature junos:celsius="50">OK   50 degrees C / 122 degrees F</temperature>
                    </temperature-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 0</fan-name>
                        <fan-speed>5120 RPM</fan-speed>
                    </fan-speed-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 1</fan-name>
                        <fan-speed>5632 RPM</fan-speed>
                    </fan-speed-reading>
                    <dc-information>
                        <dc-input>OK</dc-input>
                        <dc-detail>
                            <dc-voltage>12</dc-voltage>
                            <dc-current>9</dc-current>
                            <dc-power>108</dc-power>
                            <dc-load>12</dc-load>
                        </dc-detail>
                    </dc-information>
                </environment-component-item>
            </environment-component-information>
        </multi-routing-engine-item>
        <multi-routing-engine-item>
            
            <re-name>fpc1</re-name>
            
            <environment-component-information xmlns="http://xml.juniper.net/junos/19XX/junos-chassis">
                <environment-component-item>
                    <name>FPC 1 PEM 0</name>
                    <state>Offline</state>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 0</temperature-name>
                        <temperature  junos:celsius="41">OK   41 degrees C / 105 degrees F</temperature>
                    </temperature-reading>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 1</temperature-name>
                        <temperature  junos:celsius="47">OK   47 degrees C / 116 degrees F</temperature>
                    </temperature-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 0</fan-name>
                        <fan-speed>5120 RPM</fan-speed>
                    </fan-speed-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 1</fan-name>
                        <fan-speed>5120 RPM</fan-speed>
                    </fan-speed-reading>
                    <dc-information>
                        <dc-input>OK</dc-input>
                        <dc-detail>
                            <dc-voltage>12</dc-voltage>
                            <dc-current>8</dc-current>
                            <dc-power>96</dc-power>
                            <dc-load>11</dc-load>
                        </dc-detail>
                    </dc-information>
                </environment-component-item>
                <environment-component-item>
                    <name>FPC 1 PEM 1</name>
                    <state>Online</state>
                    <airflow-direction>
                        <airflow>Airflow</airflow>
                        <direction>Front to Back</direction>
                    </airflow-direction>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 0</temperature-name>
                        <temperature junos:celsius="40">OK   40 degrees C / 104 degrees F</temperature>
                    </temperature-reading>
                    <temperature-reading>
                        <temperature-name>Temp Sensor 1</temperature-name>
                        <temperature junos:celsius="50">OK   50 degrees C / 122 degrees F</temperature>
                    </temperature-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 0</fan-name>
                        <fan-speed>5120 RPM</fan-speed>
                    </fan-speed-reading>
                    <fan-speed-reading>
                        <fan-name>Fan 1</fan-name>
                        <fan-speed>5632 RPM</fan-speed>
                    </fan-speed-reading>
                    <dc-information>
                        <dc-input>OK</dc-input>
                        <dc-detail>
                            <dc-voltage>12</dc-voltage>
                            <dc-current>9</dc-current>
                            <dc-power>108</dc-power>
                            <dc-load>12</dc-load>
                        </dc-detail>
                    </dc-information>
                </environment-component-item>
            </environment-component-information>
        </multi-routing-engine-item>
        
    </multi-routing-engine-results>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.Results.RoutingEngines[0].EnvironmentComponentInformation)

	// test first routing engine
	assert.Equal(t, "fpc0", rpc.Results.RoutingEngines[0].Name, "re-name")

	f := rpc.Results.RoutingEngines[0].EnvironmentComponentInformation.EnvironmentComponentItem[0]

	assert.Equal(t, "FPC 0 PEM 0", f.Name, "name")

	assert.Equal(t, "Online", f.State, "state")

	assert.Equal(t, "Fan 0", f.FanSpeedReading[0].FanName, "fan-name")
	assert.Equal(t, "5120 RPM", f.FanSpeedReading[0].FanSpeed, "fan-speed")

	assert.Equal(t, float64(96), f.DcInformation.DcDetail.DcPower, "dc-power")

	// test the second routing engine
	assert.Equal(t, "fpc1", rpc.Results.RoutingEngines[1].Name, "re-name")

	f = rpc.Results.RoutingEngines[1].EnvironmentComponentInformation.EnvironmentComponentItem[1]

	assert.Equal(t, "FPC 1 PEM 1", f.Name, "name")

	assert.Equal(t, "Fan 1", f.FanSpeedReading[1].FanName, "fan-name")
	assert.Equal(t, "5632 RPM", f.FanSpeedReading[1].FanSpeed, "fan-speed")

	assert.Equal(t, float64(108), f.DcInformation.DcDetail.DcPower, "dc-power")

}

func TestParseNoMultiREOutputMXPEM(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/17XX/junos">
    <environment-component-information xmlns="http://xml.juniper.net/junos/17XX/junos-chassis">
        <environment-component-item>
            <name>PEM 0</name>
            <state>Online</state>
            <temperature-reading>
                <temperature-name>Temperature</temperature-name>
                <temperature>OK</temperature>
            </temperature-reading>
            <dc-information>
                <dc-input>OK</dc-input>
                <dc-detail>
                    <dc-voltage>56</dc-voltage>
                    <dc-current>8</dc-current>
                    <dc-power>448</dc-power>
                    <dc-load>18</dc-load>
                </dc-detail>
            </dc-information>
            <power-information>
                <voltage-title>Voltage</voltage-title>
                <voltage>
                    <reference-voltage>48.0 V input</reference-voltage>
                    <actual-voltage>56500</actual-voltage>
                </voltage>
            </power-information>
        </environment-component-item>
        <environment-component-item>
            <name>PEM 1</name>
            <state>Online</state>
            <temperature-reading>
                <temperature-name>Temperature</temperature-name>
                <temperature>OK</temperature>
            </temperature-reading>
            <dc-information>
                <dc-input>OK</dc-input>
                <dc-detail>
                    <dc-voltage>56</dc-voltage>
                    <dc-current>3</dc-current>
                    <dc-power>168</dc-power>
                    <dc-load>6</dc-load>
                </dc-detail>
            </dc-information>
            <power-information>
                <voltage-title>Voltage</voltage-title>
                <voltage>
                    <reference-voltage>48.0 V input</reference-voltage>
                    <actual-voltage>56500</actual-voltage>
                </voltage>
            </power-information>
        </environment-component-item>
        <environment-component-item>
            <name>PEM 2</name>
            <state>Online</state>
            <temperature-reading>
                <temperature-name>Temperature</temperature-name>
                <temperature>OK</temperature>
            </temperature-reading>
            <dc-information>
                <dc-input>OK</dc-input>
                <dc-detail>
                    <dc-voltage>57</dc-voltage>
                    <dc-current>7</dc-current>
                    <dc-power>399</dc-power>
                    <dc-load>16</dc-load>
                </dc-detail>
            </dc-information>
            <power-information>
                <voltage-title>Voltage</voltage-title>
                <voltage>
                    <reference-voltage>48.0 V input</reference-voltage>
                    <actual-voltage>57000</actual-voltage>
                </voltage>
            </power-information>
        </environment-component-item>
        <environment-component-item>
            <name>PEM 3</name>
            <state>Online</state>
            <temperature-reading>
                <temperature-name>Temperature</temperature-name>
                <temperature>OK</temperature>
            </temperature-reading>
            <dc-information>
                <dc-input>OK</dc-input>
                <dc-detail>
                    <dc-voltage>57</dc-voltage>
                    <dc-current>1</dc-current>
                    <dc-power>57</dc-power>
                    <dc-load>2</dc-load>
                </dc-detail>
            </dc-information>
            <power-information>
                <voltage-title>Voltage</voltage-title>
                <voltage>
                    <reference-voltage>48.0 V input</reference-voltage>
                    <actual-voltage>57000</actual-voltage>
                </voltage>
            </power-information>
        </environment-component-item>
    </environment-component-information>
    <cli>
        <banner>{master}</banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.Results.RoutingEngines[0].EnvironmentComponentInformation)

	assert.Equal(t, "N/A", rpc.Results.RoutingEngines[0].Name, "re-name")

	f := rpc.Results.RoutingEngines[0].EnvironmentComponentInformation.EnvironmentComponentItem[0]

	assert.Equal(t, "PEM 0", f.Name, "name")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, float64(448), f.DcInformation.DcDetail.DcPower, "dc-power")

}

func TestParseNoMultiREOutputMX(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/17XXX/junos">
    <environment-information xmlns="http://xml.juniper.net/junos/17XXX/junos-chassis">
        <environment-item>
            <name>PEM 0</name>
            <class>Temp</class>
            <status>OK</status>
            <temperature junos:celsius="50">50 degrees C / 122 degrees F</temperature>
        </environment-item>
        <environment-item>
            <name>PEM 1</name>
            <class>Temp</class>
            <status>OK</status>
            <temperature junos:celsius="50">50 degrees C / 122 degrees F</temperature>
        </environment-item>
    </environment-information>
    <cli>
        <banner>{master}</banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.Results.RoutingEngines[0].EnvironmentInformation)

	assert.Equal(t, "N/A", rpc.Results.RoutingEngines[0].Name, "re-name")

	f := rpc.Results.RoutingEngines[0].EnvironmentInformation.Items[0]

	assert.Equal(t, "PEM 0", f.Name, "name")
	assert.Equal(t, "OK", f.Status, "status")
	assert.Equal(t, float64(50), f.Temperature.Value, "temperature")
}

func TestParseMultiREOutputSRX(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18XXX/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>node0</re-name>
            
            <environment-information xmlns="http://xml.juniper.net/junos/18XXX/junos-chassis">
                <environment-item>
                    <name>Power Supply 0</name>
                    <class>Power</class>
                    <status>OK</status>
                </environment-item>
                <environment-item>
                    <name>Power Supply 1</name>
                    <status>OK</status>
                </environment-item>
            </environment-information>
        </multi-routing-engine-item>
        
        <multi-routing-engine-item>
            
            <re-name>node1</re-name>
            
            <environment-information xmlns="http://xml.juniper.net/junos/18XXX/junos-chassis">
                <environment-item>
                    <name>Power Supply 0</name>
                    <class>Power</class>
                    <status>OK</status>
                </environment-item>
                <environment-item>
                    <name>Power Supply 1</name>
                    <status>OK</status>
                </environment-item>
            </environment-information>
        </multi-routing-engine-item>
        
    </multi-routing-engine-results>
    <cli>
        <banner>{secondary:node1}</banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	// test first routing engine
	assert.Equal(t, "node0", rpc.Results.RoutingEngines[0].Name, "re-name")

	f := rpc.Results.RoutingEngines[0].EnvironmentInformation.Items[0]

	assert.Equal(t, "Power Supply 0", f.Name, "name")
	assert.Equal(t, "OK", f.Status, "status")

	// test the second routing engine
	assert.Equal(t, "node1", rpc.Results.RoutingEngines[1].Name, "re-name")

	f = rpc.Results.RoutingEngines[1].EnvironmentInformation.Items[1]

	assert.Equal(t, "Power Supply 1", f.Name, "name")
	assert.Equal(t, "OK", f.Status, "status")
}
