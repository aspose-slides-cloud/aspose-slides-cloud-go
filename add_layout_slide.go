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


type IAddLayoutSlide interface {

	getType() TaskType
	setType(newValue TaskType)

	getCloneFromFile() IInputFile
	setCloneFromFile(newValue IInputFile)

	getCloneFromPosition() int32
	setCloneFromPosition(newValue int32)
}

type AddLayoutSlide struct {

	Type_ TaskType `json:"Type"`

	CloneFromFile IInputFile `json:"CloneFromFile,omitempty"`

	CloneFromPosition int32 `json:"CloneFromPosition,omitempty"`
}

func (this AddLayoutSlide) getType() TaskType {
	return this.Type_
}

func (this AddLayoutSlide) setType(newValue TaskType) {
	this.Type_ = newValue
}
func (this AddLayoutSlide) getCloneFromFile() IInputFile {
	return this.CloneFromFile
}

func (this AddLayoutSlide) setCloneFromFile(newValue IInputFile) {
	this.CloneFromFile = newValue
}
func (this AddLayoutSlide) getCloneFromPosition() int32 {
	return this.CloneFromPosition
}

func (this AddLayoutSlide) setCloneFromPosition(newValue int32) {
	this.CloneFromPosition = newValue
}

func (this *AddLayoutSlide) UnmarshalJSON(b []byte) error {
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

	if valCloneFromFile, ok := objMap["CloneFromFile"]; ok {
		if valCloneFromFile != nil {
			var valueForCloneFromFile InputFile
			err = json.Unmarshal(*valCloneFromFile, &valueForCloneFromFile)
			if err != nil {
				return err
			}
			this.CloneFromFile = valueForCloneFromFile
		}
	}

	if valCloneFromPosition, ok := objMap["CloneFromPosition"]; ok {
		if valCloneFromPosition != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPosition, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}

    return nil
}
