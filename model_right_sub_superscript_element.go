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

// Specifies the Sub-Superscript object
type IRightSubSuperscriptElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base argument
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Subscript
	GetSubscript() IMathElement
	SetSubscript(newValue IMathElement)

	// Superscript
	GetSuperscript() IMathElement
	SetSuperscript(newValue IMathElement)

	// Alignment of subscript/superscript.
	GetAlignScripts() *bool
	SetAlignScripts(newValue *bool)
}

type RightSubSuperscriptElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base argument
	Base IMathElement `json:"Base,omitempty"`

	// Subscript
	Subscript IMathElement `json:"Subscript,omitempty"`

	// Superscript
	Superscript IMathElement `json:"Superscript,omitempty"`

	// Alignment of subscript/superscript.
	AlignScripts *bool `json:"AlignScripts"`
}

func NewRightSubSuperscriptElement() *RightSubSuperscriptElement {
	instance := new(RightSubSuperscriptElement)
	instance.Type_ = "RightSubSuperscriptElement"
	return instance
}

func (this *RightSubSuperscriptElement) GetType() string {
	return this.Type_
}

func (this *RightSubSuperscriptElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *RightSubSuperscriptElement) GetBase() IMathElement {
	return this.Base
}

func (this *RightSubSuperscriptElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *RightSubSuperscriptElement) GetSubscript() IMathElement {
	return this.Subscript
}

func (this *RightSubSuperscriptElement) SetSubscript(newValue IMathElement) {
	this.Subscript = newValue
}
func (this *RightSubSuperscriptElement) GetSuperscript() IMathElement {
	return this.Superscript
}

func (this *RightSubSuperscriptElement) SetSuperscript(newValue IMathElement) {
	this.Superscript = newValue
}
func (this *RightSubSuperscriptElement) GetAlignScripts() *bool {
	return this.AlignScripts
}

func (this *RightSubSuperscriptElement) SetAlignScripts(newValue *bool) {
	this.AlignScripts = newValue
}

func (this *RightSubSuperscriptElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "RightSubSuperscriptElement"
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
	
	if valAlignScripts, ok := objMap["alignScripts"]; ok {
		if valAlignScripts != nil {
			var valueForAlignScripts *bool
			err = json.Unmarshal(*valAlignScripts, &valueForAlignScripts)
			if err != nil {
				return err
			}
			this.AlignScripts = valueForAlignScripts
		}
	}
	if valAlignScriptsCap, ok := objMap["AlignScripts"]; ok {
		if valAlignScriptsCap != nil {
			var valueForAlignScripts *bool
			err = json.Unmarshal(*valAlignScriptsCap, &valueForAlignScripts)
			if err != nil {
				return err
			}
			this.AlignScripts = valueForAlignScripts
		}
	}

	return nil
}
