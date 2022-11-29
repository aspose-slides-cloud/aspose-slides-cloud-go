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
	GetType() string
	SetType(newValue string)

	// Series name.
	GetName() string
	SetName(newValue string)

	// Series name data source.
	GetDataSourceForSeriesName() IDataSource
	SetDataSourceForSeriesName(newValue IDataSource)

	// True if each data marker in the series has a different color.
	GetIsColorVaried() bool
	SetIsColorVaried(newValue bool)

	// Invert solid color for the series.
	GetInvertedSolidFillColor() string
	SetInvertedSolidFillColor(newValue string)

	// True if curve smoothing is turned on. Applies only to line and scatter connected by lines charts.
	GetSmooth() bool
	SetSmooth(newValue bool)

	// True if the series is plotted on second value axis.
	GetPlotOnSecondAxis() bool
	SetPlotOnSecondAxis(newValue bool)

	// Series order.
	GetOrder() int32
	SetOrder(newValue int32)

	// True if the series shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	GetInvertIfNegative() bool
	SetInvertIfNegative(newValue bool)

	// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
	GetExplosion() int32
	SetExplosion(newValue int32)

	// Series marker.
	GetMarker() ISeriesMarker
	SetMarker(newValue ISeriesMarker)

	// Fill properties set for the series.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Effect properties set for the series.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Line properties set for the series.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Data point type.
	GetDataPointType() string
	SetDataPointType(newValue string)

	// Gets or sets the values.
	GetDataPoints() []IOneValueChartDataPoint
	SetDataPoints(newValue []IOneValueChartDataPoint)

	// The number format for the series values.
	GetNumberFormatOfValues() string
	SetNumberFormatOfValues(newValue string)

	// Data source type for values.
	GetDataSourceForValues() IDataSource
	SetDataSourceForValues(newValue IDataSource)

	// True if inner points are shown. Applied to Waterfall series only.
	GetShowConnectorLines() bool
	SetShowConnectorLines(newValue bool)

	// Quartile method. Applied to BoxAndWhisker series only.
	GetQuartileMethod() string
	SetQuartileMethod(newValue string)

	// True if inner points are shown. Applied to BoxAndWhisker series only.
	GetShowInnerPoints() bool
	SetShowInnerPoints(newValue bool)

	// True if mean line is shown. Applied to BoxAndWhisker series only.
	GetShowMeanLine() bool
	SetShowMeanLine(newValue bool)

	// True if mean markers are shown. Applied to BoxAndWhisker series only.
	GetShowMeanMarkers() bool
	SetShowMeanMarkers(newValue bool)

	// True if outlier points are shown. Applied to BoxAndWhisker series only.
	GetShowOutlierPoints() bool
	SetShowOutlierPoints(newValue bool)
}

type OneValueSeries struct {

	// Series type.
	Type_ string `json:"Type,omitempty"`

	// Series name.
	Name string `json:"Name,omitempty"`

	// Series name data source.
	DataSourceForSeriesName IDataSource `json:"DataSourceForSeriesName,omitempty"`

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

	// Data source type for values.
	DataSourceForValues IDataSource `json:"DataSourceForValues,omitempty"`

	// True if inner points are shown. Applied to Waterfall series only.
	ShowConnectorLines bool `json:"ShowConnectorLines"`

	// Quartile method. Applied to BoxAndWhisker series only.
	QuartileMethod string `json:"QuartileMethod,omitempty"`

	// True if inner points are shown. Applied to BoxAndWhisker series only.
	ShowInnerPoints bool `json:"ShowInnerPoints"`

	// True if mean line is shown. Applied to BoxAndWhisker series only.
	ShowMeanLine bool `json:"ShowMeanLine"`

	// True if mean markers are shown. Applied to BoxAndWhisker series only.
	ShowMeanMarkers bool `json:"ShowMeanMarkers"`

	// True if outlier points are shown. Applied to BoxAndWhisker series only.
	ShowOutlierPoints bool `json:"ShowOutlierPoints"`
}

func NewOneValueSeries() *OneValueSeries {
	instance := new(OneValueSeries)
	instance.DataPointType = "OneValue"
	return instance
}

func (this *OneValueSeries) GetType() string {
	return this.Type_
}

func (this *OneValueSeries) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *OneValueSeries) GetName() string {
	return this.Name
}

func (this *OneValueSeries) SetName(newValue string) {
	this.Name = newValue
}
func (this *OneValueSeries) GetDataSourceForSeriesName() IDataSource {
	return this.DataSourceForSeriesName
}

func (this *OneValueSeries) SetDataSourceForSeriesName(newValue IDataSource) {
	this.DataSourceForSeriesName = newValue
}
func (this *OneValueSeries) GetIsColorVaried() bool {
	return this.IsColorVaried
}

func (this *OneValueSeries) SetIsColorVaried(newValue bool) {
	this.IsColorVaried = newValue
}
func (this *OneValueSeries) GetInvertedSolidFillColor() string {
	return this.InvertedSolidFillColor
}

func (this *OneValueSeries) SetInvertedSolidFillColor(newValue string) {
	this.InvertedSolidFillColor = newValue
}
func (this *OneValueSeries) GetSmooth() bool {
	return this.Smooth
}

func (this *OneValueSeries) SetSmooth(newValue bool) {
	this.Smooth = newValue
}
func (this *OneValueSeries) GetPlotOnSecondAxis() bool {
	return this.PlotOnSecondAxis
}

func (this *OneValueSeries) SetPlotOnSecondAxis(newValue bool) {
	this.PlotOnSecondAxis = newValue
}
func (this *OneValueSeries) GetOrder() int32 {
	return this.Order
}

func (this *OneValueSeries) SetOrder(newValue int32) {
	this.Order = newValue
}
func (this *OneValueSeries) GetInvertIfNegative() bool {
	return this.InvertIfNegative
}

func (this *OneValueSeries) SetInvertIfNegative(newValue bool) {
	this.InvertIfNegative = newValue
}
func (this *OneValueSeries) GetExplosion() int32 {
	return this.Explosion
}

func (this *OneValueSeries) SetExplosion(newValue int32) {
	this.Explosion = newValue
}
func (this *OneValueSeries) GetMarker() ISeriesMarker {
	return this.Marker
}

func (this *OneValueSeries) SetMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this *OneValueSeries) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *OneValueSeries) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *OneValueSeries) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *OneValueSeries) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *OneValueSeries) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *OneValueSeries) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *OneValueSeries) GetDataPointType() string {
	return this.DataPointType
}

func (this *OneValueSeries) SetDataPointType(newValue string) {
	this.DataPointType = newValue
}
func (this *OneValueSeries) GetDataPoints() []IOneValueChartDataPoint {
	return this.DataPoints
}

func (this *OneValueSeries) SetDataPoints(newValue []IOneValueChartDataPoint) {
	this.DataPoints = newValue
}
func (this *OneValueSeries) GetNumberFormatOfValues() string {
	return this.NumberFormatOfValues
}

func (this *OneValueSeries) SetNumberFormatOfValues(newValue string) {
	this.NumberFormatOfValues = newValue
}
func (this *OneValueSeries) GetDataSourceForValues() IDataSource {
	return this.DataSourceForValues
}

func (this *OneValueSeries) SetDataSourceForValues(newValue IDataSource) {
	this.DataSourceForValues = newValue
}
func (this *OneValueSeries) GetShowConnectorLines() bool {
	return this.ShowConnectorLines
}

func (this *OneValueSeries) SetShowConnectorLines(newValue bool) {
	this.ShowConnectorLines = newValue
}
func (this *OneValueSeries) GetQuartileMethod() string {
	return this.QuartileMethod
}

func (this *OneValueSeries) SetQuartileMethod(newValue string) {
	this.QuartileMethod = newValue
}
func (this *OneValueSeries) GetShowInnerPoints() bool {
	return this.ShowInnerPoints
}

func (this *OneValueSeries) SetShowInnerPoints(newValue bool) {
	this.ShowInnerPoints = newValue
}
func (this *OneValueSeries) GetShowMeanLine() bool {
	return this.ShowMeanLine
}

func (this *OneValueSeries) SetShowMeanLine(newValue bool) {
	this.ShowMeanLine = newValue
}
func (this *OneValueSeries) GetShowMeanMarkers() bool {
	return this.ShowMeanMarkers
}

func (this *OneValueSeries) SetShowMeanMarkers(newValue bool) {
	this.ShowMeanMarkers = newValue
}
func (this *OneValueSeries) GetShowOutlierPoints() bool {
	return this.ShowOutlierPoints
}

func (this *OneValueSeries) SetShowOutlierPoints(newValue bool) {
	this.ShowOutlierPoints = newValue
}

func (this *OneValueSeries) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
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
	
	if valDataSourceForSeriesName, ok := objMap["dataSourceForSeriesName"]; ok {
		if valDataSourceForSeriesName != nil {
			var valueForDataSourceForSeriesName DataSource
			err = json.Unmarshal(*valDataSourceForSeriesName, &valueForDataSourceForSeriesName)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForSeriesName)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForSeriesName, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForSeriesName = vInterfaceObject
			}
		}
	}
	if valDataSourceForSeriesNameCap, ok := objMap["DataSourceForSeriesName"]; ok {
		if valDataSourceForSeriesNameCap != nil {
			var valueForDataSourceForSeriesName DataSource
			err = json.Unmarshal(*valDataSourceForSeriesNameCap, &valueForDataSourceForSeriesName)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForSeriesNameCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForSeriesNameCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForSeriesName = vInterfaceObject
			}
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
	
	if valDataSourceForValues, ok := objMap["dataSourceForValues"]; ok {
		if valDataSourceForValues != nil {
			var valueForDataSourceForValues DataSource
			err = json.Unmarshal(*valDataSourceForValues, &valueForDataSourceForValues)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForValues)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForValues, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForValues = vInterfaceObject
			}
		}
	}
	if valDataSourceForValuesCap, ok := objMap["DataSourceForValues"]; ok {
		if valDataSourceForValuesCap != nil {
			var valueForDataSourceForValues DataSource
			err = json.Unmarshal(*valDataSourceForValuesCap, &valueForDataSourceForValues)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForValuesCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForValuesCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForValues = vInterfaceObject
			}
		}
	}
	
	if valShowConnectorLines, ok := objMap["showConnectorLines"]; ok {
		if valShowConnectorLines != nil {
			var valueForShowConnectorLines bool
			err = json.Unmarshal(*valShowConnectorLines, &valueForShowConnectorLines)
			if err != nil {
				return err
			}
			this.ShowConnectorLines = valueForShowConnectorLines
		}
	}
	if valShowConnectorLinesCap, ok := objMap["ShowConnectorLines"]; ok {
		if valShowConnectorLinesCap != nil {
			var valueForShowConnectorLines bool
			err = json.Unmarshal(*valShowConnectorLinesCap, &valueForShowConnectorLines)
			if err != nil {
				return err
			}
			this.ShowConnectorLines = valueForShowConnectorLines
		}
	}
	
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
