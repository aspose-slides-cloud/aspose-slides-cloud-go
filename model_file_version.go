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

// File Version
type IFileVersion interface {

	// File or folder name.
	GetName() string
	SetName(newValue string)

	// True if it is a folder.
	GetIsFolder() bool
	SetIsFolder(newValue bool)

	// File or folder last modified DateTime.
	GetModifiedDate() time.Time
	SetModifiedDate(newValue time.Time)

	// File or folder size.
	GetSize() int64
	SetSize(newValue int64)

	// File or folder path.
	GetPath() string
	SetPath(newValue string)

	// File Version ID.
	GetVersionId() string
	SetVersionId(newValue string)

	// Specifies whether the file is (true) or is not (false) the latest version of an file.
	GetIsLatest() bool
	SetIsLatest(newValue bool)
}

type FileVersion struct {

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

	// File Version ID.
	VersionId string `json:"VersionId,omitempty"`

	// Specifies whether the file is (true) or is not (false) the latest version of an file.
	IsLatest bool `json:"IsLatest"`
}

func NewFileVersion() *FileVersion {
	instance := new(FileVersion)
	return instance
}

func (this *FileVersion) GetName() string {
	return this.Name
}

func (this *FileVersion) SetName(newValue string) {
	this.Name = newValue
}
func (this *FileVersion) GetIsFolder() bool {
	return this.IsFolder
}

func (this *FileVersion) SetIsFolder(newValue bool) {
	this.IsFolder = newValue
}
func (this *FileVersion) GetModifiedDate() time.Time {
	return this.ModifiedDate
}

func (this *FileVersion) SetModifiedDate(newValue time.Time) {
	this.ModifiedDate = newValue
}
func (this *FileVersion) GetSize() int64 {
	return this.Size
}

func (this *FileVersion) SetSize(newValue int64) {
	this.Size = newValue
}
func (this *FileVersion) GetPath() string {
	return this.Path
}

func (this *FileVersion) SetPath(newValue string) {
	this.Path = newValue
}
func (this *FileVersion) GetVersionId() string {
	return this.VersionId
}

func (this *FileVersion) SetVersionId(newValue string) {
	this.VersionId = newValue
}
func (this *FileVersion) GetIsLatest() bool {
	return this.IsLatest
}

func (this *FileVersion) SetIsLatest(newValue bool) {
	this.IsLatest = newValue
}

func (this *FileVersion) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
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
	
	if valModifiedDate, ok := objMap["modifiedDate"]; ok {
		if valModifiedDate != nil {
			var valueForModifiedDate time.Time
			valueForModifiedDate, err = time.Parse("2006-01-02T21:36:33", string(*valModifiedDate))
			if err == nil {
				this.ModifiedDate = valueForModifiedDate
			}
		}
	}
	if valModifiedDateCap, ok := objMap["ModifiedDate"]; ok {
		if valModifiedDateCap != nil {
			var valueForModifiedDate time.Time
			valueForModifiedDate, err = time.Parse("2006-01-02T21:36:33", string(*valModifiedDateCap))
			if err == nil {
				this.ModifiedDate = valueForModifiedDate
			}
		}
	}
	
	if valSize, ok := objMap["size"]; ok {
		if valSize != nil {
			var valueForSize int64
			err = json.Unmarshal(*valSize, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
		}
	}
	if valSizeCap, ok := objMap["Size"]; ok {
		if valSizeCap != nil {
			var valueForSize int64
			err = json.Unmarshal(*valSizeCap, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
		}
	}
	
	if valPath, ok := objMap["path"]; ok {
		if valPath != nil {
			var valueForPath string
			err = json.Unmarshal(*valPath, &valueForPath)
			if err != nil {
				return err
			}
			this.Path = valueForPath
		}
	}
	if valPathCap, ok := objMap["Path"]; ok {
		if valPathCap != nil {
			var valueForPath string
			err = json.Unmarshal(*valPathCap, &valueForPath)
			if err != nil {
				return err
			}
			this.Path = valueForPath
		}
	}
	
	if valVersionId, ok := objMap["versionId"]; ok {
		if valVersionId != nil {
			var valueForVersionId string
			err = json.Unmarshal(*valVersionId, &valueForVersionId)
			if err != nil {
				return err
			}
			this.VersionId = valueForVersionId
		}
	}
	if valVersionIdCap, ok := objMap["VersionId"]; ok {
		if valVersionIdCap != nil {
			var valueForVersionId string
			err = json.Unmarshal(*valVersionIdCap, &valueForVersionId)
			if err != nil {
				return err
			}
			this.VersionId = valueForVersionId
		}
	}
	
	if valIsLatest, ok := objMap["isLatest"]; ok {
		if valIsLatest != nil {
			var valueForIsLatest bool
			err = json.Unmarshal(*valIsLatest, &valueForIsLatest)
			if err != nil {
				return err
			}
			this.IsLatest = valueForIsLatest
		}
	}
	if valIsLatestCap, ok := objMap["IsLatest"]; ok {
		if valIsLatestCap != nil {
			var valueForIsLatest bool
			err = json.Unmarshal(*valIsLatestCap, &valueForIsLatest)
			if err != nil {
				return err
			}
			this.IsLatest = valueForIsLatest
		}
	}

	return nil
}
