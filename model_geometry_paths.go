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

// Represents container for GeometryPath objects
type IGeometryPaths interface {

	// List of GeometryPath objects
	GetPaths() []IGeometryPath
	SetPaths(newValue []IGeometryPath)
}

type GeometryPaths struct {

	// List of GeometryPath objects
	Paths []IGeometryPath `json:"Paths,omitempty"`
}

func NewGeometryPaths() *GeometryPaths {
	instance := new(GeometryPaths)
	return instance
}

func (this *GeometryPaths) GetPaths() []IGeometryPath {
	return this.Paths
}

func (this *GeometryPaths) SetPaths(newValue []IGeometryPath) {
	this.Paths = newValue
}

func (this *GeometryPaths) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPaths, ok := objMap["paths"]; ok {
		if valPaths != nil {
			var valueForPaths []json.RawMessage
			err = json.Unmarshal(*valPaths, &valueForPaths)
			if err != nil {
				return err
			}
			valueForIPaths := make([]IGeometryPath, len(valueForPaths))
			for i, v := range valueForPaths {
				vObject, err := createObjectForType("GeometryPath", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPaths[i] = vObject.(IGeometryPath)
				}
			}
			this.Paths = valueForIPaths
		}
	}
	if valPathsCap, ok := objMap["Paths"]; ok {
		if valPathsCap != nil {
			var valueForPaths []json.RawMessage
			err = json.Unmarshal(*valPathsCap, &valueForPaths)
			if err != nil {
				return err
			}
			valueForIPaths := make([]IGeometryPath, len(valueForPaths))
			for i, v := range valueForPaths {
				vObject, err := createObjectForType("GeometryPath", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPaths[i] = vObject.(IGeometryPath)
				}
			}
			this.Paths = valueForIPaths
		}
	}

	return nil
}
