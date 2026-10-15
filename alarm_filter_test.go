// SPDX-License-Identifier: MIT

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileAlarmFilter(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantNil    bool
		wantErr    bool
	}{
		{"empty matches nothing", "", true, false},
		{"valid expression", "^Rescue.*", false, false},
		{"invalid expression", "[unclosed", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re, err := compileAlarmFilter(tt.expression)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "alarms.filter")
			} else {
				assert.NoError(t, err)
			}

			if tt.wantNil {
				assert.Nil(t, re)
			} else {
				assert.NotNil(t, re)
			}
		})
	}
}
