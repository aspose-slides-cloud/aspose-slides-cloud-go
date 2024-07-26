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
type ISaveSlide interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Output to save the slide to.
	GetOutput() IOutputFile
	SetOutput(newValue IOutputFile)

	// Save format.
	GetFormat() string
	SetFormat(newValue string)

	// Save options.
	GetOptions() IExportOptions
	SetOptions(newValue IExportOptions)

	// Result width for saving to an image format.
	GetWidth() int32
	SetWidth(newValue int32)

	// Result height for saving to an image format.
	GetHeight() int32
	SetHeight(newValue int32)

	// Slide index.
	GetPosition() int32
	SetPosition(newValue int32)
}

type SaveSlide struct {

	// Task type.
	Type_ string `json:"Type"`

	// Output to save the slide to.
	Output IOutputFile `json:"Output,omitempty"`

	// Save format.
	Format string `json:"Format"`

	// Save options.
	Options IExportOptions `json:"Options,omitempty"`

	// Result width for saving to an image format.
	Width int32 `json:"Width,omitempty"`

	// Result height for saving to an image format.
	Height int32 `json:"Height,omitempty"`

	// Slide index.
	Position int32 `json:"Position"`
}

func NewSaveSlide() *SaveSlide {
	instance := new(SaveSlide)
	instance.Type_ = "SaveSlide"
	instance.Format = "Jpeg"
	return instance
}

func (this *SaveSlide) GetType() string {
	return this.Type_
}

func (this *SaveSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SaveSlide) GetOutput() IOutputFile {
	return this.Output
}

func (this *SaveSlide) SetOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this *SaveSlide) GetFormat() string {
	return this.Format
}

func (this *SaveSlide) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *SaveSlide) GetOptions() IExportOptions {
	return this.Options
}

func (this *SaveSlide) SetOptions(newValue IExportOptions) {
	this.Options = newValue
}
func (this *SaveSlide) GetWidth() int32 {
	return this.Width
}

func (this *SaveSlide) SetWidth(newValue int32) {
	this.Width = newValue
}
func (this *SaveSlide) GetHeight() int32 {
	return this.Height
}

func (this *SaveSlide) SetHeight(newValue int32) {
	this.Height = newValue
}
func (this *SaveSlide) GetPosition() int32 {
	return this.Position
}

func (this *SaveSlide) SetPosition(newValue int32) {
	this.Position = newValue
}

func (this *SaveSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "SaveSlide"
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
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valPosition, ok := GetMapValue(objMap, "position"); ok {
		if valPosition != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}

	return nil
}
