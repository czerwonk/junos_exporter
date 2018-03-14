package interface_diagnostics

type InterfaceDiagnosticsDatasource interface {
	InterfaceDiagnostics() ([]*InterfaceDiagnostics, error)
}
