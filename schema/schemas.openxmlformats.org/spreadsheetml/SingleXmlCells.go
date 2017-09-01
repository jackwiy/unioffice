// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type SingleXmlCells struct {
	CT_SingleXmlCells
}

func NewSingleXmlCells() *SingleXmlCells {
	ret := &SingleXmlCells{}
	ret.CT_SingleXmlCells = *NewCT_SingleXmlCells()
	return ret
}

func (m *SingleXmlCells) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:singleXmlCells"
	return m.CT_SingleXmlCells.MarshalXML(e, start)
}

func (m *SingleXmlCells) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_SingleXmlCells = *NewCT_SingleXmlCells()
lSingleXmlCells:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "singleXmlCell":
				tmp := NewCT_SingleXmlCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SingleXmlCell = append(m.SingleXmlCell, tmp)
			default:
				log.Printf("skipping unsupported element on SingleXmlCells %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lSingleXmlCells
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the SingleXmlCells and its children
func (m *SingleXmlCells) Validate() error {
	return m.ValidateWithPath("SingleXmlCells")
}

// ValidateWithPath validates the SingleXmlCells and its children, prefixing error messages with path
func (m *SingleXmlCells) ValidateWithPath(path string) error {
	if err := m.CT_SingleXmlCells.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
