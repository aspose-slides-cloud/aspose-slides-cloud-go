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


type ITableRow interface {

	getCells() []TableCell
	setCells(newValue []TableCell)

	getMinimalHeight() float64
	setMinimalHeight(newValue float64)

	getHeight() float64
	setHeight(newValue float64)
}

type TableRow struct {

	Cells []TableCell `json:"Cells,omitempty"`

	MinimalHeight float64 `json:"MinimalHeight"`

	Height float64 `json:"Height"`
}

func (this TableRow) getCells() []TableCell {
	return this.Cells
}

func (this TableRow) setCells(newValue []TableCell) {
	this.Cells = newValue
}
func (this TableRow) getMinimalHeight() float64 {
	return this.MinimalHeight
}

func (this TableRow) setMinimalHeight(newValue float64) {
	this.MinimalHeight = newValue
}
func (this TableRow) getHeight() float64 {
	return this.Height
}

func (this TableRow) setHeight(newValue float64) {
	this.Height = newValue
}

func (this *TableRow) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valCells, ok := objMap["Cells"]; ok {
		if valCells != nil {
			var valueForCells []TableCell
			err = json.Unmarshal(*valCells, &valueForCells)
			if err != nil {
				return err
			}
			this.Cells = valueForCells
		}
	}

	if valMinimalHeight, ok := objMap["MinimalHeight"]; ok {
		if valMinimalHeight != nil {
			var valueForMinimalHeight float64
			err = json.Unmarshal(*valMinimalHeight, &valueForMinimalHeight)
			if err != nil {
				return err
			}
			this.MinimalHeight = valueForMinimalHeight
		}
	}

	if valHeight, ok := objMap["Height"]; ok {
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
