package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/czerwonk/junos_exporter/config"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/interfacelabels"
)

func TestCollectorsRegistered(t *testing.T) {
	c := &config.Config{
		Features: config.FeatureConfig{
			Alarm:               true,
			Environment:         true,
			BGP:                 true,
			OSPF:                true,
			ISIS:                true,
			NAT:                 true,
			L2Circuit:           true,
			LDP:                 true,
			Routes:              true,
			RoutingEngine:       true,
			Firewall:            true,
			Interfaces:          true,
			InterfaceDiagnostic: true,
			InterfaceQueue:      true,
			Storage:             true,
			Accounting:          true,
			IPSec:               true,
			FPC:                 true,
			RPKI:                true,
		},
	}

	cols := collectorsForDevices([]*connector.Device{&connector.Device{
		Host: "::1",
	}}, c, "", interfacelabels.NewDynamicLabels())

	assert.Equal(t, 19, len(cols.collectors), "collector count")
}

func TestCollectorsForDevices(t *testing.T) {
	c := &config.Config{
		Features: config.FeatureConfig{
			Alarm:               true,
			Environment:         true,
			BGP:                 true,
			OSPF:                true,
			ISIS:                true,
			NAT:                 true,
			L2Circuit:           true,
			LDP:                 true,
			Routes:              true,
			RoutingEngine:       true,
			Firewall:            true,
			Interfaces:          true,
			InterfaceDiagnostic: true,
			InterfaceQueue:      true,
			Storage:             true,
			Accounting:          true,
			IPSec:               true,
			FPC:                 true,
			RPKI:                true,
		},
		Devices: []*config.DeviceConfig{
			&config.DeviceConfig{
				Host: "2001:678:1e0::1",
			},
			&config.DeviceConfig{
				Host: "2001:678:1e0::2",
				Features: &config.FeatureConfig{
					Interfaces: true,
				},
			},
		},
	}

	d1 := &connector.Device{
		Host: "2001:678:1e0::1",
	}
	d2 := &connector.Device{
		Host: "2001:678:1e0::2",
	}
	cols := collectorsForDevices([]*connector.Device{d1, d2}, c, "", interfacelabels.NewDynamicLabels())

	assert.Equal(t, 19, len(cols.collectorsForDevice(d1)), "device 1 collector count")

	cd2 := cols.collectorsForDevice(d2)
	assert.Equal(t, 1, len(cd2), "device 2 collector count")
	assert.Equal(t, "Interfaces", cd2[0].Name(), "device 2 collector name")
}
