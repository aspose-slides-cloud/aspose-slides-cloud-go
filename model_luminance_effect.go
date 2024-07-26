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

// Represents a Luminance effect.
type ILuminanceEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Brightness
	GetBrightness() float64
	SetBrightness(newValue float64)

	// Contrast
	GetContrast() float64
	SetContrast(newValue float64)
}

type LuminanceEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Brightness
	Brightness float64 `json:"Brightness"`

	// Contrast
	Contrast float64 `json:"Contrast"`
}

func NewLuminanceEffect() *LuminanceEffect {
	instance := new(LuminanceEffect)
	instance.Type_ = "Luminance"
	return instance
}

func (this *LuminanceEffect) GetType() string {
	return this.Type_
}

func (this *LuminanceEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *LuminanceEffect) GetBrightness() float64 {
	return this.Brightness
}

func (this *LuminanceEffect) SetBrightness(newValue float64) {
	this.Brightness = newValue
}
func (this *LuminanceEffect) GetContrast() float64 {
	return this.Contrast
}

func (this *LuminanceEffect) SetContrast(newValue float64) {
	this.Contrast = newValue
}

func (this *LuminanceEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Luminance"
	if valType, ok := GetMapValue(objMap, "type"); ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valBrightness, ok := GetMapValue(objMap, "brightness"); ok {
		if valBrightness != nil {
			var valueForBrightness float64
			err = json.Unmarshal(*valBrightness, &valueForBrightness)
			if err != nil {
				return err
			}
			this.Brightness = valueForBrightness
		}
	}
	
	if valContrast, ok := GetMapValue(objMap, "contrast"); ok {
		if valContrast != nil {
			var valueForContrast float64
			err = json.Unmarshal(*valContrast, &valueForContrast)
			if err != nil {
				return err
			}
			this.Contrast = valueForContrast
		}
	}

	return nil
}
