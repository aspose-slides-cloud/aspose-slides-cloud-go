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
	getAlignment() string
	setAlignment(newValue string)

	// Cap style.
	getCapStyle() string
	setCapStyle(newValue string)

	// Dash style.
	getDashStyle() string
	setDashStyle(newValue string)

	// Join style.
	getJoinStyle() string
	setJoinStyle(newValue string)

	// Style.
	getStyle() string
	setStyle(newValue string)

	// Begin arrowhead.
	getBeginArrowHead() IArrowHeadProperties
	setBeginArrowHead(newValue IArrowHeadProperties)

	// End arrowhead.
	getEndArrowHead() IArrowHeadProperties
	setEndArrowHead(newValue IArrowHeadProperties)

	// Custom dash pattern.
	getCustomDashPattern() ICustomDashPattern
	setCustomDashPattern(newValue ICustomDashPattern)

	// Fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Miter limit.
	getMiterLimit() float64
	setMiterLimit(newValue float64)

	// Width.
	getWidth() float64
	setWidth(newValue float64)
}

type LineFormat struct {

	// Alignment.
	Alignment string `json:"Alignment"`

	// Cap style.
	CapStyle string `json:"CapStyle"`

	// Dash style.
	DashStyle string `json:"DashStyle"`

	// Join style.
	JoinStyle string `json:"JoinStyle"`

	// Style.
	Style string `json:"Style"`

	// Begin arrowhead.
	BeginArrowHead IArrowHeadProperties `json:"BeginArrowHead,omitempty"`

	// End arrowhead.
	EndArrowHead IArrowHeadProperties `json:"EndArrowHead,omitempty"`

	// Custom dash pattern.
	CustomDashPattern ICustomDashPattern `json:"CustomDashPattern,omitempty"`

	// Fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Miter limit.
	MiterLimit float64 `json:"MiterLimit"`

	// Width.
	Width float64 `json:"Width"`
}

func (this LineFormat) getAlignment() string {
	return this.Alignment
}

func (this LineFormat) setAlignment(newValue string) {
	this.Alignment = newValue
}
func (this LineFormat) getCapStyle() string {
	return this.CapStyle
}

func (this LineFormat) setCapStyle(newValue string) {
	this.CapStyle = newValue
}
func (this LineFormat) getDashStyle() string {
	return this.DashStyle
}

func (this LineFormat) setDashStyle(newValue string) {
	this.DashStyle = newValue
}
func (this LineFormat) getJoinStyle() string {
	return this.JoinStyle
}

func (this LineFormat) setJoinStyle(newValue string) {
	this.JoinStyle = newValue
}
func (this LineFormat) getStyle() string {
	return this.Style
}

func (this LineFormat) setStyle(newValue string) {
	this.Style = newValue
}
func (this LineFormat) getBeginArrowHead() IArrowHeadProperties {
	return this.BeginArrowHead
}

func (this LineFormat) setBeginArrowHead(newValue IArrowHeadProperties) {
	this.BeginArrowHead = newValue
}
func (this LineFormat) getEndArrowHead() IArrowHeadProperties {
	return this.EndArrowHead
}

func (this LineFormat) setEndArrowHead(newValue IArrowHeadProperties) {
	this.EndArrowHead = newValue
}
func (this LineFormat) getCustomDashPattern() ICustomDashPattern {
	return this.CustomDashPattern
}

func (this LineFormat) setCustomDashPattern(newValue ICustomDashPattern) {
	this.CustomDashPattern = newValue
}
func (this LineFormat) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this LineFormat) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this LineFormat) getMiterLimit() float64 {
	return this.MiterLimit
}

func (this LineFormat) setMiterLimit(newValue float64) {
	this.MiterLimit = newValue
}
func (this LineFormat) getWidth() float64 {
	return this.Width
}

func (this LineFormat) setWidth(newValue float64) {
	this.Width = newValue
}

func (this *LineFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Alignment = "Center"
	if valAlignment, ok := objMap["alignment"]; ok {
		if valAlignment != nil {
			var valueForAlignment string
			err = json.Unmarshal(*valAlignment, &valueForAlignment)
			if err != nil {
				return err
			}
			this.Alignment = valueForAlignment
		}
	}
	if valAlignmentCap, ok := objMap["Alignment"]; ok {
		if valAlignmentCap != nil {
			var valueForAlignment string
			err = json.Unmarshal(*valAlignmentCap, &valueForAlignment)
			if err != nil {
				return err
			}
			this.Alignment = valueForAlignment
		}
	}
	this.CapStyle = "Round"
	if valCapStyle, ok := objMap["capStyle"]; ok {
		if valCapStyle != nil {
			var valueForCapStyle string
			err = json.Unmarshal(*valCapStyle, &valueForCapStyle)
			if err != nil {
				return err
			}
			this.CapStyle = valueForCapStyle
		}
	}
	if valCapStyleCap, ok := objMap["CapStyle"]; ok {
		if valCapStyleCap != nil {
			var valueForCapStyle string
			err = json.Unmarshal(*valCapStyleCap, &valueForCapStyle)
			if err != nil {
				return err
			}
			this.CapStyle = valueForCapStyle
		}
	}
	this.DashStyle = "Solid"
	if valDashStyle, ok := objMap["dashStyle"]; ok {
		if valDashStyle != nil {
			var valueForDashStyle string
			err = json.Unmarshal(*valDashStyle, &valueForDashStyle)
			if err != nil {
				return err
			}
			this.DashStyle = valueForDashStyle
		}
	}
	if valDashStyleCap, ok := objMap["DashStyle"]; ok {
		if valDashStyleCap != nil {
			var valueForDashStyle string
			err = json.Unmarshal(*valDashStyleCap, &valueForDashStyle)
			if err != nil {
				return err
			}
			this.DashStyle = valueForDashStyle
		}
	}
	this.JoinStyle = "Round"
	if valJoinStyle, ok := objMap["joinStyle"]; ok {
		if valJoinStyle != nil {
			var valueForJoinStyle string
			err = json.Unmarshal(*valJoinStyle, &valueForJoinStyle)
			if err != nil {
				return err
			}
			this.JoinStyle = valueForJoinStyle
		}
	}
	if valJoinStyleCap, ok := objMap["JoinStyle"]; ok {
		if valJoinStyleCap != nil {
			var valueForJoinStyle string
			err = json.Unmarshal(*valJoinStyleCap, &valueForJoinStyle)
			if err != nil {
				return err
			}
			this.JoinStyle = valueForJoinStyle
		}
	}
	this.Style = "Single"
	if valStyle, ok := objMap["style"]; ok {
		if valStyle != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyle, &valueForStyle)
			if err != nil {
				return err
			}
			this.Style = valueForStyle
		}
	}
	if valStyleCap, ok := objMap["Style"]; ok {
		if valStyleCap != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyleCap, &valueForStyle)
			if err != nil {
				return err
			}
			this.Style = valueForStyle
		}
	}
	
	if valBeginArrowHead, ok := objMap["beginArrowHead"]; ok {
		if valBeginArrowHead != nil {
			var valueForBeginArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valBeginArrowHead, &valueForBeginArrowHead)
			if err != nil {
				return err
			}
			this.BeginArrowHead = valueForBeginArrowHead
		}
	}
	if valBeginArrowHeadCap, ok := objMap["BeginArrowHead"]; ok {
		if valBeginArrowHeadCap != nil {
			var valueForBeginArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valBeginArrowHeadCap, &valueForBeginArrowHead)
			if err != nil {
				return err
			}
			this.BeginArrowHead = valueForBeginArrowHead
		}
	}
	
	if valEndArrowHead, ok := objMap["endArrowHead"]; ok {
		if valEndArrowHead != nil {
			var valueForEndArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valEndArrowHead, &valueForEndArrowHead)
			if err != nil {
				return err
			}
			this.EndArrowHead = valueForEndArrowHead
		}
	}
	if valEndArrowHeadCap, ok := objMap["EndArrowHead"]; ok {
		if valEndArrowHeadCap != nil {
			var valueForEndArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valEndArrowHeadCap, &valueForEndArrowHead)
			if err != nil {
				return err
			}
			this.EndArrowHead = valueForEndArrowHead
		}
	}
	
	if valCustomDashPattern, ok := objMap["customDashPattern"]; ok {
		if valCustomDashPattern != nil {
			var valueForCustomDashPattern CustomDashPattern
			err = json.Unmarshal(*valCustomDashPattern, &valueForCustomDashPattern)
			if err != nil {
				return err
			}
			this.CustomDashPattern = valueForCustomDashPattern
		}
	}
	if valCustomDashPatternCap, ok := objMap["CustomDashPattern"]; ok {
		if valCustomDashPatternCap != nil {
			var valueForCustomDashPattern CustomDashPattern
			err = json.Unmarshal(*valCustomDashPatternCap, &valueForCustomDashPattern)
			if err != nil {
				return err
			}
			this.CustomDashPattern = valueForCustomDashPattern
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}
	
	if valMiterLimit, ok := objMap["miterLimit"]; ok {
		if valMiterLimit != nil {
			var valueForMiterLimit float64
			err = json.Unmarshal(*valMiterLimit, &valueForMiterLimit)
			if err != nil {
				return err
			}
			this.MiterLimit = valueForMiterLimit
		}
	}
	if valMiterLimitCap, ok := objMap["MiterLimit"]; ok {
		if valMiterLimitCap != nil {
			var valueForMiterLimit float64
			err = json.Unmarshal(*valMiterLimitCap, &valueForMiterLimit)
			if err != nil {
				return err
			}
			this.MiterLimit = valueForMiterLimit
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

    return nil
}
