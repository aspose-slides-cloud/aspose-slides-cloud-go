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

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	getHeight() int32
	setHeight(newValue int32)

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	getWidth() int32
	setWidth(newValue int32)

	// Export format.
	getFormat() string
	setFormat(newValue string)

	// Transition duration.
	getTransitionDuration() int32
	setTransitionDuration(newValue int32)

	// Video resolution type
	getVideoResolutionType() string
	setVideoResolutionType(newValue string)
}

type VideoExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Width int32 `json:"Width,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Transition duration.
	TransitionDuration int32 `json:"TransitionDuration,omitempty"`

	// Video resolution type
	VideoResolutionType string `json:"VideoResolutionType,omitempty"`
}

func NewVideoExportOptions() *VideoExportOptions {
	instance := new(VideoExportOptions)
	instance.VideoResolutionType = ""
	return instance
}

func (this *VideoExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *VideoExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *VideoExportOptions) getHeight() int32 {
	return this.Height
}

func (this *VideoExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *VideoExportOptions) getWidth() int32 {
	return this.Width
}

func (this *VideoExportOptions) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *VideoExportOptions) getFormat() string {
	return this.Format
}

func (this *VideoExportOptions) setFormat(newValue string) {
	this.Format = newValue
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
