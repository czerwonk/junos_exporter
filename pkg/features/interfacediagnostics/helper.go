// SPDX-License-Identifier: MIT

package interfacediagnostics

import (
	"math"
	"strconv"
	"strings"
)

func dbmStringToFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return f
	}

	return math.Inf(-1)
}

func interfaceDiagnosticsFromRPCResult(res result) []*interfaceDiagnostics {
	diagnostics := make([]*interfaceDiagnostics, 0)

	for _, diag := range res.Information.Diagnostics {
		if diag.Diagnostics.NA == "N/A" {
			continue
		}

		d := &interfaceDiagnostics{
			Index:                              "",
			Name:                               diag.Name,
			LaserBiasCurrent:                   float64(diag.Diagnostics.LaserBiasCurrent),
			LaserBiasCurrentHighAlarmThreshold: float64(diag.Diagnostics.LaserBiasCurrentHighAlarmThreshold),
			LaserBiasCurrentLowAlarmThreshold:  float64(diag.Diagnostics.LaserBiasCurrentLowAlarmThreshold),
			LaserBiasCurrentHighWarnThreshold:  float64(diag.Diagnostics.LaserBiasCurrentHighWarnThreshold),
			LaserBiasCurrentLowWarnThreshold:   float64(diag.Diagnostics.LaserBiasCurrentLowWarnThreshold),

			LaserOutputPower:                   float64(diag.Diagnostics.LaserOutputPower),
			LaserOutputPowerHighAlarmThreshold: float64(diag.Diagnostics.LaserTxOpticalPowerHighAlarmThreshold),
			LaserOutputPowerLowAlarmThreshold:  float64(diag.Diagnostics.LaserTxOpticalPowerLowAlarmThreshold),
			LaserOutputPowerHighWarnThreshold:  float64(diag.Diagnostics.LaserTxOpticalPowerHighWarnThreshold),
			LaserOutputPowerLowWarnThreshold:   float64(diag.Diagnostics.LaserTxOpticalPowerLowWarnThreshold),

			ModuleTemperature:                   float64(diag.Diagnostics.ModuleTemperature.Value),
			ModuleTemperatureHighAlarmThreshold: float64(diag.Diagnostics.ModuleTemperatureHighAlarmThreshold.Value),
			ModuleTemperatureLowAlarmThreshold:  float64(diag.Diagnostics.ModuleTemperatureLowAlarmThreshold.Value),
			ModuleTemperatureHighWarnThreshold:  float64(diag.Diagnostics.ModuleTemperatureHighWarnThreshold.Value),
			ModuleTemperatureLowWarnThreshold:   float64(diag.Diagnostics.ModuleTemperatureLowWarnThreshold.Value),

			LaserOutputPowerDbm:                   dbmStringToFloat(diag.Diagnostics.LaserOutputPowerDbm),
			LaserOutputPowerHighAlarmThresholdDbm: dbmStringToFloat(diag.Diagnostics.LaserTxOpticalPowerHighAlarmThresholdDbm),
			LaserOutputPowerLowAlarmThresholdDbm:  dbmStringToFloat(diag.Diagnostics.LaserTxOpticalPowerLowAlarmThresholdDbm),
			LaserOutputPowerHighWarnThresholdDbm:  dbmStringToFloat(diag.Diagnostics.LaserTxOpticalPowerHighWarnThresholdDbm),
			LaserOutputPowerLowWarnThresholdDbm:   dbmStringToFloat(diag.Diagnostics.LaserTxOpticalPowerLowWarnThresholdDbm),

			ModuleVoltage:                   float64(diag.Diagnostics.ModuleVoltage),
			ModuleVoltageHighAlarmThreshold: float64(diag.Diagnostics.ModuleVoltageHighAlarmThreshold),
			ModuleVoltageLowAlarmThreshold:  float64(diag.Diagnostics.ModuleVoltageLowAlarmThreshold),
			ModuleVoltageHighWarnThreshold:  float64(diag.Diagnostics.ModuleVoltageHighWarnThreshold),
			ModuleVoltageLowWarnThreshold:   float64(diag.Diagnostics.ModuleVoltageLowWarnThreshold),

			RxSignalAvgOpticalPower:               float64(diag.Diagnostics.RxSignalAvgOpticalPower),
			RxSignalAvgOpticalPowerDbm:            dbmStringToFloat(diag.Diagnostics.RxSignalAvgOpticalPowerDbm),
			LaserRxOpticalPower:                   float64(diag.Diagnostics.LaserRxOpticalPower),
			LaserRxOpticalPowerHighAlarmThreshold: float64(diag.Diagnostics.LaserRxOpticalPowerHighAlarmThreshold),
			LaserRxOpticalPowerLowAlarmThreshold:  float64(diag.Diagnostics.LaserRxOpticalPowerLowAlarmThreshold),
			LaserRxOpticalPowerHighWarnThreshold:  float64(diag.Diagnostics.LaserRxOpticalPowerHighWarnThreshold),
			LaserRxOpticalPowerLowWarnThreshold:   float64(diag.Diagnostics.LaserRxOpticalPowerLowWarnThreshold),

			LaserRxOpticalPowerDbm:                   dbmStringToFloat(diag.Diagnostics.LaserRxOpticalPowerDbm),
			LaserRxOpticalPowerHighAlarmThresholdDbm: dbmStringToFloat(diag.Diagnostics.LaserRxOpticalPowerHighAlarmThresholdDbm),
			LaserRxOpticalPowerLowAlarmThresholdDbm:  dbmStringToFloat(diag.Diagnostics.LaserRxOpticalPowerLowAlarmThresholdDbm),
			LaserRxOpticalPowerHighWarnThresholdDbm:  dbmStringToFloat(diag.Diagnostics.LaserRxOpticalPowerHighWarnThresholdDbm),
			LaserRxOpticalPowerLowWarnThresholdDbm:   dbmStringToFloat(diag.Diagnostics.LaserRxOpticalPowerLowWarnThresholdDbm),
		}

		if len(diag.Diagnostics.Lanes) > 0 {
			for _, lane := range diag.Diagnostics.Lanes {
				l := &interfaceDiagnostics{
					Index:                  lane.LaneIndex,
					Name:                   diag.Name,
					LaserBiasCurrent:       float64(lane.LaserBiasCurrent),
					LaserOutputPower:       float64(lane.LaserOutputPower),
					LaserOutputPowerDbm:    dbmStringToFloat(lane.LaserOutputPowerDbm),
					LaserRxOpticalPower:    float64(lane.LaserRxOpticalPower),
					LaserRxOpticalPowerDbm: dbmStringToFloat(lane.LaserRxOpticalPowerDbm),
				}

				d.Lanes = append(d.Lanes, l)
			}

			/* For some interfaces with 0 lanes there sometimes is  <rx-signal-avg-optical-power> instead of
			<laser-rx-optical-power> in the xml/json response and vice-versa.*/
		} else if diag.Diagnostics.LaserRxOpticalPowerDbm == "" {
			d.LaserRxOpticalPower = d.RxSignalAvgOpticalPower
			d.LaserRxOpticalPowerDbm = d.RxSignalAvgOpticalPowerDbm

		} else if diag.Diagnostics.RxSignalAvgOpticalPowerDbm == "" {
			d.RxSignalAvgOpticalPower = d.LaserRxOpticalPower
			d.RxSignalAvgOpticalPowerDbm = d.LaserRxOpticalPowerDbm
		}

		diagnostics = append(diagnostics, d)
	}

	return diagnostics
}

func interfaceMediaInfoFromRPCResult(interfaceMediaList *[]physicalInterface) map[string]*physicalInterface {
	interfaceMediaDict := make(map[string]*physicalInterface)

	for _, i := range *interfaceMediaList {
		if strings.HasPrefix(i.Name, "xe") || strings.HasPrefix(i.Name, "ge") || strings.HasPrefix(i.Name, "et") {
			iface := i
			interfaceMediaDict[slotIndex(iface.Name)] = &iface
		}
	}

	return interfaceMediaDict
}

func slotIndex(ifName string) string {
	return ifName[3:]
}
