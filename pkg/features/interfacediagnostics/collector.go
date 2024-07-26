// SPDX-License-Identifier: MIT

package interfacediagnostics

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/interfacelabels"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_interface_diagnostics_"

type interfaceDiagnosticsCollector struct {
	labels                                 *interfacelabels.DynamicLabelManager
	laserBiasCurrentDesc                   *prometheus.Desc
	laserBiasCurrentHighAlarmThresholdDesc *prometheus.Desc
	laserBiasCurrentLowAlarmThresholdDesc  *prometheus.Desc
	laserBiasCurrentHighWarnThresholdDesc  *prometheus.Desc
	laserBiasCurrentLowWarnThresholdDesc   *prometheus.Desc

	laserOutputPowerDesc                   *prometheus.Desc
	laserOutputPowerHighAlarmThresholdDesc *prometheus.Desc
	laserOutputPowerLowAlarmThresholdDesc  *prometheus.Desc
	laserOutputPowerHighWarnThresholdDesc  *prometheus.Desc
	laserOutputPowerLowWarnThresholdDesc   *prometheus.Desc

	laserOutputPowerDbmDesc                   *prometheus.Desc
	laserOutputPowerHighAlarmThresholdDbmDesc *prometheus.Desc
	laserOutputPowerLowAlarmThresholdDbmDesc  *prometheus.Desc
	laserOutputPowerHighWarnThresholdDbmDesc  *prometheus.Desc
	laserOutputPowerLowWarnThresholdDbmDesc   *prometheus.Desc

	moduleTemperatureDesc                   *prometheus.Desc
	moduleTemperatureHighAlarmThresholdDesc *prometheus.Desc
	moduleTemperatureLowAlarmThresholdDesc  *prometheus.Desc
	moduleTemperatureHighWarnThresholdDesc  *prometheus.Desc
	moduleTemperatureLowWarnThresholdDesc   *prometheus.Desc

	laserRxOpticalPowerDesc                   *prometheus.Desc
	laserRxOpticalPowerHighAlarmThresholdDesc *prometheus.Desc
	laserRxOpticalPowerLowAlarmThresholdDesc  *prometheus.Desc
	laserRxOpticalPowerHighWarnThresholdDesc  *prometheus.Desc
	laserRxOpticalPowerLowWarnThresholdDesc   *prometheus.Desc

	laserRxOpticalPowerDbmDesc                   *prometheus.Desc
	laserRxOpticalPowerHighAlarmThresholdDbmDesc *prometheus.Desc
	laserRxOpticalPowerLowAlarmThresholdDbmDesc  *prometheus.Desc
	laserRxOpticalPowerHighWarnThresholdDbmDesc  *prometheus.Desc
	laserRxOpticalPowerLowWarnThresholdDbmDesc   *prometheus.Desc

	moduleVoltageDesc                   *prometheus.Desc
	moduleVoltageHighAlarmThresholdDesc *prometheus.Desc
	moduleVoltageLowAlarmThresholdDesc  *prometheus.Desc
	moduleVoltageHighWarnThresholdDesc  *prometheus.Desc
	moduleVoltageLowWarnThresholdDesc   *prometheus.Desc

	rxSignalAvgOpticalPowerDesc    *prometheus.Desc
	rxSignalAvgOpticalPowerDbmDesc *prometheus.Desc

	transceiverDesc *prometheus.Desc
}

// NewCollector creates a new collector
func NewCollector(labels *interfacelabels.DynamicLabelManager) collector.RPCCollector {
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
	c.moduleVoltageHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_high_alarm_threshold", "Module voltage high alarm threshold", l, nil)
	c.moduleVoltageLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_low_alarm_threshold", "Module voltage low alarm threshold", l, nil)
	c.moduleVoltageHighWarnThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_high_warn_threshold", "Module voltage high warn threshold", l, nil)
	c.moduleVoltageLowWarnThresholdDesc = prometheus.NewDesc(prefix+"module_voltage_low_warn_threshold", "Module voltage low warn threshold", l, nil)

	c.moduleTemperatureDesc = prometheus.NewDesc(prefix+"temp", "Module temperature in degrees Celsius", l, nil)
	c.moduleTemperatureHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"temp_high_alarm_threshold", "Module temperature high alarm threshold in degrees Celsius", l, nil)
	c.moduleTemperatureLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"temp_low_alarm_threshold", "Module temperature low alarm threshold in degrees Celsius", l, nil)
	c.moduleTemperatureHighWarnThresholdDesc = prometheus.NewDesc(prefix+"temp_high_warn_threshold", "Module temperature high warn threshold in degrees Celsius", l, nil)
	c.moduleTemperatureLowWarnThresholdDesc = prometheus.NewDesc(prefix+"temp_low_warn_threshold", "Module temperature low warn threshold in degrees Celsius", l, nil)

	c.rxSignalAvgOpticalPowerDesc = prometheus.NewDesc(prefix+"rx_signal_avg", "Receiver signal average optical power in mW", l, nil)
	c.rxSignalAvgOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"rx_signal_avg_dbm", "Receiver signal average optical power in mW", l, nil)

	l = append(l, "lane")
	c.laserBiasCurrentDesc = prometheus.NewDesc(prefix+"laser_bias", "Laser bias current in mA", l, nil)
	c.laserBiasCurrentHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_high_alarm_threshold", "Laser bias current high alarm threshold", l, nil)
	c.laserBiasCurrentLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_low_alarm_threshold", "Laser bias current low alarm threshold", l, nil)
	c.laserBiasCurrentHighWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_high_warn_threshold", "Laser bias current high warn threshold", l, nil)
	c.laserBiasCurrentLowWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_bias_low_warn_threshold", "Laser bias current low warn threshold", l, nil)
	c.laserOutputPowerDesc = prometheus.NewDesc(prefix+"laser_output", "Laser output power in mW", l, nil)
	c.laserOutputPowerHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_output_high_alarm_threshold", "Laser output power high alarm threshold in mW", l, nil)
	c.laserOutputPowerLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_output_low_alarm_threshold", "Laser output power low alarm threshold in mW", l, nil)
	c.laserOutputPowerHighWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_output_high_warn_threshold", "Laser output power high warn threshold in mW", l, nil)
	c.laserOutputPowerLowWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_output_low_warn_threshold", "Laser output power low warn threshold in mW", l, nil)

	c.laserOutputPowerDbmDesc = prometheus.NewDesc(prefix+"laser_output_dbm", "Laser output power in dBm", l, nil)
	c.laserOutputPowerHighAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_high_alarm_threshold_dbm", "Laser output power high alarm threshold in dBm", l, nil)
	c.laserOutputPowerLowAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_low_alarm_threshold_dbm", "Laser output power low alarm threshold in dBm", l, nil)
	c.laserOutputPowerHighWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_high_warn_threshold_dbm", "Laser output power high warn threshold in dBm", l, nil)
	c.laserOutputPowerLowWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_output_low_warn_threshold_dbm", "Laser output power low warn threshold in dBm", l, nil)

	c.laserRxOpticalPowerDesc = prometheus.NewDesc(prefix+"laser_rx", "Laser rx power in mW", l, nil)
	c.laserRxOpticalPowerHighAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_high_alarm_threshold", "Laser rx power high alarm threshold in mW", l, nil)
	c.laserRxOpticalPowerLowAlarmThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_low_alarm_threshold", "Laser rx power low alarm threshold in mW", l, nil)
	c.laserRxOpticalPowerHighWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_high_warn_threshold", "Laser rx power high warn threshold in mW", l, nil)
	c.laserRxOpticalPowerLowWarnThresholdDesc = prometheus.NewDesc(prefix+"laser_rx_low_warn_threshold", "Laser rx power low warn threshold in mW", l, nil)

	c.laserRxOpticalPowerDbmDesc = prometheus.NewDesc(prefix+"laser_rx_dbm", "Laser rx power in dBm", l, nil)
	c.laserRxOpticalPowerHighAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_high_alarm_threshold_dbm", "Laser rx power high alarm threshold_dbm in dBm", l, nil)
	c.laserRxOpticalPowerLowAlarmThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_low_alarm_threshold_dbm", "Laser rx power low alarm threshold_dbm in dBm", l, nil)
	c.laserRxOpticalPowerHighWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_high_warn_threshold_dbm", "Laser rx power high warn threshold_dbm in dBm", l, nil)
	c.laserRxOpticalPowerLowWarnThresholdDbmDesc = prometheus.NewDesc(prefix+"laser_rx_low_warn_threshold_dbm", "Laser rx power low warn threshold_dbm in dBm", l, nil)

	transceiver_labels := []string{"target", "name", "serial_number", "description", "speed", "fiber_type", "vendor_name", "vendor_part_number", "wavelength"}
	c.transceiverDesc = prometheus.NewDesc("junos_interface_transceiver", "Transceiver Info", transceiver_labels, nil)
}

// Describe describes the metrics
func (c *interfaceDiagnosticsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.laserBiasCurrentDesc
	ch <- c.laserBiasCurrentHighAlarmThresholdDesc
	ch <- c.laserBiasCurrentLowAlarmThresholdDesc
	ch <- c.laserBiasCurrentHighWarnThresholdDesc
	ch <- c.laserBiasCurrentLowWarnThresholdDesc
	ch <- c.laserOutputPowerDesc
	ch <- c.laserOutputPowerHighAlarmThresholdDesc
	ch <- c.laserOutputPowerLowAlarmThresholdDesc
	ch <- c.laserOutputPowerHighWarnThresholdDesc
	ch <- c.laserOutputPowerLowWarnThresholdDesc
	ch <- c.laserOutputPowerDbmDesc
	ch <- c.laserOutputPowerHighAlarmThresholdDbmDesc
	ch <- c.laserOutputPowerLowAlarmThresholdDbmDesc
	ch <- c.laserOutputPowerHighWarnThresholdDbmDesc
	ch <- c.laserOutputPowerLowWarnThresholdDbmDesc
	ch <- c.moduleTemperatureDesc
	ch <- c.moduleTemperatureHighAlarmThresholdDesc
	ch <- c.moduleTemperatureLowAlarmThresholdDesc
	ch <- c.moduleTemperatureHighWarnThresholdDesc
	ch <- c.moduleTemperatureLowWarnThresholdDesc

	ch <- c.laserRxOpticalPowerDesc
	ch <- c.laserRxOpticalPowerHighAlarmThresholdDesc
	ch <- c.laserRxOpticalPowerLowAlarmThresholdDesc
	ch <- c.laserRxOpticalPowerHighWarnThresholdDesc
	ch <- c.laserRxOpticalPowerLowWarnThresholdDesc
	ch <- c.laserRxOpticalPowerDbmDesc
	ch <- c.laserRxOpticalPowerHighAlarmThresholdDbmDesc
	ch <- c.laserRxOpticalPowerLowAlarmThresholdDbmDesc
	ch <- c.laserRxOpticalPowerHighWarnThresholdDbmDesc
	ch <- c.laserRxOpticalPowerLowWarnThresholdDbmDesc

	ch <- c.moduleVoltageDesc
	ch <- c.moduleVoltageHighAlarmThresholdDesc
	ch <- c.moduleVoltageLowAlarmThresholdDesc
	ch <- c.moduleVoltageHighWarnThresholdDesc
	ch <- c.moduleVoltageLowWarnThresholdDesc

	ch <- c.rxSignalAvgOpticalPowerDesc
	ch <- c.rxSignalAvgOpticalPowerDbmDesc

	ch <- c.transceiverDesc
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

	diagnostics_dict := make(map[string]*interfaceDiagnostics)

	for _, d := range diagnostics {
		index := strings.Split(d.Name, "-")[1]
		diagnostics_dict[index] = d

		l := append(labelValues, d.Name)
		l = append(l, c.labels.ValuesForInterface(client.Device(), d.Name)...)

		ch <- prometheus.MustNewConstMetric(c.moduleTemperatureDesc, prometheus.GaugeValue, d.ModuleTemperature, l...)
		ch <- prometheus.MustNewConstMetric(c.moduleTemperatureHighAlarmThresholdDesc, prometheus.GaugeValue, d.ModuleTemperatureHighAlarmThreshold, l...)
		ch <- prometheus.MustNewConstMetric(c.moduleTemperatureLowAlarmThresholdDesc, prometheus.GaugeValue, d.ModuleTemperatureLowAlarmThreshold, l...)
		ch <- prometheus.MustNewConstMetric(c.moduleTemperatureHighWarnThresholdDesc, prometheus.GaugeValue, d.ModuleTemperatureHighWarnThreshold, l...)
		ch <- prometheus.MustNewConstMetric(c.moduleTemperatureLowWarnThresholdDesc, prometheus.GaugeValue, d.ModuleTemperatureLowWarnThreshold, l...)

		if d.ModuleVoltage > 0 {
			ch <- prometheus.MustNewConstMetric(c.moduleVoltageDesc, prometheus.GaugeValue, d.ModuleVoltage, l...)
			ch <- prometheus.MustNewConstMetric(c.moduleVoltageHighAlarmThresholdDesc, prometheus.GaugeValue, d.ModuleVoltageHighAlarmThreshold, l...)
			ch <- prometheus.MustNewConstMetric(c.moduleVoltageLowAlarmThresholdDesc, prometheus.GaugeValue, d.ModuleVoltageLowAlarmThreshold, l...)
			ch <- prometheus.MustNewConstMetric(c.moduleVoltageHighWarnThresholdDesc, prometheus.GaugeValue, d.ModuleVoltageHighWarnThreshold, l...)
			ch <- prometheus.MustNewConstMetric(c.moduleVoltageLowWarnThresholdDesc, prometheus.GaugeValue, d.ModuleVoltageLowWarnThreshold, l...)
		}

		if d.RxSignalAvgOpticalPower > 0 {
			ch <- prometheus.MustNewConstMetric(c.rxSignalAvgOpticalPowerDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPower, l...)
			ch <- prometheus.MustNewConstMetric(c.rxSignalAvgOpticalPowerDbmDesc, prometheus.GaugeValue, d.RxSignalAvgOpticalPowerDbm, l...)
		}

		var data []*interfaceDiagnostics
		if len(d.Lanes) > 0 {
			data = d.Lanes
		} else {
			data = []*interfaceDiagnostics{d}
		}

		for _, e := range data {
			l2 := append(l, e.Index)
			ch <- prometheus.MustNewConstMetric(c.laserBiasCurrentDesc, prometheus.GaugeValue, e.LaserBiasCurrent, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserBiasCurrentHighAlarmThresholdDesc, prometheus.GaugeValue, d.LaserBiasCurrentHighAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserBiasCurrentLowAlarmThresholdDesc, prometheus.GaugeValue, d.LaserBiasCurrentLowAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserBiasCurrentHighWarnThresholdDesc, prometheus.GaugeValue, d.LaserBiasCurrentHighWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserBiasCurrentLowWarnThresholdDesc, prometheus.GaugeValue, d.LaserBiasCurrentLowWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerDesc, prometheus.GaugeValue, e.LaserOutputPower, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerHighAlarmThresholdDesc, prometheus.GaugeValue, d.LaserOutputPowerHighAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerLowAlarmThresholdDesc, prometheus.GaugeValue, d.LaserOutputPowerLowAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerHighWarnThresholdDesc, prometheus.GaugeValue, d.LaserOutputPowerHighWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerLowWarnThresholdDesc, prometheus.GaugeValue, d.LaserOutputPowerLowWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerDbmDesc, prometheus.GaugeValue, e.LaserOutputPowerDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerHighAlarmThresholdDbmDesc, prometheus.GaugeValue, d.LaserOutputPowerHighAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerLowAlarmThresholdDbmDesc, prometheus.GaugeValue, d.LaserOutputPowerLowAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerHighWarnThresholdDbmDesc, prometheus.GaugeValue, d.LaserOutputPowerHighWarnThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserOutputPowerLowWarnThresholdDbmDesc, prometheus.GaugeValue, d.LaserOutputPowerLowWarnThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerDesc, prometheus.GaugeValue, e.LaserRxOpticalPower, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerHighAlarmThresholdDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerHighAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerLowAlarmThresholdDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerLowAlarmThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerHighWarnThresholdDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerHighWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerLowWarnThresholdDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerLowWarnThreshold, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerDbmDesc, prometheus.GaugeValue, e.LaserRxOpticalPowerDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerHighAlarmThresholdDbmDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerHighAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerLowAlarmThresholdDbmDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerLowAlarmThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerHighWarnThresholdDbmDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerHighWarnThresholdDbm, l2...)
			ch <- prometheus.MustNewConstMetric(c.laserRxOpticalPowerLowWarnThresholdDbmDesc, prometheus.GaugeValue, d.LaserRxOpticalPowerLowWarnThresholdDbm, l2...)
		}
	}

	err = c.createTransceiverMetrics(client, ch, labelValues)
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
	var x = fPCInformationStruct{}
	command := fmt.Sprintf("show chassis pic fpc-slot %s pic-slot %s", fpc, pic)
	err := client.RunCommandAndParse(command, &x)
	if err != nil {
		return nil, err
	}

	return x.FPCInformation.FPC.PicDetail.PicPortInfoList, nil
}

func (c *interfaceDiagnosticsCollector) createTransceiverMetrics(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	ifMediaDict, err := c.interfaceMediaInfo(client)
	if err != nil {
		return err
	}

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

		ch <- prometheus.MustNewConstMetric(c.transceiverDesc, prometheus.GaugeValue, oper_status, transceiver_labels...)
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
