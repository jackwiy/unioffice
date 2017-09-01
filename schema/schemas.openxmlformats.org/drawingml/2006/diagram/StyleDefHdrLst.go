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
)

type StyleDefHdrLst struct {
	CT_StyleDefinitionHeaderLst
}

func NewStyleDefHdrLst() *StyleDefHdrLst {
	ret := &StyleDefHdrLst{}
	ret.CT_StyleDefinitionHeaderLst = *NewCT_StyleDefinitionHeaderLst()
	return ret
}

func (m *StyleDefHdrLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "styleDefHdrLst"
	return m.CT_StyleDefinitionHeaderLst.MarshalXML(e, start)
}

func (m *StyleDefHdrLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_StyleDefinitionHeaderLst = *NewCT_StyleDefinitionHeaderLst()
lStyleDefHdrLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "styleDefHdr":
				tmp := NewCT_StyleDefinitionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.StyleDefHdr = append(m.StyleDefHdr, tmp)
			default:
				log.Printf("skipping unsupported element on StyleDefHdrLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lStyleDefHdrLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the StyleDefHdrLst and its children
func (m *StyleDefHdrLst) Validate() error {
	return m.ValidateWithPath("StyleDefHdrLst")
}

// ValidateWithPath validates the StyleDefHdrLst and its children, prefixing error messages with path
func (m *StyleDefHdrLst) ValidateWithPath(path string) error {
	if err := m.CT_StyleDefinitionHeaderLst.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
