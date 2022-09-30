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

	// X-value
	GetXValue() float64
	SetXValue(newValue float64)

	// Y-value
	GetYValue() float64
	SetYValue(newValue float64)

	// Bubble size.
	GetBubbleSize() float64
	SetBubbleSize(newValue float64)
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

	// X-value
	XValue float64 `json:"XValue"`

	// Y-value
	YValue float64 `json:"YValue"`

	// Bubble size.
	BubbleSize float64 `json:"BubbleSize"`
}

func NewBubbleChartDataPoint() *BubbleChartDataPoint {
	instance := new(BubbleChartDataPoint)
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
func (this *BubbleChartDataPoint) GetBubbleSize() float64 {
	return this.BubbleSize
}

func (this *BubbleChartDataPoint) SetBubbleSize(newValue float64) {
	this.BubbleSize = newValue
}

func (this *BubbleChartDataPoint) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
	
	if valThreeDFormat, ok := objMap["threeDFormat"]; ok {
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
	if valThreeDFormatCap, ok := objMap["ThreeDFormat"]; ok {
		if valThreeDFormatCap != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormatCap, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
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
