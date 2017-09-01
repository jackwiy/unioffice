// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_AdjLst struct {
	Adj []*CT_Adj
}

func NewCT_AdjLst() *CT_AdjLst {
	ret := &CT_AdjLst{}
	return ret
}

func (m *CT_AdjLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Adj != nil {
		seadj := xml.StartElement{Name: xml.Name{Local: "adj"}}
		e.EncodeElement(m.Adj, seadj)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AdjLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AdjLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "adj":
				tmp := NewCT_Adj()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Adj = append(m.Adj, tmp)
			default:
				log.Printf("skipping unsupported element on CT_AdjLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AdjLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AdjLst and its children
func (m *CT_AdjLst) Validate() error {
	return m.ValidateWithPath("CT_AdjLst")
}

// ValidateWithPath validates the CT_AdjLst and its children, prefixing error messages with path
func (m *CT_AdjLst) ValidateWithPath(path string) error {
	for i, v := range m.Adj {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Adj[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
