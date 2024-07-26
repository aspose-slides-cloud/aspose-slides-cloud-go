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

// Slides layouting options.
type ISlidesLayoutOptions interface {

	GetLayoutType() string
	SetLayoutType(newValue string)
}

type SlidesLayoutOptions struct {

	LayoutType string `json:"LayoutType,omitempty"`
}

func NewSlidesLayoutOptions() *SlidesLayoutOptions {
	instance := new(SlidesLayoutOptions)
	return instance
}

func (this *SlidesLayoutOptions) GetLayoutType() string {
	return this.LayoutType
}

func (this *SlidesLayoutOptions) SetLayoutType(newValue string) {
	this.LayoutType = newValue
}

func (this *SlidesLayoutOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valLayoutType, ok := GetMapValue(objMap, "layoutType"); ok {
		if valLayoutType != nil {
			var valueForLayoutType string
			err = json.Unmarshal(*valLayoutType, &valueForLayoutType)
			if err != nil {
				var valueForLayoutTypeInt int32
				err = json.Unmarshal(*valLayoutType, &valueForLayoutTypeInt)
				if err != nil {
					return err
				}
				this.LayoutType = string(valueForLayoutTypeInt)
			} else {
				this.LayoutType = valueForLayoutType
			}
		}
	}

	return nil
}
