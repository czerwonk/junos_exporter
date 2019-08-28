package interfacelabels

type InterfaceRPC struct {
	Information struct {
		Interfaces []PhyInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type PhyInterface struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}
