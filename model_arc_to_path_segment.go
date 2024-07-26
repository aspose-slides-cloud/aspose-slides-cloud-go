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

// Arc segment of the geometry path
type IArcToPathSegment interface {

	// Arc segment
	GetType() string
	SetType(newValue string)

	// Width of the rectangle
	GetWidth() float64
	SetWidth(newValue float64)

	// Height of the rectangle
	GetHeight() float64
	SetHeight(newValue float64)

	// Start angle
	GetStartAngle() float64
	SetStartAngle(newValue float64)

	// Sweep angle
	GetSweepAngle() float64
	SetSweepAngle(newValue float64)
}

type ArcToPathSegment struct {

	// Arc segment
	Type_ string `json:"Type"`

	// Width of the rectangle
	Width float64 `json:"Width"`

	// Height of the rectangle
	Height float64 `json:"Height"`

	// Start angle
	StartAngle float64 `json:"StartAngle"`

	// Sweep angle
	SweepAngle float64 `json:"SweepAngle"`
}

func NewArcToPathSegment() *ArcToPathSegment {
	instance := new(ArcToPathSegment)
	instance.Type_ = "ArcTo"
	return instance
}

func (this *ArcToPathSegment) GetType() string {
	return this.Type_
}

func (this *ArcToPathSegment) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ArcToPathSegment) GetWidth() float64 {
	return this.Width
}

func (this *ArcToPathSegment) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *ArcToPathSegment) GetHeight() float64 {
	return this.Height
}

func (this *ArcToPathSegment) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *ArcToPathSegment) GetStartAngle() float64 {
	return this.StartAngle
}

func (this *ArcToPathSegment) SetStartAngle(newValue float64) {
	this.StartAngle = newValue
}
func (this *ArcToPathSegment) GetSweepAngle() float64 {
	return this.SweepAngle
}

func (this *ArcToPathSegment) SetSweepAngle(newValue float64) {
	this.SweepAngle = newValue
}

func (this *ArcToPathSegment) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ArcTo"
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
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valStartAngle, ok := GetMapValue(objMap, "startAngle"); ok {
		if valStartAngle != nil {
			var valueForStartAngle float64
			err = json.Unmarshal(*valStartAngle, &valueForStartAngle)
			if err != nil {
				return err
			}
			this.StartAngle = valueForStartAngle
		}
	}
	
	if valSweepAngle, ok := GetMapValue(objMap, "sweepAngle"); ok {
		if valSweepAngle != nil {
			var valueForSweepAngle float64
			err = json.Unmarshal(*valSweepAngle, &valueForSweepAngle)
			if err != nil {
				return err
			}
			this.SweepAngle = valueForSweepAngle
		}
	}

	return nil
}
