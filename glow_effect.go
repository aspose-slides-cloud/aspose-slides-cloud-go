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

// Represents glow effect 
type IGlowEffect interface {

	// radius
	getRadius() float64
	setRadius(newValue float64)

	// color
	getColor() string
	setColor(newValue string)
}

type GlowEffect struct {

	// radius
	Radius float64 `json:"Radius"`

	// color
	Color string `json:"Color,omitempty"`
}

func (this GlowEffect) getRadius() float64 {
	return this.Radius
}

func (this GlowEffect) setRadius(newValue float64) {
	this.Radius = newValue
}
func (this GlowEffect) getColor() string {
	return this.Color
}

func (this GlowEffect) setColor(newValue string) {
	this.Color = newValue
}

func (this *GlowEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valRadius, ok := objMap["radius"]; ok {
		if valRadius != nil {
			var valueForRadius float64
			err = json.Unmarshal(*valRadius, &valueForRadius)
			if err != nil {
				return err
			}
			this.Radius = valueForRadius
		}
	}
	if valRadiusCap, ok := objMap["Radius"]; ok {
		if valRadiusCap != nil {
			var valueForRadius float64
			err = json.Unmarshal(*valRadiusCap, &valueForRadius)
			if err != nil {
				return err
			}
			this.Radius = valueForRadius
		}
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

    return nil
}
