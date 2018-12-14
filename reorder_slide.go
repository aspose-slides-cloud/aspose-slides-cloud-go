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


type IReorderSlide interface {

	getType() TaskType
	setType(newValue TaskType)

	getOldPosition() int32
	setOldPosition(newValue int32)

	getNewPosition() int32
	setNewPosition(newValue int32)
}

type ReorderSlide struct {

	Type_ TaskType `json:"Type"`

	OldPosition int32 `json:"OldPosition,omitempty"`

	NewPosition int32 `json:"NewPosition,omitempty"`
}

func (this ReorderSlide) getType() TaskType {
	return this.Type_
}

func (this ReorderSlide) setType(newValue TaskType) {
	this.Type_ = newValue
}
func (this ReorderSlide) getOldPosition() int32 {
	return this.OldPosition
}

func (this ReorderSlide) setOldPosition(newValue int32) {
	this.OldPosition = newValue
}
func (this ReorderSlide) getNewPosition() int32 {
	return this.NewPosition
}

func (this ReorderSlide) setNewPosition(newValue int32) {
	this.NewPosition = newValue
}

func (this *ReorderSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType TaskType
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valOldPosition, ok := objMap["OldPosition"]; ok {
		if valOldPosition != nil {
			var valueForOldPosition int32
			err = json.Unmarshal(*valOldPosition, &valueForOldPosition)
			if err != nil {
				return err
			}
			this.OldPosition = valueForOldPosition
		}
	}

	if valNewPosition, ok := objMap["NewPosition"]; ok {
		if valNewPosition != nil {
			var valueForNewPosition int32
			err = json.Unmarshal(*valNewPosition, &valueForNewPosition)
			if err != nil {
				return err
			}
			this.NewPosition = valueForNewPosition
		}
	}

    return nil
}
