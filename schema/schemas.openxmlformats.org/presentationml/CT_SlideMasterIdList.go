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
)

type CT_SlideMasterIdList struct {
	// Slide Master ID
	SldMasterId []*CT_SlideMasterIdListEntry
}

func NewCT_SlideMasterIdList() *CT_SlideMasterIdList {
	ret := &CT_SlideMasterIdList{}
	return ret
}

func (m *CT_SlideMasterIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.SldMasterId != nil {
		sesldMasterId := xml.StartElement{Name: xml.Name{Local: "p:sldMasterId"}}
		e.EncodeElement(m.SldMasterId, sesldMasterId)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideMasterIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideMasterIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sldMasterId":
				tmp := NewCT_SlideMasterIdListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SldMasterId = append(m.SldMasterId, tmp)
			default:
				log.Printf("skipping unsupported element on CT_SlideMasterIdList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideMasterIdList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideMasterIdList and its children
func (m *CT_SlideMasterIdList) Validate() error {
	return m.ValidateWithPath("CT_SlideMasterIdList")
}

// ValidateWithPath validates the CT_SlideMasterIdList and its children, prefixing error messages with path
func (m *CT_SlideMasterIdList) ValidateWithPath(path string) error {
	for i, v := range m.SldMasterId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SldMasterId[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
