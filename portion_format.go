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

// Represents portion format.
type IPortionFormat interface {

	// True for bold font.
	GetFontBold() string
	SetFontBold(newValue string)

	// True for italic font.
	GetFontItalic() string
	SetFontItalic(newValue string)

	// Text underline type.
	GetFontUnderline() string
	SetFontUnderline(newValue string)

	// Text strikethrough type.
	GetStrikethroughType() string
	SetStrikethroughType(newValue string)

	// Text capitalization type.
	GetTextCapType() string
	SetTextCapType(newValue string)

	// Superscript or subscript of the text.
	GetEscapement() float64
	SetEscapement(newValue float64)

	// Intercharacter spacing increment.
	GetSpacing() float64
	SetSpacing(newValue float64)

	// Font color.
	GetFontColor() string
	SetFontColor(newValue string)

	// Highlight color.
	GetHighlightColor() string
	SetHighlightColor(newValue string)

	// Font height.
	GetFontHeight() float64
	SetFontHeight(newValue float64)

	// True to normalize the text.
	GetNormaliseHeight() string
	SetNormaliseHeight(newValue string)

	// True if the text proof should be disabled.
	GetProofDisabled() string
	SetProofDisabled(newValue string)

	// True if smart tag should be cleaned.
	GetSmartTagClean() bool
	SetSmartTagClean(newValue bool)

	// Minimal font size for kerning.
	GetKerningMinimalSize() float64
	SetKerningMinimalSize(newValue float64)

	// True if numbers should ignore East-Asian specific vertical text layout.
	GetKumimoji() string
	SetKumimoji(newValue string)

	// Proving language ID.
	GetLanguageId() string
	SetLanguageId(newValue string)

	// Alternative proving language ID.
	GetAlternativeLanguageId() string
	SetAlternativeLanguageId(newValue string)

	// True if underline style has own FillFormat properties.
	GetIsHardUnderlineFill() string
	SetIsHardUnderlineFill(newValue string)

	// True if underline style has own LineFormat properties.
	GetIsHardUnderlineLine() string
	SetIsHardUnderlineLine(newValue string)

	// Fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Underline fill format.
	GetUnderlineFillFormat() IFillFormat
	SetUnderlineFillFormat(newValue IFillFormat)

	// Underline line format.
	GetUnderlineLineFormat() ILineFormat
	SetUnderlineLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	GetHyperlinkClick() IHyperlink
	SetHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	GetHyperlinkMouseOver() IHyperlink
	SetHyperlinkMouseOver(newValue IHyperlink)

	// Returns or sets the Latin font info.
	GetLatinFont() string
	SetLatinFont(newValue string)

	// Returns or sets the East Asian font info.
	GetEastAsianFont() string
	SetEastAsianFont(newValue string)

	// Returns or sets the complex script font info.
	GetComplexScriptFont() string
	SetComplexScriptFont(newValue string)
}

type PortionFormat struct {

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

	// Hyperlink defined for mouse click.
	HyperlinkClick IHyperlink `json:"HyperlinkClick,omitempty"`

	// Hyperlink defined for mouse over.
	HyperlinkMouseOver IHyperlink `json:"HyperlinkMouseOver,omitempty"`

	// Returns or sets the Latin font info.
	LatinFont string `json:"LatinFont,omitempty"`

	// Returns or sets the East Asian font info.
	EastAsianFont string `json:"EastAsianFont,omitempty"`

	// Returns or sets the complex script font info.
	ComplexScriptFont string `json:"ComplexScriptFont,omitempty"`
}

func NewPortionFormat() *PortionFormat {
	instance := new(PortionFormat)
	instance.FontBold = ""
	instance.FontItalic = ""
	instance.FontUnderline = ""
	instance.StrikethroughType = ""
	instance.TextCapType = ""
	instance.NormaliseHeight = ""
	instance.ProofDisabled = ""
	instance.Kumimoji = ""
	instance.IsHardUnderlineFill = ""
	instance.IsHardUnderlineLine = ""
	return instance
}

func (this *PortionFormat) GetFontBold() string {
	return this.FontBold
}

func (this *PortionFormat) SetFontBold(newValue string) {
	this.FontBold = newValue
}
func (this *PortionFormat) GetFontItalic() string {
	return this.FontItalic
}

func (this *PortionFormat) SetFontItalic(newValue string) {
	this.FontItalic = newValue
}
func (this *PortionFormat) GetFontUnderline() string {
	return this.FontUnderline
}

func (this *PortionFormat) SetFontUnderline(newValue string) {
	this.FontUnderline = newValue
}
func (this *PortionFormat) GetStrikethroughType() string {
	return this.StrikethroughType
}

func (this *PortionFormat) SetStrikethroughType(newValue string) {
	this.StrikethroughType = newValue
}
func (this *PortionFormat) GetTextCapType() string {
	return this.TextCapType
}

func (this *PortionFormat) SetTextCapType(newValue string) {
	this.TextCapType = newValue
}
func (this *PortionFormat) GetEscapement() float64 {
	return this.Escapement
}

func (this *PortionFormat) SetEscapement(newValue float64) {
	this.Escapement = newValue
}
func (this *PortionFormat) GetSpacing() float64 {
	return this.Spacing
}

func (this *PortionFormat) SetSpacing(newValue float64) {
	this.Spacing = newValue
}
func (this *PortionFormat) GetFontColor() string {
	return this.FontColor
}

func (this *PortionFormat) SetFontColor(newValue string) {
	this.FontColor = newValue
}
func (this *PortionFormat) GetHighlightColor() string {
	return this.HighlightColor
}

func (this *PortionFormat) SetHighlightColor(newValue string) {
	this.HighlightColor = newValue
}
func (this *PortionFormat) GetFontHeight() float64 {
	return this.FontHeight
}

func (this *PortionFormat) SetFontHeight(newValue float64) {
	this.FontHeight = newValue
}
func (this *PortionFormat) GetNormaliseHeight() string {
	return this.NormaliseHeight
}

func (this *PortionFormat) SetNormaliseHeight(newValue string) {
	this.NormaliseHeight = newValue
}
func (this *PortionFormat) GetProofDisabled() string {
	return this.ProofDisabled
}

func (this *PortionFormat) SetProofDisabled(newValue string) {
	this.ProofDisabled = newValue
}
func (this *PortionFormat) GetSmartTagClean() bool {
	return this.SmartTagClean
}

func (this *PortionFormat) SetSmartTagClean(newValue bool) {
	this.SmartTagClean = newValue
}
func (this *PortionFormat) GetKerningMinimalSize() float64 {
	return this.KerningMinimalSize
}

func (this *PortionFormat) SetKerningMinimalSize(newValue float64) {
	this.KerningMinimalSize = newValue
}
func (this *PortionFormat) GetKumimoji() string {
	return this.Kumimoji
}

func (this *PortionFormat) SetKumimoji(newValue string) {
	this.Kumimoji = newValue
}
func (this *PortionFormat) GetLanguageId() string {
	return this.LanguageId
}

func (this *PortionFormat) SetLanguageId(newValue string) {
	this.LanguageId = newValue
}
func (this *PortionFormat) GetAlternativeLanguageId() string {
	return this.AlternativeLanguageId
}

func (this *PortionFormat) SetAlternativeLanguageId(newValue string) {
	this.AlternativeLanguageId = newValue
}
func (this *PortionFormat) GetIsHardUnderlineFill() string {
	return this.IsHardUnderlineFill
}

func (this *PortionFormat) SetIsHardUnderlineFill(newValue string) {
	this.IsHardUnderlineFill = newValue
}
func (this *PortionFormat) GetIsHardUnderlineLine() string {
	return this.IsHardUnderlineLine
}

func (this *PortionFormat) SetIsHardUnderlineLine(newValue string) {
	this.IsHardUnderlineLine = newValue
}
func (this *PortionFormat) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *PortionFormat) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *PortionFormat) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *PortionFormat) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *PortionFormat) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *PortionFormat) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *PortionFormat) GetUnderlineFillFormat() IFillFormat {
	return this.UnderlineFillFormat
}

func (this *PortionFormat) SetUnderlineFillFormat(newValue IFillFormat) {
	this.UnderlineFillFormat = newValue
}
func (this *PortionFormat) GetUnderlineLineFormat() ILineFormat {
	return this.UnderlineLineFormat
}

func (this *PortionFormat) SetUnderlineLineFormat(newValue ILineFormat) {
	this.UnderlineLineFormat = newValue
}
func (this *PortionFormat) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *PortionFormat) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *PortionFormat) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *PortionFormat) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *PortionFormat) GetLatinFont() string {
	return this.LatinFont
}

func (this *PortionFormat) SetLatinFont(newValue string) {
	this.LatinFont = newValue
}
func (this *PortionFormat) GetEastAsianFont() string {
	return this.EastAsianFont
}

func (this *PortionFormat) SetEastAsianFont(newValue string) {
	this.EastAsianFont = newValue
}
func (this *PortionFormat) GetComplexScriptFont() string {
	return this.ComplexScriptFont
}

func (this *PortionFormat) SetComplexScriptFont(newValue string) {
	this.ComplexScriptFont = newValue
}

func (this *PortionFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	
	if valUnderlineFillFormat, ok := objMap["underlineFillFormat"]; ok {
		if valUnderlineFillFormat != nil {
			var valueForUnderlineFillFormat FillFormat
			err = json.Unmarshal(*valUnderlineFillFormat, &valueForUnderlineFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valUnderlineFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUnderlineFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.UnderlineFillFormat = vInterfaceObject
			}
		}
	}
	if valUnderlineFillFormatCap, ok := objMap["UnderlineFillFormat"]; ok {
		if valUnderlineFillFormatCap != nil {
			var valueForUnderlineFillFormat FillFormat
			err = json.Unmarshal(*valUnderlineFillFormatCap, &valueForUnderlineFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valUnderlineFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUnderlineFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.UnderlineFillFormat = vInterfaceObject
			}
		}
	}
	
	if valUnderlineLineFormat, ok := objMap["underlineLineFormat"]; ok {
		if valUnderlineLineFormat != nil {
			var valueForUnderlineLineFormat LineFormat
			err = json.Unmarshal(*valUnderlineLineFormat, &valueForUnderlineLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valUnderlineLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUnderlineLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.UnderlineLineFormat = vInterfaceObject
			}
		}
	}
	if valUnderlineLineFormatCap, ok := objMap["UnderlineLineFormat"]; ok {
		if valUnderlineLineFormatCap != nil {
			var valueForUnderlineLineFormat LineFormat
			err = json.Unmarshal(*valUnderlineLineFormatCap, &valueForUnderlineLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valUnderlineLineFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUnderlineLineFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.UnderlineLineFormat = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkClick, ok := objMap["hyperlinkClick"]; ok {
		if valHyperlinkClick != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClick, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClick)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClick, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	if valHyperlinkClickCap, ok := objMap["HyperlinkClick"]; ok {
		if valHyperlinkClickCap != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClickCap, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClickCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClickCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkMouseOver, ok := objMap["hyperlinkMouseOver"]; ok {
		if valHyperlinkMouseOver != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOver, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOver)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOver, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	if valHyperlinkMouseOverCap, ok := objMap["HyperlinkMouseOver"]; ok {
		if valHyperlinkMouseOverCap != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOverCap, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOverCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOverCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	
	if valLatinFont, ok := objMap["latinFont"]; ok {
		if valLatinFont != nil {
			var valueForLatinFont string
			err = json.Unmarshal(*valLatinFont, &valueForLatinFont)
			if err != nil {
				return err
			}
			this.LatinFont = valueForLatinFont
		}
	}
	if valLatinFontCap, ok := objMap["LatinFont"]; ok {
		if valLatinFontCap != nil {
			var valueForLatinFont string
			err = json.Unmarshal(*valLatinFontCap, &valueForLatinFont)
			if err != nil {
				return err
			}
			this.LatinFont = valueForLatinFont
		}
	}
	
	if valEastAsianFont, ok := objMap["eastAsianFont"]; ok {
		if valEastAsianFont != nil {
			var valueForEastAsianFont string
			err = json.Unmarshal(*valEastAsianFont, &valueForEastAsianFont)
			if err != nil {
				return err
			}
			this.EastAsianFont = valueForEastAsianFont
		}
	}
	if valEastAsianFontCap, ok := objMap["EastAsianFont"]; ok {
		if valEastAsianFontCap != nil {
			var valueForEastAsianFont string
			err = json.Unmarshal(*valEastAsianFontCap, &valueForEastAsianFont)
			if err != nil {
				return err
			}
			this.EastAsianFont = valueForEastAsianFont
		}
	}
	
	if valComplexScriptFont, ok := objMap["complexScriptFont"]; ok {
		if valComplexScriptFont != nil {
			var valueForComplexScriptFont string
			err = json.Unmarshal(*valComplexScriptFont, &valueForComplexScriptFont)
			if err != nil {
				return err
			}
			this.ComplexScriptFont = valueForComplexScriptFont
		}
	}
	if valComplexScriptFontCap, ok := objMap["ComplexScriptFont"]; ok {
		if valComplexScriptFontCap != nil {
			var valueForComplexScriptFont string
			err = json.Unmarshal(*valComplexScriptFontCap, &valueForComplexScriptFont)
			if err != nil {
				return err
			}
			this.ComplexScriptFont = valueForComplexScriptFont
		}
	}

	return nil
}
