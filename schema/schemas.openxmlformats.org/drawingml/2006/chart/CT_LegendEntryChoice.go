// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_LegendEntryChoice struct {
	Delete *CT_Boolean
	TxPr   *drawingml.CT_TextBody
}

func NewCT_LegendEntryChoice() *CT_LegendEntryChoice {
	ret := &CT_LegendEntryChoice{}
	return ret
}

func (m *CT_LegendEntryChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Delete != nil {
		sedelete := xml.StartElement{Name: xml.Name{Local: "delete"}}
		e.EncodeElement(m.Delete, sedelete)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	return nil
}

func (m *CT_LegendEntryChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_LegendEntryChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "delete":
				m.Delete = NewCT_Boolean()
				if err := d.DecodeElement(m.Delete, &el); err != nil {
					return err
				}
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_LegendEntryChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LegendEntryChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LegendEntryChoice and its children
func (m *CT_LegendEntryChoice) Validate() error {
	return m.ValidateWithPath("CT_LegendEntryChoice")
}

// ValidateWithPath validates the CT_LegendEntryChoice and its children, prefixing error messages with path
func (m *CT_LegendEntryChoice) ValidateWithPath(path string) error {
	if m.Delete != nil {
		if err := m.Delete.ValidateWithPath(path + "/Delete"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	return nil
}
