// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CmAuthorLst struct {
	CT_CommentAuthorList
}

func NewCmAuthorLst() *CmAuthorLst {
	ret := &CmAuthorLst{}
	ret.CT_CommentAuthorList = *NewCT_CommentAuthorList()
	return ret
}

func (m *CmAuthorLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:cmAuthorLst"
	return m.CT_CommentAuthorList.MarshalXML(e, start)
}

func (m *CmAuthorLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_CommentAuthorList = *NewCT_CommentAuthorList()
lCmAuthorLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cmAuthor":
				tmp := NewCT_CommentAuthor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CmAuthor = append(m.CmAuthor, tmp)
			default:
				log.Printf("skipping unsupported element on CmAuthorLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCmAuthorLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CmAuthorLst and its children
func (m *CmAuthorLst) Validate() error {
	return m.ValidateWithPath("CmAuthorLst")
}

// ValidateWithPath validates the CmAuthorLst and its children, prefixing error messages with path
func (m *CmAuthorLst) ValidateWithPath(path string) error {
	if err := m.CT_CommentAuthorList.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
