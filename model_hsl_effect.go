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

// Represents a Hue/Saturation/Luminance effect.
type IHslEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Hue
	GetHue() float64
	SetHue(newValue float64)

	// Saturation
	GetSaturation() float64
	SetSaturation(newValue float64)

	// Luminance
	GetLuminance() float64
	SetLuminance(newValue float64)
}

type HslEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Hue
	Hue float64 `json:"Hue"`

	// Saturation
	Saturation float64 `json:"Saturation"`

	// Luminance
	Luminance float64 `json:"Luminance"`
}

func NewHslEffect() *HslEffect {
	instance := new(HslEffect)
	instance.Type_ = "Hsl"
	return instance
}

func (this *HslEffect) GetType() string {
	return this.Type_
}

func (this *HslEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *HslEffect) GetHue() float64 {
	return this.Hue
}

func (this *HslEffect) SetHue(newValue float64) {
	this.Hue = newValue
}
func (this *HslEffect) GetSaturation() float64 {
	return this.Saturation
}

func (this *HslEffect) SetSaturation(newValue float64) {
	this.Saturation = newValue
}
func (this *HslEffect) GetLuminance() float64 {
	return this.Luminance
}

func (this *HslEffect) SetLuminance(newValue float64) {
	this.Luminance = newValue
}

func (this *HslEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Hsl"
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
	
	if valHue, ok := GetMapValue(objMap, "hue"); ok {
		if valHue != nil {
			var valueForHue float64
			err = json.Unmarshal(*valHue, &valueForHue)
			if err != nil {
				return err
			}
			this.Hue = valueForHue
		}
	}
	
	if valSaturation, ok := GetMapValue(objMap, "saturation"); ok {
		if valSaturation != nil {
			var valueForSaturation float64
			err = json.Unmarshal(*valSaturation, &valueForSaturation)
			if err != nil {
				return err
			}
			this.Saturation = valueForSaturation
		}
	}
	
	if valLuminance, ok := GetMapValue(objMap, "luminance"); ok {
		if valLuminance != nil {
			var valueForLuminance float64
			err = json.Unmarshal(*valLuminance, &valueForLuminance)
			if err != nil {
				return err
			}
			this.Luminance = valueForLuminance
		}
	}

	return nil
}
