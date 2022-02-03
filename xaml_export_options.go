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

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	getHeight() int32
	setHeight(newValue int32)

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	getWidth() int32
	setWidth(newValue int32)

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

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Width int32 `json:"Width,omitempty"`

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
func (this *XamlExportOptions) getHeight() int32 {
	return this.Height
}

func (this *XamlExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *XamlExportOptions) getWidth() int32 {
	return this.Width
}

func (this *XamlExportOptions) setWidth(newValue int32) {
	this.Width = newValue
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
