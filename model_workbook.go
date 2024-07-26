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

// Represents Excel spreadsheet data source.
type IWorkbook interface {

	// Data source type.
	GetType() string
	SetType(newValue string)

	// Worksheet index.
	GetWorksheetIndex() int32
	SetWorksheetIndex(newValue int32)

	// Column index of the first value.
	GetColumnIndex() int32
	SetColumnIndex(newValue int32)

	// Row index of the first value.
	GetRowIndex() int32
	SetRowIndex(newValue int32)
}

type Workbook struct {

	// Data source type.
	Type_ string `json:"Type"`

	// Worksheet index.
	WorksheetIndex int32 `json:"WorksheetIndex"`

	// Column index of the first value.
	ColumnIndex int32 `json:"ColumnIndex"`

	// Row index of the first value.
	RowIndex int32 `json:"RowIndex"`
}

func NewWorkbook() *Workbook {
	instance := new(Workbook)
	instance.Type_ = "Workbook"
	return instance
}

func (this *Workbook) GetType() string {
	return this.Type_
}

func (this *Workbook) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Workbook) GetWorksheetIndex() int32 {
	return this.WorksheetIndex
}

func (this *Workbook) SetWorksheetIndex(newValue int32) {
	this.WorksheetIndex = newValue
}
func (this *Workbook) GetColumnIndex() int32 {
	return this.ColumnIndex
}

func (this *Workbook) SetColumnIndex(newValue int32) {
	this.ColumnIndex = newValue
}
func (this *Workbook) GetRowIndex() int32 {
	return this.RowIndex
}

func (this *Workbook) SetRowIndex(newValue int32) {
	this.RowIndex = newValue
}

func (this *Workbook) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Workbook"
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
	
	if valWorksheetIndex, ok := GetMapValue(objMap, "worksheetIndex"); ok {
		if valWorksheetIndex != nil {
			var valueForWorksheetIndex int32
			err = json.Unmarshal(*valWorksheetIndex, &valueForWorksheetIndex)
			if err != nil {
				return err
			}
			this.WorksheetIndex = valueForWorksheetIndex
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

	return nil
}
