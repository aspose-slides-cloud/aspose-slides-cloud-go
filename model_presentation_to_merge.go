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

// Represents presentation to merge
type IPresentationToMerge interface {

	// Get or sets the presentation path
	GetPath() string
	SetPath(newValue string)

	// Get or sets the presentation password
	GetPassword() string
	SetPassword(newValue string)

	// Get or sets the indexes of slides to merge
	GetSlides() []int32
	SetSlides(newValue []int32)

	// Merge (request or storage). 
	GetSource() string
	SetSource(newValue string)
}

type PresentationToMerge struct {

	// Get or sets the presentation path
	Path string `json:"Path,omitempty"`

	// Get or sets the presentation password
	Password string `json:"Password,omitempty"`

	// Get or sets the indexes of slides to merge
	Slides []int32 `json:"Slides,omitempty"`

	// Merge (request or storage). 
	Source string `json:"Source,omitempty"`
}

func NewPresentationToMerge() *PresentationToMerge {
	instance := new(PresentationToMerge)
	return instance
}

func (this *PresentationToMerge) GetPath() string {
	return this.Path
}

func (this *PresentationToMerge) SetPath(newValue string) {
	this.Path = newValue
}
func (this *PresentationToMerge) GetPassword() string {
	return this.Password
}

func (this *PresentationToMerge) SetPassword(newValue string) {
	this.Password = newValue
}
func (this *PresentationToMerge) GetSlides() []int32 {
	return this.Slides
}

func (this *PresentationToMerge) SetSlides(newValue []int32) {
	this.Slides = newValue
}
func (this *PresentationToMerge) GetSource() string {
	return this.Source
}

func (this *PresentationToMerge) SetSource(newValue string) {
	this.Source = newValue
}

func (this *PresentationToMerge) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPath, ok := GetMapValue(objMap, "path"); ok {
		if valPath != nil {
			var valueForPath string
			err = json.Unmarshal(*valPath, &valueForPath)
			if err != nil {
				return err
			}
			this.Path = valueForPath
		}
	}
	
	if valPassword, ok := GetMapValue(objMap, "password"); ok {
		if valPassword != nil {
			var valueForPassword string
			err = json.Unmarshal(*valPassword, &valueForPassword)
			if err != nil {
				return err
			}
			this.Password = valueForPassword
		}
	}
	
	if valSlides, ok := GetMapValue(objMap, "slides"); ok {
		if valSlides != nil {
			var valueForSlides []int32
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = valueForSlides
		}
	}
	
	if valSource, ok := GetMapValue(objMap, "source"); ok {
		if valSource != nil {
			var valueForSource string
			err = json.Unmarshal(*valSource, &valueForSource)
			if err != nil {
				var valueForSourceInt int32
				err = json.Unmarshal(*valSource, &valueForSourceInt)
				if err != nil {
					return err
				}
				this.Source = string(valueForSourceInt)
			} else {
				this.Source = valueForSource
			}
		}
	}

	return nil
}
