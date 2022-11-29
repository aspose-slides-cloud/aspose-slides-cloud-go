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

// Delimiter element
type IDelimiterElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Arguments
	GetArguments() []IMathElement
	SetArguments(newValue []IMathElement)

	// Beginning character
	GetBeginningCharacter() string
	SetBeginningCharacter(newValue string)

	// Separator character
	GetSeparatorCharacter() string
	SetSeparatorCharacter(newValue string)

	// Ending character
	GetEndingCharacter() string
	SetEndingCharacter(newValue string)

	// Grow to match operand height
	GetGrowToMatchOperandHeight() bool
	SetGrowToMatchOperandHeight(newValue bool)

	// Delimiter shape
	GetDelimiterShape() string
	SetDelimiterShape(newValue string)
}

type DelimiterElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Arguments
	Arguments []IMathElement `json:"Arguments,omitempty"`

	// Beginning character
	BeginningCharacter string `json:"BeginningCharacter,omitempty"`

	// Separator character
	SeparatorCharacter string `json:"SeparatorCharacter,omitempty"`

	// Ending character
	EndingCharacter string `json:"EndingCharacter,omitempty"`

	// Grow to match operand height
	GrowToMatchOperandHeight bool `json:"GrowToMatchOperandHeight"`

	// Delimiter shape
	DelimiterShape string `json:"DelimiterShape,omitempty"`
}

func NewDelimiterElement() *DelimiterElement {
	instance := new(DelimiterElement)
	instance.Type_ = "Delimiter"
	return instance
}

func (this *DelimiterElement) GetType() string {
	return this.Type_
}

func (this *DelimiterElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *DelimiterElement) GetArguments() []IMathElement {
	return this.Arguments
}

func (this *DelimiterElement) SetArguments(newValue []IMathElement) {
	this.Arguments = newValue
}
func (this *DelimiterElement) GetBeginningCharacter() string {
	return this.BeginningCharacter
}

func (this *DelimiterElement) SetBeginningCharacter(newValue string) {
	this.BeginningCharacter = newValue
}
func (this *DelimiterElement) GetSeparatorCharacter() string {
	return this.SeparatorCharacter
}

func (this *DelimiterElement) SetSeparatorCharacter(newValue string) {
	this.SeparatorCharacter = newValue
}
func (this *DelimiterElement) GetEndingCharacter() string {
	return this.EndingCharacter
}

func (this *DelimiterElement) SetEndingCharacter(newValue string) {
	this.EndingCharacter = newValue
}
func (this *DelimiterElement) GetGrowToMatchOperandHeight() bool {
	return this.GrowToMatchOperandHeight
}

func (this *DelimiterElement) SetGrowToMatchOperandHeight(newValue bool) {
	this.GrowToMatchOperandHeight = newValue
}
func (this *DelimiterElement) GetDelimiterShape() string {
	return this.DelimiterShape
}

func (this *DelimiterElement) SetDelimiterShape(newValue string) {
	this.DelimiterShape = newValue
}

func (this *DelimiterElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Delimiter"
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
	
	if valArguments, ok := objMap["arguments"]; ok {
		if valArguments != nil {
			var valueForArguments []json.RawMessage
			err = json.Unmarshal(*valArguments, &valueForArguments)
			if err != nil {
				return err
			}
			valueForIArguments := make([]IMathElement, len(valueForArguments))
			for i, v := range valueForArguments {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIArguments[i] = vObject.(IMathElement)
				}
			}
			this.Arguments = valueForIArguments
		}
	}
	if valArgumentsCap, ok := objMap["Arguments"]; ok {
		if valArgumentsCap != nil {
			var valueForArguments []json.RawMessage
			err = json.Unmarshal(*valArgumentsCap, &valueForArguments)
			if err != nil {
				return err
			}
			valueForIArguments := make([]IMathElement, len(valueForArguments))
			for i, v := range valueForArguments {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIArguments[i] = vObject.(IMathElement)
				}
			}
			this.Arguments = valueForIArguments
		}
	}
	
	if valBeginningCharacter, ok := objMap["beginningCharacter"]; ok {
		if valBeginningCharacter != nil {
			var valueForBeginningCharacter string
			err = json.Unmarshal(*valBeginningCharacter, &valueForBeginningCharacter)
			if err != nil {
				return err
			}
			this.BeginningCharacter = valueForBeginningCharacter
		}
	}
	if valBeginningCharacterCap, ok := objMap["BeginningCharacter"]; ok {
		if valBeginningCharacterCap != nil {
			var valueForBeginningCharacter string
			err = json.Unmarshal(*valBeginningCharacterCap, &valueForBeginningCharacter)
			if err != nil {
				return err
			}
			this.BeginningCharacter = valueForBeginningCharacter
		}
	}
	
	if valSeparatorCharacter, ok := objMap["separatorCharacter"]; ok {
		if valSeparatorCharacter != nil {
			var valueForSeparatorCharacter string
			err = json.Unmarshal(*valSeparatorCharacter, &valueForSeparatorCharacter)
			if err != nil {
				return err
			}
			this.SeparatorCharacter = valueForSeparatorCharacter
		}
	}
	if valSeparatorCharacterCap, ok := objMap["SeparatorCharacter"]; ok {
		if valSeparatorCharacterCap != nil {
			var valueForSeparatorCharacter string
			err = json.Unmarshal(*valSeparatorCharacterCap, &valueForSeparatorCharacter)
			if err != nil {
				return err
			}
			this.SeparatorCharacter = valueForSeparatorCharacter
		}
	}
	
	if valEndingCharacter, ok := objMap["endingCharacter"]; ok {
		if valEndingCharacter != nil {
			var valueForEndingCharacter string
			err = json.Unmarshal(*valEndingCharacter, &valueForEndingCharacter)
			if err != nil {
				return err
			}
			this.EndingCharacter = valueForEndingCharacter
		}
	}
	if valEndingCharacterCap, ok := objMap["EndingCharacter"]; ok {
		if valEndingCharacterCap != nil {
			var valueForEndingCharacter string
			err = json.Unmarshal(*valEndingCharacterCap, &valueForEndingCharacter)
			if err != nil {
				return err
			}
			this.EndingCharacter = valueForEndingCharacter
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
	
	if valDelimiterShape, ok := objMap["delimiterShape"]; ok {
		if valDelimiterShape != nil {
			var valueForDelimiterShape string
			err = json.Unmarshal(*valDelimiterShape, &valueForDelimiterShape)
			if err != nil {
				var valueForDelimiterShapeInt int32
				err = json.Unmarshal(*valDelimiterShape, &valueForDelimiterShapeInt)
				if err != nil {
					return err
				}
				this.DelimiterShape = string(valueForDelimiterShapeInt)
			} else {
				this.DelimiterShape = valueForDelimiterShape
			}
		}
	}
	if valDelimiterShapeCap, ok := objMap["DelimiterShape"]; ok {
		if valDelimiterShapeCap != nil {
			var valueForDelimiterShape string
			err = json.Unmarshal(*valDelimiterShapeCap, &valueForDelimiterShape)
			if err != nil {
				var valueForDelimiterShapeInt int32
				err = json.Unmarshal(*valDelimiterShapeCap, &valueForDelimiterShapeInt)
				if err != nil {
					return err
				}
				this.DelimiterShape = string(valueForDelimiterShapeInt)
			} else {
				this.DelimiterShape = valueForDelimiterShape
			}
		}
	}

	return nil
}
