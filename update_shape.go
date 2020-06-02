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

// Update shape task.
type IUpdateShape interface {

	// Task type.
	getType() string
	setType(newValue string)

	// Shape DTO.
	getShape() IShapeBase
	setShape(newValue IShapeBase)

	// Shape path for a grouped or SmartArt shape.
	getShapePath() string
	setShapePath(newValue string)
}

type UpdateShape struct {

	// Task type.
	Type_ string `json:"Type"`

	// Shape DTO.
	Shape IShapeBase `json:"Shape,omitempty"`

	// Shape path for a grouped or SmartArt shape.
	ShapePath string `json:"ShapePath,omitempty"`
}

func NewUpdateShape() *UpdateShape {
	instance := new(UpdateShape)
	instance.Type_ = "UpdateShape"
	return instance
}

func (this *UpdateShape) getType() string {
	return this.Type_
}

func (this *UpdateShape) setType(newValue string) {
	this.Type_ = newValue
}
func (this *UpdateShape) getShape() IShapeBase {
	return this.Shape
}

func (this *UpdateShape) setShape(newValue IShapeBase) {
	this.Shape = newValue
}
func (this *UpdateShape) getShapePath() string {
	return this.ShapePath
}

func (this *UpdateShape) setShapePath(newValue string) {
	this.ShapePath = newValue
}

func (this *UpdateShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "UpdateShape"
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
	
	if valShape, ok := objMap["shape"]; ok {
		if valShape != nil {
			var valueForShape ShapeBase
			err = json.Unmarshal(*valShape, &valueForShape)
			if err != nil {
				return err
			}
			this.Shape = &valueForShape
		}
	}
	if valShapeCap, ok := objMap["Shape"]; ok {
		if valShapeCap != nil {
			var valueForShape ShapeBase
			err = json.Unmarshal(*valShapeCap, &valueForShape)
			if err != nil {
				return err
			}
			this.Shape = &valueForShape
		}
	}
	
	if valShapePath, ok := objMap["shapePath"]; ok {
		if valShapePath != nil {
			var valueForShapePath string
			err = json.Unmarshal(*valShapePath, &valueForShapePath)
			if err != nil {
				return err
			}
			this.ShapePath = valueForShapePath
		}
	}
	if valShapePathCap, ok := objMap["ShapePath"]; ok {
		if valShapePathCap != nil {
			var valueForShapePath string
			err = json.Unmarshal(*valShapePathCap, &valueForShapePath)
			if err != nil {
				return err
			}
			this.ShapePath = valueForShapePath
		}
	}

    return nil
}
