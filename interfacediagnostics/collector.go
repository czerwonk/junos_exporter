package interfacediagnostics

import (
	"encoding/xml"
	"log"
	"math"
	"strconv"
	"strings"

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

// Name returns the name of the collector
func (*interfaceDiagnosticsCollector) Name() string {
	return "Interface Diagnostics"
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

	diagnosticsSatellite, err := c.interfaceDiagnosticsSatellite(client)
	if err != nil {
		return err
	}

	diagnostics = append(diagnostics, diagnosticsSatellite...)

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
	var x = InterfaceDiagnosticsRPC{}
	err := client.RunCommandAndParse("show interfaces diagnostics optics", &x)
	if err != nil {
		return nil, err
	}

	return interfaceDiagnosticsFromRPCResult(x), nil
}

func (c *interfaceDiagnosticsCollector) interfaceDiagnosticsSatellite(client *rpc.Client) ([]*InterfaceDiagnostics, error) {
	var x = InterfaceDiagnosticsRPC{}

	// NOTE: Junos is broken and delivers incorrect XML
	// 2020/01/06 12:25:24 Output for X.X.X.X: <rpc-reply xmlns:junos="http://xml.juniper.net/junos/17.3R3/junos">
	//  <interface-information xmlns="http://xml.juniper.net/junos/17.3R3/junos-interface" junos:style="normal">
	//      <interface-information xmlns="http://xml.juniper.net/junos/17.3R3/junos-interface" junos:style="normal">
	//         [..]
	//         </physical-interface>
	//      </interface-information>
	//      <cli>
	//          <banner>{master}</banner>
	//      </cli>
	//  </rpc-reply>

	// workaround: go through all lines of the XML and remove identical, consecutive lines
	err := client.RunCommandAndParseWithParser("show interfaces diagnostics optics satellite", func(b []byte) error {
		var (
			lines     []string = strings.Split(string(b[:]), "\n")
			lineIndex int
			tmpByte   []byte
		)

		// check if satellite is enabled
		if string(b[:]) == "\nerror: syntax error, expecting <command>: satellite\n" {
			log.Printf("system doesn't seem to have satellite enabled")
			return nil
		}

		for lineIndex = range lines {
			if lineIndex == 0 {
				// add good lines to new byte buffer
				tmpByte = append(tmpByte, lines[lineIndex]...)
				continue
			}

			// check if two consecutive lines are identical (except whitespaces)
			if strings.TrimSpace(lines[lineIndex]) == strings.TrimSpace(lines[lineIndex-1]) {
				// skip the duplicate line
				continue

			} else {
				// add good lines to new byte buffer
				tmpByte = append(tmpByte, lines[lineIndex]...)
			}
		}

		return xml.Unmarshal(tmpByte, &x)
	})

	if err != nil {
		return nil, err
	}

	return interfaceDiagnosticsFromRPCResult(x), nil
}

func interfaceDiagnosticsFromRPCResult(result InterfaceDiagnosticsRPC) []*InterfaceDiagnostics {
	diagnostics := make([]*InterfaceDiagnostics, 0)

	for _, diag := range result.Information.Diagnostics {
		if diag.Diagnostics.NA == "N/A" {
			continue
		}

		d := &InterfaceDiagnostics{
			Index:                      "",
			Name:                       diag.Name,
			LaserBiasCurrent:           float64(diag.Diagnostics.LaserBiasCurrent),
			LaserOutputPower:           float64(diag.Diagnostics.LaserOutputPower),
			ModuleTemperature:          float64(diag.Diagnostics.ModuleTemperature.Value),
			LaserOutputPowerDbm:        dbmStringToFloat(diag.Diagnostics.LaserOutputPowerDbm),
			ModuleVoltage:              float64(diag.Diagnostics.ModuleVoltage),
			RxSignalAvgOpticalPower:    float64(diag.Diagnostics.RxSignalAvgOpticalPower),
			RxSignalAvgOpticalPowerDbm: dbmStringToFloat(diag.Diagnostics.RxSignalAvgOpticalPowerDbm),
			LaserRxOpticalPower:        float64(diag.Diagnostics.LaserRxOpticalPower),
			LaserRxOpticalPowerDbm:     dbmStringToFloat(diag.Diagnostics.LaserRxOpticalPowerDbm),
		}

		if len(diag.Diagnostics.Lanes) > 0 {
			for _, lane := range diag.Diagnostics.Lanes {
				l := &InterfaceDiagnostics{
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
		}

		diagnostics = append(diagnostics, d)
	}

	return diagnostics
}

func dbmStringToFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return f
	}

	return math.Inf(-1)
}
