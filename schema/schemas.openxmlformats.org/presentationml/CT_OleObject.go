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
	"strconv"
)

type CT_OleObject struct {
	// Embedded Object ProgID
	ProgIdAttr     *string
	Choice         *CT_OleObjectChoice
	Pic            *CT_Picture
	SpidAttr       *string
	NameAttr       *string
	ShowAsIconAttr *bool
	IdAttr         *string
	ImgWAttr       *int32
	ImgHAttr       *int32
}

func NewCT_OleObject() *CT_OleObject {
	ret := &CT_OleObject{}
	ret.Choice = NewCT_OleObjectChoice()
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
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.ShowAsIconAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAsIcon"},
			Value: fmt.Sprintf("%v", *m.ShowAsIconAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.ImgWAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imgW"},
			Value: fmt.Sprintf("%v", *m.ImgWAttr)})
	}
	if m.ImgHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imgH"},
			Value: fmt.Sprintf("%v", *m.ImgHAttr)})
	}
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, start)
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "p:pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_OleObjectChoice()
	for _, attr := range start.Attr {
		if attr.Name.Local == "progId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIdAttr = &parsed
		}
		if attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "showAsIcon" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAsIconAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "imgW" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgWAttr = &pt
		}
		if attr.Name.Local == "imgH" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgHAttr = &pt
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
			case "embed":
				m.Choice = NewCT_OleObjectChoice()
				if err := d.DecodeElement(&m.Choice.Embed, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "link":
				m.Choice = NewCT_OleObjectChoice()
				if err := d.DecodeElement(&m.Choice.Link, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "pic":
				m.Pic = NewCT_Picture()
				if err := d.DecodeElement(m.Pic, &el); err != nil {
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
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	if m.Pic != nil {
		if err := m.Pic.ValidateWithPath(path + "/Pic"); err != nil {
			return err
		}
	}
	if m.ImgWAttr != nil {
		if *m.ImgWAttr < 0 {
			return fmt.Errorf("%s/m.ImgWAttr must be >= 0 (have %v)", path, *m.ImgWAttr)
		}
	}
	if m.ImgHAttr != nil {
		if *m.ImgHAttr < 0 {
			return fmt.Errorf("%s/m.ImgHAttr must be >= 0 (have %v)", path, *m.ImgHAttr)
		}
	}
	return nil
}
