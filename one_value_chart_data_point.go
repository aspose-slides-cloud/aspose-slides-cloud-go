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

// One value chart data point.
type IOneValueChartDataPoint interface {

	// Value.
	getValue() float64
	setValue(newValue float64)

	// SetAsTotal. Applied to Waterfall data points only.
	getSetAsTotal() bool
	setSetAsTotal(newValue bool)
}

type OneValueChartDataPoint struct {

	// Value.
	Value float64 `json:"Value"`

	// SetAsTotal. Applied to Waterfall data points only.
	SetAsTotal bool `json:"SetAsTotal"`
}

func NewOneValueChartDataPoint() *OneValueChartDataPoint {
	instance := new(OneValueChartDataPoint)
	return instance
}

func (this *OneValueChartDataPoint) getValue() float64 {
	return this.Value
}

func (this *OneValueChartDataPoint) setValue(newValue float64) {
	this.Value = newValue
}
func (this *OneValueChartDataPoint) getSetAsTotal() bool {
	return this.SetAsTotal
}

func (this *OneValueChartDataPoint) setSetAsTotal(newValue bool) {
	this.SetAsTotal = newValue
}

func (this *OneValueChartDataPoint) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valValue, ok := objMap["value"]; ok {
		if valValue != nil {
			var valueForValue float64
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	if valValueCap, ok := objMap["Value"]; ok {
		if valValueCap != nil {
			var valueForValue float64
			err = json.Unmarshal(*valValueCap, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	
	if valSetAsTotal, ok := objMap["setAsTotal"]; ok {
		if valSetAsTotal != nil {
			var valueForSetAsTotal bool
			err = json.Unmarshal(*valSetAsTotal, &valueForSetAsTotal)
			if err != nil {
				return err
			}
			this.SetAsTotal = valueForSetAsTotal
		}
	}
	if valSetAsTotalCap, ok := objMap["SetAsTotal"]; ok {
		if valSetAsTotalCap != nil {
			var valueForSetAsTotal bool
			err = json.Unmarshal(*valSetAsTotalCap, &valueForSetAsTotal)
			if err != nil {
				return err
			}
			this.SetAsTotal = valueForSetAsTotal
		}
	}

	return nil
}
