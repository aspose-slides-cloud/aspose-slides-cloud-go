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

// Specifies the fraction object, consisting of a numerator and denominator separated by a fraction bar.
type IFractionElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Fraction type
	GetFractionType() string
	SetFractionType(newValue string)

	// Numerator
	GetNumerator() IMathElement
	SetNumerator(newValue IMathElement)

	// Denominator
	GetDenominator() IMathElement
	SetDenominator(newValue IMathElement)
}

type FractionElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Fraction type
	FractionType string `json:"FractionType,omitempty"`

	// Numerator
	Numerator IMathElement `json:"Numerator,omitempty"`

	// Denominator
	Denominator IMathElement `json:"Denominator,omitempty"`
}

func NewFractionElement() *FractionElement {
	instance := new(FractionElement)
	instance.Type_ = "Fraction"
	instance.FractionType = ""
	return instance
}

func (this *FractionElement) GetType() string {
	return this.Type_
}

func (this *FractionElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *FractionElement) GetFractionType() string {
	return this.FractionType
}

func (this *FractionElement) SetFractionType(newValue string) {
	this.FractionType = newValue
}
func (this *FractionElement) GetNumerator() IMathElement {
	return this.Numerator
}

func (this *FractionElement) SetNumerator(newValue IMathElement) {
	this.Numerator = newValue
}
func (this *FractionElement) GetDenominator() IMathElement {
	return this.Denominator
}

func (this *FractionElement) SetDenominator(newValue IMathElement) {
	this.Denominator = newValue
}

func (this *FractionElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Fraction"
	if valType, ok := objMap["type"]; ok {
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
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	this.FractionType = ""
	if valFractionType, ok := objMap["fractionType"]; ok {
		if valFractionType != nil {
			var valueForFractionType string
			err = json.Unmarshal(*valFractionType, &valueForFractionType)
			if err != nil {
				var valueForFractionTypeInt int32
				err = json.Unmarshal(*valFractionType, &valueForFractionTypeInt)
				if err != nil {
					return err
				}
				this.FractionType = string(valueForFractionTypeInt)
			} else {
				this.FractionType = valueForFractionType
			}
		}
	}
	if valFractionTypeCap, ok := objMap["FractionType"]; ok {
		if valFractionTypeCap != nil {
			var valueForFractionType string
			err = json.Unmarshal(*valFractionTypeCap, &valueForFractionType)
			if err != nil {
				var valueForFractionTypeInt int32
				err = json.Unmarshal(*valFractionTypeCap, &valueForFractionTypeInt)
				if err != nil {
					return err
				}
				this.FractionType = string(valueForFractionTypeInt)
			} else {
				this.FractionType = valueForFractionType
			}
		}
	}
	
	if valNumerator, ok := objMap["numerator"]; ok {
		if valNumerator != nil {
			var valueForNumerator MathElement
			err = json.Unmarshal(*valNumerator, &valueForNumerator)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valNumerator)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNumerator, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Numerator = vInterfaceObject
			}
		}
	}
	if valNumeratorCap, ok := objMap["Numerator"]; ok {
		if valNumeratorCap != nil {
			var valueForNumerator MathElement
			err = json.Unmarshal(*valNumeratorCap, &valueForNumerator)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valNumeratorCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNumeratorCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Numerator = vInterfaceObject
			}
		}
	}
	
	if valDenominator, ok := objMap["denominator"]; ok {
		if valDenominator != nil {
			var valueForDenominator MathElement
			err = json.Unmarshal(*valDenominator, &valueForDenominator)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valDenominator)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDenominator, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Denominator = vInterfaceObject
			}
		}
	}
	if valDenominatorCap, ok := objMap["Denominator"]; ok {
		if valDenominatorCap != nil {
			var valueForDenominator MathElement
			err = json.Unmarshal(*valDenominatorCap, &valueForDenominator)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valDenominatorCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDenominatorCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Denominator = vInterfaceObject
			}
		}
	}

	return nil
}
