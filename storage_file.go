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
)

import (
	"encoding/json"
)

// File or folder information
type IStorageFile interface {

	// File or folder name.
	getName() string
	setName(newValue string)

	// True if it is a folder.
	getIsFolder() bool
	setIsFolder(newValue bool)

	// File or folder last modified DateTime.
	getModifiedDate() time.Time
	setModifiedDate(newValue time.Time)

	// File or folder size.
	getSize() int64
	setSize(newValue int64)

	// File or folder path.
	getPath() string
	setPath(newValue string)
}

type StorageFile struct {

	// File or folder name.
	Name string `json:"Name,omitempty"`

	// True if it is a folder.
	IsFolder bool `json:"IsFolder"`

	// File or folder last modified DateTime.
	ModifiedDate time.Time `json:"ModifiedDate,omitempty"`

	// File or folder size.
	Size int64 `json:"Size"`

	// File or folder path.
	Path string `json:"Path,omitempty"`
}

func (this StorageFile) getName() string {
	return this.Name
}

func (this StorageFile) setName(newValue string) {
	this.Name = newValue
}
func (this StorageFile) getIsFolder() bool {
	return this.IsFolder
}

func (this StorageFile) setIsFolder(newValue bool) {
	this.IsFolder = newValue
}
func (this StorageFile) getModifiedDate() time.Time {
	return this.ModifiedDate
}

func (this StorageFile) setModifiedDate(newValue time.Time) {
	this.ModifiedDate = newValue
}
func (this StorageFile) getSize() int64 {
	return this.Size
}

func (this StorageFile) setSize(newValue int64) {
	this.Size = newValue
}
func (this StorageFile) getPath() string {
	return this.Path
}

func (this StorageFile) setPath(newValue string) {
	this.Path = newValue
}

func (this *StorageFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valName, ok := objMap["Name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valIsFolder, ok := objMap["IsFolder"]; ok {
		if valIsFolder != nil {
			var valueForIsFolder bool
			err = json.Unmarshal(*valIsFolder, &valueForIsFolder)
			if err != nil {
				return err
			}
			this.IsFolder = valueForIsFolder
		}
	}
	
	if valModifiedDate, ok := objMap["ModifiedDate"]; ok {
		if valModifiedDate != nil {
			var valueForModifiedDate time.Time
			err = json.Unmarshal(*valModifiedDate, &valueForModifiedDate)
			if err != nil {
				return err
			}
			this.ModifiedDate = valueForModifiedDate
		}
	}
	
	if valSize, ok := objMap["Size"]; ok {
		if valSize != nil {
			var valueForSize int64
			err = json.Unmarshal(*valSize, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
		}
	}
	
	if valPath, ok := objMap["Path"]; ok {
		if valPath != nil {
			var valueForPath string
			err = json.Unmarshal(*valPath, &valueForPath)
			if err != nil {
				return err
			}
			this.Path = valueForPath
		}
	}

    return nil
}
