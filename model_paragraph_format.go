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

// Paragraph formatting properties.
type IParagraphFormat interface {

	// Depth.
	GetDepth() int32
	SetDepth(newValue int32)

	// Text alignment.
	GetAlignment() string
	SetAlignment(newValue string)

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

	// Font alignment.
	GetFontAlignment() string
	SetFontAlignment(newValue string)

	// First line indent.
	GetIndent() float64
	SetIndent(newValue float64)

	// Determines whether the Right to Left writing is used in a paragraph. No inheritance applied.
	GetRightToLeft() string
	SetRightToLeft(newValue string)

	// Determines whether the East Asian line break is used in a paragraph. No inheritance applied.
	GetEastAsianLineBreak() string
	SetEastAsianLineBreak(newValue string)

	// Determines whether the Latin line break is used in a paragraph. No inheritance applied.
	GetLatinLineBreak() string
	SetLatinLineBreak(newValue string)

	// Determines whether the hanging punctuation is used in a paragraph. No inheritance applied.
	GetHangingPunctuation() string
	SetHangingPunctuation(newValue string)

	// Returns or sets default tabulation size with no inheritance.
	GetDefaultTabSize() float64
	SetDefaultTabSize(newValue float64)

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

type ParagraphFormat struct {

	// Depth.
	Depth int32 `json:"Depth,omitempty"`

	// Text alignment.
	Alignment string `json:"Alignment,omitempty"`

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

	// Font alignment.
	FontAlignment string `json:"FontAlignment,omitempty"`

	// First line indent.
	Indent float64 `json:"Indent,omitempty"`

	// Determines whether the Right to Left writing is used in a paragraph. No inheritance applied.
	RightToLeft string `json:"RightToLeft,omitempty"`

	// Determines whether the East Asian line break is used in a paragraph. No inheritance applied.
	EastAsianLineBreak string `json:"EastAsianLineBreak,omitempty"`

	// Determines whether the Latin line break is used in a paragraph. No inheritance applied.
	LatinLineBreak string `json:"LatinLineBreak,omitempty"`

	// Determines whether the hanging punctuation is used in a paragraph. No inheritance applied.
	HangingPunctuation string `json:"HangingPunctuation,omitempty"`

	// Returns or sets default tabulation size with no inheritance.
	DefaultTabSize float64 `json:"DefaultTabSize,omitempty"`

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

func NewParagraphFormat() *ParagraphFormat {
	instance := new(ParagraphFormat)
	return instance
}

func (this *ParagraphFormat) GetDepth() int32 {
	return this.Depth
}

func (this *ParagraphFormat) SetDepth(newValue int32) {
	this.Depth = newValue
}
func (this *ParagraphFormat) GetAlignment() string {
	return this.Alignment
}

func (this *ParagraphFormat) SetAlignment(newValue string) {
	this.Alignment = newValue
}
func (this *ParagraphFormat) GetMarginLeft() float64 {
	return this.MarginLeft
}

func (this *ParagraphFormat) SetMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this *ParagraphFormat) GetMarginRight() float64 {
	return this.MarginRight
}

func (this *ParagraphFormat) SetMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this *ParagraphFormat) GetSpaceBefore() float64 {
	return this.SpaceBefore
}

func (this *ParagraphFormat) SetSpaceBefore(newValue float64) {
	this.SpaceBefore = newValue
}
func (this *ParagraphFormat) GetSpaceAfter() float64 {
	return this.SpaceAfter
}

func (this *ParagraphFormat) SetSpaceAfter(newValue float64) {
	this.SpaceAfter = newValue
}
func (this *ParagraphFormat) GetSpaceWithin() float64 {
	return this.SpaceWithin
}

func (this *ParagraphFormat) SetSpaceWithin(newValue float64) {
	this.SpaceWithin = newValue
}
func (this *ParagraphFormat) GetFontAlignment() string {
	return this.FontAlignment
}

func (this *ParagraphFormat) SetFontAlignment(newValue string) {
	this.FontAlignment = newValue
}
func (this *ParagraphFormat) GetIndent() float64 {
	return this.Indent
}

func (this *ParagraphFormat) SetIndent(newValue float64) {
	this.Indent = newValue
}
func (this *ParagraphFormat) GetRightToLeft() string {
	return this.RightToLeft
}

func (this *ParagraphFormat) SetRightToLeft(newValue string) {
	this.RightToLeft = newValue
}
func (this *ParagraphFormat) GetEastAsianLineBreak() string {
	return this.EastAsianLineBreak
}

func (this *ParagraphFormat) SetEastAsianLineBreak(newValue string) {
	this.EastAsianLineBreak = newValue
}
func (this *ParagraphFormat) GetLatinLineBreak() string {
	return this.LatinLineBreak
}

func (this *ParagraphFormat) SetLatinLineBreak(newValue string) {
	this.LatinLineBreak = newValue
}
func (this *ParagraphFormat) GetHangingPunctuation() string {
	return this.HangingPunctuation
}

func (this *ParagraphFormat) SetHangingPunctuation(newValue string) {
	this.HangingPunctuation = newValue
}
func (this *ParagraphFormat) GetDefaultTabSize() float64 {
	return this.DefaultTabSize
}

func (this *ParagraphFormat) SetDefaultTabSize(newValue float64) {
	this.DefaultTabSize = newValue
}
func (this *ParagraphFormat) GetDefaultPortionFormat() IPortionFormat {
	return this.DefaultPortionFormat
}

func (this *ParagraphFormat) SetDefaultPortionFormat(newValue IPortionFormat) {
	this.DefaultPortionFormat = newValue
}
func (this *ParagraphFormat) GetBulletChar() string {
	return this.BulletChar
}

func (this *ParagraphFormat) SetBulletChar(newValue string) {
	this.BulletChar = newValue
}
func (this *ParagraphFormat) GetBulletHeight() float64 {
	return this.BulletHeight
}

func (this *ParagraphFormat) SetBulletHeight(newValue float64) {
	this.BulletHeight = newValue
}
func (this *ParagraphFormat) GetBulletType() string {
	return this.BulletType
}

func (this *ParagraphFormat) SetBulletType(newValue string) {
	this.BulletType = newValue
}
func (this *ParagraphFormat) GetNumberedBulletStartWith() int32 {
	return this.NumberedBulletStartWith
}

func (this *ParagraphFormat) SetNumberedBulletStartWith(newValue int32) {
	this.NumberedBulletStartWith = newValue
}
func (this *ParagraphFormat) GetNumberedBulletStyle() string {
	return this.NumberedBulletStyle
}

func (this *ParagraphFormat) SetNumberedBulletStyle(newValue string) {
	this.NumberedBulletStyle = newValue
}
func (this *ParagraphFormat) GetBulletFillFormat() IFillFormat {
	return this.BulletFillFormat
}

func (this *ParagraphFormat) SetBulletFillFormat(newValue IFillFormat) {
	this.BulletFillFormat = newValue
}

func (this *ParagraphFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
