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
	getParentCategories() []string
	setParentCategories(newValue []string)

	// Gets or sets the grouping level for the category. Used with Sunburst &amp; treemap categories; ignored for other chart types.
	getLevel() int32
	setLevel(newValue int32)

	// Category value
	getValue() string
	setValue(newValue string)

	// Get or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Gets or sets the data points for chart data
	getDataPoints() []IOneValueChartDataPoint
	setDataPoints(newValue []IOneValueChartDataPoint)
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

func (this *ChartCategory) getParentCategories() []string {
	return this.ParentCategories
}

func (this *ChartCategory) setParentCategories(newValue []string) {
	this.ParentCategories = newValue
}
func (this *ChartCategory) getLevel() int32 {
	return this.Level
}

func (this *ChartCategory) setLevel(newValue int32) {
	this.Level = newValue
}
func (this *ChartCategory) getValue() string {
	return this.Value
}

func (this *ChartCategory) setValue(newValue string) {
	this.Value = newValue
}
func (this *ChartCategory) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ChartCategory) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ChartCategory) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ChartCategory) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ChartCategory) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ChartCategory) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *ChartCategory) getDataPoints() []IOneValueChartDataPoint {
	return this.DataPoints
}

func (this *ChartCategory) setDataPoints(newValue []IOneValueChartDataPoint) {
	this.DataPoints = newValue
}

func (this *ChartCategory) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valParentCategories, ok := objMap["parentCategories"]; ok {
		if valParentCategories != nil {
			var valueForParentCategories []string
			err = json.Unmarshal(*valParentCategories, &valueForParentCategories)
			if err != nil {
				return err
			}
			this.ParentCategories = valueForParentCategories
		}
	}
	if valParentCategoriesCap, ok := objMap["ParentCategories"]; ok {
		if valParentCategoriesCap != nil {
			var valueForParentCategories []string
			err = json.Unmarshal(*valParentCategoriesCap, &valueForParentCategories)
			if err != nil {
				return err
			}
			this.ParentCategories = valueForParentCategories
		}
	}
	
	if valLevel, ok := objMap["level"]; ok {
		if valLevel != nil {
			var valueForLevel int32
			err = json.Unmarshal(*valLevel, &valueForLevel)
			if err != nil {
				return err
			}
			this.Level = valueForLevel
		}
	}
	if valLevelCap, ok := objMap["Level"]; ok {
		if valLevelCap != nil {
			var valueForLevel int32
			err = json.Unmarshal(*valLevelCap, &valueForLevel)
			if err != nil {
				return err
			}
			this.Level = valueForLevel
		}
	}
	
	if valValue, ok := objMap["value"]; ok {
		if valValue != nil {
			var valueForValue string
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	if valValueCap, ok := objMap["Value"]; ok {
		if valValueCap != nil {
			var valueForValue string
			err = json.Unmarshal(*valValueCap, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	
	if valDataPoints, ok := objMap["dataPoints"]; ok {
		if valDataPoints != nil {
			var valueForDataPoints []OneValueChartDataPoint
			err = json.Unmarshal(*valDataPoints, &valueForDataPoints)
			if err != nil {
				return err
			}
			valueForIDataPoints := make([]IOneValueChartDataPoint, len(valueForDataPoints))
			for i, v := range valueForDataPoints {
				valueForIDataPoints[i] = IOneValueChartDataPoint(&v)
			}
			this.DataPoints = valueForIDataPoints
		}
	}
	if valDataPointsCap, ok := objMap["DataPoints"]; ok {
		if valDataPointsCap != nil {
			var valueForDataPoints []OneValueChartDataPoint
			err = json.Unmarshal(*valDataPointsCap, &valueForDataPoints)
			if err != nil {
				return err
			}
			valueForIDataPoints := make([]IOneValueChartDataPoint, len(valueForDataPoints))
			for i, v := range valueForDataPoints {
				valueForIDataPoints[i] = IOneValueChartDataPoint(&v)
			}
			this.DataPoints = valueForIDataPoints
		}
	}

    return nil
}
