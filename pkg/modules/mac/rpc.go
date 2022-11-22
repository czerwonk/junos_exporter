package mac

type result struct {
	Information ethernetSwitchingTableInformation `xml:"ethernet-switching-table-information"`
}

type ethernetSwitchingTableInformation struct {
	Table ethernetSwitchingTable `xml:"ethernet-switching-table"`
}

type ethernetSwitchingTable struct {
	Entry macTableEntry `xml:"mac-table-entry"`
}

type macTableEntry struct {
	TotalCount   int64 `xml:"mac-table-total-count"`
	ReceiveCount int64 `xml:"mac-table-recieve-count"`
	DynamicCount int64 `xml:"mac-table-dynamic-count"`
	FloodCount   int64 `xml:"mac-table-flood-count"`
}
