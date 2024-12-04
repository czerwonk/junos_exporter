package arp

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	resultsData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.2R2-S1.3/junos">
    <arp-table-information xmlns="http://xml.juniper.net/junos/23.2R0/junos-arp" junos:style="no-resolve">
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>xe-0/0/5:0.0</interface-name>
            <arp-table-entry-flags>
                <none/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>bme1.0</interface-name>
            <arp-table-entry-flags>
                <permanent/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>bme0.0</interface-name>
            <arp-table-entry-flags>
                <permanent/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>bme2.0</interface-name>
            <arp-table-entry-flags>
                <permanent/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>xe-0/0/3:0.0</interface-name>
            <arp-table-entry-flags>
                <none/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>xe-0/0/3:0.0</interface-name>
            <arp-table-entry-flags>
                <none/>
            </arp-table-entry-flags>
        </arp-table-entry>  
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>fxp0.0</interface-name>
            <arp-table-entry-flags>
                <none/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-table-entry>
            <mac-address>55:55:55:55:55:55</mac-address>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>em1.32768</interface-name>
            <arp-table-entry-flags>
                <none/>
            </arp-table-entry-flags>
        </arp-table-entry>
        <arp-entry-count>7</arp-entry-count>
    </arp-table-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>
`
	var results results
	// Parse the XML data for Interfaces
	err := xml.Unmarshal([]byte(resultsData), &results)
	assert.NoError(t, err)

	// Validate the parsed data
	assert.Len(t, results.ArpTableInformation.ArpTableEntry, 8)

	map_expected := map[string]int64{
		"xe-0/0/5:0.0": 1,
		"bme1.0":       1,
		"bme0.0":       1,
		"bme2.0":       1,
		"xe-0/0/3:0.0": 2,
		"fxp0.0":       1,
		"em1.32768":    1,
	}
	map_in_test := make(map[string]float64)
	for _, a := range results.ArpTableInformation.ArpTableEntry {
		map_in_test[a.InterfaceName] += 1
	}
	assert.Equal(t, len(map_expected), len(map_in_test))
	for key, _ := range map_in_test {
		assert.Equal(t, int64(map_expected[key]), int64(map_in_test[key]))
	}
}
