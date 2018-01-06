package route

type RoutesDatasource interface {
	RoutingTables() ([]*RoutingTable, error)
}
