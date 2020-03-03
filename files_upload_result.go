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
	getUploaded() []string
	setUploaded(newValue []string)

	// List of errors.
	getErrors() []ModelError
	setErrors(newValue []ModelError)
}

type FilesUploadResult struct {

	// List of uploaded file names
	Uploaded []string `json:"Uploaded,omitempty"`

	// List of errors.
	Errors []ModelError `json:"Errors,omitempty"`
}

func (this *FilesUploadResult) getUploaded() []string {
	return this.Uploaded
}

func (this *FilesUploadResult) setUploaded(newValue []string) {
	this.Uploaded = newValue
}
func (this *FilesUploadResult) getErrors() []ModelError {
	return this.Errors
}

func (this *FilesUploadResult) setErrors(newValue []ModelError) {
	this.Errors = newValue
}

func (this *FilesUploadResult) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valUploaded, ok := objMap["uploaded"]; ok {
		if valUploaded != nil {
			var valueForUploaded []string
			err = json.Unmarshal(*valUploaded, &valueForUploaded)
			if err != nil {
				return err
			}
			this.Uploaded = valueForUploaded
		}
	}
	if valUploadedCap, ok := objMap["Uploaded"]; ok {
		if valUploadedCap != nil {
			var valueForUploaded []string
			err = json.Unmarshal(*valUploadedCap, &valueForUploaded)
			if err != nil {
				return err
			}
			this.Uploaded = valueForUploaded
		}
	}
	
	if valErrors, ok := objMap["errors"]; ok {
		if valErrors != nil {
			var valueForErrors []ModelError
			err = json.Unmarshal(*valErrors, &valueForErrors)
			if err != nil {
				return err
			}
			this.Errors = valueForErrors
		}
	}
	if valErrorsCap, ok := objMap["Errors"]; ok {
		if valErrorsCap != nil {
			var valueForErrors []ModelError
			err = json.Unmarshal(*valErrorsCap, &valueForErrors)
			if err != nil {
				return err
			}
			this.Errors = valueForErrors
		}
	}

    return nil
}
