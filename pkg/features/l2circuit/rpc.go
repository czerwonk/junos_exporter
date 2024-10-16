// SPDX-License-Identifier: MIT

package l2circuit

type result struct {
	Information neighbors `xml:"l2circuit-connection-information"`
}

type neighbors struct {
	Neighbors []neighbor `xml:"l2circuit-neighbor"`
}

type neighbor struct {
	Address     string       `xml:"neighbor-address"`
	Connections []connection `xml:"connection"`
}

type connection struct {
	ID           string `xml:"connection-id"`
	Type         string `xml:"connection-type"`
	StatusString string `xml:"connection-status"`
    LocalInterface localInterface  `xml:"local-interface"`
}
type localInterface struct {
	Name        string `xml:"interface-name"`
	Description string `xml:"interface-description"`
}
