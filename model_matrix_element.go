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

// Specifies the Matrix object,
type IMatrixElement interface {

	// Element type
	GetType() string
	SetType(newValue string)

	// Hide the placeholders for empty matrix elements
	GetHidePlaceholders() *bool
	SetHidePlaceholders(newValue *bool)

	// Specifies the vertical justification respect to surrounding text. 
	GetBaseJustification() string
	SetBaseJustification(newValue string)

	// Minimum column width in twips (1/20th of a point)
	GetMinColumnWidth() int32
	SetMinColumnWidth(newValue int32)

	// The type of horizontal spacing between columns of a matrix.
	GetColumnGapRule() string
	SetColumnGapRule(newValue string)

	// The value of horizontal spacing between columns of a matrix
	GetColumnGap() int32
	SetColumnGap(newValue int32)

	// The type of vertical spacing between rows of a matrix
	GetRowGapRule() string
	SetRowGapRule(newValue string)

	// The value of vertical spacing between rows of a matrix;             
	GetRowGap() int32
	SetRowGap(newValue int32)

	// Matrix items
	GetItems() [][]IMathElement
	SetItems(newValue [][]IMathElement)
}

type MatrixElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Hide the placeholders for empty matrix elements
	HidePlaceholders *bool `json:"HidePlaceholders"`

	// Specifies the vertical justification respect to surrounding text. 
	BaseJustification string `json:"BaseJustification,omitempty"`

	// Minimum column width in twips (1/20th of a point)
	MinColumnWidth int32 `json:"MinColumnWidth,omitempty"`

	// The type of horizontal spacing between columns of a matrix.
	ColumnGapRule string `json:"ColumnGapRule,omitempty"`

	// The value of horizontal spacing between columns of a matrix
	ColumnGap int32 `json:"ColumnGap,omitempty"`

	// The type of vertical spacing between rows of a matrix
	RowGapRule string `json:"RowGapRule,omitempty"`

	// The value of vertical spacing between rows of a matrix;             
	RowGap int32 `json:"RowGap,omitempty"`

	// Matrix items
	Items [][]IMathElement `json:"Items,omitempty"`
}

func NewMatrixElement() *MatrixElement {
	instance := new(MatrixElement)
	instance.Type_ = "Matrix"
	return instance
}

func (this *MatrixElement) GetType() string {
	return this.Type_
}

func (this *MatrixElement) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *MatrixElement) GetHidePlaceholders() *bool {
	return this.HidePlaceholders
}

func (this *MatrixElement) SetHidePlaceholders(newValue *bool) {
	this.HidePlaceholders = newValue
}
func (this *MatrixElement) GetBaseJustification() string {
	return this.BaseJustification
}

func (this *MatrixElement) SetBaseJustification(newValue string) {
	this.BaseJustification = newValue
}
func (this *MatrixElement) GetMinColumnWidth() int32 {
	return this.MinColumnWidth
}

func (this *MatrixElement) SetMinColumnWidth(newValue int32) {
	this.MinColumnWidth = newValue
}
func (this *MatrixElement) GetColumnGapRule() string {
	return this.ColumnGapRule
}

func (this *MatrixElement) SetColumnGapRule(newValue string) {
	this.ColumnGapRule = newValue
}
func (this *MatrixElement) GetColumnGap() int32 {
	return this.ColumnGap
}

func (this *MatrixElement) SetColumnGap(newValue int32) {
	this.ColumnGap = newValue
}
func (this *MatrixElement) GetRowGapRule() string {
	return this.RowGapRule
}

func (this *MatrixElement) SetRowGapRule(newValue string) {
	this.RowGapRule = newValue
}
func (this *MatrixElement) GetRowGap() int32 {
	return this.RowGap
}

func (this *MatrixElement) SetRowGap(newValue int32) {
	this.RowGap = newValue
}
func (this *MatrixElement) GetItems() [][]IMathElement {
	return this.Items
}

func (this *MatrixElement) SetItems(newValue [][]IMathElement) {
	this.Items = newValue
}

func (this *MatrixElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Matrix"
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
	
	if valHidePlaceholders, ok := GetMapValue(objMap, "hidePlaceholders"); ok {
		if valHidePlaceholders != nil {
			var valueForHidePlaceholders *bool
			err = json.Unmarshal(*valHidePlaceholders, &valueForHidePlaceholders)
			if err != nil {
				return err
			}
			this.HidePlaceholders = valueForHidePlaceholders
		}
	}
	
	if valBaseJustification, ok := GetMapValue(objMap, "baseJustification"); ok {
		if valBaseJustification != nil {
			var valueForBaseJustification string
			err = json.Unmarshal(*valBaseJustification, &valueForBaseJustification)
			if err != nil {
				var valueForBaseJustificationInt int32
				err = json.Unmarshal(*valBaseJustification, &valueForBaseJustificationInt)
				if err != nil {
					return err
				}
				this.BaseJustification = string(valueForBaseJustificationInt)
			} else {
				this.BaseJustification = valueForBaseJustification
			}
		}
	}
	
	if valMinColumnWidth, ok := GetMapValue(objMap, "minColumnWidth"); ok {
		if valMinColumnWidth != nil {
			var valueForMinColumnWidth int32
			err = json.Unmarshal(*valMinColumnWidth, &valueForMinColumnWidth)
			if err != nil {
				return err
			}
			this.MinColumnWidth = valueForMinColumnWidth
		}
	}
	
	if valColumnGapRule, ok := GetMapValue(objMap, "columnGapRule"); ok {
		if valColumnGapRule != nil {
			var valueForColumnGapRule string
			err = json.Unmarshal(*valColumnGapRule, &valueForColumnGapRule)
			if err != nil {
				var valueForColumnGapRuleInt int32
				err = json.Unmarshal(*valColumnGapRule, &valueForColumnGapRuleInt)
				if err != nil {
					return err
				}
				this.ColumnGapRule = string(valueForColumnGapRuleInt)
			} else {
				this.ColumnGapRule = valueForColumnGapRule
			}
		}
	}
	
	if valColumnGap, ok := GetMapValue(objMap, "columnGap"); ok {
		if valColumnGap != nil {
			var valueForColumnGap int32
			err = json.Unmarshal(*valColumnGap, &valueForColumnGap)
			if err != nil {
				return err
			}
			this.ColumnGap = valueForColumnGap
		}
	}
	
	if valRowGapRule, ok := GetMapValue(objMap, "rowGapRule"); ok {
		if valRowGapRule != nil {
			var valueForRowGapRule string
			err = json.Unmarshal(*valRowGapRule, &valueForRowGapRule)
			if err != nil {
				var valueForRowGapRuleInt int32
				err = json.Unmarshal(*valRowGapRule, &valueForRowGapRuleInt)
				if err != nil {
					return err
				}
				this.RowGapRule = string(valueForRowGapRuleInt)
			} else {
				this.RowGapRule = valueForRowGapRule
			}
		}
	}
	
	if valRowGap, ok := GetMapValue(objMap, "rowGap"); ok {
		if valRowGap != nil {
			var valueForRowGap int32
			err = json.Unmarshal(*valRowGap, &valueForRowGap)
			if err != nil {
				return err
			}
			this.RowGap = valueForRowGap
		}
	}
	
	if valItems, ok := GetMapValue(objMap, "items"); ok {
		if valItems != nil {
			var valueForItems []json.RawMessage
			err = json.Unmarshal(*valItems, &valueForItems)
			if err != nil {
				return err
			}
			valueForIItems := make([][]IMathElement, len(valueForItems))
			for i, v := range valueForItems {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIItems[i] = vObject.([]IMathElement)
				}
			}
			this.Items = valueForIItems
		}
	}

	return nil
}
