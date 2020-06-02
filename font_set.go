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

// Font set.
type IFontSet interface {

	// Complex script font.
	getComplexScript() string
	setComplexScript(newValue string)

	// East Asian font.
	getEastAsian() string
	setEastAsian(newValue string)

	// Latin font.
	getLatin() string
	setLatin(newValue string)
}

type FontSet struct {

	// Complex script font.
	ComplexScript string `json:"ComplexScript,omitempty"`

	// East Asian font.
	EastAsian string `json:"EastAsian,omitempty"`

	// Latin font.
	Latin string `json:"Latin,omitempty"`
}

func NewFontSet() *FontSet {
	instance := new(FontSet)
	return instance
}

func (this *FontSet) getComplexScript() string {
	return this.ComplexScript
}

func (this *FontSet) setComplexScript(newValue string) {
	this.ComplexScript = newValue
}
func (this *FontSet) getEastAsian() string {
	return this.EastAsian
}

func (this *FontSet) setEastAsian(newValue string) {
	this.EastAsian = newValue
}
func (this *FontSet) getLatin() string {
	return this.Latin
}

func (this *FontSet) setLatin(newValue string) {
	this.Latin = newValue
}

func (this *FontSet) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valComplexScript, ok := objMap["complexScript"]; ok {
		if valComplexScript != nil {
			var valueForComplexScript string
			err = json.Unmarshal(*valComplexScript, &valueForComplexScript)
			if err != nil {
				return err
			}
			this.ComplexScript = valueForComplexScript
		}
	}
	if valComplexScriptCap, ok := objMap["ComplexScript"]; ok {
		if valComplexScriptCap != nil {
			var valueForComplexScript string
			err = json.Unmarshal(*valComplexScriptCap, &valueForComplexScript)
			if err != nil {
				return err
			}
			this.ComplexScript = valueForComplexScript
		}
	}
	
	if valEastAsian, ok := objMap["eastAsian"]; ok {
		if valEastAsian != nil {
			var valueForEastAsian string
			err = json.Unmarshal(*valEastAsian, &valueForEastAsian)
			if err != nil {
				return err
			}
			this.EastAsian = valueForEastAsian
		}
	}
	if valEastAsianCap, ok := objMap["EastAsian"]; ok {
		if valEastAsianCap != nil {
			var valueForEastAsian string
			err = json.Unmarshal(*valEastAsianCap, &valueForEastAsian)
			if err != nil {
				return err
			}
			this.EastAsian = valueForEastAsian
		}
	}
	
	if valLatin, ok := objMap["latin"]; ok {
		if valLatin != nil {
			var valueForLatin string
			err = json.Unmarshal(*valLatin, &valueForLatin)
			if err != nil {
				return err
			}
			this.Latin = valueForLatin
		}
	}
	if valLatinCap, ok := objMap["Latin"]; ok {
		if valLatinCap != nil {
			var valueForLatin string
			err = json.Unmarshal(*valLatinCap, &valueForLatin)
			if err != nil {
				return err
			}
			this.Latin = valueForLatin
		}
	}

    return nil
}
