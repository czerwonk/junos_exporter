// SPDX-License-Identifier: MIT

package main

import (
	"reflect"
	"testing"

	"github.com/czerwonk/junos_exporter/internal/config"
)

// if this test fails, CLI flag defaults have drifted from config defaults
func TestDefaultsAreConsistent(t *testing.T) {
	configDefaults := config.New().Features
	flagDefaults := loadConfigFromFlags().Features

	cv := reflect.ValueOf(configDefaults)
	fv := reflect.ValueOf(flagDefaults)
	ct := cv.Type()

	for i := 0; i < ct.NumField(); i++ {
		field := ct.Field(i)
		expected := cv.Field(i).Interface()
		actual := fv.Field(i).Interface()

		if expected != actual {
			t.Errorf("feature %s: config default is %v, but flag default is %v", field.Name, expected, actual)
		}
	}
}
