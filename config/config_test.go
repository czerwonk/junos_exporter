package config

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
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
	assertFeature("InterfaceDiagnostic", c.Features.InterfaceDiagnostic, false, t)
	assertFeature("Interfacs", c.Features.Interfaces, false, t)
	assertFeature("L2Circuit", c.Features.L2Circuit, true, t)
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
	assertFeature("InterfaceDiagnostic", c.Features.InterfaceDiagnostic, true, t)
	assertFeature("Interfaces", c.Features.Interfaces, true, t)
	assertFeature("L2Circuit", c.Features.L2Circuit, false, t)
}

func assertFeature(name string, actual, expected bool, t *testing.T) {
	if actual != expected {
		t.Fatalf("feature %s should be %v, but is %v", name, expected, actual)
	}
}
