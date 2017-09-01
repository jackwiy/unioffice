// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_AbsSizeAnchor struct {
	From   *CT_Marker
	Ext    *drawingml.CT_PositiveSize2D
	Choice *EG_ObjectChoicesChoice
}

func NewCT_AbsSizeAnchor() *CT_AbsSizeAnchor {
	ret := &CT_AbsSizeAnchor{}
	ret.From = NewCT_Marker()
	ret.Ext = drawingml.NewCT_PositiveSize2D()
	return ret
}

func (m *CT_AbsSizeAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sefrom := xml.StartElement{Name: xml.Name{Local: "from"}}
	e.EncodeElement(m.From, sefrom)
	seext := xml.StartElement{Name: xml.Name{Local: "ext"}}
	e.EncodeElement(m.Ext, seext)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AbsSizeAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = NewCT_Marker()
	m.Ext = drawingml.NewCT_PositiveSize2D()
lCT_AbsSizeAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "from":
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case "ext":
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			case "sp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "grpSp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "graphicFrame":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "cxnSp":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "pic":
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
				_ = m.Choice
			default:
				log.Printf("skipping unsupported element on CT_AbsSizeAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AbsSizeAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AbsSizeAnchor and its children
func (m *CT_AbsSizeAnchor) Validate() error {
	return m.ValidateWithPath("CT_AbsSizeAnchor")
}

// ValidateWithPath validates the CT_AbsSizeAnchor and its children, prefixing error messages with path
func (m *CT_AbsSizeAnchor) ValidateWithPath(path string) error {
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
