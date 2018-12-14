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

// Represents portion resource
type IPortion interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// A list of links that originate from this document.
	getLinks() []ResourceUri
	setLinks(newValue []ResourceUri)

	getText() string
	setText(newValue string)

	getFontBold() int32
	setFontBold(newValue int32)

	getFontItalic() int32
	setFontItalic(newValue int32)

	getFontUnderline() TextUnderlineType
	setFontUnderline(newValue TextUnderlineType)

	getStrikethroughType() TextStrikethroughType
	setStrikethroughType(newValue TextStrikethroughType)

	getTextCapType() TextCapType
	setTextCapType(newValue TextCapType)

	getEscapement() float64
	setEscapement(newValue float64)

	getSpacing() float64
	setSpacing(newValue float64)

	getFontColor() string
	setFontColor(newValue string)

	getHighlightColor() string
	setHighlightColor(newValue string)

	getFontHeight() float64
	setFontHeight(newValue float64)

	getNormaliseHeight() int32
	setNormaliseHeight(newValue int32)

	getProofDisabled() int32
	setProofDisabled(newValue int32)

	getSmartTagClean() bool
	setSmartTagClean(newValue bool)

	getKerningMinimalSize() float64
	setKerningMinimalSize(newValue float64)

	getKumimoji() int32
	setKumimoji(newValue int32)

	getLanguageId() string
	setLanguageId(newValue string)

	getAlternativeLanguageId() string
	setAlternativeLanguageId(newValue string)

	getIsHardUnderlineFill() int32
	setIsHardUnderlineFill(newValue int32)

	getIsHardUnderlineLine() int32
	setIsHardUnderlineLine(newValue int32)

	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	getUnderlineFillFormat() IFillFormat
	setUnderlineFillFormat(newValue IFillFormat)

	getUnderlineLineFormat() ILineFormat
	setUnderlineLineFormat(newValue ILineFormat)
}

type Portion struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	Text string `json:"Text,omitempty"`

	FontBold int32 `json:"FontBold,omitempty"`

	FontItalic int32 `json:"FontItalic,omitempty"`

	FontUnderline TextUnderlineType `json:"FontUnderline,omitempty"`

	StrikethroughType TextStrikethroughType `json:"StrikethroughType,omitempty"`

	TextCapType TextCapType `json:"TextCapType,omitempty"`

	Escapement float64 `json:"Escapement,omitempty"`

	Spacing float64 `json:"Spacing,omitempty"`

	FontColor string `json:"FontColor,omitempty"`

	HighlightColor string `json:"HighlightColor,omitempty"`

	FontHeight float64 `json:"FontHeight,omitempty"`

	NormaliseHeight int32 `json:"NormaliseHeight,omitempty"`

	ProofDisabled int32 `json:"ProofDisabled,omitempty"`

	SmartTagClean bool `json:"SmartTagClean,omitempty"`

	KerningMinimalSize float64 `json:"KerningMinimalSize,omitempty"`

	Kumimoji int32 `json:"Kumimoji,omitempty"`

	LanguageId string `json:"LanguageId,omitempty"`

	AlternativeLanguageId string `json:"AlternativeLanguageId,omitempty"`

	IsHardUnderlineFill int32 `json:"IsHardUnderlineFill,omitempty"`

	IsHardUnderlineLine int32 `json:"IsHardUnderlineLine,omitempty"`

	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	UnderlineFillFormat IFillFormat `json:"UnderlineFillFormat,omitempty"`

	UnderlineLineFormat ILineFormat `json:"UnderlineLineFormat,omitempty"`
}

func (this Portion) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Portion) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Portion) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Portion) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Portion) getLinks() []ResourceUri {
	return this.Links
}

func (this Portion) setLinks(newValue []ResourceUri) {
	this.Links = newValue
}
func (this Portion) getText() string {
	return this.Text
}

func (this Portion) setText(newValue string) {
	this.Text = newValue
}
func (this Portion) getFontBold() int32 {
	return this.FontBold
}

func (this Portion) setFontBold(newValue int32) {
	this.FontBold = newValue
}
func (this Portion) getFontItalic() int32 {
	return this.FontItalic
}

func (this Portion) setFontItalic(newValue int32) {
	this.FontItalic = newValue
}
func (this Portion) getFontUnderline() TextUnderlineType {
	return this.FontUnderline
}

func (this Portion) setFontUnderline(newValue TextUnderlineType) {
	this.FontUnderline = newValue
}
func (this Portion) getStrikethroughType() TextStrikethroughType {
	return this.StrikethroughType
}

func (this Portion) setStrikethroughType(newValue TextStrikethroughType) {
	this.StrikethroughType = newValue
}
func (this Portion) getTextCapType() TextCapType {
	return this.TextCapType
}

func (this Portion) setTextCapType(newValue TextCapType) {
	this.TextCapType = newValue
}
func (this Portion) getEscapement() float64 {
	return this.Escapement
}

func (this Portion) setEscapement(newValue float64) {
	this.Escapement = newValue
}
func (this Portion) getSpacing() float64 {
	return this.Spacing
}

func (this Portion) setSpacing(newValue float64) {
	this.Spacing = newValue
}
func (this Portion) getFontColor() string {
	return this.FontColor
}

func (this Portion) setFontColor(newValue string) {
	this.FontColor = newValue
}
func (this Portion) getHighlightColor() string {
	return this.HighlightColor
}

func (this Portion) setHighlightColor(newValue string) {
	this.HighlightColor = newValue
}
func (this Portion) getFontHeight() float64 {
	return this.FontHeight
}

func (this Portion) setFontHeight(newValue float64) {
	this.FontHeight = newValue
}
func (this Portion) getNormaliseHeight() int32 {
	return this.NormaliseHeight
}

func (this Portion) setNormaliseHeight(newValue int32) {
	this.NormaliseHeight = newValue
}
func (this Portion) getProofDisabled() int32 {
	return this.ProofDisabled
}

func (this Portion) setProofDisabled(newValue int32) {
	this.ProofDisabled = newValue
}
func (this Portion) getSmartTagClean() bool {
	return this.SmartTagClean
}

func (this Portion) setSmartTagClean(newValue bool) {
	this.SmartTagClean = newValue
}
func (this Portion) getKerningMinimalSize() float64 {
	return this.KerningMinimalSize
}

func (this Portion) setKerningMinimalSize(newValue float64) {
	this.KerningMinimalSize = newValue
}
func (this Portion) getKumimoji() int32 {
	return this.Kumimoji
}

func (this Portion) setKumimoji(newValue int32) {
	this.Kumimoji = newValue
}
func (this Portion) getLanguageId() string {
	return this.LanguageId
}

func (this Portion) setLanguageId(newValue string) {
	this.LanguageId = newValue
}
func (this Portion) getAlternativeLanguageId() string {
	return this.AlternativeLanguageId
}

func (this Portion) setAlternativeLanguageId(newValue string) {
	this.AlternativeLanguageId = newValue
}
func (this Portion) getIsHardUnderlineFill() int32 {
	return this.IsHardUnderlineFill
}

func (this Portion) setIsHardUnderlineFill(newValue int32) {
	this.IsHardUnderlineFill = newValue
}
func (this Portion) getIsHardUnderlineLine() int32 {
	return this.IsHardUnderlineLine
}

func (this Portion) setIsHardUnderlineLine(newValue int32) {
	this.IsHardUnderlineLine = newValue
}
func (this Portion) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this Portion) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this Portion) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this Portion) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this Portion) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this Portion) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this Portion) getUnderlineFillFormat() IFillFormat {
	return this.UnderlineFillFormat
}

func (this Portion) setUnderlineFillFormat(newValue IFillFormat) {
	this.UnderlineFillFormat = newValue
}
func (this Portion) getUnderlineLineFormat() ILineFormat {
	return this.UnderlineLineFormat
}

func (this Portion) setUnderlineLineFormat(newValue ILineFormat) {
	this.UnderlineLineFormat = newValue
}

func (this *Portion) UnmarshalJSON(b []byte) error {
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

	if valText, ok := objMap["Text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}

	if valFontBold, ok := objMap["FontBold"]; ok {
		if valFontBold != nil {
			var valueForFontBold int32
			err = json.Unmarshal(*valFontBold, &valueForFontBold)
			if err != nil {
				return err
			}
			this.FontBold = valueForFontBold
		}
	}

	if valFontItalic, ok := objMap["FontItalic"]; ok {
		if valFontItalic != nil {
			var valueForFontItalic int32
			err = json.Unmarshal(*valFontItalic, &valueForFontItalic)
			if err != nil {
				return err
			}
			this.FontItalic = valueForFontItalic
		}
	}

	if valFontUnderline, ok := objMap["FontUnderline"]; ok {
		if valFontUnderline != nil {
			var valueForFontUnderline TextUnderlineType
			err = json.Unmarshal(*valFontUnderline, &valueForFontUnderline)
			if err != nil {
				return err
			}
			this.FontUnderline = valueForFontUnderline
		}
	}

	if valStrikethroughType, ok := objMap["StrikethroughType"]; ok {
		if valStrikethroughType != nil {
			var valueForStrikethroughType TextStrikethroughType
			err = json.Unmarshal(*valStrikethroughType, &valueForStrikethroughType)
			if err != nil {
				return err
			}
			this.StrikethroughType = valueForStrikethroughType
		}
	}

	if valTextCapType, ok := objMap["TextCapType"]; ok {
		if valTextCapType != nil {
			var valueForTextCapType TextCapType
			err = json.Unmarshal(*valTextCapType, &valueForTextCapType)
			if err != nil {
				return err
			}
			this.TextCapType = valueForTextCapType
		}
	}

	if valEscapement, ok := objMap["Escapement"]; ok {
		if valEscapement != nil {
			var valueForEscapement float64
			err = json.Unmarshal(*valEscapement, &valueForEscapement)
			if err != nil {
				return err
			}
			this.Escapement = valueForEscapement
		}
	}

	if valSpacing, ok := objMap["Spacing"]; ok {
		if valSpacing != nil {
			var valueForSpacing float64
			err = json.Unmarshal(*valSpacing, &valueForSpacing)
			if err != nil {
				return err
			}
			this.Spacing = valueForSpacing
		}
	}

	if valFontColor, ok := objMap["FontColor"]; ok {
		if valFontColor != nil {
			var valueForFontColor string
			err = json.Unmarshal(*valFontColor, &valueForFontColor)
			if err != nil {
				return err
			}
			this.FontColor = valueForFontColor
		}
	}

	if valHighlightColor, ok := objMap["HighlightColor"]; ok {
		if valHighlightColor != nil {
			var valueForHighlightColor string
			err = json.Unmarshal(*valHighlightColor, &valueForHighlightColor)
			if err != nil {
				return err
			}
			this.HighlightColor = valueForHighlightColor
		}
	}

	if valFontHeight, ok := objMap["FontHeight"]; ok {
		if valFontHeight != nil {
			var valueForFontHeight float64
			err = json.Unmarshal(*valFontHeight, &valueForFontHeight)
			if err != nil {
				return err
			}
			this.FontHeight = valueForFontHeight
		}
	}

	if valNormaliseHeight, ok := objMap["NormaliseHeight"]; ok {
		if valNormaliseHeight != nil {
			var valueForNormaliseHeight int32
			err = json.Unmarshal(*valNormaliseHeight, &valueForNormaliseHeight)
			if err != nil {
				return err
			}
			this.NormaliseHeight = valueForNormaliseHeight
		}
	}

	if valProofDisabled, ok := objMap["ProofDisabled"]; ok {
		if valProofDisabled != nil {
			var valueForProofDisabled int32
			err = json.Unmarshal(*valProofDisabled, &valueForProofDisabled)
			if err != nil {
				return err
			}
			this.ProofDisabled = valueForProofDisabled
		}
	}

	if valSmartTagClean, ok := objMap["SmartTagClean"]; ok {
		if valSmartTagClean != nil {
			var valueForSmartTagClean bool
			err = json.Unmarshal(*valSmartTagClean, &valueForSmartTagClean)
			if err != nil {
				return err
			}
			this.SmartTagClean = valueForSmartTagClean
		}
	}

	if valKerningMinimalSize, ok := objMap["KerningMinimalSize"]; ok {
		if valKerningMinimalSize != nil {
			var valueForKerningMinimalSize float64
			err = json.Unmarshal(*valKerningMinimalSize, &valueForKerningMinimalSize)
			if err != nil {
				return err
			}
			this.KerningMinimalSize = valueForKerningMinimalSize
		}
	}

	if valKumimoji, ok := objMap["Kumimoji"]; ok {
		if valKumimoji != nil {
			var valueForKumimoji int32
			err = json.Unmarshal(*valKumimoji, &valueForKumimoji)
			if err != nil {
				return err
			}
			this.Kumimoji = valueForKumimoji
		}
	}

	if valLanguageId, ok := objMap["LanguageId"]; ok {
		if valLanguageId != nil {
			var valueForLanguageId string
			err = json.Unmarshal(*valLanguageId, &valueForLanguageId)
			if err != nil {
				return err
			}
			this.LanguageId = valueForLanguageId
		}
	}

	if valAlternativeLanguageId, ok := objMap["AlternativeLanguageId"]; ok {
		if valAlternativeLanguageId != nil {
			var valueForAlternativeLanguageId string
			err = json.Unmarshal(*valAlternativeLanguageId, &valueForAlternativeLanguageId)
			if err != nil {
				return err
			}
			this.AlternativeLanguageId = valueForAlternativeLanguageId
		}
	}

	if valIsHardUnderlineFill, ok := objMap["IsHardUnderlineFill"]; ok {
		if valIsHardUnderlineFill != nil {
			var valueForIsHardUnderlineFill int32
			err = json.Unmarshal(*valIsHardUnderlineFill, &valueForIsHardUnderlineFill)
			if err != nil {
				return err
			}
			this.IsHardUnderlineFill = valueForIsHardUnderlineFill
		}
	}

	if valIsHardUnderlineLine, ok := objMap["IsHardUnderlineLine"]; ok {
		if valIsHardUnderlineLine != nil {
			var valueForIsHardUnderlineLine int32
			err = json.Unmarshal(*valIsHardUnderlineLine, &valueForIsHardUnderlineLine)
			if err != nil {
				return err
			}
			this.IsHardUnderlineLine = valueForIsHardUnderlineLine
		}
	}

	if valFillFormat, ok := objMap["FillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}

	if valEffectFormat, ok := objMap["EffectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}

	if valLineFormat, ok := objMap["LineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}

	if valUnderlineFillFormat, ok := objMap["UnderlineFillFormat"]; ok {
		if valUnderlineFillFormat != nil {
			var valueForUnderlineFillFormat FillFormat
			err = json.Unmarshal(*valUnderlineFillFormat, &valueForUnderlineFillFormat)
			if err != nil {
				return err
			}
			this.UnderlineFillFormat = valueForUnderlineFillFormat
		}
	}

	if valUnderlineLineFormat, ok := objMap["UnderlineLineFormat"]; ok {
		if valUnderlineLineFormat != nil {
			var valueForUnderlineLineFormat LineFormat
			err = json.Unmarshal(*valUnderlineLineFormat, &valueForUnderlineLineFormat)
			if err != nil {
				return err
			}
			this.UnderlineLineFormat = valueForUnderlineLineFormat
		}
	}

    return nil
}
