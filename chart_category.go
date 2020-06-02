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

	// Gets or sets the categories for chart data
	getCategories() []IChartCategory
	setCategories(newValue []IChartCategory)

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
}

type ChartCategory struct {

	// Gets or sets the categories for chart data
	Categories []IChartCategory `json:"Categories,omitempty"`

	// Category value
	Value string `json:"Value,omitempty"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func NewChartCategory() *ChartCategory {
	instance := new(ChartCategory)
	return instance
}

func (this *ChartCategory) getCategories() []IChartCategory {
	return this.Categories
}

func (this *ChartCategory) setCategories(newValue []IChartCategory) {
	this.Categories = newValue
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

func (this *ChartCategory) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valCategories, ok := objMap["categories"]; ok {
		if valCategories != nil {
			var valueForCategories []ChartCategory
			err = json.Unmarshal(*valCategories, &valueForCategories)
			if err != nil {
				return err
			}
			valueForICategories := make([]IChartCategory, len(valueForCategories))
			for i, v := range valueForCategories {
				valueForICategories[i] = IChartCategory(&v)
			}
			this.Categories = valueForICategories
		}
	}
	if valCategoriesCap, ok := objMap["Categories"]; ok {
		if valCategoriesCap != nil {
			var valueForCategories []ChartCategory
			err = json.Unmarshal(*valCategoriesCap, &valueForCategories)
			if err != nil {
				return err
			}
			valueForICategories := make([]IChartCategory, len(valueForCategories))
			for i, v := range valueForCategories {
				valueForICategories[i] = IChartCategory(&v)
			}
			this.Categories = valueForICategories
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

    return nil
}
