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

// Represents pipeline for one input document.
type IPipeline interface {

	// Get or sets input document.
	getInput() IInput
	setInput(newValue IInput)

	// Get or sets list of tasks representing pipeline.
	getTasks() []ITask
	setTasks(newValue []ITask)
}

type Pipeline struct {

	// Get or sets input document.
	Input IInput `json:"Input,omitempty"`

	// Get or sets list of tasks representing pipeline.
	Tasks []ITask `json:"Tasks,omitempty"`
}

func NewPipeline() *Pipeline {
	instance := new(Pipeline)
	return instance
}

func (this *Pipeline) getInput() IInput {
	return this.Input
}

func (this *Pipeline) setInput(newValue IInput) {
	this.Input = newValue
}
func (this *Pipeline) getTasks() []ITask {
	return this.Tasks
}

func (this *Pipeline) setTasks(newValue []ITask) {
	this.Tasks = newValue
}

func (this *Pipeline) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valInput, ok := objMap["input"]; ok {
		if valInput != nil {
			var valueForInput Input
			err = json.Unmarshal(*valInput, &valueForInput)
			if err != nil {
				return err
			}
			this.Input = &valueForInput
		}
	}
	if valInputCap, ok := objMap["Input"]; ok {
		if valInputCap != nil {
			var valueForInput Input
			err = json.Unmarshal(*valInputCap, &valueForInput)
			if err != nil {
				return err
			}
			this.Input = &valueForInput
		}
	}
	
	if valTasks, ok := objMap["tasks"]; ok {
		if valTasks != nil {
			var valueForTasks []Task
			err = json.Unmarshal(*valTasks, &valueForTasks)
			if err != nil {
				return err
			}
			valueForITasks := make([]ITask, len(valueForTasks))
			for i, v := range valueForTasks {
				valueForITasks[i] = ITask(&v)
			}
			this.Tasks = valueForITasks
		}
	}
	if valTasksCap, ok := objMap["Tasks"]; ok {
		if valTasksCap != nil {
			var valueForTasks []Task
			err = json.Unmarshal(*valTasksCap, &valueForTasks)
			if err != nil {
				return err
			}
			valueForITasks := make([]ITask, len(valueForTasks))
			for i, v := range valueForTasks {
				valueForITasks[i] = ITask(&v)
			}
			this.Tasks = valueForITasks
		}
	}

    return nil
}
