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


type ILineFormat interface {

	getAlignment() string
	setAlignment(newValue string)

	getCapStyle() string
	setCapStyle(newValue string)

	getDashStyle() string
	setDashStyle(newValue string)

	getJoinStyle() string
	setJoinStyle(newValue string)

	getStyle() string
	setStyle(newValue string)

	getBeginArrowHead() IArrowHeadProperties
	setBeginArrowHead(newValue IArrowHeadProperties)

	getEndArrowHead() IArrowHeadProperties
	setEndArrowHead(newValue IArrowHeadProperties)

	getCustomDashPattern() ICustomDashPattern
	setCustomDashPattern(newValue ICustomDashPattern)

	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	getMiterLimit() float64
	setMiterLimit(newValue float64)

	getWidth() float64
	setWidth(newValue float64)
}

type LineFormat struct {

	Alignment string `json:"Alignment"`

	CapStyle string `json:"CapStyle"`

	DashStyle string `json:"DashStyle"`

	JoinStyle string `json:"JoinStyle"`

	Style string `json:"Style"`

	BeginArrowHead IArrowHeadProperties `json:"BeginArrowHead,omitempty"`

	EndArrowHead IArrowHeadProperties `json:"EndArrowHead,omitempty"`

	CustomDashPattern ICustomDashPattern `json:"CustomDashPattern,omitempty"`

	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	MiterLimit float64 `json:"MiterLimit"`

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
	if valAlignment, ok := objMap["Alignment"]; ok {
		if valAlignment != nil {
			var valueForAlignment string
			err = json.Unmarshal(*valAlignment, &valueForAlignment)
			if err != nil {
				return err
			}
			this.Alignment = valueForAlignment
		}
	}
	this.CapStyle = "Round"
	if valCapStyle, ok := objMap["CapStyle"]; ok {
		if valCapStyle != nil {
			var valueForCapStyle string
			err = json.Unmarshal(*valCapStyle, &valueForCapStyle)
			if err != nil {
				return err
			}
			this.CapStyle = valueForCapStyle
		}
	}
	this.DashStyle = "Solid"
	if valDashStyle, ok := objMap["DashStyle"]; ok {
		if valDashStyle != nil {
			var valueForDashStyle string
			err = json.Unmarshal(*valDashStyle, &valueForDashStyle)
			if err != nil {
				return err
			}
			this.DashStyle = valueForDashStyle
		}
	}
	this.JoinStyle = "Round"
	if valJoinStyle, ok := objMap["JoinStyle"]; ok {
		if valJoinStyle != nil {
			var valueForJoinStyle string
			err = json.Unmarshal(*valJoinStyle, &valueForJoinStyle)
			if err != nil {
				return err
			}
			this.JoinStyle = valueForJoinStyle
		}
	}
	this.Style = "Single"
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
	
	if valBeginArrowHead, ok := objMap["BeginArrowHead"]; ok {
		if valBeginArrowHead != nil {
			var valueForBeginArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valBeginArrowHead, &valueForBeginArrowHead)
			if err != nil {
				return err
			}
			this.BeginArrowHead = valueForBeginArrowHead
		}
	}
	
	if valEndArrowHead, ok := objMap["EndArrowHead"]; ok {
		if valEndArrowHead != nil {
			var valueForEndArrowHead ArrowHeadProperties
			err = json.Unmarshal(*valEndArrowHead, &valueForEndArrowHead)
			if err != nil {
				return err
			}
			this.EndArrowHead = valueForEndArrowHead
		}
	}
	
	if valCustomDashPattern, ok := objMap["CustomDashPattern"]; ok {
		if valCustomDashPattern != nil {
			var valueForCustomDashPattern CustomDashPattern
			err = json.Unmarshal(*valCustomDashPattern, &valueForCustomDashPattern)
			if err != nil {
				return err
			}
			this.CustomDashPattern = valueForCustomDashPattern
		}
	}
	
	if valFillFormat, ok := objMap["FillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}
	
	if valMiterLimit, ok := objMap["MiterLimit"]; ok {
		if valMiterLimit != nil {
			var valueForMiterLimit float64
			err = json.Unmarshal(*valMiterLimit, &valueForMiterLimit)
			if err != nil {
				return err
			}
			this.MiterLimit = valueForMiterLimit
		}
	}
	
	if valWidth, ok := objMap["Width"]; ok {
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
