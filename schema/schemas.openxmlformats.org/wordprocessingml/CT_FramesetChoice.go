// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_FramesetChoice struct {
	Frameset []*CT_Frameset
	Frame    []*CT_Frame
}

func NewCT_FramesetChoice() *CT_FramesetChoice {
	ret := &CT_FramesetChoice{}
	return ret
}
func (m *CT_FramesetChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Frameset != nil {
		seframeset := xml.StartElement{Name: xml.Name{Local: "w:frameset"}}
		e.EncodeElement(m.Frameset, seframeset)
	}
	if m.Frame != nil {
		seframe := xml.StartElement{Name: xml.Name{Local: "w:frame"}}
		e.EncodeElement(m.Frame, seframe)
	}
	return nil
}
func (m *CT_FramesetChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FramesetChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "frameset":
				tmp := NewCT_Frameset()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Frameset = append(m.Frameset, tmp)
			case "frame":
				tmp := NewCT_Frame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Frame = append(m.Frame, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FramesetChoice
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FramesetChoice) Validate() error {
	return m.ValidateWithPath("CT_FramesetChoice")
}
func (m *CT_FramesetChoice) ValidateWithPath(path string) error {
	for i, v := range m.Frameset {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Frameset[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Frame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Frame[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}