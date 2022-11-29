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

// Specifies an N-ary mathematical object, such as Summation and Integral.
type INaryOperatorElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base argument
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Subscript argument
	GetSubscript() IMathElement
	SetSubscript(newValue IMathElement)

	// Superscript argument
	GetSuperscript() IMathElement
	SetSuperscript(newValue IMathElement)

	// Nary Operator Character
	GetOperator() string
	SetOperator(newValue string)

	// The location of limits (subscript and superscript)
	GetLimitLocation() string
	SetLimitLocation(newValue string)

	// Operator Character grows vertically to match its operand height
	GetGrowToMatchOperandHeight() bool
	SetGrowToMatchOperandHeight(newValue bool)

	// Hide Subscript
	GetHideSubscript() bool
	SetHideSubscript(newValue bool)

	// Hide Superscript
	GetHideSuperscript() bool
	SetHideSuperscript(newValue bool)
}

type NaryOperatorElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base argument
	Base IMathElement `json:"Base,omitempty"`

	// Subscript argument
	Subscript IMathElement `json:"Subscript,omitempty"`

	// Superscript argument
	Superscript IMathElement `json:"Superscript,omitempty"`

	// Nary Operator Character
	Operator string `json:"Operator,omitempty"`

	// The location of limits (subscript and superscript)
	LimitLocation string `json:"LimitLocation,omitempty"`

	// Operator Character grows vertically to match its operand height
	GrowToMatchOperandHeight bool `json:"GrowToMatchOperandHeight"`

	// Hide Subscript
	HideSubscript bool `json:"HideSubscript"`

	// Hide Superscript
	HideSuperscript bool `json:"HideSuperscript"`
}

func NewNaryOperatorElement() *NaryOperatorElement {
	instance := new(NaryOperatorElement)
	instance.Type_ = "NaryOperator"
	return instance
}

func (this *NaryOperatorElement) GetType() string {
	return this.Type_
}

func (this *NaryOperatorElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *NaryOperatorElement) GetBase() IMathElement {
	return this.Base
}

func (this *NaryOperatorElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *NaryOperatorElement) GetSubscript() IMathElement {
	return this.Subscript
}

func (this *NaryOperatorElement) SetSubscript(newValue IMathElement) {
	this.Subscript = newValue
}
func (this *NaryOperatorElement) GetSuperscript() IMathElement {
	return this.Superscript
}

func (this *NaryOperatorElement) SetSuperscript(newValue IMathElement) {
	this.Superscript = newValue
}
func (this *NaryOperatorElement) GetOperator() string {
	return this.Operator
}

func (this *NaryOperatorElement) SetOperator(newValue string) {
	this.Operator = newValue
}
func (this *NaryOperatorElement) GetLimitLocation() string {
	return this.LimitLocation
}

func (this *NaryOperatorElement) SetLimitLocation(newValue string) {
	this.LimitLocation = newValue
}
func (this *NaryOperatorElement) GetGrowToMatchOperandHeight() bool {
	return this.GrowToMatchOperandHeight
}

func (this *NaryOperatorElement) SetGrowToMatchOperandHeight(newValue bool) {
	this.GrowToMatchOperandHeight = newValue
}
func (this *NaryOperatorElement) GetHideSubscript() bool {
	return this.HideSubscript
}

func (this *NaryOperatorElement) SetHideSubscript(newValue bool) {
	this.HideSubscript = newValue
}
func (this *NaryOperatorElement) GetHideSuperscript() bool {
	return this.HideSuperscript
}

func (this *NaryOperatorElement) SetHideSuperscript(newValue bool) {
	this.HideSuperscript = newValue
}

func (this *NaryOperatorElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "NaryOperator"
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
	
	if valBase, ok := objMap["base"]; ok {
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
	if valBaseCap, ok := objMap["Base"]; ok {
		if valBaseCap != nil {
			var valueForBase MathElement
			err = json.Unmarshal(*valBaseCap, &valueForBase)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valBaseCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBaseCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Base = vInterfaceObject
			}
		}
	}
	
	if valSubscript, ok := objMap["subscript"]; ok {
		if valSubscript != nil {
			var valueForSubscript MathElement
			err = json.Unmarshal(*valSubscript, &valueForSubscript)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valSubscript)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSubscript, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Subscript = vInterfaceObject
			}
		}
	}
	if valSubscriptCap, ok := objMap["Subscript"]; ok {
		if valSubscriptCap != nil {
			var valueForSubscript MathElement
			err = json.Unmarshal(*valSubscriptCap, &valueForSubscript)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valSubscriptCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSubscriptCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Subscript = vInterfaceObject
			}
		}
	}
	
	if valSuperscript, ok := objMap["superscript"]; ok {
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
	if valSuperscriptCap, ok := objMap["Superscript"]; ok {
		if valSuperscriptCap != nil {
			var valueForSuperscript MathElement
			err = json.Unmarshal(*valSuperscriptCap, &valueForSuperscript)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("MathElement", *valSuperscriptCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSuperscriptCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IMathElement)
			if ok {
				this.Superscript = vInterfaceObject
			}
		}
	}
	
	if valOperator, ok := objMap["operator"]; ok {
		if valOperator != nil {
			var valueForOperator string
			err = json.Unmarshal(*valOperator, &valueForOperator)
			if err != nil {
				return err
			}
			this.Operator = valueForOperator
		}
	}
	if valOperatorCap, ok := objMap["Operator"]; ok {
		if valOperatorCap != nil {
			var valueForOperator string
			err = json.Unmarshal(*valOperatorCap, &valueForOperator)
			if err != nil {
				return err
			}
			this.Operator = valueForOperator
		}
	}
	
	if valLimitLocation, ok := objMap["limitLocation"]; ok {
		if valLimitLocation != nil {
			var valueForLimitLocation string
			err = json.Unmarshal(*valLimitLocation, &valueForLimitLocation)
			if err != nil {
				var valueForLimitLocationInt int32
				err = json.Unmarshal(*valLimitLocation, &valueForLimitLocationInt)
				if err != nil {
					return err
				}
				this.LimitLocation = string(valueForLimitLocationInt)
			} else {
				this.LimitLocation = valueForLimitLocation
			}
		}
	}
	if valLimitLocationCap, ok := objMap["LimitLocation"]; ok {
		if valLimitLocationCap != nil {
			var valueForLimitLocation string
			err = json.Unmarshal(*valLimitLocationCap, &valueForLimitLocation)
			if err != nil {
				var valueForLimitLocationInt int32
				err = json.Unmarshal(*valLimitLocationCap, &valueForLimitLocationInt)
				if err != nil {
					return err
				}
				this.LimitLocation = string(valueForLimitLocationInt)
			} else {
				this.LimitLocation = valueForLimitLocation
			}
		}
	}
	
	if valGrowToMatchOperandHeight, ok := objMap["growToMatchOperandHeight"]; ok {
		if valGrowToMatchOperandHeight != nil {
			var valueForGrowToMatchOperandHeight bool
			err = json.Unmarshal(*valGrowToMatchOperandHeight, &valueForGrowToMatchOperandHeight)
			if err != nil {
				return err
			}
			this.GrowToMatchOperandHeight = valueForGrowToMatchOperandHeight
		}
	}
	if valGrowToMatchOperandHeightCap, ok := objMap["GrowToMatchOperandHeight"]; ok {
		if valGrowToMatchOperandHeightCap != nil {
			var valueForGrowToMatchOperandHeight bool
			err = json.Unmarshal(*valGrowToMatchOperandHeightCap, &valueForGrowToMatchOperandHeight)
			if err != nil {
				return err
			}
			this.GrowToMatchOperandHeight = valueForGrowToMatchOperandHeight
		}
	}
	
	if valHideSubscript, ok := objMap["hideSubscript"]; ok {
		if valHideSubscript != nil {
			var valueForHideSubscript bool
			err = json.Unmarshal(*valHideSubscript, &valueForHideSubscript)
			if err != nil {
				return err
			}
			this.HideSubscript = valueForHideSubscript
		}
	}
	if valHideSubscriptCap, ok := objMap["HideSubscript"]; ok {
		if valHideSubscriptCap != nil {
			var valueForHideSubscript bool
			err = json.Unmarshal(*valHideSubscriptCap, &valueForHideSubscript)
			if err != nil {
				return err
			}
			this.HideSubscript = valueForHideSubscript
		}
	}
	
	if valHideSuperscript, ok := objMap["hideSuperscript"]; ok {
		if valHideSuperscript != nil {
			var valueForHideSuperscript bool
			err = json.Unmarshal(*valHideSuperscript, &valueForHideSuperscript)
			if err != nil {
				return err
			}
			this.HideSuperscript = valueForHideSuperscript
		}
	}
	if valHideSuperscriptCap, ok := objMap["HideSuperscript"]; ok {
		if valHideSuperscriptCap != nil {
			var valueForHideSuperscript bool
			err = json.Unmarshal(*valHideSuperscriptCap, &valueForHideSuperscript)
			if err != nil {
				return err
			}
			this.HideSuperscript = valueForHideSuperscript
		}
	}

	return nil
}
