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

	// Setting user password to protect the PDF document. 
	getDefaultRegularFont() string
	setDefaultRegularFont(newValue string)

	// Export format.
	getFormat() string
	setFormat(newValue string)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	getShowHiddenSlides() bool
	setShowHiddenSlides(newValue bool)

	// True to convert all metafiles used in a presentation to the PNG images.
	getSaveMetafilesAsPng() bool
	setSaveMetafilesAsPng(newValue bool)

	// True to draw black frame around each slide.
	getDrawSlidesFrame() bool
	setDrawSlidesFrame(newValue bool)
}

type XpsExportOptions struct {

	// Setting user password to protect the PDF document. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`

	// True to convert all metafiles used in a presentation to the PNG images.
	SaveMetafilesAsPng bool `json:"SaveMetafilesAsPng"`

	// True to draw black frame around each slide.
	DrawSlidesFrame bool `json:"DrawSlidesFrame"`
}

func NewXpsExportOptions() *XpsExportOptions {
	instance := new(XpsExportOptions)
	return instance
}

func (this *XpsExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *XpsExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *XpsExportOptions) getFormat() string {
	return this.Format
}

func (this *XpsExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *XpsExportOptions) getShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this *XpsExportOptions) setShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this *XpsExportOptions) getSaveMetafilesAsPng() bool {
	return this.SaveMetafilesAsPng
}

func (this *XpsExportOptions) setSaveMetafilesAsPng(newValue bool) {
	this.SaveMetafilesAsPng = newValue
}
func (this *XpsExportOptions) getDrawSlidesFrame() bool {
	return this.DrawSlidesFrame
}

func (this *XpsExportOptions) setDrawSlidesFrame(newValue bool) {
	this.DrawSlidesFrame = newValue
}

func (this *XpsExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valShowHiddenSlides, ok := objMap["showHiddenSlides"]; ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	if valShowHiddenSlidesCap, ok := objMap["ShowHiddenSlides"]; ok {
		if valShowHiddenSlidesCap != nil {
			var valueForShowHiddenSlides bool
			err = json.Unmarshal(*valShowHiddenSlidesCap, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	
	if valSaveMetafilesAsPng, ok := objMap["saveMetafilesAsPng"]; ok {
		if valSaveMetafilesAsPng != nil {
			var valueForSaveMetafilesAsPng bool
			err = json.Unmarshal(*valSaveMetafilesAsPng, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}
	if valSaveMetafilesAsPngCap, ok := objMap["SaveMetafilesAsPng"]; ok {
		if valSaveMetafilesAsPngCap != nil {
			var valueForSaveMetafilesAsPng bool
			err = json.Unmarshal(*valSaveMetafilesAsPngCap, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}
	
	if valDrawSlidesFrame, ok := objMap["drawSlidesFrame"]; ok {
		if valDrawSlidesFrame != nil {
			var valueForDrawSlidesFrame bool
			err = json.Unmarshal(*valDrawSlidesFrame, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
		}
	}
	if valDrawSlidesFrameCap, ok := objMap["DrawSlidesFrame"]; ok {
		if valDrawSlidesFrameCap != nil {
			var valueForDrawSlidesFrame bool
			err = json.Unmarshal(*valDrawSlidesFrameCap, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
		}
	}

    return nil
}
