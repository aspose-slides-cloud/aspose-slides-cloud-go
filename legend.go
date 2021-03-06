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

// Represents a chart legend
type ILegend interface {

	// position
	getPosition() string
	setPosition(newValue string)

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

	// true if other elements are allowed to overlay the legend
	getOverlay() bool
	setOverlay(newValue bool)

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

type Legend struct {

	// position
	Position string `json:"Position,omitempty"`

	// the X location
	X float64 `json:"X,omitempty"`

	// the Y location
	Y float64 `json:"Y,omitempty"`

	// Width
	Width float64 `json:"Width,omitempty"`

	// Height
	Height float64 `json:"Height,omitempty"`

	// true if other elements are allowed to overlay the legend
	Overlay bool `json:"Overlay"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func NewLegend() *Legend {
	instance := new(Legend)
	instance.Position = ""
	return instance
}

func (this *Legend) getPosition() string {
	return this.Position
}

func (this *Legend) setPosition(newValue string) {
	this.Position = newValue
}
func (this *Legend) getX() float64 {
	return this.X
}

func (this *Legend) setX(newValue float64) {
	this.X = newValue
}
func (this *Legend) getY() float64 {
	return this.Y
}

func (this *Legend) setY(newValue float64) {
	this.Y = newValue
}
func (this *Legend) getWidth() float64 {
	return this.Width
}

func (this *Legend) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *Legend) getHeight() float64 {
	return this.Height
}

func (this *Legend) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *Legend) getOverlay() bool {
	return this.Overlay
}

func (this *Legend) setOverlay(newValue bool) {
	this.Overlay = newValue
}
func (this *Legend) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Legend) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Legend) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Legend) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Legend) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Legend) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}

func (this *Legend) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Position = ""
	if valPosition, ok := objMap["position"]; ok {
		if valPosition != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				var valueForPositionInt int32
				err = json.Unmarshal(*valPosition, &valueForPositionInt)
				if err != nil {
					return err
				}
				this.Position = string(valueForPositionInt)
			} else {
				this.Position = valueForPosition
			}
		}
	}
	if valPositionCap, ok := objMap["Position"]; ok {
		if valPositionCap != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPositionCap, &valueForPosition)
			if err != nil {
				var valueForPositionInt int32
				err = json.Unmarshal(*valPositionCap, &valueForPositionInt)
				if err != nil {
					return err
				}
				this.Position = string(valueForPositionInt)
			} else {
				this.Position = valueForPosition
			}
		}
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
	
	if valOverlay, ok := objMap["overlay"]; ok {
		if valOverlay != nil {
			var valueForOverlay bool
			err = json.Unmarshal(*valOverlay, &valueForOverlay)
			if err != nil {
				return err
			}
			this.Overlay = valueForOverlay
		}
	}
	if valOverlayCap, ok := objMap["Overlay"]; ok {
		if valOverlayCap != nil {
			var valueForOverlay bool
			err = json.Unmarshal(*valOverlayCap, &valueForOverlay)
			if err != nil {
				return err
			}
			this.Overlay = valueForOverlay
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
