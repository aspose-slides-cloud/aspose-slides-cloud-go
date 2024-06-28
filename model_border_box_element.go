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

// Rectangular or some other border around the MathElement. 
type IBorderBoxElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Base
	GetBase() IMathElement
	SetBase(newValue IMathElement)

	// Hide Top Edge
	GetHideTop() *bool
	SetHideTop(newValue *bool)

	// Hide Bottom Edge
	GetHideBottom() *bool
	SetHideBottom(newValue *bool)

	// Hide Left Edge
	GetHideLeft() *bool
	SetHideLeft(newValue *bool)

	// Hide Right Edge
	GetHideRight() *bool
	SetHideRight(newValue *bool)

	// Strikethrough Horizontal
	GetStrikethroughHorizontal() *bool
	SetStrikethroughHorizontal(newValue *bool)

	// Strikethrough Vertical
	GetStrikethroughVertical() *bool
	SetStrikethroughVertical(newValue *bool)

	// Strikethrough Bottom-Left to Top-Right
	GetStrikethroughBottomLeftToTopRight() *bool
	SetStrikethroughBottomLeftToTopRight(newValue *bool)

	// Strikethrough Top-Left to Bottom-Right.
	GetStrikethroughTopLeftToBottomRight() *bool
	SetStrikethroughTopLeftToBottomRight(newValue *bool)
}

type BorderBoxElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Base
	Base IMathElement `json:"Base,omitempty"`

	// Hide Top Edge
	HideTop *bool `json:"HideTop"`

	// Hide Bottom Edge
	HideBottom *bool `json:"HideBottom"`

	// Hide Left Edge
	HideLeft *bool `json:"HideLeft"`

	// Hide Right Edge
	HideRight *bool `json:"HideRight"`

	// Strikethrough Horizontal
	StrikethroughHorizontal *bool `json:"StrikethroughHorizontal"`

	// Strikethrough Vertical
	StrikethroughVertical *bool `json:"StrikethroughVertical"`

	// Strikethrough Bottom-Left to Top-Right
	StrikethroughBottomLeftToTopRight *bool `json:"StrikethroughBottomLeftToTopRight"`

	// Strikethrough Top-Left to Bottom-Right.
	StrikethroughTopLeftToBottomRight *bool `json:"StrikethroughTopLeftToBottomRight"`
}

func NewBorderBoxElement() *BorderBoxElement {
	instance := new(BorderBoxElement)
	instance.Type_ = "BorderBox"
	return instance
}

func (this *BorderBoxElement) GetType() string {
	return this.Type_
}

func (this *BorderBoxElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *BorderBoxElement) GetBase() IMathElement {
	return this.Base
}

func (this *BorderBoxElement) SetBase(newValue IMathElement) {
	this.Base = newValue
}
func (this *BorderBoxElement) GetHideTop() *bool {
	return this.HideTop
}

func (this *BorderBoxElement) SetHideTop(newValue *bool) {
	this.HideTop = newValue
}
func (this *BorderBoxElement) GetHideBottom() *bool {
	return this.HideBottom
}

func (this *BorderBoxElement) SetHideBottom(newValue *bool) {
	this.HideBottom = newValue
}
func (this *BorderBoxElement) GetHideLeft() *bool {
	return this.HideLeft
}

func (this *BorderBoxElement) SetHideLeft(newValue *bool) {
	this.HideLeft = newValue
}
func (this *BorderBoxElement) GetHideRight() *bool {
	return this.HideRight
}

func (this *BorderBoxElement) SetHideRight(newValue *bool) {
	this.HideRight = newValue
}
func (this *BorderBoxElement) GetStrikethroughHorizontal() *bool {
	return this.StrikethroughHorizontal
}

func (this *BorderBoxElement) SetStrikethroughHorizontal(newValue *bool) {
	this.StrikethroughHorizontal = newValue
}
func (this *BorderBoxElement) GetStrikethroughVertical() *bool {
	return this.StrikethroughVertical
}

func (this *BorderBoxElement) SetStrikethroughVertical(newValue *bool) {
	this.StrikethroughVertical = newValue
}
func (this *BorderBoxElement) GetStrikethroughBottomLeftToTopRight() *bool {
	return this.StrikethroughBottomLeftToTopRight
}

func (this *BorderBoxElement) SetStrikethroughBottomLeftToTopRight(newValue *bool) {
	this.StrikethroughBottomLeftToTopRight = newValue
}
func (this *BorderBoxElement) GetStrikethroughTopLeftToBottomRight() *bool {
	return this.StrikethroughTopLeftToBottomRight
}

func (this *BorderBoxElement) SetStrikethroughTopLeftToBottomRight(newValue *bool) {
	this.StrikethroughTopLeftToBottomRight = newValue
}

func (this *BorderBoxElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "BorderBox"
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
	
	if valHideTop, ok := objMap["hideTop"]; ok {
		if valHideTop != nil {
			var valueForHideTop *bool
			err = json.Unmarshal(*valHideTop, &valueForHideTop)
			if err != nil {
				return err
			}
			this.HideTop = valueForHideTop
		}
	}
	if valHideTopCap, ok := objMap["HideTop"]; ok {
		if valHideTopCap != nil {
			var valueForHideTop *bool
			err = json.Unmarshal(*valHideTopCap, &valueForHideTop)
			if err != nil {
				return err
			}
			this.HideTop = valueForHideTop
		}
	}
	
	if valHideBottom, ok := objMap["hideBottom"]; ok {
		if valHideBottom != nil {
			var valueForHideBottom *bool
			err = json.Unmarshal(*valHideBottom, &valueForHideBottom)
			if err != nil {
				return err
			}
			this.HideBottom = valueForHideBottom
		}
	}
	if valHideBottomCap, ok := objMap["HideBottom"]; ok {
		if valHideBottomCap != nil {
			var valueForHideBottom *bool
			err = json.Unmarshal(*valHideBottomCap, &valueForHideBottom)
			if err != nil {
				return err
			}
			this.HideBottom = valueForHideBottom
		}
	}
	
	if valHideLeft, ok := objMap["hideLeft"]; ok {
		if valHideLeft != nil {
			var valueForHideLeft *bool
			err = json.Unmarshal(*valHideLeft, &valueForHideLeft)
			if err != nil {
				return err
			}
			this.HideLeft = valueForHideLeft
		}
	}
	if valHideLeftCap, ok := objMap["HideLeft"]; ok {
		if valHideLeftCap != nil {
			var valueForHideLeft *bool
			err = json.Unmarshal(*valHideLeftCap, &valueForHideLeft)
			if err != nil {
				return err
			}
			this.HideLeft = valueForHideLeft
		}
	}
	
	if valHideRight, ok := objMap["hideRight"]; ok {
		if valHideRight != nil {
			var valueForHideRight *bool
			err = json.Unmarshal(*valHideRight, &valueForHideRight)
			if err != nil {
				return err
			}
			this.HideRight = valueForHideRight
		}
	}
	if valHideRightCap, ok := objMap["HideRight"]; ok {
		if valHideRightCap != nil {
			var valueForHideRight *bool
			err = json.Unmarshal(*valHideRightCap, &valueForHideRight)
			if err != nil {
				return err
			}
			this.HideRight = valueForHideRight
		}
	}
	
	if valStrikethroughHorizontal, ok := objMap["strikethroughHorizontal"]; ok {
		if valStrikethroughHorizontal != nil {
			var valueForStrikethroughHorizontal *bool
			err = json.Unmarshal(*valStrikethroughHorizontal, &valueForStrikethroughHorizontal)
			if err != nil {
				return err
			}
			this.StrikethroughHorizontal = valueForStrikethroughHorizontal
		}
	}
	if valStrikethroughHorizontalCap, ok := objMap["StrikethroughHorizontal"]; ok {
		if valStrikethroughHorizontalCap != nil {
			var valueForStrikethroughHorizontal *bool
			err = json.Unmarshal(*valStrikethroughHorizontalCap, &valueForStrikethroughHorizontal)
			if err != nil {
				return err
			}
			this.StrikethroughHorizontal = valueForStrikethroughHorizontal
		}
	}
	
	if valStrikethroughVertical, ok := objMap["strikethroughVertical"]; ok {
		if valStrikethroughVertical != nil {
			var valueForStrikethroughVertical *bool
			err = json.Unmarshal(*valStrikethroughVertical, &valueForStrikethroughVertical)
			if err != nil {
				return err
			}
			this.StrikethroughVertical = valueForStrikethroughVertical
		}
	}
	if valStrikethroughVerticalCap, ok := objMap["StrikethroughVertical"]; ok {
		if valStrikethroughVerticalCap != nil {
			var valueForStrikethroughVertical *bool
			err = json.Unmarshal(*valStrikethroughVerticalCap, &valueForStrikethroughVertical)
			if err != nil {
				return err
			}
			this.StrikethroughVertical = valueForStrikethroughVertical
		}
	}
	
	if valStrikethroughBottomLeftToTopRight, ok := objMap["strikethroughBottomLeftToTopRight"]; ok {
		if valStrikethroughBottomLeftToTopRight != nil {
			var valueForStrikethroughBottomLeftToTopRight *bool
			err = json.Unmarshal(*valStrikethroughBottomLeftToTopRight, &valueForStrikethroughBottomLeftToTopRight)
			if err != nil {
				return err
			}
			this.StrikethroughBottomLeftToTopRight = valueForStrikethroughBottomLeftToTopRight
		}
	}
	if valStrikethroughBottomLeftToTopRightCap, ok := objMap["StrikethroughBottomLeftToTopRight"]; ok {
		if valStrikethroughBottomLeftToTopRightCap != nil {
			var valueForStrikethroughBottomLeftToTopRight *bool
			err = json.Unmarshal(*valStrikethroughBottomLeftToTopRightCap, &valueForStrikethroughBottomLeftToTopRight)
			if err != nil {
				return err
			}
			this.StrikethroughBottomLeftToTopRight = valueForStrikethroughBottomLeftToTopRight
		}
	}
	
	if valStrikethroughTopLeftToBottomRight, ok := objMap["strikethroughTopLeftToBottomRight"]; ok {
		if valStrikethroughTopLeftToBottomRight != nil {
			var valueForStrikethroughTopLeftToBottomRight *bool
			err = json.Unmarshal(*valStrikethroughTopLeftToBottomRight, &valueForStrikethroughTopLeftToBottomRight)
			if err != nil {
				return err
			}
			this.StrikethroughTopLeftToBottomRight = valueForStrikethroughTopLeftToBottomRight
		}
	}
	if valStrikethroughTopLeftToBottomRightCap, ok := objMap["StrikethroughTopLeftToBottomRight"]; ok {
		if valStrikethroughTopLeftToBottomRightCap != nil {
			var valueForStrikethroughTopLeftToBottomRight *bool
			err = json.Unmarshal(*valStrikethroughTopLeftToBottomRightCap, &valueForStrikethroughTopLeftToBottomRight)
			if err != nil {
				return err
			}
			this.StrikethroughTopLeftToBottomRight = valueForStrikethroughTopLeftToBottomRight
		}
	}

	return nil
}
