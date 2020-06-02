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

// Gradient stop.
type IGradientFillStop interface {

	// Color.
	getColor() string
	setColor(newValue string)

	// Stop position (0..1).
	getPosition() float64
	setPosition(newValue float64)
}

type GradientFillStop struct {

	// Color.
	Color string `json:"Color,omitempty"`

	// Stop position (0..1).
	Position float64 `json:"Position"`
}

func NewGradientFillStop() *GradientFillStop {
	instance := new(GradientFillStop)
	return instance
}

func (this *GradientFillStop) getColor() string {
	return this.Color
}

func (this *GradientFillStop) setColor(newValue string) {
	this.Color = newValue
}
func (this *GradientFillStop) getPosition() float64 {
	return this.Position
}

func (this *GradientFillStop) setPosition(newValue float64) {
	this.Position = newValue
}

func (this *GradientFillStop) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valColor, ok := objMap["color"]; ok {
		if valColor != nil {
			var valueForColor string
			err = json.Unmarshal(*valColor, &valueForColor)
			if err != nil {
				return err
			}
			this.Color = valueForColor
		}
	}
	if valColorCap, ok := objMap["Color"]; ok {
		if valColorCap != nil {
			var valueForColor string
			err = json.Unmarshal(*valColorCap, &valueForColor)
			if err != nil {
				return err
			}
			this.Color = valueForColor
		}
	}
	
	if valPosition, ok := objMap["position"]; ok {
		if valPosition != nil {
			var valueForPosition float64
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}
	if valPositionCap, ok := objMap["Position"]; ok {
		if valPositionCap != nil {
			var valueForPosition float64
			err = json.Unmarshal(*valPositionCap, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}

    return nil
}
