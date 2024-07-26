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

// Represents chart title
type IChartTitle interface {

	// Get or sets the text.
	GetText() string
	SetText(newValue string)

	// the X location
	GetX() float64
	SetX(newValue float64)

	// the Y location
	GetY() float64
	SetY(newValue float64)

	// Width
	GetWidth() float64
	SetWidth(newValue float64)

	// Height
	GetHeight() float64
	SetHeight(newValue float64)

	// true if other elements are allowed to overlay the legend
	GetOverlay() *bool
	SetOverlay(newValue *bool)

	// Get or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)
}

type ChartTitle struct {

	// Get or sets the text.
	Text string `json:"Text,omitempty"`

	// the X location
	X float64 `json:"X,omitempty"`

	// the Y location
	Y float64 `json:"Y,omitempty"`

	// Width
	Width float64 `json:"Width,omitempty"`

	// Height
	Height float64 `json:"Height,omitempty"`

	// true if other elements are allowed to overlay the legend
	Overlay *bool `json:"Overlay"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func NewChartTitle() *ChartTitle {
	instance := new(ChartTitle)
	return instance
}

func (this *ChartTitle) GetText() string {
	return this.Text
}

func (this *ChartTitle) SetText(newValue string) {
	this.Text = newValue
}
func (this *ChartTitle) GetX() float64 {
	return this.X
}

func (this *ChartTitle) SetX(newValue float64) {
	this.X = newValue
}
func (this *ChartTitle) GetY() float64 {
	return this.Y
}

func (this *ChartTitle) SetY(newValue float64) {
	this.Y = newValue
}
func (this *ChartTitle) GetWidth() float64 {
	return this.Width
}

func (this *ChartTitle) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *ChartTitle) GetHeight() float64 {
	return this.Height
}

func (this *ChartTitle) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *ChartTitle) GetOverlay() *bool {
	return this.Overlay
}

func (this *ChartTitle) SetOverlay(newValue *bool) {
	this.Overlay = newValue
}
func (this *ChartTitle) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ChartTitle) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ChartTitle) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ChartTitle) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ChartTitle) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ChartTitle) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}

func (this *ChartTitle) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valText, ok := GetMapValue(objMap, "text"); ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
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
	
	if valOverlay, ok := GetMapValue(objMap, "overlay"); ok {
		if valOverlay != nil {
			var valueForOverlay *bool
			err = json.Unmarshal(*valOverlay, &valueForOverlay)
			if err != nil {
				return err
			}
			this.Overlay = valueForOverlay
		}
	}
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}

	return nil
}
