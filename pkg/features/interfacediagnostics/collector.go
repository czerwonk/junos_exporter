// SPDX-License-Identifier: MIT

package interfacediagnostics

import (
	"encoding/xml"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/dynamiclabels"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_diagnostics_"

type description struct {
	laserBiasCurrentDesc                         *prometheus.Desc
	laserBiasCurrentHighAlarmThresholdDesc       *prometheus.Desc
	laserBiasCurrentLowAlarmThresholdDesc        *prometheus.Desc
	laserBiasCurrentHighWarnThresholdDesc        *prometheus.Desc
	laserBiasCurrentLowWarnThresholdDesc         *prometheus.Desc
	laserOutputPowerDesc                         *prometheus.Desc
	laserOutputPowerHighAlarmThresholdDesc       *prometheus.Desc
	laserOutputPowerLowAlarmThresholdDesc        *prometheus.Desc
	laserOutputPowerHighWarnThresholdDesc        *prometheus.Desc
	laserOutputPowerLowWarnThresholdDesc         *prometheus.Desc
	laserOutputPowerDbmDesc                      *prometheus.Desc
	laserOutputPowerHighAlarmThresholdDbmDesc    *prometheus.Desc
	laserOutputPowerLowAlarmThresholdDbmDesc     *prometheus.Desc
	laserOutputPowerHighWarnThresholdDbmDesc     *prometheus.Desc
	laserOutputPowerLowWarnThresholdDbmDesc      *prometheus.Desc
	moduleTemperatureDesc                        *prometheus.Desc
	moduleTemperatureHighAlarmThresholdDesc      *prometheus.Desc
	moduleTemperatureLowAlarmThresholdDesc       *prometheus.Desc
	moduleTemperatureHighWarnThresholdDesc       *prometheus.Desc
	moduleTemperatureLowWarnThresholdDesc        *prometheus.Desc
	laserRxOpticalPowerDesc                      *prometheus.Desc
	laserRxOpticalPowerHighAlarmThresholdDesc    *prometheus.Desc
	laserRxOpticalPowerLowAlarmThresholdDesc     *prometheus.Desc
	laserRxOpticalPowerHighWarnThresholdDesc     *prometheus.Desc
	laserRxOpticalPowerLowWarnThresholdDesc      *prometheus.Desc
	laserRxOpticalPowerDbmDesc                   *prometheus.Desc
	laserRxOpticalPowerHighAlarmThresholdDbmDesc *prometheus.Desc
	laserRxOpticalPowerLowAlarmThresholdDbmDesc  *prometheus.Desc
	laserRxOpticalPowerHighWarnThresholdDbmDesc  *prometheus.Desc
	laserRxOpticalPowerLowWarnThresholdDbmDesc   *prometheus.Desc
	moduleVoltageDesc                            *prometheus.Desc
	moduleVoltageHighAlarmThresholdDesc          *prometheus.Desc
	moduleVoltageLowAlarmThresholdDesc           *prometheus.Desc
	moduleVoltageHighWarnThresholdDesc           *prometheus.Desc
	moduleVoltageLowWarnThresholdDesc            *prometheus.Desc
	rxSignalAvgOpticalPowerDesc                  *prometheus.Desc
	rxSignalAvgOpticalPowerDbmDesc               *prometheus.Desc
	transceiverDesc                              *prometheus.Desc
}

func newDescriptions(dynLabels dynamiclabels.Labels) *description {
	d := &description{}

	l := []string{"target", "name"}
	l = append(l, dynLabels.Keys()...)

	d.moduleVoltageDesc = prometheus.NewDesc(prefix+"module_voltage", "Module voltage", l, nil)
	d.moduleVoltageHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_high_alarm_threshold", "Module voltage high alarm threshold", l, nil)
	d.moduleVoltageLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_low_alarm_threshold", "Module voltage low alarm threshold", l, nil)
	d.moduleVoltageHighWarnThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_high_warn_threshold", "Module voltage high warn threshold", l, nil)
	d.moduleVoltageLowWarnThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_low_warn_threshold", "Module voltage low warn threshold", l, nil)

	d.moduleTemperatureDesc = prometheus.NewDesc(prefix+"temp", "Module temperature in degrees Celsius", l, nil)
	d.moduleTemperatureHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"temp_high_alarm_threshold", "Module temperature high alarm threshold in degrees Celsius", l, nil)
	d.moduleTemperatureLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"temp_low_alarm_threshold", "Module temperature low alarm threshold in degrees Celsius", l, nil)
	d.moduleTemperatureHighWarnThresholdDesc = prometheus.NewDesc(prefix+"temp_high_warn_threshold", "Module temperature high warn threshold in degrees Celsius", l, nil)
	d.moduleTemperatureLowWarnThresholdDesc = prometheus.NewDesc(prefix+"temp_low_warn_threshold", "Module temperature low warn threshold in degrees Celsius", l, nil)

	d.rxSignalAvgOpticalPowerDesc = prometheus.NewDesc(prefix+"rx_signal_avg", "Receiver signal average optical power in mW", l, nil)
	d.rxSignalAvgOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"rx_signal_avg_dbm", "Receiver signal average optical power in mW", l, nil)

	l = append(l, "lane")
	d.laserBiasCurrentDesc = prometheus.NewDesc(prefix+"laser_bias", "Laser bias current in mA", l, nil)
	d.laserBiasCurrentHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_high_alarm_threshold", "Laser bias current high alarm threshold", l, nil)
	d.laserBiasCurrentLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_low_alarm_threshold", "Laser bias current low alarm threshold", l, nil)
	d.laserBiasCurrentHighWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_high_warn_threshold", "Laser bias current high warn threshold", l, nil)
	d.laserBiasCurrentLowWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_low_warn_threshold", "Laser bias current low warn threshold", l, nil)
	d.laserOutputPowerDesc = prometheus.NewDesc(prefix+"laser_output", "Laser output power in mW", l, nil)
	d.laserOutputPowerHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_output_high_alarm_threshold", "Laser output power high alarm threshold in mW", l, nil)
	d.laserOutputPowerLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_output_low_alarm_threshold", "Laser output power low alarm threshold in mW", l, nil)
	d.laserOutputPowerHighWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_output_high_warn_threshold", "Laser output power high warn threshold in mW", l, nil)
	d.laserOutputPowerLowWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_output_low_warn_threshold", "Laser output power low warn threshold in mW", l, nil)

	d.laserOutputPowerDbmDesc = prometheus.NewDesc(prefix+"laser_output_dbm", "Laser output power in dBm", l, nil)
	d.laserOutputPowerHighAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_high_alarm_threshold_dbm", "Laser output power high alarm threshold in dBm", l, nil)
	d.laserOutputPowerLowAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_low_alarm_threshold_dbm", "Laser output power low alarm threshold in dBm", l, nil)
	d.laserOutputPowerHighWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_high_warn_threshold_dbm", "Laser output power high warn threshold in dBm", l, nil)
	d.laserOutputPowerLowWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_low_warn_threshold_dbm", "Laser output power low warn threshold in dBm", l, nil)

	d.laserRxOpticalPowerDesc = prometheus.NewDesc(prefix+"laser_rx", "Laser rx power in mW", l, nil)
	d.laserRxOpticalPowerHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_high_alarm_threshold", "Laser rx power high alarm threshold in mW", l, nil)
	d.laserRxOpticalPowerLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_low_alarm_threshold", "Laser rx power low alarm threshold in mW", l, nil)
	d.laserRxOpticalPowerHighWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_high_warn_threshold", "Laser rx power high warn threshold in mW", l, nil)
	d.laserRxOpticalPowerLowWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_low_warn_threshold", "Laser rx power low warn threshold in mW", l, nil)

	d.laserRxOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"laser_rx_dbm", "Laser rx power in dBm", l, nil)
	d.laserRxOpticalPowerHighAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_high_alarm_threshold_dbm", "Laser rx power high alarm threshold_dbm in dBm", l, nil)
	d.laserRxOpticalPowerLowAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_low_alarm_threshold_dbm", "Laser rx power low alarm threshold_dbm in dBm", l, nil)
	d.laserRxOpticalPowerHighWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_high_warn_threshold_dbm", "Laser rx power high warn threshold_dbm in dBm", l, nil)
	d.laserRxOpticalPowerLowWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_low_warn_threshold_dbm", "Laser rx power low warn threshold_dbm in dBm", l, nil)

	transceiver_labels := []string{"target", "name", "serial_number", "description", "speed", "fiber_type", "vendor_name", "vendor_part_number", "wavelength"}
	d.transceiverDesc = prometheus.NewDesc("junos_interface_transceiver", "Transceiver Info", transceiver_labels, nil)

	return d
}

type interfaceDiagnosticsCollector struct {
	descriptionRe *regexp.Regexp
}

// NewCollector creates a new collector
func NewCollector(descriptionRe *regexp.Regexp) collector.RPCCollector {
	c := &interfaceDiagnosticsCollector{
		descriptionRe: descriptionRe,
	}

	return c
}

// Name returns the name of the collector
func (*interfaceDiagnosticsCollector) Name() string {
	return "Interface Diagnostics"
}

// Describe describes the metrics
func (c *interfaceDiagnosticsCollector) Describe(ch chan<- *prometheus.Desc) {
	d := newDescriptions(nil)

	ch <- d.laserBiasCurrentDesc
	ch <- d.laserBiasCurrentHighAlarmThresholdDesc
	ch <- d.laserBiasCurrentLowAlarmThresholdDesc
	ch <- d.laserBiasCurrentHighWarnThresholdDesc
	ch <- d.laserBiasCurrentLowWarnThresholdDesc
	ch <- d.laserOutputPowerDesc
	ch <- d.laserOutputPowerHighAlarmThresholdDesc
	ch <- d.laserOutputPowerLowAlarmThresholdDesc
	ch <- d.laserOutputPowerHighWarnThresholdDesc
	ch <- d.laserOutputPowerLowWarnThresholdDesc
	ch <- d.laserOutputPowerDbmDesc
	ch <- d.laserOutputPowerHighAlarmThresholdDbmDesc
	ch <- d.laserOutputPowerLowAlarmThresholdDbmDesc
	ch <- d.laserOutputPowerHighWarnThresholdDbmDesc
	ch <- d.laserOutputPowerLowWarnThresholdDbmDesc
	ch <- d.moduleTemperatureDesc
	ch <- d.moduleTemperatureHighAlarmThresholdDesc
	ch <- d.moduleTemperatureLowAlarmThresholdDesc
	ch <- d.moduleTemperatureHighWarnThresholdDesc
	ch <- d.moduleTemperatureLowWarnThresholdDesc

	ch <- d.laserRxOpticalPowerDesc
	ch <- d.laserRxOpticalPowerHighAlarmThresholdDesc
	ch <- d.laserRxOpticalPowerLowAlarmThresholdDesc
	ch <- d.laserRxOpticalPowerHighWarnThresholdDesc
	ch <- d.laserRxOpticalPowerLowWarnThresholdDesc
	ch <- d.laserRxOpticalPowerDbmDesc
	ch <- d.laserRxOpticalPowerHighAlarmThresholdDbmDesc
	ch <- d.laserRxOpticalPowerLowAlarmThresholdDbmDesc
	ch <- d.laserRxOpticalPowerHighWarnThresholdDbmDesc
	ch <- d.laserRxOpticalPowerLowWarnThresholdDbmDesc

	ch <- d.moduleVoltageDesc
	ch <- d.moduleVoltageHighAlarmThresholdDesc
	ch <- d.moduleVoltageLowAlarmThresholdDesc
	ch <- d.moduleVoltageHighWarnThresholdDesc
	ch <- d.moduleVoltageLowWarnThresholdDesc

	ch <- d.rxSignalAvgOpticalPowerDesc
	ch <- d.rxSignalAvgOpticalPowerDbmDesc

	ch <- d.transceiverDesc
}

// Collect collects metrics from JunOS
func (c *interfaceDiagnosticsCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	diagnostics, err := c.interfaceDiagnostics(client)
	if err != nil {
		return err
	}

	// add satellite details if feature is enabled
	if client.IsSatelliteEnabled() {
		diagnosticsSatellite, err := c.interfaceDiagnosticsSatellite(client)
		if err != nil {
			return err
		}
		diagnostics = append(diagnostics, diagnosticsSatellite...)
	}

	diagnosticsDict := make(map[string]*interfaceDiagnostics)

	ifMediaDict, err := c.interfaceMediaInfo(client)
	if err != nil {
		return err
	}

	for _, diag := range diagnostics {
		index := strings.Split(diag.Name, "-")[1]
		diagnosticsDict[index] = diag

		desc := ""
		media := ifMediaDict[slotIndex(diag.Name)]
		if media != nil {
			desc = media.Description
		}

		dynLabels := dynamiclabels.ParseDescription(desc, c.descriptionRe)
		d := newDescriptions(dynLabels)

		l := append(labelValues, diag.Name)
		l = append(l, dynLabels.Values()...)

		ch <- prometheus.MustNewConstMetric(d.moduleTemperatureDesc, prometheus.GaugeValue, diag.ModuleTemperature, l...)
		ch <- prometheus.MustNewConstMetric(d.moduleTemperatureHighAlarmThresholdDesc, prometheus.GaugeValue, diag.ModuleTemperatureHighAlarmThreshold, l...)
		ch <- prometheus.MustNewConstMetric(d.moduleTemperatureLowAlarmThresholdDesc, prometheus.GaugeValue, diag.ModuleTemperatureLowAlarmThreshold, l...)
		ch <- prometheus.MustNewConstMetric(d.moduleTemperatureHighWarnThresholdDesc, prometheus.GaugeValue, diag.ModuleTemperatureHighWarnThreshold, l...)
		ch <- prometheus.MustNewConstMetric(d.moduleTemperatureLowWarnThresholdDesc, prometheus.GaugeValue, diag.ModuleTemperatureLowWarnThreshold, l...)

		if diag.ModuleVoltage > 0 {
			ch <- prometheus.MustNewConstMetric(d.moduleVoltageDesc, prometheus.GaugeValue, diag.ModuleVoltage, l...)
			ch <- prometheus.MustNewConstMetric(d.moduleVoltageHighAlarmThresholdDesc, prometheus.GaugeValue, diag.ModuleVoltageHighAlarmThreshold, l...)
			ch <- prometheus.MustNewConstMetric(d.moduleVoltageLowAlarmThresholdDesc, prometheus.GaugeValue, diag.ModuleVoltageLowAlarmThreshold, l...)
			ch <- prometheus.MustNewConstMetric(d.moduleVoltageHighWarnThresholdDesc, prometheus.GaugeValue, diag.ModuleVoltageHighWarnThreshold, l...)
			ch <- prometheus.MustNewConstMetric(d.moduleVoltageLowWarnThresholdDesc, prometheus.GaugeValue, diag.ModuleVoltageLowWarnThreshold, l...)
		}

		if diag.RxSignalAvgOpticalPower > 0 {
			ch <- prometheus.MustNewConstMetric(d.rxSignalAvgOpticalPowerDesc, prometheus.GaugeValue, diag.RxSignalAvgOpticalPower, l...)
			ch <- prometheus.MustNewConstMetric(d.rxSignalAvgOpticalPowerDbmDesc, prometheus.GaugeValue, diag.RxSignalAvgOpticalPowerDbm, l...)
		}

		var data []*interfaceDiagnostics
		if len(diag.Lanes) > 0 {
			data = diag.Lanes
		} else {
			data = []*interfaceDiagnostics{diag}
		}

		for _, e := range data {
			l2 := append(l, e.Index)
			ch <- prometheus.MustNewConstMetric(d.laserBiasCurrentDesc, prometheus.GaugeValue, e.LaserBiasCurrent, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserBiasCurrentHighAlarmThresholdDesc, prometheus.GaugeValue, diag.LaserBiasCurrentHighAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserBiasCurrentLowAlarmThresholdDesc, prometheus.GaugeValue, diag.LaserBiasCurrentLowAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserBiasCurrentHighWarnThresholdDesc, prometheus.GaugeValue, diag.LaserBiasCurrentHighWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserBiasCurrentLowWarnThresholdDesc, prometheus.GaugeValue, diag.LaserBiasCurrentLowWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerDesc, prometheus.GaugeValue, e.LaserOutputPower, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerHighAlarmThresholdDesc, prometheus.GaugeValue, diag.LaserOutputPowerHighAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerLowAlarmThresholdDesc, prometheus.GaugeValue, diag.LaserOutputPowerLowAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerHighWarnThresholdDesc, prometheus.GaugeValue, diag.LaserOutputPowerHighWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerLowWarnThresholdDesc, prometheus.GaugeValue, diag.LaserOutputPowerLowWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerDbmDesc, prometheus.GaugeValue, e.LaserOutputPowerDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerHighAlarmThresholdDbmDesc, prometheus.GaugeValue, diag.LaserOutputPowerHighAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerLowAlarmThresholdDbmDesc, prometheus.GaugeValue, diag.LaserOutputPowerLowAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerHighWarnThresholdDbmDesc, prometheus.GaugeValue, diag.LaserOutputPowerHighWarnThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserOutputPowerLowWarnThresholdDbmDesc, prometheus.GaugeValue, diag.LaserOutputPowerLowWarnThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerDesc, prometheus.GaugeValue, e.LaserRxOpticalPower, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerHighAlarmThresholdDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerHighAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerLowAlarmThresholdDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerLowAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerHighWarnThresholdDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerHighWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerLowWarnThresholdDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerLowWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerDbmDesc, prometheus.GaugeValue, e.LaserRxOpticalPowerDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerHighAlarmThresholdDbmDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerHighAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerLowAlarmThresholdDbmDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerLowAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerHighWarnThresholdDbmDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerHighWarnThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(d.laserRxOpticalPowerLowWarnThresholdDbmDesc, prometheus.GaugeValue, diag.LaserRxOpticalPowerLowWarnThresholdDbm, l2...)
		}
	}

	err = c.createTransceiverMetrics(client, ch, labelValues, ifMediaDict)
	if err != nil {
		return err
	}

	return nil
}

func (c *interfaceDiagnosticsCollector) interfaceMediaInfo(client collector.Client) (map[string]*physicalInterface, error) {
	var x = interfacesMediaStruct{}
	err := client.RunCommandAndParse("show interfaces media", &x)
	if err != nil {
		return nil, err
	}

	return interfaceMediaInfoFromRPCResult(&x.InterfaceInformation.PhysicalInterface), nil
}

func (c *interfaceDiagnosticsCollector) chassisHardwareInfos(client collector.Client) ([]*transceiverInformation, error) {
	var x = chassisHardware{}
	err := client.RunCommandAndParse("show chassis hardware", &x)
	if err != nil {
		return nil, err
	}

	return c.transceiverInfoFromRPCResult(client, x)
}

func (c *interfaceDiagnosticsCollector) transceiverInfoFromRPCResult(client collector.Client, chassisHardware chassisHardware) ([]*transceiverInformation, error) {
	transceiverList := make([]*transceiverInformation, 0)

	var chassisModules = chassisHardware.ChassisInventory.Chassis.ChassisModule
	for _, module := range chassisModules {
		if strings.Split(module.Name, " ")[0] != "FPC" {
			continue
		}
		for _, subModule := range module.ChassisSubModule {
			if strings.Split(subModule.Name, " ")[0] != "PIC" {
				continue
			}
			fpc := strings.Split(module.Name, " ")[1]
			pic := strings.Split(subModule.Name, " ")[1]

			picPortsInformation, err := c.getPicPortsFromRPCResult(client, fpc, pic)
			if err != nil {
				return nil, err
			}

			for port, subSubModule := range subModule.ChassisSubSubModule {
				port_name := strings.Split(subSubModule.Name, " ")[1]
				subSubModule_pointer := subSubModule
				id := fpc + "/" + pic + "/" + port_name
				transceiver := transceiverInformation{
					Name:                id,
					ChassisHardwareInfo: &subSubModule_pointer,
					PicPort:             &picPortsInformation[port],
				}
				transceiverList = append(transceiverList, &transceiver)
			}
		}
	}

	return transceiverList, nil
}

func (c *interfaceDiagnosticsCollector) getPicPortsFromRPCResult(client collector.Client, fpc string, pic string) ([]picPort, error) {
	var x = fpcInformationStruct{}
	command := fmt.Sprintf("show chassis pic fpc-slot %s pic-slot %s", fpc, pic)
	err := client.RunCommandAndParse(command, &x)
	if err != nil {
		return nil, err
	}

	return x.FPCInformation.FPC.PicDetail.PicPortInfoList, nil
}

func (c *interfaceDiagnosticsCollector) createTransceiverMetrics(client collector.Client, ch chan<- prometheus.Metric, labelValues []string, ifMediaDict map[string]*physicalInterface) error {
	transceiverInfo, err := c.chassisHardwareInfos(client)
	if err != nil {
		return err
	}

	for _, t := range transceiverInfo {
		chassisInfo := t.ChassisHardwareInfo
		port_speed := "0"
		oper_status := 0.0

		if media, hit := ifMediaDict[t.Name]; hit {
			if media.OperStatus == "up" {
				oper_status = 1.0
			}
			t.Name = media.Name
			port_speed = media.Speed
		} else {
			t.Name = "slot-" + t.Name
		}

		transceiver_labels := append(labelValues, t.Name, chassisInfo.SerialNumber, chassisInfo.Description, port_speed, t.PicPort.FiberMode, strings.TrimSpace(t.PicPort.SFPVendorName), strings.TrimSpace(t.PicPort.SFPVendorPno), t.PicPort.Wavelength)

		d := newDescriptions(nil)

		ch <- prometheus.MustNewConstMetric(d.transceiverDesc, prometheus.GaugeValue, oper_status, transceiver_labels...)
	}

	return nil
}

func (c *interfaceDiagnosticsCollector) interfaceDiagnosticsSatellite(client collector.Client) ([]*interfaceDiagnostics, error) {
	var x = result{}

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
			lines     = strings.Split(string(b[:]), "\n")
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

func (c *interfaceDiagnosticsCollector) interfaceDiagnostics(client collector.Client) ([]*interfaceDiagnostics, error) {
	var x = result{}
	err := client.RunCommandAndParse("show interfaces diagnostics optics", &x)
	if err != nil {
		return nil, err
	}

	return interfaceDiagnosticsFromRPCResult(x), nil
}
