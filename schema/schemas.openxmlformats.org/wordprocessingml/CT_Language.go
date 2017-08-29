// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
)

type CT_Language struct {
	// Latin Language
	ValAttr *string
	// East Asian Language
	EastAsiaAttr *string
	// Complex Script Language
	BidiAttr *string
}

func NewCT_Language() *CT_Language {
	ret := &CT_Language{}
	return ret
}
func (m *CT_Language) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.EastAsiaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:eastAsia"},
			Value: fmt.Sprintf("%v", *m.EastAsiaAttr)})
	}
	if m.BidiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bidi"},
			Value: fmt.Sprintf("%v", *m.BidiAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Language) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
		if attr.Name.Local == "eastAsia" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EastAsiaAttr = &parsed
		}
		if attr.Name.Local == "bidi" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BidiAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Language: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Language) Validate() error {
	return m.ValidateWithPath("CT_Language")
}
func (m *CT_Language) ValidateWithPath(path string) error {
	return nil
}