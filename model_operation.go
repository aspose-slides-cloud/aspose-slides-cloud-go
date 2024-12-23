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
	"time"
	"encoding/json"
)


type IOperation interface {

	GetId() string
	SetId(newValue string)

	GetMethod() string
	SetMethod(newValue string)

	GetStatus() string
	SetStatus(newValue string)

	GetProgress() IOperationProgress
	SetProgress(newValue IOperationProgress)

	GetCreated() time.Time
	SetCreated(newValue time.Time)

	GetStarted() time.Time
	SetStarted(newValue time.Time)

	GetFailed() time.Time
	SetFailed(newValue time.Time)

	GetCanceled() time.Time
	SetCanceled(newValue time.Time)

	GetFinished() time.Time
	SetFinished(newValue time.Time)

	GetError() IOperationError
	SetError(newValue IOperationError)
}

type Operation struct {

	Id string `json:"Id"`

	Method string `json:"Method"`

	Status string `json:"Status"`

	Progress IOperationProgress `json:"Progress,omitempty"`

	Created time.Time `json:"Created,omitempty"`

	Started time.Time `json:"Started,omitempty"`

	Failed time.Time `json:"Failed,omitempty"`

	Canceled time.Time `json:"Canceled,omitempty"`

	Finished time.Time `json:"Finished,omitempty"`

	Error_ IOperationError `json:"Error,omitempty"`
}

func NewOperation() *Operation {
	instance := new(Operation)
	instance.Method = "Convert"
	instance.Status = "Created"
	return instance
}

func (this *Operation) GetId() string {
	return this.Id
}

func (this *Operation) SetId(newValue string) {
	this.Id = newValue
}
func (this *Operation) GetMethod() string {
	return this.Method
}

func (this *Operation) SetMethod(newValue string) {
	this.Method = newValue
}
func (this *Operation) GetStatus() string {
	return this.Status
}

func (this *Operation) SetStatus(newValue string) {
	this.Status = newValue
}
func (this *Operation) GetProgress() IOperationProgress {
	return this.Progress
}

func (this *Operation) SetProgress(newValue IOperationProgress) {
	this.Progress = newValue
}
func (this *Operation) GetCreated() time.Time {
	return this.Created
}

func (this *Operation) SetCreated(newValue time.Time) {
	this.Created = newValue
}
func (this *Operation) GetStarted() time.Time {
	return this.Started
}

func (this *Operation) SetStarted(newValue time.Time) {
	this.Started = newValue
}
func (this *Operation) GetFailed() time.Time {
	return this.Failed
}

func (this *Operation) SetFailed(newValue time.Time) {
	this.Failed = newValue
}
func (this *Operation) GetCanceled() time.Time {
	return this.Canceled
}

func (this *Operation) SetCanceled(newValue time.Time) {
	this.Canceled = newValue
}
func (this *Operation) GetFinished() time.Time {
	return this.Finished
}

func (this *Operation) SetFinished(newValue time.Time) {
	this.Finished = newValue
}
func (this *Operation) GetError() IOperationError {
	return this.Error_
}

func (this *Operation) SetError(newValue IOperationError) {
	this.Error_ = newValue
}

func (this *Operation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valId, ok := GetMapValue(objMap, "id"); ok {
		if valId != nil {
			var valueForId string
			err = json.Unmarshal(*valId, &valueForId)
			if err != nil {
				return err
			}
			this.Id = valueForId
		}
	}
	this.Method = "Convert"
	if valMethod, ok := GetMapValue(objMap, "method"); ok {
		if valMethod != nil {
			var valueForMethod string
			err = json.Unmarshal(*valMethod, &valueForMethod)
			if err != nil {
				var valueForMethodInt int32
				err = json.Unmarshal(*valMethod, &valueForMethodInt)
				if err != nil {
					return err
				}
				this.Method = string(valueForMethodInt)
			} else {
				this.Method = valueForMethod
			}
		}
	}
	this.Status = "Created"
	if valStatus, ok := GetMapValue(objMap, "status"); ok {
		if valStatus != nil {
			var valueForStatus string
			err = json.Unmarshal(*valStatus, &valueForStatus)
			if err != nil {
				var valueForStatusInt int32
				err = json.Unmarshal(*valStatus, &valueForStatusInt)
				if err != nil {
					return err
				}
				this.Status = string(valueForStatusInt)
			} else {
				this.Status = valueForStatus
			}
		}
	}
	
	if valProgress, ok := GetMapValue(objMap, "progress"); ok {
		if valProgress != nil {
			var valueForProgress OperationProgress
			err = json.Unmarshal(*valProgress, &valueForProgress)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("OperationProgress", *valProgress)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valProgress, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IOperationProgress)
			if ok {
				this.Progress = vInterfaceObject
			}
		}
	}
	
	if valCreated, ok := GetMapValue(objMap, "created"); ok {
		if valCreated != nil {
			var valueForCreated time.Time
			valueForCreated, err = time.Parse("2006-01-02T21:36:33", string(*valCreated))
			if err == nil {
				this.Created = valueForCreated
			}
		}
	}
	
	if valStarted, ok := GetMapValue(objMap, "started"); ok {
		if valStarted != nil {
			var valueForStarted time.Time
			valueForStarted, err = time.Parse("2006-01-02T21:36:33", string(*valStarted))
			if err == nil {
				this.Started = valueForStarted
			}
		}
	}
	
	if valFailed, ok := GetMapValue(objMap, "failed"); ok {
		if valFailed != nil {
			var valueForFailed time.Time
			valueForFailed, err = time.Parse("2006-01-02T21:36:33", string(*valFailed))
			if err == nil {
				this.Failed = valueForFailed
			}
		}
	}
	
	if valCanceled, ok := GetMapValue(objMap, "canceled"); ok {
		if valCanceled != nil {
			var valueForCanceled time.Time
			valueForCanceled, err = time.Parse("2006-01-02T21:36:33", string(*valCanceled))
			if err == nil {
				this.Canceled = valueForCanceled
			}
		}
	}
	
	if valFinished, ok := GetMapValue(objMap, "finished"); ok {
		if valFinished != nil {
			var valueForFinished time.Time
			valueForFinished, err = time.Parse("2006-01-02T21:36:33", string(*valFinished))
			if err == nil {
				this.Finished = valueForFinished
			}
		}
	}
	
	if valError, ok := GetMapValue(objMap, "error"); ok {
		if valError != nil {
			var valueForError OperationError
			err = json.Unmarshal(*valError, &valueForError)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("OperationError", *valError)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valError, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IOperationError)
			if ok {
				this.Error_ = vInterfaceObject
			}
		}
	}

	return nil
}
