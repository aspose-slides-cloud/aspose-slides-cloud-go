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

// Replace text task.
type IReplaceText interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Text to be replaced.
	GetOldText() string
	SetOldText(newValue string)

	// Text to replace with.
	GetNewText() string
	SetNewText(newValue string)

	// True to ignore case in replace pattern search.
	GetIgnoreCase() *bool
	SetIgnoreCase(newValue *bool)

	// One-based position of the slide to perform the replace in. 0 to make the replace throughout the presentation.
	GetSlidePosition() int32
	SetSlidePosition(newValue int32)
}

type ReplaceText struct {

	// Task type.
	Type_ string `json:"Type"`

	// Text to be replaced.
	OldText string `json:"OldText,omitempty"`

	// Text to replace with.
	NewText string `json:"NewText,omitempty"`

	// True to ignore case in replace pattern search.
	IgnoreCase *bool `json:"IgnoreCase"`

	// One-based position of the slide to perform the replace in. 0 to make the replace throughout the presentation.
	SlidePosition int32 `json:"SlidePosition"`
}

func NewReplaceText() *ReplaceText {
	instance := new(ReplaceText)
	instance.Type_ = "ReplaceText"
	return instance
}

func (this *ReplaceText) GetType() string {
	return this.Type_
}

func (this *ReplaceText) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ReplaceText) GetOldText() string {
	return this.OldText
}

func (this *ReplaceText) SetOldText(newValue string) {
	this.OldText = newValue
}
func (this *ReplaceText) GetNewText() string {
	return this.NewText
}

func (this *ReplaceText) SetNewText(newValue string) {
	this.NewText = newValue
}
func (this *ReplaceText) GetIgnoreCase() *bool {
	return this.IgnoreCase
}

func (this *ReplaceText) SetIgnoreCase(newValue *bool) {
	this.IgnoreCase = newValue
}
func (this *ReplaceText) GetSlidePosition() int32 {
	return this.SlidePosition
}

func (this *ReplaceText) SetSlidePosition(newValue int32) {
	this.SlidePosition = newValue
}

func (this *ReplaceText) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "ReplaceText"
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	
	if valOldText, ok := GetMapValue(objMap, "oldText"); ok {
		if valOldText != nil {
			var valueForOldText string
			err = json.Unmarshal(*valOldText, &valueForOldText)
			if err != nil {
				return err
			}
			this.OldText = valueForOldText
		}
	}
	
	if valNewText, ok := GetMapValue(objMap, "newText"); ok {
		if valNewText != nil {
			var valueForNewText string
			err = json.Unmarshal(*valNewText, &valueForNewText)
			if err != nil {
				return err
			}
			this.NewText = valueForNewText
		}
	}
	
	if valIgnoreCase, ok := GetMapValue(objMap, "ignoreCase"); ok {
		if valIgnoreCase != nil {
			var valueForIgnoreCase *bool
			err = json.Unmarshal(*valIgnoreCase, &valueForIgnoreCase)
			if err != nil {
				return err
			}
			this.IgnoreCase = valueForIgnoreCase
		}
	}
	
	if valSlidePosition, ok := GetMapValue(objMap, "slidePosition"); ok {
		if valSlidePosition != nil {
			var valueForSlidePosition int32
			err = json.Unmarshal(*valSlidePosition, &valueForSlidePosition)
			if err != nil {
				return err
			}
			this.SlidePosition = valueForSlidePosition
		}
	}

	return nil
}
