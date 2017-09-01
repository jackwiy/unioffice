// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_EmbeddedFontList struct {
	// Embedded Font
	EmbeddedFont []*CT_EmbeddedFontListEntry
}

func NewCT_EmbeddedFontList() *CT_EmbeddedFontList {
	ret := &CT_EmbeddedFontList{}
	return ret
}

func (m *CT_EmbeddedFontList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.EmbeddedFont != nil {
		seembeddedFont := xml.StartElement{Name: xml.Name{Local: "p:embeddedFont"}}
		e.EncodeElement(m.EmbeddedFont, seembeddedFont)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EmbeddedFontList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EmbeddedFontList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "embeddedFont":
				tmp := NewCT_EmbeddedFontListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.EmbeddedFont = append(m.EmbeddedFont, tmp)
			default:
				log.Printf("skipping unsupported element on CT_EmbeddedFontList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EmbeddedFontList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EmbeddedFontList and its children
func (m *CT_EmbeddedFontList) Validate() error {
	return m.ValidateWithPath("CT_EmbeddedFontList")
}

// ValidateWithPath validates the CT_EmbeddedFontList and its children, prefixing error messages with path
func (m *CT_EmbeddedFontList) ValidateWithPath(path string) error {
	for i, v := range m.EmbeddedFont {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EmbeddedFont[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
