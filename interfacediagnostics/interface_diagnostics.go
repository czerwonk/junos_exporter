package interfacediagnostics

type InterfaceDiagnostics struct {
	Index               string
	Name                string
	LaserBiasCurrent    float64
	LaserOutputPower    float64
	LaserOutputPowerDbm float64
	ModuleTemperature   float64

	LaserRxOpticalPower    float64
	LaserRxOpticalPowerDbm float64

	ModuleVoltage              float64
	RxSignalAvgOpticalPower    float64
	RxSignalAvgOpticalPowerDbm float64

	Lanes []*InterfaceDiagnostics
}
