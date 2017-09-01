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

type CT_OleObject struct {
	// Embedded Object ProgId
	ProgIdAttr *string
	// Data or View Aspect
	DvAspectAttr ST_DvAspect
	// Embedded Object's Link Moniker
	LinkAttr *string
	// Linked Embedded Object Update
	OleUpdateAttr ST_OleUpdate
	// Auto Load
	AutoLoadAttr *bool
	// Shape Id
	ShapeIdAttr uint32
	IdAttr      *string
	// Embedded Object Properties
	ObjectPr *CT_ObjectPr
}

func NewCT_OleObject() *CT_OleObject {
	ret := &CT_OleObject{}
	return ret
}

func (m *CT_OleObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ProgIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "progId"},
			Value: fmt.Sprintf("%v", *m.ProgIdAttr)})
	}
	if m.DvAspectAttr != ST_DvAspectUnset {
		attr, err := m.DvAspectAttr.MarshalXMLAttr(xml.Name{Local: "dvAspect"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LinkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "link"},
			Value: fmt.Sprintf("%v", *m.LinkAttr)})
	}
	if m.OleUpdateAttr != ST_OleUpdateUnset {
		attr, err := m.OleUpdateAttr.MarshalXMLAttr(xml.Name{Local: "oleUpdate"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AutoLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoLoad"},
			Value: fmt.Sprintf("%v", *m.AutoLoadAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shapeId"},
		Value: fmt.Sprintf("%v", m.ShapeIdAttr)})
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	if m.ObjectPr != nil {
		seobjectPr := xml.StartElement{Name: xml.Name{Local: "x:objectPr"}}
		e.EncodeElement(m.ObjectPr, seobjectPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "progId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIdAttr = &parsed
		}
		if attr.Name.Local == "dvAspect" {
			m.DvAspectAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "link" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LinkAttr = &parsed
		}
		if attr.Name.Local == "oleUpdate" {
			m.OleUpdateAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "autoLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoLoadAttr = &parsed
		}
		if attr.Name.Local == "shapeId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ShapeIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
lCT_OleObject:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "objectPr":
				m.ObjectPr = NewCT_ObjectPr()
				if err := d.DecodeElement(m.ObjectPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_OleObject %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObject
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObject and its children
func (m *CT_OleObject) Validate() error {
	return m.ValidateWithPath("CT_OleObject")
}

// ValidateWithPath validates the CT_OleObject and its children, prefixing error messages with path
func (m *CT_OleObject) ValidateWithPath(path string) error {
	if err := m.DvAspectAttr.ValidateWithPath(path + "/DvAspectAttr"); err != nil {
		return err
	}
	if err := m.OleUpdateAttr.ValidateWithPath(path + "/OleUpdateAttr"); err != nil {
		return err
	}
	if m.ObjectPr != nil {
		if err := m.ObjectPr.ValidateWithPath(path + "/ObjectPr"); err != nil {
			return err
		}
	}
	return nil
}
