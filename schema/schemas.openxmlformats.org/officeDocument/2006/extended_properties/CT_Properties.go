// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package extended_properties

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml"
)

type CT_Properties struct {
	// Name of Document Template
	Template *string
	// Name of Manager
	Manager *string
	// Name of Company
	Company *string
	// Total Number of Pages
	Pages *int32
	// Word Count
	Words *int32
	// Total Number of Characters
	Characters *int32
	// Intended Format of Presentation
	PresentationFormat *string
	// Number of Lines
	Lines *int32
	// Total Number of Paragraphs
	Paragraphs *int32
	// Slides Metadata Element
	Slides *int32
	// Number of Slides Containing Notes
	Notes *int32
	// Total Edit Time Metadata Element
	TotalTime *int32
	// Number of Hidden Slides
	HiddenSlides *int32
	// Total Number of Multimedia Clips
	MMClips *int32
	// Thumbnail Display Mode
	ScaleCrop *bool
	// Heading Pairs
	HeadingPairs *CT_VectorVariant
	// Part Titles
	TitlesOfParts *CT_VectorLpstr
	// Links Up-to-Date
	LinksUpToDate *bool
	// Number of Characters (With Spaces)
	CharactersWithSpaces *int32
	// Shared Document
	SharedDoc *bool
	// Relative Hyperlink Base
	HyperlinkBase *string
	// Hyperlink List
	HLinks *CT_VectorVariant
	// Hyperlinks Changed
	HyperlinksChanged *bool
	// Digital Signature
	DigSig *CT_DigSigBlob
	// Application Name
	Application *string
	// Application Version
	AppVersion *string
	// Document Security
	DocSecurity *int32
}

func NewCT_Properties() *CT_Properties {
	ret := &CT_Properties{}
	return ret
}

func (m *CT_Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Template != nil {
		seTemplate := xml.StartElement{Name: xml.Name{Local: "Template"}}
		gooxml.AddPreserveSpaceAttr(&seTemplate, *m.Template)
		e.EncodeElement(m.Template, seTemplate)
	}
	if m.Manager != nil {
		seManager := xml.StartElement{Name: xml.Name{Local: "Manager"}}
		gooxml.AddPreserveSpaceAttr(&seManager, *m.Manager)
		e.EncodeElement(m.Manager, seManager)
	}
	if m.Company != nil {
		seCompany := xml.StartElement{Name: xml.Name{Local: "Company"}}
		gooxml.AddPreserveSpaceAttr(&seCompany, *m.Company)
		e.EncodeElement(m.Company, seCompany)
	}
	if m.Pages != nil {
		sePages := xml.StartElement{Name: xml.Name{Local: "Pages"}}
		e.EncodeElement(m.Pages, sePages)
	}
	if m.Words != nil {
		seWords := xml.StartElement{Name: xml.Name{Local: "Words"}}
		e.EncodeElement(m.Words, seWords)
	}
	if m.Characters != nil {
		seCharacters := xml.StartElement{Name: xml.Name{Local: "Characters"}}
		e.EncodeElement(m.Characters, seCharacters)
	}
	if m.PresentationFormat != nil {
		sePresentationFormat := xml.StartElement{Name: xml.Name{Local: "PresentationFormat"}}
		gooxml.AddPreserveSpaceAttr(&sePresentationFormat, *m.PresentationFormat)
		e.EncodeElement(m.PresentationFormat, sePresentationFormat)
	}
	if m.Lines != nil {
		seLines := xml.StartElement{Name: xml.Name{Local: "Lines"}}
		e.EncodeElement(m.Lines, seLines)
	}
	if m.Paragraphs != nil {
		seParagraphs := xml.StartElement{Name: xml.Name{Local: "Paragraphs"}}
		e.EncodeElement(m.Paragraphs, seParagraphs)
	}
	if m.Slides != nil {
		seSlides := xml.StartElement{Name: xml.Name{Local: "Slides"}}
		e.EncodeElement(m.Slides, seSlides)
	}
	if m.Notes != nil {
		seNotes := xml.StartElement{Name: xml.Name{Local: "Notes"}}
		e.EncodeElement(m.Notes, seNotes)
	}
	if m.TotalTime != nil {
		seTotalTime := xml.StartElement{Name: xml.Name{Local: "TotalTime"}}
		e.EncodeElement(m.TotalTime, seTotalTime)
	}
	if m.HiddenSlides != nil {
		seHiddenSlides := xml.StartElement{Name: xml.Name{Local: "HiddenSlides"}}
		e.EncodeElement(m.HiddenSlides, seHiddenSlides)
	}
	if m.MMClips != nil {
		seMMClips := xml.StartElement{Name: xml.Name{Local: "MMClips"}}
		e.EncodeElement(m.MMClips, seMMClips)
	}
	if m.ScaleCrop != nil {
		seScaleCrop := xml.StartElement{Name: xml.Name{Local: "ScaleCrop"}}
		e.EncodeElement(m.ScaleCrop, seScaleCrop)
	}
	if m.HeadingPairs != nil {
		seHeadingPairs := xml.StartElement{Name: xml.Name{Local: "HeadingPairs"}}
		e.EncodeElement(m.HeadingPairs, seHeadingPairs)
	}
	if m.TitlesOfParts != nil {
		seTitlesOfParts := xml.StartElement{Name: xml.Name{Local: "TitlesOfParts"}}
		e.EncodeElement(m.TitlesOfParts, seTitlesOfParts)
	}
	if m.LinksUpToDate != nil {
		seLinksUpToDate := xml.StartElement{Name: xml.Name{Local: "LinksUpToDate"}}
		e.EncodeElement(m.LinksUpToDate, seLinksUpToDate)
	}
	if m.CharactersWithSpaces != nil {
		seCharactersWithSpaces := xml.StartElement{Name: xml.Name{Local: "CharactersWithSpaces"}}
		e.EncodeElement(m.CharactersWithSpaces, seCharactersWithSpaces)
	}
	if m.SharedDoc != nil {
		seSharedDoc := xml.StartElement{Name: xml.Name{Local: "SharedDoc"}}
		e.EncodeElement(m.SharedDoc, seSharedDoc)
	}
	if m.HyperlinkBase != nil {
		seHyperlinkBase := xml.StartElement{Name: xml.Name{Local: "HyperlinkBase"}}
		gooxml.AddPreserveSpaceAttr(&seHyperlinkBase, *m.HyperlinkBase)
		e.EncodeElement(m.HyperlinkBase, seHyperlinkBase)
	}
	if m.HLinks != nil {
		seHLinks := xml.StartElement{Name: xml.Name{Local: "HLinks"}}
		e.EncodeElement(m.HLinks, seHLinks)
	}
	if m.HyperlinksChanged != nil {
		seHyperlinksChanged := xml.StartElement{Name: xml.Name{Local: "HyperlinksChanged"}}
		e.EncodeElement(m.HyperlinksChanged, seHyperlinksChanged)
	}
	if m.DigSig != nil {
		seDigSig := xml.StartElement{Name: xml.Name{Local: "DigSig"}}
		e.EncodeElement(m.DigSig, seDigSig)
	}
	if m.Application != nil {
		seApplication := xml.StartElement{Name: xml.Name{Local: "Application"}}
		gooxml.AddPreserveSpaceAttr(&seApplication, *m.Application)
		e.EncodeElement(m.Application, seApplication)
	}
	if m.AppVersion != nil {
		seAppVersion := xml.StartElement{Name: xml.Name{Local: "AppVersion"}}
		gooxml.AddPreserveSpaceAttr(&seAppVersion, *m.AppVersion)
		e.EncodeElement(m.AppVersion, seAppVersion)
	}
	if m.DocSecurity != nil {
		seDocSecurity := xml.StartElement{Name: xml.Name{Local: "DocSecurity"}}
		e.EncodeElement(m.DocSecurity, seDocSecurity)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Properties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "Template":
				m.Template = new(string)
				if err := d.DecodeElement(m.Template, &el); err != nil {
					return err
				}
			case "Manager":
				m.Manager = new(string)
				if err := d.DecodeElement(m.Manager, &el); err != nil {
					return err
				}
			case "Company":
				m.Company = new(string)
				if err := d.DecodeElement(m.Company, &el); err != nil {
					return err
				}
			case "Pages":
				m.Pages = new(int32)
				if err := d.DecodeElement(m.Pages, &el); err != nil {
					return err
				}
			case "Words":
				m.Words = new(int32)
				if err := d.DecodeElement(m.Words, &el); err != nil {
					return err
				}
			case "Characters":
				m.Characters = new(int32)
				if err := d.DecodeElement(m.Characters, &el); err != nil {
					return err
				}
			case "PresentationFormat":
				m.PresentationFormat = new(string)
				if err := d.DecodeElement(m.PresentationFormat, &el); err != nil {
					return err
				}
			case "Lines":
				m.Lines = new(int32)
				if err := d.DecodeElement(m.Lines, &el); err != nil {
					return err
				}
			case "Paragraphs":
				m.Paragraphs = new(int32)
				if err := d.DecodeElement(m.Paragraphs, &el); err != nil {
					return err
				}
			case "Slides":
				m.Slides = new(int32)
				if err := d.DecodeElement(m.Slides, &el); err != nil {
					return err
				}
			case "Notes":
				m.Notes = new(int32)
				if err := d.DecodeElement(m.Notes, &el); err != nil {
					return err
				}
			case "TotalTime":
				m.TotalTime = new(int32)
				if err := d.DecodeElement(m.TotalTime, &el); err != nil {
					return err
				}
			case "HiddenSlides":
				m.HiddenSlides = new(int32)
				if err := d.DecodeElement(m.HiddenSlides, &el); err != nil {
					return err
				}
			case "MMClips":
				m.MMClips = new(int32)
				if err := d.DecodeElement(m.MMClips, &el); err != nil {
					return err
				}
			case "ScaleCrop":
				m.ScaleCrop = new(bool)
				if err := d.DecodeElement(m.ScaleCrop, &el); err != nil {
					return err
				}
			case "HeadingPairs":
				m.HeadingPairs = NewCT_VectorVariant()
				if err := d.DecodeElement(m.HeadingPairs, &el); err != nil {
					return err
				}
			case "TitlesOfParts":
				m.TitlesOfParts = NewCT_VectorLpstr()
				if err := d.DecodeElement(m.TitlesOfParts, &el); err != nil {
					return err
				}
			case "LinksUpToDate":
				m.LinksUpToDate = new(bool)
				if err := d.DecodeElement(m.LinksUpToDate, &el); err != nil {
					return err
				}
			case "CharactersWithSpaces":
				m.CharactersWithSpaces = new(int32)
				if err := d.DecodeElement(m.CharactersWithSpaces, &el); err != nil {
					return err
				}
			case "SharedDoc":
				m.SharedDoc = new(bool)
				if err := d.DecodeElement(m.SharedDoc, &el); err != nil {
					return err
				}
			case "HyperlinkBase":
				m.HyperlinkBase = new(string)
				if err := d.DecodeElement(m.HyperlinkBase, &el); err != nil {
					return err
				}
			case "HLinks":
				m.HLinks = NewCT_VectorVariant()
				if err := d.DecodeElement(m.HLinks, &el); err != nil {
					return err
				}
			case "HyperlinksChanged":
				m.HyperlinksChanged = new(bool)
				if err := d.DecodeElement(m.HyperlinksChanged, &el); err != nil {
					return err
				}
			case "DigSig":
				m.DigSig = NewCT_DigSigBlob()
				if err := d.DecodeElement(m.DigSig, &el); err != nil {
					return err
				}
			case "Application":
				m.Application = new(string)
				if err := d.DecodeElement(m.Application, &el); err != nil {
					return err
				}
			case "AppVersion":
				m.AppVersion = new(string)
				if err := d.DecodeElement(m.AppVersion, &el); err != nil {
					return err
				}
			case "DocSecurity":
				m.DocSecurity = new(int32)
				if err := d.DecodeElement(m.DocSecurity, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Properties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Properties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Properties and its children
func (m *CT_Properties) Validate() error {
	return m.ValidateWithPath("CT_Properties")
}

// ValidateWithPath validates the CT_Properties and its children, prefixing error messages with path
func (m *CT_Properties) ValidateWithPath(path string) error {
	if m.HeadingPairs != nil {
		if err := m.HeadingPairs.ValidateWithPath(path + "/HeadingPairs"); err != nil {
			return err
		}
	}
	if m.TitlesOfParts != nil {
		if err := m.TitlesOfParts.ValidateWithPath(path + "/TitlesOfParts"); err != nil {
			return err
		}
	}
	if m.HLinks != nil {
		if err := m.HLinks.ValidateWithPath(path + "/HLinks"); err != nil {
			return err
		}
	}
	if m.DigSig != nil {
		if err := m.DigSig.ValidateWithPath(path + "/DigSig"); err != nil {
			return err
		}
	}
	return nil
}
