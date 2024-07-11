package l2vpn

type l2vpnRpc struct {
	Information l2vpnInformation `xml:"l2vpn-connection-information"`
}

type l2vpnInformation struct {
	Instances []l2vpnInstance `xml:"instance"`
}

type l2vpnInstance struct {
	InstanceName  string               `xml:"instance-name"`
	ReferenceSite []l2vpnReferenceSite `xml:"reference-site"`
}

type l2vpnReferenceSite struct {
	ID          string            `xml:"local-site-id"`
	Connections []l2vpnConnection `xml:"connection"`
}

type l2vpnConnection struct {
	ID             string           `xml:"connection-id"`
	Type           string           `xml:"connection-type"`
	StatusString   string           `xml:"connection-status"`
	RemotePe       string           `xml:"remote-pe"`
	LastChange     string           `xml:"last-change"`
	UpTransitions  string           `xml:"up-transitions"`
	LocalInterface []l2vpnInterface `xml:"local-interface"`
}

type l2vpnInterface struct {
	Name string `xml:"interface-name"`
}
