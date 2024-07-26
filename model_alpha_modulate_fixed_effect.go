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

// Represents an Alpha Modulate Fixed effect.
type IAlphaModulateFixedEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Returns an amount of effect in percents.    
	GetAmount() float64
	SetAmount(newValue float64)
}

type AlphaModulateFixedEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Returns an amount of effect in percents.    
	Amount float64 `json:"Amount"`
}

func NewAlphaModulateFixedEffect() *AlphaModulateFixedEffect {
	instance := new(AlphaModulateFixedEffect)
	instance.Type_ = "AlphaModulateFixed"
	return instance
}

func (this *AlphaModulateFixedEffect) GetType() string {
	return this.Type_
}

func (this *AlphaModulateFixedEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AlphaModulateFixedEffect) GetAmount() float64 {
	return this.Amount
}

func (this *AlphaModulateFixedEffect) SetAmount(newValue float64) {
	this.Amount = newValue
}

func (this *AlphaModulateFixedEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AlphaModulateFixed"
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
