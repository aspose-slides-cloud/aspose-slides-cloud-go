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

// Provides options that control how a presentation is saved in XAML format.
type IXamlExportOptions interface {

	// Default regular font for rendering the presentation. 
	getDefaultRegularFont() string
	setDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	getFontFallbackRules() []IFontFallbackRule
	setFontFallbackRules(newValue []IFontFallbackRule)

	// Export format.
	getFormat() string
	setFormat(newValue string)

	// Export hidden slides
	getExportHiddenSlides() bool
	setExportHiddenSlides(newValue bool)
}

type XamlExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Export hidden slides
	ExportHiddenSlides bool `json:"ExportHiddenSlides"`
}

func NewXamlExportOptions() *XamlExportOptions {
	instance := new(XamlExportOptions)
	return instance
}

func (this *XamlExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *XamlExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *XamlExportOptions) getFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *XamlExportOptions) setFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *XamlExportOptions) getFormat() string {
	return this.Format
}

func (this *XamlExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *XamlExportOptions) getExportHiddenSlides() bool {
	return this.ExportHiddenSlides
}

func (this *XamlExportOptions) setExportHiddenSlides(newValue bool) {
	this.ExportHiddenSlides = newValue
}

func (this *XamlExportOptions) UnmarshalJSON(b []byte) error {
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

	return nil
}