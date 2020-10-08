package interfacediagnostics

type InterfaceDiagnosticsRPC struct {
	Information struct {
		Diagnostics []PhyDiagInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type PhyDiagInterface struct {
	Name        string                 `xml:"name"`
	Diagnostics PhyInterfaceDiagnostic `xml:"optics-diagnostics,omitempty"`
}

type PhyInterfaceDiagnostic struct {
	LaserBiasCurrent                   float64 `xml:"laser-bias-current,omitempty"`
	LaserBiasCurrentHighAlarmThreshold float64 `xml:"laser-bias-current-high-alarm-threshold,omitempty"`
	LaserBiasCurrentLowAlarmThreshold  float64 `xml:"laser-bias-current-low-alarm-threshold,omitempty"`
	LaserBiasCurrentHighWarnThreshold  float64 `xml:"laser-bias-current-high-warn-threshold,omitempty"`
	LaserBiasCurrentLowWarnThreshold   float64 `xml:"laser-bias-current-low-warn-threshold,omitempty"`

	LaserOutputPower                    float64     `xml:"laser-output-power,omitempty"`
	LaserOutputPowerDbm                 string      `xml:"laser-output-power-dbm,omitempty"`
	ModuleTemperature                   Temperature `xml:"module-temperature"`
	ModuleTemperatureHighAlarmThreshold Temperature `xml:"module-temperature-high-alarm-threshold,omitempty"`
	ModuleTemperatureLowAlarmThreshold  Temperature `xml:"module-temperature-low-alarm-threshold,omitempty"`
	ModuleTemperatureHighWarnThreshold  Temperature `xml:"module-temperature-high-warn-threshold,omitempty"`
	ModuleTemperatureLowWarnThreshold   Temperature `xml:"module-temperature-low-warn-threshold,omitempty"`

	ModuleVoltage                   float64 `xml:"module-voltage,omitempty"`
	ModuleVoltageHighAlarmThreshold float64 `xml:"module-voltage-high-alarm-threshold,omitempty"`
	ModuleVoltageLowAlarmThreshold  float64 `xml:"module-voltage-low-alarm-threshold,omitempty"`
	ModuleVoltageHighWarnThreshold  float64 `xml:"module-voltage-high-warn-threshold,omitempty"`
	ModuleVoltageLowWarnThreshold   float64 `xml:"module-voltage-low-warn-threshold,omitempty"`

	RxSignalAvgOpticalPower    float64 `xml:"rx-signal-avg-optical-power,omitempty"`
	RxSignalAvgOpticalPowerDbm string  `xml:"rx-signal-avg-optical-power-dbm,omitempty"`

	LaserRxOpticalPower                   float64 `xml:"laser-rx-optical-power,omitempty"`
	LaserRxOpticalPowerHighAlarmThreshold float64 `xml:"laser-rx-power-high-alarm-threshold,omitempty"`
	LaserRxOpticalPowerLowAlarmThreshold  float64 `xml:"laser-rx-power-low-alarm-threshold,omitempty"`
	LaserRxOpticalPowerHighWarnThreshold  float64 `xml:"laser-rx-power-high-warn-threshold,omitempty"`
	LaserRxOpticalPowerLowWarnThreshold   float64 `xml:"laser-rx-power-low-warn-threshold,omitempty"`

	LaserRxOpticalPowerDbm                   string `xml:"laser-rx-optical-power-dbm,omitempty"`
	LaserRxOpticalPowerHighAlarmThresholdDbm string `xml:"laser-rx-power-high-alarm-threshold-dbm,omitempty"`
	LaserRxOpticalPowerLowAlarmThresholdDbm  string `xml:"laser-rx-power-low-alarm-threshold-dbm,omitempty"`
	LaserRxOpticalPowerHighWarnThresholdDbm  string `xml:"laser-rx-power-high-warn-threshold-dbm,omitempty"`
	LaserRxOpticalPowerLowWarnThresholdDbm   string `xml:"laser-rx-power-low-warn-threshold-dbm,omitempty"`

	LaserTxOpticalPowerHighAlarmThreshold float64 `xml:"laser-tx-power-high-alarm-threshold,omitempty"`
	LaserTxOpticalPowerLowAlarmThreshold  float64 `xml:"laser-tx-power-low-alarm-threshold,omitempty"`
	LaserTxOpticalPowerHighWarnThreshold  float64 `xml:"laser-tx-power-high-warn-threshold,omitempty"`
	LaserTxOpticalPowerLowWarnThreshold   float64 `xml:"laser-tx-power-low-warn-threshold,omitempty"`

	LaserTxOpticalPowerHighAlarmThresholdDbm string `xml:"laser-tx-power-high-alarm-threshold-dbm,omitempty"`
	LaserTxOpticalPowerLowAlarmThresholdDbm  string `xml:"laser-tx-power-low-alarm-threshold-dbm,omitempty"`
	LaserTxOpticalPowerHighWarnThresholdDbm  string `xml:"laser-tx-power-high-warn-threshold-dbm,omitempty"`
	LaserTxOpticalPowerLowWarnThresholdDbm   string `xml:"laser-tx-power-low-warn-threshold-dbm,omitempty"`

	NA string `xml:"optic-diagnostics-not-available"`

	Lanes []LaneValue `xml:"optics-diagnostics-lane-values,omitempty"`
}

type Temperature struct {
	Value float64 `xml:"celsius,attr"`
}

type LaneValue struct {
	LaneIndex              string  `xml:"lane-index"`
	LaserBiasCurrent       float64 `xml:"laser-bias-current,omitempty"`
	LaserOutputPower       float64 `xml:"laser-output-power,omitempty"`
	LaserOutputPowerDbm    string  `xml:"laser-output-power-dbm,omitempty"`
	LaserRxOpticalPower    float64 `xml:"laser-rx-optical-power,omitempty"`
	LaserRxOpticalPowerDbm string  `xml:"laser-rx-optical-power-dbm,omitempty"`
}
