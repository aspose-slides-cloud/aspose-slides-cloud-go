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

// Represents chart category resource
type IChartCategory interface {

	// Gets or sets the parent categories. Used with Sunburst &amp; treemap categories; ignored for other chart types.
	GetParentCategories() []string
	SetParentCategories(newValue []string)

	// Gets or sets the grouping level for the category. Used with Sunburst &amp; treemap categories; ignored for other chart types.
	GetLevel() int32
	SetLevel(newValue int32)

	// Category value
	GetValue() string
	SetValue(newValue string)

	// Get or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Gets or sets the data points for chart data
	GetDataPoints() []IOneValueChartDataPoint
	SetDataPoints(newValue []IOneValueChartDataPoint)
}

type ChartCategory struct {

	// Gets or sets the parent categories. Used with Sunburst &amp; treemap categories; ignored for other chart types.
	ParentCategories []string `json:"ParentCategories,omitempty"`

	// Gets or sets the grouping level for the category. Used with Sunburst &amp; treemap categories; ignored for other chart types.
	Level int32 `json:"Level,omitempty"`

	// Category value
	Value string `json:"Value,omitempty"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Gets or sets the data points for chart data
	DataPoints []IOneValueChartDataPoint `json:"DataPoints,omitempty"`
}

func NewChartCategory() *ChartCategory {
	instance := new(ChartCategory)
	return instance
}

func (this *ChartCategory) GetParentCategories() []string {
	return this.ParentCategories
}

func (this *ChartCategory) SetParentCategories(newValue []string) {
	this.ParentCategories = newValue
}
func (this *ChartCategory) GetLevel() int32 {
	return this.Level
}

func (this *ChartCategory) SetLevel(newValue int32) {
	this.Level = newValue
}
func (this *ChartCategory) GetValue() string {
	return this.Value
}

func (this *ChartCategory) SetValue(newValue string) {
	this.Value = newValue
}
func (this *ChartCategory) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ChartCategory) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ChartCategory) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ChartCategory) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ChartCategory) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ChartCategory) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *ChartCategory) GetDataPoints() []IOneValueChartDataPoint {
	return this.DataPoints
}

func (this *ChartCategory) SetDataPoints(newValue []IOneValueChartDataPoint) {
	this.DataPoints = newValue
}

func (this *ChartCategory) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valParentCategories, ok := GetMapValue(objMap, "parentCategories"); ok {
		if valParentCategories != nil {
			var valueForParentCategories []string
			err = json.Unmarshal(*valParentCategories, &valueForParentCategories)
			if err != nil {
				return err
			}
			this.ParentCategories = valueForParentCategories
		}
	}
	
	if valLevel, ok := GetMapValue(objMap, "level"); ok {
		if valLevel != nil {
			var valueForLevel int32
			err = json.Unmarshal(*valLevel, &valueForLevel)
			if err != nil {
				return err
			}
			this.Level = valueForLevel
		}
	}
	
	if valValue, ok := GetMapValue(objMap, "value"); ok {
		if valValue != nil {
			var valueForValue string
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
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
	
	if valDataPoints, ok := GetMapValue(objMap, "dataPoints"); ok {
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

	return nil
}
