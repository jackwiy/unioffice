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

type CT_CommonSlideData struct {
	// Name
	NameAttr *string
	// Slide Background
	Bg *CT_Background
	// Shape Tree
	SpTree *CT_GroupShape
	// Customer Data List
	CustDataLst *CT_CustomerDataList
	// List of controls
	Controls *CT_ControlList
	ExtLst   *CT_ExtensionList
}

func NewCT_CommonSlideData() *CT_CommonSlideData {
	ret := &CT_CommonSlideData{}
	ret.SpTree = NewCT_GroupShape()
	return ret
}

func (m *CT_CommonSlideData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	if m.Bg != nil {
		sebg := xml.StartElement{Name: xml.Name{Local: "p:bg"}}
		e.EncodeElement(m.Bg, sebg)
	}
	sespTree := xml.StartElement{Name: xml.Name{Local: "p:spTree"}}
	e.EncodeElement(m.SpTree, sespTree)
	if m.CustDataLst != nil {
		secustDataLst := xml.StartElement{Name: xml.Name{Local: "p:custDataLst"}}
		e.EncodeElement(m.CustDataLst, secustDataLst)
	}
	if m.Controls != nil {
		secontrols := xml.StartElement{Name: xml.Name{Local: "p:controls"}}
		e.EncodeElement(m.Controls, secontrols)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommonSlideData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SpTree = NewCT_GroupShape()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
	}
lCT_CommonSlideData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bg":
				m.Bg = NewCT_Background()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case "spTree":
				if err := d.DecodeElement(m.SpTree, &el); err != nil {
					return err
				}
			case "custDataLst":
				m.CustDataLst = NewCT_CustomerDataList()
				if err := d.DecodeElement(m.CustDataLst, &el); err != nil {
					return err
				}
			case "controls":
				m.Controls = NewCT_ControlList()
				if err := d.DecodeElement(m.Controls, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_CommonSlideData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommonSlideData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommonSlideData and its children
func (m *CT_CommonSlideData) Validate() error {
	return m.ValidateWithPath("CT_CommonSlideData")
}

// ValidateWithPath validates the CT_CommonSlideData and its children, prefixing error messages with path
func (m *CT_CommonSlideData) ValidateWithPath(path string) error {
	if m.Bg != nil {
		if err := m.Bg.ValidateWithPath(path + "/Bg"); err != nil {
			return err
		}
	}
	if err := m.SpTree.ValidateWithPath(path + "/SpTree"); err != nil {
		return err
	}
	if m.CustDataLst != nil {
		if err := m.CustDataLst.ValidateWithPath(path + "/CustDataLst"); err != nil {
			return err
		}
	}
	if m.Controls != nil {
		if err := m.Controls.ValidateWithPath(path + "/Controls"); err != nil {
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
