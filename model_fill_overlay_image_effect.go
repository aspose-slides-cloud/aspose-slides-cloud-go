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
	GetType() string
	SetType(newValue string)

	// FillBlendMode.
	GetBlend() string
	SetBlend(newValue string)

	// Fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)
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

func (this *FillOverlayImageEffect) GetType() string {
	return this.Type_
}

func (this *FillOverlayImageEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *FillOverlayImageEffect) GetBlend() string {
	return this.Blend
}

func (this *FillOverlayImageEffect) SetBlend(newValue string) {
	this.Blend = newValue
}
func (this *FillOverlayImageEffect) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *FillOverlayImageEffect) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}

func (this *FillOverlayImageEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "FillOverlay"
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	this.Blend = "Darken"
	if valBlend, ok := GetMapValue(objMap, "blend"); ok {
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

	return nil
}
