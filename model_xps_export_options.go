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

// Provides options that control how a presentation is saved in XPS format.
type IXpsExportOptions interface {

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// True to delete delete all embedded binary objects.
	GetDeleteEmbeddedBinaryObjects() *bool
	SetDeleteEmbeddedBinaryObjects(newValue *bool)

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

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)

	// True to convert all metafiles used in a presentation to the PNG images.
	GetSaveMetafilesAsPng() *bool
	SetSaveMetafilesAsPng(newValue *bool)

	// True to draw black frame around each slide.
	GetDrawSlidesFrame() *bool
	SetDrawSlidesFrame(newValue *bool)
}

type XpsExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// True to delete delete all embedded binary objects.
	DeleteEmbeddedBinaryObjects *bool `json:"DeleteEmbeddedBinaryObjects"`

	// Default regular font for rendering the presentation. 
	GradientStyle string `json:"GradientStyle,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`

	// True to convert all metafiles used in a presentation to the PNG images.
	SaveMetafilesAsPng *bool `json:"SaveMetafilesAsPng"`

	// True to draw black frame around each slide.
	DrawSlidesFrame *bool `json:"DrawSlidesFrame"`
}

func NewXpsExportOptions() *XpsExportOptions {
	instance := new(XpsExportOptions)
	return instance
}

func (this *XpsExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *XpsExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *XpsExportOptions) GetDeleteEmbeddedBinaryObjects() *bool {
	return this.DeleteEmbeddedBinaryObjects
}

func (this *XpsExportOptions) SetDeleteEmbeddedBinaryObjects(newValue *bool) {
	this.DeleteEmbeddedBinaryObjects = newValue
}
func (this *XpsExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *XpsExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *XpsExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *XpsExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *XpsExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *XpsExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *XpsExportOptions) GetFormat() string {
	return this.Format
}

func (this *XpsExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *XpsExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *XpsExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}
func (this *XpsExportOptions) GetSaveMetafilesAsPng() *bool {
	return this.SaveMetafilesAsPng
}

func (this *XpsExportOptions) SetSaveMetafilesAsPng(newValue *bool) {
	this.SaveMetafilesAsPng = newValue
}
func (this *XpsExportOptions) GetDrawSlidesFrame() *bool {
	return this.DrawSlidesFrame
}

func (this *XpsExportOptions) SetDrawSlidesFrame(newValue *bool) {
	this.DrawSlidesFrame = newValue
}

func (this *XpsExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valDeleteEmbeddedBinaryObjects, ok := GetMapValue(objMap, "deleteEmbeddedBinaryObjects"); ok {
		if valDeleteEmbeddedBinaryObjects != nil {
			var valueForDeleteEmbeddedBinaryObjects *bool
			err = json.Unmarshal(*valDeleteEmbeddedBinaryObjects, &valueForDeleteEmbeddedBinaryObjects)
			if err != nil {
				return err
			}
			this.DeleteEmbeddedBinaryObjects = valueForDeleteEmbeddedBinaryObjects
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
	
	if valShowHiddenSlides, ok := GetMapValue(objMap, "showHiddenSlides"); ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides *bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	
	if valSaveMetafilesAsPng, ok := GetMapValue(objMap, "saveMetafilesAsPng"); ok {
		if valSaveMetafilesAsPng != nil {
			var valueForSaveMetafilesAsPng *bool
			err = json.Unmarshal(*valSaveMetafilesAsPng, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}
	
	if valDrawSlidesFrame, ok := GetMapValue(objMap, "drawSlidesFrame"); ok {
		if valDrawSlidesFrame != nil {
			var valueForDrawSlidesFrame *bool
			err = json.Unmarshal(*valDrawSlidesFrame, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
		}
	}

	return nil
}
