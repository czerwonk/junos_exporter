package alarm

type result struct {
	Information struct {
		Details []details `xml:"alarm-detail"`
	} `xml:"alarm-information"`
}

type details struct {
	Class       string `xml:"alarm-class"`
	Description string `xml:"alarm-description"`
	Type        string `xml:"alarm-type"`
}
