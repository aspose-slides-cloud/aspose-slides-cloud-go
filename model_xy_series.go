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

// Common properties for Bubble and Scatter series. 
type IXYSeries interface {

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
	GetIsColorVaried() *bool
	SetIsColorVaried(newValue *bool)

	// Invert solid color for the series.
	GetInvertedSolidFillColor() string
	SetInvertedSolidFillColor(newValue string)

	// True if curve smoothing is turned on. Applies only to line and scatter connected by lines charts.
	GetSmooth() *bool
	SetSmooth(newValue *bool)

	// True if the series is plotted on second value axis.
	GetPlotOnSecondAxis() *bool
	SetPlotOnSecondAxis(newValue *bool)

	// Series order.
	GetOrder() int32
	SetOrder(newValue int32)

	// True if the series shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	GetInvertIfNegative() *bool
	SetInvertIfNegative(newValue *bool)

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

	GetDataPointType() string
	SetDataPointType(newValue string)

	// The number format for the series y values.
	GetNumberFormatOfYValues() string
	SetNumberFormatOfYValues(newValue string)

	// The number format for the series x values.
	GetNumberFormatOfXValues() string
	SetNumberFormatOfXValues(newValue string)

	// Data source type for X Values.
	GetDataSourceForXValues() IDataSource
	SetDataSourceForXValues(newValue IDataSource)

	// Data source type for Y Values.
	GetDataSourceForYValues() IDataSource
	SetDataSourceForYValues(newValue IDataSource)
}

type XYSeries struct {

	// Series type.
	Type_ string `json:"Type,omitempty"`

	// Series name.
	Name string `json:"Name,omitempty"`

	// Series name data source.
	DataSourceForSeriesName IDataSource `json:"DataSourceForSeriesName,omitempty"`

	// True if each data marker in the series has a different color.
	IsColorVaried *bool `json:"IsColorVaried"`

	// Invert solid color for the series.
	InvertedSolidFillColor string `json:"InvertedSolidFillColor,omitempty"`

	// True if curve smoothing is turned on. Applies only to line and scatter connected by lines charts.
	Smooth *bool `json:"Smooth"`

	// True if the series is plotted on second value axis.
	PlotOnSecondAxis *bool `json:"PlotOnSecondAxis"`

	// Series order.
	Order int32 `json:"Order,omitempty"`

	// True if the series shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	InvertIfNegative *bool `json:"InvertIfNegative"`

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

	DataPointType string `json:"DataPointType,omitempty"`

	// The number format for the series y values.
	NumberFormatOfYValues string `json:"NumberFormatOfYValues,omitempty"`

	// The number format for the series x values.
	NumberFormatOfXValues string `json:"NumberFormatOfXValues,omitempty"`

	// Data source type for X Values.
	DataSourceForXValues IDataSource `json:"DataSourceForXValues,omitempty"`

	// Data source type for Y Values.
	DataSourceForYValues IDataSource `json:"DataSourceForYValues,omitempty"`
}

func NewXYSeries() *XYSeries {
	instance := new(XYSeries)
	return instance
}

func (this *XYSeries) GetType() string {
	return this.Type_
}

func (this *XYSeries) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *XYSeries) GetName() string {
	return this.Name
}

func (this *XYSeries) SetName(newValue string) {
	this.Name = newValue
}
func (this *XYSeries) GetDataSourceForSeriesName() IDataSource {
	return this.DataSourceForSeriesName
}

func (this *XYSeries) SetDataSourceForSeriesName(newValue IDataSource) {
	this.DataSourceForSeriesName = newValue
}
func (this *XYSeries) GetIsColorVaried() *bool {
	return this.IsColorVaried
}

func (this *XYSeries) SetIsColorVaried(newValue *bool) {
	this.IsColorVaried = newValue
}
func (this *XYSeries) GetInvertedSolidFillColor() string {
	return this.InvertedSolidFillColor
}

func (this *XYSeries) SetInvertedSolidFillColor(newValue string) {
	this.InvertedSolidFillColor = newValue
}
func (this *XYSeries) GetSmooth() *bool {
	return this.Smooth
}

func (this *XYSeries) SetSmooth(newValue *bool) {
	this.Smooth = newValue
}
func (this *XYSeries) GetPlotOnSecondAxis() *bool {
	return this.PlotOnSecondAxis
}

func (this *XYSeries) SetPlotOnSecondAxis(newValue *bool) {
	this.PlotOnSecondAxis = newValue
}
func (this *XYSeries) GetOrder() int32 {
	return this.Order
}

func (this *XYSeries) SetOrder(newValue int32) {
	this.Order = newValue
}
func (this *XYSeries) GetInvertIfNegative() *bool {
	return this.InvertIfNegative
}

func (this *XYSeries) SetInvertIfNegative(newValue *bool) {
	this.InvertIfNegative = newValue
}
func (this *XYSeries) GetExplosion() int32 {
	return this.Explosion
}

func (this *XYSeries) SetExplosion(newValue int32) {
	this.Explosion = newValue
}
func (this *XYSeries) GetMarker() ISeriesMarker {
	return this.Marker
}

func (this *XYSeries) SetMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this *XYSeries) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *XYSeries) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *XYSeries) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *XYSeries) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *XYSeries) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *XYSeries) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *XYSeries) GetDataPointType() string {
	return this.DataPointType
}

func (this *XYSeries) SetDataPointType(newValue string) {
	this.DataPointType = newValue
}
func (this *XYSeries) GetNumberFormatOfYValues() string {
	return this.NumberFormatOfYValues
}

func (this *XYSeries) SetNumberFormatOfYValues(newValue string) {
	this.NumberFormatOfYValues = newValue
}
func (this *XYSeries) GetNumberFormatOfXValues() string {
	return this.NumberFormatOfXValues
}

func (this *XYSeries) SetNumberFormatOfXValues(newValue string) {
	this.NumberFormatOfXValues = newValue
}
func (this *XYSeries) GetDataSourceForXValues() IDataSource {
	return this.DataSourceForXValues
}

func (this *XYSeries) SetDataSourceForXValues(newValue IDataSource) {
	this.DataSourceForXValues = newValue
}
func (this *XYSeries) GetDataSourceForYValues() IDataSource {
	return this.DataSourceForYValues
}

func (this *XYSeries) SetDataSourceForYValues(newValue IDataSource) {
	this.DataSourceForYValues = newValue
}

func (this *XYSeries) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valDataSourceForSeriesName, ok := GetMapValue(objMap, "dataSourceForSeriesName"); ok {
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
	
	if valIsColorVaried, ok := GetMapValue(objMap, "isColorVaried"); ok {
		if valIsColorVaried != nil {
			var valueForIsColorVaried *bool
			err = json.Unmarshal(*valIsColorVaried, &valueForIsColorVaried)
			if err != nil {
				return err
			}
			this.IsColorVaried = valueForIsColorVaried
		}
	}
	
	if valInvertedSolidFillColor, ok := GetMapValue(objMap, "invertedSolidFillColor"); ok {
		if valInvertedSolidFillColor != nil {
			var valueForInvertedSolidFillColor string
			err = json.Unmarshal(*valInvertedSolidFillColor, &valueForInvertedSolidFillColor)
			if err != nil {
				return err
			}
			this.InvertedSolidFillColor = valueForInvertedSolidFillColor
		}
	}
	
	if valSmooth, ok := GetMapValue(objMap, "smooth"); ok {
		if valSmooth != nil {
			var valueForSmooth *bool
			err = json.Unmarshal(*valSmooth, &valueForSmooth)
			if err != nil {
				return err
			}
			this.Smooth = valueForSmooth
		}
	}
	
	if valPlotOnSecondAxis, ok := GetMapValue(objMap, "plotOnSecondAxis"); ok {
		if valPlotOnSecondAxis != nil {
			var valueForPlotOnSecondAxis *bool
			err = json.Unmarshal(*valPlotOnSecondAxis, &valueForPlotOnSecondAxis)
			if err != nil {
				return err
			}
			this.PlotOnSecondAxis = valueForPlotOnSecondAxis
		}
	}
	
	if valOrder, ok := GetMapValue(objMap, "order"); ok {
		if valOrder != nil {
			var valueForOrder int32
			err = json.Unmarshal(*valOrder, &valueForOrder)
			if err != nil {
				return err
			}
			this.Order = valueForOrder
		}
	}
	
	if valInvertIfNegative, ok := GetMapValue(objMap, "invertIfNegative"); ok {
		if valInvertIfNegative != nil {
			var valueForInvertIfNegative *bool
			err = json.Unmarshal(*valInvertIfNegative, &valueForInvertIfNegative)
			if err != nil {
				return err
			}
			this.InvertIfNegative = valueForInvertIfNegative
		}
	}
	
	if valExplosion, ok := GetMapValue(objMap, "explosion"); ok {
		if valExplosion != nil {
			var valueForExplosion int32
			err = json.Unmarshal(*valExplosion, &valueForExplosion)
			if err != nil {
				return err
			}
			this.Explosion = valueForExplosion
		}
	}
	
	if valMarker, ok := GetMapValue(objMap, "marker"); ok {
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
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
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
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
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
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
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
	
	if valDataPointType, ok := GetMapValue(objMap, "dataPointType"); ok {
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
	
	if valNumberFormatOfYValues, ok := GetMapValue(objMap, "numberFormatOfYValues"); ok {
		if valNumberFormatOfYValues != nil {
			var valueForNumberFormatOfYValues string
			err = json.Unmarshal(*valNumberFormatOfYValues, &valueForNumberFormatOfYValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfYValues = valueForNumberFormatOfYValues
		}
	}
	
	if valNumberFormatOfXValues, ok := GetMapValue(objMap, "numberFormatOfXValues"); ok {
		if valNumberFormatOfXValues != nil {
			var valueForNumberFormatOfXValues string
			err = json.Unmarshal(*valNumberFormatOfXValues, &valueForNumberFormatOfXValues)
			if err != nil {
				return err
			}
			this.NumberFormatOfXValues = valueForNumberFormatOfXValues
		}
	}
	
	if valDataSourceForXValues, ok := GetMapValue(objMap, "dataSourceForXValues"); ok {
		if valDataSourceForXValues != nil {
			var valueForDataSourceForXValues DataSource
			err = json.Unmarshal(*valDataSourceForXValues, &valueForDataSourceForXValues)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForXValues)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForXValues, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForXValues = vInterfaceObject
			}
		}
	}
	
	if valDataSourceForYValues, ok := GetMapValue(objMap, "dataSourceForYValues"); ok {
		if valDataSourceForYValues != nil {
			var valueForDataSourceForYValues DataSource
			err = json.Unmarshal(*valDataSourceForYValues, &valueForDataSourceForYValues)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForYValues)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForYValues, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForYValues = vInterfaceObject
			}
		}
	}

	return nil
}
