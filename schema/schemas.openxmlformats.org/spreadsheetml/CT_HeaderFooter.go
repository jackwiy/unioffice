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

	"baliance.com/gooxml"
)

type CT_HeaderFooter struct {
	// Different Odd Even Header Footer
	DifferentOddEvenAttr *bool
	// Different First Page
	DifferentFirstAttr *bool
	// Scale Header & Footer With Document
	ScaleWithDocAttr *bool
	// Align Margins
	AlignWithMarginsAttr *bool
	// Odd Header
	OddHeader *string
	// Odd Page Footer
	OddFooter *string
	// Even Page Header
	EvenHeader *string
	// Even Page Footer
	EvenFooter *string
	// First Page Header
	FirstHeader *string
	// First Page Footer
	FirstFooter *string
}

func NewCT_HeaderFooter() *CT_HeaderFooter {
	ret := &CT_HeaderFooter{}
	return ret
}

func (m *CT_HeaderFooter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DifferentOddEvenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "differentOddEven"},
			Value: fmt.Sprintf("%v", *m.DifferentOddEvenAttr)})
	}
	if m.DifferentFirstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "differentFirst"},
			Value: fmt.Sprintf("%v", *m.DifferentFirstAttr)})
	}
	if m.ScaleWithDocAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scaleWithDoc"},
			Value: fmt.Sprintf("%v", *m.ScaleWithDocAttr)})
	}
	if m.AlignWithMarginsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alignWithMargins"},
			Value: fmt.Sprintf("%v", *m.AlignWithMarginsAttr)})
	}
	e.EncodeToken(start)
	if m.OddHeader != nil {
		seoddHeader := xml.StartElement{Name: xml.Name{Local: "x:oddHeader"}}
		gooxml.AddPreserveSpaceAttr(&seoddHeader, *m.OddHeader)
		e.EncodeElement(m.OddHeader, seoddHeader)
	}
	if m.OddFooter != nil {
		seoddFooter := xml.StartElement{Name: xml.Name{Local: "x:oddFooter"}}
		gooxml.AddPreserveSpaceAttr(&seoddFooter, *m.OddFooter)
		e.EncodeElement(m.OddFooter, seoddFooter)
	}
	if m.EvenHeader != nil {
		seevenHeader := xml.StartElement{Name: xml.Name{Local: "x:evenHeader"}}
		gooxml.AddPreserveSpaceAttr(&seevenHeader, *m.EvenHeader)
		e.EncodeElement(m.EvenHeader, seevenHeader)
	}
	if m.EvenFooter != nil {
		seevenFooter := xml.StartElement{Name: xml.Name{Local: "x:evenFooter"}}
		gooxml.AddPreserveSpaceAttr(&seevenFooter, *m.EvenFooter)
		e.EncodeElement(m.EvenFooter, seevenFooter)
	}
	if m.FirstHeader != nil {
		sefirstHeader := xml.StartElement{Name: xml.Name{Local: "x:firstHeader"}}
		gooxml.AddPreserveSpaceAttr(&sefirstHeader, *m.FirstHeader)
		e.EncodeElement(m.FirstHeader, sefirstHeader)
	}
	if m.FirstFooter != nil {
		sefirstFooter := xml.StartElement{Name: xml.Name{Local: "x:firstFooter"}}
		gooxml.AddPreserveSpaceAttr(&sefirstFooter, *m.FirstFooter)
		e.EncodeElement(m.FirstFooter, sefirstFooter)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HeaderFooter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "differentOddEven" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DifferentOddEvenAttr = &parsed
		}
		if attr.Name.Local == "differentFirst" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DifferentFirstAttr = &parsed
		}
		if attr.Name.Local == "scaleWithDoc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ScaleWithDocAttr = &parsed
		}
		if attr.Name.Local == "alignWithMargins" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AlignWithMarginsAttr = &parsed
		}
	}
lCT_HeaderFooter:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "oddHeader":
				m.OddHeader = new(string)
				if err := d.DecodeElement(m.OddHeader, &el); err != nil {
					return err
				}
			case "oddFooter":
				m.OddFooter = new(string)
				if err := d.DecodeElement(m.OddFooter, &el); err != nil {
					return err
				}
			case "evenHeader":
				m.EvenHeader = new(string)
				if err := d.DecodeElement(m.EvenHeader, &el); err != nil {
					return err
				}
			case "evenFooter":
				m.EvenFooter = new(string)
				if err := d.DecodeElement(m.EvenFooter, &el); err != nil {
					return err
				}
			case "firstHeader":
				m.FirstHeader = new(string)
				if err := d.DecodeElement(m.FirstHeader, &el); err != nil {
					return err
				}
			case "firstFooter":
				m.FirstFooter = new(string)
				if err := d.DecodeElement(m.FirstFooter, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_HeaderFooter %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_HeaderFooter
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_HeaderFooter and its children
func (m *CT_HeaderFooter) Validate() error {
	return m.ValidateWithPath("CT_HeaderFooter")
}

// ValidateWithPath validates the CT_HeaderFooter and its children, prefixing error messages with path
func (m *CT_HeaderFooter) ValidateWithPath(path string) error {
	return nil
}
