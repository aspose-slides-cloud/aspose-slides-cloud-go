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

// Represents font substitution rule.
type IFontSubstRule interface {

	// Font to substitute.
	GetSourceFont() string
	SetSourceFont(newValue string)

	// Substitution font.
	GetTargetFont() string
	SetTargetFont(newValue string)

	// Substitute when font is not found. Default: true.
	GetNotFoundOnly() bool
	SetNotFoundOnly(newValue bool)
}

type FontSubstRule struct {

	// Font to substitute.
	SourceFont string `json:"SourceFont,omitempty"`

	// Substitution font.
	TargetFont string `json:"TargetFont,omitempty"`

	// Substitute when font is not found. Default: true.
	NotFoundOnly bool `json:"NotFoundOnly"`
}

func NewFontSubstRule() *FontSubstRule {
	instance := new(FontSubstRule)
	return instance
}

func (this *FontSubstRule) GetSourceFont() string {
	return this.SourceFont
}

func (this *FontSubstRule) SetSourceFont(newValue string) {
	this.SourceFont = newValue
}
func (this *FontSubstRule) GetTargetFont() string {
	return this.TargetFont
}

func (this *FontSubstRule) SetTargetFont(newValue string) {
	this.TargetFont = newValue
}
func (this *FontSubstRule) GetNotFoundOnly() bool {
	return this.NotFoundOnly
}

func (this *FontSubstRule) SetNotFoundOnly(newValue bool) {
	this.NotFoundOnly = newValue
}

func (this *FontSubstRule) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSourceFont, ok := objMap["sourceFont"]; ok {
		if valSourceFont != nil {
			var valueForSourceFont string
			err = json.Unmarshal(*valSourceFont, &valueForSourceFont)
			if err != nil {
				return err
			}
			this.SourceFont = valueForSourceFont
		}
	}
	if valSourceFontCap, ok := objMap["SourceFont"]; ok {
		if valSourceFontCap != nil {
			var valueForSourceFont string
			err = json.Unmarshal(*valSourceFontCap, &valueForSourceFont)
			if err != nil {
				return err
			}
			this.SourceFont = valueForSourceFont
		}
	}
	
	if valTargetFont, ok := objMap["targetFont"]; ok {
		if valTargetFont != nil {
			var valueForTargetFont string
			err = json.Unmarshal(*valTargetFont, &valueForTargetFont)
			if err != nil {
				return err
			}
			this.TargetFont = valueForTargetFont
		}
	}
	if valTargetFontCap, ok := objMap["TargetFont"]; ok {
		if valTargetFontCap != nil {
			var valueForTargetFont string
			err = json.Unmarshal(*valTargetFontCap, &valueForTargetFont)
			if err != nil {
				return err
			}
			this.TargetFont = valueForTargetFont
		}
	}
	
	if valNotFoundOnly, ok := objMap["notFoundOnly"]; ok {
		if valNotFoundOnly != nil {
			var valueForNotFoundOnly bool
			err = json.Unmarshal(*valNotFoundOnly, &valueForNotFoundOnly)
			if err != nil {
				return err
			}
			this.NotFoundOnly = valueForNotFoundOnly
		}
	}
	if valNotFoundOnlyCap, ok := objMap["NotFoundOnly"]; ok {
		if valNotFoundOnlyCap != nil {
			var valueForNotFoundOnly bool
			err = json.Unmarshal(*valNotFoundOnlyCap, &valueForNotFoundOnly)
			if err != nil {
				return err
			}
			this.NotFoundOnly = valueForNotFoundOnly
		}
	}

	return nil
}
