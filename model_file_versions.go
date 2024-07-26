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

// File versions FileVersion.
type IFileVersions interface {

	// File versions FileVersion.
	GetValue() []IFileVersion
	SetValue(newValue []IFileVersion)
}

type FileVersions struct {

	// File versions FileVersion.
	Value []IFileVersion `json:"Value,omitempty"`
}

func NewFileVersions() *FileVersions {
	instance := new(FileVersions)
	return instance
}

func (this *FileVersions) GetValue() []IFileVersion {
	return this.Value
}

func (this *FileVersions) SetValue(newValue []IFileVersion) {
	this.Value = newValue
}

func (this *FileVersions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valValue, ok := GetMapValue(objMap, "value"); ok {
		if valValue != nil {
			var valueForValue []json.RawMessage
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			valueForIValue := make([]IFileVersion, len(valueForValue))
			for i, v := range valueForValue {
				vObject, err := createObjectForType("FileVersion", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIValue[i] = vObject.(IFileVersion)
				}
			}
			this.Value = valueForIValue
		}
	}

	return nil
}
