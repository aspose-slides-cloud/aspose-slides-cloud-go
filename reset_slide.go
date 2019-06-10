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


type IResetSlide interface {

	getType() string
	setType(newValue string)

	getPosition() int32
	setPosition(newValue int32)
}

type ResetSlide struct {

	Type_ string `json:"Type"`

	Position int32 `json:"Position"`
}

func (this ResetSlide) getType() string {
	return this.Type_
}

func (this ResetSlide) setType(newValue string) {
	this.Type_ = newValue
}
func (this ResetSlide) getPosition() int32 {
	return this.Position
}

func (this ResetSlide) setPosition(newValue int32) {
	this.Position = newValue
}

func (this *ResetSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valPosition, ok := objMap["Position"]; ok {
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
