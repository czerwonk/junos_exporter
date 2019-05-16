package interfacediagnostics

type InterfaceDiagnosticsRpc struct {
	Information struct {
		Diagnostics []PhyDiagInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type PhyDiagInterface struct {
	Name        string `xml:"name"`
	Diagnostics struct {
		LaserBiasCurrent    float64 `xml:"laser-bias-current,omitempty"`
		LaserOutputPower    float64 `xml:"laser-output-power,omitempty"`
		LaserOutputPowerDbm string  `xml:"laser-output-power-dbm,omitempty"`
		ModuleTemperature   struct {
			Value float64 `xml:"celsius,attr"`
		} `xml:"module-temperature"`

		ModuleVoltage              float64 `xml:"module-voltage,omitempty"`
		RxSignalAvgOpticalPower    float64 `xml:"rx-signal-avg-optical-power,omitempty"`
		RxSignalAvgOpticalPowerDbm string  `xml:"rx-signal-avg-optical-power-dbm,omitempty"`

		LaserRxOpticalPower    float64 `xml:"laser-rx-optical-power,omitempty"`
		LaserRxOpticalPowerDbm string  `xml:"laser-rx-optical-power-dbm,omitempty"`

		NA string `xml:"optic-diagnostics-not-available"`

		OpticsDiagnosticsLaneValues []LaneValue `xml:"optics-diagnostics-lane-values,omitempty"`
	} `xml:"optics-diagnostics,omitempty"`
}

type LaneValue struct {
	LaneIndex              string  `xml:"lane-index"`
	LaserBiasCurrent       float64 `xml:"laser-bias-current,omitempty"`
	LaserOutputPower       float64 `xml:"laser-output-power,omitempty"`
	LaserOutputPowerDbm    string  `xml:"laser-output-power-dbm,omitempty"`
	LaserRxOpticalPower    float64 `xml:"laser-rx-optical-power,omitempty"`
	LaserRxOpticalPowerDbm string  `xml:"laser-rx-optical-power-dbm,omitempty"`
}
