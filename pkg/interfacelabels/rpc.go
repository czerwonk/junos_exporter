// SPDX-License-Identifier: MIT

package interfacelabels

type result struct {
	Information struct {
		PhysicalInterfaces []interfaceDescription `xml:"physical-interface"`
		LogicalInterfaces  []interfaceDescription `xml:"logical-interface"`
	} `xml:"interface-information"`
}

type interfaceDescription struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}
