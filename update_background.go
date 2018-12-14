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


type IUpdateBackground interface {

	getType() TaskType
	setType(newValue TaskType)

	getSlides() []int32
	setSlides(newValue []int32)

	getBackground() ISlideBackground
	setBackground(newValue ISlideBackground)
}

type UpdateBackground struct {

	Type_ TaskType `json:"Type"`

	Slides []int32 `json:"Slides,omitempty"`

	Background ISlideBackground `json:"Background,omitempty"`
}

func (this UpdateBackground) getType() TaskType {
	return this.Type_
}

func (this UpdateBackground) setType(newValue TaskType) {
	this.Type_ = newValue
}
func (this UpdateBackground) getSlides() []int32 {
	return this.Slides
}

func (this UpdateBackground) setSlides(newValue []int32) {
	this.Slides = newValue
}
func (this UpdateBackground) getBackground() ISlideBackground {
	return this.Background
}

func (this UpdateBackground) setBackground(newValue ISlideBackground) {
	this.Background = newValue
}

func (this *UpdateBackground) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType TaskType
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valSlides, ok := objMap["Slides"]; ok {
		if valSlides != nil {
			var valueForSlides []int32
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = valueForSlides
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
