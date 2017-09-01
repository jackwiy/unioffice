// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_BlipFillProperties struct {
	DpiAttr          *uint32
	RotWithShapeAttr *bool
	Blip             *CT_Blip
	SrcRect          *CT_RelativeRect
	Tile             *CT_TileInfoProperties
	Stretch          *CT_StretchInfoProperties
}

func NewCT_BlipFillProperties() *CT_BlipFillProperties {
	ret := &CT_BlipFillProperties{}
	return ret
}

func (m *CT_BlipFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DpiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dpi"},
			Value: fmt.Sprintf("%v", *m.DpiAttr)})
	}
	if m.RotWithShapeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rotWithShape"},
			Value: fmt.Sprintf("%v", *m.RotWithShapeAttr)})
	}
	e.EncodeToken(start)
	if m.Blip != nil {
		seblip := xml.StartElement{Name: xml.Name{Local: "a:blip"}}
		e.EncodeElement(m.Blip, seblip)
	}
	if m.SrcRect != nil {
		sesrcRect := xml.StartElement{Name: xml.Name{Local: "a:srcRect"}}
		e.EncodeElement(m.SrcRect, sesrcRect)
	}
	if m.Tile != nil {
		setile := xml.StartElement{Name: xml.Name{Local: "a:tile"}}
		e.EncodeElement(m.Tile, setile)
	}
	if m.Stretch != nil {
		sestretch := xml.StartElement{Name: xml.Name{Local: "a:stretch"}}
		e.EncodeElement(m.Stretch, sestretch)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BlipFillProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dpi" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DpiAttr = &pt
		}
		if attr.Name.Local == "rotWithShape" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RotWithShapeAttr = &parsed
		}
	}
lCT_BlipFillProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "blip":
				m.Blip = NewCT_Blip()
				if err := d.DecodeElement(m.Blip, &el); err != nil {
					return err
				}
			case "srcRect":
				m.SrcRect = NewCT_RelativeRect()
				if err := d.DecodeElement(m.SrcRect, &el); err != nil {
					return err
				}
			case "tile":
				m.Tile = NewCT_TileInfoProperties()
				if err := d.DecodeElement(m.Tile, &el); err != nil {
					return err
				}
			case "stretch":
				m.Stretch = NewCT_StretchInfoProperties()
				if err := d.DecodeElement(m.Stretch, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_BlipFillProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BlipFillProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BlipFillProperties and its children
func (m *CT_BlipFillProperties) Validate() error {
	return m.ValidateWithPath("CT_BlipFillProperties")
}

// ValidateWithPath validates the CT_BlipFillProperties and its children, prefixing error messages with path
func (m *CT_BlipFillProperties) ValidateWithPath(path string) error {
	if m.Blip != nil {
		if err := m.Blip.ValidateWithPath(path + "/Blip"); err != nil {
			return err
		}
	}
	if m.SrcRect != nil {
		if err := m.SrcRect.ValidateWithPath(path + "/SrcRect"); err != nil {
			return err
		}
	}
	if m.Tile != nil {
		if err := m.Tile.ValidateWithPath(path + "/Tile"); err != nil {
			return err
		}
	}
	if m.Stretch != nil {
		if err := m.Stretch.ValidateWithPath(path + "/Stretch"); err != nil {
			return err
		}
	}
	return nil
}
