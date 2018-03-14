package routing_engine

type RoutingEngineDatasource interface {
	RouteEngineStats() (*RouteEngineStats, error)
}
