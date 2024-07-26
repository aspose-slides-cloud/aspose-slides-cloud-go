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
	GetPosition() string
	SetPosition(newValue string)

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

	// true if other elements are allowed to overlay the legend
	GetOverlay() *bool
	SetOverlay(newValue *bool)

	// Get or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Get or sets value determines the visibility of legend
	GetHasLegend() *bool
	SetHasLegend(newValue *bool)
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
	Overlay *bool `json:"Overlay"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Get or sets value determines the visibility of legend
	HasLegend *bool `json:"HasLegend"`
}

func NewLegend() *Legend {
	instance := new(Legend)
	return instance
}

func (this *Legend) GetPosition() string {
	return this.Position
}

func (this *Legend) SetPosition(newValue string) {
	this.Position = newValue
}
func (this *Legend) GetX() float64 {
	return this.X
}

func (this *Legend) SetX(newValue float64) {
	this.X = newValue
}
func (this *Legend) GetY() float64 {
	return this.Y
}

func (this *Legend) SetY(newValue float64) {
	this.Y = newValue
}
func (this *Legend) GetWidth() float64 {
	return this.Width
}

func (this *Legend) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *Legend) GetHeight() float64 {
	return this.Height
}

func (this *Legend) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *Legend) GetOverlay() *bool {
	return this.Overlay
}

func (this *Legend) SetOverlay(newValue *bool) {
	this.Overlay = newValue
}
func (this *Legend) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Legend) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Legend) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Legend) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Legend) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Legend) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Legend) GetHasLegend() *bool {
	return this.HasLegend
}

func (this *Legend) SetHasLegend(newValue *bool) {
	this.HasLegend = newValue
}

func (this *Legend) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPosition, ok := GetMapValue(objMap, "position"); ok {
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
	
	if valX, ok := GetMapValue(objMap, "x"); ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := GetMapValue(objMap, "y"); ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valOverlay, ok := GetMapValue(objMap, "overlay"); ok {
		if valOverlay != nil {
			var valueForOverlay *bool
			err = json.Unmarshal(*valOverlay, &valueForOverlay)
			if err != nil {
				return err
			}
			this.Overlay = valueForOverlay
		}
	}
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
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
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
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
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
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
	
	if valHasLegend, ok := GetMapValue(objMap, "hasLegend"); ok {
		if valHasLegend != nil {
			var valueForHasLegend *bool
			err = json.Unmarshal(*valHasLegend, &valueForHasLegend)
			if err != nil {
				return err
			}
			this.HasLegend = valueForHasLegend
		}
	}

	return nil
}
