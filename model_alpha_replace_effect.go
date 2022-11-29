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

// Represents an Alpha Replace Effect effect.
type IAlphaReplaceEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// The new opacity value.
	GetAlpha() float64
	SetAlpha(newValue float64)
}

type AlphaReplaceEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// The new opacity value.
	Alpha float64 `json:"Alpha"`
}

func NewAlphaReplaceEffect() *AlphaReplaceEffect {
	instance := new(AlphaReplaceEffect)
	instance.Type_ = "AlphaReplace"
	return instance
}

func (this *AlphaReplaceEffect) GetType() string {
	return this.Type_
}

func (this *AlphaReplaceEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AlphaReplaceEffect) GetAlpha() float64 {
	return this.Alpha
}

func (this *AlphaReplaceEffect) SetAlpha(newValue float64) {
	this.Alpha = newValue
}

func (this *AlphaReplaceEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AlphaReplace"
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
	
	if valAlpha, ok := objMap["alpha"]; ok {
		if valAlpha != nil {
			var valueForAlpha float64
			err = json.Unmarshal(*valAlpha, &valueForAlpha)
			if err != nil {
				return err
			}
			this.Alpha = valueForAlpha
		}
	}
	if valAlphaCap, ok := objMap["Alpha"]; ok {
		if valAlphaCap != nil {
			var valueForAlpha float64
			err = json.Unmarshal(*valAlphaCap, &valueForAlpha)
			if err != nil {
				return err
			}
			this.Alpha = valueForAlpha
		}
	}

	return nil
}
