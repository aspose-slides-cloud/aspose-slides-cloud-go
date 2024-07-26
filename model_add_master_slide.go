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

// Add master slide task.
type IAddMasterSlide interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Source presentation clone from.
	GetCloneFromFile() IInputFile
	SetCloneFromFile(newValue IInputFile)

	// Index of slide to clone.
	GetCloneFromPosition() int32
	SetCloneFromPosition(newValue int32)

	// True if cloned master slide is applied to all slides.
	GetApplyToAll() *bool
	SetApplyToAll(newValue *bool)
}

type AddMasterSlide struct {

	// Task type.
	Type_ string `json:"Type"`

	// Source presentation clone from.
	CloneFromFile IInputFile `json:"CloneFromFile,omitempty"`

	// Index of slide to clone.
	CloneFromPosition int32 `json:"CloneFromPosition"`

	// True if cloned master slide is applied to all slides.
	ApplyToAll *bool `json:"ApplyToAll"`
}

func NewAddMasterSlide() *AddMasterSlide {
	instance := new(AddMasterSlide)
	instance.Type_ = "AddMasterSlide"
	return instance
}

func (this *AddMasterSlide) GetType() string {
	return this.Type_
}

func (this *AddMasterSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AddMasterSlide) GetCloneFromFile() IInputFile {
	return this.CloneFromFile
}

func (this *AddMasterSlide) SetCloneFromFile(newValue IInputFile) {
	this.CloneFromFile = newValue
}
func (this *AddMasterSlide) GetCloneFromPosition() int32 {
	return this.CloneFromPosition
}

func (this *AddMasterSlide) SetCloneFromPosition(newValue int32) {
	this.CloneFromPosition = newValue
}
func (this *AddMasterSlide) GetApplyToAll() *bool {
	return this.ApplyToAll
}

func (this *AddMasterSlide) SetApplyToAll(newValue *bool) {
	this.ApplyToAll = newValue
}

func (this *AddMasterSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AddMasterSlide"
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
	
	if valCloneFromFile, ok := GetMapValue(objMap, "cloneFromFile"); ok {
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
	
	if valCloneFromPosition, ok := GetMapValue(objMap, "cloneFromPosition"); ok {
		if valCloneFromPosition != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPosition, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}
	
	if valApplyToAll, ok := GetMapValue(objMap, "applyToAll"); ok {
		if valApplyToAll != nil {
			var valueForApplyToAll *bool
			err = json.Unmarshal(*valApplyToAll, &valueForApplyToAll)
			if err != nil {
				return err
			}
			this.ApplyToAll = valueForApplyToAll
		}
	}

	return nil
}
