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

type CT_SdtRun struct {
	// Structured Document Tag Properties
	SdtPr *CT_SdtPr
	// Structured Document Tag End Character Properties
	SdtEndPr *CT_SdtEndPr
	// Inline-Level Structured Document Tag Content
	SdtContent *CT_SdtContentRun
}

func NewCT_SdtRun() *CT_SdtRun {
	ret := &CT_SdtRun{}
	return ret
}

func (m *CT_SdtRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.SdtPr != nil {
		sesdtPr := xml.StartElement{Name: xml.Name{Local: "w:sdtPr"}}
		e.EncodeElement(m.SdtPr, sesdtPr)
	}
	if m.SdtEndPr != nil {
		sesdtEndPr := xml.StartElement{Name: xml.Name{Local: "w:sdtEndPr"}}
		e.EncodeElement(m.SdtEndPr, sesdtEndPr)
	}
	if m.SdtContent != nil {
		sesdtContent := xml.StartElement{Name: xml.Name{Local: "w:sdtContent"}}
		e.EncodeElement(m.SdtContent, sesdtContent)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sdtPr":
				m.SdtPr = NewCT_SdtPr()
				if err := d.DecodeElement(m.SdtPr, &el); err != nil {
					return err
				}
			case "sdtEndPr":
				m.SdtEndPr = NewCT_SdtEndPr()
				if err := d.DecodeElement(m.SdtEndPr, &el); err != nil {
					return err
				}
			case "sdtContent":
				m.SdtContent = NewCT_SdtContentRun()
				if err := d.DecodeElement(m.SdtContent, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SdtRun %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtRun
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtRun and its children
func (m *CT_SdtRun) Validate() error {
	return m.ValidateWithPath("CT_SdtRun")
}

// ValidateWithPath validates the CT_SdtRun and its children, prefixing error messages with path
func (m *CT_SdtRun) ValidateWithPath(path string) error {
	if m.SdtPr != nil {
		if err := m.SdtPr.ValidateWithPath(path + "/SdtPr"); err != nil {
			return err
		}
	}
	if m.SdtEndPr != nil {
		if err := m.SdtEndPr.ValidateWithPath(path + "/SdtEndPr"); err != nil {
			return err
		}
	}
	if m.SdtContent != nil {
		if err := m.SdtContent.ValidateWithPath(path + "/SdtContent"); err != nil {
			return err
		}
	}
	return nil
}
