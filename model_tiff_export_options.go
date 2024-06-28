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

// Provides options that control how a presentation is saved in TIFF format.
type ITiffExportOptions interface {

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Gets of sets list of font substitution rules.
	GetFontSubstRules() []IFontSubstRule
	SetFontSubstRules(newValue []IFontSubstRule)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Gets or sets the height of slides in the output image format.
	GetHeight() int32
	SetHeight(newValue int32)

	// Gets or sets the height of slides in the output the output image format.
	GetWidth() int32
	SetWidth(newValue int32)

	// Compression type.
	GetCompression() string
	SetCompression(newValue string)

	// Horizontal resolution, in dots per inch.
	GetDpiX() int32
	SetDpiX(newValue int32)

	// Vertical resolution, in dots per inch.
	GetDpiY() int32
	SetDpiY(newValue int32)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)

	// Specifies the pixel format for the generated images. Read/write ImagePixelFormat.
	GetPixelFormat() string
	SetPixelFormat(newValue string)

	// Slides layouting options
	GetSlidesLayoutOptions() ISlidesLayoutOptions
	SetSlidesLayoutOptions(newValue ISlidesLayoutOptions)

	// Specifies the algorithm for converting a color image into a black and white image. This option will applied only if Aspose.Slides.Export.TiffOptions.CompressionType is set to Aspose.Slides.Export.TiffCompressionTypes.CCITT4 or Aspose.Slides.Export.TiffCompressionTypes.CCITT3.
	GetBwConversionMode() string
	SetBwConversionMode(newValue string)
}

type TiffExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Gets or sets the height of slides in the output image format.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output the output image format.
	Width int32 `json:"Width,omitempty"`

	// Compression type.
	Compression string `json:"Compression,omitempty"`

	// Horizontal resolution, in dots per inch.
	DpiX int32 `json:"DpiX,omitempty"`

	// Vertical resolution, in dots per inch.
	DpiY int32 `json:"DpiY,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`

	// Specifies the pixel format for the generated images. Read/write ImagePixelFormat.
	PixelFormat string `json:"PixelFormat,omitempty"`

	// Slides layouting options
	SlidesLayoutOptions ISlidesLayoutOptions `json:"SlidesLayoutOptions,omitempty"`

	// Specifies the algorithm for converting a color image into a black and white image. This option will applied only if Aspose.Slides.Export.TiffOptions.CompressionType is set to Aspose.Slides.Export.TiffCompressionTypes.CCITT4 or Aspose.Slides.Export.TiffCompressionTypes.CCITT3.
	BwConversionMode string `json:"BwConversionMode,omitempty"`
}

func NewTiffExportOptions() *TiffExportOptions {
	instance := new(TiffExportOptions)
	return instance
}

func (this *TiffExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *TiffExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *TiffExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *TiffExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *TiffExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *TiffExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *TiffExportOptions) GetFormat() string {
	return this.Format
}

func (this *TiffExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *TiffExportOptions) GetHeight() int32 {
	return this.Height
}

func (this *TiffExportOptions) SetHeight(newValue int32) {
	this.Height = newValue
}
func (this *TiffExportOptions) GetWidth() int32 {
	return this.Width
}

func (this *TiffExportOptions) SetWidth(newValue int32) {
	this.Width = newValue
}
func (this *TiffExportOptions) GetCompression() string {
	return this.Compression
}

func (this *TiffExportOptions) SetCompression(newValue string) {
	this.Compression = newValue
}
func (this *TiffExportOptions) GetDpiX() int32 {
	return this.DpiX
}

func (this *TiffExportOptions) SetDpiX(newValue int32) {
	this.DpiX = newValue
}
func (this *TiffExportOptions) GetDpiY() int32 {
	return this.DpiY
}

func (this *TiffExportOptions) SetDpiY(newValue int32) {
	this.DpiY = newValue
}
func (this *TiffExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *TiffExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}
func (this *TiffExportOptions) GetPixelFormat() string {
	return this.PixelFormat
}

func (this *TiffExportOptions) SetPixelFormat(newValue string) {
	this.PixelFormat = newValue
}
func (this *TiffExportOptions) GetSlidesLayoutOptions() ISlidesLayoutOptions {
	return this.SlidesLayoutOptions
}

func (this *TiffExportOptions) SetSlidesLayoutOptions(newValue ISlidesLayoutOptions) {
	this.SlidesLayoutOptions = newValue
}
func (this *TiffExportOptions) GetBwConversionMode() string {
	return this.BwConversionMode
}

func (this *TiffExportOptions) SetBwConversionMode(newValue string) {
	this.BwConversionMode = newValue
}

func (this *TiffExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDefaultRegularFont, ok := objMap["defaultRegularFont"]; ok {
		if valDefaultRegularFont != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFont, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	if valDefaultRegularFontCap, ok := objMap["DefaultRegularFont"]; ok {
		if valDefaultRegularFontCap != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFontCap, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	
	if valFontFallbackRules, ok := objMap["fontFallbackRules"]; ok {
		if valFontFallbackRules != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRules, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
	}
	if valFontFallbackRulesCap, ok := objMap["FontFallbackRules"]; ok {
		if valFontFallbackRulesCap != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRulesCap, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
	}
	
	if valFontSubstRules, ok := objMap["fontSubstRules"]; ok {
		if valFontSubstRules != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRules, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
		}
	}
	if valFontSubstRulesCap, ok := objMap["FontSubstRules"]; ok {
		if valFontSubstRulesCap != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRulesCap, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
		}
	}
	
	if valFormat, ok := objMap["format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	if valFormatCap, ok := objMap["Format"]; ok {
		if valFormatCap != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormatCap, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valCompression, ok := objMap["compression"]; ok {
		if valCompression != nil {
			var valueForCompression string
			err = json.Unmarshal(*valCompression, &valueForCompression)
			if err != nil {
				var valueForCompressionInt int32
				err = json.Unmarshal(*valCompression, &valueForCompressionInt)
				if err != nil {
					return err
				}
				this.Compression = string(valueForCompressionInt)
			} else {
				this.Compression = valueForCompression
			}
		}
	}
	if valCompressionCap, ok := objMap["Compression"]; ok {
		if valCompressionCap != nil {
			var valueForCompression string
			err = json.Unmarshal(*valCompressionCap, &valueForCompression)
			if err != nil {
				var valueForCompressionInt int32
				err = json.Unmarshal(*valCompressionCap, &valueForCompressionInt)
				if err != nil {
					return err
				}
				this.Compression = string(valueForCompressionInt)
			} else {
				this.Compression = valueForCompression
			}
		}
	}
	
	if valDpiX, ok := objMap["dpiX"]; ok {
		if valDpiX != nil {
			var valueForDpiX int32
			err = json.Unmarshal(*valDpiX, &valueForDpiX)
			if err != nil {
				return err
			}
			this.DpiX = valueForDpiX
		}
	}
	if valDpiXCap, ok := objMap["DpiX"]; ok {
		if valDpiXCap != nil {
			var valueForDpiX int32
			err = json.Unmarshal(*valDpiXCap, &valueForDpiX)
			if err != nil {
				return err
			}
			this.DpiX = valueForDpiX
		}
	}
	
	if valDpiY, ok := objMap["dpiY"]; ok {
		if valDpiY != nil {
			var valueForDpiY int32
			err = json.Unmarshal(*valDpiY, &valueForDpiY)
			if err != nil {
				return err
			}
			this.DpiY = valueForDpiY
		}
	}
	if valDpiYCap, ok := objMap["DpiY"]; ok {
		if valDpiYCap != nil {
			var valueForDpiY int32
			err = json.Unmarshal(*valDpiYCap, &valueForDpiY)
			if err != nil {
				return err
			}
			this.DpiY = valueForDpiY
		}
	}
	
	if valShowHiddenSlides, ok := objMap["showHiddenSlides"]; ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides *bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	if valShowHiddenSlidesCap, ok := objMap["ShowHiddenSlides"]; ok {
		if valShowHiddenSlidesCap != nil {
			var valueForShowHiddenSlides *bool
			err = json.Unmarshal(*valShowHiddenSlidesCap, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	
	if valPixelFormat, ok := objMap["pixelFormat"]; ok {
		if valPixelFormat != nil {
			var valueForPixelFormat string
			err = json.Unmarshal(*valPixelFormat, &valueForPixelFormat)
			if err != nil {
				var valueForPixelFormatInt int32
				err = json.Unmarshal(*valPixelFormat, &valueForPixelFormatInt)
				if err != nil {
					return err
				}
				this.PixelFormat = string(valueForPixelFormatInt)
			} else {
				this.PixelFormat = valueForPixelFormat
			}
		}
	}
	if valPixelFormatCap, ok := objMap["PixelFormat"]; ok {
		if valPixelFormatCap != nil {
			var valueForPixelFormat string
			err = json.Unmarshal(*valPixelFormatCap, &valueForPixelFormat)
			if err != nil {
				var valueForPixelFormatInt int32
				err = json.Unmarshal(*valPixelFormatCap, &valueForPixelFormatInt)
				if err != nil {
					return err
				}
				this.PixelFormat = string(valueForPixelFormatInt)
			} else {
				this.PixelFormat = valueForPixelFormat
			}
		}
	}
	
	if valSlidesLayoutOptions, ok := objMap["slidesLayoutOptions"]; ok {
		if valSlidesLayoutOptions != nil {
			var valueForSlidesLayoutOptions SlidesLayoutOptions
			err = json.Unmarshal(*valSlidesLayoutOptions, &valueForSlidesLayoutOptions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SlidesLayoutOptions", *valSlidesLayoutOptions)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlidesLayoutOptions, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISlidesLayoutOptions)
			if ok {
				this.SlidesLayoutOptions = vInterfaceObject
			}
		}
	}
	if valSlidesLayoutOptionsCap, ok := objMap["SlidesLayoutOptions"]; ok {
		if valSlidesLayoutOptionsCap != nil {
			var valueForSlidesLayoutOptions SlidesLayoutOptions
			err = json.Unmarshal(*valSlidesLayoutOptionsCap, &valueForSlidesLayoutOptions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SlidesLayoutOptions", *valSlidesLayoutOptionsCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlidesLayoutOptionsCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISlidesLayoutOptions)
			if ok {
				this.SlidesLayoutOptions = vInterfaceObject
			}
		}
	}
	
	if valBwConversionMode, ok := objMap["bwConversionMode"]; ok {
		if valBwConversionMode != nil {
			var valueForBwConversionMode string
			err = json.Unmarshal(*valBwConversionMode, &valueForBwConversionMode)
			if err != nil {
				var valueForBwConversionModeInt int32
				err = json.Unmarshal(*valBwConversionMode, &valueForBwConversionModeInt)
				if err != nil {
					return err
				}
				this.BwConversionMode = string(valueForBwConversionModeInt)
			} else {
				this.BwConversionMode = valueForBwConversionMode
			}
		}
	}
	if valBwConversionModeCap, ok := objMap["BwConversionMode"]; ok {
		if valBwConversionModeCap != nil {
			var valueForBwConversionMode string
			err = json.Unmarshal(*valBwConversionModeCap, &valueForBwConversionMode)
			if err != nil {
				var valueForBwConversionModeInt int32
				err = json.Unmarshal(*valBwConversionModeCap, &valueForBwConversionModeInt)
				if err != nil {
					return err
				}
				this.BwConversionMode = string(valueForBwConversionModeInt)
			} else {
				this.BwConversionMode = valueForBwConversionMode
			}
		}
	}

	return nil
}
