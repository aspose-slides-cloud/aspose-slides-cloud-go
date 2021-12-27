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

// Update background task.
type IUpdateBackground interface {

	// Task type.
	getType() string
	setType(newValue string)

	// List of slide indices.
	getSlides() []int32
	setSlides(newValue []int32)

	// Background DTO.
	getBackground() ISlideBackground
	setBackground(newValue ISlideBackground)
}

type UpdateBackground struct {

	// Task type.
	Type_ string `json:"Type"`

	// List of slide indices.
	Slides []int32 `json:"Slides,omitempty"`

	// Background DTO.
	Background ISlideBackground `json:"Background,omitempty"`
}

func NewUpdateBackground() *UpdateBackground {
	instance := new(UpdateBackground)
	instance.Type_ = "UpdateBackground"
	return instance
}

func (this *UpdateBackground) getType() string {
	return this.Type_
}

func (this *UpdateBackground) setType(newValue string) {
	this.Type_ = newValue
}
func (this *UpdateBackground) getSlides() []int32 {
	return this.Slides
}

func (this *UpdateBackground) setSlides(newValue []int32) {
	this.Slides = newValue
}
func (this *UpdateBackground) getBackground() ISlideBackground {
	return this.Background
}

func (this *UpdateBackground) setBackground(newValue ISlideBackground) {
	this.Background = newValue
}

func (this *UpdateBackground) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "UpdateBackground"
	if valType, ok := objMap["type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valSlides, ok := objMap["slides"]; ok {
		if valSlides != nil {
			var valueForSlides []int32
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = valueForSlides
		}
	}
	if valSlidesCap, ok := objMap["Slides"]; ok {
		if valSlidesCap != nil {
			var valueForSlides []int32
			err = json.Unmarshal(*valSlidesCap, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = valueForSlides
		}
	}
	
	if valBackground, ok := objMap["background"]; ok {
		if valBackground != nil {
			var valueForBackground SlideBackground
			err = json.Unmarshal(*valBackground, &valueForBackground)
			if err != nil {
				return err
			}
			this.Background = &valueForBackground
		}
	}
	if valBackgroundCap, ok := objMap["Background"]; ok {
		if valBackgroundCap != nil {
			var valueForBackground SlideBackground
			err = json.Unmarshal(*valBackgroundCap, &valueForBackground)
			if err != nil {
				return err
			}
			this.Background = &valueForBackground
		}
	}

	return nil
}
