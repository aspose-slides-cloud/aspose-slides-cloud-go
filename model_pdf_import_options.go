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

// PDF import options.
type IPdfImportOptions interface {

	// True to detect tables.
	GetDetectTables() *bool
	SetDetectTables(newValue *bool)
}

type PdfImportOptions struct {

	// True to detect tables.
	DetectTables *bool `json:"DetectTables"`
}

func NewPdfImportOptions() *PdfImportOptions {
	instance := new(PdfImportOptions)
	return instance
}

func (this *PdfImportOptions) GetDetectTables() *bool {
	return this.DetectTables
}

func (this *PdfImportOptions) SetDetectTables(newValue *bool) {
	this.DetectTables = newValue
}

func (this *PdfImportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDetectTables, ok := objMap["detectTables"]; ok {
		if valDetectTables != nil {
			var valueForDetectTables *bool
			err = json.Unmarshal(*valDetectTables, &valueForDetectTables)
			if err != nil {
				return err
			}
			this.DetectTables = valueForDetectTables
		}
	}
	if valDetectTablesCap, ok := objMap["DetectTables"]; ok {
		if valDetectTablesCap != nil {
			var valueForDetectTables *bool
			err = json.Unmarshal(*valDetectTablesCap, &valueForDetectTables)
			if err != nil {
				return err
			}
			this.DetectTables = valueForDetectTables
		}
	}

	return nil
}
