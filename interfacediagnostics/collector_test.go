package interfacediagnostics

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceDiagnosticsFromRPCResult(t *testing.T) {
	result := InterfaceDiagnosticsRPC{}

	result.Information.Diagnostics = []PhyDiagInterface{
		{
			Name: "xe-0/0/0",
			Diagnostics: PhyInterfaceDiagnostic{
				LaserBiasCurrent:    1,
				LaserOutputPower:    2,
				LaserOutputPowerDbm: "3",
				ModuleTemperature: Temperature{
					Value: 4,
				},
				ModuleVoltage:              5,
				RxSignalAvgOpticalPower:    6,
				RxSignalAvgOpticalPowerDbm: "- Inf",
				LaserRxOpticalPower:        7,
				LaserRxOpticalPowerDbm:     "- Inf",
				Lanes: []LaneValue{
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
	assert.Equal(t, float64(2), iface.LaserOutputPower)
	assert.Equal(t, float64(3), iface.LaserOutputPowerDbm)
	assert.Equal(t, float64(4), iface.ModuleTemperature)
	assert.Equal(t, float64(5), iface.ModuleVoltage)
	assert.Equal(t, float64(6), iface.RxSignalAvgOpticalPower)
	assert.Equal(t, math.Inf(-1), iface.RxSignalAvgOpticalPowerDbm)
	assert.Equal(t, float64(7), iface.LaserRxOpticalPower)
	assert.Equal(t, math.Inf(-1), iface.LaserRxOpticalPowerDbm)

	assert.Equal(t, 1, len(iface.Lanes), "lane count")

	l := iface.Lanes[0]
	assert.Equal(t, "1", l.Index, "lane index")
	assert.Equal(t, float64(8), l.LaserBiasCurrent)
	assert.Equal(t, float64(9), l.LaserOutputPower)
	assert.Equal(t, float64(10), l.LaserOutputPowerDbm)
	assert.Equal(t, float64(11), l.LaserRxOpticalPower)
	assert.Equal(t, math.Inf(-1), l.LaserRxOpticalPowerDbm)
}
