// SPDX-License-Identifier: MIT

package dot1x

type result struct {
	Results struct {
		Interfaces []dot1xInterface `xml:"interface"`
	} `xml:"dot1x-interface-information"`
}

// ProbeTestResults holds the details of a single probe test
type dot1xInterface struct {
	InterfaceName         string `xml:"interface-name"`
	UserMacAddress        string `xml:"user-mac-address"`
	AuthenticatedMethod   string `xml:"authenticated-method"`
	AuthenticatedVlan     int64  `xml:"authenticated-vlan"`
	AuthenticatedVoipVlan int64  `xml:"authenticated-voip-vlan"`
	UserName              string `xml:"user-name"`
	State                 string `xml:"state"`
}
