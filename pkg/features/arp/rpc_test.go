package arp

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestARP(t *testing.T) {
	resultsData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.2R2-S1.3/junos">
    <arp-table-information xmlns="http://xml.juniper.net/junos/23.2R0/junos-arp" junos:style="no-resolve">
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>xe-0/0/5:0.0</interface-name>
        </arp-table-entry>
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>bme1.0</interface-name>
        </arp-table-entry>
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>bme0.0</interface-name>
        </arp-table-entry>
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>bme2.0</interface-name>
        </arp-table-entry>
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>xe-0/0/3:0.0</interface-name>
        </arp-table-entry>
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>xe-0/0/3:0.0</interface-name>
        </arp-table-entry>  
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>fxp0.0</interface-name>
        </arp-table-entry>
        <arp-table-entry>
            <ip-address>55.55.55.55</ip-address>
            <interface-name>em1.32768</interface-name>
        </arp-table-entry>
        <arp-entry-count>7</arp-entry-count>
    </arp-table-information>
</rpc-reply>
`
	var results results
	// Parse the XML data for ARP
	err := xml.Unmarshal([]byte(resultsData), &results)
	assert.NoError(t, err)

	assert.Len(t, results.ArpTableInformation.ArpTableEntry, 8)

	expected := map[string]int64{
		"xe-0/0/5:0.0": 1,
		"bme1.0":       1,
		"bme0.0":       1,
		"bme2.0":       1,
		"xe-0/0/3:0.0": 2,
		"fxp0.0":       1,
		"em1.32768":    1,
	}
	inTest := make(map[string]float64)
	for _, a := range results.ArpTableInformation.ArpTableEntry {
		inTest[a.InterfaceName] += 1
	}
	assert.Equal(t, len(expected), len(inTest))
	for key, _ := range inTest {
		assert.Equal(t, int64(expected[key]), int64(inTest[key]))
	}
}
