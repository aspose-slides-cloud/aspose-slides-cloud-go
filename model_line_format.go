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

// Line format.
type ILineFormat interface {

	// Alignment.
	GetAlignment() string
	SetAlignment(newValue string)

	// Cap style.
	GetCapStyle() string
	SetCapStyle(newValue string)

	// Dash style.
	GetDashStyle() string
	SetDashStyle(newValue string)

	// Join style.
	GetJoinStyle() string
	SetJoinStyle(newValue string)

	// Style.
	GetStyle() string
	SetStyle(newValue string)

	// Sketch type.
	GetSketchType() string
	SetSketchType(newValue string)

	// Begin arrowhead.
	GetBeginArrowHead() IArrowHeadProperties
	SetBeginArrowHead(newValue IArrowHeadProperties)

	// End arrowhead.
	GetEndArrowHead() IArrowHeadProperties
	SetEndArrowHead(newValue IArrowHeadProperties)

	// Custom dash pattern.
	GetCustomDashPattern() ICustomDashPattern
	SetCustomDashPattern(newValue ICustomDashPattern)

	// Fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Miter limit.
	GetMiterLimit() float64
	SetMiterLimit(newValue float64)

	// Width.
	GetWidth() float64
	SetWidth(newValue float64)
}

type LineFormat struct {

	// Alignment.
	Alignment string `json:"Alignment,omitempty"`

	// Cap style.
	CapStyle string `json:"CapStyle,omitempty"`

	// Dash style.
	DashStyle string `json:"DashStyle,omitempty"`

	// Join style.
	JoinStyle string `json:"JoinStyle,omitempty"`

	// Style.
	Style string `json:"Style,omitempty"`

	// Sketch type.
	SketchType string `json:"SketchType,omitempty"`

	// Begin arrowhead.
	BeginArrowHead IArrowHeadProperties `json:"BeginArrowHead,omitempty"`

	// End arrowhead.
	EndArrowHead IArrowHeadProperties `json:"EndArrowHead,omitempty"`

	// Custom dash pattern.
	CustomDashPattern ICustomDashPattern `json:"CustomDashPattern,omitempty"`

	// Fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Miter limit.
	MiterLimit float64 `json:"MiterLimit,omitempty"`

	// Width.
	Width float64 `json:"Width,omitempty"`
}

func NewLineFormat() *LineFormat {
	instance := new(LineFormat)
	return instance
}

func (this *LineFormat) GetAlignment() string {
	return this.Alignment
}

func (this *LineFormat) SetAlignment(newValue string) {
	this.Alignment = newValue
}
func (this *LineFormat) GetCapStyle() string {
	return this.CapStyle
}

func (this *LineFormat) SetCapStyle(newValue string) {
	this.CapStyle = newValue
}
func (this *LineFormat) GetDashStyle() string {
	return this.DashStyle
}

func (this *LineFormat) SetDashStyle(newValue string) {
	this.DashStyle = newValue
}
func (this *LineFormat) GetJoinStyle() string {
	return this.JoinStyle
}

func (this *LineFormat) SetJoinStyle(newValue string) {
	this.JoinStyle = newValue
}
func (this *LineFormat) GetStyle() string {
	return this.Style
}

func (this *LineFormat) SetStyle(newValue string) {
	this.Style = newValue
}
func (this *LineFormat) GetSketchType() string {
	return this.SketchType
}

func (this *LineFormat) SetSketchType(newValue string) {
	this.SketchType = newValue
}
func (this *LineFormat) GetBeginArrowHead() IArrowHeadProperties {
	return this.BeginArrowHead
}

func (this *LineFormat) SetBeginArrowHead(newValue IArrowHeadProperties) {
	this.BeginArrowHead = newValue
}
func (this *LineFormat) GetEndArrowHead() IArrowHeadProperties {
	return this.EndArrowHead
}

func (this *LineFormat) SetEndArrowHead(newValue IArrowHeadProperties) {
	this.EndArrowHead = newValue
}
func (this *LineFormat) GetCustomDashPattern() ICustomDashPattern {
	return this.CustomDashPattern
}

func (this *LineFormat) SetCustomDashPattern(newValue ICustomDashPattern) {
	this.CustomDashPattern = newValue
}
func (this *LineFormat) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *LineFormat) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *LineFormat) GetMiterLimit() float64 {
	return this.MiterLimit
}

func (this *LineFormat) SetMiterLimit(newValue float64) {
	this.MiterLimit = newValue
}
func (this *LineFormat) GetWidth() float64 {
	return this.Width
}

func (this *LineFormat) SetWidth(newValue float64) {
	this.Width = newValue
}

func (this *LineFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valAlignment, ok := GetMapValue(objMap, "alignment"); ok {
		if valAlignment != nil {
			var valueForAlignment string
			err = json.Unmarshal(*valAlignment, &valueForAlignment)
			if err != nil {
				var valueForAlignmentInt int32
				err = json.Unmarshal(*valAlignment, &valueForAlignmentInt)
				if err != nil {
					return err
				}
				this.Alignment = string(valueForAlignmentInt)
			} else {
				this.Alignment = valueForAlignment
			}
		}
	}
	
	if valCapStyle, ok := GetMapValue(objMap, "capStyle"); ok {
		if valCapStyle != nil {
			var valueForCapStyle string
			err = json.Unmarshal(*valCapStyle, &valueForCapStyle)
			if err != nil {
				var valueForCapStyleInt int32
				err = json.Unmarshal(*valCapStyle, &valueForCapStyleInt)
				if err != nil {
					return err
				}
				this.CapStyle = string(valueForCapStyleInt)
			} else {
				this.CapStyle = valueForCapStyle
			}
		}
	}
	
	if valDashStyle, ok := GetMapValue(objMap, "dashStyle"); ok {
		if valDashStyle != nil {
			var valueForDashStyle string
			err = json.Unmarshal(*valDashStyle, &valueForDashStyle)
			if err != nil {
				var valueForDashStyleInt int32
				err = json.Unmarshal(*valDashStyle, &valueForDashStyleInt)
				if err != nil {
					return err
				}
				this.DashStyle = string(valueForDashStyleInt)
			} else {
				this.DashStyle = valueForDashStyle
			}
		}
	}
	
	if valJoinStyle, ok := GetMapValue(objMap, "joinStyle"); ok {
		if valJoinStyle != nil {
			var valueForJoinStyle string
			err = json.Unmarshal(*valJoinStyle, &valueForJoinStyle)
			if err != nil {
				var valueForJoinStyleInt int32
				err = json.Unmarshal(*valJoinStyle, &valueForJoinStyleInt)
				if err != nil {
					return err
				}
				this.JoinStyle = string(valueForJoinStyleInt)
			} else {
				this.JoinStyle = valueForJoinStyle
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
	
	if valSketchType, ok := GetMapValue(objMap, "sketchType"); ok {
		if valSketchType != nil {
			var valueForSketchType string
			err = json.Unmarshal(*valSketchType, &valueForSketchType)
			if err != nil {
				var valueForSketchTypeInt int32
				err = json.Unmarshal(*valSketchType, &valueForSketchTypeInt)
				if err != nil {
					return err
				}
				this.SketchType = string(valueForSketchTypeInt)
			} else {
				this.SketchType = valueForSketchType
			}
		}
	}
	
	if valBeginArrowHead, ok := GetMapValue(objMap, "beginArrowHead"); ok {
		if valBeginArrowHead != nil {
			var valueForBeginArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valBeginArrowHead, &valueForBeginArrowHead)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ArrowHeadProperties", *valBeginArrowHead)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBeginArrowHead, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IArrowHeadProperties)
			if ok {
				this.BeginArrowHead = vInterfaceObject
			}
		}
	}
	
	if valEndArrowHead, ok := GetMapValue(objMap, "endArrowHead"); ok {
		if valEndArrowHead != nil {
			var valueForEndArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valEndArrowHead, &valueForEndArrowHead)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ArrowHeadProperties", *valEndArrowHead)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEndArrowHead, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IArrowHeadProperties)
			if ok {
				this.EndArrowHead = vInterfaceObject
			}
		}
	}
	
	if valCustomDashPattern, ok := GetMapValue(objMap, "customDashPattern"); ok {
		if valCustomDashPattern != nil {
			var valueForCustomDashPattern CustomDashPattern
			err = json.Unmarshal(*valCustomDashPattern, &valueForCustomDashPattern)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("CustomDashPattern", *valCustomDashPattern)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCustomDashPattern, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ICustomDashPattern)
			if ok {
				this.CustomDashPattern = vInterfaceObject
			}
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
	
	if valMiterLimit, ok := GetMapValue(objMap, "miterLimit"); ok {
		if valMiterLimit != nil {
			var valueForMiterLimit float64
			err = json.Unmarshal(*valMiterLimit, &valueForMiterLimit)
			if err != nil {
				return err
			}
			this.MiterLimit = valueForMiterLimit
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

	return nil
}
