// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type GlossaryDocument struct {
	CT_GlossaryDocument
}

func NewGlossaryDocument() *GlossaryDocument {
	ret := &GlossaryDocument{}
	ret.CT_GlossaryDocument = *NewCT_GlossaryDocument()
	return ret
}

func (m *GlossaryDocument) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:glossaryDocument"
	return m.CT_GlossaryDocument.MarshalXML(e, start)
}

func (m *GlossaryDocument) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_GlossaryDocument = *NewCT_GlossaryDocument()
lGlossaryDocument:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "background":
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			case "docParts":
				m.DocParts = NewCT_DocParts()
				if err := d.DecodeElement(m.DocParts, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on GlossaryDocument %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lGlossaryDocument
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the GlossaryDocument and its children
func (m *GlossaryDocument) Validate() error {
	return m.ValidateWithPath("GlossaryDocument")
}

// ValidateWithPath validates the GlossaryDocument and its children, prefixing error messages with path
func (m *GlossaryDocument) ValidateWithPath(path string) error {
	if err := m.CT_GlossaryDocument.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
