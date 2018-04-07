package alarm

type AlarmRpc struct {
	Information struct {
		Details []AlarmDetails `xml:"alarm-detail"`
	} `xml:"alarm-information"`
}

type AlarmDetails struct {
	Class       string `xml:"alarm-class"`
	Description string `xml:"alarm-description"`
	Type        string `xml:"alarm-type"`
}
