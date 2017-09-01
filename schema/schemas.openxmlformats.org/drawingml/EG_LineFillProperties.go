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

type EG_LineFillProperties struct {
	NoFill    *CT_NoFillProperties
	SolidFill *CT_SolidColorFillProperties
	GradFill  *CT_GradientFillProperties
	PattFill  *CT_PatternFillProperties
}

func NewEG_LineFillProperties() *EG_LineFillProperties {
	ret := &EG_LineFillProperties{}
	return ret
}

func (m *EG_LineFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	return nil
}

func (m *EG_LineFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_LineFillProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "noFill":
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case "solidFill":
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case "gradFill":
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case "pattFill":
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_LineFillProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_LineFillProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_LineFillProperties and its children
func (m *EG_LineFillProperties) Validate() error {
	return m.ValidateWithPath("EG_LineFillProperties")
}

// ValidateWithPath validates the EG_LineFillProperties and its children, prefixing error messages with path
func (m *EG_LineFillProperties) ValidateWithPath(path string) error {
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	return nil
}
