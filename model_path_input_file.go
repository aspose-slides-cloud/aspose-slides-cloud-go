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

// Represents input file from filesystem.
type IPathInputFile interface {

	// Get or sets password to open document.
	GetPassword() string
	SetPassword(newValue string)

	// File type.
	GetType() string
	SetType(newValue string)

	// Get or sets path to file.
	GetPath() string
	SetPath(newValue string)

	// Get or sets name of storage.
	GetStorage() string
	SetStorage(newValue string)
}

type PathInputFile struct {

	// Get or sets password to open document.
	Password string `json:"Password,omitempty"`

	// File type.
	Type_ string `json:"Type"`

	// Get or sets path to file.
	Path string `json:"Path,omitempty"`

	// Get or sets name of storage.
	Storage string `json:"Storage,omitempty"`
}

func NewPathInputFile() *PathInputFile {
	instance := new(PathInputFile)
	instance.Type_ = "Path"
	return instance
}

func (this *PathInputFile) GetPassword() string {
	return this.Password
}

func (this *PathInputFile) SetPassword(newValue string) {
	this.Password = newValue
}
func (this *PathInputFile) GetType() string {
	return this.Type_
}

func (this *PathInputFile) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *PathInputFile) GetPath() string {
	return this.Path
}

func (this *PathInputFile) SetPath(newValue string) {
	this.Path = newValue
}
func (this *PathInputFile) GetStorage() string {
	return this.Storage
}

func (this *PathInputFile) SetStorage(newValue string) {
	this.Storage = newValue
}

func (this *PathInputFile) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "Path"
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
	
	if valPath, ok := GetMapValue(objMap, "path"); ok {
		if valPath != nil {
			var valueForPath string
			err = json.Unmarshal(*valPath, &valueForPath)
			if err != nil {
				return err
			}
			this.Path = valueForPath
		}
	}
	
	if valStorage, ok := GetMapValue(objMap, "storage"); ok {
		if valStorage != nil {
			var valueForStorage string
			err = json.Unmarshal(*valStorage, &valueForStorage)
			if err != nil {
				return err
			}
			this.Storage = valueForStorage
		}
	}

	return nil
}
