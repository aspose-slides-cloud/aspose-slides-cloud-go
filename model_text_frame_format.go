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

// Contains the TextFrame's formatting properties.
type ITextFrameFormat interface {

	// Represents 3d effect properties for a text.
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets text wrapping shape.
	GetTransform() string
	SetTransform(newValue string)
}

type TextFrameFormat struct {

	// Represents 3d effect properties for a text.
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets text wrapping shape.
	Transform string `json:"Transform,omitempty"`
}

func NewTextFrameFormat() *TextFrameFormat {
	instance := new(TextFrameFormat)
	return instance
}

func (this *TextFrameFormat) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *TextFrameFormat) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *TextFrameFormat) GetTransform() string {
	return this.Transform
}

func (this *TextFrameFormat) SetTransform(newValue string) {
	this.Transform = newValue
}

func (this *TextFrameFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valThreeDFormat, ok := objMap["threeDFormat"]; ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	if valThreeDFormatCap, ok := objMap["ThreeDFormat"]; ok {
		if valThreeDFormatCap != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormatCap, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	
	if valTransform, ok := objMap["transform"]; ok {
		if valTransform != nil {
			var valueForTransform string
			err = json.Unmarshal(*valTransform, &valueForTransform)
			if err != nil {
				var valueForTransformInt int32
				err = json.Unmarshal(*valTransform, &valueForTransformInt)
				if err != nil {
					return err
				}
				this.Transform = string(valueForTransformInt)
			} else {
				this.Transform = valueForTransform
			}
		}
	}
	if valTransformCap, ok := objMap["Transform"]; ok {
		if valTransformCap != nil {
			var valueForTransform string
			err = json.Unmarshal(*valTransformCap, &valueForTransform)
			if err != nil {
				var valueForTransformInt int32
				err = json.Unmarshal(*valTransformCap, &valueForTransformInt)
				if err != nil {
					return err
				}
				this.Transform = string(valueForTransformInt)
			} else {
				this.Transform = valueForTransform
			}
		}
	}

	return nil
}
