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

// Represents one cell of table.
type ITableCell interface {

	// Cell text.
	getText() string
	setText(newValue string)

	// The number of rows spanned by a merged cell.
	getRowSpan() int32
	setRowSpan(newValue int32)

	// The number of columns spanned by a merged cell.
	getColSpan() int32
	setColSpan(newValue int32)

	// The top margin of the cell.
	getMarginTop() float64
	setMarginTop(newValue float64)

	// The right margin of the cell.
	getMarginRight() float64
	setMarginRight(newValue float64)

	// The left margin of the cell.
	getMarginLeft() float64
	setMarginLeft(newValue float64)

	// The bottom margin of the cell.
	getMarginBottom() float64
	setMarginBottom(newValue float64)

	// Text anchor type.
	getTextAnchorType() string
	setTextAnchorType(newValue string)

	// The type of vertical text.
	getTextVerticalType() string
	setTextVerticalType(newValue string)

	// Fill properties set of the cell.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Line properties set for the top border of the cell.
	getBorderTop() ILineFormat
	setBorderTop(newValue ILineFormat)

	// Line properties set for the right border of the cell.
	getBorderRight() ILineFormat
	setBorderRight(newValue ILineFormat)

	// Line properties set for the left border of the cell.
	getBorderLeft() ILineFormat
	setBorderLeft(newValue ILineFormat)

	// Line properties set for the bottom border of the cell.
	getBorderBottom() ILineFormat
	setBorderBottom(newValue ILineFormat)

	// Line properties set for the diagonal up border of the cell.
	getBorderDiagonalUp() ILineFormat
	setBorderDiagonalUp(newValue ILineFormat)

	// Line properties set for the diagonal down border of the cell.
	getBorderDiagonalDown() ILineFormat
	setBorderDiagonalDown(newValue ILineFormat)
}

type TableCell struct {

	// Cell text.
	Text string `json:"Text,omitempty"`

	// The number of rows spanned by a merged cell.
	RowSpan int32 `json:"RowSpan,omitempty"`

	// The number of columns spanned by a merged cell.
	ColSpan int32 `json:"ColSpan,omitempty"`

	// The top margin of the cell.
	MarginTop float64 `json:"MarginTop,omitempty"`

	// The right margin of the cell.
	MarginRight float64 `json:"MarginRight,omitempty"`

	// The left margin of the cell.
	MarginLeft float64 `json:"MarginLeft,omitempty"`

	// The bottom margin of the cell.
	MarginBottom float64 `json:"MarginBottom,omitempty"`

	// Text anchor type.
	TextAnchorType string `json:"TextAnchorType,omitempty"`

	// The type of vertical text.
	TextVerticalType string `json:"TextVerticalType,omitempty"`

	// Fill properties set of the cell.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Line properties set for the top border of the cell.
	BorderTop ILineFormat `json:"BorderTop,omitempty"`

	// Line properties set for the right border of the cell.
	BorderRight ILineFormat `json:"BorderRight,omitempty"`

	// Line properties set for the left border of the cell.
	BorderLeft ILineFormat `json:"BorderLeft,omitempty"`

	// Line properties set for the bottom border of the cell.
	BorderBottom ILineFormat `json:"BorderBottom,omitempty"`

	// Line properties set for the diagonal up border of the cell.
	BorderDiagonalUp ILineFormat `json:"BorderDiagonalUp,omitempty"`

	// Line properties set for the diagonal down border of the cell.
	BorderDiagonalDown ILineFormat `json:"BorderDiagonalDown,omitempty"`
}

func NewTableCell() *TableCell {
	instance := new(TableCell)
	instance.TextAnchorType = ""
	instance.TextVerticalType = ""
	return instance
}

func (this *TableCell) getText() string {
	return this.Text
}

func (this *TableCell) setText(newValue string) {
	this.Text = newValue
}
func (this *TableCell) getRowSpan() int32 {
	return this.RowSpan
}

func (this *TableCell) setRowSpan(newValue int32) {
	this.RowSpan = newValue
}
func (this *TableCell) getColSpan() int32 {
	return this.ColSpan
}

func (this *TableCell) setColSpan(newValue int32) {
	this.ColSpan = newValue
}
func (this *TableCell) getMarginTop() float64 {
	return this.MarginTop
}

func (this *TableCell) setMarginTop(newValue float64) {
	this.MarginTop = newValue
}
func (this *TableCell) getMarginRight() float64 {
	return this.MarginRight
}

func (this *TableCell) setMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this *TableCell) getMarginLeft() float64 {
	return this.MarginLeft
}

func (this *TableCell) setMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this *TableCell) getMarginBottom() float64 {
	return this.MarginBottom
}

func (this *TableCell) setMarginBottom(newValue float64) {
	this.MarginBottom = newValue
}
func (this *TableCell) getTextAnchorType() string {
	return this.TextAnchorType
}

func (this *TableCell) setTextAnchorType(newValue string) {
	this.TextAnchorType = newValue
}
func (this *TableCell) getTextVerticalType() string {
	return this.TextVerticalType
}

func (this *TableCell) setTextVerticalType(newValue string) {
	this.TextVerticalType = newValue
}
func (this *TableCell) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *TableCell) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *TableCell) getBorderTop() ILineFormat {
	return this.BorderTop
}

func (this *TableCell) setBorderTop(newValue ILineFormat) {
	this.BorderTop = newValue
}
func (this *TableCell) getBorderRight() ILineFormat {
	return this.BorderRight
}

func (this *TableCell) setBorderRight(newValue ILineFormat) {
	this.BorderRight = newValue
}
func (this *TableCell) getBorderLeft() ILineFormat {
	return this.BorderLeft
}

func (this *TableCell) setBorderLeft(newValue ILineFormat) {
	this.BorderLeft = newValue
}
func (this *TableCell) getBorderBottom() ILineFormat {
	return this.BorderBottom
}

func (this *TableCell) setBorderBottom(newValue ILineFormat) {
	this.BorderBottom = newValue
}
func (this *TableCell) getBorderDiagonalUp() ILineFormat {
	return this.BorderDiagonalUp
}

func (this *TableCell) setBorderDiagonalUp(newValue ILineFormat) {
	this.BorderDiagonalUp = newValue
}
func (this *TableCell) getBorderDiagonalDown() ILineFormat {
	return this.BorderDiagonalDown
}

func (this *TableCell) setBorderDiagonalDown(newValue ILineFormat) {
	this.BorderDiagonalDown = newValue
}

func (this *TableCell) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valText, ok := objMap["text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	if valTextCap, ok := objMap["Text"]; ok {
		if valTextCap != nil {
			var valueForText string
			err = json.Unmarshal(*valTextCap, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	
	if valRowSpan, ok := objMap["rowSpan"]; ok {
		if valRowSpan != nil {
			var valueForRowSpan int32
			err = json.Unmarshal(*valRowSpan, &valueForRowSpan)
			if err != nil {
				return err
			}
			this.RowSpan = valueForRowSpan
		}
	}
	if valRowSpanCap, ok := objMap["RowSpan"]; ok {
		if valRowSpanCap != nil {
			var valueForRowSpan int32
			err = json.Unmarshal(*valRowSpanCap, &valueForRowSpan)
			if err != nil {
				return err
			}
			this.RowSpan = valueForRowSpan
		}
	}
	
	if valColSpan, ok := objMap["colSpan"]; ok {
		if valColSpan != nil {
			var valueForColSpan int32
			err = json.Unmarshal(*valColSpan, &valueForColSpan)
			if err != nil {
				return err
			}
			this.ColSpan = valueForColSpan
		}
	}
	if valColSpanCap, ok := objMap["ColSpan"]; ok {
		if valColSpanCap != nil {
			var valueForColSpan int32
			err = json.Unmarshal(*valColSpanCap, &valueForColSpan)
			if err != nil {
				return err
			}
			this.ColSpan = valueForColSpan
		}
	}
	
	if valMarginTop, ok := objMap["marginTop"]; ok {
		if valMarginTop != nil {
			var valueForMarginTop float64
			err = json.Unmarshal(*valMarginTop, &valueForMarginTop)
			if err != nil {
				return err
			}
			this.MarginTop = valueForMarginTop
		}
	}
	if valMarginTopCap, ok := objMap["MarginTop"]; ok {
		if valMarginTopCap != nil {
			var valueForMarginTop float64
			err = json.Unmarshal(*valMarginTopCap, &valueForMarginTop)
			if err != nil {
				return err
			}
			this.MarginTop = valueForMarginTop
		}
	}
	
	if valMarginRight, ok := objMap["marginRight"]; ok {
		if valMarginRight != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRight, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}
	if valMarginRightCap, ok := objMap["MarginRight"]; ok {
		if valMarginRightCap != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRightCap, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}
	
	if valMarginLeft, ok := objMap["marginLeft"]; ok {
		if valMarginLeft != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeft, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}
	if valMarginLeftCap, ok := objMap["MarginLeft"]; ok {
		if valMarginLeftCap != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeftCap, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}
	
	if valMarginBottom, ok := objMap["marginBottom"]; ok {
		if valMarginBottom != nil {
			var valueForMarginBottom float64
			err = json.Unmarshal(*valMarginBottom, &valueForMarginBottom)
			if err != nil {
				return err
			}
			this.MarginBottom = valueForMarginBottom
		}
	}
	if valMarginBottomCap, ok := objMap["MarginBottom"]; ok {
		if valMarginBottomCap != nil {
			var valueForMarginBottom float64
			err = json.Unmarshal(*valMarginBottomCap, &valueForMarginBottom)
			if err != nil {
				return err
			}
			this.MarginBottom = valueForMarginBottom
		}
	}
	this.TextAnchorType = ""
	if valTextAnchorType, ok := objMap["textAnchorType"]; ok {
		if valTextAnchorType != nil {
			var valueForTextAnchorType string
			err = json.Unmarshal(*valTextAnchorType, &valueForTextAnchorType)
			if err != nil {
				var valueForTextAnchorTypeInt int32
				err = json.Unmarshal(*valTextAnchorType, &valueForTextAnchorTypeInt)
				if err != nil {
					return err
				}
				this.TextAnchorType = string(valueForTextAnchorTypeInt)
			} else {
				this.TextAnchorType = valueForTextAnchorType
			}
		}
	}
	if valTextAnchorTypeCap, ok := objMap["TextAnchorType"]; ok {
		if valTextAnchorTypeCap != nil {
			var valueForTextAnchorType string
			err = json.Unmarshal(*valTextAnchorTypeCap, &valueForTextAnchorType)
			if err != nil {
				var valueForTextAnchorTypeInt int32
				err = json.Unmarshal(*valTextAnchorTypeCap, &valueForTextAnchorTypeInt)
				if err != nil {
					return err
				}
				this.TextAnchorType = string(valueForTextAnchorTypeInt)
			} else {
				this.TextAnchorType = valueForTextAnchorType
			}
		}
	}
	this.TextVerticalType = ""
	if valTextVerticalType, ok := objMap["textVerticalType"]; ok {
		if valTextVerticalType != nil {
			var valueForTextVerticalType string
			err = json.Unmarshal(*valTextVerticalType, &valueForTextVerticalType)
			if err != nil {
				var valueForTextVerticalTypeInt int32
				err = json.Unmarshal(*valTextVerticalType, &valueForTextVerticalTypeInt)
				if err != nil {
					return err
				}
				this.TextVerticalType = string(valueForTextVerticalTypeInt)
			} else {
				this.TextVerticalType = valueForTextVerticalType
			}
		}
	}
	if valTextVerticalTypeCap, ok := objMap["TextVerticalType"]; ok {
		if valTextVerticalTypeCap != nil {
			var valueForTextVerticalType string
			err = json.Unmarshal(*valTextVerticalTypeCap, &valueForTextVerticalType)
			if err != nil {
				var valueForTextVerticalTypeInt int32
				err = json.Unmarshal(*valTextVerticalTypeCap, &valueForTextVerticalTypeInt)
				if err != nil {
					return err
				}
				this.TextVerticalType = string(valueForTextVerticalTypeInt)
			} else {
				this.TextVerticalType = valueForTextVerticalType
			}
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	
	if valBorderTop, ok := objMap["borderTop"]; ok {
		if valBorderTop != nil {
			var valueForBorderTop LineFormat
			err = json.Unmarshal(*valBorderTop, &valueForBorderTop)
			if err != nil {
				return err
			}
			this.BorderTop = &valueForBorderTop
		}
	}
	if valBorderTopCap, ok := objMap["BorderTop"]; ok {
		if valBorderTopCap != nil {
			var valueForBorderTop LineFormat
			err = json.Unmarshal(*valBorderTopCap, &valueForBorderTop)
			if err != nil {
				return err
			}
			this.BorderTop = &valueForBorderTop
		}
	}
	
	if valBorderRight, ok := objMap["borderRight"]; ok {
		if valBorderRight != nil {
			var valueForBorderRight LineFormat
			err = json.Unmarshal(*valBorderRight, &valueForBorderRight)
			if err != nil {
				return err
			}
			this.BorderRight = &valueForBorderRight
		}
	}
	if valBorderRightCap, ok := objMap["BorderRight"]; ok {
		if valBorderRightCap != nil {
			var valueForBorderRight LineFormat
			err = json.Unmarshal(*valBorderRightCap, &valueForBorderRight)
			if err != nil {
				return err
			}
			this.BorderRight = &valueForBorderRight
		}
	}
	
	if valBorderLeft, ok := objMap["borderLeft"]; ok {
		if valBorderLeft != nil {
			var valueForBorderLeft LineFormat
			err = json.Unmarshal(*valBorderLeft, &valueForBorderLeft)
			if err != nil {
				return err
			}
			this.BorderLeft = &valueForBorderLeft
		}
	}
	if valBorderLeftCap, ok := objMap["BorderLeft"]; ok {
		if valBorderLeftCap != nil {
			var valueForBorderLeft LineFormat
			err = json.Unmarshal(*valBorderLeftCap, &valueForBorderLeft)
			if err != nil {
				return err
			}
			this.BorderLeft = &valueForBorderLeft
		}
	}
	
	if valBorderBottom, ok := objMap["borderBottom"]; ok {
		if valBorderBottom != nil {
			var valueForBorderBottom LineFormat
			err = json.Unmarshal(*valBorderBottom, &valueForBorderBottom)
			if err != nil {
				return err
			}
			this.BorderBottom = &valueForBorderBottom
		}
	}
	if valBorderBottomCap, ok := objMap["BorderBottom"]; ok {
		if valBorderBottomCap != nil {
			var valueForBorderBottom LineFormat
			err = json.Unmarshal(*valBorderBottomCap, &valueForBorderBottom)
			if err != nil {
				return err
			}
			this.BorderBottom = &valueForBorderBottom
		}
	}
	
	if valBorderDiagonalUp, ok := objMap["borderDiagonalUp"]; ok {
		if valBorderDiagonalUp != nil {
			var valueForBorderDiagonalUp LineFormat
			err = json.Unmarshal(*valBorderDiagonalUp, &valueForBorderDiagonalUp)
			if err != nil {
				return err
			}
			this.BorderDiagonalUp = &valueForBorderDiagonalUp
		}
	}
	if valBorderDiagonalUpCap, ok := objMap["BorderDiagonalUp"]; ok {
		if valBorderDiagonalUpCap != nil {
			var valueForBorderDiagonalUp LineFormat
			err = json.Unmarshal(*valBorderDiagonalUpCap, &valueForBorderDiagonalUp)
			if err != nil {
				return err
			}
			this.BorderDiagonalUp = &valueForBorderDiagonalUp
		}
	}
	
	if valBorderDiagonalDown, ok := objMap["borderDiagonalDown"]; ok {
		if valBorderDiagonalDown != nil {
			var valueForBorderDiagonalDown LineFormat
			err = json.Unmarshal(*valBorderDiagonalDown, &valueForBorderDiagonalDown)
			if err != nil {
				return err
			}
			this.BorderDiagonalDown = &valueForBorderDiagonalDown
		}
	}
	if valBorderDiagonalDownCap, ok := objMap["BorderDiagonalDown"]; ok {
		if valBorderDiagonalDownCap != nil {
			var valueForBorderDiagonalDown LineFormat
			err = json.Unmarshal(*valBorderDiagonalDownCap, &valueForBorderDiagonalDown)
			if err != nil {
				return err
			}
			this.BorderDiagonalDown = &valueForBorderDiagonalDown
		}
	}

	return nil
}
