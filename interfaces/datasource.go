package interfaces

type InterfaceStatsDatasource interface {
	InterfaceStats() ([]*InterfaceStats, error)
}
