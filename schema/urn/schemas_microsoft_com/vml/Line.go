// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/urn/schemas_microsoft_com/office/excel"
	"baliance.com/gooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
	"baliance.com/gooxml/schema/urn/schemas_microsoft_com/office/word"
)

type Line struct {
	CT_Line
}

func NewLine() *Line {
	ret := &Line{}
	ret.CT_Line = *NewCT_Line()
	return ret
}

func (m *Line) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Line.MarshalXML(e, start)
}

func (m *Line) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Line = *NewCT_Line()
	for _, attr := range start.Attr {
		if attr.Name.Local == "from" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FromAttr = &parsed
		}
		if attr.Name.Local == "to" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ToAttr = &parsed
		}
		if attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
		}
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
		}
		if attr.Name.Local == "class" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ClassAttr = &parsed
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
		}
		if attr.Name.Local == "alt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltAttr = &parsed
		}
		if attr.Name.Local == "coordsize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordsizeAttr = &parsed
		}
		if attr.Name.Local == "coordorigin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordoriginAttr = &parsed
		}
		if attr.Name.Local == "wrapcoords" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WrapcoordsAttr = &parsed
		}
		if attr.Name.Local == "print" {
			m.PrintAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
		}
		if attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
		}
		if attr.Name.Local == "oned" {
			m.OnedAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "regroupid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RegroupidAttr = &parsed
		}
		if attr.Name.Local == "doubleclicknotify" {
			m.DoubleclicknotifyAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "button" {
			m.ButtonAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "userhidden" {
			m.UserhiddenAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "bullet" {
			m.BulletAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hr" {
			m.HrAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hrstd" {
			m.HrstdAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hrnoshade" {
			m.HrnoshadeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "hrpct" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.HrpctAttr = &pt
		}
		if attr.Name.Local == "hralign" {
			m.HralignAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "allowincell" {
			m.AllowincellAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "allowoverlap" {
			m.AllowoverlapAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "userdrawn" {
			m.UserdrawnAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "bordertopcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BordertopcolorAttr = &parsed
		}
		if attr.Name.Local == "borderleftcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderleftcolorAttr = &parsed
		}
		if attr.Name.Local == "borderbottomcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderbottomcolorAttr = &parsed
		}
		if attr.Name.Local == "borderrightcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderrightcolorAttr = &parsed
		}
		if attr.Name.Local == "dgmlayout" {
			m.DgmlayoutAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "dgmnodekind" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmnodekindAttr = &parsed
		}
		if attr.Name.Local == "dgmlayoutmru" {
			m.DgmlayoutmruAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "insetmode" {
			m.InsetmodeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
		}
		if attr.Name.Local == "stroked" {
			m.StrokedAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
		}
		if attr.Name.Local == "strokeweight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokeweightAttr = &parsed
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "chromakey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ChromakeyAttr = &parsed
		}
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
		}
		if attr.Name.Local == "spt" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.SptAttr = &pt
		}
		if attr.Name.Local == "connectortype" {
			m.ConnectortypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "bwmode" {
			m.BwmodeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "bwpure" {
			m.BwpureAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "bwnormal" {
			m.BwnormalAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "oleicon" {
			m.OleiconAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ole" {
			m.OleAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "preferrelative" {
			m.PreferrelativeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cliptowrap" {
			m.CliptowrapAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "clip" {
			m.ClipAttr.UnmarshalXMLAttr(attr)
		}
	}
lLine:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "path"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Path = NewPath()
				if err := d.DecodeElement(tmpshapeelements.Path, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "formulas"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Formulas = NewFormulas()
				if err := d.DecodeElement(tmpshapeelements.Formulas, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "handles"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Handles = NewHandles()
				if err := d.DecodeElement(tmpshapeelements.Handles, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Fill = NewFill()
				if err := d.DecodeElement(tmpshapeelements.Fill, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Stroke = NewStroke()
				if err := d.DecodeElement(tmpshapeelements.Stroke, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shadow"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Shadow = NewShadow()
				if err := d.DecodeElement(tmpshapeelements.Shadow, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textbox"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textbox = NewTextbox()
				if err := d.DecodeElement(tmpshapeelements.Textbox, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textpath"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textpath = NewTextpath()
				if err := d.DecodeElement(tmpshapeelements.Textpath, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "imagedata"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Imagedata = NewImagedata()
				if err := d.DecodeElement(tmpshapeelements.Imagedata, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "skew"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Skew = NewOfcSkew()
				if err := d.DecodeElement(tmpshapeelements.Skew, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "extrusion"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Extrusion = NewOfcExtrusion()
				if err := d.DecodeElement(tmpshapeelements.Extrusion, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "callout"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Callout = NewOfcCallout()
				if err := d.DecodeElement(tmpshapeelements.Callout, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Lock = NewOfcLock()
				if err := d.DecodeElement(tmpshapeelements.Lock, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "clippath"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Clippath = NewOfcClippath()
				if err := d.DecodeElement(tmpshapeelements.Clippath, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "signatureline"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Signatureline = NewOfcSignatureline()
				if err := d.DecodeElement(tmpshapeelements.Signatureline, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "wrap"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Wrap = word.NewWrap()
				if err := d.DecodeElement(tmpshapeelements.Wrap, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "anchorlock"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Anchorlock = word.NewAnchorlock()
				if err := d.DecodeElement(tmpshapeelements.Anchorlock, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "bordertop"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Bordertop = word.NewBordertop()
				if err := d.DecodeElement(tmpshapeelements.Bordertop, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderbottom"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderbottom = word.NewBorderbottom()
				if err := d.DecodeElement(tmpshapeelements.Borderbottom, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderleft"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderleft = word.NewBorderleft()
				if err := d.DecodeElement(tmpshapeelements.Borderleft, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderright"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderright = word.NewBorderright()
				if err := d.DecodeElement(tmpshapeelements.Borderright, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ClientData"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.ClientData = excel.NewClientData()
				if err := d.DecodeElement(tmpshapeelements.ClientData, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:powerpoint", Local: "textdata"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textdata = powerpoint.NewTextdata()
				if err := d.DecodeElement(tmpshapeelements.Textdata, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			default:
				log.Printf("skipping unsupported element on Line %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lLine
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Line and its children
func (m *Line) Validate() error {
	return m.ValidateWithPath("Line")
}

// ValidateWithPath validates the Line and its children, prefixing error messages with path
func (m *Line) ValidateWithPath(path string) error {
	if err := m.CT_Line.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}