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

// Camera
type ICamera interface {

	// Camera type
	getCameraType() string
	setCameraType(newValue string)

	// Camera FOV
	getFieldOfViewAngle() float64
	setFieldOfViewAngle(newValue float64)

	// Camera zoom
	getZoom() float64
	setZoom(newValue float64)

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

type Camera struct {

	// Camera type
	CameraType string `json:"CameraType,omitempty"`

	// Camera FOV
	FieldOfViewAngle float64 `json:"FieldOfViewAngle,omitempty"`

	// Camera zoom
	Zoom float64 `json:"Zoom,omitempty"`

	// XRotation
	XRotation float64 `json:"XRotation,omitempty"`

	// YRotation
	YRotation float64 `json:"YRotation,omitempty"`

	// ZRotation
	ZRotation float64 `json:"ZRotation,omitempty"`
}

func NewCamera() *Camera {
	instance := new(Camera)
	instance.CameraType = ""
	return instance
}

func (this *Camera) getCameraType() string {
	return this.CameraType
}

func (this *Camera) setCameraType(newValue string) {
	this.CameraType = newValue
}
func (this *Camera) getFieldOfViewAngle() float64 {
	return this.FieldOfViewAngle
}

func (this *Camera) setFieldOfViewAngle(newValue float64) {
	this.FieldOfViewAngle = newValue
}
func (this *Camera) getZoom() float64 {
	return this.Zoom
}

func (this *Camera) setZoom(newValue float64) {
	this.Zoom = newValue
}
func (this *Camera) getXRotation() float64 {
	return this.XRotation
}

func (this *Camera) setXRotation(newValue float64) {
	this.XRotation = newValue
}
func (this *Camera) getYRotation() float64 {
	return this.YRotation
}

func (this *Camera) setYRotation(newValue float64) {
	this.YRotation = newValue
}
func (this *Camera) getZRotation() float64 {
	return this.ZRotation
}

func (this *Camera) setZRotation(newValue float64) {
	this.ZRotation = newValue
}

func (this *Camera) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.CameraType = ""
	if valCameraType, ok := objMap["cameraType"]; ok {
		if valCameraType != nil {
			var valueForCameraType string
			err = json.Unmarshal(*valCameraType, &valueForCameraType)
			if err != nil {
				var valueForCameraTypeInt int32
				err = json.Unmarshal(*valCameraType, &valueForCameraTypeInt)
				if err != nil {
					return err
				}
				this.CameraType = string(valueForCameraTypeInt)
			} else {
				this.CameraType = valueForCameraType
			}
		}
	}
	if valCameraTypeCap, ok := objMap["CameraType"]; ok {
		if valCameraTypeCap != nil {
			var valueForCameraType string
			err = json.Unmarshal(*valCameraTypeCap, &valueForCameraType)
			if err != nil {
				var valueForCameraTypeInt int32
				err = json.Unmarshal(*valCameraTypeCap, &valueForCameraTypeInt)
				if err != nil {
					return err
				}
				this.CameraType = string(valueForCameraTypeInt)
			} else {
				this.CameraType = valueForCameraType
			}
		}
	}
	
	if valFieldOfViewAngle, ok := objMap["fieldOfViewAngle"]; ok {
		if valFieldOfViewAngle != nil {
			var valueForFieldOfViewAngle float64
			err = json.Unmarshal(*valFieldOfViewAngle, &valueForFieldOfViewAngle)
			if err != nil {
				return err
			}
			this.FieldOfViewAngle = valueForFieldOfViewAngle
		}
	}
	if valFieldOfViewAngleCap, ok := objMap["FieldOfViewAngle"]; ok {
		if valFieldOfViewAngleCap != nil {
			var valueForFieldOfViewAngle float64
			err = json.Unmarshal(*valFieldOfViewAngleCap, &valueForFieldOfViewAngle)
			if err != nil {
				return err
			}
			this.FieldOfViewAngle = valueForFieldOfViewAngle
		}
	}
	
	if valZoom, ok := objMap["zoom"]; ok {
		if valZoom != nil {
			var valueForZoom float64
			err = json.Unmarshal(*valZoom, &valueForZoom)
			if err != nil {
				return err
			}
			this.Zoom = valueForZoom
		}
	}
	if valZoomCap, ok := objMap["Zoom"]; ok {
		if valZoomCap != nil {
			var valueForZoom float64
			err = json.Unmarshal(*valZoomCap, &valueForZoom)
			if err != nil {
				return err
			}
			this.Zoom = valueForZoom
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
