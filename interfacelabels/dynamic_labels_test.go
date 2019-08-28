package interfacelabels

import (
	"testing"

	"github.com/czerwonk/junos_exporter/connector"
	"github.com/stretchr/testify/assert"
)

func TestParseDescriptions(t *testing.T) {
	l := NewDynamicLabels()

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

	l.parseDescriptions(d1, []PhyInterface{if1})
	l.parseDescriptions(d2, []PhyInterface{if2})

	assert.Equal(t, []string{"tag1", "foo", "bar"}, l.LabelNames(), "Label names")
	assert.Equal(t, []string{"1", "x", ""}, l.ValuesForInterface(d1, if1.Name), "Values if1")
	assert.Equal(t, []string{"", "y", "123"}, l.ValuesForInterface(d2, if2.Name), "Values if2")
	assert.Equal(t, []string{"", "", ""}, l.ValuesForInterface(d2, if3.Name), "Values if3")
}
