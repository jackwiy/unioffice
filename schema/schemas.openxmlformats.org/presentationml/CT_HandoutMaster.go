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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_HandoutMaster struct {
	// Common slide data for handout master
	CSld *CT_CommonSlideData
	// Color Scheme Map
	ClrMap *drawingml.CT_ColorMapping
	// Header/Footer information for a handout master
	Hf     *CT_HeaderFooter
	ExtLst *CT_ExtensionListModify
}

func NewCT_HandoutMaster() *CT_HandoutMaster {
	ret := &CT_HandoutMaster{}
	ret.CSld = NewCT_CommonSlideData()
	ret.ClrMap = drawingml.NewCT_ColorMapping()
	return ret
}

func (m *CT_HandoutMaster) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	secSld := xml.StartElement{Name: xml.Name{Local: "p:cSld"}}
	e.EncodeElement(m.CSld, secSld)
	seclrMap := xml.StartElement{Name: xml.Name{Local: "p:clrMap"}}
	e.EncodeElement(m.ClrMap, seclrMap)
	if m.Hf != nil {
		sehf := xml.StartElement{Name: xml.Name{Local: "p:hf"}}
		e.EncodeElement(m.Hf, sehf)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HandoutMaster) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CSld = NewCT_CommonSlideData()
	m.ClrMap = drawingml.NewCT_ColorMapping()
lCT_HandoutMaster:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cSld":
				if err := d.DecodeElement(m.CSld, &el); err != nil {
					return err
				}
			case "clrMap":
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			case "hf":
				m.Hf = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.Hf, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_HandoutMaster %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_HandoutMaster
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_HandoutMaster and its children
func (m *CT_HandoutMaster) Validate() error {
	return m.ValidateWithPath("CT_HandoutMaster")
}

// ValidateWithPath validates the CT_HandoutMaster and its children, prefixing error messages with path
func (m *CT_HandoutMaster) ValidateWithPath(path string) error {
	if err := m.CSld.ValidateWithPath(path + "/CSld"); err != nil {
		return err
	}
	if err := m.ClrMap.ValidateWithPath(path + "/ClrMap"); err != nil {
		return err
	}
	if m.Hf != nil {
		if err := m.Hf.ValidateWithPath(path + "/Hf"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
