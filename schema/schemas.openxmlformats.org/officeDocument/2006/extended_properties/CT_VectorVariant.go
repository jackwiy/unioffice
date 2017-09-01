// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package extended_properties

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"
)

type CT_VectorVariant struct {
	Vector *docPropsVTypes.Vector
}

func NewCT_VectorVariant() *CT_VectorVariant {
	ret := &CT_VectorVariant{}
	ret.Vector = docPropsVTypes.NewVector()
	return ret
}

func (m *CT_VectorVariant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sevector := xml.StartElement{Name: xml.Name{Local: "vt:vector"}}
	e.EncodeElement(m.Vector, sevector)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VectorVariant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Vector = docPropsVTypes.NewVector()
lCT_VectorVariant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "vector":
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_VectorVariant %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VectorVariant
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VectorVariant and its children
func (m *CT_VectorVariant) Validate() error {
	return m.ValidateWithPath("CT_VectorVariant")
}

// ValidateWithPath validates the CT_VectorVariant and its children, prefixing error messages with path
func (m *CT_VectorVariant) ValidateWithPath(path string) error {
	if err := m.Vector.ValidateWithPath(path + "/Vector"); err != nil {
		return err
	}
	return nil
}
