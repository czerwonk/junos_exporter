package routingengine

type RoutingEngineDatasource interface {
	RouteEngineStats() (*RouteEngineStats, error)
}
