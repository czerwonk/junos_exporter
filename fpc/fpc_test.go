package fpc

import (
	"encoding/xml"
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

	r := FPCRpc{}
	err := xml.Unmarshal([]byte(body), &r)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, r.Information)

	f := r.Information.FPCs[0]
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

	r := FPCRpc{}
	err := xml.Unmarshal([]byte(body), &r)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, r.Information)

	f := r.Information.FPCs[0]
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

	r := FPCRpc{}
	err := xml.Unmarshal([]byte(body), &r)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, r.Information)

	f := r.Information.FPCs[1]
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
