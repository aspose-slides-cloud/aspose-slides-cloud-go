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
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// A list of links that originate from this document.
	getLinks() []ResourceUri
	setLinks(newValue []ResourceUri)

	getMarginLeft() float64
	setMarginLeft(newValue float64)

	getMarginRight() float64
	setMarginRight(newValue float64)

	getSpaceBefore() float64
	setSpaceBefore(newValue float64)

	getSpaceAfter() float64
	setSpaceAfter(newValue float64)

	getSpaceWithin() float64
	setSpaceWithin(newValue float64)

	getIndent() float64
	setIndent(newValue float64)

	getAlignment() TextAlignment
	setAlignment(newValue TextAlignment)

	getFontAlignment() FontAlignment
	setFontAlignment(newValue FontAlignment)

	getDefaultTabSize() float64
	setDefaultTabSize(newValue float64)

	getDepth() int32
	setDepth(newValue int32)

	getBulletChar() string
	setBulletChar(newValue string)

	getBulletHeight() float64
	setBulletHeight(newValue float64)

	getBulletType() BulletType
	setBulletType(newValue BulletType)

	getNumberedBulletStartWith() int32
	setNumberedBulletStartWith(newValue int32)

	getNumberedBulletStyle() NumberedBulletStyle
	setNumberedBulletStyle(newValue NumberedBulletStyle)

	getHangingPunctuation() int32
	setHangingPunctuation(newValue int32)

	getEastAsianLineBreak() int32
	setEastAsianLineBreak(newValue int32)

	getLatinLineBreak() int32
	setLatinLineBreak(newValue int32)

	getRightToLeft() int32
	setRightToLeft(newValue int32)

	getPortionList() []ResourceUriElement
	setPortionList(newValue []ResourceUriElement)
}

type Paragraph struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	MarginLeft float64 `json:"MarginLeft,omitempty"`

	MarginRight float64 `json:"MarginRight,omitempty"`

	SpaceBefore float64 `json:"SpaceBefore,omitempty"`

	SpaceAfter float64 `json:"SpaceAfter,omitempty"`

	SpaceWithin float64 `json:"SpaceWithin,omitempty"`

	Indent float64 `json:"Indent,omitempty"`

	Alignment TextAlignment `json:"Alignment,omitempty"`

	FontAlignment FontAlignment `json:"FontAlignment,omitempty"`

	DefaultTabSize float64 `json:"DefaultTabSize,omitempty"`

	Depth int32 `json:"Depth,omitempty"`

	BulletChar string `json:"BulletChar,omitempty"`

	BulletHeight float64 `json:"BulletHeight,omitempty"`

	BulletType BulletType `json:"BulletType,omitempty"`

	NumberedBulletStartWith int32 `json:"NumberedBulletStartWith,omitempty"`

	NumberedBulletStyle NumberedBulletStyle `json:"NumberedBulletStyle,omitempty"`

	HangingPunctuation int32 `json:"HangingPunctuation,omitempty"`

	EastAsianLineBreak int32 `json:"EastAsianLineBreak,omitempty"`

	LatinLineBreak int32 `json:"LatinLineBreak,omitempty"`

	RightToLeft int32 `json:"RightToLeft,omitempty"`

	PortionList []ResourceUriElement `json:"PortionList,omitempty"`
}

func (this Paragraph) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Paragraph) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Paragraph) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Paragraph) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Paragraph) getLinks() []ResourceUri {
	return this.Links
}

func (this Paragraph) setLinks(newValue []ResourceUri) {
	this.Links = newValue
}
func (this Paragraph) getMarginLeft() float64 {
	return this.MarginLeft
}

func (this Paragraph) setMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this Paragraph) getMarginRight() float64 {
	return this.MarginRight
}

func (this Paragraph) setMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this Paragraph) getSpaceBefore() float64 {
	return this.SpaceBefore
}

func (this Paragraph) setSpaceBefore(newValue float64) {
	this.SpaceBefore = newValue
}
func (this Paragraph) getSpaceAfter() float64 {
	return this.SpaceAfter
}

func (this Paragraph) setSpaceAfter(newValue float64) {
	this.SpaceAfter = newValue
}
func (this Paragraph) getSpaceWithin() float64 {
	return this.SpaceWithin
}

func (this Paragraph) setSpaceWithin(newValue float64) {
	this.SpaceWithin = newValue
}
func (this Paragraph) getIndent() float64 {
	return this.Indent
}

func (this Paragraph) setIndent(newValue float64) {
	this.Indent = newValue
}
func (this Paragraph) getAlignment() TextAlignment {
	return this.Alignment
}

func (this Paragraph) setAlignment(newValue TextAlignment) {
	this.Alignment = newValue
}
func (this Paragraph) getFontAlignment() FontAlignment {
	return this.FontAlignment
}

func (this Paragraph) setFontAlignment(newValue FontAlignment) {
	this.FontAlignment = newValue
}
func (this Paragraph) getDefaultTabSize() float64 {
	return this.DefaultTabSize
}

func (this Paragraph) setDefaultTabSize(newValue float64) {
	this.DefaultTabSize = newValue
}
func (this Paragraph) getDepth() int32 {
	return this.Depth
}

func (this Paragraph) setDepth(newValue int32) {
	this.Depth = newValue
}
func (this Paragraph) getBulletChar() string {
	return this.BulletChar
}

func (this Paragraph) setBulletChar(newValue string) {
	this.BulletChar = newValue
}
func (this Paragraph) getBulletHeight() float64 {
	return this.BulletHeight
}

func (this Paragraph) setBulletHeight(newValue float64) {
	this.BulletHeight = newValue
}
func (this Paragraph) getBulletType() BulletType {
	return this.BulletType
}

func (this Paragraph) setBulletType(newValue BulletType) {
	this.BulletType = newValue
}
func (this Paragraph) getNumberedBulletStartWith() int32 {
	return this.NumberedBulletStartWith
}

func (this Paragraph) setNumberedBulletStartWith(newValue int32) {
	this.NumberedBulletStartWith = newValue
}
func (this Paragraph) getNumberedBulletStyle() NumberedBulletStyle {
	return this.NumberedBulletStyle
}

func (this Paragraph) setNumberedBulletStyle(newValue NumberedBulletStyle) {
	this.NumberedBulletStyle = newValue
}
func (this Paragraph) getHangingPunctuation() int32 {
	return this.HangingPunctuation
}

func (this Paragraph) setHangingPunctuation(newValue int32) {
	this.HangingPunctuation = newValue
}
func (this Paragraph) getEastAsianLineBreak() int32 {
	return this.EastAsianLineBreak
}

func (this Paragraph) setEastAsianLineBreak(newValue int32) {
	this.EastAsianLineBreak = newValue
}
func (this Paragraph) getLatinLineBreak() int32 {
	return this.LatinLineBreak
}

func (this Paragraph) setLatinLineBreak(newValue int32) {
	this.LatinLineBreak = newValue
}
func (this Paragraph) getRightToLeft() int32 {
	return this.RightToLeft
}

func (this Paragraph) setRightToLeft(newValue int32) {
	this.RightToLeft = newValue
}
func (this Paragraph) getPortionList() []ResourceUriElement {
	return this.PortionList
}

func (this Paragraph) setPortionList(newValue []ResourceUriElement) {
	this.PortionList = newValue
}

func (this *Paragraph) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valSelfUri, ok := objMap["SelfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}

	if valAlternateLinks, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}

	if valLinks, ok := objMap["Links"]; ok {
		if valLinks != nil {
			var valueForLinks []ResourceUri
			err = json.Unmarshal(*valLinks, &valueForLinks)
			if err != nil {
				return err
			}
			this.Links = valueForLinks
		}
	}

	if valMarginLeft, ok := objMap["MarginLeft"]; ok {
		if valMarginLeft != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeft, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}

	if valMarginRight, ok := objMap["MarginRight"]; ok {
		if valMarginRight != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRight, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}

	if valSpaceBefore, ok := objMap["SpaceBefore"]; ok {
		if valSpaceBefore != nil {
			var valueForSpaceBefore float64
			err = json.Unmarshal(*valSpaceBefore, &valueForSpaceBefore)
			if err != nil {
				return err
			}
			this.SpaceBefore = valueForSpaceBefore
		}
	}

	if valSpaceAfter, ok := objMap["SpaceAfter"]; ok {
		if valSpaceAfter != nil {
			var valueForSpaceAfter float64
			err = json.Unmarshal(*valSpaceAfter, &valueForSpaceAfter)
			if err != nil {
				return err
			}
			this.SpaceAfter = valueForSpaceAfter
		}
	}

	if valSpaceWithin, ok := objMap["SpaceWithin"]; ok {
		if valSpaceWithin != nil {
			var valueForSpaceWithin float64
			err = json.Unmarshal(*valSpaceWithin, &valueForSpaceWithin)
			if err != nil {
				return err
			}
			this.SpaceWithin = valueForSpaceWithin
		}
	}

	if valIndent, ok := objMap["Indent"]; ok {
		if valIndent != nil {
			var valueForIndent float64
			err = json.Unmarshal(*valIndent, &valueForIndent)
			if err != nil {
				return err
			}
			this.Indent = valueForIndent
		}
	}

	if valAlignment, ok := objMap["Alignment"]; ok {
		if valAlignment != nil {
			var valueForAlignment TextAlignment
			err = json.Unmarshal(*valAlignment, &valueForAlignment)
			if err != nil {
				return err
			}
			this.Alignment = valueForAlignment
		}
	}

	if valFontAlignment, ok := objMap["FontAlignment"]; ok {
		if valFontAlignment != nil {
			var valueForFontAlignment FontAlignment
			err = json.Unmarshal(*valFontAlignment, &valueForFontAlignment)
			if err != nil {
				return err
			}
			this.FontAlignment = valueForFontAlignment
		}
	}

	if valDefaultTabSize, ok := objMap["DefaultTabSize"]; ok {
		if valDefaultTabSize != nil {
			var valueForDefaultTabSize float64
			err = json.Unmarshal(*valDefaultTabSize, &valueForDefaultTabSize)
			if err != nil {
				return err
			}
			this.DefaultTabSize = valueForDefaultTabSize
		}
	}

	if valDepth, ok := objMap["Depth"]; ok {
		if valDepth != nil {
			var valueForDepth int32
			err = json.Unmarshal(*valDepth, &valueForDepth)
			if err != nil {
				return err
			}
			this.Depth = valueForDepth
		}
	}

	if valBulletChar, ok := objMap["BulletChar"]; ok {
		if valBulletChar != nil {
			var valueForBulletChar string
			err = json.Unmarshal(*valBulletChar, &valueForBulletChar)
			if err != nil {
				return err
			}
			this.BulletChar = valueForBulletChar
		}
	}

	if valBulletHeight, ok := objMap["BulletHeight"]; ok {
		if valBulletHeight != nil {
			var valueForBulletHeight float64
			err = json.Unmarshal(*valBulletHeight, &valueForBulletHeight)
			if err != nil {
				return err
			}
			this.BulletHeight = valueForBulletHeight
		}
	}

	if valBulletType, ok := objMap["BulletType"]; ok {
		if valBulletType != nil {
			var valueForBulletType BulletType
			err = json.Unmarshal(*valBulletType, &valueForBulletType)
			if err != nil {
				return err
			}
			this.BulletType = valueForBulletType
		}
	}

	if valNumberedBulletStartWith, ok := objMap["NumberedBulletStartWith"]; ok {
		if valNumberedBulletStartWith != nil {
			var valueForNumberedBulletStartWith int32
			err = json.Unmarshal(*valNumberedBulletStartWith, &valueForNumberedBulletStartWith)
			if err != nil {
				return err
			}
			this.NumberedBulletStartWith = valueForNumberedBulletStartWith
		}
	}

	if valNumberedBulletStyle, ok := objMap["NumberedBulletStyle"]; ok {
		if valNumberedBulletStyle != nil {
			var valueForNumberedBulletStyle NumberedBulletStyle
			err = json.Unmarshal(*valNumberedBulletStyle, &valueForNumberedBulletStyle)
			if err != nil {
				return err
			}
			this.NumberedBulletStyle = valueForNumberedBulletStyle
		}
	}

	if valHangingPunctuation, ok := objMap["HangingPunctuation"]; ok {
		if valHangingPunctuation != nil {
			var valueForHangingPunctuation int32
			err = json.Unmarshal(*valHangingPunctuation, &valueForHangingPunctuation)
			if err != nil {
				return err
			}
			this.HangingPunctuation = valueForHangingPunctuation
		}
	}

	if valEastAsianLineBreak, ok := objMap["EastAsianLineBreak"]; ok {
		if valEastAsianLineBreak != nil {
			var valueForEastAsianLineBreak int32
			err = json.Unmarshal(*valEastAsianLineBreak, &valueForEastAsianLineBreak)
			if err != nil {
				return err
			}
			this.EastAsianLineBreak = valueForEastAsianLineBreak
		}
	}

	if valLatinLineBreak, ok := objMap["LatinLineBreak"]; ok {
		if valLatinLineBreak != nil {
			var valueForLatinLineBreak int32
			err = json.Unmarshal(*valLatinLineBreak, &valueForLatinLineBreak)
			if err != nil {
				return err
			}
			this.LatinLineBreak = valueForLatinLineBreak
		}
	}

	if valRightToLeft, ok := objMap["RightToLeft"]; ok {
		if valRightToLeft != nil {
			var valueForRightToLeft int32
			err = json.Unmarshal(*valRightToLeft, &valueForRightToLeft)
			if err != nil {
				return err
			}
			this.RightToLeft = valueForRightToLeft
		}
	}

	if valPortionList, ok := objMap["PortionList"]; ok {
		if valPortionList != nil {
			var valueForPortionList []ResourceUriElement
			err = json.Unmarshal(*valPortionList, &valueForPortionList)
			if err != nil {
				return err
			}
			this.PortionList = valueForPortionList
		}
	}

    return nil
}
