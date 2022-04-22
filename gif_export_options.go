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

// Provides options that control how a presentation is saved in Gif format.
type IGifExportOptions interface {

	// Default regular font for rendering the presentation. 
	getDefaultRegularFont() string
	setDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	getFontFallbackRules() []IFontFallbackRule
	setFontFallbackRules(newValue []IFontFallbackRule)

	// Export format.
	getFormat() string
	setFormat(newValue string)

	// Gets or sets the height of slides in the output image format.
	getHeight() int32
	setHeight(newValue int32)

	// Gets or sets the height of slides in the output the output image format.
	getWidth() int32
	setWidth(newValue int32)

	// Determines whether hidden slides will be exported.
	getExportHiddenSlides() bool
	setExportHiddenSlides(newValue bool)

	// Gets or sets transition FPS [frames/sec]
	getTransitionFps() int32
	setTransitionFps(newValue int32)

	// Gets or sets default delay time [ms].
	getDefaultDelay() int32
	setDefaultDelay(newValue int32)
}

type GifExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Gets or sets the height of slides in the output image format.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output the output image format.
	Width int32 `json:"Width,omitempty"`

	// Determines whether hidden slides will be exported.
	ExportHiddenSlides bool `json:"ExportHiddenSlides"`

	// Gets or sets transition FPS [frames/sec]
	TransitionFps int32 `json:"TransitionFps,omitempty"`

	// Gets or sets default delay time [ms].
	DefaultDelay int32 `json:"DefaultDelay,omitempty"`
}

func NewGifExportOptions() *GifExportOptions {
	instance := new(GifExportOptions)
	return instance
}

func (this *GifExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *GifExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *GifExportOptions) getFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *GifExportOptions) setFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *GifExportOptions) getFormat() string {
	return this.Format
}

func (this *GifExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *GifExportOptions) getHeight() int32 {
	return this.Height
}

func (this *GifExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *GifExportOptions) getWidth() int32 {
	return this.Width
}

func (this *GifExportOptions) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *GifExportOptions) getExportHiddenSlides() bool {
	return this.ExportHiddenSlides
}

func (this *GifExportOptions) setExportHiddenSlides(newValue bool) {
	this.ExportHiddenSlides = newValue
}
func (this *GifExportOptions) getTransitionFps() int32 {
	return this.TransitionFps
}

func (this *GifExportOptions) setTransitionFps(newValue int32) {
	this.TransitionFps = newValue
}
func (this *GifExportOptions) getDefaultDelay() int32 {
	return this.DefaultDelay
}

func (this *GifExportOptions) setDefaultDelay(newValue int32) {
	this.DefaultDelay = newValue
}

func (this *GifExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valExportHiddenSlides, ok := objMap["exportHiddenSlides"]; ok {
		if valExportHiddenSlides != nil {
			var valueForExportHiddenSlides bool
			err = json.Unmarshal(*valExportHiddenSlides, &valueForExportHiddenSlides)
			if err != nil {
				return err
			}
			this.ExportHiddenSlides = valueForExportHiddenSlides
		}
	}
	if valExportHiddenSlidesCap, ok := objMap["ExportHiddenSlides"]; ok {
		if valExportHiddenSlidesCap != nil {
			var valueForExportHiddenSlides bool
			err = json.Unmarshal(*valExportHiddenSlidesCap, &valueForExportHiddenSlides)
			if err != nil {
				return err
			}
			this.ExportHiddenSlides = valueForExportHiddenSlides
		}
	}
	
	if valTransitionFps, ok := objMap["transitionFps"]; ok {
		if valTransitionFps != nil {
			var valueForTransitionFps int32
			err = json.Unmarshal(*valTransitionFps, &valueForTransitionFps)
			if err != nil {
				return err
			}
			this.TransitionFps = valueForTransitionFps
		}
	}
	if valTransitionFpsCap, ok := objMap["TransitionFps"]; ok {
		if valTransitionFpsCap != nil {
			var valueForTransitionFps int32
			err = json.Unmarshal(*valTransitionFpsCap, &valueForTransitionFps)
			if err != nil {
				return err
			}
			this.TransitionFps = valueForTransitionFps
		}
	}
	
	if valDefaultDelay, ok := objMap["defaultDelay"]; ok {
		if valDefaultDelay != nil {
			var valueForDefaultDelay int32
			err = json.Unmarshal(*valDefaultDelay, &valueForDefaultDelay)
			if err != nil {
				return err
			}
			this.DefaultDelay = valueForDefaultDelay
		}
	}
	if valDefaultDelayCap, ok := objMap["DefaultDelay"]; ok {
		if valDefaultDelayCap != nil {
			var valueForDefaultDelay int32
			err = json.Unmarshal(*valDefaultDelayCap, &valueForDefaultDelay)
			if err != nil {
				return err
			}
			this.DefaultDelay = valueForDefaultDelay
		}
	}

	return nil
}
