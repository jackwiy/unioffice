// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
)

type CT_FontSig struct {
	// First 32 Bits of Unicode Subset Bitfield
	Usb0Attr string
	// Second 32 Bits of Unicode Subset Bitfield
	Usb1Attr string
	// Third 32 Bits of Unicode Subset Bitfield
	Usb2Attr string
	// Fourth 32 Bits of Unicode Subset Bitfield
	Usb3Attr string
	// Lower 32 Bits of Code Page Bit Field
	Csb0Attr string
	// Upper 32 Bits of Code Page Bit Field
	Csb1Attr string
}

func NewCT_FontSig() *CT_FontSig {
	ret := &CT_FontSig{}
	return ret
}
func (m *CT_FontSig) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:usb0"},
		Value: fmt.Sprintf("%v", m.Usb0Attr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:usb1"},
		Value: fmt.Sprintf("%v", m.Usb1Attr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:usb2"},
		Value: fmt.Sprintf("%v", m.Usb2Attr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:usb3"},
		Value: fmt.Sprintf("%v", m.Usb3Attr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:csb0"},
		Value: fmt.Sprintf("%v", m.Csb0Attr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:csb1"},
		Value: fmt.Sprintf("%v", m.Csb1Attr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FontSig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "usb0" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Usb0Attr = parsed
		}
		if attr.Name.Local == "usb1" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Usb1Attr = parsed
		}
		if attr.Name.Local == "usb2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Usb2Attr = parsed
		}
		if attr.Name.Local == "usb3" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Usb3Attr = parsed
		}
		if attr.Name.Local == "csb0" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Csb0Attr = parsed
		}
		if attr.Name.Local == "csb1" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Csb1Attr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FontSig: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_FontSig) Validate() error {
	return m.ValidateWithPath("CT_FontSig")
}
func (m *CT_FontSig) ValidateWithPath(path string) error {
	return nil
}