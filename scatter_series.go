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

// A scatter series
type IScatterSeries interface {

	// Data point type.
	getDataPointType() ChartDataPointType
	setDataPointType(newValue ChartDataPointType)

	// Series type.
	getType() ChartType
	setType(newValue ChartType)

	// Series name.
	getName() string
	setName(newValue string)

	// True if each data marker in the series has a different color.
	getIsColorVaried() bool
	setIsColorVaried(newValue bool)

	// Invert solid color for the series.
	getInvertedSolidFillColor() string
	setInvertedSolidFillColor(newValue string)

	// True if curve smoothing is turned on. Applies only to line and scatter connected by lines charts.
	getSmooth() bool
	setSmooth(newValue bool)

	// True if the series is plotted on second value axis.
	getPlotOnSecondAxis() bool
	setPlotOnSecondAxis(newValue bool)

	// Series order.
	getOrder() int32
	setOrder(newValue int32)

	// The number format for the series y values.
	getNumberFormatOfYValues() string
	setNumberFormatOfYValues(newValue string)

	// The number format for the series x values.
	getNumberFormatOfXValues() string
	setNumberFormatOfXValues(newValue string)

	// The number format for the series values.
	getNumberFormatOfValues() string
	setNumberFormatOfValues(newValue string)

	// The number format for the series bubble sizes.
	getNumberFormatOfBubbleSizes() string
	setNumberFormatOfBubbleSizes(newValue string)

	// True if the series shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	getInvertIfNegative() bool
	setInvertIfNegative(newValue bool)

	// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
	getExplosion() int32
	setExplosion(newValue int32)

	// Series marker.
	getMarker() ISeriesMarker
	setMarker(newValue ISeriesMarker)

	// Fill properties set for the series.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Effect properties set for the series.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Line properties set for the series.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Gets or sets the values.
	getDataPoints() []ScatterChartDataPoint
	setDataPoints(newValue []ScatterChartDataPoint)
}

type ScatterSeries struct {

	// Data point type.
	DataPointType ChartDataPointType `json:"DataPointType"`

	// Series type.
	Type_ ChartType `json:"Type"`

	// Series name.
	Name string `json:"Name,omitempty"`

	// True if each data marker in the series has a different color.
	IsColorVaried bool `json:"IsColorVaried"`

	// Invert solid color for the series.
	InvertedSolidFillColor string `json:"InvertedSolidFillColor,omitempty"`

	// True if curve smoothing is turned on. Applies only to line and scatter connected by lines charts.
	Smooth bool `json:"Smooth"`

	// True if the series is plotted on second value axis.
	PlotOnSecondAxis bool `json:"PlotOnSecondAxis"`

	// Series order.
	Order int32 `json:"Order"`

	// The number format for the series y values.
	NumberFormatOfYValues string `json:"NumberFormatOfYValues,omitempty"`

	// The number format for the series x values.
	NumberFormatOfXValues string `json:"NumberFormatOfXValues,omitempty"`

	// The number format for the series values.
	NumberFormatOfValues string `json:"NumberFormatOfValues,omitempty"`

	// The number format for the series bubble sizes.
	NumberFormatOfBubbleSizes string `json:"NumberFormatOfBubbleSizes,omitempty"`

	// True if the series shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	InvertIfNegative bool `json:"InvertIfNegative"`

	// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
	Explosion int32 `json:"Explosion"`

	// Series marker.
	Marker ISeriesMarker `json:"Marker,omitempty"`

	// Fill properties set for the series.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Effect properties set for the series.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Line properties set for the series.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Gets or sets the values.
	DataPoints []ScatterChartDataPoint `json:"DataPoints,omitempty"`
}

func (this ScatterSeries) getDataPointType() ChartDataPointType {
	return this.DataPointType
}

func (this ScatterSeries) setDataPointType(newValue ChartDataPointType) {
	this.DataPointType = newValue
}
func (this ScatterSeries) getType() ChartType {
	return this.Type_
}

func (this ScatterSeries) setType(newValue ChartType) {
	this.Type_ = newValue
}
func (this ScatterSeries) getName() string {
	return this.Name
}

func (this ScatterSeries) setName(newValue string) {
	this.Name = newValue
}
func (this ScatterSeries) getIsColorVaried() bool {
	return this.IsColorVaried
}

func (this ScatterSeries) setIsColorVaried(newValue bool) {
	this.IsColorVaried = newValue
}
func (this ScatterSeries) getInvertedSolidFillColor() string {
	return this.InvertedSolidFillColor
}

func (this ScatterSeries) setInvertedSolidFillColor(newValue string) {
	this.InvertedSolidFillColor = newValue
}
func (this ScatterSeries) getSmooth() bool {
	return this.Smooth
}

func (this ScatterSeries) setSmooth(newValue bool) {
	this.Smooth = newValue
}
func (this ScatterSeries) getPlotOnSecondAxis() bool {
	return this.PlotOnSecondAxis
}

func (this ScatterSeries) setPlotOnSecondAxis(newValue bool) {
	this.PlotOnSecondAxis = newValue
}
func (this ScatterSeries) getOrder() int32 {
	return this.Order
}

func (this ScatterSeries) setOrder(newValue int32) {
	this.Order = newValue
}
func (this ScatterSeries) getNumberFormatOfYValues() string {
	return this.NumberFormatOfYValues
}

func (this ScatterSeries) setNumberFormatOfYValues(newValue string) {
	this.NumberFormatOfYValues = newValue
}
func (this ScatterSeries) getNumberFormatOfXValues() string {
	return this.NumberFormatOfXValues
}

func (this ScatterSeries) setNumberFormatOfXValues(newValue string) {
	this.NumberFormatOfXValues = newValue
}
func (this ScatterSeries) getNumberFormatOfValues() string {
	return this.NumberFormatOfValues
}

func (this ScatterSeries) setNumberFormatOfValues(newValue string) {
	this.NumberFormatOfValues = newValue
}
func (this ScatterSeries) getNumberFormatOfBubbleSizes() string {
	return this.NumberFormatOfBubbleSizes
}

func (this ScatterSeries) setNumberFormatOfBubbleSizes(newValue string) {
	this.NumberFormatOfBubbleSizes = newValue
}
func (this ScatterSeries) getInvertIfNegative() bool {
	return this.InvertIfNegative
}

func (this ScatterSeries) setInvertIfNegative(newValue bool) {
	this.InvertIfNegative = newValue
}
func (this ScatterSeries) getExplosion() int32 {
	return this.Explosion
}

func (this ScatterSeries) setExplosion(newValue int32) {
	this.Explosion = newValue
}
func (this ScatterSeries) getMarker() ISeriesMarker {
	return this.Marker
}

func (this ScatterSeries) setMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this ScatterSeries) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this ScatterSeries) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this ScatterSeries) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this ScatterSeries) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this ScatterSeries) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this ScatterSeries) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this ScatterSeries) getDataPoints() []ScatterChartDataPoint {
	return this.DataPoints
}

func (this ScatterSeries) setDataPoints(newValue []ScatterChartDataPoint) {
	this.DataPoints = newValue
}

func (this *ScatterSeries) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valDataPointType, ok := objMap["DataPointType"]; ok {
		if valDataPointType != nil {
			var valueForDataPointType ChartDataPointType
			err = json.Unmarshal(*valDataPointType, &valueForDataPointType)
			if err != nil {
				return err
			}
			this.DataPointType = valueForDataPointType
		}
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType ChartType
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valName, ok := objMap["Name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}

	if valIsColorVaried, ok := objMap["IsColorVaried"]; ok {
		if valIsColorVaried != nil {
			var valueForIsColorVaried bool
			err = json.Unmarshal(*valIsColorVaried, &valueForIsColorVaried)
			if err != nil {
				return err
			}
			this.IsColorVaried = valueForIsColorVaried
		}
	}

	if valInvertedSolidFillColor, ok := objMap["InvertedSolidFillColor"]; ok {
		if valInvertedSolidFillColor != nil {
			var valueForInvertedSolidFillColor string
			err = json.Unmarshal(*valInvertedSolidFillColor, &valueForInvertedSolidFillColor)
			if err != nil {
				return err
			}
			this.InvertedSolidFillColor = valueForInvertedSolidFillColor
		}
	}

	if valSmooth, ok := objMap["Smooth"]; ok {
		if valSmooth != nil {
			var valueForSmooth bool
			err = json.Unmarshal(*valSmooth, &valueForSmooth)
			if err != nil {
				return err
			}
			this.Smooth = valueForSmooth
		}
	}

	if valPlotOnSecondAxis, ok := objMap["PlotOnSecondAxis"]; ok {
		if valPlotOnSecondAxis != nil {
			var valueForPlotOnSecondAxis bool
			err = json.Unmarshal(*valPlotOnSecondAxis, &valueForPlotOnSecondAxis)
			if err != nil {
				return err
			}
			this.PlotOnSecondAxis = valueForPlotOnSecondAxis
		}
	}

	if valOrder, ok := objMap["Order"]; ok {
		if valOrder != nil {
			var valueForOrder int32
			err = json.Unmarshal(*valOrder, &valueForOrder)
			if err != nil {
				return err
			}
			this.Order = valueForOrder
		}
	}

	if valNumberFormatOfYValues, ok := objMap["NumberFormatOfYValues"]; ok {
		if valNumberFormatOfYValues != nil {
			var valueForNumberFormatOfYValues string
			err = json.Unmarshal(*valNumberFormatOfYValues, &valueForNumberFormatOfYValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfYValues = valueForNumberFormatOfYValues
		}
	}

	if valNumberFormatOfXValues, ok := objMap["NumberFormatOfXValues"]; ok {
		if valNumberFormatOfXValues != nil {
			var valueForNumberFormatOfXValues string
			err = json.Unmarshal(*valNumberFormatOfXValues, &valueForNumberFormatOfXValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfXValues = valueForNumberFormatOfXValues
		}
	}

	if valNumberFormatOfValues, ok := objMap["NumberFormatOfValues"]; ok {
		if valNumberFormatOfValues != nil {
			var valueForNumberFormatOfValues string
			err = json.Unmarshal(*valNumberFormatOfValues, &valueForNumberFormatOfValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfValues = valueForNumberFormatOfValues
		}
	}

	if valNumberFormatOfBubbleSizes, ok := objMap["NumberFormatOfBubbleSizes"]; ok {
		if valNumberFormatOfBubbleSizes != nil {
			var valueForNumberFormatOfBubbleSizes string
			err = json.Unmarshal(*valNumberFormatOfBubbleSizes, &valueForNumberFormatOfBubbleSizes)
			if err != nil {
				return err
			}
			this.NumberFormatOfBubbleSizes = valueForNumberFormatOfBubbleSizes
		}
	}

	if valInvertIfNegative, ok := objMap["InvertIfNegative"]; ok {
		if valInvertIfNegative != nil {
			var valueForInvertIfNegative bool
			err = json.Unmarshal(*valInvertIfNegative, &valueForInvertIfNegative)
			if err != nil {
				return err
			}
			this.InvertIfNegative = valueForInvertIfNegative
		}
	}

	if valExplosion, ok := objMap["Explosion"]; ok {
		if valExplosion != nil {
			var valueForExplosion int32
			err = json.Unmarshal(*valExplosion, &valueForExplosion)
			if err != nil {
				return err
			}
			this.Explosion = valueForExplosion
		}
	}

	if valMarker, ok := objMap["Marker"]; ok {
		if valMarker != nil {
			var valueForMarker SeriesMarker
			err = json.Unmarshal(*valMarker, &valueForMarker)
			if err != nil {
				return err
			}
			this.Marker = valueForMarker
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

	if valEffectFormat, ok := objMap["EffectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}

	if valLineFormat, ok := objMap["LineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}

	if valDataPoints, ok := objMap["DataPoints"]; ok {
		if valDataPoints != nil {
			var valueForDataPoints []ScatterChartDataPoint
			err = json.Unmarshal(*valDataPoints, &valueForDataPoints)
			if err != nil {
				return err
			}
			this.DataPoints = valueForDataPoints
		}
	}

    return nil
}
