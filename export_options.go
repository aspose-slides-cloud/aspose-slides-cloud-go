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

// Represents export options for whole presentation.
type IExportOptions interface {

	// Setting user password to protect the PDF document. 
	getDefaultRegularFont() string
	setDefaultRegularFont(newValue string)

	getFormat() string
	setFormat(newValue string)
}

type ExportOptions struct {

	// Setting user password to protect the PDF document. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	Format string `json:"Format,omitempty"`
}

func NewExportOptions() *ExportOptions {
	instance := new(ExportOptions)
	return instance
}

func (this *ExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *ExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *ExportOptions) getFormat() string {
	return this.Format
}

func (this *ExportOptions) setFormat(newValue string) {
	this.Format = newValue
}

func (this *ExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDefaultRegularFont, ok := objMap["defaultRegularFont"]; ok {
		if valDefaultRegularFont != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFont, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	if valDefaultRegularFontCap, ok := objMap["DefaultRegularFont"]; ok {
		if valDefaultRegularFontCap != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFontCap, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	
	if valFormat, ok := objMap["format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	if valFormatCap, ok := objMap["Format"]; ok {
		if valFormatCap != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormatCap, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}

    return nil
}
