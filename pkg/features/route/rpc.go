// SPDX-License-Identifier: MIT

package route

type result struct {
	Information struct {
		Tables []routeTable `xml:"route-table"`
	} `xml:"route-summary-information"`
}

type routeTable struct {
	Name         string               `xml:"table-name"`
	MaxRoutes    int64                `xml:"prefix-max"`
	TotalRoutes  int64                `xml:"total-route-count"`
	ActiveRoutes int64                `xml:"active-route-count"`
	Protocols    []routeTableProtocol `xml:"protocols"`
}

type routeTableProtocol struct {
	Name         string `xml:"protocol-name"`
	Routes       int64  `xml:"protocol-route-count"`
	ActiveRoutes int64  `xml:"active-route-count"`
}
