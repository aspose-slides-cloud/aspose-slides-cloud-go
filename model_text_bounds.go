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

// Represents text bounds within a paragraph or portion.
type ITextBounds interface {

	// X coordinate of the text bounds.
	GetX() float64
	SetX(newValue float64)

	// X coordinate of the text bounds.             
	GetY() float64
	SetY(newValue float64)

	// Width of the text bounds.
	GetWidth() float64
	SetWidth(newValue float64)

	// Height of the text bounds.
	GetHeight() float64
	SetHeight(newValue float64)
}

type TextBounds struct {

	// X coordinate of the text bounds.
	X float64 `json:"X"`

	// X coordinate of the text bounds.             
	Y float64 `json:"Y"`

	// Width of the text bounds.
	Width float64 `json:"Width"`

	// Height of the text bounds.
	Height float64 `json:"Height"`
}

func NewTextBounds() *TextBounds {
	instance := new(TextBounds)
	return instance
}

func (this *TextBounds) GetX() float64 {
	return this.X
}

func (this *TextBounds) SetX(newValue float64) {
	this.X = newValue
}
func (this *TextBounds) GetY() float64 {
	return this.Y
}

func (this *TextBounds) SetY(newValue float64) {
	this.Y = newValue
}
func (this *TextBounds) GetWidth() float64 {
	return this.Width
}

func (this *TextBounds) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *TextBounds) GetHeight() float64 {
	return this.Height
}

func (this *TextBounds) SetHeight(newValue float64) {
	this.Height = newValue
}

func (this *TextBounds) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valX, ok := GetMapValue(objMap, "x"); ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := GetMapValue(objMap, "y"); ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	return nil
}
