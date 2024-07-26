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

// Provides options that control how a presentation is saved in Html format.
type IHtmlExportOptions interface {

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

	// Get or sets flag for save presentation as zip file
	GetSaveAsZip() *bool
	SetSaveAsZip(newValue *bool)

	// Get or set name of subdirectory in zip-file for store external files
	GetSubDirectoryName() string
	SetSubDirectoryName(newValue string)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)

	// True to make layout responsive by excluding width and height attributes from svg container.
	GetSvgResponsiveLayout() *bool
	SetSvgResponsiveLayout(newValue *bool)

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	GetJpegQuality() int32
	SetJpegQuality(newValue int32)

	// Represents the pictures compression level
	GetPicturesCompression() string
	SetPicturesCompression(newValue string)

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	GetDeletePicturesCroppedAreas() *bool
	SetDeletePicturesCroppedAreas(newValue *bool)

	// Slides layouting options
	GetSlidesLayoutOptions() ISlidesLayoutOptions
	SetSlidesLayoutOptions(newValue ISlidesLayoutOptions)
}

type HtmlExportOptions struct {

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

	// Get or sets flag for save presentation as zip file
	SaveAsZip *bool `json:"SaveAsZip"`

	// Get or set name of subdirectory in zip-file for store external files
	SubDirectoryName string `json:"SubDirectoryName,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`

	// True to make layout responsive by excluding width and height attributes from svg container.
	SvgResponsiveLayout *bool `json:"SvgResponsiveLayout"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Represents the pictures compression level
	PicturesCompression string `json:"PicturesCompression,omitempty"`

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	DeletePicturesCroppedAreas *bool `json:"DeletePicturesCroppedAreas"`

	// Slides layouting options
	SlidesLayoutOptions ISlidesLayoutOptions `json:"SlidesLayoutOptions,omitempty"`
}

func NewHtmlExportOptions() *HtmlExportOptions {
	instance := new(HtmlExportOptions)
	return instance
}

func (this *HtmlExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *HtmlExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *HtmlExportOptions) GetGradientStyle() string {
	return this.GradientStyle
}

func (this *HtmlExportOptions) SetGradientStyle(newValue string) {
	this.GradientStyle = newValue
}
func (this *HtmlExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *HtmlExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *HtmlExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *HtmlExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *HtmlExportOptions) GetFormat() string {
	return this.Format
}

func (this *HtmlExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *HtmlExportOptions) GetSaveAsZip() *bool {
	return this.SaveAsZip
}

func (this *HtmlExportOptions) SetSaveAsZip(newValue *bool) {
	this.SaveAsZip = newValue
}
func (this *HtmlExportOptions) GetSubDirectoryName() string {
	return this.SubDirectoryName
}

func (this *HtmlExportOptions) SetSubDirectoryName(newValue string) {
	this.SubDirectoryName = newValue
}
func (this *HtmlExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *HtmlExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}
func (this *HtmlExportOptions) GetSvgResponsiveLayout() *bool {
	return this.SvgResponsiveLayout
}

func (this *HtmlExportOptions) SetSvgResponsiveLayout(newValue *bool) {
	this.SvgResponsiveLayout = newValue
}
func (this *HtmlExportOptions) GetJpegQuality() int32 {
	return this.JpegQuality
}

func (this *HtmlExportOptions) SetJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this *HtmlExportOptions) GetPicturesCompression() string {
	return this.PicturesCompression
}

func (this *HtmlExportOptions) SetPicturesCompression(newValue string) {
	this.PicturesCompression = newValue
}
func (this *HtmlExportOptions) GetDeletePicturesCroppedAreas() *bool {
	return this.DeletePicturesCroppedAreas
}

func (this *HtmlExportOptions) SetDeletePicturesCroppedAreas(newValue *bool) {
	this.DeletePicturesCroppedAreas = newValue
}
func (this *HtmlExportOptions) GetSlidesLayoutOptions() ISlidesLayoutOptions {
	return this.SlidesLayoutOptions
}

func (this *HtmlExportOptions) SetSlidesLayoutOptions(newValue ISlidesLayoutOptions) {
	this.SlidesLayoutOptions = newValue
}

func (this *HtmlExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valSaveAsZip, ok := GetMapValue(objMap, "saveAsZip"); ok {
		if valSaveAsZip != nil {
			var valueForSaveAsZip *bool
			err = json.Unmarshal(*valSaveAsZip, &valueForSaveAsZip)
			if err != nil {
				return err
			}
			this.SaveAsZip = valueForSaveAsZip
		}
	}
	
	if valSubDirectoryName, ok := GetMapValue(objMap, "subDirectoryName"); ok {
		if valSubDirectoryName != nil {
			var valueForSubDirectoryName string
			err = json.Unmarshal(*valSubDirectoryName, &valueForSubDirectoryName)
			if err != nil {
				return err
			}
			this.SubDirectoryName = valueForSubDirectoryName
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
	
	if valSvgResponsiveLayout, ok := GetMapValue(objMap, "svgResponsiveLayout"); ok {
		if valSvgResponsiveLayout != nil {
			var valueForSvgResponsiveLayout *bool
			err = json.Unmarshal(*valSvgResponsiveLayout, &valueForSvgResponsiveLayout)
			if err != nil {
				return err
			}
			this.SvgResponsiveLayout = valueForSvgResponsiveLayout
		}
	}
	
	if valJpegQuality, ok := GetMapValue(objMap, "jpegQuality"); ok {
		if valJpegQuality != nil {
			var valueForJpegQuality int32
			err = json.Unmarshal(*valJpegQuality, &valueForJpegQuality)
			if err != nil {
				return err
			}
			this.JpegQuality = valueForJpegQuality
		}
	}
	
	if valPicturesCompression, ok := GetMapValue(objMap, "picturesCompression"); ok {
		if valPicturesCompression != nil {
			var valueForPicturesCompression string
			err = json.Unmarshal(*valPicturesCompression, &valueForPicturesCompression)
			if err != nil {
				var valueForPicturesCompressionInt int32
				err = json.Unmarshal(*valPicturesCompression, &valueForPicturesCompressionInt)
				if err != nil {
					return err
				}
				this.PicturesCompression = string(valueForPicturesCompressionInt)
			} else {
				this.PicturesCompression = valueForPicturesCompression
			}
		}
	}
	
	if valDeletePicturesCroppedAreas, ok := GetMapValue(objMap, "deletePicturesCroppedAreas"); ok {
		if valDeletePicturesCroppedAreas != nil {
			var valueForDeletePicturesCroppedAreas *bool
			err = json.Unmarshal(*valDeletePicturesCroppedAreas, &valueForDeletePicturesCroppedAreas)
			if err != nil {
				return err
			}
			this.DeletePicturesCroppedAreas = valueForDeletePicturesCroppedAreas
		}
	}
	
	if valSlidesLayoutOptions, ok := GetMapValue(objMap, "slidesLayoutOptions"); ok {
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

	return nil
}
