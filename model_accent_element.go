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

// Specifies the accent function, consisting of a base and a combining diacritical mark
type IAccentElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// The argument to which the accent was applied
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Accent Character
	GetCharacter() string
	SetCharacter(newValue string)
}

type AccentElement struct {

	// Element type
	Type_ string `json:"Type"`

	// The argument to which the accent was applied
	Base IMathElement `json:"Base,omitempty"`

	// Accent Character
	Character string `json:"Character,omitempty"`
}

func NewAccentElement() *AccentElement {
	instance := new(AccentElement)
	instance.Type_ = "Accent"
	return instance
}

func (this *AccentElement) GetType() string {
	return this.Type_
}

func (this *AccentElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AccentElement) GetBase() IMathElement {
	return this.Base
}

func (this *AccentElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *AccentElement) GetCharacter() string {
	return this.Character
}

func (this *AccentElement) SetCharacter(newValue string) {
	this.Character = newValue
}

func (this *AccentElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Accent"
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
	
	if valCharacter, ok := objMap["character"]; ok {
		if valCharacter != nil {
			var valueForCharacter string
			err = json.Unmarshal(*valCharacter, &valueForCharacter)
			if err != nil {
				return err
			}
			this.Character = valueForCharacter
		}
	}
	if valCharacterCap, ok := objMap["Character"]; ok {
		if valCharacterCap != nil {
			var valueForCharacter string
			err = json.Unmarshal(*valCharacterCap, &valueForCharacter)
			if err != nil {
				return err
			}
			this.Character = valueForCharacter
		}
	}

	return nil
}
