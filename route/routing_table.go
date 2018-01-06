package route

type RoutingTable struct {
	Name         string
	TotalRoutes  float64
	ActiveRoutes float64
	MaxRoutes    float64
	Protocols    []*ProtocolRouteCount
}

type ProtocolRouteCount struct {
	Name         string
	Routes       float64
	ActiveRoutes float64
}
