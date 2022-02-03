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
	getType() string
	setType(newValue string)

	// Hide the placeholders for empty matrix elements
	getHidePlaceholders() bool
	setHidePlaceholders(newValue bool)

	// Specifies the vertical justification respect to surrounding text. 
	getBaseJustification() string
	setBaseJustification(newValue string)

	// Minimum column width in twips (1/20th of a point)
	getMinColumnWidth() int32
	setMinColumnWidth(newValue int32)

	// The type of horizontal spacing between columns of a matrix.
	getColumnGapRule() string
	setColumnGapRule(newValue string)

	// The value of horizontal spacing between columns of a matrix
	getColumnGap() int32
	setColumnGap(newValue int32)

	// The type of vertical spacing between rows of a matrix
	getRowGapRule() string
	setRowGapRule(newValue string)

	// The value of vertical spacing between rows of a matrix;             
	getRowGap() int32
	setRowGap(newValue int32)

	// Matrix items
	getItems() [][]IMathElement
	setItems(newValue [][]IMathElement)
}

type MatrixElement struct {

	// Element type
	Type_ string `json:"Type"`

	// Hide the placeholders for empty matrix elements
	HidePlaceholders bool `json:"HidePlaceholders"`

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
	instance.BaseJustification = ""
	instance.ColumnGapRule = ""
	instance.RowGapRule = ""
	return instance
}

func (this *MatrixElement) getType() string {
	return this.Type_
}

func (this *MatrixElement) setType(newValue string) {
	this.Type_ = newValue
}
func (this *MatrixElement) getHidePlaceholders() bool {
	return this.HidePlaceholders
}

func (this *MatrixElement) setHidePlaceholders(newValue bool) {
	this.HidePlaceholders = newValue
}
func (this *MatrixElement) getBaseJustification() string {
	return this.BaseJustification
}

func (this *MatrixElement) setBaseJustification(newValue string) {
	this.BaseJustification = newValue
}
func (this *MatrixElement) getMinColumnWidth() int32 {
	return this.MinColumnWidth
}

func (this *MatrixElement) setMinColumnWidth(newValue int32) {
	this.MinColumnWidth = newValue
}
func (this *MatrixElement) getColumnGapRule() string {
	return this.ColumnGapRule
}

func (this *MatrixElement) setColumnGapRule(newValue string) {
	this.ColumnGapRule = newValue
}
func (this *MatrixElement) getColumnGap() int32 {
	return this.ColumnGap
}

func (this *MatrixElement) setColumnGap(newValue int32) {
	this.ColumnGap = newValue
}
func (this *MatrixElement) getRowGapRule() string {
	return this.RowGapRule
}

func (this *MatrixElement) setRowGapRule(newValue string) {
	this.RowGapRule = newValue
}
func (this *MatrixElement) getRowGap() int32 {
	return this.RowGap
}

func (this *MatrixElement) setRowGap(newValue int32) {
	this.RowGap = newValue
}
func (this *MatrixElement) getItems() [][]IMathElement {
	return this.Items
}

func (this *MatrixElement) setItems(newValue [][]IMathElement) {
	this.Items = newValue
}

func (this *MatrixElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Matrix"
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
	
	if valHidePlaceholders, ok := objMap["hidePlaceholders"]; ok {
		if valHidePlaceholders != nil {
			var valueForHidePlaceholders bool
			err = json.Unmarshal(*valHidePlaceholders, &valueForHidePlaceholders)
			if err != nil {
				return err
			}
			this.HidePlaceholders = valueForHidePlaceholders
		}
	}
	if valHidePlaceholdersCap, ok := objMap["HidePlaceholders"]; ok {
		if valHidePlaceholdersCap != nil {
			var valueForHidePlaceholders bool
			err = json.Unmarshal(*valHidePlaceholdersCap, &valueForHidePlaceholders)
			if err != nil {
				return err
			}
			this.HidePlaceholders = valueForHidePlaceholders
		}
	}
	this.BaseJustification = ""
	if valBaseJustification, ok := objMap["baseJustification"]; ok {
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
	if valBaseJustificationCap, ok := objMap["BaseJustification"]; ok {
		if valBaseJustificationCap != nil {
			var valueForBaseJustification string
			err = json.Unmarshal(*valBaseJustificationCap, &valueForBaseJustification)
			if err != nil {
				var valueForBaseJustificationInt int32
				err = json.Unmarshal(*valBaseJustificationCap, &valueForBaseJustificationInt)
				if err != nil {
					return err
				}
				this.BaseJustification = string(valueForBaseJustificationInt)
			} else {
				this.BaseJustification = valueForBaseJustification
			}
		}
	}
	
	if valMinColumnWidth, ok := objMap["minColumnWidth"]; ok {
		if valMinColumnWidth != nil {
			var valueForMinColumnWidth int32
			err = json.Unmarshal(*valMinColumnWidth, &valueForMinColumnWidth)
			if err != nil {
				return err
			}
			this.MinColumnWidth = valueForMinColumnWidth
		}
	}
	if valMinColumnWidthCap, ok := objMap["MinColumnWidth"]; ok {
		if valMinColumnWidthCap != nil {
			var valueForMinColumnWidth int32
			err = json.Unmarshal(*valMinColumnWidthCap, &valueForMinColumnWidth)
			if err != nil {
				return err
			}
			this.MinColumnWidth = valueForMinColumnWidth
		}
	}
	this.ColumnGapRule = ""
	if valColumnGapRule, ok := objMap["columnGapRule"]; ok {
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
	if valColumnGapRuleCap, ok := objMap["ColumnGapRule"]; ok {
		if valColumnGapRuleCap != nil {
			var valueForColumnGapRule string
			err = json.Unmarshal(*valColumnGapRuleCap, &valueForColumnGapRule)
			if err != nil {
				var valueForColumnGapRuleInt int32
				err = json.Unmarshal(*valColumnGapRuleCap, &valueForColumnGapRuleInt)
				if err != nil {
					return err
				}
				this.ColumnGapRule = string(valueForColumnGapRuleInt)
			} else {
				this.ColumnGapRule = valueForColumnGapRule
			}
		}
	}
	
	if valColumnGap, ok := objMap["columnGap"]; ok {
		if valColumnGap != nil {
			var valueForColumnGap int32
			err = json.Unmarshal(*valColumnGap, &valueForColumnGap)
			if err != nil {
				return err
			}
			this.ColumnGap = valueForColumnGap
		}
	}
	if valColumnGapCap, ok := objMap["ColumnGap"]; ok {
		if valColumnGapCap != nil {
			var valueForColumnGap int32
			err = json.Unmarshal(*valColumnGapCap, &valueForColumnGap)
			if err != nil {
				return err
			}
			this.ColumnGap = valueForColumnGap
		}
	}
	this.RowGapRule = ""
	if valRowGapRule, ok := objMap["rowGapRule"]; ok {
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
	if valRowGapRuleCap, ok := objMap["RowGapRule"]; ok {
		if valRowGapRuleCap != nil {
			var valueForRowGapRule string
			err = json.Unmarshal(*valRowGapRuleCap, &valueForRowGapRule)
			if err != nil {
				var valueForRowGapRuleInt int32
				err = json.Unmarshal(*valRowGapRuleCap, &valueForRowGapRuleInt)
				if err != nil {
					return err
				}
				this.RowGapRule = string(valueForRowGapRuleInt)
			} else {
				this.RowGapRule = valueForRowGapRule
			}
		}
	}
	
	if valRowGap, ok := objMap["rowGap"]; ok {
		if valRowGap != nil {
			var valueForRowGap int32
			err = json.Unmarshal(*valRowGap, &valueForRowGap)
			if err != nil {
				return err
			}
			this.RowGap = valueForRowGap
		}
	}
	if valRowGapCap, ok := objMap["RowGap"]; ok {
		if valRowGapCap != nil {
			var valueForRowGap int32
			err = json.Unmarshal(*valRowGapCap, &valueForRowGap)
			if err != nil {
				return err
			}
			this.RowGap = valueForRowGap
		}
	}
	
	if valItems, ok := objMap["items"]; ok {
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
	if valItemsCap, ok := objMap["Items"]; ok {
		if valItemsCap != nil {
			var valueForItems []json.RawMessage
			err = json.Unmarshal(*valItemsCap, &valueForItems)
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
