package poe

import "encoding/xml"

// structure for poe interface result
type poeInterfaceResult struct {
	XMLName xml.Name `xml:"rpc-reply"`
	Poe     struct {
		InterfaceInformation []InterfaceInformation `xml:"interface-information"`
	} `xml:"poe"`
}

// InterfaceInformation structure for poe information from interface
type InterfaceInformation struct {
	Name                    string `xml:"interface-name"`
	Enabled                 string `xml:"interface-enabled"`
	Status                  string `xml:"interface-status"`
	PowerLimit              string `xml:"interface-power-limit"`
	LldpNegotiationPower    string `xml:"interface-lldp-negotiation-power"`
	Priority                string `xml:"interface-priority"`
	LldpNegotiationPriority string `xml:"interface-lldp-negotiation-priority"`
	Power                   string `xml:"interface-power"`
	Asterisk                string `xml:"interface-asterisk"`
	Class                   string `xml:"interface-class"`
}
