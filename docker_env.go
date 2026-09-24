// SPDX-License-Identifier: MIT

package main

import "strings"

var legacyDockerEnvVars = []string{"SSH_KEYFILE", "CONFIG_FILE", "ALARM_FILTER", "CMD_FLAGS"}

// legacyDockerArgs rebuilds the command line of the former shell based Docker CMD
// (./junos_exporter -ssh.keyfile=$SSH_KEYFILE -config.file=$CONFIG_FILE -alarms.filter=$ALARM_FILTER $CMD_FLAGS)
// so containers configured via environment variables keep working on the shell-less distroless image.
// Like Docker replacing CMD, explicitly passed arguments disable the env var handling.
func legacyDockerArgs(args []string, lookupEnv func(string) (string, bool)) []string {
	if len(args) > 0 || !anyEnvSet(legacyDockerEnvVars, lookupEnv) {
		return args
	}

	env := func(key string) string {
		v, _ := lookupEnv(key)
		return v
	}

	legacyArgs := []string{
		"-ssh.keyfile=" + env("SSH_KEYFILE"),
		"-config.file=" + env("CONFIG_FILE"),
		"-alarms.filter=" + env("ALARM_FILTER"),
	}

	// unquoted $CMD_FLAGS was split on whitespace by the shell
	return append(legacyArgs, strings.Fields(env("CMD_FLAGS"))...)
}

func anyEnvSet(keys []string, lookupEnv func(string) (string, bool)) bool {
	for _, k := range keys {
		if _, ok := lookupEnv(k); ok {
			return true
		}
	}

	return false
}
