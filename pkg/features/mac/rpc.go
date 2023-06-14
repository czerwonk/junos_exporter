// SPDX-License-Identifier: MIT

package mac

import (
	"encoding/xml"
)

type ethernetSwitchingTableInformation struct {
	XMLName xml.Name               `xml:"ethernet-switching-table-information"`
	Table   ethernetSwitchingTable `xml:"ethernet-switching-table"`
}

type ethernetSwitchingTable struct {
	Entry macTableEntry `xml:"mac-table-entry"`
}

type macTableEntry struct {
	TotalCount int64 `xml:"mac-table-total-count"`
	Dot1XCount int64 `xml:"mac-table-dot1x-count"`
	// JunOS actually returns recieve
	ReceiveCount int64 `xml:"mac-table-recieve-count"`
	DynamicCount int64 `xml:"mac-table-dynamic-count"`
	FloodCount   int64 `xml:"mac-table-flood-count"`
}
