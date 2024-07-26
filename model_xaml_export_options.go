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
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// Default regular font for rendering the presentation. 
	GetGradientStyle() string
	SetGradientStyle(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Gets of sets list of font substitution rules.
	GetFontSubstRules() []IFontSubstRule
	SetFontSubstRules(newValue []IFontSubstRule)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Export hidden slides
	GetExportHiddenSlides() *bool
	SetExportHiddenSlides(newValue *bool)
}

type XamlExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Default regular font for rendering the presentation. 
	GradientStyle string `json:"GradientStyle,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Export hidden slides
	ExportHiddenSlides *bool `json:"ExportHiddenSlides"`
}

func NewXamlExportOptions() *XamlExportOptions {
	instance := new(XamlExportOptions)
	return instance
}

func (this *XamlExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *XamlExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *XamlExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *XamlExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *XamlExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *XamlExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *XamlExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *XamlExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *XamlExportOptions) GetFormat() string {
	return this.Format
}

func (this *XamlExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *XamlExportOptions) GetExportHiddenSlides() *bool {
	return this.ExportHiddenSlides
}

func (this *XamlExportOptions) SetExportHiddenSlides(newValue *bool) {
	this.ExportHiddenSlides = newValue
}

func (this *XamlExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDefaultRegularFont, ok := GetMapValue(objMap, "defaultRegularFont"); ok {
		if valDefaultRegularFont != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFont, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	
	if valGradientStyle, ok := GetMapValue(objMap, "gradientStyle"); ok {
		if valGradientStyle != nil {
			var valueForGradientStyle string
			err = json.Unmarshal(*valGradientStyle, &valueForGradientStyle)
			if err != nil {
				var valueForGradientStyleInt int32
				err = json.Unmarshal(*valGradientStyle, &valueForGradientStyleInt)
				if err != nil {
					return err
				}
				this.GradientStyle = string(valueForGradientStyleInt)
			} else {
				this.GradientStyle = valueForGradientStyle
			}
		}
	}
	
	if valFontFallbackRules, ok := GetMapValue(objMap, "fontFallbackRules"); ok {
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
	
	if valFontSubstRules, ok := GetMapValue(objMap, "fontSubstRules"); ok {
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
	
	if valFormat, ok := GetMapValue(objMap, "format"); ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	
	if valExportHiddenSlides, ok := GetMapValue(objMap, "exportHiddenSlides"); ok {
		if valExportHiddenSlides != nil {
			var valueForExportHiddenSlides *bool
			err = json.Unmarshal(*valExportHiddenSlides, &valueForExportHiddenSlides)
			if err != nil {
				return err
			}
			this.ExportHiddenSlides = valueForExportHiddenSlides
		}
	}

	return nil
}
