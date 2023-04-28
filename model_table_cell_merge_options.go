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

// Table cells merge options
type ITableCellMergeOptions interface {

	// Row index of the first cell
	GetFirstRowIndex() int32
	SetFirstRowIndex(newValue int32)

	// First cell index in the row
	GetFirstCellIndex() int32
	SetFirstCellIndex(newValue int32)

	// Row index of the last cell
	GetLastRowIndex() int32
	SetLastRowIndex(newValue int32)

	// Last cell index in the row
	GetLastCellIndex() int32
	SetLastCellIndex(newValue int32)

	// Allow splitting
	GetAllowSplitting() bool
	SetAllowSplitting(newValue bool)
}

type TableCellMergeOptions struct {

	// Row index of the first cell
	FirstRowIndex int32 `json:"FirstRowIndex"`

	// First cell index in the row
	FirstCellIndex int32 `json:"FirstCellIndex"`

	// Row index of the last cell
	LastRowIndex int32 `json:"LastRowIndex"`

	// Last cell index in the row
	LastCellIndex int32 `json:"LastCellIndex"`

	// Allow splitting
	AllowSplitting bool `json:"AllowSplitting"`
}

func NewTableCellMergeOptions() *TableCellMergeOptions {
	instance := new(TableCellMergeOptions)
	return instance
}

func (this *TableCellMergeOptions) GetFirstRowIndex() int32 {
	return this.FirstRowIndex
}

func (this *TableCellMergeOptions) SetFirstRowIndex(newValue int32) {
	this.FirstRowIndex = newValue
}
func (this *TableCellMergeOptions) GetFirstCellIndex() int32 {
	return this.FirstCellIndex
}

func (this *TableCellMergeOptions) SetFirstCellIndex(newValue int32) {
	this.FirstCellIndex = newValue
}
func (this *TableCellMergeOptions) GetLastRowIndex() int32 {
	return this.LastRowIndex
}

func (this *TableCellMergeOptions) SetLastRowIndex(newValue int32) {
	this.LastRowIndex = newValue
}
func (this *TableCellMergeOptions) GetLastCellIndex() int32 {
	return this.LastCellIndex
}

func (this *TableCellMergeOptions) SetLastCellIndex(newValue int32) {
	this.LastCellIndex = newValue
}
func (this *TableCellMergeOptions) GetAllowSplitting() bool {
	return this.AllowSplitting
}

func (this *TableCellMergeOptions) SetAllowSplitting(newValue bool) {
	this.AllowSplitting = newValue
}

func (this *TableCellMergeOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valFirstRowIndex, ok := objMap["firstRowIndex"]; ok {
		if valFirstRowIndex != nil {
			var valueForFirstRowIndex int32
			err = json.Unmarshal(*valFirstRowIndex, &valueForFirstRowIndex)
			if err != nil {
				return err
			}
			this.FirstRowIndex = valueForFirstRowIndex
		}
	}
	if valFirstRowIndexCap, ok := objMap["FirstRowIndex"]; ok {
		if valFirstRowIndexCap != nil {
			var valueForFirstRowIndex int32
			err = json.Unmarshal(*valFirstRowIndexCap, &valueForFirstRowIndex)
			if err != nil {
				return err
			}
			this.FirstRowIndex = valueForFirstRowIndex
		}
	}
	
	if valFirstCellIndex, ok := objMap["firstCellIndex"]; ok {
		if valFirstCellIndex != nil {
			var valueForFirstCellIndex int32
			err = json.Unmarshal(*valFirstCellIndex, &valueForFirstCellIndex)
			if err != nil {
				return err
			}
			this.FirstCellIndex = valueForFirstCellIndex
		}
	}
	if valFirstCellIndexCap, ok := objMap["FirstCellIndex"]; ok {
		if valFirstCellIndexCap != nil {
			var valueForFirstCellIndex int32
			err = json.Unmarshal(*valFirstCellIndexCap, &valueForFirstCellIndex)
			if err != nil {
				return err
			}
			this.FirstCellIndex = valueForFirstCellIndex
		}
	}
	
	if valLastRowIndex, ok := objMap["lastRowIndex"]; ok {
		if valLastRowIndex != nil {
			var valueForLastRowIndex int32
			err = json.Unmarshal(*valLastRowIndex, &valueForLastRowIndex)
			if err != nil {
				return err
			}
			this.LastRowIndex = valueForLastRowIndex
		}
	}
	if valLastRowIndexCap, ok := objMap["LastRowIndex"]; ok {
		if valLastRowIndexCap != nil {
			var valueForLastRowIndex int32
			err = json.Unmarshal(*valLastRowIndexCap, &valueForLastRowIndex)
			if err != nil {
				return err
			}
			this.LastRowIndex = valueForLastRowIndex
		}
	}
	
	if valLastCellIndex, ok := objMap["lastCellIndex"]; ok {
		if valLastCellIndex != nil {
			var valueForLastCellIndex int32
			err = json.Unmarshal(*valLastCellIndex, &valueForLastCellIndex)
			if err != nil {
				return err
			}
			this.LastCellIndex = valueForLastCellIndex
		}
	}
	if valLastCellIndexCap, ok := objMap["LastCellIndex"]; ok {
		if valLastCellIndexCap != nil {
			var valueForLastCellIndex int32
			err = json.Unmarshal(*valLastCellIndexCap, &valueForLastCellIndex)
			if err != nil {
				return err
			}
			this.LastCellIndex = valueForLastCellIndex
		}
	}
	
	if valAllowSplitting, ok := objMap["allowSplitting"]; ok {
		if valAllowSplitting != nil {
			var valueForAllowSplitting bool
			err = json.Unmarshal(*valAllowSplitting, &valueForAllowSplitting)
			if err != nil {
				return err
			}
			this.AllowSplitting = valueForAllowSplitting
		}
	}
	if valAllowSplittingCap, ok := objMap["AllowSplitting"]; ok {
		if valAllowSplittingCap != nil {
			var valueForAllowSplitting bool
			err = json.Unmarshal(*valAllowSplittingCap, &valueForAllowSplitting)
			if err != nil {
				return err
			}
			this.AllowSplitting = valueForAllowSplitting
		}
	}

	return nil
}
