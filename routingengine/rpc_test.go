package routingengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test multi routing engine
func TestParseOutputMultiRE(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/XXX/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>node0</re-name>
            
            <route-engine-information xmlns="http://xml.juniper.net/junos/XXX/junos-chassis">
                <route-engine>
                    <status>OK</status>
                    <temperature junos:celsius="35">35 degrees C / 95 degrees F</temperature>
                    <cpu-temperature junos:celsius="35">35 degrees C / 95 degrees F</cpu-temperature>
                    <memory-system-total>1905</memory-system-total>
                    <memory-system-total-used>667</memory-system-total-used>
                    <memory-system-total-util>35</memory-system-total-util>
                    <memory-buffer-utilization>34</memory-buffer-utilization>
                    <cpu-user>61</cpu-user>
                    <cpu-background>0</cpu-background>
                    <cpu-system>35</cpu-system>
                    <cpu-interrupt>4</cpu-interrupt>
                    <cpu-idle>0</cpu-idle>
                    <cpu-user1>19</cpu-user1>
                    <cpu-background1>0</cpu-background1>
                    <cpu-system1>14</cpu-system1>
                    <cpu-interrupt1>1</cpu-interrupt1>
                    <cpu-idle1>66</cpu-idle1>
                    <cpu-user2>17</cpu-user2>
                    <cpu-background2>0</cpu-background2>
                    <cpu-system2>16</cpu-system2>
                    <cpu-interrupt2>2</cpu-interrupt2>
                    <cpu-idle2>66</cpu-idle2>
                    <cpu-user3>16</cpu-user3>
                    <cpu-background3>0</cpu-background3>
                    <cpu-system3>16</cpu-system3>
                    <cpu-interrupt3>2</cpu-interrupt3>
                    <cpu-idle3>66</cpu-idle3>
                    <model>SRX Routing Engine</model>
                    <serial-number>BUILTIN</serial-number>
                    <start-time junos:seconds="1603359895">2020-10-22 09:44:55 UTC</start-time>
                    <up-time junos:seconds="23071720">267 days, 48 minutes, 40 seconds</up-time>
                    <last-reboot-reason>0x4000:VJUNOS reboot</last-reboot-reason>
                    <load-average-one>0.88</load-average-one>
                    <load-average-five>1.23</load-average-five>
                    <load-average-fifteen>1.21</load-average-fifteen>
                </route-engine>
            </route-engine-information>
        </multi-routing-engine-item>
        
        <multi-routing-engine-item>
            
            <re-name>node1</re-name>
            
            <route-engine-information xmlns="http://xml.juniper.net/junos/XXX/junos-chassis">
                <route-engine>
                    <status>OK</status>
                    <temperature junos:celsius="36">36 degrees C / 96 degrees F</temperature>
                    <cpu-temperature junos:celsius="36">36 degrees C / 96 degrees F</cpu-temperature>
                    <memory-system-total>1905</memory-system-total>
                    <memory-system-total-used>438</memory-system-total-used>
                    <memory-system-total-util>23</memory-system-total-util>
                    <memory-buffer-utilization>22</memory-buffer-utilization>
                    <cpu-user>3</cpu-user>
                    <cpu-background>0</cpu-background>
                    <cpu-system>3</cpu-system>
                    <cpu-interrupt>0</cpu-interrupt>
                    <cpu-idle>94</cpu-idle>
                    <cpu-user1>10</cpu-user1>
                    <cpu-background1>0</cpu-background1>
                    <cpu-system1>13</cpu-system1>
                    <cpu-interrupt1>2</cpu-interrupt1>
                    <cpu-idle1>76</cpu-idle1>
                    <cpu-user2>14</cpu-user2>
                    <cpu-background2>0</cpu-background2>
                    <cpu-system2>13</cpu-system2>
                    <cpu-interrupt2>1</cpu-interrupt2>
                    <cpu-idle2>72</cpu-idle2>
                    <cpu-user3>15</cpu-user3>
                    <cpu-background3>0</cpu-background3>
                    <cpu-system3>13</cpu-system3>
                    <cpu-interrupt3>1</cpu-interrupt3>
                    <cpu-idle3>72</cpu-idle3>
                    <model>SRX Routing Engine</model>
                    <serial-number>BUILTIN</serial-number>
                    <start-time junos:seconds="1603360416">2020-10-22 09:53:36 UTC</start-time>
                    <up-time junos:seconds="23071199">267 days, 39 minutes, 59 seconds</up-time>
                    <last-reboot-reason>0x4000:VJUNOS reboot</last-reboot-reason>
                    <load-average-one>0.65</load-average-one>
                    <load-average-five>0.91</load-average-five>
                    <load-average-fifteen>0.84</load-average-fifteen>
                </route-engine>
            </route-engine-information>
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

	// test routing engine 0
	assert.Equal(t, "node0", rpc.Results.RoutingEngines[0].Name, "re-name")

	// test first route engine
	assert.Equal(t, "OK", rpc.Results.RoutingEngines[0].Information.RouteEngines[0].Status, "status")

	assert.Equal(t, float64(35), rpc.Results.RoutingEngines[0].Information.RouteEngines[0].CPUTemperature.Value, "cpu-temperature")

	assert.Equal(t, float64(1905), rpc.Results.RoutingEngines[0].Information.RouteEngines[0].MemorySystemTotal, "memory-system-total")

	assert.Equal(t, float64(19), rpc.Results.RoutingEngines[0].Information.RouteEngines[0].CPUUser1, "cpu-user1")
	// test routing engine 1
	assert.Equal(t, "node1", rpc.Results.RoutingEngines[1].Name, "re-name")

	assert.Equal(t, "OK", rpc.Results.RoutingEngines[1].Information.RouteEngines[0].Status, "status")

	assert.Equal(t, float64(36), rpc.Results.RoutingEngines[1].Information.RouteEngines[0].CPUTemperature.Value, "cpu-temperature")
}

// Test no multi routing engine
func TestParseOutputNoMultiRE(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZ/junos">
    <route-engine-information xmlns="http://xml.juniper.net/junos/ZZZ/junos-chassis">
        <route-engine>
            <slot>0</slot>
            <mastership-state>master</mastership-state>
            <mastership-priority>master</mastership-priority>
            <status>OK</status>
            <temperature junos:celsius="36">36 degrees C / 96 degrees F</temperature>
            <cpu-temperature junos:celsius="31">31 degrees C / 87 degrees F</cpu-temperature>
            <memory-dram-size>32713 MB</memory-dram-size>
            <memory-installed-size>(32768 MB installed)</memory-installed-size>
            <memory-buffer-utilization>11</memory-buffer-utilization>
            <cpu-user>2</cpu-user>
            <cpu-background>0</cpu-background>
            <cpu-system>3</cpu-system>
            <cpu-interrupt>1</cpu-interrupt>
            <cpu-idle>94</cpu-idle>
            <model>Something</model>
            <serial-number>12456</serial-number>
            <start-time junos:seconds="1623283649">2021-06-10 02:07:29 CEST</start-time>
            <up-time junos:seconds="3148588">36 days, 10 hours, 36 minutes, 28 seconds</up-time>
            <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
            <load-average-one>0.23</load-average-one>
            <load-average-five>0.41</load-average-five>
            <load-average-fifteen>0.47</load-average-fifteen>
        </route-engine>
        <route-engine>
            <slot>1</slot>
            <mastership-state>backup</mastership-state>
            <mastership-priority>backup</mastership-priority>
            <status>OK</status>
            <temperature junos:celsius="34">34 degrees C / 93 degrees F</temperature>
            <cpu-temperature junos:celsius="32">32 degrees C / 89 degrees F</cpu-temperature>
            <memory-dram-size>32713 MB</memory-dram-size>
            <memory-installed-size>(32768 MB installed)</memory-installed-size>
            <memory-buffer-utilization>14</memory-buffer-utilization>
            <cpu-user>5</cpu-user>
            <cpu-background>0</cpu-background>
            <cpu-system>1</cpu-system>
            <cpu-interrupt>0</cpu-interrupt>
            <cpu-idle>94</cpu-idle>
            <cpu-user1>6</cpu-user1>
            <cpu-background1>0</cpu-background1>
            <cpu-system1>2</cpu-system1>
            <cpu-interrupt1>0</cpu-interrupt1>
            <cpu-idle1>92</cpu-idle1>
            <cpu-user2>6</cpu-user2>
            <cpu-background2>0</cpu-background2>
            <cpu-system2>2</cpu-system2>
            <cpu-interrupt2>0</cpu-interrupt2>
            <cpu-idle2>92</cpu-idle2>
            <cpu-user3>6</cpu-user3>
            <cpu-background3>0</cpu-background3>
            <cpu-system3>2</cpu-system3>
            <cpu-interrupt3>0</cpu-interrupt3>
            <cpu-idle3>92</cpu-idle3>
            <model>Something</model>
            <serial-number>12345</serial-number>
            <start-time junos:seconds="1623282379">2021-06-10 01:46:19 CEST</start-time>
            <up-time junos:seconds="3149860">36 days, 10 hours, 57 minutes, 40 seconds</up-time>
            <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
            <load-average-one>0.43</load-average-one>
            <load-average-five>0.49</load-average-five>
            <load-average-fifteen>0.44</load-average-fifteen>
        </route-engine>
    </route-engine-information>
    <cli>
        <banner>{backup}</banner>
    </cli>
</rpc-reply>`

	rpc := multiEngineResult{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "N/A", rpc.Results.RoutingEngines[0].Name, "re-name")

	// test first route engine
	assert.Equal(t, "0", rpc.Results.RoutingEngines[0].Information.RouteEngines[0].Slot, "slot")

	assert.Equal(t, float64(31), rpc.Results.RoutingEngines[0].Information.RouteEngines[0].CPUTemperature.Value, "cpu-temperature")

	// test second route engine
	assert.Equal(t, "1", rpc.Results.RoutingEngines[0].Information.RouteEngines[1].Slot, "slot")

	assert.Equal(t, uint64(3149860), rpc.Results.RoutingEngines[0].Information.RouteEngines[1].UpTime.Seconds, "up-time")

}
