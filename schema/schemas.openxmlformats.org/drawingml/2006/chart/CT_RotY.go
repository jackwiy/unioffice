// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_RotY struct {
	ValAttr *uint16
}

func NewCT_RotY() *CT_RotY {
	ret := &CT_RotY{}
	return ret
}
func (m *CT_RotY) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RotY) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint16(parsed)
			m.ValAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RotY: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RotY) Validate() error {
	return m.ValidateWithPath("CT_RotY")
}
func (m *CT_RotY) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if *m.ValAttr < 0 {
			return fmt.Errorf("%s/m.ValAttr must be >= 0 (have %v)", path, *m.ValAttr)
		}
		if *m.ValAttr > 360 {
			return fmt.Errorf("%s/m.ValAttr must be <= 360 (have %v)", path, *m.ValAttr)
		}
	}
	return nil
}