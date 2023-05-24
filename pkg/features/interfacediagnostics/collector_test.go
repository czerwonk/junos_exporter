// SPDX-License-Identifier: MIT

package interfacediagnostics

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceDiagnosticsFromRPCResult(t *testing.T) {
	result := result{}

	result.Information.Diagnostics = []phyDiagInterface{
		{
			Name: "xe-0/0/0",
			Diagnostics: phyInterfaceDiagnostic{
				LaserBiasCurrent:                         1,
				LaserBiasCurrentHighAlarmThreshold:       110,
				LaserBiasCurrentLowAlarmThreshold:        10,
				LaserBiasCurrentHighWarnThreshold:        19,
				LaserBiasCurrentLowWarnThreshold:         11,
				LaserOutputPower:                         2,
				LaserTxOpticalPowerHighAlarmThreshold:    210,
				LaserTxOpticalPowerLowAlarmThreshold:     20,
				LaserTxOpticalPowerHighWarnThreshold:     29,
				LaserTxOpticalPowerLowWarnThreshold:      21,
				LaserOutputPowerDbm:                      "3",
				LaserTxOpticalPowerHighAlarmThresholdDbm: "310",
				LaserTxOpticalPowerLowAlarmThresholdDbm:  "30",
				LaserTxOpticalPowerHighWarnThresholdDbm:  "39",
				LaserTxOpticalPowerLowWarnThresholdDbm:   "31",
				ModuleTemperature: temperatureValue{
					Value: 4,
				},
				ModuleTemperatureHighAlarmThreshold: temperatureValue{
					Value: 410,
				},
				ModuleTemperatureLowAlarmThreshold: temperatureValue{
					Value: 40,
				},
				ModuleTemperatureHighWarnThreshold: temperatureValue{
					Value: 49,
				},
				ModuleTemperatureLowWarnThreshold: temperatureValue{
					Value: 41,
				},
				ModuleVoltage:                            5,
				ModuleVoltageHighAlarmThreshold:          510,
				ModuleVoltageLowAlarmThreshold:           50,
				ModuleVoltageHighWarnThreshold:           59,
				ModuleVoltageLowWarnThreshold:            51,
				RxSignalAvgOpticalPower:                  6,
				RxSignalAvgOpticalPowerDbm:               "- Inf",
				LaserRxOpticalPower:                      7,
				LaserRxOpticalPowerHighAlarmThreshold:    710,
				LaserRxOpticalPowerLowAlarmThreshold:     70,
				LaserRxOpticalPowerHighWarnThreshold:     79,
				LaserRxOpticalPowerLowWarnThreshold:      71,
				LaserRxOpticalPowerDbm:                   "- Inf",
				LaserRxOpticalPowerHighAlarmThresholdDbm: "810",
				LaserRxOpticalPowerLowAlarmThresholdDbm:  "80",
				LaserRxOpticalPowerHighWarnThresholdDbm:  "89",
				LaserRxOpticalPowerLowWarnThresholdDbm:   "81",
				Lanes: []lane{
					{
						LaneIndex:              "1",
						LaserBiasCurrent:       8,
						LaserOutputPower:       9,
						LaserOutputPowerDbm:    "10",
						LaserRxOpticalPower:    11,
						LaserRxOpticalPowerDbm: "- Inf",
					},
				},
			},
		},
	}

	ifaces := interfaceDiagnosticsFromRPCResult(result)

	assert.Equal(t, 1, len(ifaces), "interface count")

	iface := ifaces[0]

	assert.Equal(t, float64(1), iface.LaserBiasCurrent)
	assert.Equal(t, float64(110), iface.LaserBiasCurrentHighAlarmThreshold)
	assert.Equal(t, float64(10), iface.LaserBiasCurrentLowAlarmThreshold)
	assert.Equal(t, float64(19), iface.LaserBiasCurrentHighWarnThreshold)
	assert.Equal(t, float64(11), iface.LaserBiasCurrentLowWarnThreshold)
	assert.Equal(t, float64(2), iface.LaserOutputPower)
	assert.Equal(t, float64(210), iface.LaserOutputPowerHighAlarmThreshold)
	assert.Equal(t, float64(20), iface.LaserOutputPowerLowAlarmThreshold)
	assert.Equal(t, float64(29), iface.LaserOutputPowerHighWarnThreshold)
	assert.Equal(t, float64(21), iface.LaserOutputPowerLowWarnThreshold)
	assert.Equal(t, float64(3), iface.LaserOutputPowerDbm)
	assert.Equal(t, float64(310), iface.LaserOutputPowerHighAlarmThresholdDbm)
	assert.Equal(t, float64(30), iface.LaserOutputPowerLowAlarmThresholdDbm)
	assert.Equal(t, float64(39), iface.LaserOutputPowerHighWarnThresholdDbm)
	assert.Equal(t, float64(31), iface.LaserOutputPowerLowWarnThresholdDbm)
	assert.Equal(t, float64(4), iface.ModuleTemperature)
	assert.Equal(t, float64(410), iface.ModuleTemperatureHighAlarmThreshold)
	assert.Equal(t, float64(40), iface.ModuleTemperatureLowAlarmThreshold)
	assert.Equal(t, float64(49), iface.ModuleTemperatureHighWarnThreshold)
	assert.Equal(t, float64(41), iface.ModuleTemperatureLowWarnThreshold)
	assert.Equal(t, float64(5), iface.ModuleVoltage)
	assert.Equal(t, float64(510), iface.ModuleVoltageHighAlarmThreshold)
	assert.Equal(t, float64(50), iface.ModuleVoltageLowAlarmThreshold)
	assert.Equal(t, float64(59), iface.ModuleVoltageHighWarnThreshold)
	assert.Equal(t, float64(51), iface.ModuleVoltageLowWarnThreshold)
	assert.Equal(t, float64(6), iface.RxSignalAvgOpticalPower)
	assert.Equal(t, math.Inf(-1), iface.RxSignalAvgOpticalPowerDbm)
	assert.Equal(t, float64(7), iface.LaserRxOpticalPower)
	assert.Equal(t, float64(710), iface.LaserRxOpticalPowerHighAlarmThreshold)
	assert.Equal(t, float64(70), iface.LaserRxOpticalPowerLowAlarmThreshold)
	assert.Equal(t, float64(79), iface.LaserRxOpticalPowerHighWarnThreshold)
	assert.Equal(t, float64(71), iface.LaserRxOpticalPowerLowWarnThreshold)

	assert.Equal(t, math.Inf(-1), iface.LaserRxOpticalPowerDbm)
	assert.Equal(t, float64(810), iface.LaserRxOpticalPowerHighAlarmThresholdDbm)
	assert.Equal(t, float64(80), iface.LaserRxOpticalPowerLowAlarmThresholdDbm)
	assert.Equal(t, float64(89), iface.LaserRxOpticalPowerHighWarnThresholdDbm)
	assert.Equal(t, float64(81), iface.LaserRxOpticalPowerLowWarnThresholdDbm)

	assert.Equal(t, 1, len(iface.Lanes), "lane count")

	l := iface.Lanes[0]
	assert.Equal(t, "1", l.Index, "lane index")
	assert.Equal(t, float64(8), l.LaserBiasCurrent)
	assert.Equal(t, float64(9), l.LaserOutputPower)
	assert.Equal(t, float64(10), l.LaserOutputPowerDbm)
	assert.Equal(t, float64(11), l.LaserRxOpticalPower)
	assert.Equal(t, math.Inf(-1), l.LaserRxOpticalPowerDbm)
}
