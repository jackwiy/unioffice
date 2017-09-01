// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_PatternFillProperties struct {
	PrstAttr ST_PresetPatternVal
	FgClr    *CT_Color
	BgClr    *CT_Color
}

func NewCT_PatternFillProperties() *CT_PatternFillProperties {
	ret := &CT_PatternFillProperties{}
	return ret
}

func (m *CT_PatternFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PrstAttr != ST_PresetPatternValUnset {
		attr, err := m.PrstAttr.MarshalXMLAttr(xml.Name{Local: "prst"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FgClr != nil {
		sefgClr := xml.StartElement{Name: xml.Name{Local: "a:fgClr"}}
		e.EncodeElement(m.FgClr, sefgClr)
	}
	if m.BgClr != nil {
		sebgClr := xml.StartElement{Name: xml.Name{Local: "a:bgClr"}}
		e.EncodeElement(m.BgClr, sebgClr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PatternFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "prst" {
			m.PrstAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PatternFillProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fgClr":
				m.FgClr = NewCT_Color()
				if err := d.DecodeElement(m.FgClr, &el); err != nil {
					return err
				}
			case "bgClr":
				m.BgClr = NewCT_Color()
				if err := d.DecodeElement(m.BgClr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PatternFillProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PatternFillProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PatternFillProperties and its children
func (m *CT_PatternFillProperties) Validate() error {
	return m.ValidateWithPath("CT_PatternFillProperties")
}

// ValidateWithPath validates the CT_PatternFillProperties and its children, prefixing error messages with path
func (m *CT_PatternFillProperties) ValidateWithPath(path string) error {
	if err := m.PrstAttr.ValidateWithPath(path + "/PrstAttr"); err != nil {
		return err
	}
	if m.FgClr != nil {
		if err := m.FgClr.ValidateWithPath(path + "/FgClr"); err != nil {
			return err
		}
	}
	if m.BgClr != nil {
		if err := m.BgClr.ValidateWithPath(path + "/BgClr"); err != nil {
			return err
		}
	}
	return nil
}
