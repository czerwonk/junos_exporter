package main

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type OutboundChannel struct {
	ID string `xml:"id,attr"`
}

type Interface struct {
	Name       string           `xml:"name,attr"`
	OutChannel *OutboundChannel `xml:"-"`
}

type Root struct {
	Interfaces []*Interface `xml:"interface"`
}

const example = `
<root>
	<interface name="et-0/0/0"></interface>
	<outbound-channel id="1"></outbound-channel>
	<interface name="et-0/0/1"></interface>
	<interface name="et-0/0/2"></interface>
	<outbound-channel id="3"></outbound-channel>
</root>
`

func (r *Root) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}

		switch tok.(type) {
		case xml.CharData:
			continue // ignore whitespace
		case xml.EndElement:
			return nil
		}

		elem, ok := tok.(xml.StartElement)
		if !ok {
			return fmt.Errorf("invalid token encountered: %T", tok)
		}
		switch elem.Name.Local {
		case "interface":
			var v Interface
			if err := d.DecodeElement(&v, &elem); err != nil {
				return err
			}
			r.Interfaces = append(r.Interfaces, &v)
		case "outbound-channel":
			if len(r.Interfaces) == 0 {
				// no interface encountered, skip this
				continue
			}

			var v OutboundChannel
			if err := d.DecodeElement(&v, &elem); err != nil {
				return err
			}

			currentIf := r.Interfaces[len(r.Interfaces)-1]
			currentIf.OutChannel = &v
		}
	}
}

func TestXMLUnmarshal(t *testing.T) {
	var target Root
	err := xml.Unmarshal([]byte(example), &target)
	if err != nil {
		t.Error(err)
	}

	require.Len(t, target.Interfaces, 3)
	assert.Equal(t, "et-0/0/0", target.Interfaces[0].Name)
	assert.Equal(t, "1", target.Interfaces[0].OutChannel.ID)

	assert.Equal(t, "et-0/0/1", target.Interfaces[1].Name)
	assert.Nil(t, target.Interfaces[1].OutChannel)

	assert.Equal(t, "et-0/0/2", target.Interfaces[2].Name)
	assert.Equal(t, "3", target.Interfaces[2].OutChannel.ID)
}
