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

// Save shape task.
type ISaveShape interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Format.
	GetFormat() string
	SetFormat(newValue string)

	// Shape path.
	GetShapePath() string
	SetShapePath(newValue string)

	// Output file.
	GetOutput() IOutputFile
	SetOutput(newValue IOutputFile)

	// Save options.
	GetOptions() IIShapeExportOptions
	SetOptions(newValue IIShapeExportOptions)
}

type SaveShape struct {

	// Task type.
	Type_ string `json:"Type"`

	// Format.
	Format string `json:"Format"`

	// Shape path.
	ShapePath string `json:"ShapePath,omitempty"`

	// Output file.
	Output IOutputFile `json:"Output,omitempty"`

	// Save options.
	Options IIShapeExportOptions `json:"Options,omitempty"`
}

func NewSaveShape() *SaveShape {
	instance := new(SaveShape)
	instance.Type_ = "SaveShape"
	instance.Format = "Jpeg"
	return instance
}

func (this *SaveShape) GetType() string {
	return this.Type_
}

func (this *SaveShape) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SaveShape) GetFormat() string {
	return this.Format
}

func (this *SaveShape) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *SaveShape) GetShapePath() string {
	return this.ShapePath
}

func (this *SaveShape) SetShapePath(newValue string) {
	this.ShapePath = newValue
}
func (this *SaveShape) GetOutput() IOutputFile {
	return this.Output
}

func (this *SaveShape) SetOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this *SaveShape) GetOptions() IIShapeExportOptions {
	return this.Options
}

func (this *SaveShape) SetOptions(newValue IIShapeExportOptions) {
	this.Options = newValue
}

func (this *SaveShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "SaveShape"
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
	this.Format = "Jpeg"
	if valFormat, ok := GetMapValue(objMap, "format"); ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				var valueForFormatInt int32
				err = json.Unmarshal(*valFormat, &valueForFormatInt)
				if err != nil {
					return err
				}
				this.Format = string(valueForFormatInt)
			} else {
				this.Format = valueForFormat
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
	
	if valOutput, ok := GetMapValue(objMap, "output"); ok {
		if valOutput != nil {
			var valueForOutput OutputFile
			err = json.Unmarshal(*valOutput, &valueForOutput)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("OutputFile", *valOutput)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valOutput, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IOutputFile)
			if ok {
				this.Output = vInterfaceObject
			}
		}
	}
	
	if valOptions, ok := GetMapValue(objMap, "options"); ok {
		if valOptions != nil {
			var valueForOptions IShapeExportOptions
			err = json.Unmarshal(*valOptions, &valueForOptions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("IShapeExportOptions", *valOptions)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valOptions, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IIShapeExportOptions)
			if ok {
				this.Options = vInterfaceObject
			}
		}
	}

	return nil
}
