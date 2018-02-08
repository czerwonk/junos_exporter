package interface_diagnostics

type InterfaceDiagnostics struct {
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
}
