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

// Arrow head properties.
type IArrowHeadProperties interface {

	// Length.
	GetLength() string
	SetLength(newValue string)

	// Style.
	GetStyle() string
	SetStyle(newValue string)

	// Width.
	GetWidth() string
	SetWidth(newValue string)
}

type ArrowHeadProperties struct {

	// Length.
	Length string `json:"Length,omitempty"`

	// Style.
	Style string `json:"Style,omitempty"`

	// Width.
	Width string `json:"Width,omitempty"`
}

func NewArrowHeadProperties() *ArrowHeadProperties {
	instance := new(ArrowHeadProperties)
	return instance
}

func (this *ArrowHeadProperties) GetLength() string {
	return this.Length
}

func (this *ArrowHeadProperties) SetLength(newValue string) {
	this.Length = newValue
}
func (this *ArrowHeadProperties) GetStyle() string {
	return this.Style
}

func (this *ArrowHeadProperties) SetStyle(newValue string) {
	this.Style = newValue
}
func (this *ArrowHeadProperties) GetWidth() string {
	return this.Width
}

func (this *ArrowHeadProperties) SetWidth(newValue string) {
	this.Width = newValue
}

func (this *ArrowHeadProperties) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valLength, ok := GetMapValue(objMap, "length"); ok {
		if valLength != nil {
			var valueForLength string
			err = json.Unmarshal(*valLength, &valueForLength)
			if err != nil {
				var valueForLengthInt int32
				err = json.Unmarshal(*valLength, &valueForLengthInt)
				if err != nil {
					return err
				}
				this.Length = string(valueForLengthInt)
			} else {
				this.Length = valueForLength
			}
		}
	}
	
	if valStyle, ok := GetMapValue(objMap, "style"); ok {
		if valStyle != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyle, &valueForStyle)
			if err != nil {
				var valueForStyleInt int32
				err = json.Unmarshal(*valStyle, &valueForStyleInt)
				if err != nil {
					return err
				}
				this.Style = string(valueForStyleInt)
			} else {
				this.Style = valueForStyle
			}
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth string
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				var valueForWidthInt int32
				err = json.Unmarshal(*valWidth, &valueForWidthInt)
				if err != nil {
					return err
				}
				this.Width = string(valueForWidthInt)
			} else {
				this.Width = valueForWidth
			}
		}
	}

	return nil
}
