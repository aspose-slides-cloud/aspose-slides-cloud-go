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

// Represents a Duotone effect.
type IDuotoneEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Returns target color format for dark pixels.
	GetColor1() string
	SetColor1(newValue string)

	// Returns target color format for light pixels.
	GetColor2() string
	SetColor2(newValue string)
}

type DuotoneEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Returns target color format for dark pixels.
	Color1 string `json:"Color1,omitempty"`

	// Returns target color format for light pixels.
	Color2 string `json:"Color2,omitempty"`
}

func NewDuotoneEffect() *DuotoneEffect {
	instance := new(DuotoneEffect)
	instance.Type_ = "Duotone"
	return instance
}

func (this *DuotoneEffect) GetType() string {
	return this.Type_
}

func (this *DuotoneEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *DuotoneEffect) GetColor1() string {
	return this.Color1
}

func (this *DuotoneEffect) SetColor1(newValue string) {
	this.Color1 = newValue
}
func (this *DuotoneEffect) GetColor2() string {
	return this.Color2
}

func (this *DuotoneEffect) SetColor2(newValue string) {
	this.Color2 = newValue
}

func (this *DuotoneEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Duotone"
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
	
	if valColor1, ok := GetMapValue(objMap, "color1"); ok {
		if valColor1 != nil {
			var valueForColor1 string
			err = json.Unmarshal(*valColor1, &valueForColor1)
			if err != nil {
				return err
			}
			this.Color1 = valueForColor1
		}
	}
	
	if valColor2, ok := GetMapValue(objMap, "color2"); ok {
		if valColor2 != nil {
			var valueForColor2 string
			err = json.Unmarshal(*valColor2, &valueForColor2)
			if err != nil {
				return err
			}
			this.Color2 = valueForColor2
		}
	}

	return nil
}
