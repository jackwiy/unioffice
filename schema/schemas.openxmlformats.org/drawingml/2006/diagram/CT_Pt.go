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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Pt struct {
	ModelIdAttr ST_ModelId
	TypeAttr    ST_PtType
	CxnIdAttr   *ST_ModelId
	PrSet       *CT_ElemPropSet
	SpPr        *drawingml.CT_ShapeProperties
	T           *drawingml.CT_TextBody
	ExtLst      *drawingml.CT_OfficeArtExtensionList
}

func NewCT_Pt() *CT_Pt {
	ret := &CT_Pt{}
	return ret
}

func (m *CT_Pt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "modelId"},
		Value: fmt.Sprintf("%v", m.ModelIdAttr)})
	if m.TypeAttr != ST_PtTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CxnIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cxnId"},
			Value: fmt.Sprintf("%v", *m.CxnIdAttr)})
	}
	e.EncodeToken(start)
	if m.PrSet != nil {
		seprSet := xml.StartElement{Name: xml.Name{Local: "prSet"}}
		e.EncodeElement(m.PrSet, seprSet)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "t"}}
		e.EncodeElement(m.T, set)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Pt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "modelId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.ModelIdAttr = parsed
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cxnId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.CxnIdAttr = &parsed
		}
	}
lCT_Pt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "prSet":
				m.PrSet = NewCT_ElemPropSet()
				if err := d.DecodeElement(m.PrSet, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "t":
				m.T = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Pt %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Pt
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Pt and its children
func (m *CT_Pt) Validate() error {
	return m.ValidateWithPath("CT_Pt")
}

// ValidateWithPath validates the CT_Pt and its children, prefixing error messages with path
func (m *CT_Pt) ValidateWithPath(path string) error {
	if err := m.ModelIdAttr.ValidateWithPath(path + "/ModelIdAttr"); err != nil {
		return err
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.CxnIdAttr != nil {
		if err := m.CxnIdAttr.ValidateWithPath(path + "/CxnIdAttr"); err != nil {
			return err
		}
	}
	if m.PrSet != nil {
		if err := m.PrSet.ValidateWithPath(path + "/PrSet"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.T != nil {
		if err := m.T.ValidateWithPath(path + "/T"); err != nil {
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
