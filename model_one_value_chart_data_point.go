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

	// Gets or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Gets or sets the marker.
	GetMarker() ISeriesMarker
	SetMarker(newValue ISeriesMarker)

	// Data point type.
	GetType() string
	SetType(newValue string)

	// Value.
	GetValue() float64
	SetValue(newValue float64)

	// Spreadsheet formula in A1-style.
	GetValueFormula() string
	SetValueFormula(newValue string)

	// SetAsTotal. Applied to Waterfall data points only.
	GetSetAsTotal() *bool
	SetSetAsTotal(newValue *bool)

	// True if the data point shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	GetInvertIfNegative() *bool
	SetInvertIfNegative(newValue *bool)
}

type OneValueChartDataPoint struct {

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the 3D format
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Gets or sets the marker.
	Marker ISeriesMarker `json:"Marker,omitempty"`

	// Data point type.
	Type_ string `json:"Type"`

	// Value.
	Value float64 `json:"Value,omitempty"`

	// Spreadsheet formula in A1-style.
	ValueFormula string `json:"ValueFormula,omitempty"`

	// SetAsTotal. Applied to Waterfall data points only.
	SetAsTotal *bool `json:"SetAsTotal"`

	// True if the data point shall invert its colors if the value is negative. Applies to bar, column and bubble series.
	InvertIfNegative *bool `json:"InvertIfNegative"`
}

func NewOneValueChartDataPoint() *OneValueChartDataPoint {
	instance := new(OneValueChartDataPoint)
	instance.Type_ = "OneValue"
	return instance
}

func (this *OneValueChartDataPoint) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *OneValueChartDataPoint) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *OneValueChartDataPoint) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *OneValueChartDataPoint) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *OneValueChartDataPoint) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *OneValueChartDataPoint) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *OneValueChartDataPoint) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *OneValueChartDataPoint) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *OneValueChartDataPoint) GetMarker() ISeriesMarker {
	return this.Marker
}

func (this *OneValueChartDataPoint) SetMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this *OneValueChartDataPoint) GetType() string {
	return this.Type_
}

func (this *OneValueChartDataPoint) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *OneValueChartDataPoint) GetValue() float64 {
	return this.Value
}

func (this *OneValueChartDataPoint) SetValue(newValue float64) {
	this.Value = newValue
}
func (this *OneValueChartDataPoint) GetValueFormula() string {
	return this.ValueFormula
}

func (this *OneValueChartDataPoint) SetValueFormula(newValue string) {
	this.ValueFormula = newValue
}
func (this *OneValueChartDataPoint) GetSetAsTotal() *bool {
	return this.SetAsTotal
}

func (this *OneValueChartDataPoint) SetSetAsTotal(newValue *bool) {
	this.SetAsTotal = newValue
}
func (this *OneValueChartDataPoint) GetInvertIfNegative() *bool {
	return this.InvertIfNegative
}

func (this *OneValueChartDataPoint) SetInvertIfNegative(newValue *bool) {
	this.InvertIfNegative = newValue
}

func (this *OneValueChartDataPoint) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
	
	if valThreeDFormat, ok := GetMapValue(objMap, "threeDFormat"); ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
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
	this.Type_ = "OneValue"
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
	
	if valValue, ok := GetMapValue(objMap, "value"); ok {
		if valValue != nil {
			var valueForValue float64
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	
	if valValueFormula, ok := GetMapValue(objMap, "valueFormula"); ok {
		if valValueFormula != nil {
			var valueForValueFormula string
			err = json.Unmarshal(*valValueFormula, &valueForValueFormula)
			if err != nil {
				return err
			}
			this.ValueFormula = valueForValueFormula
		}
	}
	
	if valSetAsTotal, ok := GetMapValue(objMap, "setAsTotal"); ok {
		if valSetAsTotal != nil {
			var valueForSetAsTotal *bool
			err = json.Unmarshal(*valSetAsTotal, &valueForSetAsTotal)
			if err != nil {
				return err
			}
			this.SetAsTotal = valueForSetAsTotal
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

	return nil
}
