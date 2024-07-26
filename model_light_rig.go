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

// Light rig
type ILightRig interface {

	// Light direction
	GetDirection() string
	SetDirection(newValue string)

	// Light type
	GetLightType() string
	SetLightType(newValue string)

	// XRotation
	GetXRotation() float64
	SetXRotation(newValue float64)

	// YRotation
	GetYRotation() float64
	SetYRotation(newValue float64)

	// ZRotation
	GetZRotation() float64
	SetZRotation(newValue float64)
}

type LightRig struct {

	// Light direction
	Direction string `json:"Direction,omitempty"`

	// Light type
	LightType string `json:"LightType,omitempty"`

	// XRotation
	XRotation float64 `json:"XRotation,omitempty"`

	// YRotation
	YRotation float64 `json:"YRotation,omitempty"`

	// ZRotation
	ZRotation float64 `json:"ZRotation,omitempty"`
}

func NewLightRig() *LightRig {
	instance := new(LightRig)
	return instance
}

func (this *LightRig) GetDirection() string {
	return this.Direction
}

func (this *LightRig) SetDirection(newValue string) {
	this.Direction = newValue
}
func (this *LightRig) GetLightType() string {
	return this.LightType
}

func (this *LightRig) SetLightType(newValue string) {
	this.LightType = newValue
}
func (this *LightRig) GetXRotation() float64 {
	return this.XRotation
}

func (this *LightRig) SetXRotation(newValue float64) {
	this.XRotation = newValue
}
func (this *LightRig) GetYRotation() float64 {
	return this.YRotation
}

func (this *LightRig) SetYRotation(newValue float64) {
	this.YRotation = newValue
}
func (this *LightRig) GetZRotation() float64 {
	return this.ZRotation
}

func (this *LightRig) SetZRotation(newValue float64) {
	this.ZRotation = newValue
}

func (this *LightRig) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDirection, ok := GetMapValue(objMap, "direction"); ok {
		if valDirection != nil {
			var valueForDirection string
			err = json.Unmarshal(*valDirection, &valueForDirection)
			if err != nil {
				var valueForDirectionInt int32
				err = json.Unmarshal(*valDirection, &valueForDirectionInt)
				if err != nil {
					return err
				}
				this.Direction = string(valueForDirectionInt)
			} else {
				this.Direction = valueForDirection
			}
		}
	}
	
	if valLightType, ok := GetMapValue(objMap, "lightType"); ok {
		if valLightType != nil {
			var valueForLightType string
			err = json.Unmarshal(*valLightType, &valueForLightType)
			if err != nil {
				var valueForLightTypeInt int32
				err = json.Unmarshal(*valLightType, &valueForLightTypeInt)
				if err != nil {
					return err
				}
				this.LightType = string(valueForLightTypeInt)
			} else {
				this.LightType = valueForLightType
			}
		}
	}
	
	if valXRotation, ok := GetMapValue(objMap, "xRotation"); ok {
		if valXRotation != nil {
			var valueForXRotation float64
			err = json.Unmarshal(*valXRotation, &valueForXRotation)
			if err != nil {
				return err
			}
			this.XRotation = valueForXRotation
		}
	}
	
	if valYRotation, ok := GetMapValue(objMap, "yRotation"); ok {
		if valYRotation != nil {
			var valueForYRotation float64
			err = json.Unmarshal(*valYRotation, &valueForYRotation)
			if err != nil {
				return err
			}
			this.YRotation = valueForYRotation
		}
	}
	
	if valZRotation, ok := GetMapValue(objMap, "zRotation"); ok {
		if valZRotation != nil {
			var valueForZRotation float64
			err = json.Unmarshal(*valZRotation, &valueForZRotation)
			if err != nil {
				return err
			}
			this.ZRotation = valueForZRotation
		}
	}

	return nil
}
