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

// Represents an BiLevel effect.
type IBiLevelEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Returns effect threshold.
	GetThreshold() float64
	SetThreshold(newValue float64)
}

type BiLevelEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Returns effect threshold.
	Threshold float64 `json:"Threshold"`
}

func NewBiLevelEffect() *BiLevelEffect {
	instance := new(BiLevelEffect)
	instance.Type_ = "BiLevel"
	return instance
}

func (this *BiLevelEffect) GetType() string {
	return this.Type_
}

func (this *BiLevelEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *BiLevelEffect) GetThreshold() float64 {
	return this.Threshold
}

func (this *BiLevelEffect) SetThreshold(newValue float64) {
	this.Threshold = newValue
}

func (this *BiLevelEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "BiLevel"
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
	
	if valThreshold, ok := objMap["threshold"]; ok {
		if valThreshold != nil {
			var valueForThreshold float64
			err = json.Unmarshal(*valThreshold, &valueForThreshold)
			if err != nil {
				return err
			}
			this.Threshold = valueForThreshold
		}
	}
	if valThresholdCap, ok := objMap["Threshold"]; ok {
		if valThresholdCap != nil {
			var valueForThreshold float64
			err = json.Unmarshal(*valThresholdCap, &valueForThreshold)
			if err != nil {
				return err
			}
			this.Threshold = valueForThreshold
		}
	}

	return nil
}
