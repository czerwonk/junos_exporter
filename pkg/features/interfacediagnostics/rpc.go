// SPDX-License-Identifier: MIT

package interfacediagnostics

type fpcInformationStruct struct {
	FPCInformation fpcInformation `xml:"fpc-information"`
}
type fpcInformation struct {
	FPC fpc `xml:"fpc"`
}

type fpc struct {
	PicDetail picDetail `xml:"pic-detail"`
}

type picDetail struct {
	Slot            int       `xml:"slot"`
	PicSlot         int       `xml:"pic-slot"`
	PicType         string    `xml:"pic-type"`
	State           string    `xml:"state"`
	PicVersion      string    `xml:"pic-version"`
	UpTime          string    `xml:"up-time"`
	PicPortInfoList []picPort `xml:"port-information>port"`
}

type picPort struct {
	PortNumber     int    `xml:"port-number"`
	CableType      string `xml:"cable-type"`
	FiberMode      string `xml:"fiber-mode"`
	SFPVendorName  string `xml:"sfp-vendor-name"`
	SFPVendorPno   string `xml:"sfp-vendor-pno"`
	Wavelength     string `xml:"wavelength"`
	SFPVendorFwVer string `xml:"sfp-vendor-fw-ver"`
	SFPJnprVer     string `xml:"sfp-jnpr-ver"`
}

type interfacesMediaStruct struct {
	InterfaceInformation struct {
		PhysicalInterface []physicalInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type physicalInterface struct {
	Name             string `xml:"name"`
	AdminStatus      string `xml:"admin-status"`
	OperStatus       string `xml:"oper-status"`
	LocalIndex       string `xml:"local-index"`
	SnmpIndex        string `xml:"snmp-index"`
	Description      string `xml:"description"`
	IfType           string `xml:"if-type"`
	LinkLevelType    string `xml:"link-level-type"`
	Mtu              string `xml:"mtu"`
	Speed            string `xml:"speed"`
	LinkType         string `xml:"link-type"`
	InterfaceFlapped struct {
		Seconds uint64 `xml:"seconds,attr"`
		Value   string `xml:",chardata"`
	} `xml:"interface-flapped"`
	IfDeviceFlags ifDeviceFlags `xml:"if-device-flags"`
	Stats         trafficStat   `xml:"traffic-statistics"`
}

type trafficStat struct {
	InputBytes  uint64   `xml:"input-bytes"`
	OutputBytes uint64   `xml:"output-bytes"`
	IPv6Traffic ipv6Stat `xml:"ipv6-transit-statistics"`
}
type ipv6Stat struct {
	InputBytes    uint64 `xml:"input-bytes"`
	InputPackets  uint64 `xml:"input-packets"`
	OutputBytes   uint64 `xml:"output-bytes"`
	OutputPackets uint64 `xml:"output-packets"`
}

type ifDeviceFlags struct {
	Present  bool `xml:"ifdf-present"`
	Running  bool `xml:"ifdf-running"`
	Loopback bool `xml:"ifdf-loopback"`
	Down     bool `xml:"ifdf-down"`
}

type chassisHardware struct {
	ChassisInventory chassisInventory `xml:"chassis-inventory"`
}

type chassisInventory struct {
	Chassis chassis `xml:"chassis"`
}

type chassis struct {
	Name          string          `xml:"name"`
	SerialNumber  string          `xml:"serial-number"`
	Description   string          `xml:"description"`
	ChassisModule []chassisModule `xml:"chassis-module"`
}

type chassisModule struct {
	Name             string             `xml:"name"`
	PartNumber       string             `xml:"part-number"`
	SerialNumber     string             `xml:"serial-number"`
	Description      string             `xml:"description"`
	CleiCode         string             `xml:"clei-code"`
	ModelNumber      string             `xml:"model-number"`
	ChassisSubModule []chassisSubModule `xml:"chassis-sub-module"`
}

type chassisSubModule struct {
	Name                string                `xml:"name"`
	PartNumber          string                `xml:"part-number"`
	SerialNumber        string                `xml:"serial-number"`
	Description         string                `xml:"description"`
	CleiCode            string                `xml:"clei-code"`
	ModelNumber         string                `xml:"model-number"`
	ChassisSubSubModule []chassisSubSubModule `xml:"chassis-sub-sub-module"`
}
type chassisSubSubModule struct {
	Name         string `xml:"name"`
	Version      string `xml:"version"`
	PartNumber   string `xml:"part-number"`
	SerialNumber string `xml:"serial-number"`
	Description  string `xml:"description"`
}

type result struct {
	Information struct {
		Diagnostics []phyDiagInterface `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type phyDiagInterface struct {
	Name        string                 `xml:"name"`
	Diagnostics phyInterfaceDiagnostic `xml:"optics-diagnostics,omitempty"`
}

type phyInterfaceDiagnostic struct {
	LaserBiasCurrent                   float64 `xml:"laser-bias-current,omitempty"`
	LaserBiasCurrentHighAlarmThreshold float64 `xml:"laser-bias-current-high-alarm-threshold,omitempty"`
	LaserBiasCurrentLowAlarmThreshold  float64 `xml:"laser-bias-current-low-alarm-threshold,omitempty"`
	LaserBiasCurrentHighWarnThreshold  float64 `xml:"laser-bias-current-high-warn-threshold,omitempty"`
	LaserBiasCurrentLowWarnThreshold   float64 `xml:"laser-bias-current-low-warn-threshold,omitempty"`

	LaserOutputPower                    float64          `xml:"laser-output-power,omitempty"`
	LaserOutputPowerDbm                 string           `xml:"laser-output-power-dbm,omitempty"`
	ModuleTemperature                   temperatureValue `xml:"module-temperature"`
	ModuleTemperatureHighAlarmThreshold temperatureValue `xml:"module-temperature-high-alarm-threshold,omitempty"`
	ModuleTemperatureLowAlarmThreshold  temperatureValue `xml:"module-temperature-low-alarm-threshold,omitempty"`
	ModuleTemperatureHighWarnThreshold  temperatureValue `xml:"module-temperature-high-warn-threshold,omitempty"`
	ModuleTemperatureLowWarnThreshold   temperatureValue `xml:"module-temperature-low-warn-threshold,omitempty"`

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

	Lanes []lane `xml:"optics-diagnostics-lane-values,omitempty"`
}

type temperatureValue struct {
	Value float64 `xml:"celsius,attr"`
}

type lane struct {
	LaneIndex              string  `xml:"lane-index"`
	LaserBiasCurrent       float64 `xml:"laser-bias-current,omitempty"`
	LaserOutputPower       float64 `xml:"laser-output-power,omitempty"`
	LaserOutputPowerDbm    string  `xml:"laser-output-power-dbm,omitempty"`
	LaserRxOpticalPower    float64 `xml:"laser-rx-optical-power,omitempty"`
	LaserRxOpticalPowerDbm string  `xml:"laser-rx-optical-power-dbm,omitempty"`
}
