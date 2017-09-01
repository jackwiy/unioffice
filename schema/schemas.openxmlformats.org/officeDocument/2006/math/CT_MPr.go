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

type CT_MPr struct {
	BaseJc  *CT_YAlign
	PlcHide *CT_OnOff
	RSpRule *CT_SpacingRule
	CGpRule *CT_SpacingRule
	RSp     *CT_UnSignedInteger
	CSp     *CT_UnSignedInteger
	CGp     *CT_UnSignedInteger
	Mcs     *CT_MCS
	CtrlPr  *CT_CtrlPr
}

func NewCT_MPr() *CT_MPr {
	ret := &CT_MPr{}
	return ret
}

func (m *CT_MPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.BaseJc != nil {
		sebaseJc := xml.StartElement{Name: xml.Name{Local: "m:baseJc"}}
		e.EncodeElement(m.BaseJc, sebaseJc)
	}
	if m.PlcHide != nil {
		seplcHide := xml.StartElement{Name: xml.Name{Local: "m:plcHide"}}
		e.EncodeElement(m.PlcHide, seplcHide)
	}
	if m.RSpRule != nil {
		serSpRule := xml.StartElement{Name: xml.Name{Local: "m:rSpRule"}}
		e.EncodeElement(m.RSpRule, serSpRule)
	}
	if m.CGpRule != nil {
		secGpRule := xml.StartElement{Name: xml.Name{Local: "m:cGpRule"}}
		e.EncodeElement(m.CGpRule, secGpRule)
	}
	if m.RSp != nil {
		serSp := xml.StartElement{Name: xml.Name{Local: "m:rSp"}}
		e.EncodeElement(m.RSp, serSp)
	}
	if m.CSp != nil {
		secSp := xml.StartElement{Name: xml.Name{Local: "m:cSp"}}
		e.EncodeElement(m.CSp, secSp)
	}
	if m.CGp != nil {
		secGp := xml.StartElement{Name: xml.Name{Local: "m:cGp"}}
		e.EncodeElement(m.CGp, secGp)
	}
	if m.Mcs != nil {
		semcs := xml.StartElement{Name: xml.Name{Local: "m:mcs"}}
		e.EncodeElement(m.Mcs, semcs)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "baseJc":
				m.BaseJc = NewCT_YAlign()
				if err := d.DecodeElement(m.BaseJc, &el); err != nil {
					return err
				}
			case "plcHide":
				m.PlcHide = NewCT_OnOff()
				if err := d.DecodeElement(m.PlcHide, &el); err != nil {
					return err
				}
			case "rSpRule":
				m.RSpRule = NewCT_SpacingRule()
				if err := d.DecodeElement(m.RSpRule, &el); err != nil {
					return err
				}
			case "cGpRule":
				m.CGpRule = NewCT_SpacingRule()
				if err := d.DecodeElement(m.CGpRule, &el); err != nil {
					return err
				}
			case "rSp":
				m.RSp = NewCT_UnSignedInteger()
				if err := d.DecodeElement(m.RSp, &el); err != nil {
					return err
				}
			case "cSp":
				m.CSp = NewCT_UnSignedInteger()
				if err := d.DecodeElement(m.CSp, &el); err != nil {
					return err
				}
			case "cGp":
				m.CGp = NewCT_UnSignedInteger()
				if err := d.DecodeElement(m.CGp, &el); err != nil {
					return err
				}
			case "mcs":
				m.Mcs = NewCT_MCS()
				if err := d.DecodeElement(m.Mcs, &el); err != nil {
					return err
				}
			case "ctrlPr":
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_MPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MPr and its children
func (m *CT_MPr) Validate() error {
	return m.ValidateWithPath("CT_MPr")
}

// ValidateWithPath validates the CT_MPr and its children, prefixing error messages with path
func (m *CT_MPr) ValidateWithPath(path string) error {
	if m.BaseJc != nil {
		if err := m.BaseJc.ValidateWithPath(path + "/BaseJc"); err != nil {
			return err
		}
	}
	if m.PlcHide != nil {
		if err := m.PlcHide.ValidateWithPath(path + "/PlcHide"); err != nil {
			return err
		}
	}
	if m.RSpRule != nil {
		if err := m.RSpRule.ValidateWithPath(path + "/RSpRule"); err != nil {
			return err
		}
	}
	if m.CGpRule != nil {
		if err := m.CGpRule.ValidateWithPath(path + "/CGpRule"); err != nil {
			return err
		}
	}
	if m.RSp != nil {
		if err := m.RSp.ValidateWithPath(path + "/RSp"); err != nil {
			return err
		}
	}
	if m.CSp != nil {
		if err := m.CSp.ValidateWithPath(path + "/CSp"); err != nil {
			return err
		}
	}
	if m.CGp != nil {
		if err := m.CGp.ValidateWithPath(path + "/CGp"); err != nil {
			return err
		}
	}
	if m.Mcs != nil {
		if err := m.Mcs.ValidateWithPath(path + "/Mcs"); err != nil {
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
