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

// Represents a Color Replacement effect.
type IColorReplaceEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Returns color format which will replace color of every pixel.
	GetColor() string
	SetColor(newValue string)
}

type ColorReplaceEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Returns color format which will replace color of every pixel.
	Color string `json:"Color,omitempty"`
}

func NewColorReplaceEffect() *ColorReplaceEffect {
	instance := new(ColorReplaceEffect)
	instance.Type_ = "ColorReplace"
	return instance
}

func (this *ColorReplaceEffect) GetType() string {
	return this.Type_
}

func (this *ColorReplaceEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ColorReplaceEffect) GetColor() string {
	return this.Color
}

func (this *ColorReplaceEffect) SetColor(newValue string) {
	this.Color = newValue
}

func (this *ColorReplaceEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ColorReplace"
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
