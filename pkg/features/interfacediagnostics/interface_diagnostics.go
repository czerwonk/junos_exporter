package interfacediagnostics

type interfaceDiagnostics struct {
	Index                              string
	Name                               string
	LaserBiasCurrent                   float64
	LaserBiasCurrentHighAlarmThreshold float64
	LaserBiasCurrentLowAlarmThreshold  float64
	LaserBiasCurrentHighWarnThreshold  float64
	LaserBiasCurrentLowWarnThreshold   float64

	LaserOutputPower                   float64
	LaserOutputPowerHighAlarmThreshold float64
	LaserOutputPowerLowAlarmThreshold  float64
	LaserOutputPowerHighWarnThreshold  float64
	LaserOutputPowerLowWarnThreshold   float64

	LaserOutputPowerDbm                   float64
	LaserOutputPowerHighAlarmThresholdDbm float64
	LaserOutputPowerLowAlarmThresholdDbm  float64
	LaserOutputPowerHighWarnThresholdDbm  float64
	LaserOutputPowerLowWarnThresholdDbm   float64

	ModuleTemperature                   float64
	ModuleTemperatureHighAlarmThreshold float64
	ModuleTemperatureLowAlarmThreshold  float64
	ModuleTemperatureHighWarnThreshold  float64
	ModuleTemperatureLowWarnThreshold   float64

	LaserRxOpticalPower                   float64
	LaserRxOpticalPowerHighAlarmThreshold float64
	LaserRxOpticalPowerLowAlarmThreshold  float64
	LaserRxOpticalPowerHighWarnThreshold  float64
	LaserRxOpticalPowerLowWarnThreshold   float64

	LaserRxOpticalPowerDbm                   float64
	LaserRxOpticalPowerHighAlarmThresholdDbm float64
	LaserRxOpticalPowerLowAlarmThresholdDbm  float64
	LaserRxOpticalPowerHighWarnThresholdDbm  float64
	LaserRxOpticalPowerLowWarnThresholdDbm   float64

	ModuleVoltage                   float64
	ModuleVoltageHighAlarmThreshold float64
	ModuleVoltageLowAlarmThreshold  float64
	ModuleVoltageHighWarnThreshold  float64
	ModuleVoltageLowWarnThreshold   float64
	RxSignalAvgOpticalPower         float64
	RxSignalAvgOpticalPowerDbm      float64

	Lanes []*interfaceDiagnostics
}
