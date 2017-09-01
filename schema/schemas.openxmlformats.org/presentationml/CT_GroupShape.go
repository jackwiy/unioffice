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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_GroupShape struct {
	// Non-Visual Properties for a Group Shape
	NvGrpSpPr *CT_GroupShapeNonVisual
	// Group Shape Properties
	GrpSpPr *drawingml.CT_GroupShapeProperties
	Choice  []*CT_GroupShapeChoice
	ExtLst  *CT_ExtensionListModify
}

func NewCT_GroupShape() *CT_GroupShape {
	ret := &CT_GroupShape{}
	ret.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	ret.GrpSpPr = drawingml.NewCT_GroupShapeProperties()
	return ret
}

func (m *CT_GroupShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	senvGrpSpPr := xml.StartElement{Name: xml.Name{Local: "p:nvGrpSpPr"}}
	e.EncodeElement(m.NvGrpSpPr, senvGrpSpPr)
	segrpSpPr := xml.StartElement{Name: xml.Name{Local: "p:grpSpPr"}}
	e.EncodeElement(m.GrpSpPr, segrpSpPr)
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, start)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	m.GrpSpPr = drawingml.NewCT_GroupShapeProperties()
lCT_GroupShape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nvGrpSpPr":
				if err := d.DecodeElement(m.NvGrpSpPr, &el); err != nil {
					return err
				}
			case "grpSpPr":
				if err := d.DecodeElement(m.GrpSpPr, &el); err != nil {
					return err
				}
			case "sp":
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.Sp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "grpSp":
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.GrpSp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "graphicFrame":
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.GraphicFrame, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "cxnSp":
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.CxnSp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "pic":
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.Pic, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "contentPart":
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.ContentPart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_GroupShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShape and its children
func (m *CT_GroupShape) Validate() error {
	return m.ValidateWithPath("CT_GroupShape")
}

// ValidateWithPath validates the CT_GroupShape and its children, prefixing error messages with path
func (m *CT_GroupShape) ValidateWithPath(path string) error {
	if err := m.NvGrpSpPr.ValidateWithPath(path + "/NvGrpSpPr"); err != nil {
		return err
	}
	if err := m.GrpSpPr.ValidateWithPath(path + "/GrpSpPr"); err != nil {
		return err
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
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
