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

// Save slide task.
type ISave interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Format.
	GetFormat() string
	SetFormat(newValue string)

	// Output file.
	GetOutput() IOutputFile
	SetOutput(newValue IOutputFile)

	// Save options.
	GetOptions() IExportOptions
	SetOptions(newValue IExportOptions)
}

type Save struct {

	// Task type.
	Type_ string `json:"Type"`

	// Format.
	Format string `json:"Format"`

	// Output file.
	Output IOutputFile `json:"Output,omitempty"`

	// Save options.
	Options IExportOptions `json:"Options,omitempty"`
}

func NewSave() *Save {
	instance := new(Save)
	instance.Type_ = "Save"
	instance.Format = "Pdf"
	return instance
}

func (this *Save) GetType() string {
	return this.Type_
}

func (this *Save) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Save) GetFormat() string {
	return this.Format
}

func (this *Save) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *Save) GetOutput() IOutputFile {
	return this.Output
}

func (this *Save) SetOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this *Save) GetOptions() IExportOptions {
	return this.Options
}

func (this *Save) SetOptions(newValue IExportOptions) {
	this.Options = newValue
}

func (this *Save) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Save"
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
	this.Format = "Pdf"
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
			var valueForOptions ExportOptions
			err = json.Unmarshal(*valOptions, &valueForOptions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ExportOptions", *valOptions)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valOptions, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IExportOptions)
			if ok {
				this.Options = vInterfaceObject
			}
		}
	}

	return nil
}
