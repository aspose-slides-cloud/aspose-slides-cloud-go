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
	GetType() string
	SetType(newValue string)

	// Gets or sets the back color of the pattern fill.
	GetBackColor() string
	SetBackColor(newValue string)

	// Gets or sets the fore color of the pattern fill.
	GetForeColor() string
	SetForeColor(newValue string)

	// Gets or sets the style of pattern fill.
	GetStyle() string
	SetStyle(newValue string)
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

func NewPatternFill() *PatternFill {
	instance := new(PatternFill)
	instance.Type_ = "Pattern"
	instance.Style = "Unknown"
	return instance
}

func (this *PatternFill) GetType() string {
	return this.Type_
}

func (this *PatternFill) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *PatternFill) GetBackColor() string {
	return this.BackColor
}

func (this *PatternFill) SetBackColor(newValue string) {
	this.BackColor = newValue
}
func (this *PatternFill) GetForeColor() string {
	return this.ForeColor
}

func (this *PatternFill) SetForeColor(newValue string) {
	this.ForeColor = newValue
}
func (this *PatternFill) GetStyle() string {
	return this.Style
}

func (this *PatternFill) SetStyle(newValue string) {
	this.Style = newValue
}

func (this *PatternFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Pattern"
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
	
	if valBackColor, ok := objMap["backColor"]; ok {
		if valBackColor != nil {
			var valueForBackColor string
			err = json.Unmarshal(*valBackColor, &valueForBackColor)
			if err != nil {
				return err
			}
			this.BackColor = valueForBackColor
		}
	}
	if valBackColorCap, ok := objMap["BackColor"]; ok {
		if valBackColorCap != nil {
			var valueForBackColor string
			err = json.Unmarshal(*valBackColorCap, &valueForBackColor)
			if err != nil {
				return err
			}
			this.BackColor = valueForBackColor
		}
	}
	
	if valForeColor, ok := objMap["foreColor"]; ok {
		if valForeColor != nil {
			var valueForForeColor string
			err = json.Unmarshal(*valForeColor, &valueForForeColor)
			if err != nil {
				return err
			}
			this.ForeColor = valueForForeColor
		}
	}
	if valForeColorCap, ok := objMap["ForeColor"]; ok {
		if valForeColorCap != nil {
			var valueForForeColor string
			err = json.Unmarshal(*valForeColorCap, &valueForForeColor)
			if err != nil {
				return err
			}
			this.ForeColor = valueForForeColor
		}
	}
	this.Style = "Unknown"
	if valStyle, ok := objMap["style"]; ok {
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
	if valStyleCap, ok := objMap["Style"]; ok {
		if valStyleCap != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyleCap, &valueForStyle)
			if err != nil {
				var valueForStyleInt int32
				err = json.Unmarshal(*valStyleCap, &valueForStyleInt)
				if err != nil {
					return err
				}
				this.Style = string(valueForStyleInt)
			} else {
				this.Style = valueForStyle
			}
		}
	}

	return nil
}
