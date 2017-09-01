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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_ElemPropSet struct {
	PresAssocIDAttr          *ST_ModelId
	PresNameAttr             *string
	PresStyleLblAttr         *string
	PresStyleIdxAttr         *int32
	PresStyleCntAttr         *int32
	LoTypeIdAttr             *string
	LoCatIdAttr              *string
	QsTypeIdAttr             *string
	QsCatIdAttr              *string
	CsTypeIdAttr             *string
	CsCatIdAttr              *string
	Coherent3DOffAttr        *bool
	PhldrTAttr               *string
	PhldrAttr                *bool
	CustAngAttr              *int32
	CustFlipVertAttr         *bool
	CustFlipHorAttr          *bool
	CustSzXAttr              *int32
	CustSzYAttr              *int32
	CustScaleXAttr           *ST_PrSetCustVal
	CustScaleYAttr           *ST_PrSetCustVal
	CustTAttr                *bool
	CustLinFactXAttr         *ST_PrSetCustVal
	CustLinFactYAttr         *ST_PrSetCustVal
	CustLinFactNeighborXAttr *ST_PrSetCustVal
	CustLinFactNeighborYAttr *ST_PrSetCustVal
	CustRadScaleRadAttr      *ST_PrSetCustVal
	CustRadScaleIncAttr      *ST_PrSetCustVal
	PresLayoutVars           *CT_LayoutVariablePropertySet
	Style                    *drawingml.CT_ShapeStyle
}

func NewCT_ElemPropSet() *CT_ElemPropSet {
	ret := &CT_ElemPropSet{}
	return ret
}

func (m *CT_ElemPropSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PresAssocIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presAssocID"},
			Value: fmt.Sprintf("%v", *m.PresAssocIDAttr)})
	}
	if m.PresNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presName"},
			Value: fmt.Sprintf("%v", *m.PresNameAttr)})
	}
	if m.PresStyleLblAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presStyleLbl"},
			Value: fmt.Sprintf("%v", *m.PresStyleLblAttr)})
	}
	if m.PresStyleIdxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presStyleIdx"},
			Value: fmt.Sprintf("%v", *m.PresStyleIdxAttr)})
	}
	if m.PresStyleCntAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presStyleCnt"},
			Value: fmt.Sprintf("%v", *m.PresStyleCntAttr)})
	}
	if m.LoTypeIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "loTypeId"},
			Value: fmt.Sprintf("%v", *m.LoTypeIdAttr)})
	}
	if m.LoCatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "loCatId"},
			Value: fmt.Sprintf("%v", *m.LoCatIdAttr)})
	}
	if m.QsTypeIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "qsTypeId"},
			Value: fmt.Sprintf("%v", *m.QsTypeIdAttr)})
	}
	if m.QsCatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "qsCatId"},
			Value: fmt.Sprintf("%v", *m.QsCatIdAttr)})
	}
	if m.CsTypeIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "csTypeId"},
			Value: fmt.Sprintf("%v", *m.CsTypeIdAttr)})
	}
	if m.CsCatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "csCatId"},
			Value: fmt.Sprintf("%v", *m.CsCatIdAttr)})
	}
	if m.Coherent3DOffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coherent3DOff"},
			Value: fmt.Sprintf("%v", *m.Coherent3DOffAttr)})
	}
	if m.PhldrTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "phldrT"},
			Value: fmt.Sprintf("%v", *m.PhldrTAttr)})
	}
	if m.PhldrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "phldr"},
			Value: fmt.Sprintf("%v", *m.PhldrAttr)})
	}
	if m.CustAngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custAng"},
			Value: fmt.Sprintf("%v", *m.CustAngAttr)})
	}
	if m.CustFlipVertAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custFlipVert"},
			Value: fmt.Sprintf("%v", *m.CustFlipVertAttr)})
	}
	if m.CustFlipHorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custFlipHor"},
			Value: fmt.Sprintf("%v", *m.CustFlipHorAttr)})
	}
	if m.CustSzXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custSzX"},
			Value: fmt.Sprintf("%v", *m.CustSzXAttr)})
	}
	if m.CustSzYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custSzY"},
			Value: fmt.Sprintf("%v", *m.CustSzYAttr)})
	}
	if m.CustScaleXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custScaleX"},
			Value: fmt.Sprintf("%v", *m.CustScaleXAttr)})
	}
	if m.CustScaleYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custScaleY"},
			Value: fmt.Sprintf("%v", *m.CustScaleYAttr)})
	}
	if m.CustTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custT"},
			Value: fmt.Sprintf("%v", *m.CustTAttr)})
	}
	if m.CustLinFactXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custLinFactX"},
			Value: fmt.Sprintf("%v", *m.CustLinFactXAttr)})
	}
	if m.CustLinFactYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custLinFactY"},
			Value: fmt.Sprintf("%v", *m.CustLinFactYAttr)})
	}
	if m.CustLinFactNeighborXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custLinFactNeighborX"},
			Value: fmt.Sprintf("%v", *m.CustLinFactNeighborXAttr)})
	}
	if m.CustLinFactNeighborYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custLinFactNeighborY"},
			Value: fmt.Sprintf("%v", *m.CustLinFactNeighborYAttr)})
	}
	if m.CustRadScaleRadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custRadScaleRad"},
			Value: fmt.Sprintf("%v", *m.CustRadScaleRadAttr)})
	}
	if m.CustRadScaleIncAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "custRadScaleInc"},
			Value: fmt.Sprintf("%v", *m.CustRadScaleIncAttr)})
	}
	e.EncodeToken(start)
	if m.PresLayoutVars != nil {
		sepresLayoutVars := xml.StartElement{Name: xml.Name{Local: "presLayoutVars"}}
		e.EncodeElement(m.PresLayoutVars, sepresLayoutVars)
	}
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ElemPropSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "presAssocID" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.PresAssocIDAttr = &parsed
		}
		if attr.Name.Local == "presName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PresNameAttr = &parsed
		}
		if attr.Name.Local == "presStyleLbl" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PresStyleLblAttr = &parsed
		}
		if attr.Name.Local == "presStyleIdx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.PresStyleIdxAttr = &pt
		}
		if attr.Name.Local == "presStyleCnt" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.PresStyleCntAttr = &pt
		}
		if attr.Name.Local == "loTypeId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LoTypeIdAttr = &parsed
		}
		if attr.Name.Local == "loCatId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LoCatIdAttr = &parsed
		}
		if attr.Name.Local == "qsTypeId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.QsTypeIdAttr = &parsed
		}
		if attr.Name.Local == "qsCatId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.QsCatIdAttr = &parsed
		}
		if attr.Name.Local == "csTypeId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsTypeIdAttr = &parsed
		}
		if attr.Name.Local == "csCatId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsCatIdAttr = &parsed
		}
		if attr.Name.Local == "coherent3DOff" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Coherent3DOffAttr = &parsed
		}
		if attr.Name.Local == "phldrT" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PhldrTAttr = &parsed
		}
		if attr.Name.Local == "phldr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PhldrAttr = &parsed
		}
		if attr.Name.Local == "custAng" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.CustAngAttr = &pt
		}
		if attr.Name.Local == "custFlipVert" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustFlipVertAttr = &parsed
		}
		if attr.Name.Local == "custFlipHor" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustFlipHorAttr = &parsed
		}
		if attr.Name.Local == "custSzX" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.CustSzXAttr = &pt
		}
		if attr.Name.Local == "custSzY" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.CustSzYAttr = &pt
		}
		if attr.Name.Local == "custScaleX" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustScaleXAttr = &parsed
		}
		if attr.Name.Local == "custScaleY" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustScaleYAttr = &parsed
		}
		if attr.Name.Local == "custT" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustTAttr = &parsed
		}
		if attr.Name.Local == "custLinFactX" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustLinFactXAttr = &parsed
		}
		if attr.Name.Local == "custLinFactY" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustLinFactYAttr = &parsed
		}
		if attr.Name.Local == "custLinFactNeighborX" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustLinFactNeighborXAttr = &parsed
		}
		if attr.Name.Local == "custLinFactNeighborY" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustLinFactNeighborYAttr = &parsed
		}
		if attr.Name.Local == "custRadScaleRad" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustRadScaleRadAttr = &parsed
		}
		if attr.Name.Local == "custRadScaleInc" {
			parsed, err := ParseUnionST_PrSetCustVal(attr.Value)
			if err != nil {
				return err
			}
			m.CustRadScaleIncAttr = &parsed
		}
	}
lCT_ElemPropSet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "presLayoutVars":
				m.PresLayoutVars = NewCT_LayoutVariablePropertySet()
				if err := d.DecodeElement(m.PresLayoutVars, &el); err != nil {
					return err
				}
			case "style":
				m.Style = drawingml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ElemPropSet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ElemPropSet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ElemPropSet and its children
func (m *CT_ElemPropSet) Validate() error {
	return m.ValidateWithPath("CT_ElemPropSet")
}

// ValidateWithPath validates the CT_ElemPropSet and its children, prefixing error messages with path
func (m *CT_ElemPropSet) ValidateWithPath(path string) error {
	if m.PresAssocIDAttr != nil {
		if err := m.PresAssocIDAttr.ValidateWithPath(path + "/PresAssocIDAttr"); err != nil {
			return err
		}
	}
	if m.CustScaleXAttr != nil {
		if err := m.CustScaleXAttr.ValidateWithPath(path + "/CustScaleXAttr"); err != nil {
			return err
		}
	}
	if m.CustScaleYAttr != nil {
		if err := m.CustScaleYAttr.ValidateWithPath(path + "/CustScaleYAttr"); err != nil {
			return err
		}
	}
	if m.CustLinFactXAttr != nil {
		if err := m.CustLinFactXAttr.ValidateWithPath(path + "/CustLinFactXAttr"); err != nil {
			return err
		}
	}
	if m.CustLinFactYAttr != nil {
		if err := m.CustLinFactYAttr.ValidateWithPath(path + "/CustLinFactYAttr"); err != nil {
			return err
		}
	}
	if m.CustLinFactNeighborXAttr != nil {
		if err := m.CustLinFactNeighborXAttr.ValidateWithPath(path + "/CustLinFactNeighborXAttr"); err != nil {
			return err
		}
	}
	if m.CustLinFactNeighborYAttr != nil {
		if err := m.CustLinFactNeighborYAttr.ValidateWithPath(path + "/CustLinFactNeighborYAttr"); err != nil {
			return err
		}
	}
	if m.CustRadScaleRadAttr != nil {
		if err := m.CustRadScaleRadAttr.ValidateWithPath(path + "/CustRadScaleRadAttr"); err != nil {
			return err
		}
	}
	if m.CustRadScaleIncAttr != nil {
		if err := m.CustRadScaleIncAttr.ValidateWithPath(path + "/CustRadScaleIncAttr"); err != nil {
			return err
		}
	}
	if m.PresLayoutVars != nil {
		if err := m.PresLayoutVars.ValidateWithPath(path + "/PresLayoutVars"); err != nil {
			return err
		}
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	return nil
}
