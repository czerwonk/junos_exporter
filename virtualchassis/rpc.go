package virtualchassis

type virtualChassisRpc struct {
	VirtualChassisInformation struct {
		VirtualChassisIdInformation struct {
			VirtualChassisId        string `xml:"virtual-chassis-id"`
			VirtualChassisMode      string `xml:"virtual-chassis-mode"`
		} `xml:"virtual-chassis-id-information"`
		MemberList struct {
			Member             []vcmembers `xml:"member"`
		} `xml:"member-list"`
	} `xml:"virtual-chassis-information"`
}

type vcmembers struct {
	Status         string `xml:"member-status"`
	Id             string `xml:"member-id"`
	FpcSlot        string `xml:"fpc-slot"`
	SerialNumber   string `xml:"member-serial-number"`
	Model          string `xml:"member-model"`
	Role           string `xml:"member-role"`
}
