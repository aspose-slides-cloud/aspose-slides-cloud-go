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

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Text.
	getText() string
	setText(newValue string)

	// True for bold font.
	getFontBold() string
	setFontBold(newValue string)

	// True for italic font.
	getFontItalic() string
	setFontItalic(newValue string)

	// Text underline type.
	getFontUnderline() string
	setFontUnderline(newValue string)

	// Text strikethrough type.
	getStrikethroughType() string
	setStrikethroughType(newValue string)

	// Text capitalization type.
	getTextCapType() string
	setTextCapType(newValue string)

	// Superscript or subscript of the text.
	getEscapement() float64
	setEscapement(newValue float64)

	// Intercharacter spacing increment.
	getSpacing() float64
	setSpacing(newValue float64)

	// Font color.
	getFontColor() string
	setFontColor(newValue string)

	// Highlight color.
	getHighlightColor() string
	setHighlightColor(newValue string)

	// Font height.
	getFontHeight() float64
	setFontHeight(newValue float64)

	// True to normalize the text.
	getNormaliseHeight() string
	setNormaliseHeight(newValue string)

	// True if the text proof should be disabled.
	getProofDisabled() string
	setProofDisabled(newValue string)

	// True if smart tag should be cleaned.
	getSmartTagClean() bool
	setSmartTagClean(newValue bool)

	// Minimal font size for kerning.
	getKerningMinimalSize() float64
	setKerningMinimalSize(newValue float64)

	// True if numbers should ignore East-Asian specific vertical text layout.
	getKumimoji() string
	setKumimoji(newValue string)

	// Proving language ID.
	getLanguageId() string
	setLanguageId(newValue string)

	// Alternative proving language ID.
	getAlternativeLanguageId() string
	setAlternativeLanguageId(newValue string)

	// True if underline style has own FillFormat properties.
	getIsHardUnderlineFill() string
	setIsHardUnderlineFill(newValue string)

	// True if underline style has own LineFormat properties.
	getIsHardUnderlineLine() string
	setIsHardUnderlineLine(newValue string)

	// Fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Underline fill format.
	getUnderlineFillFormat() IFillFormat
	setUnderlineFillFormat(newValue IFillFormat)

	// Underline line format.
	getUnderlineLineFormat() ILineFormat
	setUnderlineLineFormat(newValue ILineFormat)
}

type Portion struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Text.
	Text string `json:"Text,omitempty"`

	// True for bold font.
	FontBold string `json:"FontBold,omitempty"`

	// True for italic font.
	FontItalic string `json:"FontItalic,omitempty"`

	// Text underline type.
	FontUnderline string `json:"FontUnderline,omitempty"`

	// Text strikethrough type.
	StrikethroughType string `json:"StrikethroughType,omitempty"`

	// Text capitalization type.
	TextCapType string `json:"TextCapType,omitempty"`

	// Superscript or subscript of the text.
	Escapement float64 `json:"Escapement,omitempty"`

	// Intercharacter spacing increment.
	Spacing float64 `json:"Spacing,omitempty"`

	// Font color.
	FontColor string `json:"FontColor,omitempty"`

	// Highlight color.
	HighlightColor string `json:"HighlightColor,omitempty"`

	// Font height.
	FontHeight float64 `json:"FontHeight,omitempty"`

	// True to normalize the text.
	NormaliseHeight string `json:"NormaliseHeight,omitempty"`

	// True if the text proof should be disabled.
	ProofDisabled string `json:"ProofDisabled,omitempty"`

	// True if smart tag should be cleaned.
	SmartTagClean bool `json:"SmartTagClean"`

	// Minimal font size for kerning.
	KerningMinimalSize float64 `json:"KerningMinimalSize,omitempty"`

	// True if numbers should ignore East-Asian specific vertical text layout.
	Kumimoji string `json:"Kumimoji,omitempty"`

	// Proving language ID.
	LanguageId string `json:"LanguageId,omitempty"`

	// Alternative proving language ID.
	AlternativeLanguageId string `json:"AlternativeLanguageId,omitempty"`

	// True if underline style has own FillFormat properties.
	IsHardUnderlineFill string `json:"IsHardUnderlineFill,omitempty"`

	// True if underline style has own LineFormat properties.
	IsHardUnderlineLine string `json:"IsHardUnderlineLine,omitempty"`

	// Fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Underline fill format.
	UnderlineFillFormat IFillFormat `json:"UnderlineFillFormat,omitempty"`

	// Underline line format.
	UnderlineLineFormat ILineFormat `json:"UnderlineLineFormat,omitempty"`
}

func (this *Portion) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Portion) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Portion) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Portion) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Portion) getText() string {
	return this.Text
}

func (this *Portion) setText(newValue string) {
	this.Text = newValue
}
func (this *Portion) getFontBold() string {
	return this.FontBold
}

func (this *Portion) setFontBold(newValue string) {
	this.FontBold = newValue
}
func (this *Portion) getFontItalic() string {
	return this.FontItalic
}

func (this *Portion) setFontItalic(newValue string) {
	this.FontItalic = newValue
}
func (this *Portion) getFontUnderline() string {
	return this.FontUnderline
}

func (this *Portion) setFontUnderline(newValue string) {
	this.FontUnderline = newValue
}
func (this *Portion) getStrikethroughType() string {
	return this.StrikethroughType
}

func (this *Portion) setStrikethroughType(newValue string) {
	this.StrikethroughType = newValue
}
func (this *Portion) getTextCapType() string {
	return this.TextCapType
}

func (this *Portion) setTextCapType(newValue string) {
	this.TextCapType = newValue
}
func (this *Portion) getEscapement() float64 {
	return this.Escapement
}

func (this *Portion) setEscapement(newValue float64) {
	this.Escapement = newValue
}
func (this *Portion) getSpacing() float64 {
	return this.Spacing
}

func (this *Portion) setSpacing(newValue float64) {
	this.Spacing = newValue
}
func (this *Portion) getFontColor() string {
	return this.FontColor
}

func (this *Portion) setFontColor(newValue string) {
	this.FontColor = newValue
}
func (this *Portion) getHighlightColor() string {
	return this.HighlightColor
}

func (this *Portion) setHighlightColor(newValue string) {
	this.HighlightColor = newValue
}
func (this *Portion) getFontHeight() float64 {
	return this.FontHeight
}

func (this *Portion) setFontHeight(newValue float64) {
	this.FontHeight = newValue
}
func (this *Portion) getNormaliseHeight() string {
	return this.NormaliseHeight
}

func (this *Portion) setNormaliseHeight(newValue string) {
	this.NormaliseHeight = newValue
}
func (this *Portion) getProofDisabled() string {
	return this.ProofDisabled
}

func (this *Portion) setProofDisabled(newValue string) {
	this.ProofDisabled = newValue
}
func (this *Portion) getSmartTagClean() bool {
	return this.SmartTagClean
}

func (this *Portion) setSmartTagClean(newValue bool) {
	this.SmartTagClean = newValue
}
func (this *Portion) getKerningMinimalSize() float64 {
	return this.KerningMinimalSize
}

func (this *Portion) setKerningMinimalSize(newValue float64) {
	this.KerningMinimalSize = newValue
}
func (this *Portion) getKumimoji() string {
	return this.Kumimoji
}

func (this *Portion) setKumimoji(newValue string) {
	this.Kumimoji = newValue
}
func (this *Portion) getLanguageId() string {
	return this.LanguageId
}

func (this *Portion) setLanguageId(newValue string) {
	this.LanguageId = newValue
}
func (this *Portion) getAlternativeLanguageId() string {
	return this.AlternativeLanguageId
}

func (this *Portion) setAlternativeLanguageId(newValue string) {
	this.AlternativeLanguageId = newValue
}
func (this *Portion) getIsHardUnderlineFill() string {
	return this.IsHardUnderlineFill
}

func (this *Portion) setIsHardUnderlineFill(newValue string) {
	this.IsHardUnderlineFill = newValue
}
func (this *Portion) getIsHardUnderlineLine() string {
	return this.IsHardUnderlineLine
}

func (this *Portion) setIsHardUnderlineLine(newValue string) {
	this.IsHardUnderlineLine = newValue
}
func (this *Portion) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Portion) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Portion) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Portion) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Portion) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Portion) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Portion) getUnderlineFillFormat() IFillFormat {
	return this.UnderlineFillFormat
}

func (this *Portion) setUnderlineFillFormat(newValue IFillFormat) {
	this.UnderlineFillFormat = newValue
}
func (this *Portion) getUnderlineLineFormat() ILineFormat {
	return this.UnderlineLineFormat
}

func (this *Portion) setUnderlineLineFormat(newValue ILineFormat) {
	this.UnderlineLineFormat = newValue
}

func (this *Portion) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["selfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valText, ok := objMap["text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	if valTextCap, ok := objMap["Text"]; ok {
		if valTextCap != nil {
			var valueForText string
			err = json.Unmarshal(*valTextCap, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	this.FontBold = ""
	if valFontBold, ok := objMap["fontBold"]; ok {
		if valFontBold != nil {
			var valueForFontBold string
			err = json.Unmarshal(*valFontBold, &valueForFontBold)
			if err != nil {
				var valueForFontBoldInt int32
				err = json.Unmarshal(*valFontBold, &valueForFontBoldInt)
				if err != nil {
					return err
				}
				this.FontBold = string(valueForFontBoldInt)
			} else {
				this.FontBold = valueForFontBold
			}
		}
	}
	if valFontBoldCap, ok := objMap["FontBold"]; ok {
		if valFontBoldCap != nil {
			var valueForFontBold string
			err = json.Unmarshal(*valFontBoldCap, &valueForFontBold)
			if err != nil {
				var valueForFontBoldInt int32
				err = json.Unmarshal(*valFontBoldCap, &valueForFontBoldInt)
				if err != nil {
					return err
				}
				this.FontBold = string(valueForFontBoldInt)
			} else {
				this.FontBold = valueForFontBold
			}
		}
	}
	this.FontItalic = ""
	if valFontItalic, ok := objMap["fontItalic"]; ok {
		if valFontItalic != nil {
			var valueForFontItalic string
			err = json.Unmarshal(*valFontItalic, &valueForFontItalic)
			if err != nil {
				var valueForFontItalicInt int32
				err = json.Unmarshal(*valFontItalic, &valueForFontItalicInt)
				if err != nil {
					return err
				}
				this.FontItalic = string(valueForFontItalicInt)
			} else {
				this.FontItalic = valueForFontItalic
			}
		}
	}
	if valFontItalicCap, ok := objMap["FontItalic"]; ok {
		if valFontItalicCap != nil {
			var valueForFontItalic string
			err = json.Unmarshal(*valFontItalicCap, &valueForFontItalic)
			if err != nil {
				var valueForFontItalicInt int32
				err = json.Unmarshal(*valFontItalicCap, &valueForFontItalicInt)
				if err != nil {
					return err
				}
				this.FontItalic = string(valueForFontItalicInt)
			} else {
				this.FontItalic = valueForFontItalic
			}
		}
	}
	this.FontUnderline = ""
	if valFontUnderline, ok := objMap["fontUnderline"]; ok {
		if valFontUnderline != nil {
			var valueForFontUnderline string
			err = json.Unmarshal(*valFontUnderline, &valueForFontUnderline)
			if err != nil {
				var valueForFontUnderlineInt int32
				err = json.Unmarshal(*valFontUnderline, &valueForFontUnderlineInt)
				if err != nil {
					return err
				}
				this.FontUnderline = string(valueForFontUnderlineInt)
			} else {
				this.FontUnderline = valueForFontUnderline
			}
		}
	}
	if valFontUnderlineCap, ok := objMap["FontUnderline"]; ok {
		if valFontUnderlineCap != nil {
			var valueForFontUnderline string
			err = json.Unmarshal(*valFontUnderlineCap, &valueForFontUnderline)
			if err != nil {
				var valueForFontUnderlineInt int32
				err = json.Unmarshal(*valFontUnderlineCap, &valueForFontUnderlineInt)
				if err != nil {
					return err
				}
				this.FontUnderline = string(valueForFontUnderlineInt)
			} else {
				this.FontUnderline = valueForFontUnderline
			}
		}
	}
	this.StrikethroughType = ""
	if valStrikethroughType, ok := objMap["strikethroughType"]; ok {
		if valStrikethroughType != nil {
			var valueForStrikethroughType string
			err = json.Unmarshal(*valStrikethroughType, &valueForStrikethroughType)
			if err != nil {
				var valueForStrikethroughTypeInt int32
				err = json.Unmarshal(*valStrikethroughType, &valueForStrikethroughTypeInt)
				if err != nil {
					return err
				}
				this.StrikethroughType = string(valueForStrikethroughTypeInt)
			} else {
				this.StrikethroughType = valueForStrikethroughType
			}
		}
	}
	if valStrikethroughTypeCap, ok := objMap["StrikethroughType"]; ok {
		if valStrikethroughTypeCap != nil {
			var valueForStrikethroughType string
			err = json.Unmarshal(*valStrikethroughTypeCap, &valueForStrikethroughType)
			if err != nil {
				var valueForStrikethroughTypeInt int32
				err = json.Unmarshal(*valStrikethroughTypeCap, &valueForStrikethroughTypeInt)
				if err != nil {
					return err
				}
				this.StrikethroughType = string(valueForStrikethroughTypeInt)
			} else {
				this.StrikethroughType = valueForStrikethroughType
			}
		}
	}
	this.TextCapType = ""
	if valTextCapType, ok := objMap["textCapType"]; ok {
		if valTextCapType != nil {
			var valueForTextCapType string
			err = json.Unmarshal(*valTextCapType, &valueForTextCapType)
			if err != nil {
				var valueForTextCapTypeInt int32
				err = json.Unmarshal(*valTextCapType, &valueForTextCapTypeInt)
				if err != nil {
					return err
				}
				this.TextCapType = string(valueForTextCapTypeInt)
			} else {
				this.TextCapType = valueForTextCapType
			}
		}
	}
	if valTextCapTypeCap, ok := objMap["TextCapType"]; ok {
		if valTextCapTypeCap != nil {
			var valueForTextCapType string
			err = json.Unmarshal(*valTextCapTypeCap, &valueForTextCapType)
			if err != nil {
				var valueForTextCapTypeInt int32
				err = json.Unmarshal(*valTextCapTypeCap, &valueForTextCapTypeInt)
				if err != nil {
					return err
				}
				this.TextCapType = string(valueForTextCapTypeInt)
			} else {
				this.TextCapType = valueForTextCapType
			}
		}
	}
	
	if valEscapement, ok := objMap["escapement"]; ok {
		if valEscapement != nil {
			var valueForEscapement float64
			err = json.Unmarshal(*valEscapement, &valueForEscapement)
			if err != nil {
				return err
			}
			this.Escapement = valueForEscapement
		}
	}
	if valEscapementCap, ok := objMap["Escapement"]; ok {
		if valEscapementCap != nil {
			var valueForEscapement float64
			err = json.Unmarshal(*valEscapementCap, &valueForEscapement)
			if err != nil {
				return err
			}
			this.Escapement = valueForEscapement
		}
	}
	
	if valSpacing, ok := objMap["spacing"]; ok {
		if valSpacing != nil {
			var valueForSpacing float64
			err = json.Unmarshal(*valSpacing, &valueForSpacing)
			if err != nil {
				return err
			}
			this.Spacing = valueForSpacing
		}
	}
	if valSpacingCap, ok := objMap["Spacing"]; ok {
		if valSpacingCap != nil {
			var valueForSpacing float64
			err = json.Unmarshal(*valSpacingCap, &valueForSpacing)
			if err != nil {
				return err
			}
			this.Spacing = valueForSpacing
		}
	}
	
	if valFontColor, ok := objMap["fontColor"]; ok {
		if valFontColor != nil {
			var valueForFontColor string
			err = json.Unmarshal(*valFontColor, &valueForFontColor)
			if err != nil {
				return err
			}
			this.FontColor = valueForFontColor
		}
	}
	if valFontColorCap, ok := objMap["FontColor"]; ok {
		if valFontColorCap != nil {
			var valueForFontColor string
			err = json.Unmarshal(*valFontColorCap, &valueForFontColor)
			if err != nil {
				return err
			}
			this.FontColor = valueForFontColor
		}
	}
	
	if valHighlightColor, ok := objMap["highlightColor"]; ok {
		if valHighlightColor != nil {
			var valueForHighlightColor string
			err = json.Unmarshal(*valHighlightColor, &valueForHighlightColor)
			if err != nil {
				return err
			}
			this.HighlightColor = valueForHighlightColor
		}
	}
	if valHighlightColorCap, ok := objMap["HighlightColor"]; ok {
		if valHighlightColorCap != nil {
			var valueForHighlightColor string
			err = json.Unmarshal(*valHighlightColorCap, &valueForHighlightColor)
			if err != nil {
				return err
			}
			this.HighlightColor = valueForHighlightColor
		}
	}
	
	if valFontHeight, ok := objMap["fontHeight"]; ok {
		if valFontHeight != nil {
			var valueForFontHeight float64
			err = json.Unmarshal(*valFontHeight, &valueForFontHeight)
			if err != nil {
				return err
			}
			this.FontHeight = valueForFontHeight
		}
	}
	if valFontHeightCap, ok := objMap["FontHeight"]; ok {
		if valFontHeightCap != nil {
			var valueForFontHeight float64
			err = json.Unmarshal(*valFontHeightCap, &valueForFontHeight)
			if err != nil {
				return err
			}
			this.FontHeight = valueForFontHeight
		}
	}
	this.NormaliseHeight = ""
	if valNormaliseHeight, ok := objMap["normaliseHeight"]; ok {
		if valNormaliseHeight != nil {
			var valueForNormaliseHeight string
			err = json.Unmarshal(*valNormaliseHeight, &valueForNormaliseHeight)
			if err != nil {
				var valueForNormaliseHeightInt int32
				err = json.Unmarshal(*valNormaliseHeight, &valueForNormaliseHeightInt)
				if err != nil {
					return err
				}
				this.NormaliseHeight = string(valueForNormaliseHeightInt)
			} else {
				this.NormaliseHeight = valueForNormaliseHeight
			}
		}
	}
	if valNormaliseHeightCap, ok := objMap["NormaliseHeight"]; ok {
		if valNormaliseHeightCap != nil {
			var valueForNormaliseHeight string
			err = json.Unmarshal(*valNormaliseHeightCap, &valueForNormaliseHeight)
			if err != nil {
				var valueForNormaliseHeightInt int32
				err = json.Unmarshal(*valNormaliseHeightCap, &valueForNormaliseHeightInt)
				if err != nil {
					return err
				}
				this.NormaliseHeight = string(valueForNormaliseHeightInt)
			} else {
				this.NormaliseHeight = valueForNormaliseHeight
			}
		}
	}
	this.ProofDisabled = ""
	if valProofDisabled, ok := objMap["proofDisabled"]; ok {
		if valProofDisabled != nil {
			var valueForProofDisabled string
			err = json.Unmarshal(*valProofDisabled, &valueForProofDisabled)
			if err != nil {
				var valueForProofDisabledInt int32
				err = json.Unmarshal(*valProofDisabled, &valueForProofDisabledInt)
				if err != nil {
					return err
				}
				this.ProofDisabled = string(valueForProofDisabledInt)
			} else {
				this.ProofDisabled = valueForProofDisabled
			}
		}
	}
	if valProofDisabledCap, ok := objMap["ProofDisabled"]; ok {
		if valProofDisabledCap != nil {
			var valueForProofDisabled string
			err = json.Unmarshal(*valProofDisabledCap, &valueForProofDisabled)
			if err != nil {
				var valueForProofDisabledInt int32
				err = json.Unmarshal(*valProofDisabledCap, &valueForProofDisabledInt)
				if err != nil {
					return err
				}
				this.ProofDisabled = string(valueForProofDisabledInt)
			} else {
				this.ProofDisabled = valueForProofDisabled
			}
		}
	}
	
	if valSmartTagClean, ok := objMap["smartTagClean"]; ok {
		if valSmartTagClean != nil {
			var valueForSmartTagClean bool
			err = json.Unmarshal(*valSmartTagClean, &valueForSmartTagClean)
			if err != nil {
				return err
			}
			this.SmartTagClean = valueForSmartTagClean
		}
	}
	if valSmartTagCleanCap, ok := objMap["SmartTagClean"]; ok {
		if valSmartTagCleanCap != nil {
			var valueForSmartTagClean bool
			err = json.Unmarshal(*valSmartTagCleanCap, &valueForSmartTagClean)
			if err != nil {
				return err
			}
			this.SmartTagClean = valueForSmartTagClean
		}
	}
	
	if valKerningMinimalSize, ok := objMap["kerningMinimalSize"]; ok {
		if valKerningMinimalSize != nil {
			var valueForKerningMinimalSize float64
			err = json.Unmarshal(*valKerningMinimalSize, &valueForKerningMinimalSize)
			if err != nil {
				return err
			}
			this.KerningMinimalSize = valueForKerningMinimalSize
		}
	}
	if valKerningMinimalSizeCap, ok := objMap["KerningMinimalSize"]; ok {
		if valKerningMinimalSizeCap != nil {
			var valueForKerningMinimalSize float64
			err = json.Unmarshal(*valKerningMinimalSizeCap, &valueForKerningMinimalSize)
			if err != nil {
				return err
			}
			this.KerningMinimalSize = valueForKerningMinimalSize
		}
	}
	this.Kumimoji = ""
	if valKumimoji, ok := objMap["kumimoji"]; ok {
		if valKumimoji != nil {
			var valueForKumimoji string
			err = json.Unmarshal(*valKumimoji, &valueForKumimoji)
			if err != nil {
				var valueForKumimojiInt int32
				err = json.Unmarshal(*valKumimoji, &valueForKumimojiInt)
				if err != nil {
					return err
				}
				this.Kumimoji = string(valueForKumimojiInt)
			} else {
				this.Kumimoji = valueForKumimoji
			}
		}
	}
	if valKumimojiCap, ok := objMap["Kumimoji"]; ok {
		if valKumimojiCap != nil {
			var valueForKumimoji string
			err = json.Unmarshal(*valKumimojiCap, &valueForKumimoji)
			if err != nil {
				var valueForKumimojiInt int32
				err = json.Unmarshal(*valKumimojiCap, &valueForKumimojiInt)
				if err != nil {
					return err
				}
				this.Kumimoji = string(valueForKumimojiInt)
			} else {
				this.Kumimoji = valueForKumimoji
			}
		}
	}
	
	if valLanguageId, ok := objMap["languageId"]; ok {
		if valLanguageId != nil {
			var valueForLanguageId string
			err = json.Unmarshal(*valLanguageId, &valueForLanguageId)
			if err != nil {
				return err
			}
			this.LanguageId = valueForLanguageId
		}
	}
	if valLanguageIdCap, ok := objMap["LanguageId"]; ok {
		if valLanguageIdCap != nil {
			var valueForLanguageId string
			err = json.Unmarshal(*valLanguageIdCap, &valueForLanguageId)
			if err != nil {
				return err
			}
			this.LanguageId = valueForLanguageId
		}
	}
	
	if valAlternativeLanguageId, ok := objMap["alternativeLanguageId"]; ok {
		if valAlternativeLanguageId != nil {
			var valueForAlternativeLanguageId string
			err = json.Unmarshal(*valAlternativeLanguageId, &valueForAlternativeLanguageId)
			if err != nil {
				return err
			}
			this.AlternativeLanguageId = valueForAlternativeLanguageId
		}
	}
	if valAlternativeLanguageIdCap, ok := objMap["AlternativeLanguageId"]; ok {
		if valAlternativeLanguageIdCap != nil {
			var valueForAlternativeLanguageId string
			err = json.Unmarshal(*valAlternativeLanguageIdCap, &valueForAlternativeLanguageId)
			if err != nil {
				return err
			}
			this.AlternativeLanguageId = valueForAlternativeLanguageId
		}
	}
	this.IsHardUnderlineFill = ""
	if valIsHardUnderlineFill, ok := objMap["isHardUnderlineFill"]; ok {
		if valIsHardUnderlineFill != nil {
			var valueForIsHardUnderlineFill string
			err = json.Unmarshal(*valIsHardUnderlineFill, &valueForIsHardUnderlineFill)
			if err != nil {
				var valueForIsHardUnderlineFillInt int32
				err = json.Unmarshal(*valIsHardUnderlineFill, &valueForIsHardUnderlineFillInt)
				if err != nil {
					return err
				}
				this.IsHardUnderlineFill = string(valueForIsHardUnderlineFillInt)
			} else {
				this.IsHardUnderlineFill = valueForIsHardUnderlineFill
			}
		}
	}
	if valIsHardUnderlineFillCap, ok := objMap["IsHardUnderlineFill"]; ok {
		if valIsHardUnderlineFillCap != nil {
			var valueForIsHardUnderlineFill string
			err = json.Unmarshal(*valIsHardUnderlineFillCap, &valueForIsHardUnderlineFill)
			if err != nil {
				var valueForIsHardUnderlineFillInt int32
				err = json.Unmarshal(*valIsHardUnderlineFillCap, &valueForIsHardUnderlineFillInt)
				if err != nil {
					return err
				}
				this.IsHardUnderlineFill = string(valueForIsHardUnderlineFillInt)
			} else {
				this.IsHardUnderlineFill = valueForIsHardUnderlineFill
			}
		}
	}
	this.IsHardUnderlineLine = ""
	if valIsHardUnderlineLine, ok := objMap["isHardUnderlineLine"]; ok {
		if valIsHardUnderlineLine != nil {
			var valueForIsHardUnderlineLine string
			err = json.Unmarshal(*valIsHardUnderlineLine, &valueForIsHardUnderlineLine)
			if err != nil {
				var valueForIsHardUnderlineLineInt int32
				err = json.Unmarshal(*valIsHardUnderlineLine, &valueForIsHardUnderlineLineInt)
				if err != nil {
					return err
				}
				this.IsHardUnderlineLine = string(valueForIsHardUnderlineLineInt)
			} else {
				this.IsHardUnderlineLine = valueForIsHardUnderlineLine
			}
		}
	}
	if valIsHardUnderlineLineCap, ok := objMap["IsHardUnderlineLine"]; ok {
		if valIsHardUnderlineLineCap != nil {
			var valueForIsHardUnderlineLine string
			err = json.Unmarshal(*valIsHardUnderlineLineCap, &valueForIsHardUnderlineLine)
			if err != nil {
				var valueForIsHardUnderlineLineInt int32
				err = json.Unmarshal(*valIsHardUnderlineLineCap, &valueForIsHardUnderlineLineInt)
				if err != nil {
					return err
				}
				this.IsHardUnderlineLine = string(valueForIsHardUnderlineLineInt)
			} else {
				this.IsHardUnderlineLine = valueForIsHardUnderlineLine
			}
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	
	if valUnderlineFillFormat, ok := objMap["underlineFillFormat"]; ok {
		if valUnderlineFillFormat != nil {
			var valueForUnderlineFillFormat FillFormat
			err = json.Unmarshal(*valUnderlineFillFormat, &valueForUnderlineFillFormat)
			if err != nil {
				return err
			}
			this.UnderlineFillFormat = &valueForUnderlineFillFormat
		}
	}
	if valUnderlineFillFormatCap, ok := objMap["UnderlineFillFormat"]; ok {
		if valUnderlineFillFormatCap != nil {
			var valueForUnderlineFillFormat FillFormat
			err = json.Unmarshal(*valUnderlineFillFormatCap, &valueForUnderlineFillFormat)
			if err != nil {
				return err
			}
			this.UnderlineFillFormat = &valueForUnderlineFillFormat
		}
	}
	
	if valUnderlineLineFormat, ok := objMap["underlineLineFormat"]; ok {
		if valUnderlineLineFormat != nil {
			var valueForUnderlineLineFormat LineFormat
			err = json.Unmarshal(*valUnderlineLineFormat, &valueForUnderlineLineFormat)
			if err != nil {
				return err
			}
			this.UnderlineLineFormat = &valueForUnderlineLineFormat
		}
	}
	if valUnderlineLineFormatCap, ok := objMap["UnderlineLineFormat"]; ok {
		if valUnderlineLineFormatCap != nil {
			var valueForUnderlineLineFormat LineFormat
			err = json.Unmarshal(*valUnderlineLineFormatCap, &valueForUnderlineLineFormat)
			if err != nil {
				return err
			}
			this.UnderlineLineFormat = &valueForUnderlineLineFormat
		}
	}

    return nil
}
