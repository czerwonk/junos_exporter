package interfacediagnostics

import (
	"strconv"

	"github.com/czerwonk/junos_exporter/rpc"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_diagnostics_"

var (
	laserBiasCurrentDesc    *prometheus.Desc
	laserOutputPowerDesc    *prometheus.Desc
	laserOutputPowerDbmDesc *prometheus.Desc
	moduleTemperatureDesc   *prometheus.Desc

	laserRxOpticalPowerDesc    *prometheus.Desc
	laserRxOpticalPowerDbmDesc *prometheus.Desc

	moduleVoltageDesc              *prometheus.Desc
	rxSignalAvgOpticalPowerDesc    *prometheus.Desc
	rxSignalAvgOpticalPowerDbmDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "name"}
	moduleVoltageDesc = prometheus.NewDesc(prefix+"module_voltage", "Module voltage", l, nil)
	moduleTemperatureDesc = prometheus.NewDesc(prefix+"temp", "Module temperature in degrees Celsius", l, nil)
	rxSignalAvgOpticalPowerDesc = prometheus.NewDesc(prefix+"rx_signal_avg", "Receiver signal average optical power in mW", l, nil)
	rxSignalAvgOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"rx_signal_avg_dbm", "Receiver signal average optical power in mW", l, nil)

	l = append(l, "lane")
	laserBiasCurrentDesc = prometheus.NewDesc(prefix+"laser_bias", "Laser bias current in mA", l, nil)
	laserOutputPowerDesc = prometheus.NewDesc(prefix+"laser_output", "Laser output power in mW", l, nil)
	laserOutputPowerDbmDesc = prometheus.NewDesc(prefix+"laser_output_dbm", "Laser output power in dBm", l, nil)
	laserRxOpticalPowerDesc = prometheus.NewDesc(prefix+"laser_rx", "Laser rx power in mW", l, nil)
	laserRxOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"laser_rx_dbm", "Laser rx power in dBm", l, nil)
}

type interfaceDiagnosticsCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &interfaceDiagnosticsCollector{}
}

// Describe describes the metrics
func (*interfaceDiagnosticsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- laserBiasCurrentDesc
	ch <- laserOutputPowerDesc
	ch <- laserOutputPowerDbmDesc
	ch <- moduleTemperatureDesc

	ch <- laserRxOpticalPowerDesc
	ch <- laserRxOpticalPowerDbmDesc

	ch <- moduleVoltageDesc
	ch <- rxSignalAvgOpticalPowerDesc
	ch <- rxSignalAvgOpticalPowerDbmDesc
}

// Collect collects metrics from JunOS
func (c *interfaceDiagnosticsCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	diagnostics, err := c.interfaceDiagnostics(client)
	if err != nil {
		return err
	}

	for _, d := range diagnostics {
		l := append(labelValues, d.Name)
		ch <- prometheus.MustNewConstMetric(moduleTemperatureDesc, prometheus.GaugeValue, d.ModuleTemperature, l...)
		if d.ModuleVoltage > 0 {
			ch <- prometheus.MustNewConstMetric(moduleVoltageDesc, prometheus.GaugeValue, d.ModuleVoltage, l...)
			ch <- prometheus.MustNewConstMetric(rxSignalAvgOpticalPowerDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPower, l...)
			ch <- prometheus.MustNewConstMetric(rxSignalAvgOpticalPowerDbmDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPowerDbm, l...)
		}

		var data []*InterfaceDiagnostics
		if len(d.Lanes) > 0 {
			data = d.Lanes
		} else {
			data = []*InterfaceDiagnostics{d}
		}

		for _, e := range data {
			l2 := append(l, e.Index)
			ch <- prometheus.MustNewConstMetric(laserBiasCurrentDesc, prometheus.GaugeValue, e.LaserBiasCurrent, l2...)
			ch <- prometheus.MustNewConstMetric(laserOutputPowerDesc, prometheus.GaugeValue, e.LaserOutputPower, l2...)
			ch <- prometheus.MustNewConstMetric(laserOutputPowerDbmDesc, prometheus.GaugeValue, e.LaserOutputPowerDbm, l2...)
			if d.ModuleVoltage <= 0 {
				ch <- prometheus.MustNewConstMetric(laserRxOpticalPowerDesc, prometheus.GaugeValue, e.LaserRxOpticalPower, l2...)
				ch <- prometheus.MustNewConstMetric(laserRxOpticalPowerDbmDesc, prometheus.GaugeValue, e.LaserRxOpticalPowerDbm, l2...)
			}
		}
	}

	return nil
}

func (c *interfaceDiagnosticsCollector) interfaceDiagnostics(client *rpc.Client) ([]*InterfaceDiagnostics, error) {
	var x = InterfaceDiagnosticsRpc{}
	err := client.RunCommandAndParse("show interfaces diagnostics optics", &x)
	if err != nil {
		return nil, err
	}

	diagnostics := make([]*InterfaceDiagnostics, 0)
	for _, diag := range x.Information.Diagnostics {
		if diag.Diagnostics.NA == "N/A" {
			continue
		}
		d := &InterfaceDiagnostics{
			Index:             "",
			Name:              diag.Name,
			LaserBiasCurrent:  float64(diag.Diagnostics.LaserBiasCurrent),
			LaserOutputPower:  float64(diag.Diagnostics.LaserOutputPower),
			ModuleTemperature: float64(diag.Diagnostics.ModuleTemperature.Value),
		}
		f, err := strconv.ParseFloat(diag.Diagnostics.LaserOutputPowerDbm, 64)
		if err == nil {
			d.LaserOutputPowerDbm = f
		}

		if diag.Diagnostics.ModuleVoltage > 0 {
			d.ModuleVoltage = float64(diag.Diagnostics.ModuleVoltage)
			d.RxSignalAvgOpticalPower = float64(diag.Diagnostics.RxSignalAvgOpticalPower)
			f, err = strconv.ParseFloat(diag.Diagnostics.RxSignalAvgOpticalPowerDbm, 64)
			if err == nil {
				d.RxSignalAvgOpticalPowerDbm = f
			}
		} else {
			d.LaserRxOpticalPower = float64(diag.Diagnostics.LaserRxOpticalPower)
			f, err = strconv.ParseFloat(diag.Diagnostics.LaserRxOpticalPowerDbm, 64)
			if err == nil {
				d.LaserRxOpticalPowerDbm = f
			}
		}

		if len(diag.Diagnostics.OpticsDiagnosticsLaneValues) > 0 {
			for _, lane := range diag.Diagnostics.OpticsDiagnosticsLaneValues {
				l := &InterfaceDiagnostics{
					Index:            lane.LaneIndex,
					Name:             diag.Name,
					LaserBiasCurrent: float64(lane.LaserBiasCurrent),
					LaserOutputPower: float64(lane.LaserOutputPower),
				}
				f, err := strconv.ParseFloat(lane.LaserOutputPowerDbm, 64)
				if err == nil {
					l.LaserOutputPowerDbm = f
				}
				if diag.Diagnostics.ModuleVoltage <= 0 {
					l.LaserRxOpticalPower = float64(lane.LaserRxOpticalPower)
					f, err = strconv.ParseFloat(lane.LaserRxOpticalPowerDbm, 64)
					if err == nil {
						l.LaserRxOpticalPowerDbm = f
					}
				}
				d.Lanes = append(d.Lanes, l)
			}
		}

		diagnostics = append(diagnostics, d)
	}

	return diagnostics, nil
}
