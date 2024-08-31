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


type IOperationError interface {

	GetCode() string
	SetCode(newValue string)

	GetDescription() string
	SetDescription(newValue string)

	GetHttpStatusCode() int32
	SetHttpStatusCode(newValue int32)

	GetMessage() string
	SetMessage(newValue string)
}

type OperationError struct {

	Code string `json:"Code,omitempty"`

	Description string `json:"Description,omitempty"`

	HttpStatusCode int32 `json:"HttpStatusCode"`

	Message string `json:"Message,omitempty"`
}

func NewOperationError() *OperationError {
	instance := new(OperationError)
	return instance
}

func (this *OperationError) GetCode() string {
	return this.Code
}

func (this *OperationError) SetCode(newValue string) {
	this.Code = newValue
}
func (this *OperationError) GetDescription() string {
	return this.Description
}

func (this *OperationError) SetDescription(newValue string) {
	this.Description = newValue
}
func (this *OperationError) GetHttpStatusCode() int32 {
	return this.HttpStatusCode
}

func (this *OperationError) SetHttpStatusCode(newValue int32) {
	this.HttpStatusCode = newValue
}
func (this *OperationError) GetMessage() string {
	return this.Message
}

func (this *OperationError) SetMessage(newValue string) {
	this.Message = newValue
}

func (this *OperationError) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valCode, ok := GetMapValue(objMap, "code"); ok {
		if valCode != nil {
			var valueForCode string
			err = json.Unmarshal(*valCode, &valueForCode)
			if err != nil {
				return err
			}
			this.Code = valueForCode
		}
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
	
	if valHttpStatusCode, ok := GetMapValue(objMap, "httpStatusCode"); ok {
		if valHttpStatusCode != nil {
			var valueForHttpStatusCode int32
			err = json.Unmarshal(*valHttpStatusCode, &valueForHttpStatusCode)
			if err != nil {
				return err
			}
			this.HttpStatusCode = valueForHttpStatusCode
		}
	}
	
	if valMessage, ok := GetMapValue(objMap, "message"); ok {
		if valMessage != nil {
			var valueForMessage string
			err = json.Unmarshal(*valMessage, &valueForMessage)
			if err != nil {
				return err
			}
			this.Message = valueForMessage
		}
	}

	return nil
}
