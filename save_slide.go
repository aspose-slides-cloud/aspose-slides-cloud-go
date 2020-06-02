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
	getType() string
	setType(newValue string)

	// Output to save the slide to.
	getOutput() IOutputFile
	setOutput(newValue IOutputFile)

	// Save format.
	getFormat() string
	setFormat(newValue string)

	// Save options.
	getOptions() IExportOptions
	setOptions(newValue IExportOptions)

	// Result width for saving to an image format.
	getWidth() int32
	setWidth(newValue int32)

	// Result height for saving to an image format.
	getHeight() int32
	setHeight(newValue int32)

	// Slide index.
	getPosition() int32
	setPosition(newValue int32)
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

func (this *SaveSlide) getType() string {
	return this.Type_
}

func (this *SaveSlide) setType(newValue string) {
	this.Type_ = newValue
}
func (this *SaveSlide) getOutput() IOutputFile {
	return this.Output
}

func (this *SaveSlide) setOutput(newValue IOutputFile) {
	this.Output = newValue
}
func (this *SaveSlide) getFormat() string {
	return this.Format
}

func (this *SaveSlide) setFormat(newValue string) {
	this.Format = newValue
}
func (this *SaveSlide) getOptions() IExportOptions {
	return this.Options
}

func (this *SaveSlide) setOptions(newValue IExportOptions) {
	this.Options = newValue
}
func (this *SaveSlide) getWidth() int32 {
	return this.Width
}

func (this *SaveSlide) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *SaveSlide) getHeight() int32 {
	return this.Height
}

func (this *SaveSlide) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *SaveSlide) getPosition() int32 {
	return this.Position
}

func (this *SaveSlide) setPosition(newValue int32) {
	this.Position = newValue
}

func (this *SaveSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "SaveSlide"
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
	
	if valOutput, ok := objMap["output"]; ok {
		if valOutput != nil {
			var valueForOutput OutputFile
			err = json.Unmarshal(*valOutput, &valueForOutput)
			if err != nil {
				return err
			}
			this.Output = &valueForOutput
		}
	}
	if valOutputCap, ok := objMap["Output"]; ok {
		if valOutputCap != nil {
			var valueForOutput OutputFile
			err = json.Unmarshal(*valOutputCap, &valueForOutput)
			if err != nil {
				return err
			}
			this.Output = &valueForOutput
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
	
	if valOptions, ok := objMap["options"]; ok {
		if valOptions != nil {
			var valueForOptions ExportOptions
			err = json.Unmarshal(*valOptions, &valueForOptions)
			if err != nil {
				return err
			}
			this.Options = &valueForOptions
		}
	}
	if valOptionsCap, ok := objMap["Options"]; ok {
		if valOptionsCap != nil {
			var valueForOptions ExportOptions
			err = json.Unmarshal(*valOptionsCap, &valueForOptions)
			if err != nil {
				return err
			}
			this.Options = &valueForOptions
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valPosition, ok := objMap["position"]; ok {
		if valPosition != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}
	if valPositionCap, ok := objMap["Position"]; ok {
		if valPositionCap != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPositionCap, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}

    return nil
}
