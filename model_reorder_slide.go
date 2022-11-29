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

// Reorder slide task.
type IReorderSlide interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Old position.
	GetOldPosition() int32
	SetOldPosition(newValue int32)

	// New position.
	GetNewPosition() int32
	SetNewPosition(newValue int32)
}

type ReorderSlide struct {

	// Task type.
	Type_ string `json:"Type"`

	// Old position.
	OldPosition int32 `json:"OldPosition"`

	// New position.
	NewPosition int32 `json:"NewPosition"`
}

func NewReorderSlide() *ReorderSlide {
	instance := new(ReorderSlide)
	instance.Type_ = "ReoderSlide"
	return instance
}

func (this *ReorderSlide) GetType() string {
	return this.Type_
}

func (this *ReorderSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ReorderSlide) GetOldPosition() int32 {
	return this.OldPosition
}

func (this *ReorderSlide) SetOldPosition(newValue int32) {
	this.OldPosition = newValue
}
func (this *ReorderSlide) GetNewPosition() int32 {
	return this.NewPosition
}

func (this *ReorderSlide) SetNewPosition(newValue int32) {
	this.NewPosition = newValue
}

func (this *ReorderSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ReoderSlide"
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
	
	if valOldPosition, ok := objMap["oldPosition"]; ok {
		if valOldPosition != nil {
			var valueForOldPosition int32
			err = json.Unmarshal(*valOldPosition, &valueForOldPosition)
			if err != nil {
				return err
			}
			this.OldPosition = valueForOldPosition
		}
	}
	if valOldPositionCap, ok := objMap["OldPosition"]; ok {
		if valOldPositionCap != nil {
			var valueForOldPosition int32
			err = json.Unmarshal(*valOldPositionCap, &valueForOldPosition)
			if err != nil {
				return err
			}
			this.OldPosition = valueForOldPosition
		}
	}
	
	if valNewPosition, ok := objMap["newPosition"]; ok {
		if valNewPosition != nil {
			var valueForNewPosition int32
			err = json.Unmarshal(*valNewPosition, &valueForNewPosition)
			if err != nil {
				return err
			}
			this.NewPosition = valueForNewPosition
		}
	}
	if valNewPositionCap, ok := objMap["NewPosition"]; ok {
		if valNewPositionCap != nil {
			var valueForNewPosition int32
			err = json.Unmarshal(*valNewPositionCap, &valueForNewPosition)
			if err != nil {
				return err
			}
			this.NewPosition = valueForNewPosition
		}
	}

	return nil
}
