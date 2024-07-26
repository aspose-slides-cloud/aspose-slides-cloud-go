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
	"time"
	"encoding/json"
)

// The error details
type IErrorDetails interface {

	// The request id
	GetRequestId() string
	SetRequestId(newValue string)

	// Date
	GetDate() time.Time
	SetDate(newValue time.Time)
}

type ErrorDetails struct {

	// The request id
	RequestId string `json:"RequestId,omitempty"`

	// Date
	Date time.Time `json:"Date"`
}

func NewErrorDetails() *ErrorDetails {
	instance := new(ErrorDetails)
	return instance
}

func (this *ErrorDetails) GetRequestId() string {
	return this.RequestId
}

func (this *ErrorDetails) SetRequestId(newValue string) {
	this.RequestId = newValue
}
func (this *ErrorDetails) GetDate() time.Time {
	return this.Date
}

func (this *ErrorDetails) SetDate(newValue time.Time) {
	this.Date = newValue
}

func (this *ErrorDetails) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valRequestId, ok := GetMapValue(objMap, "requestId"); ok {
		if valRequestId != nil {
			var valueForRequestId string
			err = json.Unmarshal(*valRequestId, &valueForRequestId)
			if err != nil {
				return err
			}
			this.RequestId = valueForRequestId
		}
	}
	
	if valDate, ok := GetMapValue(objMap, "date"); ok {
		if valDate != nil {
			var valueForDate time.Time
			valueForDate, err = time.Parse("2006-01-02T21:36:33", string(*valDate))
			if err == nil {
				this.Date = valueForDate
			}
		}
	}

	return nil
}
