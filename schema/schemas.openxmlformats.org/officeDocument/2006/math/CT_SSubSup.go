// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_SSubSup struct {
	SSubSupPr *CT_SSubSupPr
	E         *CT_OMathArg
	Sub       *CT_OMathArg
	Sup       *CT_OMathArg
}

func NewCT_SSubSup() *CT_SSubSup {
	ret := &CT_SSubSup{}
	ret.E = NewCT_OMathArg()
	ret.Sub = NewCT_OMathArg()
	ret.Sup = NewCT_OMathArg()
	return ret
}

func (m *CT_SSubSup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.SSubSupPr != nil {
		sesSubSupPr := xml.StartElement{Name: xml.Name{Local: "m:sSubSupPr"}}
		e.EncodeElement(m.SSubSupPr, sesSubSupPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	sesub := xml.StartElement{Name: xml.Name{Local: "m:sub"}}
	e.EncodeElement(m.Sub, sesub)
	sesup := xml.StartElement{Name: xml.Name{Local: "m:sup"}}
	e.EncodeElement(m.Sup, sesup)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SSubSup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
	m.Sub = NewCT_OMathArg()
	m.Sup = NewCT_OMathArg()
lCT_SSubSup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sSubSupPr":
				m.SSubSupPr = NewCT_SSubSupPr()
				if err := d.DecodeElement(m.SSubSupPr, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			case "sub":
				if err := d.DecodeElement(m.Sub, &el); err != nil {
					return err
				}
			case "sup":
				if err := d.DecodeElement(m.Sup, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SSubSup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SSubSup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSubSup and its children
func (m *CT_SSubSup) Validate() error {
	return m.ValidateWithPath("CT_SSubSup")
}

// ValidateWithPath validates the CT_SSubSup and its children, prefixing error messages with path
func (m *CT_SSubSup) ValidateWithPath(path string) error {
	if m.SSubSupPr != nil {
		if err := m.SSubSupPr.ValidateWithPath(path + "/SSubSupPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	if err := m.Sub.ValidateWithPath(path + "/Sub"); err != nil {
		return err
	}
	if err := m.Sup.ValidateWithPath(path + "/Sup"); err != nil {
		return err
	}
	return nil
}
