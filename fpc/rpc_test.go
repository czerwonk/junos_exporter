package fpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEXOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/15.1R7/junos">
    <fpc-information xmlns="http://xml.juniper.net/junos/15.1R7/junos-chassis" junos:style="verbose">
        <fpc>
            <slot>0</slot>
            <state>Online</state>
            <memory-dram-size>1024</memory-dram-size>
            <start-time junos:seconds="1540974142">2018-10-31 08:22:22 UTC</start-time>
            <up-time junos:seconds="18847524">218 days, 3 hours, 25 minutes, 24 seconds</up-time>
        </fpc>
    </fpc-information>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs)
	assert.Equal(t, "N/A", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC[0]
	assert.Equal(t, 0, f.Slot, "slot")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, uint(1024), f.MemoryDramSize, "memory-dram-size")
	assert.Equal(t, uint64(1540974142), f.StartTime.Seconds, "start-time")
	assert.Equal(t, uint64(18847524), f.UpTime.Seconds, "up-time")
}

func TestParseQFXOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R1/junos">
    <fpc-information xmlns="http://xml.juniper.net/junos/18.4R1/junos-chassis" junos:style="verbose">
        <fpc>
            <slot>0</slot>
            <state>Online</state>
            <temperature junos:celsius="39">39 degrees C / 102 degrees F</temperature>
            <memory-dram-size>16384</memory-dram-size>
            <memory-sram-size>1024</memory-sram-size>
            <memory-sdram-size>2048</memory-sdram-size>
            <start-time junos:seconds="1559804691">2019-06-06 07:04:51 UTC</start-time>
            <up-time junos:seconds="16737">4 hours, 38 minutes, 57 seconds</up-time>
        </fpc>
    </fpc-information>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs)
	assert.Equal(t, "N/A", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC[0]

	assert.Equal(t, 0, f.Slot, "slot")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, uint(16384), f.MemoryDramSize, "memory-dram-size")
	assert.Equal(t, uint(2048), f.MemorySdramSize, "memory-sdram-size")
	assert.Equal(t, uint(1024), f.MemorySramSize, "memory-sram-size")
	assert.Equal(t, uint64(1559804691), f.StartTime.Seconds, "start-time")
	assert.Equal(t, uint64(16737), f.UpTime.Seconds, "up-time")
}

func TestParseMXOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/17.3R3/junos">
    <fpc-information xmlns="http://xml.juniper.net/junos/17.3R3/junos-chassis" junos:style="verbose">
        <fpc>
            <slot>1</slot>
            <state>Online</state>
            <temperature junos:celsius="30">30 degrees C / 86 degrees F</temperature>
            <memory-dram-size>2048</memory-dram-size>
            <memory-rldram-size>1036</memory-rldram-size>
            <memory-ddr-dram-size>11264</memory-ddr-dram-size>
            <start-time junos:seconds="1550101033">2019-02-13 23:37:13 UTC</start-time>
            <up-time junos:seconds="9720385">112 days, 12 hours, 6 minutes, 25 seconds</up-time>
            <max-power-consumption>584</max-power-consumption>
        </fpc>
        <fpc>
            <slot>2</slot>
            <state>Online</state>
            <temperature junos:celsius="30">30 degrees C / 86 degrees F</temperature>
            <memory-dram-size>2048</memory-dram-size>
            <memory-rldram-size>1036</memory-rldram-size>
            <memory-ddr-dram-size>11264</memory-ddr-dram-size>
            <start-time junos:seconds="1550101046">2019-02-13 23:37:26 UTC</start-time>
            <up-time junos:seconds="9720372">112 days, 12 hours, 6 minutes, 12 seconds</up-time>
            <max-power-consumption>584</max-power-consumption>
		</fpc>
	</fpc-information>
	<cli>
		<banner>{master:0}</banner>
	</cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs)
	assert.Equal(t, "N/A", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC[1]

	assert.Equal(t, 2, f.Slot, "slot")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, 30, f.Temperature.Celsius, "temperature")
	assert.Equal(t, uint(2048), f.MemoryDramSize, "memory-dram-size")
	assert.Equal(t, uint(1036), f.MemoryRldramSize, "memory-rl-dram-size")
	assert.Equal(t, uint(11264), f.MemoryDdrDramSize, "memory-ddr-dram-size")
	assert.Equal(t, uint64(1550101046), f.StartTime.Seconds, "start-time")
	assert.Equal(t, uint64(9720372), f.UpTime.Seconds, "up-time")
	assert.Equal(t, uint(584), f.MaxPowerConsumption, "max-power-consumption")
}

// Test multi routing engine also
func TestParseSRXOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R3/junos">
    <multi-routing-engine-results>
        
        <multi-routing-engine-item>
            
            <re-name>node0</re-name>
            
            <fpc-information xmlns="http://xml.juniper.net/junos/18.4R3/junos-chassis" junos:style="verbose">
                <fpc>
                    <slot>0</slot>
                    <state>Online</state>
                    <memory-dram-size>2048</memory-dram-size>
                    <memory-sram-size>0</memory-sram-size>
                    <memory-sdram-size>0</memory-sdram-size>
                    <temperature junos:celsius="37">37 degrees C / 98 degrees F</temperature>
                    <start-time junos:seconds="1598014826">2020-08-21 13:00:26 UTC</start-time>
                    <up-time junos:seconds="18494181">214 days, 1 hour, 16 minutes, 21 seconds</up-time>
                </fpc>
                <fpc>
                    <slot>0</slot>
                    <state>Online</state>
                    <memory-dram-size>2048</memory-dram-size>
                    <memory-sram-size>0</memory-sram-size>
                    <memory-sdram-size>0</memory-sdram-size>
                    <temperature junos:celsius="38">38 degrees C / 98 degrees F</temperature>
                    <start-time junos:seconds="1598014826">2020-08-21 13:00:26 UTC</start-time>
                    <up-time junos:seconds="18494181">214 days, 1 hour, 16 minutes, 21 seconds</up-time>
                </fpc>
            </fpc-information>
        </multi-routing-engine-item>
        
        <multi-routing-engine-item>
            
            <re-name>node1</re-name>
            
            <fpc-information xmlns="http://xml.juniper.net/junos/18.4R3/junos-chassis" junos:style="verbose">
                <fpc>
                    <slot>0</slot>
                    <state>Online</state>
                    <memory-dram-size>2048</memory-dram-size>
                    <memory-sram-size>0</memory-sram-size>
                    <memory-sdram-size>0</memory-sdram-size>
                    <temperature junos:celsius="38">38 degrees C / 100 degrees F</temperature>
                    <start-time junos:seconds="1598015225">2020-08-21 13:07:05 UTC</start-time>
                    <up-time junos:seconds="18493783">214 days, 1 hour, 9 minutes, 43 seconds</up-time>
                </fpc>
            </fpc-information>
        </multi-routing-engine-item>
        
    </multi-routing-engine-results>
    <cli>
        <banner>{primary:node0}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC)

	assert.Equal(t, "node0", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")
	assert.Equal(t, "node1", rpc.MultiRoutingEngineResults.RoutingEngine[1].Name, "re-name")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC[1]

	assert.Equal(t, 0, f.Slot, "slot")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, 38, f.Temperature.Celsius, "temperature")
	assert.Equal(t, uint(2048), f.MemoryDramSize, "memory-dram-size")
	assert.Equal(t, uint(0), f.MemorySramSize, "memory-rl-dram-size")
	assert.Equal(t, uint(0), f.MemorySdramSize, "memory-ddr-dram-size")
	assert.Equal(t, uint64(1598014826), f.StartTime.Seconds, "start-time")
	assert.Equal(t, uint64(18494181), f.UpTime.Seconds, "up-time")
}

func TestParseMXPicOutput(t *testing.T) {

	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/17.3R3/junos">
    <fpc-information xmlns="http://xml.juniper.net/junos/17.3R3/junos-chassis" junos:style="pic-style">
        <fpc>
            <slot>0</slot>
            <state>Online</state>
            <description>Desc1</description>
            <pic>
                <pic-slot>2</pic-slot>
                <pic-state>Online</pic-state>
                <pic-type>10X10GE SFPP</pic-type>
            </pic>
        </fpc>
        <fpc>
            <slot>1</slot>
            <state>Online</state>
            <description>Desc2</description>
            <pic>
                <pic-slot>2</pic-slot>
                <pic-state>Online</pic-state>
                <pic-type>10X10GE SFPP</pic-type>
            </pic>
        </fpc>
        <fpc>
            <slot>3</slot>
            <state>Online</state>
            <description>Desc3</description>
            <pic>
                <pic-slot>0</pic-slot>
                <pic-state>Online</pic-state>
                <pic-type>MRATE-6xQSFPP-XGE-XLGE-CGE</pic-type>
            </pic>
            <pic>
                <pic-slot>1</pic-slot>
                <pic-state>Online</pic-state>
                <pic-type>MRATE-6xQSFPP-XGE-XLGE-CGE</pic-type>
            </pic>
        </fpc>
    </fpc-information>
    <cli>
        <banner>{master}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC)

	assert.Equal(t, "N/A", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC[2]

	assert.Equal(t, 3, f.Slot, "slot")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, "Desc3", f.Description, "description")

	p := f.Pics[1]

	assert.Equal(t, 1, p.PicSlot, "pic-slot")
	assert.Equal(t, "Online", p.PicState, "pic-state")
	assert.Equal(t, "MRATE-6xQSFPP-XGE-XLGE-CGE", p.PicType, "pic-type")
}

// Test multi routing engine also
func TestParseSRXPicOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R3/junos">
    <multi-routing-engine-results>
        <multi-routing-engine-item>
            <re-name>node0</re-name>
            <fpc-information xmlns="http://xml.juniper.net/junos/18.4R3/junos-chassis" junos:style="pic-style">
                <fpc>
                    <slot>0</slot>
                    <state>Online</state>
                    <description>FEB</description>
                    <pic>
                        <pic-slot>0</pic-slot>
                        <pic-state>Online</pic-state>
                        <pic-type>12x1G-T-4x1G-SFP-4x10G</pic-type>
                    </pic>
                    <pic>
                        <pic-slot>1</pic-slot>
                        <pic-state>Online</pic-state>
                        <pic-type>12x1G-T-4x1G-SFP-4x10G</pic-type>
                    </pic>
                </fpc>
                <fpc>
                    <slot>1</slot>
                    <state>Online</state>
                    <description>FEB</description>
                    <pic>
                        <pic-slot>0</pic-slot>
                        <pic-state>Online</pic-state>
                        <pic-type>N/A</pic-type>
                    </pic>
                </fpc>
            </fpc-information>
        </multi-routing-engine-item>
        <multi-routing-engine-item>
            <re-name>node1</re-name>
            <fpc-information xmlns="http://xml.juniper.net/junos/18.4R3/junos-chassis" junos:style="pic-style">
                <fpc>
                    <slot>0</slot>
                    <state>Online</state>
                    <description>FEB</description>
                    <pic>
                        <pic-slot>0</pic-slot>
                        <pic-state>Online</pic-state>
                        <pic-type>12x1G-T-4x1G-SFP-4x10G</pic-type>
                    </pic>
                </fpc>
            </fpc-information>
        </multi-routing-engine-item>
    </multi-routing-engine-results>
    <cli>
        <banner>{primary:node0}</banner>
    </cli>
</rpc-reply>`

	rpc := RpcReply{}
	err := parseXML([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC)

	assert.Equal(t, "node0", rpc.MultiRoutingEngineResults.RoutingEngine[0].Name, "re-name")
	assert.Equal(t, "node1", rpc.MultiRoutingEngineResults.RoutingEngine[1].Name, "re-name")

	f := rpc.MultiRoutingEngineResults.RoutingEngine[0].FPCs.FPC[1]

	assert.Equal(t, 1, f.Slot, "slot")
	assert.Equal(t, "Online", f.State, "state")
	assert.Equal(t, "FEB", f.Description, "description")

	p := f.Pics[0]

	assert.Equal(t, 0, p.PicSlot, "pic-slot")
	assert.Equal(t, "Online", p.PicState, "pic-state")
	assert.Equal(t, "N/A", p.PicType, "pic-type")
}
