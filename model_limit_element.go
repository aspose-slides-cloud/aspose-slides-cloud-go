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

// Specifies the Limit object
type ILimitElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Limit
	GetLimit() IMathElement
	SetLimit(newValue IMathElement)

	// Specifies upper or lower limit
	GetUpperLimit() *bool
	SetUpperLimit(newValue *bool)
}

type LimitElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base
	Base IMathElement `json:"Base,omitempty"`

	// Limit
	Limit IMathElement `json:"Limit,omitempty"`

	// Specifies upper or lower limit
	UpperLimit *bool `json:"UpperLimit"`
}

func NewLimitElement() *LimitElement {
	instance := new(LimitElement)
	instance.Type_ = "Limit"
	return instance
}

func (this *LimitElement) GetType() string {
	return this.Type_
}

func (this *LimitElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *LimitElement) GetBase() IMathElement {
	return this.Base
}

func (this *LimitElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *LimitElement) GetLimit() IMathElement {
	return this.Limit
}

func (this *LimitElement) SetLimit(newValue IMathElement) {
	this.Limit = newValue
}
func (this *LimitElement) GetUpperLimit() *bool {
	return this.UpperLimit
}

func (this *LimitElement) SetUpperLimit(newValue *bool) {
	this.UpperLimit = newValue
}

func (this *LimitElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Limit"
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
	
	if valLimit, ok := GetMapValue(objMap, "limit"); ok {
		if valLimit != nil {
			var valueForLimit MathElement
			err = json.Unmarshal(*valLimit, &valueForLimit)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valLimit)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLimit, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Limit = vInterfaceObject
			}
		}
	}
	
	if valUpperLimit, ok := GetMapValue(objMap, "upperLimit"); ok {
		if valUpperLimit != nil {
			var valueForUpperLimit *bool
			err = json.Unmarshal(*valUpperLimit, &valueForUpperLimit)
			if err != nil {
				return err
			}
			this.UpperLimit = valueForUpperLimit
		}
	}

	return nil
}
