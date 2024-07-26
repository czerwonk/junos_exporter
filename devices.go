// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/czerwonk/junos_exporter/internal/config"
	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/pkg/errors"
)

func devicesForConfig(cfg *config.Config) ([]*connector.Device, error) {
	if cfg.Devices == nil {
		if cfg.Targets == nil {
			cfg.Targets = strings.Split(strings.Trim(*sshHosts, " "), ",")
		}

		cfg.Devices = devicesFromTargets(cfg.Targets)
	}

	devs := make([]*connector.Device, 0)
	for _, d := range cfg.Devices {
		if d.IsHostPattern {
			continue
		}

		dev, err := deviceFromDeviceConfig(d, d.Host, cfg)
		if err != nil {
			return nil, err
		}

		devs = append(devs, dev)
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

func deviceFromDeviceConfig(device *config.DeviceConfig, hostname string, cfg *config.Config) (*connector.Device, error) {
	auth, err := authForDevice(device, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "could not initialize config for device %s", device.Host)
	}

	// check whether there is a device specific regex otherwise fallback to global regex
	if len(device.IfDescRegStr) == 0 {
		device.IfDescReg = cfg.IfDescReg
	} else {
		re, err := regexp.Compile(device.IfDescRegStr)
		if err != nil {
			return nil, fmt.Errorf("unable to compile device description regex for %q: %q: %w", hostname, device.IfDescRegStr, err)
		}

		device.IfDescReg = re
	}

	return &connector.Device{
		Host: hostname,
		Auth: auth,
	}, nil
}

func authForDevice(device *config.DeviceConfig, cfg *config.Config) (connector.AuthMethod, error) {
	user := *sshUsername
	if device.Username != "" {
		user = device.Username
	}

	if device.KeyFile != "" {
		return authForKeyFile(user, device.KeyFile, device.KeyPassphrase)
	}

	if *sshKeyFile != "" {
		return authForKeyFile(user, *sshKeyFile, *sshKeyPassphrase)
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

func authForKeyFile(username, keyFile, keyPassphrase string) (connector.AuthMethod, error) {
	f, err := os.Open(keyFile)
	if err != nil {
		return nil, errors.Wrap(err, "could not open ssh key file")
	}
	defer f.Close()

	auth, err := connector.AuthByKey(username, f, keyPassphrase)
	if err != nil {
		return nil, errors.Wrap(err, "could not load ssh private key file")
	}

	return auth, nil
}
