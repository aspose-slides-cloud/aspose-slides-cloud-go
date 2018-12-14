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


type ISave interface {

	getType() TaskType
	setType(newValue TaskType)

	getFormat() ExportFormat
	setFormat(newValue ExportFormat)

	getOutput() IOutputFile
	setOutput(newValue IOutputFile)

	getOptions() IExportOptions
	setOptions(newValue IExportOptions)
}

type Save struct {

	Type_ TaskType `json:"Type"`

	Format ExportFormat `json:"Format,omitempty"`

	Output IOutputFile `json:"Output,omitempty"`

	Options IExportOptions `json:"Options,omitempty"`
}

func (this Save) getType() TaskType {
	return this.Type_
}

func (this Save) setType(newValue TaskType) {
	this.Type_ = newValue
}
func (this Save) getFormat() ExportFormat {
	return this.Format
}

func (this Save) setFormat(newValue ExportFormat) {
	this.Format = newValue
}
func (this Save) getOutput() IOutputFile {
	return this.Output
}

func (this Save) setOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this Save) getOptions() IExportOptions {
	return this.Options
}

func (this Save) setOptions(newValue IExportOptions) {
	this.Options = newValue
}

func (this *Save) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType TaskType
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valFormat, ok := objMap["Format"]; ok {
		if valFormat != nil {
			var valueForFormat ExportFormat
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
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
			var valueForOptions ExportOptions
			err = json.Unmarshal(*valOptions, &valueForOptions)
			if err != nil {
				return err
			}
			this.Options = valueForOptions
		}
	}

    return nil
}
