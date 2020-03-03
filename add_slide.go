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
	getType() string
	setType(newValue string)

	// File to clone a slide from.
	getCloneFromFile() IInputFile
	setCloneFromFile(newValue IInputFile)

	// Position of the slide to clone.
	getCloneFromPosition() int32
	setCloneFromPosition(newValue int32)

	// Position at which to insert the slide.
	getPosition() int32
	setPosition(newValue int32)

	// Alias of layout (href, index or type). If value is null a blank slide is added.
	getLayoutAlias() string
	setLayoutAlias(newValue string)
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

func (this *AddSlide) getType() string {
	return this.Type_
}

func (this *AddSlide) setType(newValue string) {
	this.Type_ = newValue
}
func (this *AddSlide) getCloneFromFile() IInputFile {
	return this.CloneFromFile
}

func (this *AddSlide) setCloneFromFile(newValue IInputFile) {
	this.CloneFromFile = newValue
}
func (this *AddSlide) getCloneFromPosition() int32 {
	return this.CloneFromPosition
}

func (this *AddSlide) setCloneFromPosition(newValue int32) {
	this.CloneFromPosition = newValue
}
func (this *AddSlide) getPosition() int32 {
	return this.Position
}

func (this *AddSlide) setPosition(newValue int32) {
	this.Position = newValue
}
func (this *AddSlide) getLayoutAlias() string {
	return this.LayoutAlias
}

func (this *AddSlide) setLayoutAlias(newValue string) {
	this.LayoutAlias = newValue
}

func (this *AddSlide) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "AddSlide"
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
	
	if valCloneFromFile, ok := objMap["cloneFromFile"]; ok {
		if valCloneFromFile != nil {
			var valueForCloneFromFile InputFile
			err = json.Unmarshal(*valCloneFromFile, &valueForCloneFromFile)
			if err != nil {
				return err
			}
			this.CloneFromFile = &valueForCloneFromFile
		}
	}
	if valCloneFromFileCap, ok := objMap["CloneFromFile"]; ok {
		if valCloneFromFileCap != nil {
			var valueForCloneFromFile InputFile
			err = json.Unmarshal(*valCloneFromFileCap, &valueForCloneFromFile)
			if err != nil {
				return err
			}
			this.CloneFromFile = &valueForCloneFromFile
		}
	}
	
	if valCloneFromPosition, ok := objMap["cloneFromPosition"]; ok {
		if valCloneFromPosition != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPosition, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}
	if valCloneFromPositionCap, ok := objMap["CloneFromPosition"]; ok {
		if valCloneFromPositionCap != nil {
			var valueForCloneFromPosition int32
			err = json.Unmarshal(*valCloneFromPositionCap, &valueForCloneFromPosition)
			if err != nil {
				return err
			}
			this.CloneFromPosition = valueForCloneFromPosition
		}
	}
	
	if valPosition, ok := objMap["position"]; ok {
		if valPosition != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}
	if valPositionCap, ok := objMap["Position"]; ok {
		if valPositionCap != nil {
			var valueForPosition int32
			err = json.Unmarshal(*valPositionCap, &valueForPosition)
			if err != nil {
				return err
			}
			this.Position = valueForPosition
		}
	}
	
	if valLayoutAlias, ok := objMap["layoutAlias"]; ok {
		if valLayoutAlias != nil {
			var valueForLayoutAlias string
			err = json.Unmarshal(*valLayoutAlias, &valueForLayoutAlias)
			if err != nil {
				return err
			}
			this.LayoutAlias = valueForLayoutAlias
		}
	}
	if valLayoutAliasCap, ok := objMap["LayoutAlias"]; ok {
		if valLayoutAliasCap != nil {
			var valueForLayoutAlias string
			err = json.Unmarshal(*valLayoutAliasCap, &valueForLayoutAlias)
			if err != nil {
				return err
			}
			this.LayoutAlias = valueForLayoutAlias
		}
	}

    return nil
}
