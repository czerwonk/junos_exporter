// SPDX-License-Identifier: MIT

package config

import (
	"fmt"
	"io"
	"regexp"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration for the exporter
type Config struct {
	Password    string          `yaml:"password"`
	Targets     []string        `yaml:"targets,omitempty"`
	Devices     []*DeviceConfig `yaml:"devices,omitempty"`
	Features    FeatureConfig   `yaml:"features,omitempty"`
	LSEnabled   bool            `yaml:"logical_systems,omitempty"`
	IfDescReStr string          `yaml:"interface_description_regex,omitempty"`
	IfDescReg   *regexp.Regexp  `yaml:"-"`
}

func (c *Config) load(dynamicIfaceLabels bool) error {
	if c.IfDescReStr != "" && dynamicIfaceLabels {
		re, err := regexp.Compile(c.IfDescReStr)
		if err != nil {
			return fmt.Errorf("unable to compile interfce description regex %q: %w", c.IfDescReStr, err)
		}

		c.IfDescReg = re
	}

	for _, d := range c.Devices {
		if d.IfDescRegStr != "" && dynamicIfaceLabels {
			re, err := regexp.Compile(c.IfDescReStr)
			if err != nil {
				return fmt.Errorf("unable to compile interfce description regex %q: %w", c.IfDescReStr, err)
			}

			d.IfDescReg = re
		}
	}

	return nil
}

// DeviceConfig is the config representation of 1 device
type DeviceConfig struct {
	Host          string         `yaml:"host"`
	Username      string         `yaml:"username,omitempty"`
	Password      string         `yaml:"password,omitempty"`
	KeyFile       string         `yaml:"key_file,omitempty"`
	KeyPassphrase string         `yaml:"key_passphrase,omitempty"`
	Features      *FeatureConfig `yaml:"features,omitempty"`
	IfDescRegStr  string         `yaml:"interface_description_regex,omitempty"`
	IfDescReg     *regexp.Regexp `yaml:"-"`
	IsHostPattern bool           `yaml:"host_pattern,omitempty"`
	HostPattern   *regexp.Regexp
}

// FeatureConfig is the list of collectors enabled or disabled
type FeatureConfig struct {
	Alarm               bool `yaml:"alarm,omitempty"`
	Environment         bool `yaml:"environment,omitempty"`
	BFD                 bool `yaml:"bfd,omitempty"`
	BGP                 bool `yaml:"bgp,omitempty"`
	OSPF                bool `yaml:"ospf,omitempty"`
	ISIS                bool `yaml:"isis,omitempty"`
	NAT                 bool `yaml:"nat,omitempty"`
	NAT2                bool `yaml:"nat2,omitempty"`
	L2Circuit           bool `yaml:"l2circuit,omitempty"`
	L2Vpn               bool `yaml:"l2vpn,omitempty"`
	LACP                bool `yaml:"lacp,omitempty"`
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
	Security            bool `yaml:"security,omitempty"`
	SecurityIKE         bool `yaml:"security_ike,omitempty"`
	SecurityPolicies    bool `yaml:"security_policies,omitempty"`
	FPC                 bool `yaml:"fpc,omitempty"`
	RPKI                bool `yaml:"rpki,omitempty"`
	RPM                 bool `yaml:"rpm,omitempty"`
	Satellite           bool `yaml:"satellite,omitempty"`
	System              bool `yaml:"system,omitempty"`
	Power               bool `yaml:"power,omitempty"`
	MAC                 bool `yaml:"mac,omitempty"`
	MPLSLSP             bool `yaml:"mpls_lsp,omitempty"`
	VPWS                bool `yaml:"vpws,omitempty"`
	VRRP                bool `yaml:"vrrp,omitempty"`
	License             bool `yaml:"license,omitempty"`
	Subscriber          bool `yaml:"subscriber,omitempty"`
	MACSec              bool `yaml:"macsec,omitempty"`
	ARP                 bool `yaml:"arp,omitempty"`
	Poe                 bool `yaml:"poe,omitempty"`
	DDOSProtection      bool `yaml:"ddos_protection,omitempty"`
	KRT                 bool `yaml:"krt,omitempty"`
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
func Load(reader io.Reader, dynamicIfaceLabels bool) (*Config, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	c := New()
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	err = c.load(dynamicIfaceLabels)
	if err != nil {
		return nil, err
	}

	for _, device := range c.Devices {
		if device.IsHostPattern {
			hostPattern, err := regexp.Compile(device.Host)
			if err != nil {
				return nil, err
			}
			device.HostPattern = hostPattern
		}
	}

	return c, nil
}

func setDefaultValues(c *Config) {
	c.Password = ""
	c.LSEnabled = false
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
	f.Security = false
	f.SecurityPolicies = false
	f.Storage = false
	f.Accounting = false
	f.FPC = false
	f.L2Circuit = false
	f.L2Vpn = false
	f.RPKI = false
	f.RPM = false
	f.Satellite = false
	f.Power = false
	f.MAC = false
	f.MPLSLSP = false
	f.VPWS = false
	f.VRRP = false
	f.BFD = false
	f.License = false
	f.MACSec = true
	f.Poe = false
}

// FeaturesForDevice gets the feature set configured for a device
func (c *Config) FeaturesForDevice(host string) *FeatureConfig {
	d := c.FindDeviceConfig(host)

	if d != nil && d.Features != nil {
		return d.Features
	}

	return &c.Features
}

func (c *Config) FindDeviceConfig(host string) *DeviceConfig {
	for _, dc := range c.Devices {
		if dc.HostPattern != nil {
			if dc.HostPattern.MatchString(host) {
				return dc
			}
		} else {
			if dc.Host == host {
				return dc
			}
		}
	}

	return nil
}
