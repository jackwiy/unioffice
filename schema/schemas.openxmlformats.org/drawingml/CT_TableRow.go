// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TableRow struct {
	HAttr  ST_Coordinate
	Tc     []*CT_TableCell
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_TableRow() *CT_TableRow {
	ret := &CT_TableRow{}
	return ret
}
func (m *CT_TableRow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "h"},
		Value: fmt.Sprintf("%v", m.HAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	if m.Tc != nil {
		setc := xml.StartElement{Name: xml.Name{Local: "a:tc"}}
		e.EncodeElement(m.Tc, setc)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TableRow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "h" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.HAttr = parsed
		}
	}
lCT_TableRow:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tc":
				tmp := NewCT_TableCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tc = append(m.Tc, tmp)
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableRow
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TableRow) Validate() error {
	return m.ValidateWithPath("CT_TableRow")
}
func (m *CT_TableRow) ValidateWithPath(path string) error {
	if err := m.HAttr.ValidateWithPath(path + "/HAttr"); err != nil {
		return err
	}
	for i, v := range m.Tc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tc[%d]", path, i)); err != nil {
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