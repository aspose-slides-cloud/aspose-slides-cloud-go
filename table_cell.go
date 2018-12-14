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
	getTextAnchorType() TextAnchorType
	setTextAnchorType(newValue TextAnchorType)

	// The type of vertical text.
	getTextVerticalType() TextVerticalType
	setTextVerticalType(newValue TextVerticalType)

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
	RowSpan int32 `json:"RowSpan"`

	// The number of columns spanned by a merged cell.
	ColSpan int32 `json:"ColSpan"`

	// The top margin of the cell.
	MarginTop float64 `json:"MarginTop"`

	// The right margin of the cell.
	MarginRight float64 `json:"MarginRight"`

	// The left margin of the cell.
	MarginLeft float64 `json:"MarginLeft"`

	// The bottom margin of the cell.
	MarginBottom float64 `json:"MarginBottom"`

	// Text anchor type.
	TextAnchorType TextAnchorType `json:"TextAnchorType"`

	// The type of vertical text.
	TextVerticalType TextVerticalType `json:"TextVerticalType"`

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

func (this TableCell) getText() string {
	return this.Text
}

func (this TableCell) setText(newValue string) {
	this.Text = newValue
}
func (this TableCell) getRowSpan() int32 {
	return this.RowSpan
}

func (this TableCell) setRowSpan(newValue int32) {
	this.RowSpan = newValue
}
func (this TableCell) getColSpan() int32 {
	return this.ColSpan
}

func (this TableCell) setColSpan(newValue int32) {
	this.ColSpan = newValue
}
func (this TableCell) getMarginTop() float64 {
	return this.MarginTop
}

func (this TableCell) setMarginTop(newValue float64) {
	this.MarginTop = newValue
}
func (this TableCell) getMarginRight() float64 {
	return this.MarginRight
}

func (this TableCell) setMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this TableCell) getMarginLeft() float64 {
	return this.MarginLeft
}

func (this TableCell) setMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this TableCell) getMarginBottom() float64 {
	return this.MarginBottom
}

func (this TableCell) setMarginBottom(newValue float64) {
	this.MarginBottom = newValue
}
func (this TableCell) getTextAnchorType() TextAnchorType {
	return this.TextAnchorType
}

func (this TableCell) setTextAnchorType(newValue TextAnchorType) {
	this.TextAnchorType = newValue
}
func (this TableCell) getTextVerticalType() TextVerticalType {
	return this.TextVerticalType
}

func (this TableCell) setTextVerticalType(newValue TextVerticalType) {
	this.TextVerticalType = newValue
}
func (this TableCell) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this TableCell) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this TableCell) getBorderTop() ILineFormat {
	return this.BorderTop
}

func (this TableCell) setBorderTop(newValue ILineFormat) {
	this.BorderTop = newValue
}
func (this TableCell) getBorderRight() ILineFormat {
	return this.BorderRight
}

func (this TableCell) setBorderRight(newValue ILineFormat) {
	this.BorderRight = newValue
}
func (this TableCell) getBorderLeft() ILineFormat {
	return this.BorderLeft
}

func (this TableCell) setBorderLeft(newValue ILineFormat) {
	this.BorderLeft = newValue
}
func (this TableCell) getBorderBottom() ILineFormat {
	return this.BorderBottom
}

func (this TableCell) setBorderBottom(newValue ILineFormat) {
	this.BorderBottom = newValue
}
func (this TableCell) getBorderDiagonalUp() ILineFormat {
	return this.BorderDiagonalUp
}

func (this TableCell) setBorderDiagonalUp(newValue ILineFormat) {
	this.BorderDiagonalUp = newValue
}
func (this TableCell) getBorderDiagonalDown() ILineFormat {
	return this.BorderDiagonalDown
}

func (this TableCell) setBorderDiagonalDown(newValue ILineFormat) {
	this.BorderDiagonalDown = newValue
}

func (this *TableCell) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valText, ok := objMap["Text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}

	if valRowSpan, ok := objMap["RowSpan"]; ok {
		if valRowSpan != nil {
			var valueForRowSpan int32
			err = json.Unmarshal(*valRowSpan, &valueForRowSpan)
			if err != nil {
				return err
			}
			this.RowSpan = valueForRowSpan
		}
	}

	if valColSpan, ok := objMap["ColSpan"]; ok {
		if valColSpan != nil {
			var valueForColSpan int32
			err = json.Unmarshal(*valColSpan, &valueForColSpan)
			if err != nil {
				return err
			}
			this.ColSpan = valueForColSpan
		}
	}

	if valMarginTop, ok := objMap["MarginTop"]; ok {
		if valMarginTop != nil {
			var valueForMarginTop float64
			err = json.Unmarshal(*valMarginTop, &valueForMarginTop)
			if err != nil {
				return err
			}
			this.MarginTop = valueForMarginTop
		}
	}

	if valMarginRight, ok := objMap["MarginRight"]; ok {
		if valMarginRight != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRight, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}

	if valMarginLeft, ok := objMap["MarginLeft"]; ok {
		if valMarginLeft != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeft, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}

	if valMarginBottom, ok := objMap["MarginBottom"]; ok {
		if valMarginBottom != nil {
			var valueForMarginBottom float64
			err = json.Unmarshal(*valMarginBottom, &valueForMarginBottom)
			if err != nil {
				return err
			}
			this.MarginBottom = valueForMarginBottom
		}
	}

	if valTextAnchorType, ok := objMap["TextAnchorType"]; ok {
		if valTextAnchorType != nil {
			var valueForTextAnchorType TextAnchorType
			err = json.Unmarshal(*valTextAnchorType, &valueForTextAnchorType)
			if err != nil {
				return err
			}
			this.TextAnchorType = valueForTextAnchorType
		}
	}

	if valTextVerticalType, ok := objMap["TextVerticalType"]; ok {
		if valTextVerticalType != nil {
			var valueForTextVerticalType TextVerticalType
			err = json.Unmarshal(*valTextVerticalType, &valueForTextVerticalType)
			if err != nil {
				return err
			}
			this.TextVerticalType = valueForTextVerticalType
		}
	}

	if valFillFormat, ok := objMap["FillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}

	if valBorderTop, ok := objMap["BorderTop"]; ok {
		if valBorderTop != nil {
			var valueForBorderTop LineFormat
			err = json.Unmarshal(*valBorderTop, &valueForBorderTop)
			if err != nil {
				return err
			}
			this.BorderTop = valueForBorderTop
		}
	}

	if valBorderRight, ok := objMap["BorderRight"]; ok {
		if valBorderRight != nil {
			var valueForBorderRight LineFormat
			err = json.Unmarshal(*valBorderRight, &valueForBorderRight)
			if err != nil {
				return err
			}
			this.BorderRight = valueForBorderRight
		}
	}

	if valBorderLeft, ok := objMap["BorderLeft"]; ok {
		if valBorderLeft != nil {
			var valueForBorderLeft LineFormat
			err = json.Unmarshal(*valBorderLeft, &valueForBorderLeft)
			if err != nil {
				return err
			}
			this.BorderLeft = valueForBorderLeft
		}
	}

	if valBorderBottom, ok := objMap["BorderBottom"]; ok {
		if valBorderBottom != nil {
			var valueForBorderBottom LineFormat
			err = json.Unmarshal(*valBorderBottom, &valueForBorderBottom)
			if err != nil {
				return err
			}
			this.BorderBottom = valueForBorderBottom
		}
	}

	if valBorderDiagonalUp, ok := objMap["BorderDiagonalUp"]; ok {
		if valBorderDiagonalUp != nil {
			var valueForBorderDiagonalUp LineFormat
			err = json.Unmarshal(*valBorderDiagonalUp, &valueForBorderDiagonalUp)
			if err != nil {
				return err
			}
			this.BorderDiagonalUp = valueForBorderDiagonalUp
		}
	}

	if valBorderDiagonalDown, ok := objMap["BorderDiagonalDown"]; ok {
		if valBorderDiagonalDown != nil {
			var valueForBorderDiagonalDown LineFormat
			err = json.Unmarshal(*valBorderDiagonalDown, &valueForBorderDiagonalDown)
			if err != nil {
				return err
			}
			this.BorderDiagonalDown = valueForBorderDiagonalDown
		}
	}

    return nil
}
