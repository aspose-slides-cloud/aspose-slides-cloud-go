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

// Provides options that control how a presentation is saved in PPTX format.
type IPptxExportOptions interface {

	// Export format.
	getFormat() string
	setFormat(newValue string)

	// The conformance class to which the PresentationML document conforms. Read/write Conformance.
	getConformance() string
	setConformance(newValue string)
}

type PptxExportOptions struct {

	// Export format.
	Format string `json:"Format,omitempty"`

	// The conformance class to which the PresentationML document conforms. Read/write Conformance.
	Conformance string `json:"Conformance"`
}

func (this PptxExportOptions) getFormat() string {
	return this.Format
}

func (this PptxExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this PptxExportOptions) getConformance() string {
	return this.Conformance
}

func (this PptxExportOptions) setConformance(newValue string) {
	this.Conformance = newValue
}

func (this *PptxExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
	this.Conformance = "Ecma376_2006"
	if valConformance, ok := objMap["conformance"]; ok {
		if valConformance != nil {
			var valueForConformance string
			err = json.Unmarshal(*valConformance, &valueForConformance)
			if err != nil {
				var valueForConformanceInt int32
				err = json.Unmarshal(*valConformance, &valueForConformanceInt)
				if err != nil {
					return err
				}
				this.Conformance = string(valueForConformanceInt)
			} else {
				this.Conformance = valueForConformance
			}
		}
	}
	if valConformanceCap, ok := objMap["Conformance"]; ok {
		if valConformanceCap != nil {
			var valueForConformance string
			err = json.Unmarshal(*valConformanceCap, &valueForConformance)
			if err != nil {
				var valueForConformanceInt int32
				err = json.Unmarshal(*valConformanceCap, &valueForConformanceInt)
				if err != nil {
					return err
				}
				this.Conformance = string(valueForConformanceInt)
			} else {
				this.Conformance = valueForConformance
			}
		}
	}

    return nil
}
