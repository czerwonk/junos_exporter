// SPDX-License-Identifier: MIT

package mac

type oldResult struct {
	Information oldEthernetSwitchingTableInformation `xml:"ethernet-switching-table-information"`
}

type oldEthernetSwitchingTableInformation struct {
	Table oldEthernetSwitchingTable `xml:"ethernet-switching-table"`
}

type oldEthernetSwitchingTable struct {
	Entry oldMacTableEntry `xml:"mac-table-entry"`
}

type oldMacTableEntry struct {
	TotalCount   int64 `xml:"mac-table-total-count"`
	ReceiveCount int64 `xml:"mac-table-recieve-count"`
	DynamicCount int64 `xml:"mac-table-dynamic-count"`
	FloodCount   int64 `xml:"mac-table-flood-count"`
}

type newResult struct {
	Macdb newL2ngMacdb `xml:"l2ng-l2ald-rtb-macdb"`
}

type newL2ngMacdb struct {
	TableSummary newL2ngTableSummary `xml:"l2ng-l2ald-ethernet-switching-table-summary"`
}

type newL2ngTableSummary struct {
	TotalMacCount  int64 `xml:"l2ng-l2-total-mac-count"`
	TotalSMacCount int64 `xml:"l2ng-l2-total-smac-count"`
}
