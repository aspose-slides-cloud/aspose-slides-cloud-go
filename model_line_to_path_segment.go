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

// Line segment of the geometry path
type ILineToPathSegment interface {

	// Line segment
	GetType() string
	SetType(newValue string)

	// X coordinate of the end point of the line
	GetX() float64
	SetX(newValue float64)

	// Y coordinate of the end point of the line
	GetY() float64
	SetY(newValue float64)
}

type LineToPathSegment struct {

	// Line segment
	Type_ string `json:"Type"`

	// X coordinate of the end point of the line
	X float64 `json:"X"`

	// Y coordinate of the end point of the line
	Y float64 `json:"Y"`
}

func NewLineToPathSegment() *LineToPathSegment {
	instance := new(LineToPathSegment)
	instance.Type_ = "LineTo"
	return instance
}

func (this *LineToPathSegment) GetType() string {
	return this.Type_
}

func (this *LineToPathSegment) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *LineToPathSegment) GetX() float64 {
	return this.X
}

func (this *LineToPathSegment) SetX(newValue float64) {
	this.X = newValue
}
func (this *LineToPathSegment) GetY() float64 {
	return this.Y
}

func (this *LineToPathSegment) SetY(newValue float64) {
	this.Y = newValue
}

func (this *LineToPathSegment) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "LineTo"
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
	
	if valX, ok := GetMapValue(objMap, "x"); ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := GetMapValue(objMap, "y"); ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}

	return nil
}
