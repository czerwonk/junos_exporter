// SPDX-License-Identifier: MIT

package main

import (
	"reflect"
	"testing"
)

func TestLegacyDockerArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		env  map[string]string
		want []string
	}{
		{
			name: "no env vars set keeps empty args",
			args: []string{},
			env:  map[string]string{},
			want: []string{},
		},
		{
			name: "explicit args take precedence over env vars",
			args: []string{"-config.file=/other.yml"},
			env:  map[string]string{"CONFIG_FILE": "/config.yml", "SSH_KEYFILE": "/key"},
			want: []string{"-config.file=/other.yml"},
		},
		{
			name: "image defaults",
			args: []string{},
			env:  map[string]string{"SSH_KEYFILE": "", "CONFIG_FILE": "/config.yml", "ALARM_FILTER": "", "CMD_FLAGS": ""},
			want: []string{"-ssh.keyfile=", "-config.file=/config.yml", "-alarms.filter="},
		},
		{
			name: "all env vars set",
			args: []string{},
			env: map[string]string{
				"SSH_KEYFILE":  "/ssh-keyfile",
				"CONFIG_FILE":  "/config.yml",
				"ALARM_FILTER": "Management Ethernet",
				"CMD_FLAGS":    " -debug  -bgp.enabled=false\t-web.listen-address=:9999 ",
			},
			want: []string{
				"-ssh.keyfile=/ssh-keyfile",
				"-config.file=/config.yml",
				"-alarms.filter=Management Ethernet",
				"-debug",
				"-bgp.enabled=false",
				"-web.listen-address=:9999",
			},
		},
		{
			name: "single env var set expands others to empty values",
			args: []string{},
			env:  map[string]string{"SSH_KEYFILE": "/ssh-keyfile"},
			want: []string{"-ssh.keyfile=/ssh-keyfile", "-config.file=", "-alarms.filter="},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookupEnv := func(key string) (string, bool) {
				v, ok := tt.env[key]
				return v, ok
			}

			got := legacyDockerArgs(tt.args, lookupEnv)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("legacyDockerArgs() = %q, want %q", got, tt.want)
			}
		})
	}
}
