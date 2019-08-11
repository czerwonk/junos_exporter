package config

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldParse(t *testing.T) {
	b, err := ioutil.ReadFile("tests/config1.yml")
	if err != nil {
		t.Fatal(err)
	}

	c, err := Load(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	et := []string{"router1", "router2"}
	if !reflect.DeepEqual(c.Targets, et) {
		t.Fatalf("expected targets %v, got %v", et, c.Targets)
	}

	assertFeature("BGP", c.Features.BGP, true, t)
	assertFeature("OSPF", c.Features.OSPF, false, t)
	assertFeature("ISIS", c.Features.ISIS, true, t)
	assertFeature("Routes", c.Features.Routes, true, t)
	assertFeature("RoutingEngine", c.Features.RoutingEngine, true, t)
	assertFeature("Environment", c.Features.Environment, false, t)
	assertFeature("Firewall", c.Features.Firewall, false, t)
	assertFeature("InterfaceDiagnostic", c.Features.InterfaceDiagnostic, false, t)
	assertFeature("Interfacs", c.Features.Interfaces, false, t)
	assertFeature("L2Circuit", c.Features.L2Circuit, true, t)
	assertFeature("Storage", c.Features.Storage, false, t)
	assertFeature("FPC", c.Features.FPC, true, t)
}

func TestShouldUseDefaults(t *testing.T) {
	b, err := ioutil.ReadFile("tests/config2.yml")
	if err != nil {
		t.Fatal(err)
	}

	c, err := Load(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	assertFeature("BGP", c.Features.BGP, true, t)
	assertFeature("OSPF", c.Features.OSPF, true, t)
	assertFeature("ISIS", c.Features.ISIS, false, t)
	assertFeature("Routes", c.Features.Routes, true, t)
	assertFeature("RoutingEngine", c.Features.RoutingEngine, true, t)
	assertFeature("Environment", c.Features.Environment, true, t)
	assertFeature("Firewall", c.Features.Firewall, false, t)
	assertFeature("InterfaceDiagnostic", c.Features.InterfaceDiagnostic, true, t)
	assertFeature("Interfaces", c.Features.Interfaces, true, t)
	assertFeature("L2Circuit", c.Features.L2Circuit, false, t)
	assertFeature("Storage", c.Features.Storage, true, t)
	assertFeature("FPC", c.Features.FPC, false, t)
}

func TestShouldParseDevices(t *testing.T) {
	b, err := ioutil.ReadFile("tests/config3.yml")
	if err != nil {
		t.Fatal(err)
	}

	c, err := Load(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(c.Devices), "devices")

	d1 := c.Devices[0]
	assert.Equal(t, "router1", d1.Host, "Device 1: Host")
	assert.Equal(t, "keyfile_user", d1.Username, "Device 1: Username")
	assert.Equal(t, "/path/to/key", d1.KeyFile, "Device 1: Keyfile")

	d2 := c.Devices[1]
	assert.Equal(t, "router2", d2.Host, "Device 2: Host")
	assert.Equal(t, "password_user", d2.Username, "Device 2: Username")
	assert.Equal(t, "secret", d2.Password, "Device 2: Password")
}

func assertFeature(name string, actual, expected bool, t *testing.T) {
	if actual != expected {
		t.Fatalf("feature %s should be %v, but is %v", name, expected, actual)
	}
}
