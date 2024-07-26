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

// Error
type IModelError interface {

	// Code             
	GetCode() string
	SetCode(newValue string)

	// Message             
	GetMessage() string
	SetMessage(newValue string)

	// Description             
	GetDescription() string
	SetDescription(newValue string)

	// Inner Error             
	GetInnerError() IErrorDetails
	SetInnerError(newValue IErrorDetails)
}

type ModelError struct {

	// Code             
	Code string `json:"Code,omitempty"`

	// Message             
	Message string `json:"Message,omitempty"`

	// Description             
	Description string `json:"Description,omitempty"`

	// Inner Error             
	InnerError IErrorDetails `json:"InnerError,omitempty"`
}

func NewModelError() *ModelError {
	instance := new(ModelError)
	return instance
}

func (this *ModelError) GetCode() string {
	return this.Code
}

func (this *ModelError) SetCode(newValue string) {
	this.Code = newValue
}
func (this *ModelError) GetMessage() string {
	return this.Message
}

func (this *ModelError) SetMessage(newValue string) {
	this.Message = newValue
}
func (this *ModelError) GetDescription() string {
	return this.Description
}

func (this *ModelError) SetDescription(newValue string) {
	this.Description = newValue
}
func (this *ModelError) GetInnerError() IErrorDetails {
	return this.InnerError
}

func (this *ModelError) SetInnerError(newValue IErrorDetails) {
	this.InnerError = newValue
}

func (this *ModelError) UnmarshalJSON(b []byte) error {
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
	
	if valInnerError, ok := GetMapValue(objMap, "innerError"); ok {
		if valInnerError != nil {
			var valueForInnerError ErrorDetails
			err = json.Unmarshal(*valInnerError, &valueForInnerError)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ErrorDetails", *valInnerError)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valInnerError, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IErrorDetails)
			if ok {
				this.InnerError = vInterfaceObject
			}
		}
	}

	return nil
}
