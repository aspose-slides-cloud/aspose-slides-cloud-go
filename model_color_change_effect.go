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

// Represents a Color Change effect.
type IColorChangeEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Color which will be replaced.
	GetFromColor() string
	SetFromColor(newValue string)

	// Color which will replace.
	GetToColor() string
	SetToColor(newValue string)
}

type ColorChangeEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Color which will be replaced.
	FromColor string `json:"FromColor,omitempty"`

	// Color which will replace.
	ToColor string `json:"ToColor,omitempty"`
}

func NewColorChangeEffect() *ColorChangeEffect {
	instance := new(ColorChangeEffect)
	instance.Type_ = "ColorChange"
	return instance
}

func (this *ColorChangeEffect) GetType() string {
	return this.Type_
}

func (this *ColorChangeEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ColorChangeEffect) GetFromColor() string {
	return this.FromColor
}

func (this *ColorChangeEffect) SetFromColor(newValue string) {
	this.FromColor = newValue
}
func (this *ColorChangeEffect) GetToColor() string {
	return this.ToColor
}

func (this *ColorChangeEffect) SetToColor(newValue string) {
	this.ToColor = newValue
}

func (this *ColorChangeEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ColorChange"
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
	
	if valFromColor, ok := GetMapValue(objMap, "fromColor"); ok {
		if valFromColor != nil {
			var valueForFromColor string
			err = json.Unmarshal(*valFromColor, &valueForFromColor)
			if err != nil {
				return err
			}
			this.FromColor = valueForFromColor
		}
	}
	
	if valToColor, ok := GetMapValue(objMap, "toColor"); ok {
		if valToColor != nil {
			var valueForToColor string
			err = json.Unmarshal(*valToColor, &valueForToColor)
			if err != nil {
				return err
			}
			this.ToColor = valueForToColor
		}
	}

	return nil
}
