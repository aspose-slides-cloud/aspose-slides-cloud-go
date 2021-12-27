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

// Provides options that control how a presentation is saved in PPTX format.
type IPptxExportOptions interface {

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

	// The conformance class to which the PresentationML document conforms. Read/write Conformance.
	getConformance() string
	setConformance(newValue string)
}

type PptxExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Width int32 `json:"Width,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// The conformance class to which the PresentationML document conforms. Read/write Conformance.
	Conformance string `json:"Conformance,omitempty"`
}

func NewPptxExportOptions() *PptxExportOptions {
	instance := new(PptxExportOptions)
	instance.Conformance = ""
	return instance
}

func (this *PptxExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *PptxExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *PptxExportOptions) getHeight() int32 {
	return this.Height
}

func (this *PptxExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *PptxExportOptions) getWidth() int32 {
	return this.Width
}

func (this *PptxExportOptions) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *PptxExportOptions) getFormat() string {
	return this.Format
}

func (this *PptxExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *PptxExportOptions) getConformance() string {
	return this.Conformance
}

func (this *PptxExportOptions) setConformance(newValue string) {
	this.Conformance = newValue
}

func (this *PptxExportOptions) UnmarshalJSON(b []byte) error {
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
	this.Conformance = ""
	if valConformance, ok := objMap["conformance"]; ok {
		if valConformance != nil {
			var valueForConformance string
			err = json.Unmarshal(*valConformance, &valueForConformance)
			if err != nil {
				var valueForConformanceInt int32
				err = json.Unmarshal(*valConformance, &valueForConformanceInt)
				if err != nil {
					return err
				}
				this.Conformance = string(valueForConformanceInt)
			} else {
				this.Conformance = valueForConformance
			}
		}
	}
	if valConformanceCap, ok := objMap["Conformance"]; ok {
		if valConformanceCap != nil {
			var valueForConformance string
			err = json.Unmarshal(*valConformanceCap, &valueForConformance)
			if err != nil {
				var valueForConformanceInt int32
				err = json.Unmarshal(*valConformanceCap, &valueForConformanceInt)
				if err != nil {
					return err
				}
				this.Conformance = string(valueForConformanceInt)
			} else {
				this.Conformance = valueForConformance
			}
		}
	}

	return nil
}
