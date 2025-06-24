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
