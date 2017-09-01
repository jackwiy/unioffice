// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_TcPrInner struct {
	// Table Cell Conditional Formatting
	CnfStyle *CT_Cnf
	// Preferred Table Cell Width
	TcW *CT_TblWidth
	// Grid Columns Spanned by Current Table Cell
	GridSpan *CT_DecimalNumber
	// Horizontally Merged Cell
	HMerge *CT_HMerge
	// Vertically Merged Cell
	VMerge *CT_VMerge
	// Table Cell Borders
	TcBorders *CT_TcBorders
	// Table Cell Shading
	Shd *CT_Shd
	// Don't Wrap Cell Content
	NoWrap *CT_OnOff
	// Single Table Cell Margins
	TcMar *CT_TcMar
	// Table Cell Text Flow Direction
	TextDirection *CT_TextDirection
	// Fit Text Within Cell
	TcFitText *CT_OnOff
	// Table Cell Vertical Alignment
	VAlign *CT_VerticalJc
	// Ignore End Of Cell Marker In Row Height Calculation
	HideMark *CT_OnOff
	// Header Cells Associated With Table Cell
	Headers *CT_Headers
	// Table Cell Insertion
	CellIns *CT_TrackChange
	// Table Cell Deletion
	CellDel *CT_TrackChange
	// Vertically Merged/Split Table Cells
	CellMerge *CT_CellMergeTrackChange
}

func NewCT_TcPrInner() *CT_TcPrInner {
	ret := &CT_TcPrInner{}
	return ret
}

func (m *CT_TcPrInner) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.CnfStyle != nil {
		secnfStyle := xml.StartElement{Name: xml.Name{Local: "w:cnfStyle"}}
		e.EncodeElement(m.CnfStyle, secnfStyle)
	}
	if m.TcW != nil {
		setcW := xml.StartElement{Name: xml.Name{Local: "w:tcW"}}
		e.EncodeElement(m.TcW, setcW)
	}
	if m.GridSpan != nil {
		segridSpan := xml.StartElement{Name: xml.Name{Local: "w:gridSpan"}}
		e.EncodeElement(m.GridSpan, segridSpan)
	}
	if m.HMerge != nil {
		sehMerge := xml.StartElement{Name: xml.Name{Local: "w:hMerge"}}
		e.EncodeElement(m.HMerge, sehMerge)
	}
	if m.VMerge != nil {
		sevMerge := xml.StartElement{Name: xml.Name{Local: "w:vMerge"}}
		e.EncodeElement(m.VMerge, sevMerge)
	}
	if m.TcBorders != nil {
		setcBorders := xml.StartElement{Name: xml.Name{Local: "w:tcBorders"}}
		e.EncodeElement(m.TcBorders, setcBorders)
	}
	if m.Shd != nil {
		seshd := xml.StartElement{Name: xml.Name{Local: "w:shd"}}
		e.EncodeElement(m.Shd, seshd)
	}
	if m.NoWrap != nil {
		senoWrap := xml.StartElement{Name: xml.Name{Local: "w:noWrap"}}
		e.EncodeElement(m.NoWrap, senoWrap)
	}
	if m.TcMar != nil {
		setcMar := xml.StartElement{Name: xml.Name{Local: "w:tcMar"}}
		e.EncodeElement(m.TcMar, setcMar)
	}
	if m.TextDirection != nil {
		setextDirection := xml.StartElement{Name: xml.Name{Local: "w:textDirection"}}
		e.EncodeElement(m.TextDirection, setextDirection)
	}
	if m.TcFitText != nil {
		setcFitText := xml.StartElement{Name: xml.Name{Local: "w:tcFitText"}}
		e.EncodeElement(m.TcFitText, setcFitText)
	}
	if m.VAlign != nil {
		sevAlign := xml.StartElement{Name: xml.Name{Local: "w:vAlign"}}
		e.EncodeElement(m.VAlign, sevAlign)
	}
	if m.HideMark != nil {
		sehideMark := xml.StartElement{Name: xml.Name{Local: "w:hideMark"}}
		e.EncodeElement(m.HideMark, sehideMark)
	}
	if m.Headers != nil {
		seheaders := xml.StartElement{Name: xml.Name{Local: "w:headers"}}
		e.EncodeElement(m.Headers, seheaders)
	}
	if m.CellIns != nil {
		secellIns := xml.StartElement{Name: xml.Name{Local: "w:cellIns"}}
		e.EncodeElement(m.CellIns, secellIns)
	}
	if m.CellDel != nil {
		secellDel := xml.StartElement{Name: xml.Name{Local: "w:cellDel"}}
		e.EncodeElement(m.CellDel, secellDel)
	}
	if m.CellMerge != nil {
		secellMerge := xml.StartElement{Name: xml.Name{Local: "w:cellMerge"}}
		e.EncodeElement(m.CellMerge, secellMerge)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TcPrInner) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TcPrInner:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cnfStyle":
				m.CnfStyle = NewCT_Cnf()
				if err := d.DecodeElement(m.CnfStyle, &el); err != nil {
					return err
				}
			case "tcW":
				m.TcW = NewCT_TblWidth()
				if err := d.DecodeElement(m.TcW, &el); err != nil {
					return err
				}
			case "gridSpan":
				m.GridSpan = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.GridSpan, &el); err != nil {
					return err
				}
			case "hMerge":
				m.HMerge = NewCT_HMerge()
				if err := d.DecodeElement(m.HMerge, &el); err != nil {
					return err
				}
			case "vMerge":
				m.VMerge = NewCT_VMerge()
				if err := d.DecodeElement(m.VMerge, &el); err != nil {
					return err
				}
			case "tcBorders":
				m.TcBorders = NewCT_TcBorders()
				if err := d.DecodeElement(m.TcBorders, &el); err != nil {
					return err
				}
			case "shd":
				m.Shd = NewCT_Shd()
				if err := d.DecodeElement(m.Shd, &el); err != nil {
					return err
				}
			case "noWrap":
				m.NoWrap = NewCT_OnOff()
				if err := d.DecodeElement(m.NoWrap, &el); err != nil {
					return err
				}
			case "tcMar":
				m.TcMar = NewCT_TcMar()
				if err := d.DecodeElement(m.TcMar, &el); err != nil {
					return err
				}
			case "textDirection":
				m.TextDirection = NewCT_TextDirection()
				if err := d.DecodeElement(m.TextDirection, &el); err != nil {
					return err
				}
			case "tcFitText":
				m.TcFitText = NewCT_OnOff()
				if err := d.DecodeElement(m.TcFitText, &el); err != nil {
					return err
				}
			case "vAlign":
				m.VAlign = NewCT_VerticalJc()
				if err := d.DecodeElement(m.VAlign, &el); err != nil {
					return err
				}
			case "hideMark":
				m.HideMark = NewCT_OnOff()
				if err := d.DecodeElement(m.HideMark, &el); err != nil {
					return err
				}
			case "headers":
				m.Headers = NewCT_Headers()
				if err := d.DecodeElement(m.Headers, &el); err != nil {
					return err
				}
			case "cellIns":
				m.CellIns = NewCT_TrackChange()
				if err := d.DecodeElement(m.CellIns, &el); err != nil {
					return err
				}
			case "cellDel":
				m.CellDel = NewCT_TrackChange()
				if err := d.DecodeElement(m.CellDel, &el); err != nil {
					return err
				}
			case "cellMerge":
				m.CellMerge = NewCT_CellMergeTrackChange()
				if err := d.DecodeElement(m.CellMerge, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TcPrInner %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TcPrInner
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TcPrInner and its children
func (m *CT_TcPrInner) Validate() error {
	return m.ValidateWithPath("CT_TcPrInner")
}

// ValidateWithPath validates the CT_TcPrInner and its children, prefixing error messages with path
func (m *CT_TcPrInner) ValidateWithPath(path string) error {
	if m.CnfStyle != nil {
		if err := m.CnfStyle.ValidateWithPath(path + "/CnfStyle"); err != nil {
			return err
		}
	}
	if m.TcW != nil {
		if err := m.TcW.ValidateWithPath(path + "/TcW"); err != nil {
			return err
		}
	}
	if m.GridSpan != nil {
		if err := m.GridSpan.ValidateWithPath(path + "/GridSpan"); err != nil {
			return err
		}
	}
	if m.HMerge != nil {
		if err := m.HMerge.ValidateWithPath(path + "/HMerge"); err != nil {
			return err
		}
	}
	if m.VMerge != nil {
		if err := m.VMerge.ValidateWithPath(path + "/VMerge"); err != nil {
			return err
		}
	}
	if m.TcBorders != nil {
		if err := m.TcBorders.ValidateWithPath(path + "/TcBorders"); err != nil {
			return err
		}
	}
	if m.Shd != nil {
		if err := m.Shd.ValidateWithPath(path + "/Shd"); err != nil {
			return err
		}
	}
	if m.NoWrap != nil {
		if err := m.NoWrap.ValidateWithPath(path + "/NoWrap"); err != nil {
			return err
		}
	}
	if m.TcMar != nil {
		if err := m.TcMar.ValidateWithPath(path + "/TcMar"); err != nil {
			return err
		}
	}
	if m.TextDirection != nil {
		if err := m.TextDirection.ValidateWithPath(path + "/TextDirection"); err != nil {
			return err
		}
	}
	if m.TcFitText != nil {
		if err := m.TcFitText.ValidateWithPath(path + "/TcFitText"); err != nil {
			return err
		}
	}
	if m.VAlign != nil {
		if err := m.VAlign.ValidateWithPath(path + "/VAlign"); err != nil {
			return err
		}
	}
	if m.HideMark != nil {
		if err := m.HideMark.ValidateWithPath(path + "/HideMark"); err != nil {
			return err
		}
	}
	if m.Headers != nil {
		if err := m.Headers.ValidateWithPath(path + "/Headers"); err != nil {
			return err
		}
	}
	if m.CellIns != nil {
		if err := m.CellIns.ValidateWithPath(path + "/CellIns"); err != nil {
			return err
		}
	}
	if m.CellDel != nil {
		if err := m.CellDel.ValidateWithPath(path + "/CellDel"); err != nil {
			return err
		}
	}
	if m.CellMerge != nil {
		if err := m.CellMerge.ValidateWithPath(path + "/CellMerge"); err != nil {
			return err
		}
	}
	return nil
}
