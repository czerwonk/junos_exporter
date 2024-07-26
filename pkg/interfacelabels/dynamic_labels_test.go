// SPDX-License-Identifier: MIT

package interfacelabels

import (
	"regexp"
	"testing"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/stretchr/testify/assert"
)

func TestParseDescriptions(t *testing.T) {
	t.Run("Test default", func(t *testing.T) {
		l := NewDynamicLabelManager()
		regex := DefaultInterfaceDescRegex()

		if1 := interfaceDescription{
			Name:        "xe-0/0/0",
			Description: "Name1 [tag1] [foo=x]",
		}
		if2 := interfaceDescription{
			Name:        "xe-0/0/1",
			Description: "Name2 [foo=y] [bar=123]",
		}
		if3 := interfaceDescription{
			Name: "xe-0/0/3",
		}
		if4 := interfaceDescription{
			Name:        "irb.216",
			Description: "Internal: Test-Network [vrf=AS4711]",
		}

		d1 := &connector.Device{Host: "device1"}
		d2 := &connector.Device{Host: "device2"}

		l.parseDescriptions(d1, []interfaceDescription{if1}, regex)
		l.parseDescriptions(d2, []interfaceDescription{if2, if3, if4}, regex)

		assert.Equal(t, []string{"tag1", "foo", "bar", "vrf"}, l.LabelNames(), "Label names")
		assert.Equal(t, []string{"1", "x", "", ""}, l.ValuesForInterface(d1, if1.Name), "Values if1")
		assert.Equal(t, []string{"", "y", "123", ""}, l.ValuesForInterface(d2, if2.Name), "Values if2")
		assert.Equal(t, []string{"", "", "", ""}, l.ValuesForInterface(d2, if3.Name), "Values if3")
		assert.Equal(t, []string{"", "", "", "AS4711"}, l.ValuesForInterface(d2, if4.Name), "Values if4")
	})

	t.Run("Test custom regex", func(t *testing.T) {
		l := NewDynamicLabelManager()
		regex := regexp.MustCompile(`[[\s]([^=\[\]]+)(=[^,\]]+)?[,\]]`)

		if1 := interfaceDescription{
			Name:        "xe-0/0/0",
			Description: "Name1 [foo=x, bar=y, thisisatag]",
		}
		if2 := interfaceDescription{
			Name:        "xe-0/0/1",
			Description: "Name2 [onlyatag]",
		}
		if3 := interfaceDescription{
			Name:        "xe-0/0/3",
			Description: "Name2 [foo=x, bar=y, this=is]",
		}

		d1 := &connector.Device{Host: "device1"}
		d2 := &connector.Device{Host: "device2"}
		d3 := &connector.Device{Host: "device3"}

		l.parseDescriptions(d1, []interfaceDescription{if1}, regex)
		l.parseDescriptions(d2, []interfaceDescription{if2}, regex)
		l.parseDescriptions(d3, []interfaceDescription{if3}, regex)

		assert.Equal(t, []string{"foo", "bar", "thisisatag", "onlyatag", "this"}, l.LabelNames(), "Label names")
		assert.Equal(t, []string{"x", "y", "1", "", ""}, l.ValuesForInterface(d1, if1.Name), "Values if1")
		assert.Equal(t, []string{"", "", "", "1", ""}, l.ValuesForInterface(d2, if2.Name), "Values if2")
		assert.Equal(t, []string{"x", "y", "", "", "is"}, l.ValuesForInterface(d3, if3.Name), "Values if3")
	})
}
