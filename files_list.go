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

// Files list
type IFilesList interface {

	// Files and folders contained by folder StorageFile.
	getValue() []IStorageFile
	setValue(newValue []IStorageFile)
}

type FilesList struct {

	// Files and folders contained by folder StorageFile.
	Value []IStorageFile `json:"Value,omitempty"`
}

func NewFilesList() *FilesList {
	instance := new(FilesList)
	return instance
}

func (this *FilesList) getValue() []IStorageFile {
	return this.Value
}

func (this *FilesList) setValue(newValue []IStorageFile) {
	this.Value = newValue
}

func (this *FilesList) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valValue, ok := objMap["value"]; ok {
		if valValue != nil {
			var valueForValue []json.RawMessage
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			valueForIValue := make([]IStorageFile, len(valueForValue))
			for i, v := range valueForValue {
				vObject, err := createObjectForType("StorageFile", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIValue[i] = vObject.(IStorageFile)
				}
			}
			this.Value = valueForIValue
		}
	}
	if valValueCap, ok := objMap["Value"]; ok {
		if valValueCap != nil {
			var valueForValue []json.RawMessage
			err = json.Unmarshal(*valValueCap, &valueForValue)
			if err != nil {
				return err
			}
			valueForIValue := make([]IStorageFile, len(valueForValue))
			for i, v := range valueForValue {
				vObject, err := createObjectForType("StorageFile", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIValue[i] = vObject.(IStorageFile)
				}
			}
			this.Value = valueForIValue
		}
	}

	return nil
}
