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

// Custom dash pattern.
type ICustomDashPattern interface {

	// Pattern items.
	getItems() []float64
	setItems(newValue []float64)
}

type CustomDashPattern struct {

	// Pattern items.
	Items []float64 `json:"Items,omitempty"`
}

func NewCustomDashPattern() *CustomDashPattern {
	instance := new(CustomDashPattern)
	return instance
}

func (this *CustomDashPattern) getItems() []float64 {
	return this.Items
}

func (this *CustomDashPattern) setItems(newValue []float64) {
	this.Items = newValue
}

func (this *CustomDashPattern) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valItems, ok := objMap["items"]; ok {
		if valItems != nil {
			var valueForItems []float64
			err = json.Unmarshal(*valItems, &valueForItems)
			if err != nil {
				return err
			}
			this.Items = valueForItems
		}
	}
	if valItemsCap, ok := objMap["Items"]; ok {
		if valItemsCap != nil {
			var valueForItems []float64
			err = json.Unmarshal(*valItemsCap, &valueForItems)
			if err != nil {
				return err
			}
			this.Items = valueForItems
		}
	}

    return nil
}
