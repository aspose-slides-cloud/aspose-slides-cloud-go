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
	GetText() string
	SetText(newValue string)

	// The number of rows spanned by a merged cell.
	GetRowSpan() int32
	SetRowSpan(newValue int32)

	// The number of columns spanned by a merged cell.
	GetColSpan() int32
	SetColSpan(newValue int32)

	// The top margin of the cell.
	GetMarginTop() float64
	SetMarginTop(newValue float64)

	// The right margin of the cell.
	GetMarginRight() float64
	SetMarginRight(newValue float64)

	// The left margin of the cell.
	GetMarginLeft() float64
	SetMarginLeft(newValue float64)

	// The bottom margin of the cell.
	GetMarginBottom() float64
	SetMarginBottom(newValue float64)

	// Transparency.
	GetTransparency() float64
	SetTransparency(newValue float64)

	// Text anchor type.
	GetTextAnchorType() string
	SetTextAnchorType(newValue string)

	// The type of vertical text.
	GetTextVerticalType() string
	SetTextVerticalType(newValue string)

	// Fill properties set of the cell.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Line properties set for the top border of the cell.
	GetBorderTop() ILineFormat
	SetBorderTop(newValue ILineFormat)

	// Line properties set for the right border of the cell.
	GetBorderRight() ILineFormat
	SetBorderRight(newValue ILineFormat)

	// Line properties set for the left border of the cell.
	GetBorderLeft() ILineFormat
	SetBorderLeft(newValue ILineFormat)

	// Line properties set for the bottom border of the cell.
	GetBorderBottom() ILineFormat
	SetBorderBottom(newValue ILineFormat)

	// Line properties set for the diagonal up border of the cell.
	GetBorderDiagonalUp() ILineFormat
	SetBorderDiagonalUp(newValue ILineFormat)

	// Line properties set for the diagonal down border of the cell.
	GetBorderDiagonalDown() ILineFormat
	SetBorderDiagonalDown(newValue ILineFormat)

	// Cell column index
	GetColumnIndex() int32
	SetColumnIndex(newValue int32)

	// Cell row index
	GetRowIndex() int32
	SetRowIndex(newValue int32)

	// Returns TextFrame's formatting properties.
	GetTextFrameFormat() ITextFrameFormat
	SetTextFrameFormat(newValue ITextFrameFormat)

	// Get or sets list to paragraphs list
	GetParagraphs() IResourceUri
	SetParagraphs(newValue IResourceUri)
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

	// Transparency.
	Transparency float64 `json:"Transparency,omitempty"`

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

	// Cell column index
	ColumnIndex int32 `json:"ColumnIndex,omitempty"`

	// Cell row index
	RowIndex int32 `json:"RowIndex,omitempty"`

	// Returns TextFrame's formatting properties.
	TextFrameFormat ITextFrameFormat `json:"TextFrameFormat,omitempty"`

	// Get or sets list to paragraphs list
	Paragraphs IResourceUri `json:"Paragraphs,omitempty"`
}

func NewTableCell() *TableCell {
	instance := new(TableCell)
	return instance
}

func (this *TableCell) GetText() string {
	return this.Text
}

func (this *TableCell) SetText(newValue string) {
	this.Text = newValue
}
func (this *TableCell) GetRowSpan() int32 {
	return this.RowSpan
}

func (this *TableCell) SetRowSpan(newValue int32) {
	this.RowSpan = newValue
}
func (this *TableCell) GetColSpan() int32 {
	return this.ColSpan
}

func (this *TableCell) SetColSpan(newValue int32) {
	this.ColSpan = newValue
}
func (this *TableCell) GetMarginTop() float64 {
	return this.MarginTop
}

func (this *TableCell) SetMarginTop(newValue float64) {
	this.MarginTop = newValue
}
func (this *TableCell) GetMarginRight() float64 {
	return this.MarginRight
}

func (this *TableCell) SetMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this *TableCell) GetMarginLeft() float64 {
	return this.MarginLeft
}

func (this *TableCell) SetMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this *TableCell) GetMarginBottom() float64 {
	return this.MarginBottom
}

func (this *TableCell) SetMarginBottom(newValue float64) {
	this.MarginBottom = newValue
}
func (this *TableCell) GetTransparency() float64 {
	return this.Transparency
}

func (this *TableCell) SetTransparency(newValue float64) {
	this.Transparency = newValue
}
func (this *TableCell) GetTextAnchorType() string {
	return this.TextAnchorType
}

func (this *TableCell) SetTextAnchorType(newValue string) {
	this.TextAnchorType = newValue
}
func (this *TableCell) GetTextVerticalType() string {
	return this.TextVerticalType
}

func (this *TableCell) SetTextVerticalType(newValue string) {
	this.TextVerticalType = newValue
}
func (this *TableCell) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *TableCell) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *TableCell) GetBorderTop() ILineFormat {
	return this.BorderTop
}

func (this *TableCell) SetBorderTop(newValue ILineFormat) {
	this.BorderTop = newValue
}
func (this *TableCell) GetBorderRight() ILineFormat {
	return this.BorderRight
}

func (this *TableCell) SetBorderRight(newValue ILineFormat) {
	this.BorderRight = newValue
}
func (this *TableCell) GetBorderLeft() ILineFormat {
	return this.BorderLeft
}

func (this *TableCell) SetBorderLeft(newValue ILineFormat) {
	this.BorderLeft = newValue
}
func (this *TableCell) GetBorderBottom() ILineFormat {
	return this.BorderBottom
}

func (this *TableCell) SetBorderBottom(newValue ILineFormat) {
	this.BorderBottom = newValue
}
func (this *TableCell) GetBorderDiagonalUp() ILineFormat {
	return this.BorderDiagonalUp
}

func (this *TableCell) SetBorderDiagonalUp(newValue ILineFormat) {
	this.BorderDiagonalUp = newValue
}
func (this *TableCell) GetBorderDiagonalDown() ILineFormat {
	return this.BorderDiagonalDown
}

func (this *TableCell) SetBorderDiagonalDown(newValue ILineFormat) {
	this.BorderDiagonalDown = newValue
}
func (this *TableCell) GetColumnIndex() int32 {
	return this.ColumnIndex
}

func (this *TableCell) SetColumnIndex(newValue int32) {
	this.ColumnIndex = newValue
}
func (this *TableCell) GetRowIndex() int32 {
	return this.RowIndex
}

func (this *TableCell) SetRowIndex(newValue int32) {
	this.RowIndex = newValue
}
func (this *TableCell) GetTextFrameFormat() ITextFrameFormat {
	return this.TextFrameFormat
}

func (this *TableCell) SetTextFrameFormat(newValue ITextFrameFormat) {
	this.TextFrameFormat = newValue
}
func (this *TableCell) GetParagraphs() IResourceUri {
	return this.Paragraphs
}

func (this *TableCell) SetParagraphs(newValue IResourceUri) {
	this.Paragraphs = newValue
}

func (this *TableCell) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valText, ok := GetMapValue(objMap, "text"); ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	
	if valRowSpan, ok := GetMapValue(objMap, "rowSpan"); ok {
		if valRowSpan != nil {
			var valueForRowSpan int32
			err = json.Unmarshal(*valRowSpan, &valueForRowSpan)
			if err != nil {
				return err
			}
			this.RowSpan = valueForRowSpan
		}
	}
	
	if valColSpan, ok := GetMapValue(objMap, "colSpan"); ok {
		if valColSpan != nil {
			var valueForColSpan int32
			err = json.Unmarshal(*valColSpan, &valueForColSpan)
			if err != nil {
				return err
			}
			this.ColSpan = valueForColSpan
		}
	}
	
	if valMarginTop, ok := GetMapValue(objMap, "marginTop"); ok {
		if valMarginTop != nil {
			var valueForMarginTop float64
			err = json.Unmarshal(*valMarginTop, &valueForMarginTop)
			if err != nil {
				return err
			}
			this.MarginTop = valueForMarginTop
		}
	}
	
	if valMarginRight, ok := GetMapValue(objMap, "marginRight"); ok {
		if valMarginRight != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRight, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}
	
	if valMarginLeft, ok := GetMapValue(objMap, "marginLeft"); ok {
		if valMarginLeft != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeft, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}
	
	if valMarginBottom, ok := GetMapValue(objMap, "marginBottom"); ok {
		if valMarginBottom != nil {
			var valueForMarginBottom float64
			err = json.Unmarshal(*valMarginBottom, &valueForMarginBottom)
			if err != nil {
				return err
			}
			this.MarginBottom = valueForMarginBottom
		}
	}
	
	if valTransparency, ok := GetMapValue(objMap, "transparency"); ok {
		if valTransparency != nil {
			var valueForTransparency float64
			err = json.Unmarshal(*valTransparency, &valueForTransparency)
			if err != nil {
				return err
			}
			this.Transparency = valueForTransparency
		}
	}
	
	if valTextAnchorType, ok := GetMapValue(objMap, "textAnchorType"); ok {
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
	
	if valTextVerticalType, ok := GetMapValue(objMap, "textVerticalType"); ok {
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
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valBorderTop, ok := GetMapValue(objMap, "borderTop"); ok {
		if valBorderTop != nil {
			var valueForBorderTop LineFormat
			err = json.Unmarshal(*valBorderTop, &valueForBorderTop)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valBorderTop)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBorderTop, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.BorderTop = vInterfaceObject
			}
		}
	}
	
	if valBorderRight, ok := GetMapValue(objMap, "borderRight"); ok {
		if valBorderRight != nil {
			var valueForBorderRight LineFormat
			err = json.Unmarshal(*valBorderRight, &valueForBorderRight)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valBorderRight)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBorderRight, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.BorderRight = vInterfaceObject
			}
		}
	}
	
	if valBorderLeft, ok := GetMapValue(objMap, "borderLeft"); ok {
		if valBorderLeft != nil {
			var valueForBorderLeft LineFormat
			err = json.Unmarshal(*valBorderLeft, &valueForBorderLeft)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valBorderLeft)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBorderLeft, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.BorderLeft = vInterfaceObject
			}
		}
	}
	
	if valBorderBottom, ok := GetMapValue(objMap, "borderBottom"); ok {
		if valBorderBottom != nil {
			var valueForBorderBottom LineFormat
			err = json.Unmarshal(*valBorderBottom, &valueForBorderBottom)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valBorderBottom)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBorderBottom, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.BorderBottom = vInterfaceObject
			}
		}
	}
	
	if valBorderDiagonalUp, ok := GetMapValue(objMap, "borderDiagonalUp"); ok {
		if valBorderDiagonalUp != nil {
			var valueForBorderDiagonalUp LineFormat
			err = json.Unmarshal(*valBorderDiagonalUp, &valueForBorderDiagonalUp)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valBorderDiagonalUp)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBorderDiagonalUp, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.BorderDiagonalUp = vInterfaceObject
			}
		}
	}
	
	if valBorderDiagonalDown, ok := GetMapValue(objMap, "borderDiagonalDown"); ok {
		if valBorderDiagonalDown != nil {
			var valueForBorderDiagonalDown LineFormat
			err = json.Unmarshal(*valBorderDiagonalDown, &valueForBorderDiagonalDown)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valBorderDiagonalDown)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBorderDiagonalDown, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.BorderDiagonalDown = vInterfaceObject
			}
		}
	}
	
	if valColumnIndex, ok := GetMapValue(objMap, "columnIndex"); ok {
		if valColumnIndex != nil {
			var valueForColumnIndex int32
			err = json.Unmarshal(*valColumnIndex, &valueForColumnIndex)
			if err != nil {
				return err
			}
			this.ColumnIndex = valueForColumnIndex
		}
	}
	
	if valRowIndex, ok := GetMapValue(objMap, "rowIndex"); ok {
		if valRowIndex != nil {
			var valueForRowIndex int32
			err = json.Unmarshal(*valRowIndex, &valueForRowIndex)
			if err != nil {
				return err
			}
			this.RowIndex = valueForRowIndex
		}
	}
	
	if valTextFrameFormat, ok := GetMapValue(objMap, "textFrameFormat"); ok {
		if valTextFrameFormat != nil {
			var valueForTextFrameFormat TextFrameFormat
			err = json.Unmarshal(*valTextFrameFormat, &valueForTextFrameFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("TextFrameFormat", *valTextFrameFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valTextFrameFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ITextFrameFormat)
			if ok {
				this.TextFrameFormat = vInterfaceObject
			}
		}
	}
	
	if valParagraphs, ok := GetMapValue(objMap, "paragraphs"); ok {
		if valParagraphs != nil {
			var valueForParagraphs ResourceUri
			err = json.Unmarshal(*valParagraphs, &valueForParagraphs)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valParagraphs)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valParagraphs, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Paragraphs = vInterfaceObject
			}
		}
	}

	return nil
}
