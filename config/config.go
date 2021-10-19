package config

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration for the exporter
type Config struct {
	Password  string          `yaml:"password"`
	Targets   []string        `yaml:"targets,omitempty"`
	Devices   []*DeviceConfig `yaml:"devices,omitempty"`
	Features  FeatureConfig   `yaml:"features,omitempty"`
	LSEnabled bool            `yaml:"logical_systems,omitempty"`
	IfDescReg string          `yaml:"interface_description_regex,omitempty"`
}

// DeviceConfig is the config representation of 1 device
type DeviceConfig struct {
	Host      string         `yaml:"host"`
	Username  string         `yaml:"username,omitempty"`
	Password  string         `yaml:"password,omitempty"`
	KeyFile   string         `yaml:"key_file,omitempty"`
	Features  *FeatureConfig `yaml:"features,omitempty"`
	IfDescReg string         `yaml:"interface_description_regex,omitempty"`
}

// FeatureConfig is the list of collectors enabled or disabled
type FeatureConfig struct {
	Alarm               bool `yaml:"alarm,omitempty"`
	Environment         bool `yaml:"environment,omitempty"`
	BGP                 bool `yaml:"bgp,omitempty"`
	OSPF                bool `yaml:"ospf,omitempty"`
	ISIS                bool `yaml:"isis,omitempty"`
	NAT                 bool `yaml:"nat,omitempty"`
	L2Circuit           bool `yaml:"l2circuit,omitempty"`
	LDP                 bool `yaml:"ldp,omitempty"`
	Routes              bool `yaml:"routes,omitempty"`
	RoutingEngine       bool `yaml:"routing_engine,omitempty"`
	Firewall            bool `yaml:"firewall,omitempty"`
	Interfaces          bool `yaml:"interfaces,omitempty"`
	InterfaceDiagnostic bool `yaml:"interface_diagnostic,omitempty"`
	InterfaceQueue      bool `yaml:"interface_queue,omitempty"`
	Storage             bool `yaml:"storage,omitempty"`
	Accounting          bool `yaml:"accounting,omitempty"`
	IPSec               bool `yaml:"ipsec,omitempty"`
	FPC                 bool `yaml:"fpc,omitempty"`
	RPKI                bool `yaml:"rpki,omitempty"`
	RPM                 bool `yaml:"rpm,omitempty"`
	Satellite           bool `yaml:"satellite,omitempty"`
	System              bool `yaml:"system,omitempty"`
	Power               bool `yaml:"power,omitempty"`
	MAC                 bool `yaml:"mac,omitempty"`
	VRRP                bool `yaml:"vrrp,omitempty"`
}

// New creates a new config
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
	c.Password = ""
	c.LSEnabled = false
	c.IfDescReg = ""
	f := &c.Features
	f.Alarm = true
	f.BGP = true
	f.Environment = true
	f.Interfaces = true
	f.InterfaceDiagnostic = true
	f.InterfaceQueue = true
	f.IPSec = false
	f.OSPF = true
	f.ISIS = true
	f.LDP = true
	f.Routes = true
	f.Firewall = true
	f.RoutingEngine = true
	f.Storage = false
	f.Accounting = false
	f.FPC = false
	f.L2Circuit = false
	f.RPKI = false
	f.RPM = false
	f.Satellite = false
	f.Power = false
	f.MAC = false
	f.VRRP = false
}

// FeaturesForDevice gets the feature set configured for a device
func (c *Config) FeaturesForDevice(host string) *FeatureConfig {
	d := c.findDeviceConfig(host)

	if d != nil && d.Features != nil {
		return d.Features
	}

	return &c.Features
}

func (c *Config) findDeviceConfig(host string) *DeviceConfig {
	for _, dc := range c.Devices {
		if dc.Host == host {
			return dc
		}
	}

	return nil
}
