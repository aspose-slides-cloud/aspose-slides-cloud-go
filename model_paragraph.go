/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package asposeslidescloud
import (
	"encoding/json"
)

// Represents paragraph resource
type IParagraph interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Left margin.
	GetMarginLeft() float64
	SetMarginLeft(newValue float64)

	// Right margin.
	GetMarginRight() float64
	SetMarginRight(newValue float64)

	// Left spacing.
	GetSpaceBefore() float64
	SetSpaceBefore(newValue float64)

	// Right spacing.
	GetSpaceAfter() float64
	SetSpaceAfter(newValue float64)

	// Spacing between lines.
	GetSpaceWithin() float64
	SetSpaceWithin(newValue float64)

	// First line indent.
	GetIndent() float64
	SetIndent(newValue float64)

	// Text alignment.
	GetAlignment() string
	SetAlignment(newValue string)

	// Font alignment.
	GetFontAlignment() string
	SetFontAlignment(newValue string)

	// Default tabulation size.
	GetDefaultTabSize() float64
	SetDefaultTabSize(newValue float64)

	// Depth.
	GetDepth() int32
	SetDepth(newValue int32)

	// True if hanging punctuation is used with the paragraph.
	GetHangingPunctuation() string
	SetHangingPunctuation(newValue string)

	// True if East Asian line break is used with the paragraph.
	GetEastAsianLineBreak() string
	SetEastAsianLineBreak(newValue string)

	// True if Latin line break is used with the paragraph.
	GetLatinLineBreak() string
	SetLatinLineBreak(newValue string)

	// True if right to left direction is used with the paragraph.
	GetRightToLeft() string
	SetRightToLeft(newValue string)

	// List of portion links.
	GetPortionList() []IPortion
	SetPortionList(newValue []IPortion)

	// Default portion format.
	GetDefaultPortionFormat() IPortionFormat
	SetDefaultPortionFormat(newValue IPortionFormat)

	// Bullet char.
	GetBulletChar() string
	SetBulletChar(newValue string)

	// Bullet height.
	GetBulletHeight() float64
	SetBulletHeight(newValue float64)

	// Bullet type.
	GetBulletType() string
	SetBulletType(newValue string)

	// Starting number for a numbered bullet.
	GetNumberedBulletStartWith() int32
	SetNumberedBulletStartWith(newValue int32)

	// Numbered bullet style.
	GetNumberedBulletStyle() string
	SetNumberedBulletStyle(newValue string)

	// Bullet fill format.
	GetBulletFillFormat() IFillFormat
	SetBulletFillFormat(newValue IFillFormat)
}

type Paragraph struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Left margin.
	MarginLeft float64 `json:"MarginLeft,omitempty"`

	// Right margin.
	MarginRight float64 `json:"MarginRight,omitempty"`

	// Left spacing.
	SpaceBefore float64 `json:"SpaceBefore,omitempty"`

	// Right spacing.
	SpaceAfter float64 `json:"SpaceAfter,omitempty"`

	// Spacing between lines.
	SpaceWithin float64 `json:"SpaceWithin,omitempty"`

	// First line indent.
	Indent float64 `json:"Indent,omitempty"`

	// Text alignment.
	Alignment string `json:"Alignment,omitempty"`

	// Font alignment.
	FontAlignment string `json:"FontAlignment,omitempty"`

	// Default tabulation size.
	DefaultTabSize float64 `json:"DefaultTabSize,omitempty"`

	// Depth.
	Depth int32 `json:"Depth,omitempty"`

	// True if hanging punctuation is used with the paragraph.
	HangingPunctuation string `json:"HangingPunctuation,omitempty"`

	// True if East Asian line break is used with the paragraph.
	EastAsianLineBreak string `json:"EastAsianLineBreak,omitempty"`

	// True if Latin line break is used with the paragraph.
	LatinLineBreak string `json:"LatinLineBreak,omitempty"`

	// True if right to left direction is used with the paragraph.
	RightToLeft string `json:"RightToLeft,omitempty"`

	// List of portion links.
	PortionList []IPortion `json:"PortionList,omitempty"`

	// Default portion format.
	DefaultPortionFormat IPortionFormat `json:"DefaultPortionFormat,omitempty"`

	// Bullet char.
	BulletChar string `json:"BulletChar,omitempty"`

	// Bullet height.
	BulletHeight float64 `json:"BulletHeight,omitempty"`

	// Bullet type.
	BulletType string `json:"BulletType,omitempty"`

	// Starting number for a numbered bullet.
	NumberedBulletStartWith int32 `json:"NumberedBulletStartWith,omitempty"`

	// Numbered bullet style.
	NumberedBulletStyle string `json:"NumberedBulletStyle,omitempty"`

	// Bullet fill format.
	BulletFillFormat IFillFormat `json:"BulletFillFormat,omitempty"`
}

func NewParagraph() *Paragraph {
	instance := new(Paragraph)
	return instance
}

func (this *Paragraph) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Paragraph) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Paragraph) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Paragraph) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Paragraph) GetMarginLeft() float64 {
	return this.MarginLeft
}

func (this *Paragraph) SetMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this *Paragraph) GetMarginRight() float64 {
	return this.MarginRight
}

func (this *Paragraph) SetMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this *Paragraph) GetSpaceBefore() float64 {
	return this.SpaceBefore
}

func (this *Paragraph) SetSpaceBefore(newValue float64) {
	this.SpaceBefore = newValue
}
func (this *Paragraph) GetSpaceAfter() float64 {
	return this.SpaceAfter
}

func (this *Paragraph) SetSpaceAfter(newValue float64) {
	this.SpaceAfter = newValue
}
func (this *Paragraph) GetSpaceWithin() float64 {
	return this.SpaceWithin
}

func (this *Paragraph) SetSpaceWithin(newValue float64) {
	this.SpaceWithin = newValue
}
func (this *Paragraph) GetIndent() float64 {
	return this.Indent
}

func (this *Paragraph) SetIndent(newValue float64) {
	this.Indent = newValue
}
func (this *Paragraph) GetAlignment() string {
	return this.Alignment
}

func (this *Paragraph) SetAlignment(newValue string) {
	this.Alignment = newValue
}
func (this *Paragraph) GetFontAlignment() string {
	return this.FontAlignment
}

func (this *Paragraph) SetFontAlignment(newValue string) {
	this.FontAlignment = newValue
}
func (this *Paragraph) GetDefaultTabSize() float64 {
	return this.DefaultTabSize
}

func (this *Paragraph) SetDefaultTabSize(newValue float64) {
	this.DefaultTabSize = newValue
}
func (this *Paragraph) GetDepth() int32 {
	return this.Depth
}

func (this *Paragraph) SetDepth(newValue int32) {
	this.Depth = newValue
}
func (this *Paragraph) GetHangingPunctuation() string {
	return this.HangingPunctuation
}

func (this *Paragraph) SetHangingPunctuation(newValue string) {
	this.HangingPunctuation = newValue
}
func (this *Paragraph) GetEastAsianLineBreak() string {
	return this.EastAsianLineBreak
}

func (this *Paragraph) SetEastAsianLineBreak(newValue string) {
	this.EastAsianLineBreak = newValue
}
func (this *Paragraph) GetLatinLineBreak() string {
	return this.LatinLineBreak
}

func (this *Paragraph) SetLatinLineBreak(newValue string) {
	this.LatinLineBreak = newValue
}
func (this *Paragraph) GetRightToLeft() string {
	return this.RightToLeft
}

func (this *Paragraph) SetRightToLeft(newValue string) {
	this.RightToLeft = newValue
}
func (this *Paragraph) GetPortionList() []IPortion {
	return this.PortionList
}

func (this *Paragraph) SetPortionList(newValue []IPortion) {
	this.PortionList = newValue
}
func (this *Paragraph) GetDefaultPortionFormat() IPortionFormat {
	return this.DefaultPortionFormat
}

func (this *Paragraph) SetDefaultPortionFormat(newValue IPortionFormat) {
	this.DefaultPortionFormat = newValue
}
func (this *Paragraph) GetBulletChar() string {
	return this.BulletChar
}

func (this *Paragraph) SetBulletChar(newValue string) {
	this.BulletChar = newValue
}
func (this *Paragraph) GetBulletHeight() float64 {
	return this.BulletHeight
}

func (this *Paragraph) SetBulletHeight(newValue float64) {
	this.BulletHeight = newValue
}
func (this *Paragraph) GetBulletType() string {
	return this.BulletType
}

func (this *Paragraph) SetBulletType(newValue string) {
	this.BulletType = newValue
}
func (this *Paragraph) GetNumberedBulletStartWith() int32 {
	return this.NumberedBulletStartWith
}

func (this *Paragraph) SetNumberedBulletStartWith(newValue int32) {
	this.NumberedBulletStartWith = newValue
}
func (this *Paragraph) GetNumberedBulletStyle() string {
	return this.NumberedBulletStyle
}

func (this *Paragraph) SetNumberedBulletStyle(newValue string) {
	this.NumberedBulletStyle = newValue
}
func (this *Paragraph) GetBulletFillFormat() IFillFormat {
	return this.BulletFillFormat
}

func (this *Paragraph) SetBulletFillFormat(newValue IFillFormat) {
	this.BulletFillFormat = newValue
}

func (this *Paragraph) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := GetMapValue(objMap, "selfUri"); ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	
	if valAlternateLinks, ok := GetMapValue(objMap, "alternateLinks"); ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valMarginLeft, ok := GetMapValue(objMap, "marginLeft"); ok {
		if valMarginLeft != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeft, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}
	
	if valMarginRight, ok := GetMapValue(objMap, "marginRight"); ok {
		if valMarginRight != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRight, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}
	
	if valSpaceBefore, ok := GetMapValue(objMap, "spaceBefore"); ok {
		if valSpaceBefore != nil {
			var valueForSpaceBefore float64
			err = json.Unmarshal(*valSpaceBefore, &valueForSpaceBefore)
			if err != nil {
				return err
			}
			this.SpaceBefore = valueForSpaceBefore
		}
	}
	
	if valSpaceAfter, ok := GetMapValue(objMap, "spaceAfter"); ok {
		if valSpaceAfter != nil {
			var valueForSpaceAfter float64
			err = json.Unmarshal(*valSpaceAfter, &valueForSpaceAfter)
			if err != nil {
				return err
			}
			this.SpaceAfter = valueForSpaceAfter
		}
	}
	
	if valSpaceWithin, ok := GetMapValue(objMap, "spaceWithin"); ok {
		if valSpaceWithin != nil {
			var valueForSpaceWithin float64
			err = json.Unmarshal(*valSpaceWithin, &valueForSpaceWithin)
			if err != nil {
				return err
			}
			this.SpaceWithin = valueForSpaceWithin
		}
	}
	
	if valIndent, ok := GetMapValue(objMap, "indent"); ok {
		if valIndent != nil {
			var valueForIndent float64
			err = json.Unmarshal(*valIndent, &valueForIndent)
			if err != nil {
				return err
			}
			this.Indent = valueForIndent
		}
	}
	
	if valAlignment, ok := GetMapValue(objMap, "alignment"); ok {
		if valAlignment != nil {
			var valueForAlignment string
			err = json.Unmarshal(*valAlignment, &valueForAlignment)
			if err != nil {
				var valueForAlignmentInt int32
				err = json.Unmarshal(*valAlignment, &valueForAlignmentInt)
				if err != nil {
					return err
				}
				this.Alignment = string(valueForAlignmentInt)
			} else {
				this.Alignment = valueForAlignment
			}
		}
	}
	
	if valFontAlignment, ok := GetMapValue(objMap, "fontAlignment"); ok {
		if valFontAlignment != nil {
			var valueForFontAlignment string
			err = json.Unmarshal(*valFontAlignment, &valueForFontAlignment)
			if err != nil {
				var valueForFontAlignmentInt int32
				err = json.Unmarshal(*valFontAlignment, &valueForFontAlignmentInt)
				if err != nil {
					return err
				}
				this.FontAlignment = string(valueForFontAlignmentInt)
			} else {
				this.FontAlignment = valueForFontAlignment
			}
		}
	}
	
	if valDefaultTabSize, ok := GetMapValue(objMap, "defaultTabSize"); ok {
		if valDefaultTabSize != nil {
			var valueForDefaultTabSize float64
			err = json.Unmarshal(*valDefaultTabSize, &valueForDefaultTabSize)
			if err != nil {
				return err
			}
			this.DefaultTabSize = valueForDefaultTabSize
		}
	}
	
	if valDepth, ok := GetMapValue(objMap, "depth"); ok {
		if valDepth != nil {
			var valueForDepth int32
			err = json.Unmarshal(*valDepth, &valueForDepth)
			if err != nil {
				return err
			}
			this.Depth = valueForDepth
		}
	}
	
	if valHangingPunctuation, ok := GetMapValue(objMap, "hangingPunctuation"); ok {
		if valHangingPunctuation != nil {
			var valueForHangingPunctuation string
			err = json.Unmarshal(*valHangingPunctuation, &valueForHangingPunctuation)
			if err != nil {
				var valueForHangingPunctuationInt int32
				err = json.Unmarshal(*valHangingPunctuation, &valueForHangingPunctuationInt)
				if err != nil {
					return err
				}
				this.HangingPunctuation = string(valueForHangingPunctuationInt)
			} else {
				this.HangingPunctuation = valueForHangingPunctuation
			}
		}
	}
	
	if valEastAsianLineBreak, ok := GetMapValue(objMap, "eastAsianLineBreak"); ok {
		if valEastAsianLineBreak != nil {
			var valueForEastAsianLineBreak string
			err = json.Unmarshal(*valEastAsianLineBreak, &valueForEastAsianLineBreak)
			if err != nil {
				var valueForEastAsianLineBreakInt int32
				err = json.Unmarshal(*valEastAsianLineBreak, &valueForEastAsianLineBreakInt)
				if err != nil {
					return err
				}
				this.EastAsianLineBreak = string(valueForEastAsianLineBreakInt)
			} else {
				this.EastAsianLineBreak = valueForEastAsianLineBreak
			}
		}
	}
	
	if valLatinLineBreak, ok := GetMapValue(objMap, "latinLineBreak"); ok {
		if valLatinLineBreak != nil {
			var valueForLatinLineBreak string
			err = json.Unmarshal(*valLatinLineBreak, &valueForLatinLineBreak)
			if err != nil {
				var valueForLatinLineBreakInt int32
				err = json.Unmarshal(*valLatinLineBreak, &valueForLatinLineBreakInt)
				if err != nil {
					return err
				}
				this.LatinLineBreak = string(valueForLatinLineBreakInt)
			} else {
				this.LatinLineBreak = valueForLatinLineBreak
			}
		}
	}
	
	if valRightToLeft, ok := GetMapValue(objMap, "rightToLeft"); ok {
		if valRightToLeft != nil {
			var valueForRightToLeft string
			err = json.Unmarshal(*valRightToLeft, &valueForRightToLeft)
			if err != nil {
				var valueForRightToLeftInt int32
				err = json.Unmarshal(*valRightToLeft, &valueForRightToLeftInt)
				if err != nil {
					return err
				}
				this.RightToLeft = string(valueForRightToLeftInt)
			} else {
				this.RightToLeft = valueForRightToLeft
			}
		}
	}
	
	if valPortionList, ok := GetMapValue(objMap, "portionList"); ok {
		if valPortionList != nil {
			var valueForPortionList []json.RawMessage
			err = json.Unmarshal(*valPortionList, &valueForPortionList)
			if err != nil {
				return err
			}
			valueForIPortionList := make([]IPortion, len(valueForPortionList))
			for i, v := range valueForPortionList {
				vObject, err := createObjectForType("Portion", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPortionList[i] = vObject.(IPortion)
				}
			}
			this.PortionList = valueForIPortionList
		}
	}
	
	if valDefaultPortionFormat, ok := GetMapValue(objMap, "defaultPortionFormat"); ok {
		if valDefaultPortionFormat != nil {
			var valueForDefaultPortionFormat PortionFormat
			err = json.Unmarshal(*valDefaultPortionFormat, &valueForDefaultPortionFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PortionFormat", *valDefaultPortionFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDefaultPortionFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPortionFormat)
			if ok {
				this.DefaultPortionFormat = vInterfaceObject
			}
		}
	}
	
	if valBulletChar, ok := GetMapValue(objMap, "bulletChar"); ok {
		if valBulletChar != nil {
			var valueForBulletChar string
			err = json.Unmarshal(*valBulletChar, &valueForBulletChar)
			if err != nil {
				return err
			}
			this.BulletChar = valueForBulletChar
		}
	}
	
	if valBulletHeight, ok := GetMapValue(objMap, "bulletHeight"); ok {
		if valBulletHeight != nil {
			var valueForBulletHeight float64
			err = json.Unmarshal(*valBulletHeight, &valueForBulletHeight)
			if err != nil {
				return err
			}
			this.BulletHeight = valueForBulletHeight
		}
	}
	
	if valBulletType, ok := GetMapValue(objMap, "bulletType"); ok {
		if valBulletType != nil {
			var valueForBulletType string
			err = json.Unmarshal(*valBulletType, &valueForBulletType)
			if err != nil {
				var valueForBulletTypeInt int32
				err = json.Unmarshal(*valBulletType, &valueForBulletTypeInt)
				if err != nil {
					return err
				}
				this.BulletType = string(valueForBulletTypeInt)
			} else {
				this.BulletType = valueForBulletType
			}
		}
	}
	
	if valNumberedBulletStartWith, ok := GetMapValue(objMap, "numberedBulletStartWith"); ok {
		if valNumberedBulletStartWith != nil {
			var valueForNumberedBulletStartWith int32
			err = json.Unmarshal(*valNumberedBulletStartWith, &valueForNumberedBulletStartWith)
			if err != nil {
				return err
			}
			this.NumberedBulletStartWith = valueForNumberedBulletStartWith
		}
	}
	
	if valNumberedBulletStyle, ok := GetMapValue(objMap, "numberedBulletStyle"); ok {
		if valNumberedBulletStyle != nil {
			var valueForNumberedBulletStyle string
			err = json.Unmarshal(*valNumberedBulletStyle, &valueForNumberedBulletStyle)
			if err != nil {
				var valueForNumberedBulletStyleInt int32
				err = json.Unmarshal(*valNumberedBulletStyle, &valueForNumberedBulletStyleInt)
				if err != nil {
					return err
				}
				this.NumberedBulletStyle = string(valueForNumberedBulletStyleInt)
			} else {
				this.NumberedBulletStyle = valueForNumberedBulletStyle
			}
		}
	}
	
	if valBulletFillFormat, ok := GetMapValue(objMap, "bulletFillFormat"); ok {
		if valBulletFillFormat != nil {
			var valueForBulletFillFormat FillFormat
			err = json.Unmarshal(*valBulletFillFormat, &valueForBulletFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valBulletFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBulletFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.BulletFillFormat = vInterfaceObject
			}
		}
	}

	return nil
}
