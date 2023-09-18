// SPDX-License-Identifier: MIT

package bgp

import "strings"

func groupForPeer(p peer, groups groupMap) string {
	if len(p.Group) > 0 {
		return p.Group
	}

	return groups[p.GroupIndex].Name
}

func formatPolicy(s string) string {
	return strings.Trim(s, "\n ")
}

func bgpStateToNumber(bgpState string) float64 {
	switch bgpState {
	case "Active":
		return 1
	case "Connect":
		return 2
	case "Established":
		return 3
	case "Idle":
		return 4
	case "Openconfirm":
		return 5
	case "OpenSent":
		return 6
	case "route reflector client":
		return 7
	default:
		return 0
	}
}
