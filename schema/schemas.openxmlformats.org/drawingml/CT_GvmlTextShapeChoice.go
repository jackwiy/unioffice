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

type CT_GvmlTextShapeChoice struct {
	UseSpRect *CT_GvmlUseShapeRectangle
	Xfrm      *CT_Transform2D
}

func NewCT_GvmlTextShapeChoice() *CT_GvmlTextShapeChoice {
	ret := &CT_GvmlTextShapeChoice{}
	return ret
}

func (m *CT_GvmlTextShapeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.UseSpRect != nil {
		seuseSpRect := xml.StartElement{Name: xml.Name{Local: "a:useSpRect"}}
		e.EncodeElement(m.UseSpRect, seuseSpRect)
	}
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	return nil
}

func (m *CT_GvmlTextShapeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GvmlTextShapeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "useSpRect":
				m.UseSpRect = NewCT_GvmlUseShapeRectangle()
				if err := d.DecodeElement(m.UseSpRect, &el); err != nil {
					return err
				}
			case "xfrm":
				m.Xfrm = NewCT_Transform2D()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_GvmlTextShapeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlTextShapeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlTextShapeChoice and its children
func (m *CT_GvmlTextShapeChoice) Validate() error {
	return m.ValidateWithPath("CT_GvmlTextShapeChoice")
}

// ValidateWithPath validates the CT_GvmlTextShapeChoice and its children, prefixing error messages with path
func (m *CT_GvmlTextShapeChoice) ValidateWithPath(path string) error {
	if m.UseSpRect != nil {
		if err := m.UseSpRect.ValidateWithPath(path + "/UseSpRect"); err != nil {
			return err
		}
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
			return err
		}
	}
	return nil
}
