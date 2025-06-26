// SPDX-License-Identifier: MIT

package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

type Client interface {
	// RunCommandAndParse runs a command on JunOS and unmarshals the XML result
	RunCommandAndParse(cmd string, obj interface{}) error

	// RunCommandAndParseWithParser runs a command on JunOS and unmarshals the XML result using the specified parser function
	RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error

	// IsSatelliteEnabled returns if sattelite features are enabled on the device
	IsSatelliteEnabled() bool

	IsScrapingLicenseEnabled() bool

	// Device returns device information for the connected device
	Device() *connector.Device

	// Ctx returns the context the client is running in
	Context() context.Context
}

// RPCCollector collects metrics from JunOS using rpc.Client
type RPCCollector interface {
	// Name returns an human readable name for logging and debugging purposes
	Name() string

	// Describe describes the metrics
	Describe(ch chan<- *prometheus.Desc)

	// Collect collects metrics from JunOS
	Collect(client Client, ch chan<- prometheus.Metric, labelValues []string) error
}
