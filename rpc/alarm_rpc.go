package rpc

type AlarmRpc struct {
	Information struct {
		Details []AlarmDetails `xml:"alarm-detail"`
  } `xml:"alarm-information"`
}

type AlarmDetails struct {
	Class string `xml:"alarm-class"`
}
