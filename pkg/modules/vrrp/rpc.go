package vrrp

type result struct {
	Information struct {
		Interfaces []iface `xml:"vrrp-interface"`
	} `xml:"vrrp-information"`
}

type iface struct {
	Interface             string `xml:"interface"`
	InterfaceState        string `xml:"interface-state"`
	Group                 string `xml:"group"`
	VrrpState             string `xml:"vrrp-state"`
	VrrpMode              string `xml:"vrrp-mode"`
	LocalInterfaceAddress string `xml:"local-interface-address"`
	VirtualIPAddress      string `xml:"virtual-ip-address"`
}
