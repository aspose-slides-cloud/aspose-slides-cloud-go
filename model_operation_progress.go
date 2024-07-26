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

// Operation progress.
type IOperationProgress interface {

	// Description.
	GetDescription() string
	SetDescription(newValue string)

	// Current Step Index.
	GetStepIndex() int32
	SetStepIndex(newValue int32)

	// Current Step Index.
	GetStepCount() int32
	SetStepCount(newValue int32)
}

type OperationProgress struct {

	// Description.
	Description string `json:"Description,omitempty"`

	// Current Step Index.
	StepIndex int32 `json:"StepIndex"`

	// Current Step Index.
	StepCount int32 `json:"StepCount"`
}

func NewOperationProgress() *OperationProgress {
	instance := new(OperationProgress)
	return instance
}

func (this *OperationProgress) GetDescription() string {
	return this.Description
}

func (this *OperationProgress) SetDescription(newValue string) {
	this.Description = newValue
}
func (this *OperationProgress) GetStepIndex() int32 {
	return this.StepIndex
}

func (this *OperationProgress) SetStepIndex(newValue int32) {
	this.StepIndex = newValue
}
func (this *OperationProgress) GetStepCount() int32 {
	return this.StepCount
}

func (this *OperationProgress) SetStepCount(newValue int32) {
	this.StepCount = newValue
}

func (this *OperationProgress) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDescription, ok := GetMapValue(objMap, "description"); ok {
		if valDescription != nil {
			var valueForDescription string
			err = json.Unmarshal(*valDescription, &valueForDescription)
			if err != nil {
				return err
			}
			this.Description = valueForDescription
		}
	}
	
	if valStepIndex, ok := GetMapValue(objMap, "stepIndex"); ok {
		if valStepIndex != nil {
			var valueForStepIndex int32
			err = json.Unmarshal(*valStepIndex, &valueForStepIndex)
			if err != nil {
				return err
			}
			this.StepIndex = valueForStepIndex
		}
	}
	
	if valStepCount, ok := GetMapValue(objMap, "stepCount"); ok {
		if valStepCount != nil {
			var valueForStepCount int32
			err = json.Unmarshal(*valStepCount, &valueForStepCount)
			if err != nil {
				return err
			}
			this.StepCount = valueForStepCount
		}
	}

	return nil
}
