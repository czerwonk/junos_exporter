package interfacelabels

import (
	"regexp"
	"testing"

	"github.com/czerwonk/junos_exporter/connector"
	"github.com/stretchr/testify/assert"
)

func TestParseDescriptions(t *testing.T) {

	t.Run("Test default", func(t *testing.T) {
		l := NewDynamicLabels()
		regex := regexp.MustCompile(`\[([^=\]]+)(=[^\]]+)?\]`)

		if1 := PhyInterface{
			Name:        "xe-0/0/0",
			Description: "Name1 [tag1] [foo=x]",
		}
		if2 := PhyInterface{
			Name:        "xe-0/0/1",
			Description: "Name2 [foo=y] [bar=123]",
		}
		if3 := PhyInterface{
			Name: "xe-0/0/3",
		}

		d1 := &connector.Device{Host: "device1"}
		d2 := &connector.Device{Host: "device2"}

		l.parseDescriptions(d1, []PhyInterface{if1}, regex)
		l.parseDescriptions(d2, []PhyInterface{if2}, regex)

		assert.Equal(t, []string{"tag1", "foo", "bar"}, l.LabelNames(), "Label names")
		assert.Equal(t, []string{"1", "x", ""}, l.ValuesForInterface(d1, if1.Name), "Values if1")
		assert.Equal(t, []string{"", "y", "123"}, l.ValuesForInterface(d2, if2.Name), "Values if2")
		assert.Equal(t, []string{"", "", ""}, l.ValuesForInterface(d2, if3.Name), "Values if3")
	})

	t.Run("Test custom regex", func(t *testing.T) {
		l := NewDynamicLabels()
		regex := regexp.MustCompile(`[[\s]([^=\[\]]+)(=[^,\]]+)?[,\]]`)

		if1 := PhyInterface{
			Name:        "xe-0/0/0",
			Description: "Name1 [foo=x, bar=y, thisisatag]",
		}
		if2 := PhyInterface{
			Name:        "xe-0/0/1",
			Description: "Name2 [onlyatag]",
		}
		if3 := PhyInterface{
			Name:        "xe-0/0/3",
			Description: "Name2 [foo=x, bar=y, this=is]",
		}

		d1 := &connector.Device{Host: "device1"}
		d2 := &connector.Device{Host: "device2"}
		d3 := &connector.Device{Host: "device3"}

		l.parseDescriptions(d1, []PhyInterface{if1}, regex)
		l.parseDescriptions(d2, []PhyInterface{if2}, regex)
		l.parseDescriptions(d3, []PhyInterface{if3}, regex)

		assert.Equal(t, []string{"foo", "bar", "thisisatag", "onlyatag", "this"}, l.LabelNames(), "Label names")
		assert.Equal(t, []string{"x", "y", "1", "", ""}, l.ValuesForInterface(d1, if1.Name), "Values if1")
		assert.Equal(t, []string{"", "", "", "1", ""}, l.ValuesForInterface(d2, if2.Name), "Values if2")
		assert.Equal(t, []string{"x", "y", "", "", "is"}, l.ValuesForInterface(d3, if3.Name), "Values if3")
	})
}
