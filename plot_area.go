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
	GetX() float64
	SetX(newValue float64)

	// the Y location
	GetY() float64
	SetY(newValue float64)

	// Width
	GetWidth() float64
	SetWidth(newValue float64)

	// Height
	GetHeight() float64
	SetHeight(newValue float64)

	// If layout of the plot area is defined manually specifies whether to layout the plot area by its inside (not including axis and axis labels) or outside.
	GetLayoutTargetType() string
	SetLayoutTargetType(newValue string)

	// Get or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)
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

func (this *PlotArea) GetX() float64 {
	return this.X
}

func (this *PlotArea) SetX(newValue float64) {
	this.X = newValue
}
func (this *PlotArea) GetY() float64 {
	return this.Y
}

func (this *PlotArea) SetY(newValue float64) {
	this.Y = newValue
}
func (this *PlotArea) GetWidth() float64 {
	return this.Width
}

func (this *PlotArea) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *PlotArea) GetHeight() float64 {
	return this.Height
}

func (this *PlotArea) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *PlotArea) GetLayoutTargetType() string {
	return this.LayoutTargetType
}

func (this *PlotArea) SetLayoutTargetType(newValue string) {
	this.LayoutTargetType = newValue
}
func (this *PlotArea) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *PlotArea) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *PlotArea) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *PlotArea) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *PlotArea) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *PlotArea) SetLineFormat(newValue ILineFormat) {
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
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}

	return nil
}
