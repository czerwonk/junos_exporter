// SPDX-License-Identifier: MIT

package lldp

type result struct {
	Information struct {
		Neighbors []neighbor `xml:"lldp-neighbor-information"`
	} `xml:"lldp-neighbors-information"`
}

type neighbor struct {
	LocalPortID              string `xml:"lldp-local-port-id"`
	LocalParentInterfaceName string `xml:"lldp-local-parent-interface-name"`
	RemoteChassisID          string `xml:"lldp-remote-chassis-id"`
	RemotePortID             string `xml:"lldp-remote-port-id"`
	RemoteSystemName         string `xml:"lldp-remote-system-name"`
}

type localResult struct {
	Information struct {
		LocalInfo []localInfo `xml:"lldp-local-info"`
	} `xml:"lldp-local-info"`
}

type localInfo struct {
	LocalInterfaces []localInterface `xml:"lldp-local-interface-info"`
}

type localInterface struct {
	InterfaceName        string `xml:"lldp-local-interface-name"`
	ParentInterfaceName  string `xml:"lldp-parent-local-interface-name"`
	InterfaceID          string `xml:"lldp-local-interface-id"`
	InterfaceDescription string `xml:"lldp-local-interface-description"`
	InterfaceStatus      string `xml:"lldp-local-interface-status"`
} 
