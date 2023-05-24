// SPDX-License-Identifier: MIT

package interfacelabels

type result struct {
	Information struct {
		Interfaces []phyInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type phyInterface struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}
