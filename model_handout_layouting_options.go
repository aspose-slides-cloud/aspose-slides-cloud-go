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

// Handout layouting options
type IHandoutLayoutingOptions interface {

	GetLayoutType() string
	SetLayoutType(newValue string)

	// Specified how many pages and in what sequence will be placed on the page.
	GetHandout() string
	SetHandout(newValue string)

	// True to print the displayed slide numbers.
	GetPrintSlideNumbers() *bool
	SetPrintSlideNumbers(newValue *bool)

	// True to display comments on slide.
	GetPrintComments() *bool
	SetPrintComments(newValue *bool)

	// True to draw frames around the displayed slides.
	GetPrintFrameSlide() *bool
	SetPrintFrameSlide(newValue *bool)
}

type HandoutLayoutingOptions struct {

	LayoutType string `json:"LayoutType"`

	// Specified how many pages and in what sequence will be placed on the page.
	Handout string `json:"Handout,omitempty"`

	// True to print the displayed slide numbers.
	PrintSlideNumbers *bool `json:"PrintSlideNumbers"`

	// True to display comments on slide.
	PrintComments *bool `json:"PrintComments"`

	// True to draw frames around the displayed slides.
	PrintFrameSlide *bool `json:"PrintFrameSlide"`
}

func NewHandoutLayoutingOptions() *HandoutLayoutingOptions {
	instance := new(HandoutLayoutingOptions)
	instance.LayoutType = "Handout"
	return instance
}

func (this *HandoutLayoutingOptions) GetLayoutType() string {
	return this.LayoutType
}

func (this *HandoutLayoutingOptions) SetLayoutType(newValue string) {
	this.LayoutType = newValue
}
func (this *HandoutLayoutingOptions) GetHandout() string {
	return this.Handout
}

func (this *HandoutLayoutingOptions) SetHandout(newValue string) {
	this.Handout = newValue
}
func (this *HandoutLayoutingOptions) GetPrintSlideNumbers() *bool {
	return this.PrintSlideNumbers
}

func (this *HandoutLayoutingOptions) SetPrintSlideNumbers(newValue *bool) {
	this.PrintSlideNumbers = newValue
}
func (this *HandoutLayoutingOptions) GetPrintComments() *bool {
	return this.PrintComments
}

func (this *HandoutLayoutingOptions) SetPrintComments(newValue *bool) {
	this.PrintComments = newValue
}
func (this *HandoutLayoutingOptions) GetPrintFrameSlide() *bool {
	return this.PrintFrameSlide
}

func (this *HandoutLayoutingOptions) SetPrintFrameSlide(newValue *bool) {
	this.PrintFrameSlide = newValue
}

func (this *HandoutLayoutingOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.LayoutType = "Handout"
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
	
	if valHandout, ok := GetMapValue(objMap, "handout"); ok {
		if valHandout != nil {
			var valueForHandout string
			err = json.Unmarshal(*valHandout, &valueForHandout)
			if err != nil {
				var valueForHandoutInt int32
				err = json.Unmarshal(*valHandout, &valueForHandoutInt)
				if err != nil {
					return err
				}
				this.Handout = string(valueForHandoutInt)
			} else {
				this.Handout = valueForHandout
			}
		}
	}
	
	if valPrintSlideNumbers, ok := GetMapValue(objMap, "printSlideNumbers"); ok {
		if valPrintSlideNumbers != nil {
			var valueForPrintSlideNumbers *bool
			err = json.Unmarshal(*valPrintSlideNumbers, &valueForPrintSlideNumbers)
			if err != nil {
				return err
			}
			this.PrintSlideNumbers = valueForPrintSlideNumbers
		}
	}
	
	if valPrintComments, ok := GetMapValue(objMap, "printComments"); ok {
		if valPrintComments != nil {
			var valueForPrintComments *bool
			err = json.Unmarshal(*valPrintComments, &valueForPrintComments)
			if err != nil {
				return err
			}
			this.PrintComments = valueForPrintComments
		}
	}
	
	if valPrintFrameSlide, ok := GetMapValue(objMap, "printFrameSlide"); ok {
		if valPrintFrameSlide != nil {
			var valueForPrintFrameSlide *bool
			err = json.Unmarshal(*valPrintFrameSlide, &valueForPrintFrameSlide)
			if err != nil {
				return err
			}
			this.PrintFrameSlide = valueForPrintFrameSlide
		}
	}

	return nil
}
