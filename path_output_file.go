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

	getType() OutputFileType
	setType(newValue OutputFileType)

	// Get or sets path to file.
	getPath() string
	setPath(newValue string)

	// Get or sets name of storage.
	getStorage() string
	setStorage(newValue string)
}

type PathOutputFile struct {

	Type_ OutputFileType `json:"Type"`

	// Get or sets path to file.
	Path string `json:"Path,omitempty"`

	// Get or sets name of storage.
	Storage string `json:"Storage,omitempty"`
}

func (this PathOutputFile) getType() OutputFileType {
	return this.Type_
}

func (this PathOutputFile) setType(newValue OutputFileType) {
	this.Type_ = newValue
}
func (this PathOutputFile) getPath() string {
	return this.Path
}

func (this PathOutputFile) setPath(newValue string) {
	this.Path = newValue
}
func (this PathOutputFile) getStorage() string {
	return this.Storage
}

func (this PathOutputFile) setStorage(newValue string) {
	this.Storage = newValue
}

func (this *PathOutputFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType OutputFileType
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
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

	if valStorage, ok := objMap["Storage"]; ok {
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
