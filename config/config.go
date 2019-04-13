package config

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration for the exporter
type Config struct {
	Targets  []string `yaml:"targets"`
	Features struct {
		Environment         bool `yaml:"environment,omitempty"`
		BGP                 bool `yaml:"bgp,omitempty"`
		OSPF                bool `yaml:"ospf,omitempty"`
		Firewall            bool `yaml:"firewall,omitempty"`
		ISIS                bool `yaml:"isis,omitempty"`
		L2Circuit           bool `yaml:"l2circuit,omitempty"`
		Routes              bool `yaml:"routes,omitempty"`
		RoutingEngine       bool `yaml:"routing_engine,omitempty"`
		Interfaces          bool `yaml:"interfaces,omitempty"`
		InterfaceDiagnostic bool `yaml:"interface_diagnostic,omitempty"`
		Storage             bool `yaml:"storage,omitempty"`
	} `yaml:"features,omitempty"`
}

func New() *Config {
	c := &Config{
		Targets: make([]string, 0),
	}
	setDefaultValues(c)

	return c
}

// Load loads a config from reader
func Load(reader io.Reader) (*Config, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	c := New()
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func setDefaultValues(c *Config) {
	f := &c.Features
	f.BGP = true
	f.Environment = true
	f.Interfaces = true
	f.InterfaceDiagnostic = true
	f.OSPF = true
	f.Routes = true
	f.RoutingEngine = true
	f.Storage = true
}
