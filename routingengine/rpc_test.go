package routingengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test multi routing engine
func TestParseOutputMultiRE(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZZ/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>node0</re-name>
            
            <route-engine-information xmlns="http://xml.juniper.net/junos/ZZZZ/junos-chassis">
                <route-engine>
                    <status>OK</status>
                    <temperature junos:celsius="44">44 degrees C / 111 degrees F</temperature>
                    <cpu-temperature junos:celsius="42">42 degrees C / 107 degrees F</cpu-temperature>
                    <memory-system-total>2048</memory-system-total>
                    <memory-system-total-used>1024</memory-system-total-used>
                    <memory-system-total-util>50</memory-system-total-util>
                    <memory-control-plane>1072</memory-control-plane>
                    <memory-control-plane-used>461</memory-control-plane-used>
                    <memory-control-plane-util>43</memory-control-plane-util>
                    <memory-data-plane>976</memory-data-plane>
                    <memory-data-plane-used>566</memory-data-plane-used>
                    <memory-data-plane-util>58</memory-data-plane-util>
                    <cpu-user>6</cpu-user>
                    <cpu-background>0</cpu-background>
                    <cpu-system>6</cpu-system>
                    <cpu-interrupt>0</cpu-interrupt>
                    <cpu-idle>88</cpu-idle>
                    <model>XXXX</model>
                    <serial-number>WWWW</serial-number>
                    <start-time junos:seconds="1554927200">2019-04-10 20:13:20 GMT</start-time>
                    <up-time junos:seconds="21161026">244 days, 22 hours, 3 minutes, 46 seconds</up-time>
                    <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
                    <load-average-one>0.10</load-average-one>
                    <load-average-five>0.10</load-average-five>
                    <load-average-fifteen>0.08</load-average-fifteen>
                </route-engine>
                <route-engine>
                    <status>Failed</status>
                    <temperature junos:celsius="45">45 degrees C / 113 degrees F</temperature>
                    <cpu-temperature junos:celsius="45">45 degrees C / 113 degrees F</cpu-temperature>
                    <memory-system-total>2048</memory-system-total>
                    <memory-system-total-used>1024</memory-system-total-used>
                    <memory-system-total-util>50</memory-system-total-util>
                    <memory-control-plane>1072</memory-control-plane>
                    <memory-control-plane-used>461</memory-control-plane-used>
                    <memory-control-plane-util>43</memory-control-plane-util>
                    <memory-data-plane>976</memory-data-plane>
                    <memory-data-plane-used>566</memory-data-plane-used>
                    <memory-data-plane-util>58</memory-data-plane-util>
                    <cpu-user>6</cpu-user>
                    <cpu-background>0</cpu-background>
                    <cpu-system>6</cpu-system>
                    <cpu-interrupt>0</cpu-interrupt>
                    <cpu-idle>88</cpu-idle>
                    <model>XXXX</model>
                    <serial-number>WWWW</serial-number>
                    <start-time junos:seconds="1554927200">2019-04-10 20:13:20 GMT</start-time>
                    <up-time junos:seconds="21161026">244 days, 22 hours, 3 minutes, 46 seconds</up-time>
                    <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
                    <load-average-one>0.10</load-average-one>
                    <load-average-five>0.10</load-average-five>
                    <load-average-fifteen>0.08</load-average-fifteen>
                </route-engine>
            </route-engine-information>
        </multi-routing-engine-item>
        
        <multi-routing-engine-item>
            
            <re-name>node1</re-name>
            
            <route-engine-information xmlns="http://xml.juniper.net/junos/ZZZZ/junos-chassis">
                <route-engine>
                    <status>OK</status>
                    <temperature junos:celsius="45">45 degrees C / 113 degrees F</temperature>
                    <cpu-temperature junos:celsius="44">44 degrees C / 111 degrees F</cpu-temperature>
                    <memory-system-total>2048</memory-system-total>
                    <memory-system-total-used>1065</memory-system-total-used>
                    <memory-system-total-util>52</memory-system-total-util>
                    <memory-control-plane>1072</memory-control-plane>
                    <memory-control-plane-used>504</memory-control-plane-used>
                    <memory-control-plane-util>47</memory-control-plane-util>
                    <memory-data-plane>976</memory-data-plane>
                    <memory-data-plane-used>566</memory-data-plane-used>
                    <memory-data-plane-util>58</memory-data-plane-util>
                    <cpu-user>25</cpu-user>
                    <cpu-background>0</cpu-background>
                    <cpu-system>26</cpu-system>
                    <cpu-interrupt>0</cpu-interrupt>
                    <cpu-idle>49</cpu-idle>
                    <model>XXXX</model>
                    <serial-number>YYYY</serial-number>
                    <start-time junos:seconds="1554927139">2019-04-10 20:12:19 GMT</start-time>
                    <up-time junos:seconds="21161085">244 days, 22 hours, 4 minutes, 45 seconds</up-time>
                    <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
                    <load-average-one>0.73</load-average-one>
                    <load-average-five>0.81</load-average-five>
                    <load-average-fifteen>0.74</load-average-fifteen>
                </route-engine>
            </route-engine-information>
        </multi-routing-engine-item>
        
    </multi-routing-engine-results>
    <cli>
        <banner>{primary:node1}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	// test routing engine 0
	assert.Equal(t, "node0", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	// test first route engine
	assert.Equal(t, "OK", rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[0].Status, "status")

	assert.Equal(t, float64(42), rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[0].CPUTemperature.Value, "cpu-temperature")

	// test second route engine
	assert.Equal(t, "Failed", rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[1].Status, "status")

	assert.Equal(t, float64(45), rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[1].CPUTemperature.Value, "cpu-temperature")

	// test routing engine 1
	assert.Equal(t, "node1", rpc.MultiRoutingEngineResults.RoutingEngine[1].Name, "re-name")

	// test first route engine
	assert.Equal(t, "OK", rpc.MultiRoutingEngineResults.RoutingEngine[1].RouteEngineInformation.RouteEngines[0].Status, "status")

	assert.Equal(t, float64(44), rpc.MultiRoutingEngineResults.RoutingEngine[1].RouteEngineInformation.RouteEngines[0].CPUTemperature.Value, "cpu-temperature")
}

// Test no multi routing engine
func TestParseOutputNoMultiRE(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZZ/junos">
    <route-engine-information xmlns="http://xml.juniper.net/junos/ZZZZ/junos-chassis">
        <route-engine>
            <slot>0</slot>
            <mastership-state>master</mastership-state>
            <mastership-priority>master</mastership-priority>
            <status>OK</status>
            <temperature junos:celsius="42">42 degrees C / 107 degrees F</temperature>
            <cpu-temperature junos:celsius="40">40 degrees C / 104 degrees F</cpu-temperature>
            <memory-dram-size>32713 MB</memory-dram-size>
            <memory-installed-size>(32768 MB installed)</memory-installed-size>
            <memory-buffer-utilization>10</memory-buffer-utilization>
            <cpu-user>1</cpu-user>
            <cpu-background>0</cpu-background>
            <cpu-system>1</cpu-system>
            <cpu-interrupt>0</cpu-interrupt>
            <cpu-idle>98</cpu-idle>
            <cpu-user1>3</cpu-user1>
            <cpu-background1>0</cpu-background1>
            <cpu-system1>3</cpu-system1>
            <cpu-interrupt1>0</cpu-interrupt1>
            <cpu-idle1>94</cpu-idle1>
            <cpu-user2>8</cpu-user2>
            <cpu-background2>0</cpu-background2>
            <cpu-system2>3</cpu-system2>
            <cpu-interrupt2>0</cpu-interrupt2>
            <cpu-idle2>89</cpu-idle2>
            <cpu-user3>6</cpu-user3>
            <cpu-background3>0</cpu-background3>
            <cpu-system3>3</cpu-system3>
            <cpu-interrupt3>0</cpu-interrupt3>
            <cpu-idle3>91</cpu-idle3>
            <model>XXXX</model>
            <serial-number>YYYY</serial-number>
            <up-time junos:seconds="3095349">35 days, 19 hours, 49 minutes, 9 seconds</up-time>
            <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
            <load-average-one>0.70</load-average-one>
            <load-average-five>0.64</load-average-five>
            <load-average-fifteen>0.48</load-average-fifteen>
        </route-engine>
        <route-engine>
            <slot>1</slot>
            <mastership-state>backup</mastership-state>
            <mastership-priority>backup</mastership-priority>
            <status>OK</status>
            <temperature junos:celsius="42">42 degrees C / 107 degrees F</temperature>
            <cpu-temperature junos:celsius="39">39 degrees C / 102 degrees F</cpu-temperature>
            <memory-dram-size>32713 MB</memory-dram-size>
            <memory-installed-size>(32768 MB installed)</memory-installed-size>
            <memory-buffer-utilization>9</memory-buffer-utilization>
            <cpu-user>0</cpu-user>
            <cpu-background>0</cpu-background>
            <cpu-system>0</cpu-system>
            <cpu-interrupt>0</cpu-interrupt>
            <cpu-idle>99</cpu-idle>
            <model>XXXX</model>
            <serial-number>YYYY</serial-number>
            <up-time junos:seconds="3095344">35 days, 19 hours, 49 minutes, 4 seconds</up-time>
            <last-reboot-reason>Router rebooted after a normal shutdown.</last-reboot-reason>
            <load-average-one>0.59</load-average-one>
            <load-average-five>0.53</load-average-five>
            <load-average-fifteen>0.39</load-average-fifteen>
        </route-engine>
    </route-engine-information>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "N/A", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	// test first route engine
	assert.Equal(t, "0", rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[0].Slot, "slot")

	assert.Equal(t, float64(40), rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[0].CPUTemperature.Value, "cpu-temperature")

	// test second route engine
	assert.Equal(t, "1", rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[1].Slot, "slot")

	assert.Equal(t, uint64(3095344), rpc.MultiRoutingEngineResults.RoutingEngine[0].RouteEngineInformation.RouteEngines[1].UpTime.Seconds, "up-time")

}
