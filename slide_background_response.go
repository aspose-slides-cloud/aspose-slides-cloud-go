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

// Represents response for  DTO
type ISlideBackgroundResponse interface {

	getCode() int32
	setCode(newValue int32)

	getStatus() string
	setStatus(newValue string)

	// Get or sets slide's background DTO
	getBackground() ISlideBackground
	setBackground(newValue ISlideBackground)
}

type SlideBackgroundResponse struct {

	Code int32 `json:"Code"`

	Status string `json:"Status,omitempty"`

	// Get or sets slide's background DTO
	Background ISlideBackground `json:"Background,omitempty"`
}

func (this SlideBackgroundResponse) getCode() int32 {
	return this.Code
}

func (this SlideBackgroundResponse) setCode(newValue int32) {
	this.Code = newValue
}
func (this SlideBackgroundResponse) getStatus() string {
	return this.Status
}

func (this SlideBackgroundResponse) setStatus(newValue string) {
	this.Status = newValue
}
func (this SlideBackgroundResponse) getBackground() ISlideBackground {
	return this.Background
}

func (this SlideBackgroundResponse) setBackground(newValue ISlideBackground) {
	this.Background = newValue
}

func (this *SlideBackgroundResponse) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valCode, ok := objMap["Code"]; ok {
		if valCode != nil {
			var valueForCode int32
			err = json.Unmarshal(*valCode, &valueForCode)
			if err != nil {
				return err
			}
			this.Code = valueForCode
		}
	}

	if valStatus, ok := objMap["Status"]; ok {
		if valStatus != nil {
			var valueForStatus string
			err = json.Unmarshal(*valStatus, &valueForStatus)
			if err != nil {
				return err
			}
			this.Status = valueForStatus
		}
	}

	if valBackground, ok := objMap["Background"]; ok {
		if valBackground != nil {
			var valueForBackground SlideBackground
			err = json.Unmarshal(*valBackground, &valueForBackground)
			if err != nil {
				return err
			}
			this.Background = valueForBackground
		}
	}

    return nil
}
