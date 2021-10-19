package vrrp

type VrrpRpc struct {
	Information struct {
		Interfaces []VrrpInterface `xml:"vrrp-interface"`
	} `xml:"vrrp-information"`
}

type VrrpInterface struct {
	Interface             string `xml:"interface"`
	InterfaceState        string `xml:"interface-state"`
        Group                 string `xml:"group"`
	VrrpState             string `xml:"vrrp-state"`
	VrrpMode              string `xml:"vrrp-mode"`
	LocalInterfaceAddress string `xml:"local-interface-address"`
	VirtualIpAddress      string `xml:"virtual-ip-address"`
}
