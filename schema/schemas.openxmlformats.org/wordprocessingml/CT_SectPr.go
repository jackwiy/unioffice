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

type CT_SectPr struct {
	EG_HdrFtrReferences []*EG_HdrFtrReferences
	// Section-Wide Footnote Properties
	FootnotePr *CT_FtnProps
	// Section-Wide Endnote Properties
	EndnotePr *CT_EdnProps
	// Section Type
	Type *CT_SectType
	// Page Size
	PgSz *CT_PageSz
	// Page Margins
	PgMar *CT_PageMar
	// Paper Source Information
	PaperSrc *CT_PaperSource
	// Page Borders
	PgBorders *CT_PageBorders
	// Line Numbering Settings
	LnNumType *CT_LineNumber
	// Page Numbering Settings
	PgNumType *CT_PageNumber
	// Column Definitions
	Cols *CT_Columns
	// Only Allow Editing of Form Fields
	FormProt *CT_OnOff
	// Vertical Text Alignment on Page
	VAlign *CT_VerticalJc
	// Suppress Endnotes In Document
	NoEndnote *CT_OnOff
	// Different First Page Headers and Footers
	TitlePg *CT_OnOff
	// Text Flow Direction
	TextDirection *CT_TextDirection
	// Right to Left Section Layout
	Bidi *CT_OnOff
	// Gutter on Right Side of Page
	RtlGutter *CT_OnOff
	// Document Grid
	DocGrid *CT_DocGrid
	// Reference to Printer Settings Data
	PrinterSettings *CT_Rel
	// Revision Information for Section Properties
	SectPrChange *CT_SectPrChange
	RsidRPrAttr  *string
	RsidDelAttr  *string
	RsidRAttr    *string
	RsidSectAttr *string
}

func NewCT_SectPr() *CT_SectPr {
	ret := &CT_SectPr{}
	return ret
}

func (m *CT_SectPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RsidRPrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"},
			Value: fmt.Sprintf("%v", *m.RsidRPrAttr)})
	}
	if m.RsidDelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"},
			Value: fmt.Sprintf("%v", *m.RsidDelAttr)})
	}
	if m.RsidRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"},
			Value: fmt.Sprintf("%v", *m.RsidRAttr)})
	}
	if m.RsidSectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidSect"},
			Value: fmt.Sprintf("%v", *m.RsidSectAttr)})
	}
	e.EncodeToken(start)
	if m.EG_HdrFtrReferences != nil {
		for _, c := range m.EG_HdrFtrReferences {
			c.MarshalXML(e, start)
		}
	}
	if m.FootnotePr != nil {
		sefootnotePr := xml.StartElement{Name: xml.Name{Local: "w:footnotePr"}}
		e.EncodeElement(m.FootnotePr, sefootnotePr)
	}
	if m.EndnotePr != nil {
		seendnotePr := xml.StartElement{Name: xml.Name{Local: "w:endnotePr"}}
		e.EncodeElement(m.EndnotePr, seendnotePr)
	}
	if m.Type != nil {
		setype := xml.StartElement{Name: xml.Name{Local: "w:type"}}
		e.EncodeElement(m.Type, setype)
	}
	if m.PgSz != nil {
		sepgSz := xml.StartElement{Name: xml.Name{Local: "w:pgSz"}}
		e.EncodeElement(m.PgSz, sepgSz)
	}
	if m.PgMar != nil {
		sepgMar := xml.StartElement{Name: xml.Name{Local: "w:pgMar"}}
		e.EncodeElement(m.PgMar, sepgMar)
	}
	if m.PaperSrc != nil {
		sepaperSrc := xml.StartElement{Name: xml.Name{Local: "w:paperSrc"}}
		e.EncodeElement(m.PaperSrc, sepaperSrc)
	}
	if m.PgBorders != nil {
		sepgBorders := xml.StartElement{Name: xml.Name{Local: "w:pgBorders"}}
		e.EncodeElement(m.PgBorders, sepgBorders)
	}
	if m.LnNumType != nil {
		selnNumType := xml.StartElement{Name: xml.Name{Local: "w:lnNumType"}}
		e.EncodeElement(m.LnNumType, selnNumType)
	}
	if m.PgNumType != nil {
		sepgNumType := xml.StartElement{Name: xml.Name{Local: "w:pgNumType"}}
		e.EncodeElement(m.PgNumType, sepgNumType)
	}
	if m.Cols != nil {
		secols := xml.StartElement{Name: xml.Name{Local: "w:cols"}}
		e.EncodeElement(m.Cols, secols)
	}
	if m.FormProt != nil {
		seformProt := xml.StartElement{Name: xml.Name{Local: "w:formProt"}}
		e.EncodeElement(m.FormProt, seformProt)
	}
	if m.VAlign != nil {
		sevAlign := xml.StartElement{Name: xml.Name{Local: "w:vAlign"}}
		e.EncodeElement(m.VAlign, sevAlign)
	}
	if m.NoEndnote != nil {
		senoEndnote := xml.StartElement{Name: xml.Name{Local: "w:noEndnote"}}
		e.EncodeElement(m.NoEndnote, senoEndnote)
	}
	if m.TitlePg != nil {
		setitlePg := xml.StartElement{Name: xml.Name{Local: "w:titlePg"}}
		e.EncodeElement(m.TitlePg, setitlePg)
	}
	if m.TextDirection != nil {
		setextDirection := xml.StartElement{Name: xml.Name{Local: "w:textDirection"}}
		e.EncodeElement(m.TextDirection, setextDirection)
	}
	if m.Bidi != nil {
		sebidi := xml.StartElement{Name: xml.Name{Local: "w:bidi"}}
		e.EncodeElement(m.Bidi, sebidi)
	}
	if m.RtlGutter != nil {
		sertlGutter := xml.StartElement{Name: xml.Name{Local: "w:rtlGutter"}}
		e.EncodeElement(m.RtlGutter, sertlGutter)
	}
	if m.DocGrid != nil {
		sedocGrid := xml.StartElement{Name: xml.Name{Local: "w:docGrid"}}
		e.EncodeElement(m.DocGrid, sedocGrid)
	}
	if m.PrinterSettings != nil {
		seprinterSettings := xml.StartElement{Name: xml.Name{Local: "w:printerSettings"}}
		e.EncodeElement(m.PrinterSettings, seprinterSettings)
	}
	if m.SectPrChange != nil {
		sesectPrChange := xml.StartElement{Name: xml.Name{Local: "w:sectPrChange"}}
		e.EncodeElement(m.SectPrChange, sesectPrChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SectPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rsidRPr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRPrAttr = &parsed
		}
		if attr.Name.Local == "rsidDel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidDelAttr = &parsed
		}
		if attr.Name.Local == "rsidR" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRAttr = &parsed
		}
		if attr.Name.Local == "rsidSect" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidSectAttr = &parsed
		}
	}
lCT_SectPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "headerReference":
				tmphdrftrreferences := NewEG_HdrFtrReferences()
				tmphdrftrreferences.HeaderReference = NewCT_HdrFtrRef()
				if err := d.DecodeElement(tmphdrftrreferences.HeaderReference, &el); err != nil {
					return err
				}
				m.EG_HdrFtrReferences = append(m.EG_HdrFtrReferences, tmphdrftrreferences)
			case "footerReference":
				tmphdrftrreferences := NewEG_HdrFtrReferences()
				tmphdrftrreferences.FooterReference = NewCT_HdrFtrRef()
				if err := d.DecodeElement(tmphdrftrreferences.FooterReference, &el); err != nil {
					return err
				}
				m.EG_HdrFtrReferences = append(m.EG_HdrFtrReferences, tmphdrftrreferences)
			case "footnotePr":
				m.FootnotePr = NewCT_FtnProps()
				if err := d.DecodeElement(m.FootnotePr, &el); err != nil {
					return err
				}
			case "endnotePr":
				m.EndnotePr = NewCT_EdnProps()
				if err := d.DecodeElement(m.EndnotePr, &el); err != nil {
					return err
				}
			case "type":
				m.Type = NewCT_SectType()
				if err := d.DecodeElement(m.Type, &el); err != nil {
					return err
				}
			case "pgSz":
				m.PgSz = NewCT_PageSz()
				if err := d.DecodeElement(m.PgSz, &el); err != nil {
					return err
				}
			case "pgMar":
				m.PgMar = NewCT_PageMar()
				if err := d.DecodeElement(m.PgMar, &el); err != nil {
					return err
				}
			case "paperSrc":
				m.PaperSrc = NewCT_PaperSource()
				if err := d.DecodeElement(m.PaperSrc, &el); err != nil {
					return err
				}
			case "pgBorders":
				m.PgBorders = NewCT_PageBorders()
				if err := d.DecodeElement(m.PgBorders, &el); err != nil {
					return err
				}
			case "lnNumType":
				m.LnNumType = NewCT_LineNumber()
				if err := d.DecodeElement(m.LnNumType, &el); err != nil {
					return err
				}
			case "pgNumType":
				m.PgNumType = NewCT_PageNumber()
				if err := d.DecodeElement(m.PgNumType, &el); err != nil {
					return err
				}
			case "cols":
				m.Cols = NewCT_Columns()
				if err := d.DecodeElement(m.Cols, &el); err != nil {
					return err
				}
			case "formProt":
				m.FormProt = NewCT_OnOff()
				if err := d.DecodeElement(m.FormProt, &el); err != nil {
					return err
				}
			case "vAlign":
				m.VAlign = NewCT_VerticalJc()
				if err := d.DecodeElement(m.VAlign, &el); err != nil {
					return err
				}
			case "noEndnote":
				m.NoEndnote = NewCT_OnOff()
				if err := d.DecodeElement(m.NoEndnote, &el); err != nil {
					return err
				}
			case "titlePg":
				m.TitlePg = NewCT_OnOff()
				if err := d.DecodeElement(m.TitlePg, &el); err != nil {
					return err
				}
			case "textDirection":
				m.TextDirection = NewCT_TextDirection()
				if err := d.DecodeElement(m.TextDirection, &el); err != nil {
					return err
				}
			case "bidi":
				m.Bidi = NewCT_OnOff()
				if err := d.DecodeElement(m.Bidi, &el); err != nil {
					return err
				}
			case "rtlGutter":
				m.RtlGutter = NewCT_OnOff()
				if err := d.DecodeElement(m.RtlGutter, &el); err != nil {
					return err
				}
			case "docGrid":
				m.DocGrid = NewCT_DocGrid()
				if err := d.DecodeElement(m.DocGrid, &el); err != nil {
					return err
				}
			case "printerSettings":
				m.PrinterSettings = NewCT_Rel()
				if err := d.DecodeElement(m.PrinterSettings, &el); err != nil {
					return err
				}
			case "sectPrChange":
				m.SectPrChange = NewCT_SectPrChange()
				if err := d.DecodeElement(m.SectPrChange, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SectPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SectPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SectPr and its children
func (m *CT_SectPr) Validate() error {
	return m.ValidateWithPath("CT_SectPr")
}

// ValidateWithPath validates the CT_SectPr and its children, prefixing error messages with path
func (m *CT_SectPr) ValidateWithPath(path string) error {
	for i, v := range m.EG_HdrFtrReferences {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_HdrFtrReferences[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.FootnotePr != nil {
		if err := m.FootnotePr.ValidateWithPath(path + "/FootnotePr"); err != nil {
			return err
		}
	}
	if m.EndnotePr != nil {
		if err := m.EndnotePr.ValidateWithPath(path + "/EndnotePr"); err != nil {
			return err
		}
	}
	if m.Type != nil {
		if err := m.Type.ValidateWithPath(path + "/Type"); err != nil {
			return err
		}
	}
	if m.PgSz != nil {
		if err := m.PgSz.ValidateWithPath(path + "/PgSz"); err != nil {
			return err
		}
	}
	if m.PgMar != nil {
		if err := m.PgMar.ValidateWithPath(path + "/PgMar"); err != nil {
			return err
		}
	}
	if m.PaperSrc != nil {
		if err := m.PaperSrc.ValidateWithPath(path + "/PaperSrc"); err != nil {
			return err
		}
	}
	if m.PgBorders != nil {
		if err := m.PgBorders.ValidateWithPath(path + "/PgBorders"); err != nil {
			return err
		}
	}
	if m.LnNumType != nil {
		if err := m.LnNumType.ValidateWithPath(path + "/LnNumType"); err != nil {
			return err
		}
	}
	if m.PgNumType != nil {
		if err := m.PgNumType.ValidateWithPath(path + "/PgNumType"); err != nil {
			return err
		}
	}
	if m.Cols != nil {
		if err := m.Cols.ValidateWithPath(path + "/Cols"); err != nil {
			return err
		}
	}
	if m.FormProt != nil {
		if err := m.FormProt.ValidateWithPath(path + "/FormProt"); err != nil {
			return err
		}
	}
	if m.VAlign != nil {
		if err := m.VAlign.ValidateWithPath(path + "/VAlign"); err != nil {
			return err
		}
	}
	if m.NoEndnote != nil {
		if err := m.NoEndnote.ValidateWithPath(path + "/NoEndnote"); err != nil {
			return err
		}
	}
	if m.TitlePg != nil {
		if err := m.TitlePg.ValidateWithPath(path + "/TitlePg"); err != nil {
			return err
		}
	}
	if m.TextDirection != nil {
		if err := m.TextDirection.ValidateWithPath(path + "/TextDirection"); err != nil {
			return err
		}
	}
	if m.Bidi != nil {
		if err := m.Bidi.ValidateWithPath(path + "/Bidi"); err != nil {
			return err
		}
	}
	if m.RtlGutter != nil {
		if err := m.RtlGutter.ValidateWithPath(path + "/RtlGutter"); err != nil {
			return err
		}
	}
	if m.DocGrid != nil {
		if err := m.DocGrid.ValidateWithPath(path + "/DocGrid"); err != nil {
			return err
		}
	}
	if m.PrinterSettings != nil {
		if err := m.PrinterSettings.ValidateWithPath(path + "/PrinterSettings"); err != nil {
			return err
		}
	}
	if m.SectPrChange != nil {
		if err := m.SectPrChange.ValidateWithPath(path + "/SectPrChange"); err != nil {
			return err
		}
	}
	return nil
}
