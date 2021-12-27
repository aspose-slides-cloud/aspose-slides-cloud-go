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

// Object exists
type IObjectExist interface {

	// Indicates that the file or folder exists.
	getExists() bool
	setExists(newValue bool)

	// True if it is a folder, false if it is a file.
	getIsFolder() bool
	setIsFolder(newValue bool)
}

type ObjectExist struct {

	// Indicates that the file or folder exists.
	Exists bool `json:"Exists"`

	// True if it is a folder, false if it is a file.
	IsFolder bool `json:"IsFolder"`
}

func NewObjectExist() *ObjectExist {
	instance := new(ObjectExist)
	return instance
}

func (this *ObjectExist) getExists() bool {
	return this.Exists
}

func (this *ObjectExist) setExists(newValue bool) {
	this.Exists = newValue
}
func (this *ObjectExist) getIsFolder() bool {
	return this.IsFolder
}

func (this *ObjectExist) setIsFolder(newValue bool) {
	this.IsFolder = newValue
}

func (this *ObjectExist) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valExists, ok := objMap["exists"]; ok {
		if valExists != nil {
			var valueForExists bool
			err = json.Unmarshal(*valExists, &valueForExists)
			if err != nil {
				return err
			}
			this.Exists = valueForExists
		}
	}
	if valExistsCap, ok := objMap["Exists"]; ok {
		if valExistsCap != nil {
			var valueForExists bool
			err = json.Unmarshal(*valExistsCap, &valueForExists)
			if err != nil {
				return err
			}
			this.Exists = valueForExists
		}
	}
	
	if valIsFolder, ok := objMap["isFolder"]; ok {
		if valIsFolder != nil {
			var valueForIsFolder bool
			err = json.Unmarshal(*valIsFolder, &valueForIsFolder)
			if err != nil {
				return err
			}
			this.IsFolder = valueForIsFolder
		}
	}
	if valIsFolderCap, ok := objMap["IsFolder"]; ok {
		if valIsFolderCap != nil {
			var valueForIsFolder bool
			err = json.Unmarshal(*valIsFolderCap, &valueForIsFolder)
			if err != nil {
				return err
			}
			this.IsFolder = valueForIsFolder
		}
	}

	return nil
}
