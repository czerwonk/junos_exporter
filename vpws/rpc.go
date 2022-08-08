package vpws

type vpwsRpc struct {
	Information struct {
		VpwsInstances []vpwsInstance `xml:"evpn-vpws-instance"`
	} `xml:"evpn-vpws-information"`
}

type vpwsInstance struct {
	Name              string `xml:"evpn-vpws-instance-name"`
	RD                string `xml:"route-distinguisher"`
	LocalInterfaces   int64  `xml:"local-interfaces"`
	LocalInterfacesUp int64  `xml:"local-interfaces-up"`

	VpwsInterfaces []vpwsInterface `xml:"evpn-vpws-interface-status-table>evpn-vpws-interface"`
}

type vpwsInterface struct {
	Name   string `xml:"evpn-vpws-interface-name"`
	Esi    string `xml:"evpn-vpws-interface-esi"`
	Mode   string `xml:"evpn-vpws-interface-mode"`
	Role   string `xml:"evpn-vpws-interface-role"`
	Status string `xml:"evpn-vpws-interface-status"`

	LocalStatus struct {
		Sid       string          `xml:"evpn-vpws-sid-local-value"`
		SidPeInfo []vpwsSidPeInfo `xml:"evpn-vpws-sid-pe-status-table>evpn-vpws-sid-pe-info"`
	} `xml:"evpn-vpws-service-id-local-status-table>evpn-vpws-sid-local"`

	RemoteStatus struct {
		Sid                      string `xml:"evpn-vpws-sid-remote-value"`
		LocalInterfaceName       string `xml:"evpn-vpws-sid-local-interface-name"`
		LocalInterfaceStatus     string `xml:"evpn-vpws-sid-local-interface-status"`
		SidPeInfo        []vpwsSidPeInfo `xml:"evpn-vpws-sid-pe-status-table>evpn-vpws-sid-pe-info"`
	} `xml:"evpn-vpws-service-id-remote-status-table>evpn-vpws-sid-remote"`
}

type vpwsSidPeInfo struct {
	Esi    string `xml:"evpn-vpws-sid-interface-esi"`
	IP     string `xml:"evpn-vpws-sid-pe-ipaddr"`
	Mode   string `xml:"evpn-vpws-sid-pe-mode"`
	Role   string `xml:"evpn-vpws-sid-pe-role"`
	Status string `xml:"evpn-vpws-sid-pe-status"`
}
