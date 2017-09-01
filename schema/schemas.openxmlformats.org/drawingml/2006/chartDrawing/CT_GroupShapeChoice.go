// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_GroupShapeChoice struct {
	Sp           []*CT_Shape
	GrpSp        []*CT_GroupShape
	GraphicFrame []*CT_GraphicFrame
	CxnSp        []*CT_Connector
	Pic          []*CT_Picture
}

func NewCT_GroupShapeChoice() *CT_GroupShapeChoice {
	ret := &CT_GroupShapeChoice{}
	return ret
}

func (m *CT_GroupShapeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Sp != nil {
		sesp := xml.StartElement{Name: xml.Name{Local: "sp"}}
		e.EncodeElement(m.Sp, sesp)
	}
	if m.GrpSp != nil {
		segrpSp := xml.StartElement{Name: xml.Name{Local: "grpSp"}}
		e.EncodeElement(m.GrpSp, segrpSp)
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "graphicFrame"}}
		e.EncodeElement(m.GraphicFrame, segraphicFrame)
	}
	if m.CxnSp != nil {
		secxnSp := xml.StartElement{Name: xml.Name{Local: "cxnSp"}}
		e.EncodeElement(m.CxnSp, secxnSp)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	return nil
}

func (m *CT_GroupShapeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GroupShapeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sp":
				tmp := NewCT_Shape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Sp = append(m.Sp, tmp)
			case "grpSp":
				tmp := NewCT_GroupShape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GrpSp = append(m.GrpSp, tmp)
			case "graphicFrame":
				tmp := NewCT_GraphicFrame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GraphicFrame = append(m.GraphicFrame, tmp)
			case "cxnSp":
				tmp := NewCT_Connector()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CxnSp = append(m.CxnSp, tmp)
			case "pic":
				tmp := NewCT_Picture()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pic = append(m.Pic, tmp)
			default:
				log.Printf("skipping unsupported element on CT_GroupShapeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupShapeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShapeChoice and its children
func (m *CT_GroupShapeChoice) Validate() error {
	return m.ValidateWithPath("CT_GroupShapeChoice")
}

// ValidateWithPath validates the CT_GroupShapeChoice and its children, prefixing error messages with path
func (m *CT_GroupShapeChoice) ValidateWithPath(path string) error {
	for i, v := range m.Sp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Sp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GrpSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GrpSp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GraphicFrame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GraphicFrame[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CxnSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CxnSp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Pic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pic[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
