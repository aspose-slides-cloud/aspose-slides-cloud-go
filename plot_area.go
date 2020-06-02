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

// Represents the plot area
type IPlotArea interface {

	// the X location
	getX() float64
	setX(newValue float64)

	// the Y location
	getY() float64
	setY(newValue float64)

	// Width
	getWidth() float64
	setWidth(newValue float64)

	// Height
	getHeight() float64
	setHeight(newValue float64)

	// If layout of the plot area is defined manually specifies whether to layout the plot area by its inside (not including axis and axis labels) or outside.
	getLayoutTargetType() string
	setLayoutTargetType(newValue string)

	// Get or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)
}

type PlotArea struct {

	// the X location
	X float64 `json:"X,omitempty"`

	// the Y location
	Y float64 `json:"Y,omitempty"`

	// Width
	Width float64 `json:"Width,omitempty"`

	// Height
	Height float64 `json:"Height,omitempty"`

	// If layout of the plot area is defined manually specifies whether to layout the plot area by its inside (not including axis and axis labels) or outside.
	LayoutTargetType string `json:"LayoutTargetType,omitempty"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func NewPlotArea() *PlotArea {
	instance := new(PlotArea)
	instance.LayoutTargetType = ""
	return instance
}

func (this *PlotArea) getX() float64 {
	return this.X
}

func (this *PlotArea) setX(newValue float64) {
	this.X = newValue
}
func (this *PlotArea) getY() float64 {
	return this.Y
}

func (this *PlotArea) setY(newValue float64) {
	this.Y = newValue
}
func (this *PlotArea) getWidth() float64 {
	return this.Width
}

func (this *PlotArea) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *PlotArea) getHeight() float64 {
	return this.Height
}

func (this *PlotArea) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *PlotArea) getLayoutTargetType() string {
	return this.LayoutTargetType
}

func (this *PlotArea) setLayoutTargetType(newValue string) {
	this.LayoutTargetType = newValue
}
func (this *PlotArea) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *PlotArea) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *PlotArea) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *PlotArea) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *PlotArea) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *PlotArea) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}

func (this *PlotArea) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valX, ok := objMap["x"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	if valXCap, ok := objMap["X"]; ok {
		if valXCap != nil {
			var valueForX float64
			err = json.Unmarshal(*valXCap, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := objMap["y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	if valYCap, ok := objMap["Y"]; ok {
		if valYCap != nil {
			var valueForY float64
			err = json.Unmarshal(*valYCap, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	this.LayoutTargetType = ""
	if valLayoutTargetType, ok := objMap["layoutTargetType"]; ok {
		if valLayoutTargetType != nil {
			var valueForLayoutTargetType string
			err = json.Unmarshal(*valLayoutTargetType, &valueForLayoutTargetType)
			if err != nil {
				var valueForLayoutTargetTypeInt int32
				err = json.Unmarshal(*valLayoutTargetType, &valueForLayoutTargetTypeInt)
				if err != nil {
					return err
				}
				this.LayoutTargetType = string(valueForLayoutTargetTypeInt)
			} else {
				this.LayoutTargetType = valueForLayoutTargetType
			}
		}
	}
	if valLayoutTargetTypeCap, ok := objMap["LayoutTargetType"]; ok {
		if valLayoutTargetTypeCap != nil {
			var valueForLayoutTargetType string
			err = json.Unmarshal(*valLayoutTargetTypeCap, &valueForLayoutTargetType)
			if err != nil {
				var valueForLayoutTargetTypeInt int32
				err = json.Unmarshal(*valLayoutTargetTypeCap, &valueForLayoutTargetTypeInt)
				if err != nil {
					return err
				}
				this.LayoutTargetType = string(valueForLayoutTargetTypeInt)
			} else {
				this.LayoutTargetType = valueForLayoutTargetType
			}
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}

    return nil
}
