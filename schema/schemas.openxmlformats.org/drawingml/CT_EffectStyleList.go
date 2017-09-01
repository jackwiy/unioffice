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

type CT_EffectStyleList struct {
	EffectStyle []*CT_EffectStyleItem
}

func NewCT_EffectStyleList() *CT_EffectStyleList {
	ret := &CT_EffectStyleList{}
	return ret
}

func (m *CT_EffectStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seeffectStyle := xml.StartElement{Name: xml.Name{Local: "a:effectStyle"}}
	e.EncodeElement(m.EffectStyle, seeffectStyle)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EffectStyleList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EffectStyleList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "effectStyle":
				tmp := NewCT_EffectStyleItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.EffectStyle = append(m.EffectStyle, tmp)
			default:
				log.Printf("skipping unsupported element on CT_EffectStyleList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EffectStyleList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EffectStyleList and its children
func (m *CT_EffectStyleList) Validate() error {
	return m.ValidateWithPath("CT_EffectStyleList")
}

// ValidateWithPath validates the CT_EffectStyleList and its children, prefixing error messages with path
func (m *CT_EffectStyleList) ValidateWithPath(path string) error {
	for i, v := range m.EffectStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EffectStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
