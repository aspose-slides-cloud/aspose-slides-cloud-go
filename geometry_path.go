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

// Represents GeometryPath of the shape
type IGeometryPath interface {

	// Path fill mode
	GetFillMode() string
	SetFillMode(newValue string)

	// Stroke
	GetStroke() bool
	SetStroke(newValue bool)

	// List of PathSegmen objects
	GetPathData() []IPathSegment
	SetPathData(newValue []IPathSegment)
}

type GeometryPath struct {

	// Path fill mode
	FillMode string `json:"FillMode,omitempty"`

	// Stroke
	Stroke bool `json:"Stroke"`

	// List of PathSegmen objects
	PathData []IPathSegment `json:"PathData,omitempty"`
}

func NewGeometryPath() *GeometryPath {
	instance := new(GeometryPath)
	instance.FillMode = ""
	return instance
}

func (this *GeometryPath) GetFillMode() string {
	return this.FillMode
}

func (this *GeometryPath) SetFillMode(newValue string) {
	this.FillMode = newValue
}
func (this *GeometryPath) GetStroke() bool {
	return this.Stroke
}

func (this *GeometryPath) SetStroke(newValue bool) {
	this.Stroke = newValue
}
func (this *GeometryPath) GetPathData() []IPathSegment {
	return this.PathData
}

func (this *GeometryPath) SetPathData(newValue []IPathSegment) {
	this.PathData = newValue
}

func (this *GeometryPath) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.FillMode = ""
	if valFillMode, ok := objMap["fillMode"]; ok {
		if valFillMode != nil {
			var valueForFillMode string
			err = json.Unmarshal(*valFillMode, &valueForFillMode)
			if err != nil {
				var valueForFillModeInt int32
				err = json.Unmarshal(*valFillMode, &valueForFillModeInt)
				if err != nil {
					return err
				}
				this.FillMode = string(valueForFillModeInt)
			} else {
				this.FillMode = valueForFillMode
			}
		}
	}
	if valFillModeCap, ok := objMap["FillMode"]; ok {
		if valFillModeCap != nil {
			var valueForFillMode string
			err = json.Unmarshal(*valFillModeCap, &valueForFillMode)
			if err != nil {
				var valueForFillModeInt int32
				err = json.Unmarshal(*valFillModeCap, &valueForFillModeInt)
				if err != nil {
					return err
				}
				this.FillMode = string(valueForFillModeInt)
			} else {
				this.FillMode = valueForFillMode
			}
		}
	}
	
	if valStroke, ok := objMap["stroke"]; ok {
		if valStroke != nil {
			var valueForStroke bool
			err = json.Unmarshal(*valStroke, &valueForStroke)
			if err != nil {
				return err
			}
			this.Stroke = valueForStroke
		}
	}
	if valStrokeCap, ok := objMap["Stroke"]; ok {
		if valStrokeCap != nil {
			var valueForStroke bool
			err = json.Unmarshal(*valStrokeCap, &valueForStroke)
			if err != nil {
				return err
			}
			this.Stroke = valueForStroke
		}
	}
	
	if valPathData, ok := objMap["pathData"]; ok {
		if valPathData != nil {
			var valueForPathData []json.RawMessage
			err = json.Unmarshal(*valPathData, &valueForPathData)
			if err != nil {
				return err
			}
			valueForIPathData := make([]IPathSegment, len(valueForPathData))
			for i, v := range valueForPathData {
				vObject, err := createObjectForType("PathSegment", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPathData[i] = vObject.(IPathSegment)
				}
			}
			this.PathData = valueForIPathData
		}
	}
	if valPathDataCap, ok := objMap["PathData"]; ok {
		if valPathDataCap != nil {
			var valueForPathData []json.RawMessage
			err = json.Unmarshal(*valPathDataCap, &valueForPathData)
			if err != nil {
				return err
			}
			valueForIPathData := make([]IPathSegment, len(valueForPathData))
			for i, v := range valueForPathData {
				vObject, err := createObjectForType("PathSegment", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPathData[i] = vObject.(IPathSegment)
				}
			}
			this.PathData = valueForIPathData
		}
	}

	return nil
}
