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

// Represents a Tint effect.
type ITintEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Hue
	GetHue() float64
	SetHue(newValue float64)

	// Amount
	GetAmount() float64
	SetAmount(newValue float64)
}

type TintEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Hue
	Hue float64 `json:"Hue"`

	// Amount
	Amount float64 `json:"Amount"`
}

func NewTintEffect() *TintEffect {
	instance := new(TintEffect)
	instance.Type_ = "Tint"
	return instance
}

func (this *TintEffect) GetType() string {
	return this.Type_
}

func (this *TintEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *TintEffect) GetHue() float64 {
	return this.Hue
}

func (this *TintEffect) SetHue(newValue float64) {
	this.Hue = newValue
}
func (this *TintEffect) GetAmount() float64 {
	return this.Amount
}

func (this *TintEffect) SetAmount(newValue float64) {
	this.Amount = newValue
}

func (this *TintEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Tint"
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
	
	if valAmount, ok := GetMapValue(objMap, "amount"); ok {
		if valAmount != nil {
			var valueForAmount float64
			err = json.Unmarshal(*valAmount, &valueForAmount)
			if err != nil {
				return err
			}
			this.Amount = valueForAmount
		}
	}

	return nil
}
