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

// Table Row.
type ITableRow interface {

	// Cells for the row.
	GetCells() []ITableCell
	SetCells(newValue []ITableCell)

	// Minimal height of the row.
	GetMinimalHeight() float64
	SetMinimalHeight(newValue float64)

	// Height of the row.
	GetHeight() float64
	SetHeight(newValue float64)
}

type TableRow struct {

	// Cells for the row.
	Cells []ITableCell `json:"Cells,omitempty"`

	// Minimal height of the row.
	MinimalHeight float64 `json:"MinimalHeight,omitempty"`

	// Height of the row.
	Height float64 `json:"Height,omitempty"`
}

func NewTableRow() *TableRow {
	instance := new(TableRow)
	return instance
}

func (this *TableRow) GetCells() []ITableCell {
	return this.Cells
}

func (this *TableRow) SetCells(newValue []ITableCell) {
	this.Cells = newValue
}
func (this *TableRow) GetMinimalHeight() float64 {
	return this.MinimalHeight
}

func (this *TableRow) SetMinimalHeight(newValue float64) {
	this.MinimalHeight = newValue
}
func (this *TableRow) GetHeight() float64 {
	return this.Height
}

func (this *TableRow) SetHeight(newValue float64) {
	this.Height = newValue
}

func (this *TableRow) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valCells, ok := GetMapValue(objMap, "cells"); ok {
		if valCells != nil {
			var valueForCells []json.RawMessage
			err = json.Unmarshal(*valCells, &valueForCells)
			if err != nil {
				return err
			}
			valueForICells := make([]ITableCell, len(valueForCells))
			for i, v := range valueForCells {
				vObject, err := createObjectForType("TableCell", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForICells[i] = vObject.(ITableCell)
				}
			}
			this.Cells = valueForICells
		}
	}
	
	if valMinimalHeight, ok := GetMapValue(objMap, "minimalHeight"); ok {
		if valMinimalHeight != nil {
			var valueForMinimalHeight float64
			err = json.Unmarshal(*valMinimalHeight, &valueForMinimalHeight)
			if err != nil {
				return err
			}
			this.MinimalHeight = valueForMinimalHeight
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	return nil
}
