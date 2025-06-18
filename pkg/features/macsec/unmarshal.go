package macsec

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

func ParseShowSecurityMacsecConnections(input []byte) (*ShowSecMacsecConns, error) {
	res := &ShowSecMacsecConns{}

	err := xml.Unmarshal(input, res)
	fmt.Println(input)
	if err != nil {
		return nil, fmt.Errorf("xml.unmarshal failed: %v", err)
	}

	res.MacsecConnectionInformation = make([]*MacsecConnectionInformation, 0)

	d := xml.NewDecoder(bytes.NewBuffer(res.InnerXML))
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read token: %v", err)
		}

		switch e := t.(type) {
		case xml.StartElement:
			if e.Name.Local == "macsec-interface-common-information" {
				mici, err := unmarshalMacsecInterfaceCommonInformation(d, &e)
				if err != nil {
					return nil, fmt.Errorf("unable to unmarshal macsec-interface-common-information")
				}

				res.MacsecConnectionInformation = append(res.MacsecConnectionInformation, &MacsecConnectionInformation{
					MacsecInterfaceCommonInformation: mici,
				})
			}

			if e.Name.Local == "outbound-secure-channel" {
				osc, err := unmarshalOutboundSecureChannel(d, &e)
				if err != nil {
					return nil, fmt.Errorf("unable to unmarshal outbound-secure-channel")
				}

				n := len(res.MacsecConnectionInformation)
				if n == 0 {
					return nil, fmt.Errorf("found outbound-secure-channel before macsec-interface-common-information. XML invalid")
				}

				res.MacsecConnectionInformation[n-1].OutboundSecureChannel = osc
			}

			if e.Name.Local == "inbound-secure-channel" {
				isc, err := unmarshalInboundSecureChannel(d, &e)
				if err != nil {
					return nil, fmt.Errorf("unable to unmarshal inbound-secure-channel")
				}

				n := len(res.MacsecConnectionInformation)
				if n == 0 {
					return nil, fmt.Errorf("found inbound-secure-channel before macsec-interface-common-information. XML invalid")
				}

				res.MacsecConnectionInformation[n-1].InboundSecureChannel = isc
			}
		}

	}

	return res, nil
}

func unmarshalMacsecInterfaceCommonInformation(d *xml.Decoder, start *xml.StartElement) (*MacsecInterfaceCommonInformation, error) {
	mici := &MacsecInterfaceCommonInformation{}
	err := d.DecodeElement(mici, start)
	if err != nil {
		return nil, err
	}

	return mici, nil
}

func unmarshalOutboundSecureChannel(d *xml.Decoder, start *xml.StartElement) (*OutboundSecureChannel, error) {
	osc := &OutboundSecureChannel{}
	err := d.DecodeElement(osc, start)
	if err != nil {
		return nil, err
	}

	return osc, nil
}

func unmarshalInboundSecureChannel(d *xml.Decoder, start *xml.StartElement) (*InboundSecureChannel, error) {
	isc := &InboundSecureChannel{}
	err := d.DecodeElement(isc, start)
	if err != nil {
		return nil, err
	}

	return isc, nil
}
