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

// Represents a Fill Overlay effect. A fill overlay may be used to specify an additional fill for an object and blend the two fills together.
type IFillOverlayImageEffect interface {

	// Image transform effect type
	getType() string
	setType(newValue string)

	// FillBlendMode.
	getBlend() string
	setBlend(newValue string)

	// Fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)
}

type FillOverlayImageEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// FillBlendMode.
	Blend string `json:"Blend"`

	// Fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`
}

func NewFillOverlayImageEffect() *FillOverlayImageEffect {
	instance := new(FillOverlayImageEffect)
	instance.Type_ = "FillOverlay"
	instance.Blend = "Darken"
	return instance
}

func (this *FillOverlayImageEffect) getType() string {
	return this.Type_
}

func (this *FillOverlayImageEffect) setType(newValue string) {
	this.Type_ = newValue
}
func (this *FillOverlayImageEffect) getBlend() string {
	return this.Blend
}

func (this *FillOverlayImageEffect) setBlend(newValue string) {
	this.Blend = newValue
}
func (this *FillOverlayImageEffect) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *FillOverlayImageEffect) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}

func (this *FillOverlayImageEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "FillOverlay"
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
	this.Blend = "Darken"
	if valBlend, ok := objMap["blend"]; ok {
		if valBlend != nil {
			var valueForBlend string
			err = json.Unmarshal(*valBlend, &valueForBlend)
			if err != nil {
				var valueForBlendInt int32
				err = json.Unmarshal(*valBlend, &valueForBlendInt)
				if err != nil {
					return err
				}
				this.Blend = string(valueForBlendInt)
			} else {
				this.Blend = valueForBlend
			}
		}
	}
	if valBlendCap, ok := objMap["Blend"]; ok {
		if valBlendCap != nil {
			var valueForBlend string
			err = json.Unmarshal(*valBlendCap, &valueForBlend)
			if err != nil {
				var valueForBlendInt int32
				err = json.Unmarshal(*valBlendCap, &valueForBlendInt)
				if err != nil {
					return err
				}
				this.Blend = string(valueForBlendInt)
			} else {
				this.Blend = valueForBlend
			}
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
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
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}

	return nil
}
