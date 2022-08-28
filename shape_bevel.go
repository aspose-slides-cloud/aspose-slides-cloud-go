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

// ShapeBevel
type IShapeBevel interface {

	// Bevel type
	GetBevelType() string
	SetBevelType(newValue string)

	// Bevel width
	GetWidth() float64
	SetWidth(newValue float64)

	// Bevel height
	GetHeight() float64
	SetHeight(newValue float64)
}

type ShapeBevel struct {

	// Bevel type
	BevelType string `json:"BevelType,omitempty"`

	// Bevel width
	Width float64 `json:"Width,omitempty"`

	// Bevel height
	Height float64 `json:"Height,omitempty"`
}

func NewShapeBevel() *ShapeBevel {
	instance := new(ShapeBevel)
	instance.BevelType = ""
	return instance
}

func (this *ShapeBevel) GetBevelType() string {
	return this.BevelType
}

func (this *ShapeBevel) SetBevelType(newValue string) {
	this.BevelType = newValue
}
func (this *ShapeBevel) GetWidth() float64 {
	return this.Width
}

func (this *ShapeBevel) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *ShapeBevel) GetHeight() float64 {
	return this.Height
}

func (this *ShapeBevel) SetHeight(newValue float64) {
	this.Height = newValue
}

func (this *ShapeBevel) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.BevelType = ""
	if valBevelType, ok := objMap["bevelType"]; ok {
		if valBevelType != nil {
			var valueForBevelType string
			err = json.Unmarshal(*valBevelType, &valueForBevelType)
			if err != nil {
				var valueForBevelTypeInt int32
				err = json.Unmarshal(*valBevelType, &valueForBevelTypeInt)
				if err != nil {
					return err
				}
				this.BevelType = string(valueForBevelTypeInt)
			} else {
				this.BevelType = valueForBevelType
			}
		}
	}
	if valBevelTypeCap, ok := objMap["BevelType"]; ok {
		if valBevelTypeCap != nil {
			var valueForBevelType string
			err = json.Unmarshal(*valBevelTypeCap, &valueForBevelType)
			if err != nil {
				var valueForBevelTypeInt int32
				err = json.Unmarshal(*valBevelTypeCap, &valueForBevelTypeInt)
				if err != nil {
					return err
				}
				this.BevelType = string(valueForBevelTypeInt)
			} else {
				this.BevelType = valueForBevelType
			}
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
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	return nil
}
