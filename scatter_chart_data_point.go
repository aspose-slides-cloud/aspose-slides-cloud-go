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

// Scatter chart (two-dimensional) data point
type IScatterChartDataPoint interface {

	// X-value
	getXValue() float64
	setXValue(newValue float64)

	// Y-value
	getYValue() float64
	setYValue(newValue float64)
}

type ScatterChartDataPoint struct {

	// X-value
	XValue float64 `json:"XValue"`

	// Y-value
	YValue float64 `json:"YValue"`
}

func (this ScatterChartDataPoint) getXValue() float64 {
	return this.XValue
}

func (this ScatterChartDataPoint) setXValue(newValue float64) {
	this.XValue = newValue
}
func (this ScatterChartDataPoint) getYValue() float64 {
	return this.YValue
}

func (this ScatterChartDataPoint) setYValue(newValue float64) {
	this.YValue = newValue
}

func (this *ScatterChartDataPoint) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valXValue, ok := objMap["XValue"]; ok {
		if valXValue != nil {
			var valueForXValue float64
			err = json.Unmarshal(*valXValue, &valueForXValue)
			if err != nil {
				return err
			}
			this.XValue = valueForXValue
		}
	}
	
	if valYValue, ok := objMap["YValue"]; ok {
		if valYValue != nil {
			var valueForYValue float64
			err = json.Unmarshal(*valYValue, &valueForYValue)
			if err != nil {
				return err
			}
			this.YValue = valueForYValue
		}
	}

    return nil
}
