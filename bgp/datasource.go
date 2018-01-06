package bgp

type BgpDatasource interface {
	BgpSessions() ([]*BgpSession, error)
}
