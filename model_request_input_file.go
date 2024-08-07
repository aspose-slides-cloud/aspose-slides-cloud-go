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

// Represents input file from multipart request.
type IRequestInputFile interface {

	// Get or sets password to open document.
	GetPassword() string
	SetPassword(newValue string)

	// File type.
	GetType() string
	SetType(newValue string)

	// Get or sets index of file from request.
	GetIndex() int32
	SetIndex(newValue int32)
}

type RequestInputFile struct {

	// Get or sets password to open document.
	Password string `json:"Password,omitempty"`

	// File type.
	Type_ string `json:"Type"`

	// Get or sets index of file from request.
	Index int32 `json:"Index"`
}

func NewRequestInputFile() *RequestInputFile {
	instance := new(RequestInputFile)
	instance.Type_ = "Request"
	return instance
}

func (this *RequestInputFile) GetPassword() string {
	return this.Password
}

func (this *RequestInputFile) SetPassword(newValue string) {
	this.Password = newValue
}
func (this *RequestInputFile) GetType() string {
	return this.Type_
}

func (this *RequestInputFile) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *RequestInputFile) GetIndex() int32 {
	return this.Index
}

func (this *RequestInputFile) SetIndex(newValue int32) {
	this.Index = newValue
}

func (this *RequestInputFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPassword, ok := GetMapValue(objMap, "password"); ok {
		if valPassword != nil {
			var valueForPassword string
			err = json.Unmarshal(*valPassword, &valueForPassword)
			if err != nil {
				return err
			}
			this.Password = valueForPassword
		}
	}
	this.Type_ = "Request"
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
	
	if valIndex, ok := GetMapValue(objMap, "index"); ok {
		if valIndex != nil {
			var valueForIndex int32
			err = json.Unmarshal(*valIndex, &valueForIndex)
			if err != nil {
				return err
			}
			this.Index = valueForIndex
		}
	}

	return nil
}
