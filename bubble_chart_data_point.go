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

// Bubble chart data point.
type IBubbleChartDataPoint interface {

	// X-value
	getXValue() float64
	setXValue(newValue float64)

	// Y-value
	getYValue() float64
	setYValue(newValue float64)

	// Bubble size.
	getBubbleSize() float64
	setBubbleSize(newValue float64)
}

type BubbleChartDataPoint struct {

	// X-value
	XValue float64 `json:"XValue"`

	// Y-value
	YValue float64 `json:"YValue"`

	// Bubble size.
	BubbleSize float64 `json:"BubbleSize"`
}

func (this *BubbleChartDataPoint) getXValue() float64 {
	return this.XValue
}

func (this *BubbleChartDataPoint) setXValue(newValue float64) {
	this.XValue = newValue
}
func (this *BubbleChartDataPoint) getYValue() float64 {
	return this.YValue
}

func (this *BubbleChartDataPoint) setYValue(newValue float64) {
	this.YValue = newValue
}
func (this *BubbleChartDataPoint) getBubbleSize() float64 {
	return this.BubbleSize
}

func (this *BubbleChartDataPoint) setBubbleSize(newValue float64) {
	this.BubbleSize = newValue
}

func (this *BubbleChartDataPoint) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valXValue, ok := objMap["xValue"]; ok {
		if valXValue != nil {
			var valueForXValue float64
			err = json.Unmarshal(*valXValue, &valueForXValue)
			if err != nil {
				return err
			}
			this.XValue = valueForXValue
		}
	}
	if valXValueCap, ok := objMap["XValue"]; ok {
		if valXValueCap != nil {
			var valueForXValue float64
			err = json.Unmarshal(*valXValueCap, &valueForXValue)
			if err != nil {
				return err
			}
			this.XValue = valueForXValue
		}
	}
	
	if valYValue, ok := objMap["yValue"]; ok {
		if valYValue != nil {
			var valueForYValue float64
			err = json.Unmarshal(*valYValue, &valueForYValue)
			if err != nil {
				return err
			}
			this.YValue = valueForYValue
		}
	}
	if valYValueCap, ok := objMap["YValue"]; ok {
		if valYValueCap != nil {
			var valueForYValue float64
			err = json.Unmarshal(*valYValueCap, &valueForYValue)
			if err != nil {
				return err
			}
			this.YValue = valueForYValue
		}
	}
	
	if valBubbleSize, ok := objMap["bubbleSize"]; ok {
		if valBubbleSize != nil {
			var valueForBubbleSize float64
			err = json.Unmarshal(*valBubbleSize, &valueForBubbleSize)
			if err != nil {
				return err
			}
			this.BubbleSize = valueForBubbleSize
		}
	}
	if valBubbleSizeCap, ok := objMap["BubbleSize"]; ok {
		if valBubbleSizeCap != nil {
			var valueForBubbleSize float64
			err = json.Unmarshal(*valBubbleSizeCap, &valueForBubbleSize)
			if err != nil {
				return err
			}
			this.BubbleSize = valueForBubbleSize
		}
	}

    return nil
}
