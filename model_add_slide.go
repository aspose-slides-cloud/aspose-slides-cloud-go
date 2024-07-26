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

// Add slide task.
type IAddSlide interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// File to clone a slide from.
	GetCloneFromFile() IInputFile
	SetCloneFromFile(newValue IInputFile)

	// Position of the slide to clone.
	GetCloneFromPosition() int32
	SetCloneFromPosition(newValue int32)

	// Position at which to insert the slide.
	GetPosition() int32
	SetPosition(newValue int32)

	// Alias of layout (href, index or type). If value is null a blank slide is added.
	GetLayoutAlias() string
	SetLayoutAlias(newValue string)
}

type AddSlide struct {

	// Task type.
	Type_ string `json:"Type"`

	// File to clone a slide from.
	CloneFromFile IInputFile `json:"CloneFromFile,omitempty"`

	// Position of the slide to clone.
	CloneFromPosition int32 `json:"CloneFromPosition"`

	// Position at which to insert the slide.
	Position int32 `json:"Position"`

	// Alias of layout (href, index or type). If value is null a blank slide is added.
	LayoutAlias string `json:"LayoutAlias,omitempty"`
}

func NewAddSlide() *AddSlide {
	instance := new(AddSlide)
	instance.Type_ = "AddSlide"
	return instance
}

func (this *AddSlide) GetType() string {
	return this.Type_
}

func (this *AddSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AddSlide) GetCloneFromFile() IInputFile {
	return this.CloneFromFile
}

func (this *AddSlide) SetCloneFromFile(newValue IInputFile) {
	this.CloneFromFile = newValue
}
func (this *AddSlide) GetCloneFromPosition() int32 {
	return this.CloneFromPosition
}

func (this *AddSlide) SetCloneFromPosition(newValue int32) {
	this.CloneFromPosition = newValue
}
func (this *AddSlide) GetPosition() int32 {
	return this.Position
}

func (this *AddSlide) SetPosition(newValue int32) {
	this.Position = newValue
}
func (this *AddSlide) GetLayoutAlias() string {
	return this.LayoutAlias
}

func (this *AddSlide) SetLayoutAlias(newValue string) {
	this.LayoutAlias = newValue
}

func (this *AddSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AddSlide"
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
	
	if valCloneFromFile, ok := GetMapValue(objMap, "cloneFromFile"); ok {
		if valCloneFromFile != nil {
			var valueForCloneFromFile InputFile
			err = json.Unmarshal(*valCloneFromFile, &valueForCloneFromFile)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("InputFile", *valCloneFromFile)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCloneFromFile, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IInputFile)
			if ok {
				this.CloneFromFile = vInterfaceObject
			}
		}
	}
	
	if valCloneFromPosition, ok := GetMapValue(objMap, "cloneFromPosition"); ok {
		if valCloneFromPosition != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPosition, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}
	
	if valPosition, ok := GetMapValue(objMap, "position"); ok {
		if valPosition != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}
	
	if valLayoutAlias, ok := GetMapValue(objMap, "layoutAlias"); ok {
		if valLayoutAlias != nil {
			var valueForLayoutAlias string
			err = json.Unmarshal(*valLayoutAlias, &valueForLayoutAlias)
			if err != nil {
				return err
			}
			this.LayoutAlias = valueForLayoutAlias
		}
	}

	return nil
}
