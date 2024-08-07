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

// Represents base64 inline encoded file.
type IBase64InputFile interface {

	// Get or sets password to open document.
	GetPassword() string
	SetPassword(newValue string)

	// File type.
	GetType() string
	SetType(newValue string)

	// Get or sets base64 data.
	GetData() string
	SetData(newValue string)
}

type Base64InputFile struct {

	// Get or sets password to open document.
	Password string `json:"Password,omitempty"`

	// File type.
	Type_ string `json:"Type"`

	// Get or sets base64 data.
	Data string `json:"Data,omitempty"`
}

func NewBase64InputFile() *Base64InputFile {
	instance := new(Base64InputFile)
	instance.Type_ = "Base64"
	return instance
}

func (this *Base64InputFile) GetPassword() string {
	return this.Password
}

func (this *Base64InputFile) SetPassword(newValue string) {
	this.Password = newValue
}
func (this *Base64InputFile) GetType() string {
	return this.Type_
}

func (this *Base64InputFile) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Base64InputFile) GetData() string {
	return this.Data
}

func (this *Base64InputFile) SetData(newValue string) {
	this.Data = newValue
}

func (this *Base64InputFile) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "Base64"
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
	
	if valData, ok := GetMapValue(objMap, "data"); ok {
		if valData != nil {
			var valueForData string
			err = json.Unmarshal(*valData, &valueForData)
			if err != nil {
				return err
			}
			this.Data = valueForData
		}
	}

	return nil
}
