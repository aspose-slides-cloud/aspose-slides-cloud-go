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

// Storage exists
type IStorageExist interface {

	// Shows that the storage exists.             
	GetExists() *bool
	SetExists(newValue *bool)
}

type StorageExist struct {

	// Shows that the storage exists.             
	Exists *bool `json:"Exists"`
}

func NewStorageExist() *StorageExist {
	instance := new(StorageExist)
	return instance
}

func (this *StorageExist) GetExists() *bool {
	return this.Exists
}

func (this *StorageExist) SetExists(newValue *bool) {
	this.Exists = newValue
}

func (this *StorageExist) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valExists, ok := GetMapValue(objMap, "exists"); ok {
		if valExists != nil {
			var valueForExists *bool
			err = json.Unmarshal(*valExists, &valueForExists)
			if err != nil {
				return err
			}
			this.Exists = valueForExists
		}
	}

	return nil
}
