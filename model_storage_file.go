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

// File or folder information
type IStorageFile interface {

	// File or folder name.
	GetName() string
	SetName(newValue string)

	// True if it is a folder.
	GetIsFolder() *bool
	SetIsFolder(newValue *bool)

	// File or folder last modified DateTime.
	GetModifiedDate() time.Time
	SetModifiedDate(newValue time.Time)

	// File or folder size.
	GetSize() int64
	SetSize(newValue int64)

	// File or folder path.
	GetPath() string
	SetPath(newValue string)
}

type StorageFile struct {

	// File or folder name.
	Name string `json:"Name,omitempty"`

	// True if it is a folder.
	IsFolder *bool `json:"IsFolder"`

	// File or folder last modified DateTime.
	ModifiedDate time.Time `json:"ModifiedDate,omitempty"`

	// File or folder size.
	Size int64 `json:"Size"`

	// File or folder path.
	Path string `json:"Path,omitempty"`
}

func NewStorageFile() *StorageFile {
	instance := new(StorageFile)
	return instance
}

func (this *StorageFile) GetName() string {
	return this.Name
}

func (this *StorageFile) SetName(newValue string) {
	this.Name = newValue
}
func (this *StorageFile) GetIsFolder() *bool {
	return this.IsFolder
}

func (this *StorageFile) SetIsFolder(newValue *bool) {
	this.IsFolder = newValue
}
func (this *StorageFile) GetModifiedDate() time.Time {
	return this.ModifiedDate
}

func (this *StorageFile) SetModifiedDate(newValue time.Time) {
	this.ModifiedDate = newValue
}
func (this *StorageFile) GetSize() int64 {
	return this.Size
}

func (this *StorageFile) SetSize(newValue int64) {
	this.Size = newValue
}
func (this *StorageFile) GetPath() string {
	return this.Path
}

func (this *StorageFile) SetPath(newValue string) {
	this.Path = newValue
}

func (this *StorageFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valIsFolder, ok := GetMapValue(objMap, "isFolder"); ok {
		if valIsFolder != nil {
			var valueForIsFolder *bool
			err = json.Unmarshal(*valIsFolder, &valueForIsFolder)
			if err != nil {
				return err
			}
			this.IsFolder = valueForIsFolder
		}
	}
	
	if valModifiedDate, ok := GetMapValue(objMap, "modifiedDate"); ok {
		if valModifiedDate != nil {
			var valueForModifiedDate time.Time
			valueForModifiedDate, err = time.Parse("2006-01-02T21:36:33", string(*valModifiedDate))
			if err == nil {
				this.ModifiedDate = valueForModifiedDate
			}
		}
	}
	
	if valSize, ok := GetMapValue(objMap, "size"); ok {
		if valSize != nil {
			var valueForSize int64
			err = json.Unmarshal(*valSize, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
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

	return nil
}
