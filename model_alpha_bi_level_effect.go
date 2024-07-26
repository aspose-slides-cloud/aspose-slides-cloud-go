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

// Represents an Alpha Bi-Level effect.
type IAlphaBiLevelEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Returns effect threshold.
	GetThreshold() float64
	SetThreshold(newValue float64)
}

type AlphaBiLevelEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Returns effect threshold.
	Threshold float64 `json:"Threshold"`
}

func NewAlphaBiLevelEffect() *AlphaBiLevelEffect {
	instance := new(AlphaBiLevelEffect)
	instance.Type_ = "AlphaBiLevel"
	return instance
}

func (this *AlphaBiLevelEffect) GetType() string {
	return this.Type_
}

func (this *AlphaBiLevelEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AlphaBiLevelEffect) GetThreshold() float64 {
	return this.Threshold
}

func (this *AlphaBiLevelEffect) SetThreshold(newValue float64) {
	this.Threshold = newValue
}

func (this *AlphaBiLevelEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AlphaBiLevel"
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
	
	if valThreshold, ok := GetMapValue(objMap, "threshold"); ok {
		if valThreshold != nil {
			var valueForThreshold float64
			err = json.Unmarshal(*valThreshold, &valueForThreshold)
			if err != nil {
				return err
			}
			this.Threshold = valueForThreshold
		}
	}

	return nil
}
