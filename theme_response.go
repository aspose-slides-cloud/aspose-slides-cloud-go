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

// Represents response for   DTO
type IThemeResponse interface {

	getCode() int32
	setCode(newValue int32)

	getStatus() string
	setStatus(newValue string)

	// Gets or sets the  DTO
	getTheme() ITheme
	setTheme(newValue ITheme)
}

type ThemeResponse struct {

	Code int32 `json:"Code"`

	Status string `json:"Status,omitempty"`

	// Gets or sets the  DTO
	Theme ITheme `json:"Theme,omitempty"`
}

func (this ThemeResponse) getCode() int32 {
	return this.Code
}

func (this ThemeResponse) setCode(newValue int32) {
	this.Code = newValue
}
func (this ThemeResponse) getStatus() string {
	return this.Status
}

func (this ThemeResponse) setStatus(newValue string) {
	this.Status = newValue
}
func (this ThemeResponse) getTheme() ITheme {
	return this.Theme
}

func (this ThemeResponse) setTheme(newValue ITheme) {
	this.Theme = newValue
}

func (this *ThemeResponse) UnmarshalJSON(b []byte) error {
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

	if valTheme, ok := objMap["Theme"]; ok {
		if valTheme != nil {
			var valueForTheme Theme
			err = json.Unmarshal(*valTheme, &valueForTheme)
			if err != nil {
				return err
			}
			this.Theme = valueForTheme
		}
	}

    return nil
}
