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
	getType() string
	setType(newValue string)

	// Format.
	getFormat() string
	setFormat(newValue string)

	// Shape path.
	getShapePath() string
	setShapePath(newValue string)

	// Output file.
	getOutput() IOutputFile
	setOutput(newValue IOutputFile)

	// Save options.
	getOptions() IIShapeExportOptions
	setOptions(newValue IIShapeExportOptions)
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

func (this *SaveShape) getType() string {
	return this.Type_
}

func (this *SaveShape) setType(newValue string) {
	this.Type_ = newValue
}
func (this *SaveShape) getFormat() string {
	return this.Format
}

func (this *SaveShape) setFormat(newValue string) {
	this.Format = newValue
}
func (this *SaveShape) getShapePath() string {
	return this.ShapePath
}

func (this *SaveShape) setShapePath(newValue string) {
	this.ShapePath = newValue
}
func (this *SaveShape) getOutput() IOutputFile {
	return this.Output
}

func (this *SaveShape) setOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this *SaveShape) getOptions() IIShapeExportOptions {
	return this.Options
}

func (this *SaveShape) setOptions(newValue IIShapeExportOptions) {
	this.Options = newValue
}

func (this *SaveShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "SaveShape"
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
	this.Format = "Jpeg"
	if valFormat, ok := objMap["format"]; ok {
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
	if valFormatCap, ok := objMap["Format"]; ok {
		if valFormatCap != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormatCap, &valueForFormat)
			if err != nil {
				var valueForFormatInt int32
				err = json.Unmarshal(*valFormatCap, &valueForFormatInt)
				if err != nil {
					return err
				}
				this.Format = string(valueForFormatInt)
			} else {
				this.Format = valueForFormat
			}
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
	
	if valOutput, ok := objMap["output"]; ok {
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
	if valOutputCap, ok := objMap["Output"]; ok {
		if valOutputCap != nil {
			var valueForOutput OutputFile
			err = json.Unmarshal(*valOutputCap, &valueForOutput)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("OutputFile", *valOutputCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valOutputCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IOutputFile)
			if ok {
				this.Output = vInterfaceObject
			}
		}
	}
	
	if valOptions, ok := objMap["options"]; ok {
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
	if valOptionsCap, ok := objMap["Options"]; ok {
		if valOptionsCap != nil {
			var valueForOptions IShapeExportOptions
			err = json.Unmarshal(*valOptionsCap, &valueForOptions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("IShapeExportOptions", *valOptionsCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valOptionsCap, &vObject)
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
