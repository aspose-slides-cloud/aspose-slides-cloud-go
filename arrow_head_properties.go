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


type IArrowHeadProperties interface {

	getLength() string
	setLength(newValue string)

	getStyle() string
	setStyle(newValue string)

	getWidth() string
	setWidth(newValue string)
}

type ArrowHeadProperties struct {

	Length string `json:"Length"`

	Style string `json:"Style"`

	Width string `json:"Width"`
}

func (this ArrowHeadProperties) getLength() string {
	return this.Length
}

func (this ArrowHeadProperties) setLength(newValue string) {
	this.Length = newValue
}
func (this ArrowHeadProperties) getStyle() string {
	return this.Style
}

func (this ArrowHeadProperties) setStyle(newValue string) {
	this.Style = newValue
}
func (this ArrowHeadProperties) getWidth() string {
	return this.Width
}

func (this ArrowHeadProperties) setWidth(newValue string) {
	this.Width = newValue
}

func (this *ArrowHeadProperties) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valLength, ok := objMap["Length"]; ok {
		if valLength != nil {
			var valueForLength string
			err = json.Unmarshal(*valLength, &valueForLength)
			if err != nil {
				return err
			}
			this.Length = valueForLength
		}
	}

	if valStyle, ok := objMap["Style"]; ok {
		if valStyle != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyle, &valueForStyle)
			if err != nil {
				return err
			}
			this.Style = valueForStyle
		}
	}

	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth string
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

    return nil
}
