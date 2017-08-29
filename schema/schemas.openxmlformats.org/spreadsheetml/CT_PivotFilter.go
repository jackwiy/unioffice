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

type CT_PivotFilter struct {
	// Field Index
	FldAttr uint32
	// Member Property Field Id
	MpFldAttr *uint32
	// Pivot Filter Type
	TypeAttr ST_PivotFilterType
	// Evaluation Order
	EvalOrderAttr *int32
	// Pivot Filter Id
	IdAttr uint32
	// Measure Index
	IMeasureHierAttr *uint32
	// Measure Field Index
	IMeasureFldAttr *uint32
	// Pivot Filter Name
	NameAttr *string
	// Pivot Filter Description
	DescriptionAttr *string
	// Label Pivot
	StringValue1Attr *string
	// Label Pivot Filter String Value 2
	StringValue2Attr *string
	// Auto Filter
	AutoFilter *CT_AutoFilter
	ExtLst     *CT_ExtensionList
}

func NewCT_PivotFilter() *CT_PivotFilter {
	ret := &CT_PivotFilter{}
	ret.AutoFilter = NewCT_AutoFilter()
	return ret
}
func (m *CT_PivotFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fld"},
		Value: fmt.Sprintf("%v", m.FldAttr)})
	if m.MpFldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mpFld"},
			Value: fmt.Sprintf("%v", *m.MpFldAttr)})
	}
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.EvalOrderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "evalOrder"},
			Value: fmt.Sprintf("%v", *m.EvalOrderAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.IMeasureHierAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iMeasureHier"},
			Value: fmt.Sprintf("%v", *m.IMeasureHierAttr)})
	}
	if m.IMeasureFldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iMeasureFld"},
			Value: fmt.Sprintf("%v", *m.IMeasureFldAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.DescriptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "description"},
			Value: fmt.Sprintf("%v", *m.DescriptionAttr)})
	}
	if m.StringValue1Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stringValue1"},
			Value: fmt.Sprintf("%v", *m.StringValue1Attr)})
	}
	if m.StringValue2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stringValue2"},
			Value: fmt.Sprintf("%v", *m.StringValue2Attr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	seautoFilter := xml.StartElement{Name: xml.Name{Local: "x:autoFilter"}}
	e.EncodeElement(m.AutoFilter, seautoFilter)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AutoFilter = NewCT_AutoFilter()
	for _, attr := range start.Attr {
		if attr.Name.Local == "fld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FldAttr = uint32(parsed)
		}
		if attr.Name.Local == "mpFld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.MpFldAttr = &pt
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "evalOrder" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.EvalOrderAttr = &pt
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "iMeasureHier" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.IMeasureHierAttr = &pt
		}
		if attr.Name.Local == "iMeasureFld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.IMeasureFldAttr = &pt
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "description" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DescriptionAttr = &parsed
		}
		if attr.Name.Local == "stringValue1" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StringValue1Attr = &parsed
		}
		if attr.Name.Local == "stringValue2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StringValue2Attr = &parsed
		}
	}
lCT_PivotFilter:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "autoFilter":
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotFilter
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotFilter) Validate() error {
	return m.ValidateWithPath("CT_PivotFilter")
}
func (m *CT_PivotFilter) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_PivotFilterTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.AutoFilter.ValidateWithPath(path + "/AutoFilter"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}