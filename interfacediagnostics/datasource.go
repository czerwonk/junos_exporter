package interfacediagnostics

type InterfaceDiagnosticsDatasource interface {
	InterfaceDiagnostics() ([]*InterfaceDiagnostics, error)
}
