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

// Represents font fallback rule.             
type IFontFallbackRule interface {

	// First index of continuous unicode range.
	GetRangeStartIndex() int32
	SetRangeStartIndex(newValue int32)

	// Last index of continuous unicode range.
	GetRangeEndIndex() int32
	SetRangeEndIndex(newValue int32)

	// List of fallback font links.
	GetFallbackFontList() []string
	SetFallbackFontList(newValue []string)
}

type FontFallbackRule struct {

	// First index of continuous unicode range.
	RangeStartIndex int32 `json:"RangeStartIndex"`

	// Last index of continuous unicode range.
	RangeEndIndex int32 `json:"RangeEndIndex"`

	// List of fallback font links.
	FallbackFontList []string `json:"FallbackFontList,omitempty"`
}

func NewFontFallbackRule() *FontFallbackRule {
	instance := new(FontFallbackRule)
	return instance
}

func (this *FontFallbackRule) GetRangeStartIndex() int32 {
	return this.RangeStartIndex
}

func (this *FontFallbackRule) SetRangeStartIndex(newValue int32) {
	this.RangeStartIndex = newValue
}
func (this *FontFallbackRule) GetRangeEndIndex() int32 {
	return this.RangeEndIndex
}

func (this *FontFallbackRule) SetRangeEndIndex(newValue int32) {
	this.RangeEndIndex = newValue
}
func (this *FontFallbackRule) GetFallbackFontList() []string {
	return this.FallbackFontList
}

func (this *FontFallbackRule) SetFallbackFontList(newValue []string) {
	this.FallbackFontList = newValue
}

func (this *FontFallbackRule) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valRangeStartIndex, ok := GetMapValue(objMap, "rangeStartIndex"); ok {
		if valRangeStartIndex != nil {
			var valueForRangeStartIndex int32
			err = json.Unmarshal(*valRangeStartIndex, &valueForRangeStartIndex)
			if err != nil {
				return err
			}
			this.RangeStartIndex = valueForRangeStartIndex
		}
	}
	
	if valRangeEndIndex, ok := GetMapValue(objMap, "rangeEndIndex"); ok {
		if valRangeEndIndex != nil {
			var valueForRangeEndIndex int32
			err = json.Unmarshal(*valRangeEndIndex, &valueForRangeEndIndex)
			if err != nil {
				return err
			}
			this.RangeEndIndex = valueForRangeEndIndex
		}
	}
	
	if valFallbackFontList, ok := GetMapValue(objMap, "fallbackFontList"); ok {
		if valFallbackFontList != nil {
			var valueForFallbackFontList []string
			err = json.Unmarshal(*valFallbackFontList, &valueForFallbackFontList)
			if err != nil {
				return err
			}
			this.FallbackFontList = valueForFallbackFontList
		}
	}

	return nil
}
