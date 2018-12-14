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

	getLength() LineArrowheadLength
	setLength(newValue LineArrowheadLength)

	getStyle() LineArrowheadStyle
	setStyle(newValue LineArrowheadStyle)

	getWidth() LineArrowheadWidth
	setWidth(newValue LineArrowheadWidth)
}

type ArrowHeadProperties struct {

	Length LineArrowheadLength `json:"Length"`

	Style LineArrowheadStyle `json:"Style"`

	Width LineArrowheadWidth `json:"Width"`
}

func (this ArrowHeadProperties) getLength() LineArrowheadLength {
	return this.Length
}

func (this ArrowHeadProperties) setLength(newValue LineArrowheadLength) {
	this.Length = newValue
}
func (this ArrowHeadProperties) getStyle() LineArrowheadStyle {
	return this.Style
}

func (this ArrowHeadProperties) setStyle(newValue LineArrowheadStyle) {
	this.Style = newValue
}
func (this ArrowHeadProperties) getWidth() LineArrowheadWidth {
	return this.Width
}

func (this ArrowHeadProperties) setWidth(newValue LineArrowheadWidth) {
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
			var valueForLength LineArrowheadLength
			err = json.Unmarshal(*valLength, &valueForLength)
			if err != nil {
				return err
			}
			this.Length = valueForLength
		}
	}

	if valStyle, ok := objMap["Style"]; ok {
		if valStyle != nil {
			var valueForStyle LineArrowheadStyle
			err = json.Unmarshal(*valStyle, &valueForStyle)
			if err != nil {
				return err
			}
			this.Style = valueForStyle
		}
	}

	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth LineArrowheadWidth
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

    return nil
}
