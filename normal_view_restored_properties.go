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

// The sizing of the slide region.
type INormalViewRestoredProperties interface {

	// True if the size of the side content region should compensate for the new size when resizing the window containing the view within the application.
	getAutoAdjust() bool
	setAutoAdjust(newValue bool)

	// The size of the slide region.
	getDimensionSize() float64
	setDimensionSize(newValue float64)
}

type NormalViewRestoredProperties struct {

	// True if the size of the side content region should compensate for the new size when resizing the window containing the view within the application.
	AutoAdjust bool `json:"AutoAdjust"`

	// The size of the slide region.
	DimensionSize float64 `json:"DimensionSize,omitempty"`
}

func NewNormalViewRestoredProperties() *NormalViewRestoredProperties {
	instance := new(NormalViewRestoredProperties)
	return instance
}

func (this *NormalViewRestoredProperties) getAutoAdjust() bool {
	return this.AutoAdjust
}

func (this *NormalViewRestoredProperties) setAutoAdjust(newValue bool) {
	this.AutoAdjust = newValue
}
func (this *NormalViewRestoredProperties) getDimensionSize() float64 {
	return this.DimensionSize
}

func (this *NormalViewRestoredProperties) setDimensionSize(newValue float64) {
	this.DimensionSize = newValue
}

func (this *NormalViewRestoredProperties) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valAutoAdjust, ok := objMap["autoAdjust"]; ok {
		if valAutoAdjust != nil {
			var valueForAutoAdjust bool
			err = json.Unmarshal(*valAutoAdjust, &valueForAutoAdjust)
			if err != nil {
				return err
			}
			this.AutoAdjust = valueForAutoAdjust
		}
	}
	if valAutoAdjustCap, ok := objMap["AutoAdjust"]; ok {
		if valAutoAdjustCap != nil {
			var valueForAutoAdjust bool
			err = json.Unmarshal(*valAutoAdjustCap, &valueForAutoAdjust)
			if err != nil {
				return err
			}
			this.AutoAdjust = valueForAutoAdjust
		}
	}
	
	if valDimensionSize, ok := objMap["dimensionSize"]; ok {
		if valDimensionSize != nil {
			var valueForDimensionSize float64
			err = json.Unmarshal(*valDimensionSize, &valueForDimensionSize)
			if err != nil {
				return err
			}
			this.DimensionSize = valueForDimensionSize
		}
	}
	if valDimensionSizeCap, ok := objMap["DimensionSize"]; ok {
		if valDimensionSizeCap != nil {
			var valueForDimensionSize float64
			err = json.Unmarshal(*valDimensionSizeCap, &valueForDimensionSize)
			if err != nil {
				return err
			}
			this.DimensionSize = valueForDimensionSize
		}
	}

    return nil
}
