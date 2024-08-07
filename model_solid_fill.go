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

// Represents solid fill format 
type ISolidFill interface {

	// Fill type.
	GetType() string
	SetType(newValue string)

	// Color.
	GetColor() string
	SetColor(newValue string)
}

type SolidFill struct {

	// Fill type.
	Type_ string `json:"Type"`

	// Color.
	Color string `json:"Color,omitempty"`
}

func NewSolidFill() *SolidFill {
	instance := new(SolidFill)
	instance.Type_ = "Solid"
	return instance
}

func (this *SolidFill) GetType() string {
	return this.Type_
}

func (this *SolidFill) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SolidFill) GetColor() string {
	return this.Color
}

func (this *SolidFill) SetColor(newValue string) {
	this.Color = newValue
}

func (this *SolidFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Solid"
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
	
	if valColor, ok := GetMapValue(objMap, "color"); ok {
		if valColor != nil {
			var valueForColor string
			err = json.Unmarshal(*valColor, &valueForColor)
			if err != nil {
				return err
			}
			this.Color = valueForColor
		}
	}

	return nil
}
