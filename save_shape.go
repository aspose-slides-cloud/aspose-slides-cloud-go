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

func (this SaveShape) getType() string {
	return this.Type_
}

func (this SaveShape) setType(newValue string) {
	this.Type_ = newValue
}
func (this SaveShape) getFormat() string {
	return this.Format
}

func (this SaveShape) setFormat(newValue string) {
	this.Format = newValue
}
func (this SaveShape) getShapePath() string {
	return this.ShapePath
}

func (this SaveShape) setShapePath(newValue string) {
	this.ShapePath = newValue
}
func (this SaveShape) getOutput() IOutputFile {
	return this.Output
}

func (this SaveShape) setOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this SaveShape) getOptions() IIShapeExportOptions {
	return this.Options
}

func (this SaveShape) setOptions(newValue IIShapeExportOptions) {
	this.Options = newValue
}

func (this *SaveShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

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

	if valFormat, ok := objMap["Format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
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

	if valOutput, ok := objMap["Output"]; ok {
		if valOutput != nil {
			var valueForOutput OutputFile
			err = json.Unmarshal(*valOutput, &valueForOutput)
			if err != nil {
				return err
			}
			this.Output = valueForOutput
		}
	}

	if valOptions, ok := objMap["Options"]; ok {
		if valOptions != nil {
			var valueForOptions IShapeExportOptions
			err = json.Unmarshal(*valOptions, &valueForOptions)
			if err != nil {
				return err
			}
			this.Options = valueForOptions
		}
	}

    return nil
}
