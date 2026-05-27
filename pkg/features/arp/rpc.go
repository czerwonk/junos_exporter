package arp

type results struct {
	Text                string `xml:",chardata"`
	ArpTableInformation struct {
		Text          string `xml:",chardata"`
		ArpTableEntry []struct {
			InterfaceName      string `xml:"interface-name"`
			ArpTableEntryFlags struct {
				Text string `xml:",chardata"`
			} `xml:"arp-table-entry-flags"`
		} `xml:"arp-table-entry"`
		ArpEntryCount string `xml:"arp-entry-count"`
	} `xml:"arp-table-information"`
}
