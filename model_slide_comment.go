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

// Represents comment of slide
type ISlideComment interface {

	// Author.
	GetAuthor() string
	SetAuthor(newValue string)

	// Text.
	GetText() string
	SetText(newValue string)

	// Creation time.
	GetCreatedTime() string
	SetCreatedTime(newValue string)

	// Child comments.
	GetChildComments() []ISlideCommentBase
	SetChildComments(newValue []ISlideCommentBase)

	// Slide comment type. 
	GetType() string
	SetType(newValue string)
}

type SlideComment struct {

	// Author.
	Author string `json:"Author,omitempty"`

	// Text.
	Text string `json:"Text,omitempty"`

	// Creation time.
	CreatedTime string `json:"CreatedTime,omitempty"`

	// Child comments.
	ChildComments []ISlideCommentBase `json:"ChildComments,omitempty"`

	// Slide comment type. 
	Type_ string `json:"Type"`
}

func NewSlideComment() *SlideComment {
	instance := new(SlideComment)
	instance.Type_ = "Regular"
	return instance
}

func (this *SlideComment) GetAuthor() string {
	return this.Author
}

func (this *SlideComment) SetAuthor(newValue string) {
	this.Author = newValue
}
func (this *SlideComment) GetText() string {
	return this.Text
}

func (this *SlideComment) SetText(newValue string) {
	this.Text = newValue
}
func (this *SlideComment) GetCreatedTime() string {
	return this.CreatedTime
}

func (this *SlideComment) SetCreatedTime(newValue string) {
	this.CreatedTime = newValue
}
func (this *SlideComment) GetChildComments() []ISlideCommentBase {
	return this.ChildComments
}

func (this *SlideComment) SetChildComments(newValue []ISlideCommentBase) {
	this.ChildComments = newValue
}
func (this *SlideComment) GetType() string {
	return this.Type_
}

func (this *SlideComment) SetType(newValue string) {
	this.Type_ = newValue
}

func (this *SlideComment) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valAuthor, ok := GetMapValue(objMap, "author"); ok {
		if valAuthor != nil {
			var valueForAuthor string
			err = json.Unmarshal(*valAuthor, &valueForAuthor)
			if err != nil {
				return err
			}
			this.Author = valueForAuthor
		}
	}
	
	if valText, ok := GetMapValue(objMap, "text"); ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	
	if valCreatedTime, ok := GetMapValue(objMap, "createdTime"); ok {
		if valCreatedTime != nil {
			var valueForCreatedTime string
			err = json.Unmarshal(*valCreatedTime, &valueForCreatedTime)
			if err != nil {
				return err
			}
			this.CreatedTime = valueForCreatedTime
		}
	}
	
	if valChildComments, ok := GetMapValue(objMap, "childComments"); ok {
		if valChildComments != nil {
			var valueForChildComments []json.RawMessage
			err = json.Unmarshal(*valChildComments, &valueForChildComments)
			if err != nil {
				return err
			}
			valueForIChildComments := make([]ISlideCommentBase, len(valueForChildComments))
			for i, v := range valueForChildComments {
				vObject, err := createObjectForType("SlideCommentBase", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIChildComments[i] = vObject.(ISlideCommentBase)
				}
			}
			this.ChildComments = valueForIChildComments
		}
	}
	this.Type_ = "Regular"
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

	return nil
}
