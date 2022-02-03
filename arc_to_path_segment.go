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
	getType() string
	setType(newValue string)

	// Width of the rectangle
	getWidth() float64
	setWidth(newValue float64)

	// Height of the rectangle
	getHeight() float64
	setHeight(newValue float64)

	// Start angle
	getStartAngle() float64
	setStartAngle(newValue float64)

	// Sweep angle
	getSweepAngle() float64
	setSweepAngle(newValue float64)
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

func (this *ArcToPathSegment) getType() string {
	return this.Type_
}

func (this *ArcToPathSegment) setType(newValue string) {
	this.Type_ = newValue
}
func (this *ArcToPathSegment) getWidth() float64 {
	return this.Width
}

func (this *ArcToPathSegment) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *ArcToPathSegment) getHeight() float64 {
	return this.Height
}

func (this *ArcToPathSegment) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *ArcToPathSegment) getStartAngle() float64 {
	return this.StartAngle
}

func (this *ArcToPathSegment) setStartAngle(newValue float64) {
	this.StartAngle = newValue
}
func (this *ArcToPathSegment) getSweepAngle() float64 {
	return this.SweepAngle
}

func (this *ArcToPathSegment) setSweepAngle(newValue float64) {
	this.SweepAngle = newValue
}

func (this *ArcToPathSegment) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ArcTo"
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
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valStartAngle, ok := objMap["startAngle"]; ok {
		if valStartAngle != nil {
			var valueForStartAngle float64
			err = json.Unmarshal(*valStartAngle, &valueForStartAngle)
			if err != nil {
				return err
			}
			this.StartAngle = valueForStartAngle
		}
	}
	if valStartAngleCap, ok := objMap["StartAngle"]; ok {
		if valStartAngleCap != nil {
			var valueForStartAngle float64
			err = json.Unmarshal(*valStartAngleCap, &valueForStartAngle)
			if err != nil {
				return err
			}
			this.StartAngle = valueForStartAngle
		}
	}
	
	if valSweepAngle, ok := objMap["sweepAngle"]; ok {
		if valSweepAngle != nil {
			var valueForSweepAngle float64
			err = json.Unmarshal(*valSweepAngle, &valueForSweepAngle)
			if err != nil {
				return err
			}
			this.SweepAngle = valueForSweepAngle
		}
	}
	if valSweepAngleCap, ok := objMap["SweepAngle"]; ok {
		if valSweepAngleCap != nil {
			var valueForSweepAngle float64
			err = json.Unmarshal(*valSweepAngleCap, &valueForSweepAngle)
			if err != nil {
				return err
			}
			this.SweepAngle = valueForSweepAngle
		}
	}

	return nil
}
