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

// Provides options that control how a presentation is saved in an video format.
type IVideoExportOptions interface {

	// Default regular font for rendering the presentation. 
	getDefaultRegularFont() string
	setDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	getFontFallbackRules() []IFontFallbackRule
	setFontFallbackRules(newValue []IFontFallbackRule)

	// Export format.
	getFormat() string
	setFormat(newValue string)

	// Slides transition duration.
	getSlidesTransitionDuration() int32
	setSlidesTransitionDuration(newValue int32)

	// Video transition type
	getTransitionType() string
	setTransitionType(newValue string)

	// Duration of transition defined in TransitionType property.
	getTransitionDuration() int32
	setTransitionDuration(newValue int32)

	// Video resolution type
	getVideoResolutionType() string
	setVideoResolutionType(newValue string)
}

type VideoExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Slides transition duration.
	SlidesTransitionDuration int32 `json:"SlidesTransitionDuration,omitempty"`

	// Video transition type
	TransitionType string `json:"TransitionType,omitempty"`

	// Duration of transition defined in TransitionType property.
	TransitionDuration int32 `json:"TransitionDuration,omitempty"`

	// Video resolution type
	VideoResolutionType string `json:"VideoResolutionType,omitempty"`
}

func NewVideoExportOptions() *VideoExportOptions {
	instance := new(VideoExportOptions)
	instance.TransitionType = ""
	instance.VideoResolutionType = ""
	return instance
}

func (this *VideoExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *VideoExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *VideoExportOptions) getFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *VideoExportOptions) setFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *VideoExportOptions) getFormat() string {
	return this.Format
}

func (this *VideoExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *VideoExportOptions) getSlidesTransitionDuration() int32 {
	return this.SlidesTransitionDuration
}

func (this *VideoExportOptions) setSlidesTransitionDuration(newValue int32) {
	this.SlidesTransitionDuration = newValue
}
func (this *VideoExportOptions) getTransitionType() string {
	return this.TransitionType
}

func (this *VideoExportOptions) setTransitionType(newValue string) {
	this.TransitionType = newValue
}
func (this *VideoExportOptions) getTransitionDuration() int32 {
	return this.TransitionDuration
}

func (this *VideoExportOptions) setTransitionDuration(newValue int32) {
	this.TransitionDuration = newValue
}
func (this *VideoExportOptions) getVideoResolutionType() string {
	return this.VideoResolutionType
}

func (this *VideoExportOptions) setVideoResolutionType(newValue string) {
	this.VideoResolutionType = newValue
}

func (this *VideoExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valSlidesTransitionDuration, ok := objMap["slidesTransitionDuration"]; ok {
		if valSlidesTransitionDuration != nil {
			var valueForSlidesTransitionDuration int32
			err = json.Unmarshal(*valSlidesTransitionDuration, &valueForSlidesTransitionDuration)
			if err != nil {
				return err
			}
			this.SlidesTransitionDuration = valueForSlidesTransitionDuration
		}
	}
	if valSlidesTransitionDurationCap, ok := objMap["SlidesTransitionDuration"]; ok {
		if valSlidesTransitionDurationCap != nil {
			var valueForSlidesTransitionDuration int32
			err = json.Unmarshal(*valSlidesTransitionDurationCap, &valueForSlidesTransitionDuration)
			if err != nil {
				return err
			}
			this.SlidesTransitionDuration = valueForSlidesTransitionDuration
		}
	}
	this.TransitionType = ""
	if valTransitionType, ok := objMap["transitionType"]; ok {
		if valTransitionType != nil {
			var valueForTransitionType string
			err = json.Unmarshal(*valTransitionType, &valueForTransitionType)
			if err != nil {
				var valueForTransitionTypeInt int32
				err = json.Unmarshal(*valTransitionType, &valueForTransitionTypeInt)
				if err != nil {
					return err
				}
				this.TransitionType = string(valueForTransitionTypeInt)
			} else {
				this.TransitionType = valueForTransitionType
			}
		}
	}
	if valTransitionTypeCap, ok := objMap["TransitionType"]; ok {
		if valTransitionTypeCap != nil {
			var valueForTransitionType string
			err = json.Unmarshal(*valTransitionTypeCap, &valueForTransitionType)
			if err != nil {
				var valueForTransitionTypeInt int32
				err = json.Unmarshal(*valTransitionTypeCap, &valueForTransitionTypeInt)
				if err != nil {
					return err
				}
				this.TransitionType = string(valueForTransitionTypeInt)
			} else {
				this.TransitionType = valueForTransitionType
			}
		}
	}
	
	if valTransitionDuration, ok := objMap["transitionDuration"]; ok {
		if valTransitionDuration != nil {
			var valueForTransitionDuration int32
			err = json.Unmarshal(*valTransitionDuration, &valueForTransitionDuration)
			if err != nil {
				return err
			}
			this.TransitionDuration = valueForTransitionDuration
		}
	}
	if valTransitionDurationCap, ok := objMap["TransitionDuration"]; ok {
		if valTransitionDurationCap != nil {
			var valueForTransitionDuration int32
			err = json.Unmarshal(*valTransitionDurationCap, &valueForTransitionDuration)
			if err != nil {
				return err
			}
			this.TransitionDuration = valueForTransitionDuration
		}
	}
	this.VideoResolutionType = ""
	if valVideoResolutionType, ok := objMap["videoResolutionType"]; ok {
		if valVideoResolutionType != nil {
			var valueForVideoResolutionType string
			err = json.Unmarshal(*valVideoResolutionType, &valueForVideoResolutionType)
			if err != nil {
				var valueForVideoResolutionTypeInt int32
				err = json.Unmarshal(*valVideoResolutionType, &valueForVideoResolutionTypeInt)
				if err != nil {
					return err
				}
				this.VideoResolutionType = string(valueForVideoResolutionTypeInt)
			} else {
				this.VideoResolutionType = valueForVideoResolutionType
			}
		}
	}
	if valVideoResolutionTypeCap, ok := objMap["VideoResolutionType"]; ok {
		if valVideoResolutionTypeCap != nil {
			var valueForVideoResolutionType string
			err = json.Unmarshal(*valVideoResolutionTypeCap, &valueForVideoResolutionType)
			if err != nil {
				var valueForVideoResolutionTypeInt int32
				err = json.Unmarshal(*valVideoResolutionTypeCap, &valueForVideoResolutionTypeInt)
				if err != nil {
					return err
				}
				this.VideoResolutionType = string(valueForVideoResolutionTypeInt)
			} else {
				this.VideoResolutionType = valueForVideoResolutionType
			}
		}
	}

	return nil
}
