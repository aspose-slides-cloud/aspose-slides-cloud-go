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

// Reset slide task.
type IResetSlide interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Slide position.
	GetPosition() int32
	SetPosition(newValue int32)
}

type ResetSlide struct {

	// Task type.
	Type_ string `json:"Type"`

	// Slide position.
	Position int32 `json:"Position"`
}

func NewResetSlide() *ResetSlide {
	instance := new(ResetSlide)
	instance.Type_ = "ResetSlide"
	return instance
}

func (this *ResetSlide) GetType() string {
	return this.Type_
}

func (this *ResetSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ResetSlide) GetPosition() int32 {
	return this.Position
}

func (this *ResetSlide) SetPosition(newValue int32) {
	this.Position = newValue
}

func (this *ResetSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ResetSlide"
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
	
	if valPosition, ok := GetMapValue(objMap, "position"); ok {
		if valPosition != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}

	return nil
}
