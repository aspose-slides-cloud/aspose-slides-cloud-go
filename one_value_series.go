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

// One value series.
type IOneValueSeries interface {

	// Series type.
	getType() string
	setType(newValue string)

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

	// Data point type.
	getDataPointType() string
	setDataPointType(newValue string)

	// Gets or sets the values.
	getDataPoints() []IOneValueChartDataPoint
	setDataPoints(newValue []IOneValueChartDataPoint)
}

type OneValueSeries struct {

	// Series type.
	Type_ string `json:"Type,omitempty"`

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
	Order int32 `json:"Order,omitempty"`

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
	Explosion int32 `json:"Explosion,omitempty"`

	// Series marker.
	Marker ISeriesMarker `json:"Marker,omitempty"`

	// Fill properties set for the series.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Effect properties set for the series.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Line properties set for the series.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Data point type.
	DataPointType string `json:"DataPointType"`

	// Gets or sets the values.
	DataPoints []IOneValueChartDataPoint `json:"DataPoints,omitempty"`
}

func NewOneValueSeries() *OneValueSeries {
	instance := new(OneValueSeries)
	instance.Type_ = ""
	instance.DataPointType = "OneValue"
	return instance
}

func (this *OneValueSeries) getType() string {
	return this.Type_
}

func (this *OneValueSeries) setType(newValue string) {
	this.Type_ = newValue
}
func (this *OneValueSeries) getName() string {
	return this.Name
}

func (this *OneValueSeries) setName(newValue string) {
	this.Name = newValue
}
func (this *OneValueSeries) getIsColorVaried() bool {
	return this.IsColorVaried
}

func (this *OneValueSeries) setIsColorVaried(newValue bool) {
	this.IsColorVaried = newValue
}
func (this *OneValueSeries) getInvertedSolidFillColor() string {
	return this.InvertedSolidFillColor
}

func (this *OneValueSeries) setInvertedSolidFillColor(newValue string) {
	this.InvertedSolidFillColor = newValue
}
func (this *OneValueSeries) getSmooth() bool {
	return this.Smooth
}

func (this *OneValueSeries) setSmooth(newValue bool) {
	this.Smooth = newValue
}
func (this *OneValueSeries) getPlotOnSecondAxis() bool {
	return this.PlotOnSecondAxis
}

func (this *OneValueSeries) setPlotOnSecondAxis(newValue bool) {
	this.PlotOnSecondAxis = newValue
}
func (this *OneValueSeries) getOrder() int32 {
	return this.Order
}

func (this *OneValueSeries) setOrder(newValue int32) {
	this.Order = newValue
}
func (this *OneValueSeries) getNumberFormatOfYValues() string {
	return this.NumberFormatOfYValues
}

func (this *OneValueSeries) setNumberFormatOfYValues(newValue string) {
	this.NumberFormatOfYValues = newValue
}
func (this *OneValueSeries) getNumberFormatOfXValues() string {
	return this.NumberFormatOfXValues
}

func (this *OneValueSeries) setNumberFormatOfXValues(newValue string) {
	this.NumberFormatOfXValues = newValue
}
func (this *OneValueSeries) getNumberFormatOfValues() string {
	return this.NumberFormatOfValues
}

func (this *OneValueSeries) setNumberFormatOfValues(newValue string) {
	this.NumberFormatOfValues = newValue
}
func (this *OneValueSeries) getNumberFormatOfBubbleSizes() string {
	return this.NumberFormatOfBubbleSizes
}

func (this *OneValueSeries) setNumberFormatOfBubbleSizes(newValue string) {
	this.NumberFormatOfBubbleSizes = newValue
}
func (this *OneValueSeries) getInvertIfNegative() bool {
	return this.InvertIfNegative
}

func (this *OneValueSeries) setInvertIfNegative(newValue bool) {
	this.InvertIfNegative = newValue
}
func (this *OneValueSeries) getExplosion() int32 {
	return this.Explosion
}

func (this *OneValueSeries) setExplosion(newValue int32) {
	this.Explosion = newValue
}
func (this *OneValueSeries) getMarker() ISeriesMarker {
	return this.Marker
}

func (this *OneValueSeries) setMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this *OneValueSeries) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *OneValueSeries) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *OneValueSeries) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *OneValueSeries) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *OneValueSeries) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *OneValueSeries) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *OneValueSeries) getDataPointType() string {
	return this.DataPointType
}

func (this *OneValueSeries) setDataPointType(newValue string) {
	this.DataPointType = newValue
}
func (this *OneValueSeries) getDataPoints() []IOneValueChartDataPoint {
	return this.DataPoints
}

func (this *OneValueSeries) setDataPoints(newValue []IOneValueChartDataPoint) {
	this.DataPoints = newValue
}

func (this *OneValueSeries) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = ""
	if valType, ok := objMap["type"]; ok {
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
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valIsColorVaried, ok := objMap["isColorVaried"]; ok {
		if valIsColorVaried != nil {
			var valueForIsColorVaried bool
			err = json.Unmarshal(*valIsColorVaried, &valueForIsColorVaried)
			if err != nil {
				return err
			}
			this.IsColorVaried = valueForIsColorVaried
		}
	}
	if valIsColorVariedCap, ok := objMap["IsColorVaried"]; ok {
		if valIsColorVariedCap != nil {
			var valueForIsColorVaried bool
			err = json.Unmarshal(*valIsColorVariedCap, &valueForIsColorVaried)
			if err != nil {
				return err
			}
			this.IsColorVaried = valueForIsColorVaried
		}
	}
	
	if valInvertedSolidFillColor, ok := objMap["invertedSolidFillColor"]; ok {
		if valInvertedSolidFillColor != nil {
			var valueForInvertedSolidFillColor string
			err = json.Unmarshal(*valInvertedSolidFillColor, &valueForInvertedSolidFillColor)
			if err != nil {
				return err
			}
			this.InvertedSolidFillColor = valueForInvertedSolidFillColor
		}
	}
	if valInvertedSolidFillColorCap, ok := objMap["InvertedSolidFillColor"]; ok {
		if valInvertedSolidFillColorCap != nil {
			var valueForInvertedSolidFillColor string
			err = json.Unmarshal(*valInvertedSolidFillColorCap, &valueForInvertedSolidFillColor)
			if err != nil {
				return err
			}
			this.InvertedSolidFillColor = valueForInvertedSolidFillColor
		}
	}
	
	if valSmooth, ok := objMap["smooth"]; ok {
		if valSmooth != nil {
			var valueForSmooth bool
			err = json.Unmarshal(*valSmooth, &valueForSmooth)
			if err != nil {
				return err
			}
			this.Smooth = valueForSmooth
		}
	}
	if valSmoothCap, ok := objMap["Smooth"]; ok {
		if valSmoothCap != nil {
			var valueForSmooth bool
			err = json.Unmarshal(*valSmoothCap, &valueForSmooth)
			if err != nil {
				return err
			}
			this.Smooth = valueForSmooth
		}
	}
	
	if valPlotOnSecondAxis, ok := objMap["plotOnSecondAxis"]; ok {
		if valPlotOnSecondAxis != nil {
			var valueForPlotOnSecondAxis bool
			err = json.Unmarshal(*valPlotOnSecondAxis, &valueForPlotOnSecondAxis)
			if err != nil {
				return err
			}
			this.PlotOnSecondAxis = valueForPlotOnSecondAxis
		}
	}
	if valPlotOnSecondAxisCap, ok := objMap["PlotOnSecondAxis"]; ok {
		if valPlotOnSecondAxisCap != nil {
			var valueForPlotOnSecondAxis bool
			err = json.Unmarshal(*valPlotOnSecondAxisCap, &valueForPlotOnSecondAxis)
			if err != nil {
				return err
			}
			this.PlotOnSecondAxis = valueForPlotOnSecondAxis
		}
	}
	
	if valOrder, ok := objMap["order"]; ok {
		if valOrder != nil {
			var valueForOrder int32
			err = json.Unmarshal(*valOrder, &valueForOrder)
			if err != nil {
				return err
			}
			this.Order = valueForOrder
		}
	}
	if valOrderCap, ok := objMap["Order"]; ok {
		if valOrderCap != nil {
			var valueForOrder int32
			err = json.Unmarshal(*valOrderCap, &valueForOrder)
			if err != nil {
				return err
			}
			this.Order = valueForOrder
		}
	}
	
	if valNumberFormatOfYValues, ok := objMap["numberFormatOfYValues"]; ok {
		if valNumberFormatOfYValues != nil {
			var valueForNumberFormatOfYValues string
			err = json.Unmarshal(*valNumberFormatOfYValues, &valueForNumberFormatOfYValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfYValues = valueForNumberFormatOfYValues
		}
	}
	if valNumberFormatOfYValuesCap, ok := objMap["NumberFormatOfYValues"]; ok {
		if valNumberFormatOfYValuesCap != nil {
			var valueForNumberFormatOfYValues string
			err = json.Unmarshal(*valNumberFormatOfYValuesCap, &valueForNumberFormatOfYValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfYValues = valueForNumberFormatOfYValues
		}
	}
	
	if valNumberFormatOfXValues, ok := objMap["numberFormatOfXValues"]; ok {
		if valNumberFormatOfXValues != nil {
			var valueForNumberFormatOfXValues string
			err = json.Unmarshal(*valNumberFormatOfXValues, &valueForNumberFormatOfXValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfXValues = valueForNumberFormatOfXValues
		}
	}
	if valNumberFormatOfXValuesCap, ok := objMap["NumberFormatOfXValues"]; ok {
		if valNumberFormatOfXValuesCap != nil {
			var valueForNumberFormatOfXValues string
			err = json.Unmarshal(*valNumberFormatOfXValuesCap, &valueForNumberFormatOfXValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfXValues = valueForNumberFormatOfXValues
		}
	}
	
	if valNumberFormatOfValues, ok := objMap["numberFormatOfValues"]; ok {
		if valNumberFormatOfValues != nil {
			var valueForNumberFormatOfValues string
			err = json.Unmarshal(*valNumberFormatOfValues, &valueForNumberFormatOfValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfValues = valueForNumberFormatOfValues
		}
	}
	if valNumberFormatOfValuesCap, ok := objMap["NumberFormatOfValues"]; ok {
		if valNumberFormatOfValuesCap != nil {
			var valueForNumberFormatOfValues string
			err = json.Unmarshal(*valNumberFormatOfValuesCap, &valueForNumberFormatOfValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfValues = valueForNumberFormatOfValues
		}
	}
	
	if valNumberFormatOfBubbleSizes, ok := objMap["numberFormatOfBubbleSizes"]; ok {
		if valNumberFormatOfBubbleSizes != nil {
			var valueForNumberFormatOfBubbleSizes string
			err = json.Unmarshal(*valNumberFormatOfBubbleSizes, &valueForNumberFormatOfBubbleSizes)
			if err != nil {
				return err
			}
			this.NumberFormatOfBubbleSizes = valueForNumberFormatOfBubbleSizes
		}
	}
	if valNumberFormatOfBubbleSizesCap, ok := objMap["NumberFormatOfBubbleSizes"]; ok {
		if valNumberFormatOfBubbleSizesCap != nil {
			var valueForNumberFormatOfBubbleSizes string
			err = json.Unmarshal(*valNumberFormatOfBubbleSizesCap, &valueForNumberFormatOfBubbleSizes)
			if err != nil {
				return err
			}
			this.NumberFormatOfBubbleSizes = valueForNumberFormatOfBubbleSizes
		}
	}
	
	if valInvertIfNegative, ok := objMap["invertIfNegative"]; ok {
		if valInvertIfNegative != nil {
			var valueForInvertIfNegative bool
			err = json.Unmarshal(*valInvertIfNegative, &valueForInvertIfNegative)
			if err != nil {
				return err
			}
			this.InvertIfNegative = valueForInvertIfNegative
		}
	}
	if valInvertIfNegativeCap, ok := objMap["InvertIfNegative"]; ok {
		if valInvertIfNegativeCap != nil {
			var valueForInvertIfNegative bool
			err = json.Unmarshal(*valInvertIfNegativeCap, &valueForInvertIfNegative)
			if err != nil {
				return err
			}
			this.InvertIfNegative = valueForInvertIfNegative
		}
	}
	
	if valExplosion, ok := objMap["explosion"]; ok {
		if valExplosion != nil {
			var valueForExplosion int32
			err = json.Unmarshal(*valExplosion, &valueForExplosion)
			if err != nil {
				return err
			}
			this.Explosion = valueForExplosion
		}
	}
	if valExplosionCap, ok := objMap["Explosion"]; ok {
		if valExplosionCap != nil {
			var valueForExplosion int32
			err = json.Unmarshal(*valExplosionCap, &valueForExplosion)
			if err != nil {
				return err
			}
			this.Explosion = valueForExplosion
		}
	}
	
	if valMarker, ok := objMap["marker"]; ok {
		if valMarker != nil {
			var valueForMarker SeriesMarker
			err = json.Unmarshal(*valMarker, &valueForMarker)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SeriesMarker", *valMarker)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMarker, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISeriesMarker)
			if ok {
				this.Marker = vInterfaceObject
			}
		}
	}
	if valMarkerCap, ok := objMap["Marker"]; ok {
		if valMarkerCap != nil {
			var valueForMarker SeriesMarker
			err = json.Unmarshal(*valMarkerCap, &valueForMarker)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SeriesMarker", *valMarkerCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMarkerCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISeriesMarker)
			if ok {
				this.Marker = vInterfaceObject
			}
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
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
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	this.DataPointType = "OneValue"
	if valDataPointType, ok := objMap["dataPointType"]; ok {
		if valDataPointType != nil {
			var valueForDataPointType string
			err = json.Unmarshal(*valDataPointType, &valueForDataPointType)
			if err != nil {
				var valueForDataPointTypeInt int32
				err = json.Unmarshal(*valDataPointType, &valueForDataPointTypeInt)
				if err != nil {
					return err
				}
				this.DataPointType = string(valueForDataPointTypeInt)
			} else {
				this.DataPointType = valueForDataPointType
			}
		}
	}
	if valDataPointTypeCap, ok := objMap["DataPointType"]; ok {
		if valDataPointTypeCap != nil {
			var valueForDataPointType string
			err = json.Unmarshal(*valDataPointTypeCap, &valueForDataPointType)
			if err != nil {
				var valueForDataPointTypeInt int32
				err = json.Unmarshal(*valDataPointTypeCap, &valueForDataPointTypeInt)
				if err != nil {
					return err
				}
				this.DataPointType = string(valueForDataPointTypeInt)
			} else {
				this.DataPointType = valueForDataPointType
			}
		}
	}
	
	if valDataPoints, ok := objMap["dataPoints"]; ok {
		if valDataPoints != nil {
			var valueForDataPoints []json.RawMessage
			err = json.Unmarshal(*valDataPoints, &valueForDataPoints)
			if err != nil {
				return err
			}
			valueForIDataPoints := make([]IOneValueChartDataPoint, len(valueForDataPoints))
			for i, v := range valueForDataPoints {
				vObject, err := createObjectForType("OneValueChartDataPoint", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIDataPoints[i] = vObject.(IOneValueChartDataPoint)
				}
			}
			this.DataPoints = valueForIDataPoints
		}
	}
	if valDataPointsCap, ok := objMap["DataPoints"]; ok {
		if valDataPointsCap != nil {
			var valueForDataPoints []json.RawMessage
			err = json.Unmarshal(*valDataPointsCap, &valueForDataPoints)
			if err != nil {
				return err
			}
			valueForIDataPoints := make([]IOneValueChartDataPoint, len(valueForDataPoints))
			for i, v := range valueForDataPoints {
				vObject, err := createObjectForType("OneValueChartDataPoint", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIDataPoints[i] = vObject.(IOneValueChartDataPoint)
				}
			}
			this.DataPoints = valueForIDataPoints
		}
	}

	return nil
}
