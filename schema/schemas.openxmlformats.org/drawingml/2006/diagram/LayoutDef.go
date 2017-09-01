// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type LayoutDef struct {
	CT_DiagramDefinition
}

func NewLayoutDef() *LayoutDef {
	ret := &LayoutDef{}
	ret.CT_DiagramDefinition = *NewCT_DiagramDefinition()
	return ret
}

func (m *LayoutDef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "layoutDef"
	return m.CT_DiagramDefinition.MarshalXML(e, start)
}

func (m *LayoutDef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_DiagramDefinition = *NewCT_DiagramDefinition()
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueIdAttr = &parsed
		}
		if attr.Name.Local == "minVer" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MinVerAttr = &parsed
		}
		if attr.Name.Local == "defStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefStyleAttr = &parsed
		}
	}
lLayoutDef:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "title":
				tmp := NewCT_Name()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Title = append(m.Title, tmp)
			case "desc":
				tmp := NewCT_Description()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Desc = append(m.Desc, tmp)
			case "catLst":
				m.CatLst = NewCT_Categories()
				if err := d.DecodeElement(m.CatLst, &el); err != nil {
					return err
				}
			case "sampData":
				m.SampData = NewCT_SampleData()
				if err := d.DecodeElement(m.SampData, &el); err != nil {
					return err
				}
			case "styleData":
				m.StyleData = NewCT_SampleData()
				if err := d.DecodeElement(m.StyleData, &el); err != nil {
					return err
				}
			case "clrData":
				m.ClrData = NewCT_SampleData()
				if err := d.DecodeElement(m.ClrData, &el); err != nil {
					return err
				}
			case "layoutNode":
				if err := d.DecodeElement(m.LayoutNode, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on LayoutDef %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lLayoutDef
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the LayoutDef and its children
func (m *LayoutDef) Validate() error {
	return m.ValidateWithPath("LayoutDef")
}

// ValidateWithPath validates the LayoutDef and its children, prefixing error messages with path
func (m *LayoutDef) ValidateWithPath(path string) error {
	if err := m.CT_DiagramDefinition.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
