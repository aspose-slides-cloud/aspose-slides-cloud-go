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

// Add layout slide task.
type IAddLayoutSlide interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Source file.
	GetCloneFromFile() IInputFile
	SetCloneFromFile(newValue IInputFile)

	// Source layout slide position.
	GetCloneFromPosition() int32
	SetCloneFromPosition(newValue int32)
}

type AddLayoutSlide struct {

	// Task type.
	Type_ string `json:"Type"`

	// Source file.
	CloneFromFile IInputFile `json:"CloneFromFile,omitempty"`

	// Source layout slide position.
	CloneFromPosition int32 `json:"CloneFromPosition"`
}

func NewAddLayoutSlide() *AddLayoutSlide {
	instance := new(AddLayoutSlide)
	instance.Type_ = "AddLayoutSlide"
	return instance
}

func (this *AddLayoutSlide) GetType() string {
	return this.Type_
}

func (this *AddLayoutSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AddLayoutSlide) GetCloneFromFile() IInputFile {
	return this.CloneFromFile
}

func (this *AddLayoutSlide) SetCloneFromFile(newValue IInputFile) {
	this.CloneFromFile = newValue
}
func (this *AddLayoutSlide) GetCloneFromPosition() int32 {
	return this.CloneFromPosition
}

func (this *AddLayoutSlide) SetCloneFromPosition(newValue int32) {
	this.CloneFromPosition = newValue
}

func (this *AddLayoutSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AddLayoutSlide"
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
	
	if valCloneFromFile, ok := objMap["cloneFromFile"]; ok {
		if valCloneFromFile != nil {
			var valueForCloneFromFile InputFile
			err = json.Unmarshal(*valCloneFromFile, &valueForCloneFromFile)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("InputFile", *valCloneFromFile)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCloneFromFile, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IInputFile)
			if ok {
				this.CloneFromFile = vInterfaceObject
			}
		}
	}
	if valCloneFromFileCap, ok := objMap["CloneFromFile"]; ok {
		if valCloneFromFileCap != nil {
			var valueForCloneFromFile InputFile
			err = json.Unmarshal(*valCloneFromFileCap, &valueForCloneFromFile)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("InputFile", *valCloneFromFileCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCloneFromFileCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IInputFile)
			if ok {
				this.CloneFromFile = vInterfaceObject
			}
		}
	}
	
	if valCloneFromPosition, ok := objMap["cloneFromPosition"]; ok {
		if valCloneFromPosition != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPosition, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}
	if valCloneFromPositionCap, ok := objMap["CloneFromPosition"]; ok {
		if valCloneFromPositionCap != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPositionCap, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}

	return nil
}
