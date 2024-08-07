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

// File upload result
type IFilesUploadResult interface {

	// List of uploaded file names
	GetUploaded() []string
	SetUploaded(newValue []string)

	// List of errors.
	GetErrors() []ModelError
	SetErrors(newValue []ModelError)
}

type FilesUploadResult struct {

	// List of uploaded file names
	Uploaded []string `json:"Uploaded,omitempty"`

	// List of errors.
	Errors []ModelError `json:"Errors,omitempty"`
}

func NewFilesUploadResult() *FilesUploadResult {
	instance := new(FilesUploadResult)
	return instance
}

func (this *FilesUploadResult) GetUploaded() []string {
	return this.Uploaded
}

func (this *FilesUploadResult) SetUploaded(newValue []string) {
	this.Uploaded = newValue
}
func (this *FilesUploadResult) GetErrors() []ModelError {
	return this.Errors
}

func (this *FilesUploadResult) SetErrors(newValue []ModelError) {
	this.Errors = newValue
}

func (this *FilesUploadResult) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valUploaded, ok := GetMapValue(objMap, "uploaded"); ok {
		if valUploaded != nil {
			var valueForUploaded []string
			err = json.Unmarshal(*valUploaded, &valueForUploaded)
			if err != nil {
				return err
			}
			this.Uploaded = valueForUploaded
		}
	}
	
	if valErrors, ok := GetMapValue(objMap, "errors"); ok {
		if valErrors != nil {
			var valueForErrors []ModelError
			err = json.Unmarshal(*valErrors, &valueForErrors)
			if err != nil {
				return err
			}
			this.Errors = valueForErrors
		}
	}

	return nil
}
