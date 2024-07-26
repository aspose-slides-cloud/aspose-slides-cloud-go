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

// Radical function
type IRadicalElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base argument
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Degree argument
	GetDegree() IMathElement
	SetDegree(newValue IMathElement)

	// Hide degree
	GetHideDegree() *bool
	SetHideDegree(newValue *bool)
}

type RadicalElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base argument
	Base IMathElement `json:"Base,omitempty"`

	// Degree argument
	Degree IMathElement `json:"Degree,omitempty"`

	// Hide degree
	HideDegree *bool `json:"HideDegree"`
}

func NewRadicalElement() *RadicalElement {
	instance := new(RadicalElement)
	instance.Type_ = "Radical"
	return instance
}

func (this *RadicalElement) GetType() string {
	return this.Type_
}

func (this *RadicalElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *RadicalElement) GetBase() IMathElement {
	return this.Base
}

func (this *RadicalElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *RadicalElement) GetDegree() IMathElement {
	return this.Degree
}

func (this *RadicalElement) SetDegree(newValue IMathElement) {
	this.Degree = newValue
}
func (this *RadicalElement) GetHideDegree() *bool {
	return this.HideDegree
}

func (this *RadicalElement) SetHideDegree(newValue *bool) {
	this.HideDegree = newValue
}

func (this *RadicalElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Radical"
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
	
	if valBase, ok := GetMapValue(objMap, "base"); ok {
		if valBase != nil {
			var valueForBase MathElement
			err = json.Unmarshal(*valBase, &valueForBase)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valBase)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBase, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Base = vInterfaceObject
			}
		}
	}
	
	if valDegree, ok := GetMapValue(objMap, "degree"); ok {
		if valDegree != nil {
			var valueForDegree MathElement
			err = json.Unmarshal(*valDegree, &valueForDegree)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valDegree)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDegree, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Degree = vInterfaceObject
			}
		}
	}
	
	if valHideDegree, ok := GetMapValue(objMap, "hideDegree"); ok {
		if valHideDegree != nil {
			var valueForHideDegree *bool
			err = json.Unmarshal(*valHideDegree, &valueForHideDegree)
			if err != nil {
				return err
			}
			this.HideDegree = valueForHideDegree
		}
	}

	return nil
}
