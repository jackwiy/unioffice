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

type CT_BoxPr struct {
	OpEmu   *CT_OnOff
	NoBreak *CT_OnOff
	Diff    *CT_OnOff
	Brk     *CT_ManualBreak
	Aln     *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func NewCT_BoxPr() *CT_BoxPr {
	ret := &CT_BoxPr{}
	return ret
}

func (m *CT_BoxPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.OpEmu != nil {
		seopEmu := xml.StartElement{Name: xml.Name{Local: "m:opEmu"}}
		e.EncodeElement(m.OpEmu, seopEmu)
	}
	if m.NoBreak != nil {
		senoBreak := xml.StartElement{Name: xml.Name{Local: "m:noBreak"}}
		e.EncodeElement(m.NoBreak, senoBreak)
	}
	if m.Diff != nil {
		sediff := xml.StartElement{Name: xml.Name{Local: "m:diff"}}
		e.EncodeElement(m.Diff, sediff)
	}
	if m.Brk != nil {
		sebrk := xml.StartElement{Name: xml.Name{Local: "m:brk"}}
		e.EncodeElement(m.Brk, sebrk)
	}
	if m.Aln != nil {
		sealn := xml.StartElement{Name: xml.Name{Local: "m:aln"}}
		e.EncodeElement(m.Aln, sealn)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BoxPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BoxPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "opEmu":
				m.OpEmu = NewCT_OnOff()
				if err := d.DecodeElement(m.OpEmu, &el); err != nil {
					return err
				}
			case "noBreak":
				m.NoBreak = NewCT_OnOff()
				if err := d.DecodeElement(m.NoBreak, &el); err != nil {
					return err
				}
			case "diff":
				m.Diff = NewCT_OnOff()
				if err := d.DecodeElement(m.Diff, &el); err != nil {
					return err
				}
			case "brk":
				m.Brk = NewCT_ManualBreak()
				if err := d.DecodeElement(m.Brk, &el); err != nil {
					return err
				}
			case "aln":
				m.Aln = NewCT_OnOff()
				if err := d.DecodeElement(m.Aln, &el); err != nil {
					return err
				}
			case "ctrlPr":
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_BoxPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BoxPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BoxPr and its children
func (m *CT_BoxPr) Validate() error {
	return m.ValidateWithPath("CT_BoxPr")
}

// ValidateWithPath validates the CT_BoxPr and its children, prefixing error messages with path
func (m *CT_BoxPr) ValidateWithPath(path string) error {
	if m.OpEmu != nil {
		if err := m.OpEmu.ValidateWithPath(path + "/OpEmu"); err != nil {
			return err
		}
	}
	if m.NoBreak != nil {
		if err := m.NoBreak.ValidateWithPath(path + "/NoBreak"); err != nil {
			return err
		}
	}
	if m.Diff != nil {
		if err := m.Diff.ValidateWithPath(path + "/Diff"); err != nil {
			return err
		}
	}
	if m.Brk != nil {
		if err := m.Brk.ValidateWithPath(path + "/Brk"); err != nil {
			return err
		}
	}
	if m.Aln != nil {
		if err := m.Aln.ValidateWithPath(path + "/Aln"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
