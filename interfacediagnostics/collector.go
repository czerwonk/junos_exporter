package interfacediagnostics

import (
	"strconv"

	"github.com/czerwonk/junos_exporter/interfacelabels"
	"github.com/czerwonk/junos_exporter/rpc"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_diagnostics_"

type interfaceDiagnosticsCollector struct {
	labels                         *interfacelabels.DynamicLabels
	laserBiasCurrentDesc           *prometheus.Desc
	laserOutputPowerDesc           *prometheus.Desc
	laserOutputPowerDbmDesc        *prometheus.Desc
	moduleTemperatureDesc          *prometheus.Desc
	laserRxOpticalPowerDesc        *prometheus.Desc
	laserRxOpticalPowerDbmDesc     *prometheus.Desc
	moduleVoltageDesc              *prometheus.Desc
	rxSignalAvgOpticalPowerDesc    *prometheus.Desc
	rxSignalAvgOpticalPowerDbmDesc *prometheus.Desc
}

// NewCollector creates a new collector
func NewCollector(labels *interfacelabels.DynamicLabels) collector.RPCCollector {
	c := &interfaceDiagnosticsCollector{
		labels: labels,
	}
	c.init()

	return c
}

func (c *interfaceDiagnosticsCollector) init() {
	l := []string{"target", "name"}
	l = append(l, c.labels.LabelNames()...)

	c.moduleVoltageDesc = prometheus.NewDesc(prefix+"module_voltage", "Module voltage", l, nil)
	c.moduleTemperatureDesc = prometheus.NewDesc(prefix+"temp", "Module temperature in degrees Celsius", l, nil)
	c.rxSignalAvgOpticalPowerDesc = prometheus.NewDesc(prefix+"rx_signal_avg", "Receiver signal average optical power in mW", l, nil)
	c.rxSignalAvgOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"rx_signal_avg_dbm", "Receiver signal average optical power in mW", l, nil)

	l = append(l, "lane")
	c.laserBiasCurrentDesc = prometheus.NewDesc(prefix+"laser_bias", "Laser bias current in mA", l, nil)
	c.laserOutputPowerDesc = prometheus.NewDesc(prefix+"laser_output", "Laser output power in mW", l, nil)
	c.laserOutputPowerDbmDesc = prometheus.NewDesc(prefix+"laser_output_dbm", "Laser output power in dBm", l, nil)
	c.laserRxOpticalPowerDesc = prometheus.NewDesc(prefix+"laser_rx", "Laser rx power in mW", l, nil)
	c.laserRxOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"laser_rx_dbm", "Laser rx power in dBm", l, nil)
}

// Describe describes the metrics
func (c *interfaceDiagnosticsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.laserBiasCurrentDesc
	ch <- c.laserOutputPowerDesc
	ch <- c.laserOutputPowerDbmDesc
	ch <- c.moduleTemperatureDesc
	ch <- c.laserRxOpticalPowerDesc
	ch <- c.laserRxOpticalPowerDbmDesc
	ch <- c.moduleVoltageDesc
	ch <- c.rxSignalAvgOpticalPowerDesc
	ch <- c.rxSignalAvgOpticalPowerDbmDesc
}

// Collect collects metrics from JunOS
func (c *interfaceDiagnosticsCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	diagnostics, err := c.interfaceDiagnostics(client)
	if err != nil {
		return err
	}

	for _, d := range diagnostics {
		l := append(labelValues, d.Name)
		l = append(l, c.labels.ValuesForInterface(client.Device(), d.Name)...)

		ch <- prometheus.MustNewConstMetric(c.moduleTemperatureDesc, prometheus.GaugeValue, d.ModuleTemperature, l...)
		if d.ModuleVoltage > 0 {
			ch <- prometheus.MustNewConstMetric(c.moduleVoltageDesc, prometheus.GaugeValue, d.ModuleVoltage, l...)
		}
		if d.RxSignalAvgOpticalPower > 0 {
			ch <- prometheus.MustNewConstMetric(c.rxSignalAvgOpticalPowerDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPower, l...)
			ch <- prometheus.MustNewConstMetric(c.rxSignalAvgOpticalPowerDbmDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPowerDbm, l...)
		}

		var data []*InterfaceDiagnostics
		if len(d.Lanes) > 0 {
			data = d.Lanes
		} else {
			data = []*InterfaceDiagnostics{d}
		}

		for _, e := range data {
			l2 := append(l, e.Index)
			ch <- prometheus.MustNewConstMetric(c.laserBiasCurrentDesc, prometheus.GaugeValue, e.LaserBiasCurrent, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerDesc, prometheus.GaugeValue, e.LaserOutputPower, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerDbmDesc, prometheus.GaugeValue, e.LaserOutputPowerDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerDesc, prometheus.GaugeValue, e.LaserRxOpticalPower, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerDbmDesc, prometheus.GaugeValue, e.LaserRxOpticalPowerDbm, l2...)
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

		d.ModuleVoltage = float64(diag.Diagnostics.ModuleVoltage)
		d.RxSignalAvgOpticalPower = float64(diag.Diagnostics.RxSignalAvgOpticalPower)
		f, err = strconv.ParseFloat(diag.Diagnostics.RxSignalAvgOpticalPowerDbm, 64)
		if err == nil {
			d.RxSignalAvgOpticalPowerDbm = f
		}
		d.LaserRxOpticalPower = float64(diag.Diagnostics.LaserRxOpticalPower)
		f, err = strconv.ParseFloat(diag.Diagnostics.LaserRxOpticalPowerDbm, 64)
		if err == nil {
			d.LaserRxOpticalPowerDbm = f
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
				l.LaserRxOpticalPower = float64(lane.LaserRxOpticalPower)
				f, err = strconv.ParseFloat(lane.LaserRxOpticalPowerDbm, 64)
				if err == nil {
					l.LaserRxOpticalPowerDbm = f
				}
				d.Lanes = append(d.Lanes, l)
			}
		}

		diagnostics = append(diagnostics, d)
	}

	return diagnostics, nil
}
