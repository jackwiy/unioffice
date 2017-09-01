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

type EG_Effect struct {
	Cont         *CT_EffectContainer
	Effect       *CT_EffectReference
	AlphaBiLevel *CT_AlphaBiLevelEffect
	AlphaCeiling *CT_AlphaCeilingEffect
	AlphaFloor   *CT_AlphaFloorEffect
	AlphaInv     *CT_AlphaInverseEffect
	AlphaMod     *CT_AlphaModulateEffect
	AlphaModFix  *CT_AlphaModulateFixedEffect
	AlphaOutset  *CT_AlphaOutsetEffect
	AlphaRepl    *CT_AlphaReplaceEffect
	BiLevel      *CT_BiLevelEffect
	Blend        *CT_BlendEffect
	Blur         *CT_BlurEffect
	ClrChange    *CT_ColorChangeEffect
	ClrRepl      *CT_ColorReplaceEffect
	Duotone      *CT_DuotoneEffect
	Fill         *CT_FillEffect
	FillOverlay  *CT_FillOverlayEffect
	Glow         *CT_GlowEffect
	Grayscl      *CT_GrayscaleEffect
	Hsl          *CT_HSLEffect
	InnerShdw    *CT_InnerShadowEffect
	Lum          *CT_LuminanceEffect
	OuterShdw    *CT_OuterShadowEffect
	PrstShdw     *CT_PresetShadowEffect
	Reflection   *CT_ReflectionEffect
	RelOff       *CT_RelativeOffsetEffect
	SoftEdge     *CT_SoftEdgesEffect
	Tint         *CT_TintEffect
	Xfrm         *CT_TransformEffect
}

func NewEG_Effect() *EG_Effect {
	ret := &EG_Effect{}
	return ret
}

func (m *EG_Effect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Cont != nil {
		secont := xml.StartElement{Name: xml.Name{Local: "a:cont"}}
		e.EncodeElement(m.Cont, secont)
	}
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "a:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.AlphaBiLevel != nil {
		sealphaBiLevel := xml.StartElement{Name: xml.Name{Local: "a:alphaBiLevel"}}
		e.EncodeElement(m.AlphaBiLevel, sealphaBiLevel)
	}
	if m.AlphaCeiling != nil {
		sealphaCeiling := xml.StartElement{Name: xml.Name{Local: "a:alphaCeiling"}}
		e.EncodeElement(m.AlphaCeiling, sealphaCeiling)
	}
	if m.AlphaFloor != nil {
		sealphaFloor := xml.StartElement{Name: xml.Name{Local: "a:alphaFloor"}}
		e.EncodeElement(m.AlphaFloor, sealphaFloor)
	}
	if m.AlphaInv != nil {
		sealphaInv := xml.StartElement{Name: xml.Name{Local: "a:alphaInv"}}
		e.EncodeElement(m.AlphaInv, sealphaInv)
	}
	if m.AlphaMod != nil {
		sealphaMod := xml.StartElement{Name: xml.Name{Local: "a:alphaMod"}}
		e.EncodeElement(m.AlphaMod, sealphaMod)
	}
	if m.AlphaModFix != nil {
		sealphaModFix := xml.StartElement{Name: xml.Name{Local: "a:alphaModFix"}}
		e.EncodeElement(m.AlphaModFix, sealphaModFix)
	}
	if m.AlphaOutset != nil {
		sealphaOutset := xml.StartElement{Name: xml.Name{Local: "a:alphaOutset"}}
		e.EncodeElement(m.AlphaOutset, sealphaOutset)
	}
	if m.AlphaRepl != nil {
		sealphaRepl := xml.StartElement{Name: xml.Name{Local: "a:alphaRepl"}}
		e.EncodeElement(m.AlphaRepl, sealphaRepl)
	}
	if m.BiLevel != nil {
		sebiLevel := xml.StartElement{Name: xml.Name{Local: "a:biLevel"}}
		e.EncodeElement(m.BiLevel, sebiLevel)
	}
	if m.Blend != nil {
		seblend := xml.StartElement{Name: xml.Name{Local: "a:blend"}}
		e.EncodeElement(m.Blend, seblend)
	}
	if m.Blur != nil {
		seblur := xml.StartElement{Name: xml.Name{Local: "a:blur"}}
		e.EncodeElement(m.Blur, seblur)
	}
	if m.ClrChange != nil {
		seclrChange := xml.StartElement{Name: xml.Name{Local: "a:clrChange"}}
		e.EncodeElement(m.ClrChange, seclrChange)
	}
	if m.ClrRepl != nil {
		seclrRepl := xml.StartElement{Name: xml.Name{Local: "a:clrRepl"}}
		e.EncodeElement(m.ClrRepl, seclrRepl)
	}
	if m.Duotone != nil {
		seduotone := xml.StartElement{Name: xml.Name{Local: "a:duotone"}}
		e.EncodeElement(m.Duotone, seduotone)
	}
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "a:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.FillOverlay != nil {
		sefillOverlay := xml.StartElement{Name: xml.Name{Local: "a:fillOverlay"}}
		e.EncodeElement(m.FillOverlay, sefillOverlay)
	}
	if m.Glow != nil {
		seglow := xml.StartElement{Name: xml.Name{Local: "a:glow"}}
		e.EncodeElement(m.Glow, seglow)
	}
	if m.Grayscl != nil {
		segrayscl := xml.StartElement{Name: xml.Name{Local: "a:grayscl"}}
		e.EncodeElement(m.Grayscl, segrayscl)
	}
	if m.Hsl != nil {
		sehsl := xml.StartElement{Name: xml.Name{Local: "a:hsl"}}
		e.EncodeElement(m.Hsl, sehsl)
	}
	if m.InnerShdw != nil {
		seinnerShdw := xml.StartElement{Name: xml.Name{Local: "a:innerShdw"}}
		e.EncodeElement(m.InnerShdw, seinnerShdw)
	}
	if m.Lum != nil {
		selum := xml.StartElement{Name: xml.Name{Local: "a:lum"}}
		e.EncodeElement(m.Lum, selum)
	}
	if m.OuterShdw != nil {
		seouterShdw := xml.StartElement{Name: xml.Name{Local: "a:outerShdw"}}
		e.EncodeElement(m.OuterShdw, seouterShdw)
	}
	if m.PrstShdw != nil {
		seprstShdw := xml.StartElement{Name: xml.Name{Local: "a:prstShdw"}}
		e.EncodeElement(m.PrstShdw, seprstShdw)
	}
	if m.Reflection != nil {
		sereflection := xml.StartElement{Name: xml.Name{Local: "a:reflection"}}
		e.EncodeElement(m.Reflection, sereflection)
	}
	if m.RelOff != nil {
		serelOff := xml.StartElement{Name: xml.Name{Local: "a:relOff"}}
		e.EncodeElement(m.RelOff, serelOff)
	}
	if m.SoftEdge != nil {
		sesoftEdge := xml.StartElement{Name: xml.Name{Local: "a:softEdge"}}
		e.EncodeElement(m.SoftEdge, sesoftEdge)
	}
	if m.Tint != nil {
		setint := xml.StartElement{Name: xml.Name{Local: "a:tint"}}
		e.EncodeElement(m.Tint, setint)
	}
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	return nil
}

func (m *EG_Effect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Effect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cont":
				m.Cont = NewCT_EffectContainer()
				if err := d.DecodeElement(m.Cont, &el); err != nil {
					return err
				}
			case "effect":
				m.Effect = NewCT_EffectReference()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case "alphaBiLevel":
				m.AlphaBiLevel = NewCT_AlphaBiLevelEffect()
				if err := d.DecodeElement(m.AlphaBiLevel, &el); err != nil {
					return err
				}
			case "alphaCeiling":
				m.AlphaCeiling = NewCT_AlphaCeilingEffect()
				if err := d.DecodeElement(m.AlphaCeiling, &el); err != nil {
					return err
				}
			case "alphaFloor":
				m.AlphaFloor = NewCT_AlphaFloorEffect()
				if err := d.DecodeElement(m.AlphaFloor, &el); err != nil {
					return err
				}
			case "alphaInv":
				m.AlphaInv = NewCT_AlphaInverseEffect()
				if err := d.DecodeElement(m.AlphaInv, &el); err != nil {
					return err
				}
			case "alphaMod":
				m.AlphaMod = NewCT_AlphaModulateEffect()
				if err := d.DecodeElement(m.AlphaMod, &el); err != nil {
					return err
				}
			case "alphaModFix":
				m.AlphaModFix = NewCT_AlphaModulateFixedEffect()
				if err := d.DecodeElement(m.AlphaModFix, &el); err != nil {
					return err
				}
			case "alphaOutset":
				m.AlphaOutset = NewCT_AlphaOutsetEffect()
				if err := d.DecodeElement(m.AlphaOutset, &el); err != nil {
					return err
				}
			case "alphaRepl":
				m.AlphaRepl = NewCT_AlphaReplaceEffect()
				if err := d.DecodeElement(m.AlphaRepl, &el); err != nil {
					return err
				}
			case "biLevel":
				m.BiLevel = NewCT_BiLevelEffect()
				if err := d.DecodeElement(m.BiLevel, &el); err != nil {
					return err
				}
			case "blend":
				m.Blend = NewCT_BlendEffect()
				if err := d.DecodeElement(m.Blend, &el); err != nil {
					return err
				}
			case "blur":
				m.Blur = NewCT_BlurEffect()
				if err := d.DecodeElement(m.Blur, &el); err != nil {
					return err
				}
			case "clrChange":
				m.ClrChange = NewCT_ColorChangeEffect()
				if err := d.DecodeElement(m.ClrChange, &el); err != nil {
					return err
				}
			case "clrRepl":
				m.ClrRepl = NewCT_ColorReplaceEffect()
				if err := d.DecodeElement(m.ClrRepl, &el); err != nil {
					return err
				}
			case "duotone":
				m.Duotone = NewCT_DuotoneEffect()
				if err := d.DecodeElement(m.Duotone, &el); err != nil {
					return err
				}
			case "fill":
				m.Fill = NewCT_FillEffect()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case "fillOverlay":
				m.FillOverlay = NewCT_FillOverlayEffect()
				if err := d.DecodeElement(m.FillOverlay, &el); err != nil {
					return err
				}
			case "glow":
				m.Glow = NewCT_GlowEffect()
				if err := d.DecodeElement(m.Glow, &el); err != nil {
					return err
				}
			case "grayscl":
				m.Grayscl = NewCT_GrayscaleEffect()
				if err := d.DecodeElement(m.Grayscl, &el); err != nil {
					return err
				}
			case "hsl":
				m.Hsl = NewCT_HSLEffect()
				if err := d.DecodeElement(m.Hsl, &el); err != nil {
					return err
				}
			case "innerShdw":
				m.InnerShdw = NewCT_InnerShadowEffect()
				if err := d.DecodeElement(m.InnerShdw, &el); err != nil {
					return err
				}
			case "lum":
				m.Lum = NewCT_LuminanceEffect()
				if err := d.DecodeElement(m.Lum, &el); err != nil {
					return err
				}
			case "outerShdw":
				m.OuterShdw = NewCT_OuterShadowEffect()
				if err := d.DecodeElement(m.OuterShdw, &el); err != nil {
					return err
				}
			case "prstShdw":
				m.PrstShdw = NewCT_PresetShadowEffect()
				if err := d.DecodeElement(m.PrstShdw, &el); err != nil {
					return err
				}
			case "reflection":
				m.Reflection = NewCT_ReflectionEffect()
				if err := d.DecodeElement(m.Reflection, &el); err != nil {
					return err
				}
			case "relOff":
				m.RelOff = NewCT_RelativeOffsetEffect()
				if err := d.DecodeElement(m.RelOff, &el); err != nil {
					return err
				}
			case "softEdge":
				m.SoftEdge = NewCT_SoftEdgesEffect()
				if err := d.DecodeElement(m.SoftEdge, &el); err != nil {
					return err
				}
			case "tint":
				m.Tint = NewCT_TintEffect()
				if err := d.DecodeElement(m.Tint, &el); err != nil {
					return err
				}
			case "xfrm":
				m.Xfrm = NewCT_TransformEffect()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_Effect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Effect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Effect and its children
func (m *EG_Effect) Validate() error {
	return m.ValidateWithPath("EG_Effect")
}

// ValidateWithPath validates the EG_Effect and its children, prefixing error messages with path
func (m *EG_Effect) ValidateWithPath(path string) error {
	if m.Cont != nil {
		if err := m.Cont.ValidateWithPath(path + "/Cont"); err != nil {
			return err
		}
	}
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.AlphaBiLevel != nil {
		if err := m.AlphaBiLevel.ValidateWithPath(path + "/AlphaBiLevel"); err != nil {
			return err
		}
	}
	if m.AlphaCeiling != nil {
		if err := m.AlphaCeiling.ValidateWithPath(path + "/AlphaCeiling"); err != nil {
			return err
		}
	}
	if m.AlphaFloor != nil {
		if err := m.AlphaFloor.ValidateWithPath(path + "/AlphaFloor"); err != nil {
			return err
		}
	}
	if m.AlphaInv != nil {
		if err := m.AlphaInv.ValidateWithPath(path + "/AlphaInv"); err != nil {
			return err
		}
	}
	if m.AlphaMod != nil {
		if err := m.AlphaMod.ValidateWithPath(path + "/AlphaMod"); err != nil {
			return err
		}
	}
	if m.AlphaModFix != nil {
		if err := m.AlphaModFix.ValidateWithPath(path + "/AlphaModFix"); err != nil {
			return err
		}
	}
	if m.AlphaOutset != nil {
		if err := m.AlphaOutset.ValidateWithPath(path + "/AlphaOutset"); err != nil {
			return err
		}
	}
	if m.AlphaRepl != nil {
		if err := m.AlphaRepl.ValidateWithPath(path + "/AlphaRepl"); err != nil {
			return err
		}
	}
	if m.BiLevel != nil {
		if err := m.BiLevel.ValidateWithPath(path + "/BiLevel"); err != nil {
			return err
		}
	}
	if m.Blend != nil {
		if err := m.Blend.ValidateWithPath(path + "/Blend"); err != nil {
			return err
		}
	}
	if m.Blur != nil {
		if err := m.Blur.ValidateWithPath(path + "/Blur"); err != nil {
			return err
		}
	}
	if m.ClrChange != nil {
		if err := m.ClrChange.ValidateWithPath(path + "/ClrChange"); err != nil {
			return err
		}
	}
	if m.ClrRepl != nil {
		if err := m.ClrRepl.ValidateWithPath(path + "/ClrRepl"); err != nil {
			return err
		}
	}
	if m.Duotone != nil {
		if err := m.Duotone.ValidateWithPath(path + "/Duotone"); err != nil {
			return err
		}
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.FillOverlay != nil {
		if err := m.FillOverlay.ValidateWithPath(path + "/FillOverlay"); err != nil {
			return err
		}
	}
	if m.Glow != nil {
		if err := m.Glow.ValidateWithPath(path + "/Glow"); err != nil {
			return err
		}
	}
	if m.Grayscl != nil {
		if err := m.Grayscl.ValidateWithPath(path + "/Grayscl"); err != nil {
			return err
		}
	}
	if m.Hsl != nil {
		if err := m.Hsl.ValidateWithPath(path + "/Hsl"); err != nil {
			return err
		}
	}
	if m.InnerShdw != nil {
		if err := m.InnerShdw.ValidateWithPath(path + "/InnerShdw"); err != nil {
			return err
		}
	}
	if m.Lum != nil {
		if err := m.Lum.ValidateWithPath(path + "/Lum"); err != nil {
			return err
		}
	}
	if m.OuterShdw != nil {
		if err := m.OuterShdw.ValidateWithPath(path + "/OuterShdw"); err != nil {
			return err
		}
	}
	if m.PrstShdw != nil {
		if err := m.PrstShdw.ValidateWithPath(path + "/PrstShdw"); err != nil {
			return err
		}
	}
	if m.Reflection != nil {
		if err := m.Reflection.ValidateWithPath(path + "/Reflection"); err != nil {
			return err
		}
	}
	if m.RelOff != nil {
		if err := m.RelOff.ValidateWithPath(path + "/RelOff"); err != nil {
			return err
		}
	}
	if m.SoftEdge != nil {
		if err := m.SoftEdge.ValidateWithPath(path + "/SoftEdge"); err != nil {
			return err
		}
	}
	if m.Tint != nil {
		if err := m.Tint.ValidateWithPath(path + "/Tint"); err != nil {
			return err
		}
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
			return err
		}
	}
	return nil
}
