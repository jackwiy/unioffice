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
)

type CT_RelSizeAnchor struct {
	From   *CT_Marker
	To     *CT_Marker
	Choice *EG_ObjectChoicesChoice
}

func NewCT_RelSizeAnchor() *CT_RelSizeAnchor {
	ret := &CT_RelSizeAnchor{}
	ret.From = NewCT_Marker()
	ret.To = NewCT_Marker()
	return ret
}

func (m *CT_RelSizeAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	sefrom := xml.StartElement{Name: xml.Name{Local: "from"}}
	e.EncodeElement(m.From, sefrom)
	seto := xml.StartElement{Name: xml.Name{Local: "to"}}
	e.EncodeElement(m.To, seto)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RelSizeAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = NewCT_Marker()
	m.To = NewCT_Marker()
lCT_RelSizeAnchor:
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
			case "to":
				if err := d.DecodeElement(m.To, &el); err != nil {
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
				log.Printf("skipping unsupported element on CT_RelSizeAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RelSizeAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RelSizeAnchor and its children
func (m *CT_RelSizeAnchor) Validate() error {
	return m.ValidateWithPath("CT_RelSizeAnchor")
}

// ValidateWithPath validates the CT_RelSizeAnchor and its children, prefixing error messages with path
func (m *CT_RelSizeAnchor) ValidateWithPath(path string) error {
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.To.ValidateWithPath(path + "/To"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
