// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
	"strconv"
)

type PivotTableDefinition struct {
	CT_pivotTableDefinition
}

func NewPivotTableDefinition() *PivotTableDefinition {
	ret := &PivotTableDefinition{}
	ret.CT_pivotTableDefinition = *NewCT_pivotTableDefinition()
	return ret
}

func (m *PivotTableDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:pivotTableDefinition"
	return m.CT_pivotTableDefinition.MarshalXML(e, start)
}

func (m *PivotTableDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_pivotTableDefinition = *NewCT_pivotTableDefinition()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "cacheId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.CacheIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "dataOnRows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataOnRowsAttr = &parsed
		}
		if attr.Name.Local == "dataPosition" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DataPositionAttr = &pt
		}
		if attr.Name.Local == "dataCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataCaptionAttr = parsed
		}
		if attr.Name.Local == "grandTotalCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GrandTotalCaptionAttr = &parsed
		}
		if attr.Name.Local == "errorCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ErrorCaptionAttr = &parsed
		}
		if attr.Name.Local == "showError" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowErrorAttr = &parsed
		}
		if attr.Name.Local == "missingCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MissingCaptionAttr = &parsed
		}
		if attr.Name.Local == "showMissing" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMissingAttr = &parsed
		}
		if attr.Name.Local == "pageStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PageStyleAttr = &parsed
		}
		if attr.Name.Local == "pivotTableStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PivotTableStyleAttr = &parsed
		}
		if attr.Name.Local == "vacatedStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VacatedStyleAttr = &parsed
		}
		if attr.Name.Local == "tag" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TagAttr = &parsed
		}
		if attr.Name.Local == "updatedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.UpdatedVersionAttr = &pt
		}
		if attr.Name.Local == "minRefreshableVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.MinRefreshableVersionAttr = &pt
		}
		if attr.Name.Local == "asteriskTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AsteriskTotalsAttr = &parsed
		}
		if attr.Name.Local == "showItems" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowItemsAttr = &parsed
		}
		if attr.Name.Local == "editData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EditDataAttr = &parsed
		}
		if attr.Name.Local == "disableFieldList" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableFieldListAttr = &parsed
		}
		if attr.Name.Local == "showCalcMbrs" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCalcMbrsAttr = &parsed
		}
		if attr.Name.Local == "visualTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VisualTotalsAttr = &parsed
		}
		if attr.Name.Local == "showMultipleLabel" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMultipleLabelAttr = &parsed
		}
		if attr.Name.Local == "showDataDropDown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDataDropDownAttr = &parsed
		}
		if attr.Name.Local == "showDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDrillAttr = &parsed
		}
		if attr.Name.Local == "printDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PrintDrillAttr = &parsed
		}
		if attr.Name.Local == "showMemberPropertyTips" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMemberPropertyTipsAttr = &parsed
		}
		if attr.Name.Local == "showDataTips" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDataTipsAttr = &parsed
		}
		if attr.Name.Local == "enableWizard" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableWizardAttr = &parsed
		}
		if attr.Name.Local == "enableDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableDrillAttr = &parsed
		}
		if attr.Name.Local == "enableFieldProperties" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableFieldPropertiesAttr = &parsed
		}
		if attr.Name.Local == "preserveFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveFormattingAttr = &parsed
		}
		if attr.Name.Local == "useAutoFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseAutoFormattingAttr = &parsed
		}
		if attr.Name.Local == "pageWrap" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PageWrapAttr = &pt
		}
		if attr.Name.Local == "pageOverThenDown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PageOverThenDownAttr = &parsed
		}
		if attr.Name.Local == "subtotalHiddenItems" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SubtotalHiddenItemsAttr = &parsed
		}
		if attr.Name.Local == "rowGrandTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowGrandTotalsAttr = &parsed
		}
		if attr.Name.Local == "colGrandTotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ColGrandTotalsAttr = &parsed
		}
		if attr.Name.Local == "fieldPrintTitles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FieldPrintTitlesAttr = &parsed
		}
		if attr.Name.Local == "itemPrintTitles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ItemPrintTitlesAttr = &parsed
		}
		if attr.Name.Local == "mergeItem" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MergeItemAttr = &parsed
		}
		if attr.Name.Local == "showDropZones" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDropZonesAttr = &parsed
		}
		if attr.Name.Local == "createdVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.CreatedVersionAttr = &pt
		}
		if attr.Name.Local == "indent" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IndentAttr = &pt
		}
		if attr.Name.Local == "showEmptyRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowEmptyRowAttr = &parsed
		}
		if attr.Name.Local == "showEmptyCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowEmptyColAttr = &parsed
		}
		if attr.Name.Local == "showHeaders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowHeadersAttr = &parsed
		}
		if attr.Name.Local == "compact" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompactAttr = &parsed
		}
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
		}
		if attr.Name.Local == "outlineData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineDataAttr = &parsed
		}
		if attr.Name.Local == "compactData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompactDataAttr = &parsed
		}
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
		}
		if attr.Name.Local == "gridDropZones" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GridDropZonesAttr = &parsed
		}
		if attr.Name.Local == "immersive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ImmersiveAttr = &parsed
		}
		if attr.Name.Local == "multipleFieldFilters" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MultipleFieldFiltersAttr = &parsed
		}
		if attr.Name.Local == "chartFormat" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ChartFormatAttr = &pt
		}
		if attr.Name.Local == "rowHeaderCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RowHeaderCaptionAttr = &parsed
		}
		if attr.Name.Local == "colHeaderCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColHeaderCaptionAttr = &parsed
		}
		if attr.Name.Local == "fieldListSortAscending" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FieldListSortAscendingAttr = &parsed
		}
		if attr.Name.Local == "mdxSubqueries" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MdxSubqueriesAttr = &parsed
		}
		if attr.Name.Local == "customListSort" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomListSortAttr = &parsed
		}
		if attr.Name.Local == "autoFormatId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AutoFormatIdAttr = &pt
		}
		if attr.Name.Local == "applyNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyBorderFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyFontFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyPatternFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyPatternFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyAlignmentFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentFormatsAttr = &parsed
		}
		if attr.Name.Local == "applyWidthHeightFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyWidthHeightFormatsAttr = &parsed
		}
	}
lPivotTableDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "location":
				if err := d.DecodeElement(m.Location, &el); err != nil {
					return err
				}
			case "pivotFields":
				m.PivotFields = NewCT_PivotFields()
				if err := d.DecodeElement(m.PivotFields, &el); err != nil {
					return err
				}
			case "rowFields":
				m.RowFields = NewCT_RowFields()
				if err := d.DecodeElement(m.RowFields, &el); err != nil {
					return err
				}
			case "rowItems":
				m.RowItems = NewCT_rowItems()
				if err := d.DecodeElement(m.RowItems, &el); err != nil {
					return err
				}
			case "colFields":
				m.ColFields = NewCT_ColFields()
				if err := d.DecodeElement(m.ColFields, &el); err != nil {
					return err
				}
			case "colItems":
				m.ColItems = NewCT_colItems()
				if err := d.DecodeElement(m.ColItems, &el); err != nil {
					return err
				}
			case "pageFields":
				m.PageFields = NewCT_PageFields()
				if err := d.DecodeElement(m.PageFields, &el); err != nil {
					return err
				}
			case "dataFields":
				m.DataFields = NewCT_DataFields()
				if err := d.DecodeElement(m.DataFields, &el); err != nil {
					return err
				}
			case "formats":
				m.Formats = NewCT_Formats()
				if err := d.DecodeElement(m.Formats, &el); err != nil {
					return err
				}
			case "conditionalFormats":
				m.ConditionalFormats = NewCT_ConditionalFormats()
				if err := d.DecodeElement(m.ConditionalFormats, &el); err != nil {
					return err
				}
			case "chartFormats":
				m.ChartFormats = NewCT_ChartFormats()
				if err := d.DecodeElement(m.ChartFormats, &el); err != nil {
					return err
				}
			case "pivotHierarchies":
				m.PivotHierarchies = NewCT_PivotHierarchies()
				if err := d.DecodeElement(m.PivotHierarchies, &el); err != nil {
					return err
				}
			case "pivotTableStyleInfo":
				m.PivotTableStyleInfo = NewCT_PivotTableStyle()
				if err := d.DecodeElement(m.PivotTableStyleInfo, &el); err != nil {
					return err
				}
			case "filters":
				m.Filters = NewCT_PivotFilters()
				if err := d.DecodeElement(m.Filters, &el); err != nil {
					return err
				}
			case "rowHierarchiesUsage":
				m.RowHierarchiesUsage = NewCT_RowHierarchiesUsage()
				if err := d.DecodeElement(m.RowHierarchiesUsage, &el); err != nil {
					return err
				}
			case "colHierarchiesUsage":
				m.ColHierarchiesUsage = NewCT_ColHierarchiesUsage()
				if err := d.DecodeElement(m.ColHierarchiesUsage, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on PivotTableDefinition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPivotTableDefinition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the PivotTableDefinition and its children
func (m *PivotTableDefinition) Validate() error {
	return m.ValidateWithPath("PivotTableDefinition")
}

// ValidateWithPath validates the PivotTableDefinition and its children, prefixing error messages with path
func (m *PivotTableDefinition) ValidateWithPath(path string) error {
	if err := m.CT_pivotTableDefinition.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
