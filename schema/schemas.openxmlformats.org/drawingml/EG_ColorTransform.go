// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_ColorTransform struct {
	Tint     *CT_PositiveFixedPercentage
	Shade    *CT_PositiveFixedPercentage
	Comp     *CT_ComplementTransform
	Inv      *CT_InverseTransform
	Gray     *CT_GrayscaleTransform
	Alpha    *CT_PositiveFixedPercentage
	AlphaOff *CT_FixedPercentage
	AlphaMod *CT_PositivePercentage
	Hue      *CT_PositiveFixedAngle
	HueOff   *CT_Angle
	HueMod   *CT_PositivePercentage
	Sat      *CT_Percentage
	SatOff   *CT_Percentage
	SatMod   *CT_Percentage
	Lum      *CT_Percentage
	LumOff   *CT_Percentage
	LumMod   *CT_Percentage
	Red      *CT_Percentage
	RedOff   *CT_Percentage
	RedMod   *CT_Percentage
	Green    *CT_Percentage
	GreenOff *CT_Percentage
	GreenMod *CT_Percentage
	Blue     *CT_Percentage
	BlueOff  *CT_Percentage
	BlueMod  *CT_Percentage
	Gamma    *CT_GammaTransform
	InvGamma *CT_InverseGammaTransform
}

func NewEG_ColorTransform() *EG_ColorTransform {
	ret := &EG_ColorTransform{}
	return ret
}

func (m *EG_ColorTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Tint != nil {
		setint := xml.StartElement{Name: xml.Name{Local: "a:tint"}}
		e.EncodeElement(m.Tint, setint)
	}
	if m.Shade != nil {
		seshade := xml.StartElement{Name: xml.Name{Local: "a:shade"}}
		e.EncodeElement(m.Shade, seshade)
	}
	if m.Comp != nil {
		secomp := xml.StartElement{Name: xml.Name{Local: "a:comp"}}
		e.EncodeElement(m.Comp, secomp)
	}
	if m.Inv != nil {
		seinv := xml.StartElement{Name: xml.Name{Local: "a:inv"}}
		e.EncodeElement(m.Inv, seinv)
	}
	if m.Gray != nil {
		segray := xml.StartElement{Name: xml.Name{Local: "a:gray"}}
		e.EncodeElement(m.Gray, segray)
	}
	if m.Alpha != nil {
		sealpha := xml.StartElement{Name: xml.Name{Local: "a:alpha"}}
		e.EncodeElement(m.Alpha, sealpha)
	}
	if m.AlphaOff != nil {
		sealphaOff := xml.StartElement{Name: xml.Name{Local: "a:alphaOff"}}
		e.EncodeElement(m.AlphaOff, sealphaOff)
	}
	if m.AlphaMod != nil {
		sealphaMod := xml.StartElement{Name: xml.Name{Local: "a:alphaMod"}}
		e.EncodeElement(m.AlphaMod, sealphaMod)
	}
	if m.Hue != nil {
		sehue := xml.StartElement{Name: xml.Name{Local: "a:hue"}}
		e.EncodeElement(m.Hue, sehue)
	}
	if m.HueOff != nil {
		sehueOff := xml.StartElement{Name: xml.Name{Local: "a:hueOff"}}
		e.EncodeElement(m.HueOff, sehueOff)
	}
	if m.HueMod != nil {
		sehueMod := xml.StartElement{Name: xml.Name{Local: "a:hueMod"}}
		e.EncodeElement(m.HueMod, sehueMod)
	}
	if m.Sat != nil {
		sesat := xml.StartElement{Name: xml.Name{Local: "a:sat"}}
		e.EncodeElement(m.Sat, sesat)
	}
	if m.SatOff != nil {
		sesatOff := xml.StartElement{Name: xml.Name{Local: "a:satOff"}}
		e.EncodeElement(m.SatOff, sesatOff)
	}
	if m.SatMod != nil {
		sesatMod := xml.StartElement{Name: xml.Name{Local: "a:satMod"}}
		e.EncodeElement(m.SatMod, sesatMod)
	}
	if m.Lum != nil {
		selum := xml.StartElement{Name: xml.Name{Local: "a:lum"}}
		e.EncodeElement(m.Lum, selum)
	}
	if m.LumOff != nil {
		selumOff := xml.StartElement{Name: xml.Name{Local: "a:lumOff"}}
		e.EncodeElement(m.LumOff, selumOff)
	}
	if m.LumMod != nil {
		selumMod := xml.StartElement{Name: xml.Name{Local: "a:lumMod"}}
		e.EncodeElement(m.LumMod, selumMod)
	}
	if m.Red != nil {
		sered := xml.StartElement{Name: xml.Name{Local: "a:red"}}
		e.EncodeElement(m.Red, sered)
	}
	if m.RedOff != nil {
		seredOff := xml.StartElement{Name: xml.Name{Local: "a:redOff"}}
		e.EncodeElement(m.RedOff, seredOff)
	}
	if m.RedMod != nil {
		seredMod := xml.StartElement{Name: xml.Name{Local: "a:redMod"}}
		e.EncodeElement(m.RedMod, seredMod)
	}
	if m.Green != nil {
		segreen := xml.StartElement{Name: xml.Name{Local: "a:green"}}
		e.EncodeElement(m.Green, segreen)
	}
	if m.GreenOff != nil {
		segreenOff := xml.StartElement{Name: xml.Name{Local: "a:greenOff"}}
		e.EncodeElement(m.GreenOff, segreenOff)
	}
	if m.GreenMod != nil {
		segreenMod := xml.StartElement{Name: xml.Name{Local: "a:greenMod"}}
		e.EncodeElement(m.GreenMod, segreenMod)
	}
	if m.Blue != nil {
		seblue := xml.StartElement{Name: xml.Name{Local: "a:blue"}}
		e.EncodeElement(m.Blue, seblue)
	}
	if m.BlueOff != nil {
		seblueOff := xml.StartElement{Name: xml.Name{Local: "a:blueOff"}}
		e.EncodeElement(m.BlueOff, seblueOff)
	}
	if m.BlueMod != nil {
		seblueMod := xml.StartElement{Name: xml.Name{Local: "a:blueMod"}}
		e.EncodeElement(m.BlueMod, seblueMod)
	}
	if m.Gamma != nil {
		segamma := xml.StartElement{Name: xml.Name{Local: "a:gamma"}}
		e.EncodeElement(m.Gamma, segamma)
	}
	if m.InvGamma != nil {
		seinvGamma := xml.StartElement{Name: xml.Name{Local: "a:invGamma"}}
		e.EncodeElement(m.InvGamma, seinvGamma)
	}
	return nil
}

func (m *EG_ColorTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ColorTransform:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tint":
				m.Tint = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(m.Tint, &el); err != nil {
					return err
				}
			case "shade":
				m.Shade = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(m.Shade, &el); err != nil {
					return err
				}
			case "comp":
				m.Comp = NewCT_ComplementTransform()
				if err := d.DecodeElement(m.Comp, &el); err != nil {
					return err
				}
			case "inv":
				m.Inv = NewCT_InverseTransform()
				if err := d.DecodeElement(m.Inv, &el); err != nil {
					return err
				}
			case "gray":
				m.Gray = NewCT_GrayscaleTransform()
				if err := d.DecodeElement(m.Gray, &el); err != nil {
					return err
				}
			case "alpha":
				m.Alpha = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(m.Alpha, &el); err != nil {
					return err
				}
			case "alphaOff":
				m.AlphaOff = NewCT_FixedPercentage()
				if err := d.DecodeElement(m.AlphaOff, &el); err != nil {
					return err
				}
			case "alphaMod":
				m.AlphaMod = NewCT_PositivePercentage()
				if err := d.DecodeElement(m.AlphaMod, &el); err != nil {
					return err
				}
			case "hue":
				m.Hue = NewCT_PositiveFixedAngle()
				if err := d.DecodeElement(m.Hue, &el); err != nil {
					return err
				}
			case "hueOff":
				m.HueOff = NewCT_Angle()
				if err := d.DecodeElement(m.HueOff, &el); err != nil {
					return err
				}
			case "hueMod":
				m.HueMod = NewCT_PositivePercentage()
				if err := d.DecodeElement(m.HueMod, &el); err != nil {
					return err
				}
			case "sat":
				m.Sat = NewCT_Percentage()
				if err := d.DecodeElement(m.Sat, &el); err != nil {
					return err
				}
			case "satOff":
				m.SatOff = NewCT_Percentage()
				if err := d.DecodeElement(m.SatOff, &el); err != nil {
					return err
				}
			case "satMod":
				m.SatMod = NewCT_Percentage()
				if err := d.DecodeElement(m.SatMod, &el); err != nil {
					return err
				}
			case "lum":
				m.Lum = NewCT_Percentage()
				if err := d.DecodeElement(m.Lum, &el); err != nil {
					return err
				}
			case "lumOff":
				m.LumOff = NewCT_Percentage()
				if err := d.DecodeElement(m.LumOff, &el); err != nil {
					return err
				}
			case "lumMod":
				m.LumMod = NewCT_Percentage()
				if err := d.DecodeElement(m.LumMod, &el); err != nil {
					return err
				}
			case "red":
				m.Red = NewCT_Percentage()
				if err := d.DecodeElement(m.Red, &el); err != nil {
					return err
				}
			case "redOff":
				m.RedOff = NewCT_Percentage()
				if err := d.DecodeElement(m.RedOff, &el); err != nil {
					return err
				}
			case "redMod":
				m.RedMod = NewCT_Percentage()
				if err := d.DecodeElement(m.RedMod, &el); err != nil {
					return err
				}
			case "green":
				m.Green = NewCT_Percentage()
				if err := d.DecodeElement(m.Green, &el); err != nil {
					return err
				}
			case "greenOff":
				m.GreenOff = NewCT_Percentage()
				if err := d.DecodeElement(m.GreenOff, &el); err != nil {
					return err
				}
			case "greenMod":
				m.GreenMod = NewCT_Percentage()
				if err := d.DecodeElement(m.GreenMod, &el); err != nil {
					return err
				}
			case "blue":
				m.Blue = NewCT_Percentage()
				if err := d.DecodeElement(m.Blue, &el); err != nil {
					return err
				}
			case "blueOff":
				m.BlueOff = NewCT_Percentage()
				if err := d.DecodeElement(m.BlueOff, &el); err != nil {
					return err
				}
			case "blueMod":
				m.BlueMod = NewCT_Percentage()
				if err := d.DecodeElement(m.BlueMod, &el); err != nil {
					return err
				}
			case "gamma":
				m.Gamma = NewCT_GammaTransform()
				if err := d.DecodeElement(m.Gamma, &el); err != nil {
					return err
				}
			case "invGamma":
				m.InvGamma = NewCT_InverseGammaTransform()
				if err := d.DecodeElement(m.InvGamma, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_ColorTransform %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ColorTransform
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ColorTransform and its children
func (m *EG_ColorTransform) Validate() error {
	return m.ValidateWithPath("EG_ColorTransform")
}

// ValidateWithPath validates the EG_ColorTransform and its children, prefixing error messages with path
func (m *EG_ColorTransform) ValidateWithPath(path string) error {
	if m.Tint != nil {
		if err := m.Tint.ValidateWithPath(path + "/Tint"); err != nil {
			return err
		}
	}
	if m.Shade != nil {
		if err := m.Shade.ValidateWithPath(path + "/Shade"); err != nil {
			return err
		}
	}
	if m.Comp != nil {
		if err := m.Comp.ValidateWithPath(path + "/Comp"); err != nil {
			return err
		}
	}
	if m.Inv != nil {
		if err := m.Inv.ValidateWithPath(path + "/Inv"); err != nil {
			return err
		}
	}
	if m.Gray != nil {
		if err := m.Gray.ValidateWithPath(path + "/Gray"); err != nil {
			return err
		}
	}
	if m.Alpha != nil {
		if err := m.Alpha.ValidateWithPath(path + "/Alpha"); err != nil {
			return err
		}
	}
	if m.AlphaOff != nil {
		if err := m.AlphaOff.ValidateWithPath(path + "/AlphaOff"); err != nil {
			return err
		}
	}
	if m.AlphaMod != nil {
		if err := m.AlphaMod.ValidateWithPath(path + "/AlphaMod"); err != nil {
			return err
		}
	}
	if m.Hue != nil {
		if err := m.Hue.ValidateWithPath(path + "/Hue"); err != nil {
			return err
		}
	}
	if m.HueOff != nil {
		if err := m.HueOff.ValidateWithPath(path + "/HueOff"); err != nil {
			return err
		}
	}
	if m.HueMod != nil {
		if err := m.HueMod.ValidateWithPath(path + "/HueMod"); err != nil {
			return err
		}
	}
	if m.Sat != nil {
		if err := m.Sat.ValidateWithPath(path + "/Sat"); err != nil {
			return err
		}
	}
	if m.SatOff != nil {
		if err := m.SatOff.ValidateWithPath(path + "/SatOff"); err != nil {
			return err
		}
	}
	if m.SatMod != nil {
		if err := m.SatMod.ValidateWithPath(path + "/SatMod"); err != nil {
			return err
		}
	}
	if m.Lum != nil {
		if err := m.Lum.ValidateWithPath(path + "/Lum"); err != nil {
			return err
		}
	}
	if m.LumOff != nil {
		if err := m.LumOff.ValidateWithPath(path + "/LumOff"); err != nil {
			return err
		}
	}
	if m.LumMod != nil {
		if err := m.LumMod.ValidateWithPath(path + "/LumMod"); err != nil {
			return err
		}
	}
	if m.Red != nil {
		if err := m.Red.ValidateWithPath(path + "/Red"); err != nil {
			return err
		}
	}
	if m.RedOff != nil {
		if err := m.RedOff.ValidateWithPath(path + "/RedOff"); err != nil {
			return err
		}
	}
	if m.RedMod != nil {
		if err := m.RedMod.ValidateWithPath(path + "/RedMod"); err != nil {
			return err
		}
	}
	if m.Green != nil {
		if err := m.Green.ValidateWithPath(path + "/Green"); err != nil {
			return err
		}
	}
	if m.GreenOff != nil {
		if err := m.GreenOff.ValidateWithPath(path + "/GreenOff"); err != nil {
			return err
		}
	}
	if m.GreenMod != nil {
		if err := m.GreenMod.ValidateWithPath(path + "/GreenMod"); err != nil {
			return err
		}
	}
	if m.Blue != nil {
		if err := m.Blue.ValidateWithPath(path + "/Blue"); err != nil {
			return err
		}
	}
	if m.BlueOff != nil {
		if err := m.BlueOff.ValidateWithPath(path + "/BlueOff"); err != nil {
			return err
		}
	}
	if m.BlueMod != nil {
		if err := m.BlueMod.ValidateWithPath(path + "/BlueMod"); err != nil {
			return err
		}
	}
	if m.Gamma != nil {
		if err := m.Gamma.ValidateWithPath(path + "/Gamma"); err != nil {
			return err
		}
	}
	if m.InvGamma != nil {
		if err := m.InvGamma.ValidateWithPath(path + "/InvGamma"); err != nil {
			return err
		}
	}
	return nil
}
