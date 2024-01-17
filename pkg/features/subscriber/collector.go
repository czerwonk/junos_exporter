// SPDX-License-Identifier: MIT

package subscriber

import (
	"errors"
	"fmt"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_subscriber_info"

var subscriberInfo *prometheus.Desc

func init() {
	l := []string{"target", "interface", "agent_circuit_id", "agent_remote_id", "underlying_ifd"}
	subscriberInfo = prometheus.NewDesc(prefix+"", "Subscriber Detail", l, nil)
}

// Name implements collector.RPCCollector.
func (*subcsribers_information) Name() string {
	return "Subscriber Detail"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &subcsribers_information{}
}

// Describe describes the metrics
func (*subcsribers_information) Describe(ch chan<- *prometheus.Desc) {
	ch <- subscriberInfo
}

// Collect collects metrics from JunOS
func (c *subcsribers_information) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = subcsribers_information{}
	err := client.RunCommandAndParse("show subscribers client-type dhcp detail", &x) //TODO: see if client-type dhcp can be left out
	if err != nil {
		return err
	}

	logicalInterfaceMap, err := getLogicalInterfaceInformation(client)
	if err != nil {
		return err
	}

	for _, subscriber := range x.SubscribersInformation.Subscriber {
		underlying_interface, err := findUnderlyingInterface(client, subscriber.UnderlyingInterface, logicalInterfaceMap, 2)
		if err != nil {
			fmt.Println(err)
		}

		labels := append(labelValues, subscriber.Interface, subscriber.AgentCircuitId, subscriber.AgentRemoteId, underlying_interface)
		ch <- prometheus.MustNewConstMetric(subscriberInfo, prometheus.CounterValue, 1, labels...)
	}
	return nil
}

func getLogicalInterfaceInformation(client collector.Client) (map[string]string, error) {

	var interfaceInformation = &InterfaceInformation{}
	var interfaceMap = make(map[string]string)

	err := client.RunCommandAndParse("show interfaces demux0 brief", interfaceInformation)
	if err != nil {
		return nil, err
	}

	for _, logicalInterface := range interfaceInformation.LogicalInterfaces {
		interfaceMap[logicalInterface.Name] = logicalInterface.DemuxUnderlyingIfName
	}

	return interfaceMap, nil
}

func findUnderlyingInterface(client collector.Client, ifName string, logicalIfMap map[string]string, maxDepth int) (string, error) {

	if !(strings.HasPrefix(ifName, "demux")) {
		return ifName, nil
	}

	if maxDepth < 0 {
		return "", errors.New("no underlying interface found, max treshold reached")
	}

	logicalIfName, exists := logicalIfMap[ifName]
	if !exists {
		return "", errors.New("no underlying interface found")
	}

	return findUnderlyingInterface(client, logicalIfName, logicalIfMap, maxDepth-1)

}
