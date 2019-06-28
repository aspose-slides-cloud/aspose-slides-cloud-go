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


type IUpdateShape interface {

	getType() string
	setType(newValue string)

	getShape() IShapeBase
	setShape(newValue IShapeBase)

	getShapePath() string
	setShapePath(newValue string)
}

type UpdateShape struct {

	Type_ string `json:"Type"`

	Shape IShapeBase `json:"Shape,omitempty"`

	ShapePath string `json:"ShapePath,omitempty"`
}

func (this UpdateShape) getType() string {
	return this.Type_
}

func (this UpdateShape) setType(newValue string) {
	this.Type_ = newValue
}
func (this UpdateShape) getShape() IShapeBase {
	return this.Shape
}

func (this UpdateShape) setShape(newValue IShapeBase) {
	this.Shape = newValue
}
func (this UpdateShape) getShapePath() string {
	return this.ShapePath
}

func (this UpdateShape) setShapePath(newValue string) {
	this.ShapePath = newValue
}

func (this *UpdateShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Save"
	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}
	
	if valShape, ok := objMap["Shape"]; ok {
		if valShape != nil {
			var valueForShape ShapeBase
			err = json.Unmarshal(*valShape, &valueForShape)
			if err != nil {
				return err
			}
			this.Shape = valueForShape
		}
	}
	
	if valShapePath, ok := objMap["ShapePath"]; ok {
		if valShapePath != nil {
			var valueForShapePath string
			err = json.Unmarshal(*valShapePath, &valueForShapePath)
			if err != nil {
				return err
			}
			this.ShapePath = valueForShapePath
		}
	}

    return nil
}
