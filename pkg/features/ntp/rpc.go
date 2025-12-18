package ntp

import "encoding/xml"

type rpcReply struct {
	XMLName xml.Name `xml:"rpc-reply"`

	Output struct {
		Text string `xml:",chardata"`
	} `xml:"output"`

	NtpStatus struct {
		AssocID   string `xml:"associd"`
		Status    string `xml:"status"`
		Leap      string `xml:"leap"`
		Stratum   string `xml:"stratum"`
		Precision string `xml:"precision"`
		Rootdelay string `xml:"rootdelay"`
		Refid     string `xml:"refid"`
		Offset    string `xml:"offset"`
		SysJitter string `xml:"sys-jitter"`
		ClkJitter string `xml:"clk-jitter"`
		Tc        string `xml:"tc"`
	} `xml:"ntp-status"`
}
