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
	getDirection() string
	setDirection(newValue string)

	// Light type
	getLightType() string
	setLightType(newValue string)

	// XRotation
	getXRotation() float64
	setXRotation(newValue float64)

	// YRotation
	getYRotation() float64
	setYRotation(newValue float64)

	// ZRotation
	getZRotation() float64
	setZRotation(newValue float64)
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
	instance.Direction = ""
	instance.LightType = ""
	return instance
}

func (this *LightRig) getDirection() string {
	return this.Direction
}

func (this *LightRig) setDirection(newValue string) {
	this.Direction = newValue
}
func (this *LightRig) getLightType() string {
	return this.LightType
}

func (this *LightRig) setLightType(newValue string) {
	this.LightType = newValue
}
func (this *LightRig) getXRotation() float64 {
	return this.XRotation
}

func (this *LightRig) setXRotation(newValue float64) {
	this.XRotation = newValue
}
func (this *LightRig) getYRotation() float64 {
	return this.YRotation
}

func (this *LightRig) setYRotation(newValue float64) {
	this.YRotation = newValue
}
func (this *LightRig) getZRotation() float64 {
	return this.ZRotation
}

func (this *LightRig) setZRotation(newValue float64) {
	this.ZRotation = newValue
}

func (this *LightRig) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Direction = ""
	if valDirection, ok := objMap["direction"]; ok {
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
	if valDirectionCap, ok := objMap["Direction"]; ok {
		if valDirectionCap != nil {
			var valueForDirection string
			err = json.Unmarshal(*valDirectionCap, &valueForDirection)
			if err != nil {
				var valueForDirectionInt int32
				err = json.Unmarshal(*valDirectionCap, &valueForDirectionInt)
				if err != nil {
					return err
				}
				this.Direction = string(valueForDirectionInt)
			} else {
				this.Direction = valueForDirection
			}
		}
	}
	this.LightType = ""
	if valLightType, ok := objMap["lightType"]; ok {
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
	if valLightTypeCap, ok := objMap["LightType"]; ok {
		if valLightTypeCap != nil {
			var valueForLightType string
			err = json.Unmarshal(*valLightTypeCap, &valueForLightType)
			if err != nil {
				var valueForLightTypeInt int32
				err = json.Unmarshal(*valLightTypeCap, &valueForLightTypeInt)
				if err != nil {
					return err
				}
				this.LightType = string(valueForLightTypeInt)
			} else {
				this.LightType = valueForLightType
			}
		}
	}
	
	if valXRotation, ok := objMap["xRotation"]; ok {
		if valXRotation != nil {
			var valueForXRotation float64
			err = json.Unmarshal(*valXRotation, &valueForXRotation)
			if err != nil {
				return err
			}
			this.XRotation = valueForXRotation
		}
	}
	if valXRotationCap, ok := objMap["XRotation"]; ok {
		if valXRotationCap != nil {
			var valueForXRotation float64
			err = json.Unmarshal(*valXRotationCap, &valueForXRotation)
			if err != nil {
				return err
			}
			this.XRotation = valueForXRotation
		}
	}
	
	if valYRotation, ok := objMap["yRotation"]; ok {
		if valYRotation != nil {
			var valueForYRotation float64
			err = json.Unmarshal(*valYRotation, &valueForYRotation)
			if err != nil {
				return err
			}
			this.YRotation = valueForYRotation
		}
	}
	if valYRotationCap, ok := objMap["YRotation"]; ok {
		if valYRotationCap != nil {
			var valueForYRotation float64
			err = json.Unmarshal(*valYRotationCap, &valueForYRotation)
			if err != nil {
				return err
			}
			this.YRotation = valueForYRotation
		}
	}
	
	if valZRotation, ok := objMap["zRotation"]; ok {
		if valZRotation != nil {
			var valueForZRotation float64
			err = json.Unmarshal(*valZRotation, &valueForZRotation)
			if err != nil {
				return err
			}
			this.ZRotation = valueForZRotation
		}
	}
	if valZRotationCap, ok := objMap["ZRotation"]; ok {
		if valZRotationCap != nil {
			var valueForZRotation float64
			err = json.Unmarshal(*valZRotationCap, &valueForZRotation)
			if err != nil {
				return err
			}
			this.ZRotation = valueForZRotation
		}
	}

	return nil
}
