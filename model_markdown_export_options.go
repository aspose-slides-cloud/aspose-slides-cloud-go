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

// Provides options that control how a presentation is saved in Html5 format.
type IMarkdownExportOptions interface {

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

	// Specifies markdown specification to convert presentation. Default is TextOnly.
	GetExportType() string
	SetExportType(newValue string)

	// Specifies markdown specification to convert presentation. Default is MultiMarkdown.
	GetFlavor() string
	SetFlavor(newValue string)

	// Specifies whether the generated document should have new lines of \\\\r(Macintosh), \\\\n(Unix) or \\\\r\\\\n(Windows). Default is Unix.
	GetNewLineType() string
	SetNewLineType(newValue string)

	// Specifies folder name to save images. Default is Images. 
	GetImagesSaveFolderName() string
	SetImagesSaveFolderName(newValue string)

	// Specifies whether the generated document should include slide number. Default is false. 
	GetShowSlideNumber() *bool
	SetShowSlideNumber(newValue *bool)

	// Specifies whether the generated document should include comments. Default is false. 
	GetShowComments() *bool
	SetShowComments(newValue *bool)

	// Specifies whether the generated document should include hidden slides. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)
}

type MarkdownExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Specifies markdown specification to convert presentation. Default is TextOnly.
	ExportType string `json:"ExportType,omitempty"`

	// Specifies markdown specification to convert presentation. Default is MultiMarkdown.
	Flavor string `json:"Flavor,omitempty"`

	// Specifies whether the generated document should have new lines of \\\\r(Macintosh), \\\\n(Unix) or \\\\r\\\\n(Windows). Default is Unix.
	NewLineType string `json:"NewLineType,omitempty"`

	// Specifies folder name to save images. Default is Images. 
	ImagesSaveFolderName string `json:"ImagesSaveFolderName,omitempty"`

	// Specifies whether the generated document should include slide number. Default is false. 
	ShowSlideNumber *bool `json:"ShowSlideNumber"`

	// Specifies whether the generated document should include comments. Default is false. 
	ShowComments *bool `json:"ShowComments"`

	// Specifies whether the generated document should include hidden slides. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`
}

func NewMarkdownExportOptions() *MarkdownExportOptions {
	instance := new(MarkdownExportOptions)
	return instance
}

func (this *MarkdownExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *MarkdownExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *MarkdownExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *MarkdownExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *MarkdownExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *MarkdownExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *MarkdownExportOptions) GetFormat() string {
	return this.Format
}

func (this *MarkdownExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *MarkdownExportOptions) GetExportType() string {
	return this.ExportType
}

func (this *MarkdownExportOptions) SetExportType(newValue string) {
	this.ExportType = newValue
}
func (this *MarkdownExportOptions) GetFlavor() string {
	return this.Flavor
}

func (this *MarkdownExportOptions) SetFlavor(newValue string) {
	this.Flavor = newValue
}
func (this *MarkdownExportOptions) GetNewLineType() string {
	return this.NewLineType
}

func (this *MarkdownExportOptions) SetNewLineType(newValue string) {
	this.NewLineType = newValue
}
func (this *MarkdownExportOptions) GetImagesSaveFolderName() string {
	return this.ImagesSaveFolderName
}

func (this *MarkdownExportOptions) SetImagesSaveFolderName(newValue string) {
	this.ImagesSaveFolderName = newValue
}
func (this *MarkdownExportOptions) GetShowSlideNumber() *bool {
	return this.ShowSlideNumber
}

func (this *MarkdownExportOptions) SetShowSlideNumber(newValue *bool) {
	this.ShowSlideNumber = newValue
}
func (this *MarkdownExportOptions) GetShowComments() *bool {
	return this.ShowComments
}

func (this *MarkdownExportOptions) SetShowComments(newValue *bool) {
	this.ShowComments = newValue
}
func (this *MarkdownExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *MarkdownExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}

func (this *MarkdownExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valExportType, ok := objMap["exportType"]; ok {
		if valExportType != nil {
			var valueForExportType string
			err = json.Unmarshal(*valExportType, &valueForExportType)
			if err != nil {
				var valueForExportTypeInt int32
				err = json.Unmarshal(*valExportType, &valueForExportTypeInt)
				if err != nil {
					return err
				}
				this.ExportType = string(valueForExportTypeInt)
			} else {
				this.ExportType = valueForExportType
			}
		}
	}
	if valExportTypeCap, ok := objMap["ExportType"]; ok {
		if valExportTypeCap != nil {
			var valueForExportType string
			err = json.Unmarshal(*valExportTypeCap, &valueForExportType)
			if err != nil {
				var valueForExportTypeInt int32
				err = json.Unmarshal(*valExportTypeCap, &valueForExportTypeInt)
				if err != nil {
					return err
				}
				this.ExportType = string(valueForExportTypeInt)
			} else {
				this.ExportType = valueForExportType
			}
		}
	}
	
	if valFlavor, ok := objMap["flavor"]; ok {
		if valFlavor != nil {
			var valueForFlavor string
			err = json.Unmarshal(*valFlavor, &valueForFlavor)
			if err != nil {
				var valueForFlavorInt int32
				err = json.Unmarshal(*valFlavor, &valueForFlavorInt)
				if err != nil {
					return err
				}
				this.Flavor = string(valueForFlavorInt)
			} else {
				this.Flavor = valueForFlavor
			}
		}
	}
	if valFlavorCap, ok := objMap["Flavor"]; ok {
		if valFlavorCap != nil {
			var valueForFlavor string
			err = json.Unmarshal(*valFlavorCap, &valueForFlavor)
			if err != nil {
				var valueForFlavorInt int32
				err = json.Unmarshal(*valFlavorCap, &valueForFlavorInt)
				if err != nil {
					return err
				}
				this.Flavor = string(valueForFlavorInt)
			} else {
				this.Flavor = valueForFlavor
			}
		}
	}
	
	if valNewLineType, ok := objMap["newLineType"]; ok {
		if valNewLineType != nil {
			var valueForNewLineType string
			err = json.Unmarshal(*valNewLineType, &valueForNewLineType)
			if err != nil {
				var valueForNewLineTypeInt int32
				err = json.Unmarshal(*valNewLineType, &valueForNewLineTypeInt)
				if err != nil {
					return err
				}
				this.NewLineType = string(valueForNewLineTypeInt)
			} else {
				this.NewLineType = valueForNewLineType
			}
		}
	}
	if valNewLineTypeCap, ok := objMap["NewLineType"]; ok {
		if valNewLineTypeCap != nil {
			var valueForNewLineType string
			err = json.Unmarshal(*valNewLineTypeCap, &valueForNewLineType)
			if err != nil {
				var valueForNewLineTypeInt int32
				err = json.Unmarshal(*valNewLineTypeCap, &valueForNewLineTypeInt)
				if err != nil {
					return err
				}
				this.NewLineType = string(valueForNewLineTypeInt)
			} else {
				this.NewLineType = valueForNewLineType
			}
		}
	}
	
	if valImagesSaveFolderName, ok := objMap["imagesSaveFolderName"]; ok {
		if valImagesSaveFolderName != nil {
			var valueForImagesSaveFolderName string
			err = json.Unmarshal(*valImagesSaveFolderName, &valueForImagesSaveFolderName)
			if err != nil {
				return err
			}
			this.ImagesSaveFolderName = valueForImagesSaveFolderName
		}
	}
	if valImagesSaveFolderNameCap, ok := objMap["ImagesSaveFolderName"]; ok {
		if valImagesSaveFolderNameCap != nil {
			var valueForImagesSaveFolderName string
			err = json.Unmarshal(*valImagesSaveFolderNameCap, &valueForImagesSaveFolderName)
			if err != nil {
				return err
			}
			this.ImagesSaveFolderName = valueForImagesSaveFolderName
		}
	}
	
	if valShowSlideNumber, ok := objMap["showSlideNumber"]; ok {
		if valShowSlideNumber != nil {
			var valueForShowSlideNumber *bool
			err = json.Unmarshal(*valShowSlideNumber, &valueForShowSlideNumber)
			if err != nil {
				return err
			}
			this.ShowSlideNumber = valueForShowSlideNumber
		}
	}
	if valShowSlideNumberCap, ok := objMap["ShowSlideNumber"]; ok {
		if valShowSlideNumberCap != nil {
			var valueForShowSlideNumber *bool
			err = json.Unmarshal(*valShowSlideNumberCap, &valueForShowSlideNumber)
			if err != nil {
				return err
			}
			this.ShowSlideNumber = valueForShowSlideNumber
		}
	}
	
	if valShowComments, ok := objMap["showComments"]; ok {
		if valShowComments != nil {
			var valueForShowComments *bool
			err = json.Unmarshal(*valShowComments, &valueForShowComments)
			if err != nil {
				return err
			}
			this.ShowComments = valueForShowComments
		}
	}
	if valShowCommentsCap, ok := objMap["ShowComments"]; ok {
		if valShowCommentsCap != nil {
			var valueForShowComments *bool
			err = json.Unmarshal(*valShowCommentsCap, &valueForShowComments)
			if err != nil {
				return err
			}
			this.ShowComments = valueForShowComments
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

	return nil
}
