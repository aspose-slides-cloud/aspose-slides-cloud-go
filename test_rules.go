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

type TestRules struct {
    Files []FileRule `json:"Files,omitempty"`
    Values []ValueRule `json:"Values,omitempty"`
    OkToNotFail []OkToNotFailRule `json:"OKToNotFail,omitempty"`
    Results []ResultRule `json:"Results,omitempty"`
}

type ITestRule interface {
    getInvalid() *bool
    getParameter() string
    getMethod() string
    getType() string
    getLanguage() string
}

type FileRule struct {
    Invalid *bool `json:"Invalid,omitempty"`
    Parameter string `json:"Parameter,omitempty"`
    Method string `json:"Method,omitempty"`
    Language string `json:"Language,omitempty"`
    Action string `json:"Action,omitempty"`
    File string `json:"File,omitempty"`
    Folder string `json:"Folder,omitempty"`
    Path string `json:"Path,omitempty"`
    Type string `json:"Type,omitempty"`
    ActualName string
}

func (this FileRule) getInvalid() *bool {
	return this.Invalid
}

func (this FileRule) getParameter() string {
	return this.Parameter
}

func (this FileRule) getMethod() string {
	return this.Method
}

func (this FileRule) getType() string {
	return this.Type
}

func (this FileRule) getLanguage() string {
	return this.Language
}

type ValueRule struct {
    Invalid *bool `json:"Invalid,omitempty"`
    Parameter string `json:"Parameter,omitempty"`
    Method string `json:"Method,omitempty"`
    Language string `json:"Language,omitempty"`
    Value interface{} `json:"Value,omitempty"`
    ValueSet bool
    InvalidValue interface{} `json:"InvalidValue,omitempty"`
    InvalidValueSet bool
    Type string `json:"Type,omitempty"`
}

func (this ValueRule) getInvalid() *bool {
	return this.Invalid
}

func (this ValueRule) getParameter() string {
	return this.Parameter
}

func (this ValueRule) getMethod() string {
	return this.Method
}

func (this ValueRule) getType() string {
	return this.Type
}

func (this ValueRule) getLanguage() string {
	return this.Language
}

func (this *ValueRule) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valInvalid, ok := objMap["Invalid"]; ok {
		if valInvalid != nil {
			var valueForInvalid *bool
			err = json.Unmarshal(*valInvalid, &valueForInvalid)
			if err != nil {
				return err
			}
			this.Invalid = valueForInvalid
		}
	}

	if valParameter, ok := objMap["Parameter"]; ok {
		if valParameter != nil {
			var valueForParameter string
			err = json.Unmarshal(*valParameter, &valueForParameter)
			if err != nil {
				return err
			}
			this.Parameter = valueForParameter
		}
	}

	if valMethod, ok := objMap["Method"]; ok {
		if valMethod != nil {
			var valueForMethod string
			err = json.Unmarshal(*valMethod, &valueForMethod)
			if err != nil {
				return err
			}
			this.Method = valueForMethod
		}
	}

	if valLanguage, ok := objMap["Language"]; ok {
		if valLanguage != nil {
			var valueForLanguage string
			err = json.Unmarshal(*valLanguage, &valueForLanguage)
			if err != nil {
				return err
			}
			this.Language = valueForLanguage
		}
	}

	if valValue, ok := objMap["Value"]; ok {
		this.ValueSet = true
		if valValue != nil {
			var valueForValue interface{}
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}

	if valInvalidValue, ok := objMap["InvalidValue"]; ok {
		this.InvalidValueSet = true
		if valInvalidValue != nil {
			var valueForInvalidValue interface{}
			err = json.Unmarshal(*valInvalidValue, &valueForInvalidValue)
			if err != nil {
				return err
			}
			this.InvalidValue = valueForInvalidValue
		}
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type = valueForType
		}
	}

    return nil
}

type OkToNotFailRule struct {
    Invalid *bool `json:"Invalid,omitempty"`
    Parameter string `json:"Parameter,omitempty"`
    Method string `json:"Method,omitempty"`
    Type string `json:"Type,omitempty"`
    Language string `json:"Language,omitempty"`
}

func (this OkToNotFailRule) getInvalid() *bool {
	return this.Invalid
}

func (this OkToNotFailRule) getParameter() string {
	return this.Parameter
}

func (this OkToNotFailRule) getMethod() string {
	return this.Method
}

func (this OkToNotFailRule) getType() string {
	return this.Type
}

func (this OkToNotFailRule) getLanguage() string {
	return this.Language
}

type ResultRule struct {
    Invalid *bool `json:"Invalid,omitempty"`
    Parameter string `json:"Parameter,omitempty"`
    Method string `json:"Method,omitempty"`
    Type string `json:"Type,omitempty"`
    Language string `json:"Language,omitempty"`
    Code int32 `json:"Code,omitempty"`
    Message string `json:"Message,omitempty"`
}

func (this ResultRule) getInvalid() *bool {
	return this.Invalid
}

func (this ResultRule) getParameter() string {
	return this.Parameter
}

func (this ResultRule) getMethod() string {
	return this.Method
}

func (this ResultRule) getType() string {
	return this.Type
}

func (this ResultRule) getLanguage() string {
	return this.Language
}
