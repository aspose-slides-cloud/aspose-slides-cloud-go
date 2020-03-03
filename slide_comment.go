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
	getAuthor() string
	setAuthor(newValue string)

	// Text.
	getText() string
	setText(newValue string)

	// Creation time.
	getCreatedTime() string
	setCreatedTime(newValue string)

	// Child comments.
	getChildComments() []ISlideComment
	setChildComments(newValue []ISlideComment)
}

type SlideComment struct {

	// Author.
	Author string `json:"Author,omitempty"`

	// Text.
	Text string `json:"Text,omitempty"`

	// Creation time.
	CreatedTime string `json:"CreatedTime,omitempty"`

	// Child comments.
	ChildComments []ISlideComment `json:"ChildComments,omitempty"`
}

func (this *SlideComment) getAuthor() string {
	return this.Author
}

func (this *SlideComment) setAuthor(newValue string) {
	this.Author = newValue
}
func (this *SlideComment) getText() string {
	return this.Text
}

func (this *SlideComment) setText(newValue string) {
	this.Text = newValue
}
func (this *SlideComment) getCreatedTime() string {
	return this.CreatedTime
}

func (this *SlideComment) setCreatedTime(newValue string) {
	this.CreatedTime = newValue
}
func (this *SlideComment) getChildComments() []ISlideComment {
	return this.ChildComments
}

func (this *SlideComment) setChildComments(newValue []ISlideComment) {
	this.ChildComments = newValue
}

func (this *SlideComment) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valAuthor, ok := objMap["author"]; ok {
		if valAuthor != nil {
			var valueForAuthor string
			err = json.Unmarshal(*valAuthor, &valueForAuthor)
			if err != nil {
				return err
			}
			this.Author = valueForAuthor
		}
	}
	if valAuthorCap, ok := objMap["Author"]; ok {
		if valAuthorCap != nil {
			var valueForAuthor string
			err = json.Unmarshal(*valAuthorCap, &valueForAuthor)
			if err != nil {
				return err
			}
			this.Author = valueForAuthor
		}
	}
	
	if valText, ok := objMap["text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	if valTextCap, ok := objMap["Text"]; ok {
		if valTextCap != nil {
			var valueForText string
			err = json.Unmarshal(*valTextCap, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	
	if valCreatedTime, ok := objMap["createdTime"]; ok {
		if valCreatedTime != nil {
			var valueForCreatedTime string
			err = json.Unmarshal(*valCreatedTime, &valueForCreatedTime)
			if err != nil {
				return err
			}
			this.CreatedTime = valueForCreatedTime
		}
	}
	if valCreatedTimeCap, ok := objMap["CreatedTime"]; ok {
		if valCreatedTimeCap != nil {
			var valueForCreatedTime string
			err = json.Unmarshal(*valCreatedTimeCap, &valueForCreatedTime)
			if err != nil {
				return err
			}
			this.CreatedTime = valueForCreatedTime
		}
	}
	
	if valChildComments, ok := objMap["childComments"]; ok {
		if valChildComments != nil {
			var valueForChildComments []SlideComment
			err = json.Unmarshal(*valChildComments, &valueForChildComments)
			if err != nil {
				return err
			}
			valueForIChildComments := make([]ISlideComment, len(valueForChildComments))
			for i, v := range valueForChildComments {
				valueForIChildComments[i] = ISlideComment(&v)
			}
			this.ChildComments = valueForIChildComments
		}
	}
	if valChildCommentsCap, ok := objMap["ChildComments"]; ok {
		if valChildCommentsCap != nil {
			var valueForChildComments []SlideComment
			err = json.Unmarshal(*valChildCommentsCap, &valueForChildComments)
			if err != nil {
				return err
			}
			valueForIChildComments := make([]ISlideComment, len(valueForChildComments))
			for i, v := range valueForChildComments {
				valueForIChildComments[i] = ISlideComment(&v)
			}
			this.ChildComments = valueForIChildComments
		}
	}

    return nil
}
