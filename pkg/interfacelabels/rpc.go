// SPDX-License-Identifier: MIT

package interfacelabels

import (
	"encoding/xml"
)

type result struct {
	XMLName    xml.Name       `xml:"interface-information"`
	Interfaces []phyInterface `xml:"physical-interface"`
}

type phyInterface struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
}
