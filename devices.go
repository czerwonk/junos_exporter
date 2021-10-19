package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/czerwonk/junos_exporter/config"
	"github.com/czerwonk/junos_exporter/connector"
	"github.com/pkg/errors"
)

func devicesForConfig(cfg *config.Config) ([]*connector.Device, error) {
	if cfg.Devices == nil {
		if cfg.Targets == nil {
			cfg.Targets = strings.Split(strings.Trim(*sshHosts, " "), ",")
		}

		cfg.Devices = devicesFromTargets(cfg.Targets)
	}

	devs := make([]*connector.Device, len(cfg.Devices))
	var err error
	for i, d := range cfg.Devices {
		devs[i], err = deviceFromDeviceConfig(d, cfg)
		if err != nil {
			return nil, err
		}
	}

	return devs, nil
}

func devicesFromTargets(targets []string) []*config.DeviceConfig {
	devices := make([]*config.DeviceConfig, len(targets))
	for i, t := range targets {
		devices[i] = &config.DeviceConfig{
			Host: t,
		}
	}

	return devices
}

func deviceFromDeviceConfig(device *config.DeviceConfig, cfg *config.Config) (*connector.Device, error) {
	auth, err := authForDevice(device, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "could not initialize config for device %s", device.Host)
	}

	// check whether there is a device specific regex otherwise fallback to global regex
	if len(device.IfDescReg) == 0 {
		device.IfDescReg = cfg.IfDescReg
	} else {
		regexp.MustCompile(device.IfDescReg)
	}

	return &connector.Device{
		Host: device.Host,
		Auth: auth,
	}, nil
}

func authForDevice(device *config.DeviceConfig, cfg *config.Config) (connector.AuthMethod, error) {
	user := *sshUsername
	if device.Username != "" {
		user = device.Username
	}

	if device.KeyFile != "" {
		return authForKeyFile(user, device.KeyFile)
	}

	if *sshKeyFile != "" {
		return authForKeyFile(user, *sshKeyFile)
	}

	if device.Password != "" {
		return connector.AuthByPassword(user, device.Password), nil
	}

	if cfg.Password != "" {
		return connector.AuthByPassword(user, cfg.Password), nil
	}

	if *sshPassword != "" {
		return connector.AuthByPassword(user, *sshPassword), nil
	}

	return nil, errors.New("no valid authentication method available")
}

func authForKeyFile(username, keyFile string) (connector.AuthMethod, error) {
	f, err := os.Open(keyFile)
	if err != nil {
		return nil, errors.Wrap(err, "could not open ssh key file")
	}
	defer f.Close()

	auth, err := connector.AuthByKey(username, f)
	if err != nil {
		return nil, errors.Wrap(err, "could not load ssh private key file")
	}

	return auth, nil
}
