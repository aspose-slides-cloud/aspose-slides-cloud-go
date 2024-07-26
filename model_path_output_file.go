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

// Represents fileSystem file with path.
type IPathOutputFile interface {

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

type PathOutputFile struct {

	// File type.
	Type_ string `json:"Type"`

	// Get or sets path to file.
	Path string `json:"Path,omitempty"`

	// Get or sets name of storage.
	Storage string `json:"Storage,omitempty"`
}

func NewPathOutputFile() *PathOutputFile {
	instance := new(PathOutputFile)
	instance.Type_ = "Path"
	return instance
}

func (this *PathOutputFile) GetType() string {
	return this.Type_
}

func (this *PathOutputFile) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *PathOutputFile) GetPath() string {
	return this.Path
}

func (this *PathOutputFile) SetPath(newValue string) {
	this.Path = newValue
}
func (this *PathOutputFile) GetStorage() string {
	return this.Storage
}

func (this *PathOutputFile) SetStorage(newValue string) {
	this.Storage = newValue
}

func (this *PathOutputFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
