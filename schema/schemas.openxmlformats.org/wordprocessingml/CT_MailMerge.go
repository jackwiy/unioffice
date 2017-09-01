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

type CT_MailMerge struct {
	// Source Document Type
	MainDocumentType *CT_MailMergeDocType
	// Query Contains Link to External Query File
	LinkToQuery *CT_OnOff
	// Data Source Type
	DataType *CT_MailMergeDataType
	// Data Source Connection String
	ConnectString *CT_String
	// Query For Data Source Records To Merge
	Query *CT_String
	// Data Source File Path
	DataSource *CT_Rel
	// Header Definition File Path
	HeaderSource *CT_Rel
	// Remove Blank Lines from Merged Documents
	DoNotSuppressBlankLines *CT_OnOff
	// Merged Document Destination
	Destination *CT_MailMergeDest
	// Column Containing E-mail Address
	AddressFieldName *CT_String
	// Merged E-mail or Fax Subject Line
	MailSubject *CT_String
	// Merged Document To E-Mail Attachment
	MailAsAttachment *CT_OnOff
	// View Merged Data Within Document
	ViewMergedData *CT_OnOff
	// Record Currently Displayed In Merged Document
	ActiveRecord *CT_DecimalNumber
	// Mail Merge Error Reporting Setting
	CheckErrors *CT_DecimalNumber
	// Office Data Source Object Settings
	Odso *CT_Odso
}

func NewCT_MailMerge() *CT_MailMerge {
	ret := &CT_MailMerge{}
	ret.MainDocumentType = NewCT_MailMergeDocType()
	ret.DataType = NewCT_MailMergeDataType()
	return ret
}

func (m *CT_MailMerge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	semainDocumentType := xml.StartElement{Name: xml.Name{Local: "w:mainDocumentType"}}
	e.EncodeElement(m.MainDocumentType, semainDocumentType)
	if m.LinkToQuery != nil {
		selinkToQuery := xml.StartElement{Name: xml.Name{Local: "w:linkToQuery"}}
		e.EncodeElement(m.LinkToQuery, selinkToQuery)
	}
	sedataType := xml.StartElement{Name: xml.Name{Local: "w:dataType"}}
	e.EncodeElement(m.DataType, sedataType)
	if m.ConnectString != nil {
		seconnectString := xml.StartElement{Name: xml.Name{Local: "w:connectString"}}
		e.EncodeElement(m.ConnectString, seconnectString)
	}
	if m.Query != nil {
		sequery := xml.StartElement{Name: xml.Name{Local: "w:query"}}
		e.EncodeElement(m.Query, sequery)
	}
	if m.DataSource != nil {
		sedataSource := xml.StartElement{Name: xml.Name{Local: "w:dataSource"}}
		e.EncodeElement(m.DataSource, sedataSource)
	}
	if m.HeaderSource != nil {
		seheaderSource := xml.StartElement{Name: xml.Name{Local: "w:headerSource"}}
		e.EncodeElement(m.HeaderSource, seheaderSource)
	}
	if m.DoNotSuppressBlankLines != nil {
		sedoNotSuppressBlankLines := xml.StartElement{Name: xml.Name{Local: "w:doNotSuppressBlankLines"}}
		e.EncodeElement(m.DoNotSuppressBlankLines, sedoNotSuppressBlankLines)
	}
	if m.Destination != nil {
		sedestination := xml.StartElement{Name: xml.Name{Local: "w:destination"}}
		e.EncodeElement(m.Destination, sedestination)
	}
	if m.AddressFieldName != nil {
		seaddressFieldName := xml.StartElement{Name: xml.Name{Local: "w:addressFieldName"}}
		e.EncodeElement(m.AddressFieldName, seaddressFieldName)
	}
	if m.MailSubject != nil {
		semailSubject := xml.StartElement{Name: xml.Name{Local: "w:mailSubject"}}
		e.EncodeElement(m.MailSubject, semailSubject)
	}
	if m.MailAsAttachment != nil {
		semailAsAttachment := xml.StartElement{Name: xml.Name{Local: "w:mailAsAttachment"}}
		e.EncodeElement(m.MailAsAttachment, semailAsAttachment)
	}
	if m.ViewMergedData != nil {
		seviewMergedData := xml.StartElement{Name: xml.Name{Local: "w:viewMergedData"}}
		e.EncodeElement(m.ViewMergedData, seviewMergedData)
	}
	if m.ActiveRecord != nil {
		seactiveRecord := xml.StartElement{Name: xml.Name{Local: "w:activeRecord"}}
		e.EncodeElement(m.ActiveRecord, seactiveRecord)
	}
	if m.CheckErrors != nil {
		secheckErrors := xml.StartElement{Name: xml.Name{Local: "w:checkErrors"}}
		e.EncodeElement(m.CheckErrors, secheckErrors)
	}
	if m.Odso != nil {
		seodso := xml.StartElement{Name: xml.Name{Local: "w:odso"}}
		e.EncodeElement(m.Odso, seodso)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MailMerge) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.MainDocumentType = NewCT_MailMergeDocType()
	m.DataType = NewCT_MailMergeDataType()
lCT_MailMerge:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "mainDocumentType":
				if err := d.DecodeElement(m.MainDocumentType, &el); err != nil {
					return err
				}
			case "linkToQuery":
				m.LinkToQuery = NewCT_OnOff()
				if err := d.DecodeElement(m.LinkToQuery, &el); err != nil {
					return err
				}
			case "dataType":
				if err := d.DecodeElement(m.DataType, &el); err != nil {
					return err
				}
			case "connectString":
				m.ConnectString = NewCT_String()
				if err := d.DecodeElement(m.ConnectString, &el); err != nil {
					return err
				}
			case "query":
				m.Query = NewCT_String()
				if err := d.DecodeElement(m.Query, &el); err != nil {
					return err
				}
			case "dataSource":
				m.DataSource = NewCT_Rel()
				if err := d.DecodeElement(m.DataSource, &el); err != nil {
					return err
				}
			case "headerSource":
				m.HeaderSource = NewCT_Rel()
				if err := d.DecodeElement(m.HeaderSource, &el); err != nil {
					return err
				}
			case "doNotSuppressBlankLines":
				m.DoNotSuppressBlankLines = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotSuppressBlankLines, &el); err != nil {
					return err
				}
			case "destination":
				m.Destination = NewCT_MailMergeDest()
				if err := d.DecodeElement(m.Destination, &el); err != nil {
					return err
				}
			case "addressFieldName":
				m.AddressFieldName = NewCT_String()
				if err := d.DecodeElement(m.AddressFieldName, &el); err != nil {
					return err
				}
			case "mailSubject":
				m.MailSubject = NewCT_String()
				if err := d.DecodeElement(m.MailSubject, &el); err != nil {
					return err
				}
			case "mailAsAttachment":
				m.MailAsAttachment = NewCT_OnOff()
				if err := d.DecodeElement(m.MailAsAttachment, &el); err != nil {
					return err
				}
			case "viewMergedData":
				m.ViewMergedData = NewCT_OnOff()
				if err := d.DecodeElement(m.ViewMergedData, &el); err != nil {
					return err
				}
			case "activeRecord":
				m.ActiveRecord = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.ActiveRecord, &el); err != nil {
					return err
				}
			case "checkErrors":
				m.CheckErrors = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.CheckErrors, &el); err != nil {
					return err
				}
			case "odso":
				m.Odso = NewCT_Odso()
				if err := d.DecodeElement(m.Odso, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_MailMerge %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MailMerge
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MailMerge and its children
func (m *CT_MailMerge) Validate() error {
	return m.ValidateWithPath("CT_MailMerge")
}

// ValidateWithPath validates the CT_MailMerge and its children, prefixing error messages with path
func (m *CT_MailMerge) ValidateWithPath(path string) error {
	if err := m.MainDocumentType.ValidateWithPath(path + "/MainDocumentType"); err != nil {
		return err
	}
	if m.LinkToQuery != nil {
		if err := m.LinkToQuery.ValidateWithPath(path + "/LinkToQuery"); err != nil {
			return err
		}
	}
	if err := m.DataType.ValidateWithPath(path + "/DataType"); err != nil {
		return err
	}
	if m.ConnectString != nil {
		if err := m.ConnectString.ValidateWithPath(path + "/ConnectString"); err != nil {
			return err
		}
	}
	if m.Query != nil {
		if err := m.Query.ValidateWithPath(path + "/Query"); err != nil {
			return err
		}
	}
	if m.DataSource != nil {
		if err := m.DataSource.ValidateWithPath(path + "/DataSource"); err != nil {
			return err
		}
	}
	if m.HeaderSource != nil {
		if err := m.HeaderSource.ValidateWithPath(path + "/HeaderSource"); err != nil {
			return err
		}
	}
	if m.DoNotSuppressBlankLines != nil {
		if err := m.DoNotSuppressBlankLines.ValidateWithPath(path + "/DoNotSuppressBlankLines"); err != nil {
			return err
		}
	}
	if m.Destination != nil {
		if err := m.Destination.ValidateWithPath(path + "/Destination"); err != nil {
			return err
		}
	}
	if m.AddressFieldName != nil {
		if err := m.AddressFieldName.ValidateWithPath(path + "/AddressFieldName"); err != nil {
			return err
		}
	}
	if m.MailSubject != nil {
		if err := m.MailSubject.ValidateWithPath(path + "/MailSubject"); err != nil {
			return err
		}
	}
	if m.MailAsAttachment != nil {
		if err := m.MailAsAttachment.ValidateWithPath(path + "/MailAsAttachment"); err != nil {
			return err
		}
	}
	if m.ViewMergedData != nil {
		if err := m.ViewMergedData.ValidateWithPath(path + "/ViewMergedData"); err != nil {
			return err
		}
	}
	if m.ActiveRecord != nil {
		if err := m.ActiveRecord.ValidateWithPath(path + "/ActiveRecord"); err != nil {
			return err
		}
	}
	if m.CheckErrors != nil {
		if err := m.CheckErrors.ValidateWithPath(path + "/CheckErrors"); err != nil {
			return err
		}
	}
	if m.Odso != nil {
		if err := m.Odso.ValidateWithPath(path + "/Odso"); err != nil {
			return err
		}
	}
	return nil
}
