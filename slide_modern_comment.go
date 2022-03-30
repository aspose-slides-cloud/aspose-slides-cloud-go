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

// Represents modern comment of slide
type ISlideModernComment interface {

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
	getChildComments() []ISlideCommentBase
	setChildComments(newValue []ISlideCommentBase)

	// Slide modern comment type. 
	getType() string
	setType(newValue string)

	// Returns or sets starting position of text selection in text frame if the comment associated with AutoShape. Read/write Int32.
	getTextSelectionStart() int32
	setTextSelectionStart(newValue int32)

	// Returns or sets text selection length in text frame if the comment associated with AutoShape. Read/write Int32.
	getTextSelectionLength() int32
	setTextSelectionLength(newValue int32)

	// Returns or sets the status of the comment. Read/write ModernCommentStatus.
	getStatus() string
	setStatus(newValue string)
}

type SlideModernComment struct {

	// Author.
	Author string `json:"Author,omitempty"`

	// Text.
	Text string `json:"Text,omitempty"`

	// Creation time.
	CreatedTime string `json:"CreatedTime,omitempty"`

	// Child comments.
	ChildComments []ISlideCommentBase `json:"ChildComments,omitempty"`

	// Slide modern comment type. 
	Type_ string `json:"Type"`

	// Returns or sets starting position of text selection in text frame if the comment associated with AutoShape. Read/write Int32.
	TextSelectionStart int32 `json:"TextSelectionStart,omitempty"`

	// Returns or sets text selection length in text frame if the comment associated with AutoShape. Read/write Int32.
	TextSelectionLength int32 `json:"TextSelectionLength,omitempty"`

	// Returns or sets the status of the comment. Read/write ModernCommentStatus.
	Status string `json:"Status,omitempty"`
}

func NewSlideModernComment() *SlideModernComment {
	instance := new(SlideModernComment)
	instance.Type_ = "Modern"
	instance.Status = ""
	return instance
}

func (this *SlideModernComment) getAuthor() string {
	return this.Author
}

func (this *SlideModernComment) setAuthor(newValue string) {
	this.Author = newValue
}
func (this *SlideModernComment) getText() string {
	return this.Text
}

func (this *SlideModernComment) setText(newValue string) {
	this.Text = newValue
}
func (this *SlideModernComment) getCreatedTime() string {
	return this.CreatedTime
}

func (this *SlideModernComment) setCreatedTime(newValue string) {
	this.CreatedTime = newValue
}
func (this *SlideModernComment) getChildComments() []ISlideCommentBase {
	return this.ChildComments
}

func (this *SlideModernComment) setChildComments(newValue []ISlideCommentBase) {
	this.ChildComments = newValue
}
func (this *SlideModernComment) getType() string {
	return this.Type_
}

func (this *SlideModernComment) setType(newValue string) {
	this.Type_ = newValue
}
func (this *SlideModernComment) getTextSelectionStart() int32 {
	return this.TextSelectionStart
}

func (this *SlideModernComment) setTextSelectionStart(newValue int32) {
	this.TextSelectionStart = newValue
}
func (this *SlideModernComment) getTextSelectionLength() int32 {
	return this.TextSelectionLength
}

func (this *SlideModernComment) setTextSelectionLength(newValue int32) {
	this.TextSelectionLength = newValue
}
func (this *SlideModernComment) getStatus() string {
	return this.Status
}

func (this *SlideModernComment) setStatus(newValue string) {
	this.Status = newValue
}

func (this *SlideModernComment) UnmarshalJSON(b []byte) error {
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
	if valChildCommentsCap, ok := objMap["ChildComments"]; ok {
		if valChildCommentsCap != nil {
			var valueForChildComments []json.RawMessage
			err = json.Unmarshal(*valChildCommentsCap, &valueForChildComments)
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
	this.Type_ = "Modern"
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
	
	if valTextSelectionStart, ok := objMap["textSelectionStart"]; ok {
		if valTextSelectionStart != nil {
			var valueForTextSelectionStart int32
			err = json.Unmarshal(*valTextSelectionStart, &valueForTextSelectionStart)
			if err != nil {
				return err
			}
			this.TextSelectionStart = valueForTextSelectionStart
		}
	}
	if valTextSelectionStartCap, ok := objMap["TextSelectionStart"]; ok {
		if valTextSelectionStartCap != nil {
			var valueForTextSelectionStart int32
			err = json.Unmarshal(*valTextSelectionStartCap, &valueForTextSelectionStart)
			if err != nil {
				return err
			}
			this.TextSelectionStart = valueForTextSelectionStart
		}
	}
	
	if valTextSelectionLength, ok := objMap["textSelectionLength"]; ok {
		if valTextSelectionLength != nil {
			var valueForTextSelectionLength int32
			err = json.Unmarshal(*valTextSelectionLength, &valueForTextSelectionLength)
			if err != nil {
				return err
			}
			this.TextSelectionLength = valueForTextSelectionLength
		}
	}
	if valTextSelectionLengthCap, ok := objMap["TextSelectionLength"]; ok {
		if valTextSelectionLengthCap != nil {
			var valueForTextSelectionLength int32
			err = json.Unmarshal(*valTextSelectionLengthCap, &valueForTextSelectionLength)
			if err != nil {
				return err
			}
			this.TextSelectionLength = valueForTextSelectionLength
		}
	}
	this.Status = ""
	if valStatus, ok := objMap["status"]; ok {
		if valStatus != nil {
			var valueForStatus string
			err = json.Unmarshal(*valStatus, &valueForStatus)
			if err != nil {
				var valueForStatusInt int32
				err = json.Unmarshal(*valStatus, &valueForStatusInt)
				if err != nil {
					return err
				}
				this.Status = string(valueForStatusInt)
			} else {
				this.Status = valueForStatus
			}
		}
	}
	if valStatusCap, ok := objMap["Status"]; ok {
		if valStatusCap != nil {
			var valueForStatus string
			err = json.Unmarshal(*valStatusCap, &valueForStatus)
			if err != nil {
				var valueForStatusInt int32
				err = json.Unmarshal(*valStatusCap, &valueForStatusInt)
				if err != nil {
					return err
				}
				this.Status = string(valueForStatusInt)
			} else {
				this.Status = valueForStatus
			}
		}
	}

	return nil
}
