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

// Represents a comment author
type ICommentAuthor interface {

	// Name
	GetName() string
	SetName(newValue string)

	// Initials
	GetInitials() string
	SetInitials(newValue string)
}

type CommentAuthor struct {

	// Name
	Name string `json:"Name,omitempty"`

	// Initials
	Initials string `json:"Initials,omitempty"`
}

func NewCommentAuthor() *CommentAuthor {
	instance := new(CommentAuthor)
	return instance
}

func (this *CommentAuthor) GetName() string {
	return this.Name
}

func (this *CommentAuthor) SetName(newValue string) {
	this.Name = newValue
}
func (this *CommentAuthor) GetInitials() string {
	return this.Initials
}

func (this *CommentAuthor) SetInitials(newValue string) {
	this.Initials = newValue
}

func (this *CommentAuthor) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valInitials, ok := objMap["initials"]; ok {
		if valInitials != nil {
			var valueForInitials string
			err = json.Unmarshal(*valInitials, &valueForInitials)
			if err != nil {
				return err
			}
			this.Initials = valueForInitials
		}
	}
	if valInitialsCap, ok := objMap["Initials"]; ok {
		if valInitialsCap != nil {
			var valueForInitials string
			err = json.Unmarshal(*valInitialsCap, &valueForInitials)
			if err != nil {
				return err
			}
			this.Initials = valueForInitials
		}
	}

	return nil
}
