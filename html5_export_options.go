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
type IHtml5ExportOptions interface {

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

	// Gets or sets transitions animation option.
	getAnimateTransitions() bool
	setAnimateTransitions(newValue bool)

	// Gets or sets shapes animation option.
	getAnimateShapes() bool
	setAnimateShapes(newValue bool)
}

type Html5ExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Width int32 `json:"Width,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Gets or sets transitions animation option.
	AnimateTransitions bool `json:"AnimateTransitions"`

	// Gets or sets shapes animation option.
	AnimateShapes bool `json:"AnimateShapes"`
}

func NewHtml5ExportOptions() *Html5ExportOptions {
	instance := new(Html5ExportOptions)
	return instance
}

func (this *Html5ExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *Html5ExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *Html5ExportOptions) getHeight() int32 {
	return this.Height
}

func (this *Html5ExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *Html5ExportOptions) getWidth() int32 {
	return this.Width
}

func (this *Html5ExportOptions) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *Html5ExportOptions) getFormat() string {
	return this.Format
}

func (this *Html5ExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *Html5ExportOptions) getAnimateTransitions() bool {
	return this.AnimateTransitions
}

func (this *Html5ExportOptions) setAnimateTransitions(newValue bool) {
	this.AnimateTransitions = newValue
}
func (this *Html5ExportOptions) getAnimateShapes() bool {
	return this.AnimateShapes
}

func (this *Html5ExportOptions) setAnimateShapes(newValue bool) {
	this.AnimateShapes = newValue
}

func (this *Html5ExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valAnimateTransitions, ok := objMap["animateTransitions"]; ok {
		if valAnimateTransitions != nil {
			var valueForAnimateTransitions bool
			err = json.Unmarshal(*valAnimateTransitions, &valueForAnimateTransitions)
			if err != nil {
				return err
			}
			this.AnimateTransitions = valueForAnimateTransitions
		}
	}
	if valAnimateTransitionsCap, ok := objMap["AnimateTransitions"]; ok {
		if valAnimateTransitionsCap != nil {
			var valueForAnimateTransitions bool
			err = json.Unmarshal(*valAnimateTransitionsCap, &valueForAnimateTransitions)
			if err != nil {
				return err
			}
			this.AnimateTransitions = valueForAnimateTransitions
		}
	}
	
	if valAnimateShapes, ok := objMap["animateShapes"]; ok {
		if valAnimateShapes != nil {
			var valueForAnimateShapes bool
			err = json.Unmarshal(*valAnimateShapes, &valueForAnimateShapes)
			if err != nil {
				return err
			}
			this.AnimateShapes = valueForAnimateShapes
		}
	}
	if valAnimateShapesCap, ok := objMap["AnimateShapes"]; ok {
		if valAnimateShapesCap != nil {
			var valueForAnimateShapes bool
			err = json.Unmarshal(*valAnimateShapesCap, &valueForAnimateShapes)
			if err != nil {
				return err
			}
			this.AnimateShapes = valueForAnimateShapes
		}
	}

	return nil
}
