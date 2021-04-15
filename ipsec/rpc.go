package ipsec

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name  string   `xml:"re-name"`
	IpSec IpSecRpc `xml:"ipsec-security-associations-information"`
}

type IpSecRpc struct {
	ActiveTunnels        int                             `xml:"total-active-tunnels"`
	SecurityAssociations []IpsecSecurityAssociationBlock `xml:"ipsec-security-associations-block"`
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

type RpcReplyNoRE struct {
	XMLName xml.Name `xml:"rpc-reply"`
	IpSec   IpSecRpc `xml:"ipsec-security-associations-information"`
}

// ConfigurationSecurityIpsec is used for xml unmarshalling
// In order to get the number of configured VPNs
type ConfigurationSecurityIpsec struct {
	Configuration struct {
		Security struct {
			Ipsec struct {
				Proposal struct {
					Text                    string `xml:",chardata"`
					Name                    string `xml:"name"`
					Protocol                string `xml:"protocol"`
					AuthenticationAlgorithm string `xml:"authentication-algorithm"`
					EncryptionAlgorithm     string `xml:"encryption-algorithm"`
					LifetimeSeconds         string `xml:"lifetime-seconds"`
				} `xml:"proposal"`
				Policy struct {
					Name      string `xml:"name"`
					Proposals string `xml:"proposals"`
				} `xml:"policy"`
				Vpn []struct {
					Name          string `xml:"name"`
					BindInterface string `xml:"bind-interface"`
					Ike           struct {
						Gateway     string `xml:"gateway"`
						IpsecPolicy string `xml:"ipsec-policy"`
					} `xml:"ike"`
					EstablishTunnels string `xml:"establish-tunnels"`
				} `xml:"vpn"`
			} `xml:"ipsec"`
		} `xml:"security"`
	} `xml:"configuration"`
}
