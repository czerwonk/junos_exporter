package ospf

type OspfDatasource interface {
	OspfAreas() ([]*OspfArea, error)
}
