package environment

type EnvironmentDatasource interface {
	EnvironmentItems() ([]*EnvironmentItem, error)
}
