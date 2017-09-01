// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_ControlPr struct {
	// Locked Flag
	LockedAttr *bool
	// Default Size Flag
	DefaultSizeAttr *bool
	// Print Flag
	PrintAttr *bool
	// Disabled Flag
	DisabledAttr *bool
	// Recalculation Flag
	RecalcAlwaysAttr *bool
	// UI Object Flag
	UiObjectAttr *bool
	// Automatic Fill Flag
	AutoFillAttr *bool
	// Automatic Line Flag
	AutoLineAttr *bool
	// Automatic Size Flag
	AutoPictAttr *bool
	// Custom Function
	MacroAttr *string
	// Alternative Text
	AltTextAttr *string
	// Linked Formula
	LinkedCellAttr *string
	// List Items Source Range
	ListFillRangeAttr *string
	// Image Format
	CfAttr *string
	IdAttr *string
	// Object Cell Anchor
	Anchor *CT_ObjectAnchor
}

func NewCT_ControlPr() *CT_ControlPr {
	ret := &CT_ControlPr{}
	ret.Anchor = NewCT_ObjectAnchor()
	return ret
}

func (m *CT_ControlPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"},
			Value: fmt.Sprintf("%v", *m.LockedAttr)})
	}
	if m.DefaultSizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultSize"},
			Value: fmt.Sprintf("%v", *m.DefaultSizeAttr)})
	}
	if m.PrintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "print"},
			Value: fmt.Sprintf("%v", *m.PrintAttr)})
	}
	if m.DisabledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disabled"},
			Value: fmt.Sprintf("%v", *m.DisabledAttr)})
	}
	if m.RecalcAlwaysAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "recalcAlways"},
			Value: fmt.Sprintf("%v", *m.RecalcAlwaysAttr)})
	}
	if m.UiObjectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiObject"},
			Value: fmt.Sprintf("%v", *m.UiObjectAttr)})
	}
	if m.AutoFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFill"},
			Value: fmt.Sprintf("%v", *m.AutoFillAttr)})
	}
	if m.AutoLineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoLine"},
			Value: fmt.Sprintf("%v", *m.AutoLineAttr)})
	}
	if m.AutoPictAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoPict"},
			Value: fmt.Sprintf("%v", *m.AutoPictAttr)})
	}
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.AltTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "altText"},
			Value: fmt.Sprintf("%v", *m.AltTextAttr)})
	}
	if m.LinkedCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "linkedCell"},
			Value: fmt.Sprintf("%v", *m.LinkedCellAttr)})
	}
	if m.ListFillRangeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "listFillRange"},
			Value: fmt.Sprintf("%v", *m.ListFillRangeAttr)})
	}
	if m.CfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cf"},
			Value: fmt.Sprintf("%v", *m.CfAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	seanchor := xml.StartElement{Name: xml.Name{Local: "x:anchor"}}
	e.EncodeElement(m.Anchor, seanchor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ControlPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Anchor = NewCT_ObjectAnchor()
	for _, attr := range start.Attr {
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = &parsed
		}
		if attr.Name.Local == "defaultSize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultSizeAttr = &parsed
		}
		if attr.Name.Local == "print" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PrintAttr = &parsed
		}
		if attr.Name.Local == "disabled" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisabledAttr = &parsed
		}
		if attr.Name.Local == "recalcAlways" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RecalcAlwaysAttr = &parsed
		}
		if attr.Name.Local == "uiObject" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiObjectAttr = &parsed
		}
		if attr.Name.Local == "autoFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoFillAttr = &parsed
		}
		if attr.Name.Local == "autoLine" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoLineAttr = &parsed
		}
		if attr.Name.Local == "autoPict" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoPictAttr = &parsed
		}
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
		}
		if attr.Name.Local == "altText" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltTextAttr = &parsed
		}
		if attr.Name.Local == "linkedCell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LinkedCellAttr = &parsed
		}
		if attr.Name.Local == "listFillRange" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ListFillRangeAttr = &parsed
		}
		if attr.Name.Local == "cf" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CfAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
lCT_ControlPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "anchor":
				if err := d.DecodeElement(m.Anchor, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ControlPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ControlPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ControlPr and its children
func (m *CT_ControlPr) Validate() error {
	return m.ValidateWithPath("CT_ControlPr")
}

// ValidateWithPath validates the CT_ControlPr and its children, prefixing error messages with path
func (m *CT_ControlPr) ValidateWithPath(path string) error {
	if err := m.Anchor.ValidateWithPath(path + "/Anchor"); err != nil {
		return err
	}
	return nil
}
