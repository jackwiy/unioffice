// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_ChartsheetPr struct {
	// Published
	PublishedAttr *bool
	// Code Name
	CodeNameAttr *string
	TabColor     *CT_Color
}

func NewCT_ChartsheetPr() *CT_ChartsheetPr {
	ret := &CT_ChartsheetPr{}
	return ret
}

func (m *CT_ChartsheetPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "published"},
			Value: fmt.Sprintf("%v", *m.PublishedAttr)})
	}
	if m.CodeNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codeName"},
			Value: fmt.Sprintf("%v", *m.CodeNameAttr)})
	}
	e.EncodeToken(start)
	if m.TabColor != nil {
		setabColor := xml.StartElement{Name: xml.Name{Local: "x:tabColor"}}
		e.EncodeElement(m.TabColor, setabColor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartsheetPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
		}
		if attr.Name.Local == "codeName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CodeNameAttr = &parsed
		}
	}
lCT_ChartsheetPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tabColor":
				m.TabColor = NewCT_Color()
				if err := d.DecodeElement(m.TabColor, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ChartsheetPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartsheetPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartsheetPr and its children
func (m *CT_ChartsheetPr) Validate() error {
	return m.ValidateWithPath("CT_ChartsheetPr")
}

// ValidateWithPath validates the CT_ChartsheetPr and its children, prefixing error messages with path
func (m *CT_ChartsheetPr) ValidateWithPath(path string) error {
	if m.TabColor != nil {
		if err := m.TabColor.ValidateWithPath(path + "/TabColor"); err != nil {
			return err
		}
	}
	return nil
}
