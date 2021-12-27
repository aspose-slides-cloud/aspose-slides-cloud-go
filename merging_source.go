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

// Merging source.
type IMergingSource interface {

	// Source file.
	getInput() IInputFile
	setInput(newValue IInputFile)

	// Indices of slides to be merged.
	getSlides() []int32
	setSlides(newValue []int32)
}

type MergingSource struct {

	// Source file.
	Input IInputFile `json:"Input,omitempty"`

	// Indices of slides to be merged.
	Slides []int32 `json:"Slides,omitempty"`
}

func NewMergingSource() *MergingSource {
	instance := new(MergingSource)
	return instance
}

func (this *MergingSource) getInput() IInputFile {
	return this.Input
}

func (this *MergingSource) setInput(newValue IInputFile) {
	this.Input = newValue
}
func (this *MergingSource) getSlides() []int32 {
	return this.Slides
}

func (this *MergingSource) setSlides(newValue []int32) {
	this.Slides = newValue
}

func (this *MergingSource) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valInput, ok := objMap["input"]; ok {
		if valInput != nil {
			var valueForInput InputFile
			err = json.Unmarshal(*valInput, &valueForInput)
			if err != nil {
				return err
			}
			this.Input = &valueForInput
		}
	}
	if valInputCap, ok := objMap["Input"]; ok {
		if valInputCap != nil {
			var valueForInput InputFile
			err = json.Unmarshal(*valInputCap, &valueForInput)
			if err != nil {
				return err
			}
			this.Input = &valueForInput
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

	return nil
}
