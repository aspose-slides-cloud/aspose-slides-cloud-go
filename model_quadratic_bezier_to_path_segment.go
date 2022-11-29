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

// Quadratic Bezier curve segment of the geometry path
type IQuadraticBezierToPathSegment interface {

	// Quadratic Bezier curve segment
	GetType() string
	SetType(newValue string)

	// X coordinate of direction point
	GetX1() float64
	SetX1(newValue float64)

	// Y coordinate of direction point
	GetY1() float64
	SetY1(newValue float64)

	// X coordinate of end point
	GetX2() float64
	SetX2(newValue float64)

	// Y coordinate of end point
	GetY2() float64
	SetY2(newValue float64)
}

type QuadraticBezierToPathSegment struct {

	// Quadratic Bezier curve segment
	Type_ string `json:"Type"`

	// X coordinate of direction point
	X1 float64 `json:"X1"`

	// Y coordinate of direction point
	Y1 float64 `json:"Y1"`

	// X coordinate of end point
	X2 float64 `json:"X2"`

	// Y coordinate of end point
	Y2 float64 `json:"Y2"`
}

func NewQuadraticBezierToPathSegment() *QuadraticBezierToPathSegment {
	instance := new(QuadraticBezierToPathSegment)
	instance.Type_ = "QuadBezierTo"
	return instance
}

func (this *QuadraticBezierToPathSegment) GetType() string {
	return this.Type_
}

func (this *QuadraticBezierToPathSegment) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *QuadraticBezierToPathSegment) GetX1() float64 {
	return this.X1
}

func (this *QuadraticBezierToPathSegment) SetX1(newValue float64) {
	this.X1 = newValue
}
func (this *QuadraticBezierToPathSegment) GetY1() float64 {
	return this.Y1
}

func (this *QuadraticBezierToPathSegment) SetY1(newValue float64) {
	this.Y1 = newValue
}
func (this *QuadraticBezierToPathSegment) GetX2() float64 {
	return this.X2
}

func (this *QuadraticBezierToPathSegment) SetX2(newValue float64) {
	this.X2 = newValue
}
func (this *QuadraticBezierToPathSegment) GetY2() float64 {
	return this.Y2
}

func (this *QuadraticBezierToPathSegment) SetY2(newValue float64) {
	this.Y2 = newValue
}

func (this *QuadraticBezierToPathSegment) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "QuadBezierTo"
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
	
	if valX1, ok := objMap["x1"]; ok {
		if valX1 != nil {
			var valueForX1 float64
			err = json.Unmarshal(*valX1, &valueForX1)
			if err != nil {
				return err
			}
			this.X1 = valueForX1
		}
	}
	if valX1Cap, ok := objMap["X1"]; ok {
		if valX1Cap != nil {
			var valueForX1 float64
			err = json.Unmarshal(*valX1Cap, &valueForX1)
			if err != nil {
				return err
			}
			this.X1 = valueForX1
		}
	}
	
	if valY1, ok := objMap["y1"]; ok {
		if valY1 != nil {
			var valueForY1 float64
			err = json.Unmarshal(*valY1, &valueForY1)
			if err != nil {
				return err
			}
			this.Y1 = valueForY1
		}
	}
	if valY1Cap, ok := objMap["Y1"]; ok {
		if valY1Cap != nil {
			var valueForY1 float64
			err = json.Unmarshal(*valY1Cap, &valueForY1)
			if err != nil {
				return err
			}
			this.Y1 = valueForY1
		}
	}
	
	if valX2, ok := objMap["x2"]; ok {
		if valX2 != nil {
			var valueForX2 float64
			err = json.Unmarshal(*valX2, &valueForX2)
			if err != nil {
				return err
			}
			this.X2 = valueForX2
		}
	}
	if valX2Cap, ok := objMap["X2"]; ok {
		if valX2Cap != nil {
			var valueForX2 float64
			err = json.Unmarshal(*valX2Cap, &valueForX2)
			if err != nil {
				return err
			}
			this.X2 = valueForX2
		}
	}
	
	if valY2, ok := objMap["y2"]; ok {
		if valY2 != nil {
			var valueForY2 float64
			err = json.Unmarshal(*valY2, &valueForY2)
			if err != nil {
				return err
			}
			this.Y2 = valueForY2
		}
	}
	if valY2Cap, ok := objMap["Y2"]; ok {
		if valY2Cap != nil {
			var valueForY2 float64
			err = json.Unmarshal(*valY2Cap, &valueForY2)
			if err != nil {
				return err
			}
			this.Y2 = valueForY2
		}
	}

	return nil
}
