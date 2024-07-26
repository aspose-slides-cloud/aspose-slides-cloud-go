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

// Subscript object
type ISuperscriptElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base argument
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Superscript
	GetSuperscript() IMathElement
	SetSuperscript(newValue IMathElement)
}

type SuperscriptElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base argument
	Base IMathElement `json:"Base,omitempty"`

	// Superscript
	Superscript IMathElement `json:"Superscript,omitempty"`
}

func NewSuperscriptElement() *SuperscriptElement {
	instance := new(SuperscriptElement)
	instance.Type_ = "SuperscriptElement"
	return instance
}

func (this *SuperscriptElement) GetType() string {
	return this.Type_
}

func (this *SuperscriptElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SuperscriptElement) GetBase() IMathElement {
	return this.Base
}

func (this *SuperscriptElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *SuperscriptElement) GetSuperscript() IMathElement {
	return this.Superscript
}

func (this *SuperscriptElement) SetSuperscript(newValue IMathElement) {
	this.Superscript = newValue
}

func (this *SuperscriptElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "SuperscriptElement"
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
	
	if valSuperscript, ok := GetMapValue(objMap, "superscript"); ok {
		if valSuperscript != nil {
			var valueForSuperscript MathElement
			err = json.Unmarshal(*valSuperscript, &valueForSuperscript)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valSuperscript)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSuperscript, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Superscript = vInterfaceObject
			}
		}
	}

	return nil
}
