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

// Specifies a grouping symbol above or below an expression, usually to highlight the relationship between elements 
type IGroupingCharacterElement interface {

	// Element type
	getType() string
	setType(newValue string)

	// Base
	getBase() IMathElement
	setBase(newValue IMathElement)

	// Grouping character
	getCharacter() string
	setCharacter(newValue string)

	// Position of grouping character.
	getPosition() string
	setPosition(newValue string)

	// Vertical justification of group character.
	getVerticalJustification() string
	setVerticalJustification(newValue string)
}

type GroupingCharacterElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base
	Base IMathElement `json:"Base,omitempty"`

	// Grouping character
	Character string `json:"Character,omitempty"`

	// Position of grouping character.
	Position string `json:"Position,omitempty"`

	// Vertical justification of group character.
	VerticalJustification string `json:"VerticalJustification,omitempty"`
}

func NewGroupingCharacterElement() *GroupingCharacterElement {
	instance := new(GroupingCharacterElement)
	instance.Type_ = "GroupingCharacter"
	instance.Position = ""
	instance.VerticalJustification = ""
	return instance
}

func (this *GroupingCharacterElement) getType() string {
	return this.Type_
}

func (this *GroupingCharacterElement) setType(newValue string) {
	this.Type_ = newValue
}
func (this *GroupingCharacterElement) getBase() IMathElement {
	return this.Base
}

func (this *GroupingCharacterElement) setBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *GroupingCharacterElement) getCharacter() string {
	return this.Character
}

func (this *GroupingCharacterElement) setCharacter(newValue string) {
	this.Character = newValue
}
func (this *GroupingCharacterElement) getPosition() string {
	return this.Position
}

func (this *GroupingCharacterElement) setPosition(newValue string) {
	this.Position = newValue
}
func (this *GroupingCharacterElement) getVerticalJustification() string {
	return this.VerticalJustification
}

func (this *GroupingCharacterElement) setVerticalJustification(newValue string) {
	this.VerticalJustification = newValue
}

func (this *GroupingCharacterElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "GroupingCharacter"
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
	this.Position = ""
	if valPosition, ok := objMap["position"]; ok {
		if valPosition != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				var valueForPositionInt int32
				err = json.Unmarshal(*valPosition, &valueForPositionInt)
				if err != nil {
					return err
				}
				this.Position = string(valueForPositionInt)
			} else {
				this.Position = valueForPosition
			}
		}
	}
	if valPositionCap, ok := objMap["Position"]; ok {
		if valPositionCap != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPositionCap, &valueForPosition)
			if err != nil {
				var valueForPositionInt int32
				err = json.Unmarshal(*valPositionCap, &valueForPositionInt)
				if err != nil {
					return err
				}
				this.Position = string(valueForPositionInt)
			} else {
				this.Position = valueForPosition
			}
		}
	}
	this.VerticalJustification = ""
	if valVerticalJustification, ok := objMap["verticalJustification"]; ok {
		if valVerticalJustification != nil {
			var valueForVerticalJustification string
			err = json.Unmarshal(*valVerticalJustification, &valueForVerticalJustification)
			if err != nil {
				var valueForVerticalJustificationInt int32
				err = json.Unmarshal(*valVerticalJustification, &valueForVerticalJustificationInt)
				if err != nil {
					return err
				}
				this.VerticalJustification = string(valueForVerticalJustificationInt)
			} else {
				this.VerticalJustification = valueForVerticalJustification
			}
		}
	}
	if valVerticalJustificationCap, ok := objMap["VerticalJustification"]; ok {
		if valVerticalJustificationCap != nil {
			var valueForVerticalJustification string
			err = json.Unmarshal(*valVerticalJustificationCap, &valueForVerticalJustification)
			if err != nil {
				var valueForVerticalJustificationInt int32
				err = json.Unmarshal(*valVerticalJustificationCap, &valueForVerticalJustificationInt)
				if err != nil {
					return err
				}
				this.VerticalJustification = string(valueForVerticalJustificationInt)
			} else {
				this.VerticalJustification = valueForVerticalJustification
			}
		}
	}

	return nil
}
