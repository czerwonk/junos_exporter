package interface_diagnostics

import "github.com/prometheus/client_golang/prometheus"

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
	laserBiasCurrentDesc = prometheus.NewDesc(prefix+"laser_bias", "Laser bias current in mA", l, nil)
	laserOutputPowerDesc = prometheus.NewDesc(prefix+"laser_output", "Laser output power in mW", l, nil)
	laserOutputPowerDbmDesc = prometheus.NewDesc(prefix+"laser_output_dbm", "Laser output power in dBm", l, nil)
	moduleTemperatureDesc = prometheus.NewDesc(prefix+"temp", "Module temperature in degrees Celsius", l, nil)

	laserRxOpticalPowerDesc = prometheus.NewDesc(prefix+"laser_rx", "Laser rx power in mW", l, nil)
	laserRxOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"laser_rx_dbm", "Laser rx power in dBm", l, nil)

	moduleVoltageDesc = prometheus.NewDesc(prefix+"module_voltage", "Module voltage", l, nil)
	rxSignalAvgOpticalPowerDesc = prometheus.NewDesc(prefix+"rx_signal_avg", "Receiver signal average optical power in mW", l, nil)
	rxSignalAvgOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"rx_signal_avg_dbm", "Receiver signal average optical power in mW", l, nil)
}

type InterfaceDiagnosticsCollector struct {
}

func (*InterfaceDiagnosticsCollector) Describe(ch chan<- *prometheus.Desc) {
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

func (c *InterfaceDiagnosticsCollector) Collect(datasource InterfaceDiagnosticsDatasource, ch chan<- prometheus.Metric, labelValues []string) error {
	diagnostics, err := datasource.InterfaceDiagnostics()
	if err != nil {
		return err
	}

	for _, d := range diagnostics {
		l := append(labelValues, d.Name)
		ch <- prometheus.MustNewConstMetric(laserBiasCurrentDesc, prometheus.GaugeValue, d.LaserBiasCurrent, l...)
		ch <- prometheus.MustNewConstMetric(laserOutputPowerDesc, prometheus.GaugeValue, d.LaserOutputPower, l...)
		ch <- prometheus.MustNewConstMetric(laserOutputPowerDbmDesc, prometheus.GaugeValue, d.LaserOutputPowerDbm, l...)
		ch <- prometheus.MustNewConstMetric(moduleTemperatureDesc, prometheus.GaugeValue, d.ModuleTemperature, l...)

		if d.ModuleVoltage > 0 {
			ch <- prometheus.MustNewConstMetric(moduleVoltageDesc, prometheus.GaugeValue, d.ModuleVoltage, l...)
			ch <- prometheus.MustNewConstMetric(rxSignalAvgOpticalPowerDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPower, l...)
			ch <- prometheus.MustNewConstMetric(rxSignalAvgOpticalPowerDbmDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPowerDbm, l...)
		} else {
			ch <- prometheus.MustNewConstMetric(laserRxOpticalPowerDesc, prometheus.GaugeValue, d.LaserRxOpticalPower, l...)
			ch <- prometheus.MustNewConstMetric(laserRxOpticalPowerDbmDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerDbm, l...)
		}
	}

	return nil
}
