package main

import (
	"encoding/xml"
	"fmt"
	"testing"
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
		elem, ok := tok.(xml.StartElement)
		if !ok {
			fmt.Printf("invalid token encountered: %T\n", tok)
			continue
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

	fmt.Println("found interfaces:")
	for _, intf := range target.Interfaces {
		fmt.Printf("%#v\n", intf)
		fmt.Printf("outbound: %#v\n", intf.OutChannel)
	}
}
