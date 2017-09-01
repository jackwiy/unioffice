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
)

type CT_Macrosheet struct {
	// Sheet Properties
	SheetPr *CT_SheetPr
	// Macro Sheet Dimensions
	Dimension *CT_SheetDimension
	// Macro Sheet Views
	SheetViews *CT_SheetViews
	// Sheet Format Properties
	SheetFormatPr *CT_SheetFormatPr
	// Column Information
	Cols []*CT_Cols
	// Sheet Data
	SheetData *CT_SheetData
	// Sheet Protection Options
	SheetProtection *CT_SheetProtection
	// AutoFilter
	AutoFilter *CT_AutoFilter
	// Sort State
	SortState *CT_SortState
	// Data Consolidation
	DataConsolidate *CT_DataConsolidate
	// Custom Sheet Views
	CustomSheetViews *CT_CustomSheetViews
	// Phonetic Properties
	PhoneticPr *CT_PhoneticPr
	// Conditional Formatting
	ConditionalFormatting []*CT_ConditionalFormatting
	// Print Options
	PrintOptions *CT_PrintOptions
	// Page Margins
	PageMargins *CT_PageMargins
	// Page Setup Settings
	PageSetup *CT_PageSetup
	// Header Footer Settings
	HeaderFooter *CT_HeaderFooter
	// Horizontal Page Breaks (Row)
	RowBreaks *CT_PageBreak
	// Vertical Page Breaks
	ColBreaks *CT_PageBreak
	// Custom Properties
	CustomProperties *CT_CustomProperties
	// Drawing
	Drawing *CT_Drawing
	// Legacy Drawing Reference
	LegacyDrawing *CT_LegacyDrawing
	// Legacy Drawing Header Footer
	LegacyDrawingHF *CT_LegacyDrawing
	DrawingHF       *CT_DrawingHF
	// Background Image
	Picture *CT_SheetBackgroundPicture
	// Embedded Objects
	OleObjects *CT_OleObjects
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Macrosheet() *CT_Macrosheet {
	ret := &CT_Macrosheet{}
	ret.SheetData = NewCT_SheetData()
	return ret
}

func (m *CT_Macrosheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "x:CT_Macrosheet"
	e.EncodeToken(start)
	if m.SheetPr != nil {
		sesheetPr := xml.StartElement{Name: xml.Name{Local: "x:sheetPr"}}
		e.EncodeElement(m.SheetPr, sesheetPr)
	}
	if m.Dimension != nil {
		sedimension := xml.StartElement{Name: xml.Name{Local: "x:dimension"}}
		e.EncodeElement(m.Dimension, sedimension)
	}
	if m.SheetViews != nil {
		sesheetViews := xml.StartElement{Name: xml.Name{Local: "x:sheetViews"}}
		e.EncodeElement(m.SheetViews, sesheetViews)
	}
	if m.SheetFormatPr != nil {
		sesheetFormatPr := xml.StartElement{Name: xml.Name{Local: "x:sheetFormatPr"}}
		e.EncodeElement(m.SheetFormatPr, sesheetFormatPr)
	}
	if m.Cols != nil {
		secols := xml.StartElement{Name: xml.Name{Local: "x:cols"}}
		e.EncodeElement(m.Cols, secols)
	}
	sesheetData := xml.StartElement{Name: xml.Name{Local: "x:sheetData"}}
	e.EncodeElement(m.SheetData, sesheetData)
	if m.SheetProtection != nil {
		sesheetProtection := xml.StartElement{Name: xml.Name{Local: "x:sheetProtection"}}
		e.EncodeElement(m.SheetProtection, sesheetProtection)
	}
	if m.AutoFilter != nil {
		seautoFilter := xml.StartElement{Name: xml.Name{Local: "x:autoFilter"}}
		e.EncodeElement(m.AutoFilter, seautoFilter)
	}
	if m.SortState != nil {
		sesortState := xml.StartElement{Name: xml.Name{Local: "x:sortState"}}
		e.EncodeElement(m.SortState, sesortState)
	}
	if m.DataConsolidate != nil {
		sedataConsolidate := xml.StartElement{Name: xml.Name{Local: "x:dataConsolidate"}}
		e.EncodeElement(m.DataConsolidate, sedataConsolidate)
	}
	if m.CustomSheetViews != nil {
		secustomSheetViews := xml.StartElement{Name: xml.Name{Local: "x:customSheetViews"}}
		e.EncodeElement(m.CustomSheetViews, secustomSheetViews)
	}
	if m.PhoneticPr != nil {
		sephoneticPr := xml.StartElement{Name: xml.Name{Local: "x:phoneticPr"}}
		e.EncodeElement(m.PhoneticPr, sephoneticPr)
	}
	if m.ConditionalFormatting != nil {
		seconditionalFormatting := xml.StartElement{Name: xml.Name{Local: "x:conditionalFormatting"}}
		e.EncodeElement(m.ConditionalFormatting, seconditionalFormatting)
	}
	if m.PrintOptions != nil {
		seprintOptions := xml.StartElement{Name: xml.Name{Local: "x:printOptions"}}
		e.EncodeElement(m.PrintOptions, seprintOptions)
	}
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "x:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "x:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "x:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	if m.RowBreaks != nil {
		serowBreaks := xml.StartElement{Name: xml.Name{Local: "x:rowBreaks"}}
		e.EncodeElement(m.RowBreaks, serowBreaks)
	}
	if m.ColBreaks != nil {
		secolBreaks := xml.StartElement{Name: xml.Name{Local: "x:colBreaks"}}
		e.EncodeElement(m.ColBreaks, secolBreaks)
	}
	if m.CustomProperties != nil {
		secustomProperties := xml.StartElement{Name: xml.Name{Local: "x:customProperties"}}
		e.EncodeElement(m.CustomProperties, secustomProperties)
	}
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "x:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	if m.LegacyDrawing != nil {
		selegacyDrawing := xml.StartElement{Name: xml.Name{Local: "x:legacyDrawing"}}
		e.EncodeElement(m.LegacyDrawing, selegacyDrawing)
	}
	if m.LegacyDrawingHF != nil {
		selegacyDrawingHF := xml.StartElement{Name: xml.Name{Local: "x:legacyDrawingHF"}}
		e.EncodeElement(m.LegacyDrawingHF, selegacyDrawingHF)
	}
	if m.DrawingHF != nil {
		sedrawingHF := xml.StartElement{Name: xml.Name{Local: "x:drawingHF"}}
		e.EncodeElement(m.DrawingHF, sedrawingHF)
	}
	if m.Picture != nil {
		sepicture := xml.StartElement{Name: xml.Name{Local: "x:picture"}}
		e.EncodeElement(m.Picture, sepicture)
	}
	if m.OleObjects != nil {
		seoleObjects := xml.StartElement{Name: xml.Name{Local: "x:oleObjects"}}
		e.EncodeElement(m.OleObjects, seoleObjects)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Macrosheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SheetData = NewCT_SheetData()
lCT_Macrosheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetPr":
				m.SheetPr = NewCT_SheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case "dimension":
				m.Dimension = NewCT_SheetDimension()
				if err := d.DecodeElement(m.Dimension, &el); err != nil {
					return err
				}
			case "sheetViews":
				m.SheetViews = NewCT_SheetViews()
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case "sheetFormatPr":
				m.SheetFormatPr = NewCT_SheetFormatPr()
				if err := d.DecodeElement(m.SheetFormatPr, &el); err != nil {
					return err
				}
			case "cols":
				tmp := NewCT_Cols()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cols = append(m.Cols, tmp)
			case "sheetData":
				if err := d.DecodeElement(m.SheetData, &el); err != nil {
					return err
				}
			case "sheetProtection":
				m.SheetProtection = NewCT_SheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case "autoFilter":
				m.AutoFilter = NewCT_AutoFilter()
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case "sortState":
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case "dataConsolidate":
				m.DataConsolidate = NewCT_DataConsolidate()
				if err := d.DecodeElement(m.DataConsolidate, &el); err != nil {
					return err
				}
			case "customSheetViews":
				m.CustomSheetViews = NewCT_CustomSheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case "phoneticPr":
				m.PhoneticPr = NewCT_PhoneticPr()
				if err := d.DecodeElement(m.PhoneticPr, &el); err != nil {
					return err
				}
			case "conditionalFormatting":
				tmp := NewCT_ConditionalFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConditionalFormatting = append(m.ConditionalFormatting, tmp)
			case "printOptions":
				m.PrintOptions = NewCT_PrintOptions()
				if err := d.DecodeElement(m.PrintOptions, &el); err != nil {
					return err
				}
			case "pageMargins":
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case "pageSetup":
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case "headerFooter":
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case "rowBreaks":
				m.RowBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.RowBreaks, &el); err != nil {
					return err
				}
			case "colBreaks":
				m.ColBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.ColBreaks, &el); err != nil {
					return err
				}
			case "customProperties":
				m.CustomProperties = NewCT_CustomProperties()
				if err := d.DecodeElement(m.CustomProperties, &el); err != nil {
					return err
				}
			case "drawing":
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case "legacyDrawing":
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case "legacyDrawingHF":
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case "drawingHF":
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case "picture":
				m.Picture = NewCT_SheetBackgroundPicture()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case "oleObjects":
				m.OleObjects = NewCT_OleObjects()
				if err := d.DecodeElement(m.OleObjects, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Macrosheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Macrosheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Macrosheet and its children
func (m *CT_Macrosheet) Validate() error {
	return m.ValidateWithPath("CT_Macrosheet")
}

// ValidateWithPath validates the CT_Macrosheet and its children, prefixing error messages with path
func (m *CT_Macrosheet) ValidateWithPath(path string) error {
	if m.SheetPr != nil {
		if err := m.SheetPr.ValidateWithPath(path + "/SheetPr"); err != nil {
			return err
		}
	}
	if m.Dimension != nil {
		if err := m.Dimension.ValidateWithPath(path + "/Dimension"); err != nil {
			return err
		}
	}
	if m.SheetViews != nil {
		if err := m.SheetViews.ValidateWithPath(path + "/SheetViews"); err != nil {
			return err
		}
	}
	if m.SheetFormatPr != nil {
		if err := m.SheetFormatPr.ValidateWithPath(path + "/SheetFormatPr"); err != nil {
			return err
		}
	}
	for i, v := range m.Cols {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cols[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.SheetData.ValidateWithPath(path + "/SheetData"); err != nil {
		return err
	}
	if m.SheetProtection != nil {
		if err := m.SheetProtection.ValidateWithPath(path + "/SheetProtection"); err != nil {
			return err
		}
	}
	if m.AutoFilter != nil {
		if err := m.AutoFilter.ValidateWithPath(path + "/AutoFilter"); err != nil {
			return err
		}
	}
	if m.SortState != nil {
		if err := m.SortState.ValidateWithPath(path + "/SortState"); err != nil {
			return err
		}
	}
	if m.DataConsolidate != nil {
		if err := m.DataConsolidate.ValidateWithPath(path + "/DataConsolidate"); err != nil {
			return err
		}
	}
	if m.CustomSheetViews != nil {
		if err := m.CustomSheetViews.ValidateWithPath(path + "/CustomSheetViews"); err != nil {
			return err
		}
	}
	if m.PhoneticPr != nil {
		if err := m.PhoneticPr.ValidateWithPath(path + "/PhoneticPr"); err != nil {
			return err
		}
	}
	for i, v := range m.ConditionalFormatting {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ConditionalFormatting[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.PrintOptions != nil {
		if err := m.PrintOptions.ValidateWithPath(path + "/PrintOptions"); err != nil {
			return err
		}
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	if m.RowBreaks != nil {
		if err := m.RowBreaks.ValidateWithPath(path + "/RowBreaks"); err != nil {
			return err
		}
	}
	if m.ColBreaks != nil {
		if err := m.ColBreaks.ValidateWithPath(path + "/ColBreaks"); err != nil {
			return err
		}
	}
	if m.CustomProperties != nil {
		if err := m.CustomProperties.ValidateWithPath(path + "/CustomProperties"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawing != nil {
		if err := m.LegacyDrawing.ValidateWithPath(path + "/LegacyDrawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawingHF != nil {
		if err := m.LegacyDrawingHF.ValidateWithPath(path + "/LegacyDrawingHF"); err != nil {
			return err
		}
	}
	if m.DrawingHF != nil {
		if err := m.DrawingHF.ValidateWithPath(path + "/DrawingHF"); err != nil {
			return err
		}
	}
	if m.Picture != nil {
		if err := m.Picture.ValidateWithPath(path + "/Picture"); err != nil {
			return err
		}
	}
	if m.OleObjects != nil {
		if err := m.OleObjects.ValidateWithPath(path + "/OleObjects"); err != nil {
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
