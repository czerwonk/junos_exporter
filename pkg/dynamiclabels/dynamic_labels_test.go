// SPDX-License-Identifier: MIT

package dynamiclabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDescriptions(t *testing.T) {
	tests := []struct {
		name        string
		description string
		keys        []string
		values      []string
	}{
		{
			name:        "tags and kv pairs",
			description: "Name1 [tag1] [foo=x]",
			keys:        []string{"tag1", "foo"},
			values:      []string{"1", "x"},
		},
		{
			name:        "kv pairs",
			description: "Name2 [foo=y] [bar=123]",
			keys:        []string{"foo", "bar"},
			values:      []string{"y", "123"},
		},
		{
			name:        "empty",
			description: "",
			keys:        []string{},
			values:      []string{},
		},
		{
			name:        "more kv pairs",
			description: "Internal: Test-Network [vrf=AS4711]",
			keys:        []string{"vrf"},
			values:      []string{"AS4711"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ifLabels := ParseDescription(test.description, DefaultInterfaceDescRegex())
			assert.Equal(t, test.keys, ifLabels.Keys(), test.name)
			assert.Equal(t, test.values, ifLabels.Values(), test.name)
		})
	}
}
