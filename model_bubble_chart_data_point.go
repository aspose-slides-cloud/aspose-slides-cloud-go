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

	// X-value
	GetXValue() float64
	SetXValue(newValue float64)

	// Y-value
	GetYValue() float64
	SetYValue(newValue float64)

	// Spreadsheet formula in A1-style.
	GetXValueFormula() string
	SetXValueFormula(newValue string)

	// Spreadsheet formula in A1-style.
	GetYValueFormula() string
	SetYValueFormula(newValue string)

	// Bubble size.
	GetBubbleSize() float64
	SetBubbleSize(newValue float64)

	// Spreadsheet formula in A1-style.
	GetBubbleSizeFormula() string
	SetBubbleSizeFormula(newValue string)
}

type BubbleChartDataPoint struct {

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

	// X-value
	XValue float64 `json:"XValue,omitempty"`

	// Y-value
	YValue float64 `json:"YValue,omitempty"`

	// Spreadsheet formula in A1-style.
	XValueFormula string `json:"XValueFormula,omitempty"`

	// Spreadsheet formula in A1-style.
	YValueFormula string `json:"YValueFormula,omitempty"`

	// Bubble size.
	BubbleSize float64 `json:"BubbleSize,omitempty"`

	// Spreadsheet formula in A1-style.
	BubbleSizeFormula string `json:"BubbleSizeFormula,omitempty"`
}

func NewBubbleChartDataPoint() *BubbleChartDataPoint {
	instance := new(BubbleChartDataPoint)
	instance.Type_ = "Bubble"
	return instance
}

func (this *BubbleChartDataPoint) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *BubbleChartDataPoint) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *BubbleChartDataPoint) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *BubbleChartDataPoint) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *BubbleChartDataPoint) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *BubbleChartDataPoint) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *BubbleChartDataPoint) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *BubbleChartDataPoint) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *BubbleChartDataPoint) GetMarker() ISeriesMarker {
	return this.Marker
}

func (this *BubbleChartDataPoint) SetMarker(newValue ISeriesMarker) {
	this.Marker = newValue
}
func (this *BubbleChartDataPoint) GetType() string {
	return this.Type_
}

func (this *BubbleChartDataPoint) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *BubbleChartDataPoint) GetXValue() float64 {
	return this.XValue
}

func (this *BubbleChartDataPoint) SetXValue(newValue float64) {
	this.XValue = newValue
}
func (this *BubbleChartDataPoint) GetYValue() float64 {
	return this.YValue
}

func (this *BubbleChartDataPoint) SetYValue(newValue float64) {
	this.YValue = newValue
}
func (this *BubbleChartDataPoint) GetXValueFormula() string {
	return this.XValueFormula
}

func (this *BubbleChartDataPoint) SetXValueFormula(newValue string) {
	this.XValueFormula = newValue
}
func (this *BubbleChartDataPoint) GetYValueFormula() string {
	return this.YValueFormula
}

func (this *BubbleChartDataPoint) SetYValueFormula(newValue string) {
	this.YValueFormula = newValue
}
func (this *BubbleChartDataPoint) GetBubbleSize() float64 {
	return this.BubbleSize
}

func (this *BubbleChartDataPoint) SetBubbleSize(newValue float64) {
	this.BubbleSize = newValue
}
func (this *BubbleChartDataPoint) GetBubbleSizeFormula() string {
	return this.BubbleSizeFormula
}

func (this *BubbleChartDataPoint) SetBubbleSizeFormula(newValue string) {
	this.BubbleSizeFormula = newValue
}

func (this *BubbleChartDataPoint) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "Bubble"
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
	
	if valXValue, ok := GetMapValue(objMap, "xValue"); ok {
		if valXValue != nil {
			var valueForXValue float64
			err = json.Unmarshal(*valXValue, &valueForXValue)
			if err != nil {
				return err
			}
			this.XValue = valueForXValue
		}
	}
	
	if valYValue, ok := GetMapValue(objMap, "yValue"); ok {
		if valYValue != nil {
			var valueForYValue float64
			err = json.Unmarshal(*valYValue, &valueForYValue)
			if err != nil {
				return err
			}
			this.YValue = valueForYValue
		}
	}
	
	if valXValueFormula, ok := GetMapValue(objMap, "xValueFormula"); ok {
		if valXValueFormula != nil {
			var valueForXValueFormula string
			err = json.Unmarshal(*valXValueFormula, &valueForXValueFormula)
			if err != nil {
				return err
			}
			this.XValueFormula = valueForXValueFormula
		}
	}
	
	if valYValueFormula, ok := GetMapValue(objMap, "yValueFormula"); ok {
		if valYValueFormula != nil {
			var valueForYValueFormula string
			err = json.Unmarshal(*valYValueFormula, &valueForYValueFormula)
			if err != nil {
				return err
			}
			this.YValueFormula = valueForYValueFormula
		}
	}
	
	if valBubbleSize, ok := GetMapValue(objMap, "bubbleSize"); ok {
		if valBubbleSize != nil {
			var valueForBubbleSize float64
			err = json.Unmarshal(*valBubbleSize, &valueForBubbleSize)
			if err != nil {
				return err
			}
			this.BubbleSize = valueForBubbleSize
		}
	}
	
	if valBubbleSizeFormula, ok := GetMapValue(objMap, "bubbleSizeFormula"); ok {
		if valBubbleSizeFormula != nil {
			var valueForBubbleSizeFormula string
			err = json.Unmarshal(*valBubbleSizeFormula, &valueForBubbleSizeFormula)
			if err != nil {
				return err
			}
			this.BubbleSizeFormula = valueForBubbleSizeFormula
		}
	}

	return nil
}
