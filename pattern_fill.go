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

// Represents Pattern Fill
type IPatternFill interface {

	// Type of fill.
	getType() string
	setType(newValue string)

	// Gets or sets the back color of the pattern fill.
	getBackColor() string
	setBackColor(newValue string)

	// Gets or sets the fore color of the pattern fill.
	getForeColor() string
	setForeColor(newValue string)

	// Gets or sets the style of pattern fill.
	getStyle() string
	setStyle(newValue string)
}

type PatternFill struct {

	// Type of fill.
	Type_ string `json:"Type"`

	// Gets or sets the back color of the pattern fill.
	BackColor string `json:"BackColor,omitempty"`

	// Gets or sets the fore color of the pattern fill.
	ForeColor string `json:"ForeColor,omitempty"`

	// Gets or sets the style of pattern fill.
	Style string `json:"Style"`
}

func (this PatternFill) getType() string {
	return this.Type_
}

func (this PatternFill) setType(newValue string) {
	this.Type_ = newValue
}
func (this PatternFill) getBackColor() string {
	return this.BackColor
}

func (this PatternFill) setBackColor(newValue string) {
	this.BackColor = newValue
}
func (this PatternFill) getForeColor() string {
	return this.ForeColor
}

func (this PatternFill) setForeColor(newValue string) {
	this.ForeColor = newValue
}
func (this PatternFill) getStyle() string {
	return this.Style
}

func (this PatternFill) setStyle(newValue string) {
	this.Style = newValue
}

func (this *PatternFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valBackColor, ok := objMap["BackColor"]; ok {
		if valBackColor != nil {
			var valueForBackColor string
			err = json.Unmarshal(*valBackColor, &valueForBackColor)
			if err != nil {
				return err
			}
			this.BackColor = valueForBackColor
		}
	}

	if valForeColor, ok := objMap["ForeColor"]; ok {
		if valForeColor != nil {
			var valueForForeColor string
			err = json.Unmarshal(*valForeColor, &valueForForeColor)
			if err != nil {
				return err
			}
			this.ForeColor = valueForForeColor
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

    return nil
}
