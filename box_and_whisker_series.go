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
type IBoxAndWhiskerSeries interface {

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

	// The number format for the series values.
	getNumberFormatOfValues() string
	setNumberFormatOfValues(newValue string)

	// Quartile method.
	getQuartileMethod() string
	setQuartileMethod(newValue string)

	// True if inner points are shown.
	getShowInnerPoints() bool
	setShowInnerPoints(newValue bool)

	// True if mean line is shown.
	getShowMeanLine() bool
	setShowMeanLine(newValue bool)

	// True if mean markers are shown.
	getShowMeanMarkers() bool
	setShowMeanMarkers(newValue bool)

	// True if outlier points are shown.
	getShowOutlierPoints() bool
	setShowOutlierPoints(newValue bool)
}

type BoxAndWhiskerSeries struct {

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

	// The number format for the series values.
	NumberFormatOfValues string `json:"NumberFormatOfValues,omitempty"`

	// Quartile method.
	QuartileMethod string `json:"QuartileMethod,omitempty"`

	// True if inner points are shown.
	ShowInnerPoints bool `json:"ShowInnerPoints"`

	// True if mean line is shown.
	ShowMeanLine bool `json:"ShowMeanLine"`

	// True if mean markers are shown.
	ShowMeanMarkers bool `json:"ShowMeanMarkers"`

	// True if outlier points are shown.
	ShowOutlierPoints bool `json:"ShowOutlierPoints"`
}

func NewBoxAndWhiskerSeries() *BoxAndWhiskerSeries {
	instance := new(BoxAndWhiskerSeries)
	instance.Type_ = ""
	instance.DataPointType = "OneValue"
	instance.QuartileMethod = ""
	return instance
}

func (this *BoxAndWhiskerSeries) getType() string {
	return this.Type_
}

func (this *BoxAndWhiskerSeries) setType(newValue string) {
	this.Type_ = newValue
}
func (this *BoxAndWhiskerSeries) getName() string {
	return this.Name
}

func (this *BoxAndWhiskerSeries) setName(newValue string) {
	this.Name = newValue
}
func (this *BoxAndWhiskerSeries) getIsColorVaried() bool {
	return this.IsColorVaried
}

func (this *BoxAndWhiskerSeries) setIsColorVaried(newValue bool) {
	this.IsColorVaried = newValue
}
func (this *BoxAndWhiskerSeries) getInvertedSolidFillColor() string {
	return this.InvertedSolidFillColor
}

func (this *BoxAndWhiskerSeries) setInvertedSolidFillColor(newValue string) {
	this.InvertedSolidFillColor = newValue
}
func (this *BoxAndWhiskerSeries) getSmooth() bool {
	return this.Smooth
}

func (this *BoxAndWhiskerSeries) setSmooth(newValue bool) {
	this.Smooth = newValue
}
func (this *BoxAndWhiskerSeries) getPlotOnSecondAxis() bool {
	return this.PlotOnSecondAxis
}

func (this *BoxAndWhiskerSeries) setPlotOnSecondAxis(newValue bool) {
	this.PlotOnSecondAxis = newValue
}
func (this *BoxAndWhiskerSeries) getOrder() int32 {
	return this.Order
}

func (this *BoxAndWhiskerSeries) setOrder(newValue int32) {
	this.Order = newValue
}
func (this *BoxAndWhiskerSeries) getInvertIfNegative() bool {
	return this.InvertIfNegative
}

func (this *BoxAndWhiskerSeries) setInvertIfNegative(newValue bool) {
	this.InvertIfNegative = newValue
}
func (this *BoxAndWhiskerSeries) getExplosion() int32 {
	return this.Explosion
}

func (this *BoxAndWhiskerSeries) setExplosion(newValue int32) {
	this.Explosion = newValue
}
func (this *BoxAndWhiskerSeries) getMarker() ISeriesMarker {
	return this.Marker
}

func (this *BoxAndWhiskerSeries) setMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this *BoxAndWhiskerSeries) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *BoxAndWhiskerSeries) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *BoxAndWhiskerSeries) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *BoxAndWhiskerSeries) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *BoxAndWhiskerSeries) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *BoxAndWhiskerSeries) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *BoxAndWhiskerSeries) getDataPointType() string {
	return this.DataPointType
}

func (this *BoxAndWhiskerSeries) setDataPointType(newValue string) {
	this.DataPointType = newValue
}
func (this *BoxAndWhiskerSeries) getDataPoints() []IOneValueChartDataPoint {
	return this.DataPoints
}

func (this *BoxAndWhiskerSeries) setDataPoints(newValue []IOneValueChartDataPoint) {
	this.DataPoints = newValue
}
func (this *BoxAndWhiskerSeries) getNumberFormatOfValues() string {
	return this.NumberFormatOfValues
}

func (this *BoxAndWhiskerSeries) setNumberFormatOfValues(newValue string) {
	this.NumberFormatOfValues = newValue
}
func (this *BoxAndWhiskerSeries) getQuartileMethod() string {
	return this.QuartileMethod
}

func (this *BoxAndWhiskerSeries) setQuartileMethod(newValue string) {
	this.QuartileMethod = newValue
}
func (this *BoxAndWhiskerSeries) getShowInnerPoints() bool {
	return this.ShowInnerPoints
}

func (this *BoxAndWhiskerSeries) setShowInnerPoints(newValue bool) {
	this.ShowInnerPoints = newValue
}
func (this *BoxAndWhiskerSeries) getShowMeanLine() bool {
	return this.ShowMeanLine
}

func (this *BoxAndWhiskerSeries) setShowMeanLine(newValue bool) {
	this.ShowMeanLine = newValue
}
func (this *BoxAndWhiskerSeries) getShowMeanMarkers() bool {
	return this.ShowMeanMarkers
}

func (this *BoxAndWhiskerSeries) setShowMeanMarkers(newValue bool) {
	this.ShowMeanMarkers = newValue
}
func (this *BoxAndWhiskerSeries) getShowOutlierPoints() bool {
	return this.ShowOutlierPoints
}

func (this *BoxAndWhiskerSeries) setShowOutlierPoints(newValue bool) {
	this.ShowOutlierPoints = newValue
}

func (this *BoxAndWhiskerSeries) UnmarshalJSON(b []byte) error {
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
	this.QuartileMethod = ""
	if valQuartileMethod, ok := objMap["quartileMethod"]; ok {
		if valQuartileMethod != nil {
			var valueForQuartileMethod string
			err = json.Unmarshal(*valQuartileMethod, &valueForQuartileMethod)
			if err != nil {
				var valueForQuartileMethodInt int32
				err = json.Unmarshal(*valQuartileMethod, &valueForQuartileMethodInt)
				if err != nil {
					return err
				}
				this.QuartileMethod = string(valueForQuartileMethodInt)
			} else {
				this.QuartileMethod = valueForQuartileMethod
			}
		}
	}
	if valQuartileMethodCap, ok := objMap["QuartileMethod"]; ok {
		if valQuartileMethodCap != nil {
			var valueForQuartileMethod string
			err = json.Unmarshal(*valQuartileMethodCap, &valueForQuartileMethod)
			if err != nil {
				var valueForQuartileMethodInt int32
				err = json.Unmarshal(*valQuartileMethodCap, &valueForQuartileMethodInt)
				if err != nil {
					return err
				}
				this.QuartileMethod = string(valueForQuartileMethodInt)
			} else {
				this.QuartileMethod = valueForQuartileMethod
			}
		}
	}
	
	if valShowInnerPoints, ok := objMap["showInnerPoints"]; ok {
		if valShowInnerPoints != nil {
			var valueForShowInnerPoints bool
			err = json.Unmarshal(*valShowInnerPoints, &valueForShowInnerPoints)
			if err != nil {
				return err
			}
			this.ShowInnerPoints = valueForShowInnerPoints
		}
	}
	if valShowInnerPointsCap, ok := objMap["ShowInnerPoints"]; ok {
		if valShowInnerPointsCap != nil {
			var valueForShowInnerPoints bool
			err = json.Unmarshal(*valShowInnerPointsCap, &valueForShowInnerPoints)
			if err != nil {
				return err
			}
			this.ShowInnerPoints = valueForShowInnerPoints
		}
	}
	
	if valShowMeanLine, ok := objMap["showMeanLine"]; ok {
		if valShowMeanLine != nil {
			var valueForShowMeanLine bool
			err = json.Unmarshal(*valShowMeanLine, &valueForShowMeanLine)
			if err != nil {
				return err
			}
			this.ShowMeanLine = valueForShowMeanLine
		}
	}
	if valShowMeanLineCap, ok := objMap["ShowMeanLine"]; ok {
		if valShowMeanLineCap != nil {
			var valueForShowMeanLine bool
			err = json.Unmarshal(*valShowMeanLineCap, &valueForShowMeanLine)
			if err != nil {
				return err
			}
			this.ShowMeanLine = valueForShowMeanLine
		}
	}
	
	if valShowMeanMarkers, ok := objMap["showMeanMarkers"]; ok {
		if valShowMeanMarkers != nil {
			var valueForShowMeanMarkers bool
			err = json.Unmarshal(*valShowMeanMarkers, &valueForShowMeanMarkers)
			if err != nil {
				return err
			}
			this.ShowMeanMarkers = valueForShowMeanMarkers
		}
	}
	if valShowMeanMarkersCap, ok := objMap["ShowMeanMarkers"]; ok {
		if valShowMeanMarkersCap != nil {
			var valueForShowMeanMarkers bool
			err = json.Unmarshal(*valShowMeanMarkersCap, &valueForShowMeanMarkers)
			if err != nil {
				return err
			}
			this.ShowMeanMarkers = valueForShowMeanMarkers
		}
	}
	
	if valShowOutlierPoints, ok := objMap["showOutlierPoints"]; ok {
		if valShowOutlierPoints != nil {
			var valueForShowOutlierPoints bool
			err = json.Unmarshal(*valShowOutlierPoints, &valueForShowOutlierPoints)
			if err != nil {
				return err
			}
			this.ShowOutlierPoints = valueForShowOutlierPoints
		}
	}
	if valShowOutlierPointsCap, ok := objMap["ShowOutlierPoints"]; ok {
		if valShowOutlierPointsCap != nil {
			var valueForShowOutlierPoints bool
			err = json.Unmarshal(*valShowOutlierPointsCap, &valueForShowOutlierPoints)
			if err != nil {
				return err
			}
			this.ShowOutlierPoints = valueForShowOutlierPoints
		}
	}

	return nil
}
