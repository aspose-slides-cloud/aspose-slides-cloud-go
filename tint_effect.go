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
	getType() string
	setType(newValue string)

	// Hue
	getHue() float64
	setHue(newValue float64)

	// Amount
	getAmount() float64
	setAmount(newValue float64)
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

func (this *TintEffect) getType() string {
	return this.Type_
}

func (this *TintEffect) setType(newValue string) {
	this.Type_ = newValue
}
func (this *TintEffect) getHue() float64 {
	return this.Hue
}

func (this *TintEffect) setHue(newValue float64) {
	this.Hue = newValue
}
func (this *TintEffect) getAmount() float64 {
	return this.Amount
}

func (this *TintEffect) setAmount(newValue float64) {
	this.Amount = newValue
}

func (this *TintEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Tint"
	if valType, ok := objMap["type"]; ok {
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
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valHue, ok := objMap["hue"]; ok {
		if valHue != nil {
			var valueForHue float64
			err = json.Unmarshal(*valHue, &valueForHue)
			if err != nil {
				return err
			}
			this.Hue = valueForHue
		}
	}
	if valHueCap, ok := objMap["Hue"]; ok {
		if valHueCap != nil {
			var valueForHue float64
			err = json.Unmarshal(*valHueCap, &valueForHue)
			if err != nil {
				return err
			}
			this.Hue = valueForHue
		}
	}
	
	if valAmount, ok := objMap["amount"]; ok {
		if valAmount != nil {
			var valueForAmount float64
			err = json.Unmarshal(*valAmount, &valueForAmount)
			if err != nil {
				return err
			}
			this.Amount = valueForAmount
		}
	}
	if valAmountCap, ok := objMap["Amount"]; ok {
		if valAmountCap != nil {
			var valueForAmount float64
			err = json.Unmarshal(*valAmountCap, &valueForAmount)
			if err != nil {
				return err
			}
			this.Amount = valueForAmount
		}
	}

	return nil
}
