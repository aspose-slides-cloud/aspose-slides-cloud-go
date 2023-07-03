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

// VBA reference
type IVbaReference interface {

	// Name
	GetName() string
	SetName(newValue string)

	// Library ID
	GetLibId() string
	SetLibId(newValue string)
}

type VbaReference struct {

	// Name
	Name string `json:"Name,omitempty"`

	// Library ID
	LibId string `json:"LibId,omitempty"`
}

func NewVbaReference() *VbaReference {
	instance := new(VbaReference)
	return instance
}

func (this *VbaReference) GetName() string {
	return this.Name
}

func (this *VbaReference) SetName(newValue string) {
	this.Name = newValue
}
func (this *VbaReference) GetLibId() string {
	return this.LibId
}

func (this *VbaReference) SetLibId(newValue string) {
	this.LibId = newValue
}

func (this *VbaReference) UnmarshalJSON(b []byte) error {
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
	
	if valLibId, ok := objMap["libId"]; ok {
		if valLibId != nil {
			var valueForLibId string
			err = json.Unmarshal(*valLibId, &valueForLibId)
			if err != nil {
				return err
			}
			this.LibId = valueForLibId
		}
	}
	if valLibIdCap, ok := objMap["LibId"]; ok {
		if valLibIdCap != nil {
			var valueForLibId string
			err = json.Unmarshal(*valLibIdCap, &valueForLibId)
			if err != nil {
				return err
			}
			this.LibId = valueForLibId
		}
	}

	return nil
}
