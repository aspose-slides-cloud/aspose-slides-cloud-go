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

// Box element.
type IBoxElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Operator emulator
	GetOperatorEmulator() *bool
	SetOperatorEmulator(newValue *bool)

	// No break
	GetNoBreak() *bool
	SetNoBreak(newValue *bool)

	// Differential
	GetDifferential() *bool
	SetDifferential(newValue *bool)

	// Alignment point
	GetAlignmentPoint() *bool
	SetAlignmentPoint(newValue *bool)

	// Explicit break
	GetExplicitBreak() int32
	SetExplicitBreak(newValue int32)
}

type BoxElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base
	Base IMathElement `json:"Base,omitempty"`

	// Operator emulator
	OperatorEmulator *bool `json:"OperatorEmulator"`

	// No break
	NoBreak *bool `json:"NoBreak"`

	// Differential
	Differential *bool `json:"Differential"`

	// Alignment point
	AlignmentPoint *bool `json:"AlignmentPoint"`

	// Explicit break
	ExplicitBreak int32 `json:"ExplicitBreak,omitempty"`
}

func NewBoxElement() *BoxElement {
	instance := new(BoxElement)
	instance.Type_ = "Box"
	return instance
}

func (this *BoxElement) GetType() string {
	return this.Type_
}

func (this *BoxElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *BoxElement) GetBase() IMathElement {
	return this.Base
}

func (this *BoxElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *BoxElement) GetOperatorEmulator() *bool {
	return this.OperatorEmulator
}

func (this *BoxElement) SetOperatorEmulator(newValue *bool) {
	this.OperatorEmulator = newValue
}
func (this *BoxElement) GetNoBreak() *bool {
	return this.NoBreak
}

func (this *BoxElement) SetNoBreak(newValue *bool) {
	this.NoBreak = newValue
}
func (this *BoxElement) GetDifferential() *bool {
	return this.Differential
}

func (this *BoxElement) SetDifferential(newValue *bool) {
	this.Differential = newValue
}
func (this *BoxElement) GetAlignmentPoint() *bool {
	return this.AlignmentPoint
}

func (this *BoxElement) SetAlignmentPoint(newValue *bool) {
	this.AlignmentPoint = newValue
}
func (this *BoxElement) GetExplicitBreak() int32 {
	return this.ExplicitBreak
}

func (this *BoxElement) SetExplicitBreak(newValue int32) {
	this.ExplicitBreak = newValue
}

func (this *BoxElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Box"
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
	
	if valOperatorEmulator, ok := objMap["operatorEmulator"]; ok {
		if valOperatorEmulator != nil {
			var valueForOperatorEmulator *bool
			err = json.Unmarshal(*valOperatorEmulator, &valueForOperatorEmulator)
			if err != nil {
				return err
			}
			this.OperatorEmulator = valueForOperatorEmulator
		}
	}
	if valOperatorEmulatorCap, ok := objMap["OperatorEmulator"]; ok {
		if valOperatorEmulatorCap != nil {
			var valueForOperatorEmulator *bool
			err = json.Unmarshal(*valOperatorEmulatorCap, &valueForOperatorEmulator)
			if err != nil {
				return err
			}
			this.OperatorEmulator = valueForOperatorEmulator
		}
	}
	
	if valNoBreak, ok := objMap["noBreak"]; ok {
		if valNoBreak != nil {
			var valueForNoBreak *bool
			err = json.Unmarshal(*valNoBreak, &valueForNoBreak)
			if err != nil {
				return err
			}
			this.NoBreak = valueForNoBreak
		}
	}
	if valNoBreakCap, ok := objMap["NoBreak"]; ok {
		if valNoBreakCap != nil {
			var valueForNoBreak *bool
			err = json.Unmarshal(*valNoBreakCap, &valueForNoBreak)
			if err != nil {
				return err
			}
			this.NoBreak = valueForNoBreak
		}
	}
	
	if valDifferential, ok := objMap["differential"]; ok {
		if valDifferential != nil {
			var valueForDifferential *bool
			err = json.Unmarshal(*valDifferential, &valueForDifferential)
			if err != nil {
				return err
			}
			this.Differential = valueForDifferential
		}
	}
	if valDifferentialCap, ok := objMap["Differential"]; ok {
		if valDifferentialCap != nil {
			var valueForDifferential *bool
			err = json.Unmarshal(*valDifferentialCap, &valueForDifferential)
			if err != nil {
				return err
			}
			this.Differential = valueForDifferential
		}
	}
	
	if valAlignmentPoint, ok := objMap["alignmentPoint"]; ok {
		if valAlignmentPoint != nil {
			var valueForAlignmentPoint *bool
			err = json.Unmarshal(*valAlignmentPoint, &valueForAlignmentPoint)
			if err != nil {
				return err
			}
			this.AlignmentPoint = valueForAlignmentPoint
		}
	}
	if valAlignmentPointCap, ok := objMap["AlignmentPoint"]; ok {
		if valAlignmentPointCap != nil {
			var valueForAlignmentPoint *bool
			err = json.Unmarshal(*valAlignmentPointCap, &valueForAlignmentPoint)
			if err != nil {
				return err
			}
			this.AlignmentPoint = valueForAlignmentPoint
		}
	}
	
	if valExplicitBreak, ok := objMap["explicitBreak"]; ok {
		if valExplicitBreak != nil {
			var valueForExplicitBreak int32
			err = json.Unmarshal(*valExplicitBreak, &valueForExplicitBreak)
			if err != nil {
				return err
			}
			this.ExplicitBreak = valueForExplicitBreak
		}
	}
	if valExplicitBreakCap, ok := objMap["ExplicitBreak"]; ok {
		if valExplicitBreakCap != nil {
			var valueForExplicitBreak int32
			err = json.Unmarshal(*valExplicitBreakCap, &valueForExplicitBreak)
			if err != nil {
				return err
			}
			this.ExplicitBreak = valueForExplicitBreak
		}
	}

	return nil
}
