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

// Add shape task.
type IAddShape interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Shape DTO.
	GetShape() IShapeBase
	SetShape(newValue IShapeBase)

	// Shape path for a grouped shape or smart art shape.
	GetShapePath() string
	SetShapePath(newValue string)
}

type AddShape struct {

	// Task type.
	Type_ string `json:"Type"`

	// Shape DTO.
	Shape IShapeBase `json:"Shape,omitempty"`

	// Shape path for a grouped shape or smart art shape.
	ShapePath string `json:"ShapePath,omitempty"`
}

func NewAddShape() *AddShape {
	instance := new(AddShape)
	instance.Type_ = "AddShape"
	return instance
}

func (this *AddShape) GetType() string {
	return this.Type_
}

func (this *AddShape) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AddShape) GetShape() IShapeBase {
	return this.Shape
}

func (this *AddShape) SetShape(newValue IShapeBase) {
	this.Shape = newValue
}
func (this *AddShape) GetShapePath() string {
	return this.ShapePath
}

func (this *AddShape) SetShapePath(newValue string) {
	this.ShapePath = newValue
}

func (this *AddShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AddShape"
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
	
	if valShape, ok := GetMapValue(objMap, "shape"); ok {
		if valShape != nil {
			var valueForShape ShapeBase
			err = json.Unmarshal(*valShape, &valueForShape)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ShapeBase", *valShape)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valShape, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IShapeBase)
			if ok {
				this.Shape = vInterfaceObject
			}
		}
	}
	
	if valShapePath, ok := GetMapValue(objMap, "shapePath"); ok {
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
