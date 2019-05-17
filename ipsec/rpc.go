package ipsec

// IpsecRpc is the root element for xml unmarshalling
type IpsecRpc struct {
	Information struct {
		ActiveTunnels        int                             `xml:"total-active-tunnels"`
		SecurityAssociations []IpsecSecurityAssociationBlock `xml:"ipsec-security-associations-block"`
	} `xml:"ipsec-security-associations-information"`
}

// IpsecSecurityAssociationBlock is used for xml unmarshalling
type IpsecSecurityAssociationBlock struct {
	State                string                     `xml:"sa-block-state"`
	SecurityAssociations []IpsecSecurityAssociation `xml:"ipsec-security-associations"`
}

// IpsecSecurityAssociation is used for xml unmarshalling
type IpsecSecurityAssociation struct {
	Direction              string `xml:"sa-direction"`
	TunnelIndex            int64  `xml:"sa-tunnel-index"`
	Spi                    string `xml:"sa-spi"`
	AuxSpi                 string `xml:"sa-aux-spi"`
	RemoteGateway          string `xml:"sa-remote-gateway"`
	Port                   int    `xml:"sa-port"`
	MonitoringState        string `xml:"sa-vpn-monitoring-state"`
	Protocol               string `xml:"sa-protocol"`
	EspEncryptionAlgorithm string `xml:"sa-esp-encryption-algorithm"`
	HmacAlgorithm          string `xml:"sa-hmac-algorithm"`
	HardLifetime           string `xml:"sa-hard-lifetime"`
	LifesizeRemaining      string `xml:"sa-lifesize-remaining"`
	VirtualSystem          string `xml:"sa-virtual-system"`
}
