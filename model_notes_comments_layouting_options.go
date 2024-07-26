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

// Notes &amp; comments layouting options.
type INotesCommentsLayoutingOptions interface {

	GetLayoutType() string
	SetLayoutType(newValue string)

	// Gets or sets the position of the notes on the page.
	GetNotesPosition() string
	SetNotesPosition(newValue string)

	// Gets or sets the position of the comments on the page.
	GetCommentsPosition() string
	SetCommentsPosition(newValue string)

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	GetCommentsAreaWidth() int32
	SetCommentsAreaWidth(newValue int32)

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	GetCommentsAreaColor() string
	SetCommentsAreaColor(newValue string)

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	GetShowCommentsByNoAuthor() *bool
	SetShowCommentsByNoAuthor(newValue *bool)
}

type NotesCommentsLayoutingOptions struct {

	LayoutType string `json:"LayoutType"`

	// Gets or sets the position of the notes on the page.
	NotesPosition string `json:"NotesPosition,omitempty"`

	// Gets or sets the position of the comments on the page.
	CommentsPosition string `json:"CommentsPosition,omitempty"`

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	CommentsAreaWidth int32 `json:"CommentsAreaWidth,omitempty"`

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	CommentsAreaColor string `json:"CommentsAreaColor,omitempty"`

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	ShowCommentsByNoAuthor *bool `json:"ShowCommentsByNoAuthor"`
}

func NewNotesCommentsLayoutingOptions() *NotesCommentsLayoutingOptions {
	instance := new(NotesCommentsLayoutingOptions)
	instance.LayoutType = "NotesComments"
	return instance
}

func (this *NotesCommentsLayoutingOptions) GetLayoutType() string {
	return this.LayoutType
}

func (this *NotesCommentsLayoutingOptions) SetLayoutType(newValue string) {
	this.LayoutType = newValue
}
func (this *NotesCommentsLayoutingOptions) GetNotesPosition() string {
	return this.NotesPosition
}

func (this *NotesCommentsLayoutingOptions) SetNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this *NotesCommentsLayoutingOptions) GetCommentsPosition() string {
	return this.CommentsPosition
}

func (this *NotesCommentsLayoutingOptions) SetCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this *NotesCommentsLayoutingOptions) GetCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this *NotesCommentsLayoutingOptions) SetCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this *NotesCommentsLayoutingOptions) GetCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this *NotesCommentsLayoutingOptions) SetCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this *NotesCommentsLayoutingOptions) GetShowCommentsByNoAuthor() *bool {
	return this.ShowCommentsByNoAuthor
}

func (this *NotesCommentsLayoutingOptions) SetShowCommentsByNoAuthor(newValue *bool) {
	this.ShowCommentsByNoAuthor = newValue
}

func (this *NotesCommentsLayoutingOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.LayoutType = "NotesComments"
	if valLayoutType, ok := GetMapValue(objMap, "layoutType"); ok {
		if valLayoutType != nil {
			var valueForLayoutType string
			err = json.Unmarshal(*valLayoutType, &valueForLayoutType)
			if err != nil {
				var valueForLayoutTypeInt int32
				err = json.Unmarshal(*valLayoutType, &valueForLayoutTypeInt)
				if err != nil {
					return err
				}
				this.LayoutType = string(valueForLayoutTypeInt)
			} else {
				this.LayoutType = valueForLayoutType
			}
		}
	}
	
	if valNotesPosition, ok := GetMapValue(objMap, "notesPosition"); ok {
		if valNotesPosition != nil {
			var valueForNotesPosition string
			err = json.Unmarshal(*valNotesPosition, &valueForNotesPosition)
			if err != nil {
				var valueForNotesPositionInt int32
				err = json.Unmarshal(*valNotesPosition, &valueForNotesPositionInt)
				if err != nil {
					return err
				}
				this.NotesPosition = string(valueForNotesPositionInt)
			} else {
				this.NotesPosition = valueForNotesPosition
			}
		}
	}
	
	if valCommentsPosition, ok := GetMapValue(objMap, "commentsPosition"); ok {
		if valCommentsPosition != nil {
			var valueForCommentsPosition string
			err = json.Unmarshal(*valCommentsPosition, &valueForCommentsPosition)
			if err != nil {
				var valueForCommentsPositionInt int32
				err = json.Unmarshal(*valCommentsPosition, &valueForCommentsPositionInt)
				if err != nil {
					return err
				}
				this.CommentsPosition = string(valueForCommentsPositionInt)
			} else {
				this.CommentsPosition = valueForCommentsPosition
			}
		}
	}
	
	if valCommentsAreaWidth, ok := GetMapValue(objMap, "commentsAreaWidth"); ok {
		if valCommentsAreaWidth != nil {
			var valueForCommentsAreaWidth int32
			err = json.Unmarshal(*valCommentsAreaWidth, &valueForCommentsAreaWidth)
			if err != nil {
				return err
			}
			this.CommentsAreaWidth = valueForCommentsAreaWidth
		}
	}
	
	if valCommentsAreaColor, ok := GetMapValue(objMap, "commentsAreaColor"); ok {
		if valCommentsAreaColor != nil {
			var valueForCommentsAreaColor string
			err = json.Unmarshal(*valCommentsAreaColor, &valueForCommentsAreaColor)
			if err != nil {
				return err
			}
			this.CommentsAreaColor = valueForCommentsAreaColor
		}
	}
	
	if valShowCommentsByNoAuthor, ok := GetMapValue(objMap, "showCommentsByNoAuthor"); ok {
		if valShowCommentsByNoAuthor != nil {
			var valueForShowCommentsByNoAuthor *bool
			err = json.Unmarshal(*valShowCommentsByNoAuthor, &valueForShowCommentsByNoAuthor)
			if err != nil {
				return err
			}
			this.ShowCommentsByNoAuthor = valueForShowCommentsByNoAuthor
		}
	}

	return nil
}
