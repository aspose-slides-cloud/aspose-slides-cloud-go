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
	Position string `json:"Position"`

	// the X location
	X float64 `json:"X"`

	// the Y location
	Y float64 `json:"Y"`

	// Width
	Width float64 `json:"Width"`

	// Height
	Height float64 `json:"Height"`

	// true if other elements are allowed to overlay the legend
	Overlay bool `json:"Overlay"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func (this Legend) getPosition() string {
	return this.Position
}

func (this Legend) setPosition(newValue string) {
	this.Position = newValue
}
func (this Legend) getX() float64 {
	return this.X
}

func (this Legend) setX(newValue float64) {
	this.X = newValue
}
func (this Legend) getY() float64 {
	return this.Y
}

func (this Legend) setY(newValue float64) {
	this.Y = newValue
}
func (this Legend) getWidth() float64 {
	return this.Width
}

func (this Legend) setWidth(newValue float64) {
	this.Width = newValue
}
func (this Legend) getHeight() float64 {
	return this.Height
}

func (this Legend) setHeight(newValue float64) {
	this.Height = newValue
}
func (this Legend) getOverlay() bool {
	return this.Overlay
}

func (this Legend) setOverlay(newValue bool) {
	this.Overlay = newValue
}
func (this Legend) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this Legend) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this Legend) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this Legend) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this Legend) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this Legend) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}

func (this *Legend) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valPosition, ok := objMap["Position"]; ok {
		if valPosition != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}

	if valX, ok := objMap["X"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}

	if valY, ok := objMap["Y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}

	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

	if valHeight, ok := objMap["Height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	if valOverlay, ok := objMap["Overlay"]; ok {
		if valOverlay != nil {
			var valueForOverlay bool
			err = json.Unmarshal(*valOverlay, &valueForOverlay)
			if err != nil {
				return err
			}
			this.Overlay = valueForOverlay
		}
	}

	if valFillFormat, ok := objMap["FillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}

	if valEffectFormat, ok := objMap["EffectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}

	if valLineFormat, ok := objMap["LineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}

    return nil
}
