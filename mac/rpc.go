package mac

type MacRpc struct {
	Information EthernetSwitchingTableInformation `xml:"ethernet-switching-table-information"`
}

type EthernetSwitchingTableInformation struct {
	Table EthernetSwitchingTable `xml:"ethernet-switching-table"`
}

type EthernetSwitchingTable struct {
	Entry MacTableEntry `xml:"mac-table-entry"`
}

type MacTableEntry struct {
	TotalCount   int64 `xml:"mac-table-total-count"`
	ReceiveCount int64 `xml:"mac-table-recieve-count"`
	DynamicCount int64 `xml:"mac-table-dynamic-count"`
	FloodCount   int64 `xml:"mac-table-flood-count"`
}

